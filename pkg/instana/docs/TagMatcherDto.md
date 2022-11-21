# TagMatcherDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | **string** |  | 
**Key** | **string** |  | 
**Operator** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewTagMatcherDTO

`func NewTagMatcherDTO(entity string, key string, operator string, ) *TagMatcherDTO`

NewTagMatcherDTO instantiates a new TagMatcherDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagMatcherDTOWithDefaults

`func NewTagMatcherDTOWithDefaults() *TagMatcherDTO`

NewTagMatcherDTOWithDefaults instantiates a new TagMatcherDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *TagMatcherDTO) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *TagMatcherDTO) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *TagMatcherDTO) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetKey

`func (o *TagMatcherDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TagMatcherDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TagMatcherDTO) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *TagMatcherDTO) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TagMatcherDTO) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TagMatcherDTO) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *TagMatcherDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagMatcherDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagMatcherDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TagMatcherDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


