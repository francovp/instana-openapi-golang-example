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

// MobileAppMonitoringBeacon struct for MobileAppMonitoringBeacon
type MobileAppMonitoringBeacon struct {
	AccuracyRadius            *int64             `json:"accuracyRadius,omitempty"`
	AgentVersion              *string            `json:"agentVersion,omitempty"`
	AppBuild                  *string            `json:"appBuild,omitempty"`
	AppVersion                *string            `json:"appVersion,omitempty"`
	BackendTraceId            *string            `json:"backendTraceId,omitempty"`
	BatchSize                 *int64             `json:"batchSize,omitempty"`
	BeaconId                  string             `json:"beaconId"`
	BundleIdentifier          *string            `json:"bundleIdentifier,omitempty"`
	Carrier                   *string            `json:"carrier,omitempty"`
	City                      *string            `json:"city,omitempty"`
	ClockSkew                 *int64             `json:"clockSkew,omitempty"`
	ConnectionType            *string            `json:"connectionType,omitempty"`
	Continent                 *string            `json:"continent,omitempty"`
	ContinentCode             *string            `json:"continentCode,omitempty"`
	Country                   *string            `json:"country,omitempty"`
	CountryCode               *string            `json:"countryCode,omitempty"`
	CustomEventName           *string            `json:"customEventName,omitempty"`
	DecodedBodySize           *int64             `json:"decodedBodySize,omitempty"`
	DeviceHardware            *string            `json:"deviceHardware,omitempty"`
	DeviceManufacturer        *string            `json:"deviceManufacturer,omitempty"`
	DeviceModel               *string            `json:"deviceModel,omitempty"`
	Duration                  *int64             `json:"duration,omitempty"`
	EffectiveConnectionType   *string            `json:"effectiveConnectionType,omitempty"`
	EncodedBodySize           *int64             `json:"encodedBodySize,omitempty"`
	Environment               *string            `json:"environment,omitempty"`
	ErrorCount                *int64             `json:"errorCount,omitempty"`
	ErrorId                   *string            `json:"errorId,omitempty"`
	ErrorMessage              *string            `json:"errorMessage,omitempty"`
	ErrorType                 *string            `json:"errorType,omitempty"`
	GooglePlayServicesMissing *bool              `json:"googlePlayServicesMissing,omitempty"`
	HttpCallHeaders           *map[string]string `json:"httpCallHeaders,omitempty"`
	HttpCallMethod            *string            `json:"httpCallMethod,omitempty"`
	HttpCallOrigin            *string            `json:"httpCallOrigin,omitempty"`
	HttpCallPath              *string            `json:"httpCallPath,omitempty"`
	HttpCallStatus            *int32             `json:"httpCallStatus,omitempty"`
	HttpCallUrl               *string            `json:"httpCallUrl,omitempty"`
	IngestionTime             *int64             `json:"ingestionTime,omitempty"`
	Latitude                  *float64           `json:"latitude,omitempty"`
	Longitude                 *float64           `json:"longitude,omitempty"`
	Meta                      *map[string]string `json:"meta,omitempty"`
	MobileAppId               string             `json:"mobileAppId"`
	MobileAppLabel            *string            `json:"mobileAppLabel,omitempty"`
	OsName                    *string            `json:"osName,omitempty"`
	OsVersion                 *string            `json:"osVersion,omitempty"`
	Platform                  *string            `json:"platform,omitempty"`
	Region                    *string            `json:"region,omitempty"`
	Rooted                    *bool              `json:"rooted,omitempty"`
	SessionId                 string             `json:"sessionId"`
	StackTrace                *string            `json:"stackTrace,omitempty"`
	Subdivision               *string            `json:"subdivision,omitempty"`
	SubdivisionCode           *string            `json:"subdivisionCode,omitempty"`
	Tenant                    *string            `json:"tenant,omitempty"`
	Timestamp                 *int64             `json:"timestamp,omitempty"`
	TransferSize              *int64             `json:"transferSize,omitempty"`
	Type                      string             `json:"type"`
	Unit                      *string            `json:"unit,omitempty"`
	UserEmail                 *string            `json:"userEmail,omitempty"`
	UserId                    *string            `json:"userId,omitempty"`
	UserIp                    *string            `json:"userIp,omitempty"`
	UserLanguages             []string           `json:"userLanguages,omitempty"`
	UserName                  *string            `json:"userName,omitempty"`
	View                      *string            `json:"view,omitempty"`
	ViewportHeight            *int32             `json:"viewportHeight,omitempty"`
	ViewportWidth             *int32             `json:"viewportWidth,omitempty"`
}

// NewMobileAppMonitoringBeacon instantiates a new MobileAppMonitoringBeacon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileAppMonitoringBeacon(beaconId string, mobileAppId string, sessionId string, type_ string) *MobileAppMonitoringBeacon {
	this := MobileAppMonitoringBeacon{}
	this.BeaconId = beaconId
	this.MobileAppId = mobileAppId
	this.SessionId = sessionId
	this.Type = type_
	return &this
}

// NewMobileAppMonitoringBeaconWithDefaults instantiates a new MobileAppMonitoringBeacon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileAppMonitoringBeaconWithDefaults() *MobileAppMonitoringBeacon {
	this := MobileAppMonitoringBeacon{}
	return &this
}

