package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatManager_ActiveUsers(t *testing.T) {
	configureHTTPTestRecordings(t)

	activeUsers, err := api.Stat.ActiveUsers(context.Background())
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, activeUsers, 0)
}

func TestStatManager_Daily(t *testing.T) {
	configureHTTPTestRecordings(t)

	daily, err := api.Stat.Daily(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, daily)
}
