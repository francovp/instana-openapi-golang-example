# ApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGrantingToken** | **string** |  | 
**CanConfigureAgentRunMode** | Pointer to **bool** |  | [optional] 
**CanConfigureAgents** | Pointer to **bool** |  | [optional] 
**CanConfigureApiTokens** | Pointer to **bool** |  | [optional] 
**CanConfigureApplications** | Pointer to **bool** |  | [optional] 
**CanConfigureAuthenticationMethods** | Pointer to **bool** |  | [optional] 
**CanConfigureAutomationActions** | Pointer to **bool** |  | [optional] 
**CanConfigureCustomAlerts** | Pointer to **bool** |  | [optional] 
**CanConfigureEumApplications** | Pointer to **bool** |  | [optional] 
**CanConfigureGlobalAlertConfigs** | Pointer to **bool** |  | [optional] 
**CanConfigureGlobalAlertPayload** | Pointer to **bool** |  | [optional] 
**CanConfigureIntegrations** | Pointer to **bool** |  | [optional] 
**CanConfigureLogManagement** | Pointer to **bool** |  | [optional] 
**CanConfigureMobileAppMonitoring** | Pointer to **bool** |  | [optional] 
**CanConfigurePersonalApiTokens** | Pointer to **bool** |  | [optional] 
**CanConfigureReleases** | Pointer to **bool** |  | [optional] 
**CanConfigureServiceLevelIndicators** | Pointer to **bool** |  | [optional] 
**CanConfigureServiceMapping** | Pointer to **bool** |  | [optional] 
**CanConfigureSessionSettings** | Pointer to **bool** |  | [optional] 
**CanConfigureTeams** | Pointer to **bool** |  | [optional] 
**CanConfigureUsers** | Pointer to **bool** |  | [optional] 
**CanCreatePublicCustomDashboards** | Pointer to **bool** |  | [optional] 
**CanEditAllAccessibleCustomDashboards** | Pointer to **bool** |  | [optional] 
**CanInstallNewAgents** | Pointer to **bool** |  | [optional] 
**CanRunAutomationActions** | Pointer to **bool** |  | [optional] 
**CanSeeOnPremLicenseInformation** | Pointer to **bool** |  | [optional] 
**CanSeeUsageInformation** | Pointer to **bool** |  | [optional] 
**CanViewAccountAndBillingInformation** | Pointer to **bool** |  | [optional] 
**CanViewAuditLog** | Pointer to **bool** |  | [optional] 
**CanViewLogs** | Pointer to **bool** |  | [optional] 
**CanViewTraceDetails** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InternalId** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewApiToken

`func NewApiToken(accessGrantingToken string, internalId string, name string, ) *ApiToken`

NewApiToken instantiates a new ApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenWithDefaults

`func NewApiTokenWithDefaults() *ApiToken`

NewApiTokenWithDefaults instantiates a new ApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGrantingToken

`func (o *ApiToken) GetAccessGrantingToken() string`

GetAccessGrantingToken returns the AccessGrantingToken field if non-nil, zero value otherwise.

### GetAccessGrantingTokenOk

`func (o *ApiToken) GetAccessGrantingTokenOk() (*string, bool)`

GetAccessGrantingTokenOk returns a tuple with the AccessGrantingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGrantingToken

`func (o *ApiToken) SetAccessGrantingToken(v string)`

SetAccessGrantingToken sets AccessGrantingToken field to given value.


### GetCanConfigureAgentRunMode

`func (o *ApiToken) GetCanConfigureAgentRunMode() bool`

GetCanConfigureAgentRunMode returns the CanConfigureAgentRunMode field if non-nil, zero value otherwise.

### GetCanConfigureAgentRunModeOk

`func (o *ApiToken) GetCanConfigureAgentRunModeOk() (*bool, bool)`

GetCanConfigureAgentRunModeOk returns a tuple with the CanConfigureAgentRunMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureAgentRunMode

`func (o *ApiToken) SetCanConfigureAgentRunMode(v bool)`

SetCanConfigureAgentRunMode sets CanConfigureAgentRunMode field to given value.

### HasCanConfigureAgentRunMode

