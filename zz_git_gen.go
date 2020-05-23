// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

/*
GitCreateBlobReq builds requests for "git/create-blob"

Create a blob.

  POST /repos/{owner}/{repo}/git/blobs

https://developer.github.com/v3/git/blobs/#create-a-blob
*/
type GitCreateBlobReq struct {
	Owner       string
	Repo        string
	RequestBody GitCreateBlobReqBody
}

func (r GitCreateBlobReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/blobs", r.Owner, r.Repo)
}

func (r GitCreateBlobReq) method() string {
	return "POST"
}

func (r GitCreateBlobReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitCreateBlobReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitCreateBlobReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), r.RequestBody, opt)
}

/*
GitCreateBlobReqBody is a request body for git/create-blob

API documentation: https://developer.github.com/v3/git/blobs/#create-a-blob
*/
type GitCreateBlobReqBody struct {

	// The new blob's content.
	Content *string `json:"content"`

	/*
	   The encoding used for `content`. Currently, `"utf-8"` and `"base64"` are
	   supported.
	*/
	Encoding *string `json:"encoding,omitempty"`
}

/*
GitCreateBlobResponseBody201 is a response body for git/create-blob

API documentation: https://developer.github.com/v3/git/blobs/#create-a-blob
*/
type GitCreateBlobResponseBody201 struct {
	Sha string `json:"sha,omitempty"`
	Url string `json:"url,omitempty"`
}

/*
GitCreateCommitReq builds requests for "git/create-commit"

Create a commit.

  POST /repos/{owner}/{repo}/git/commits

https://developer.github.com/v3/git/commits/#create-a-commit
*/
type GitCreateCommitReq struct {
	Owner       string
	Repo        string
	RequestBody GitCreateCommitReqBody
}

func (r GitCreateCommitReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/commits", r.Owner, r.Repo)
}

func (r GitCreateCommitReq) method() string {
	return "POST"
}

func (r GitCreateCommitReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitCreateCommitReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitCreateCommitReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), r.RequestBody, opt)
}

// GitCreateCommitReqBodyAuthor is a value for GitCreateCommitReqBody's Author field
type GitCreateCommitReqBodyAuthor struct {

	/*
	   Indicates when this commit was authored (or committed). This is a timestamp in
	   [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format:
	   `YYYY-MM-DDTHH:MM:SSZ`.
	*/
	Date *string `json:"date,omitempty"`

	// The email of the author (or committer) of the commit
	Email *string `json:"email,omitempty"`

	// The name of the author (or committer) of the commit
	Name *string `json:"name,omitempty"`
}

// GitCreateCommitReqBodyCommitter is a value for GitCreateCommitReqBody's Committer field
type GitCreateCommitReqBodyCommitter struct {

	/*
	   Indicates when this commit was authored (or committed). This is a timestamp in
	   [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format:
	   `YYYY-MM-DDTHH:MM:SSZ`.
	*/
	Date *string `json:"date,omitempty"`

	// The email of the author (or committer) of the commit
	Email *string `json:"email,omitempty"`

	// The name of the author (or committer) of the commit
	Name *string `json:"name,omitempty"`
}

