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

// ApiToken struct for ApiToken
type ApiToken struct {
	AccessGrantingToken                  string  `json:"accessGrantingToken"`
	CanConfigureAgentRunMode             *bool   `json:"canConfigureAgentRunMode,omitempty"`
	CanConfigureAgents                   *bool   `json:"canConfigureAgents,omitempty"`
	CanConfigureApiTokens                *bool   `json:"canConfigureApiTokens,omitempty"`
	CanConfigureApplications             *bool   `json:"canConfigureApplications,omitempty"`
	CanConfigureAuthenticationMethods    *bool   `json:"canConfigureAuthenticationMethods,omitempty"`
	CanConfigureAutomationActions        *bool   `json:"canConfigureAutomationActions,omitempty"`
	CanConfigureCustomAlerts             *bool   `json:"canConfigureCustomAlerts,omitempty"`
	CanConfigureEumApplications          *bool   `json:"canConfigureEumApplications,omitempty"`
	CanConfigureGlobalAlertConfigs       *bool   `json:"canConfigureGlobalAlertConfigs,omitempty"`
	CanConfigureGlobalAlertPayload       *bool   `json:"canConfigureGlobalAlertPayload,omitempty"`
	CanConfigureIntegrations             *bool   `json:"canConfigureIntegrations,omitempty"`
	CanConfigureLogManagement            *bool   `json:"canConfigureLogManagement,omitempty"`
	CanConfigureMobileAppMonitoring      *bool   `json:"canConfigureMobileAppMonitoring,omitempty"`
	CanConfigurePersonalApiTokens        *bool   `json:"canConfigurePersonalApiTokens,omitempty"`
	CanConfigureReleases                 *bool   `json:"canConfigureReleases,omitempty"`
	CanConfigureServiceLevelIndicators   *bool   `json:"canConfigureServiceLevelIndicators,omitempty"`
	CanConfigureServiceMapping           *bool   `json:"canConfigureServiceMapping,omitempty"`
	CanConfigureSessionSettings          *bool   `json:"canConfigureSessionSettings,omitempty"`
	CanConfigureTeams                    *bool   `json:"canConfigureTeams,omitempty"`
	CanConfigureUsers                    *bool   `json:"canConfigureUsers,omitempty"`
	CanCreatePublicCustomDashboards      *bool   `json:"canCreatePublicCustomDashboards,omitempty"`
	CanEditAllAccessibleCustomDashboards *bool   `json:"canEditAllAccessibleCustomDashboards,omitempty"`
	CanInstallNewAgents                  *bool   `json:"canInstallNewAgents,omitempty"`
	CanRunAutomationActions              *bool   `json:"canRunAutomationActions,omitempty"`
	CanSeeOnPremLicenseInformation       *bool   `json:"canSeeOnPremLicenseInformation,omitempty"`
	CanSeeUsageInformation               *bool   `json:"canSeeUsageInformation,omitempty"`
	CanViewAccountAndBillingInformation  *bool   `json:"canViewAccountAndBillingInformation,omitempty"`
	CanViewAuditLog                      *bool   `json:"canViewAuditLog,omitempty"`
	CanViewLogs                          *bool   `json:"canViewLogs,omitempty"`
	CanViewTraceDetails                  *bool   `json:"canViewTraceDetails,omitempty"`
	Id                                   *string `json:"id,omitempty"`
	InternalId                           string  `json:"internalId"`
	Name                                 string  `json:"name"`
}

// NewApiToken instantiates a new ApiToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiToken(accessGrantingToken string, internalId string, name string) *ApiToken {
	this := ApiToken{}
	this.AccessGrantingToken = accessGrantingToken
	this.InternalId = internalId
	this.Name = name
	return &this
}

// NewApiTokenWithDefaults instantiates a new ApiToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenWithDefaults() *ApiToken {
	this := ApiToken{}
	return &this
}

// GetAccessGrantingToken returns the AccessGrantingToken field value
func (o *ApiToken) GetAccessGrantingToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessGrantingToken
}

// GetAccessGrantingTokenOk returns a tuple with the AccessGrantingToken field value
// and a boolean to check if the value has been set.
func (o *ApiToken) GetAccessGrantingTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessGrantingToken, true
}

// SetAccessGrantingToken sets field value
func (o *ApiToken) SetAccessGrantingToken(v string) {
	o.AccessGrantingToken = v
}

// GetCanConfigureAgentRunMode returns the CanConfigureAgentRunMode field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureAgentRunMode() bool {
	if o == nil || isNil(o.CanConfigureAgentRunMode) {
		var ret bool
		return ret
	}
	return *o.CanConfigureAgentRunMode
}

// GetCanConfigureAgentRunModeOk returns a tuple with the CanConfigureAgentRunMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureAgentRunModeOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureAgentRunMode) {
		return nil, false
	}
	return o.CanConfigureAgentRunMode, true
}

// HasCanConfigureAgentRunMode returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureAgentRunMode() bool {
	if o != nil && !isNil(o.CanConfigureAgentRunMode) {
		return true
	}

	return false
}

