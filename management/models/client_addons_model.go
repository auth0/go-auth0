/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"bytes"
	"encoding/json"
)

// ClientAddons Addons enabled for this client and their associated configurations.
type ClientAddons struct {
	Aws       ClientAddonsAws       `json:"aws"`
	AzureBlob ClientAddonsAzureBlob `json:"azure_blob"`
	AzureSb   ClientAddonsAzureSb   `json:"azure_sb"`
	Rms       ClientAddonsRms       `json:"rms"`
	Mscrm     ClientAddonsMscrm     `json:"mscrm"`
	Slack     ClientAddonsSlack     `json:"slack"`
	Sentry    ClientAddonsSentry    `json:"sentry"`
	// Box SSO indicator (no configuration settings needed for Box SSO).
	Box map[string]interface{} `json:"box"`
	// CloudBees SSO indicator (no configuration settings needed for CloudBees SSO).
	Cloudbees map[string]interface{} `json:"cloudbees"`
	// Concur SSO indicator (no configuration settings needed for Concur SSO).
	Concur map[string]interface{} `json:"concur"`
	// Dropbox SSO indicator (no configuration settings needed for Dropbox SSO).
	Dropbox              map[string]interface{}           `json:"dropbox"`
	Echosign             ClientAddonsEchosign             `json:"echosign"`
	Egnyte               ClientAddonsEgnyte               `json:"egnyte"`
	Firebase             ClientAddonsFirebase             `json:"firebase"`
	Newrelic             ClientAddonsNewrelic             `json:"newrelic"`
	Office365            ClientAddonsOffice365            `json:"office365"`
	Salesforce           ClientAddonsSalesforce           `json:"salesforce"`
	SalesforceApi        ClientAddonsSalesforceApi        `json:"salesforce_api"`
	SalesforceSandboxApi ClientAddonsSalesforceSandboxApi `json:"salesforce_sandbox_api"`
	Samlp                ClientAddonsSamlp                `json:"samlp"`
	Layer                ClientAddonsLayer                `json:"layer"`
	SapApi               ClientAddonsSapApi               `json:"sap_api"`
	Sharepoint           ClientAddonsSharepoint           `json:"sharepoint"`
	Springcm             ClientAddonsSpringcm             `json:"springcm"`
	Wams                 ClientAddonsWams                 `json:"wams"`
	// WS-Fed (WIF) addon indicator. Actual configuration is stored in `callback` and `client_aliases` properties on the client.
	Wsfed          map[string]interface{}     `json:"wsfed"`
	Zendesk        ClientAddonsZendesk        `json:"zendesk"`
	Zoom           ClientAddonsZoom           `json:"zoom"`
	SsoIntegration ClientAddonsSsoIntegration `json:"sso_integration"`
	// Okta Access Gateway SSO configuration
	Oag map[string]interface{} `json:"oag"`
}

type _ClientAddons ClientAddons

// GetAws returns the Aws field value
func (o *ClientAddons) GetAws() ClientAddonsAws {
	if o == nil {
		var ret ClientAddonsAws
		return ret
	}

	return o.Aws
}

// GetAwsOk returns a tuple with the Aws field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetAwsOk() (*ClientAddonsAws, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aws, true
}

// SetAws sets field value
func (o *ClientAddons) SetAws(v ClientAddonsAws) {
	o.Aws = v
}

// GetAzureBlob returns the AzureBlob field value
func (o *ClientAddons) GetAzureBlob() ClientAddonsAzureBlob {
	if o == nil {
		var ret ClientAddonsAzureBlob
		return ret
	}

	return o.AzureBlob
}

// GetAzureBlobOk returns a tuple with the AzureBlob field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetAzureBlobOk() (*ClientAddonsAzureBlob, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureBlob, true
}

// SetAzureBlob sets field value
func (o *ClientAddons) SetAzureBlob(v ClientAddonsAzureBlob) {
	o.AzureBlob = v
}

