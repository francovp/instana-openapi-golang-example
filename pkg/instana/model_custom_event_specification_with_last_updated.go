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

// CustomEventSpecificationWithLastUpdated struct for CustomEventSpecificationWithLastUpdated
type CustomEventSpecificationWithLastUpdated struct {
	Actions                  []Action       `json:"actions,omitempty"`
	ApplicationAlertConfigId *string        `json:"applicationAlertConfigId,omitempty"`
	Deleted                  *bool          `json:"deleted,omitempty"`
	Description              *string        `json:"description,omitempty"`
	Enabled                  *bool          `json:"enabled,omitempty"`
	EntityType               string         `json:"entityType"`
	ExpirationTime           *int64         `json:"expirationTime,omitempty"`
	Id                       string         `json:"id"`
	LastUpdated              *int64         `json:"lastUpdated,omitempty"`
	Migrated                 *bool          `json:"migrated,omitempty"`
	Name                     string         `json:"name"`
	Query                    *string        `json:"query,omitempty"`
	Rules                    []AbstractRule `json:"rules"`
	Triggering               *bool          `json:"triggering,omitempty"`
	ValidVersion             *int32         `json:"validVersion,omitempty"`
}

// NewCustomEventSpecificationWithLastUpdated instantiates a new CustomEventSpecificationWithLastUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEventSpecificationWithLastUpdated(entityType string, id string, name string, rules []AbstractRule) *CustomEventSpecificationWithLastUpdated {
	this := CustomEventSpecificationWithLastUpdated{}
	this.EntityType = entityType
	this.Id = id
	this.Name = name
	this.Rules = rules
	return &this
}

