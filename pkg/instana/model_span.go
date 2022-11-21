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

// Span struct for Span
type Span struct {
	BatchSelfTime      *int64                 `json:"batchSelfTime,omitempty"`
	BatchSize          *int32                 `json:"batchSize,omitempty"`
	CalculatedSelfTime *int64                 `json:"calculatedSelfTime,omitempty"`
	CallId             string                 `json:"callId"`
	ChildSpans         []Span                 `json:"childSpans"`
	Data               map[string]interface{} `json:"data"`
	Destination        *SpanRelation          `json:"destination,omitempty"`
	Duration           *int64                 `json:"duration,omitempty"`
	ErrorCount         *int32                 `json:"errorCount,omitempty"`
	ForeignParentId    *string                `json:"foreignParentId,omitempty"`
	Id                 string                 `json:"id"`
	IsSynthetic        *bool                  `json:"isSynthetic,omitempty"`
	Kind               string                 `json:"kind"`
	Label              string                 `json:"label"`
	Name               string                 `json:"name"`
	ParentId           *string                `json:"parentId,omitempty"`
	Source             *SpanRelation          `json:"source,omitempty"`
	StackTrace         []StackTraceItem       `json:"stackTrace"`
	Start              *int64                 `json:"start,omitempty"`
}

// NewSpan instantiates a new Span object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpan(callId string, childSpans []Span, data map[string]interface{}, id string, kind string, label string, name string, stackTrace []StackTraceItem) *Span {
	this := Span{}
	this.CallId = callId
	this.ChildSpans = childSpans
	this.Data = data
	this.Id = id
	this.Kind = kind
	this.Label = label
	this.Name = name
	this.StackTrace = stackTrace
	return &this
}

// NewSpanWithDefaults instantiates a new Span object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanWithDefaults() *Span {
	this := Span{}
	return &this
}

// GetBatchSelfTime returns the BatchSelfTime field value if set, zero value otherwise.
func (o *Span) GetBatchSelfTime() int64 {
	if o == nil || isNil(o.BatchSelfTime) {
		var ret int64
		return ret
	}
	return *o.BatchSelfTime
}

// GetBatchSelfTimeOk returns a tuple with the BatchSelfTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetBatchSelfTimeOk() (*int64, bool) {
	if o == nil || isNil(o.BatchSelfTime) {
		return nil, false
	}
	return o.BatchSelfTime, true
}

// HasBatchSelfTime returns a boolean if a field has been set.
func (o *Span) HasBatchSelfTime() bool {
	if o != nil && !isNil(o.BatchSelfTime) {
		return true
	}

	return false
}

