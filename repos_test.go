package octo_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/components"
)

func TestReposGetContent(t *testing.T) {
	t.Run("file", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), patAuth())

		response, err := client.ReposGetContent(ctx, &octo.ReposGetContentReq{
			Owner: "WillAbides",
			Repo:  "octo-go",
			Path:  "generator/main.go",
		})
		require.NoError(t, err)
		gotVal, ok := response.Data.Value().(components.ContentFile)
		require.True(t, ok)
		require.Equal(t, "main.go", gotVal.Name)
	})

	t.Run("directory", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), patAuth())

		response, err := client.ReposGetContent(ctx, &octo.ReposGetContentReq{
			Owner: "WillAbides",
			Repo:  "octo-go",
			Path:  "generator",
		})
		require.NoError(t, err)
		fmt.Printf("%T\n", response.Data.Value())
		gotVal, ok := response.Data.Value().(components.ContentDirectory)
		require.True(t, ok)
		require.Greater(t, len(gotVal), 1)
	})
}

func TestReposDownloadTarballArchive(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), patAuth())

	resp, err := client.ReposDownloadTarballArchive(ctx, &octo.ReposDownloadTarballArchiveReq{
		Owner: "octocat",
		Repo:  "Hello-World",
		Ref:   "master",
	})
	require.NoError(t, err)
	g, err := ioutil.ReadAll(resp.HTTPResponse().Body)
	require.NoError(t, err)
	require.True(t, len(g) > 100)
}

func TestReposCompareCommits(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), patAuth())
	got, err := client.ReposCompareCommits(ctx, &octo.ReposCompareCommitsReq{
		Owner: "WillAbides",
		Repo:  "octo-go",
		Base:  "2c80673f15b275cbb22fc167cb1b9aa43ba7a27a",
		Head:  "ceac7c6d9a134326a0871174423bea19acbb122a",
	})
	require.NoError(t, err)
	require.Equal(t, "ahead", got.Data.Status)
}

func TestResposUploadReleaseAsset(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), patAuth())
		releaseResp, err := client.ReposGetReleaseByTag(ctx, &octo.ReposGetReleaseByTagReq{
			Owner: "octo-cli-testorg",
			Repo:  "scratch",
			Tag:   "v0.0.2",
		})
		require.NoError(t, err)
		licenseBytes, err := ioutil.ReadFile("LICENSE")
		require.NoError(t, err)
		uploadResp, err := client.ReposUploadReleaseAsset(ctx, &octo.ReposUploadReleaseAssetReq{
			URL:               releaseResp.Data.UploadUrl,
			Name:              octo.String("LICENSE"),
			RequestBody:       bytes.NewBuffer(licenseBytes),
			ContentTypeHeader: octo.String("text/plain"),
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, err := client.ReposDeleteReleaseAsset(ctx, &octo.ReposDeleteReleaseAssetReq{
				Owner:   "octo-cli-testorg",
				Repo:    "scratch",
				AssetId: uploadResp.Data.Id,
			})
			require.NoError(t, err)
		})

		require.Equal(t, "LICENSE", uploadResp.Data.Name)
	})
}
