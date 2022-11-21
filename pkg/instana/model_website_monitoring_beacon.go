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

// WebsiteMonitoringBeacon struct for WebsiteMonitoringBeacon
type WebsiteMonitoringBeacon struct {
	AccuracyRadius               *int64             `json:"accuracyRadius,omitempty"`
	AccurateTimingsAvailable     *bool              `json:"accurateTimingsAvailable,omitempty"`
	AppCacheTime                 *int64             `json:"appCacheTime,omitempty"`
	BackendTime                  *int64             `json:"backendTime,omitempty"`
	BackendTraceId               *string            `json:"backendTraceId,omitempty"`
	BatchSize                    *int64             `json:"batchSize,omitempty"`
	BeaconId                     string             `json:"beaconId"`
	BrowserName                  *string            `json:"browserName,omitempty"`
	BrowserVersion               *string            `json:"browserVersion,omitempty"`
	CacheInteraction             *string            `json:"cacheInteraction,omitempty"`
	ChildrenTime                 *int64             `json:"childrenTime,omitempty"`
	City                         *string            `json:"city,omitempty"`
	ClockSkew                    *int64             `json:"clockSkew,omitempty"`
	ComponentStack               *string            `json:"componentStack,omitempty"`
	ConnectionType               *string            `json:"connectionType,omitempty"`
	Continent                    *string            `json:"continent,omitempty"`
	ContinentCode                *string            `json:"continentCode,omitempty"`
	Country                      *string            `json:"country,omitempty"`
	CountryCode                  *string            `json:"countryCode,omitempty"`
	CspBlockedUri                *string            `json:"cspBlockedUri,omitempty"`
	CspColumnNumber              *int64             `json:"cspColumnNumber,omitempty"`
	CspDisposition               *string            `json:"cspDisposition,omitempty"`
	CspEffectiveDirective        *string            `json:"cspEffectiveDirective,omitempty"`
	CspLineNumber                *int64             `json:"cspLineNumber,omitempty"`
	CspOriginalPolicy            *string            `json:"cspOriginalPolicy,omitempty"`
	CspSample                    *string            `json:"cspSample,omitempty"`
	CspSourceFile                *string            `json:"cspSourceFile,omitempty"`
	CumulativeLayoutShift        *float64           `json:"cumulativeLayoutShift,omitempty"`
	CustomEventName              *string            `json:"customEventName,omitempty"`
	DecodedBodySize              *int64             `json:"decodedBodySize,omitempty"`
	Deprecations                 []string           `json:"deprecations,omitempty"`
	DeviceType                   *string            `json:"deviceType,omitempty"`
	DnsTime                      *int64             `json:"dnsTime,omitempty"`
	DomTime                      *int64             `json:"domTime,omitempty"`
	Duration                     *int64             `json:"duration,omitempty"`
	EncodedBodySize              *int64             `json:"encodedBodySize,omitempty"`
	ErrorCount                   *int64             `json:"errorCount,omitempty"`
	ErrorId                      *string            `json:"errorId,omitempty"`
	ErrorMessage                 *string            `json:"errorMessage,omitempty"`
	ErrorType                    *string            `json:"errorType,omitempty"`
	FirstContentfulPaintTime     *int64             `json:"firstContentfulPaintTime,omitempty"`
	FirstInputDelayTime          *int64             `json:"firstInputDelayTime,omitempty"`
	FirstPaintTime               *int64             `json:"firstPaintTime,omitempty"`
	FrontendTime                 *int64             `json:"frontendTime,omitempty"`
	GraphqlOperationName         *string            `json:"graphqlOperationName,omitempty"`
	GraphqlOperationType         *string            `json:"graphqlOperationType,omitempty"`
	HttpCallAsynchronous         *bool              `json:"httpCallAsynchronous,omitempty"`
	HttpCallCorrelationAttempted *bool              `json:"httpCallCorrelationAttempted,omitempty"`
	HttpCallHeaders              *map[string]string `json:"httpCallHeaders,omitempty"`
	HttpCallMethod               *string            `json:"httpCallMethod,omitempty"`
	HttpCallOrigin               *string            `json:"httpCallOrigin,omitempty"`
	HttpCallPath                 *string            `json:"httpCallPath,omitempty"`
	HttpCallStatus               *int32             `json:"httpCallStatus,omitempty"`
	HttpCallUrl                  *string            `json:"httpCallUrl,omitempty"`
	Initiator                    *string            `json:"initiator,omitempty"`
	LargestContentfulPaintTime   *int64             `json:"largestContentfulPaintTime,omitempty"`
	Latitude                     *float64           `json:"latitude,omitempty"`
	LocationOrigin               string             `json:"locationOrigin"`
	LocationPath                 *string            `json:"locationPath,omitempty"`
	LocationUrl                  string             `json:"locationUrl"`
	Longitude                    *float64           `json:"longitude,omitempty"`
	Meta                         *map[string]string `json:"meta,omitempty"`
	OnLoadTime                   *int64             `json:"onLoadTime,omitempty"`
	OsName                       *string            `json:"osName,omitempty"`
	OsVersion                    *string            `json:"osVersion,omitempty"`
	Page                         *string            `json:"page,omitempty"`
	PageLoadId                   string             `json:"pageLoadId"`
	ParsedStackTrace             []StackTraceLine   `json:"parsedStackTrace,omitempty"`
	Phase                        *string            `json:"phase,omitempty"`
	ProcessingTime               *int64             `json:"processingTime,omitempty"`
	RedirectTime                 *int64             `json:"redirectTime,omitempty"`
	RequestTime                  *int64             `json:"requestTime,omitempty"`
	ResourceType                 *string            `json:"resourceType,omitempty"`
	ResponseTime                 *int64             `json:"responseTime,omitempty"`
	SessionId                    *string            `json:"sessionId,omitempty"`
	SnippetVersion               *string            `json:"snippetVersion,omitempty"`
	SslTime                      *int64             `json:"sslTime,omitempty"`
	StackTrace                   *string            `json:"stackTrace,omitempty"`
	StackTraceParsingStatus      *int32             `json:"stackTraceParsingStatus,omitempty"`
	StackTraceReadability        *int32             `json:"stackTraceReadability,omitempty"`
	Subdivision                  *string            `json:"subdivision,omitempty"`
	SubdivisionCode              *string            `json:"subdivisionCode,omitempty"`
	TcpTime                      *int64             `json:"tcpTime,omitempty"`
	Timestamp                    *int64             `json:"timestamp,omitempty"`
	TransferSize                 *int64             `json:"transferSize,omitempty"`
	Type                         string             `json:"type"`
	UnloadTime                   *int64             `json:"unloadTime,omitempty"`
	UserEmail                    *string            `json:"userEmail,omitempty"`
	UserId                       *string            `json:"userId,omitempty"`
	UserIp                       *string            `json:"userIp,omitempty"`
	UserLanguages                []string           `json:"userLanguages,omitempty"`
	UserName                     *string            `json:"userName,omitempty"`
	WebsiteId                    string             `json:"websiteId"`
	WebsiteLabel                 string             `json:"websiteLabel"`
	WindowHeight                 *int32             `json:"windowHeight,omitempty"`
	WindowHidden                 *bool              `json:"windowHidden,omitempty"`
	WindowWidth                  *int32             `json:"windowWidth,omitempty"`
}

// NewWebsiteMonitoringBeacon instantiates a new WebsiteMonitoringBeacon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsiteMonitoringBeacon(beaconId string, locationOrigin string, locationUrl string, pageLoadId string, type_ string, websiteId string, websiteLabel string) *WebsiteMonitoringBeacon {
	this := WebsiteMonitoringBeacon{}
	this.BeaconId = beaconId
	this.LocationOrigin = locationOrigin
	this.LocationUrl = locationUrl
	this.PageLoadId = pageLoadId
	this.Type = type_
	this.WebsiteId = websiteId
	this.WebsiteLabel = websiteLabel
	return &this
}

// NewWebsiteMonitoringBeaconWithDefaults instantiates a new WebsiteMonitoringBeacon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsiteMonitoringBeaconWithDefaults() *WebsiteMonitoringBeacon {
	this := WebsiteMonitoringBeacon{}
	return &this
}

// GetAccuracyRadius returns the AccuracyRadius field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetAccuracyRadius() int64 {
	if o == nil || isNil(o.AccuracyRadius) {
		var ret int64
		return ret
	}
	return *o.AccuracyRadius
}

// GetAccuracyRadiusOk returns a tuple with the AccuracyRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetAccuracyRadiusOk() (*int64, bool) {
	if o == nil || isNil(o.AccuracyRadius) {
		return nil, false
	}
	return o.AccuracyRadius, true
}

// HasAccuracyRadius returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasAccuracyRadius() bool {
	if o != nil && !isNil(o.AccuracyRadius) {
		return true
	}

	return false
}

// SetAccuracyRadius gets a reference to the given int64 and assigns it to the AccuracyRadius field.
func (o *WebsiteMonitoringBeacon) SetAccuracyRadius(v int64) {
	o.AccuracyRadius = &v
}

// GetAccurateTimingsAvailable returns the AccurateTimingsAvailable field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetAccurateTimingsAvailable() bool {
	if o == nil || isNil(o.AccurateTimingsAvailable) {
		var ret bool
		return ret
	}
	return *o.AccurateTimingsAvailable
}

// GetAccurateTimingsAvailableOk returns a tuple with the AccurateTimingsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetAccurateTimingsAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.AccurateTimingsAvailable) {
		return nil, false
	}
	return o.AccurateTimingsAvailable, true
}

// HasAccurateTimingsAvailable returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasAccurateTimingsAvailable() bool {
	if o != nil && !isNil(o.AccurateTimingsAvailable) {
		return true
	}

	return false
}

// SetAccurateTimingsAvailable gets a reference to the given bool and assigns it to the AccurateTimingsAvailable field.
func (o *WebsiteMonitoringBeacon) SetAccurateTimingsAvailable(v bool) {
	o.AccurateTimingsAvailable = &v
}

