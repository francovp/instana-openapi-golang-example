# WebsiteTimeBasedSliEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeaconType** | **string** |  | 
**FilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**WebsiteId** | Pointer to **string** |  | [optional] 

## Methods

### NewWebsiteTimeBasedSliEntity

`func NewWebsiteTimeBasedSliEntity(beaconType string, ) *WebsiteTimeBasedSliEntity`

NewWebsiteTimeBasedSliEntity instantiates a new WebsiteTimeBasedSliEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteTimeBasedSliEntityWithDefaults

`func NewWebsiteTimeBasedSliEntityWithDefaults() *WebsiteTimeBasedSliEntity`

NewWebsiteTimeBasedSliEntityWithDefaults instantiates a new WebsiteTimeBasedSliEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeaconType

`func (o *WebsiteTimeBasedSliEntity) GetBeaconType() string`

GetBeaconType returns the BeaconType field if non-nil, zero value otherwise.

### GetBeaconTypeOk

`func (o *WebsiteTimeBasedSliEntity) GetBeaconTypeOk() (*string, bool)`

GetBeaconTypeOk returns a tuple with the BeaconType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconType

`func (o *WebsiteTimeBasedSliEntity) SetBeaconType(v string)`

SetBeaconType sets BeaconType field to given value.


### GetFilterExpression

`func (o *WebsiteTimeBasedSliEntity) GetFilterExpression() TagFilterExpressionElement`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *WebsiteTimeBasedSliEntity) GetFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *WebsiteTimeBasedSliEntity) SetFilterExpression(v TagFilterExpressionElement)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *WebsiteTimeBasedSliEntity) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetWebsiteId

`func (o *WebsiteTimeBasedSliEntity) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *WebsiteTimeBasedSliEntity) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *WebsiteTimeBasedSliEntity) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.

### HasWebsiteId

`func (o *WebsiteTimeBasedSliEntity) HasWebsiteId() bool`

HasWebsiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


