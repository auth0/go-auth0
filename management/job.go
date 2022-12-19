package management

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
	"time"
)

// Job is used for importing/exporting users or for
// sending email address verification emails.
//
// See: https://auth0.com/docs/manage-users/user-migration/bulk-user-imports
type Job struct {
	// The job's identifier. Useful to retrieve its status
	ID *string `json:"id,omitempty"`
	// The job's status
	Status *string `json:"status,omitempty"`
	// The type of job
	Type *string `json:"type,omitempty"`
	// The date when the job was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The user_id of the user to whom the email will be sent
	UserID *string `json:"user_id,omitempty"`
	// The id of the client, if not provided the global one will be used
	ClientID *string `json:"client_id,omitempty"`

	// The id of the connection.
	ConnectionID *string `json:"connection_id,omitempty"`
	// The url to download the result of the job.
	Location *string `json:"location,omitempty"`
	// The percentage of the work done so far.
	PercentageDone *int `json:"percentage_done,omitempty"`
	// Estimated amount of time remaining to finish the job.
	TimeLeftSeconds *int `json:"time_left_seconds,omitempty"`
	// The format of the file. Valid values are: "json" and "csv".
	Format *string `json:"format,omitempty"`
	// Limit the number of records.
	Limit *int `json:"limit,omitempty"`
	// A list of fields to be included in the CSV. If omitted, a set of
	// predefined fields will be exported.
	Fields []map[string]interface{} `json:"fields,omitempty"`

	// A list of users. Used when importing users in bulk.
	Users []map[string]interface{} `json:"users,omitempty"`
	// If false, users will only be inserted. If there are already user(s) with
	// the same emails as one or more of those being inserted, they will fail.
	// If this value is set to true and the user being imported already exists,
	// the user will be updated with the new information.
	Upsert *bool `json:"upsert,omitempty"`
	// Optional user defined string that can be used for correlating multiple
	// jobs, and is returned as part of the job status response.
	ExternalID *string `json:"external_id,omitempty"`
	// When true, sends a completion email to all tenant owners when the job is
	// finished. The default is true, so you must explicitly set this parameter
	// to false if you do not want emails sent.
	SendCompletionEmail *bool `json:"send_completion_email,omitempty"`
}

// JobManager manages Auth0 Job resources.
type JobManager struct {
	*Management
}

func newJobManager(m *Management) *JobManager {
	return &JobManager{m}
}

// VerifyEmail sends an email to the specified user that asks them to
// click a link to verify their email address.
func (m *JobManager) VerifyEmail(j *Job, opts ...RequestOption) error {
	return m.Request("POST", m.URI("jobs", "verification-email"), j, opts...)
}

// Read retrieves a job. Useful to check its status.
//
// See: https://auth0.com/docs/api/management/v2#!/Jobs/get_jobs_by_id
func (m *JobManager) Read(id string, opts ...RequestOption) (j *Job, err error) {
	err = m.Request("GET", m.URI("jobs", id), &j)
	return
}

// ExportUsers exports all users to a file via a long-running job.
//
// See: https://auth0.com/docs/api/management/v2#!/Jobs/post_users_exports
func (m *JobManager) ExportUsers(j *Job, opts ...RequestOption) error {
	return m.Request("POST", m.URI("jobs", "users-exports"), j, opts...)
}

// ImportUsers imports users from a formatted file into a connection via a long-running job.
//
// See: https://auth0.com/docs/api/management/v2#!/Jobs/post_users_imports
func (m *JobManager) ImportUsers(j *Job, opts ...RequestOption) error {
	var payload bytes.Buffer
	mp := multipart.NewWriter(&payload)

	if err := mp.WriteField("connection_id", j.GetConnectionID()); err != nil {
		return err
	}
	if err := mp.WriteField("upsert", strconv.FormatBool(j.GetUpsert())); err != nil {
		return err
	}
	if err := mp.WriteField("external_id", j.GetExternalID()); err != nil {
		return err
	}
	if err := mp.WriteField("send_completion_email", strconv.FormatBool(j.GetSendCompletionEmail())); err != nil {
		return err
	}

	if j.Users != nil {
		usersJSON, err := json.Marshal(j.Users)
		if err != nil {
			return err
		}

		header := textproto.MIMEHeader{}
		header.Set("Content-Disposition", `form-data; name="users"; filename="users.json"`)
		header.Set("Content-Type", "application/json")

		writer, err := mp.CreatePart(header)
		if err != nil {
			return err
		}

		if _, err := writer.Write(usersJSON); err != nil {
			return err
		}
	}
	if err := mp.Close(); err != nil {
		return err
	}

	request, err := http.NewRequest("POST", m.URI("jobs", "users-imports"), &payload)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", mp.FormDataContentType())

	for _, option := range opts {
		option.apply(request)
	}

	response, err := m.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= http.StatusBadRequest {
		return newError(response)
	}

	if response.StatusCode != http.StatusNoContent {
		return json.NewDecoder(response.Body).Decode(j)
	}

	return nil
}
