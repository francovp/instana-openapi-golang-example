# ServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**MatchSpecification** | [**[]ServiceMatchingRule**](ServiceMatchingRule.md) |  | 
**Name** | **string** |  | 

## Methods

### NewServiceConfig

`func NewServiceConfig(enabled bool, id string, label string, matchSpecification []ServiceMatchingRule, name string, ) *ServiceConfig`

NewServiceConfig instantiates a new ServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceConfigWithDefaults

`func NewServiceConfigWithDefaults() *ServiceConfig`

NewServiceConfigWithDefaults instantiates a new ServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ServiceConfig) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServiceConfig) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServiceConfig) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServiceConfig) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *ServiceConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceConfig) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *ServiceConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceConfig) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMatchSpecification

`func (o *ServiceConfig) GetMatchSpecification() []ServiceMatchingRule`

GetMatchSpecification returns the MatchSpecification field if non-nil, zero value otherwise.

### GetMatchSpecificationOk

`func (o *ServiceConfig) GetMatchSpecificationOk() (*[]ServiceMatchingRule, bool)`

GetMatchSpecificationOk returns a tuple with the MatchSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSpecification

`func (o *ServiceConfig) SetMatchSpecification(v []ServiceMatchingRule)`

SetMatchSpecification sets MatchSpecification field to given value.


### GetName

`func (o *ServiceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceConfig) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


