// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	"net/http"
	"net/url"
)

/*
RateLimitGet performs requests for "rate-limit/get"

Get rate limit status for the authenticated user.

  GET /rate_limit

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user
*/
func RateLimitGet(ctx context.Context, req *RateLimitGetReq, opt ...RequestOption) (*RateLimitGetResponse, error) {
	if req == nil {
		req = new(RateLimitGetReq)
	}
	resp := &RateLimitGetResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = new(RateLimitGetResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
RateLimitGet performs requests for "rate-limit/get"

Get rate limit status for the authenticated user.

  GET /rate_limit

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user
*/
func (c Client) RateLimitGet(ctx context.Context, req *RateLimitGetReq, opt ...RequestOption) (*RateLimitGetResponse, error) {
	return RateLimitGet(ctx, req, append(c, opt...)...)
}

/*
RateLimitGetReq is request data for Client.RateLimitGet

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user
*/
type RateLimitGetReq struct {
	_url string
}

func (r *RateLimitGetReq) url() string {
	return r._url
}

func (r *RateLimitGetReq) urlPath() string {
	return fmt.Sprintf("/rate_limit")
}

func (r *RateLimitGetReq) method() string {
	return "GET"
}

func (r *RateLimitGetReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *RateLimitGetReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *RateLimitGetReq) body() interface{} {
	return nil
}

func (r *RateLimitGetReq) dataStatuses() []int {
	return []int{200}
}

func (r *RateLimitGetReq) validStatuses() []int {
	return []int{200}
}

func (r *RateLimitGetReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *RateLimitGetReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RateLimitGetReq) Rel(link RelName, resp *RateLimitGetResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
RateLimitGetResponseBody is a response body for RateLimitGet

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user
*/
type RateLimitGetResponseBody components.RateLimitOverview

/*
RateLimitGetResponse is a response for RateLimitGet

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user
*/
type RateLimitGetResponse struct {
	response
	request *RateLimitGetReq
	Data    *RateLimitGetResponseBody
}
