package management

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

const successfulAPIOperation = "sapi"

func TestLogManager(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Limit results to 5 entries, starting from the first page.
	logs, err := api.Log.List(context.Background(), Page(1), PerPage(5))
	assert.NoError(t, err)
	assert.Greater(t, len(logs), 0, "can't seem to find any logs, have you ran any tests before?")

	actualLog, err := api.Log.Read(context.Background(), logs[0].GetID())
	assert.NoError(t, err)
	assert.Equal(t, logs[0], actualLog)

	// Search by type "Success API Operation" and limit results to 5 entries.
	logs, err = api.Log.List(
		context.Background(),
		Parameter("q", fmt.Sprintf(`type:%q`, successfulAPIOperation)),
		PerPage(5),
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, logs)
}

func TestLogManager_CheckpointPagination(t *testing.T) {
	configureHTTPTestRecordings(t)

	logList, err := api.Log.List(context.Background(), Page(1), PerPage(10))
	assert.NoError(t, err)
	assert.Greater(t, len(logList), 0, "can't seem to find any logs, have you ran any tests before?")

	from := logList[9].GetID()

	// Take the first 2 entries from the 10th log
	logs, err := api.Log.List(context.Background(), Take(2), From(from))
	assert.NoError(t, err)
	assert.Len(t, logs, 2)

	log1 := logList[8]
	log2 := logList[7]
	assert.Equal(t, log1.GetID(), logs[0].GetID())
	assert.Equal(t, log2.GetID(), logs[1].GetID())
}

func TestLogManagerScope_UnmarshalJSON(t *testing.T) {
	for expectedAsString, expected := range map[string]*Log{
		`{}`: {},
		`{"scope": "openid profile email"}`: {
			Scope: auth0.String("openid profile email"),
		},
		`{"scope": ["openid", "profile", "email"]}`: {
			Scope: auth0.String("openid profile email"),
		},
		`{"scope": []}`: {
			Scope: auth0.String(""),
		},
	} {
		var actual *Log
		err := json.Unmarshal([]byte(expectedAsString), &actual)
		assert.NoError(t, err)
		assert.Equal(t, actual, expected)
	}

	t.Run("Throws an unexpected type error", func(t *testing.T) {
		var actual *Log
		err := json.Unmarshal([]byte(`{"scope": 1}`), &actual)
		assert.EqualError(t, err, "unexpected type for field scope: float64")
	})
}
