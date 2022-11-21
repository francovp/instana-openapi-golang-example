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

// EventResult struct for EventResult
type EventResult struct {
	Detail         *string                  `json:"detail,omitempty"`
	End            *int64                   `json:"end,omitempty"`
	EntityLabel    *string                  `json:"entityLabel,omitempty"`
	EntityName     *string                  `json:"entityName,omitempty"`
	EntityType     *string                  `json:"entityType,omitempty"`
	EventId        *string                  `json:"eventId,omitempty"`
	FixSuggestion  *string                  `json:"fixSuggestion,omitempty"`
	Metrics        []map[string]interface{} `json:"metrics,omitempty"`
	Problem        *string                  `json:"problem,omitempty"`
	RecentEvents   []map[string]interface{} `json:"recentEvents,omitempty"`
	Severity       *int32                   `json:"severity,omitempty"`
	SnapshotId     *string                  `json:"snapshotId,omitempty"`
	Start          *int64                   `json:"start,omitempty"`
	State          *string                  `json:"state,omitempty"`
	TriggeringTime *int64                   `json:"triggeringTime,omitempty"`
	Type           *string                  `json:"type,omitempty"`
}

// NewEventResult instantiates a new EventResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResult() *EventResult {
	this := EventResult{}
	return &this
}

// NewEventResultWithDefaults instantiates a new EventResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResultWithDefaults() *EventResult {
	this := EventResult{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *EventResult) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *EventResult) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *EventResult) SetDetail(v string) {
	o.Detail = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *EventResult) GetEnd() int64 {
	if o == nil || isNil(o.End) {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetEndOk() (*int64, bool) {
	if o == nil || isNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *EventResult) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *EventResult) SetEnd(v int64) {
	o.End = &v
}

// GetEntityLabel returns the EntityLabel field value if set, zero value otherwise.
func (o *EventResult) GetEntityLabel() string {
	if o == nil || isNil(o.EntityLabel) {
		var ret string
		return ret
	}
	return *o.EntityLabel
}

// GetEntityLabelOk returns a tuple with the EntityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetEntityLabelOk() (*string, bool) {
	if o == nil || isNil(o.EntityLabel) {
		return nil, false
	}
	return o.EntityLabel, true
}

// HasEntityLabel returns a boolean if a field has been set.
func (o *EventResult) HasEntityLabel() bool {
	if o != nil && !isNil(o.EntityLabel) {
		return true
	}

	return false
}

