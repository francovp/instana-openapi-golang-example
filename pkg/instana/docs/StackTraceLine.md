# StackTraceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **int32** |  | [optional] 
**File** | **string** |  | 
**Line** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TranslationExplanation** | Pointer to **string** |  | [optional] 
**TranslationStatus** | Pointer to **int32** |  | [optional] 

## Methods

### NewStackTraceLine

`func NewStackTraceLine(file string, ) *StackTraceLine`

NewStackTraceLine instantiates a new StackTraceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackTraceLineWithDefaults

`func NewStackTraceLineWithDefaults() *StackTraceLine`

NewStackTraceLineWithDefaults instantiates a new StackTraceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *StackTraceLine) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *StackTraceLine) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *StackTraceLine) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *StackTraceLine) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetFile

`func (o *StackTraceLine) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *StackTraceLine) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *StackTraceLine) SetFile(v string)`

SetFile sets File field to given value.


### GetLine

`func (o *StackTraceLine) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *StackTraceLine) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *StackTraceLine) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *StackTraceLine) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetName

`func (o *StackTraceLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackTraceLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackTraceLine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StackTraceLine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTranslationExplanation

`func (o *StackTraceLine) GetTranslationExplanation() string`

GetTranslationExplanation returns the TranslationExplanation field if non-nil, zero value otherwise.

### GetTranslationExplanationOk

`func (o *StackTraceLine) GetTranslationExplanationOk() (*string, bool)`

GetTranslationExplanationOk returns a tuple with the TranslationExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationExplanation

`func (o *StackTraceLine) SetTranslationExplanation(v string)`

SetTranslationExplanation sets TranslationExplanation field to given value.

### HasTranslationExplanation

`func (o *StackTraceLine) HasTranslationExplanation() bool`

HasTranslationExplanation returns a boolean if a field has been set.

### GetTranslationStatus

`func (o *StackTraceLine) GetTranslationStatus() int32`

GetTranslationStatus returns the TranslationStatus field if non-nil, zero value otherwise.

### GetTranslationStatusOk

`func (o *StackTraceLine) GetTranslationStatusOk() (*int32, bool)`

GetTranslationStatusOk returns a tuple with the TranslationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationStatus

`func (o *StackTraceLine) SetTranslationStatus(v int32)`

SetTranslationStatus sets TranslationStatus field to given value.

### HasTranslationStatus

`func (o *StackTraceLine) HasTranslationStatus() bool`

HasTranslationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


