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

// GlobalApplicationAlertConfigWithMetadata struct for GlobalApplicationAlertConfigWithMetadata
type GlobalApplicationAlertConfigWithMetadata struct {
	AlertChannelIds     []string                   `json:"alertChannelIds"`
	ApplicationIds      []string                   `json:"applicationIds,omitempty"`
	Applications        map[string]ApplicationNode `json:"applications"`
	BoundaryScope       string                     `json:"boundaryScope"`
	BuiltIn             *bool                      `json:"builtIn,omitempty"`
	Created             *int64                     `json:"created,omitempty"`
	CustomPayloadFields []CustomPayloadField       `json:"customPayloadFields"`
	Description         string                     `json:"description"`
	Enabled             *bool                      `json:"enabled,omitempty"`
	EvaluationType      string                     `json:"evaluationType"`
	Granularity         int32                      `json:"granularity"`
	Id                  string                     `json:"id"`
	IncludeInternal     *bool                      `json:"includeInternal,omitempty"`
	IncludeSynthetic    *bool                      `json:"includeSynthetic,omitempty"`
	Name                string                     `json:"name"`
	ReadOnly            *bool                      `json:"readOnly,omitempty"`
	Rule                ApplicationAlertRule       `json:"rule"`
	Severity            *int32                     `json:"severity,omitempty"`
	TagFilterExpression TagFilterExpressionElement `json:"tagFilterExpression"`
	Threshold           Threshold                  `json:"threshold"`
	TimeThreshold       ApplicationTimeThreshold   `json:"timeThreshold"`
	Triggering          *bool                      `json:"triggering,omitempty"`
}

// NewGlobalApplicationAlertConfigWithMetadata instantiates a new GlobalApplicationAlertConfigWithMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalApplicationAlertConfigWithMetadata(alertChannelIds []string, applications map[string]ApplicationNode, boundaryScope string, customPayloadFields []CustomPayloadField, description string, evaluationType string, granularity int32, id string, name string, rule ApplicationAlertRule, tagFilterExpression TagFilterExpressionElement, threshold Threshold, timeThreshold ApplicationTimeThreshold) *GlobalApplicationAlertConfigWithMetadata {
	this := GlobalApplicationAlertConfigWithMetadata{}
	this.AlertChannelIds = alertChannelIds
	this.Applications = applications
	this.BoundaryScope = boundaryScope
	this.CustomPayloadFields = customPayloadFields
	this.Description = description
	this.EvaluationType = evaluationType
	this.Granularity = granularity
	this.Id = id
	this.Name = name
	this.Rule = rule
	this.TagFilterExpression = tagFilterExpression
	this.Threshold = threshold
	this.TimeThreshold = timeThreshold
	return &this
}

// NewGlobalApplicationAlertConfigWithMetadataWithDefaults instantiates a new GlobalApplicationAlertConfigWithMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalApplicationAlertConfigWithMetadataWithDefaults() *GlobalApplicationAlertConfigWithMetadata {
	this := GlobalApplicationAlertConfigWithMetadata{}
	var granularity int32 = 600000
	this.Granularity = granularity
	return &this
}

// GetAlertChannelIds returns the AlertChannelIds field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetAlertChannelIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlertChannelIds
}

// GetAlertChannelIdsOk returns a tuple with the AlertChannelIds field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetAlertChannelIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertChannelIds, true
}

// SetAlertChannelIds sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetAlertChannelIds(v []string) {
	o.AlertChannelIds = v
}

// GetApplicationIds returns the ApplicationIds field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetApplicationIds() []string {
	if o == nil || isNil(o.ApplicationIds) {
		var ret []string
		return ret
	}
	return o.ApplicationIds
}

// GetApplicationIdsOk returns a tuple with the ApplicationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetApplicationIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ApplicationIds) {
		return nil, false
	}
	return o.ApplicationIds, true
}

// HasApplicationIds returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasApplicationIds() bool {
	if o != nil && !isNil(o.ApplicationIds) {
		return true
	}

	return false
}

// SetApplicationIds gets a reference to the given []string and assigns it to the ApplicationIds field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetApplicationIds(v []string) {
	o.ApplicationIds = v
}

// GetApplications returns the Applications field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetApplications() map[string]ApplicationNode {
	if o == nil {
		var ret map[string]ApplicationNode
		return ret
	}

	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetApplicationsOk() (*map[string]ApplicationNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Applications, true
}