/*
GitCreateCommitReqBody is a request body for git/create-commit

API documentation: https://developer.github.com/v3/git/commits/#create-a-commit
*/
type GitCreateCommitReqBody struct {

	/*
	   Information about the author of the commit. By default, the `author` will be the
	   authenticated user and the current date. See the `author` and `committer` object
	   below for details.
	*/
	Author *GitCreateCommitReqBodyAuthor `json:"author,omitempty"`

	/*
	   Information about the person who is making the commit. By default, `committer`
	   will use the information set in `author`. See the `author` and `committer`
	   object below for details.
	*/
	Committer *GitCreateCommitReqBodyCommitter `json:"committer,omitempty"`

	// The commit message
	Message *string `json:"message"`

	/*
	   The SHAs of the commits that were the parents of this commit. If omitted or
	   empty, the commit will be written as a root commit. For a single parent, an
	   array of one SHA should be provided; for a merge commit, an array of more than
	   one should be provided.
	*/
	Parents []string `json:"parents"`

	/*
	   The [PGP signature](https://en.wikipedia.org/wiki/Pretty_Good_Privacy) of the
	   commit. GitHub adds the signature to the `gpgsig` header of the created commit.
	   For a commit signature to be verifiable by Git or GitHub, it must be an
	   ASCII-armored detached PGP signature over the string commit as it would be
	   written to the object database. To pass a `signature` parameter, you need to
	   first manually create a valid PGP signature, which can be complicated. You may
	   find it easier to [use the command
	   line](https://git-scm.com/book/id/v2/Git-Tools-Signing-Your-Work) to create
	   signed commits.
	*/
	Signature *string `json:"signature,omitempty"`

	// The SHA of the tree object this commit points to
	Tree *string `json:"tree"`
}

/*
GitCreateCommitResponseBody201 is a response body for git/create-commit

API documentation: https://developer.github.com/v3/git/commits/#create-a-commit
*/
type GitCreateCommitResponseBody201 struct {
	Author struct {
		Date  string `json:"date,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"author,omitempty"`
	Committer struct {
		Date  string `json:"date,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"committer,omitempty"`
	Message string `json:"message,omitempty"`
	NodeId  string `json:"node_id,omitempty"`
	Parents []struct {
		Sha string `json:"sha,omitempty"`
		Url string `json:"url,omitempty"`
	} `json:"parents,omitempty"`
	Sha  string `json:"sha,omitempty"`
	Tree struct {
		Sha string `json:"sha,omitempty"`
		Url string `json:"url,omitempty"`
	} `json:"tree,omitempty"`
	Url          string `json:"url,omitempty"`
	Verification struct {
		Payload   string `json:"payload,omitempty"`
		Reason    string `json:"reason,omitempty"`
		Signature string `json:"signature,omitempty"`
		Verified  bool   `json:"verified,omitempty"`
	} `json:"verification,omitempty"`
}

/*
GitCreateRefReq builds requests for "git/create-ref"

Create a reference.

  POST /repos/{owner}/{repo}/git/refs

https://developer.github.com/v3/git/refs/#create-a-reference
*/
type GitCreateRefReq struct {
	Owner       string
	Repo        string
	RequestBody GitCreateRefReqBody
}

func (r GitCreateRefReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/refs", r.Owner, r.Repo)
}

func (r GitCreateRefReq) method() string {
	return "POST"
}

func (r GitCreateRefReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitCreateRefReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitCreateRefReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), r.RequestBody, opt)
}

/*
GitCreateRefReqBody is a request body for git/create-ref

API documentation: https://developer.github.com/v3/git/refs/#create-a-reference
*/
type GitCreateRefReqBody struct {

	/*
	   The name of the fully qualified reference (ie: `refs/heads/master`). If it
	   doesn't start with 'refs' and have at least two slashes, it will be rejected.
	*/
	Ref *string `json:"ref"`

	// The SHA1 value for this reference.
	Sha *string `json:"sha"`
}

/*
GitCreateRefResponseBody201 is a response body for git/create-ref

API documentation: https://developer.github.com/v3/git/refs/#create-a-reference
*/
type GitCreateRefResponseBody201 struct {
	NodeId string `json:"node_id,omitempty"`
	Object struct {
		Sha  string `json:"sha,omitempty"`
		Type string `json:"type,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"object,omitempty"`
	Ref string `json:"ref,omitempty"`
	Url string `json:"url,omitempty"`
}

/*
GitCreateTagReq builds requests for "git/create-tag"

Create a tag object.

  POST /repos/{owner}/{repo}/git/tags

https://developer.github.com/v3/git/tags/#create-a-tag-object
*/
type GitCreateTagReq struct {
	Owner       string
	Repo        string
	RequestBody GitCreateTagReqBody
}

func (r GitCreateTagReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/tags", r.Owner, r.Repo)
}

func (r GitCreateTagReq) method() string {
	return "POST"
}

func (r GitCreateTagReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitCreateTagReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitCreateTagReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), r.RequestBody, opt)
}

