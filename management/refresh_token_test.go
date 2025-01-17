package management

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

// TestRefreshTokenManager_Read tests the RefreshTokenManager Read method.
// This E2E test is skipped because refresh tokens cannot be created without UI interaction.
func TestRefreshTokenManager_Read(t *testing.T) {
	skipTestIfRunningE2E(t)
	configureHTTPTestRecordings(t)

	// RecordingNote: This test recording was manually generated to match these details.
	// If any changes occur here, the test recording will need manual modification.
	expectedRefreshToken := &RefreshToken{
		ID:        auth0.String("RefreshTokenID"),
		UserID:    auth0.String("UserID"),
		CreatedAt: auth0.Time(time.Date(2024, 5, 1, 13, 0, 30, 38000000, time.UTC)),
		ClientID:  auth0.String("CLIENTID"),
		Rotating:  auth0.Bool(false),
		ResourceServer: []*RefreshTokenResourceServer{
			{
				Audience: auth0.String("https://go-auth0-dev.eu.auth0.com.us.auth0.com/api/v2/"),
				Scopes:   auth0.String("openid profile offline_access"),
			},
		},
	}
	actualRefreshToken, err := api.RefreshToken.Read(context.Background(), expectedRefreshToken.GetID())
	require.NoError(t, err)
	require.Equal(t, expectedRefreshToken, actualRefreshToken)
}

// TestRefreshTokenManager_Delete tests the RefreshTokenManager Delete method.
// This E2E test is skipped because refresh tokens cannot be created without UI interaction.
func TestRefreshTokenManager_Delete(t *testing.T) {
	skipTestIfRunningE2E(t)
	configureHTTPTestRecordings(t)

	// RecordingNote: This test recording was manually generated to match these details.
	// If any changes occur here, the test recording will need manual modification.
	expectedRefreshToken := &RefreshToken{
		ID:        auth0.String("RefreshTokenID"),
		UserID:    auth0.String("UserID"),
		CreatedAt: auth0.Time(time.Date(2024, 5, 1, 13, 0, 30, 38000000, time.UTC)),
		ClientID:  auth0.String("CLIENTID"),
		Rotating:  auth0.Bool(false),
		ResourceServer: []*RefreshTokenResourceServer{
			{
				Audience: auth0.String("https://go-auth0-dev.eu.auth0.com.us.auth0.com/api/v2/"),
				Scopes:   auth0.String("openid profile offline_access"),
			},
		},
	}
	actualRefreshToken, err := api.RefreshToken.Read(context.Background(), expectedRefreshToken.GetID())
	require.NoError(t, err)

	err = api.RefreshToken.Delete(context.Background(), actualRefreshToken.GetID())
	require.NoError(t, err)
}