// GetAppCacheTime returns the AppCacheTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetAppCacheTime() int64 {
	if o == nil || isNil(o.AppCacheTime) {
		var ret int64
		return ret
	}
	return *o.AppCacheTime
}

// GetAppCacheTimeOk returns a tuple with the AppCacheTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetAppCacheTimeOk() (*int64, bool) {
	if o == nil || isNil(o.AppCacheTime) {
		return nil, false
	}
	return o.AppCacheTime, true
}

// HasAppCacheTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasAppCacheTime() bool {
	if o != nil && !isNil(o.AppCacheTime) {
		return true
	}

	return false
}

// SetAppCacheTime gets a reference to the given int64 and assigns it to the AppCacheTime field.
func (o *WebsiteMonitoringBeacon) SetAppCacheTime(v int64) {
	o.AppCacheTime = &v
}

// GetBackendTime returns the BackendTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetBackendTime() int64 {
	if o == nil || isNil(o.BackendTime) {
		var ret int64
		return ret
	}
	return *o.BackendTime
}

// GetBackendTimeOk returns a tuple with the BackendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetBackendTimeOk() (*int64, bool) {
	if o == nil || isNil(o.BackendTime) {
		return nil, false
	}
	return o.BackendTime, true
}

// HasBackendTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasBackendTime() bool {
	if o != nil && !isNil(o.BackendTime) {
		return true
	}

	return false
}

// SetBackendTime gets a reference to the given int64 and assigns it to the BackendTime field.
func (o *WebsiteMonitoringBeacon) SetBackendTime(v int64) {
	o.BackendTime = &v
}

// GetBackendTraceId returns the BackendTraceId field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetBackendTraceId() string {
	if o == nil || isNil(o.BackendTraceId) {
		var ret string
		return ret
	}
	return *o.BackendTraceId
}

// GetBackendTraceIdOk returns a tuple with the BackendTraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetBackendTraceIdOk() (*string, bool) {
	if o == nil || isNil(o.BackendTraceId) {
		return nil, false
	}
	return o.BackendTraceId, true
}

// HasBackendTraceId returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasBackendTraceId() bool {
	if o != nil && !isNil(o.BackendTraceId) {
		return true
	}

	return false
}

// SetBackendTraceId gets a reference to the given string and assigns it to the BackendTraceId field.
func (o *WebsiteMonitoringBeacon) SetBackendTraceId(v string) {
	o.BackendTraceId = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetBatchSize() int64 {
	if o == nil || isNil(o.BatchSize) {
		var ret int64
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetBatchSizeOk() (*int64, bool) {
	if o == nil || isNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasBatchSize() bool {
	if o != nil && !isNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int64 and assigns it to the BatchSize field.
func (o *WebsiteMonitoringBeacon) SetBatchSize(v int64) {
	o.BatchSize = &v
}

// GetBeaconId returns the BeaconId field value
func (o *WebsiteMonitoringBeacon) GetBeaconId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeaconId
}

// GetBeaconIdOk returns a tuple with the BeaconId field value
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetBeaconIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeaconId, true
}

// SetBeaconId sets field value
func (o *WebsiteMonitoringBeacon) SetBeaconId(v string) {
	o.BeaconId = v
}

// GetBrowserName returns the BrowserName field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetBrowserName() string {
	if o == nil || isNil(o.BrowserName) {
		var ret string
		return ret
	}
	return *o.BrowserName
}

// GetBrowserNameOk returns a tuple with the BrowserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetBrowserNameOk() (*string, bool) {
	if o == nil || isNil(o.BrowserName) {
		return nil, false
	}
	return o.BrowserName, true
}

// HasBrowserName returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasBrowserName() bool {
	if o != nil && !isNil(o.BrowserName) {
		return true
	}

	return false
}

// SetBrowserName gets a reference to the given string and assigns it to the BrowserName field.
func (o *WebsiteMonitoringBeacon) SetBrowserName(v string) {
	o.BrowserName = &v
}

// GetBrowserVersion returns the BrowserVersion field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetBrowserVersion() string {
	if o == nil || isNil(o.BrowserVersion) {
		var ret string
		return ret
	}
	return *o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetBrowserVersionOk() (*string, bool) {
	if o == nil || isNil(o.BrowserVersion) {
		return nil, false
	}
	return o.BrowserVersion, true
}

// HasBrowserVersion returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasBrowserVersion() bool {
	if o != nil && !isNil(o.BrowserVersion) {
		return true
	}

	return false
}

// SetBrowserVersion gets a reference to the given string and assigns it to the BrowserVersion field.
func (o *WebsiteMonitoringBeacon) SetBrowserVersion(v string) {
	o.BrowserVersion = &v
}

// GetCacheInteraction returns the CacheInteraction field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCacheInteraction() string {
	if o == nil || isNil(o.CacheInteraction) {
		var ret string
		return ret
	}
	return *o.CacheInteraction
}

// GetCacheInteractionOk returns a tuple with the CacheInteraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCacheInteractionOk() (*string, bool) {
	if o == nil || isNil(o.CacheInteraction) {
		return nil, false
	}
	return o.CacheInteraction, true
}

// HasCacheInteraction returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCacheInteraction() bool {
	if o != nil && !isNil(o.CacheInteraction) {
		return true
	}

	return false
}

// SetCacheInteraction gets a reference to the given string and assigns it to the CacheInteraction field.
func (o *WebsiteMonitoringBeacon) SetCacheInteraction(v string) {
	o.CacheInteraction = &v
}

// GetChildrenTime returns the ChildrenTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetChildrenTime() int64 {
	if o == nil || isNil(o.ChildrenTime) {
		var ret int64
		return ret
	}
	return *o.ChildrenTime
}

// GetChildrenTimeOk returns a tuple with the ChildrenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetChildrenTimeOk() (*int64, bool) {
	if o == nil || isNil(o.ChildrenTime) {
		return nil, false
	}
	return o.ChildrenTime, true
}

// HasChildrenTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasChildrenTime() bool {
	if o != nil && !isNil(o.ChildrenTime) {
		return true
	}

	return false
}

// SetChildrenTime gets a reference to the given int64 and assigns it to the ChildrenTime field.
func (o *WebsiteMonitoringBeacon) SetChildrenTime(v int64) {
	o.ChildrenTime = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *WebsiteMonitoringBeacon) SetCity(v string) {
	o.City = &v
}

// GetClockSkew returns the ClockSkew field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetClockSkew() int64 {
	if o == nil || isNil(o.ClockSkew) {
		var ret int64
		return ret
	}
	return *o.ClockSkew
}

// GetClockSkewOk returns a tuple with the ClockSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetClockSkewOk() (*int64, bool) {
	if o == nil || isNil(o.ClockSkew) {
		return nil, false
	}
	return o.ClockSkew, true
}

// HasClockSkew returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasClockSkew() bool {
	if o != nil && !isNil(o.ClockSkew) {
		return true
	}

	return false
}

// SetClockSkew gets a reference to the given int64 and assigns it to the ClockSkew field.
func (o *WebsiteMonitoringBeacon) SetClockSkew(v int64) {
	o.ClockSkew = &v
}

// GetComponentStack returns the ComponentStack field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetComponentStack() string {
	if o == nil || isNil(o.ComponentStack) {
		var ret string
		return ret
	}
	return *o.ComponentStack
}

// GetComponentStackOk returns a tuple with the ComponentStack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetComponentStackOk() (*string, bool) {
	if o == nil || isNil(o.ComponentStack) {
		return nil, false
	}
	return o.ComponentStack, true
}

// HasComponentStack returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasComponentStack() bool {
	if o != nil && !isNil(o.ComponentStack) {
		return true
	}

	return false
}

// SetComponentStack gets a reference to the given string and assigns it to the ComponentStack field.
func (o *WebsiteMonitoringBeacon) SetComponentStack(v string) {
	o.ComponentStack = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetConnectionType() string {
	if o == nil || isNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetConnectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasConnectionType() bool {
	if o != nil && !isNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *WebsiteMonitoringBeacon) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetContinent() string {
	if o == nil || isNil(o.Continent) {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetContinentOk() (*string, bool) {
	if o == nil || isNil(o.Continent) {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasContinent() bool {
	if o != nil && !isNil(o.Continent) {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *WebsiteMonitoringBeacon) SetContinent(v string) {
	o.Continent = &v
}

// GetContinentCode returns the ContinentCode field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetContinentCode() string {
	if o == nil || isNil(o.ContinentCode) {
		var ret string
		return ret
	}
	return *o.ContinentCode
}

// GetContinentCodeOk returns a tuple with the ContinentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetContinentCodeOk() (*string, bool) {
	if o == nil || isNil(o.ContinentCode) {
		return nil, false
	}
	return o.ContinentCode, true
}

// HasContinentCode returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasContinentCode() bool {
	if o != nil && !isNil(o.ContinentCode) {
		return true
	}

	return false
}

// SetContinentCode gets a reference to the given string and assigns it to the ContinentCode field.
func (o *WebsiteMonitoringBeacon) SetContinentCode(v string) {
	o.ContinentCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *WebsiteMonitoringBeacon) SetCountry(v string) {
	o.Country = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCountryCode() string {
	if o == nil || isNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCountryCodeOk() (*string, bool) {
	if o == nil || isNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCountryCode() bool {
	if o != nil && !isNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *WebsiteMonitoringBeacon) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCspBlockedUri returns the CspBlockedUri field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCspBlockedUri() string {
	if o == nil || isNil(o.CspBlockedUri) {
		var ret string
		return ret
	}
	return *o.CspBlockedUri
}

// GetCspBlockedUriOk returns a tuple with the CspBlockedUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCspBlockedUriOk() (*string, bool) {
	if o == nil || isNil(o.CspBlockedUri) {
		return nil, false
	}
	return o.CspBlockedUri, true
}

// HasCspBlockedUri returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCspBlockedUri() bool {
	if o != nil && !isNil(o.CspBlockedUri) {
		return true
	}

	return false
}

// SetCspBlockedUri gets a reference to the given string and assigns it to the CspBlockedUri field.
func (o *WebsiteMonitoringBeacon) SetCspBlockedUri(v string) {
	o.CspBlockedUri = &v
}

// GetCspColumnNumber returns the CspColumnNumber field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCspColumnNumber() int64 {
	if o == nil || isNil(o.CspColumnNumber) {
		var ret int64
		return ret
	}
	return *o.CspColumnNumber
}

// GetCspColumnNumberOk returns a tuple with the CspColumnNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCspColumnNumberOk() (*int64, bool) {
	if o == nil || isNil(o.CspColumnNumber) {
		return nil, false
	}
	return o.CspColumnNumber, true
}

// HasCspColumnNumber returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCspColumnNumber() bool {
	if o != nil && !isNil(o.CspColumnNumber) {
		return true
	}

	return false
}

// SetCspColumnNumber gets a reference to the given int64 and assigns it to the CspColumnNumber field.
func (o *WebsiteMonitoringBeacon) SetCspColumnNumber(v int64) {
	o.CspColumnNumber = &v
}

// GetCspDisposition returns the CspDisposition field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCspDisposition() string {
	if o == nil || isNil(o.CspDisposition) {
		var ret string
		return ret
	}
	return *o.CspDisposition
}

// GetCspDispositionOk returns a tuple with the CspDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCspDispositionOk() (*string, bool) {
	if o == nil || isNil(o.CspDisposition) {
		return nil, false
	}
	return o.CspDisposition, true
}