// GetAzureSb returns the AzureSb field value
func (o *ClientAddons) GetAzureSb() ClientAddonsAzureSb {
	if o == nil {
		var ret ClientAddonsAzureSb
		return ret
	}

	return o.AzureSb
}

// GetAzureSbOk returns a tuple with the AzureSb field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetAzureSbOk() (*ClientAddonsAzureSb, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureSb, true
}

// SetAzureSb sets field value
func (o *ClientAddons) SetAzureSb(v ClientAddonsAzureSb) {
	o.AzureSb = v
}

// GetRms returns the Rms field value
func (o *ClientAddons) GetRms() ClientAddonsRms {
	if o == nil {
		var ret ClientAddonsRms
		return ret
	}

	return o.Rms
}

// GetRmsOk returns a tuple with the Rms field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetRmsOk() (*ClientAddonsRms, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rms, true
}

// SetRms sets field value
func (o *ClientAddons) SetRms(v ClientAddonsRms) {
	o.Rms = v
}

// GetMscrm returns the Mscrm field value
func (o *ClientAddons) GetMscrm() ClientAddonsMscrm {
	if o == nil {
		var ret ClientAddonsMscrm
		return ret
	}

	return o.Mscrm
}

// GetMscrmOk returns a tuple with the Mscrm field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetMscrmOk() (*ClientAddonsMscrm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mscrm, true
}

// SetMscrm sets field value
func (o *ClientAddons) SetMscrm(v ClientAddonsMscrm) {
	o.Mscrm = v
}

// GetSlack returns the Slack field value
func (o *ClientAddons) GetSlack() ClientAddonsSlack {
	if o == nil {
		var ret ClientAddonsSlack
		return ret
	}

	return o.Slack
}

// GetSlackOk returns a tuple with the Slack field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSlackOk() (*ClientAddonsSlack, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slack, true
}

// SetSlack sets field value
func (o *ClientAddons) SetSlack(v ClientAddonsSlack) {
	o.Slack = v
}

// GetSentry returns the Sentry field value
func (o *ClientAddons) GetSentry() ClientAddonsSentry {
	if o == nil {
		var ret ClientAddonsSentry
		return ret
	}

	return o.Sentry
}

// GetSentryOk returns a tuple with the Sentry field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSentryOk() (*ClientAddonsSentry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sentry, true
}

// SetSentry sets field value
func (o *ClientAddons) SetSentry(v ClientAddonsSentry) {
	o.Sentry = v
}

// GetBox returns the Box field value
func (o *ClientAddons) GetBox() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Box
}

// GetBoxOk returns a tuple with the Box field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetBoxOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Box, true
}

// SetBox sets field value
func (o *ClientAddons) SetBox(v map[string]interface{}) {
	o.Box = v
}

// GetCloudbees returns the Cloudbees field value
func (o *ClientAddons) GetCloudbees() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Cloudbees
}

// GetCloudbeesOk returns a tuple with the Cloudbees field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetCloudbeesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Cloudbees, true
}

// SetCloudbees sets field value
func (o *ClientAddons) SetCloudbees(v map[string]interface{}) {
	o.Cloudbees = v
}

// GetConcur returns the Concur field value
func (o *ClientAddons) GetConcur() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Concur
}

// GetConcurOk returns a tuple with the Concur field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetConcurOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Concur, true
}

// SetConcur sets field value
func (o *ClientAddons) SetConcur(v map[string]interface{}) {
	o.Concur = v
}

// GetDropbox returns the Dropbox field value
func (o *ClientAddons) GetDropbox() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Dropbox
}

// GetDropboxOk returns a tuple with the Dropbox field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetDropboxOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Dropbox, true
}

// SetDropbox sets field value
func (o *ClientAddons) SetDropbox(v map[string]interface{}) {
	o.Dropbox = v
}

// GetEchosign returns the Echosign field value
func (o *ClientAddons) GetEchosign() ClientAddonsEchosign {
	if o == nil {
		var ret ClientAddonsEchosign
		return ret
	}

	return o.Echosign
}

// GetEchosignOk returns a tuple with the Echosign field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetEchosignOk() (*ClientAddonsEchosign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Echosign, true
}

