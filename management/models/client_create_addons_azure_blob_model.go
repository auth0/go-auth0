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

// ClientCreateAddonsAzureBlob Azure Blob Storage addon configuration.
type ClientCreateAddonsAzureBlob struct {
	// Your Azure storage account name. Usually first segment in your Azure storage URL. e.g. `https://acme-org.blob.core.windows.net` would be the account name `acme-org`.
	AccountName *string `json:"accountName,omitempty"`
	// Access key associated with this storage account.
	StorageAccessKey *string `json:"storageAccessKey,omitempty"`
	// Container to request a token for. e.g. `my-container`.
	ContainerName *string `json:"containerName,omitempty"`
	// Entity to request a token for. e.g. `my-blob`. If blank the computed SAS will apply to the entire storage container.
	BlobName *string `json:"blobName,omitempty"`
	// Expiration in minutes for the generated token (default of 5 minutes).
	Expiration *int32 `json:"expiration,omitempty"`
	// Shared access policy identifier defined in your storage account resource.
	SignedIdentifier *string `json:"signedIdentifier,omitempty"`
	// Indicates if the issued token has permission to read the content, properties, metadata and block list. Use the blob as the source of a copy operation.
	BlobRead *bool `json:"blob_read,omitempty"`
	// Indicates if the issued token has permission to create or write content, properties, metadata, or block list. Snapshot or lease the blob. Resize the blob (page blob only). Use the blob as the destination of a copy operation within the same account.
	BlobWrite *bool `json:"blob_write,omitempty"`
	// Indicates if the issued token has permission to delete the blob.
	BlobDelete *bool `json:"blob_delete,omitempty"`
	// Indicates if the issued token has permission to read the content, properties, metadata or block list of any blob in the container. Use any blob in the container as the source of a copy operation
	ContainerRead *bool `json:"container_read,omitempty"`
	// Indicates that for any blob in the container if the issued token has permission to create or write content, properties, metadata, or block list. Snapshot or lease the blob. Resize the blob (page blob only). Use the blob as the destination of a copy operation within the same account.
	ContainerWrite *bool `json:"container_write,omitempty"`
	// Indicates if issued token has permission to delete any blob in the container.
	ContainerDelete *bool `json:"container_delete,omitempty"`
	// Indicates if the issued token has permission to list blobs in the container.
	ContainerList        *bool `json:"container_list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientCreateAddonsAzureBlob ClientCreateAddonsAzureBlob

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ClientCreateAddonsAzureBlob) SetAccountName(v string) {
	o.AccountName = &v
}

// GetStorageAccessKey returns the StorageAccessKey field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetStorageAccessKey() string {
	if o == nil || IsNil(o.StorageAccessKey) {
		var ret string
		return ret
	}
	return *o.StorageAccessKey
}

// GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetStorageAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StorageAccessKey) {
		return nil, false
	}
	return o.StorageAccessKey, true
}

// HasStorageAccessKey returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasStorageAccessKey() bool {
	if o != nil && !IsNil(o.StorageAccessKey) {
		return true
	}

	return false
}

// SetStorageAccessKey gets a reference to the given string and assigns it to the StorageAccessKey field.
func (o *ClientCreateAddonsAzureBlob) SetStorageAccessKey(v string) {
	o.StorageAccessKey = &v
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetContainerName() string {
	if o == nil || IsNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetContainerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasContainerName() bool {
	if o != nil && !IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *ClientCreateAddonsAzureBlob) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetBlobName returns the BlobName field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetBlobName() string {
	if o == nil || IsNil(o.BlobName) {
		var ret string
		return ret
	}
	return *o.BlobName
}

// GetBlobNameOk returns a tuple with the BlobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetBlobNameOk() (*string, bool) {
	if o == nil || IsNil(o.BlobName) {
		return nil, false
	}
	return o.BlobName, true
}

// HasBlobName returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasBlobName() bool {
	if o != nil && !IsNil(o.BlobName) {
		return true
	}

	return false
}

// SetBlobName gets a reference to the given string and assigns it to the BlobName field.
func (o *ClientCreateAddonsAzureBlob) SetBlobName(v string) {
	o.BlobName = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetExpiration() int32 {
	if o == nil || IsNil(o.Expiration) {
		var ret int32
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetExpirationOk() (*int32, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int32 and assigns it to the Expiration field.
func (o *ClientCreateAddonsAzureBlob) SetExpiration(v int32) {
	o.Expiration = &v
}

// GetSignedIdentifier returns the SignedIdentifier field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetSignedIdentifier() string {
	if o == nil || IsNil(o.SignedIdentifier) {
		var ret string
		return ret
	}
	return *o.SignedIdentifier
}

// GetSignedIdentifierOk returns a tuple with the SignedIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetSignedIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.SignedIdentifier) {
		return nil, false
	}
	return o.SignedIdentifier, true
}

// HasSignedIdentifier returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasSignedIdentifier() bool {
	if o != nil && !IsNil(o.SignedIdentifier) {
		return true
	}

	return false
}

// SetSignedIdentifier gets a reference to the given string and assigns it to the SignedIdentifier field.
func (o *ClientCreateAddonsAzureBlob) SetSignedIdentifier(v string) {
	o.SignedIdentifier = &v
}

// GetBlobRead returns the BlobRead field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetBlobRead() bool {
	if o == nil || IsNil(o.BlobRead) {
		var ret bool
		return ret
	}
	return *o.BlobRead
}

// GetBlobReadOk returns a tuple with the BlobRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetBlobReadOk() (*bool, bool) {
	if o == nil || IsNil(o.BlobRead) {
		return nil, false
	}
	return o.BlobRead, true
}

// HasBlobRead returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasBlobRead() bool {
	if o != nil && !IsNil(o.BlobRead) {
		return true
	}

	return false
}

// SetBlobRead gets a reference to the given bool and assigns it to the BlobRead field.
func (o *ClientCreateAddonsAzureBlob) SetBlobRead(v bool) {
	o.BlobRead = &v
}

// GetBlobWrite returns the BlobWrite field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetBlobWrite() bool {
	if o == nil || IsNil(o.BlobWrite) {
		var ret bool
		return ret
	}
	return *o.BlobWrite
}

// GetBlobWriteOk returns a tuple with the BlobWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetBlobWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.BlobWrite) {
		return nil, false
	}
	return o.BlobWrite, true
}

// HasBlobWrite returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasBlobWrite() bool {
	if o != nil && !IsNil(o.BlobWrite) {
		return true
	}

	return false
}

// SetBlobWrite gets a reference to the given bool and assigns it to the BlobWrite field.
func (o *ClientCreateAddonsAzureBlob) SetBlobWrite(v bool) {
	o.BlobWrite = &v
}

// GetBlobDelete returns the BlobDelete field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetBlobDelete() bool {
	if o == nil || IsNil(o.BlobDelete) {
		var ret bool
		return ret
	}
	return *o.BlobDelete
}

// GetBlobDeleteOk returns a tuple with the BlobDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetBlobDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.BlobDelete) {
		return nil, false
	}
	return o.BlobDelete, true
}

// HasBlobDelete returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasBlobDelete() bool {
	if o != nil && !IsNil(o.BlobDelete) {
		return true
	}

	return false
}

// SetBlobDelete gets a reference to the given bool and assigns it to the BlobDelete field.
func (o *ClientCreateAddonsAzureBlob) SetBlobDelete(v bool) {
	o.BlobDelete = &v
}

// GetContainerRead returns the ContainerRead field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetContainerRead() bool {
	if o == nil || IsNil(o.ContainerRead) {
		var ret bool
		return ret
	}
	return *o.ContainerRead
}

// GetContainerReadOk returns a tuple with the ContainerRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetContainerReadOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainerRead) {
		return nil, false
	}
	return o.ContainerRead, true
}

// HasContainerRead returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasContainerRead() bool {
	if o != nil && !IsNil(o.ContainerRead) {
		return true
	}

	return false
}

// SetContainerRead gets a reference to the given bool and assigns it to the ContainerRead field.
func (o *ClientCreateAddonsAzureBlob) SetContainerRead(v bool) {
	o.ContainerRead = &v
}

// GetContainerWrite returns the ContainerWrite field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetContainerWrite() bool {
	if o == nil || IsNil(o.ContainerWrite) {
		var ret bool
		return ret
	}
	return *o.ContainerWrite
}

// GetContainerWriteOk returns a tuple with the ContainerWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetContainerWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainerWrite) {
		return nil, false
	}
	return o.ContainerWrite, true
}

// HasContainerWrite returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasContainerWrite() bool {
	if o != nil && !IsNil(o.ContainerWrite) {
		return true
	}

	return false
}

// SetContainerWrite gets a reference to the given bool and assigns it to the ContainerWrite field.
func (o *ClientCreateAddonsAzureBlob) SetContainerWrite(v bool) {
	o.ContainerWrite = &v
}

// GetContainerDelete returns the ContainerDelete field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetContainerDelete() bool {
	if o == nil || IsNil(o.ContainerDelete) {
		var ret bool
		return ret
	}
	return *o.ContainerDelete
}

// GetContainerDeleteOk returns a tuple with the ContainerDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetContainerDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainerDelete) {
		return nil, false
	}
	return o.ContainerDelete, true
}

// HasContainerDelete returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasContainerDelete() bool {
	if o != nil && !IsNil(o.ContainerDelete) {
		return true
	}

	return false
}

// SetContainerDelete gets a reference to the given bool and assigns it to the ContainerDelete field.
func (o *ClientCreateAddonsAzureBlob) SetContainerDelete(v bool) {
	o.ContainerDelete = &v
}

// GetContainerList returns the ContainerList field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureBlob) GetContainerList() bool {
	if o == nil || IsNil(o.ContainerList) {
		var ret bool
		return ret
	}
	return *o.ContainerList
}

// GetContainerListOk returns a tuple with the ContainerList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureBlob) GetContainerListOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainerList) {
		return nil, false
	}
	return o.ContainerList, true
}

// HasContainerList returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureBlob) HasContainerList() bool {
	if o != nil && !IsNil(o.ContainerList) {
		return true
	}

	return false
}

// SetContainerList gets a reference to the given bool and assigns it to the ContainerList field.
func (o *ClientCreateAddonsAzureBlob) SetContainerList(v bool) {
	o.ContainerList = &v
}

func (o ClientCreateAddonsAzureBlob) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCreateAddonsAzureBlob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.StorageAccessKey) {
		toSerialize["storageAccessKey"] = o.StorageAccessKey
	}
	if !IsNil(o.ContainerName) {
		toSerialize["containerName"] = o.ContainerName
	}
	if !IsNil(o.BlobName) {
		toSerialize["blobName"] = o.BlobName
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.SignedIdentifier) {
		toSerialize["signedIdentifier"] = o.SignedIdentifier
	}
	if !IsNil(o.BlobRead) {
		toSerialize["blob_read"] = o.BlobRead
	}
	if !IsNil(o.BlobWrite) {
		toSerialize["blob_write"] = o.BlobWrite
	}
	if !IsNil(o.BlobDelete) {
		toSerialize["blob_delete"] = o.BlobDelete
	}
	if !IsNil(o.ContainerRead) {
		toSerialize["container_read"] = o.ContainerRead
	}
	if !IsNil(o.ContainerWrite) {
		toSerialize["container_write"] = o.ContainerWrite
	}
	if !IsNil(o.ContainerDelete) {
		toSerialize["container_delete"] = o.ContainerDelete
	}
	if !IsNil(o.ContainerList) {
		toSerialize["container_list"] = o.ContainerList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientCreateAddonsAzureBlob) UnmarshalJSON(data []byte) (err error) {
	varClientCreateAddonsAzureBlob := _ClientCreateAddonsAzureBlob{}

	err = json.Unmarshal(data, &varClientCreateAddonsAzureBlob)

	if err != nil {
		return err
	}

	*o = ClientCreateAddonsAzureBlob(varClientCreateAddonsAzureBlob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "storageAccessKey")
		delete(additionalProperties, "containerName")
		delete(additionalProperties, "blobName")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "signedIdentifier")
		delete(additionalProperties, "blob_read")
		delete(additionalProperties, "blob_write")
		delete(additionalProperties, "blob_delete")
		delete(additionalProperties, "container_read")
		delete(additionalProperties, "container_write")
		delete(additionalProperties, "container_delete")
		delete(additionalProperties, "container_list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientCreateAddonsAzureBlob struct {
	value *ClientCreateAddonsAzureBlob
	isSet bool
}

func (v NullableClientCreateAddonsAzureBlob) Get() *ClientCreateAddonsAzureBlob {
	return v.value
}

func (v *NullableClientCreateAddonsAzureBlob) Set(val *ClientCreateAddonsAzureBlob) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCreateAddonsAzureBlob) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCreateAddonsAzureBlob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCreateAddonsAzureBlob(val *ClientCreateAddonsAzureBlob) *NullableClientCreateAddonsAzureBlob {
	return &NullableClientCreateAddonsAzureBlob{value: val, isSet: true}
}

func (v NullableClientCreateAddonsAzureBlob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCreateAddonsAzureBlob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
