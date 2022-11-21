# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]ApplicationScope**](ApplicationScope.md) |  | [optional] 
**Name** | **string** |  | 
**Services** | Pointer to [**[]ServiceScope**](ServiceScope.md) |  | [optional] 
**Start** | **int64** |  | 

## Methods

### NewRelease

`func NewRelease(name string, start int64, ) *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *Release) GetApplications() []ApplicationScope`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *Release) GetApplicationsOk() (*[]ApplicationScope, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *Release) SetApplications(v []ApplicationScope)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *Release) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetName

`func (o *Release) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Release) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Release) SetName(v string)`

SetName sets Name field to given value.


### GetServices

`func (o *Release) GetServices() []ServiceScope`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Release) GetServicesOk() (*[]ServiceScope, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Release) SetServices(v []ServiceScope)`

SetServices sets Services field to given value.

### HasServices

`func (o *Release) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStart

`func (o *Release) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Release) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Release) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


