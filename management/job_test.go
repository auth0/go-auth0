package management

import (
	"testing"

	"github.com/auth0/go-auth0"
)

func TestJob(t *testing.T) {
	var err error

	c, err := m.Connection.ReadByName("Username-Password-Authentication")
	if err != nil {
		t.Error(err)
	}
	connectionID := auth0.StringValue(c.ID)

	u := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("example@example.com"),
		Username:   auth0.String("example"),
		Password:   auth0.String("I have a password and its a secret"),
	}
	err = m.User.Create(u)
	if err != nil {
		t.Error(err)
	}
	userID := auth0.StringValue(u.ID)

	defer m.User.Delete(userID)

	var jobID string

	t.Run("VerifyEmail", func(t *testing.T) {
		job := &Job{
			UserID: auth0.String(userID),
		}

		err = m.Job.VerifyEmail(job)
		if err != nil {
			t.Error(err)
		}

		jobID = auth0.StringValue(job.ID) // save for use in subsequent tests

		if auth0.StringValue(job.Type) != "verification_email" {
			t.Errorf("expected job type to be verification_email, got %s", auth0.StringValue(job.Type))
		}
	})

	t.Run("Read", func(t *testing.T) {
		job, err := m.Job.Read(jobID)
		if err != nil {
			t.Error(err)
		}
		t.Log(job)
	})

	t.Run("ExportUsers", func(t *testing.T) {
		job := &Job{
			ConnectionID: auth0.String(connectionID),
			Format:       auth0.String("json"),
			Limit:        auth0.Int(5),
			Fields: []map[string]interface{}{
				{"name": "name"},
				{"name": "email"},
				{"name": "identities[0].connection"},
			},
		}
		err := m.Job.ExportUsers(job)
		if err != nil {
			t.Error(err)
		}
		t.Log(job)
	})

	t.Run("ImportUsers", func(t *testing.T) {
		job := &Job{
			ConnectionID:        auth0.String(connectionID),
			Upsert:              auth0.Bool(true),
			SendCompletionEmail: auth0.Bool(false),
			Users: []map[string]interface{}{
				{"email": "alex@example.com", "email_verified": true},
			},
		}
		err = m.Job.ImportUsers(job)
		if err != nil {
			t.Error(err)
		}
		t.Log(job)
	})
}