// HasCspDisposition returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCspDisposition() bool {
	if o != nil && !isNil(o.CspDisposition) {
		return true
	}

	return false
}

// SetCspDisposition gets a reference to the given string and assigns it to the CspDisposition field.
func (o *WebsiteMonitoringBeacon) SetCspDisposition(v string) {
	o.CspDisposition = &v
}

// GetCspEffectiveDirective returns the CspEffectiveDirective field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCspEffectiveDirective() string {
	if o == nil || isNil(o.CspEffectiveDirective) {
		var ret string
		return ret
	}
	return *o.CspEffectiveDirective
}

// GetCspEffectiveDirectiveOk returns a tuple with the CspEffectiveDirective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCspEffectiveDirectiveOk() (*string, bool) {
	if o == nil || isNil(o.CspEffectiveDirective) {
		return nil, false
	}
	return o.CspEffectiveDirective, true
}

// HasCspEffectiveDirective returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCspEffectiveDirective() bool {
	if o != nil && !isNil(o.CspEffectiveDirective) {
		return true
	}

	return false
}

// SetCspEffectiveDirective gets a reference to the given string and assigns it to the CspEffectiveDirective field.
func (o *WebsiteMonitoringBeacon) SetCspEffectiveDirective(v string) {
	o.CspEffectiveDirective = &v
}

// GetCspLineNumber returns the CspLineNumber field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCspLineNumber() int64 {
	if o == nil || isNil(o.CspLineNumber) {
		var ret int64
		return ret
	}
	return *o.CspLineNumber
}

// GetCspLineNumberOk returns a tuple with the CspLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCspLineNumberOk() (*int64, bool) {
	if o == nil || isNil(o.CspLineNumber) {
		return nil, false
	}
	return o.CspLineNumber, true
}

// HasCspLineNumber returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCspLineNumber() bool {
	if o != nil && !isNil(o.CspLineNumber) {
		return true
	}

	return false
}

// SetCspLineNumber gets a reference to the given int64 and assigns it to the CspLineNumber field.
func (o *WebsiteMonitoringBeacon) SetCspLineNumber(v int64) {
	o.CspLineNumber = &v
}

// GetCspOriginalPolicy returns the CspOriginalPolicy field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCspOriginalPolicy() string {
	if o == nil || isNil(o.CspOriginalPolicy) {
		var ret string
		return ret
	}
	return *o.CspOriginalPolicy
}

// GetCspOriginalPolicyOk returns a tuple with the CspOriginalPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCspOriginalPolicyOk() (*string, bool) {
	if o == nil || isNil(o.CspOriginalPolicy) {
		return nil, false
	}
	return o.CspOriginalPolicy, true
}

// HasCspOriginalPolicy returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCspOriginalPolicy() bool {
	if o != nil && !isNil(o.CspOriginalPolicy) {
		return true
	}

	return false
}

// SetCspOriginalPolicy gets a reference to the given string and assigns it to the CspOriginalPolicy field.
func (o *WebsiteMonitoringBeacon) SetCspOriginalPolicy(v string) {
	o.CspOriginalPolicy = &v
}

// GetCspSample returns the CspSample field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCspSample() string {
	if o == nil || isNil(o.CspSample) {
		var ret string
		return ret
	}
	return *o.CspSample
}

// GetCspSampleOk returns a tuple with the CspSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCspSampleOk() (*string, bool) {
	if o == nil || isNil(o.CspSample) {
		return nil, false
	}
	return o.CspSample, true
}

// HasCspSample returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCspSample() bool {
	if o != nil && !isNil(o.CspSample) {
		return true
	}

	return false
}

// SetCspSample gets a reference to the given string and assigns it to the CspSample field.
func (o *WebsiteMonitoringBeacon) SetCspSample(v string) {
	o.CspSample = &v
}

// GetCspSourceFile returns the CspSourceFile field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCspSourceFile() string {
	if o == nil || isNil(o.CspSourceFile) {
		var ret string
		return ret
	}
	return *o.CspSourceFile
}

// GetCspSourceFileOk returns a tuple with the CspSourceFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCspSourceFileOk() (*string, bool) {
	if o == nil || isNil(o.CspSourceFile) {
		return nil, false
	}
	return o.CspSourceFile, true
}

// HasCspSourceFile returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCspSourceFile() bool {
	if o != nil && !isNil(o.CspSourceFile) {
		return true
	}

	return false
}

// SetCspSourceFile gets a reference to the given string and assigns it to the CspSourceFile field.
func (o *WebsiteMonitoringBeacon) SetCspSourceFile(v string) {
	o.CspSourceFile = &v
}

// GetCumulativeLayoutShift returns the CumulativeLayoutShift field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCumulativeLayoutShift() float64 {
	if o == nil || isNil(o.CumulativeLayoutShift) {
		var ret float64
		return ret
	}
	return *o.CumulativeLayoutShift
}

// GetCumulativeLayoutShiftOk returns a tuple with the CumulativeLayoutShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCumulativeLayoutShiftOk() (*float64, bool) {
	if o == nil || isNil(o.CumulativeLayoutShift) {
		return nil, false
	}
	return o.CumulativeLayoutShift, true
}

// HasCumulativeLayoutShift returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCumulativeLayoutShift() bool {
	if o != nil && !isNil(o.CumulativeLayoutShift) {
		return true
	}

	return false
}

// SetCumulativeLayoutShift gets a reference to the given float64 and assigns it to the CumulativeLayoutShift field.
func (o *WebsiteMonitoringBeacon) SetCumulativeLayoutShift(v float64) {
	o.CumulativeLayoutShift = &v
}

// GetCustomEventName returns the CustomEventName field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetCustomEventName() string {
	if o == nil || isNil(o.CustomEventName) {
		var ret string
		return ret
	}
	return *o.CustomEventName
}

// GetCustomEventNameOk returns a tuple with the CustomEventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetCustomEventNameOk() (*string, bool) {
	if o == nil || isNil(o.CustomEventName) {
		return nil, false
	}
	return o.CustomEventName, true
}

// HasCustomEventName returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasCustomEventName() bool {
	if o != nil && !isNil(o.CustomEventName) {
		return true
	}

	return false
}

// SetCustomEventName gets a reference to the given string and assigns it to the CustomEventName field.
func (o *WebsiteMonitoringBeacon) SetCustomEventName(v string) {
	o.CustomEventName = &v
}

// GetDecodedBodySize returns the DecodedBodySize field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetDecodedBodySize() int64 {
	if o == nil || isNil(o.DecodedBodySize) {
		var ret int64
		return ret
	}
	return *o.DecodedBodySize
}

// GetDecodedBodySizeOk returns a tuple with the DecodedBodySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetDecodedBodySizeOk() (*int64, bool) {
	if o == nil || isNil(o.DecodedBodySize) {
		return nil, false
	}
	return o.DecodedBodySize, true
}

// HasDecodedBodySize returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasDecodedBodySize() bool {
	if o != nil && !isNil(o.DecodedBodySize) {
		return true
	}

	return false
}

// SetDecodedBodySize gets a reference to the given int64 and assigns it to the DecodedBodySize field.
func (o *WebsiteMonitoringBeacon) SetDecodedBodySize(v int64) {
	o.DecodedBodySize = &v
}

// GetDeprecations returns the Deprecations field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetDeprecations() []string {
	if o == nil || isNil(o.Deprecations) {
		var ret []string
		return ret
	}
	return o.Deprecations
}

// GetDeprecationsOk returns a tuple with the Deprecations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetDeprecationsOk() ([]string, bool) {
	if o == nil || isNil(o.Deprecations) {
		return nil, false
	}
	return o.Deprecations, true
}

// HasDeprecations returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasDeprecations() bool {
	if o != nil && !isNil(o.Deprecations) {
		return true
	}

	return false
}

