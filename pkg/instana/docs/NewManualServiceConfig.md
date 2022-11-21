# NewManualServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExistingServiceId** | Pointer to **string** |  | [optional] 
**TagFilterExpression** | [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | 
**UnmonitoredServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewNewManualServiceConfig

`func NewNewManualServiceConfig(tagFilterExpression TagFilterExpressionElement, ) *NewManualServiceConfig`

NewNewManualServiceConfig instantiates a new NewManualServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewManualServiceConfigWithDefaults

`func NewNewManualServiceConfigWithDefaults() *NewManualServiceConfig`

NewNewManualServiceConfigWithDefaults instantiates a new NewManualServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NewManualServiceConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewManualServiceConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewManualServiceConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewManualServiceConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *NewManualServiceConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewManualServiceConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewManualServiceConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NewManualServiceConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExistingServiceId

`func (o *NewManualServiceConfig) GetExistingServiceId() string`

GetExistingServiceId returns the ExistingServiceId field if non-nil, zero value otherwise.

### GetExistingServiceIdOk

`func (o *NewManualServiceConfig) GetExistingServiceIdOk() (*string, bool)`

GetExistingServiceIdOk returns a tuple with the ExistingServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingServiceId

`func (o *NewManualServiceConfig) SetExistingServiceId(v string)`

SetExistingServiceId sets ExistingServiceId field to given value.

### HasExistingServiceId

`func (o *NewManualServiceConfig) HasExistingServiceId() bool`

HasExistingServiceId returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *NewManualServiceConfig) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *NewManualServiceConfig) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *NewManualServiceConfig) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.


### GetUnmonitoredServiceName

`func (o *NewManualServiceConfig) GetUnmonitoredServiceName() string`

GetUnmonitoredServiceName returns the UnmonitoredServiceName field if non-nil, zero value otherwise.

### GetUnmonitoredServiceNameOk

`func (o *NewManualServiceConfig) GetUnmonitoredServiceNameOk() (*string, bool)`

GetUnmonitoredServiceNameOk returns a tuple with the UnmonitoredServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmonitoredServiceName

`func (o *NewManualServiceConfig) SetUnmonitoredServiceName(v string)`

SetUnmonitoredServiceName sets UnmonitoredServiceName field to given value.

### HasUnmonitoredServiceName

`func (o *NewManualServiceConfig) HasUnmonitoredServiceName() bool`

HasUnmonitoredServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


