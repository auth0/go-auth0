package management

import (
	"context"
	"fmt"
	"slices"
)

// CustomPrompt to be used on authentication pages.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#-signup-prompt-entry-points
type CustomPrompt struct {
	FormContentStart      string `json:"form-content-start,omitempty"`
	FormContentEnd        string `json:"form-content-end,omitempty"`
	FormFooterStart       string `json:"form-footer-start,omitempty"`
	FormFooterEnd         string `json:"form-footer-end,omitempty"`
	SecondaryActionsStart string `json:"secondary-actions-start,omitempty"`
	SecondaryActionsEnd   string `json:"secondary-actions-end,omitempty"`

	// Prompt state for custom prompt. Options are:
	//	- signup
	//	- signup-id
	//  - signup-password
	Prompt CustomPromptType `json:"-"`
}

// CustomPromptManager manages Auth0 CustomPrompt resources.
type CustomPromptManager manager

type CustomPromptType string

const (
	CustomPromptSignup         CustomPromptType = "signup"
	CustomPromptSignupID       CustomPromptType = "signup-id"
	CustomPromptSignupPassword CustomPromptType = "signup-password"
)

var validPrompts = []CustomPromptType{CustomPromptSignup, CustomPromptSignupID, CustomPromptSignupPassword}

// Create a new custom prompt partial.
func (m *CustomPromptManager) Create(ctx context.Context, c *CustomPrompt, opts ...RequestOption) error {
	if !slices.Contains(validPrompts, c.Prompt) {
		return fmt.Errorf("invalid custom prompt: %s", c.Prompt)
	}
	return m.management.Request(ctx, "POST", m.management.URI("prompts", string(c.Prompt), "partials"), c, opts...)
}

// Update a custom prompt partial.
func (m *CustomPromptManager) Update(ctx context.Context, c *CustomPrompt, opts ...RequestOption) error {
	if !slices.Contains(validPrompts, c.Prompt) {
		return fmt.Errorf("invalid custom prompt: %s", c.Prompt)
	}
	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Prompt), "partials"), c, opts...)
}

// Read a custom prompt partial.
func (m *CustomPromptManager) Read(ctx context.Context, prompt CustomPromptType, opts ...RequestOption) (c *CustomPrompt, err error) {
	if !slices.Contains(validPrompts, CustomPromptType(prompt)) {
		return nil, fmt.Errorf("invalid custom prompt: %s", prompt)
	}
	err = m.management.Request(ctx, "GET", m.management.URI("prompts", string(prompt), "partials"), &c, opts...)
	return
}

// Delete a custom prompt partial.
func (m *CustomPromptManager) Delete(ctx context.Context, c *CustomPrompt, opts ...RequestOption) error {
	if !slices.Contains(validPrompts, c.Prompt) {
		return fmt.Errorf("invalid custom prompt: %s", c.Prompt)
	}
	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Prompt), "partials"), &CustomPrompt{}, opts...)
}