// SetDeprecations gets a reference to the given []string and assigns it to the Deprecations field.
func (o *WebsiteMonitoringBeacon) SetDeprecations(v []string) {
	o.Deprecations = v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetDeviceType() string {
	if o == nil || isNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetDeviceTypeOk() (*string, bool) {
	if o == nil || isNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasDeviceType() bool {
	if o != nil && !isNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *WebsiteMonitoringBeacon) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetDnsTime returns the DnsTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetDnsTime() int64 {
	if o == nil || isNil(o.DnsTime) {
		var ret int64
		return ret
	}
	return *o.DnsTime
}

// GetDnsTimeOk returns a tuple with the DnsTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetDnsTimeOk() (*int64, bool) {
	if o == nil || isNil(o.DnsTime) {
		return nil, false
	}
	return o.DnsTime, true
}

// HasDnsTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasDnsTime() bool {
	if o != nil && !isNil(o.DnsTime) {
		return true
	}

	return false
}

// SetDnsTime gets a reference to the given int64 and assigns it to the DnsTime field.
func (o *WebsiteMonitoringBeacon) SetDnsTime(v int64) {
	o.DnsTime = &v
}

// GetDomTime returns the DomTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetDomTime() int64 {
	if o == nil || isNil(o.DomTime) {
		var ret int64
		return ret
	}
	return *o.DomTime
}

// GetDomTimeOk returns a tuple with the DomTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetDomTimeOk() (*int64, bool) {
	if o == nil || isNil(o.DomTime) {
		return nil, false
	}
	return o.DomTime, true
}

// HasDomTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasDomTime() bool {
	if o != nil && !isNil(o.DomTime) {
		return true
	}

	return false
}

// SetDomTime gets a reference to the given int64 and assigns it to the DomTime field.
func (o *WebsiteMonitoringBeacon) SetDomTime(v int64) {
	o.DomTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetDuration() int64 {
	if o == nil || isNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetDurationOk() (*int64, bool) {
	if o == nil || isNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *WebsiteMonitoringBeacon) SetDuration(v int64) {
	o.Duration = &v
}

// GetEncodedBodySize returns the EncodedBodySize field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetEncodedBodySize() int64 {
	if o == nil || isNil(o.EncodedBodySize) {
		var ret int64
		return ret
	}
	return *o.EncodedBodySize
}

// GetEncodedBodySizeOk returns a tuple with the EncodedBodySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetEncodedBodySizeOk() (*int64, bool) {
	if o == nil || isNil(o.EncodedBodySize) {
		return nil, false
	}
	return o.EncodedBodySize, true
}

// HasEncodedBodySize returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasEncodedBodySize() bool {
	if o != nil && !isNil(o.EncodedBodySize) {
		return true
	}

	return false
}

// SetEncodedBodySize gets a reference to the given int64 and assigns it to the EncodedBodySize field.
func (o *WebsiteMonitoringBeacon) SetEncodedBodySize(v int64) {
	o.EncodedBodySize = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetErrorCount() int64 {
	if o == nil || isNil(o.ErrorCount) {
		var ret int64
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetErrorCountOk() (*int64, bool) {
	if o == nil || isNil(o.ErrorCount) {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasErrorCount() bool {
	if o != nil && !isNil(o.ErrorCount) {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int64 and assigns it to the ErrorCount field.
func (o *WebsiteMonitoringBeacon) SetErrorCount(v int64) {
	o.ErrorCount = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetErrorId() string {
	if o == nil || isNil(o.ErrorId) {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetErrorIdOk() (*string, bool) {
	if o == nil || isNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasErrorId() bool {
	if o != nil && !isNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *WebsiteMonitoringBeacon) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *WebsiteMonitoringBeacon) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetErrorType() string {
	if o == nil || isNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetErrorTypeOk() (*string, bool) {
	if o == nil || isNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasErrorType() bool {
	if o != nil && !isNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *WebsiteMonitoringBeacon) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetFirstContentfulPaintTime returns the FirstContentfulPaintTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetFirstContentfulPaintTime() int64 {
	if o == nil || isNil(o.FirstContentfulPaintTime) {
		var ret int64
		return ret
	}
	return *o.FirstContentfulPaintTime
}

// GetFirstContentfulPaintTimeOk returns a tuple with the FirstContentfulPaintTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetFirstContentfulPaintTimeOk() (*int64, bool) {
	if o == nil || isNil(o.FirstContentfulPaintTime) {
		return nil, false
	}
	return o.FirstContentfulPaintTime, true
}

// HasFirstContentfulPaintTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasFirstContentfulPaintTime() bool {
	if o != nil && !isNil(o.FirstContentfulPaintTime) {
		return true
	}

	return false
}

// SetFirstContentfulPaintTime gets a reference to the given int64 and assigns it to the FirstContentfulPaintTime field.
func (o *WebsiteMonitoringBeacon) SetFirstContentfulPaintTime(v int64) {
	o.FirstContentfulPaintTime = &v
}

// GetFirstInputDelayTime returns the FirstInputDelayTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetFirstInputDelayTime() int64 {
	if o == nil || isNil(o.FirstInputDelayTime) {
		var ret int64
		return ret
	}
	return *o.FirstInputDelayTime
}

// GetFirstInputDelayTimeOk returns a tuple with the FirstInputDelayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetFirstInputDelayTimeOk() (*int64, bool) {
	if o == nil || isNil(o.FirstInputDelayTime) {
		return nil, false
	}
	return o.FirstInputDelayTime, true
}

// HasFirstInputDelayTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasFirstInputDelayTime() bool {
	if o != nil && !isNil(o.FirstInputDelayTime) {
		return true
	}

	return false
}

// SetFirstInputDelayTime gets a reference to the given int64 and assigns it to the FirstInputDelayTime field.
func (o *WebsiteMonitoringBeacon) SetFirstInputDelayTime(v int64) {
	o.FirstInputDelayTime = &v
}

// GetFirstPaintTime returns the FirstPaintTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetFirstPaintTime() int64 {
	if o == nil || isNil(o.FirstPaintTime) {
		var ret int64
		return ret
	}
	return *o.FirstPaintTime
}

// GetFirstPaintTimeOk returns a tuple with the FirstPaintTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetFirstPaintTimeOk() (*int64, bool) {
	if o == nil || isNil(o.FirstPaintTime) {
		return nil, false
	}
	return o.FirstPaintTime, true
}

// HasFirstPaintTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasFirstPaintTime() bool {
	if o != nil && !isNil(o.FirstPaintTime) {
		return true
	}

	return false
}

// SetFirstPaintTime gets a reference to the given int64 and assigns it to the FirstPaintTime field.
func (o *WebsiteMonitoringBeacon) SetFirstPaintTime(v int64) {
	o.FirstPaintTime = &v
}

// GetFrontendTime returns the FrontendTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetFrontendTime() int64 {
	if o == nil || isNil(o.FrontendTime) {
		var ret int64
		return ret
	}
	return *o.FrontendTime
}

// GetFrontendTimeOk returns a tuple with the FrontendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetFrontendTimeOk() (*int64, bool) {
	if o == nil || isNil(o.FrontendTime) {
		return nil, false
	}
	return o.FrontendTime, true
}

// HasFrontendTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasFrontendTime() bool {
	if o != nil && !isNil(o.FrontendTime) {
		return true
	}

	return false
}

// SetFrontendTime gets a reference to the given int64 and assigns it to the FrontendTime field.
func (o *WebsiteMonitoringBeacon) SetFrontendTime(v int64) {
	o.FrontendTime = &v
}

// GetGraphqlOperationName returns the GraphqlOperationName field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetGraphqlOperationName() string {
	if o == nil || isNil(o.GraphqlOperationName) {
		var ret string
		return ret
	}
	return *o.GraphqlOperationName
}

// GetGraphqlOperationNameOk returns a tuple with the GraphqlOperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetGraphqlOperationNameOk() (*string, bool) {
	if o == nil || isNil(o.GraphqlOperationName) {
		return nil, false
	}
	return o.GraphqlOperationName, true
}

// HasGraphqlOperationName returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasGraphqlOperationName() bool {
	if o != nil && !isNil(o.GraphqlOperationName) {
		return true
	}

	return false
}

// SetGraphqlOperationName gets a reference to the given string and assigns it to the GraphqlOperationName field.
func (o *WebsiteMonitoringBeacon) SetGraphqlOperationName(v string) {
	o.GraphqlOperationName = &v
}

// GetGraphqlOperationType returns the GraphqlOperationType field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetGraphqlOperationType() string {
	if o == nil || isNil(o.GraphqlOperationType) {
		var ret string
		return ret
	}
	return *o.GraphqlOperationType
}

// GetGraphqlOperationTypeOk returns a tuple with the GraphqlOperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetGraphqlOperationTypeOk() (*string, bool) {
	if o == nil || isNil(o.GraphqlOperationType) {
		return nil, false
	}
	return o.GraphqlOperationType, true
}

// HasGraphqlOperationType returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasGraphqlOperationType() bool {
	if o != nil && !isNil(o.GraphqlOperationType) {
		return true
	}

	return false
}

// SetGraphqlOperationType gets a reference to the given string and assigns it to the GraphqlOperationType field.
func (o *WebsiteMonitoringBeacon) SetGraphqlOperationType(v string) {
	o.GraphqlOperationType = &v
}

// GetHttpCallAsynchronous returns the HttpCallAsynchronous field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetHttpCallAsynchronous() bool {
	if o == nil || isNil(o.HttpCallAsynchronous) {
		var ret bool
		return ret
	}
	return *o.HttpCallAsynchronous
}

// GetHttpCallAsynchronousOk returns a tuple with the HttpCallAsynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetHttpCallAsynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.HttpCallAsynchronous) {
		return nil, false
	}
	return o.HttpCallAsynchronous, true
}

// HasHttpCallAsynchronous returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasHttpCallAsynchronous() bool {
	if o != nil && !isNil(o.HttpCallAsynchronous) {
		return true
	}

	return false
}

// SetHttpCallAsynchronous gets a reference to the given bool and assigns it to the HttpCallAsynchronous field.
func (o *WebsiteMonitoringBeacon) SetHttpCallAsynchronous(v bool) {
	o.HttpCallAsynchronous = &v
}