`func (o *ApiToken) HasCanConfigureAgentRunMode() bool`

HasCanConfigureAgentRunMode returns a boolean if a field has been set.

### GetCanConfigureAgents

`func (o *ApiToken) GetCanConfigureAgents() bool`

GetCanConfigureAgents returns the CanConfigureAgents field if non-nil, zero value otherwise.

### GetCanConfigureAgentsOk

`func (o *ApiToken) GetCanConfigureAgentsOk() (*bool, bool)`

GetCanConfigureAgentsOk returns a tuple with the CanConfigureAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureAgents

`func (o *ApiToken) SetCanConfigureAgents(v bool)`

SetCanConfigureAgents sets CanConfigureAgents field to given value.

### HasCanConfigureAgents

`func (o *ApiToken) HasCanConfigureAgents() bool`

HasCanConfigureAgents returns a boolean if a field has been set.

### GetCanConfigureApiTokens

`func (o *ApiToken) GetCanConfigureApiTokens() bool`

GetCanConfigureApiTokens returns the CanConfigureApiTokens field if non-nil, zero value otherwise.

### GetCanConfigureApiTokensOk

`func (o *ApiToken) GetCanConfigureApiTokensOk() (*bool, bool)`

GetCanConfigureApiTokensOk returns a tuple with the CanConfigureApiTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureApiTokens

`func (o *ApiToken) SetCanConfigureApiTokens(v bool)`

SetCanConfigureApiTokens sets CanConfigureApiTokens field to given value.

### HasCanConfigureApiTokens

`func (o *ApiToken) HasCanConfigureApiTokens() bool`

HasCanConfigureApiTokens returns a boolean if a field has been set.

### GetCanConfigureApplications

`func (o *ApiToken) GetCanConfigureApplications() bool`

GetCanConfigureApplications returns the CanConfigureApplications field if non-nil, zero value otherwise.

### GetCanConfigureApplicationsOk

`func (o *ApiToken) GetCanConfigureApplicationsOk() (*bool, bool)`

GetCanConfigureApplicationsOk returns a tuple with the CanConfigureApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureApplications

`func (o *ApiToken) SetCanConfigureApplications(v bool)`

SetCanConfigureApplications sets CanConfigureApplications field to given value.

### HasCanConfigureApplications

`func (o *ApiToken) HasCanConfigureApplications() bool`

HasCanConfigureApplications returns a boolean if a field has been set.

### GetCanConfigureAuthenticationMethods

`func (o *ApiToken) GetCanConfigureAuthenticationMethods() bool`

GetCanConfigureAuthenticationMethods returns the CanConfigureAuthenticationMethods field if non-nil, zero value otherwise.

### GetCanConfigureAuthenticationMethodsOk

`func (o *ApiToken) GetCanConfigureAuthenticationMethodsOk() (*bool, bool)`

GetCanConfigureAuthenticationMethodsOk returns a tuple with the CanConfigureAuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureAuthenticationMethods

`func (o *ApiToken) SetCanConfigureAuthenticationMethods(v bool)`

SetCanConfigureAuthenticationMethods sets CanConfigureAuthenticationMethods field to given value.

### HasCanConfigureAuthenticationMethods

`func (o *ApiToken) HasCanConfigureAuthenticationMethods() bool`

HasCanConfigureAuthenticationMethods returns a boolean if a field has been set.

### GetCanConfigureAutomationActions

`func (o *ApiToken) GetCanConfigureAutomationActions() bool`

GetCanConfigureAutomationActions returns the CanConfigureAutomationActions field if non-nil, zero value otherwise.

### GetCanConfigureAutomationActionsOk

`func (o *ApiToken) GetCanConfigureAutomationActionsOk() (*bool, bool)`

GetCanConfigureAutomationActionsOk returns a tuple with the CanConfigureAutomationActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureAutomationActions

`func (o *ApiToken) SetCanConfigureAutomationActions(v bool)`

SetCanConfigureAutomationActions sets CanConfigureAutomationActions field to given value.

### HasCanConfigureAutomationActions

`func (o *ApiToken) HasCanConfigureAutomationActions() bool`

