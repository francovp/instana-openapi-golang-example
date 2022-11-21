# ChangeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**Author**](Author.md) |  | 
**ChangeType** | **string** |  | 

## Methods

### NewChangeSummary

`func NewChangeSummary(author Author, changeType string, ) *ChangeSummary`

NewChangeSummary instantiates a new ChangeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeSummaryWithDefaults

`func NewChangeSummaryWithDefaults() *ChangeSummary`

NewChangeSummaryWithDefaults instantiates a new ChangeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ChangeSummary) GetAuthor() Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ChangeSummary) GetAuthorOk() (*Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ChangeSummary) SetAuthor(v Author)`

SetAuthor sets Author field to given value.


### GetChangeType

`func (o *ChangeSummary) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *ChangeSummary) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *ChangeSummary) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


