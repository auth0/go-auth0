package management

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestJobManager_VerifyEmail(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	job := &Job{
		UserID:   user.ID,
		Identity: user.Identities[0],
	}

	err := api.Job.VerifyEmail(job)
	assert.NoError(t, err)

	actualJob, err := api.Job.Read(job.GetID())
	assert.NoError(t, err)
	assert.Equal(t, job.GetID(), actualJob.GetID())
}

func TestJobManager_ExportUsers(t *testing.T) {
	configureHTTPTestRecordings(t)

	givenAUser(t)
	conn, err := api.Connection.ReadByName("Username-Password-Authentication")
	assert.NoError(t, err)

	job := &Job{
		ConnectionID: conn.ID,
		Format:       auth0.String("json"),
		Limit:        auth0.Int(5),
		Fields: []map[string]interface{}{
			{"name": "name"},
			{"name": "email"},
			{"name": "identities[0].connection"},
		},
	}

	err = api.Job.ExportUsers(job)
	assert.NoError(t, err)
}

func TestJobManager_ImportUsers(t *testing.T) {
	configureHTTPTestRecordings(t)

	conn, err := api.Connection.ReadByName("Username-Password-Authentication")
	require.NoError(t, err)

	job := &Job{
		ConnectionID:        conn.ID,
		Upsert:              auth0.Bool(true),
		SendCompletionEmail: auth0.Bool(false),
		Users: []map[string]interface{}{
			{
				"email":          "auzironian@example.com",
				"email_verified": true,
			},
		},
	}
	err = api.Job.ImportUsers(job)
	assert.NoError(t, err)

	// Let's give the ImportUsers job enough time to complete,
	// so we can ensure the Read Job has the summary field set.
	time.Sleep(time.Second * 2)

	actualJob, err := api.Job.Read(job.GetID())
	assert.NoError(t, err)
	assert.NotEmpty(t, actualJob.GetSummary())
	assert.Equal(t, 0, actualJob.GetSummary().GetFailed())
	assert.Equal(t, 0, actualJob.GetSummary().GetUpdated())
	assert.Equal(t, 1, actualJob.GetSummary().GetInserted())
	assert.Equal(t, 1, actualJob.GetSummary().GetTotal())

	t.Cleanup(func() {
		users, err := api.User.ListByEmail("auzironian@example.com")
		assert.NoError(t, err)
		assert.Len(t, users, 1)

		cleanupUser(t, users[0].GetID())
	})
}

func TestJobManager_ReadErrors(t *testing.T) {
	configureHTTPTestRecordings(t)

	alreadyExistingUser := givenAUser(t)
	conn, err := api.Connection.ReadByName("Username-Password-Authentication")
	require.NoError(t, err)

	job := &Job{
		ConnectionID: conn.ID,
		Users: []map[string]interface{}{
			{
				"email":          alreadyExistingUser.GetEmail(),
				"email_verified": true,
			},
		},
	}
	err = api.Job.ImportUsers(job)
	assert.NoError(t, err)

	// Let's give the ImportUsers job enough time to complete.
	time.Sleep(time.Second * 2)

	expectedJobErrors := JobError{
		User: map[string]interface{}{
			"email":          alreadyExistingUser.GetEmail(),
			"email_verified": true,
		},
		Errors: []JobUserErrors{
			{
				Code:    "DUPLICATED_USER",
				Message: "The user already exist and upsert parameter is false",
			},
		},
	}

	actualJobErrors, err := api.Job.ReadErrors(job.GetID())
	assert.NoError(t, err)
	assert.Len(t, actualJobErrors, 1)
	assert.Equal(t, expectedJobErrors, actualJobErrors[0])
}
