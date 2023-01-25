package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatManager_ActiveUsers(t *testing.T) {
	setupHTTPRecordings(t)

	activeUsers, err := api.Stat.ActiveUsers()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, activeUsers, 0)
}

func TestStatManager_Daily(t *testing.T) {
	setupHTTPRecordings(t)

	daily, err := api.Stat.Daily()
	assert.NoError(t, err)
	assert.NotEmpty(t, daily)
}
