package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrantManager_List(t *testing.T) {
	setupHTTPRecordings(t)

	grantList, err := m.Grant.List()
	assert.NoError(t, err)
	assert.IsType(t, &GrantList{}, grantList)
	assert.Len(t, grantList.Grants, 0)
}
