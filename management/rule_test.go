package management

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestRuleManager_Create(t *testing.T) {
	setupHTTPRecordings(t)

	rule := &Rule{
		Name:    auth0.String("test-rule"),
		Script:  auth0.String("function (user, context, callback) { callback(null, user, context); }"),
		Enabled: auth0.Bool(false),
	}

	err := m.Rule.Create(rule)
	assert.NoError(t, err)
	assert.NotEmpty(t, rule.GetID())

	t.Cleanup(func() {
		cleanupRule(t, rule.GetID())
	})
}

func TestRuleManager_Read(t *testing.T) {
	setupHTTPRecordings(t)

	expectedRule := givenARule(t)

	actualRule, err := m.Rule.Read(expectedRule.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedRule, actualRule)
}

func TestRuleManager_Update(t *testing.T) {
	setupHTTPRecordings(t)

	rule := givenARule(t)
	updatedRule := &Rule{
		Order:   auth0.Int(5),
		Enabled: auth0.Bool(true),
	}

	err := m.Rule.Update(rule.GetID(), updatedRule)
	assert.NoError(t, err)

	actualRule, err := m.Rule.Read(rule.GetID())
	assert.NoError(t, err)
	assert.Equal(t, updatedRule.GetOrder(), actualRule.GetOrder())
	assert.Equal(t, updatedRule.GetEnabled(), actualRule.GetEnabled())
}

func TestRuleManager_Delete(t *testing.T) {
	setupHTTPRecordings(t)

	rule := givenARule(t)

	err := m.Rule.Delete(rule.GetID())
	assert.NoError(t, err)

	actualRule, err := m.Rule.Read(rule.GetID())
	assert.Empty(t, actualRule)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestRuleManager_List(t *testing.T) {
	setupHTTPRecordings(t)

	rule := givenARule(t)

	ruleList, err := m.Rule.List(IncludeFields("id"))

	assert.NoError(t, err)
	assert.Contains(t, ruleList.Rules, &Rule{ID: rule.ID})
}

func givenARule(t *testing.T) *Rule {
	t.Helper()

	rule := &Rule{
		Name:    auth0.String(fmt.Sprintf("test-rule%d", rand.Intn(999))),
		Script:  auth0.String("function (user, context, callback) { callback(null, user, context); }"),
		Enabled: auth0.Bool(false),
	}

	err := m.Rule.Create(rule)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupRule(t, rule.GetID())
	})

	return rule
}

func cleanupRule(t *testing.T, ruleID string) {
	t.Helper()

	err := m.Rule.Delete(ruleID)
	require.NoError(t, err)
}
