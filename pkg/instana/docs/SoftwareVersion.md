# SoftwareVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Origin** | **string** |  | 
**Type** | **string** |  | 
**UsedBy** | [**[]SoftwareUser**](SoftwareUser.md) |  | 
**Version** | **string** |  | 

## Methods

### NewSoftwareVersion

`func NewSoftwareVersion(name string, origin string, type_ string, usedBy []SoftwareUser, version string, ) *SoftwareVersion`

NewSoftwareVersion instantiates a new SoftwareVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareVersionWithDefaults

`func NewSoftwareVersionWithDefaults() *SoftwareVersion`

NewSoftwareVersionWithDefaults instantiates a new SoftwareVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SoftwareVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwareVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwareVersion) SetName(v string)`

SetName sets Name field to given value.


### GetOrigin

`func (o *SoftwareVersion) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SoftwareVersion) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SoftwareVersion) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetType

`func (o *SoftwareVersion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SoftwareVersion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SoftwareVersion) SetType(v string)`

SetType sets Type field to given value.


### GetUsedBy

`func (o *SoftwareVersion) GetUsedBy() []SoftwareUser`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *SoftwareVersion) GetUsedByOk() (*[]SoftwareUser, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *SoftwareVersion) SetUsedBy(v []SoftwareUser)`

SetUsedBy sets UsedBy field to given value.


### GetVersion

`func (o *SoftwareVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwareVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwareVersion) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


