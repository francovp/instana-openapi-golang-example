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

// BuiltInEventSpecification struct for BuiltInEventSpecification
type BuiltInEventSpecification struct {
	Description   *string      `json:"description,omitempty"`
	Enabled       *bool        `json:"enabled,omitempty"`
	Hidden        *bool        `json:"hidden,omitempty"`
	HyperParams   []HyperParam `json:"hyperParams"`
	Id            string       `json:"id"`
	Name          string       `json:"name"`
	RuleInputs    []RuleInput  `json:"ruleInputs"`
	Severity      *int32       `json:"severity,omitempty"`
	ShortPluginId string       `json:"shortPluginId"`
	Triggering    *bool        `json:"triggering,omitempty"`
}

// NewBuiltInEventSpecification instantiates a new BuiltInEventSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuiltInEventSpecification(hyperParams []HyperParam, id string, name string, ruleInputs []RuleInput, shortPluginId string) *BuiltInEventSpecification {
	this := BuiltInEventSpecification{}
	this.HyperParams = hyperParams
	this.Id = id
	this.Name = name
	this.RuleInputs = ruleInputs
	this.ShortPluginId = shortPluginId
	return &this
}

// NewBuiltInEventSpecificationWithDefaults instantiates a new BuiltInEventSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuiltInEventSpecificationWithDefaults() *BuiltInEventSpecification {
	this := BuiltInEventSpecification{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BuiltInEventSpecification) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BuiltInEventSpecification) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BuiltInEventSpecification) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BuiltInEventSpecification) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BuiltInEventSpecification) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BuiltInEventSpecification) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BuiltInEventSpecification) GetHidden() bool {
	if o == nil || isNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetHiddenOk() (*bool, bool) {
	if o == nil || isNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BuiltInEventSpecification) HasHidden() bool {
	if o != nil && !isNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BuiltInEventSpecification) SetHidden(v bool) {
	o.Hidden = &v
}

// GetHyperParams returns the HyperParams field value
func (o *BuiltInEventSpecification) GetHyperParams() []HyperParam {
	if o == nil {
		var ret []HyperParam
		return ret
	}

	return o.HyperParams
}

// GetHyperParamsOk returns a tuple with the HyperParams field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetHyperParamsOk() ([]HyperParam, bool) {
	if o == nil {
		return nil, false
	}
	return o.HyperParams, true
}

// SetHyperParams sets field value
func (o *BuiltInEventSpecification) SetHyperParams(v []HyperParam) {
	o.HyperParams = v
}

// GetId returns the Id field value
func (o *BuiltInEventSpecification) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BuiltInEventSpecification) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BuiltInEventSpecification) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BuiltInEventSpecification) SetName(v string) {
	o.Name = v
}

// GetRuleInputs returns the RuleInputs field value
func (o *BuiltInEventSpecification) GetRuleInputs() []RuleInput {
	if o == nil {
		var ret []RuleInput
		return ret
	}

	return o.RuleInputs
}

// GetRuleInputsOk returns a tuple with the RuleInputs field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetRuleInputsOk() ([]RuleInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleInputs, true
}

// SetRuleInputs sets field value
func (o *BuiltInEventSpecification) SetRuleInputs(v []RuleInput) {
	o.RuleInputs = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *BuiltInEventSpecification) GetSeverity() int32 {
	if o == nil || isNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetSeverityOk() (*int32, bool) {
	if o == nil || isNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *BuiltInEventSpecification) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *BuiltInEventSpecification) SetSeverity(v int32) {
	o.Severity = &v
}

// GetShortPluginId returns the ShortPluginId field value
func (o *BuiltInEventSpecification) GetShortPluginId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortPluginId
}

// GetShortPluginIdOk returns a tuple with the ShortPluginId field value
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetShortPluginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortPluginId, true
}

// SetShortPluginId sets field value
func (o *BuiltInEventSpecification) SetShortPluginId(v string) {
	o.ShortPluginId = v
}

// GetTriggering returns the Triggering field value if set, zero value otherwise.
func (o *BuiltInEventSpecification) GetTriggering() bool {
	if o == nil || isNil(o.Triggering) {
		var ret bool
		return ret
	}
	return *o.Triggering
}

// GetTriggeringOk returns a tuple with the Triggering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltInEventSpecification) GetTriggeringOk() (*bool, bool) {
	if o == nil || isNil(o.Triggering) {
		return nil, false
	}
	return o.Triggering, true
}

// HasTriggering returns a boolean if a field has been set.
func (o *BuiltInEventSpecification) HasTriggering() bool {
	if o != nil && !isNil(o.Triggering) {
		return true
	}

	return false
}

// SetTriggering gets a reference to the given bool and assigns it to the Triggering field.
func (o *BuiltInEventSpecification) SetTriggering(v bool) {
	o.Triggering = &v
}

func (o BuiltInEventSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if true {
		toSerialize["hyperParams"] = o.HyperParams
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ruleInputs"] = o.RuleInputs
	}
	if !isNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["shortPluginId"] = o.ShortPluginId
	}
	if !isNil(o.Triggering) {
		toSerialize["triggering"] = o.Triggering
	}
	return json.Marshal(toSerialize)
}

type NullableBuiltInEventSpecification struct {
	value *BuiltInEventSpecification
	isSet bool
}

func (v NullableBuiltInEventSpecification) Get() *BuiltInEventSpecification {
	return v.value
}

func (v *NullableBuiltInEventSpecification) Set(val *BuiltInEventSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableBuiltInEventSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableBuiltInEventSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuiltInEventSpecification(val *BuiltInEventSpecification) *NullableBuiltInEventSpecification {
	return &NullableBuiltInEventSpecification{value: val, isSet: true}
}

func (v NullableBuiltInEventSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuiltInEventSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
