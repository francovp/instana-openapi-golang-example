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

// SyntheticLocation struct for SyntheticLocation
type SyntheticLocation struct {
	CreatedAt            *int64                        `json:"createdAt,omitempty"`
	CustomProperties     *map[string]string            `json:"customProperties,omitempty"`
	Description          *string                       `json:"description,omitempty"`
	DisplayLabel         *string                       `json:"displayLabel,omitempty"`
	GeoPoint             SyntheticGeoPoint             `json:"geoPoint"`
	Id                   *string                       `json:"id,omitempty"`
	Label                string                        `json:"label"`
	LocationType         string                        `json:"locationType"`
	ModifiedAt           *int64                        `json:"modifiedAt,omitempty"`
	ObservedAt           *int64                        `json:"observedAt,omitempty"`
	PlaybackCapabilities SyntheticPlaybackCapabilities `json:"playbackCapabilities"`
	PopVersion           *string                       `json:"popVersion,omitempty"`
	Status               *string                       `json:"status,omitempty"`
	TotalTests           *int32                        `json:"totalTests,omitempty"`
}

// NewSyntheticLocation instantiates a new SyntheticLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticLocation(geoPoint SyntheticGeoPoint, label string, locationType string, playbackCapabilities SyntheticPlaybackCapabilities) *SyntheticLocation {
	this := SyntheticLocation{}
	this.GeoPoint = geoPoint
	this.Label = label
	this.LocationType = locationType
	this.PlaybackCapabilities = playbackCapabilities
	return &this
}

// NewSyntheticLocationWithDefaults instantiates a new SyntheticLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticLocationWithDefaults() *SyntheticLocation {
	this := SyntheticLocation{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SyntheticLocation) GetCreatedAt() int64 {
	if o == nil || isNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetCreatedAtOk() (*int64, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SyntheticLocation) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *SyntheticLocation) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *SyntheticLocation) GetCustomProperties() map[string]string {
	if o == nil || isNil(o.CustomProperties) {
		var ret map[string]string
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetCustomPropertiesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *SyntheticLocation) HasCustomProperties() bool {
	if o != nil && !isNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]string and assigns it to the CustomProperties field.
func (o *SyntheticLocation) SetCustomProperties(v map[string]string) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyntheticLocation) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyntheticLocation) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyntheticLocation) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *SyntheticLocation) GetDisplayLabel() string {
	if o == nil || isNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetDisplayLabelOk() (*string, bool) {
	if o == nil || isNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *SyntheticLocation) HasDisplayLabel() bool {
	if o != nil && !isNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *SyntheticLocation) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetGeoPoint returns the GeoPoint field value
func (o *SyntheticLocation) GetGeoPoint() SyntheticGeoPoint {
	if o == nil {
		var ret SyntheticGeoPoint
		return ret
	}

	return o.GeoPoint
}

// GetGeoPointOk returns a tuple with the GeoPoint field value
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetGeoPointOk() (*SyntheticGeoPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeoPoint, true
}

// SetGeoPoint sets field value
func (o *SyntheticLocation) SetGeoPoint(v SyntheticGeoPoint) {
	o.GeoPoint = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticLocation) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticLocation) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticLocation) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value
func (o *SyntheticLocation) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SyntheticLocation) SetLabel(v string) {
	o.Label = v
}

// GetLocationType returns the LocationType field value
func (o *SyntheticLocation) GetLocationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetLocationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationType, true
}

