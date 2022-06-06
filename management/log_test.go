package management

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const successfulAPIOperation = "sapi"

func TestLogManager(t *testing.T) {
	setupHTTPRecordings(t)

	// Limit results to 5 entries, starting from the first page.
	logs, err := m.Log.List(Page(1), PerPage(5))
	assert.NoError(t, err)
	assert.Greater(t, len(logs), 0, "can't seem to find any logs, have you ran any tests before?")

	actualLog, err := m.Log.Read(logs[0].GetID())
	assert.NoError(t, err)
	assert.Equal(t, logs[0], actualLog)

	// Search by type "Success API Operation" and limit results to 5 entries.
	logs, err = m.Log.List(
		Parameter("q", fmt.Sprintf(`type:%q`, successfulAPIOperation)),
		PerPage(5),
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, logs)
}
