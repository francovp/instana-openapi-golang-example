# Threshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewThreshold

`func NewThreshold(operator string, type_ string, ) *Threshold`

NewThreshold instantiates a new Threshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdWithDefaults

`func NewThresholdWithDefaults() *Threshold`

NewThresholdWithDefaults instantiates a new Threshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *Threshold) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Threshold) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Threshold) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetType

`func (o *Threshold) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Threshold) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Threshold) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


