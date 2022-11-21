# KubernetesPhysicalContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**SnapshotPreview**](SnapshotPreview.md) |  | [optional] 
**Namespace** | Pointer to [**SnapshotPreview**](SnapshotPreview.md) |  | [optional] 
**Node** | Pointer to [**SnapshotPreview**](SnapshotPreview.md) |  | [optional] 
**Pod** | Pointer to [**SnapshotPreview**](SnapshotPreview.md) |  | [optional] 

## Methods

### NewKubernetesPhysicalContext

`func NewKubernetesPhysicalContext() *KubernetesPhysicalContext`

NewKubernetesPhysicalContext instantiates a new KubernetesPhysicalContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesPhysicalContextWithDefaults

`func NewKubernetesPhysicalContextWithDefaults() *KubernetesPhysicalContext`

NewKubernetesPhysicalContextWithDefaults instantiates a new KubernetesPhysicalContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *KubernetesPhysicalContext) GetCluster() SnapshotPreview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *KubernetesPhysicalContext) GetClusterOk() (*SnapshotPreview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *KubernetesPhysicalContext) SetCluster(v SnapshotPreview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *KubernetesPhysicalContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetNamespace

`func (o *KubernetesPhysicalContext) GetNamespace() SnapshotPreview`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubernetesPhysicalContext) GetNamespaceOk() (*SnapshotPreview, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubernetesPhysicalContext) SetNamespace(v SnapshotPreview)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubernetesPhysicalContext) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNode

`func (o *KubernetesPhysicalContext) GetNode() SnapshotPreview`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *KubernetesPhysicalContext) GetNodeOk() (*SnapshotPreview, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *KubernetesPhysicalContext) SetNode(v SnapshotPreview)`

SetNode sets Node field to given value.

### HasNode

`func (o *KubernetesPhysicalContext) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPod

`func (o *KubernetesPhysicalContext) GetPod() SnapshotPreview`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *KubernetesPhysicalContext) GetPodOk() (*SnapshotPreview, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *KubernetesPhysicalContext) SetPod(v SnapshotPreview)`

SetPod sets Pod field to given value.

### HasPod

`func (o *KubernetesPhysicalContext) HasPod() bool`

HasPod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