// SetEchosign sets field value
func (o *ClientAddons) SetEchosign(v ClientAddonsEchosign) {
	o.Echosign = v
}

// GetEgnyte returns the Egnyte field value
func (o *ClientAddons) GetEgnyte() ClientAddonsEgnyte {
	if o == nil {
		var ret ClientAddonsEgnyte
		return ret
	}

	return o.Egnyte
}

// GetEgnyteOk returns a tuple with the Egnyte field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetEgnyteOk() (*ClientAddonsEgnyte, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Egnyte, true
}

// SetEgnyte sets field value
func (o *ClientAddons) SetEgnyte(v ClientAddonsEgnyte) {
	o.Egnyte = v
}

// GetFirebase returns the Firebase field value
func (o *ClientAddons) GetFirebase() ClientAddonsFirebase {
	if o == nil {
		var ret ClientAddonsFirebase
		return ret
	}

	return o.Firebase
}

// GetFirebaseOk returns a tuple with the Firebase field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetFirebaseOk() (*ClientAddonsFirebase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Firebase, true
}

// SetFirebase sets field value
func (o *ClientAddons) SetFirebase(v ClientAddonsFirebase) {
	o.Firebase = v
}

// GetNewrelic returns the Newrelic field value
func (o *ClientAddons) GetNewrelic() ClientAddonsNewrelic {
	if o == nil {
		var ret ClientAddonsNewrelic
		return ret
	}

	return o.Newrelic
}

// GetNewrelicOk returns a tuple with the Newrelic field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetNewrelicOk() (*ClientAddonsNewrelic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Newrelic, true
}

// SetNewrelic sets field value
func (o *ClientAddons) SetNewrelic(v ClientAddonsNewrelic) {
	o.Newrelic = v
}

// GetOffice365 returns the Office365 field value
func (o *ClientAddons) GetOffice365() ClientAddonsOffice365 {
	if o == nil {
		var ret ClientAddonsOffice365
		return ret
	}

	return o.Office365
}

// GetOffice365Ok returns a tuple with the Office365 field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetOffice365Ok() (*ClientAddonsOffice365, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Office365, true
}

// SetOffice365 sets field value
func (o *ClientAddons) SetOffice365(v ClientAddonsOffice365) {
	o.Office365 = v
}

// GetSalesforce returns the Salesforce field value
func (o *ClientAddons) GetSalesforce() ClientAddonsSalesforce {
	if o == nil {
		var ret ClientAddonsSalesforce
		return ret
	}

	return o.Salesforce
}

// GetSalesforceOk returns a tuple with the Salesforce field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSalesforceOk() (*ClientAddonsSalesforce, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Salesforce, true
}

// SetSalesforce sets field value
func (o *ClientAddons) SetSalesforce(v ClientAddonsSalesforce) {
	o.Salesforce = v
}

// GetSalesforceApi returns the SalesforceApi field value
func (o *ClientAddons) GetSalesforceApi() ClientAddonsSalesforceApi {
	if o == nil {
		var ret ClientAddonsSalesforceApi
		return ret
	}

	return o.SalesforceApi
}

// GetSalesforceApiOk returns a tuple with the SalesforceApi field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSalesforceApiOk() (*ClientAddonsSalesforceApi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesforceApi, true
}

// SetSalesforceApi sets field value
func (o *ClientAddons) SetSalesforceApi(v ClientAddonsSalesforceApi) {
	o.SalesforceApi = v
}

// GetSalesforceSandboxApi returns the SalesforceSandboxApi field value
func (o *ClientAddons) GetSalesforceSandboxApi() ClientAddonsSalesforceSandboxApi {
	if o == nil {
		var ret ClientAddonsSalesforceSandboxApi
		return ret
	}

	return o.SalesforceSandboxApi
}

// GetSalesforceSandboxApiOk returns a tuple with the SalesforceSandboxApi field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSalesforceSandboxApiOk() (*ClientAddonsSalesforceSandboxApi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesforceSandboxApi, true
}