// GetHttpCallCorrelationAttempted returns the HttpCallCorrelationAttempted field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetHttpCallCorrelationAttempted() bool {
	if o == nil || isNil(o.HttpCallCorrelationAttempted) {
		var ret bool
		return ret
	}
	return *o.HttpCallCorrelationAttempted
}

// GetHttpCallCorrelationAttemptedOk returns a tuple with the HttpCallCorrelationAttempted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetHttpCallCorrelationAttemptedOk() (*bool, bool) {
	if o == nil || isNil(o.HttpCallCorrelationAttempted) {
		return nil, false
	}
	return o.HttpCallCorrelationAttempted, true
}

// HasHttpCallCorrelationAttempted returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasHttpCallCorrelationAttempted() bool {
	if o != nil && !isNil(o.HttpCallCorrelationAttempted) {
		return true
	}

	return false
}

// SetHttpCallCorrelationAttempted gets a reference to the given bool and assigns it to the HttpCallCorrelationAttempted field.
func (o *WebsiteMonitoringBeacon) SetHttpCallCorrelationAttempted(v bool) {
	o.HttpCallCorrelationAttempted = &v
}

// GetHttpCallHeaders returns the HttpCallHeaders field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetHttpCallHeaders() map[string]string {
	if o == nil || isNil(o.HttpCallHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.HttpCallHeaders
}

// GetHttpCallHeadersOk returns a tuple with the HttpCallHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetHttpCallHeadersOk() (*map[string]string, bool) {
	if o == nil || isNil(o.HttpCallHeaders) {
		return nil, false
	}
	return o.HttpCallHeaders, true
}

// HasHttpCallHeaders returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasHttpCallHeaders() bool {
	if o != nil && !isNil(o.HttpCallHeaders) {
		return true
	}

	return false
}

// SetHttpCallHeaders gets a reference to the given map[string]string and assigns it to the HttpCallHeaders field.
func (o *WebsiteMonitoringBeacon) SetHttpCallHeaders(v map[string]string) {
	o.HttpCallHeaders = &v
}

// GetHttpCallMethod returns the HttpCallMethod field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetHttpCallMethod() string {
	if o == nil || isNil(o.HttpCallMethod) {
		var ret string
		return ret
	}
	return *o.HttpCallMethod
}

// GetHttpCallMethodOk returns a tuple with the HttpCallMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetHttpCallMethodOk() (*string, bool) {
	if o == nil || isNil(o.HttpCallMethod) {
		return nil, false
	}
	return o.HttpCallMethod, true
}

// HasHttpCallMethod returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasHttpCallMethod() bool {
	if o != nil && !isNil(o.HttpCallMethod) {
		return true
	}

	return false
}

// SetHttpCallMethod gets a reference to the given string and assigns it to the HttpCallMethod field.
func (o *WebsiteMonitoringBeacon) SetHttpCallMethod(v string) {
	o.HttpCallMethod = &v
}

// GetHttpCallOrigin returns the HttpCallOrigin field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetHttpCallOrigin() string {
	if o == nil || isNil(o.HttpCallOrigin) {
		var ret string
		return ret
	}
	return *o.HttpCallOrigin
}

// GetHttpCallOriginOk returns a tuple with the HttpCallOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetHttpCallOriginOk() (*string, bool) {
	if o == nil || isNil(o.HttpCallOrigin) {
		return nil, false
	}
	return o.HttpCallOrigin, true
}

// HasHttpCallOrigin returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasHttpCallOrigin() bool {
	if o != nil && !isNil(o.HttpCallOrigin) {
		return true
	}

	return false
}

// SetHttpCallOrigin gets a reference to the given string and assigns it to the HttpCallOrigin field.
func (o *WebsiteMonitoringBeacon) SetHttpCallOrigin(v string) {
	o.HttpCallOrigin = &v
}

// GetHttpCallPath returns the HttpCallPath field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetHttpCallPath() string {
	if o == nil || isNil(o.HttpCallPath) {
		var ret string
		return ret
	}
	return *o.HttpCallPath
}

// GetHttpCallPathOk returns a tuple with the HttpCallPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetHttpCallPathOk() (*string, bool) {
	if o == nil || isNil(o.HttpCallPath) {
		return nil, false
	}
	return o.HttpCallPath, true
}

// HasHttpCallPath returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasHttpCallPath() bool {
	if o != nil && !isNil(o.HttpCallPath) {
		return true
	}

	return false
}

// SetHttpCallPath gets a reference to the given string and assigns it to the HttpCallPath field.
func (o *WebsiteMonitoringBeacon) SetHttpCallPath(v string) {
	o.HttpCallPath = &v
}

// GetHttpCallStatus returns the HttpCallStatus field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetHttpCallStatus() int32 {
	if o == nil || isNil(o.HttpCallStatus) {
		var ret int32
		return ret
	}
	return *o.HttpCallStatus
}

// GetHttpCallStatusOk returns a tuple with the HttpCallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetHttpCallStatusOk() (*int32, bool) {
	if o == nil || isNil(o.HttpCallStatus) {
		return nil, false
	}
	return o.HttpCallStatus, true
}

// HasHttpCallStatus returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasHttpCallStatus() bool {
	if o != nil && !isNil(o.HttpCallStatus) {
		return true
	}

	return false
}

// SetHttpCallStatus gets a reference to the given int32 and assigns it to the HttpCallStatus field.
func (o *WebsiteMonitoringBeacon) SetHttpCallStatus(v int32) {
	o.HttpCallStatus = &v
}

// GetHttpCallUrl returns the HttpCallUrl field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetHttpCallUrl() string {
	if o == nil || isNil(o.HttpCallUrl) {
		var ret string
		return ret
	}
	return *o.HttpCallUrl
}

// GetHttpCallUrlOk returns a tuple with the HttpCallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetHttpCallUrlOk() (*string, bool) {
	if o == nil || isNil(o.HttpCallUrl) {
		return nil, false
	}
	return o.HttpCallUrl, true
}

// HasHttpCallUrl returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasHttpCallUrl() bool {
	if o != nil && !isNil(o.HttpCallUrl) {
		return true
	}

	return false
}

// SetHttpCallUrl gets a reference to the given string and assigns it to the HttpCallUrl field.
func (o *WebsiteMonitoringBeacon) SetHttpCallUrl(v string) {
	o.HttpCallUrl = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetInitiator() string {
	if o == nil || isNil(o.Initiator) {
		var ret string
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetInitiatorOk() (*string, bool) {
	if o == nil || isNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasInitiator() bool {
	if o != nil && !isNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given string and assigns it to the Initiator field.
func (o *WebsiteMonitoringBeacon) SetInitiator(v string) {
	o.Initiator = &v
}

// GetLargestContentfulPaintTime returns the LargestContentfulPaintTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetLargestContentfulPaintTime() int64 {
	if o == nil || isNil(o.LargestContentfulPaintTime) {
		var ret int64
		return ret
	}
	return *o.LargestContentfulPaintTime
}

// GetLargestContentfulPaintTimeOk returns a tuple with the LargestContentfulPaintTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetLargestContentfulPaintTimeOk() (*int64, bool) {
	if o == nil || isNil(o.LargestContentfulPaintTime) {
		return nil, false
	}
	return o.LargestContentfulPaintTime, true
}

// HasLargestContentfulPaintTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasLargestContentfulPaintTime() bool {
	if o != nil && !isNil(o.LargestContentfulPaintTime) {
		return true
	}

	return false
}

// SetLargestContentfulPaintTime gets a reference to the given int64 and assigns it to the LargestContentfulPaintTime field.
func (o *WebsiteMonitoringBeacon) SetLargestContentfulPaintTime(v int64) {
	o.LargestContentfulPaintTime = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetLatitude() float64 {
	if o == nil || isNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetLatitudeOk() (*float64, bool) {
	if o == nil || isNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasLatitude() bool {
	if o != nil && !isNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *WebsiteMonitoringBeacon) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLocationOrigin returns the LocationOrigin field value
func (o *WebsiteMonitoringBeacon) GetLocationOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocationOrigin
}

// GetLocationOriginOk returns a tuple with the LocationOrigin field value
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetLocationOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationOrigin, true
}

// SetLocationOrigin sets field value
func (o *WebsiteMonitoringBeacon) SetLocationOrigin(v string) {
	o.LocationOrigin = v
}

// GetLocationPath returns the LocationPath field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetLocationPath() string {
	if o == nil || isNil(o.LocationPath) {
		var ret string
		return ret
	}
	return *o.LocationPath
}

// GetLocationPathOk returns a tuple with the LocationPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetLocationPathOk() (*string, bool) {
	if o == nil || isNil(o.LocationPath) {
		return nil, false
	}
	return o.LocationPath, true
}

// HasLocationPath returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasLocationPath() bool {
	if o != nil && !isNil(o.LocationPath) {
		return true
	}

	return false
}

// SetLocationPath gets a reference to the given string and assigns it to the LocationPath field.
func (o *WebsiteMonitoringBeacon) SetLocationPath(v string) {
	o.LocationPath = &v
}

// GetLocationUrl returns the LocationUrl field value
func (o *WebsiteMonitoringBeacon) GetLocationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocationUrl
}

// GetLocationUrlOk returns a tuple with the LocationUrl field value
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetLocationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationUrl, true
}

