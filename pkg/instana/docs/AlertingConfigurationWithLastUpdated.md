# AlertingConfigurationWithLastUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertName** | **string** |  | 
**CustomPayloadFields** | [**[]StaticStringField**](StaticStringField.md) |  | 
**EventFilteringConfiguration** | [**EventFilteringConfiguration**](EventFilteringConfiguration.md) |  | 
**Id** | **string** |  | 
**IntegrationIds** | **[]string** |  | 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**MuteUntil** | Pointer to **int64** |  | [optional] 

## Methods

### NewAlertingConfigurationWithLastUpdated

`func NewAlertingConfigurationWithLastUpdated(alertName string, customPayloadFields []StaticStringField, eventFilteringConfiguration EventFilteringConfiguration, id string, integrationIds []string, ) *AlertingConfigurationWithLastUpdated`

NewAlertingConfigurationWithLastUpdated instantiates a new AlertingConfigurationWithLastUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingConfigurationWithLastUpdatedWithDefaults

`func NewAlertingConfigurationWithLastUpdatedWithDefaults() *AlertingConfigurationWithLastUpdated`

NewAlertingConfigurationWithLastUpdatedWithDefaults instantiates a new AlertingConfigurationWithLastUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertName

`func (o *AlertingConfigurationWithLastUpdated) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *AlertingConfigurationWithLastUpdated) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *AlertingConfigurationWithLastUpdated) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.


### GetCustomPayloadFields

`func (o *AlertingConfigurationWithLastUpdated) GetCustomPayloadFields() []StaticStringField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *AlertingConfigurationWithLastUpdated) GetCustomPayloadFieldsOk() (*[]StaticStringField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *AlertingConfigurationWithLastUpdated) SetCustomPayloadFields(v []StaticStringField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetEventFilteringConfiguration

`func (o *AlertingConfigurationWithLastUpdated) GetEventFilteringConfiguration() EventFilteringConfiguration`

GetEventFilteringConfiguration returns the EventFilteringConfiguration field if non-nil, zero value otherwise.

### GetEventFilteringConfigurationOk

`func (o *AlertingConfigurationWithLastUpdated) GetEventFilteringConfigurationOk() (*EventFilteringConfiguration, bool)`

GetEventFilteringConfigurationOk returns a tuple with the EventFilteringConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilteringConfiguration

`func (o *AlertingConfigurationWithLastUpdated) SetEventFilteringConfiguration(v EventFilteringConfiguration)`

SetEventFilteringConfiguration sets EventFilteringConfiguration field to given value.


### GetId

`func (o *AlertingConfigurationWithLastUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingConfigurationWithLastUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingConfigurationWithLastUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationIds

`func (o *AlertingConfigurationWithLastUpdated) GetIntegrationIds() []string`

GetIntegrationIds returns the IntegrationIds field if non-nil, zero value otherwise.

### GetIntegrationIdsOk

`func (o *AlertingConfigurationWithLastUpdated) GetIntegrationIdsOk() (*[]string, bool)`

GetIntegrationIdsOk returns a tuple with the IntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIds

`func (o *AlertingConfigurationWithLastUpdated) SetIntegrationIds(v []string)`

SetIntegrationIds sets IntegrationIds field to given value.


### GetLastUpdated

`func (o *AlertingConfigurationWithLastUpdated) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AlertingConfigurationWithLastUpdated) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AlertingConfigurationWithLastUpdated) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AlertingConfigurationWithLastUpdated) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMuteUntil

`func (o *AlertingConfigurationWithLastUpdated) GetMuteUntil() int64`

GetMuteUntil returns the MuteUntil field if non-nil, zero value otherwise.

### GetMuteUntilOk

`func (o *AlertingConfigurationWithLastUpdated) GetMuteUntilOk() (*int64, bool)`

GetMuteUntilOk returns a tuple with the MuteUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteUntil

`func (o *AlertingConfigurationWithLastUpdated) SetMuteUntil(v int64)`

SetMuteUntil sets MuteUntil field to given value.

### HasMuteUntil

`func (o *AlertingConfigurationWithLastUpdated) HasMuteUntil() bool`

HasMuteUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