HasCanConfigureAutomationActions returns a boolean if a field has been set.

### GetCanConfigureCustomAlerts

`func (o *ApiToken) GetCanConfigureCustomAlerts() bool`

GetCanConfigureCustomAlerts returns the CanConfigureCustomAlerts field if non-nil, zero value otherwise.

### GetCanConfigureCustomAlertsOk

`func (o *ApiToken) GetCanConfigureCustomAlertsOk() (*bool, bool)`

GetCanConfigureCustomAlertsOk returns a tuple with the CanConfigureCustomAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureCustomAlerts

`func (o *ApiToken) SetCanConfigureCustomAlerts(v bool)`

SetCanConfigureCustomAlerts sets CanConfigureCustomAlerts field to given value.

### HasCanConfigureCustomAlerts

`func (o *ApiToken) HasCanConfigureCustomAlerts() bool`

HasCanConfigureCustomAlerts returns a boolean if a field has been set.

### GetCanConfigureEumApplications

`func (o *ApiToken) GetCanConfigureEumApplications() bool`

GetCanConfigureEumApplications returns the CanConfigureEumApplications field if non-nil, zero value otherwise.

### GetCanConfigureEumApplicationsOk

`func (o *ApiToken) GetCanConfigureEumApplicationsOk() (*bool, bool)`

GetCanConfigureEumApplicationsOk returns a tuple with the CanConfigureEumApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureEumApplications

`func (o *ApiToken) SetCanConfigureEumApplications(v bool)`

SetCanConfigureEumApplications sets CanConfigureEumApplications field to given value.

### HasCanConfigureEumApplications

`func (o *ApiToken) HasCanConfigureEumApplications() bool`

HasCanConfigureEumApplications returns a boolean if a field has been set.

### GetCanConfigureGlobalAlertConfigs

`func (o *ApiToken) GetCanConfigureGlobalAlertConfigs() bool`

GetCanConfigureGlobalAlertConfigs returns the CanConfigureGlobalAlertConfigs field if non-nil, zero value otherwise.

### GetCanConfigureGlobalAlertConfigsOk

`func (o *ApiToken) GetCanConfigureGlobalAlertConfigsOk() (*bool, bool)`

GetCanConfigureGlobalAlertConfigsOk returns a tuple with the CanConfigureGlobalAlertConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureGlobalAlertConfigs

`func (o *ApiToken) SetCanConfigureGlobalAlertConfigs(v bool)`

SetCanConfigureGlobalAlertConfigs sets CanConfigureGlobalAlertConfigs field to given value.

### HasCanConfigureGlobalAlertConfigs

`func (o *ApiToken) HasCanConfigureGlobalAlertConfigs() bool`

HasCanConfigureGlobalAlertConfigs returns a boolean if a field has been set.

### GetCanConfigureGlobalAlertPayload

`func (o *ApiToken) GetCanConfigureGlobalAlertPayload() bool`

GetCanConfigureGlobalAlertPayload returns the CanConfigureGlobalAlertPayload field if non-nil, zero value otherwise.

### GetCanConfigureGlobalAlertPayloadOk

`func (o *ApiToken) GetCanConfigureGlobalAlertPayloadOk() (*bool, bool)`

GetCanConfigureGlobalAlertPayloadOk returns a tuple with the CanConfigureGlobalAlertPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureGlobalAlertPayload

`func (o *ApiToken) SetCanConfigureGlobalAlertPayload(v bool)`

SetCanConfigureGlobalAlertPayload sets CanConfigureGlobalAlertPayload field to given value.

### HasCanConfigureGlobalAlertPayload

`func (o *ApiToken) HasCanConfigureGlobalAlertPayload() bool`

HasCanConfigureGlobalAlertPayload returns a boolean if a field has been set.

### GetCanConfigureIntegrations

`func (o *ApiToken) GetCanConfigureIntegrations() bool`

GetCanConfigureIntegrations returns the CanConfigureIntegrations field if non-nil, zero value otherwise.

### GetCanConfigureIntegrationsOk

`func (o *ApiToken) GetCanConfigureIntegrationsOk() (*bool, bool)`

