# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IsSynthetic** | Pointer to **bool** |  | [optional] 
**Label** | **string** |  | 
**ServiceId** | **string** |  | 
**Synthetic** | Pointer to **bool** |  | [optional] 
**SyntheticType** | Pointer to **string** |  | [optional] 
**Technologies** | **[]string** |  | 
**Type** | **string** |  | 

## Methods

### NewEndpoint

`func NewEndpoint(id string, label string, serviceId string, technologies []string, type_ string, ) *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *Endpoint) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Endpoint) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Endpoint) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *Endpoint) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetId

`func (o *Endpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Endpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Endpoint) SetId(v string)`

SetId sets Id field to given value.


### GetIsSynthetic

`func (o *Endpoint) GetIsSynthetic() bool`

GetIsSynthetic returns the IsSynthetic field if non-nil, zero value otherwise.

### GetIsSyntheticOk

`func (o *Endpoint) GetIsSyntheticOk() (*bool, bool)`

GetIsSyntheticOk returns a tuple with the IsSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynthetic

`func (o *Endpoint) SetIsSynthetic(v bool)`

SetIsSynthetic sets IsSynthetic field to given value.

### HasIsSynthetic

`func (o *Endpoint) HasIsSynthetic() bool`

HasIsSynthetic returns a boolean if a field has been set.

### GetLabel

`func (o *Endpoint) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Endpoint) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Endpoint) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetServiceId

`func (o *Endpoint) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Endpoint) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Endpoint) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSynthetic

`func (o *Endpoint) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *Endpoint) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *Endpoint) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *Endpoint) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetSyntheticType

`func (o *Endpoint) GetSyntheticType() string`

GetSyntheticType returns the SyntheticType field if non-nil, zero value otherwise.

### GetSyntheticTypeOk

`func (o *Endpoint) GetSyntheticTypeOk() (*string, bool)`

GetSyntheticTypeOk returns a tuple with the SyntheticType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticType

`func (o *Endpoint) SetSyntheticType(v string)`

SetSyntheticType sets SyntheticType field to given value.

### HasSyntheticType

`func (o *Endpoint) HasSyntheticType() bool`

HasSyntheticType returns a boolean if a field has been set.

### GetTechnologies

`func (o *Endpoint) GetTechnologies() []string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *Endpoint) GetTechnologiesOk() (*[]string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *Endpoint) SetTechnologies(v []string)`

SetTechnologies sets Technologies field to given value.


### GetType

`func (o *Endpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Endpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Endpoint) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


