// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

/*
MarkdownRender performs requests for "markdown/render"

Render an arbitrary Markdown document.

  POST /markdown

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document
*/
func (c *Client) MarkdownRender(ctx context.Context, req *MarkdownRenderReq, opt ...RequestOption) (*MarkdownRenderResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &MarkdownRenderResponse{
		request:  req,
		response: *r,
	}
	err = r.decodeBody(nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
MarkdownRenderReq is request data for Client.MarkdownRender

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document
*/
type MarkdownRenderReq struct {
	pgURL       string
	RequestBody MarkdownRenderReqBody
}

func (r *MarkdownRenderReq) pagingURL() string {
	return r.pgURL
}

func (r *MarkdownRenderReq) urlPath() string {
	return fmt.Sprintf("/markdown")
}

func (r *MarkdownRenderReq) method() string {
	return "POST"
}

func (r *MarkdownRenderReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *MarkdownRenderReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *MarkdownRenderReq) body() interface{} {
	return r.RequestBody
}

func (r *MarkdownRenderReq) dataStatuses() []int {
	return []int{}
}

func (r *MarkdownRenderReq) validStatuses() []int {
	return []int{200}
}

func (r *MarkdownRenderReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{attrRegular}
}

// httpRequest creates an http request
func (r *MarkdownRenderReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *MarkdownRenderReq) Rel(link RelName, resp *MarkdownRenderResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
MarkdownRenderReqBody is a request body for markdown/render

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document
*/
type MarkdownRenderReqBody struct {

	/*
	   The repository context to use when creating references in `gfm` mode. Omit this
	   parameter when using `markdown` mode.
	*/
	Context *string `json:"context,omitempty"`

	/*
	   The rendering mode. Can be either:
	   \* `markdown` to render a document in plain Markdown, just like README.md files
	   are rendered.
	   \* `gfm` to render a document in [GitHub Flavored
	   Markdown](https://github.github.com/gfm/), which creates links for user mentions
	   as well as references to SHA-1 hashes, issues, and pull requests.
	*/
	Mode *string `json:"mode,omitempty"`

	// The Markdown text to render in HTML. Markdown content must be 400 KB or less.
	Text *string `json:"text"`
}

/*
MarkdownRenderResponse is a response for MarkdownRender

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document
*/
type MarkdownRenderResponse struct {
	response
	request *MarkdownRenderReq
}

/*
MarkdownRenderRaw performs requests for "markdown/render-raw"

Render a Markdown document in raw mode.

  POST /markdown/raw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
func (c *Client) MarkdownRenderRaw(ctx context.Context, req *MarkdownRenderRawReq, opt ...RequestOption) (*MarkdownRenderRawResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &MarkdownRenderRawResponse{
		request:  req,
		response: *r,
	}
	err = r.decodeBody(nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
MarkdownRenderRawReq is request data for Client.MarkdownRenderRaw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
type MarkdownRenderRawReq struct {
	pgURL       string
	RequestBody MarkdownRenderRawReqBody

	// Setting content-type header is required for this endpoint
	ContentTypeHeader *string
}

func (r *MarkdownRenderRawReq) pagingURL() string {
	return r.pgURL
}

func (r *MarkdownRenderRawReq) urlPath() string {
	return fmt.Sprintf("/markdown/raw")
}

func (r *MarkdownRenderRawReq) method() string {
	return "POST"
}

func (r *MarkdownRenderRawReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *MarkdownRenderRawReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"content-type": r.ContentTypeHeader}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *MarkdownRenderRawReq) body() interface{} {
	return r.RequestBody
}

func (r *MarkdownRenderRawReq) dataStatuses() []int {
	return []int{}
}

func (r *MarkdownRenderRawReq) validStatuses() []int {
	return []int{200}
}

func (r *MarkdownRenderRawReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{attrRegular}
}

// httpRequest creates an http request
func (r *MarkdownRenderRawReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *MarkdownRenderRawReq) Rel(link RelName, resp *MarkdownRenderRawResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
MarkdownRenderRawReqBody is a request body for markdown/render-raw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
type MarkdownRenderRawReqBody string

/*
MarkdownRenderRawResponse is a response for MarkdownRenderRaw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
type MarkdownRenderRawResponse struct {
	response
	request *MarkdownRenderRawReq
}
