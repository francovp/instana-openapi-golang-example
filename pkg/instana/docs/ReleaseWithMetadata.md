# ReleaseWithMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]ApplicationScopeWithMetadata**](ApplicationScopeWithMetadata.md) |  | [optional] 
**Id** | **string** |  | 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Scopes** | Pointer to [**[]ReleaseScope**](ReleaseScope.md) |  | [optional] 
**Services** | Pointer to [**[]ServiceScopeWithMetadata**](ServiceScopeWithMetadata.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewReleaseWithMetadata

`func NewReleaseWithMetadata(id string, name string, ) *ReleaseWithMetadata`

NewReleaseWithMetadata instantiates a new ReleaseWithMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithMetadataWithDefaults

`func NewReleaseWithMetadataWithDefaults() *ReleaseWithMetadata`

NewReleaseWithMetadataWithDefaults instantiates a new ReleaseWithMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ReleaseWithMetadata) GetApplications() []ApplicationScopeWithMetadata`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ReleaseWithMetadata) GetApplicationsOk() (*[]ApplicationScopeWithMetadata, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ReleaseWithMetadata) SetApplications(v []ApplicationScopeWithMetadata)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ReleaseWithMetadata) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetId

`func (o *ReleaseWithMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseWithMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseWithMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdated

`func (o *ReleaseWithMetadata) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ReleaseWithMetadata) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ReleaseWithMetadata) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ReleaseWithMetadata) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ReleaseWithMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleaseWithMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleaseWithMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetScopes

`func (o *ReleaseWithMetadata) GetScopes() []ReleaseScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ReleaseWithMetadata) GetScopesOk() (*[]ReleaseScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ReleaseWithMetadata) SetScopes(v []ReleaseScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ReleaseWithMetadata) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetServices

`func (o *ReleaseWithMetadata) GetServices() []ServiceScopeWithMetadata`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ReleaseWithMetadata) GetServicesOk() (*[]ServiceScopeWithMetadata, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ReleaseWithMetadata) SetServices(v []ServiceScopeWithMetadata)`

SetServices sets Services field to given value.

### HasServices

`func (o *ReleaseWithMetadata) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStart

`func (o *ReleaseWithMetadata) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ReleaseWithMetadata) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ReleaseWithMetadata) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ReleaseWithMetadata) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