// GitCreateTagReqBodyTagger is a value for GitCreateTagReqBody's Tagger field
type GitCreateTagReqBodyTagger struct {

	/*
	   When this object was tagged. This is a timestamp in [ISO
	   8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
	*/
	Date *string `json:"date,omitempty"`

	// The email of the author of the tag
	Email *string `json:"email,omitempty"`

	// The name of the author of the tag
	Name *string `json:"name,omitempty"`
}

/*
GitCreateTagReqBody is a request body for git/create-tag

API documentation: https://developer.github.com/v3/git/tags/#create-a-tag-object
*/
type GitCreateTagReqBody struct {

	// The tag message.
	Message *string `json:"message"`

	// The SHA of the git object this is tagging.
	Object *string `json:"object"`

	// The tag's name. This is typically a version (e.g., "v0.0.1").
	Tag *string `json:"tag"`

	// An object with information about the individual creating the tag.
	Tagger *GitCreateTagReqBodyTagger `json:"tagger,omitempty"`

	/*
	   The type of the object we're tagging. Normally this is a `commit` but it can
	   also be a `tree` or a `blob`.
	*/
	Type *string `json:"type"`
}

/*
GitCreateTagResponseBody201 is a response body for git/create-tag

API documentation: https://developer.github.com/v3/git/tags/#create-a-tag-object
*/
type GitCreateTagResponseBody201 struct {
	Message string `json:"message,omitempty"`
	NodeId  string `json:"node_id,omitempty"`
	Object  struct {
		Sha  string `json:"sha,omitempty"`
		Type string `json:"type,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"object,omitempty"`
	Sha    string `json:"sha,omitempty"`
	Tag    string `json:"tag,omitempty"`
	Tagger struct {
		Date  string `json:"date,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"tagger,omitempty"`
	Url          string `json:"url,omitempty"`
	Verification struct {
		Payload   string `json:"payload,omitempty"`
		Reason    string `json:"reason,omitempty"`
		Signature string `json:"signature,omitempty"`
		Verified  bool   `json:"verified,omitempty"`
	} `json:"verification,omitempty"`
}

/*
GitCreateTreeReq builds requests for "git/create-tree"

Create a tree.

  POST /repos/{owner}/{repo}/git/trees

https://developer.github.com/v3/git/trees/#create-a-tree
*/
type GitCreateTreeReq struct {
	Owner       string
	Repo        string
	RequestBody GitCreateTreeReqBody
}

func (r GitCreateTreeReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/trees", r.Owner, r.Repo)
}

func (r GitCreateTreeReq) method() string {
	return "POST"
}

func (r GitCreateTreeReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitCreateTreeReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitCreateTreeReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), r.RequestBody, opt)
}

// GitCreateTreeReqBodyTree is a value for GitCreateTreeReqBody's Tree field
type GitCreateTreeReqBodyTree struct {

	/*
	   The content you want this file to have. GitHub will write this blob out and use
	   that SHA for this entry. Use either this, or `tree.sha`.

	   **Note:** Use either `tree.sha` or `content` to specify the contents of the
	   entry. Using both `tree.sha` and `content` will return an error.
	*/
	Content *string `json:"content,omitempty"`

	/*
	   The file mode; one of `100644` for file (blob), `100755` for executable (blob),
	   `040000` for subdirectory (tree), `160000` for submodule (commit), or `120000`
	   for a blob that specifies the path of a symlink.
	*/
	Mode *string `json:"mode,omitempty"`

	// The file referenced in the tree.
	Path *string `json:"path,omitempty"`

	/*
	   The SHA1 checksum ID of the object in the tree. Also called `tree.sha`. If the
	   value is `null` then the file will be deleted.

	   **Note:** Use either `tree.sha` or `content` to specify the contents of the
	   entry. Using both `tree.sha` and `content` will return an error.
	*/
	Sha *string `json:"sha,omitempty"`

	// Either `blob`, `tree`, or `commit`.
	Type *string `json:"type,omitempty"`
}