// SetCanConfigureAgentRunMode gets a reference to the given bool and assigns it to the CanConfigureAgentRunMode field.
func (o *ApiToken) SetCanConfigureAgentRunMode(v bool) {
	o.CanConfigureAgentRunMode = &v
}

// GetCanConfigureAgents returns the CanConfigureAgents field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureAgents() bool {
	if o == nil || isNil(o.CanConfigureAgents) {
		var ret bool
		return ret
	}
	return *o.CanConfigureAgents
}

// GetCanConfigureAgentsOk returns a tuple with the CanConfigureAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureAgentsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureAgents) {
		return nil, false
	}
	return o.CanConfigureAgents, true
}

// HasCanConfigureAgents returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureAgents() bool {
	if o != nil && !isNil(o.CanConfigureAgents) {
		return true
	}

	return false
}

// SetCanConfigureAgents gets a reference to the given bool and assigns it to the CanConfigureAgents field.
func (o *ApiToken) SetCanConfigureAgents(v bool) {
	o.CanConfigureAgents = &v
}

// GetCanConfigureApiTokens returns the CanConfigureApiTokens field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureApiTokens() bool {
	if o == nil || isNil(o.CanConfigureApiTokens) {
		var ret bool
		return ret
	}
	return *o.CanConfigureApiTokens
}

// GetCanConfigureApiTokensOk returns a tuple with the CanConfigureApiTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureApiTokensOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureApiTokens) {
		return nil, false
	}
	return o.CanConfigureApiTokens, true
}

// HasCanConfigureApiTokens returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureApiTokens() bool {
	if o != nil && !isNil(o.CanConfigureApiTokens) {
		return true
	}

	return false
}

// SetCanConfigureApiTokens gets a reference to the given bool and assigns it to the CanConfigureApiTokens field.
func (o *ApiToken) SetCanConfigureApiTokens(v bool) {
	o.CanConfigureApiTokens = &v
}

// GetCanConfigureApplications returns the CanConfigureApplications field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureApplications() bool {
	if o == nil || isNil(o.CanConfigureApplications) {
		var ret bool
		return ret
	}
	return *o.CanConfigureApplications
}

// GetCanConfigureApplicationsOk returns a tuple with the CanConfigureApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureApplicationsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureApplications) {
		return nil, false
	}
	return o.CanConfigureApplications, true
}

// HasCanConfigureApplications returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureApplications() bool {
	if o != nil && !isNil(o.CanConfigureApplications) {
		return true
	}

	return false
}

// SetCanConfigureApplications gets a reference to the given bool and assigns it to the CanConfigureApplications field.
func (o *ApiToken) SetCanConfigureApplications(v bool) {
	o.CanConfigureApplications = &v
}

// GetCanConfigureAuthenticationMethods returns the CanConfigureAuthenticationMethods field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureAuthenticationMethods() bool {
	if o == nil || isNil(o.CanConfigureAuthenticationMethods) {
		var ret bool
		return ret
	}
	return *o.CanConfigureAuthenticationMethods
}

// GetCanConfigureAuthenticationMethodsOk returns a tuple with the CanConfigureAuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureAuthenticationMethodsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureAuthenticationMethods) {
		return nil, false
	}
	return o.CanConfigureAuthenticationMethods, true
}

// HasCanConfigureAuthenticationMethods returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureAuthenticationMethods() bool {
	if o != nil && !isNil(o.CanConfigureAuthenticationMethods) {
		return true
	}

	return false
}

// SetCanConfigureAuthenticationMethods gets a reference to the given bool and assigns it to the CanConfigureAuthenticationMethods field.
func (o *ApiToken) SetCanConfigureAuthenticationMethods(v bool) {
	o.CanConfigureAuthenticationMethods = &v
}

// GetCanConfigureAutomationActions returns the CanConfigureAutomationActions field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureAutomationActions() bool {
	if o == nil || isNil(o.CanConfigureAutomationActions) {
		var ret bool
		return ret
	}
	return *o.CanConfigureAutomationActions
}

// GetCanConfigureAutomationActionsOk returns a tuple with the CanConfigureAutomationActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureAutomationActionsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureAutomationActions) {
		return nil, false
	}
	return o.CanConfigureAutomationActions, true
}

// HasCanConfigureAutomationActions returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureAutomationActions() bool {
	if o != nil && !isNil(o.CanConfigureAutomationActions) {
		return true
	}

	return false
}

// SetCanConfigureAutomationActions gets a reference to the given bool and assigns it to the CanConfigureAutomationActions field.
func (o *ApiToken) SetCanConfigureAutomationActions(v bool) {
	o.CanConfigureAutomationActions = &v
}

// GetCanConfigureCustomAlerts returns the CanConfigureCustomAlerts field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureCustomAlerts() bool {
	if o == nil || isNil(o.CanConfigureCustomAlerts) {
		var ret bool
		return ret
	}
	return *o.CanConfigureCustomAlerts
}

