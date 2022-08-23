package management

import (
	"encoding/json"
	"fmt"
	"io"
)

// Error is an interface describing any error which
// could be returned by the Auth0 Management API.
type Error interface {
	// Status returns the status code returned by
	// the server together with the present error.
	Status() int
	error
}

type managementError struct {
	StatusCode int    `json:"statusCode"`
	Err        string `json:"error"`
	Message    string `json:"message"`
}

func newError(r io.Reader) error {
	m := &managementError{}
	if err := json.NewDecoder(r).Decode(m); err != nil {
		return err
	}

	return m
}

// Error formats the error into a string representation.
func (m *managementError) Error() string {
	return fmt.Sprintf("%d %s: %s", m.StatusCode, m.Err, m.Message)
}

// Status returns the status code of the error.
func (m *managementError) Status() int {
	return m.StatusCode
}
