package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlacklist(t *testing.T) {
	setupHTTPRecordings(t)

	client := givenAClient(t)

	blackListToken := &BlacklistToken{
		Audience: client.GetClientID(),
		JTI:      "test",
	}

	err := m.Blacklist.Create(blackListToken)
	assert.NoError(t, err)

	blackList, err := m.Blacklist.List(Parameter("aud", client.GetClientID()))
	assert.NoError(t, err)
	assert.Len(t, blackList, 1)
	assert.Equal(t, client.GetClientID(), blackList[0].Audience)
}
