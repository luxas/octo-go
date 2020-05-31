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
LicensesGet performs requests for "licenses/get"

Get an individual license.

  GET /licenses/{license}

https://developer.github.com/v3/licenses/#get-an-individual-license
*/
func (c *Client) LicensesGet(ctx context.Context, req *LicensesGetReq, opt ...RequestOption) (*LicensesGetResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &LicensesGetResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(LicensesGetResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesGetReq is request data for Client.LicensesGet

https://developer.github.com/v3/licenses/#get-an-individual-license
*/
type LicensesGetReq struct {
	pgURL   string
	License string
}

func (r *LicensesGetReq) pagingURL() string {
	return r.pgURL
}

func (r *LicensesGetReq) urlPath() string {
	return fmt.Sprintf("/licenses/%v", r.License)
}

func (r *LicensesGetReq) method() string {
	return "GET"
}

func (r *LicensesGetReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *LicensesGetReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *LicensesGetReq) body() interface{} {
	return nil
}

func (r *LicensesGetReq) dataStatuses() []int {
	return []int{200}
}

func (r *LicensesGetReq) validStatuses() []int {
	return []int{200}
}

func (r *LicensesGetReq) endpointAttribute() endpointAttribute {
	return attrRegular
}

// httpRequest creates an http request
func (r *LicensesGetReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *LicensesGetReq) Rel(link RelName, resp *LicensesGetResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
LicensesGetResponseBody is a response body for LicensesGet

https://developer.github.com/v3/licenses/#get-an-individual-license
*/
type LicensesGetResponseBody struct {
	components.License
}

/*
LicensesGetResponse is a response for LicensesGet

https://developer.github.com/v3/licenses/#get-an-individual-license
*/
type LicensesGetResponse struct {
	response
	request *LicensesGetReq
	Data    *LicensesGetResponseBody
}

/*
LicensesGetForRepo performs requests for "licenses/get-for-repo"

Get the contents of a repository's license.

  GET /repos/{owner}/{repo}/license

https://developer.github.com/v3/licenses/#get-the-contents-of-a-repositorys-license
*/
func (c *Client) LicensesGetForRepo(ctx context.Context, req *LicensesGetForRepoReq, opt ...RequestOption) (*LicensesGetForRepoResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &LicensesGetForRepoResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(LicensesGetForRepoResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesGetForRepoReq is request data for Client.LicensesGetForRepo

https://developer.github.com/v3/licenses/#get-the-contents-of-a-repositorys-license
*/
type LicensesGetForRepoReq struct {
	pgURL string
	Owner string
	Repo  string
}

func (r *LicensesGetForRepoReq) pagingURL() string {
	return r.pgURL
}

func (r *LicensesGetForRepoReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/license", r.Owner, r.Repo)
}

func (r *LicensesGetForRepoReq) method() string {
	return "GET"
}

func (r *LicensesGetForRepoReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *LicensesGetForRepoReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *LicensesGetForRepoReq) body() interface{} {
	return nil
}

func (r *LicensesGetForRepoReq) dataStatuses() []int {
	return []int{200}
}

func (r *LicensesGetForRepoReq) validStatuses() []int {
	return []int{200}
}

func (r *LicensesGetForRepoReq) endpointAttribute() endpointAttribute {
	return attrRegular
}

// httpRequest creates an http request
func (r *LicensesGetForRepoReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *LicensesGetForRepoReq) Rel(link RelName, resp *LicensesGetForRepoResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
LicensesGetForRepoResponseBody is a response body for LicensesGetForRepo

https://developer.github.com/v3/licenses/#get-the-contents-of-a-repositorys-license
*/
type LicensesGetForRepoResponseBody struct {
	components.LicenseContent
}

/*
LicensesGetForRepoResponse is a response for LicensesGetForRepo

https://developer.github.com/v3/licenses/#get-the-contents-of-a-repositorys-license
*/
type LicensesGetForRepoResponse struct {
	response
	request *LicensesGetForRepoReq
	Data    *LicensesGetForRepoResponseBody
}

/*
LicensesListCommonlyUsed performs requests for "licenses/list-commonly-used"

List commonly used licenses.

  GET /licenses

https://developer.github.com/v3/licenses/#list-commonly-used-licenses
*/
func (c *Client) LicensesListCommonlyUsed(ctx context.Context, req *LicensesListCommonlyUsedReq, opt ...RequestOption) (*LicensesListCommonlyUsedResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &LicensesListCommonlyUsedResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(LicensesListCommonlyUsedResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesListCommonlyUsedReq is request data for Client.LicensesListCommonlyUsed

https://developer.github.com/v3/licenses/#list-commonly-used-licenses
*/
type LicensesListCommonlyUsedReq struct {
	pgURL string
}

func (r *LicensesListCommonlyUsedReq) pagingURL() string {
	return r.pgURL
}

func (r *LicensesListCommonlyUsedReq) urlPath() string {
	return fmt.Sprintf("/licenses")
}

func (r *LicensesListCommonlyUsedReq) method() string {
	return "GET"
}

func (r *LicensesListCommonlyUsedReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *LicensesListCommonlyUsedReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *LicensesListCommonlyUsedReq) body() interface{} {
	return nil
}

func (r *LicensesListCommonlyUsedReq) dataStatuses() []int {
	return []int{200}
}

func (r *LicensesListCommonlyUsedReq) validStatuses() []int {
	return []int{200}
}

func (r *LicensesListCommonlyUsedReq) endpointAttribute() endpointAttribute {
	return attrRegular
}

// httpRequest creates an http request
func (r *LicensesListCommonlyUsedReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *LicensesListCommonlyUsedReq) Rel(link RelName, resp *LicensesListCommonlyUsedResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
LicensesListCommonlyUsedResponseBody is a response body for LicensesListCommonlyUsed

https://developer.github.com/v3/licenses/#list-commonly-used-licenses
*/
type LicensesListCommonlyUsedResponseBody []struct {
	components.LicenseSimple
}

/*
LicensesListCommonlyUsedResponse is a response for LicensesListCommonlyUsed

https://developer.github.com/v3/licenses/#list-commonly-used-licenses
*/
type LicensesListCommonlyUsedResponse struct {
	response
	request *LicensesListCommonlyUsedReq
	Data    *LicensesListCommonlyUsedResponseBody
}
