# BinaryOperatorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conjunction** | **string** |  | 
**Left** | [**MatchExpressionDTO**](MatchExpressionDTO.md) |  | 
**Right** | [**MatchExpressionDTO**](MatchExpressionDTO.md) |  | 

## Methods

### NewBinaryOperatorDTO

`func NewBinaryOperatorDTO(conjunction string, left MatchExpressionDTO, right MatchExpressionDTO, ) *BinaryOperatorDTO`

NewBinaryOperatorDTO instantiates a new BinaryOperatorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryOperatorDTOWithDefaults

`func NewBinaryOperatorDTOWithDefaults() *BinaryOperatorDTO`

NewBinaryOperatorDTOWithDefaults instantiates a new BinaryOperatorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConjunction

`func (o *BinaryOperatorDTO) GetConjunction() string`

GetConjunction returns the Conjunction field if non-nil, zero value otherwise.

### GetConjunctionOk

`func (o *BinaryOperatorDTO) GetConjunctionOk() (*string, bool)`

GetConjunctionOk returns a tuple with the Conjunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjunction

`func (o *BinaryOperatorDTO) SetConjunction(v string)`

SetConjunction sets Conjunction field to given value.


### GetLeft

`func (o *BinaryOperatorDTO) GetLeft() MatchExpressionDTO`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *BinaryOperatorDTO) GetLeftOk() (*MatchExpressionDTO, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *BinaryOperatorDTO) SetLeft(v MatchExpressionDTO)`

SetLeft sets Left field to given value.


### GetRight

`func (o *BinaryOperatorDTO) GetRight() MatchExpressionDTO`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *BinaryOperatorDTO) GetRightOk() (*MatchExpressionDTO, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *BinaryOperatorDTO) SetRight(v MatchExpressionDTO)`

SetRight sets Right field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


