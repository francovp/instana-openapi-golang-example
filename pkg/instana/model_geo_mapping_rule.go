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

// GeoMappingRule struct for GeoMappingRule
type GeoMappingRule struct {
	AccuracyRadius           *int64           `json:"accuracyRadius,omitempty"`
	Cidr                     string           `json:"cidr"`
	City                     *string          `json:"city,omitempty"`
	Continent                *string          `json:"continent,omitempty"`
	ContinentCode            *string          `json:"continentCode,omitempty"`
	Country                  *string          `json:"country,omitempty"`
	CountryCode              *string          `json:"countryCode,omitempty"`
	Latitude                 *float64         `json:"latitude,omitempty"`
	LeastSpecificSubdivision *GeoSubdivision  `json:"leastSpecificSubdivision,omitempty"`
	Longitude                *float64         `json:"longitude,omitempty"`
	Subdivisions             []GeoSubdivision `json:"subdivisions"`
}

// NewGeoMappingRule instantiates a new GeoMappingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoMappingRule(cidr string, subdivisions []GeoSubdivision) *GeoMappingRule {
	this := GeoMappingRule{}
	this.Cidr = cidr
	this.Subdivisions = subdivisions
	return &this
}

// NewGeoMappingRuleWithDefaults instantiates a new GeoMappingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoMappingRuleWithDefaults() *GeoMappingRule {
	this := GeoMappingRule{}
	return &this
}

// GetAccuracyRadius returns the AccuracyRadius field value if set, zero value otherwise.
func (o *GeoMappingRule) GetAccuracyRadius() int64 {
	if o == nil || isNil(o.AccuracyRadius) {
		var ret int64
		return ret
	}
	return *o.AccuracyRadius
}

// GetAccuracyRadiusOk returns a tuple with the AccuracyRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetAccuracyRadiusOk() (*int64, bool) {
	if o == nil || isNil(o.AccuracyRadius) {
		return nil, false
	}
	return o.AccuracyRadius, true
}

// HasAccuracyRadius returns a boolean if a field has been set.
func (o *GeoMappingRule) HasAccuracyRadius() bool {
	if o != nil && !isNil(o.AccuracyRadius) {
		return true
	}

	return false
}

// SetAccuracyRadius gets a reference to the given int64 and assigns it to the AccuracyRadius field.
func (o *GeoMappingRule) SetAccuracyRadius(v int64) {
	o.AccuracyRadius = &v
}

// GetCidr returns the Cidr field value
func (o *GeoMappingRule) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *GeoMappingRule) SetCidr(v string) {
	o.Cidr = v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *GeoMappingRule) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *GeoMappingRule) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *GeoMappingRule) SetCity(v string) {
	o.City = &v
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *GeoMappingRule) GetContinent() string {
	if o == nil || isNil(o.Continent) {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetContinentOk() (*string, bool) {
	if o == nil || isNil(o.Continent) {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *GeoMappingRule) HasContinent() bool {
	if o != nil && !isNil(o.Continent) {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *GeoMappingRule) SetContinent(v string) {
	o.Continent = &v
}

// GetContinentCode returns the ContinentCode field value if set, zero value otherwise.
func (o *GeoMappingRule) GetContinentCode() string {
	if o == nil || isNil(o.ContinentCode) {
		var ret string
		return ret
	}
	return *o.ContinentCode
}

// GetContinentCodeOk returns a tuple with the ContinentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetContinentCodeOk() (*string, bool) {
	if o == nil || isNil(o.ContinentCode) {
		return nil, false
	}
	return o.ContinentCode, true
}

// HasContinentCode returns a boolean if a field has been set.
func (o *GeoMappingRule) HasContinentCode() bool {
	if o != nil && !isNil(o.ContinentCode) {
		return true
	}

	return false
}

// SetContinentCode gets a reference to the given string and assigns it to the ContinentCode field.
func (o *GeoMappingRule) SetContinentCode(v string) {
	o.ContinentCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GeoMappingRule) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GeoMappingRule) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GeoMappingRule) SetCountry(v string) {
	o.Country = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *GeoMappingRule) GetCountryCode() string {
	if o == nil || isNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetCountryCodeOk() (*string, bool) {
	if o == nil || isNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *GeoMappingRule) HasCountryCode() bool {
	if o != nil && !isNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *GeoMappingRule) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *GeoMappingRule) GetLatitude() float64 {
	if o == nil || isNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetLatitudeOk() (*float64, bool) {
	if o == nil || isNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *GeoMappingRule) HasLatitude() bool {
	if o != nil && !isNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *GeoMappingRule) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLeastSpecificSubdivision returns the LeastSpecificSubdivision field value if set, zero value otherwise.
func (o *GeoMappingRule) GetLeastSpecificSubdivision() GeoSubdivision {
	if o == nil || isNil(o.LeastSpecificSubdivision) {
		var ret GeoSubdivision
		return ret
	}
	return *o.LeastSpecificSubdivision
}

// GetLeastSpecificSubdivisionOk returns a tuple with the LeastSpecificSubdivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetLeastSpecificSubdivisionOk() (*GeoSubdivision, bool) {
	if o == nil || isNil(o.LeastSpecificSubdivision) {
		return nil, false
	}
	return o.LeastSpecificSubdivision, true
}

// HasLeastSpecificSubdivision returns a boolean if a field has been set.
func (o *GeoMappingRule) HasLeastSpecificSubdivision() bool {
	if o != nil && !isNil(o.LeastSpecificSubdivision) {
		return true
	}

	return false
}

// SetLeastSpecificSubdivision gets a reference to the given GeoSubdivision and assigns it to the LeastSpecificSubdivision field.
func (o *GeoMappingRule) SetLeastSpecificSubdivision(v GeoSubdivision) {
	o.LeastSpecificSubdivision = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *GeoMappingRule) GetLongitude() float64 {
	if o == nil || isNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetLongitudeOk() (*float64, bool) {
	if o == nil || isNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *GeoMappingRule) HasLongitude() bool {
	if o != nil && !isNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *GeoMappingRule) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetSubdivisions returns the Subdivisions field value
func (o *GeoMappingRule) GetSubdivisions() []GeoSubdivision {
	if o == nil {
		var ret []GeoSubdivision
		return ret
	}

	return o.Subdivisions
}

// GetSubdivisionsOk returns a tuple with the Subdivisions field value
// and a boolean to check if the value has been set.
func (o *GeoMappingRule) GetSubdivisionsOk() ([]GeoSubdivision, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subdivisions, true
}

// SetSubdivisions sets field value
func (o *GeoMappingRule) SetSubdivisions(v []GeoSubdivision) {
	o.Subdivisions = v
}

func (o GeoMappingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccuracyRadius) {
		toSerialize["accuracyRadius"] = o.AccuracyRadius
	}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
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
	if !isNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !isNil(o.LeastSpecificSubdivision) {
		toSerialize["leastSpecificSubdivision"] = o.LeastSpecificSubdivision
	}
	if !isNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if true {
		toSerialize["subdivisions"] = o.Subdivisions
	}
	return json.Marshal(toSerialize)
}

type NullableGeoMappingRule struct {
	value *GeoMappingRule
	isSet bool
}

func (v NullableGeoMappingRule) Get() *GeoMappingRule {
	return v.value
}

func (v *NullableGeoMappingRule) Set(val *GeoMappingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoMappingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoMappingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoMappingRule(val *GeoMappingRule) *NullableGeoMappingRule {
	return &NullableGeoMappingRule{value: val, isSet: true}
}

func (v NullableGeoMappingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoMappingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