// GetAccuracyRadius returns the AccuracyRadius field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetAccuracyRadius() int64 {
	if o == nil || isNil(o.AccuracyRadius) {
		var ret int64
		return ret
	}
	return *o.AccuracyRadius
}

// GetAccuracyRadiusOk returns a tuple with the AccuracyRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetAccuracyRadiusOk() (*int64, bool) {
	if o == nil || isNil(o.AccuracyRadius) {
		return nil, false
	}
	return o.AccuracyRadius, true
}

// HasAccuracyRadius returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasAccuracyRadius() bool {
	if o != nil && !isNil(o.AccuracyRadius) {
		return true
	}

	return false
}

// SetAccuracyRadius gets a reference to the given int64 and assigns it to the AccuracyRadius field.
func (o *MobileAppMonitoringBeacon) SetAccuracyRadius(v int64) {
	o.AccuracyRadius = &v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetAgentVersion() string {
	if o == nil || isNil(o.AgentVersion) {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetAgentVersionOk() (*string, bool) {
	if o == nil || isNil(o.AgentVersion) {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasAgentVersion() bool {
	if o != nil && !isNil(o.AgentVersion) {
		return true
	}

	return false
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *MobileAppMonitoringBeacon) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetAppBuild returns the AppBuild field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetAppBuild() string {
	if o == nil || isNil(o.AppBuild) {
		var ret string
		return ret
	}
	return *o.AppBuild
}

// GetAppBuildOk returns a tuple with the AppBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetAppBuildOk() (*string, bool) {
	if o == nil || isNil(o.AppBuild) {
		return nil, false
	}
	return o.AppBuild, true
}

// HasAppBuild returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasAppBuild() bool {
	if o != nil && !isNil(o.AppBuild) {
		return true
	}

	return false
}

// SetAppBuild gets a reference to the given string and assigns it to the AppBuild field.
func (o *MobileAppMonitoringBeacon) SetAppBuild(v string) {
	o.AppBuild = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetAppVersion() string {
	if o == nil || isNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetAppVersionOk() (*string, bool) {
	if o == nil || isNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasAppVersion() bool {
	if o != nil && !isNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *MobileAppMonitoringBeacon) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBackendTraceId returns the BackendTraceId field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetBackendTraceId() string {
	if o == nil || isNil(o.BackendTraceId) {
		var ret string
		return ret
	}
	return *o.BackendTraceId
}

// GetBackendTraceIdOk returns a tuple with the BackendTraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetBackendTraceIdOk() (*string, bool) {
	if o == nil || isNil(o.BackendTraceId) {
		return nil, false
	}
	return o.BackendTraceId, true
}

// HasBackendTraceId returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasBackendTraceId() bool {
	if o != nil && !isNil(o.BackendTraceId) {
		return true
	}

	return false
}

// SetBackendTraceId gets a reference to the given string and assigns it to the BackendTraceId field.
func (o *MobileAppMonitoringBeacon) SetBackendTraceId(v string) {
	o.BackendTraceId = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetBatchSize() int64 {
	if o == nil || isNil(o.BatchSize) {
		var ret int64
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetBatchSizeOk() (*int64, bool) {
	if o == nil || isNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasBatchSize() bool {
	if o != nil && !isNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int64 and assigns it to the BatchSize field.
func (o *MobileAppMonitoringBeacon) SetBatchSize(v int64) {
	o.BatchSize = &v
}

// GetBeaconId returns the BeaconId field value
func (o *MobileAppMonitoringBeacon) GetBeaconId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeaconId
}

// GetBeaconIdOk returns a tuple with the BeaconId field value
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetBeaconIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeaconId, true
}

// SetBeaconId sets field value
func (o *MobileAppMonitoringBeacon) SetBeaconId(v string) {
	o.BeaconId = v
}

// GetBundleIdentifier returns the BundleIdentifier field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetBundleIdentifier() string {
	if o == nil || isNil(o.BundleIdentifier) {
		var ret string
		return ret
	}
	return *o.BundleIdentifier
}

// GetBundleIdentifierOk returns a tuple with the BundleIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetBundleIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.BundleIdentifier) {
		return nil, false
	}
	return o.BundleIdentifier, true
}

// HasBundleIdentifier returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasBundleIdentifier() bool {
	if o != nil && !isNil(o.BundleIdentifier) {
		return true
	}

	return false
}

// SetBundleIdentifier gets a reference to the given string and assigns it to the BundleIdentifier field.
func (o *MobileAppMonitoringBeacon) SetBundleIdentifier(v string) {
	o.BundleIdentifier = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetCarrier() string {
	if o == nil || isNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetCarrierOk() (*string, bool) {
	if o == nil || isNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasCarrier() bool {
	if o != nil && !isNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *MobileAppMonitoringBeacon) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *MobileAppMonitoringBeacon) SetCity(v string) {
	o.City = &v
}

// GetClockSkew returns the ClockSkew field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetClockSkew() int64 {
	if o == nil || isNil(o.ClockSkew) {
		var ret int64
		return ret
	}
	return *o.ClockSkew
}

// GetClockSkewOk returns a tuple with the ClockSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetClockSkewOk() (*int64, bool) {
	if o == nil || isNil(o.ClockSkew) {
		return nil, false
	}
	return o.ClockSkew, true
}

// HasClockSkew returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasClockSkew() bool {
	if o != nil && !isNil(o.ClockSkew) {
		return true
	}

	return false
}

// SetClockSkew gets a reference to the given int64 and assigns it to the ClockSkew field.
func (o *MobileAppMonitoringBeacon) SetClockSkew(v int64) {
	o.ClockSkew = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetConnectionType() string {
	if o == nil || isNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetConnectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasConnectionType() bool {
	if o != nil && !isNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *MobileAppMonitoringBeacon) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetContinent() string {
	if o == nil || isNil(o.Continent) {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetContinentOk() (*string, bool) {
	if o == nil || isNil(o.Continent) {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasContinent() bool {
	if o != nil && !isNil(o.Continent) {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *MobileAppMonitoringBeacon) SetContinent(v string) {
	o.Continent = &v
}

// GetContinentCode returns the ContinentCode field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetContinentCode() string {
	if o == nil || isNil(o.ContinentCode) {
		var ret string
		return ret
	}
	return *o.ContinentCode
}

// GetContinentCodeOk returns a tuple with the ContinentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetContinentCodeOk() (*string, bool) {
	if o == nil || isNil(o.ContinentCode) {
		return nil, false
	}
	return o.ContinentCode, true
}

// HasContinentCode returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasContinentCode() bool {
	if o != nil && !isNil(o.ContinentCode) {
		return true
	}

	return false
}

// SetContinentCode gets a reference to the given string and assigns it to the ContinentCode field.
func (o *MobileAppMonitoringBeacon) SetContinentCode(v string) {
	o.ContinentCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *MobileAppMonitoringBeacon) SetCountry(v string) {
	o.Country = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetCountryCode() string {
	if o == nil || isNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetCountryCodeOk() (*string, bool) {
	if o == nil || isNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasCountryCode() bool {
	if o != nil && !isNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *MobileAppMonitoringBeacon) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCustomEventName returns the CustomEventName field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetCustomEventName() string {
	if o == nil || isNil(o.CustomEventName) {
		var ret string
		return ret
	}
	return *o.CustomEventName
}

// GetCustomEventNameOk returns a tuple with the CustomEventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetCustomEventNameOk() (*string, bool) {
	if o == nil || isNil(o.CustomEventName) {
		return nil, false
	}
	return o.CustomEventName, true
}

// HasCustomEventName returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasCustomEventName() bool {
	if o != nil && !isNil(o.CustomEventName) {
		return true
	}

	return false
}

// SetCustomEventName gets a reference to the given string and assigns it to the CustomEventName field.
func (o *MobileAppMonitoringBeacon) SetCustomEventName(v string) {
	o.CustomEventName = &v
}

// GetDecodedBodySize returns the DecodedBodySize field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetDecodedBodySize() int64 {
	if o == nil || isNil(o.DecodedBodySize) {
		var ret int64
		return ret
	}
	return *o.DecodedBodySize
}

// GetDecodedBodySizeOk returns a tuple with the DecodedBodySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetDecodedBodySizeOk() (*int64, bool) {
	if o == nil || isNil(o.DecodedBodySize) {
		return nil, false
	}
	return o.DecodedBodySize, true
}

// HasDecodedBodySize returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasDecodedBodySize() bool {
	if o != nil && !isNil(o.DecodedBodySize) {
		return true
	}

	return false
}

// SetDecodedBodySize gets a reference to the given int64 and assigns it to the DecodedBodySize field.
func (o *MobileAppMonitoringBeacon) SetDecodedBodySize(v int64) {
	o.DecodedBodySize = &v
}

// GetDeviceHardware returns the DeviceHardware field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetDeviceHardware() string {
	if o == nil || isNil(o.DeviceHardware) {
		var ret string
		return ret
	}
	return *o.DeviceHardware
}

// GetDeviceHardwareOk returns a tuple with the DeviceHardware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetDeviceHardwareOk() (*string, bool) {
	if o == nil || isNil(o.DeviceHardware) {
		return nil, false
	}
	return o.DeviceHardware, true
}

// HasDeviceHardware returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasDeviceHardware() bool {
	if o != nil && !isNil(o.DeviceHardware) {
		return true
	}

	return false
}

// SetDeviceHardware gets a reference to the given string and assigns it to the DeviceHardware field.
func (o *MobileAppMonitoringBeacon) SetDeviceHardware(v string) {
	o.DeviceHardware = &v
}

// GetDeviceManufacturer returns the DeviceManufacturer field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetDeviceManufacturer() string {
	if o == nil || isNil(o.DeviceManufacturer) {
		var ret string
		return ret
	}
	return *o.DeviceManufacturer
}

