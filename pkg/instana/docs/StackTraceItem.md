# StackTraceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to **string** |  | [optional] 
**Line** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 

## Methods

### NewStackTraceItem

`func NewStackTraceItem() *StackTraceItem`

NewStackTraceItem instantiates a new StackTraceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackTraceItemWithDefaults

`func NewStackTraceItemWithDefaults() *StackTraceItem`

NewStackTraceItemWithDefaults instantiates a new StackTraceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *StackTraceItem) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *StackTraceItem) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *StackTraceItem) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *StackTraceItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetLine

`func (o *StackTraceItem) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *StackTraceItem) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *StackTraceItem) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *StackTraceItem) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetMethod

`func (o *StackTraceItem) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *StackTraceItem) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *StackTraceItem) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *StackTraceItem) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