// SetApplications sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetApplications(v map[string]ApplicationNode) {
	o.Applications = v
}

// GetBoundaryScope returns the BoundaryScope field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetBoundaryScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoundaryScope
}

// GetBoundaryScopeOk returns a tuple with the BoundaryScope field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetBoundaryScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundaryScope, true
}

// SetBoundaryScope sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetBoundaryScope(v string) {
	o.BoundaryScope = v
}

// GetBuiltIn returns the BuiltIn field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetBuiltIn() bool {
	if o == nil || isNil(o.BuiltIn) {
		var ret bool
		return ret
	}
	return *o.BuiltIn
}

// GetBuiltInOk returns a tuple with the BuiltIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetBuiltInOk() (*bool, bool) {
	if o == nil || isNil(o.BuiltIn) {
		return nil, false
	}
	return o.BuiltIn, true
}

// HasBuiltIn returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasBuiltIn() bool {
	if o != nil && !isNil(o.BuiltIn) {
		return true
	}

	return false
}

// SetBuiltIn gets a reference to the given bool and assigns it to the BuiltIn field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetBuiltIn(v bool) {
	o.BuiltIn = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetCreated() int64 {
	if o == nil || isNil(o.Created) {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetCreatedOk() (*int64, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetCreated(v int64) {
	o.Created = &v
}

// GetCustomPayloadFields returns the CustomPayloadFields field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetCustomPayloadFields() []CustomPayloadField {
	if o == nil {
		var ret []CustomPayloadField
		return ret
	}

	return o.CustomPayloadFields
}

// GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetCustomPayloadFieldsOk() ([]CustomPayloadField, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomPayloadFields, true
}

// SetCustomPayloadFields sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetCustomPayloadFields(v []CustomPayloadField) {
	o.CustomPayloadFields = v
}

// GetDescription returns the Description field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvaluationType returns the EvaluationType field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetEvaluationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationType
}

// GetEvaluationTypeOk returns a tuple with the EvaluationType field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetEvaluationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationType, true
}

// SetEvaluationType sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetEvaluationType(v string) {
	o.EvaluationType = v
}

// GetGranularity returns the Granularity field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetGranularity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetGranularityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetGranularity(v int32) {
	o.Granularity = v
}

// GetId returns the Id field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetId(v string) {
	o.Id = v
}

// GetIncludeInternal returns the IncludeInternal field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetIncludeInternal() bool {
	if o == nil || isNil(o.IncludeInternal) {
		var ret bool
		return ret
	}
	return *o.IncludeInternal
}

// GetIncludeInternalOk returns a tuple with the IncludeInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetIncludeInternalOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeInternal) {
		return nil, false
	}
	return o.IncludeInternal, true
}

// HasIncludeInternal returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasIncludeInternal() bool {
	if o != nil && !isNil(o.IncludeInternal) {
		return true
	}

	return false
}

// SetIncludeInternal gets a reference to the given bool and assigns it to the IncludeInternal field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetIncludeInternal(v bool) {
	o.IncludeInternal = &v
}

// GetIncludeSynthetic returns the IncludeSynthetic field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetIncludeSynthetic() bool {
	if o == nil || isNil(o.IncludeSynthetic) {
		var ret bool
		return ret
	}
	return *o.IncludeSynthetic
}

// GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetIncludeSyntheticOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeSynthetic) {
		return nil, false
	}
	return o.IncludeSynthetic, true
}

// HasIncludeSynthetic returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasIncludeSynthetic() bool {
	if o != nil && !isNil(o.IncludeSynthetic) {
		return true
	}

	return false
}

// SetIncludeSynthetic gets a reference to the given bool and assigns it to the IncludeSynthetic field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetIncludeSynthetic(v bool) {
	o.IncludeSynthetic = &v
}