// GetDeviceManufacturerOk returns a tuple with the DeviceManufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetDeviceManufacturerOk() (*string, bool) {
	if o == nil || isNil(o.DeviceManufacturer) {
		return nil, false
	}
	return o.DeviceManufacturer, true
}

// HasDeviceManufacturer returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasDeviceManufacturer() bool {
	if o != nil && !isNil(o.DeviceManufacturer) {
		return true
	}

	return false
}

// SetDeviceManufacturer gets a reference to the given string and assigns it to the DeviceManufacturer field.
func (o *MobileAppMonitoringBeacon) SetDeviceManufacturer(v string) {
	o.DeviceManufacturer = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetDeviceModel() string {
	if o == nil || isNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetDeviceModelOk() (*string, bool) {
	if o == nil || isNil(o.DeviceModel) {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasDeviceModel() bool {
	if o != nil && !isNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *MobileAppMonitoringBeacon) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetDuration() int64 {
	if o == nil || isNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetDurationOk() (*int64, bool) {
	if o == nil || isNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *MobileAppMonitoringBeacon) SetDuration(v int64) {
	o.Duration = &v
}

// GetEffectiveConnectionType returns the EffectiveConnectionType field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetEffectiveConnectionType() string {
	if o == nil || isNil(o.EffectiveConnectionType) {
		var ret string
		return ret
	}
	return *o.EffectiveConnectionType
}

// GetEffectiveConnectionTypeOk returns a tuple with the EffectiveConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetEffectiveConnectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.EffectiveConnectionType) {
		return nil, false
	}
	return o.EffectiveConnectionType, true
}

// HasEffectiveConnectionType returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasEffectiveConnectionType() bool {
	if o != nil && !isNil(o.EffectiveConnectionType) {
		return true
	}

	return false
}

// SetEffectiveConnectionType gets a reference to the given string and assigns it to the EffectiveConnectionType field.
func (o *MobileAppMonitoringBeacon) SetEffectiveConnectionType(v string) {
	o.EffectiveConnectionType = &v
}

// GetEncodedBodySize returns the EncodedBodySize field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetEncodedBodySize() int64 {
	if o == nil || isNil(o.EncodedBodySize) {
		var ret int64
		return ret
	}
	return *o.EncodedBodySize
}