// SetBatchSelfTime gets a reference to the given int64 and assigns it to the BatchSelfTime field.
func (o *Span) SetBatchSelfTime(v int64) {
	o.BatchSelfTime = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *Span) GetBatchSize() int32 {
	if o == nil || isNil(o.BatchSize) {
		var ret int32
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetBatchSizeOk() (*int32, bool) {
	if o == nil || isNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *Span) HasBatchSize() bool {
	if o != nil && !isNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.
func (o *Span) SetBatchSize(v int32) {
	o.BatchSize = &v
}

// GetCalculatedSelfTime returns the CalculatedSelfTime field value if set, zero value otherwise.
func (o *Span) GetCalculatedSelfTime() int64 {
	if o == nil || isNil(o.CalculatedSelfTime) {
		var ret int64
		return ret
	}
	return *o.CalculatedSelfTime
}

// GetCalculatedSelfTimeOk returns a tuple with the CalculatedSelfTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetCalculatedSelfTimeOk() (*int64, bool) {
	if o == nil || isNil(o.CalculatedSelfTime) {
		return nil, false
	}
	return o.CalculatedSelfTime, true
}

// HasCalculatedSelfTime returns a boolean if a field has been set.
func (o *Span) HasCalculatedSelfTime() bool {
	if o != nil && !isNil(o.CalculatedSelfTime) {
		return true
	}

	return false
}

// SetCalculatedSelfTime gets a reference to the given int64 and assigns it to the CalculatedSelfTime field.
func (o *Span) SetCalculatedSelfTime(v int64) {
	o.CalculatedSelfTime = &v
}

// GetCallId returns the CallId field value
func (o *Span) GetCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value
// and a boolean to check if the value has been set.
func (o *Span) GetCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallId, true
}

// SetCallId sets field value
func (o *Span) SetCallId(v string) {
	o.CallId = v
}

// GetChildSpans returns the ChildSpans field value
func (o *Span) GetChildSpans() []Span {
	if o == nil {
		var ret []Span
		return ret
	}

	return o.ChildSpans
}

// GetChildSpansOk returns a tuple with the ChildSpans field value
// and a boolean to check if the value has been set.
func (o *Span) GetChildSpansOk() ([]Span, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChildSpans, true
}

// SetChildSpans sets field value
func (o *Span) SetChildSpans(v []Span) {
	o.ChildSpans = v
}

// GetData returns the Data field value
func (o *Span) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Span) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *Span) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *Span) GetDestination() SpanRelation {
	if o == nil || isNil(o.Destination) {
		var ret SpanRelation
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetDestinationOk() (*SpanRelation, bool) {
	if o == nil || isNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *Span) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given SpanRelation and assigns it to the Destination field.
func (o *Span) SetDestination(v SpanRelation) {
	o.Destination = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Span) GetDuration() int64 {
	if o == nil || isNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetDurationOk() (*int64, bool) {
	if o == nil || isNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Span) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *Span) SetDuration(v int64) {
	o.Duration = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *Span) GetErrorCount() int32 {
	if o == nil || isNil(o.ErrorCount) {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetErrorCountOk() (*int32, bool) {
	if o == nil || isNil(o.ErrorCount) {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *Span) HasErrorCount() bool {
	if o != nil && !isNil(o.ErrorCount) {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *Span) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

// GetForeignParentId returns the ForeignParentId field value if set, zero value otherwise.
func (o *Span) GetForeignParentId() string {
	if o == nil || isNil(o.ForeignParentId) {
		var ret string
		return ret
	}
	return *o.ForeignParentId
}

// GetForeignParentIdOk returns a tuple with the ForeignParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetForeignParentIdOk() (*string, bool) {
	if o == nil || isNil(o.ForeignParentId) {
		return nil, false
	}
	return o.ForeignParentId, true
}

// HasForeignParentId returns a boolean if a field has been set.
func (o *Span) HasForeignParentId() bool {
	if o != nil && !isNil(o.ForeignParentId) {
		return true
	}

	return false
}

// SetForeignParentId gets a reference to the given string and assigns it to the ForeignParentId field.
func (o *Span) SetForeignParentId(v string) {
	o.ForeignParentId = &v
}

// GetId returns the Id field value
func (o *Span) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Span) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Span) SetId(v string) {
	o.Id = v
}

// GetIsSynthetic returns the IsSynthetic field value if set, zero value otherwise.
func (o *Span) GetIsSynthetic() bool {
	if o == nil || isNil(o.IsSynthetic) {
		var ret bool
		return ret
	}
	return *o.IsSynthetic
}

// GetIsSyntheticOk returns a tuple with the IsSynthetic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetIsSyntheticOk() (*bool, bool) {
	if o == nil || isNil(o.IsSynthetic) {
		return nil, false
	}
	return o.IsSynthetic, true
}

// HasIsSynthetic returns a boolean if a field has been set.
func (o *Span) HasIsSynthetic() bool {
	if o != nil && !isNil(o.IsSynthetic) {
		return true
	}

	return false
}

// SetIsSynthetic gets a reference to the given bool and assigns it to the IsSynthetic field.
func (o *Span) SetIsSynthetic(v bool) {
	o.IsSynthetic = &v
}

// GetKind returns the Kind field value
func (o *Span) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Span) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Span) SetKind(v string) {
	o.Kind = v
}

// GetLabel returns the Label field value
func (o *Span) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Span) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Span) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *Span) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Span) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Span) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Span) GetParentId() string {
	if o == nil || isNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetParentIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Span) HasParentId() bool {
	if o != nil && !isNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *Span) SetParentId(v string) {
	o.ParentId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Span) GetSource() SpanRelation {
	if o == nil || isNil(o.Source) {
		var ret SpanRelation
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetSourceOk() (*SpanRelation, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Span) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given SpanRelation and assigns it to the Source field.
func (o *Span) SetSource(v SpanRelation) {
	o.Source = &v
}

// GetStackTrace returns the StackTrace field value
func (o *Span) GetStackTrace() []StackTraceItem {
	if o == nil {
		var ret []StackTraceItem
		return ret
	}

	return o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value
// and a boolean to check if the value has been set.
func (o *Span) GetStackTraceOk() ([]StackTraceItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.StackTrace, true
}

// SetStackTrace sets field value
func (o *Span) SetStackTrace(v []StackTraceItem) {
	o.StackTrace = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Span) GetStart() int64 {
	if o == nil || isNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Span) GetStartOk() (*int64, bool) {
	if o == nil || isNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Span) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *Span) SetStart(v int64) {
	o.Start = &v
}

func (o Span) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BatchSelfTime) {
		toSerialize["batchSelfTime"] = o.BatchSelfTime
	}
	if !isNil(o.BatchSize) {
		toSerialize["batchSize"] = o.BatchSize
	}
	if !isNil(o.CalculatedSelfTime) {
		toSerialize["calculatedSelfTime"] = o.CalculatedSelfTime
	}
	if true {
		toSerialize["callId"] = o.CallId
	}
	if true {
		toSerialize["childSpans"] = o.ChildSpans
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.ErrorCount) {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if !isNil(o.ForeignParentId) {
		toSerialize["foreignParentId"] = o.ForeignParentId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IsSynthetic) {
		toSerialize["isSynthetic"] = o.IsSynthetic
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["stackTrace"] = o.StackTrace
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableSpan struct {
	value *Span
	isSet bool
}

func (v NullableSpan) Get() *Span {
	return v.value
}

func (v *NullableSpan) Set(val *Span) {
	v.value = val
	v.isSet = true
}

func (v NullableSpan) IsSet() bool {
	return v.isSet
}

func (v *NullableSpan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpan(val *Span) *NullableSpan {
	return &NullableSpan{value: val, isSet: true}
}

func (v NullableSpan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
