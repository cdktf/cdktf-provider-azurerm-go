package linuxwebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/linuxwebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxWebAppSiteConfigOutputReference interface {
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
	ApplicationStack() LinuxWebAppSiteConfigApplicationStackOutputReference
	ApplicationStackInput() *LinuxWebAppSiteConfigApplicationStack
	AutoHealEnabled() interface{}
	SetAutoHealEnabled(val interface{})
	AutoHealEnabledInput() interface{}
	AutoHealSetting() LinuxWebAppSiteConfigAutoHealSettingOutputReference
	AutoHealSettingInput() *LinuxWebAppSiteConfigAutoHealSetting
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
	Cors() LinuxWebAppSiteConfigCorsOutputReference
	CorsInput() *LinuxWebAppSiteConfigCors
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultDocuments() *[]*string
	SetDefaultDocuments(val *[]*string)
	DefaultDocumentsInput() *[]*string
	DetailedErrorLoggingEnabled() cdktf.IResolvable
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
	InternalValue() *LinuxWebAppSiteConfig
	SetInternalValue(val *LinuxWebAppSiteConfig)
	IpRestriction() LinuxWebAppSiteConfigIpRestrictionList
	IpRestrictionInput() interface{}
	LinuxFxVersion() *string
	LoadBalancingMode() *string
	SetLoadBalancingMode(val *string)
	LoadBalancingModeInput() *string
	LocalMysqlEnabled() interface{}
	SetLocalMysqlEnabled(val interface{})
	LocalMysqlEnabledInput() interface{}
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
	ScmIpRestriction() LinuxWebAppSiteConfigScmIpRestrictionList
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
	PutApplicationStack(value *LinuxWebAppSiteConfigApplicationStack)
	PutAutoHealSetting(value *LinuxWebAppSiteConfigAutoHealSetting)
	PutCors(value *LinuxWebAppSiteConfigCors)
	PutIpRestriction(value interface{})
	PutScmIpRestriction(value interface{})
	ResetAlwaysOn()
	ResetApiDefinitionUrl()
	ResetApiManagementApiId()
	ResetAppCommandLine()
	ResetApplicationStack()
	ResetAutoHealEnabled()
	ResetAutoHealSetting()
	ResetContainerRegistryManagedIdentityClientId()
	ResetContainerRegistryUseManagedIdentity()
	ResetCors()
	ResetDefaultDocuments()
	ResetFtpsState()
	ResetHealthCheckEvictionTimeInMin()
	ResetHealthCheckPath()
	ResetHttp2Enabled()
	ResetIpRestriction()
	ResetLoadBalancingMode()
	ResetLocalMysqlEnabled()
	ResetManagedPipelineMode()
	ResetMinimumTlsVersion()
	ResetRemoteDebuggingEnabled()
	ResetRemoteDebuggingVersion()
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

// The jsii proxy struct for LinuxWebAppSiteConfigOutputReference
type jsiiProxy_LinuxWebAppSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) AlwaysOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) AlwaysOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ApiDefinitionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ApiDefinitionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ApiManagementApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ApiManagementApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) AppCommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) AppCommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ApplicationStack() LinuxWebAppSiteConfigApplicationStackOutputReference {
	var returns LinuxWebAppSiteConfigApplicationStackOutputReference
	_jsii_.Get(
		j,
		"applicationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ApplicationStackInput() *LinuxWebAppSiteConfigApplicationStack {
	var returns *LinuxWebAppSiteConfigApplicationStack
	_jsii_.Get(
		j,
		"applicationStackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) AutoHealEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) AutoHealEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) AutoHealSetting() LinuxWebAppSiteConfigAutoHealSettingOutputReference {
	var returns LinuxWebAppSiteConfigAutoHealSettingOutputReference
	_jsii_.Get(
		j,
		"autoHealSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) AutoHealSettingInput() *LinuxWebAppSiteConfigAutoHealSetting {
	var returns *LinuxWebAppSiteConfigAutoHealSetting
	_jsii_.Get(
		j,
		"autoHealSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ContainerRegistryManagedIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryManagedIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ContainerRegistryManagedIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryManagedIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ContainerRegistryUseManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryUseManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ContainerRegistryUseManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryUseManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) Cors() LinuxWebAppSiteConfigCorsOutputReference {
	var returns LinuxWebAppSiteConfigCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) CorsInput() *LinuxWebAppSiteConfigCors {
	var returns *LinuxWebAppSiteConfigCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) DefaultDocuments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) DefaultDocumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) DetailedErrorLoggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"detailedErrorLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) FtpsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) FtpsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) HealthCheckEvictionTimeInMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) HealthCheckEvictionTimeInMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) Http2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) Http2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) InternalValue() *LinuxWebAppSiteConfig {
	var returns *LinuxWebAppSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) IpRestriction() LinuxWebAppSiteConfigIpRestrictionList {
	var returns LinuxWebAppSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) IpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) LinuxFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) LoadBalancingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) LoadBalancingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) LocalMysqlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localMysqlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) LocalMysqlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localMysqlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ManagedPipelineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ManagedPipelineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) MinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) RemoteDebuggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) RemoteDebuggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) RemoteDebuggingVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) RemoteDebuggingVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ScmIpRestriction() LinuxWebAppSiteConfigScmIpRestrictionList {
	var returns LinuxWebAppSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ScmIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ScmMinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ScmMinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ScmUseMainIpRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ScmUseMainIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) Use32BitWorker() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) Use32BitWorkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) VnetRouteAllEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) VnetRouteAllEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) WebsocketsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) WebsocketsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) WorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference) WorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCountInput",
		&returns,
	)
	return returns
}


func NewLinuxWebAppSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxWebAppSiteConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxWebAppSiteConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxWebAppSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebApp.LinuxWebAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxWebAppSiteConfigOutputReference_Override(l LinuxWebAppSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebApp.LinuxWebAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetAlwaysOn(val interface{}) {
	if err := j.validateSetAlwaysOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysOn",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetApiDefinitionUrl(val *string) {
	if err := j.validateSetApiDefinitionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiDefinitionUrl",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetApiManagementApiId(val *string) {
	if err := j.validateSetApiManagementApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementApiId",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetAppCommandLine(val *string) {
	if err := j.validateSetAppCommandLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appCommandLine",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetAutoHealEnabled(val interface{}) {
	if err := j.validateSetAutoHealEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHealEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetContainerRegistryManagedIdentityClientId(val *string) {
	if err := j.validateSetContainerRegistryManagedIdentityClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryManagedIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetContainerRegistryUseManagedIdentity(val interface{}) {
	if err := j.validateSetContainerRegistryUseManagedIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryUseManagedIdentity",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetDefaultDocuments(val *[]*string) {
	if err := j.validateSetDefaultDocumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDocuments",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetFtpsState(val *string) {
	if err := j.validateSetFtpsStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpsState",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetHealthCheckEvictionTimeInMin(val *float64) {
	if err := j.validateSetHealthCheckEvictionTimeInMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckEvictionTimeInMin",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetHttp2Enabled(val interface{}) {
	if err := j.validateSetHttp2EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Enabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetInternalValue(val *LinuxWebAppSiteConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetLoadBalancingMode(val *string) {
	if err := j.validateSetLoadBalancingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingMode",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetLocalMysqlEnabled(val interface{}) {
	if err := j.validateSetLocalMysqlEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localMysqlEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetManagedPipelineMode(val *string) {
	if err := j.validateSetManagedPipelineModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedPipelineMode",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetMinimumTlsVersion(val *string) {
	if err := j.validateSetMinimumTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetRemoteDebuggingEnabled(val interface{}) {
	if err := j.validateSetRemoteDebuggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDebuggingEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetRemoteDebuggingVersion(val *string) {
	if err := j.validateSetRemoteDebuggingVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDebuggingVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetScmMinimumTlsVersion(val *string) {
	if err := j.validateSetScmMinimumTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmMinimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetScmUseMainIpRestriction(val interface{}) {
	if err := j.validateSetScmUseMainIpRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmUseMainIpRestriction",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetUse32BitWorker(val interface{}) {
	if err := j.validateSetUse32BitWorkerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"use32BitWorker",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetVnetRouteAllEnabled(val interface{}) {
	if err := j.validateSetVnetRouteAllEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetRouteAllEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetWebsocketsEnabled(val interface{}) {
	if err := j.validateSetWebsocketsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websocketsEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigOutputReference)SetWorkerCount(val *float64) {
	if err := j.validateSetWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerCount",
		val,
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) PutApplicationStack(value *LinuxWebAppSiteConfigApplicationStack) {
	if err := l.validatePutApplicationStackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putApplicationStack",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) PutAutoHealSetting(value *LinuxWebAppSiteConfigAutoHealSetting) {
	if err := l.validatePutAutoHealSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAutoHealSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) PutCors(value *LinuxWebAppSiteConfigCors) {
	if err := l.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCors",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) PutIpRestriction(value interface{}) {
	if err := l.validatePutIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIpRestriction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) PutScmIpRestriction(value interface{}) {
	if err := l.validatePutScmIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putScmIpRestriction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetAlwaysOn() {
	_jsii_.InvokeVoid(
		l,
		"resetAlwaysOn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetApiDefinitionUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetApiDefinitionUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetApiManagementApiId() {
	_jsii_.InvokeVoid(
		l,
		"resetApiManagementApiId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetAppCommandLine() {
	_jsii_.InvokeVoid(
		l,
		"resetAppCommandLine",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetApplicationStack() {
	_jsii_.InvokeVoid(
		l,
		"resetApplicationStack",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetAutoHealEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetAutoHealEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetAutoHealSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetAutoHealSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetContainerRegistryManagedIdentityClientId() {
	_jsii_.InvokeVoid(
		l,
		"resetContainerRegistryManagedIdentityClientId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetContainerRegistryUseManagedIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetContainerRegistryUseManagedIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		l,
		"resetCors",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetDefaultDocuments() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultDocuments",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetFtpsState() {
	_jsii_.InvokeVoid(
		l,
		"resetFtpsState",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetHealthCheckEvictionTimeInMin() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckEvictionTimeInMin",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetHttp2Enabled() {
	_jsii_.InvokeVoid(
		l,
		"resetHttp2Enabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetLoadBalancingMode() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadBalancingMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetLocalMysqlEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetLocalMysqlEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetManagedPipelineMode() {
	_jsii_.InvokeVoid(
		l,
		"resetManagedPipelineMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetMinimumTlsVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetRemoteDebuggingEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetRemoteDebuggingEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetRemoteDebuggingVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetRemoteDebuggingVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetScmIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetScmIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetScmMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetScmMinimumTlsVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetScmUseMainIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetScmUseMainIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetUse32BitWorker() {
	_jsii_.InvokeVoid(
		l,
		"resetUse32BitWorker",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetVnetRouteAllEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetVnetRouteAllEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetWebsocketsEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetWebsocketsEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ResetWorkerCount() {
	_jsii_.InvokeVoid(
		l,
		"resetWorkerCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LinuxWebAppSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

