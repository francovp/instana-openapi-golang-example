# AccessRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | **string** |  | 
**RelatedId** | Pointer to **string** |  | [optional] 
**RelationType** | **string** |  | 

## Methods

### NewAccessRule

`func NewAccessRule(accessType string, relationType string, ) *AccessRule`

NewAccessRule instantiates a new AccessRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRuleWithDefaults

`func NewAccessRuleWithDefaults() *AccessRule`

NewAccessRuleWithDefaults instantiates a new AccessRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *AccessRule) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AccessRule) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AccessRule) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetRelatedId

`func (o *AccessRule) GetRelatedId() string`

GetRelatedId returns the RelatedId field if non-nil, zero value otherwise.

### GetRelatedIdOk

`func (o *AccessRule) GetRelatedIdOk() (*string, bool)`

GetRelatedIdOk returns a tuple with the RelatedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedId

`func (o *AccessRule) SetRelatedId(v string)`

SetRelatedId sets RelatedId field to given value.

### HasRelatedId

`func (o *AccessRule) HasRelatedId() bool`

HasRelatedId returns a boolean if a field has been set.

### GetRelationType

`func (o *AccessRule) GetRelationType() string`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *AccessRule) GetRelationTypeOk() (*string, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *AccessRule) SetRelationType(v string)`

SetRelationType sets RelationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


