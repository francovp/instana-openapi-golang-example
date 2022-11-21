# CustomPayloadWithLastUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]CustomPayloadField**](CustomPayloadField.md) |  | 
**LastUpdated** | Pointer to **int64** |  | [optional] 

## Methods

### NewCustomPayloadWithLastUpdated

`func NewCustomPayloadWithLastUpdated(fields []CustomPayloadField, ) *CustomPayloadWithLastUpdated`

NewCustomPayloadWithLastUpdated instantiates a new CustomPayloadWithLastUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPayloadWithLastUpdatedWithDefaults

`func NewCustomPayloadWithLastUpdatedWithDefaults() *CustomPayloadWithLastUpdated`

NewCustomPayloadWithLastUpdatedWithDefaults instantiates a new CustomPayloadWithLastUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *CustomPayloadWithLastUpdated) GetFields() []CustomPayloadField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomPayloadWithLastUpdated) GetFieldsOk() (*[]CustomPayloadField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomPayloadWithLastUpdated) SetFields(v []CustomPayloadField)`

SetFields sets Fields field to given value.


### GetLastUpdated

`func (o *CustomPayloadWithLastUpdated) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CustomPayloadWithLastUpdated) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CustomPayloadWithLastUpdated) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CustomPayloadWithLastUpdated) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


