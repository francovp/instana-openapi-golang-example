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

// SyntheticTest struct for SyntheticTest
type SyntheticTest struct {
	Active                bool                       `json:"active"`
	ApplicationId         *string                    `json:"applicationId,omitempty"`
	ApplicationLabel      *string                    `json:"applicationLabel,omitempty"`
	Configuration         SyntheticTypeConfiguration `json:"configuration"`
	CreatedAt             *int64                     `json:"createdAt,omitempty"`
	CreatedBy             *string                    `json:"createdBy,omitempty"`
	CustomProperties      *map[string]string         `json:"customProperties,omitempty"`
	Description           *string                    `json:"description,omitempty"`
	Id                    *string                    `json:"id,omitempty"`
	Label                 string                     `json:"label"`
	LocationDisplayLabels []string                   `json:"locationDisplayLabels,omitempty"`
	LocationLabels        []string                   `json:"locationLabels,omitempty"`
	Locations             []string                   `json:"locations"`
	ModifiedAt            *int64                     `json:"modifiedAt,omitempty"`
	ModifiedBy            *string                    `json:"modifiedBy,omitempty"`
	PlaybackMode          string                     `json:"playbackMode"`
	TenantId              *string                    `json:"tenantId,omitempty"`
	TestFrequency         int32                      `json:"testFrequency"`
}

// NewSyntheticTest instantiates a new SyntheticTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticTest(active bool, configuration SyntheticTypeConfiguration, label string, locations []string, playbackMode string, testFrequency int32) *SyntheticTest {
	this := SyntheticTest{}
	this.Active = active
	this.Configuration = configuration
	this.Label = label
	this.Locations = locations
	this.PlaybackMode = playbackMode
	this.TestFrequency = testFrequency
	return &this
}

// NewSyntheticTestWithDefaults instantiates a new SyntheticTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticTestWithDefaults() *SyntheticTest {
	this := SyntheticTest{}
	return &this
}

// GetActive returns the Active field value
func (o *SyntheticTest) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *SyntheticTest) SetActive(v bool) {
	o.Active = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *SyntheticTest) GetApplicationId() string {
	if o == nil || isNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetApplicationIdOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *SyntheticTest) HasApplicationId() bool {
	if o != nil && !isNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *SyntheticTest) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetApplicationLabel returns the ApplicationLabel field value if set, zero value otherwise.
func (o *SyntheticTest) GetApplicationLabel() string {
	if o == nil || isNil(o.ApplicationLabel) {
		var ret string
		return ret
	}
	return *o.ApplicationLabel
}

// GetApplicationLabelOk returns a tuple with the ApplicationLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetApplicationLabelOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationLabel) {
		return nil, false
	}
	return o.ApplicationLabel, true
}

// HasApplicationLabel returns a boolean if a field has been set.
func (o *SyntheticTest) HasApplicationLabel() bool {
	if o != nil && !isNil(o.ApplicationLabel) {
		return true
	}

	return false
}

// SetApplicationLabel gets a reference to the given string and assigns it to the ApplicationLabel field.
func (o *SyntheticTest) SetApplicationLabel(v string) {
	o.ApplicationLabel = &v
}

// GetConfiguration returns the Configuration field value
func (o *SyntheticTest) GetConfiguration() SyntheticTypeConfiguration {
	if o == nil {
		var ret SyntheticTypeConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetConfigurationOk() (*SyntheticTypeConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *SyntheticTest) SetConfiguration(v SyntheticTypeConfiguration) {
	o.Configuration = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SyntheticTest) GetCreatedAt() int64 {
	if o == nil || isNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetCreatedAtOk() (*int64, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SyntheticTest) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *SyntheticTest) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SyntheticTest) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetCreatedByOk() (*string, bool) {
	if o == nil || isNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SyntheticTest) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SyntheticTest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *SyntheticTest) GetCustomProperties() map[string]string {
	if o == nil || isNil(o.CustomProperties) {
		var ret map[string]string
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetCustomPropertiesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *SyntheticTest) HasCustomProperties() bool {
	if o != nil && !isNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]string and assigns it to the CustomProperties field.
func (o *SyntheticTest) SetCustomProperties(v map[string]string) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyntheticTest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyntheticTest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyntheticTest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticTest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticTest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticTest) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value
func (o *SyntheticTest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SyntheticTest) SetLabel(v string) {
	o.Label = v
}

