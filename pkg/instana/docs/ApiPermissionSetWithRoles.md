# ApiPermissionSetWithRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | [**[]ScopeBinding**](ScopeBinding.md) |  | 
**InfraDfqFilter** | Pointer to [**ScopeBinding**](ScopeBinding.md) |  | [optional] 
**KubernetesClusterUUIDs** | [**[]ScopeBinding**](ScopeBinding.md) |  | 
**KubernetesNamespaceUIDs** | [**[]ScopeBinding**](ScopeBinding.md) |  | 
**MobileAppIds** | [**[]ScopeBinding**](ScopeBinding.md) |  | 
**Permissions** | **[]string** |  | 
**WebsiteIds** | [**[]ScopeBinding**](ScopeBinding.md) |  | 

## Methods

### NewApiPermissionSetWithRoles

`func NewApiPermissionSetWithRoles(applicationIds []ScopeBinding, kubernetesClusterUUIDs []ScopeBinding, kubernetesNamespaceUIDs []ScopeBinding, mobileAppIds []ScopeBinding, permissions []string, websiteIds []ScopeBinding, ) *ApiPermissionSetWithRoles`

NewApiPermissionSetWithRoles instantiates a new ApiPermissionSetWithRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPermissionSetWithRolesWithDefaults

`func NewApiPermissionSetWithRolesWithDefaults() *ApiPermissionSetWithRoles`

NewApiPermissionSetWithRolesWithDefaults instantiates a new ApiPermissionSetWithRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIds

`func (o *ApiPermissionSetWithRoles) GetApplicationIds() []ScopeBinding`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *ApiPermissionSetWithRoles) GetApplicationIdsOk() (*[]ScopeBinding, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *ApiPermissionSetWithRoles) SetApplicationIds(v []ScopeBinding)`

SetApplicationIds sets ApplicationIds field to given value.


### GetInfraDfqFilter

`func (o *ApiPermissionSetWithRoles) GetInfraDfqFilter() ScopeBinding`

GetInfraDfqFilter returns the InfraDfqFilter field if non-nil, zero value otherwise.

### GetInfraDfqFilterOk

`func (o *ApiPermissionSetWithRoles) GetInfraDfqFilterOk() (*ScopeBinding, bool)`

GetInfraDfqFilterOk returns a tuple with the InfraDfqFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraDfqFilter

`func (o *ApiPermissionSetWithRoles) SetInfraDfqFilter(v ScopeBinding)`

SetInfraDfqFilter sets InfraDfqFilter field to given value.

### HasInfraDfqFilter

`func (o *ApiPermissionSetWithRoles) HasInfraDfqFilter() bool`

HasInfraDfqFilter returns a boolean if a field has been set.

### GetKubernetesClusterUUIDs

`func (o *ApiPermissionSetWithRoles) GetKubernetesClusterUUIDs() []ScopeBinding`

GetKubernetesClusterUUIDs returns the KubernetesClusterUUIDs field if non-nil, zero value otherwise.

### GetKubernetesClusterUUIDsOk

`func (o *ApiPermissionSetWithRoles) GetKubernetesClusterUUIDsOk() (*[]ScopeBinding, bool)`

GetKubernetesClusterUUIDsOk returns a tuple with the KubernetesClusterUUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterUUIDs

`func (o *ApiPermissionSetWithRoles) SetKubernetesClusterUUIDs(v []ScopeBinding)`

SetKubernetesClusterUUIDs sets KubernetesClusterUUIDs field to given value.


### GetKubernetesNamespaceUIDs

`func (o *ApiPermissionSetWithRoles) GetKubernetesNamespaceUIDs() []ScopeBinding`

GetKubernetesNamespaceUIDs returns the KubernetesNamespaceUIDs field if non-nil, zero value otherwise.

### GetKubernetesNamespaceUIDsOk

`func (o *ApiPermissionSetWithRoles) GetKubernetesNamespaceUIDsOk() (*[]ScopeBinding, bool)`

GetKubernetesNamespaceUIDsOk returns a tuple with the KubernetesNamespaceUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespaceUIDs

`func (o *ApiPermissionSetWithRoles) SetKubernetesNamespaceUIDs(v []ScopeBinding)`

SetKubernetesNamespaceUIDs sets KubernetesNamespaceUIDs field to given value.


### GetMobileAppIds

`func (o *ApiPermissionSetWithRoles) GetMobileAppIds() []ScopeBinding`

GetMobileAppIds returns the MobileAppIds field if non-nil, zero value otherwise.

### GetMobileAppIdsOk

`func (o *ApiPermissionSetWithRoles) GetMobileAppIdsOk() (*[]ScopeBinding, bool)`

GetMobileAppIdsOk returns a tuple with the MobileAppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppIds

`func (o *ApiPermissionSetWithRoles) SetMobileAppIds(v []ScopeBinding)`

SetMobileAppIds sets MobileAppIds field to given value.


### GetPermissions

`func (o *ApiPermissionSetWithRoles) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiPermissionSetWithRoles) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiPermissionSetWithRoles) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetWebsiteIds

`func (o *ApiPermissionSetWithRoles) GetWebsiteIds() []ScopeBinding`

GetWebsiteIds returns the WebsiteIds field if non-nil, zero value otherwise.

### GetWebsiteIdsOk

`func (o *ApiPermissionSetWithRoles) GetWebsiteIdsOk() (*[]ScopeBinding, bool)`

GetWebsiteIdsOk returns a tuple with the WebsiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteIds

`func (o *ApiPermissionSetWithRoles) SetWebsiteIds(v []ScopeBinding)`

SetWebsiteIds sets WebsiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