// GetEncodedBodySizeOk returns a tuple with the EncodedBodySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetEncodedBodySizeOk() (*int64, bool) {
	if o == nil || isNil(o.EncodedBodySize) {
		return nil, false
	}
	return o.EncodedBodySize, true
}

// HasEncodedBodySize returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasEncodedBodySize() bool {
	if o != nil && !isNil(o.EncodedBodySize) {
		return true
	}

	return false
}

// SetEncodedBodySize gets a reference to the given int64 and assigns it to the EncodedBodySize field.
func (o *MobileAppMonitoringBeacon) SetEncodedBodySize(v int64) {
	o.EncodedBodySize = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetEnvironment() string {
	if o == nil || isNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetEnvironmentOk() (*string, bool) {
	if o == nil || isNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *MobileAppMonitoringBeacon) SetEnvironment(v string) {
	o.Environment = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetErrorCount() int64 {
	if o == nil || isNil(o.ErrorCount) {
		var ret int64
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetErrorCountOk() (*int64, bool) {
	if o == nil || isNil(o.ErrorCount) {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasErrorCount() bool {
	if o != nil && !isNil(o.ErrorCount) {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int64 and assigns it to the ErrorCount field.
func (o *MobileAppMonitoringBeacon) SetErrorCount(v int64) {
	o.ErrorCount = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetErrorId() string {
	if o == nil || isNil(o.ErrorId) {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetErrorIdOk() (*string, bool) {
	if o == nil || isNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasErrorId() bool {
	if o != nil && !isNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *MobileAppMonitoringBeacon) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *MobileAppMonitoringBeacon) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetErrorType() string {
	if o == nil || isNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetErrorTypeOk() (*string, bool) {
	if o == nil || isNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasErrorType() bool {
	if o != nil && !isNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *MobileAppMonitoringBeacon) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetGooglePlayServicesMissing returns the GooglePlayServicesMissing field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetGooglePlayServicesMissing() bool {
	if o == nil || isNil(o.GooglePlayServicesMissing) {
		var ret bool
		return ret
	}
	return *o.GooglePlayServicesMissing
}

// GetGooglePlayServicesMissingOk returns a tuple with the GooglePlayServicesMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetGooglePlayServicesMissingOk() (*bool, bool) {
	if o == nil || isNil(o.GooglePlayServicesMissing) {
		return nil, false
	}
	return o.GooglePlayServicesMissing, true
}

// HasGooglePlayServicesMissing returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasGooglePlayServicesMissing() bool {
	if o != nil && !isNil(o.GooglePlayServicesMissing) {
		return true
	}

	return false
}

// SetGooglePlayServicesMissing gets a reference to the given bool and assigns it to the GooglePlayServicesMissing field.
func (o *MobileAppMonitoringBeacon) SetGooglePlayServicesMissing(v bool) {
	o.GooglePlayServicesMissing = &v
}

// GetHttpCallHeaders returns the HttpCallHeaders field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetHttpCallHeaders() map[string]string {
	if o == nil || isNil(o.HttpCallHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.HttpCallHeaders
}

// GetHttpCallHeadersOk returns a tuple with the HttpCallHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetHttpCallHeadersOk() (*map[string]string, bool) {
	if o == nil || isNil(o.HttpCallHeaders) {
		return nil, false
	}
	return o.HttpCallHeaders, true
}

// HasHttpCallHeaders returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasHttpCallHeaders() bool {
	if o != nil && !isNil(o.HttpCallHeaders) {
		return true
	}

	return false
}

// SetHttpCallHeaders gets a reference to the given map[string]string and assigns it to the HttpCallHeaders field.
func (o *MobileAppMonitoringBeacon) SetHttpCallHeaders(v map[string]string) {
	o.HttpCallHeaders = &v
}

// GetHttpCallMethod returns the HttpCallMethod field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetHttpCallMethod() string {
	if o == nil || isNil(o.HttpCallMethod) {
		var ret string
		return ret
	}
	return *o.HttpCallMethod
}

// GetHttpCallMethodOk returns a tuple with the HttpCallMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetHttpCallMethodOk() (*string, bool) {
	if o == nil || isNil(o.HttpCallMethod) {
		return nil, false
	}
	return o.HttpCallMethod, true
}

// HasHttpCallMethod returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasHttpCallMethod() bool {
	if o != nil && !isNil(o.HttpCallMethod) {
		return true
	}

	return false
}

// SetHttpCallMethod gets a reference to the given string and assigns it to the HttpCallMethod field.
func (o *MobileAppMonitoringBeacon) SetHttpCallMethod(v string) {
	o.HttpCallMethod = &v
}

// GetHttpCallOrigin returns the HttpCallOrigin field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetHttpCallOrigin() string {
	if o == nil || isNil(o.HttpCallOrigin) {
		var ret string
		return ret
	}
	return *o.HttpCallOrigin
}

// GetHttpCallOriginOk returns a tuple with the HttpCallOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetHttpCallOriginOk() (*string, bool) {
	if o == nil || isNil(o.HttpCallOrigin) {
		return nil, false
	}
	return o.HttpCallOrigin, true
}

// HasHttpCallOrigin returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasHttpCallOrigin() bool {
	if o != nil && !isNil(o.HttpCallOrigin) {
		return true
	}

	return false
}

// SetHttpCallOrigin gets a reference to the given string and assigns it to the HttpCallOrigin field.
func (o *MobileAppMonitoringBeacon) SetHttpCallOrigin(v string) {
	o.HttpCallOrigin = &v
}

// GetHttpCallPath returns the HttpCallPath field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetHttpCallPath() string {
	if o == nil || isNil(o.HttpCallPath) {
		var ret string
		return ret
	}
	return *o.HttpCallPath
}

// GetHttpCallPathOk returns a tuple with the HttpCallPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetHttpCallPathOk() (*string, bool) {
	if o == nil || isNil(o.HttpCallPath) {
		return nil, false
	}
	return o.HttpCallPath, true
}

// HasHttpCallPath returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasHttpCallPath() bool {
	if o != nil && !isNil(o.HttpCallPath) {
		return true
	}

	return false
}

// SetHttpCallPath gets a reference to the given string and assigns it to the HttpCallPath field.
func (o *MobileAppMonitoringBeacon) SetHttpCallPath(v string) {
	o.HttpCallPath = &v
}

// GetHttpCallStatus returns the HttpCallStatus field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetHttpCallStatus() int32 {
	if o == nil || isNil(o.HttpCallStatus) {
		var ret int32
		return ret
	}
	return *o.HttpCallStatus
}

