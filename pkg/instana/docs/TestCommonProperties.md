# TestCommonProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ApplicationLabel** | Pointer to **string** |  | [optional] 
**Frequency** | **int32** |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**LocationDisplayLabels** | Pointer to **[]string** |  | [optional] 
**LocationIds** | Pointer to **[]string** |  | [optional] 
**LocationLabels** | Pointer to **[]string** |  | [optional] 
**LocationStatusList** | Pointer to [**[]LocationStatus**](LocationStatus.md) |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewTestCommonProperties

`func NewTestCommonProperties(active bool, frequency int32, id string, label string, type_ string, ) *TestCommonProperties`

NewTestCommonProperties instantiates a new TestCommonProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCommonPropertiesWithDefaults

`func NewTestCommonPropertiesWithDefaults() *TestCommonProperties`

NewTestCommonPropertiesWithDefaults instantiates a new TestCommonProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *TestCommonProperties) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TestCommonProperties) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TestCommonProperties) SetActive(v bool)`

SetActive sets Active field to given value.


### GetApplicationId

`func (o *TestCommonProperties) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *TestCommonProperties) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *TestCommonProperties) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *TestCommonProperties) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationLabel

`func (o *TestCommonProperties) GetApplicationLabel() string`

GetApplicationLabel returns the ApplicationLabel field if non-nil, zero value otherwise.

### GetApplicationLabelOk

`func (o *TestCommonProperties) GetApplicationLabelOk() (*string, bool)`

GetApplicationLabelOk returns a tuple with the ApplicationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLabel

`func (o *TestCommonProperties) SetApplicationLabel(v string)`

SetApplicationLabel sets ApplicationLabel field to given value.

### HasApplicationLabel

`func (o *TestCommonProperties) HasApplicationLabel() bool`

HasApplicationLabel returns a boolean if a field has been set.

### GetFrequency

`func (o *TestCommonProperties) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *TestCommonProperties) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *TestCommonProperties) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.


### GetId

`func (o *TestCommonProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestCommonProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestCommonProperties) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *TestCommonProperties) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TestCommonProperties) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TestCommonProperties) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLocationDisplayLabels

`func (o *TestCommonProperties) GetLocationDisplayLabels() []string`

GetLocationDisplayLabels returns the LocationDisplayLabels field if non-nil, zero value otherwise.

### GetLocationDisplayLabelsOk

`func (o *TestCommonProperties) GetLocationDisplayLabelsOk() (*[]string, bool)`

GetLocationDisplayLabelsOk returns a tuple with the LocationDisplayLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDisplayLabels

`func (o *TestCommonProperties) SetLocationDisplayLabels(v []string)`

SetLocationDisplayLabels sets LocationDisplayLabels field to given value.

### HasLocationDisplayLabels

`func (o *TestCommonProperties) HasLocationDisplayLabels() bool`

HasLocationDisplayLabels returns a boolean if a field has been set.

### GetLocationIds

`func (o *TestCommonProperties) GetLocationIds() []string`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *TestCommonProperties) GetLocationIdsOk() (*[]string, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *TestCommonProperties) SetLocationIds(v []string)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *TestCommonProperties) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetLocationLabels

`func (o *TestCommonProperties) GetLocationLabels() []string`

GetLocationLabels returns the LocationLabels field if non-nil, zero value otherwise.

### GetLocationLabelsOk

`func (o *TestCommonProperties) GetLocationLabelsOk() (*[]string, bool)`

GetLocationLabelsOk returns a tuple with the LocationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationLabels

`func (o *TestCommonProperties) SetLocationLabels(v []string)`

SetLocationLabels sets LocationLabels field to given value.

### HasLocationLabels

`func (o *TestCommonProperties) HasLocationLabels() bool`

HasLocationLabels returns a boolean if a field has been set.

### GetLocationStatusList

`func (o *TestCommonProperties) GetLocationStatusList() []LocationStatus`

GetLocationStatusList returns the LocationStatusList field if non-nil, zero value otherwise.

### GetLocationStatusListOk

`func (o *TestCommonProperties) GetLocationStatusListOk() (*[]LocationStatus, bool)`

GetLocationStatusListOk returns a tuple with the LocationStatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationStatusList

`func (o *TestCommonProperties) SetLocationStatusList(v []LocationStatus)`

SetLocationStatusList sets LocationStatusList field to given value.

### HasLocationStatusList

`func (o *TestCommonProperties) HasLocationStatusList() bool`

HasLocationStatusList returns a boolean if a field has been set.

### GetServiceId

`func (o *TestCommonProperties) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TestCommonProperties) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TestCommonProperties) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *TestCommonProperties) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetType

`func (o *TestCommonProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCommonProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCommonProperties) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