// GetCanConfigureCustomAlertsOk returns a tuple with the CanConfigureCustomAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureCustomAlertsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureCustomAlerts) {
		return nil, false
	}
	return o.CanConfigureCustomAlerts, true
}

// HasCanConfigureCustomAlerts returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureCustomAlerts() bool {
	if o != nil && !isNil(o.CanConfigureCustomAlerts) {
		return true
	}

	return false
}

// SetCanConfigureCustomAlerts gets a reference to the given bool and assigns it to the CanConfigureCustomAlerts field.
func (o *ApiToken) SetCanConfigureCustomAlerts(v bool) {
	o.CanConfigureCustomAlerts = &v
}

// GetCanConfigureEumApplications returns the CanConfigureEumApplications field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureEumApplications() bool {
	if o == nil || isNil(o.CanConfigureEumApplications) {
		var ret bool
		return ret
	}
	return *o.CanConfigureEumApplications
}

// GetCanConfigureEumApplicationsOk returns a tuple with the CanConfigureEumApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureEumApplicationsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureEumApplications) {
		return nil, false
	}
	return o.CanConfigureEumApplications, true
}

// HasCanConfigureEumApplications returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureEumApplications() bool {
	if o != nil && !isNil(o.CanConfigureEumApplications) {
		return true
	}

	return false
}

// SetCanConfigureEumApplications gets a reference to the given bool and assigns it to the CanConfigureEumApplications field.
func (o *ApiToken) SetCanConfigureEumApplications(v bool) {
	o.CanConfigureEumApplications = &v
}

// GetCanConfigureGlobalAlertConfigs returns the CanConfigureGlobalAlertConfigs field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureGlobalAlertConfigs() bool {
	if o == nil || isNil(o.CanConfigureGlobalAlertConfigs) {
		var ret bool
		return ret
	}
	return *o.CanConfigureGlobalAlertConfigs
}

// GetCanConfigureGlobalAlertConfigsOk returns a tuple with the CanConfigureGlobalAlertConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureGlobalAlertConfigsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureGlobalAlertConfigs) {
		return nil, false
	}
	return o.CanConfigureGlobalAlertConfigs, true
}

// HasCanConfigureGlobalAlertConfigs returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureGlobalAlertConfigs() bool {
	if o != nil && !isNil(o.CanConfigureGlobalAlertConfigs) {
		return true
	}

	return false
}

// SetCanConfigureGlobalAlertConfigs gets a reference to the given bool and assigns it to the CanConfigureGlobalAlertConfigs field.
func (o *ApiToken) SetCanConfigureGlobalAlertConfigs(v bool) {
	o.CanConfigureGlobalAlertConfigs = &v
}

// GetCanConfigureGlobalAlertPayload returns the CanConfigureGlobalAlertPayload field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureGlobalAlertPayload() bool {
	if o == nil || isNil(o.CanConfigureGlobalAlertPayload) {
		var ret bool
		return ret
	}
	return *o.CanConfigureGlobalAlertPayload
}

// GetCanConfigureGlobalAlertPayloadOk returns a tuple with the CanConfigureGlobalAlertPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureGlobalAlertPayloadOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureGlobalAlertPayload) {
		return nil, false
	}
	return o.CanConfigureGlobalAlertPayload, true
}

// HasCanConfigureGlobalAlertPayload returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureGlobalAlertPayload() bool {
	if o != nil && !isNil(o.CanConfigureGlobalAlertPayload) {
		return true
	}

	return false
}

// SetCanConfigureGlobalAlertPayload gets a reference to the given bool and assigns it to the CanConfigureGlobalAlertPayload field.
func (o *ApiToken) SetCanConfigureGlobalAlertPayload(v bool) {
	o.CanConfigureGlobalAlertPayload = &v
}

// GetCanConfigureIntegrations returns the CanConfigureIntegrations field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureIntegrations() bool {
	if o == nil || isNil(o.CanConfigureIntegrations) {
		var ret bool
		return ret
	}
	return *o.CanConfigureIntegrations
}

// GetCanConfigureIntegrationsOk returns a tuple with the CanConfigureIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureIntegrationsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureIntegrations) {
		return nil, false
	}
	return o.CanConfigureIntegrations, true
}

// HasCanConfigureIntegrations returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureIntegrations() bool {
	if o != nil && !isNil(o.CanConfigureIntegrations) {
		return true
	}

	return false
}

// SetCanConfigureIntegrations gets a reference to the given bool and assigns it to the CanConfigureIntegrations field.
func (o *ApiToken) SetCanConfigureIntegrations(v bool) {
	o.CanConfigureIntegrations = &v
}

// GetCanConfigureLogManagement returns the CanConfigureLogManagement field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureLogManagement() bool {
	if o == nil || isNil(o.CanConfigureLogManagement) {
		var ret bool
		return ret
	}
	return *o.CanConfigureLogManagement
}

// GetCanConfigureLogManagementOk returns a tuple with the CanConfigureLogManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureLogManagementOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureLogManagement) {
		return nil, false
	}
	return o.CanConfigureLogManagement, true
}

