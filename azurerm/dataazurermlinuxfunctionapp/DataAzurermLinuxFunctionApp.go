// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermlinuxfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dataazurermlinuxfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/data-sources/linux_function_app azurerm_linux_function_app}.
type DataAzurermLinuxFunctionApp interface {
	cdktf.TerraformDataSource
	AppSettings() cdktf.StringMap
	AuthSettings() DataAzurermLinuxFunctionAppAuthSettingsList
	AuthSettingsV2() DataAzurermLinuxFunctionAppAuthSettingsV2List
	Availability() *string
	Backup() DataAzurermLinuxFunctionAppBackupList
	BuiltinLoggingEnabled() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificateEnabled() cdktf.IResolvable
	ClientCertificateExclusionPaths() *string
	ClientCertificateMode() *string
	ConnectionString() DataAzurermLinuxFunctionAppConnectionStringList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentShareForceDisabled() cdktf.IResolvable
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomDomainVerificationId() *string
	DailyMemoryTimeQuota() *float64
	DefaultHostname() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FtpPublishBasicAuthenticationEnabled() cdktf.IResolvable
	FunctionsExtensionVersion() *string
	HostingEnvironmentId() *string
	HttpsOnly() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	Identity() DataAzurermLinuxFunctionAppIdentityList
	IdInput() *string
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
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
	PublicNetworkAccessEnabled() cdktf.IResolvable
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ServicePlanId() *string
	SiteConfig() DataAzurermLinuxFunctionAppSiteConfigList
	SiteCredential() DataAzurermLinuxFunctionAppSiteCredentialList
	StickySettings() DataAzurermLinuxFunctionAppStickySettingsList
	StorageAccountAccessKey() *string
	StorageAccountName() *string
	StorageKeyVaultSecretId() *string
	StorageUsesManagedIdentity() cdktf.IResolvable
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermLinuxFunctionAppTimeoutsOutputReference
	TimeoutsInput() interface{}
	Usage() *string
	VirtualNetworkBackupRestoreEnabled() cdktf.IResolvable
	VirtualNetworkSubnetId() *string
	WebdeployPublishBasicAuthenticationEnabled() cdktf.IResolvable
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DataAzurermLinuxFunctionAppTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAzurermLinuxFunctionApp
type jsiiProxy_DataAzurermLinuxFunctionApp struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) AppSettings() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) AuthSettings() DataAzurermLinuxFunctionAppAuthSettingsList {
	var returns DataAzurermLinuxFunctionAppAuthSettingsList
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) AuthSettingsV2() DataAzurermLinuxFunctionAppAuthSettingsV2List {
	var returns DataAzurermLinuxFunctionAppAuthSettingsV2List
	_jsii_.Get(
		j,
		"authSettingsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Availability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Backup() DataAzurermLinuxFunctionAppBackupList {
	var returns DataAzurermLinuxFunctionAppBackupList
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) BuiltinLoggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"builtinLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ClientCertificateEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ClientCertificateExclusionPaths() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ClientCertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ConnectionString() DataAzurermLinuxFunctionAppConnectionStringList {
	var returns DataAzurermLinuxFunctionAppConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ContentShareForceDisabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"contentShareForceDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) DailyMemoryTimeQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) FtpPublishBasicAuthenticationEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) FunctionsExtensionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionsExtensionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) HostingEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) HttpsOnly() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Identity() DataAzurermLinuxFunctionAppIdentityList {
	var returns DataAzurermLinuxFunctionAppIdentityList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) OutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) PossibleOutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) PublicNetworkAccessEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) ServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) SiteConfig() DataAzurermLinuxFunctionAppSiteConfigList {
	var returns DataAzurermLinuxFunctionAppSiteConfigList
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) SiteCredential() DataAzurermLinuxFunctionAppSiteCredentialList {
	var returns DataAzurermLinuxFunctionAppSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) StickySettings() DataAzurermLinuxFunctionAppStickySettingsList {
	var returns DataAzurermLinuxFunctionAppStickySettingsList
	_jsii_.Get(
		j,
		"stickySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) StorageKeyVaultSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyVaultSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) StorageUsesManagedIdentity() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"storageUsesManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Timeouts() DataAzurermLinuxFunctionAppTimeoutsOutputReference {
	var returns DataAzurermLinuxFunctionAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) Usage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) VirtualNetworkBackupRestoreEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"virtualNetworkBackupRestoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp) WebdeployPublishBasicAuthenticationEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/data-sources/linux_function_app azurerm_linux_function_app} Data Source.
func NewDataAzurermLinuxFunctionApp(scope constructs.Construct, id *string, config *DataAzurermLinuxFunctionAppConfig) DataAzurermLinuxFunctionApp {
	_init_.Initialize()

	if err := validateNewDataAzurermLinuxFunctionAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermLinuxFunctionApp{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermLinuxFunctionApp.DataAzurermLinuxFunctionApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/data-sources/linux_function_app azurerm_linux_function_app} Data Source.
func NewDataAzurermLinuxFunctionApp_Override(d DataAzurermLinuxFunctionApp, scope constructs.Construct, id *string, config *DataAzurermLinuxFunctionAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermLinuxFunctionApp.DataAzurermLinuxFunctionApp",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxFunctionApp)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermLinuxFunctionApp resource upon running "cdktf plan <stack-name>".
func DataAzurermLinuxFunctionApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermLinuxFunctionApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermLinuxFunctionApp.DataAzurermLinuxFunctionApp",
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
func DataAzurermLinuxFunctionApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermLinuxFunctionApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermLinuxFunctionApp.DataAzurermLinuxFunctionApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermLinuxFunctionApp_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermLinuxFunctionApp_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermLinuxFunctionApp.DataAzurermLinuxFunctionApp",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermLinuxFunctionApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermLinuxFunctionApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermLinuxFunctionApp.DataAzurermLinuxFunctionApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermLinuxFunctionApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermLinuxFunctionApp.DataAzurermLinuxFunctionApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) PutTimeouts(value *DataAzurermLinuxFunctionAppTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxFunctionApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