GetCanConfigureIntegrationsOk returns a tuple with the CanConfigureIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureIntegrations

`func (o *ApiToken) SetCanConfigureIntegrations(v bool)`

SetCanConfigureIntegrations sets CanConfigureIntegrations field to given value.

### HasCanConfigureIntegrations

`func (o *ApiToken) HasCanConfigureIntegrations() bool`

HasCanConfigureIntegrations returns a boolean if a field has been set.

### GetCanConfigureLogManagement

`func (o *ApiToken) GetCanConfigureLogManagement() bool`

GetCanConfigureLogManagement returns the CanConfigureLogManagement field if non-nil, zero value otherwise.

### GetCanConfigureLogManagementOk

`func (o *ApiToken) GetCanConfigureLogManagementOk() (*bool, bool)`

GetCanConfigureLogManagementOk returns a tuple with the CanConfigureLogManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureLogManagement

`func (o *ApiToken) SetCanConfigureLogManagement(v bool)`

SetCanConfigureLogManagement sets CanConfigureLogManagement field to given value.

### HasCanConfigureLogManagement

`func (o *ApiToken) HasCanConfigureLogManagement() bool`

HasCanConfigureLogManagement returns a boolean if a field has been set.

### GetCanConfigureMobileAppMonitoring

`func (o *ApiToken) GetCanConfigureMobileAppMonitoring() bool`

GetCanConfigureMobileAppMonitoring returns the CanConfigureMobileAppMonitoring field if non-nil, zero value otherwise.

### GetCanConfigureMobileAppMonitoringOk

`func (o *ApiToken) GetCanConfigureMobileAppMonitoringOk() (*bool, bool)`

GetCanConfigureMobileAppMonitoringOk returns a tuple with the CanConfigureMobileAppMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureMobileAppMonitoring

`func (o *ApiToken) SetCanConfigureMobileAppMonitoring(v bool)`

SetCanConfigureMobileAppMonitoring sets CanConfigureMobileAppMonitoring field to given value.

### HasCanConfigureMobileAppMonitoring

`func (o *ApiToken) HasCanConfigureMobileAppMonitoring() bool`

HasCanConfigureMobileAppMonitoring returns a boolean if a field has been set.

### GetCanConfigurePersonalApiTokens

`func (o *ApiToken) GetCanConfigurePersonalApiTokens() bool`

GetCanConfigurePersonalApiTokens returns the CanConfigurePersonalApiTokens field if non-nil, zero value otherwise.

### GetCanConfigurePersonalApiTokensOk

`func (o *ApiToken) GetCanConfigurePersonalApiTokensOk() (*bool, bool)`

GetCanConfigurePersonalApiTokensOk returns a tuple with the CanConfigurePersonalApiTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigurePersonalApiTokens

`func (o *ApiToken) SetCanConfigurePersonalApiTokens(v bool)`

SetCanConfigurePersonalApiTokens sets CanConfigurePersonalApiTokens field to given value.

### HasCanConfigurePersonalApiTokens

`func (o *ApiToken) HasCanConfigurePersonalApiTokens() bool`

HasCanConfigurePersonalApiTokens returns a boolean if a field has been set.

### GetCanConfigureReleases

`func (o *ApiToken) GetCanConfigureReleases() bool`

GetCanConfigureReleases returns the CanConfigureReleases field if non-nil, zero value otherwise.

### GetCanConfigureReleasesOk

`func (o *ApiToken) GetCanConfigureReleasesOk() (*bool, bool)`

GetCanConfigureReleasesOk returns a tuple with the CanConfigureReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureReleases

`func (o *ApiToken) SetCanConfigureReleases(v bool)`

SetCanConfigureReleases sets CanConfigureReleases field to given value.

### HasCanConfigureReleases

`func (o *ApiToken) HasCanConfigureReleases() bool`

HasCanConfigureReleases returns a boolean if a field has been set.

### GetCanConfigureServiceLevelIndicators

`func (o *ApiToken) GetCanConfigureServiceLevelIndicators() bool`

GetCanConfigureServiceLevelIndicators returns the CanConfigureServiceLevelIndicators field if non-nil, zero value otherwise.

