package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlacklist(t *testing.T) {
	configureHTTPTestRecordings(t)

	client := givenAClient(t)

	blackListToken := &BlacklistToken{
		Audience: client.GetClientID(),
		JTI:      "test",
	}

	err := api.Blacklist.Create(context.Background(), blackListToken)
	assert.NoError(t, err)

	blackList, err := api.Blacklist.List(context.Background(), Parameter("aud", client.GetClientID()))
	assert.NoError(t, err)
	assert.Len(t, blackList, 1)
	assert.Equal(t, client.GetClientID(), blackList[0].Audience)
}
