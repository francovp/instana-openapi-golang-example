# PhysicalContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloudfoundry** | Pointer to [**CloudfoundryPhysicalContext**](CloudfoundryPhysicalContext.md) |  | [optional] 
**Cluster** | Pointer to [**SnapshotPreview**](SnapshotPreview.md) |  | [optional] 
**Container** | Pointer to [**SnapshotPreview**](SnapshotPreview.md) |  | [optional] 
**Host** | Pointer to [**SnapshotPreview**](SnapshotPreview.md) |  | [optional] 
**Kubernetes** | Pointer to [**KubernetesPhysicalContext**](KubernetesPhysicalContext.md) |  | [optional] 
**Process** | Pointer to [**SnapshotPreview**](SnapshotPreview.md) |  | [optional] 

## Methods

### NewPhysicalContext

`func NewPhysicalContext() *PhysicalContext`

NewPhysicalContext instantiates a new PhysicalContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalContextWithDefaults

`func NewPhysicalContextWithDefaults() *PhysicalContext`

NewPhysicalContextWithDefaults instantiates a new PhysicalContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudfoundry

`func (o *PhysicalContext) GetCloudfoundry() CloudfoundryPhysicalContext`

GetCloudfoundry returns the Cloudfoundry field if non-nil, zero value otherwise.

### GetCloudfoundryOk

`func (o *PhysicalContext) GetCloudfoundryOk() (*CloudfoundryPhysicalContext, bool)`

GetCloudfoundryOk returns a tuple with the Cloudfoundry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudfoundry

`func (o *PhysicalContext) SetCloudfoundry(v CloudfoundryPhysicalContext)`

SetCloudfoundry sets Cloudfoundry field to given value.

### HasCloudfoundry

`func (o *PhysicalContext) HasCloudfoundry() bool`

HasCloudfoundry returns a boolean if a field has been set.

### GetCluster

`func (o *PhysicalContext) GetCluster() SnapshotPreview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PhysicalContext) GetClusterOk() (*SnapshotPreview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PhysicalContext) SetCluster(v SnapshotPreview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PhysicalContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetContainer

`func (o *PhysicalContext) GetContainer() SnapshotPreview`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *PhysicalContext) GetContainerOk() (*SnapshotPreview, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *PhysicalContext) SetContainer(v SnapshotPreview)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *PhysicalContext) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetHost

`func (o *PhysicalContext) GetHost() SnapshotPreview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PhysicalContext) GetHostOk() (*SnapshotPreview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PhysicalContext) SetHost(v SnapshotPreview)`

SetHost sets Host field to given value.

### HasHost

`func (o *PhysicalContext) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKubernetes

`func (o *PhysicalContext) GetKubernetes() KubernetesPhysicalContext`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *PhysicalContext) GetKubernetesOk() (*KubernetesPhysicalContext, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *PhysicalContext) SetKubernetes(v KubernetesPhysicalContext)`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *PhysicalContext) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

### GetProcess

`func (o *PhysicalContext) GetProcess() SnapshotPreview`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *PhysicalContext) GetProcessOk() (*SnapshotPreview, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *PhysicalContext) SetProcess(v SnapshotPreview)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *PhysicalContext) HasProcess() bool`

HasProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