// SetLocationUrl sets field value
func (o *WebsiteMonitoringBeacon) SetLocationUrl(v string) {
	o.LocationUrl = v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetLongitude() float64 {
	if o == nil || isNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetLongitudeOk() (*float64, bool) {
	if o == nil || isNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasLongitude() bool {
	if o != nil && !isNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *WebsiteMonitoringBeacon) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetMeta() map[string]string {
	if o == nil || isNil(o.Meta) {
		var ret map[string]string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetMetaOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]string and assigns it to the Meta field.
func (o *WebsiteMonitoringBeacon) SetMeta(v map[string]string) {
	o.Meta = &v
}

// GetOnLoadTime returns the OnLoadTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetOnLoadTime() int64 {
	if o == nil || isNil(o.OnLoadTime) {
		var ret int64
		return ret
	}
	return *o.OnLoadTime
}

// GetOnLoadTimeOk returns a tuple with the OnLoadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetOnLoadTimeOk() (*int64, bool) {
	if o == nil || isNil(o.OnLoadTime) {
		return nil, false
	}
	return o.OnLoadTime, true
}

// HasOnLoadTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasOnLoadTime() bool {
	if o != nil && !isNil(o.OnLoadTime) {
		return true
	}

	return false
}

// SetOnLoadTime gets a reference to the given int64 and assigns it to the OnLoadTime field.
func (o *WebsiteMonitoringBeacon) SetOnLoadTime(v int64) {
	o.OnLoadTime = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetOsName() string {
	if o == nil || isNil(o.OsName) {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetOsNameOk() (*string, bool) {
	if o == nil || isNil(o.OsName) {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasOsName() bool {
	if o != nil && !isNil(o.OsName) {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *WebsiteMonitoringBeacon) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetOsVersion() string {
	if o == nil || isNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetOsVersionOk() (*string, bool) {
	if o == nil || isNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasOsVersion() bool {
	if o != nil && !isNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *WebsiteMonitoringBeacon) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetPage() string {
	if o == nil || isNil(o.Page) {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetPageOk() (*string, bool) {
	if o == nil || isNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *WebsiteMonitoringBeacon) SetPage(v string) {
	o.Page = &v
}

// GetPageLoadId returns the PageLoadId field value
func (o *WebsiteMonitoringBeacon) GetPageLoadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PageLoadId
}

// GetPageLoadIdOk returns a tuple with the PageLoadId field value
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetPageLoadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageLoadId, true
}

// SetPageLoadId sets field value
func (o *WebsiteMonitoringBeacon) SetPageLoadId(v string) {
	o.PageLoadId = v
}

// GetParsedStackTrace returns the ParsedStackTrace field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetParsedStackTrace() []StackTraceLine {
	if o == nil || isNil(o.ParsedStackTrace) {
		var ret []StackTraceLine
		return ret
	}
	return o.ParsedStackTrace
}

// GetParsedStackTraceOk returns a tuple with the ParsedStackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetParsedStackTraceOk() ([]StackTraceLine, bool) {
	if o == nil || isNil(o.ParsedStackTrace) {
		return nil, false
	}
	return o.ParsedStackTrace, true
}

// HasParsedStackTrace returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasParsedStackTrace() bool {
	if o != nil && !isNil(o.ParsedStackTrace) {
		return true
	}

	return false
}

// SetParsedStackTrace gets a reference to the given []StackTraceLine and assigns it to the ParsedStackTrace field.
func (o *WebsiteMonitoringBeacon) SetParsedStackTrace(v []StackTraceLine) {
	o.ParsedStackTrace = v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetPhase() string {
	if o == nil || isNil(o.Phase) {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetPhaseOk() (*string, bool) {
	if o == nil || isNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasPhase() bool {
	if o != nil && !isNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *WebsiteMonitoringBeacon) SetPhase(v string) {
	o.Phase = &v
}

// GetProcessingTime returns the ProcessingTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetProcessingTime() int64 {
	if o == nil || isNil(o.ProcessingTime) {
		var ret int64
		return ret
	}
	return *o.ProcessingTime
}

// GetProcessingTimeOk returns a tuple with the ProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.ProcessingTime) {
		return nil, false
	}
	return o.ProcessingTime, true
}

// HasProcessingTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasProcessingTime() bool {
	if o != nil && !isNil(o.ProcessingTime) {
		return true
	}

	return false
}

// SetProcessingTime gets a reference to the given int64 and assigns it to the ProcessingTime field.
func (o *WebsiteMonitoringBeacon) SetProcessingTime(v int64) {
	o.ProcessingTime = &v
}

// GetRedirectTime returns the RedirectTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetRedirectTime() int64 {
	if o == nil || isNil(o.RedirectTime) {
		var ret int64
		return ret
	}
	return *o.RedirectTime
}

// GetRedirectTimeOk returns a tuple with the RedirectTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetRedirectTimeOk() (*int64, bool) {
	if o == nil || isNil(o.RedirectTime) {
		return nil, false
	}
	return o.RedirectTime, true
}

// HasRedirectTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasRedirectTime() bool {
	if o != nil && !isNil(o.RedirectTime) {
		return true
	}

	return false
}

// SetRedirectTime gets a reference to the given int64 and assigns it to the RedirectTime field.
func (o *WebsiteMonitoringBeacon) SetRedirectTime(v int64) {
	o.RedirectTime = &v
}

// GetRequestTime returns the RequestTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetRequestTime() int64 {
	if o == nil || isNil(o.RequestTime) {
		var ret int64
		return ret
	}
	return *o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetRequestTimeOk() (*int64, bool) {
	if o == nil || isNil(o.RequestTime) {
		return nil, false
	}
	return o.RequestTime, true
}

// HasRequestTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasRequestTime() bool {
	if o != nil && !isNil(o.RequestTime) {
		return true
	}

	return false
}

// SetRequestTime gets a reference to the given int64 and assigns it to the RequestTime field.
func (o *WebsiteMonitoringBeacon) SetRequestTime(v int64) {
	o.RequestTime = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetResourceType() string {
	if o == nil || isNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *WebsiteMonitoringBeacon) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetResponseTime() int64 {
	if o == nil || isNil(o.ResponseTime) {
		var ret int64
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetResponseTimeOk() (*int64, bool) {
	if o == nil || isNil(o.ResponseTime) {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasResponseTime() bool {
	if o != nil && !isNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given int64 and assigns it to the ResponseTime field.
func (o *WebsiteMonitoringBeacon) SetResponseTime(v int64) {
	o.ResponseTime = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetSessionId() string {
	if o == nil || isNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetSessionIdOk() (*string, bool) {
	if o == nil || isNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasSessionId() bool {
	if o != nil && !isNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *WebsiteMonitoringBeacon) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSnippetVersion returns the SnippetVersion field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetSnippetVersion() string {
	if o == nil || isNil(o.SnippetVersion) {
		var ret string
		return ret
	}
	return *o.SnippetVersion
}

// GetSnippetVersionOk returns a tuple with the SnippetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetSnippetVersionOk() (*string, bool) {
	if o == nil || isNil(o.SnippetVersion) {
		return nil, false
	}
	return o.SnippetVersion, true
}

// HasSnippetVersion returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasSnippetVersion() bool {
	if o != nil && !isNil(o.SnippetVersion) {
		return true
	}

	return false
}

// SetSnippetVersion gets a reference to the given string and assigns it to the SnippetVersion field.
func (o *WebsiteMonitoringBeacon) SetSnippetVersion(v string) {
	o.SnippetVersion = &v
}

// GetSslTime returns the SslTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetSslTime() int64 {
	if o == nil || isNil(o.SslTime) {
		var ret int64
		return ret
	}
	return *o.SslTime
}

// GetSslTimeOk returns a tuple with the SslTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetSslTimeOk() (*int64, bool) {
	if o == nil || isNil(o.SslTime) {
		return nil, false
	}
	return o.SslTime, true
}

// HasSslTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasSslTime() bool {
	if o != nil && !isNil(o.SslTime) {
		return true
	}

	return false
}

// SetSslTime gets a reference to the given int64 and assigns it to the SslTime field.
func (o *WebsiteMonitoringBeacon) SetSslTime(v int64) {
	o.SslTime = &v
}

// GetStackTrace returns the StackTrace field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetStackTrace() string {
	if o == nil || isNil(o.StackTrace) {
		var ret string
		return ret
	}
	return *o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetStackTraceOk() (*string, bool) {
	if o == nil || isNil(o.StackTrace) {
		return nil, false
	}
	return o.StackTrace, true
}

// HasStackTrace returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasStackTrace() bool {
	if o != nil && !isNil(o.StackTrace) {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given string and assigns it to the StackTrace field.
func (o *WebsiteMonitoringBeacon) SetStackTrace(v string) {
	o.StackTrace = &v
}

// GetStackTraceParsingStatus returns the StackTraceParsingStatus field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetStackTraceParsingStatus() int32 {
	if o == nil || isNil(o.StackTraceParsingStatus) {
		var ret int32
		return ret
	}
	return *o.StackTraceParsingStatus
}

// GetStackTraceParsingStatusOk returns a tuple with the StackTraceParsingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetStackTraceParsingStatusOk() (*int32, bool) {
	if o == nil || isNil(o.StackTraceParsingStatus) {
		return nil, false
	}
	return o.StackTraceParsingStatus, true
}

// HasStackTraceParsingStatus returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasStackTraceParsingStatus() bool {
	if o != nil && !isNil(o.StackTraceParsingStatus) {
		return true
	}

	return false
}

// SetStackTraceParsingStatus gets a reference to the given int32 and assigns it to the StackTraceParsingStatus field.
func (o *WebsiteMonitoringBeacon) SetStackTraceParsingStatus(v int32) {
	o.StackTraceParsingStatus = &v
}

// GetStackTraceReadability returns the StackTraceReadability field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetStackTraceReadability() int32 {
	if o == nil || isNil(o.StackTraceReadability) {
		var ret int32
		return ret
	}
	return *o.StackTraceReadability
}

// GetStackTraceReadabilityOk returns a tuple with the StackTraceReadability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetStackTraceReadabilityOk() (*int32, bool) {
	if o == nil || isNil(o.StackTraceReadability) {
		return nil, false
	}
	return o.StackTraceReadability, true
}

// HasStackTraceReadability returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasStackTraceReadability() bool {
	if o != nil && !isNil(o.StackTraceReadability) {
		return true
	}

	return false
}

// SetStackTraceReadability gets a reference to the given int32 and assigns it to the StackTraceReadability field.
func (o *WebsiteMonitoringBeacon) SetStackTraceReadability(v int32) {
	o.StackTraceReadability = &v
}

// GetSubdivision returns the Subdivision field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetSubdivision() string {
	if o == nil || isNil(o.Subdivision) {
		var ret string
		return ret
	}
	return *o.Subdivision
}