// NewCustomEventSpecificationWithLastUpdatedWithDefaults instantiates a new CustomEventSpecificationWithLastUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEventSpecificationWithLastUpdatedWithDefaults() *CustomEventSpecificationWithLastUpdated {
	this := CustomEventSpecificationWithLastUpdated{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetActions() []Action {
	if o == nil || isNil(o.Actions) {
		var ret []Action
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetActionsOk() ([]Action, bool) {
	if o == nil || isNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasActions() bool {
	if o != nil && !isNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []Action and assigns it to the Actions field.
func (o *CustomEventSpecificationWithLastUpdated) SetActions(v []Action) {
	o.Actions = v
}

// GetApplicationAlertConfigId returns the ApplicationAlertConfigId field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetApplicationAlertConfigId() string {
	if o == nil || isNil(o.ApplicationAlertConfigId) {
		var ret string
		return ret
	}
	return *o.ApplicationAlertConfigId
}

// GetApplicationAlertConfigIdOk returns a tuple with the ApplicationAlertConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetApplicationAlertConfigIdOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationAlertConfigId) {
		return nil, false
	}
	return o.ApplicationAlertConfigId, true
}

// HasApplicationAlertConfigId returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasApplicationAlertConfigId() bool {
	if o != nil && !isNil(o.ApplicationAlertConfigId) {
		return true
	}

	return false
}

// SetApplicationAlertConfigId gets a reference to the given string and assigns it to the ApplicationAlertConfigId field.
func (o *CustomEventSpecificationWithLastUpdated) SetApplicationAlertConfigId(v string) {
	o.ApplicationAlertConfigId = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetDeleted() bool {
	if o == nil || isNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasDeleted() bool {
	if o != nil && !isNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CustomEventSpecificationWithLastUpdated) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomEventSpecificationWithLastUpdated) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomEventSpecificationWithLastUpdated) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEntityType returns the EntityType field value
func (o *CustomEventSpecificationWithLastUpdated) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *CustomEventSpecificationWithLastUpdated) SetEntityType(v string) {
	o.EntityType = v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetExpirationTime() int64 {
	if o == nil || isNil(o.ExpirationTime) {
		var ret int64
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetExpirationTimeOk() (*int64, bool) {
	if o == nil || isNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasExpirationTime() bool {
	if o != nil && !isNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given int64 and assigns it to the ExpirationTime field.
func (o *CustomEventSpecificationWithLastUpdated) SetExpirationTime(v int64) {
	o.ExpirationTime = &v
}

// GetId returns the Id field value
func (o *CustomEventSpecificationWithLastUpdated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomEventSpecificationWithLastUpdated) SetId(v string) {
	o.Id = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetLastUpdated() int64 {
	if o == nil || isNil(o.LastUpdated) {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || isNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasLastUpdated() bool {
	if o != nil && !isNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *CustomEventSpecificationWithLastUpdated) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

// GetMigrated returns the Migrated field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetMigrated() bool {
	if o == nil || isNil(o.Migrated) {
		var ret bool
		return ret
	}
	return *o.Migrated
}

// GetMigratedOk returns a tuple with the Migrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetMigratedOk() (*bool, bool) {
	if o == nil || isNil(o.Migrated) {
		return nil, false
	}
	return o.Migrated, true
}

// HasMigrated returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasMigrated() bool {
	if o != nil && !isNil(o.Migrated) {
		return true
	}

	return false
}

// SetMigrated gets a reference to the given bool and assigns it to the Migrated field.
func (o *CustomEventSpecificationWithLastUpdated) SetMigrated(v bool) {
	o.Migrated = &v
}

// GetName returns the Name field value
func (o *CustomEventSpecificationWithLastUpdated) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomEventSpecificationWithLastUpdated) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CustomEventSpecificationWithLastUpdated) SetQuery(v string) {
	o.Query = &v
}

// GetRules returns the Rules field value
func (o *CustomEventSpecificationWithLastUpdated) GetRules() []AbstractRule {
	if o == nil {
		var ret []AbstractRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetRulesOk() ([]AbstractRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *CustomEventSpecificationWithLastUpdated) SetRules(v []AbstractRule) {
	o.Rules = v
}

// GetTriggering returns the Triggering field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetTriggering() bool {
	if o == nil || isNil(o.Triggering) {
		var ret bool
		return ret
	}
	return *o.Triggering
}

// GetTriggeringOk returns a tuple with the Triggering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetTriggeringOk() (*bool, bool) {
	if o == nil || isNil(o.Triggering) {
		return nil, false
	}
	return o.Triggering, true
}

// HasTriggering returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasTriggering() bool {
	if o != nil && !isNil(o.Triggering) {
		return true
	}

	return false
}

// SetTriggering gets a reference to the given bool and assigns it to the Triggering field.
func (o *CustomEventSpecificationWithLastUpdated) SetTriggering(v bool) {
	o.Triggering = &v
}

// GetValidVersion returns the ValidVersion field value if set, zero value otherwise.
func (o *CustomEventSpecificationWithLastUpdated) GetValidVersion() int32 {
	if o == nil || isNil(o.ValidVersion) {
		var ret int32
		return ret
	}
	return *o.ValidVersion
}

// GetValidVersionOk returns a tuple with the ValidVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSpecificationWithLastUpdated) GetValidVersionOk() (*int32, bool) {
	if o == nil || isNil(o.ValidVersion) {
		return nil, false
	}
	return o.ValidVersion, true
}

// HasValidVersion returns a boolean if a field has been set.
func (o *CustomEventSpecificationWithLastUpdated) HasValidVersion() bool {
	if o != nil && !isNil(o.ValidVersion) {
		return true
	}

	return false
}

// SetValidVersion gets a reference to the given int32 and assigns it to the ValidVersion field.
func (o *CustomEventSpecificationWithLastUpdated) SetValidVersion(v int32) {
	o.ValidVersion = &v
}

func (o CustomEventSpecificationWithLastUpdated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !isNil(o.ApplicationAlertConfigId) {
		toSerialize["applicationAlertConfigId"] = o.ApplicationAlertConfigId
	}
	if !isNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["entityType"] = o.EntityType
	}
	if !isNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !isNil(o.Migrated) {
		toSerialize["migrated"] = o.Migrated
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.Triggering) {
		toSerialize["triggering"] = o.Triggering
	}
	if !isNil(o.ValidVersion) {
		toSerialize["validVersion"] = o.ValidVersion
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEventSpecificationWithLastUpdated struct {
	value *CustomEventSpecificationWithLastUpdated
	isSet bool
}

func (v NullableCustomEventSpecificationWithLastUpdated) Get() *CustomEventSpecificationWithLastUpdated {
	return v.value
}

func (v *NullableCustomEventSpecificationWithLastUpdated) Set(val *CustomEventSpecificationWithLastUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEventSpecificationWithLastUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEventSpecificationWithLastUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEventSpecificationWithLastUpdated(val *CustomEventSpecificationWithLastUpdated) *NullableCustomEventSpecificationWithLastUpdated {
	return &NullableCustomEventSpecificationWithLastUpdated{value: val, isSet: true}
}

func (v NullableCustomEventSpecificationWithLastUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEventSpecificationWithLastUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