// HasCanConfigureLogManagement returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureLogManagement() bool {
	if o != nil && !isNil(o.CanConfigureLogManagement) {
		return true
	}

	return false
}

// SetCanConfigureLogManagement gets a reference to the given bool and assigns it to the CanConfigureLogManagement field.
func (o *ApiToken) SetCanConfigureLogManagement(v bool) {
	o.CanConfigureLogManagement = &v
}

// GetCanConfigureMobileAppMonitoring returns the CanConfigureMobileAppMonitoring field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureMobileAppMonitoring() bool {
	if o == nil || isNil(o.CanConfigureMobileAppMonitoring) {
		var ret bool
		return ret
	}
	return *o.CanConfigureMobileAppMonitoring
}

// GetCanConfigureMobileAppMonitoringOk returns a tuple with the CanConfigureMobileAppMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureMobileAppMonitoringOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureMobileAppMonitoring) {
		return nil, false
	}
	return o.CanConfigureMobileAppMonitoring, true
}

// HasCanConfigureMobileAppMonitoring returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureMobileAppMonitoring() bool {
	if o != nil && !isNil(o.CanConfigureMobileAppMonitoring) {
		return true
	}

	return false
}

// SetCanConfigureMobileAppMonitoring gets a reference to the given bool and assigns it to the CanConfigureMobileAppMonitoring field.
func (o *ApiToken) SetCanConfigureMobileAppMonitoring(v bool) {
	o.CanConfigureMobileAppMonitoring = &v
}

// GetCanConfigurePersonalApiTokens returns the CanConfigurePersonalApiTokens field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigurePersonalApiTokens() bool {
	if o == nil || isNil(o.CanConfigurePersonalApiTokens) {
		var ret bool
		return ret
	}
	return *o.CanConfigurePersonalApiTokens
}

// GetCanConfigurePersonalApiTokensOk returns a tuple with the CanConfigurePersonalApiTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigurePersonalApiTokensOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigurePersonalApiTokens) {
		return nil, false
	}
	return o.CanConfigurePersonalApiTokens, true
}

// HasCanConfigurePersonalApiTokens returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigurePersonalApiTokens() bool {
	if o != nil && !isNil(o.CanConfigurePersonalApiTokens) {
		return true
	}

	return false
}

// SetCanConfigurePersonalApiTokens gets a reference to the given bool and assigns it to the CanConfigurePersonalApiTokens field.
func (o *ApiToken) SetCanConfigurePersonalApiTokens(v bool) {
	o.CanConfigurePersonalApiTokens = &v
}

// GetCanConfigureReleases returns the CanConfigureReleases field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureReleases() bool {
	if o == nil || isNil(o.CanConfigureReleases) {
		var ret bool
		return ret
	}
	return *o.CanConfigureReleases
}

// GetCanConfigureReleasesOk returns a tuple with the CanConfigureReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureReleasesOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureReleases) {
		return nil, false
	}
	return o.CanConfigureReleases, true
}

// HasCanConfigureReleases returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureReleases() bool {
	if o != nil && !isNil(o.CanConfigureReleases) {
		return true
	}

	return false
}

// SetCanConfigureReleases gets a reference to the given bool and assigns it to the CanConfigureReleases field.
func (o *ApiToken) SetCanConfigureReleases(v bool) {
	o.CanConfigureReleases = &v
}

// GetCanConfigureServiceLevelIndicators returns the CanConfigureServiceLevelIndicators field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureServiceLevelIndicators() bool {
	if o == nil || isNil(o.CanConfigureServiceLevelIndicators) {
		var ret bool
		return ret
	}
	return *o.CanConfigureServiceLevelIndicators
}

// GetCanConfigureServiceLevelIndicatorsOk returns a tuple with the CanConfigureServiceLevelIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureServiceLevelIndicatorsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureServiceLevelIndicators) {
		return nil, false
	}
	return o.CanConfigureServiceLevelIndicators, true
}

// HasCanConfigureServiceLevelIndicators returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureServiceLevelIndicators() bool {
	if o != nil && !isNil(o.CanConfigureServiceLevelIndicators) {
		return true
	}

	return false
}

// SetCanConfigureServiceLevelIndicators gets a reference to the given bool and assigns it to the CanConfigureServiceLevelIndicators field.
func (o *ApiToken) SetCanConfigureServiceLevelIndicators(v bool) {
	o.CanConfigureServiceLevelIndicators = &v
}

// GetCanConfigureServiceMapping returns the CanConfigureServiceMapping field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureServiceMapping() bool {
	if o == nil || isNil(o.CanConfigureServiceMapping) {
		var ret bool
		return ret
	}
	return *o.CanConfigureServiceMapping
}

// GetCanConfigureServiceMappingOk returns a tuple with the CanConfigureServiceMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureServiceMappingOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureServiceMapping) {
		return nil, false
	}
	return o.CanConfigureServiceMapping, true
}

// HasCanConfigureServiceMapping returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureServiceMapping() bool {
	if o != nil && !isNil(o.CanConfigureServiceMapping) {
		return true
	}

	return false
}

