# HttpActionConfigurationAllOf

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
**Url** | Pointer to **string** |  | [optional] 
**ValidationString** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpActionConfigurationAllOf

`func NewHttpActionConfigurationAllOf() *HttpActionConfigurationAllOf`

NewHttpActionConfigurationAllOf instantiates a new HttpActionConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpActionConfigurationAllOfWithDefaults

`func NewHttpActionConfigurationAllOfWithDefaults() *HttpActionConfigurationAllOf`

NewHttpActionConfigurationAllOfWithDefaults instantiates a new HttpActionConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *HttpActionConfigurationAllOf) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *HttpActionConfigurationAllOf) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *HttpActionConfigurationAllOf) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *HttpActionConfigurationAllOf) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetBody

`func (o *HttpActionConfigurationAllOf) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HttpActionConfigurationAllOf) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HttpActionConfigurationAllOf) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HttpActionConfigurationAllOf) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetExpectExists

`func (o *HttpActionConfigurationAllOf) GetExpectExists() []string`

GetExpectExists returns the ExpectExists field if non-nil, zero value otherwise.

### GetExpectExistsOk

`func (o *HttpActionConfigurationAllOf) GetExpectExistsOk() (*[]string, bool)`

GetExpectExistsOk returns a tuple with the ExpectExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectExists

`func (o *HttpActionConfigurationAllOf) SetExpectExists(v []string)`

SetExpectExists sets ExpectExists field to given value.

### HasExpectExists

`func (o *HttpActionConfigurationAllOf) HasExpectExists() bool`

HasExpectExists returns a boolean if a field has been set.

### GetExpectJson

`func (o *HttpActionConfigurationAllOf) GetExpectJson() map[string]interface{}`

GetExpectJson returns the ExpectJson field if non-nil, zero value otherwise.

### GetExpectJsonOk

`func (o *HttpActionConfigurationAllOf) GetExpectJsonOk() (*map[string]interface{}, bool)`

GetExpectJsonOk returns a tuple with the ExpectJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectJson

`func (o *HttpActionConfigurationAllOf) SetExpectJson(v map[string]interface{})`

SetExpectJson sets ExpectJson field to given value.

### HasExpectJson

`func (o *HttpActionConfigurationAllOf) HasExpectJson() bool`

HasExpectJson returns a boolean if a field has been set.

### GetExpectMatch

`func (o *HttpActionConfigurationAllOf) GetExpectMatch() string`

GetExpectMatch returns the ExpectMatch field if non-nil, zero value otherwise.

### GetExpectMatchOk

`func (o *HttpActionConfigurationAllOf) GetExpectMatchOk() (*string, bool)`

GetExpectMatchOk returns a tuple with the ExpectMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectMatch

`func (o *HttpActionConfigurationAllOf) SetExpectMatch(v string)`

SetExpectMatch sets ExpectMatch field to given value.

### HasExpectMatch

`func (o *HttpActionConfigurationAllOf) HasExpectMatch() bool`

HasExpectMatch returns a boolean if a field has been set.

### GetExpectNotEmpty

`func (o *HttpActionConfigurationAllOf) GetExpectNotEmpty() []string`

GetExpectNotEmpty returns the ExpectNotEmpty field if non-nil, zero value otherwise.

### GetExpectNotEmptyOk

`func (o *HttpActionConfigurationAllOf) GetExpectNotEmptyOk() (*[]string, bool)`

GetExpectNotEmptyOk returns a tuple with the ExpectNotEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectNotEmpty

`func (o *HttpActionConfigurationAllOf) SetExpectNotEmpty(v []string)`

SetExpectNotEmpty sets ExpectNotEmpty field to given value.

### HasExpectNotEmpty

`func (o *HttpActionConfigurationAllOf) HasExpectNotEmpty() bool`

HasExpectNotEmpty returns a boolean if a field has been set.

### GetExpectStatus

`func (o *HttpActionConfigurationAllOf) GetExpectStatus() int32`

GetExpectStatus returns the ExpectStatus field if non-nil, zero value otherwise.

### GetExpectStatusOk

`func (o *HttpActionConfigurationAllOf) GetExpectStatusOk() (*int32, bool)`

GetExpectStatusOk returns a tuple with the ExpectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectStatus

`func (o *HttpActionConfigurationAllOf) SetExpectStatus(v int32)`

SetExpectStatus sets ExpectStatus field to given value.

### HasExpectStatus

`func (o *HttpActionConfigurationAllOf) HasExpectStatus() bool`

HasExpectStatus returns a boolean if a field has been set.

### GetFollowRedirect

`func (o *HttpActionConfigurationAllOf) GetFollowRedirect() bool`

GetFollowRedirect returns the FollowRedirect field if non-nil, zero value otherwise.

### GetFollowRedirectOk

`func (o *HttpActionConfigurationAllOf) GetFollowRedirectOk() (*bool, bool)`

GetFollowRedirectOk returns a tuple with the FollowRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirect

`func (o *HttpActionConfigurationAllOf) SetFollowRedirect(v bool)`

SetFollowRedirect sets FollowRedirect field to given value.

### HasFollowRedirect

`func (o *HttpActionConfigurationAllOf) HasFollowRedirect() bool`

HasFollowRedirect returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpActionConfigurationAllOf) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpActionConfigurationAllOf) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpActionConfigurationAllOf) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpActionConfigurationAllOf) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetOperation

`func (o *HttpActionConfigurationAllOf) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *HttpActionConfigurationAllOf) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *HttpActionConfigurationAllOf) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *HttpActionConfigurationAllOf) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetUrl

`func (o *HttpActionConfigurationAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpActionConfigurationAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpActionConfigurationAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpActionConfigurationAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetValidationString

`func (o *HttpActionConfigurationAllOf) GetValidationString() string`

GetValidationString returns the ValidationString field if non-nil, zero value otherwise.

### GetValidationStringOk

`func (o *HttpActionConfigurationAllOf) GetValidationStringOk() (*string, bool)`

GetValidationStringOk returns a tuple with the ValidationString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationString

`func (o *HttpActionConfigurationAllOf) SetValidationString(v string)`

SetValidationString sets ValidationString field to given value.

### HasValidationString

`func (o *HttpActionConfigurationAllOf) HasValidationString() bool`

HasValidationString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