### GetCanConfigureServiceLevelIndicatorsOk

`func (o *ApiToken) GetCanConfigureServiceLevelIndicatorsOk() (*bool, bool)`

GetCanConfigureServiceLevelIndicatorsOk returns a tuple with the CanConfigureServiceLevelIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureServiceLevelIndicators

`func (o *ApiToken) SetCanConfigureServiceLevelIndicators(v bool)`

SetCanConfigureServiceLevelIndicators sets CanConfigureServiceLevelIndicators field to given value.

### HasCanConfigureServiceLevelIndicators

`func (o *ApiToken) HasCanConfigureServiceLevelIndicators() bool`

HasCanConfigureServiceLevelIndicators returns a boolean if a field has been set.

### GetCanConfigureServiceMapping

`func (o *ApiToken) GetCanConfigureServiceMapping() bool`

GetCanConfigureServiceMapping returns the CanConfigureServiceMapping field if non-nil, zero value otherwise.

### GetCanConfigureServiceMappingOk

`func (o *ApiToken) GetCanConfigureServiceMappingOk() (*bool, bool)`

GetCanConfigureServiceMappingOk returns a tuple with the CanConfigureServiceMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureServiceMapping

`func (o *ApiToken) SetCanConfigureServiceMapping(v bool)`

SetCanConfigureServiceMapping sets CanConfigureServiceMapping field to given value.

### HasCanConfigureServiceMapping

`func (o *ApiToken) HasCanConfigureServiceMapping() bool`

HasCanConfigureServiceMapping returns a boolean if a field has been set.

### GetCanConfigureSessionSettings

`func (o *ApiToken) GetCanConfigureSessionSettings() bool`

GetCanConfigureSessionSettings returns the CanConfigureSessionSettings field if non-nil, zero value otherwise.

### GetCanConfigureSessionSettingsOk

`func (o *ApiToken) GetCanConfigureSessionSettingsOk() (*bool, bool)`

GetCanConfigureSessionSettingsOk returns a tuple with the CanConfigureSessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureSessionSettings

`func (o *ApiToken) SetCanConfigureSessionSettings(v bool)`

SetCanConfigureSessionSettings sets CanConfigureSessionSettings field to given value.

### HasCanConfigureSessionSettings

`func (o *ApiToken) HasCanConfigureSessionSettings() bool`

HasCanConfigureSessionSettings returns a boolean if a field has been set.

### GetCanConfigureTeams

`func (o *ApiToken) GetCanConfigureTeams() bool`

GetCanConfigureTeams returns the CanConfigureTeams field if non-nil, zero value otherwise.

### GetCanConfigureTeamsOk

`func (o *ApiToken) GetCanConfigureTeamsOk() (*bool, bool)`

GetCanConfigureTeamsOk returns a tuple with the CanConfigureTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureTeams

`func (o *ApiToken) SetCanConfigureTeams(v bool)`

SetCanConfigureTeams sets CanConfigureTeams field to given value.

### HasCanConfigureTeams

`func (o *ApiToken) HasCanConfigureTeams() bool`

HasCanConfigureTeams returns a boolean if a field has been set.

### GetCanConfigureUsers

`func (o *ApiToken) GetCanConfigureUsers() bool`

GetCanConfigureUsers returns the CanConfigureUsers field if non-nil, zero value otherwise.

### GetCanConfigureUsersOk

`func (o *ApiToken) GetCanConfigureUsersOk() (*bool, bool)`

GetCanConfigureUsersOk returns a tuple with the CanConfigureUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureUsers

`func (o *ApiToken) SetCanConfigureUsers(v bool)`

SetCanConfigureUsers sets CanConfigureUsers field to given value.

### HasCanConfigureUsers

`func (o *ApiToken) HasCanConfigureUsers() bool`

HasCanConfigureUsers returns a boolean if a field has been set.

### GetCanCreatePublicCustomDashboards

`func (o *ApiToken) GetCanCreatePublicCustomDashboards() bool`

GetCanCreatePublicCustomDashboards returns the CanCreatePublicCustomDashboards field if non-nil, zero value otherwise.

### GetCanCreatePublicCustomDashboardsOk

