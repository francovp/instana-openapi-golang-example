# TestResultDetailData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Har** | Pointer to **map[string]interface{}** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**Logs** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SubtransactionAvgMetrics** | Pointer to **map[string]interface{}** |  | [optional] 
**Subtransactions** | Pointer to [**[]TestResultSubtransaction**](TestResultSubtransaction.md) |  | [optional] 
**TestId** | Pointer to **string** |  | [optional] 
**TestResultId** | Pointer to **string** |  | [optional] 
**Videos** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTestResultDetailData

`func NewTestResultDetailData() *TestResultDetailData`

NewTestResultDetailData instantiates a new TestResultDetailData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultDetailDataWithDefaults

`func NewTestResultDetailDataWithDefaults() *TestResultDetailData`

NewTestResultDetailDataWithDefaults instantiates a new TestResultDetailData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHar

`func (o *TestResultDetailData) GetHar() map[string]interface{}`

GetHar returns the Har field if non-nil, zero value otherwise.

### GetHarOk

`func (o *TestResultDetailData) GetHarOk() (*map[string]interface{}, bool)`

GetHarOk returns a tuple with the Har field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHar

`func (o *TestResultDetailData) SetHar(v map[string]interface{})`

SetHar sets Har field to given value.

### HasHar

`func (o *TestResultDetailData) HasHar() bool`

HasHar returns a boolean if a field has been set.

### GetImages

`func (o *TestResultDetailData) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *TestResultDetailData) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *TestResultDetailData) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *TestResultDetailData) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetLogs

`func (o *TestResultDetailData) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *TestResultDetailData) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *TestResultDetailData) SetLogs(v string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *TestResultDetailData) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetName

`func (o *TestResultDetailData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestResultDetailData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestResultDetailData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestResultDetailData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubtransactionAvgMetrics

`func (o *TestResultDetailData) GetSubtransactionAvgMetrics() map[string]interface{}`

GetSubtransactionAvgMetrics returns the SubtransactionAvgMetrics field if non-nil, zero value otherwise.

### GetSubtransactionAvgMetricsOk

`func (o *TestResultDetailData) GetSubtransactionAvgMetricsOk() (*map[string]interface{}, bool)`

GetSubtransactionAvgMetricsOk returns a tuple with the SubtransactionAvgMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtransactionAvgMetrics

`func (o *TestResultDetailData) SetSubtransactionAvgMetrics(v map[string]interface{})`

SetSubtransactionAvgMetrics sets SubtransactionAvgMetrics field to given value.

### HasSubtransactionAvgMetrics

`func (o *TestResultDetailData) HasSubtransactionAvgMetrics() bool`

HasSubtransactionAvgMetrics returns a boolean if a field has been set.

### GetSubtransactions

`func (o *TestResultDetailData) GetSubtransactions() []TestResultSubtransaction`

GetSubtransactions returns the Subtransactions field if non-nil, zero value otherwise.

### GetSubtransactionsOk

`func (o *TestResultDetailData) GetSubtransactionsOk() (*[]TestResultSubtransaction, bool)`

GetSubtransactionsOk returns a tuple with the Subtransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtransactions

`func (o *TestResultDetailData) SetSubtransactions(v []TestResultSubtransaction)`

SetSubtransactions sets Subtransactions field to given value.

### HasSubtransactions

`func (o *TestResultDetailData) HasSubtransactions() bool`

HasSubtransactions returns a boolean if a field has been set.

### GetTestId

`func (o *TestResultDetailData) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *TestResultDetailData) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *TestResultDetailData) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *TestResultDetailData) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTestResultId

`func (o *TestResultDetailData) GetTestResultId() string`

GetTestResultId returns the TestResultId field if non-nil, zero value otherwise.

### GetTestResultIdOk

`func (o *TestResultDetailData) GetTestResultIdOk() (*string, bool)`

GetTestResultIdOk returns a tuple with the TestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultId

`func (o *TestResultDetailData) SetTestResultId(v string)`

SetTestResultId sets TestResultId field to given value.

### HasTestResultId

`func (o *TestResultDetailData) HasTestResultId() bool`

HasTestResultId returns a boolean if a field has been set.

### GetVideos

`func (o *TestResultDetailData) GetVideos() []string`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *TestResultDetailData) GetVideosOk() (*[]string, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *TestResultDetailData) SetVideos(v []string)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *TestResultDetailData) HasVideos() bool`

HasVideos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