// GetSubdivisionOk returns a tuple with the Subdivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetSubdivisionOk() (*string, bool) {
	if o == nil || isNil(o.Subdivision) {
		return nil, false
	}
	return o.Subdivision, true
}

// HasSubdivision returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasSubdivision() bool {
	if o != nil && !isNil(o.Subdivision) {
		return true
	}

	return false
}

// SetSubdivision gets a reference to the given string and assigns it to the Subdivision field.
func (o *WebsiteMonitoringBeacon) SetSubdivision(v string) {
	o.Subdivision = &v
}

// GetSubdivisionCode returns the SubdivisionCode field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetSubdivisionCode() string {
	if o == nil || isNil(o.SubdivisionCode) {
		var ret string
		return ret
	}
	return *o.SubdivisionCode
}

// GetSubdivisionCodeOk returns a tuple with the SubdivisionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetSubdivisionCodeOk() (*string, bool) {
	if o == nil || isNil(o.SubdivisionCode) {
		return nil, false
	}
	return o.SubdivisionCode, true
}

// HasSubdivisionCode returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasSubdivisionCode() bool {
	if o != nil && !isNil(o.SubdivisionCode) {
		return true
	}

	return false
}

// SetSubdivisionCode gets a reference to the given string and assigns it to the SubdivisionCode field.
func (o *WebsiteMonitoringBeacon) SetSubdivisionCode(v string) {
	o.SubdivisionCode = &v
}

// GetTcpTime returns the TcpTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetTcpTime() int64 {
	if o == nil || isNil(o.TcpTime) {
		var ret int64
		return ret
	}
	return *o.TcpTime
}

// GetTcpTimeOk returns a tuple with the TcpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetTcpTimeOk() (*int64, bool) {
	if o == nil || isNil(o.TcpTime) {
		return nil, false
	}
	return o.TcpTime, true
}

// HasTcpTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasTcpTime() bool {
	if o != nil && !isNil(o.TcpTime) {
		return true
	}

	return false
}

// SetTcpTime gets a reference to the given int64 and assigns it to the TcpTime field.
func (o *WebsiteMonitoringBeacon) SetTcpTime(v int64) {
	o.TcpTime = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetTimestamp() int64 {
	if o == nil || isNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *WebsiteMonitoringBeacon) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTransferSize returns the TransferSize field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetTransferSize() int64 {
	if o == nil || isNil(o.TransferSize) {
		var ret int64
		return ret
	}
	return *o.TransferSize
}

// GetTransferSizeOk returns a tuple with the TransferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetTransferSizeOk() (*int64, bool) {
	if o == nil || isNil(o.TransferSize) {
		return nil, false
	}
	return o.TransferSize, true
}

// HasTransferSize returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasTransferSize() bool {
	if o != nil && !isNil(o.TransferSize) {
		return true
	}

	return false
}

// SetTransferSize gets a reference to the given int64 and assigns it to the TransferSize field.
func (o *WebsiteMonitoringBeacon) SetTransferSize(v int64) {
	o.TransferSize = &v
}

// GetType returns the Type field value
func (o *WebsiteMonitoringBeacon) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebsiteMonitoringBeacon) SetType(v string) {
	o.Type = v
}

// GetUnloadTime returns the UnloadTime field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetUnloadTime() int64 {
	if o == nil || isNil(o.UnloadTime) {
		var ret int64
		return ret
	}
	return *o.UnloadTime
}

// GetUnloadTimeOk returns a tuple with the UnloadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetUnloadTimeOk() (*int64, bool) {
	if o == nil || isNil(o.UnloadTime) {
		return nil, false
	}
	return o.UnloadTime, true
}

// HasUnloadTime returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasUnloadTime() bool {
	if o != nil && !isNil(o.UnloadTime) {
		return true
	}

	return false
}

// SetUnloadTime gets a reference to the given int64 and assigns it to the UnloadTime field.
func (o *WebsiteMonitoringBeacon) SetUnloadTime(v int64) {
	o.UnloadTime = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetUserEmail() string {
	if o == nil || isNil(o.UserEmail) {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetUserEmailOk() (*string, bool) {
	if o == nil || isNil(o.UserEmail) {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasUserEmail() bool {
	if o != nil && !isNil(o.UserEmail) {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *WebsiteMonitoringBeacon) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *WebsiteMonitoringBeacon) SetUserId(v string) {
	o.UserId = &v
}

// GetUserIp returns the UserIp field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetUserIp() string {
	if o == nil || isNil(o.UserIp) {
		var ret string
		return ret
	}
	return *o.UserIp
}

// GetUserIpOk returns a tuple with the UserIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetUserIpOk() (*string, bool) {
	if o == nil || isNil(o.UserIp) {
		return nil, false
	}
	return o.UserIp, true
}

// HasUserIp returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasUserIp() bool {
	if o != nil && !isNil(o.UserIp) {
		return true
	}

	return false
}

// SetUserIp gets a reference to the given string and assigns it to the UserIp field.
func (o *WebsiteMonitoringBeacon) SetUserIp(v string) {
	o.UserIp = &v
}

// GetUserLanguages returns the UserLanguages field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetUserLanguages() []string {
	if o == nil || isNil(o.UserLanguages) {
		var ret []string
		return ret
	}
	return o.UserLanguages
}

// GetUserLanguagesOk returns a tuple with the UserLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetUserLanguagesOk() ([]string, bool) {
	if o == nil || isNil(o.UserLanguages) {
		return nil, false
	}
	return o.UserLanguages, true
}

// HasUserLanguages returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasUserLanguages() bool {
	if o != nil && !isNil(o.UserLanguages) {
		return true
	}

	return false
}

// SetUserLanguages gets a reference to the given []string and assigns it to the UserLanguages field.
func (o *WebsiteMonitoringBeacon) SetUserLanguages(v []string) {
	o.UserLanguages = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetUserName() string {
	if o == nil || isNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetUserNameOk() (*string, bool) {
	if o == nil || isNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasUserName() bool {
	if o != nil && !isNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *WebsiteMonitoringBeacon) SetUserName(v string) {
	o.UserName = &v
}

// GetWebsiteId returns the WebsiteId field value
func (o *WebsiteMonitoringBeacon) GetWebsiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebsiteId
}

// GetWebsiteIdOk returns a tuple with the WebsiteId field value
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetWebsiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsiteId, true
}

// SetWebsiteId sets field value
func (o *WebsiteMonitoringBeacon) SetWebsiteId(v string) {
	o.WebsiteId = v
}

// GetWebsiteLabel returns the WebsiteLabel field value
func (o *WebsiteMonitoringBeacon) GetWebsiteLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebsiteLabel
}

// GetWebsiteLabelOk returns a tuple with the WebsiteLabel field value
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetWebsiteLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsiteLabel, true
}

// SetWebsiteLabel sets field value
func (o *WebsiteMonitoringBeacon) SetWebsiteLabel(v string) {
	o.WebsiteLabel = v
}

// GetWindowHeight returns the WindowHeight field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetWindowHeight() int32 {
	if o == nil || isNil(o.WindowHeight) {
		var ret int32
		return ret
	}
	return *o.WindowHeight
}

// GetWindowHeightOk returns a tuple with the WindowHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetWindowHeightOk() (*int32, bool) {
	if o == nil || isNil(o.WindowHeight) {
		return nil, false
	}
	return o.WindowHeight, true
}

// HasWindowHeight returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasWindowHeight() bool {
	if o != nil && !isNil(o.WindowHeight) {
		return true
	}

	return false
}

// SetWindowHeight gets a reference to the given int32 and assigns it to the WindowHeight field.
func (o *WebsiteMonitoringBeacon) SetWindowHeight(v int32) {
	o.WindowHeight = &v
}

// GetWindowHidden returns the WindowHidden field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetWindowHidden() bool {
	if o == nil || isNil(o.WindowHidden) {
		var ret bool
		return ret
	}
	return *o.WindowHidden
}

// GetWindowHiddenOk returns a tuple with the WindowHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetWindowHiddenOk() (*bool, bool) {
	if o == nil || isNil(o.WindowHidden) {
		return nil, false
	}
	return o.WindowHidden, true
}

// HasWindowHidden returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasWindowHidden() bool {
	if o != nil && !isNil(o.WindowHidden) {
		return true
	}

	return false
}

// SetWindowHidden gets a reference to the given bool and assigns it to the WindowHidden field.
func (o *WebsiteMonitoringBeacon) SetWindowHidden(v bool) {
	o.WindowHidden = &v
}

// GetWindowWidth returns the WindowWidth field value if set, zero value otherwise.
func (o *WebsiteMonitoringBeacon) GetWindowWidth() int32 {
	if o == nil || isNil(o.WindowWidth) {
		var ret int32
		return ret
	}
	return *o.WindowWidth
}

// GetWindowWidthOk returns a tuple with the WindowWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteMonitoringBeacon) GetWindowWidthOk() (*int32, bool) {
	if o == nil || isNil(o.WindowWidth) {
		return nil, false
	}
	return o.WindowWidth, true
}

// HasWindowWidth returns a boolean if a field has been set.
func (o *WebsiteMonitoringBeacon) HasWindowWidth() bool {
	if o != nil && !isNil(o.WindowWidth) {
		return true
	}

	return false
}

// SetWindowWidth gets a reference to the given int32 and assigns it to the WindowWidth field.
func (o *WebsiteMonitoringBeacon) SetWindowWidth(v int32) {
	o.WindowWidth = &v
}