`func (o *ApiToken) GetCanCreatePublicCustomDashboardsOk() (*bool, bool)`

GetCanCreatePublicCustomDashboardsOk returns a tuple with the CanCreatePublicCustomDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreatePublicCustomDashboards

`func (o *ApiToken) SetCanCreatePublicCustomDashboards(v bool)`

SetCanCreatePublicCustomDashboards sets CanCreatePublicCustomDashboards field to given value.

### HasCanCreatePublicCustomDashboards

`func (o *ApiToken) HasCanCreatePublicCustomDashboards() bool`

HasCanCreatePublicCustomDashboards returns a boolean if a field has been set.

### GetCanEditAllAccessibleCustomDashboards

`func (o *ApiToken) GetCanEditAllAccessibleCustomDashboards() bool`

GetCanEditAllAccessibleCustomDashboards returns the CanEditAllAccessibleCustomDashboards field if non-nil, zero value otherwise.

### GetCanEditAllAccessibleCustomDashboardsOk

`func (o *ApiToken) GetCanEditAllAccessibleCustomDashboardsOk() (*bool, bool)`

GetCanEditAllAccessibleCustomDashboardsOk returns a tuple with the CanEditAllAccessibleCustomDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditAllAccessibleCustomDashboards

`func (o *ApiToken) SetCanEditAllAccessibleCustomDashboards(v bool)`

SetCanEditAllAccessibleCustomDashboards sets CanEditAllAccessibleCustomDashboards field to given value.

### HasCanEditAllAccessibleCustomDashboards

`func (o *ApiToken) HasCanEditAllAccessibleCustomDashboards() bool`

HasCanEditAllAccessibleCustomDashboards returns a boolean if a field has been set.

### GetCanInstallNewAgents

`func (o *ApiToken) GetCanInstallNewAgents() bool`

GetCanInstallNewAgents returns the CanInstallNewAgents field if non-nil, zero value otherwise.

### GetCanInstallNewAgentsOk

`func (o *ApiToken) GetCanInstallNewAgentsOk() (*bool, bool)`

GetCanInstallNewAgentsOk returns a tuple with the CanInstallNewAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInstallNewAgents

`func (o *ApiToken) SetCanInstallNewAgents(v bool)`

SetCanInstallNewAgents sets CanInstallNewAgents field to given value.

### HasCanInstallNewAgents

`func (o *ApiToken) HasCanInstallNewAgents() bool`

HasCanInstallNewAgents returns a boolean if a field has been set.

### GetCanRunAutomationActions

`func (o *ApiToken) GetCanRunAutomationActions() bool`

GetCanRunAutomationActions returns the CanRunAutomationActions field if non-nil, zero value otherwise.

### GetCanRunAutomationActionsOk

`func (o *ApiToken) GetCanRunAutomationActionsOk() (*bool, bool)`

GetCanRunAutomationActionsOk returns a tuple with the CanRunAutomationActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRunAutomationActions

`func (o *ApiToken) SetCanRunAutomationActions(v bool)`

SetCanRunAutomationActions sets CanRunAutomationActions field to given value.

### HasCanRunAutomationActions

`func (o *ApiToken) HasCanRunAutomationActions() bool`

HasCanRunAutomationActions returns a boolean if a field has been set.

### GetCanSeeOnPremLicenseInformation

`func (o *ApiToken) GetCanSeeOnPremLicenseInformation() bool`

GetCanSeeOnPremLicenseInformation returns the CanSeeOnPremLicenseInformation field if non-nil, zero value otherwise.

### GetCanSeeOnPremLicenseInformationOk

`func (o *ApiToken) GetCanSeeOnPremLicenseInformationOk() (*bool, bool)`

GetCanSeeOnPremLicenseInformationOk returns a tuple with the CanSeeOnPremLicenseInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeeOnPremLicenseInformation

`func (o *ApiToken) SetCanSeeOnPremLicenseInformation(v bool)`

SetCanSeeOnPremLicenseInformation sets CanSeeOnPremLicenseInformation field to given value.

### HasCanSeeOnPremLicenseInformation

`func (o *ApiToken) HasCanSeeOnPremLicenseInformation() bool`

