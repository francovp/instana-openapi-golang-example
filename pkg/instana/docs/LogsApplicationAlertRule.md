# LogsApplicationAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** |  | 
**Loglevel** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Operator** | **string** |  | 

## Methods

### NewLogsApplicationAlertRule

`func NewLogsApplicationAlertRule(level string, operator string, ) *LogsApplicationAlertRule`

NewLogsApplicationAlertRule instantiates a new LogsApplicationAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsApplicationAlertRuleWithDefaults

`func NewLogsApplicationAlertRuleWithDefaults() *LogsApplicationAlertRule`

NewLogsApplicationAlertRuleWithDefaults instantiates a new LogsApplicationAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *LogsApplicationAlertRule) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogsApplicationAlertRule) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogsApplicationAlertRule) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetLoglevel

`func (o *LogsApplicationAlertRule) GetLoglevel() string`

GetLoglevel returns the Loglevel field if non-nil, zero value otherwise.

### GetLoglevelOk

`func (o *LogsApplicationAlertRule) GetLoglevelOk() (*string, bool)`

GetLoglevelOk returns a tuple with the Loglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoglevel

`func (o *LogsApplicationAlertRule) SetLoglevel(v string)`

SetLoglevel sets Loglevel field to given value.

### HasLoglevel

`func (o *LogsApplicationAlertRule) HasLoglevel() bool`

HasLoglevel returns a boolean if a field has been set.

### GetMessage

`func (o *LogsApplicationAlertRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogsApplicationAlertRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogsApplicationAlertRule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogsApplicationAlertRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOperator

`func (o *LogsApplicationAlertRule) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LogsApplicationAlertRule) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LogsApplicationAlertRule) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


