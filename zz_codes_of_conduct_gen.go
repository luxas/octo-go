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
CodesOfConductGetAllCodesOfConduct performs requests for "codes-of-conduct/get-all-codes-of-conduct"

List all codes of conduct.

  GET /codes_of_conduct

https://developer.github.com/v3/codes_of_conduct/#list-all-codes-of-conduct
*/
func (c *Client) CodesOfConductGetAllCodesOfConduct(ctx context.Context, req *CodesOfConductGetAllCodesOfConductReq, opt ...RequestOption) (*CodesOfConductGetAllCodesOfConductResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &CodesOfConductGetAllCodesOfConductResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(CodesOfConductGetAllCodesOfConductResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodesOfConductGetAllCodesOfConductReq is request data for Client.CodesOfConductGetAllCodesOfConduct

https://developer.github.com/v3/codes_of_conduct/#list-all-codes-of-conduct
*/
type CodesOfConductGetAllCodesOfConductReq struct {
	pgURL string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

func (r *CodesOfConductGetAllCodesOfConductReq) pagingURL() string {
	return r.pgURL
}

func (r *CodesOfConductGetAllCodesOfConductReq) urlPath() string {
	return fmt.Sprintf("/codes_of_conduct")
}

func (r *CodesOfConductGetAllCodesOfConductReq) method() string {
	return "GET"
}

func (r *CodesOfConductGetAllCodesOfConductReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *CodesOfConductGetAllCodesOfConductReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"scarlet-witch": r.ScarletWitchPreview}
	if requiredPreviews {
		previewVals["scarlet-witch"] = true
	}
	if allPreviews {
		previewVals["scarlet-witch"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *CodesOfConductGetAllCodesOfConductReq) body() interface{} {
	return nil
}

func (r *CodesOfConductGetAllCodesOfConductReq) dataStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetAllCodesOfConductReq) validStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetAllCodesOfConductReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *CodesOfConductGetAllCodesOfConductReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodesOfConductGetAllCodesOfConductReq) Rel(link RelName, resp *CodesOfConductGetAllCodesOfConductResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
CodesOfConductGetAllCodesOfConductResponseBody is a response body for CodesOfConductGetAllCodesOfConduct

https://developer.github.com/v3/codes_of_conduct/#list-all-codes-of-conduct
*/
type CodesOfConductGetAllCodesOfConductResponseBody []struct {
	components.CodeOfConductSimple
}

/*
CodesOfConductGetAllCodesOfConductResponse is a response for CodesOfConductGetAllCodesOfConduct

https://developer.github.com/v3/codes_of_conduct/#list-all-codes-of-conduct
*/
type CodesOfConductGetAllCodesOfConductResponse struct {
	response
	request *CodesOfConductGetAllCodesOfConductReq
	Data    *CodesOfConductGetAllCodesOfConductResponseBody
}

/*
CodesOfConductGetConductCode performs requests for "codes-of-conduct/get-conduct-code"

Get an individual code of conduct.

  GET /codes_of_conduct/{key}

https://developer.github.com/v3/codes_of_conduct/#get-an-individual-code-of-conduct
*/
func (c *Client) CodesOfConductGetConductCode(ctx context.Context, req *CodesOfConductGetConductCodeReq, opt ...RequestOption) (*CodesOfConductGetConductCodeResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &CodesOfConductGetConductCodeResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(CodesOfConductGetConductCodeResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodesOfConductGetConductCodeReq is request data for Client.CodesOfConductGetConductCode

https://developer.github.com/v3/codes_of_conduct/#get-an-individual-code-of-conduct
*/
type CodesOfConductGetConductCodeReq struct {
	pgURL string
	Key   string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

func (r *CodesOfConductGetConductCodeReq) pagingURL() string {
	return r.pgURL
}

func (r *CodesOfConductGetConductCodeReq) urlPath() string {
	return fmt.Sprintf("/codes_of_conduct/%v", r.Key)
}

func (r *CodesOfConductGetConductCodeReq) method() string {
	return "GET"
}

func (r *CodesOfConductGetConductCodeReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *CodesOfConductGetConductCodeReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"scarlet-witch": r.ScarletWitchPreview}
	if requiredPreviews {
		previewVals["scarlet-witch"] = true
	}
	if allPreviews {
		previewVals["scarlet-witch"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *CodesOfConductGetConductCodeReq) body() interface{} {
	return nil
}

func (r *CodesOfConductGetConductCodeReq) dataStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetConductCodeReq) validStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetConductCodeReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *CodesOfConductGetConductCodeReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodesOfConductGetConductCodeReq) Rel(link RelName, resp *CodesOfConductGetConductCodeResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
CodesOfConductGetConductCodeResponseBody is a response body for CodesOfConductGetConductCode

https://developer.github.com/v3/codes_of_conduct/#get-an-individual-code-of-conduct
*/
type CodesOfConductGetConductCodeResponseBody struct {
	components.CodeOfConduct
}

/*
CodesOfConductGetConductCodeResponse is a response for CodesOfConductGetConductCode

https://developer.github.com/v3/codes_of_conduct/#get-an-individual-code-of-conduct
*/
type CodesOfConductGetConductCodeResponse struct {
	response
	request *CodesOfConductGetConductCodeReq
	Data    *CodesOfConductGetConductCodeResponseBody
}

/*
CodesOfConductGetForRepo performs requests for "codes-of-conduct/get-for-repo"

Get the contents of a repository's code of conduct.

  GET /repos/{owner}/{repo}/community/code_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-the-contents-of-a-repositorys-code-of-conduct
*/
func (c *Client) CodesOfConductGetForRepo(ctx context.Context, req *CodesOfConductGetForRepoReq, opt ...RequestOption) (*CodesOfConductGetForRepoResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &CodesOfConductGetForRepoResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(CodesOfConductGetForRepoResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodesOfConductGetForRepoReq is request data for Client.CodesOfConductGetForRepo

https://developer.github.com/v3/codes_of_conduct/#get-the-contents-of-a-repositorys-code-of-conduct
*/
type CodesOfConductGetForRepoReq struct {
	pgURL string
	Owner string
	Repo  string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

func (r *CodesOfConductGetForRepoReq) pagingURL() string {
	return r.pgURL
}

func (r *CodesOfConductGetForRepoReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/community/code_of_conduct", r.Owner, r.Repo)
}

func (r *CodesOfConductGetForRepoReq) method() string {
	return "GET"
}

func (r *CodesOfConductGetForRepoReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *CodesOfConductGetForRepoReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"scarlet-witch": r.ScarletWitchPreview}
	if requiredPreviews {
		previewVals["scarlet-witch"] = true
	}
	if allPreviews {
		previewVals["scarlet-witch"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *CodesOfConductGetForRepoReq) body() interface{} {
	return nil
}

func (r *CodesOfConductGetForRepoReq) dataStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetForRepoReq) validStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetForRepoReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *CodesOfConductGetForRepoReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodesOfConductGetForRepoReq) Rel(link RelName, resp *CodesOfConductGetForRepoResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
CodesOfConductGetForRepoResponseBody is a response body for CodesOfConductGetForRepo

https://developer.github.com/v3/codes_of_conduct/#get-the-contents-of-a-repositorys-code-of-conduct
*/
type CodesOfConductGetForRepoResponseBody struct {
	components.CodeOfConduct
}

/*
CodesOfConductGetForRepoResponse is a response for CodesOfConductGetForRepo

https://developer.github.com/v3/codes_of_conduct/#get-the-contents-of-a-repositorys-code-of-conduct
*/
type CodesOfConductGetForRepoResponse struct {
	response
	request *CodesOfConductGetForRepoReq
	Data    *CodesOfConductGetForRepoResponseBody
}