func (o WebsiteMonitoringBeacon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccuracyRadius) {
		toSerialize["accuracyRadius"] = o.AccuracyRadius
	}
	if !isNil(o.AccurateTimingsAvailable) {
		toSerialize["accurateTimingsAvailable"] = o.AccurateTimingsAvailable
	}
	if !isNil(o.AppCacheTime) {
		toSerialize["appCacheTime"] = o.AppCacheTime
	}
	if !isNil(o.BackendTime) {
		toSerialize["backendTime"] = o.BackendTime
	}
	if !isNil(o.BackendTraceId) {
		toSerialize["backendTraceId"] = o.BackendTraceId
	}
	if !isNil(o.BatchSize) {
		toSerialize["batchSize"] = o.BatchSize
	}
	if true {
		toSerialize["beaconId"] = o.BeaconId
	}
	if !isNil(o.BrowserName) {
		toSerialize["browserName"] = o.BrowserName
	}
	if !isNil(o.BrowserVersion) {
		toSerialize["browserVersion"] = o.BrowserVersion
	}
	if !isNil(o.CacheInteraction) {
		toSerialize["cacheInteraction"] = o.CacheInteraction
	}
	if !isNil(o.ChildrenTime) {
		toSerialize["childrenTime"] = o.ChildrenTime
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.ClockSkew) {
		toSerialize["clockSkew"] = o.ClockSkew
	}
	if !isNil(o.ComponentStack) {
		toSerialize["componentStack"] = o.ComponentStack
	}
	if !isNil(o.ConnectionType) {
		toSerialize["connectionType"] = o.ConnectionType
	}
	if !isNil(o.Continent) {
		toSerialize["continent"] = o.Continent
	}
	if !isNil(o.ContinentCode) {
		toSerialize["continentCode"] = o.ContinentCode
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !isNil(o.CspBlockedUri) {
		toSerialize["cspBlockedUri"] = o.CspBlockedUri
	}
	if !isNil(o.CspColumnNumber) {
		toSerialize["cspColumnNumber"] = o.CspColumnNumber
	}
	if !isNil(o.CspDisposition) {
		toSerialize["cspDisposition"] = o.CspDisposition
	}
	if !isNil(o.CspEffectiveDirective) {
		toSerialize["cspEffectiveDirective"] = o.CspEffectiveDirective
	}
	if !isNil(o.CspLineNumber) {
		toSerialize["cspLineNumber"] = o.CspLineNumber
	}
	if !isNil(o.CspOriginalPolicy) {
		toSerialize["cspOriginalPolicy"] = o.CspOriginalPolicy
	}
	if !isNil(o.CspSample) {
		toSerialize["cspSample"] = o.CspSample
	}
	if !isNil(o.CspSourceFile) {
		toSerialize["cspSourceFile"] = o.CspSourceFile
	}
	if !isNil(o.CumulativeLayoutShift) {
		toSerialize["cumulativeLayoutShift"] = o.CumulativeLayoutShift
	}
	if !isNil(o.CustomEventName) {
		toSerialize["customEventName"] = o.CustomEventName
	}
	if !isNil(o.DecodedBodySize) {
		toSerialize["decodedBodySize"] = o.DecodedBodySize
	}
	if !isNil(o.Deprecations) {
		toSerialize["deprecations"] = o.Deprecations
	}
	if !isNil(o.DeviceType) {
		toSerialize["deviceType"] = o.DeviceType
	}
	if !isNil(o.DnsTime) {
		toSerialize["dnsTime"] = o.DnsTime
	}
	if !isNil(o.DomTime) {
		toSerialize["domTime"] = o.DomTime
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.EncodedBodySize) {
		toSerialize["encodedBodySize"] = o.EncodedBodySize
	}
	if !isNil(o.ErrorCount) {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if !isNil(o.ErrorId) {
		toSerialize["errorId"] = o.ErrorId
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !isNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	if !isNil(o.FirstContentfulPaintTime) {
		toSerialize["firstContentfulPaintTime"] = o.FirstContentfulPaintTime
	}
	if !isNil(o.FirstInputDelayTime) {
		toSerialize["firstInputDelayTime"] = o.FirstInputDelayTime
	}
	if !isNil(o.FirstPaintTime) {
		toSerialize["firstPaintTime"] = o.FirstPaintTime
	}
	if !isNil(o.FrontendTime) {
		toSerialize["frontendTime"] = o.FrontendTime
	}
	if !isNil(o.GraphqlOperationName) {
		toSerialize["graphqlOperationName"] = o.GraphqlOperationName
	}
	if !isNil(o.GraphqlOperationType) {
		toSerialize["graphqlOperationType"] = o.GraphqlOperationType
	}
	if !isNil(o.HttpCallAsynchronous) {
		toSerialize["httpCallAsynchronous"] = o.HttpCallAsynchronous
	}
	if !isNil(o.HttpCallCorrelationAttempted) {
		toSerialize["httpCallCorrelationAttempted"] = o.HttpCallCorrelationAttempted
	}
	if !isNil(o.HttpCallHeaders) {
		toSerialize["httpCallHeaders"] = o.HttpCallHeaders
	}
	if !isNil(o.HttpCallMethod) {
		toSerialize["httpCallMethod"] = o.HttpCallMethod
	}
	if !isNil(o.HttpCallOrigin) {
		toSerialize["httpCallOrigin"] = o.HttpCallOrigin
	}
	if !isNil(o.HttpCallPath) {
		toSerialize["httpCallPath"] = o.HttpCallPath
	}
	if !isNil(o.HttpCallStatus) {
		toSerialize["httpCallStatus"] = o.HttpCallStatus
	}
	if !isNil(o.HttpCallUrl) {
		toSerialize["httpCallUrl"] = o.HttpCallUrl
	}
	if !isNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	if !isNil(o.LargestContentfulPaintTime) {
		toSerialize["largestContentfulPaintTime"] = o.LargestContentfulPaintTime
	}
	if !isNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if true {
		toSerialize["locationOrigin"] = o.LocationOrigin
	}
	if !isNil(o.LocationPath) {
		toSerialize["locationPath"] = o.LocationPath
	}
	if true {
		toSerialize["locationUrl"] = o.LocationUrl
	}
	if !isNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.OnLoadTime) {
		toSerialize["onLoadTime"] = o.OnLoadTime
	}
	if !isNil(o.OsName) {
		toSerialize["osName"] = o.OsName
	}
	if !isNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["pageLoadId"] = o.PageLoadId
	}
	if !isNil(o.ParsedStackTrace) {
		toSerialize["parsedStackTrace"] = o.ParsedStackTrace
	}
	if !isNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	if !isNil(o.ProcessingTime) {
		toSerialize["processingTime"] = o.ProcessingTime
	}
	if !isNil(o.RedirectTime) {
		toSerialize["redirectTime"] = o.RedirectTime
	}
	if !isNil(o.RequestTime) {
		toSerialize["requestTime"] = o.RequestTime
	}
	if !isNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !isNil(o.ResponseTime) {
		toSerialize["responseTime"] = o.ResponseTime
	}
	if !isNil(o.SessionId) {
		toSerialize["sessionId"] = o.SessionId
	}
	if !isNil(o.SnippetVersion) {
		toSerialize["snippetVersion"] = o.SnippetVersion
	}
	if !isNil(o.SslTime) {
		toSerialize["sslTime"] = o.SslTime
	}
	if !isNil(o.StackTrace) {
		toSerialize["stackTrace"] = o.StackTrace
	}
	if !isNil(o.StackTraceParsingStatus) {
		toSerialize["stackTraceParsingStatus"] = o.StackTraceParsingStatus
	}
	if !isNil(o.StackTraceReadability) {
		toSerialize["stackTraceReadability"] = o.StackTraceReadability
	}
	if !isNil(o.Subdivision) {
		toSerialize["subdivision"] = o.Subdivision
	}
	if !isNil(o.SubdivisionCode) {
		toSerialize["subdivisionCode"] = o.SubdivisionCode
	}
	if !isNil(o.TcpTime) {
		toSerialize["tcpTime"] = o.TcpTime
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.TransferSize) {
		toSerialize["transferSize"] = o.TransferSize
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UnloadTime) {
		toSerialize["unloadTime"] = o.UnloadTime
	}
	if !isNil(o.UserEmail) {
		toSerialize["userEmail"] = o.UserEmail
	}
	if !isNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !isNil(o.UserIp) {
		toSerialize["userIp"] = o.UserIp
	}
	if !isNil(o.UserLanguages) {
		toSerialize["userLanguages"] = o.UserLanguages
	}
	if !isNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if true {
		toSerialize["websiteId"] = o.WebsiteId
	}
	if true {
		toSerialize["websiteLabel"] = o.WebsiteLabel
	}
	if !isNil(o.WindowHeight) {
		toSerialize["windowHeight"] = o.WindowHeight
	}
	if !isNil(o.WindowHidden) {
		toSerialize["windowHidden"] = o.WindowHidden
	}
	if !isNil(o.WindowWidth) {
		toSerialize["windowWidth"] = o.WindowWidth
	}
	return json.Marshal(toSerialize)
}

type NullableWebsiteMonitoringBeacon struct {
	value *WebsiteMonitoringBeacon
	isSet bool
}

func (v NullableWebsiteMonitoringBeacon) Get() *WebsiteMonitoringBeacon {
	return v.value
}

func (v *NullableWebsiteMonitoringBeacon) Set(val *WebsiteMonitoringBeacon) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsiteMonitoringBeacon) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsiteMonitoringBeacon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsiteMonitoringBeacon(val *WebsiteMonitoringBeacon) *NullableWebsiteMonitoringBeacon {
	return &NullableWebsiteMonitoringBeacon{value: val, isSet: true}
}

func (v NullableWebsiteMonitoringBeacon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsiteMonitoringBeacon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
