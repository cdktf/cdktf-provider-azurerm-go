package dataazurermwindowswebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/dataazurermwindowswebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/windows_web_app azurerm_windows_web_app}.
type DataAzurermWindowsWebApp interface {
	cdktf.TerraformDataSource
	AppSettings() cdktf.StringMap
	AuthSettings() DataAzurermWindowsWebAppAuthSettingsList
	AuthSettingsV2() DataAzurermWindowsWebAppAuthSettingsV2List
	Backup() DataAzurermWindowsWebAppBackupList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientAffinityEnabled() cdktf.IResolvable
	ClientCertificateEnabled() cdktf.IResolvable
	ClientCertificateExclusionPaths() *string
	ClientCertificateMode() *string
	ConnectionString() DataAzurermWindowsWebAppConnectionStringList
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
	Enabled() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostingEnvironmentId() *string
	HttpsOnly() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	Identity() DataAzurermWindowsWebAppIdentityList
	IdInput() *string
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	Logs() DataAzurermWindowsWebAppLogsList
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
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ServicePlanId() *string
	SiteConfig() DataAzurermWindowsWebAppSiteConfigList
	SiteCredential() DataAzurermWindowsWebAppSiteCredentialList
	StickySettings() DataAzurermWindowsWebAppStickySettingsList
	StorageAccount() DataAzurermWindowsWebAppStorageAccountList
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermWindowsWebAppTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkSubnetId() *string
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
	PutTimeouts(value *DataAzurermWindowsWebAppTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAzurermWindowsWebApp
type jsiiProxy_DataAzurermWindowsWebApp struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) AppSettings() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) AuthSettings() DataAzurermWindowsWebAppAuthSettingsList {
	var returns DataAzurermWindowsWebAppAuthSettingsList
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) AuthSettingsV2() DataAzurermWindowsWebAppAuthSettingsV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2List
	_jsii_.Get(
		j,
		"authSettingsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Backup() DataAzurermWindowsWebAppBackupList {
	var returns DataAzurermWindowsWebAppBackupList
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ClientAffinityEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"clientAffinityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ClientCertificateEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ClientCertificateExclusionPaths() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ClientCertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ConnectionString() DataAzurermWindowsWebAppConnectionStringList {
	var returns DataAzurermWindowsWebAppConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) HostingEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) HttpsOnly() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Identity() DataAzurermWindowsWebAppIdentityList {
	var returns DataAzurermWindowsWebAppIdentityList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Logs() DataAzurermWindowsWebAppLogsList {
	var returns DataAzurermWindowsWebAppLogsList
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) OutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) PossibleOutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) ServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) SiteConfig() DataAzurermWindowsWebAppSiteConfigList {
	var returns DataAzurermWindowsWebAppSiteConfigList
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) SiteCredential() DataAzurermWindowsWebAppSiteCredentialList {
	var returns DataAzurermWindowsWebAppSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) StickySettings() DataAzurermWindowsWebAppStickySettingsList {
	var returns DataAzurermWindowsWebAppStickySettingsList
	_jsii_.Get(
		j,
		"stickySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) StorageAccount() DataAzurermWindowsWebAppStorageAccountList {
	var returns DataAzurermWindowsWebAppStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) Timeouts() DataAzurermWindowsWebAppTimeoutsOutputReference {
	var returns DataAzurermWindowsWebAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebApp) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/windows_web_app azurerm_windows_web_app} Data Source.
func NewDataAzurermWindowsWebApp(scope constructs.Construct, id *string, config *DataAzurermWindowsWebAppConfig) DataAzurermWindowsWebApp {
	_init_.Initialize()

	if err := validateNewDataAzurermWindowsWebAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermWindowsWebApp{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/windows_web_app azurerm_windows_web_app} Data Source.
func NewDataAzurermWindowsWebApp_Override(d DataAzurermWindowsWebApp, scope constructs.Construct, id *string, config *DataAzurermWindowsWebAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebApp",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebApp)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
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
func DataAzurermWindowsWebApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermWindowsWebApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermWindowsWebApp_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermWindowsWebApp_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebApp",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermWindowsWebApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermWindowsWebApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermWindowsWebApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermWindowsWebApp) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) PutTimeouts(value *DataAzurermWindowsWebAppTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

