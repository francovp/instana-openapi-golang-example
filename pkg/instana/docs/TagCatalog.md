# TagCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagTree** | [**[]TagTreeLevel**](TagTreeLevel.md) |  | 
**Tags** | [**[]Tag**](Tag.md) |  | 

## Methods

### NewTagCatalog

`func NewTagCatalog(tagTree []TagTreeLevel, tags []Tag, ) *TagCatalog`

NewTagCatalog instantiates a new TagCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCatalogWithDefaults

`func NewTagCatalogWithDefaults() *TagCatalog`

NewTagCatalogWithDefaults instantiates a new TagCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagTree

`func (o *TagCatalog) GetTagTree() []TagTreeLevel`

GetTagTree returns the TagTree field if non-nil, zero value otherwise.

### GetTagTreeOk

`func (o *TagCatalog) GetTagTreeOk() (*[]TagTreeLevel, bool)`

GetTagTreeOk returns a tuple with the TagTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTree

`func (o *TagCatalog) SetTagTree(v []TagTreeLevel)`

SetTagTree sets TagTree field to given value.


### GetTags

`func (o *TagCatalog) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagCatalog) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagCatalog) SetTags(v []Tag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