// SetCanConfigureServiceMapping gets a reference to the given bool and assigns it to the CanConfigureServiceMapping field.
func (o *ApiToken) SetCanConfigureServiceMapping(v bool) {
	o.CanConfigureServiceMapping = &v
}

// GetCanConfigureSessionSettings returns the CanConfigureSessionSettings field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureSessionSettings() bool {
	if o == nil || isNil(o.CanConfigureSessionSettings) {
		var ret bool
		return ret
	}
	return *o.CanConfigureSessionSettings
}

// GetCanConfigureSessionSettingsOk returns a tuple with the CanConfigureSessionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureSessionSettingsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureSessionSettings) {
		return nil, false
	}
	return o.CanConfigureSessionSettings, true
}

// HasCanConfigureSessionSettings returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureSessionSettings() bool {
	if o != nil && !isNil(o.CanConfigureSessionSettings) {
		return true
	}

	return false
}

// SetCanConfigureSessionSettings gets a reference to the given bool and assigns it to the CanConfigureSessionSettings field.
func (o *ApiToken) SetCanConfigureSessionSettings(v bool) {
	o.CanConfigureSessionSettings = &v
}

// GetCanConfigureTeams returns the CanConfigureTeams field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureTeams() bool {
	if o == nil || isNil(o.CanConfigureTeams) {
		var ret bool
		return ret
	}
	return *o.CanConfigureTeams
}

// GetCanConfigureTeamsOk returns a tuple with the CanConfigureTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureTeamsOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureTeams) {
		return nil, false
	}
	return o.CanConfigureTeams, true
}

// HasCanConfigureTeams returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureTeams() bool {
	if o != nil && !isNil(o.CanConfigureTeams) {
		return true
	}

	return false
}

// SetCanConfigureTeams gets a reference to the given bool and assigns it to the CanConfigureTeams field.
func (o *ApiToken) SetCanConfigureTeams(v bool) {
	o.CanConfigureTeams = &v
}

// GetCanConfigureUsers returns the CanConfigureUsers field value if set, zero value otherwise.
func (o *ApiToken) GetCanConfigureUsers() bool {
	if o == nil || isNil(o.CanConfigureUsers) {
		var ret bool
		return ret
	}
	return *o.CanConfigureUsers
}

// GetCanConfigureUsersOk returns a tuple with the CanConfigureUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanConfigureUsersOk() (*bool, bool) {
	if o == nil || isNil(o.CanConfigureUsers) {
		return nil, false
	}
	return o.CanConfigureUsers, true
}

// HasCanConfigureUsers returns a boolean if a field has been set.
func (o *ApiToken) HasCanConfigureUsers() bool {
	if o != nil && !isNil(o.CanConfigureUsers) {
		return true
	}

	return false
}

// SetCanConfigureUsers gets a reference to the given bool and assigns it to the CanConfigureUsers field.
func (o *ApiToken) SetCanConfigureUsers(v bool) {
	o.CanConfigureUsers = &v
}

// GetCanCreatePublicCustomDashboards returns the CanCreatePublicCustomDashboards field value if set, zero value otherwise.
func (o *ApiToken) GetCanCreatePublicCustomDashboards() bool {
	if o == nil || isNil(o.CanCreatePublicCustomDashboards) {
		var ret bool
		return ret
	}
	return *o.CanCreatePublicCustomDashboards
}

// GetCanCreatePublicCustomDashboardsOk returns a tuple with the CanCreatePublicCustomDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanCreatePublicCustomDashboardsOk() (*bool, bool) {
	if o == nil || isNil(o.CanCreatePublicCustomDashboards) {
		return nil, false
	}
	return o.CanCreatePublicCustomDashboards, true
}

// HasCanCreatePublicCustomDashboards returns a boolean if a field has been set.
func (o *ApiToken) HasCanCreatePublicCustomDashboards() bool {
	if o != nil && !isNil(o.CanCreatePublicCustomDashboards) {
		return true
	}

	return false
}

// SetCanCreatePublicCustomDashboards gets a reference to the given bool and assigns it to the CanCreatePublicCustomDashboards field.
func (o *ApiToken) SetCanCreatePublicCustomDashboards(v bool) {
	o.CanCreatePublicCustomDashboards = &v
}

// GetCanEditAllAccessibleCustomDashboards returns the CanEditAllAccessibleCustomDashboards field value if set, zero value otherwise.
func (o *ApiToken) GetCanEditAllAccessibleCustomDashboards() bool {
	if o == nil || isNil(o.CanEditAllAccessibleCustomDashboards) {
		var ret bool
		return ret
	}
	return *o.CanEditAllAccessibleCustomDashboards
}

// GetCanEditAllAccessibleCustomDashboardsOk returns a tuple with the CanEditAllAccessibleCustomDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanEditAllAccessibleCustomDashboardsOk() (*bool, bool) {
	if o == nil || isNil(o.CanEditAllAccessibleCustomDashboards) {
		return nil, false
	}
	return o.CanEditAllAccessibleCustomDashboards, true
}

