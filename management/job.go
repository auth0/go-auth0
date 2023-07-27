package management

import (
	"bytes"
	"context"
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
	// The job's identifier. Useful to retrieve its status.
	ID *string `json:"id,omitempty"`
	// The job's status.
	Status *string `json:"status,omitempty"`
	// The type of job.
	Type *string `json:"type,omitempty"`
	// The date when the job was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The user_id of the user to whom the email will be sent.
	UserID *string `json:"user_id,omitempty"`
	// The ID of the client, if not provided the global one will be used.
	ClientID *string `json:"client_id,omitempty"`
	// The identity of the user. This must be provided to verify primary social, enterprise and passwordless email identities. Also, is needed to verify secondary identities.
	Identity *UserIdentity `json:"-"`
	// The ID of the connection.
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
	// A list of fields to be included in the CSV. If omitted,
	// a set of predefined fields will be exported.
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
	// If a job is completed, the job status response will include a summary.
	Summary *JobSummary `json:"summary,omitempty"`
	// The ID of the Organization. If provided, organization parameters will be made available to the email template and organization branding will
	// be applied to the prompt. In addition, the redirect link in the prompt will include organization_id and organization_name query string parameters.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// MarshalJSON is a custom serializer for the Job type.
func (j *Job) MarshalJSON() ([]byte, error) {
	type job Job
	type identity struct {
		UserID   *string `json:"user_id,omitempty"`
		Provider *string `json:"provider,omitempty"`
	}
	type jobWrapper struct {
		*job
		Identity *identity `json:"identity,omitempty"`
	}

	w := &jobWrapper{job: (*job)(j)}
	if j.Identity != nil {
		w.Identity = &identity{
			UserID:   j.Identity.UserID,
			Provider: j.Identity.Provider,
		}
	}

	return json.Marshal(w)
}

// JobSummary includes totals of successful,
// failed, inserted, and updated records.
type JobSummary struct {
	Failed   *int `json:"failed,omitempty"`
	Updated  *int `json:"updated,omitempty"`
	Inserted *int `json:"inserted,omitempty"`
	Total    *int `json:"total,omitempty"`
}

// JobError is used to check for errors during jobs.
//
// See: https://auth0.com/docs/manage-users/user-migration/bulk-user-imports#retrieve-failed-entries
type JobError struct {
	User   map[string]interface{} `json:"user,omitempty"`
	Errors []JobUserErrors        `json:"errors,omitempty"`
}

// JobUserErrors holds errors for the specific user during a job.
type JobUserErrors struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Path    string `json:"path,omitempty"`
}

// JobManager manages Auth0 Job resources.
type JobManager manager

// Read retrieves a job. Useful to check its status.
//
// See: https://auth0.com/docs/api/management/v2#!/Jobs/get_jobs_by_id
func (m *JobManager) Read(ctx context.Context, id string, opts ...RequestOption) (j *Job, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("jobs", id), &j, opts...)
	return
}

// ReadErrors retrieves error details of a failed job.
//
// See: https://auth0.com/docs/api/management/v2#!/Jobs/get_errors
func (m *JobManager) ReadErrors(ctx context.Context, id string, opts ...RequestOption) (jobErrors []JobError, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("jobs", id, "errors"), &jobErrors, opts...)
	return
}

// VerifyEmail sends an email to the specified user that asks them to
// click a link to verify their email address.
func (m *JobManager) VerifyEmail(ctx context.Context, j *Job, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("jobs", "verification-email"), j, opts...)
}

// ExportUsers exports all users to a file via a long-running job.
//
// See: https://auth0.com/docs/api/management/v2#!/Jobs/post_users_exports
func (m *JobManager) ExportUsers(ctx context.Context, j *Job, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("jobs", "users-exports"), j, opts...)
}

// ImportUsers imports users from a formatted file into a connection via a long-running job.
//
// See: https://auth0.com/docs/api/management/v2#!/Jobs/post_users_imports
func (m *JobManager) ImportUsers(ctx context.Context, j *Job, opts ...RequestOption) error {
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

	request, err := http.NewRequestWithContext(ctx, "POST", m.management.URI("jobs", "users-imports"), &payload)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", mp.FormDataContentType())

	for _, option := range opts {
		option.apply(request)
	}

	response, err := m.management.Do(request)
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
