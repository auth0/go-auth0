package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestSessionManager_Read(t *testing.T) {
	skipTestIfRunningE2E(t)
	configureHTTPTestRecordings(t)

	// RecordingNote: This test recording was manually generated to match these details.
	// If any changes occur here, the test recording will need manual modification.
	expectedSession := &Session{
		ID:               auth0.String("SessionID"),
		UserID:           auth0.String("UserID"),
		CreatedAt:        auth0.String("2024-05-01T13:00:30.380Z"),
		UpdatedAt:        auth0.String("2024-05-01T13:00:30.380Z"),
		AuthenticatedAt:  auth0.String("2024-05-01T13:00:30.380Z"),
		IdleExpiresAt:    auth0.String("2024-05-01T13:00:30.380Z"),
		ExpiresAt:        auth0.String("2024-05-01T13:00:30.380Z"),
		LastInteractedAt: auth0.String("2024-05-01T13:00:30.380Z"),
		Device: &SessionDevice{
			InitialUserAgent: auth0.String("InitialUserAgent"),
			InitialIP:        auth0.String("InitialIP"),
			InitialASN:       auth0.String("InitialASN"),
			LastUserAgent:    auth0.String("LastUserAgent"),
			LastIP:           auth0.String("LastIP"),
			LastASN:          auth0.String("LastASN"),
		},
		Clients: []*SessionClient{
			{
				ClientID: auth0.String("ClientID"),
			},
		},
		Authentication: &SessionAuthentication{
			Methods: []*SessionAuthenticationMethod{
				{
					Name:      auth0.String("Name"),
					Timestamp: auth0.String("2024-05-01T13:00:30.380Z"),
					Type:      auth0.String("Type"),
				},
			},
		},
	}
	actualSession, err := api.Session.Read(context.Background(), expectedSession.GetID())
	require.NoError(t, err)
	require.Equal(t, expectedSession, actualSession)
}

func TestSessionManager_Delete(t *testing.T) {
	skipTestIfRunningE2E(t)
	configureHTTPTestRecordings(t)

	err := api.Session.Delete(context.Background(), "SessionID")
	require.NoError(t, err)
}

func TestSessionManager_Revoke(t *testing.T) {
	skipTestIfRunningE2E(t)
	configureHTTPTestRecordings(t)

	err := api.Session.Revoke(context.Background(), "SessionID")
	require.NoError(t, err)
}