HasCanSeeOnPremLicenseInformation returns a boolean if a field has been set.

### GetCanSeeUsageInformation

`func (o *ApiToken) GetCanSeeUsageInformation() bool`

GetCanSeeUsageInformation returns the CanSeeUsageInformation field if non-nil, zero value otherwise.

### GetCanSeeUsageInformationOk

`func (o *ApiToken) GetCanSeeUsageInformationOk() (*bool, bool)`

GetCanSeeUsageInformationOk returns a tuple with the CanSeeUsageInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeeUsageInformation

`func (o *ApiToken) SetCanSeeUsageInformation(v bool)`

SetCanSeeUsageInformation sets CanSeeUsageInformation field to given value.

### HasCanSeeUsageInformation

`func (o *ApiToken) HasCanSeeUsageInformation() bool`

HasCanSeeUsageInformation returns a boolean if a field has been set.

### GetCanViewAccountAndBillingInformation

`func (o *ApiToken) GetCanViewAccountAndBillingInformation() bool`

GetCanViewAccountAndBillingInformation returns the CanViewAccountAndBillingInformation field if non-nil, zero value otherwise.

### GetCanViewAccountAndBillingInformationOk

`func (o *ApiToken) GetCanViewAccountAndBillingInformationOk() (*bool, bool)`

GetCanViewAccountAndBillingInformationOk returns a tuple with the CanViewAccountAndBillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewAccountAndBillingInformation

`func (o *ApiToken) SetCanViewAccountAndBillingInformation(v bool)`

SetCanViewAccountAndBillingInformation sets CanViewAccountAndBillingInformation field to given value.

### HasCanViewAccountAndBillingInformation

`func (o *ApiToken) HasCanViewAccountAndBillingInformation() bool`

HasCanViewAccountAndBillingInformation returns a boolean if a field has been set.

### GetCanViewAuditLog

`func (o *ApiToken) GetCanViewAuditLog() bool`

GetCanViewAuditLog returns the CanViewAuditLog field if non-nil, zero value otherwise.

### GetCanViewAuditLogOk

`func (o *ApiToken) GetCanViewAuditLogOk() (*bool, bool)`

GetCanViewAuditLogOk returns a tuple with the CanViewAuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewAuditLog

`func (o *ApiToken) SetCanViewAuditLog(v bool)`

SetCanViewAuditLog sets CanViewAuditLog field to given value.

### HasCanViewAuditLog

`func (o *ApiToken) HasCanViewAuditLog() bool`

HasCanViewAuditLog returns a boolean if a field has been set.

### GetCanViewLogs

`func (o *ApiToken) GetCanViewLogs() bool`

GetCanViewLogs returns the CanViewLogs field if non-nil, zero value otherwise.

### GetCanViewLogsOk

`func (o *ApiToken) GetCanViewLogsOk() (*bool, bool)`

GetCanViewLogsOk returns a tuple with the CanViewLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewLogs

`func (o *ApiToken) SetCanViewLogs(v bool)`

SetCanViewLogs sets CanViewLogs field to given value.

### HasCanViewLogs

`func (o *ApiToken) HasCanViewLogs() bool`

HasCanViewLogs returns a boolean if a field has been set.

### GetCanViewTraceDetails

`func (o *ApiToken) GetCanViewTraceDetails() bool`

GetCanViewTraceDetails returns the CanViewTraceDetails field if non-nil, zero value otherwise.

### GetCanViewTraceDetailsOk

`func (o *ApiToken) GetCanViewTraceDetailsOk() (*bool, bool)`

GetCanViewTraceDetailsOk returns a tuple with the CanViewTraceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewTraceDetails

`func (o *ApiToken) SetCanViewTraceDetails(v bool)`

SetCanViewTraceDetails sets CanViewTraceDetails field to given value.

### HasCanViewTraceDetails

`func (o *ApiToken) HasCanViewTraceDetails() bool`

HasCanViewTraceDetails returns a boolean if a field has been set.

### GetId

`func (o *ApiToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *ApiToken) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ApiToken) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ApiToken) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.


### GetName

`func (o *ApiToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiToken) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


