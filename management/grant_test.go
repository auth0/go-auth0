package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrantManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	grantList, err := api.Grant.List()
	assert.NoError(t, err)
	assert.IsType(t, &GrantList{}, grantList)
	assert.NotNil(t, grantList.Grants)
}
