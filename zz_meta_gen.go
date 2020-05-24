// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

/*
MetaGetReq builds requests for "meta/get"

Get.

  GET /meta

https://developer.github.com/v3/meta/#meta
*/
type MetaGetReq struct{}

func (r *MetaGetReq) urlPath() string {
	return fmt.Sprintf("/meta")
}

func (r *MetaGetReq) method() string {
	return "GET"
}

func (r *MetaGetReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *MetaGetReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *MetaGetReq) body() interface{} {
	return nil
}

// HTTPRequest creates an http request
func (r *MetaGetReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
MetaGetResponseBody200 is a response body for meta/get

API documentation: https://developer.github.com/v3/meta/#meta
*/
type MetaGetResponseBody200 struct {
	Api                []string `json:"api,omitempty"`
	Git                []string `json:"git,omitempty"`
	Hooks              []string `json:"hooks,omitempty"`
	Importer           []string `json:"importer,omitempty"`
	Pages              []string `json:"pages,omitempty"`
	SshKeyFingerprints struct {
		MD5DSA    string `json:"MD5_DSA,omitempty"`
		MD5RSA    string `json:"MD5_RSA,omitempty"`
		SHA256DSA string `json:"SHA256_DSA,omitempty"`
		SHA256RSA string `json:"SHA256_RSA,omitempty"`
	} `json:"ssh_key_fingerprints,omitempty"`
	VerifiablePasswordAuthentication bool     `json:"verifiable_password_authentication,omitempty"`
	Web                              []string `json:"web,omitempty"`
}
