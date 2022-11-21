# ManualServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExistingServiceId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**TagFilterExpression** | [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | 
**UnmonitoredServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewManualServiceConfig

`func NewManualServiceConfig(id string, tagFilterExpression TagFilterExpressionElement, ) *ManualServiceConfig`

NewManualServiceConfig instantiates a new ManualServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualServiceConfigWithDefaults

`func NewManualServiceConfigWithDefaults() *ManualServiceConfig`

NewManualServiceConfigWithDefaults instantiates a new ManualServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManualServiceConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManualServiceConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManualServiceConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManualServiceConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ManualServiceConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManualServiceConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManualServiceConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ManualServiceConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExistingServiceId

`func (o *ManualServiceConfig) GetExistingServiceId() string`

GetExistingServiceId returns the ExistingServiceId field if non-nil, zero value otherwise.

### GetExistingServiceIdOk

`func (o *ManualServiceConfig) GetExistingServiceIdOk() (*string, bool)`

GetExistingServiceIdOk returns a tuple with the ExistingServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingServiceId

`func (o *ManualServiceConfig) SetExistingServiceId(v string)`

SetExistingServiceId sets ExistingServiceId field to given value.

### HasExistingServiceId

`func (o *ManualServiceConfig) HasExistingServiceId() bool`

HasExistingServiceId returns a boolean if a field has been set.

### GetId

`func (o *ManualServiceConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualServiceConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualServiceConfig) SetId(v string)`

SetId sets Id field to given value.


### GetTagFilterExpression

`func (o *ManualServiceConfig) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *ManualServiceConfig) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *ManualServiceConfig) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.


### GetUnmonitoredServiceName

`func (o *ManualServiceConfig) GetUnmonitoredServiceName() string`

GetUnmonitoredServiceName returns the UnmonitoredServiceName field if non-nil, zero value otherwise.

### GetUnmonitoredServiceNameOk

`func (o *ManualServiceConfig) GetUnmonitoredServiceNameOk() (*string, bool)`

GetUnmonitoredServiceNameOk returns a tuple with the UnmonitoredServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmonitoredServiceName

`func (o *ManualServiceConfig) SetUnmonitoredServiceName(v string)`

SetUnmonitoredServiceName sets UnmonitoredServiceName field to given value.

### HasUnmonitoredServiceName

`func (o *ManualServiceConfig) HasUnmonitoredServiceName() bool`

HasUnmonitoredServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


