# CustomDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRules** | [**[]AccessRule**](AccessRule.md) |  | 
**Id** | **string** |  | 
**Title** | **string** |  | 
**Widgets** | [**[]Widget**](Widget.md) |  | 

## Methods

### NewCustomDashboard

`func NewCustomDashboard(accessRules []AccessRule, id string, title string, widgets []Widget, ) *CustomDashboard`

NewCustomDashboard instantiates a new CustomDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDashboardWithDefaults

`func NewCustomDashboardWithDefaults() *CustomDashboard`

NewCustomDashboardWithDefaults instantiates a new CustomDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRules

`func (o *CustomDashboard) GetAccessRules() []AccessRule`

GetAccessRules returns the AccessRules field if non-nil, zero value otherwise.

### GetAccessRulesOk

`func (o *CustomDashboard) GetAccessRulesOk() (*[]AccessRule, bool)`

GetAccessRulesOk returns a tuple with the AccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRules

`func (o *CustomDashboard) SetAccessRules(v []AccessRule)`

SetAccessRules sets AccessRules field to given value.


### GetId

`func (o *CustomDashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDashboard) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *CustomDashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomDashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomDashboard) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWidgets

`func (o *CustomDashboard) GetWidgets() []Widget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *CustomDashboard) GetWidgetsOk() (*[]Widget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *CustomDashboard) SetWidgets(v []Widget)`

SetWidgets sets Widgets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