// HasCanEditAllAccessibleCustomDashboards returns a boolean if a field has been set.
func (o *ApiToken) HasCanEditAllAccessibleCustomDashboards() bool {
	if o != nil && !isNil(o.CanEditAllAccessibleCustomDashboards) {
		return true
	}

	return false
}

// SetCanEditAllAccessibleCustomDashboards gets a reference to the given bool and assigns it to the CanEditAllAccessibleCustomDashboards field.
func (o *ApiToken) SetCanEditAllAccessibleCustomDashboards(v bool) {
	o.CanEditAllAccessibleCustomDashboards = &v
}

// GetCanInstallNewAgents returns the CanInstallNewAgents field value if set, zero value otherwise.
func (o *ApiToken) GetCanInstallNewAgents() bool {
	if o == nil || isNil(o.CanInstallNewAgents) {
		var ret bool
		return ret
	}
	return *o.CanInstallNewAgents
}

// GetCanInstallNewAgentsOk returns a tuple with the CanInstallNewAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanInstallNewAgentsOk() (*bool, bool) {
	if o == nil || isNil(o.CanInstallNewAgents) {
		return nil, false
	}
	return o.CanInstallNewAgents, true
}

// HasCanInstallNewAgents returns a boolean if a field has been set.
func (o *ApiToken) HasCanInstallNewAgents() bool {
	if o != nil && !isNil(o.CanInstallNewAgents) {
		return true
	}

	return false
}

// SetCanInstallNewAgents gets a reference to the given bool and assigns it to the CanInstallNewAgents field.
func (o *ApiToken) SetCanInstallNewAgents(v bool) {
	o.CanInstallNewAgents = &v
}

// GetCanRunAutomationActions returns the CanRunAutomationActions field value if set, zero value otherwise.
func (o *ApiToken) GetCanRunAutomationActions() bool {
	if o == nil || isNil(o.CanRunAutomationActions) {
		var ret bool
		return ret
	}
	return *o.CanRunAutomationActions
}

// GetCanRunAutomationActionsOk returns a tuple with the CanRunAutomationActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanRunAutomationActionsOk() (*bool, bool) {
	if o == nil || isNil(o.CanRunAutomationActions) {
		return nil, false
	}
	return o.CanRunAutomationActions, true
}

// HasCanRunAutomationActions returns a boolean if a field has been set.
func (o *ApiToken) HasCanRunAutomationActions() bool {
	if o != nil && !isNil(o.CanRunAutomationActions) {
		return true
	}

	return false
}

// SetCanRunAutomationActions gets a reference to the given bool and assigns it to the CanRunAutomationActions field.
func (o *ApiToken) SetCanRunAutomationActions(v bool) {
	o.CanRunAutomationActions = &v
}

// GetCanSeeOnPremLicenseInformation returns the CanSeeOnPremLicenseInformation field value if set, zero value otherwise.
func (o *ApiToken) GetCanSeeOnPremLicenseInformation() bool {
	if o == nil || isNil(o.CanSeeOnPremLicenseInformation) {
		var ret bool
		return ret
	}
	return *o.CanSeeOnPremLicenseInformation
}

// GetCanSeeOnPremLicenseInformationOk returns a tuple with the CanSeeOnPremLicenseInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanSeeOnPremLicenseInformationOk() (*bool, bool) {
	if o == nil || isNil(o.CanSeeOnPremLicenseInformation) {
		return nil, false
	}
	return o.CanSeeOnPremLicenseInformation, true
}

// HasCanSeeOnPremLicenseInformation returns a boolean if a field has been set.
func (o *ApiToken) HasCanSeeOnPremLicenseInformation() bool {
	if o != nil && !isNil(o.CanSeeOnPremLicenseInformation) {
		return true
	}

	return false
}

// SetCanSeeOnPremLicenseInformation gets a reference to the given bool and assigns it to the CanSeeOnPremLicenseInformation field.
func (o *ApiToken) SetCanSeeOnPremLicenseInformation(v bool) {
	o.CanSeeOnPremLicenseInformation = &v
}

// GetCanSeeUsageInformation returns the CanSeeUsageInformation field value if set, zero value otherwise.
func (o *ApiToken) GetCanSeeUsageInformation() bool {
	if o == nil || isNil(o.CanSeeUsageInformation) {
		var ret bool
		return ret
	}
	return *o.CanSeeUsageInformation
}

// GetCanSeeUsageInformationOk returns a tuple with the CanSeeUsageInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanSeeUsageInformationOk() (*bool, bool) {
	if o == nil || isNil(o.CanSeeUsageInformation) {
		return nil, false
	}
	return o.CanSeeUsageInformation, true
}

// HasCanSeeUsageInformation returns a boolean if a field has been set.
func (o *ApiToken) HasCanSeeUsageInformation() bool {
	if o != nil && !isNil(o.CanSeeUsageInformation) {
		return true
	}

	return false
}

// SetCanSeeUsageInformation gets a reference to the given bool and assigns it to the CanSeeUsageInformation field.
func (o *ApiToken) SetCanSeeUsageInformation(v bool) {
	o.CanSeeUsageInformation = &v
}