// SetSalesforceSandboxApi sets field value
func (o *ClientAddons) SetSalesforceSandboxApi(v ClientAddonsSalesforceSandboxApi) {
	o.SalesforceSandboxApi = v
}

// GetSamlp returns the Samlp field value
func (o *ClientAddons) GetSamlp() ClientAddonsSamlp {
	if o == nil {
		var ret ClientAddonsSamlp
		return ret
	}

	return o.Samlp
}

// GetSamlpOk returns a tuple with the Samlp field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSamlpOk() (*ClientAddonsSamlp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Samlp, true
}

// SetSamlp sets field value
func (o *ClientAddons) SetSamlp(v ClientAddonsSamlp) {
	o.Samlp = v
}

// GetLayer returns the Layer field value
func (o *ClientAddons) GetLayer() ClientAddonsLayer {
	if o == nil {
		var ret ClientAddonsLayer
		return ret
	}

	return o.Layer
}

// GetLayerOk returns a tuple with the Layer field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetLayerOk() (*ClientAddonsLayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layer, true
}

// SetLayer sets field value
func (o *ClientAddons) SetLayer(v ClientAddonsLayer) {
	o.Layer = v
}

// GetSapApi returns the SapApi field value
func (o *ClientAddons) GetSapApi() ClientAddonsSapApi {
	if o == nil {
		var ret ClientAddonsSapApi
		return ret
	}

	return o.SapApi
}

// GetSapApiOk returns a tuple with the SapApi field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSapApiOk() (*ClientAddonsSapApi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SapApi, true
}

// SetSapApi sets field value
func (o *ClientAddons) SetSapApi(v ClientAddonsSapApi) {
	o.SapApi = v
}

// GetSharepoint returns the Sharepoint field value
func (o *ClientAddons) GetSharepoint() ClientAddonsSharepoint {
	if o == nil {
		var ret ClientAddonsSharepoint
		return ret
	}

	return o.Sharepoint
}

// GetSharepointOk returns a tuple with the Sharepoint field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSharepointOk() (*ClientAddonsSharepoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sharepoint, true
}

// SetSharepoint sets field value
func (o *ClientAddons) SetSharepoint(v ClientAddonsSharepoint) {
	o.Sharepoint = v
}

// GetSpringcm returns the Springcm field value
func (o *ClientAddons) GetSpringcm() ClientAddonsSpringcm {
	if o == nil {
		var ret ClientAddonsSpringcm
		return ret
	}

	return o.Springcm
}

// GetSpringcmOk returns a tuple with the Springcm field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSpringcmOk() (*ClientAddonsSpringcm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Springcm, true
}

// SetSpringcm sets field value
func (o *ClientAddons) SetSpringcm(v ClientAddonsSpringcm) {
	o.Springcm = v
}

// GetWams returns the Wams field value
func (o *ClientAddons) GetWams() ClientAddonsWams {
	if o == nil {
		var ret ClientAddonsWams
		return ret
	}

	return o.Wams
}

// GetWamsOk returns a tuple with the Wams field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetWamsOk() (*ClientAddonsWams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wams, true
}

// SetWams sets field value
func (o *ClientAddons) SetWams(v ClientAddonsWams) {
	o.Wams = v
}

// GetWsfed returns the Wsfed field value
func (o *ClientAddons) GetWsfed() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Wsfed
}

// GetWsfedOk returns a tuple with the Wsfed field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetWsfedOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Wsfed, true
}

// SetWsfed sets field value
func (o *ClientAddons) SetWsfed(v map[string]interface{}) {
	o.Wsfed = v
}

// GetZendesk returns the Zendesk field value
func (o *ClientAddons) GetZendesk() ClientAddonsZendesk {
	if o == nil {
		var ret ClientAddonsZendesk
		return ret
	}

	return o.Zendesk
}

// GetZendeskOk returns a tuple with the Zendesk field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetZendeskOk() (*ClientAddonsZendesk, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zendesk, true
}

