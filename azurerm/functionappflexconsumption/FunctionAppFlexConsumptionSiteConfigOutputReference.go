// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/functionappflexconsumption/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppFlexConsumptionSiteConfigOutputReference interface {
	cdktf.ComplexObject
	ApiDefinitionUrl() *string
	SetApiDefinitionUrl(val *string)
	ApiDefinitionUrlInput() *string
	ApiManagementApiId() *string
	SetApiManagementApiId(val *string)
	ApiManagementApiIdInput() *string
	AppCommandLine() *string
	SetAppCommandLine(val *string)
	AppCommandLineInput() *string
	ApplicationInsightsConnectionString() *string
	SetApplicationInsightsConnectionString(val *string)
	ApplicationInsightsConnectionStringInput() *string
	ApplicationInsightsKey() *string
	SetApplicationInsightsKey(val *string)
	ApplicationInsightsKeyInput() *string
	AppServiceLogs() FunctionAppFlexConsumptionSiteConfigAppServiceLogsOutputReference
	AppServiceLogsInput() *FunctionAppFlexConsumptionSiteConfigAppServiceLogs
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ContainerRegistryManagedIdentityClientId() *string
	SetContainerRegistryManagedIdentityClientId(val *string)
	ContainerRegistryManagedIdentityClientIdInput() *string
	ContainerRegistryUseManagedIdentity() interface{}
	SetContainerRegistryUseManagedIdentity(val interface{})
	ContainerRegistryUseManagedIdentityInput() interface{}
	Cors() FunctionAppFlexConsumptionSiteConfigCorsOutputReference
	CorsInput() *FunctionAppFlexConsumptionSiteConfigCors
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultDocuments() *[]*string
	SetDefaultDocuments(val *[]*string)
	DefaultDocumentsInput() *[]*string
	DetailedErrorLoggingEnabled() cdktf.IResolvable
	ElasticInstanceMinimum() *float64
	SetElasticInstanceMinimum(val *float64)
	ElasticInstanceMinimumInput() *float64
	// Experimental.
	Fqn() *string
	HealthCheckEvictionTimeInMin() *float64
	SetHealthCheckEvictionTimeInMin(val *float64)
	HealthCheckEvictionTimeInMinInput() *float64
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	Http2Enabled() interface{}
	SetHttp2Enabled(val interface{})
	Http2EnabledInput() interface{}
	InternalValue() *FunctionAppFlexConsumptionSiteConfig
	SetInternalValue(val *FunctionAppFlexConsumptionSiteConfig)
	IpRestriction() FunctionAppFlexConsumptionSiteConfigIpRestrictionList
	IpRestrictionDefaultAction() *string
	SetIpRestrictionDefaultAction(val *string)
	IpRestrictionDefaultActionInput() *string
	IpRestrictionInput() interface{}
	LoadBalancingMode() *string
	SetLoadBalancingMode(val *string)
	LoadBalancingModeInput() *string
	ManagedPipelineMode() *string
	SetManagedPipelineMode(val *string)
	ManagedPipelineModeInput() *string
	MinimumTlsVersion() *string
	SetMinimumTlsVersion(val *string)
	MinimumTlsVersionInput() *string
	RemoteDebuggingEnabled() interface{}
	SetRemoteDebuggingEnabled(val interface{})
	RemoteDebuggingEnabledInput() interface{}
	RemoteDebuggingVersion() *string
	SetRemoteDebuggingVersion(val *string)
	RemoteDebuggingVersionInput() *string
	RuntimeScaleMonitoringEnabled() interface{}
	SetRuntimeScaleMonitoringEnabled(val interface{})
	RuntimeScaleMonitoringEnabledInput() interface{}
	ScmIpRestriction() FunctionAppFlexConsumptionSiteConfigScmIpRestrictionList
	ScmIpRestrictionDefaultAction() *string
	SetScmIpRestrictionDefaultAction(val *string)
	ScmIpRestrictionDefaultActionInput() *string
	ScmIpRestrictionInput() interface{}
	ScmMinimumTlsVersion() *string
	SetScmMinimumTlsVersion(val *string)
	ScmMinimumTlsVersionInput() *string
	ScmType() *string
	ScmUseMainIpRestriction() interface{}
	SetScmUseMainIpRestriction(val interface{})
	ScmUseMainIpRestrictionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Use32BitWorker() interface{}
	SetUse32BitWorker(val interface{})
	Use32BitWorkerInput() interface{}
	VnetRouteAllEnabled() interface{}
	SetVnetRouteAllEnabled(val interface{})
	VnetRouteAllEnabledInput() interface{}
	WebsocketsEnabled() interface{}
	SetWebsocketsEnabled(val interface{})
	WebsocketsEnabledInput() interface{}
	WorkerCount() *float64
	SetWorkerCount(val *float64)
	WorkerCountInput() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAppServiceLogs(value *FunctionAppFlexConsumptionSiteConfigAppServiceLogs)
	PutCors(value *FunctionAppFlexConsumptionSiteConfigCors)
	PutIpRestriction(value interface{})
	PutScmIpRestriction(value interface{})
	ResetApiDefinitionUrl()
	ResetApiManagementApiId()
	ResetAppCommandLine()
	ResetApplicationInsightsConnectionString()
	ResetApplicationInsightsKey()
	ResetAppServiceLogs()
	ResetContainerRegistryManagedIdentityClientId()
	ResetContainerRegistryUseManagedIdentity()
	ResetCors()
	ResetDefaultDocuments()
	ResetElasticInstanceMinimum()
	ResetHealthCheckEvictionTimeInMin()
	ResetHealthCheckPath()
	ResetHttp2Enabled()
	ResetIpRestriction()
	ResetIpRestrictionDefaultAction()
	ResetLoadBalancingMode()
	ResetManagedPipelineMode()
	ResetMinimumTlsVersion()
	ResetRemoteDebuggingEnabled()
	ResetRemoteDebuggingVersion()
	ResetRuntimeScaleMonitoringEnabled()
	ResetScmIpRestriction()
	ResetScmIpRestrictionDefaultAction()
	ResetScmMinimumTlsVersion()
	ResetScmUseMainIpRestriction()
	ResetUse32BitWorker()
	ResetVnetRouteAllEnabled()
	ResetWebsocketsEnabled()
	ResetWorkerCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionAppFlexConsumptionSiteConfigOutputReference
type jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ApiDefinitionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ApiDefinitionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ApiManagementApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ApiManagementApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) AppCommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) AppCommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ApplicationInsightsConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ApplicationInsightsConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ApplicationInsightsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ApplicationInsightsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) AppServiceLogs() FunctionAppFlexConsumptionSiteConfigAppServiceLogsOutputReference {
	var returns FunctionAppFlexConsumptionSiteConfigAppServiceLogsOutputReference
	_jsii_.Get(
		j,
		"appServiceLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) AppServiceLogsInput() *FunctionAppFlexConsumptionSiteConfigAppServiceLogs {
	var returns *FunctionAppFlexConsumptionSiteConfigAppServiceLogs
	_jsii_.Get(
		j,
		"appServiceLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ContainerRegistryManagedIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryManagedIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ContainerRegistryManagedIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryManagedIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ContainerRegistryUseManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryUseManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ContainerRegistryUseManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryUseManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) Cors() FunctionAppFlexConsumptionSiteConfigCorsOutputReference {
	var returns FunctionAppFlexConsumptionSiteConfigCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) CorsInput() *FunctionAppFlexConsumptionSiteConfigCors {
	var returns *FunctionAppFlexConsumptionSiteConfigCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) DefaultDocuments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) DefaultDocumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) DetailedErrorLoggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"detailedErrorLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ElasticInstanceMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ElasticInstanceMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) HealthCheckEvictionTimeInMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) HealthCheckEvictionTimeInMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) Http2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) Http2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) InternalValue() *FunctionAppFlexConsumptionSiteConfig {
	var returns *FunctionAppFlexConsumptionSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) IpRestriction() FunctionAppFlexConsumptionSiteConfigIpRestrictionList {
	var returns FunctionAppFlexConsumptionSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) IpRestrictionDefaultAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRestrictionDefaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) IpRestrictionDefaultActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRestrictionDefaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) IpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) LoadBalancingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) LoadBalancingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ManagedPipelineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ManagedPipelineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) MinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) RemoteDebuggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) RemoteDebuggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) RemoteDebuggingVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) RemoteDebuggingVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) RuntimeScaleMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) RuntimeScaleMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmIpRestriction() FunctionAppFlexConsumptionSiteConfigScmIpRestrictionList {
	var returns FunctionAppFlexConsumptionSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmIpRestrictionDefaultAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmIpRestrictionDefaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmIpRestrictionDefaultActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmIpRestrictionDefaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmMinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmMinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmUseMainIpRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ScmUseMainIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) Use32BitWorker() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) Use32BitWorkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) VnetRouteAllEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) VnetRouteAllEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) WebsocketsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) WebsocketsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) WorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) WorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCountInput",
		&returns,
	)
	return returns
}


func NewFunctionAppFlexConsumptionSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FunctionAppFlexConsumptionSiteConfigOutputReference {
	_init_.Initialize()

	if err := validateNewFunctionAppFlexConsumptionSiteConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumptionSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFunctionAppFlexConsumptionSiteConfigOutputReference_Override(f FunctionAppFlexConsumptionSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumptionSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetApiDefinitionUrl(val *string) {
	if err := j.validateSetApiDefinitionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiDefinitionUrl",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetApiManagementApiId(val *string) {
	if err := j.validateSetApiManagementApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementApiId",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetAppCommandLine(val *string) {
	if err := j.validateSetAppCommandLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appCommandLine",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetApplicationInsightsConnectionString(val *string) {
	if err := j.validateSetApplicationInsightsConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationInsightsConnectionString",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetApplicationInsightsKey(val *string) {
	if err := j.validateSetApplicationInsightsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationInsightsKey",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetContainerRegistryManagedIdentityClientId(val *string) {
	if err := j.validateSetContainerRegistryManagedIdentityClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryManagedIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetContainerRegistryUseManagedIdentity(val interface{}) {
	if err := j.validateSetContainerRegistryUseManagedIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryUseManagedIdentity",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetDefaultDocuments(val *[]*string) {
	if err := j.validateSetDefaultDocumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDocuments",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetElasticInstanceMinimum(val *float64) {
	if err := j.validateSetElasticInstanceMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticInstanceMinimum",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetHealthCheckEvictionTimeInMin(val *float64) {
	if err := j.validateSetHealthCheckEvictionTimeInMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckEvictionTimeInMin",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetHttp2Enabled(val interface{}) {
	if err := j.validateSetHttp2EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Enabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetInternalValue(val *FunctionAppFlexConsumptionSiteConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetIpRestrictionDefaultAction(val *string) {
	if err := j.validateSetIpRestrictionDefaultActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRestrictionDefaultAction",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetLoadBalancingMode(val *string) {
	if err := j.validateSetLoadBalancingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingMode",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetManagedPipelineMode(val *string) {
	if err := j.validateSetManagedPipelineModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedPipelineMode",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetMinimumTlsVersion(val *string) {
	if err := j.validateSetMinimumTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetRemoteDebuggingEnabled(val interface{}) {
	if err := j.validateSetRemoteDebuggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDebuggingEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetRemoteDebuggingVersion(val *string) {
	if err := j.validateSetRemoteDebuggingVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDebuggingVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetRuntimeScaleMonitoringEnabled(val interface{}) {
	if err := j.validateSetRuntimeScaleMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeScaleMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetScmIpRestrictionDefaultAction(val *string) {
	if err := j.validateSetScmIpRestrictionDefaultActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmIpRestrictionDefaultAction",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetScmMinimumTlsVersion(val *string) {
	if err := j.validateSetScmMinimumTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmMinimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetScmUseMainIpRestriction(val interface{}) {
	if err := j.validateSetScmUseMainIpRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmUseMainIpRestriction",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetUse32BitWorker(val interface{}) {
	if err := j.validateSetUse32BitWorkerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"use32BitWorker",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetVnetRouteAllEnabled(val interface{}) {
	if err := j.validateSetVnetRouteAllEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetRouteAllEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetWebsocketsEnabled(val interface{}) {
	if err := j.validateSetWebsocketsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websocketsEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference)SetWorkerCount(val *float64) {
	if err := j.validateSetWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerCount",
		val,
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) PutAppServiceLogs(value *FunctionAppFlexConsumptionSiteConfigAppServiceLogs) {
	if err := f.validatePutAppServiceLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAppServiceLogs",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) PutCors(value *FunctionAppFlexConsumptionSiteConfigCors) {
	if err := f.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCors",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) PutIpRestriction(value interface{}) {
	if err := f.validatePutIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putIpRestriction",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) PutScmIpRestriction(value interface{}) {
	if err := f.validatePutScmIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putScmIpRestriction",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetApiDefinitionUrl() {
	_jsii_.InvokeVoid(
		f,
		"resetApiDefinitionUrl",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetApiManagementApiId() {
	_jsii_.InvokeVoid(
		f,
		"resetApiManagementApiId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetAppCommandLine() {
	_jsii_.InvokeVoid(
		f,
		"resetAppCommandLine",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetApplicationInsightsConnectionString() {
	_jsii_.InvokeVoid(
		f,
		"resetApplicationInsightsConnectionString",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetApplicationInsightsKey() {
	_jsii_.InvokeVoid(
		f,
		"resetApplicationInsightsKey",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetAppServiceLogs() {
	_jsii_.InvokeVoid(
		f,
		"resetAppServiceLogs",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetContainerRegistryManagedIdentityClientId() {
	_jsii_.InvokeVoid(
		f,
		"resetContainerRegistryManagedIdentityClientId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetContainerRegistryUseManagedIdentity() {
	_jsii_.InvokeVoid(
		f,
		"resetContainerRegistryUseManagedIdentity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		f,
		"resetCors",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetDefaultDocuments() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaultDocuments",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetElasticInstanceMinimum() {
	_jsii_.InvokeVoid(
		f,
		"resetElasticInstanceMinimum",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetHealthCheckEvictionTimeInMin() {
	_jsii_.InvokeVoid(
		f,
		"resetHealthCheckEvictionTimeInMin",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		f,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetHttp2Enabled() {
	_jsii_.InvokeVoid(
		f,
		"resetHttp2Enabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetIpRestriction() {
	_jsii_.InvokeVoid(
		f,
		"resetIpRestriction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetIpRestrictionDefaultAction() {
	_jsii_.InvokeVoid(
		f,
		"resetIpRestrictionDefaultAction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetLoadBalancingMode() {
	_jsii_.InvokeVoid(
		f,
		"resetLoadBalancingMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetManagedPipelineMode() {
	_jsii_.InvokeVoid(
		f,
		"resetManagedPipelineMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetMinimumTlsVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetRemoteDebuggingEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetRemoteDebuggingEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetRemoteDebuggingVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetRemoteDebuggingVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetRuntimeScaleMonitoringEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetRuntimeScaleMonitoringEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetScmIpRestriction() {
	_jsii_.InvokeVoid(
		f,
		"resetScmIpRestriction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetScmIpRestrictionDefaultAction() {
	_jsii_.InvokeVoid(
		f,
		"resetScmIpRestrictionDefaultAction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetScmMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetScmMinimumTlsVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetScmUseMainIpRestriction() {
	_jsii_.InvokeVoid(
		f,
		"resetScmUseMainIpRestriction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetUse32BitWorker() {
	_jsii_.InvokeVoid(
		f,
		"resetUse32BitWorker",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetVnetRouteAllEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetVnetRouteAllEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetWebsocketsEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetWebsocketsEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ResetWorkerCount() {
	_jsii_.InvokeVoid(
		f,
		"resetWorkerCount",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

