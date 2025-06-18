package management

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

/*
Flow tests.
*/
func TestFlowManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	action := httpAction{
		ID:     "http_request_Id_123",
		Type:   "HTTP",
		Action: "SEND_REQUEST",
		Params: struct {
			URL    string `json:"url"`
			Method string `json:"method"`
		}{
			URL:    "https://api-endpoint.com/api/v1/resource",
			Method: "GET",
		},
	}

	flow := &Flow{
		Name:    auth0.String("test-flow"),
		Actions: []interface{}{action},
	}

	err := api.Flow.Create(context.Background(), flow)

	assert.NoError(t, err)
	assert.NotEmpty(t, flow.GetID())

	t.Cleanup(func() {
		cleanupFlow(t, flow.GetID())
	})
}

func TestFlowManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedFlow := givenAFlow(t)

	actualFlow, err := api.Flow.Read(context.Background(), expectedFlow.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedFlow, actualFlow)
}

func TestFlowManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedFlow := givenAFlow(t)
	updatedFlow := &Flow{
		Name: auth0.String("update-test-flow"),
	}
	err := api.Flow.Update(context.Background(), expectedFlow.GetID(), updatedFlow)

	assert.NoError(t, err)
	assert.Equal(t, "update-test-flow", updatedFlow.GetName())
}

func TestFlowManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedFlow := givenAFlow(t)

	err := api.Flow.Delete(context.Background(), expectedFlow.GetID())
	assert.NoError(t, err)
}

func TestFlowManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	flow := givenAFlow(t)
	flow.Actions = nil

	flowList, err := api.Flow.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(flowList.Flows), 0)
	assert.Contains(t, flowList.Flows, flow)
}

func TestFlowManager_MarshalJSON(t *testing.T) {
	for flow, expected := range map[*Flow]string{
		{}: `{}`,
		{
			Name: auth0.String("test-flow"),
		}: `{"name":"test-flow"}`,
		{
			ID:        auth0.String("some-id"),
			CreatedAt: auth0.Time(time.Now()),
			UpdatedAt: auth0.Time(time.Now()),
		}: `{}`,
	} {
		payload, err := json.Marshal(flow)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

type httpAction struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Action string `json:"action"`
	Params struct {
		URL    string `json:"url"`
		Method string `json:"method"`
	} `json:"params"`
}

func givenAFlow(t *testing.T) *Flow {
	t.Helper()

	flow := &Flow{
		Name: auth0.String("test-flow"),
	}

	err := api.Flow.Create(context.Background(), flow)
	assert.NoError(t, err)
	t.Cleanup(func() {
		cleanupFlow(t, flow.GetID())
	})

	return flow
}

func cleanupFlow(t *testing.T, flowID string) {
	t.Helper()

	err := api.Flow.Delete(context.Background(), flowID)
	if err != nil {
		var managementErr Error
		ok := errors.As(err, &managementErr)
		// We don't want to fail the test if the resource is already deleted.
		// clean up, therefore we only raise non-404 errors.
		// If `err` doesn't cast to management.Error, we raise it immediately.
		if !ok || managementErr.Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}

/*
Flow Vault Connection tests.
*/
func TestFlowVaultConnectionManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	flowVaultConnection := &FlowVaultConnection{
		AppID: auth0.String("HTTP"),
		Name:  auth0.String("test-vault-connection"),
		Setup: &map[string]interface{}{
			"token": "my-token",
			"type":  "BEARER",
		},
	}

	err := api.Flow.Vault.CreateConnection(context.Background(), flowVaultConnection)

	assert.NoError(t, err)
	assert.NotEmpty(t, flowVaultConnection.GetID())

	t.Cleanup(func() {
		cleanupFlowVaultConnection(t, flowVaultConnection.GetID())
	})
}

func TestFlowVaultConnectionManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedFlowVaultConnection := givenAFlowVaultConnection(t)

	actualFlowVaultConnection, err := api.Flow.Vault.GetConnection(context.Background(), expectedFlowVaultConnection.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedFlowVaultConnection, actualFlowVaultConnection)
}

func TestFlowVaultConnectionManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedFlowVaultConnection := givenAFlowVaultConnection(t)
	updatedFlowVaultConnection := &FlowVaultConnection{
		Name: auth0.String("new-connection-name"),
	}

	err := api.Flow.Vault.UpdateConnection(context.Background(), expectedFlowVaultConnection.GetID(), updatedFlowVaultConnection)

	assert.NoError(t, err)
	assert.Equal(t, "new-connection-name", updatedFlowVaultConnection.GetName())
	assert.Equal(t, expectedFlowVaultConnection.GetAppID(), updatedFlowVaultConnection.GetAppID())
}

func TestFlowVaultConnectionManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedFlowVaultConnection := givenAFlowVaultConnection(t)

	err := api.Flow.Vault.DeleteConnection(context.Background(), expectedFlowVaultConnection.GetID())
	assert.NoError(t, err)
}

func TestFlowVaultConnectionManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	flowVaultConnection := givenAFlowVaultConnection(t)

	flowVaultConnectionList, err := api.Flow.Vault.GetConnectionList(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(flowVaultConnectionList.Connections), 0)
	assert.Contains(t, flowVaultConnectionList.Connections, flowVaultConnection)
}

func TestFlowVaultConnectionManager_MarshalJSON(t *testing.T) {
	for connection, expected := range map[*FlowVaultConnection]string{
		{}: `{}`,
		{
			AppID: auth0.String("HTTP"),
			Name:  auth0.String("test-flow-vault-connection"),
			Setup: &map[string]interface{}{
				"key": "value",
			},
		}: `{"app_id":"HTTP","name":"test-flow-vault-connection","setup":{"key":"value"}}`,
		{
			AppID: auth0.String("AUTH0"),
			Name:  auth0.String("auth0-flow-vault-connection"),
		}: `{"app_id":"AUTH0","name":"auth0-flow-vault-connection"}`,
		{
			ID:        auth0.String("some-id"),
			CreatedAt: auth0.Time(time.Now()),
			UpdatedAt: auth0.Time(time.Now()),
		}: `{}`,
	} {
		payload, err := json.Marshal(connection)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func givenAFlowVaultConnection(t *testing.T) *FlowVaultConnection {
	flowVaultConnection := &FlowVaultConnection{
		AppID: auth0.String("HTTP"),
		Name:  auth0.String("test-vault-connection"),
	}

	err := api.Flow.Vault.CreateConnection(context.Background(), flowVaultConnection)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		cleanupFlowVaultConnection(t, flowVaultConnection.GetID())
	})

	return flowVaultConnection
}

func cleanupFlowVaultConnection(t *testing.T, id string) {
	t.Helper()

	err := api.Flow.Vault.DeleteConnection(context.Background(), id)
	if err != nil {
		var managementErr Error
		ok := errors.As(err, &managementErr)
		// We don't want to fail the test if the resource is already deleted.
		// clean up, therefore we only raise non-404 errors.
		// If `err` doesn't cast to management.Error, we raise it immediately.
		if !ok || managementErr.Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
