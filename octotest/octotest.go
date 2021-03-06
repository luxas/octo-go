package octotest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"

	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/octotest/internal"
)

const (
	// got a request that isn't matched by any matcher
	StatusUnexpectedRequest = internal.StatusUnexpectedRequest
	// couldn't clone the request body
	StatusErrorCloningRequestBody = internal.StatusErrorCloningRequestBody
)

// Server is a test server that will serve requests configured with Expect()
type Server struct {
	opts    []octo.RequestOption
	server  *httptest.Server
	mu      sync.Mutex
	finish  func()
	client  octo.Client
	handler *internal.RequestHandler
	
	httpClient *http.Client
	baseURL *net.URL
}

// Finish stops the underlying server
func (s *Server) Finish() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.finish()
}

// Client returns an octo.Client configured for this server
func (s *Server) Client() octo.Client {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.client
}

// HTTPClient returns an *http.Client configured for this server
func (s *Server) HTTPClient() *http.Client {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.httpClient
}

// BaseURL returns the base URL this server is configured for
func (s *Server) BaseURL() *net.URL {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.baseURL
}

func (s *Server) handle(w http.ResponseWriter, req *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handler.ServeHTTP(w, req)
}

// New returns a new *Server
func New(opt ...octo.RequestOption) *Server {
	s := &Server{
		opts:    opt,
		handler: new(internal.RequestHandler),
	}
	testServer := httptest.NewServer(http.HandlerFunc(s.handle))
	s.server = testServer
	baseURL, err := url.Parse(s.server.URL)
	if err != nil {
		panic(fmt.Sprintf("error parsing server url: %v", err))
	}
	s.baseURL = baseURL
	s.httpClient = s.server.Client()
	
	s.client = s.opts
	s.client = append(s.client, octo.WithHTTPClient(s.httpClient), octo.WithBaseURL(*s.baseURL))
	s.finish = func() {
		s.server.Close()
	}
	return s
}

// Expect configures the Server to expect a request matching request and respond with response
func (s *Server) Expect(request HTTPRequester, response http.Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	wantReq, err := request.HTTPRequest(context.Background(), s.opts...)
	if err != nil {
		panic(err)
	}
	wantBody := ""
	if wantReq.Body != nil {
		wb, err := ioutil.ReadAll(wantReq.Body)
		if err != nil {
			panic(err)
		}
		err = wantReq.Body.Close()
		if err != nil {
			panic(err)
		}
		wantBody = string(wb)
	}
	matcher := internal.MatchAll(
		internal.MatchRequestPath(wantReq.URL.Path),
		internal.MatchRequestQuery(wantReq.URL.RawQuery),
		internal.MatchBody(wantBody),
	)
	s.handler.Expect(&internal.ExpectedRequest{
		Matcher: matcher,
		Handler: response,
	})
}

// HTTPRequester is an interface that is met by all of octo's requests
type HTTPRequester interface {
	HTTPRequest(ctx context.Context, opt ...octo.RequestOption) (*http.Request, error)
}

// HTTPResponder is a responder for an expected request
type HTTPResponder struct {
	StatusCode int
	Body       []byte
	Header     http.Header
}

func (r *HTTPResponder) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	for k, vals := range r.Header {
		for _, val := range vals {
			w.Header().Add(k, val)
		}
	}
	if w.Header().Get("content-type") == "" {
		w.Header().Set("content-type", "application/json")
	}
	w.WriteHeader(r.StatusCode)
	_, err := w.Write(r.Body)
	if err != nil {
		panic(err)
	}
}

// JSONResponder is a responder that will respond with the given status code and body.
func JSONResponder(statusCode int, body interface{}) *HTTPResponder {
	bd, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return &HTTPResponder{
		StatusCode: statusCode,
		Body:       bd,
	}
}

// RelLinkHandler adds a rel link to a response. This is useful for testing paging through results.
func RelLinkHandler(relName octo.RelName, handler http.Handler, relLinkRequester HTTPRequester, server *Server) http.HandlerFunc {
	relReq, err := relLinkRequester.HTTPRequest(context.Background(), server.Client()...)
	if err != nil {
		panic(fmt.Sprintf("error from relLinkRequester.HTTPRequest(ctx): %v", err))
	}
	linkURL := relReq.URL.String()
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Link", fmt.Sprintf(`<%s>; rel="%s"`, linkURL, string(relName)))
		handler.ServeHTTP(w, req)
	}
}
