/*
Introduction to Instana public APIs

Searching for answers and best pratices? Check our [IBM Instana Community](https://community.ibm.com/community/user/aiops/communities/community-home?CommunityKey=58f324a3-3104-41be-9510-5b7c413cc48f).  ## Agent REST API ### Event SDK REST Web Service Using the Event SDK REST Web Service, it is possible to integrate custom health checks and other event sources into Instana. Each one running the Instana Agent can be used to feed in manual events. The agent has an endpoint which listens on `http://localhost:42699/com.instana.plugin.generic.event` and accepts the following JSON via a POST request:  ```json {     \"title\": <string>,     \"text\": <string>,     \"severity\": <integer> , -1, 5 or 10     \"timestamp\": <integer>, timestamp in milliseconds from epoch     \"duration\": <integer>, duration in milliseconds } ```  *Title* and *text* are used for display purposes.  *Severity* is an optional integer of -1, 5 and 10. A value of -1 or EMPTY will generate a Change. A value of 5 will generate a *warning Issue*, and a value of 10 will generate a *critical Issue*.  When absent, the event is treated as a change without severity. *Timestamp* is the timestamp of the event, but it is optional, in which case the current time is used. *Duration* can be used to mark a timespan for the event. It also is optional, in which case the event will be marked as \"instant\" rather than \"from-to.\"  The endpoint also accepts a batch of events, which then need to be given as an array:  ```json [     {     // event as above     },     {     // event as above     } ] ```  #### Ruby Example  ```ruby duration = (Time.now.to_f * 1000).floor - deploy_start_time_in_ms payload = {} payload[:title] = 'Deployed MyApp' payload[:text] = 'pglombardo deployed MyApp@revision' payload[:duration] = duration  uri = URI('http://localhost:42699/com.instana.plugin.generic.event') req = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json') req.body = payload.to_json Net::HTTP.start(uri.hostname, uri.port) do |http|     http.request(req) end ```  #### Curl Example  ```bash curl -XPOST http://localhost:42699/com.instana.plugin.generic.event -H \"Content-Type: application/json\" -d '{\"title\":\"Custom API Events \", \"text\": \"Failure Redeploying Service Duration\", \"duration\": 5000, \"severity\": -1}' ```  #### PowerShell Example  For Powershell you can either use the standard Cmdlets `Invoke-WebRequest` or `Invoke-RestMethod`. The parameters to be provided are basically the same.  ```bash Invoke-RestMethod     -Uri http://localhost:42699/com.instana.plugin.generic.event     -Method POST     -Body '{\"title\":\"PowerShell Event \", \"text\": \"You used PowerShell to create this event!\", \"duration\": 5000, \"severity\": -1}' ```  ```bash Invoke-WebRequest     -Uri http://localhost:42699/com.instana.plugin.generic.event     -Method Post     -Body '{\"title\":\"PowerShell Event \", \"text\": \"You used PowerShell to create this event!\", \"duration\": 5000, \"severity\": -1}' ``` ### Trace SDK REST Web Service Using the Trace SDK REST Web Service, it is possible to integrate Instana into any application regardless of language. Each active Instana Agent can be used to feed manually captured traces into the Web Service, which can be joined with automatically captured traces or be completely separate. The Agent offers an endpoint that listens on `http://localhost:42699/com.instana.plugin.generic.trace` and accepts the following JSON via a POST request:  ```json {   \"spanId\": <string>,   \"parentId\": <string>,   \"traceId\": <string>,   \"timestamp\": <64 bit long>,   \"duration\": <64 bit long>,   \"name\": <string>,   \"type\": <string>,   \"error\": <boolean>,   \"data\": {     <string> : <string>   } } ```  spanId is the unique identifier for any particular span. The trace is defined by a root span, that is, a span that does not have a parentId. The traceId needs to be identical for all spans that belong to the same trace, and is allowed to be overlapping with a spanId. traceId, spanId and parentId are 64-bit unique values encoded as hex string like b0789916ff8f319f. Spans form a hierarchy by referencing the spanId of the parent as parentId. An example of a span hierarchy in a trace is shown below:  ```text   root (spanId=1, traceId=1, type=Entry)     child A (spanId=2, traceId=1, parentId=1, type=Exit)       child A (spanId=3, traceId=1, parentId=2, type=Entry)         child B (spanId=4, traceId=1, parentId=3, type=Exit)   child B (spanId=5, traceId=1, parentId=4, type=Entry) ```  The timestamp and duration fields are in milliseconds. The timestamp must be the epoch timestamp coordinated to Coordinated Universal Time.  The name field can be any string that is used to visualize and group traces, and can contain any text. However, simplicity is recommended.  The type field is optional, when absent is treated as ENTRY. Options are ENTRY, EXIT, INTERMEDIATE, or EUM. Setting the type is important for the UI. It is assumed that after an ENTRY, child spans are INTERMEDIATE or EXIT. After an EXIT an ENTRY should follow. This is visualized as a remote call.  The data field is optional and can contain arbitrary key-value pairs. The behavior of supplying duplicate keys is unspecified.  The error field is optional and can be set to true to indicate an erroneous span.  The endpoint also accepts a batch of spans, which then need to be given as an array:  ```json [   {     // span as above   },   {     // span as above   } ] ```  For traces received via the Trace SDK Web Service the same rules regarding [Conversion and Service/Endpoint Naming](https://www.ibm.com/docs/en/obi/current?topic=references-java-trace-sdk#conversion-and-naming) are applied as for the Java Trace SDK. In particular these key-value pairs in data are used for naming: service, endpoint and call.name.  Note: The optional [Instana Agent Service](https://www.ibm.com/docs/en/obi/current?topic=requirements-installing-host-agent-kubernetes#instana-agent-service) provided on Kubernetes via the Instana Agent Helm Chart is very useful in combination with the Trace Web SDK support, as it ensures that the data is pushed to the Instana Agent running on the same Kubernetes node, ensuring the Instana Agent can correctly fill in the infrastructure correlation data.  #### Curl Example  The following example shows how to send to the host agent data about a matching ENTRY and EXIT call, which simulates a process that receives an HTTP GET request targeted to the https://orders.happyshop.com/my/service/asdasd URL and routes it to an upstream service at the https://crm.internal/orders/asdasd URL.  ```bash #!/bin/bash  curl -0 -v -X POST 'http://localhost:42699/com.instana.plugin.generic.trace' -H 'Content-Type: application/json; charset=utf-8' -d @- <<EOF [   {     \"spanId\": \"8165b19a37094800\",     \"traceId\": \"1368e0592a91fe00\",     \"timestamp\": 1591346182000,     \"duration\": 134,     \"name\": \"GET /my/service/asdasd\",     \"type\": \"ENTRY\",     \"error\": false,     \"data\": {       \"http.url\": \"https://orders.happyshop.com/my/service/asdasd\",       \"http.method\": \"GET\",       \"http.status_code\": 200,       \"http.path\": \"/my/service/asdasd\",       \"http.host\": \"orders.happyshop.com\"     }   },   {     \"spanId\": \"7ddf6b31b320cc00\",     \"parentId\": \"8165b19a37094800\",     \"traceId\": \"1368e0592a91fe00\",     \"timestamp\": 1591346182010,     \"duration\": 97,     \"name\": \"GET /orders/asdasd\",     \"type\": \"EXIT\",     \"error\": false,     \"data\": {       \"http.url\": \"https://crm.internal/orders/asdasd\",       \"http.method\": \"GET\",       \"http.status_code\": 200,       \"http.path\": \"/orders/asdasd\",       \"http.host\": \"crm.internal\"     }   } ] EOF ```  #### Limitations  Adhere to the following rate limits for the trace web service:  - Maximum API calls/sec: 20   - Maximum payload per POST request: A span must not exceed 4 KiB. The request size must not exceed 4 MiB.   - Maximum batch size (spans/array): 1000    ## Backend REST API The Instana API allows retrieval and configuration of key data points. Among others, this API enables automatic reaction and further analysis of identified incidents as well as reporting capabilities.  The API documentation refers to two crucial parameters that you need to know about before reading further:  - `base`: This is the base URL of a tenant unit, e.g. `https://test-example.instana.io`. This is the same URL that is used to access the Instana user interface. - `apiToken`: Requests against the Instana API require valid API tokens. An initial API token can be generated via the Instana user interface. Any additional API tokens can be generated via the API itself.  ### Example Here is an Example to use the REST API with Curl. First lets get all the available metrics with possible aggregations with a GET call.  ```bash curl --request GET \\   --url https://test-instana.instana.io/api/application-monitoring/catalog/metrics \\   --header 'authorization: apiToken xxxxxxxxxxxxxxxx' ```  Next we can get every call grouped by the endpoint name that has an error count greater then zero. As a metric we could get the mean error rate for example.  ```bash curl --request POST \\   --url https://test-instana.instana.io/api/application-monitoring/analyze/call-groups \\   --header 'authorization: apiToken xxxxxxxxxxxxxxxx' \\   --header 'content-type: application/json' \\   --data '{   \"group\":{       \"groupbyTag\":\"endpoint.name\"   },   \"tagFilters\":[    {     \"name\":\"call.error.count\",     \"value\":\"0\",     \"operator\":\"GREATER_THAN\"    }   ],   \"metrics\":[    {     \"metric\":\"errors\",     \"aggregation\":\"MEAN\"    }   ]   }' ```   ### Rate Limiting A rate limit is applied to API usage. Up to 5,000 calls per hour can be made. How many remaining calls can be made and when this call limit resets, can inspected via three headers that are part of the responses of the API server.  **X-RateLimit-Limit:** Shows the maximum number of calls that may be executed per hour.  **X-RateLimit-Remaining:** How many calls may still be executed within the current hour.  **X-RateLimit-Reset:** Time when the remaining calls will be reset to the limit. For compatibility reasons with other rate limited APIs, this date is not the date in milliseconds, but instead in seconds since 1970-01-01T00:00:00+00:00.  ## Generating REST API clients  The API is specified using the [OpenAPI v3](https://github.com/OAI/OpenAPI-Specification) (previously known as Swagger) format. You can download the current specification at our [GitHub API documentation](https://instana.github.io/openapi/openapi.yaml).  OpenAPI tries to solve the issue of ever-evolving APIs and clients lagging behind. Please make sure that you always use the latest version of the generator, as a number of improvements are regularly made. To generate a client library for your language, you can use the [OpenAPI client generators](https://github.com/OpenAPITools/openapi-generator).  ### Go For example, to generate a client library for Go to interact with our backend, you can use the following script; mind replacing the values of the `UNIT_NAME` and `TENANT_NAME` environment variables using those for your tenant unit:  ```bash #!/bin/bash  ### This script assumes you have the `java` and `wget` commands on the path  export UNIT_NAME='myunit' # for example: prod export TENANT_NAME='mytenant' # for example: awesomecompany  //Download the generator to your current working directory: wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.1/openapi-generator-cli-4.3.1.jar -O openapi-generator-cli.jar --server-variables \"tenant=${TENANT_NAME},unit=${UNIT_NAME}\"  //generate a client library that you can vendor into your repository java -jar openapi-generator-cli.jar generate -i https://instana.github.io/openapi/openapi.yaml -g go \\     -o pkg/instana/openapi \\     --skip-validate-spec  //(optional) format the Go code according to the Go code standard gofmt -s -w pkg/instana/openapi ```  The generated clients contain comprehensive READMEs, and you can start right away using the client from the example above:  ```go import instana \"./pkg/instana/openapi\"  // readTags will read all available application monitoring tags along with their type and category func readTags() {  configuration := instana.NewConfiguration()  configuration.Host = \"tenant-unit.instana.io\"  configuration.BasePath = \"https://tenant-unit.instana.io\"   client := instana.NewAPIClient(configuration)  auth := context.WithValue(context.Background(), instana.ContextAPIKey, instana.APIKey{   Key:    apiKey,   Prefix: \"apiToken\",  })   tags, _, err := client.ApplicationCatalogApi.GetApplicationTagCatalog(auth)  if err != nil {   fmt.Fatalf(\"Error calling the API, aborting.\")  }   for _, tag := range tags {   fmt.Printf(\"%s (%s): %s\\n\", tag.Category, tag.Type, tag.Name)  } } ```  ### Java Download the latest openapi generator cli: ``` wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.1/openapi-generator-cli-4.3.1.jar -O openapi-generator-cli.jar ```  A list for calls for different java http client implementations, which creates a valid generated source code for our spec. ``` //Nativ Java HTTP Client java -jar openapi-generator-cli.jar generate -i https://instana.github.io/openapi/openapi.yaml -g java -o pkg/instana/openapi --skip-validate-spec  -p dateLibrary=java8 --library native  //Spring WebClient java -jar openapi-generator-cli.jar generate -i https://instana.github.io/openapi/openapi.yaml -g java -o pkg/instana/openapi --skip-validate-spec  -p dateLibrary=java8,hideGenerationTimestamp=true --library webclient  //Spring RestTemplate java -jar openapi-generator-cli.jar generate -i https://instana.github.io/openapi/openapi.yaml -g java -o pkg/instana/openapi --skip-validate-spec  -p dateLibrary=java8,hideGenerationTimestamp=true --library resttemplate  ```

API version: 1.236.406
Contact: support@instana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instana

import (
	"encoding/json"
)

// ValidatedAlertingConfiguration struct for ValidatedAlertingConfiguration
type ValidatedAlertingConfiguration struct {
	AlertChannelNames           []string                    `json:"alertChannelNames,omitempty"`
	AlertName                   string                      `json:"alertName"`
	ApplicationNames            []string                    `json:"applicationNames,omitempty"`
	CustomPayloadFields         []StaticStringField         `json:"customPayloadFields"`
	EventFilteringConfiguration EventFilteringConfiguration `json:"eventFilteringConfiguration"`
	Id                          string                      `json:"id"`
	IntegrationIds              []string                    `json:"integrationIds"`
	Invalid                     *bool                       `json:"invalid,omitempty"`
	LastUpdated                 *int64                      `json:"lastUpdated,omitempty"`
	MuteUntil                   *int64                      `json:"muteUntil,omitempty"`
}

// NewValidatedAlertingConfiguration instantiates a new ValidatedAlertingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatedAlertingConfiguration(alertName string, customPayloadFields []StaticStringField, eventFilteringConfiguration EventFilteringConfiguration, id string, integrationIds []string) *ValidatedAlertingConfiguration {
	this := ValidatedAlertingConfiguration{}
	this.AlertName = alertName
	this.CustomPayloadFields = customPayloadFields
	this.EventFilteringConfiguration = eventFilteringConfiguration
	this.Id = id
	this.IntegrationIds = integrationIds
	return &this
}

// NewValidatedAlertingConfigurationWithDefaults instantiates a new ValidatedAlertingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatedAlertingConfigurationWithDefaults() *ValidatedAlertingConfiguration {
	this := ValidatedAlertingConfiguration{}
	return &this
}

// GetAlertChannelNames returns the AlertChannelNames field value if set, zero value otherwise.
func (o *ValidatedAlertingConfiguration) GetAlertChannelNames() []string {
	if o == nil || isNil(o.AlertChannelNames) {
		var ret []string
		return ret
	}
	return o.AlertChannelNames
}

// GetAlertChannelNamesOk returns a tuple with the AlertChannelNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetAlertChannelNamesOk() ([]string, bool) {
	if o == nil || isNil(o.AlertChannelNames) {
		return nil, false
	}
	return o.AlertChannelNames, true
}

// HasAlertChannelNames returns a boolean if a field has been set.
func (o *ValidatedAlertingConfiguration) HasAlertChannelNames() bool {
	if o != nil && !isNil(o.AlertChannelNames) {
		return true
	}

	return false
}

// SetAlertChannelNames gets a reference to the given []string and assigns it to the AlertChannelNames field.
func (o *ValidatedAlertingConfiguration) SetAlertChannelNames(v []string) {
	o.AlertChannelNames = v
}

// GetAlertName returns the AlertName field value
func (o *ValidatedAlertingConfiguration) GetAlertName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetAlertNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertName, true
}

// SetAlertName sets field value
func (o *ValidatedAlertingConfiguration) SetAlertName(v string) {
	o.AlertName = v
}

// GetApplicationNames returns the ApplicationNames field value if set, zero value otherwise.
func (o *ValidatedAlertingConfiguration) GetApplicationNames() []string {
	if o == nil || isNil(o.ApplicationNames) {
		var ret []string
		return ret
	}
	return o.ApplicationNames
}

// GetApplicationNamesOk returns a tuple with the ApplicationNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetApplicationNamesOk() ([]string, bool) {
	if o == nil || isNil(o.ApplicationNames) {
		return nil, false
	}
	return o.ApplicationNames, true
}

// HasApplicationNames returns a boolean if a field has been set.
func (o *ValidatedAlertingConfiguration) HasApplicationNames() bool {
	if o != nil && !isNil(o.ApplicationNames) {
		return true
	}

	return false
}

// SetApplicationNames gets a reference to the given []string and assigns it to the ApplicationNames field.
func (o *ValidatedAlertingConfiguration) SetApplicationNames(v []string) {
	o.ApplicationNames = v
}

// GetCustomPayloadFields returns the CustomPayloadFields field value
func (o *ValidatedAlertingConfiguration) GetCustomPayloadFields() []StaticStringField {
	if o == nil {
		var ret []StaticStringField
		return ret
	}

	return o.CustomPayloadFields
}

// GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field value
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetCustomPayloadFieldsOk() ([]StaticStringField, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomPayloadFields, true
}

// SetCustomPayloadFields sets field value
func (o *ValidatedAlertingConfiguration) SetCustomPayloadFields(v []StaticStringField) {
	o.CustomPayloadFields = v
}

// GetEventFilteringConfiguration returns the EventFilteringConfiguration field value
func (o *ValidatedAlertingConfiguration) GetEventFilteringConfiguration() EventFilteringConfiguration {
	if o == nil {
		var ret EventFilteringConfiguration
		return ret
	}

	return o.EventFilteringConfiguration
}

// GetEventFilteringConfigurationOk returns a tuple with the EventFilteringConfiguration field value
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetEventFilteringConfigurationOk() (*EventFilteringConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventFilteringConfiguration, true
}

// SetEventFilteringConfiguration sets field value
func (o *ValidatedAlertingConfiguration) SetEventFilteringConfiguration(v EventFilteringConfiguration) {
	o.EventFilteringConfiguration = v
}

// GetId returns the Id field value
func (o *ValidatedAlertingConfiguration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ValidatedAlertingConfiguration) SetId(v string) {
	o.Id = v
}

// GetIntegrationIds returns the IntegrationIds field value
func (o *ValidatedAlertingConfiguration) GetIntegrationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IntegrationIds
}

// GetIntegrationIdsOk returns a tuple with the IntegrationIds field value
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetIntegrationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationIds, true
}

// SetIntegrationIds sets field value
func (o *ValidatedAlertingConfiguration) SetIntegrationIds(v []string) {
	o.IntegrationIds = v
}

// GetInvalid returns the Invalid field value if set, zero value otherwise.
func (o *ValidatedAlertingConfiguration) GetInvalid() bool {
	if o == nil || isNil(o.Invalid) {
		var ret bool
		return ret
	}
	return *o.Invalid
}

// GetInvalidOk returns a tuple with the Invalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetInvalidOk() (*bool, bool) {
	if o == nil || isNil(o.Invalid) {
		return nil, false
	}
	return o.Invalid, true
}

// HasInvalid returns a boolean if a field has been set.
func (o *ValidatedAlertingConfiguration) HasInvalid() bool {
	if o != nil && !isNil(o.Invalid) {
		return true
	}

	return false
}

// SetInvalid gets a reference to the given bool and assigns it to the Invalid field.
func (o *ValidatedAlertingConfiguration) SetInvalid(v bool) {
	o.Invalid = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ValidatedAlertingConfiguration) GetLastUpdated() int64 {
	if o == nil || isNil(o.LastUpdated) {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || isNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ValidatedAlertingConfiguration) HasLastUpdated() bool {
	if o != nil && !isNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *ValidatedAlertingConfiguration) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

// GetMuteUntil returns the MuteUntil field value if set, zero value otherwise.
func (o *ValidatedAlertingConfiguration) GetMuteUntil() int64 {
	if o == nil || isNil(o.MuteUntil) {
		var ret int64
		return ret
	}
	return *o.MuteUntil
}

// GetMuteUntilOk returns a tuple with the MuteUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedAlertingConfiguration) GetMuteUntilOk() (*int64, bool) {
	if o == nil || isNil(o.MuteUntil) {
		return nil, false
	}
	return o.MuteUntil, true
}

// HasMuteUntil returns a boolean if a field has been set.
func (o *ValidatedAlertingConfiguration) HasMuteUntil() bool {
	if o != nil && !isNil(o.MuteUntil) {
		return true
	}

	return false
}

// SetMuteUntil gets a reference to the given int64 and assigns it to the MuteUntil field.
func (o *ValidatedAlertingConfiguration) SetMuteUntil(v int64) {
	o.MuteUntil = &v
}

func (o ValidatedAlertingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AlertChannelNames) {
		toSerialize["alertChannelNames"] = o.AlertChannelNames
	}
	if true {
		toSerialize["alertName"] = o.AlertName
	}
	if !isNil(o.ApplicationNames) {
		toSerialize["applicationNames"] = o.ApplicationNames
	}
	if true {
		toSerialize["customPayloadFields"] = o.CustomPayloadFields
	}
	if true {
		toSerialize["eventFilteringConfiguration"] = o.EventFilteringConfiguration
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["integrationIds"] = o.IntegrationIds
	}
	if !isNil(o.Invalid) {
		toSerialize["invalid"] = o.Invalid
	}
	if !isNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !isNil(o.MuteUntil) {
		toSerialize["muteUntil"] = o.MuteUntil
	}
	return json.Marshal(toSerialize)
}

type NullableValidatedAlertingConfiguration struct {
	value *ValidatedAlertingConfiguration
	isSet bool
}

func (v NullableValidatedAlertingConfiguration) Get() *ValidatedAlertingConfiguration {
	return v.value
}

func (v *NullableValidatedAlertingConfiguration) Set(val *ValidatedAlertingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatedAlertingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatedAlertingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatedAlertingConfiguration(val *ValidatedAlertingConfiguration) *NullableValidatedAlertingConfiguration {
	return &NullableValidatedAlertingConfiguration{value: val, isSet: true}
}

func (v NullableValidatedAlertingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatedAlertingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
