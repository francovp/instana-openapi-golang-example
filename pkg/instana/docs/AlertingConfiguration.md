# AlertingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertName** | **string** |  | 
**CustomPayloadFields** | [**[]StaticStringField**](StaticStringField.md) |  | 
**EventFilteringConfiguration** | [**EventFilteringConfiguration**](EventFilteringConfiguration.md) |  | 
**Id** | **string** |  | 
**IntegrationIds** | **[]string** |  | 
**MuteUntil** | Pointer to **int64** |  | [optional] 

## Methods

### NewAlertingConfiguration

`func NewAlertingConfiguration(alertName string, customPayloadFields []StaticStringField, eventFilteringConfiguration EventFilteringConfiguration, id string, integrationIds []string, ) *AlertingConfiguration`

NewAlertingConfiguration instantiates a new AlertingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingConfigurationWithDefaults

`func NewAlertingConfigurationWithDefaults() *AlertingConfiguration`

NewAlertingConfigurationWithDefaults instantiates a new AlertingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertName

`func (o *AlertingConfiguration) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *AlertingConfiguration) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *AlertingConfiguration) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.


### GetCustomPayloadFields

`func (o *AlertingConfiguration) GetCustomPayloadFields() []StaticStringField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *AlertingConfiguration) GetCustomPayloadFieldsOk() (*[]StaticStringField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *AlertingConfiguration) SetCustomPayloadFields(v []StaticStringField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetEventFilteringConfiguration

`func (o *AlertingConfiguration) GetEventFilteringConfiguration() EventFilteringConfiguration`

GetEventFilteringConfiguration returns the EventFilteringConfiguration field if non-nil, zero value otherwise.

### GetEventFilteringConfigurationOk

`func (o *AlertingConfiguration) GetEventFilteringConfigurationOk() (*EventFilteringConfiguration, bool)`

GetEventFilteringConfigurationOk returns a tuple with the EventFilteringConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilteringConfiguration

`func (o *AlertingConfiguration) SetEventFilteringConfiguration(v EventFilteringConfiguration)`

SetEventFilteringConfiguration sets EventFilteringConfiguration field to given value.


### GetId

`func (o *AlertingConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationIds

`func (o *AlertingConfiguration) GetIntegrationIds() []string`

GetIntegrationIds returns the IntegrationIds field if non-nil, zero value otherwise.

### GetIntegrationIdsOk

`func (o *AlertingConfiguration) GetIntegrationIdsOk() (*[]string, bool)`

GetIntegrationIdsOk returns a tuple with the IntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIds

`func (o *AlertingConfiguration) SetIntegrationIds(v []string)`

SetIntegrationIds sets IntegrationIds field to given value.


### GetMuteUntil

`func (o *AlertingConfiguration) GetMuteUntil() int64`

GetMuteUntil returns the MuteUntil field if non-nil, zero value otherwise.

### GetMuteUntilOk

`func (o *AlertingConfiguration) GetMuteUntilOk() (*int64, bool)`

GetMuteUntilOk returns a tuple with the MuteUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteUntil

`func (o *AlertingConfiguration) SetMuteUntil(v int64)`

SetMuteUntil sets MuteUntil field to given value.

### HasMuteUntil

`func (o *AlertingConfiguration) HasMuteUntil() bool`

HasMuteUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