/*
GitCreateTreeReqBody is a request body for git/create-tree

API documentation: https://developer.github.com/v3/git/trees/#create-a-tree
*/
type GitCreateTreeReqBody struct {

	/*
	   The SHA1 of the tree you want to update with new data. If you don't set this,
	   the commit will be created on top of everything; however, it will only contain
	   your change, the rest of your files will show up as deleted.
	*/
	BaseTree *string `json:"base_tree,omitempty"`

	// Objects (of `path`, `mode`, `type`, and `sha`) specifying a tree structure.
	Tree []GitCreateTreeReqBodyTree `json:"tree"`
}

/*
GitCreateTreeResponseBody201 is a response body for git/create-tree

API documentation: https://developer.github.com/v3/git/trees/#create-a-tree
*/
type GitCreateTreeResponseBody201 struct {
	Sha  string `json:"sha,omitempty"`
	Tree []struct {
		Mode string      `json:"mode,omitempty"`
		Path string      `json:"path,omitempty"`
		Sha  string      `json:"sha,omitempty"`
		Size json.Number `json:"size,omitempty"`
		Type string      `json:"type,omitempty"`
		Url  string      `json:"url,omitempty"`
	} `json:"tree,omitempty"`
	Url string `json:"url,omitempty"`
}

/*
GitDeleteRefReq builds requests for "git/delete-ref"

Delete a reference.

  DELETE /repos/{owner}/{repo}/git/refs/{ref}

https://developer.github.com/v3/git/refs/#delete-a-reference
*/
type GitDeleteRefReq struct {
	Owner string
	Repo  string
	Ref   string
}

func (r GitDeleteRefReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/refs/%v", r.Owner, r.Repo, r.Ref)
}

func (r GitDeleteRefReq) method() string {
	return "DELETE"
}

