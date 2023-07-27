package management

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestRuleConfigManager_Upsert(t *testing.T) {
	configureHTTPTestRecordings(t)

	key := "foo"
	ruleConfig := &RuleConfig{Value: auth0.String("bar")}

	err := api.RuleConfig.Upsert(context.Background(), key, ruleConfig)
	assert.NoError(t, err)
	assert.Equal(t, key, ruleConfig.GetKey())

	t.Cleanup(func() {
		cleanupRuleConfig(t, ruleConfig.GetKey())
	})
}

func TestRuleConfigManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expected := givenARuleConfig(t)

	actual, err := api.RuleConfig.Read(context.Background(), expected.GetKey())

	assert.NoError(t, err)
	assert.Equal(t, expected.GetKey(), actual.GetKey())
}

func TestRuleConfigManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	ruleConfig := givenARuleConfig(t)

	err := api.RuleConfig.Delete(context.Background(), ruleConfig.GetKey())
	assert.NoError(t, err)

	actualRuleConfig, err := api.RuleConfig.Read(context.Background(), ruleConfig.GetKey())
	assert.Empty(t, actualRuleConfig)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestRuleConfigManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	ruleConfig := givenARuleConfig(t)

	ruleConfigs, err := api.RuleConfig.List(context.Background())

	assert.NoError(t, err)
	assert.Len(t, ruleConfigs, 1)
	assert.Equal(t, ruleConfig.GetKey(), ruleConfigs[0].GetKey())
}

func givenARuleConfig(t *testing.T) *RuleConfig {
	t.Helper()

	key := "foo"
	ruleConfig := &RuleConfig{Value: auth0.String("bar")}

	err := api.RuleConfig.Upsert(context.Background(), key, ruleConfig)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupRuleConfig(t, ruleConfig.GetKey())
	})

	return ruleConfig
}

func cleanupRuleConfig(t *testing.T, key string) {
	t.Helper()

	err := api.RuleConfig.Delete(context.Background(), key)
	require.NoError(t, err)
}
