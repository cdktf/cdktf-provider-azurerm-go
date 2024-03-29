// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotsecuritysolution

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		reflect.TypeOf((*IotSecuritySolution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalWorkspace", GoGetter: "AdditionalWorkspace"},
			_jsii_.MemberProperty{JsiiProperty: "additionalWorkspaceInput", GoGetter: "AdditionalWorkspaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disabledDataSources", GoGetter: "DisabledDataSources"},
			_jsii_.MemberProperty{JsiiProperty: "disabledDataSourcesInput", GoGetter: "DisabledDataSourcesInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventsToExport", GoGetter: "EventsToExport"},
			_jsii_.MemberProperty{JsiiProperty: "eventsToExportInput", GoGetter: "EventsToExportInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iothubIds", GoGetter: "IothubIds"},
			_jsii_.MemberProperty{JsiiProperty: "iothubIdsInput", GoGetter: "IothubIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "logAnalyticsWorkspaceId", GoGetter: "LogAnalyticsWorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "logAnalyticsWorkspaceIdInput", GoGetter: "LogAnalyticsWorkspaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "logUnmaskedIpsEnabled", GoGetter: "LogUnmaskedIpsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "logUnmaskedIpsEnabledInput", GoGetter: "LogUnmaskedIpsEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAdditionalWorkspace", GoMethod: "PutAdditionalWorkspace"},
			_jsii_.MemberMethod{JsiiMethod: "putRecommendationsEnabled", GoMethod: "PutRecommendationsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "queryForResources", GoGetter: "QueryForResources"},
			_jsii_.MemberProperty{JsiiProperty: "queryForResourcesInput", GoGetter: "QueryForResourcesInput"},
			_jsii_.MemberProperty{JsiiProperty: "querySubscriptionIds", GoGetter: "QuerySubscriptionIds"},
			_jsii_.MemberProperty{JsiiProperty: "querySubscriptionIdsInput", GoGetter: "QuerySubscriptionIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recommendationsEnabled", GoGetter: "RecommendationsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "recommendationsEnabledInput", GoGetter: "RecommendationsEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalWorkspace", GoMethod: "ResetAdditionalWorkspace"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisabledDataSources", GoMethod: "ResetDisabledDataSources"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventsToExport", GoMethod: "ResetEventsToExport"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogAnalyticsWorkspaceId", GoMethod: "ResetLogAnalyticsWorkspaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogUnmaskedIpsEnabled", GoMethod: "ResetLogUnmaskedIpsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryForResources", GoMethod: "ResetQueryForResources"},
			_jsii_.MemberMethod{JsiiMethod: "resetQuerySubscriptionIds", GoMethod: "ResetQuerySubscriptionIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecommendationsEnabled", GoMethod: "ResetRecommendationsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_IotSecuritySolution{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionAdditionalWorkspace",
		reflect.TypeOf((*IotSecuritySolutionAdditionalWorkspace)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionAdditionalWorkspaceList",
		reflect.TypeOf((*IotSecuritySolutionAdditionalWorkspaceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionAdditionalWorkspaceOutputReference",
		reflect.TypeOf((*IotSecuritySolutionAdditionalWorkspaceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypes", GoGetter: "DataTypes"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypesInput", GoGetter: "DataTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdInput", GoGetter: "WorkspaceIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionConfig",
		reflect.TypeOf((*IotSecuritySolutionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionRecommendationsEnabled",
		reflect.TypeOf((*IotSecuritySolutionRecommendationsEnabled)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionRecommendationsEnabledOutputReference",
		reflect.TypeOf((*IotSecuritySolutionRecommendationsEnabledOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acrAuthentication", GoGetter: "AcrAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "acrAuthenticationInput", GoGetter: "AcrAuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "agentSendUnutilizedMsg", GoGetter: "AgentSendUnutilizedMsg"},
			_jsii_.MemberProperty{JsiiProperty: "agentSendUnutilizedMsgInput", GoGetter: "AgentSendUnutilizedMsgInput"},
			_jsii_.MemberProperty{JsiiProperty: "baseline", GoGetter: "Baseline"},
			_jsii_.MemberProperty{JsiiProperty: "baselineInput", GoGetter: "BaselineInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "edgeHubMemOptimize", GoGetter: "EdgeHubMemOptimize"},
			_jsii_.MemberProperty{JsiiProperty: "edgeHubMemOptimizeInput", GoGetter: "EdgeHubMemOptimizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "edgeLoggingOption", GoGetter: "EdgeLoggingOption"},
			_jsii_.MemberProperty{JsiiProperty: "edgeLoggingOptionInput", GoGetter: "EdgeLoggingOptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "inconsistentModuleSettings", GoGetter: "InconsistentModuleSettings"},
			_jsii_.MemberProperty{JsiiProperty: "inconsistentModuleSettingsInput", GoGetter: "InconsistentModuleSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "installAgent", GoGetter: "InstallAgent"},
			_jsii_.MemberProperty{JsiiProperty: "installAgentInput", GoGetter: "InstallAgentInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipFilterDenyAll", GoGetter: "IpFilterDenyAll"},
			_jsii_.MemberProperty{JsiiProperty: "ipFilterDenyAllInput", GoGetter: "IpFilterDenyAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipFilterPermissiveRule", GoGetter: "IpFilterPermissiveRule"},
			_jsii_.MemberProperty{JsiiProperty: "ipFilterPermissiveRuleInput", GoGetter: "IpFilterPermissiveRuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "openPorts", GoGetter: "OpenPorts"},
			_jsii_.MemberProperty{JsiiProperty: "openPortsInput", GoGetter: "OpenPortsInput"},
			_jsii_.MemberProperty{JsiiProperty: "permissiveFirewallPolicy", GoGetter: "PermissiveFirewallPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "permissiveFirewallPolicyInput", GoGetter: "PermissiveFirewallPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "permissiveInputFirewallRules", GoGetter: "PermissiveInputFirewallRules"},
			_jsii_.MemberProperty{JsiiProperty: "permissiveInputFirewallRulesInput", GoGetter: "PermissiveInputFirewallRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "permissiveOutputFirewallRules", GoGetter: "PermissiveOutputFirewallRules"},
			_jsii_.MemberProperty{JsiiProperty: "permissiveOutputFirewallRulesInput", GoGetter: "PermissiveOutputFirewallRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "privilegedDockerOptions", GoGetter: "PrivilegedDockerOptions"},
			_jsii_.MemberProperty{JsiiProperty: "privilegedDockerOptionsInput", GoGetter: "PrivilegedDockerOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcrAuthentication", GoMethod: "ResetAcrAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetAgentSendUnutilizedMsg", GoMethod: "ResetAgentSendUnutilizedMsg"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseline", GoMethod: "ResetBaseline"},
			_jsii_.MemberMethod{JsiiMethod: "resetEdgeHubMemOptimize", GoMethod: "ResetEdgeHubMemOptimize"},
			_jsii_.MemberMethod{JsiiMethod: "resetEdgeLoggingOption", GoMethod: "ResetEdgeLoggingOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetInconsistentModuleSettings", GoMethod: "ResetInconsistentModuleSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstallAgent", GoMethod: "ResetInstallAgent"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpFilterDenyAll", GoMethod: "ResetIpFilterDenyAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpFilterPermissiveRule", GoMethod: "ResetIpFilterPermissiveRule"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenPorts", GoMethod: "ResetOpenPorts"},
			_jsii_.MemberMethod{JsiiMethod: "resetPermissiveFirewallPolicy", GoMethod: "ResetPermissiveFirewallPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetPermissiveInputFirewallRules", GoMethod: "ResetPermissiveInputFirewallRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetPermissiveOutputFirewallRules", GoMethod: "ResetPermissiveOutputFirewallRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivilegedDockerOptions", GoMethod: "ResetPrivilegedDockerOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedCredentials", GoMethod: "ResetSharedCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetVulnerableTlsCipherSuite", GoMethod: "ResetVulnerableTlsCipherSuite"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sharedCredentials", GoGetter: "SharedCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "sharedCredentialsInput", GoGetter: "SharedCredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vulnerableTlsCipherSuite", GoGetter: "VulnerableTlsCipherSuite"},
			_jsii_.MemberProperty{JsiiProperty: "vulnerableTlsCipherSuiteInput", GoGetter: "VulnerableTlsCipherSuiteInput"},
		},
		func() interface{} {
			j := jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionTimeouts",
		reflect.TypeOf((*IotSecuritySolutionTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionTimeoutsOutputReference",
		reflect.TypeOf((*IotSecuritySolutionTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_IotSecuritySolutionTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