// GetHttpCallStatusOk returns a tuple with the HttpCallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetHttpCallStatusOk() (*int32, bool) {
	if o == nil || isNil(o.HttpCallStatus) {
		return nil, false
	}
	return o.HttpCallStatus, true
}

// HasHttpCallStatus returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasHttpCallStatus() bool {
	if o != nil && !isNil(o.HttpCallStatus) {
		return true
	}

	return false
}

// SetHttpCallStatus gets a reference to the given int32 and assigns it to the HttpCallStatus field.
func (o *MobileAppMonitoringBeacon) SetHttpCallStatus(v int32) {
	o.HttpCallStatus = &v
}

// GetHttpCallUrl returns the HttpCallUrl field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetHttpCallUrl() string {
	if o == nil || isNil(o.HttpCallUrl) {
		var ret string
		return ret
	}
	return *o.HttpCallUrl
}

// GetHttpCallUrlOk returns a tuple with the HttpCallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetHttpCallUrlOk() (*string, bool) {
	if o == nil || isNil(o.HttpCallUrl) {
		return nil, false
	}
	return o.HttpCallUrl, true
}

// HasHttpCallUrl returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasHttpCallUrl() bool {
	if o != nil && !isNil(o.HttpCallUrl) {
		return true
	}

	return false
}

// SetHttpCallUrl gets a reference to the given string and assigns it to the HttpCallUrl field.
func (o *MobileAppMonitoringBeacon) SetHttpCallUrl(v string) {
	o.HttpCallUrl = &v
}

// GetIngestionTime returns the IngestionTime field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetIngestionTime() int64 {
	if o == nil || isNil(o.IngestionTime) {
		var ret int64
		return ret
	}
	return *o.IngestionTime
}

// GetIngestionTimeOk returns a tuple with the IngestionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetIngestionTimeOk() (*int64, bool) {
	if o == nil || isNil(o.IngestionTime) {
		return nil, false
	}
	return o.IngestionTime, true
}

