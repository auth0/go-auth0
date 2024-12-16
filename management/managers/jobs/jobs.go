package jobs

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/auth0/go-auth0/management/models"
)

// Manager defines the base manager interface
type Manager struct {
	management *management.Management
}

// NewManager creates a new manager for  operations
func NewManager(mgmt *management.Management) *Manager {
	return &Manager{
		management: mgmt,
	}
}

// GetErrors Get job error details
//
// https://auth0.com/docs/api/management/v2/#!/Jobs/get_errors
func (m *Manager) GetErrors(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetErrors200Response, error) {
	var localVarReturnValue *models.GetErrors200Response
	err := m.management.Request(ctx, "GET", m.management.URI("jobs", management.SafeString(id), "errors"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetJobsById Get a job
//
// https://auth0.com/docs/api/management/v2/#!/Jobs/get_jobs_by_id
func (m *Manager) GetJobsById(ctx context.Context, id string, opts ...management.RequestOption) (*models.Job, error) {
	var localVarReturnValue *models.Job
	err := m.management.Request(ctx, "GET", m.management.URI("jobs", management.SafeString(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostUsersExports Create export users job
//
// https://auth0.com/docs/api/management/v2/#!/Jobs/post_users_exports
func (m *Manager) PostUsersExports(ctx context.Context, postUsersExportsRequest *models.PostUsersExportsRequest, opts ...management.RequestOption) (*models.Job, error) {
	var localVarReturnValue *models.Job
	err := m.management.Request(ctx, "POST", m.management.URI("jobs", "users-exports"), postUsersExportsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostUsersImports Create import users job
//
// https://auth0.com/docs/api/management/v2/#!/Jobs/post_users_imports
func (m *Manager) PostUsersImports(ctx context.Context, opts ...management.RequestOption) (*models.Job, error) {
	var localVarReturnValue *models.Job
	err := m.management.Request(ctx, "POST", m.management.URI("jobs", "users-imports"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostVerificationEmail Send an email address verification email
//
// https://auth0.com/docs/api/management/v2/#!/Jobs/post_verification_email
func (m *Manager) PostVerificationEmail(ctx context.Context, postVerificationEmailRequest *models.PostVerificationEmailRequest, opts ...management.RequestOption) (*models.Job, error) {
	var localVarReturnValue *models.Job
	err := m.management.Request(ctx, "POST", m.management.URI("jobs", "verification-email"), postVerificationEmailRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
