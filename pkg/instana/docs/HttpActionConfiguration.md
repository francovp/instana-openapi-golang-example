# HttpActionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInsecure** | Pointer to **bool** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**ExpectExists** | Pointer to **[]string** |  | [optional] 
**ExpectJson** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpectMatch** | Pointer to **string** |  | [optional] 
**ExpectNotEmpty** | Pointer to **[]string** |  | [optional] 
**ExpectStatus** | Pointer to **int32** |  | [optional] 
**FollowRedirect** | Pointer to **bool** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**ValidationString** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpActionConfiguration

`func NewHttpActionConfiguration(url string, ) *HttpActionConfiguration`

NewHttpActionConfiguration instantiates a new HttpActionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpActionConfigurationWithDefaults

`func NewHttpActionConfigurationWithDefaults() *HttpActionConfiguration`

NewHttpActionConfigurationWithDefaults instantiates a new HttpActionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *HttpActionConfiguration) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *HttpActionConfiguration) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *HttpActionConfiguration) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *HttpActionConfiguration) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetBody

`func (o *HttpActionConfiguration) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HttpActionConfiguration) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HttpActionConfiguration) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HttpActionConfiguration) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetExpectExists

`func (o *HttpActionConfiguration) GetExpectExists() []string`

GetExpectExists returns the ExpectExists field if non-nil, zero value otherwise.

### GetExpectExistsOk

`func (o *HttpActionConfiguration) GetExpectExistsOk() (*[]string, bool)`

GetExpectExistsOk returns a tuple with the ExpectExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectExists

`func (o *HttpActionConfiguration) SetExpectExists(v []string)`

SetExpectExists sets ExpectExists field to given value.

### HasExpectExists

`func (o *HttpActionConfiguration) HasExpectExists() bool`

HasExpectExists returns a boolean if a field has been set.

### GetExpectJson

`func (o *HttpActionConfiguration) GetExpectJson() map[string]interface{}`

GetExpectJson returns the ExpectJson field if non-nil, zero value otherwise.

### GetExpectJsonOk

`func (o *HttpActionConfiguration) GetExpectJsonOk() (*map[string]interface{}, bool)`

GetExpectJsonOk returns a tuple with the ExpectJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectJson

`func (o *HttpActionConfiguration) SetExpectJson(v map[string]interface{})`

SetExpectJson sets ExpectJson field to given value.

### HasExpectJson

`func (o *HttpActionConfiguration) HasExpectJson() bool`

HasExpectJson returns a boolean if a field has been set.

### GetExpectMatch

`func (o *HttpActionConfiguration) GetExpectMatch() string`

GetExpectMatch returns the ExpectMatch field if non-nil, zero value otherwise.

### GetExpectMatchOk

`func (o *HttpActionConfiguration) GetExpectMatchOk() (*string, bool)`

GetExpectMatchOk returns a tuple with the ExpectMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectMatch

`func (o *HttpActionConfiguration) SetExpectMatch(v string)`

SetExpectMatch sets ExpectMatch field to given value.

### HasExpectMatch

`func (o *HttpActionConfiguration) HasExpectMatch() bool`

HasExpectMatch returns a boolean if a field has been set.

### GetExpectNotEmpty

`func (o *HttpActionConfiguration) GetExpectNotEmpty() []string`

GetExpectNotEmpty returns the ExpectNotEmpty field if non-nil, zero value otherwise.

### GetExpectNotEmptyOk

`func (o *HttpActionConfiguration) GetExpectNotEmptyOk() (*[]string, bool)`

GetExpectNotEmptyOk returns a tuple with the ExpectNotEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectNotEmpty

`func (o *HttpActionConfiguration) SetExpectNotEmpty(v []string)`

SetExpectNotEmpty sets ExpectNotEmpty field to given value.

### HasExpectNotEmpty

`func (o *HttpActionConfiguration) HasExpectNotEmpty() bool`

HasExpectNotEmpty returns a boolean if a field has been set.

### GetExpectStatus

`func (o *HttpActionConfiguration) GetExpectStatus() int32`

GetExpectStatus returns the ExpectStatus field if non-nil, zero value otherwise.

### GetExpectStatusOk

`func (o *HttpActionConfiguration) GetExpectStatusOk() (*int32, bool)`

GetExpectStatusOk returns a tuple with the ExpectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectStatus

`func (o *HttpActionConfiguration) SetExpectStatus(v int32)`

SetExpectStatus sets ExpectStatus field to given value.

### HasExpectStatus

`func (o *HttpActionConfiguration) HasExpectStatus() bool`

HasExpectStatus returns a boolean if a field has been set.

### GetFollowRedirect

`func (o *HttpActionConfiguration) GetFollowRedirect() bool`

GetFollowRedirect returns the FollowRedirect field if non-nil, zero value otherwise.

### GetFollowRedirectOk

`func (o *HttpActionConfiguration) GetFollowRedirectOk() (*bool, bool)`

GetFollowRedirectOk returns a tuple with the FollowRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirect

`func (o *HttpActionConfiguration) SetFollowRedirect(v bool)`

SetFollowRedirect sets FollowRedirect field to given value.

### HasFollowRedirect

`func (o *HttpActionConfiguration) HasFollowRedirect() bool`

HasFollowRedirect returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpActionConfiguration) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpActionConfiguration) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpActionConfiguration) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpActionConfiguration) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetOperation

`func (o *HttpActionConfiguration) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *HttpActionConfiguration) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *HttpActionConfiguration) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *HttpActionConfiguration) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetUrl

`func (o *HttpActionConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpActionConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpActionConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetValidationString

`func (o *HttpActionConfiguration) GetValidationString() string`

GetValidationString returns the ValidationString field if non-nil, zero value otherwise.

### GetValidationStringOk

`func (o *HttpActionConfiguration) GetValidationStringOk() (*string, bool)`

GetValidationStringOk returns a tuple with the ValidationString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationString

`func (o *HttpActionConfiguration) SetValidationString(v string)`

SetValidationString sets ValidationString field to given value.

### HasValidationString

`func (o *HttpActionConfiguration) HasValidationString() bool`

HasValidationString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