// GetName returns the Name field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetName(v string) {
	o.Name = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetReadOnly() bool {
	if o == nil || isNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetReadOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasReadOnly() bool {
	if o != nil && !isNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRule returns the Rule field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetRule() ApplicationAlertRule {
	if o == nil {
		var ret ApplicationAlertRule
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetRuleOk() (*ApplicationAlertRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetRule(v ApplicationAlertRule) {
	o.Rule = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetSeverity() int32 {
	if o == nil || isNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetSeverityOk() (*int32, bool) {
	if o == nil || isNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetSeverity(v int32) {
	o.Severity = &v
}

// GetTagFilterExpression returns the TagFilterExpression field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetTagFilterExpression() TagFilterExpressionElement {
	if o == nil {
		var ret TagFilterExpressionElement
		return ret
	}

	return o.TagFilterExpression
}

// GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagFilterExpression, true
}

// SetTagFilterExpression sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetTagFilterExpression(v TagFilterExpressionElement) {
	o.TagFilterExpression = v
}

// GetThreshold returns the Threshold field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetThreshold() Threshold {
	if o == nil {
		var ret Threshold
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetThresholdOk() (*Threshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetThreshold(v Threshold) {
	o.Threshold = v
}

// GetTimeThreshold returns the TimeThreshold field value
func (o *GlobalApplicationAlertConfigWithMetadata) GetTimeThreshold() ApplicationTimeThreshold {
	if o == nil {
		var ret ApplicationTimeThreshold
		return ret
	}

	return o.TimeThreshold
}

// GetTimeThresholdOk returns a tuple with the TimeThreshold field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetTimeThresholdOk() (*ApplicationTimeThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeThreshold, true
}

// SetTimeThreshold sets field value
func (o *GlobalApplicationAlertConfigWithMetadata) SetTimeThreshold(v ApplicationTimeThreshold) {
	o.TimeThreshold = v
}

// GetTriggering returns the Triggering field value if set, zero value otherwise.
func (o *GlobalApplicationAlertConfigWithMetadata) GetTriggering() bool {
	if o == nil || isNil(o.Triggering) {
		var ret bool
		return ret
	}
	return *o.Triggering
}

// GetTriggeringOk returns a tuple with the Triggering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) GetTriggeringOk() (*bool, bool) {
	if o == nil || isNil(o.Triggering) {
		return nil, false
	}
	return o.Triggering, true
}

// HasTriggering returns a boolean if a field has been set.
func (o *GlobalApplicationAlertConfigWithMetadata) HasTriggering() bool {
	if o != nil && !isNil(o.Triggering) {
		return true
	}

	return false
}

// SetTriggering gets a reference to the given bool and assigns it to the Triggering field.
func (o *GlobalApplicationAlertConfigWithMetadata) SetTriggering(v bool) {
	o.Triggering = &v
}

func (o GlobalApplicationAlertConfigWithMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alertChannelIds"] = o.AlertChannelIds
	}
	if !isNil(o.ApplicationIds) {
		toSerialize["applicationIds"] = o.ApplicationIds
	}
	if true {
		toSerialize["applications"] = o.Applications
	}
	if true {
		toSerialize["boundaryScope"] = o.BoundaryScope
	}
	if !isNil(o.BuiltIn) {
		toSerialize["builtIn"] = o.BuiltIn
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["customPayloadFields"] = o.CustomPayloadFields
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["evaluationType"] = o.EvaluationType
	}
	if true {
		toSerialize["granularity"] = o.Granularity
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IncludeInternal) {
		toSerialize["includeInternal"] = o.IncludeInternal
	}
	if !isNil(o.IncludeSynthetic) {
		toSerialize["includeSynthetic"] = o.IncludeSynthetic
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if true {
		toSerialize["rule"] = o.Rule
	}
	if !isNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["tagFilterExpression"] = o.TagFilterExpression
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["timeThreshold"] = o.TimeThreshold
	}
	if !isNil(o.Triggering) {
		toSerialize["triggering"] = o.Triggering
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalApplicationAlertConfigWithMetadata struct {
	value *GlobalApplicationAlertConfigWithMetadata
	isSet bool
}

func (v NullableGlobalApplicationAlertConfigWithMetadata) Get() *GlobalApplicationAlertConfigWithMetadata {
	return v.value
}

func (v *NullableGlobalApplicationAlertConfigWithMetadata) Set(val *GlobalApplicationAlertConfigWithMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalApplicationAlertConfigWithMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalApplicationAlertConfigWithMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalApplicationAlertConfigWithMetadata(val *GlobalApplicationAlertConfigWithMetadata) *NullableGlobalApplicationAlertConfigWithMetadata {
	return &NullableGlobalApplicationAlertConfigWithMetadata{value: val, isSet: true}
}

func (v NullableGlobalApplicationAlertConfigWithMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalApplicationAlertConfigWithMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
