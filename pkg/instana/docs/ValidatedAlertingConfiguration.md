# ValidatedAlertingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelNames** | Pointer to **[]string** |  | [optional] 
**AlertName** | **string** |  | 
**ApplicationNames** | Pointer to **[]string** |  | [optional] 
**CustomPayloadFields** | [**[]StaticStringField**](StaticStringField.md) |  | 
**EventFilteringConfiguration** | [**EventFilteringConfiguration**](EventFilteringConfiguration.md) |  | 
**Id** | **string** |  | 
**IntegrationIds** | **[]string** |  | 
**Invalid** | Pointer to **bool** |  | [optional] 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**MuteUntil** | Pointer to **int64** |  | [optional] 

## Methods

### NewValidatedAlertingConfiguration

`func NewValidatedAlertingConfiguration(alertName string, customPayloadFields []StaticStringField, eventFilteringConfiguration EventFilteringConfiguration, id string, integrationIds []string, ) *ValidatedAlertingConfiguration`

NewValidatedAlertingConfiguration instantiates a new ValidatedAlertingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatedAlertingConfigurationWithDefaults

`func NewValidatedAlertingConfigurationWithDefaults() *ValidatedAlertingConfiguration`

NewValidatedAlertingConfigurationWithDefaults instantiates a new ValidatedAlertingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertChannelNames

`func (o *ValidatedAlertingConfiguration) GetAlertChannelNames() []string`

GetAlertChannelNames returns the AlertChannelNames field if non-nil, zero value otherwise.

### GetAlertChannelNamesOk

`func (o *ValidatedAlertingConfiguration) GetAlertChannelNamesOk() (*[]string, bool)`

GetAlertChannelNamesOk returns a tuple with the AlertChannelNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannelNames

`func (o *ValidatedAlertingConfiguration) SetAlertChannelNames(v []string)`

SetAlertChannelNames sets AlertChannelNames field to given value.

### HasAlertChannelNames

`func (o *ValidatedAlertingConfiguration) HasAlertChannelNames() bool`

HasAlertChannelNames returns a boolean if a field has been set.

### GetAlertName

`func (o *ValidatedAlertingConfiguration) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *ValidatedAlertingConfiguration) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *ValidatedAlertingConfiguration) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.


### GetApplicationNames

`func (o *ValidatedAlertingConfiguration) GetApplicationNames() []string`

GetApplicationNames returns the ApplicationNames field if non-nil, zero value otherwise.

### GetApplicationNamesOk

`func (o *ValidatedAlertingConfiguration) GetApplicationNamesOk() (*[]string, bool)`

GetApplicationNamesOk returns a tuple with the ApplicationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNames

`func (o *ValidatedAlertingConfiguration) SetApplicationNames(v []string)`

SetApplicationNames sets ApplicationNames field to given value.

### HasApplicationNames

`func (o *ValidatedAlertingConfiguration) HasApplicationNames() bool`

HasApplicationNames returns a boolean if a field has been set.

### GetCustomPayloadFields

`func (o *ValidatedAlertingConfiguration) GetCustomPayloadFields() []StaticStringField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *ValidatedAlertingConfiguration) GetCustomPayloadFieldsOk() (*[]StaticStringField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *ValidatedAlertingConfiguration) SetCustomPayloadFields(v []StaticStringField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetEventFilteringConfiguration

`func (o *ValidatedAlertingConfiguration) GetEventFilteringConfiguration() EventFilteringConfiguration`

GetEventFilteringConfiguration returns the EventFilteringConfiguration field if non-nil, zero value otherwise.

### GetEventFilteringConfigurationOk

`func (o *ValidatedAlertingConfiguration) GetEventFilteringConfigurationOk() (*EventFilteringConfiguration, bool)`

GetEventFilteringConfigurationOk returns a tuple with the EventFilteringConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilteringConfiguration

`func (o *ValidatedAlertingConfiguration) SetEventFilteringConfiguration(v EventFilteringConfiguration)`

SetEventFilteringConfiguration sets EventFilteringConfiguration field to given value.


### GetId

`func (o *ValidatedAlertingConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValidatedAlertingConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValidatedAlertingConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationIds

`func (o *ValidatedAlertingConfiguration) GetIntegrationIds() []string`

GetIntegrationIds returns the IntegrationIds field if non-nil, zero value otherwise.

### GetIntegrationIdsOk

`func (o *ValidatedAlertingConfiguration) GetIntegrationIdsOk() (*[]string, bool)`

GetIntegrationIdsOk returns a tuple with the IntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIds

`func (o *ValidatedAlertingConfiguration) SetIntegrationIds(v []string)`

SetIntegrationIds sets IntegrationIds field to given value.


### GetInvalid

`func (o *ValidatedAlertingConfiguration) GetInvalid() bool`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *ValidatedAlertingConfiguration) GetInvalidOk() (*bool, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *ValidatedAlertingConfiguration) SetInvalid(v bool)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *ValidatedAlertingConfiguration) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ValidatedAlertingConfiguration) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ValidatedAlertingConfiguration) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ValidatedAlertingConfiguration) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ValidatedAlertingConfiguration) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMuteUntil

`func (o *ValidatedAlertingConfiguration) GetMuteUntil() int64`

GetMuteUntil returns the MuteUntil field if non-nil, zero value otherwise.

### GetMuteUntilOk

`func (o *ValidatedAlertingConfiguration) GetMuteUntilOk() (*int64, bool)`

GetMuteUntilOk returns a tuple with the MuteUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteUntil

`func (o *ValidatedAlertingConfiguration) SetMuteUntil(v int64)`

SetMuteUntil sets MuteUntil field to given value.

### HasMuteUntil

`func (o *ValidatedAlertingConfiguration) HasMuteUntil() bool`

HasMuteUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


