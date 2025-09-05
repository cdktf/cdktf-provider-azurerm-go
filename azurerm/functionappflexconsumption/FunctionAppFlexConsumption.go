// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/functionappflexconsumption/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/function_app_flex_consumption azurerm_function_app_flex_consumption}.
type FunctionAppFlexConsumption interface {
	cdktf.TerraformResource
	AlwaysReady() FunctionAppFlexConsumptionAlwaysReadyList
	AlwaysReadyInput() interface{}
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() FunctionAppFlexConsumptionAuthSettingsOutputReference
	AuthSettingsInput() *FunctionAppFlexConsumptionAuthSettings
	AuthSettingsV2() FunctionAppFlexConsumptionAuthSettingsV2OutputReference
	AuthSettingsV2Input() *FunctionAppFlexConsumptionAuthSettingsV2
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificateEnabled() interface{}
	SetClientCertificateEnabled(val interface{})
	ClientCertificateEnabledInput() interface{}
	ClientCertificateExclusionPaths() *string
	SetClientCertificateExclusionPaths(val *string)
	ClientCertificateExclusionPathsInput() *string
	ClientCertificateMode() *string
	SetClientCertificateMode(val *string)
	ClientCertificateModeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionString() FunctionAppFlexConsumptionConnectionStringList
	ConnectionStringInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomDomainVerificationId() *string
	DefaultHostname() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostingEnvironmentId() *string
	HttpsOnly() interface{}
	SetHttpsOnly(val interface{})
	HttpsOnlyInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() FunctionAppFlexConsumptionIdentityOutputReference
	IdentityInput() *FunctionAppFlexConsumptionIdentity
	IdInput() *string
	InstanceMemoryInMb() *float64
	SetInstanceMemoryInMb(val *float64)
	InstanceMemoryInMbInput() *float64
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaximumInstanceCount() *float64
	SetMaximumInstanceCount(val *float64)
	MaximumInstanceCountInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutboundIpAddresses() *string
	OutboundIpAddressList() *[]*string
	PossibleOutboundIpAddresses() *string
	PossibleOutboundIpAddressList() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RuntimeName() *string
	SetRuntimeName(val *string)
	RuntimeNameInput() *string
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	ServicePlanId() *string
	SetServicePlanId(val *string)
	ServicePlanIdInput() *string
	SiteConfig() FunctionAppFlexConsumptionSiteConfigOutputReference
	SiteConfigInput() *FunctionAppFlexConsumptionSiteConfig
	SiteCredential() FunctionAppFlexConsumptionSiteCredentialList
	StickySettings() FunctionAppFlexConsumptionStickySettingsOutputReference
	StickySettingsInput() *FunctionAppFlexConsumptionStickySettings
	StorageAccessKey() *string
	SetStorageAccessKey(val *string)
	StorageAccessKeyInput() *string
	StorageAuthenticationType() *string
	SetStorageAuthenticationType(val *string)
	StorageAuthenticationTypeInput() *string
	StorageContainerEndpoint() *string
	SetStorageContainerEndpoint(val *string)
	StorageContainerEndpointInput() *string
	StorageContainerType() *string
	SetStorageContainerType(val *string)
	StorageContainerTypeInput() *string
	StorageUserAssignedIdentityId() *string
	SetStorageUserAssignedIdentityId(val *string)
	StorageUserAssignedIdentityIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FunctionAppFlexConsumptionTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
	WebdeployPublishBasicAuthenticationEnabled() interface{}
	SetWebdeployPublishBasicAuthenticationEnabled(val interface{})
	WebdeployPublishBasicAuthenticationEnabledInput() interface{}
	ZipDeployFile() *string
	SetZipDeployFile(val *string)
	ZipDeployFileInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAlwaysReady(value interface{})
	PutAuthSettings(value *FunctionAppFlexConsumptionAuthSettings)
	PutAuthSettingsV2(value *FunctionAppFlexConsumptionAuthSettingsV2)
	PutConnectionString(value interface{})
	PutIdentity(value *FunctionAppFlexConsumptionIdentity)
	PutSiteConfig(value *FunctionAppFlexConsumptionSiteConfig)
	PutStickySettings(value *FunctionAppFlexConsumptionStickySettings)
	PutTimeouts(value *FunctionAppFlexConsumptionTimeouts)
	ResetAlwaysReady()
	ResetAppSettings()
	ResetAuthSettings()
	ResetAuthSettingsV2()
	ResetClientCertificateEnabled()
	ResetClientCertificateExclusionPaths()
	ResetClientCertificateMode()
	ResetConnectionString()
	ResetEnabled()
	ResetHttpsOnly()
	ResetId()
	ResetIdentity()
	ResetInstanceMemoryInMb()
	ResetMaximumInstanceCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetStickySettings()
	ResetStorageAccessKey()
	ResetStorageUserAssignedIdentityId()
	ResetTags()
	ResetTimeouts()
	ResetVirtualNetworkSubnetId()
	ResetWebdeployPublishBasicAuthenticationEnabled()
	ResetZipDeployFile()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FunctionAppFlexConsumption
type jsiiProxy_FunctionAppFlexConsumption struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FunctionAppFlexConsumption) AlwaysReady() FunctionAppFlexConsumptionAlwaysReadyList {
	var returns FunctionAppFlexConsumptionAlwaysReadyList
	_jsii_.Get(
		j,
		"alwaysReady",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) AlwaysReadyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysReadyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) AppSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) AppSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) AuthSettings() FunctionAppFlexConsumptionAuthSettingsOutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) AuthSettingsInput() *FunctionAppFlexConsumptionAuthSettings {
	var returns *FunctionAppFlexConsumptionAuthSettings
	_jsii_.Get(
		j,
		"authSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) AuthSettingsV2() FunctionAppFlexConsumptionAuthSettingsV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2OutputReference
	_jsii_.Get(
		j,
		"authSettingsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) AuthSettingsV2Input() *FunctionAppFlexConsumptionAuthSettingsV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2
	_jsii_.Get(
		j,
		"authSettingsV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ClientCertificateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ClientCertificateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ClientCertificateExclusionPaths() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ClientCertificateExclusionPathsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ClientCertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ClientCertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ConnectionString() FunctionAppFlexConsumptionConnectionStringList {
	var returns FunctionAppFlexConsumptionConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ConnectionStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) HostingEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Identity() FunctionAppFlexConsumptionIdentityOutputReference {
	var returns FunctionAppFlexConsumptionIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) IdentityInput() *FunctionAppFlexConsumptionIdentity {
	var returns *FunctionAppFlexConsumptionIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) InstanceMemoryInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceMemoryInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) InstanceMemoryInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceMemoryInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) MaximumInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) MaximumInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) OutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) PossibleOutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) RuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) RuntimeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ServicePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) SiteConfig() FunctionAppFlexConsumptionSiteConfigOutputReference {
	var returns FunctionAppFlexConsumptionSiteConfigOutputReference
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) SiteConfigInput() *FunctionAppFlexConsumptionSiteConfig {
	var returns *FunctionAppFlexConsumptionSiteConfig
	_jsii_.Get(
		j,
		"siteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) SiteCredential() FunctionAppFlexConsumptionSiteCredentialList {
	var returns FunctionAppFlexConsumptionSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StickySettings() FunctionAppFlexConsumptionStickySettingsOutputReference {
	var returns FunctionAppFlexConsumptionStickySettingsOutputReference
	_jsii_.Get(
		j,
		"stickySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StickySettingsInput() *FunctionAppFlexConsumptionStickySettings {
	var returns *FunctionAppFlexConsumptionStickySettings
	_jsii_.Get(
		j,
		"stickySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageAuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAuthenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageAuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAuthenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageContainerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageContainerEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageContainerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageContainerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageUserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageUserAssignedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) StorageUserAssignedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageUserAssignedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) Timeouts() FunctionAppFlexConsumptionTimeoutsOutputReference {
	var returns FunctionAppFlexConsumptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) WebdeployPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) WebdeployPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ZipDeployFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipDeployFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumption) ZipDeployFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipDeployFileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/function_app_flex_consumption azurerm_function_app_flex_consumption} Resource.
func NewFunctionAppFlexConsumption(scope constructs.Construct, id *string, config *FunctionAppFlexConsumptionConfig) FunctionAppFlexConsumption {
	_init_.Initialize()

	if err := validateNewFunctionAppFlexConsumptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionAppFlexConsumption{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumption",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/function_app_flex_consumption azurerm_function_app_flex_consumption} Resource.
func NewFunctionAppFlexConsumption_Override(f FunctionAppFlexConsumption, scope constructs.Construct, id *string, config *FunctionAppFlexConsumptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumption",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetAppSettings(val *map[string]*string) {
	if err := j.validateSetAppSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetClientCertificateEnabled(val interface{}) {
	if err := j.validateSetClientCertificateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetClientCertificateExclusionPaths(val *string) {
	if err := j.validateSetClientCertificateExclusionPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateExclusionPaths",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetClientCertificateMode(val *string) {
	if err := j.validateSetClientCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateMode",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetHttpsOnly(val interface{}) {
	if err := j.validateSetHttpsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetInstanceMemoryInMb(val *float64) {
	if err := j.validateSetInstanceMemoryInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMemoryInMb",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetMaximumInstanceCount(val *float64) {
	if err := j.validateSetMaximumInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumInstanceCount",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetRuntimeName(val *string) {
	if err := j.validateSetRuntimeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeName",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetServicePlanId(val *string) {
	if err := j.validateSetServicePlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePlanId",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetStorageAccessKey(val *string) {
	if err := j.validateSetStorageAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccessKey",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetStorageAuthenticationType(val *string) {
	if err := j.validateSetStorageAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAuthenticationType",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetStorageContainerEndpoint(val *string) {
	if err := j.validateSetStorageContainerEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageContainerEndpoint",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetStorageContainerType(val *string) {
	if err := j.validateSetStorageContainerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageContainerType",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetStorageUserAssignedIdentityId(val *string) {
	if err := j.validateSetStorageUserAssignedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageUserAssignedIdentityId",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetVirtualNetworkSubnetId(val *string) {
	if err := j.validateSetVirtualNetworkSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetWebdeployPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetWebdeployPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumption)SetZipDeployFile(val *string) {
	if err := j.validateSetZipDeployFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipDeployFile",
		val,
	)
}

// Generates CDKTF code for importing a FunctionAppFlexConsumption resource upon running "cdktf plan <stack-name>".
func FunctionAppFlexConsumption_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFunctionAppFlexConsumption_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumption",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func FunctionAppFlexConsumption_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionAppFlexConsumption_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumption",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionAppFlexConsumption_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionAppFlexConsumption_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumption",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionAppFlexConsumption_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionAppFlexConsumption_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumption",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FunctionAppFlexConsumption_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumption",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumption) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FunctionAppFlexConsumption) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumption) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) PutAlwaysReady(value interface{}) {
	if err := f.validatePutAlwaysReadyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAlwaysReady",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) PutAuthSettings(value *FunctionAppFlexConsumptionAuthSettings) {
	if err := f.validatePutAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) PutAuthSettingsV2(value *FunctionAppFlexConsumptionAuthSettingsV2) {
	if err := f.validatePutAuthSettingsV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAuthSettingsV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) PutConnectionString(value interface{}) {
	if err := f.validatePutConnectionStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) PutIdentity(value *FunctionAppFlexConsumptionIdentity) {
	if err := f.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putIdentity",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) PutSiteConfig(value *FunctionAppFlexConsumptionSiteConfig) {
	if err := f.validatePutSiteConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) PutStickySettings(value *FunctionAppFlexConsumptionStickySettings) {
	if err := f.validatePutStickySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putStickySettings",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) PutTimeouts(value *FunctionAppFlexConsumptionTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetAlwaysReady() {
	_jsii_.InvokeVoid(
		f,
		"resetAlwaysReady",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetAppSettings() {
	_jsii_.InvokeVoid(
		f,
		"resetAppSettings",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetAuthSettings() {
	_jsii_.InvokeVoid(
		f,
		"resetAuthSettings",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetAuthSettingsV2() {
	_jsii_.InvokeVoid(
		f,
		"resetAuthSettingsV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetClientCertificateEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetClientCertificateEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetClientCertificateExclusionPaths() {
	_jsii_.InvokeVoid(
		f,
		"resetClientCertificateExclusionPaths",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetClientCertificateMode() {
	_jsii_.InvokeVoid(
		f,
		"resetClientCertificateMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetConnectionString() {
	_jsii_.InvokeVoid(
		f,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		f,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetIdentity() {
	_jsii_.InvokeVoid(
		f,
		"resetIdentity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetInstanceMemoryInMb() {
	_jsii_.InvokeVoid(
		f,
		"resetInstanceMemoryInMb",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetMaximumInstanceCount() {
	_jsii_.InvokeVoid(
		f,
		"resetMaximumInstanceCount",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetStickySettings() {
	_jsii_.InvokeVoid(
		f,
		"resetStickySettings",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetStorageAccessKey() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageAccessKey",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetStorageUserAssignedIdentityId() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageUserAssignedIdentityId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		f,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetWebdeployPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetWebdeployPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ResetZipDeployFile() {
	_jsii_.InvokeVoid(
		f,
		"resetZipDeployFile",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumption) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumption) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumption) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