// GetCanViewAccountAndBillingInformation returns the CanViewAccountAndBillingInformation field value if set, zero value otherwise.
func (o *ApiToken) GetCanViewAccountAndBillingInformation() bool {
	if o == nil || isNil(o.CanViewAccountAndBillingInformation) {
		var ret bool
		return ret
	}
	return *o.CanViewAccountAndBillingInformation
}

// GetCanViewAccountAndBillingInformationOk returns a tuple with the CanViewAccountAndBillingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanViewAccountAndBillingInformationOk() (*bool, bool) {
	if o == nil || isNil(o.CanViewAccountAndBillingInformation) {
		return nil, false
	}
	return o.CanViewAccountAndBillingInformation, true
}

// HasCanViewAccountAndBillingInformation returns a boolean if a field has been set.
func (o *ApiToken) HasCanViewAccountAndBillingInformation() bool {
	if o != nil && !isNil(o.CanViewAccountAndBillingInformation) {
		return true
	}

	return false
}

// SetCanViewAccountAndBillingInformation gets a reference to the given bool and assigns it to the CanViewAccountAndBillingInformation field.
func (o *ApiToken) SetCanViewAccountAndBillingInformation(v bool) {
	o.CanViewAccountAndBillingInformation = &v
}

// GetCanViewAuditLog returns the CanViewAuditLog field value if set, zero value otherwise.
func (o *ApiToken) GetCanViewAuditLog() bool {
	if o == nil || isNil(o.CanViewAuditLog) {
		var ret bool
		return ret
	}
	return *o.CanViewAuditLog
}

// GetCanViewAuditLogOk returns a tuple with the CanViewAuditLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanViewAuditLogOk() (*bool, bool) {
	if o == nil || isNil(o.CanViewAuditLog) {
		return nil, false
	}
	return o.CanViewAuditLog, true
}

// HasCanViewAuditLog returns a boolean if a field has been set.
func (o *ApiToken) HasCanViewAuditLog() bool {
	if o != nil && !isNil(o.CanViewAuditLog) {
		return true
	}

	return false
}

// SetCanViewAuditLog gets a reference to the given bool and assigns it to the CanViewAuditLog field.
func (o *ApiToken) SetCanViewAuditLog(v bool) {
	o.CanViewAuditLog = &v
}

// GetCanViewLogs returns the CanViewLogs field value if set, zero value otherwise.
func (o *ApiToken) GetCanViewLogs() bool {
	if o == nil || isNil(o.CanViewLogs) {
		var ret bool
		return ret
	}
	return *o.CanViewLogs
}

// GetCanViewLogsOk returns a tuple with the CanViewLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanViewLogsOk() (*bool, bool) {
	if o == nil || isNil(o.CanViewLogs) {
		return nil, false
	}
	return o.CanViewLogs, true
}

// HasCanViewLogs returns a boolean if a field has been set.
func (o *ApiToken) HasCanViewLogs() bool {
	if o != nil && !isNil(o.CanViewLogs) {
		return true
	}

	return false
}

// SetCanViewLogs gets a reference to the given bool and assigns it to the CanViewLogs field.
func (o *ApiToken) SetCanViewLogs(v bool) {
	o.CanViewLogs = &v
}

// GetCanViewTraceDetails returns the CanViewTraceDetails field value if set, zero value otherwise.
func (o *ApiToken) GetCanViewTraceDetails() bool {
	if o == nil || isNil(o.CanViewTraceDetails) {
		var ret bool
		return ret
	}
	return *o.CanViewTraceDetails
}

// GetCanViewTraceDetailsOk returns a tuple with the CanViewTraceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCanViewTraceDetailsOk() (*bool, bool) {
	if o == nil || isNil(o.CanViewTraceDetails) {
		return nil, false
	}
	return o.CanViewTraceDetails, true
}

// HasCanViewTraceDetails returns a boolean if a field has been set.
func (o *ApiToken) HasCanViewTraceDetails() bool {
	if o != nil && !isNil(o.CanViewTraceDetails) {
		return true
	}

	return false
}

// SetCanViewTraceDetails gets a reference to the given bool and assigns it to the CanViewTraceDetails field.
func (o *ApiToken) SetCanViewTraceDetails(v bool) {
	o.CanViewTraceDetails = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiToken) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiToken) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiToken) SetId(v string) {
	o.Id = &v
}

// GetInternalId returns the InternalId field value
func (o *ApiToken) GetInternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value
// and a boolean to check if the value has been set.
func (o *ApiToken) GetInternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalId, true
}

// SetInternalId sets field value
func (o *ApiToken) SetInternalId(v string) {
	o.InternalId = v
}

// GetName returns the Name field value
func (o *ApiToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiToken) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiToken) SetName(v string) {
	o.Name = v
}