func (r GitDeleteRefReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitDeleteRefReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitDeleteRefReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
GitGetBlobReq builds requests for "git/get-blob"

Get a blob.

  GET /repos/{owner}/{repo}/git/blobs/{file_sha}

https://developer.github.com/v3/git/blobs/#get-a-blob
*/
type GitGetBlobReq struct {
	Owner   string
	Repo    string
	FileSha string
}

func (r GitGetBlobReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/blobs/%v", r.Owner, r.Repo, r.FileSha)
}

func (r GitGetBlobReq) method() string {
	return "GET"
}

func (r GitGetBlobReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitGetBlobReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitGetBlobReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
GitGetBlobResponseBody200 is a response body for git/get-blob

API documentation: https://developer.github.com/v3/git/blobs/#get-a-blob
*/
type GitGetBlobResponseBody200 struct {
	Content  string      `json:"content,omitempty"`
	Encoding string      `json:"encoding,omitempty"`
	Sha      string      `json:"sha,omitempty"`
	Size     json.Number `json:"size,omitempty"`
	Url      string      `json:"url,omitempty"`
}

/*
GitGetCommitReq builds requests for "git/get-commit"

Get a commit.

  GET /repos/{owner}/{repo}/git/commits/{commit_sha}

https://developer.github.com/v3/git/commits/#get-a-commit
*/
type GitGetCommitReq struct {
	Owner     string
	Repo      string
	CommitSha string
}

func (r GitGetCommitReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/commits/%v", r.Owner, r.Repo, r.CommitSha)
}

func (r GitGetCommitReq) method() string {
	return "GET"
}

func (r GitGetCommitReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitGetCommitReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitGetCommitReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
GitGetCommitResponseBody200 is a response body for git/get-commit

API documentation: https://developer.github.com/v3/git/commits/#get-a-commit
*/
type GitGetCommitResponseBody200 struct {
	Author struct {
		Date  string `json:"date,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"author,omitempty"`
	Committer struct {
		Date  string `json:"date,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"committer,omitempty"`
	Message string `json:"message,omitempty"`
	Parents []struct {
		Sha string `json:"sha,omitempty"`
		Url string `json:"url,omitempty"`
	} `json:"parents,omitempty"`
	Sha  string `json:"sha,omitempty"`
	Tree struct {
		Sha string `json:"sha,omitempty"`
		Url string `json:"url,omitempty"`
	} `json:"tree,omitempty"`
	Url          string `json:"url,omitempty"`
	Verification struct {
		Payload   string `json:"payload,omitempty"`
		Reason    string `json:"reason,omitempty"`
		Signature string `json:"signature,omitempty"`
		Verified  bool   `json:"verified,omitempty"`
	} `json:"verification,omitempty"`
}

/*
GitGetRefReq builds requests for "git/get-ref"

Get a single reference.

  GET /repos/{owner}/{repo}/git/ref/{ref}

https://developer.github.com/v3/git/refs/#get-a-single-reference
*/
type GitGetRefReq struct {
	Owner string
	Repo  string
	Ref   string
}

func (r GitGetRefReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/ref/%v", r.Owner, r.Repo, r.Ref)
}

func (r GitGetRefReq) method() string {
	return "GET"
}

func (r GitGetRefReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitGetRefReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitGetRefReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
GitGetRefResponseBody200 is a response body for git/get-ref

API documentation: https://developer.github.com/v3/git/refs/#get-a-single-reference
*/
type GitGetRefResponseBody200 struct {
	NodeId string `json:"node_id,omitempty"`
	Object struct {
		Sha  string `json:"sha,omitempty"`
		Type string `json:"type,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"object,omitempty"`
	Ref string `json:"ref,omitempty"`
	Url string `json:"url,omitempty"`
}

/*
GitGetTagReq builds requests for "git/get-tag"

Get a tag.

  GET /repos/{owner}/{repo}/git/tags/{tag_sha}

https://developer.github.com/v3/git/tags/#get-a-tag
*/
type GitGetTagReq struct {
	Owner  string
	Repo   string
	TagSha string
}

func (r GitGetTagReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/tags/%v", r.Owner, r.Repo, r.TagSha)
}

func (r GitGetTagReq) method() string {
	return "GET"
}

func (r GitGetTagReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitGetTagReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitGetTagReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
GitGetTagResponseBody200 is a response body for git/get-tag

API documentation: https://developer.github.com/v3/git/tags/#get-a-tag
*/
type GitGetTagResponseBody200 struct {
	Message string `json:"message,omitempty"`
	NodeId  string `json:"node_id,omitempty"`
	Object  struct {
		Sha  string `json:"sha,omitempty"`
		Type string `json:"type,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"object,omitempty"`
	Sha    string `json:"sha,omitempty"`
	Tag    string `json:"tag,omitempty"`
	Tagger struct {
		Date  string `json:"date,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"tagger,omitempty"`
	Url          string `json:"url,omitempty"`
	Verification struct {
		Payload   string `json:"payload,omitempty"`
		Reason    string `json:"reason,omitempty"`
		Signature string `json:"signature,omitempty"`
		Verified  bool   `json:"verified,omitempty"`
	} `json:"verification,omitempty"`
}

/*
GitGetTreeReq builds requests for "git/get-tree"

Get a tree.

  GET /repos/{owner}/{repo}/git/trees/{tree_sha}

https://developer.github.com/v3/git/trees/#get-a-tree
*/
type GitGetTreeReq struct {

	// owner parameter
	Owner string

	// repo parameter
	Repo string

	// tree_sha parameter
	TreeSha string

	// recursive parameter
	Recursive *int64
}

func (r GitGetTreeReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/trees/%v", r.Owner, r.Repo, r.TreeSha)
}

func (r GitGetTreeReq) method() string {
	return "GET"
}

func (r GitGetTreeReq) urlQuery() url.Values {
	query := url.Values{}
	if r.Recursive != nil {
		query.Set("recursive", strconv.FormatInt(*r.Recursive, 10))
	}
	return query
}

func (r GitGetTreeReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitGetTreeReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
GitGetTreeResponseBody200 is a response body for git/get-tree

API documentation: https://developer.github.com/v3/git/trees/#get-a-tree
*/
type GitGetTreeResponseBody200 struct {
	Sha  string `json:"sha,omitempty"`
	Tree []struct {
		Mode string      `json:"mode"`
		Path string      `json:"path"`
		Sha  string      `json:"sha"`
		Size json.Number `json:"size"`
		Type string      `json:"type"`
		Url  string      `json:"url"`
	} `json:"tree,omitempty"`
	Truncated bool   `json:"truncated,omitempty"`
	Url       string `json:"url,omitempty"`
}

/*
GitListMatchingRefsReq builds requests for "git/list-matching-refs"

List matching references.

  GET /repos/{owner}/{repo}/git/matching-refs/{ref}

https://developer.github.com/v3/git/refs/#list-matching-references
*/
type GitListMatchingRefsReq struct {
	Owner string
	Repo  string
	Ref   string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64
}

func (r GitListMatchingRefsReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/matching-refs/%v", r.Owner, r.Repo, r.Ref)
}

func (r GitListMatchingRefsReq) method() string {
	return "GET"
}

func (r GitListMatchingRefsReq) urlQuery() url.Values {
	query := url.Values{}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}
	return query
}

func (r GitListMatchingRefsReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitListMatchingRefsReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
GitListMatchingRefsResponseBody200 is a response body for git/list-matching-refs

API documentation: https://developer.github.com/v3/git/refs/#list-matching-references
*/
type GitListMatchingRefsResponseBody200 []struct {
	NodeId string `json:"node_id,omitempty"`
	Object struct {
		Sha  string `json:"sha,omitempty"`
		Type string `json:"type,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"object,omitempty"`
	Ref string `json:"ref,omitempty"`
	Url string `json:"url,omitempty"`
}

/*
GitUpdateRefReq builds requests for "git/update-ref"

Update a reference.

  PATCH /repos/{owner}/{repo}/git/refs/{ref}

https://developer.github.com/v3/git/refs/#update-a-reference
*/
type GitUpdateRefReq struct {
	Owner       string
	Repo        string
	Ref         string
	RequestBody GitUpdateRefReqBody
}

func (r GitUpdateRefReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/git/refs/%v", r.Owner, r.Repo, r.Ref)
}

func (r GitUpdateRefReq) method() string {
	return "PATCH"
}

func (r GitUpdateRefReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r GitUpdateRefReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r GitUpdateRefReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), r.RequestBody, opt)
}

/*
GitUpdateRefReqBody is a request body for git/update-ref

API documentation: https://developer.github.com/v3/git/refs/#update-a-reference
*/
type GitUpdateRefReqBody struct {

	/*
	   Indicates whether to force the update or to make sure the update is a
	   fast-forward update. Leaving this out or setting it to `false` will make sure
	   you're not overwriting work.
	*/
	Force *bool `json:"force,omitempty"`

	// The SHA1 value to set this reference to
	Sha *string `json:"sha"`
}

/*
GitUpdateRefResponseBody200 is a response body for git/update-ref

API documentation: https://developer.github.com/v3/git/refs/#update-a-reference
*/
type GitUpdateRefResponseBody200 struct {
	NodeId string `json:"node_id,omitempty"`
	Object struct {
		Sha  string `json:"sha,omitempty"`
		Type string `json:"type,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"object,omitempty"`
	Ref string `json:"ref,omitempty"`
	Url string `json:"url,omitempty"`
}
