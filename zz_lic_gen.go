// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	"net/http"
	"net/url"
	"strconv"
)

/*
LicensesGet performs requests for "licenses/get"

Get a license.

  GET /licenses/{license}

https://developer.github.com/v3/licenses/#get-a-license
*/
func LicensesGet(ctx context.Context, req *LicensesGetReq, opt ...RequestOption) (*LicensesGetResponse, error) {
	if req == nil {
		req = new(LicensesGetReq)
	}
	resp := &LicensesGetResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.License{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesGet performs requests for "licenses/get"

Get a license.

  GET /licenses/{license}

https://developer.github.com/v3/licenses/#get-a-license
*/
func (c Client) LicensesGet(ctx context.Context, req *LicensesGetReq, opt ...RequestOption) (*LicensesGetResponse, error) {
	return LicensesGet(ctx, req, append(c, opt...)...)
}

/*
LicensesGetReq is request data for Client.LicensesGet

https://developer.github.com/v3/licenses/#get-a-license
*/
type LicensesGetReq struct {
	_url string

	// license parameter
	License string
}

func (r *LicensesGetReq) url() string {
	return r._url
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
	headerVals := map[string]*string{"accept": String("application/json")}
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
	return []int{200, 304}
}

func (r *LicensesGetReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *LicensesGetReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
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
	r._url = u
	return true
}

/*
LicensesGetResponse is a response for LicensesGet

https://developer.github.com/v3/licenses/#get-a-license
*/
type LicensesGetResponse struct {
	response
	request *LicensesGetReq
	Data    components.License
}

/*
LicensesGetAllCommonlyUsed performs requests for "licenses/get-all-commonly-used"

Get all commonly used licenses.

  GET /licenses

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
func LicensesGetAllCommonlyUsed(ctx context.Context, req *LicensesGetAllCommonlyUsedReq, opt ...RequestOption) (*LicensesGetAllCommonlyUsedResponse, error) {
	if req == nil {
		req = new(LicensesGetAllCommonlyUsedReq)
	}
	resp := &LicensesGetAllCommonlyUsedResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = LicensesGetAllCommonlyUsedResponseBody{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesGetAllCommonlyUsed performs requests for "licenses/get-all-commonly-used"

Get all commonly used licenses.

  GET /licenses

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
func (c Client) LicensesGetAllCommonlyUsed(ctx context.Context, req *LicensesGetAllCommonlyUsedReq, opt ...RequestOption) (*LicensesGetAllCommonlyUsedResponse, error) {
	return LicensesGetAllCommonlyUsed(ctx, req, append(c, opt...)...)
}

/*
LicensesGetAllCommonlyUsedReq is request data for Client.LicensesGetAllCommonlyUsed

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
type LicensesGetAllCommonlyUsedReq struct {
	_url     string
	Featured *bool

	// Results per page (max 100)
	PerPage *int64
}

func (r *LicensesGetAllCommonlyUsedReq) url() string {
	return r._url
}

func (r *LicensesGetAllCommonlyUsedReq) urlPath() string {
	return fmt.Sprintf("/licenses")
}

func (r *LicensesGetAllCommonlyUsedReq) method() string {
	return "GET"
}

func (r *LicensesGetAllCommonlyUsedReq) urlQuery() url.Values {
	query := url.Values{}
	if r.Featured != nil {
		query.Set("featured", strconv.FormatBool(*r.Featured))
	}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	return query
}

func (r *LicensesGetAllCommonlyUsedReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *LicensesGetAllCommonlyUsedReq) body() interface{} {
	return nil
}

func (r *LicensesGetAllCommonlyUsedReq) dataStatuses() []int {
	return []int{200}
}

func (r *LicensesGetAllCommonlyUsedReq) validStatuses() []int {
	return []int{200, 304}
}

func (r *LicensesGetAllCommonlyUsedReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *LicensesGetAllCommonlyUsedReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *LicensesGetAllCommonlyUsedReq) Rel(link RelName, resp *LicensesGetAllCommonlyUsedResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
LicensesGetAllCommonlyUsedResponseBody is a response body for LicensesGetAllCommonlyUsed

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
type LicensesGetAllCommonlyUsedResponseBody []components.LicenseSimple

/*
LicensesGetAllCommonlyUsedResponse is a response for LicensesGetAllCommonlyUsed

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
type LicensesGetAllCommonlyUsedResponse struct {
	response
	request *LicensesGetAllCommonlyUsedReq
	Data    LicensesGetAllCommonlyUsedResponseBody
}

/*
LicensesGetForRepo performs requests for "licenses/get-for-repo"

Get the license for a repository.

  GET /repos/{owner}/{repo}/license

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository
*/
func LicensesGetForRepo(ctx context.Context, req *LicensesGetForRepoReq, opt ...RequestOption) (*LicensesGetForRepoResponse, error) {
	if req == nil {
		req = new(LicensesGetForRepoReq)
	}
	resp := &LicensesGetForRepoResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.LicenseContent{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesGetForRepo performs requests for "licenses/get-for-repo"

Get the license for a repository.

  GET /repos/{owner}/{repo}/license

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository
*/
func (c Client) LicensesGetForRepo(ctx context.Context, req *LicensesGetForRepoReq, opt ...RequestOption) (*LicensesGetForRepoResponse, error) {
	return LicensesGetForRepo(ctx, req, append(c, opt...)...)
}

/*
LicensesGetForRepoReq is request data for Client.LicensesGetForRepo

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository
*/
type LicensesGetForRepoReq struct {
	_url  string
	Owner string
	Repo  string
}

func (r *LicensesGetForRepoReq) url() string {
	return r._url
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
	headerVals := map[string]*string{"accept": String("application/json")}
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

func (r *LicensesGetForRepoReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *LicensesGetForRepoReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
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
	r._url = u
	return true
}

/*
LicensesGetForRepoResponse is a response for LicensesGetForRepo

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository
*/
type LicensesGetForRepoResponse struct {
	response
	request *LicensesGetForRepoReq
	Data    components.LicenseContent
}