// SetZendesk sets field value
func (o *ClientAddons) SetZendesk(v ClientAddonsZendesk) {
	o.Zendesk = v
}

// GetZoom returns the Zoom field value
func (o *ClientAddons) GetZoom() ClientAddonsZoom {
	if o == nil {
		var ret ClientAddonsZoom
		return ret
	}

	return o.Zoom
}

// GetZoomOk returns a tuple with the Zoom field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetZoomOk() (*ClientAddonsZoom, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zoom, true
}

// SetZoom sets field value
func (o *ClientAddons) SetZoom(v ClientAddonsZoom) {
	o.Zoom = v
}

// GetSsoIntegration returns the SsoIntegration field value
func (o *ClientAddons) GetSsoIntegration() ClientAddonsSsoIntegration {
	if o == nil {
		var ret ClientAddonsSsoIntegration
		return ret
	}

	return o.SsoIntegration
}

// GetSsoIntegrationOk returns a tuple with the SsoIntegration field value
// and a boolean to check if the value has been set.
func (o *ClientAddons) GetSsoIntegrationOk() (*ClientAddonsSsoIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsoIntegration, true
}

// SetSsoIntegration sets field value
func (o *ClientAddons) SetSsoIntegration(v ClientAddonsSsoIntegration) {
	o.SsoIntegration = v
}

// GetOag returns the Oag field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ClientAddons) GetOag() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Oag
}

// GetOagOk returns a tuple with the Oag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientAddons) GetOagOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Oag) {
		return map[string]interface{}{}, false
	}
	return o.Oag, true
}

// SetOag sets field value
func (o *ClientAddons) SetOag(v map[string]interface{}) {
	o.Oag = v
}

func (o ClientAddons) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAddons) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws"] = o.Aws
	toSerialize["azure_blob"] = o.AzureBlob
	toSerialize["azure_sb"] = o.AzureSb
	toSerialize["rms"] = o.Rms
	toSerialize["mscrm"] = o.Mscrm
	toSerialize["slack"] = o.Slack
	toSerialize["sentry"] = o.Sentry
	toSerialize["box"] = o.Box
	toSerialize["cloudbees"] = o.Cloudbees
	toSerialize["concur"] = o.Concur
	toSerialize["dropbox"] = o.Dropbox
	toSerialize["echosign"] = o.Echosign
	toSerialize["egnyte"] = o.Egnyte
	toSerialize["firebase"] = o.Firebase
	toSerialize["newrelic"] = o.Newrelic
	toSerialize["office365"] = o.Office365
	toSerialize["salesforce"] = o.Salesforce
	toSerialize["salesforce_api"] = o.SalesforceApi
	toSerialize["salesforce_sandbox_api"] = o.SalesforceSandboxApi
	toSerialize["samlp"] = o.Samlp
	toSerialize["layer"] = o.Layer
	toSerialize["sap_api"] = o.SapApi
	toSerialize["sharepoint"] = o.Sharepoint
	toSerialize["springcm"] = o.Springcm
	toSerialize["wams"] = o.Wams
	toSerialize["wsfed"] = o.Wsfed
	toSerialize["zendesk"] = o.Zendesk
	toSerialize["zoom"] = o.Zoom
	toSerialize["sso_integration"] = o.SsoIntegration
	if o.Oag != nil {
		toSerialize["oag"] = o.Oag
	}
	return toSerialize, nil
}

func (o *ClientAddons) UnmarshalJSON(data []byte) (err error) {
	varClientAddons := _ClientAddons{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClientAddons)

	if err != nil {
		return err
	}

	*o = ClientAddons(varClientAddons)

	return err
}

type NullableClientAddons struct {
	value *ClientAddons
	isSet bool
}

func (v NullableClientAddons) Get() *ClientAddons {
	return v.value
}

func (v *NullableClientAddons) Set(val *ClientAddons) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAddons) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAddons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAddons(val *ClientAddons) *NullableClientAddons {
	return &NullableClientAddons{value: val, isSet: true}
}

func (v NullableClientAddons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAddons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