func (o ApiToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessGrantingToken"] = o.AccessGrantingToken
	}
	if !isNil(o.CanConfigureAgentRunMode) {
		toSerialize["canConfigureAgentRunMode"] = o.CanConfigureAgentRunMode
	}
	if !isNil(o.CanConfigureAgents) {
		toSerialize["canConfigureAgents"] = o.CanConfigureAgents
	}
	if !isNil(o.CanConfigureApiTokens) {
		toSerialize["canConfigureApiTokens"] = o.CanConfigureApiTokens
	}
	if !isNil(o.CanConfigureApplications) {
		toSerialize["canConfigureApplications"] = o.CanConfigureApplications
	}
	if !isNil(o.CanConfigureAuthenticationMethods) {
		toSerialize["canConfigureAuthenticationMethods"] = o.CanConfigureAuthenticationMethods
	}
	if !isNil(o.CanConfigureAutomationActions) {
		toSerialize["canConfigureAutomationActions"] = o.CanConfigureAutomationActions
	}
	if !isNil(o.CanConfigureCustomAlerts) {
		toSerialize["canConfigureCustomAlerts"] = o.CanConfigureCustomAlerts
	}
	if !isNil(o.CanConfigureEumApplications) {
		toSerialize["canConfigureEumApplications"] = o.CanConfigureEumApplications
	}
	if !isNil(o.CanConfigureGlobalAlertConfigs) {
		toSerialize["canConfigureGlobalAlertConfigs"] = o.CanConfigureGlobalAlertConfigs
	}
	if !isNil(o.CanConfigureGlobalAlertPayload) {
		toSerialize["canConfigureGlobalAlertPayload"] = o.CanConfigureGlobalAlertPayload
	}
	if !isNil(o.CanConfigureIntegrations) {
		toSerialize["canConfigureIntegrations"] = o.CanConfigureIntegrations
	}
	if !isNil(o.CanConfigureLogManagement) {
		toSerialize["canConfigureLogManagement"] = o.CanConfigureLogManagement
	}
	if !isNil(o.CanConfigureMobileAppMonitoring) {
		toSerialize["canConfigureMobileAppMonitoring"] = o.CanConfigureMobileAppMonitoring
	}
	if !isNil(o.CanConfigurePersonalApiTokens) {
		toSerialize["canConfigurePersonalApiTokens"] = o.CanConfigurePersonalApiTokens
	}
	if !isNil(o.CanConfigureReleases) {
		toSerialize["canConfigureReleases"] = o.CanConfigureReleases
	}
	if !isNil(o.CanConfigureServiceLevelIndicators) {
		toSerialize["canConfigureServiceLevelIndicators"] = o.CanConfigureServiceLevelIndicators
	}
	if !isNil(o.CanConfigureServiceMapping) {
		toSerialize["canConfigureServiceMapping"] = o.CanConfigureServiceMapping
	}
	if !isNil(o.CanConfigureSessionSettings) {
		toSerialize["canConfigureSessionSettings"] = o.CanConfigureSessionSettings
	}
	if !isNil(o.CanConfigureTeams) {
		toSerialize["canConfigureTeams"] = o.CanConfigureTeams
	}
	if !isNil(o.CanConfigureUsers) {
		toSerialize["canConfigureUsers"] = o.CanConfigureUsers
	}
	if !isNil(o.CanCreatePublicCustomDashboards) {
		toSerialize["canCreatePublicCustomDashboards"] = o.CanCreatePublicCustomDashboards
	}
	if !isNil(o.CanEditAllAccessibleCustomDashboards) {
		toSerialize["canEditAllAccessibleCustomDashboards"] = o.CanEditAllAccessibleCustomDashboards
	}
	if !isNil(o.CanInstallNewAgents) {
		toSerialize["canInstallNewAgents"] = o.CanInstallNewAgents
	}
	if !isNil(o.CanRunAutomationActions) {
		toSerialize["canRunAutomationActions"] = o.CanRunAutomationActions
	}
	if !isNil(o.CanSeeOnPremLicenseInformation) {
		toSerialize["canSeeOnPremLicenseInformation"] = o.CanSeeOnPremLicenseInformation
	}
	if !isNil(o.CanSeeUsageInformation) {
		toSerialize["canSeeUsageInformation"] = o.CanSeeUsageInformation
	}
	if !isNil(o.CanViewAccountAndBillingInformation) {
		toSerialize["canViewAccountAndBillingInformation"] = o.CanViewAccountAndBillingInformation
	}
	if !isNil(o.CanViewAuditLog) {
		toSerialize["canViewAuditLog"] = o.CanViewAuditLog
	}
	if !isNil(o.CanViewLogs) {
		toSerialize["canViewLogs"] = o.CanViewLogs
	}
	if !isNil(o.CanViewTraceDetails) {
		toSerialize["canViewTraceDetails"] = o.CanViewTraceDetails
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["internalId"] = o.InternalId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableApiToken struct {
	value *ApiToken
	isSet bool
}

func (v NullableApiToken) Get() *ApiToken {
	return v.value
}

func (v *NullableApiToken) Set(val *ApiToken) {
	v.value = val
	v.isSet = true
}

func (v NullableApiToken) IsSet() bool {
	return v.isSet
}

func (v *NullableApiToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiToken(val *ApiToken) *NullableApiToken {
	return &NullableApiToken{value: val, isSet: true}
}

func (v NullableApiToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
