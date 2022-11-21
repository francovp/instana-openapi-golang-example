# TestResultMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**TestId** | **string** |  | 
**TestResultId** | Pointer to **string** |  | [optional] 

## Methods

### NewTestResultMetadata

`func NewTestResultMetadata(testId string, ) *TestResultMetadata`

NewTestResultMetadata instantiates a new TestResultMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultMetadataWithDefaults

`func NewTestResultMetadataWithDefaults() *TestResultMetadata`

NewTestResultMetadataWithDefaults instantiates a new TestResultMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *TestResultMetadata) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TestResultMetadata) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TestResultMetadata) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TestResultMetadata) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTestId

`func (o *TestResultMetadata) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *TestResultMetadata) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *TestResultMetadata) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetTestResultId

`func (o *TestResultMetadata) GetTestResultId() string`

GetTestResultId returns the TestResultId field if non-nil, zero value otherwise.

### GetTestResultIdOk

`func (o *TestResultMetadata) GetTestResultIdOk() (*string, bool)`

GetTestResultIdOk returns a tuple with the TestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultId

`func (o *TestResultMetadata) SetTestResultId(v string)`

SetTestResultId sets TestResultId field to given value.

### HasTestResultId

`func (o *TestResultMetadata) HasTestResultId() bool`

HasTestResultId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


