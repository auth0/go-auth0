package management

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestRuleConfigManager_Upsert(t *testing.T) {
	setupHTTPRecordings(t)

	key := "foo"
	ruleConfig := &RuleConfig{Value: auth0.String("bar")}

	err := m.RuleConfig.Upsert(key, ruleConfig)
	assert.NoError(t, err)
	assert.Equal(t, key, ruleConfig.GetKey())

	t.Cleanup(func() {
		cleanupRuleConfig(t, ruleConfig.GetKey())
	})
}

func TestRuleConfigManager_Read(t *testing.T) {
	setupHTTPRecordings(t)

	expected := givenARuleConfig(t)

	actual, err := m.RuleConfig.Read(expected.GetKey())

	assert.NoError(t, err)
	assert.Equal(t, expected.GetKey(), actual.GetKey())
}

func TestRuleConfigManager_Delete(t *testing.T) {
	setupHTTPRecordings(t)

	ruleConfig := givenARuleConfig(t)

	err := m.RuleConfig.Delete(ruleConfig.GetKey())
	assert.NoError(t, err)

	actualRuleConfig, err := m.RuleConfig.Read(ruleConfig.GetKey())
	assert.Empty(t, actualRuleConfig)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestRuleConfigManager_List(t *testing.T) {
	setupHTTPRecordings(t)

	ruleConfig := givenARuleConfig(t)

	ruleConfigs, err := m.RuleConfig.List()

	assert.NoError(t, err)
	assert.Len(t, ruleConfigs, 1)
	assert.Equal(t, ruleConfig.GetKey(), ruleConfigs[0].GetKey())
}

func givenARuleConfig(t *testing.T) *RuleConfig {
	t.Helper()

	key := "foo"
	ruleConfig := &RuleConfig{Value: auth0.String("bar")}

	err := m.RuleConfig.Upsert(key, ruleConfig)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupRuleConfig(t, ruleConfig.GetKey())
	})

	return ruleConfig
}

func cleanupRuleConfig(t *testing.T, key string) {
	t.Helper()

	err := m.RuleConfig.Delete(key)
	require.NoError(t, err)
}
