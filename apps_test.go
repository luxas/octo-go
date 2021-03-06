package octo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
)

func TestAppsCreateInstallationAccessToken(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), appAuth(t), octo.WithRequiredPreviews())
	token, err := client.AppsCreateInstallationAccessToken(ctx, &octo.AppsCreateInstallationAccessTokenReq{
		InstallationId: appInstallationID,
	})
	require.NoError(t, err)

	// revoke the token so that we don't inadvertently commit a valid token to git
	_, err = client.AppsRevokeInstallationAccessToken(ctx, nil, octo.WithPATAuth(token.Data.Token))
	require.NoError(t, err)
}

func TestAppsGetRepoInstallation(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), appAuth(t))
		installation, err := client.AppsGetRepoInstallation(ctx, &octo.AppsGetRepoInstallationReq{
			Owner:             "WillAbides",
			Repo:              "octo-go",
			MachineManPreview: true,
		})
		require.NoError(t, err)
		require.Equal(t, appInstallationID, installation.Data.Id)
	})

	t.Run("no installation", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), appAuth(t))
		installation, err := client.AppsGetRepoInstallation(ctx, &octo.AppsGetRepoInstallationReq{
			Owner:             "torvalds",
			Repo:              "linux",
			MachineManPreview: true,
		})
		require.Error(t, err)
		require.Empty(t, installation.Data)
	})
}

func TestAppsGetUserInstallation(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), appAuth(t))
		installation, err := client.AppsGetUserInstallation(ctx, &octo.AppsGetUserInstallationReq{
			Username:          "WillAbides",
			MachineManPreview: true,
		})
		require.NoError(t, err)
		require.Equal(t, appInstallationID, installation.Data.Id)
	})
}
