package linuxfunctionappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/linuxfunctionappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxFunctionAppSlotSiteConfigOutputReference interface {
	cdktf.ComplexObject
	AlwaysOn() interface{}
	SetAlwaysOn(val interface{})
	AlwaysOnInput() interface{}
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
	ApplicationStack() LinuxFunctionAppSlotSiteConfigApplicationStackOutputReference
	ApplicationStackInput() *LinuxFunctionAppSlotSiteConfigApplicationStack
	AppScaleLimit() *float64
	SetAppScaleLimit(val *float64)
	AppScaleLimitInput() *float64
	AppServiceLogs() LinuxFunctionAppSlotSiteConfigAppServiceLogsOutputReference
	AppServiceLogsInput() *LinuxFunctionAppSlotSiteConfigAppServiceLogs
	AutoSwapSlotName() *string
	SetAutoSwapSlotName(val *string)
	AutoSwapSlotNameInput() *string
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
	Cors() LinuxFunctionAppSlotSiteConfigCorsOutputReference
	CorsInput() *LinuxFunctionAppSlotSiteConfigCors
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
	FtpsState() *string
	SetFtpsState(val *string)
	FtpsStateInput() *string
	HealthCheckEvictionTimeInMin() *float64
	SetHealthCheckEvictionTimeInMin(val *float64)
	HealthCheckEvictionTimeInMinInput() *float64
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	Http2Enabled() interface{}
	SetHttp2Enabled(val interface{})
	Http2EnabledInput() interface{}
	InternalValue() *LinuxFunctionAppSlotSiteConfig
	SetInternalValue(val *LinuxFunctionAppSlotSiteConfig)
	IpRestriction() LinuxFunctionAppSlotSiteConfigIpRestrictionList
	IpRestrictionInput() interface{}
	LinuxFxVersion() *string
	LoadBalancingMode() *string
	SetLoadBalancingMode(val *string)
	LoadBalancingModeInput() *string
	ManagedPipelineMode() *string
	SetManagedPipelineMode(val *string)
	ManagedPipelineModeInput() *string
	MinimumTlsVersion() *string
	SetMinimumTlsVersion(val *string)
	MinimumTlsVersionInput() *string
	PreWarmedInstanceCount() *float64
	SetPreWarmedInstanceCount(val *float64)
	PreWarmedInstanceCountInput() *float64
	RemoteDebuggingEnabled() interface{}
	SetRemoteDebuggingEnabled(val interface{})
	RemoteDebuggingEnabledInput() interface{}
	RemoteDebuggingVersion() *string
	SetRemoteDebuggingVersion(val *string)
	RemoteDebuggingVersionInput() *string
	RuntimeScaleMonitoringEnabled() interface{}
	SetRuntimeScaleMonitoringEnabled(val interface{})
	RuntimeScaleMonitoringEnabledInput() interface{}
	ScmIpRestriction() LinuxFunctionAppSlotSiteConfigScmIpRestrictionList
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
	PutApplicationStack(value *LinuxFunctionAppSlotSiteConfigApplicationStack)
	PutAppServiceLogs(value *LinuxFunctionAppSlotSiteConfigAppServiceLogs)
	PutCors(value *LinuxFunctionAppSlotSiteConfigCors)
	PutIpRestriction(value interface{})
	PutScmIpRestriction(value interface{})
	ResetAlwaysOn()
	ResetApiDefinitionUrl()
	ResetApiManagementApiId()
	ResetAppCommandLine()
	ResetApplicationInsightsConnectionString()
	ResetApplicationInsightsKey()
	ResetApplicationStack()
	ResetAppScaleLimit()
	ResetAppServiceLogs()
	ResetAutoSwapSlotName()
	ResetContainerRegistryManagedIdentityClientId()
	ResetContainerRegistryUseManagedIdentity()
	ResetCors()
	ResetDefaultDocuments()
	ResetElasticInstanceMinimum()
	ResetFtpsState()
	ResetHealthCheckEvictionTimeInMin()
	ResetHealthCheckPath()
	ResetHttp2Enabled()
	ResetIpRestriction()
	ResetLoadBalancingMode()
	ResetManagedPipelineMode()
	ResetMinimumTlsVersion()
	ResetPreWarmedInstanceCount()
	ResetRemoteDebuggingEnabled()
	ResetRemoteDebuggingVersion()
	ResetRuntimeScaleMonitoringEnabled()
	ResetScmIpRestriction()
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

// The jsii proxy struct for LinuxFunctionAppSlotSiteConfigOutputReference
type jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AlwaysOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AlwaysOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApiDefinitionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApiDefinitionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApiManagementApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApiManagementApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AppCommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AppCommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApplicationInsightsConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApplicationInsightsConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApplicationInsightsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApplicationInsightsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApplicationStack() LinuxFunctionAppSlotSiteConfigApplicationStackOutputReference {
	var returns LinuxFunctionAppSlotSiteConfigApplicationStackOutputReference
	_jsii_.Get(
		j,
		"applicationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ApplicationStackInput() *LinuxFunctionAppSlotSiteConfigApplicationStack {
	var returns *LinuxFunctionAppSlotSiteConfigApplicationStack
	_jsii_.Get(
		j,
		"applicationStackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AppScaleLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appScaleLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AppScaleLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appScaleLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AppServiceLogs() LinuxFunctionAppSlotSiteConfigAppServiceLogsOutputReference {
	var returns LinuxFunctionAppSlotSiteConfigAppServiceLogsOutputReference
	_jsii_.Get(
		j,
		"appServiceLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AppServiceLogsInput() *LinuxFunctionAppSlotSiteConfigAppServiceLogs {
	var returns *LinuxFunctionAppSlotSiteConfigAppServiceLogs
	_jsii_.Get(
		j,
		"appServiceLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AutoSwapSlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSwapSlotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) AutoSwapSlotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSwapSlotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ContainerRegistryManagedIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryManagedIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ContainerRegistryManagedIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryManagedIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ContainerRegistryUseManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryUseManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ContainerRegistryUseManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryUseManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) Cors() LinuxFunctionAppSlotSiteConfigCorsOutputReference {
	var returns LinuxFunctionAppSlotSiteConfigCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) CorsInput() *LinuxFunctionAppSlotSiteConfigCors {
	var returns *LinuxFunctionAppSlotSiteConfigCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) DefaultDocuments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) DefaultDocumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) DetailedErrorLoggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"detailedErrorLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ElasticInstanceMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ElasticInstanceMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) FtpsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) FtpsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) HealthCheckEvictionTimeInMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) HealthCheckEvictionTimeInMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) Http2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) Http2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) InternalValue() *LinuxFunctionAppSlotSiteConfig {
	var returns *LinuxFunctionAppSlotSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) IpRestriction() LinuxFunctionAppSlotSiteConfigIpRestrictionList {
	var returns LinuxFunctionAppSlotSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) IpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) LinuxFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) LoadBalancingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) LoadBalancingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ManagedPipelineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ManagedPipelineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) MinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) PreWarmedInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preWarmedInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) PreWarmedInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preWarmedInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) RemoteDebuggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) RemoteDebuggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) RemoteDebuggingVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) RemoteDebuggingVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) RuntimeScaleMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) RuntimeScaleMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ScmIpRestriction() LinuxFunctionAppSlotSiteConfigScmIpRestrictionList {
	var returns LinuxFunctionAppSlotSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ScmIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ScmMinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ScmMinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ScmUseMainIpRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ScmUseMainIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) Use32BitWorker() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) Use32BitWorkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) VnetRouteAllEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) VnetRouteAllEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) WebsocketsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) WebsocketsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) WorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) WorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCountInput",
		&returns,
	)
	return returns
}


func NewLinuxFunctionAppSlotSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxFunctionAppSlotSiteConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxFunctionAppSlotSiteConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlotSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxFunctionAppSlotSiteConfigOutputReference_Override(l LinuxFunctionAppSlotSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlotSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetAlwaysOn(val interface{}) {
	if err := j.validateSetAlwaysOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysOn",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetApiDefinitionUrl(val *string) {
	if err := j.validateSetApiDefinitionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiDefinitionUrl",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetApiManagementApiId(val *string) {
	if err := j.validateSetApiManagementApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementApiId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetAppCommandLine(val *string) {
	if err := j.validateSetAppCommandLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appCommandLine",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetApplicationInsightsConnectionString(val *string) {
	if err := j.validateSetApplicationInsightsConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationInsightsConnectionString",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetApplicationInsightsKey(val *string) {
	if err := j.validateSetApplicationInsightsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationInsightsKey",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetAppScaleLimit(val *float64) {
	if err := j.validateSetAppScaleLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appScaleLimit",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetAutoSwapSlotName(val *string) {
	if err := j.validateSetAutoSwapSlotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSwapSlotName",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetContainerRegistryManagedIdentityClientId(val *string) {
	if err := j.validateSetContainerRegistryManagedIdentityClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryManagedIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetContainerRegistryUseManagedIdentity(val interface{}) {
	if err := j.validateSetContainerRegistryUseManagedIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryUseManagedIdentity",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetDefaultDocuments(val *[]*string) {
	if err := j.validateSetDefaultDocumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDocuments",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetElasticInstanceMinimum(val *float64) {
	if err := j.validateSetElasticInstanceMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticInstanceMinimum",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetFtpsState(val *string) {
	if err := j.validateSetFtpsStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpsState",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetHealthCheckEvictionTimeInMin(val *float64) {
	if err := j.validateSetHealthCheckEvictionTimeInMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckEvictionTimeInMin",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetHttp2Enabled(val interface{}) {
	if err := j.validateSetHttp2EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Enabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetInternalValue(val *LinuxFunctionAppSlotSiteConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetLoadBalancingMode(val *string) {
	if err := j.validateSetLoadBalancingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingMode",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetManagedPipelineMode(val *string) {
	if err := j.validateSetManagedPipelineModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedPipelineMode",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetMinimumTlsVersion(val *string) {
	if err := j.validateSetMinimumTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetPreWarmedInstanceCount(val *float64) {
	if err := j.validateSetPreWarmedInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preWarmedInstanceCount",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetRemoteDebuggingEnabled(val interface{}) {
	if err := j.validateSetRemoteDebuggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDebuggingEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetRemoteDebuggingVersion(val *string) {
	if err := j.validateSetRemoteDebuggingVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDebuggingVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetRuntimeScaleMonitoringEnabled(val interface{}) {
	if err := j.validateSetRuntimeScaleMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeScaleMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetScmMinimumTlsVersion(val *string) {
	if err := j.validateSetScmMinimumTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmMinimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetScmUseMainIpRestriction(val interface{}) {
	if err := j.validateSetScmUseMainIpRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmUseMainIpRestriction",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetUse32BitWorker(val interface{}) {
	if err := j.validateSetUse32BitWorkerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"use32BitWorker",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetVnetRouteAllEnabled(val interface{}) {
	if err := j.validateSetVnetRouteAllEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetRouteAllEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetWebsocketsEnabled(val interface{}) {
	if err := j.validateSetWebsocketsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websocketsEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference)SetWorkerCount(val *float64) {
	if err := j.validateSetWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerCount",
		val,
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) PutApplicationStack(value *LinuxFunctionAppSlotSiteConfigApplicationStack) {
	if err := l.validatePutApplicationStackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putApplicationStack",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) PutAppServiceLogs(value *LinuxFunctionAppSlotSiteConfigAppServiceLogs) {
	if err := l.validatePutAppServiceLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAppServiceLogs",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) PutCors(value *LinuxFunctionAppSlotSiteConfigCors) {
	if err := l.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCors",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) PutIpRestriction(value interface{}) {
	if err := l.validatePutIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIpRestriction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) PutScmIpRestriction(value interface{}) {
	if err := l.validatePutScmIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putScmIpRestriction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetAlwaysOn() {
	_jsii_.InvokeVoid(
		l,
		"resetAlwaysOn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetApiDefinitionUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetApiDefinitionUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetApiManagementApiId() {
	_jsii_.InvokeVoid(
		l,
		"resetApiManagementApiId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetAppCommandLine() {
	_jsii_.InvokeVoid(
		l,
		"resetAppCommandLine",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetApplicationInsightsConnectionString() {
	_jsii_.InvokeVoid(
		l,
		"resetApplicationInsightsConnectionString",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetApplicationInsightsKey() {
	_jsii_.InvokeVoid(
		l,
		"resetApplicationInsightsKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetApplicationStack() {
	_jsii_.InvokeVoid(
		l,
		"resetApplicationStack",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetAppScaleLimit() {
	_jsii_.InvokeVoid(
		l,
		"resetAppScaleLimit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetAppServiceLogs() {
	_jsii_.InvokeVoid(
		l,
		"resetAppServiceLogs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetAutoSwapSlotName() {
	_jsii_.InvokeVoid(
		l,
		"resetAutoSwapSlotName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetContainerRegistryManagedIdentityClientId() {
	_jsii_.InvokeVoid(
		l,
		"resetContainerRegistryManagedIdentityClientId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetContainerRegistryUseManagedIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetContainerRegistryUseManagedIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		l,
		"resetCors",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetDefaultDocuments() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultDocuments",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetElasticInstanceMinimum() {
	_jsii_.InvokeVoid(
		l,
		"resetElasticInstanceMinimum",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetFtpsState() {
	_jsii_.InvokeVoid(
		l,
		"resetFtpsState",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetHealthCheckEvictionTimeInMin() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckEvictionTimeInMin",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetHttp2Enabled() {
	_jsii_.InvokeVoid(
		l,
		"resetHttp2Enabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetLoadBalancingMode() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadBalancingMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetManagedPipelineMode() {
	_jsii_.InvokeVoid(
		l,
		"resetManagedPipelineMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetMinimumTlsVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetPreWarmedInstanceCount() {
	_jsii_.InvokeVoid(
		l,
		"resetPreWarmedInstanceCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetRemoteDebuggingEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetRemoteDebuggingEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetRemoteDebuggingVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetRemoteDebuggingVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetRuntimeScaleMonitoringEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntimeScaleMonitoringEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetScmIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetScmIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetScmMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetScmMinimumTlsVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetScmUseMainIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetScmUseMainIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetUse32BitWorker() {
	_jsii_.InvokeVoid(
		l,
		"resetUse32BitWorker",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetVnetRouteAllEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetVnetRouteAllEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetWebsocketsEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetWebsocketsEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ResetWorkerCount() {
	_jsii_.InvokeVoid(
		l,
		"resetWorkerCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

