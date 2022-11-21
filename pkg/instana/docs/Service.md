# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | **[]string** |  | 
**EntityType** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Label** | **string** |  | 
**MaxSeverity** | Pointer to **float64** |  | [optional] 
**NumberOfOpenIssues** | Pointer to **int32** |  | [optional] 
**Technologies** | **[]string** |  | 
**Types** | **[]string** |  | 

## Methods

### NewService

`func NewService(applications []string, id string, label string, technologies []string, types []string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *Service) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *Service) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *Service) SetApplications(v []string)`

SetApplications sets Applications field to given value.


### GetEntityType

`func (o *Service) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Service) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Service) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *Service) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *Service) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Service) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Service) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMaxSeverity

`func (o *Service) GetMaxSeverity() float64`

GetMaxSeverity returns the MaxSeverity field if non-nil, zero value otherwise.

### GetMaxSeverityOk

`func (o *Service) GetMaxSeverityOk() (*float64, bool)`

GetMaxSeverityOk returns a tuple with the MaxSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSeverity

`func (o *Service) SetMaxSeverity(v float64)`

SetMaxSeverity sets MaxSeverity field to given value.

### HasMaxSeverity

`func (o *Service) HasMaxSeverity() bool`

HasMaxSeverity returns a boolean if a field has been set.

### GetNumberOfOpenIssues

`func (o *Service) GetNumberOfOpenIssues() int32`

GetNumberOfOpenIssues returns the NumberOfOpenIssues field if non-nil, zero value otherwise.

### GetNumberOfOpenIssuesOk

`func (o *Service) GetNumberOfOpenIssuesOk() (*int32, bool)`

GetNumberOfOpenIssuesOk returns a tuple with the NumberOfOpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOpenIssues

`func (o *Service) SetNumberOfOpenIssues(v int32)`

SetNumberOfOpenIssues sets NumberOfOpenIssues field to given value.

### HasNumberOfOpenIssues

`func (o *Service) HasNumberOfOpenIssues() bool`

HasNumberOfOpenIssues returns a boolean if a field has been set.

### GetTechnologies

`func (o *Service) GetTechnologies() []string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *Service) GetTechnologiesOk() (*[]string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *Service) SetTechnologies(v []string)`

SetTechnologies sets Technologies field to given value.


### GetTypes

`func (o *Service) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Service) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Service) SetTypes(v []string)`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


