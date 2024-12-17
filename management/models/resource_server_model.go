/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// ResourceServer struct for ResourceServer
type ResourceServer struct {
	// ID of the API (resource server).
	Id *string `json:"id,omitempty"`
	// Friendly name for this resource server. Can not contain `<` or `>` characters.
	Name *string `json:"name,omitempty"`
	// Whether this is an Auth0 system API (true) or a custom API (false).
	IsSystem *bool `json:"is_system,omitempty"`
	// Unique identifier for the API used as the audience parameter on authorization calls. Can not be changed once set.
	Identifier *string `json:"identifier,omitempty"`
	// List of permissions (scopes) that this API uses.
	Scopes     []Scope                    `json:"scopes,omitempty"`
	SigningAlg *ClientJwtConfigurationAlg `json:"signing_alg,omitempty"`
	// Secret used to sign tokens when using symmetric algorithms (HS256).
	SigningSecret *string `json:"signing_secret,omitempty"`
	// Whether refresh tokens can be issued for this API (true) or not (false).
	AllowOfflineAccess *bool `json:"allow_offline_access,omitempty"`
	// Whether to skip user consent for applications flagged as first party (true) or not (false).
	SkipConsentForVerifiableFirstPartyClients *bool `json:"skip_consent_for_verifiable_first_party_clients,omitempty"`
	// Expiration value (in seconds) for access tokens issued for this API from the token endpoint.
	TokenLifetime *int32 `json:"token_lifetime,omitempty"`
	// Expiration value (in seconds) for access tokens issued for this API via Implicit or Hybrid Flows. Cannot be greater than the `token_lifetime` value.
	TokenLifetimeForWeb *int32 `json:"token_lifetime_for_web,omitempty"`
	// Whether authorization polices are enforced (true) or unenforced (false).
	EnforcePolicies      *bool                                   `json:"enforce_policies,omitempty"`
	TokenDialect         *ResourceServerTokenDialect             `json:"token_dialect,omitempty"`
	Client               map[string]interface{}                  `json:"client,omitempty"`
	TokenEncryption      NullableResourceServerTokenEncryption   `json:"token_encryption,omitempty"`
	ConsentPolicy        NullableResourceServerConsentPolicy     `json:"consent_policy,omitempty"`
	AuthorizationDetails []interface{}                           `json:"authorization_details,omitempty"`
	ProofOfPossession    NullableResourceServerProofOfPossession `json:"proof_of_possession,omitempty"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceServer) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceServer) SetName(v string) {
	o.Name = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *ResourceServer) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetIsSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *ResourceServer) HasIsSystem() bool {
	if o != nil && !IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *ResourceServer) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ResourceServer) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ResourceServer) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *ResourceServer) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ResourceServer) GetScopes() []Scope {
	if o == nil || IsNil(o.Scopes) {
		var ret []Scope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetScopesOk() ([]Scope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ResourceServer) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *ResourceServer) SetScopes(v []Scope) {
	o.Scopes = v
}

// GetSigningAlg returns the SigningAlg field value if set, zero value otherwise.
func (o *ResourceServer) GetSigningAlg() ClientJwtConfigurationAlg {
	if o == nil || IsNil(o.SigningAlg) {
		var ret ClientJwtConfigurationAlg
		return ret
	}
	return *o.SigningAlg
}

// GetSigningAlgOk returns a tuple with the SigningAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetSigningAlgOk() (*ClientJwtConfigurationAlg, bool) {
	if o == nil || IsNil(o.SigningAlg) {
		return nil, false
	}
	return o.SigningAlg, true
}

// HasSigningAlg returns a boolean if a field has been set.
func (o *ResourceServer) HasSigningAlg() bool {
	if o != nil && !IsNil(o.SigningAlg) {
		return true
	}

	return false
}

// SetSigningAlg gets a reference to the given ClientJwtConfigurationAlg and assigns it to the SigningAlg field.
func (o *ResourceServer) SetSigningAlg(v ClientJwtConfigurationAlg) {
	o.SigningAlg = &v
}

// GetSigningSecret returns the SigningSecret field value if set, zero value otherwise.
func (o *ResourceServer) GetSigningSecret() string {
	if o == nil || IsNil(o.SigningSecret) {
		var ret string
		return ret
	}
	return *o.SigningSecret
}

// GetSigningSecretOk returns a tuple with the SigningSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetSigningSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SigningSecret) {
		return nil, false
	}
	return o.SigningSecret, true
}

// HasSigningSecret returns a boolean if a field has been set.
func (o *ResourceServer) HasSigningSecret() bool {
	if o != nil && !IsNil(o.SigningSecret) {
		return true
	}

	return false
}

// SetSigningSecret gets a reference to the given string and assigns it to the SigningSecret field.
func (o *ResourceServer) SetSigningSecret(v string) {
	o.SigningSecret = &v
}

// GetAllowOfflineAccess returns the AllowOfflineAccess field value if set, zero value otherwise.
func (o *ResourceServer) GetAllowOfflineAccess() bool {
	if o == nil || IsNil(o.AllowOfflineAccess) {
		var ret bool
		return ret
	}
	return *o.AllowOfflineAccess
}

// GetAllowOfflineAccessOk returns a tuple with the AllowOfflineAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetAllowOfflineAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowOfflineAccess) {
		return nil, false
	}
	return o.AllowOfflineAccess, true
}

// HasAllowOfflineAccess returns a boolean if a field has been set.
func (o *ResourceServer) HasAllowOfflineAccess() bool {
	if o != nil && !IsNil(o.AllowOfflineAccess) {
		return true
	}

	return false
}

// SetAllowOfflineAccess gets a reference to the given bool and assigns it to the AllowOfflineAccess field.
func (o *ResourceServer) SetAllowOfflineAccess(v bool) {
	o.AllowOfflineAccess = &v
}

// GetSkipConsentForVerifiableFirstPartyClients returns the SkipConsentForVerifiableFirstPartyClients field value if set, zero value otherwise.
func (o *ResourceServer) GetSkipConsentForVerifiableFirstPartyClients() bool {
	if o == nil || IsNil(o.SkipConsentForVerifiableFirstPartyClients) {
		var ret bool
		return ret
	}
	return *o.SkipConsentForVerifiableFirstPartyClients
}

// GetSkipConsentForVerifiableFirstPartyClientsOk returns a tuple with the SkipConsentForVerifiableFirstPartyClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetSkipConsentForVerifiableFirstPartyClientsOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipConsentForVerifiableFirstPartyClients) {
		return nil, false
	}
	return o.SkipConsentForVerifiableFirstPartyClients, true
}

// HasSkipConsentForVerifiableFirstPartyClients returns a boolean if a field has been set.
func (o *ResourceServer) HasSkipConsentForVerifiableFirstPartyClients() bool {
	if o != nil && !IsNil(o.SkipConsentForVerifiableFirstPartyClients) {
		return true
	}

	return false
}

// SetSkipConsentForVerifiableFirstPartyClients gets a reference to the given bool and assigns it to the SkipConsentForVerifiableFirstPartyClients field.
func (o *ResourceServer) SetSkipConsentForVerifiableFirstPartyClients(v bool) {
	o.SkipConsentForVerifiableFirstPartyClients = &v
}

// GetTokenLifetime returns the TokenLifetime field value if set, zero value otherwise.
func (o *ResourceServer) GetTokenLifetime() int32 {
	if o == nil || IsNil(o.TokenLifetime) {
		var ret int32
		return ret
	}
	return *o.TokenLifetime
}

// GetTokenLifetimeOk returns a tuple with the TokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetTokenLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.TokenLifetime) {
		return nil, false
	}
	return o.TokenLifetime, true
}

// HasTokenLifetime returns a boolean if a field has been set.
func (o *ResourceServer) HasTokenLifetime() bool {
	if o != nil && !IsNil(o.TokenLifetime) {
		return true
	}

	return false
}

// SetTokenLifetime gets a reference to the given int32 and assigns it to the TokenLifetime field.
func (o *ResourceServer) SetTokenLifetime(v int32) {
	o.TokenLifetime = &v
}

// GetTokenLifetimeForWeb returns the TokenLifetimeForWeb field value if set, zero value otherwise.
func (o *ResourceServer) GetTokenLifetimeForWeb() int32 {
	if o == nil || IsNil(o.TokenLifetimeForWeb) {
		var ret int32
		return ret
	}
	return *o.TokenLifetimeForWeb
}

// GetTokenLifetimeForWebOk returns a tuple with the TokenLifetimeForWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetTokenLifetimeForWebOk() (*int32, bool) {
	if o == nil || IsNil(o.TokenLifetimeForWeb) {
		return nil, false
	}
	return o.TokenLifetimeForWeb, true
}

// HasTokenLifetimeForWeb returns a boolean if a field has been set.
func (o *ResourceServer) HasTokenLifetimeForWeb() bool {
	if o != nil && !IsNil(o.TokenLifetimeForWeb) {
		return true
	}

	return false
}

// SetTokenLifetimeForWeb gets a reference to the given int32 and assigns it to the TokenLifetimeForWeb field.
func (o *ResourceServer) SetTokenLifetimeForWeb(v int32) {
	o.TokenLifetimeForWeb = &v
}

// GetEnforcePolicies returns the EnforcePolicies field value if set, zero value otherwise.
func (o *ResourceServer) GetEnforcePolicies() bool {
	if o == nil || IsNil(o.EnforcePolicies) {
		var ret bool
		return ret
	}
	return *o.EnforcePolicies
}

// GetEnforcePoliciesOk returns a tuple with the EnforcePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetEnforcePoliciesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforcePolicies) {
		return nil, false
	}
	return o.EnforcePolicies, true
}

// HasEnforcePolicies returns a boolean if a field has been set.
func (o *ResourceServer) HasEnforcePolicies() bool {
	if o != nil && !IsNil(o.EnforcePolicies) {
		return true
	}

	return false
}

// SetEnforcePolicies gets a reference to the given bool and assigns it to the EnforcePolicies field.
func (o *ResourceServer) SetEnforcePolicies(v bool) {
	o.EnforcePolicies = &v
}

// GetTokenDialect returns the TokenDialect field value if set, zero value otherwise.
func (o *ResourceServer) GetTokenDialect() ResourceServerTokenDialect {
	if o == nil || IsNil(o.TokenDialect) {
		var ret ResourceServerTokenDialect
		return ret
	}
	return *o.TokenDialect
}

// GetTokenDialectOk returns a tuple with the TokenDialect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetTokenDialectOk() (*ResourceServerTokenDialect, bool) {
	if o == nil || IsNil(o.TokenDialect) {
		return nil, false
	}
	return o.TokenDialect, true
}

// HasTokenDialect returns a boolean if a field has been set.
func (o *ResourceServer) HasTokenDialect() bool {
	if o != nil && !IsNil(o.TokenDialect) {
		return true
	}

	return false
}

// SetTokenDialect gets a reference to the given ResourceServerTokenDialect and assigns it to the TokenDialect field.
func (o *ResourceServer) SetTokenDialect(v ResourceServerTokenDialect) {
	o.TokenDialect = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *ResourceServer) GetClient() map[string]interface{} {
	if o == nil || IsNil(o.Client) {
		var ret map[string]interface{}
		return ret
	}
	return o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetClientOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Client) {
		return map[string]interface{}{}, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *ResourceServer) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given map[string]interface{} and assigns it to the Client field.
func (o *ResourceServer) SetClient(v map[string]interface{}) {
	o.Client = v
}

// GetTokenEncryption returns the TokenEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceServer) GetTokenEncryption() ResourceServerTokenEncryption {
	if o == nil || IsNil(o.TokenEncryption.Get()) {
		var ret ResourceServerTokenEncryption
		return ret
	}
	return *o.TokenEncryption.Get()
}

// GetTokenEncryptionOk returns a tuple with the TokenEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceServer) GetTokenEncryptionOk() (*ResourceServerTokenEncryption, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenEncryption.Get(), o.TokenEncryption.IsSet()
}

// HasTokenEncryption returns a boolean if a field has been set.
func (o *ResourceServer) HasTokenEncryption() bool {
	if o != nil && o.TokenEncryption.IsSet() {
		return true
	}

	return false
}

// SetTokenEncryption gets a reference to the given NullableResourceServerTokenEncryption and assigns it to the TokenEncryption field.
func (o *ResourceServer) SetTokenEncryption(v ResourceServerTokenEncryption) {
	o.TokenEncryption.Set(&v)
}

// SetTokenEncryptionNil sets the value for TokenEncryption to be an explicit nil
func (o *ResourceServer) SetTokenEncryptionNil() {
	o.TokenEncryption.Set(nil)
}

// UnsetTokenEncryption ensures that no value is present for TokenEncryption, not even an explicit nil
func (o *ResourceServer) UnsetTokenEncryption() {
	o.TokenEncryption.Unset()
}

// GetConsentPolicy returns the ConsentPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceServer) GetConsentPolicy() ResourceServerConsentPolicy {
	if o == nil || IsNil(o.ConsentPolicy.Get()) {
		var ret ResourceServerConsentPolicy
		return ret
	}
	return *o.ConsentPolicy.Get()
}

// GetConsentPolicyOk returns a tuple with the ConsentPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceServer) GetConsentPolicyOk() (*ResourceServerConsentPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsentPolicy.Get(), o.ConsentPolicy.IsSet()
}

// HasConsentPolicy returns a boolean if a field has been set.
func (o *ResourceServer) HasConsentPolicy() bool {
	if o != nil && o.ConsentPolicy.IsSet() {
		return true
	}

	return false
}

// SetConsentPolicy gets a reference to the given NullableResourceServerConsentPolicy and assigns it to the ConsentPolicy field.
func (o *ResourceServer) SetConsentPolicy(v ResourceServerConsentPolicy) {
	o.ConsentPolicy.Set(&v)
}

// SetConsentPolicyNil sets the value for ConsentPolicy to be an explicit nil
func (o *ResourceServer) SetConsentPolicyNil() {
	o.ConsentPolicy.Set(nil)
}

// UnsetConsentPolicy ensures that no value is present for ConsentPolicy, not even an explicit nil
func (o *ResourceServer) UnsetConsentPolicy() {
	o.ConsentPolicy.Unset()
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *ResourceServer) GetAuthorizationDetails() []interface{} {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret []interface{}
		return ret
	}
	return o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServer) GetAuthorizationDetailsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *ResourceServer) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given []interface{} and assigns it to the AuthorizationDetails field.
func (o *ResourceServer) SetAuthorizationDetails(v []interface{}) {
	o.AuthorizationDetails = v
}

// GetProofOfPossession returns the ProofOfPossession field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceServer) GetProofOfPossession() ResourceServerProofOfPossession {
	if o == nil || IsNil(o.ProofOfPossession.Get()) {
		var ret ResourceServerProofOfPossession
		return ret
	}
	return *o.ProofOfPossession.Get()
}

// GetProofOfPossessionOk returns a tuple with the ProofOfPossession field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceServer) GetProofOfPossessionOk() (*ResourceServerProofOfPossession, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProofOfPossession.Get(), o.ProofOfPossession.IsSet()
}

// HasProofOfPossession returns a boolean if a field has been set.
func (o *ResourceServer) HasProofOfPossession() bool {
	if o != nil && o.ProofOfPossession.IsSet() {
		return true
	}

	return false
}

// SetProofOfPossession gets a reference to the given NullableResourceServerProofOfPossession and assigns it to the ProofOfPossession field.
func (o *ResourceServer) SetProofOfPossession(v ResourceServerProofOfPossession) {
	o.ProofOfPossession.Set(&v)
}

// SetProofOfPossessionNil sets the value for ProofOfPossession to be an explicit nil
func (o *ResourceServer) SetProofOfPossessionNil() {
	o.ProofOfPossession.Set(nil)
}

// UnsetProofOfPossession ensures that no value is present for ProofOfPossession, not even an explicit nil
func (o *ResourceServer) UnsetProofOfPossession() {
	o.ProofOfPossession.Unset()
}

func (o ResourceServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsSystem) {
		toSerialize["is_system"] = o.IsSystem
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.SigningAlg) {
		toSerialize["signing_alg"] = o.SigningAlg
	}
	if !IsNil(o.SigningSecret) {
		toSerialize["signing_secret"] = o.SigningSecret
	}
	if !IsNil(o.AllowOfflineAccess) {
		toSerialize["allow_offline_access"] = o.AllowOfflineAccess
	}
	if !IsNil(o.SkipConsentForVerifiableFirstPartyClients) {
		toSerialize["skip_consent_for_verifiable_first_party_clients"] = o.SkipConsentForVerifiableFirstPartyClients
	}
	if !IsNil(o.TokenLifetime) {
		toSerialize["token_lifetime"] = o.TokenLifetime
	}
	if !IsNil(o.TokenLifetimeForWeb) {
		toSerialize["token_lifetime_for_web"] = o.TokenLifetimeForWeb
	}
	if !IsNil(o.EnforcePolicies) {
		toSerialize["enforce_policies"] = o.EnforcePolicies
	}
	if !IsNil(o.TokenDialect) {
		toSerialize["token_dialect"] = o.TokenDialect
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if o.TokenEncryption.IsSet() {
		toSerialize["token_encryption"] = o.TokenEncryption.Get()
	}
	if o.ConsentPolicy.IsSet() {
		toSerialize["consent_policy"] = o.ConsentPolicy.Get()
	}
	if !IsNil(o.AuthorizationDetails) {
		toSerialize["authorization_details"] = o.AuthorizationDetails
	}
	if o.ProofOfPossession.IsSet() {
		toSerialize["proof_of_possession"] = o.ProofOfPossession.Get()
	}
	return toSerialize, nil
}

type NullableResourceServer struct {
	value *ResourceServer
	isSet bool
}

func (v NullableResourceServer) Get() *ResourceServer {
	return v.value
}

func (v *NullableResourceServer) Set(val *ResourceServer) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceServer) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceServer(val *ResourceServer) *NullableResourceServer {
	return &NullableResourceServer{value: val, isSet: true}
}

func (v NullableResourceServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