// GetLocationDisplayLabels returns the LocationDisplayLabels field value if set, zero value otherwise.
func (o *SyntheticTest) GetLocationDisplayLabels() []string {
	if o == nil || isNil(o.LocationDisplayLabels) {
		var ret []string
		return ret
	}
	return o.LocationDisplayLabels
}

// GetLocationDisplayLabelsOk returns a tuple with the LocationDisplayLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetLocationDisplayLabelsOk() ([]string, bool) {
	if o == nil || isNil(o.LocationDisplayLabels) {
		return nil, false
	}
	return o.LocationDisplayLabels, true
}

// HasLocationDisplayLabels returns a boolean if a field has been set.
func (o *SyntheticTest) HasLocationDisplayLabels() bool {
	if o != nil && !isNil(o.LocationDisplayLabels) {
		return true
	}

	return false
}

// SetLocationDisplayLabels gets a reference to the given []string and assigns it to the LocationDisplayLabels field.
func (o *SyntheticTest) SetLocationDisplayLabels(v []string) {
	o.LocationDisplayLabels = v
}

// GetLocationLabels returns the LocationLabels field value if set, zero value otherwise.
func (o *SyntheticTest) GetLocationLabels() []string {
	if o == nil || isNil(o.LocationLabels) {
		var ret []string
		return ret
	}
	return o.LocationLabels
}

// GetLocationLabelsOk returns a tuple with the LocationLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetLocationLabelsOk() ([]string, bool) {
	if o == nil || isNil(o.LocationLabels) {
		return nil, false
	}
	return o.LocationLabels, true
}

// HasLocationLabels returns a boolean if a field has been set.
func (o *SyntheticTest) HasLocationLabels() bool {
	if o != nil && !isNil(o.LocationLabels) {
		return true
	}

	return false
}

// SetLocationLabels gets a reference to the given []string and assigns it to the LocationLabels field.
func (o *SyntheticTest) SetLocationLabels(v []string) {
	o.LocationLabels = v
}

// GetLocations returns the Locations field value
func (o *SyntheticTest) GetLocations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetLocationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locations, true
}

// SetLocations sets field value
func (o *SyntheticTest) SetLocations(v []string) {
	o.Locations = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SyntheticTest) GetModifiedAt() int64 {
	if o == nil || isNil(o.ModifiedAt) {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetModifiedAtOk() (*int64, bool) {
	if o == nil || isNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SyntheticTest) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *SyntheticTest) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *SyntheticTest) GetModifiedBy() string {
	if o == nil || isNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetModifiedByOk() (*string, bool) {
	if o == nil || isNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *SyntheticTest) HasModifiedBy() bool {
	if o != nil && !isNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *SyntheticTest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetPlaybackMode returns the PlaybackMode field value
func (o *SyntheticTest) GetPlaybackMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlaybackMode
}

// GetPlaybackModeOk returns a tuple with the PlaybackMode field value
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetPlaybackModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlaybackMode, true
}

// SetPlaybackMode sets field value
func (o *SyntheticTest) SetPlaybackMode(v string) {
	o.PlaybackMode = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *SyntheticTest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *SyntheticTest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *SyntheticTest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTestFrequency returns the TestFrequency field value
func (o *SyntheticTest) GetTestFrequency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TestFrequency
}

// GetTestFrequencyOk returns a tuple with the TestFrequency field value
// and a boolean to check if the value has been set.
func (o *SyntheticTest) GetTestFrequencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestFrequency, true
}

// SetTestFrequency sets field value
func (o *SyntheticTest) SetTestFrequency(v int32) {
	o.TestFrequency = v
}

func (o SyntheticTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !isNil(o.ApplicationLabel) {
		toSerialize["applicationLabel"] = o.ApplicationLabel
	}
	if true {
		toSerialize["configuration"] = o.Configuration
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !isNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.LocationDisplayLabels) {
		toSerialize["locationDisplayLabels"] = o.LocationDisplayLabels
	}
	if !isNil(o.LocationLabels) {
		toSerialize["locationLabels"] = o.LocationLabels
	}
	if true {
		toSerialize["locations"] = o.Locations
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !isNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if true {
		toSerialize["playbackMode"] = o.PlaybackMode
	}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["testFrequency"] = o.TestFrequency
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticTest struct {
	value *SyntheticTest
	isSet bool
}

func (v NullableSyntheticTest) Get() *SyntheticTest {
	return v.value
}

func (v *NullableSyntheticTest) Set(val *SyntheticTest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticTest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticTest(val *SyntheticTest) *NullableSyntheticTest {
	return &NullableSyntheticTest{value: val, isSet: true}
}

func (v NullableSyntheticTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