// SetLocationType sets field value
func (o *SyntheticLocation) SetLocationType(v string) {
	o.LocationType = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SyntheticLocation) GetModifiedAt() int64 {
	if o == nil || isNil(o.ModifiedAt) {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetModifiedAtOk() (*int64, bool) {
	if o == nil || isNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SyntheticLocation) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *SyntheticLocation) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *SyntheticLocation) GetObservedAt() int64 {
	if o == nil || isNil(o.ObservedAt) {
		var ret int64
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetObservedAtOk() (*int64, bool) {
	if o == nil || isNil(o.ObservedAt) {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *SyntheticLocation) HasObservedAt() bool {
	if o != nil && !isNil(o.ObservedAt) {
		return true
	}

	return false
}

// SetObservedAt gets a reference to the given int64 and assigns it to the ObservedAt field.
func (o *SyntheticLocation) SetObservedAt(v int64) {
	o.ObservedAt = &v
}

// GetPlaybackCapabilities returns the PlaybackCapabilities field value
func (o *SyntheticLocation) GetPlaybackCapabilities() SyntheticPlaybackCapabilities {
	if o == nil {
		var ret SyntheticPlaybackCapabilities
		return ret
	}

	return o.PlaybackCapabilities
}

// GetPlaybackCapabilitiesOk returns a tuple with the PlaybackCapabilities field value
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetPlaybackCapabilitiesOk() (*SyntheticPlaybackCapabilities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlaybackCapabilities, true
}

// SetPlaybackCapabilities sets field value
func (o *SyntheticLocation) SetPlaybackCapabilities(v SyntheticPlaybackCapabilities) {
	o.PlaybackCapabilities = v
}

// GetPopVersion returns the PopVersion field value if set, zero value otherwise.
func (o *SyntheticLocation) GetPopVersion() string {
	if o == nil || isNil(o.PopVersion) {
		var ret string
		return ret
	}
	return *o.PopVersion
}

// GetPopVersionOk returns a tuple with the PopVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetPopVersionOk() (*string, bool) {
	if o == nil || isNil(o.PopVersion) {
		return nil, false
	}
	return o.PopVersion, true
}

// HasPopVersion returns a boolean if a field has been set.
func (o *SyntheticLocation) HasPopVersion() bool {
	if o != nil && !isNil(o.PopVersion) {
		return true
	}

	return false
}

// SetPopVersion gets a reference to the given string and assigns it to the PopVersion field.
func (o *SyntheticLocation) SetPopVersion(v string) {
	o.PopVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticLocation) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticLocation) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticLocation) SetStatus(v string) {
	o.Status = &v
}

// GetTotalTests returns the TotalTests field value if set, zero value otherwise.
func (o *SyntheticLocation) GetTotalTests() int32 {
	if o == nil || isNil(o.TotalTests) {
		var ret int32
		return ret
	}
	return *o.TotalTests
}

// GetTotalTestsOk returns a tuple with the TotalTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticLocation) GetTotalTestsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalTests) {
		return nil, false
	}
	return o.TotalTests, true
}

// HasTotalTests returns a boolean if a field has been set.
func (o *SyntheticLocation) HasTotalTests() bool {
	if o != nil && !isNil(o.TotalTests) {
		return true
	}

	return false
}

// SetTotalTests gets a reference to the given int32 and assigns it to the TotalTests field.
func (o *SyntheticLocation) SetTotalTests(v int32) {
	o.TotalTests = &v
}

func (o SyntheticLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DisplayLabel) {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	if true {
		toSerialize["geoPoint"] = o.GeoPoint
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["locationType"] = o.LocationType
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !isNil(o.ObservedAt) {
		toSerialize["observedAt"] = o.ObservedAt
	}
	if true {
		toSerialize["playbackCapabilities"] = o.PlaybackCapabilities
	}
	if !isNil(o.PopVersion) {
		toSerialize["popVersion"] = o.PopVersion
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TotalTests) {
		toSerialize["totalTests"] = o.TotalTests
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticLocation struct {
	value *SyntheticLocation
	isSet bool
}

func (v NullableSyntheticLocation) Get() *SyntheticLocation {
	return v.value
}

func (v *NullableSyntheticLocation) Set(val *SyntheticLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticLocation(val *SyntheticLocation) *NullableSyntheticLocation {
	return &NullableSyntheticLocation{value: val, isSet: true}
}

func (v NullableSyntheticLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
