package management

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const successfulAPIOperation = "sapi"

func TestLogManager(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Limit results to 5 entries, starting from the first page.
	logs, err := api.Log.List(Page(1), PerPage(5))
	assert.NoError(t, err)
	assert.Greater(t, len(logs), 0, "can't seem to find any logs, have you ran any tests before?")

	actualLog, err := api.Log.Read(logs[0].GetID())
	assert.NoError(t, err)
	assert.Equal(t, logs[0], actualLog)

	// Search by type "Success API Operation" and limit results to 5 entries.
	logs, err = api.Log.List(
		Parameter("q", fmt.Sprintf(`type:%q`, successfulAPIOperation)),
		PerPage(5),
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, logs)
}

func TestLogManager_CheckpointPagination(t *testing.T) {
	configureHTTPTestRecordings(t)

	logList, err := api.Log.List(Page(1), PerPage(5))
	assert.NoError(t, err)
	assert.Greater(t, len(logList), 0, "can't seem to find any logs, have you ran any tests before?")

	from := logList[4].GetID()

	// Take the first 2 entries from the 5th log
	logs, err := api.Log.List(Take(2), From(from))
	assert.NoError(t, err)
	assert.Len(t, logs, 2)

	log1 := logList[3]
	log2 := logList[2]
	assert.Equal(t, log1.GetID(), logs[0].GetID())
	assert.Equal(t, log2.GetID(), logs[1].GetID())
}