// SetEntityLabel gets a reference to the given string and assigns it to the EntityLabel field.
func (o *EventResult) SetEntityLabel(v string) {
	o.EntityLabel = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *EventResult) GetEntityName() string {
	if o == nil || isNil(o.EntityName) {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetEntityNameOk() (*string, bool) {
	if o == nil || isNil(o.EntityName) {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *EventResult) HasEntityName() bool {
	if o != nil && !isNil(o.EntityName) {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *EventResult) SetEntityName(v string) {
	o.EntityName = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *EventResult) GetEntityType() string {
	if o == nil || isNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetEntityTypeOk() (*string, bool) {
	if o == nil || isNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *EventResult) HasEntityType() bool {
	if o != nil && !isNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *EventResult) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EventResult) GetEventId() string {
	if o == nil || isNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetEventIdOk() (*string, bool) {
	if o == nil || isNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EventResult) HasEventId() bool {
	if o != nil && !isNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *EventResult) SetEventId(v string) {
	o.EventId = &v
}

// GetFixSuggestion returns the FixSuggestion field value if set, zero value otherwise.
func (o *EventResult) GetFixSuggestion() string {
	if o == nil || isNil(o.FixSuggestion) {
		var ret string
		return ret
	}
	return *o.FixSuggestion
}

// GetFixSuggestionOk returns a tuple with the FixSuggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetFixSuggestionOk() (*string, bool) {
	if o == nil || isNil(o.FixSuggestion) {
		return nil, false
	}
	return o.FixSuggestion, true
}

// HasFixSuggestion returns a boolean if a field has been set.
func (o *EventResult) HasFixSuggestion() bool {
	if o != nil && !isNil(o.FixSuggestion) {
		return true
	}

	return false
}

// SetFixSuggestion gets a reference to the given string and assigns it to the FixSuggestion field.
func (o *EventResult) SetFixSuggestion(v string) {
	o.FixSuggestion = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *EventResult) GetMetrics() []map[string]interface{} {
	if o == nil || isNil(o.Metrics) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetMetricsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *EventResult) HasMetrics() bool {
	if o != nil && !isNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []map[string]interface{} and assigns it to the Metrics field.
func (o *EventResult) SetMetrics(v []map[string]interface{}) {
	o.Metrics = v
}

// GetProblem returns the Problem field value if set, zero value otherwise.
func (o *EventResult) GetProblem() string {
	if o == nil || isNil(o.Problem) {
		var ret string
		return ret
	}
	return *o.Problem
}

// GetProblemOk returns a tuple with the Problem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetProblemOk() (*string, bool) {
	if o == nil || isNil(o.Problem) {
		return nil, false
	}
	return o.Problem, true
}

// HasProblem returns a boolean if a field has been set.
func (o *EventResult) HasProblem() bool {
	if o != nil && !isNil(o.Problem) {
		return true
	}

	return false
}

// SetProblem gets a reference to the given string and assigns it to the Problem field.
func (o *EventResult) SetProblem(v string) {
	o.Problem = &v
}

// GetRecentEvents returns the RecentEvents field value if set, zero value otherwise.
func (o *EventResult) GetRecentEvents() []map[string]interface{} {
	if o == nil || isNil(o.RecentEvents) {
		var ret []map[string]interface{}
		return ret
	}
	return o.RecentEvents
}

// GetRecentEventsOk returns a tuple with the RecentEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetRecentEventsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.RecentEvents) {
		return nil, false
	}
	return o.RecentEvents, true
}

// HasRecentEvents returns a boolean if a field has been set.
func (o *EventResult) HasRecentEvents() bool {
	if o != nil && !isNil(o.RecentEvents) {
		return true
	}

	return false
}

// SetRecentEvents gets a reference to the given []map[string]interface{} and assigns it to the RecentEvents field.
func (o *EventResult) SetRecentEvents(v []map[string]interface{}) {
	o.RecentEvents = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *EventResult) GetSeverity() int32 {
	if o == nil || isNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetSeverityOk() (*int32, bool) {
	if o == nil || isNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *EventResult) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *EventResult) SetSeverity(v int32) {
	o.Severity = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *EventResult) GetSnapshotId() string {
	if o == nil || isNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetSnapshotIdOk() (*string, bool) {
	if o == nil || isNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *EventResult) HasSnapshotId() bool {
	if o != nil && !isNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *EventResult) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *EventResult) GetStart() int64 {
	if o == nil || isNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetStartOk() (*int64, bool) {
	if o == nil || isNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *EventResult) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *EventResult) SetStart(v int64) {
	o.Start = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EventResult) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EventResult) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EventResult) SetState(v string) {
	o.State = &v
}

// GetTriggeringTime returns the TriggeringTime field value if set, zero value otherwise.
func (o *EventResult) GetTriggeringTime() int64 {
	if o == nil || isNil(o.TriggeringTime) {
		var ret int64
		return ret
	}
	return *o.TriggeringTime
}

// GetTriggeringTimeOk returns a tuple with the TriggeringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetTriggeringTimeOk() (*int64, bool) {
	if o == nil || isNil(o.TriggeringTime) {
		return nil, false
	}
	return o.TriggeringTime, true
}

// HasTriggeringTime returns a boolean if a field has been set.
func (o *EventResult) HasTriggeringTime() bool {
	if o != nil && !isNil(o.TriggeringTime) {
		return true
	}

	return false
}

// SetTriggeringTime gets a reference to the given int64 and assigns it to the TriggeringTime field.
func (o *EventResult) SetTriggeringTime(v int64) {
	o.TriggeringTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventResult) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResult) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventResult) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventResult) SetType(v string) {
	o.Type = &v
}

func (o EventResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.EntityLabel) {
		toSerialize["entityLabel"] = o.EntityLabel
	}
	if !isNil(o.EntityName) {
		toSerialize["entityName"] = o.EntityName
	}
	if !isNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !isNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !isNil(o.FixSuggestion) {
		toSerialize["fixSuggestion"] = o.FixSuggestion
	}
	if !isNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !isNil(o.Problem) {
		toSerialize["problem"] = o.Problem
	}
	if !isNil(o.RecentEvents) {
		toSerialize["recentEvents"] = o.RecentEvents
	}
	if !isNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !isNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.TriggeringTime) {
		toSerialize["triggeringTime"] = o.TriggeringTime
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEventResult struct {
	value *EventResult
	isSet bool
}

func (v NullableEventResult) Get() *EventResult {
	return v.value
}

func (v *NullableEventResult) Set(val *EventResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResult(val *EventResult) *NullableEventResult {
	return &NullableEventResult{value: val, isSet: true}
}

func (v NullableEventResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