// HasIngestionTime returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasIngestionTime() bool {
	if o != nil && !isNil(o.IngestionTime) {
		return true
	}

	return false
}

// SetIngestionTime gets a reference to the given int64 and assigns it to the IngestionTime field.
func (o *MobileAppMonitoringBeacon) SetIngestionTime(v int64) {
	o.IngestionTime = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetLatitude() float64 {
	if o == nil || isNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetLatitudeOk() (*float64, bool) {
	if o == nil || isNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasLatitude() bool {
	if o != nil && !isNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *MobileAppMonitoringBeacon) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetLongitude() float64 {
	if o == nil || isNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetLongitudeOk() (*float64, bool) {
	if o == nil || isNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasLongitude() bool {
	if o != nil && !isNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *MobileAppMonitoringBeacon) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetMeta() map[string]string {
	if o == nil || isNil(o.Meta) {
		var ret map[string]string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetMetaOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]string and assigns it to the Meta field.
func (o *MobileAppMonitoringBeacon) SetMeta(v map[string]string) {
	o.Meta = &v
}

// GetMobileAppId returns the MobileAppId field value
func (o *MobileAppMonitoringBeacon) GetMobileAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MobileAppId
}

// GetMobileAppIdOk returns a tuple with the MobileAppId field value
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetMobileAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MobileAppId, true
}

// SetMobileAppId sets field value
func (o *MobileAppMonitoringBeacon) SetMobileAppId(v string) {
	o.MobileAppId = v
}

// GetMobileAppLabel returns the MobileAppLabel field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetMobileAppLabel() string {
	if o == nil || isNil(o.MobileAppLabel) {
		var ret string
		return ret
	}
	return *o.MobileAppLabel
}

// GetMobileAppLabelOk returns a tuple with the MobileAppLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetMobileAppLabelOk() (*string, bool) {
	if o == nil || isNil(o.MobileAppLabel) {
		return nil, false
	}
	return o.MobileAppLabel, true
}

// HasMobileAppLabel returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasMobileAppLabel() bool {
	if o != nil && !isNil(o.MobileAppLabel) {
		return true
	}

	return false
}

// SetMobileAppLabel gets a reference to the given string and assigns it to the MobileAppLabel field.
func (o *MobileAppMonitoringBeacon) SetMobileAppLabel(v string) {
	o.MobileAppLabel = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetOsName() string {
	if o == nil || isNil(o.OsName) {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetOsNameOk() (*string, bool) {
	if o == nil || isNil(o.OsName) {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasOsName() bool {
	if o != nil && !isNil(o.OsName) {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *MobileAppMonitoringBeacon) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetOsVersion() string {
	if o == nil || isNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetOsVersionOk() (*string, bool) {
	if o == nil || isNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasOsVersion() bool {
	if o != nil && !isNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *MobileAppMonitoringBeacon) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetPlatform() string {
	if o == nil || isNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetPlatformOk() (*string, bool) {
	if o == nil || isNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasPlatform() bool {
	if o != nil && !isNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *MobileAppMonitoringBeacon) SetPlatform(v string) {
	o.Platform = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *MobileAppMonitoringBeacon) SetRegion(v string) {
	o.Region = &v
}

// GetRooted returns the Rooted field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetRooted() bool {
	if o == nil || isNil(o.Rooted) {
		var ret bool
		return ret
	}
	return *o.Rooted
}

// GetRootedOk returns a tuple with the Rooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetRootedOk() (*bool, bool) {
	if o == nil || isNil(o.Rooted) {
		return nil, false
	}
	return o.Rooted, true
}

// HasRooted returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasRooted() bool {
	if o != nil && !isNil(o.Rooted) {
		return true
	}

	return false
}

// SetRooted gets a reference to the given bool and assigns it to the Rooted field.
func (o *MobileAppMonitoringBeacon) SetRooted(v bool) {
	o.Rooted = &v
}

// GetSessionId returns the SessionId field value
func (o *MobileAppMonitoringBeacon) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *MobileAppMonitoringBeacon) SetSessionId(v string) {
	o.SessionId = v
}

// GetStackTrace returns the StackTrace field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetStackTrace() string {
	if o == nil || isNil(o.StackTrace) {
		var ret string
		return ret
	}
	return *o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetStackTraceOk() (*string, bool) {
	if o == nil || isNil(o.StackTrace) {
		return nil, false
	}
	return o.StackTrace, true
}

// HasStackTrace returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasStackTrace() bool {
	if o != nil && !isNil(o.StackTrace) {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given string and assigns it to the StackTrace field.
func (o *MobileAppMonitoringBeacon) SetStackTrace(v string) {
	o.StackTrace = &v
}

// GetSubdivision returns the Subdivision field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetSubdivision() string {
	if o == nil || isNil(o.Subdivision) {
		var ret string
		return ret
	}
	return *o.Subdivision
}

// GetSubdivisionOk returns a tuple with the Subdivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetSubdivisionOk() (*string, bool) {
	if o == nil || isNil(o.Subdivision) {
		return nil, false
	}
	return o.Subdivision, true
}

// HasSubdivision returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasSubdivision() bool {
	if o != nil && !isNil(o.Subdivision) {
		return true
	}

	return false
}

// SetSubdivision gets a reference to the given string and assigns it to the Subdivision field.
func (o *MobileAppMonitoringBeacon) SetSubdivision(v string) {
	o.Subdivision = &v
}

// GetSubdivisionCode returns the SubdivisionCode field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetSubdivisionCode() string {
	if o == nil || isNil(o.SubdivisionCode) {
		var ret string
		return ret
	}
	return *o.SubdivisionCode
}

// GetSubdivisionCodeOk returns a tuple with the SubdivisionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetSubdivisionCodeOk() (*string, bool) {
	if o == nil || isNil(o.SubdivisionCode) {
		return nil, false
	}
	return o.SubdivisionCode, true
}

// HasSubdivisionCode returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasSubdivisionCode() bool {
	if o != nil && !isNil(o.SubdivisionCode) {
		return true
	}

	return false
}

// SetSubdivisionCode gets a reference to the given string and assigns it to the SubdivisionCode field.
func (o *MobileAppMonitoringBeacon) SetSubdivisionCode(v string) {
	o.SubdivisionCode = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetTenant() string {
	if o == nil || isNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetTenantOk() (*string, bool) {
	if o == nil || isNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasTenant() bool {
	if o != nil && !isNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *MobileAppMonitoringBeacon) SetTenant(v string) {
	o.Tenant = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetTimestamp() int64 {
	if o == nil || isNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *MobileAppMonitoringBeacon) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTransferSize returns the TransferSize field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetTransferSize() int64 {
	if o == nil || isNil(o.TransferSize) {
		var ret int64
		return ret
	}
	return *o.TransferSize
}

// GetTransferSizeOk returns a tuple with the TransferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetTransferSizeOk() (*int64, bool) {
	if o == nil || isNil(o.TransferSize) {
		return nil, false
	}
	return o.TransferSize, true
}

// HasTransferSize returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasTransferSize() bool {
	if o != nil && !isNil(o.TransferSize) {
		return true
	}

	return false
}

// SetTransferSize gets a reference to the given int64 and assigns it to the TransferSize field.
func (o *MobileAppMonitoringBeacon) SetTransferSize(v int64) {
	o.TransferSize = &v
}

// GetType returns the Type field value
func (o *MobileAppMonitoringBeacon) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MobileAppMonitoringBeacon) SetType(v string) {
	o.Type = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetUnit() string {
	if o == nil || isNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetUnitOk() (*string, bool) {
	if o == nil || isNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MobileAppMonitoringBeacon) SetUnit(v string) {
	o.Unit = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetUserEmail() string {
	if o == nil || isNil(o.UserEmail) {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetUserEmailOk() (*string, bool) {
	if o == nil || isNil(o.UserEmail) {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasUserEmail() bool {
	if o != nil && !isNil(o.UserEmail) {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *MobileAppMonitoringBeacon) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *MobileAppMonitoringBeacon) SetUserId(v string) {
	o.UserId = &v
}

// GetUserIp returns the UserIp field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetUserIp() string {
	if o == nil || isNil(o.UserIp) {
		var ret string
		return ret
	}
	return *o.UserIp
}

// GetUserIpOk returns a tuple with the UserIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetUserIpOk() (*string, bool) {
	if o == nil || isNil(o.UserIp) {
		return nil, false
	}
	return o.UserIp, true
}

// HasUserIp returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasUserIp() bool {
	if o != nil && !isNil(o.UserIp) {
		return true
	}

	return false
}

// SetUserIp gets a reference to the given string and assigns it to the UserIp field.
func (o *MobileAppMonitoringBeacon) SetUserIp(v string) {
	o.UserIp = &v
}

// GetUserLanguages returns the UserLanguages field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetUserLanguages() []string {
	if o == nil || isNil(o.UserLanguages) {
		var ret []string
		return ret
	}
	return o.UserLanguages
}

// GetUserLanguagesOk returns a tuple with the UserLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetUserLanguagesOk() ([]string, bool) {
	if o == nil || isNil(o.UserLanguages) {
		return nil, false
	}
	return o.UserLanguages, true
}

// HasUserLanguages returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasUserLanguages() bool {
	if o != nil && !isNil(o.UserLanguages) {
		return true
	}

	return false
}

// SetUserLanguages gets a reference to the given []string and assigns it to the UserLanguages field.
func (o *MobileAppMonitoringBeacon) SetUserLanguages(v []string) {
	o.UserLanguages = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetUserName() string {
	if o == nil || isNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetUserNameOk() (*string, bool) {
	if o == nil || isNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasUserName() bool {
	if o != nil && !isNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *MobileAppMonitoringBeacon) SetUserName(v string) {
	o.UserName = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetView() string {
	if o == nil || isNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetViewOk() (*string, bool) {
	if o == nil || isNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasView() bool {
	if o != nil && !isNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *MobileAppMonitoringBeacon) SetView(v string) {
	o.View = &v
}

// GetViewportHeight returns the ViewportHeight field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetViewportHeight() int32 {
	if o == nil || isNil(o.ViewportHeight) {
		var ret int32
		return ret
	}
	return *o.ViewportHeight
}

// GetViewportHeightOk returns a tuple with the ViewportHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetViewportHeightOk() (*int32, bool) {
	if o == nil || isNil(o.ViewportHeight) {
		return nil, false
	}
	return o.ViewportHeight, true
}

// HasViewportHeight returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasViewportHeight() bool {
	if o != nil && !isNil(o.ViewportHeight) {
		return true
	}

	return false
}

// SetViewportHeight gets a reference to the given int32 and assigns it to the ViewportHeight field.
func (o *MobileAppMonitoringBeacon) SetViewportHeight(v int32) {
	o.ViewportHeight = &v
}

// GetViewportWidth returns the ViewportWidth field value if set, zero value otherwise.
func (o *MobileAppMonitoringBeacon) GetViewportWidth() int32 {
	if o == nil || isNil(o.ViewportWidth) {
		var ret int32
		return ret
	}
	return *o.ViewportWidth
}

// GetViewportWidthOk returns a tuple with the ViewportWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppMonitoringBeacon) GetViewportWidthOk() (*int32, bool) {
	if o == nil || isNil(o.ViewportWidth) {
		return nil, false
	}
	return o.ViewportWidth, true
}

// HasViewportWidth returns a boolean if a field has been set.
func (o *MobileAppMonitoringBeacon) HasViewportWidth() bool {
	if o != nil && !isNil(o.ViewportWidth) {
		return true
	}

	return false
}

// SetViewportWidth gets a reference to the given int32 and assigns it to the ViewportWidth field.
func (o *MobileAppMonitoringBeacon) SetViewportWidth(v int32) {
	o.ViewportWidth = &v
}

func (o MobileAppMonitoringBeacon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccuracyRadius) {
		toSerialize["accuracyRadius"] = o.AccuracyRadius
	}
	if !isNil(o.AgentVersion) {
		toSerialize["agentVersion"] = o.AgentVersion
	}
	if !isNil(o.AppBuild) {
		toSerialize["appBuild"] = o.AppBuild
	}
	if !isNil(o.AppVersion) {
		toSerialize["appVersion"] = o.AppVersion
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
	if !isNil(o.BundleIdentifier) {
		toSerialize["bundleIdentifier"] = o.BundleIdentifier
	}
	if !isNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.ClockSkew) {
		toSerialize["clockSkew"] = o.ClockSkew
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
	if !isNil(o.CustomEventName) {
		toSerialize["customEventName"] = o.CustomEventName
	}
	if !isNil(o.DecodedBodySize) {
		toSerialize["decodedBodySize"] = o.DecodedBodySize
	}
	if !isNil(o.DeviceHardware) {
		toSerialize["deviceHardware"] = o.DeviceHardware
	}
	if !isNil(o.DeviceManufacturer) {
		toSerialize["deviceManufacturer"] = o.DeviceManufacturer
	}
	if !isNil(o.DeviceModel) {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.EffectiveConnectionType) {
		toSerialize["effectiveConnectionType"] = o.EffectiveConnectionType
	}
	if !isNil(o.EncodedBodySize) {
		toSerialize["encodedBodySize"] = o.EncodedBodySize
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
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
	if !isNil(o.GooglePlayServicesMissing) {
		toSerialize["googlePlayServicesMissing"] = o.GooglePlayServicesMissing
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
	if !isNil(o.IngestionTime) {
		toSerialize["ingestionTime"] = o.IngestionTime
	}
	if !isNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !isNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if true {
		toSerialize["mobileAppId"] = o.MobileAppId
	}
	if !isNil(o.MobileAppLabel) {
		toSerialize["mobileAppLabel"] = o.MobileAppLabel
	}
	if !isNil(o.OsName) {
		toSerialize["osName"] = o.OsName
	}
	if !isNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !isNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Rooted) {
		toSerialize["rooted"] = o.Rooted
	}
	if true {
		toSerialize["sessionId"] = o.SessionId
	}
	if !isNil(o.StackTrace) {
		toSerialize["stackTrace"] = o.StackTrace
	}
	if !isNil(o.Subdivision) {
		toSerialize["subdivision"] = o.Subdivision
	}
	if !isNil(o.SubdivisionCode) {
		toSerialize["subdivisionCode"] = o.SubdivisionCode
	}
	if !isNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
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
	if !isNil(o.Unit) {
		toSerialize["unit"] = o.Unit
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
	if !isNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !isNil(o.ViewportHeight) {
		toSerialize["viewportHeight"] = o.ViewportHeight
	}
	if !isNil(o.ViewportWidth) {
		toSerialize["viewportWidth"] = o.ViewportWidth
	}
	return json.Marshal(toSerialize)
}

type NullableMobileAppMonitoringBeacon struct {
	value *MobileAppMonitoringBeacon
	isSet bool
}

func (v NullableMobileAppMonitoringBeacon) Get() *MobileAppMonitoringBeacon {
	return v.value
}

func (v *NullableMobileAppMonitoringBeacon) Set(val *MobileAppMonitoringBeacon) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileAppMonitoringBeacon) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileAppMonitoringBeacon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileAppMonitoringBeacon(val *MobileAppMonitoringBeacon) *NullableMobileAppMonitoringBeacon {
	return &NullableMobileAppMonitoringBeacon{value: val, isSet: true}
}

func (v NullableMobileAppMonitoringBeacon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileAppMonitoringBeacon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
