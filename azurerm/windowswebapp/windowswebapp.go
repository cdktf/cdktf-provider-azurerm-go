package windowswebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/windowswebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app azurerm_windows_web_app}.
type WindowsWebApp interface {
	cdktf.TerraformResource
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() WindowsWebAppAuthSettingsOutputReference
	AuthSettingsInput() *WindowsWebAppAuthSettings
	Backup() WindowsWebAppBackupOutputReference
	BackupInput() *WindowsWebAppBackup
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientAffinityEnabled() interface{}
	SetClientAffinityEnabled(val interface{})
	ClientAffinityEnabledInput() interface{}
	ClientCertificateEnabled() interface{}
	SetClientCertificateEnabled(val interface{})
	ClientCertificateEnabledInput() interface{}
	ClientCertificateMode() *string
	SetClientCertificateMode(val *string)
	ClientCertificateModeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionString() WindowsWebAppConnectionStringList
	ConnectionStringInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	HttpsOnly() interface{}
	SetHttpsOnly(val interface{})
	HttpsOnlyInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() WindowsWebAppIdentityOutputReference
	IdentityInput() *WindowsWebAppIdentity
	IdInput() *string
	KeyVaultReferenceIdentityId() *string
	SetKeyVaultReferenceIdentityId(val *string)
	KeyVaultReferenceIdentityIdInput() *string
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Logs() WindowsWebAppLogsOutputReference
	LogsInput() *WindowsWebAppLogs
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
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ServicePlanId() *string
	SetServicePlanId(val *string)
	ServicePlanIdInput() *string
	SiteConfig() WindowsWebAppSiteConfigOutputReference
	SiteConfigInput() *WindowsWebAppSiteConfig
	SiteCredential() WindowsWebAppSiteCredentialList
	StickySettings() WindowsWebAppStickySettingsOutputReference
	StickySettingsInput() *WindowsWebAppStickySettings
	StorageAccount() WindowsWebAppStorageAccountList
	StorageAccountInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WindowsWebAppTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
	ZipDeployFile() *string
	SetZipDeployFile(val *string)
	ZipDeployFileInput() *string
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
	PutAuthSettings(value *WindowsWebAppAuthSettings)
	PutBackup(value *WindowsWebAppBackup)
	PutConnectionString(value interface{})
	PutIdentity(value *WindowsWebAppIdentity)
	PutLogs(value *WindowsWebAppLogs)
	PutSiteConfig(value *WindowsWebAppSiteConfig)
	PutStickySettings(value *WindowsWebAppStickySettings)
	PutStorageAccount(value interface{})
	PutTimeouts(value *WindowsWebAppTimeouts)
	ResetAppSettings()
	ResetAuthSettings()
	ResetBackup()
	ResetClientAffinityEnabled()
	ResetClientCertificateEnabled()
	ResetClientCertificateMode()
	ResetConnectionString()
	ResetEnabled()
	ResetHttpsOnly()
	ResetId()
	ResetIdentity()
	ResetKeyVaultReferenceIdentityId()
	ResetLogs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStickySettings()
	ResetStorageAccount()
	ResetTags()
	ResetTimeouts()
	ResetVirtualNetworkSubnetId()
	ResetZipDeployFile()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for WindowsWebApp
type jsiiProxy_WindowsWebApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WindowsWebApp) AppSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) AppSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) AuthSettings() WindowsWebAppAuthSettingsOutputReference {
	var returns WindowsWebAppAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) AuthSettingsInput() *WindowsWebAppAuthSettings {
	var returns *WindowsWebAppAuthSettings
	_jsii_.Get(
		j,
		"authSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Backup() WindowsWebAppBackupOutputReference {
	var returns WindowsWebAppBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) BackupInput() *WindowsWebAppBackup {
	var returns *WindowsWebAppBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ClientAffinityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAffinityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ClientAffinityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAffinityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ClientCertificateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ClientCertificateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ClientCertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ClientCertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ConnectionString() WindowsWebAppConnectionStringList {
	var returns WindowsWebAppConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ConnectionStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Identity() WindowsWebAppIdentityOutputReference {
	var returns WindowsWebAppIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) IdentityInput() *WindowsWebAppIdentity {
	var returns *WindowsWebAppIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) KeyVaultReferenceIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) KeyVaultReferenceIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Logs() WindowsWebAppLogsOutputReference {
	var returns WindowsWebAppLogsOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) LogsInput() *WindowsWebAppLogs {
	var returns *WindowsWebAppLogs
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) OutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) PossibleOutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ServicePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) SiteConfig() WindowsWebAppSiteConfigOutputReference {
	var returns WindowsWebAppSiteConfigOutputReference
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) SiteConfigInput() *WindowsWebAppSiteConfig {
	var returns *WindowsWebAppSiteConfig
	_jsii_.Get(
		j,
		"siteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) SiteCredential() WindowsWebAppSiteCredentialList {
	var returns WindowsWebAppSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) StickySettings() WindowsWebAppStickySettingsOutputReference {
	var returns WindowsWebAppStickySettingsOutputReference
	_jsii_.Get(
		j,
		"stickySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) StickySettingsInput() *WindowsWebAppStickySettings {
	var returns *WindowsWebAppStickySettings
	_jsii_.Get(
		j,
		"stickySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) StorageAccount() WindowsWebAppStorageAccountList {
	var returns WindowsWebAppStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) Timeouts() WindowsWebAppTimeoutsOutputReference {
	var returns WindowsWebAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ZipDeployFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipDeployFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ZipDeployFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipDeployFileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app azurerm_windows_web_app} Resource.
func NewWindowsWebApp(scope constructs.Construct, id *string, config *WindowsWebAppConfig) WindowsWebApp {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebApp{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app azurerm_windows_web_app} Resource.
func NewWindowsWebApp_Override(w WindowsWebApp, scope constructs.Construct, id *string, config *WindowsWebAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetAppSettings(val *map[string]*string) {
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetClientAffinityEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"clientAffinityEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetClientCertificateEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetClientCertificateMode(val *string) {
	_jsii_.Set(
		j,
		"clientCertificateMode",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetHttpsOnly(val interface{}) {
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetKeyVaultReferenceIdentityId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultReferenceIdentityId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetServicePlanId(val *string) {
	_jsii_.Set(
		j,
		"servicePlanId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetVirtualNetworkSubnetId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp) SetZipDeployFile(val *string) {
	_jsii_.Set(
		j,
		"zipDeployFile",
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
func WindowsWebApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WindowsWebApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WindowsWebApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WindowsWebApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutAuthSettings(value *WindowsWebAppAuthSettings) {
	_jsii_.InvokeVoid(
		w,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutBackup(value *WindowsWebAppBackup) {
	_jsii_.InvokeVoid(
		w,
		"putBackup",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutConnectionString(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutIdentity(value *WindowsWebAppIdentity) {
	_jsii_.InvokeVoid(
		w,
		"putIdentity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutLogs(value *WindowsWebAppLogs) {
	_jsii_.InvokeVoid(
		w,
		"putLogs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutSiteConfig(value *WindowsWebAppSiteConfig) {
	_jsii_.InvokeVoid(
		w,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutStickySettings(value *WindowsWebAppStickySettings) {
	_jsii_.InvokeVoid(
		w,
		"putStickySettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutStorageAccount(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutTimeouts(value *WindowsWebAppTimeouts) {
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetAppSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetAppSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetAuthSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetAuthSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetBackup() {
	_jsii_.InvokeVoid(
		w,
		"resetBackup",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetClientAffinityEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetClientAffinityEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetClientCertificateEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetClientCertificateEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetClientCertificateMode() {
	_jsii_.InvokeVoid(
		w,
		"resetClientCertificateMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetConnectionString() {
	_jsii_.InvokeVoid(
		w,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetKeyVaultReferenceIdentityId() {
	_jsii_.InvokeVoid(
		w,
		"resetKeyVaultReferenceIdentityId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetLogs() {
	_jsii_.InvokeVoid(
		w,
		"resetLogs",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetStickySettings() {
	_jsii_.InvokeVoid(
		w,
		"resetStickySettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		w,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) ResetZipDeployFile() {
	_jsii_.InvokeVoid(
		w,
		"resetZipDeployFile",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppAuthSettings struct {
	// Should the Authentication / Authorization feature be enabled?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#enabled WindowsWebApp#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#active_directory WindowsWebApp#active_directory}
	ActiveDirectory *WindowsWebAppAuthSettingsActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Specifies a map of Login Parameters to send to the OpenID Connect authorization endpoint when a user logs in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#additional_login_parameters WindowsWebApp#additional_login_parameters}
	AdditionalLoginParameters *map[string]*string `field:"optional" json:"additionalLoginParameters" yaml:"additionalLoginParameters"`
	// Specifies a list of External URLs that can be redirected to as part of logging in or logging out of the Windows Web App.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#allowed_external_redirect_urls WindowsWebApp#allowed_external_redirect_urls}
	AllowedExternalRedirectUrls *[]*string `field:"optional" json:"allowedExternalRedirectUrls" yaml:"allowedExternalRedirectUrls"`
	// The default authentication provider to use when multiple providers are configured.
	//
	// Possible values include: `AzureActiveDirectory`, `Facebook`, `Google`, `MicrosoftAccount`, `Twitter`, `Github`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#default_provider WindowsWebApp#default_provider}
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// facebook block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#facebook WindowsWebApp#facebook}
	Facebook *WindowsWebAppAuthSettingsFacebook `field:"optional" json:"facebook" yaml:"facebook"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#github WindowsWebApp#github}
	Github *WindowsWebAppAuthSettingsGithub `field:"optional" json:"github" yaml:"github"`
	// google block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#google WindowsWebApp#google}
	Google *WindowsWebAppAuthSettingsGoogle `field:"optional" json:"google" yaml:"google"`
	// The OpenID Connect Issuer URI that represents the entity which issues access tokens.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#issuer WindowsWebApp#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// microsoft block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#microsoft WindowsWebApp#microsoft}
	Microsoft *WindowsWebAppAuthSettingsMicrosoft `field:"optional" json:"microsoft" yaml:"microsoft"`
	// The RuntimeVersion of the Authentication / Authorization feature in use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#runtime_version WindowsWebApp#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// The number of hours after session token expiration that a session token can be used to call the token refresh API.
	//
	// Defaults to `72` hours.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#token_refresh_extension_hours WindowsWebApp#token_refresh_extension_hours}
	TokenRefreshExtensionHours *float64 `field:"optional" json:"tokenRefreshExtensionHours" yaml:"tokenRefreshExtensionHours"`
	// Should the Windows Web App durably store platform-specific security tokens that are obtained during login flows? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#token_store_enabled WindowsWebApp#token_store_enabled}
	TokenStoreEnabled interface{} `field:"optional" json:"tokenStoreEnabled" yaml:"tokenStoreEnabled"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#twitter WindowsWebApp#twitter}
	Twitter *WindowsWebAppAuthSettingsTwitter `field:"optional" json:"twitter" yaml:"twitter"`
	// The action to take when an unauthenticated client attempts to access the app. Possible values include: `RedirectToLoginPage`, `AllowAnonymous`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#unauthenticated_client_action WindowsWebApp#unauthenticated_client_action}
	UnauthenticatedClientAction *string `field:"optional" json:"unauthenticatedClientAction" yaml:"unauthenticatedClientAction"`
}

type WindowsWebAppAuthSettingsActiveDirectory struct {
	// The ID of the Client to use to authenticate with Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_id WindowsWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Specifies a list of Allowed audience values to consider when validating JWTs issued by Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#allowed_audiences WindowsWebApp#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// The Client Secret for the Client ID. Cannot be used with `client_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_secret WindowsWebApp#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The App Setting name that contains the client secret of the Client. Cannot be used with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_secret_setting_name WindowsWebApp#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
}

type WindowsWebAppAuthSettingsActiveDirectoryOutputReference interface {
	cdktf.ComplexObject
	AllowedAudiences() *[]*string
	SetAllowedAudiences(val *[]*string)
	AllowedAudiencesInput() *[]*string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	ClientSecretSettingName() *string
	SetClientSecretSettingName(val *string)
	ClientSecretSettingNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppAuthSettingsActiveDirectory
	SetInternalValue(val *WindowsWebAppAuthSettingsActiveDirectory)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAllowedAudiences()
	ResetClientSecret()
	ResetClientSecretSettingName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppAuthSettingsActiveDirectoryOutputReference
type jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) AllowedAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) AllowedAudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ClientSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ClientSecretSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) InternalValue() *WindowsWebAppAuthSettingsActiveDirectory {
	var returns *WindowsWebAppAuthSettingsActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsActiveDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsActiveDirectoryOutputReference_Override(w WindowsWebAppAuthSettingsActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetAllowedAudiences(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedAudiences",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetClientSecretSettingName(val *string) {
	_jsii_.Set(
		j,
		"clientSecretSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetInternalValue(val *WindowsWebAppAuthSettingsActiveDirectory) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ResetAllowedAudiences() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedAudiences",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ResetClientSecretSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecretSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppAuthSettingsFacebook struct {
	// The App ID of the Facebook app used for login.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#app_id WindowsWebApp#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The App Secret of the Facebook app used for Facebook Login. Cannot be specified with `app_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#app_secret WindowsWebApp#app_secret}
	AppSecret *string `field:"optional" json:"appSecret" yaml:"appSecret"`
	// The app setting name that contains the `app_secret` value used for Facebook Login. Cannot be specified with `app_secret`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#app_secret_setting_name WindowsWebApp#app_secret_setting_name}
	AppSecretSettingName *string `field:"optional" json:"appSecretSettingName" yaml:"appSecretSettingName"`
	// Specifies a list of OAuth 2.0 scopes to be requested as part of Facebook Login authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#oauth_scopes WindowsWebApp#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

type WindowsWebAppAuthSettingsFacebookOutputReference interface {
	cdktf.ComplexObject
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	AppSecret() *string
	SetAppSecret(val *string)
	AppSecretInput() *string
	AppSecretSettingName() *string
	SetAppSecretSettingName(val *string)
	AppSecretSettingNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppAuthSettingsFacebook
	SetInternalValue(val *WindowsWebAppAuthSettingsFacebook)
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAppSecret()
	ResetAppSecretSettingName()
	ResetOauthScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppAuthSettingsFacebookOutputReference
type jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) AppSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) AppSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) AppSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) AppSecretSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSecretSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) InternalValue() *WindowsWebAppAuthSettingsFacebook {
	var returns *WindowsWebAppAuthSettingsFacebook
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsFacebookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsFacebookOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsFacebookOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsFacebookOutputReference_Override(w WindowsWebAppAuthSettingsFacebookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsFacebookOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetAppSecret(val *string) {
	_jsii_.Set(
		j,
		"appSecret",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetAppSecretSettingName(val *string) {
	_jsii_.Set(
		j,
		"appSecretSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetInternalValue(val *WindowsWebAppAuthSettingsFacebook) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetOauthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) ResetAppSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetAppSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) ResetAppSecretSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetAppSecretSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		w,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsFacebookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppAuthSettingsGithub struct {
	// The ID of the GitHub app used for login.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_id WindowsWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Client Secret of the GitHub app used for GitHub Login. Cannot be specified with `client_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_secret WindowsWebApp#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The app setting name that contains the `client_secret` value used for GitHub Login. Cannot be specified with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_secret_setting_name WindowsWebApp#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// Specifies a list of OAuth 2.0 scopes that will be requested as part of GitHub Login authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#oauth_scopes WindowsWebApp#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

type WindowsWebAppAuthSettingsGithubOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	ClientSecretSettingName() *string
	SetClientSecretSettingName(val *string)
	ClientSecretSettingNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppAuthSettingsGithub
	SetInternalValue(val *WindowsWebAppAuthSettingsGithub)
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetClientSecret()
	ResetClientSecretSettingName()
	ResetOauthScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppAuthSettingsGithubOutputReference
type jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ClientSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ClientSecretSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) InternalValue() *WindowsWebAppAuthSettingsGithub {
	var returns *WindowsWebAppAuthSettingsGithub
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsGithubOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsGithubOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsGithubOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsGithubOutputReference_Override(w WindowsWebAppAuthSettingsGithubOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsGithubOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetClientSecretSettingName(val *string) {
	_jsii_.Set(
		j,
		"clientSecretSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetInternalValue(val *WindowsWebAppAuthSettingsGithub) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetOauthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ResetClientSecretSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecretSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		w,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGithubOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppAuthSettingsGoogle struct {
	// The OpenID Connect Client ID for the Google web application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_id WindowsWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret associated with the Google web application.  Cannot be specified with `client_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_secret WindowsWebApp#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The app setting name that contains the `client_secret` value used for Google Login. Cannot be specified with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_secret_setting_name WindowsWebApp#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// Specifies a list of OAuth 2.0 scopes that will be requested as part of Google Sign-In authentication. If not specified, "openid", "profile", and "email" are used as default scopes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#oauth_scopes WindowsWebApp#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

type WindowsWebAppAuthSettingsGoogleOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	ClientSecretSettingName() *string
	SetClientSecretSettingName(val *string)
	ClientSecretSettingNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppAuthSettingsGoogle
	SetInternalValue(val *WindowsWebAppAuthSettingsGoogle)
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetClientSecret()
	ResetClientSecretSettingName()
	ResetOauthScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppAuthSettingsGoogleOutputReference
type jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ClientSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ClientSecretSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) InternalValue() *WindowsWebAppAuthSettingsGoogle {
	var returns *WindowsWebAppAuthSettingsGoogle
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsGoogleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsGoogleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsGoogleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsGoogleOutputReference_Override(w WindowsWebAppAuthSettingsGoogleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsGoogleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetClientSecretSettingName(val *string) {
	_jsii_.Set(
		j,
		"clientSecretSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetInternalValue(val *WindowsWebAppAuthSettingsGoogle) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetOauthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ResetClientSecretSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecretSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		w,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsGoogleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppAuthSettingsMicrosoft struct {
	// The OAuth 2.0 client ID that was created for the app used for authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_id WindowsWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret that was created for the app used for authentication. Cannot be specified with `client_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_secret WindowsWebApp#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The app setting name containing the OAuth 2.0 client secret that was created for the app used for authentication. Cannot be specified with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_secret_setting_name WindowsWebApp#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// The list of OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication. If not specified, `wl.basic` is used as the default scope.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#oauth_scopes WindowsWebApp#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

type WindowsWebAppAuthSettingsMicrosoftOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	ClientSecretSettingName() *string
	SetClientSecretSettingName(val *string)
	ClientSecretSettingNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppAuthSettingsMicrosoft
	SetInternalValue(val *WindowsWebAppAuthSettingsMicrosoft)
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetClientSecret()
	ResetClientSecretSettingName()
	ResetOauthScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppAuthSettingsMicrosoftOutputReference
type jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ClientSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ClientSecretSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) InternalValue() *WindowsWebAppAuthSettingsMicrosoft {
	var returns *WindowsWebAppAuthSettingsMicrosoft
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsMicrosoftOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsMicrosoftOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsMicrosoftOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsMicrosoftOutputReference_Override(w WindowsWebAppAuthSettingsMicrosoftOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsMicrosoftOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetClientSecretSettingName(val *string) {
	_jsii_.Set(
		j,
		"clientSecretSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetInternalValue(val *WindowsWebAppAuthSettingsMicrosoft) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetOauthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ResetClientSecretSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecretSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		w,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsMicrosoftOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppAuthSettingsOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() WindowsWebAppAuthSettingsActiveDirectoryOutputReference
	ActiveDirectoryInput() *WindowsWebAppAuthSettingsActiveDirectory
	AdditionalLoginParameters() *map[string]*string
	SetAdditionalLoginParameters(val *map[string]*string)
	AdditionalLoginParametersInput() *map[string]*string
	AllowedExternalRedirectUrls() *[]*string
	SetAllowedExternalRedirectUrls(val *[]*string)
	AllowedExternalRedirectUrlsInput() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultProvider() *string
	SetDefaultProvider(val *string)
	DefaultProviderInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Facebook() WindowsWebAppAuthSettingsFacebookOutputReference
	FacebookInput() *WindowsWebAppAuthSettingsFacebook
	// Experimental.
	Fqn() *string
	Github() WindowsWebAppAuthSettingsGithubOutputReference
	GithubInput() *WindowsWebAppAuthSettingsGithub
	Google() WindowsWebAppAuthSettingsGoogleOutputReference
	GoogleInput() *WindowsWebAppAuthSettingsGoogle
	InternalValue() *WindowsWebAppAuthSettings
	SetInternalValue(val *WindowsWebAppAuthSettings)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	Microsoft() WindowsWebAppAuthSettingsMicrosoftOutputReference
	MicrosoftInput() *WindowsWebAppAuthSettingsMicrosoft
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenRefreshExtensionHours() *float64
	SetTokenRefreshExtensionHours(val *float64)
	TokenRefreshExtensionHoursInput() *float64
	TokenStoreEnabled() interface{}
	SetTokenStoreEnabled(val interface{})
	TokenStoreEnabledInput() interface{}
	Twitter() WindowsWebAppAuthSettingsTwitterOutputReference
	TwitterInput() *WindowsWebAppAuthSettingsTwitter
	UnauthenticatedClientAction() *string
	SetUnauthenticatedClientAction(val *string)
	UnauthenticatedClientActionInput() *string
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
	PutActiveDirectory(value *WindowsWebAppAuthSettingsActiveDirectory)
	PutFacebook(value *WindowsWebAppAuthSettingsFacebook)
	PutGithub(value *WindowsWebAppAuthSettingsGithub)
	PutGoogle(value *WindowsWebAppAuthSettingsGoogle)
	PutMicrosoft(value *WindowsWebAppAuthSettingsMicrosoft)
	PutTwitter(value *WindowsWebAppAuthSettingsTwitter)
	ResetActiveDirectory()
	ResetAdditionalLoginParameters()
	ResetAllowedExternalRedirectUrls()
	ResetDefaultProvider()
	ResetFacebook()
	ResetGithub()
	ResetGoogle()
	ResetIssuer()
	ResetMicrosoft()
	ResetRuntimeVersion()
	ResetTokenRefreshExtensionHours()
	ResetTokenStoreEnabled()
	ResetTwitter()
	ResetUnauthenticatedClientAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppAuthSettingsOutputReference
type jsiiProxy_WindowsWebAppAuthSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ActiveDirectory() WindowsWebAppAuthSettingsActiveDirectoryOutputReference {
	var returns WindowsWebAppAuthSettingsActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ActiveDirectoryInput() *WindowsWebAppAuthSettingsActiveDirectory {
	var returns *WindowsWebAppAuthSettingsActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) AdditionalLoginParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) AdditionalLoginParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) AllowedExternalRedirectUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) AllowedExternalRedirectUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Facebook() WindowsWebAppAuthSettingsFacebookOutputReference {
	var returns WindowsWebAppAuthSettingsFacebookOutputReference
	_jsii_.Get(
		j,
		"facebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) FacebookInput() *WindowsWebAppAuthSettingsFacebook {
	var returns *WindowsWebAppAuthSettingsFacebook
	_jsii_.Get(
		j,
		"facebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Github() WindowsWebAppAuthSettingsGithubOutputReference {
	var returns WindowsWebAppAuthSettingsGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GithubInput() *WindowsWebAppAuthSettingsGithub {
	var returns *WindowsWebAppAuthSettingsGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Google() WindowsWebAppAuthSettingsGoogleOutputReference {
	var returns WindowsWebAppAuthSettingsGoogleOutputReference
	_jsii_.Get(
		j,
		"google",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GoogleInput() *WindowsWebAppAuthSettingsGoogle {
	var returns *WindowsWebAppAuthSettingsGoogle
	_jsii_.Get(
		j,
		"googleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) InternalValue() *WindowsWebAppAuthSettings {
	var returns *WindowsWebAppAuthSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Microsoft() WindowsWebAppAuthSettingsMicrosoftOutputReference {
	var returns WindowsWebAppAuthSettingsMicrosoftOutputReference
	_jsii_.Get(
		j,
		"microsoft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) MicrosoftInput() *WindowsWebAppAuthSettingsMicrosoft {
	var returns *WindowsWebAppAuthSettingsMicrosoft
	_jsii_.Get(
		j,
		"microsoftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) TokenRefreshExtensionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) TokenRefreshExtensionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) TokenStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) TokenStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Twitter() WindowsWebAppAuthSettingsTwitterOutputReference {
	var returns WindowsWebAppAuthSettingsTwitterOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) TwitterInput() *WindowsWebAppAuthSettingsTwitter {
	var returns *WindowsWebAppAuthSettingsTwitter
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) UnauthenticatedClientAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) UnauthenticatedClientActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientActionInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppAuthSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsOutputReference_Override(w WindowsWebAppAuthSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetAdditionalLoginParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"additionalLoginParameters",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetAllowedExternalRedirectUrls(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedExternalRedirectUrls",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetDefaultProvider(val *string) {
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetInternalValue(val *WindowsWebAppAuthSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetIssuer(val *string) {
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetRuntimeVersion(val *string) {
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetTokenRefreshExtensionHours(val *float64) {
	_jsii_.Set(
		j,
		"tokenRefreshExtensionHours",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetTokenStoreEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tokenStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) SetUnauthenticatedClientAction(val *string) {
	_jsii_.Set(
		j,
		"unauthenticatedClientAction",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) PutActiveDirectory(value *WindowsWebAppAuthSettingsActiveDirectory) {
	_jsii_.InvokeVoid(
		w,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) PutFacebook(value *WindowsWebAppAuthSettingsFacebook) {
	_jsii_.InvokeVoid(
		w,
		"putFacebook",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) PutGithub(value *WindowsWebAppAuthSettingsGithub) {
	_jsii_.InvokeVoid(
		w,
		"putGithub",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) PutGoogle(value *WindowsWebAppAuthSettingsGoogle) {
	_jsii_.InvokeVoid(
		w,
		"putGoogle",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) PutMicrosoft(value *WindowsWebAppAuthSettingsMicrosoft) {
	_jsii_.InvokeVoid(
		w,
		"putMicrosoft",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) PutTwitter(value *WindowsWebAppAuthSettingsTwitter) {
	_jsii_.InvokeVoid(
		w,
		"putTwitter",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		w,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetAdditionalLoginParameters() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalLoginParameters",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetAllowedExternalRedirectUrls() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedExternalRedirectUrls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		w,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetFacebook() {
	_jsii_.InvokeVoid(
		w,
		"resetFacebook",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		w,
		"resetGithub",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetGoogle() {
	_jsii_.InvokeVoid(
		w,
		"resetGoogle",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		w,
		"resetIssuer",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetMicrosoft() {
	_jsii_.InvokeVoid(
		w,
		"resetMicrosoft",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetTokenRefreshExtensionHours() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenRefreshExtensionHours",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetTokenStoreEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenStoreEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		w,
		"resetTwitter",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ResetUnauthenticatedClientAction() {
	_jsii_.InvokeVoid(
		w,
		"resetUnauthenticatedClientAction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppAuthSettingsTwitter struct {
	// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#consumer_key WindowsWebApp#consumer_key}
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// The OAuth 1.0a consumer secret of the Twitter application used for sign-in. Cannot be specified with `consumer_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#consumer_secret WindowsWebApp#consumer_secret}
	ConsumerSecret *string `field:"optional" json:"consumerSecret" yaml:"consumerSecret"`
	// The app setting name that contains the OAuth 1.0a consumer secret of the Twitter application used for sign-in. Cannot be specified with `consumer_secret`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#consumer_secret_setting_name WindowsWebApp#consumer_secret_setting_name}
	ConsumerSecretSettingName *string `field:"optional" json:"consumerSecretSettingName" yaml:"consumerSecretSettingName"`
}

type WindowsWebAppAuthSettingsTwitterOutputReference interface {
	cdktf.ComplexObject
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
	ConsumerKey() *string
	SetConsumerKey(val *string)
	ConsumerKeyInput() *string
	ConsumerSecret() *string
	SetConsumerSecret(val *string)
	ConsumerSecretInput() *string
	ConsumerSecretSettingName() *string
	SetConsumerSecretSettingName(val *string)
	ConsumerSecretSettingNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppAuthSettingsTwitter
	SetInternalValue(val *WindowsWebAppAuthSettingsTwitter)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetConsumerSecret()
	ResetConsumerSecretSettingName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppAuthSettingsTwitterOutputReference
type jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ConsumerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ConsumerKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ConsumerSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ConsumerSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ConsumerSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ConsumerSecretSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerSecretSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) InternalValue() *WindowsWebAppAuthSettingsTwitter {
	var returns *WindowsWebAppAuthSettingsTwitter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsTwitterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsTwitterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsTwitterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsTwitterOutputReference_Override(w WindowsWebAppAuthSettingsTwitterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsTwitterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) SetConsumerKey(val *string) {
	_jsii_.Set(
		j,
		"consumerKey",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) SetConsumerSecret(val *string) {
	_jsii_.Set(
		j,
		"consumerSecret",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) SetConsumerSecretSettingName(val *string) {
	_jsii_.Set(
		j,
		"consumerSecretSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) SetInternalValue(val *WindowsWebAppAuthSettingsTwitter) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ResetConsumerSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetConsumerSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ResetConsumerSecretSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetConsumerSecretSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsTwitterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppBackup struct {
	// The name which should be used for this Backup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#name WindowsWebApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#schedule WindowsWebApp#schedule}
	Schedule *WindowsWebAppBackupSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// The SAS URL to the container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#storage_account_url WindowsWebApp#storage_account_url}
	StorageAccountUrl *string `field:"required" json:"storageAccountUrl" yaml:"storageAccountUrl"`
	// Should this backup job be enabled?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#enabled WindowsWebApp#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

type WindowsWebAppBackupOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppBackup
	SetInternalValue(val *WindowsWebAppBackup)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Schedule() WindowsWebAppBackupScheduleOutputReference
	ScheduleInput() *WindowsWebAppBackupSchedule
	StorageAccountUrl() *string
	SetStorageAccountUrl(val *string)
	StorageAccountUrlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutSchedule(value *WindowsWebAppBackupSchedule)
	ResetEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppBackupOutputReference
type jsiiProxy_WindowsWebAppBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) InternalValue() *WindowsWebAppBackup {
	var returns *WindowsWebAppBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) Schedule() WindowsWebAppBackupScheduleOutputReference {
	var returns WindowsWebAppBackupScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) ScheduleInput() *WindowsWebAppBackupSchedule {
	var returns *WindowsWebAppBackupSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) StorageAccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) StorageAccountUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppBackupOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppBackupOutputReference_Override(w WindowsWebAppBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) SetInternalValue(val *WindowsWebAppBackup) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) SetStorageAccountUrl(val *string) {
	_jsii_.Set(
		j,
		"storageAccountUrl",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) PutSchedule(value *WindowsWebAppBackupSchedule) {
	_jsii_.InvokeVoid(
		w,
		"putSchedule",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppBackupSchedule struct {
	// How often the backup should be executed (e.g. for weekly backup, this should be set to `7` and `frequency_unit` should be set to `Day`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#frequency_interval WindowsWebApp#frequency_interval}
	FrequencyInterval *float64 `field:"required" json:"frequencyInterval" yaml:"frequencyInterval"`
	// The unit of time for how often the backup should take place. Possible values include: `Day` and `Hour`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#frequency_unit WindowsWebApp#frequency_unit}
	FrequencyUnit *string `field:"required" json:"frequencyUnit" yaml:"frequencyUnit"`
	// Should the service keep at least one backup, regardless of age of backup. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#keep_at_least_one_backup WindowsWebApp#keep_at_least_one_backup}
	KeepAtLeastOneBackup interface{} `field:"optional" json:"keepAtLeastOneBackup" yaml:"keepAtLeastOneBackup"`
	// After how many days backups should be deleted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#retention_period_days WindowsWebApp#retention_period_days}
	RetentionPeriodDays *float64 `field:"optional" json:"retentionPeriodDays" yaml:"retentionPeriodDays"`
	// When the schedule should start working in RFC-3339 format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#start_time WindowsWebApp#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

type WindowsWebAppBackupScheduleOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	FrequencyInterval() *float64
	SetFrequencyInterval(val *float64)
	FrequencyIntervalInput() *float64
	FrequencyUnit() *string
	SetFrequencyUnit(val *string)
	FrequencyUnitInput() *string
	InternalValue() *WindowsWebAppBackupSchedule
	SetInternalValue(val *WindowsWebAppBackupSchedule)
	KeepAtLeastOneBackup() interface{}
	SetKeepAtLeastOneBackup(val interface{})
	KeepAtLeastOneBackupInput() interface{}
	LastExecutionTime() *string
	RetentionPeriodDays() *float64
	SetRetentionPeriodDays(val *float64)
	RetentionPeriodDaysInput() *float64
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetKeepAtLeastOneBackup()
	ResetRetentionPeriodDays()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppBackupScheduleOutputReference
type jsiiProxy_WindowsWebAppBackupScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) FrequencyInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) FrequencyIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) FrequencyUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) FrequencyUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) InternalValue() *WindowsWebAppBackupSchedule {
	var returns *WindowsWebAppBackupSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) KeepAtLeastOneBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepAtLeastOneBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) KeepAtLeastOneBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepAtLeastOneBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) LastExecutionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastExecutionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) RetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) RetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppBackupScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppBackupScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppBackupScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppBackupScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppBackupScheduleOutputReference_Override(w WindowsWebAppBackupScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppBackupScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetFrequencyInterval(val *float64) {
	_jsii_.Set(
		j,
		"frequencyInterval",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetFrequencyUnit(val *string) {
	_jsii_.Set(
		j,
		"frequencyUnit",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetInternalValue(val *WindowsWebAppBackupSchedule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetKeepAtLeastOneBackup(val interface{}) {
	_jsii_.Set(
		j,
		"keepAtLeastOneBackup",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetRetentionPeriodDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionPeriodDays",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) ResetKeepAtLeastOneBackup() {
	_jsii_.InvokeVoid(
		w,
		"resetKeepAtLeastOneBackup",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) ResetRetentionPeriodDays() {
	_jsii_.InvokeVoid(
		w,
		"resetRetentionPeriodDays",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		w,
		"resetStartTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppBackupScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#location WindowsWebApp#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#name WindowsWebApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#resource_group_name WindowsWebApp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#service_plan_id WindowsWebApp#service_plan_id}.
	ServicePlanId *string `field:"required" json:"servicePlanId" yaml:"servicePlanId"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#site_config WindowsWebApp#site_config}
	SiteConfig *WindowsWebAppSiteConfig `field:"required" json:"siteConfig" yaml:"siteConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#app_settings WindowsWebApp#app_settings}.
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#auth_settings WindowsWebApp#auth_settings}
	AuthSettings *WindowsWebAppAuthSettings `field:"optional" json:"authSettings" yaml:"authSettings"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#backup WindowsWebApp#backup}
	Backup *WindowsWebAppBackup `field:"optional" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_affinity_enabled WindowsWebApp#client_affinity_enabled}.
	ClientAffinityEnabled interface{} `field:"optional" json:"clientAffinityEnabled" yaml:"clientAffinityEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_certificate_enabled WindowsWebApp#client_certificate_enabled}.
	ClientCertificateEnabled interface{} `field:"optional" json:"clientCertificateEnabled" yaml:"clientCertificateEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#client_certificate_mode WindowsWebApp#client_certificate_mode}.
	ClientCertificateMode *string `field:"optional" json:"clientCertificateMode" yaml:"clientCertificateMode"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#connection_string WindowsWebApp#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#enabled WindowsWebApp#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#https_only WindowsWebApp#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#id WindowsWebApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#identity WindowsWebApp#identity}
	Identity *WindowsWebAppIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#key_vault_reference_identity_id WindowsWebApp#key_vault_reference_identity_id}.
	KeyVaultReferenceIdentityId *string `field:"optional" json:"keyVaultReferenceIdentityId" yaml:"keyVaultReferenceIdentityId"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#logs WindowsWebApp#logs}
	Logs *WindowsWebAppLogs `field:"optional" json:"logs" yaml:"logs"`
	// sticky_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#sticky_settings WindowsWebApp#sticky_settings}
	StickySettings *WindowsWebAppStickySettings `field:"optional" json:"stickySettings" yaml:"stickySettings"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#storage_account WindowsWebApp#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#tags WindowsWebApp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#timeouts WindowsWebApp#timeouts}
	Timeouts *WindowsWebAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#virtual_network_subnet_id WindowsWebApp#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
	// The local path and filename of the Zip packaged application to deploy to this Windows Web App.
	//
	// **Note:** Using this value requires `WEBSITE_RUN_FROM_PACKAGE=1` on the App in `app_settings`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#zip_deploy_file WindowsWebApp#zip_deploy_file}
	ZipDeployFile *string `field:"optional" json:"zipDeployFile" yaml:"zipDeployFile"`
}

type WindowsWebAppConnectionString struct {
	// The name which should be used for this Connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#name WindowsWebApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of database. Possible values include: `MySQL`, `SQLServer`, `SQLAzure`, `Custom`, `NotificationHub`, `ServiceBus`, `EventHub`, `APIHub`, `DocDb`, `RedisCache`, and `PostgreSQL`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#type WindowsWebApp#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The connection string value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#value WindowsWebApp#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

type WindowsWebAppConnectionStringList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppConnectionStringOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppConnectionStringList
type jsiiProxy_WindowsWebAppConnectionStringList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppConnectionStringList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppConnectionStringList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppConnectionStringList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppConnectionStringList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppConnectionStringList_Override(w WindowsWebAppConnectionStringList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppConnectionStringList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppConnectionStringList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringList) Get(index *float64) WindowsWebAppConnectionStringOutputReference {
	var returns WindowsWebAppConnectionStringOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppConnectionStringOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppConnectionStringOutputReference
type jsiiProxy_WindowsWebAppConnectionStringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppConnectionStringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppConnectionStringOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppConnectionStringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppConnectionStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppConnectionStringOutputReference_Override(w WindowsWebAppConnectionStringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppConnectionStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppConnectionStringOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppConnectionStringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#type WindowsWebApp#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#identity_ids WindowsWebApp#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type WindowsWebAppIdentityOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *WindowsWebAppIdentity
	SetInternalValue(val *WindowsWebAppIdentity)
	PrincipalId() *string
	TenantId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetIdentityIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppIdentityOutputReference
type jsiiProxy_WindowsWebAppIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) InternalValue() *WindowsWebAppIdentity {
	var returns *WindowsWebAppIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppIdentityOutputReference_Override(w WindowsWebAppIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) SetInternalValue(val *WindowsWebAppIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppLogs struct {
	// application_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#application_logs WindowsWebApp#application_logs}
	ApplicationLogs *WindowsWebAppLogsApplicationLogs `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#detailed_error_messages WindowsWebApp#detailed_error_messages}.
	DetailedErrorMessages interface{} `field:"optional" json:"detailedErrorMessages" yaml:"detailedErrorMessages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#failed_request_tracing WindowsWebApp#failed_request_tracing}.
	FailedRequestTracing interface{} `field:"optional" json:"failedRequestTracing" yaml:"failedRequestTracing"`
	// http_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#http_logs WindowsWebApp#http_logs}
	HttpLogs *WindowsWebAppLogsHttpLogs `field:"optional" json:"httpLogs" yaml:"httpLogs"`
}

type WindowsWebAppLogsApplicationLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#file_system_level WindowsWebApp#file_system_level}.
	FileSystemLevel *string `field:"required" json:"fileSystemLevel" yaml:"fileSystemLevel"`
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#azure_blob_storage WindowsWebApp#azure_blob_storage}
	AzureBlobStorage *WindowsWebAppLogsApplicationLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
}

type WindowsWebAppLogsApplicationLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#level WindowsWebApp#level}.
	Level *string `field:"required" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#retention_in_days WindowsWebApp#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#sas_url WindowsWebApp#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
}

type WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppLogsApplicationLogsAzureBlobStorage
	SetInternalValue(val *WindowsWebAppLogsApplicationLogsAzureBlobStorage)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	SasUrl() *string
	SetSasUrl(val *string)
	SasUrlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference
type jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) InternalValue() *WindowsWebAppLogsApplicationLogsAzureBlobStorage {
	var returns *WindowsWebAppLogsApplicationLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SasUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SasUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference_Override(w WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SetInternalValue(val *WindowsWebAppLogsApplicationLogsAzureBlobStorage) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SetLevel(val *string) {
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SetRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SetSasUrl(val *string) {
	_jsii_.Set(
		j,
		"sasUrl",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppLogsApplicationLogsOutputReference interface {
	cdktf.ComplexObject
	AzureBlobStorage() WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference
	AzureBlobStorageInput() *WindowsWebAppLogsApplicationLogsAzureBlobStorage
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FileSystemLevel() *string
	SetFileSystemLevel(val *string)
	FileSystemLevelInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppLogsApplicationLogs
	SetInternalValue(val *WindowsWebAppLogsApplicationLogs)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAzureBlobStorage(value *WindowsWebAppLogsApplicationLogsAzureBlobStorage)
	ResetAzureBlobStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppLogsApplicationLogsOutputReference
type jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) AzureBlobStorage() WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference {
	var returns WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference
	_jsii_.Get(
		j,
		"azureBlobStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) AzureBlobStorageInput() *WindowsWebAppLogsApplicationLogsAzureBlobStorage {
	var returns *WindowsWebAppLogsApplicationLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"azureBlobStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) FileSystemLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) FileSystemLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) InternalValue() *WindowsWebAppLogsApplicationLogs {
	var returns *WindowsWebAppLogsApplicationLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppLogsApplicationLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppLogsApplicationLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsApplicationLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppLogsApplicationLogsOutputReference_Override(w WindowsWebAppLogsApplicationLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsApplicationLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) SetFileSystemLevel(val *string) {
	_jsii_.Set(
		j,
		"fileSystemLevel",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) SetInternalValue(val *WindowsWebAppLogsApplicationLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) PutAzureBlobStorage(value *WindowsWebAppLogsApplicationLogsAzureBlobStorage) {
	_jsii_.InvokeVoid(
		w,
		"putAzureBlobStorage",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) ResetAzureBlobStorage() {
	_jsii_.InvokeVoid(
		w,
		"resetAzureBlobStorage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppLogsHttpLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#azure_blob_storage WindowsWebApp#azure_blob_storage}
	AzureBlobStorage *WindowsWebAppLogsHttpLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// file_system block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#file_system WindowsWebApp#file_system}
	FileSystem *WindowsWebAppLogsHttpLogsFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}

type WindowsWebAppLogsHttpLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#sas_url WindowsWebApp#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#retention_in_days WindowsWebApp#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
}

type WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppLogsHttpLogsAzureBlobStorage
	SetInternalValue(val *WindowsWebAppLogsHttpLogsAzureBlobStorage)
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	SasUrl() *string
	SetSasUrl(val *string)
	SasUrlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetRetentionInDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference
type jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) InternalValue() *WindowsWebAppLogsHttpLogsAzureBlobStorage {
	var returns *WindowsWebAppLogsHttpLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SasUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SasUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference_Override(w WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SetInternalValue(val *WindowsWebAppLogsHttpLogsAzureBlobStorage) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SetRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SetSasUrl(val *string) {
	_jsii_.Set(
		j,
		"sasUrl",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		w,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppLogsHttpLogsFileSystem struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#retention_in_days WindowsWebApp#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#retention_in_mb WindowsWebApp#retention_in_mb}.
	RetentionInMb *float64 `field:"required" json:"retentionInMb" yaml:"retentionInMb"`
}

type WindowsWebAppLogsHttpLogsFileSystemOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppLogsHttpLogsFileSystem
	SetInternalValue(val *WindowsWebAppLogsHttpLogsFileSystem)
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	RetentionInMb() *float64
	SetRetentionInMb(val *float64)
	RetentionInMbInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppLogsHttpLogsFileSystemOutputReference
type jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) InternalValue() *WindowsWebAppLogsHttpLogsFileSystem {
	var returns *WindowsWebAppLogsHttpLogsFileSystem
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) RetentionInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) RetentionInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppLogsHttpLogsFileSystemOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppLogsHttpLogsFileSystemOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsHttpLogsFileSystemOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppLogsHttpLogsFileSystemOutputReference_Override(w WindowsWebAppLogsHttpLogsFileSystemOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsHttpLogsFileSystemOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) SetInternalValue(val *WindowsWebAppLogsHttpLogsFileSystem) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) SetRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) SetRetentionInMb(val *float64) {
	_jsii_.Set(
		j,
		"retentionInMb",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsFileSystemOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppLogsHttpLogsOutputReference interface {
	cdktf.ComplexObject
	AzureBlobStorage() WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference
	AzureBlobStorageInput() *WindowsWebAppLogsHttpLogsAzureBlobStorage
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FileSystem() WindowsWebAppLogsHttpLogsFileSystemOutputReference
	FileSystemInput() *WindowsWebAppLogsHttpLogsFileSystem
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppLogsHttpLogs
	SetInternalValue(val *WindowsWebAppLogsHttpLogs)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAzureBlobStorage(value *WindowsWebAppLogsHttpLogsAzureBlobStorage)
	PutFileSystem(value *WindowsWebAppLogsHttpLogsFileSystem)
	ResetAzureBlobStorage()
	ResetFileSystem()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppLogsHttpLogsOutputReference
type jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) AzureBlobStorage() WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference {
	var returns WindowsWebAppLogsHttpLogsAzureBlobStorageOutputReference
	_jsii_.Get(
		j,
		"azureBlobStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) AzureBlobStorageInput() *WindowsWebAppLogsHttpLogsAzureBlobStorage {
	var returns *WindowsWebAppLogsHttpLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"azureBlobStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) FileSystem() WindowsWebAppLogsHttpLogsFileSystemOutputReference {
	var returns WindowsWebAppLogsHttpLogsFileSystemOutputReference
	_jsii_.Get(
		j,
		"fileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) FileSystemInput() *WindowsWebAppLogsHttpLogsFileSystem {
	var returns *WindowsWebAppLogsHttpLogsFileSystem
	_jsii_.Get(
		j,
		"fileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) InternalValue() *WindowsWebAppLogsHttpLogs {
	var returns *WindowsWebAppLogsHttpLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppLogsHttpLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppLogsHttpLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsHttpLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppLogsHttpLogsOutputReference_Override(w WindowsWebAppLogsHttpLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsHttpLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) SetInternalValue(val *WindowsWebAppLogsHttpLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) PutAzureBlobStorage(value *WindowsWebAppLogsHttpLogsAzureBlobStorage) {
	_jsii_.InvokeVoid(
		w,
		"putAzureBlobStorage",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) PutFileSystem(value *WindowsWebAppLogsHttpLogsFileSystem) {
	_jsii_.InvokeVoid(
		w,
		"putFileSystem",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) ResetAzureBlobStorage() {
	_jsii_.InvokeVoid(
		w,
		"resetAzureBlobStorage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) ResetFileSystem() {
	_jsii_.InvokeVoid(
		w,
		"resetFileSystem",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsHttpLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppLogsOutputReference interface {
	cdktf.ComplexObject
	ApplicationLogs() WindowsWebAppLogsApplicationLogsOutputReference
	ApplicationLogsInput() *WindowsWebAppLogsApplicationLogs
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DetailedErrorMessages() interface{}
	SetDetailedErrorMessages(val interface{})
	DetailedErrorMessagesInput() interface{}
	FailedRequestTracing() interface{}
	SetFailedRequestTracing(val interface{})
	FailedRequestTracingInput() interface{}
	// Experimental.
	Fqn() *string
	HttpLogs() WindowsWebAppLogsHttpLogsOutputReference
	HttpLogsInput() *WindowsWebAppLogsHttpLogs
	InternalValue() *WindowsWebAppLogs
	SetInternalValue(val *WindowsWebAppLogs)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutApplicationLogs(value *WindowsWebAppLogsApplicationLogs)
	PutHttpLogs(value *WindowsWebAppLogsHttpLogs)
	ResetApplicationLogs()
	ResetDetailedErrorMessages()
	ResetFailedRequestTracing()
	ResetHttpLogs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppLogsOutputReference
type jsiiProxy_WindowsWebAppLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) ApplicationLogs() WindowsWebAppLogsApplicationLogsOutputReference {
	var returns WindowsWebAppLogsApplicationLogsOutputReference
	_jsii_.Get(
		j,
		"applicationLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) ApplicationLogsInput() *WindowsWebAppLogsApplicationLogs {
	var returns *WindowsWebAppLogsApplicationLogs
	_jsii_.Get(
		j,
		"applicationLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) DetailedErrorMessages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedErrorMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) DetailedErrorMessagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedErrorMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) FailedRequestTracing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedRequestTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) FailedRequestTracingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedRequestTracingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) HttpLogs() WindowsWebAppLogsHttpLogsOutputReference {
	var returns WindowsWebAppLogsHttpLogsOutputReference
	_jsii_.Get(
		j,
		"httpLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) HttpLogsInput() *WindowsWebAppLogsHttpLogs {
	var returns *WindowsWebAppLogsHttpLogs
	_jsii_.Get(
		j,
		"httpLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) InternalValue() *WindowsWebAppLogs {
	var returns *WindowsWebAppLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppLogsOutputReference_Override(w WindowsWebAppLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) SetDetailedErrorMessages(val interface{}) {
	_jsii_.Set(
		j,
		"detailedErrorMessages",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) SetFailedRequestTracing(val interface{}) {
	_jsii_.Set(
		j,
		"failedRequestTracing",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) SetInternalValue(val *WindowsWebAppLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) PutApplicationLogs(value *WindowsWebAppLogsApplicationLogs) {
	_jsii_.InvokeVoid(
		w,
		"putApplicationLogs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) PutHttpLogs(value *WindowsWebAppLogsHttpLogs) {
	_jsii_.InvokeVoid(
		w,
		"putHttpLogs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) ResetApplicationLogs() {
	_jsii_.InvokeVoid(
		w,
		"resetApplicationLogs",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) ResetDetailedErrorMessages() {
	_jsii_.InvokeVoid(
		w,
		"resetDetailedErrorMessages",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) ResetFailedRequestTracing() {
	_jsii_.InvokeVoid(
		w,
		"resetFailedRequestTracing",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) ResetHttpLogs() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpLogs",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#always_on WindowsWebApp#always_on}.
	AlwaysOn interface{} `field:"optional" json:"alwaysOn" yaml:"alwaysOn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#api_definition_url WindowsWebApp#api_definition_url}.
	ApiDefinitionUrl *string `field:"optional" json:"apiDefinitionUrl" yaml:"apiDefinitionUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#api_management_api_id WindowsWebApp#api_management_api_id}.
	ApiManagementApiId *string `field:"optional" json:"apiManagementApiId" yaml:"apiManagementApiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#app_command_line WindowsWebApp#app_command_line}.
	AppCommandLine *string `field:"optional" json:"appCommandLine" yaml:"appCommandLine"`
	// application_stack block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#application_stack WindowsWebApp#application_stack}
	ApplicationStack *WindowsWebAppSiteConfigApplicationStack `field:"optional" json:"applicationStack" yaml:"applicationStack"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#auto_heal_enabled WindowsWebApp#auto_heal_enabled}.
	AutoHealEnabled interface{} `field:"optional" json:"autoHealEnabled" yaml:"autoHealEnabled"`
	// auto_heal_setting block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#auto_heal_setting WindowsWebApp#auto_heal_setting}
	AutoHealSetting *WindowsWebAppSiteConfigAutoHealSetting `field:"optional" json:"autoHealSetting" yaml:"autoHealSetting"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#container_registry_managed_identity_client_id WindowsWebApp#container_registry_managed_identity_client_id}.
	ContainerRegistryManagedIdentityClientId *string `field:"optional" json:"containerRegistryManagedIdentityClientId" yaml:"containerRegistryManagedIdentityClientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#container_registry_use_managed_identity WindowsWebApp#container_registry_use_managed_identity}.
	ContainerRegistryUseManagedIdentity interface{} `field:"optional" json:"containerRegistryUseManagedIdentity" yaml:"containerRegistryUseManagedIdentity"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#cors WindowsWebApp#cors}
	Cors *WindowsWebAppSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#default_documents WindowsWebApp#default_documents}.
	DefaultDocuments *[]*string `field:"optional" json:"defaultDocuments" yaml:"defaultDocuments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#ftps_state WindowsWebApp#ftps_state}.
	FtpsState *string `field:"optional" json:"ftpsState" yaml:"ftpsState"`
	// The amount of time in minutes that a node is unhealthy before being removed from the load balancer.
	//
	// Possible values are between `2` and `10`. Defaults to `10`. Only valid in conjunction with `health_check_path`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#health_check_eviction_time_in_min WindowsWebApp#health_check_eviction_time_in_min}
	HealthCheckEvictionTimeInMin *float64 `field:"optional" json:"healthCheckEvictionTimeInMin" yaml:"healthCheckEvictionTimeInMin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#health_check_path WindowsWebApp#health_check_path}.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#http2_enabled WindowsWebApp#http2_enabled}.
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#ip_restriction WindowsWebApp#ip_restriction}.
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#load_balancing_mode WindowsWebApp#load_balancing_mode}.
	LoadBalancingMode *string `field:"optional" json:"loadBalancingMode" yaml:"loadBalancingMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#local_mysql_enabled WindowsWebApp#local_mysql_enabled}.
	LocalMysqlEnabled interface{} `field:"optional" json:"localMysqlEnabled" yaml:"localMysqlEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#managed_pipeline_mode WindowsWebApp#managed_pipeline_mode}.
	ManagedPipelineMode *string `field:"optional" json:"managedPipelineMode" yaml:"managedPipelineMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#minimum_tls_version WindowsWebApp#minimum_tls_version}.
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#remote_debugging_enabled WindowsWebApp#remote_debugging_enabled}.
	RemoteDebuggingEnabled interface{} `field:"optional" json:"remoteDebuggingEnabled" yaml:"remoteDebuggingEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#remote_debugging_version WindowsWebApp#remote_debugging_version}.
	RemoteDebuggingVersion *string `field:"optional" json:"remoteDebuggingVersion" yaml:"remoteDebuggingVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#scm_ip_restriction WindowsWebApp#scm_ip_restriction}.
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#scm_minimum_tls_version WindowsWebApp#scm_minimum_tls_version}.
	ScmMinimumTlsVersion *string `field:"optional" json:"scmMinimumTlsVersion" yaml:"scmMinimumTlsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#scm_use_main_ip_restriction WindowsWebApp#scm_use_main_ip_restriction}.
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#use_32_bit_worker WindowsWebApp#use_32_bit_worker}.
	Use32BitWorker interface{} `field:"optional" json:"use32BitWorker" yaml:"use32BitWorker"`
	// virtual_application block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#virtual_application WindowsWebApp#virtual_application}
	VirtualApplication interface{} `field:"optional" json:"virtualApplication" yaml:"virtualApplication"`
	// Should all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#vnet_route_all_enabled WindowsWebApp#vnet_route_all_enabled}
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#websockets_enabled WindowsWebApp#websockets_enabled}.
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#worker_count WindowsWebApp#worker_count}.
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
}

type WindowsWebAppSiteConfigApplicationStack struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#current_stack WindowsWebApp#current_stack}.
	CurrentStack *string `field:"optional" json:"currentStack" yaml:"currentStack"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#docker_container_name WindowsWebApp#docker_container_name}.
	DockerContainerName *string `field:"optional" json:"dockerContainerName" yaml:"dockerContainerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#docker_container_registry WindowsWebApp#docker_container_registry}.
	DockerContainerRegistry *string `field:"optional" json:"dockerContainerRegistry" yaml:"dockerContainerRegistry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#docker_container_tag WindowsWebApp#docker_container_tag}.
	DockerContainerTag *string `field:"optional" json:"dockerContainerTag" yaml:"dockerContainerTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#dotnet_version WindowsWebApp#dotnet_version}.
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#java_container WindowsWebApp#java_container}.
	JavaContainer *string `field:"optional" json:"javaContainer" yaml:"javaContainer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#java_container_version WindowsWebApp#java_container_version}.
	JavaContainerVersion *string `field:"optional" json:"javaContainerVersion" yaml:"javaContainerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#java_version WindowsWebApp#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#node_version WindowsWebApp#node_version}.
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#php_version WindowsWebApp#php_version}.
	PhpVersion *string `field:"optional" json:"phpVersion" yaml:"phpVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#python_version WindowsWebApp#python_version}.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
}

type WindowsWebAppSiteConfigApplicationStackOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CurrentStack() *string
	SetCurrentStack(val *string)
	CurrentStackInput() *string
	DockerContainerName() *string
	SetDockerContainerName(val *string)
	DockerContainerNameInput() *string
	DockerContainerRegistry() *string
	SetDockerContainerRegistry(val *string)
	DockerContainerRegistryInput() *string
	DockerContainerTag() *string
	SetDockerContainerTag(val *string)
	DockerContainerTagInput() *string
	DotnetVersion() *string
	SetDotnetVersion(val *string)
	DotnetVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteConfigApplicationStack
	SetInternalValue(val *WindowsWebAppSiteConfigApplicationStack)
	JavaContainer() *string
	SetJavaContainer(val *string)
	JavaContainerInput() *string
	JavaContainerVersion() *string
	SetJavaContainerVersion(val *string)
	JavaContainerVersionInput() *string
	JavaVersion() *string
	SetJavaVersion(val *string)
	JavaVersionInput() *string
	NodeVersion() *string
	SetNodeVersion(val *string)
	NodeVersionInput() *string
	PhpVersion() *string
	SetPhpVersion(val *string)
	PhpVersionInput() *string
	PythonVersion() *string
	SetPythonVersion(val *string)
	PythonVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetCurrentStack()
	ResetDockerContainerName()
	ResetDockerContainerRegistry()
	ResetDockerContainerTag()
	ResetDotnetVersion()
	ResetJavaContainer()
	ResetJavaContainerVersion()
	ResetJavaVersion()
	ResetNodeVersion()
	ResetPhpVersion()
	ResetPythonVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigApplicationStackOutputReference
type jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) CurrentStack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) CurrentStackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentStackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) DockerContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) DockerContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) DockerContainerRegistry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) DockerContainerRegistryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) DockerContainerTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) DockerContainerTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) DotnetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) DotnetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) InternalValue() *WindowsWebAppSiteConfigApplicationStack {
	var returns *WindowsWebAppSiteConfigApplicationStack
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) JavaContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) JavaContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) JavaContainerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) JavaContainerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) JavaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) JavaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) NodeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) PhpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) PhpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) PythonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) PythonVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigApplicationStackOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigApplicationStackOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigApplicationStackOutputReference_Override(w WindowsWebAppSiteConfigApplicationStackOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetCurrentStack(val *string) {
	_jsii_.Set(
		j,
		"currentStack",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetDockerContainerName(val *string) {
	_jsii_.Set(
		j,
		"dockerContainerName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetDockerContainerRegistry(val *string) {
	_jsii_.Set(
		j,
		"dockerContainerRegistry",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetDockerContainerTag(val *string) {
	_jsii_.Set(
		j,
		"dockerContainerTag",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetDotnetVersion(val *string) {
	_jsii_.Set(
		j,
		"dotnetVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetInternalValue(val *WindowsWebAppSiteConfigApplicationStack) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetJavaContainer(val *string) {
	_jsii_.Set(
		j,
		"javaContainer",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetJavaContainerVersion(val *string) {
	_jsii_.Set(
		j,
		"javaContainerVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetJavaVersion(val *string) {
	_jsii_.Set(
		j,
		"javaVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetNodeVersion(val *string) {
	_jsii_.Set(
		j,
		"nodeVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetPhpVersion(val *string) {
	_jsii_.Set(
		j,
		"phpVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetPythonVersion(val *string) {
	_jsii_.Set(
		j,
		"pythonVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetCurrentStack() {
	_jsii_.InvokeVoid(
		w,
		"resetCurrentStack",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetDockerContainerName() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerContainerName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetDockerContainerRegistry() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerContainerRegistry",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetDockerContainerTag() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerContainerTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetDotnetVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetDotnetVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetJavaContainer() {
	_jsii_.InvokeVoid(
		w,
		"resetJavaContainer",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetJavaContainerVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetJavaContainerVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetJavaVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetJavaVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetNodeVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetNodeVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetPhpVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetPhpVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ResetPythonVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetPythonVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigApplicationStackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#action WindowsWebApp#action}
	Action *WindowsWebAppSiteConfigAutoHealSettingAction `field:"required" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#trigger WindowsWebApp#trigger}
	Trigger *WindowsWebAppSiteConfigAutoHealSettingTrigger `field:"required" json:"trigger" yaml:"trigger"`
}

type WindowsWebAppSiteConfigAutoHealSettingAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#action_type WindowsWebApp#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// custom_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#custom_action WindowsWebApp#custom_action}
	CustomAction *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction `field:"optional" json:"customAction" yaml:"customAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#minimum_process_execution_time WindowsWebApp#minimum_process_execution_time}.
	MinimumProcessExecutionTime *string `field:"optional" json:"minimumProcessExecutionTime" yaml:"minimumProcessExecutionTime"`
}

type WindowsWebAppSiteConfigAutoHealSettingActionCustomAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#executable WindowsWebApp#executable}.
	Executable *string `field:"required" json:"executable" yaml:"executable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#parameters WindowsWebApp#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

type WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Executable() *string
	SetExecutable(val *string)
	ExecutableInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction
	SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction)
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference
type jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) Executable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) ExecutableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) InternalValue() *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction {
	var returns *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference_Override(w WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) SetExecutable(val *string) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) SetParameters(val *string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		w,
		"resetParameters",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigAutoHealSettingActionOutputReference interface {
	cdktf.ComplexObject
	ActionType() *string
	SetActionType(val *string)
	ActionTypeInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomAction() WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference
	CustomActionInput() *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteConfigAutoHealSettingAction
	SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingAction)
	MinimumProcessExecutionTime() *string
	SetMinimumProcessExecutionTime(val *string)
	MinimumProcessExecutionTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutCustomAction(value *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction)
	ResetCustomAction()
	ResetMinimumProcessExecutionTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigAutoHealSettingActionOutputReference
type jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) ActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) ActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) CustomAction() WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference {
	var returns WindowsWebAppSiteConfigAutoHealSettingActionCustomActionOutputReference
	_jsii_.Get(
		j,
		"customAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) CustomActionInput() *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction {
	var returns *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction
	_jsii_.Get(
		j,
		"customActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) InternalValue() *WindowsWebAppSiteConfigAutoHealSettingAction {
	var returns *WindowsWebAppSiteConfigAutoHealSettingAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) MinimumProcessExecutionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProcessExecutionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) MinimumProcessExecutionTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProcessExecutionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigAutoHealSettingActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigAutoHealSettingActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigAutoHealSettingActionOutputReference_Override(w WindowsWebAppSiteConfigAutoHealSettingActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) SetActionType(val *string) {
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) SetMinimumProcessExecutionTime(val *string) {
	_jsii_.Set(
		j,
		"minimumProcessExecutionTime",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) PutCustomAction(value *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction) {
	_jsii_.InvokeVoid(
		w,
		"putCustomAction",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) ResetCustomAction() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomAction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) ResetMinimumProcessExecutionTime() {
	_jsii_.InvokeVoid(
		w,
		"resetMinimumProcessExecutionTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigAutoHealSettingOutputReference interface {
	cdktf.ComplexObject
	Action() WindowsWebAppSiteConfigAutoHealSettingActionOutputReference
	ActionInput() *WindowsWebAppSiteConfigAutoHealSettingAction
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteConfigAutoHealSetting
	SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSetting)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trigger() WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference
	TriggerInput() *WindowsWebAppSiteConfigAutoHealSettingTrigger
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
	PutAction(value *WindowsWebAppSiteConfigAutoHealSettingAction)
	PutTrigger(value *WindowsWebAppSiteConfigAutoHealSettingTrigger)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigAutoHealSettingOutputReference
type jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) Action() WindowsWebAppSiteConfigAutoHealSettingActionOutputReference {
	var returns WindowsWebAppSiteConfigAutoHealSettingActionOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) ActionInput() *WindowsWebAppSiteConfigAutoHealSettingAction {
	var returns *WindowsWebAppSiteConfigAutoHealSettingAction
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) InternalValue() *WindowsWebAppSiteConfigAutoHealSetting {
	var returns *WindowsWebAppSiteConfigAutoHealSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) Trigger() WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference {
	var returns WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) TriggerInput() *WindowsWebAppSiteConfigAutoHealSettingTrigger {
	var returns *WindowsWebAppSiteConfigAutoHealSettingTrigger
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigAutoHealSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigAutoHealSettingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigAutoHealSettingOutputReference_Override(w WindowsWebAppSiteConfigAutoHealSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSetting) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) PutAction(value *WindowsWebAppSiteConfigAutoHealSettingAction) {
	_jsii_.InvokeVoid(
		w,
		"putAction",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) PutTrigger(value *WindowsWebAppSiteConfigAutoHealSettingTrigger) {
	_jsii_.InvokeVoid(
		w,
		"putTrigger",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigAutoHealSettingTrigger struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#private_memory_kb WindowsWebApp#private_memory_kb}.
	PrivateMemoryKb *float64 `field:"optional" json:"privateMemoryKb" yaml:"privateMemoryKb"`
	// requests block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#requests WindowsWebApp#requests}
	Requests *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests `field:"optional" json:"requests" yaml:"requests"`
	// slow_request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#slow_request WindowsWebApp#slow_request}
	SlowRequest *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest `field:"optional" json:"slowRequest" yaml:"slowRequest"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#status_code WindowsWebApp#status_code}
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

type WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteConfigAutoHealSettingTrigger
	SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingTrigger)
	PrivateMemoryKb() *float64
	SetPrivateMemoryKb(val *float64)
	PrivateMemoryKbInput() *float64
	Requests() WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference
	RequestsInput() *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests
	SlowRequest() WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference
	SlowRequestInput() *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest
	StatusCode() WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList
	StatusCodeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutRequests(value *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests)
	PutSlowRequest(value *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest)
	PutStatusCode(value interface{})
	ResetPrivateMemoryKb()
	ResetRequests()
	ResetSlowRequest()
	ResetStatusCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference
type jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) InternalValue() *WindowsWebAppSiteConfigAutoHealSettingTrigger {
	var returns *WindowsWebAppSiteConfigAutoHealSettingTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) PrivateMemoryKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateMemoryKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) PrivateMemoryKbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateMemoryKbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) Requests() WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference {
	var returns WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) RequestsInput() *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests {
	var returns *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests
	_jsii_.Get(
		j,
		"requestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) SlowRequest() WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference {
	var returns WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference
	_jsii_.Get(
		j,
		"slowRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) SlowRequestInput() *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest {
	var returns *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest
	_jsii_.Get(
		j,
		"slowRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) StatusCode() WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList {
	var returns WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) StatusCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference_Override(w WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingTrigger) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) SetPrivateMemoryKb(val *float64) {
	_jsii_.Set(
		j,
		"privateMemoryKb",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) PutRequests(value *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests) {
	_jsii_.InvokeVoid(
		w,
		"putRequests",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) PutSlowRequest(value *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest) {
	_jsii_.InvokeVoid(
		w,
		"putSlowRequest",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) PutStatusCode(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putStatusCode",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) ResetPrivateMemoryKb() {
	_jsii_.InvokeVoid(
		w,
		"resetPrivateMemoryKb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) ResetRequests() {
	_jsii_.InvokeVoid(
		w,
		"resetRequests",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) ResetSlowRequest() {
	_jsii_.InvokeVoid(
		w,
		"resetSlowRequest",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) ResetStatusCode() {
	_jsii_.InvokeVoid(
		w,
		"resetStatusCode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigAutoHealSettingTriggerRequests struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#count WindowsWebApp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#interval WindowsWebApp#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
}

type WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference interface {
	cdktf.ComplexObject
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests
	SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests)
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference
type jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) InternalValue() *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests {
	var returns *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference_Override(w WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) SetInterval(val *string) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerRequestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#count WindowsWebApp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#interval WindowsWebApp#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#time_taken WindowsWebApp#time_taken}.
	TimeTaken *string `field:"required" json:"timeTaken" yaml:"timeTaken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#path WindowsWebApp#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

type WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference interface {
	cdktf.ComplexObject
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest
	SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest)
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeTaken() *string
	SetTimeTaken(val *string)
	TimeTakenInput() *string
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
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference
type jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) InternalValue() *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest {
	var returns *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) TimeTaken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeTaken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) TimeTakenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeTakenInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference_Override(w WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetInternalValue(val *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetInterval(val *string) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) SetTimeTaken(val *string) {
	_jsii_.Set(
		j,
		"timeTaken",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		w,
		"resetPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#count WindowsWebApp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#interval WindowsWebApp#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#status_code_range WindowsWebApp#status_code_range}.
	StatusCodeRange *string `field:"required" json:"statusCodeRange" yaml:"statusCodeRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#path WindowsWebApp#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#sub_status WindowsWebApp#sub_status}.
	SubStatus *float64 `field:"optional" json:"subStatus" yaml:"subStatus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#win32_status WindowsWebApp#win32_status}.
	Win32Status *string `field:"optional" json:"win32Status" yaml:"win32Status"`
}

type WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList
type jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList_Override(w WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) Get(index *float64) WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference {
	var returns WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference interface {
	cdktf.ComplexObject
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	StatusCodeRange() *string
	SetStatusCodeRange(val *string)
	StatusCodeRangeInput() *string
	SubStatus() *float64
	SetSubStatus(val *float64)
	SubStatusInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Win32Status() *string
	SetWin32Status(val *string)
	Win32StatusInput() *string
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
	ResetPath()
	ResetSubStatus()
	ResetWin32Status()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference
type jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) StatusCodeRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCodeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) StatusCodeRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCodeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SubStatus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SubStatusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) Win32Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"win32Status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) Win32StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"win32StatusInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference_Override(w WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetInterval(val *string) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetStatusCodeRange(val *string) {
	_jsii_.Set(
		j,
		"statusCodeRange",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetSubStatus(val *float64) {
	_jsii_.Set(
		j,
		"subStatus",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) SetWin32Status(val *string) {
	_jsii_.Set(
		j,
		"win32Status",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		w,
		"resetPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) ResetSubStatus() {
	_jsii_.InvokeVoid(
		w,
		"resetSubStatus",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) ResetWin32Status() {
	_jsii_.InvokeVoid(
		w,
		"resetWin32Status",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigAutoHealSettingTriggerStatusCodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigCors struct {
	// Specifies a list of origins that should be allowed to make cross-origin calls.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#allowed_origins WindowsWebApp#allowed_origins}
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Are credentials allowed in CORS requests? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#support_credentials WindowsWebApp#support_credentials}
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

type WindowsWebAppSiteConfigCorsOutputReference interface {
	cdktf.ComplexObject
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteConfigCors
	SetInternalValue(val *WindowsWebAppSiteConfigCors)
	SupportCredentials() interface{}
	SetSupportCredentials(val interface{})
	SupportCredentialsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetSupportCredentials()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigCorsOutputReference
type jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) InternalValue() *WindowsWebAppSiteConfigCors {
	var returns *WindowsWebAppSiteConfigCors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SupportCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SupportCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigCorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigCorsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigCorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigCorsOutputReference_Override(w WindowsWebAppSiteConfigCorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigCorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SetAllowedOrigins(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SetInternalValue(val *WindowsWebAppSiteConfigCors) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SetSupportCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"supportCredentials",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) ResetSupportCredentials() {
	_jsii_.InvokeVoid(
		w,
		"resetSupportCredentials",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigCorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigIpRestriction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#action WindowsWebApp#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#headers WindowsWebApp#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#ip_address WindowsWebApp#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#name WindowsWebApp#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#priority WindowsWebApp#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#service_tag WindowsWebApp#service_tag}.
	ServiceTag *string `field:"optional" json:"serviceTag" yaml:"serviceTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#virtual_network_subnet_id WindowsWebApp#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
}

type WindowsWebAppSiteConfigIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#x_azure_fdid WindowsWebApp#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#x_fd_health_probe WindowsWebApp#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#x_forwarded_for WindowsWebApp#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#x_forwarded_host WindowsWebApp#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}

type WindowsWebAppSiteConfigIpRestrictionHeadersList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigIpRestrictionHeadersList
type jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigIpRestrictionHeadersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppSiteConfigIpRestrictionHeadersList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigIpRestrictionHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigIpRestrictionHeadersList_Override(w WindowsWebAppSiteConfigIpRestrictionHeadersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigIpRestrictionHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) Get(index *float64) WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference {
	var returns WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XAzureFdid() *[]*string
	SetXAzureFdid(val *[]*string)
	XAzureFdidInput() *[]*string
	XFdHealthProbe() *[]*string
	SetXFdHealthProbe(val *[]*string)
	XFdHealthProbeInput() *[]*string
	XForwardedFor() *[]*string
	SetXForwardedFor(val *[]*string)
	XForwardedForInput() *[]*string
	XForwardedHost() *[]*string
	SetXForwardedHost(val *[]*string)
	XForwardedHostInput() *[]*string
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
	ResetXAzureFdid()
	ResetXFdHealthProbe()
	ResetXForwardedFor()
	ResetXForwardedHost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference
type jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) XAzureFdid() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) XAzureFdidInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) XFdHealthProbe() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) XFdHealthProbeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) XForwardedFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) XForwardedForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedForInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) XForwardedHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) XForwardedHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHostInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigIpRestrictionHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigIpRestrictionHeadersOutputReference_Override(w WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetXAzureFdid(val *[]*string) {
	_jsii_.Set(
		j,
		"xAzureFdid",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetXFdHealthProbe(val *[]*string) {
	_jsii_.Set(
		j,
		"xFdHealthProbe",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetXForwardedFor(val *[]*string) {
	_jsii_.Set(
		j,
		"xForwardedFor",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) SetXForwardedHost(val *[]*string) {
	_jsii_.Set(
		j,
		"xForwardedHost",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) ResetXAzureFdid() {
	_jsii_.InvokeVoid(
		w,
		"resetXAzureFdid",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) ResetXFdHealthProbe() {
	_jsii_.InvokeVoid(
		w,
		"resetXFdHealthProbe",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) ResetXForwardedFor() {
	_jsii_.InvokeVoid(
		w,
		"resetXForwardedFor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) ResetXForwardedHost() {
	_jsii_.InvokeVoid(
		w,
		"resetXForwardedHost",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigIpRestrictionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppSiteConfigIpRestrictionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigIpRestrictionList
type jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigIpRestrictionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppSiteConfigIpRestrictionList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigIpRestrictionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigIpRestrictionList_Override(w WindowsWebAppSiteConfigIpRestrictionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigIpRestrictionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) Get(index *float64) WindowsWebAppSiteConfigIpRestrictionOutputReference {
	var returns WindowsWebAppSiteConfigIpRestrictionOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigIpRestrictionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() WindowsWebAppSiteConfigIpRestrictionHeadersList
	HeadersInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ServiceTag() *string
	SetServiceTag(val *string)
	ServiceTagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
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
	PutHeaders(value interface{})
	ResetAction()
	ResetHeaders()
	ResetIpAddress()
	ResetName()
	ResetPriority()
	ResetServiceTag()
	ResetVirtualNetworkSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigIpRestrictionOutputReference
type jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) Headers() WindowsWebAppSiteConfigIpRestrictionHeadersList {
	var returns WindowsWebAppSiteConfigIpRestrictionHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ServiceTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ServiceTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigIpRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSiteConfigIpRestrictionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigIpRestrictionOutputReference_Override(w WindowsWebAppSiteConfigIpRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetIpAddress(val *string) {
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetServiceTag(val *string) {
	_jsii_.Set(
		j,
		"serviceTag",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) SetVirtualNetworkSubnetId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) PutHeaders(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		w,
		"resetAction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		w,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		w,
		"resetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		w,
		"resetPriority",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ResetServiceTag() {
	_jsii_.InvokeVoid(
		w,
		"resetServiceTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigIpRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigOutputReference interface {
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
	ApplicationStack() WindowsWebAppSiteConfigApplicationStackOutputReference
	ApplicationStackInput() *WindowsWebAppSiteConfigApplicationStack
	AutoHealEnabled() interface{}
	SetAutoHealEnabled(val interface{})
	AutoHealEnabledInput() interface{}
	AutoHealSetting() WindowsWebAppSiteConfigAutoHealSettingOutputReference
	AutoHealSettingInput() *WindowsWebAppSiteConfigAutoHealSetting
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
	Cors() WindowsWebAppSiteConfigCorsOutputReference
	CorsInput() *WindowsWebAppSiteConfigCors
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
	InternalValue() *WindowsWebAppSiteConfig
	SetInternalValue(val *WindowsWebAppSiteConfig)
	IpRestriction() WindowsWebAppSiteConfigIpRestrictionList
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
	ScmIpRestriction() WindowsWebAppSiteConfigScmIpRestrictionList
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
	VirtualApplication() WindowsWebAppSiteConfigVirtualApplicationList
	VirtualApplicationInput() interface{}
	VnetRouteAllEnabled() interface{}
	SetVnetRouteAllEnabled(val interface{})
	VnetRouteAllEnabledInput() interface{}
	WebsocketsEnabled() interface{}
	SetWebsocketsEnabled(val interface{})
	WebsocketsEnabledInput() interface{}
	WindowsFxVersion() *string
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
	PutApplicationStack(value *WindowsWebAppSiteConfigApplicationStack)
	PutAutoHealSetting(value *WindowsWebAppSiteConfigAutoHealSetting)
	PutCors(value *WindowsWebAppSiteConfigCors)
	PutIpRestriction(value interface{})
	PutScmIpRestriction(value interface{})
	PutVirtualApplication(value interface{})
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
	ResetVirtualApplication()
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

// The jsii proxy struct for WindowsWebAppSiteConfigOutputReference
type jsiiProxy_WindowsWebAppSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) AlwaysOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) AlwaysOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ApiDefinitionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ApiDefinitionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ApiManagementApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ApiManagementApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) AppCommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) AppCommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ApplicationStack() WindowsWebAppSiteConfigApplicationStackOutputReference {
	var returns WindowsWebAppSiteConfigApplicationStackOutputReference
	_jsii_.Get(
		j,
		"applicationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ApplicationStackInput() *WindowsWebAppSiteConfigApplicationStack {
	var returns *WindowsWebAppSiteConfigApplicationStack
	_jsii_.Get(
		j,
		"applicationStackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) AutoHealEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) AutoHealEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) AutoHealSetting() WindowsWebAppSiteConfigAutoHealSettingOutputReference {
	var returns WindowsWebAppSiteConfigAutoHealSettingOutputReference
	_jsii_.Get(
		j,
		"autoHealSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) AutoHealSettingInput() *WindowsWebAppSiteConfigAutoHealSetting {
	var returns *WindowsWebAppSiteConfigAutoHealSetting
	_jsii_.Get(
		j,
		"autoHealSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ContainerRegistryManagedIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryManagedIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ContainerRegistryManagedIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryManagedIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ContainerRegistryUseManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryUseManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ContainerRegistryUseManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryUseManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) Cors() WindowsWebAppSiteConfigCorsOutputReference {
	var returns WindowsWebAppSiteConfigCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) CorsInput() *WindowsWebAppSiteConfigCors {
	var returns *WindowsWebAppSiteConfigCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) DefaultDocuments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) DefaultDocumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) DetailedErrorLoggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"detailedErrorLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) FtpsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) FtpsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) HealthCheckEvictionTimeInMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) HealthCheckEvictionTimeInMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) Http2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) Http2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) InternalValue() *WindowsWebAppSiteConfig {
	var returns *WindowsWebAppSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) IpRestriction() WindowsWebAppSiteConfigIpRestrictionList {
	var returns WindowsWebAppSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) IpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) LinuxFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) LoadBalancingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) LoadBalancingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) LocalMysqlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localMysqlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) LocalMysqlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localMysqlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ManagedPipelineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ManagedPipelineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) MinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) RemoteDebuggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) RemoteDebuggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) RemoteDebuggingVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) RemoteDebuggingVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ScmIpRestriction() WindowsWebAppSiteConfigScmIpRestrictionList {
	var returns WindowsWebAppSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ScmIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ScmMinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ScmMinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ScmUseMainIpRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ScmUseMainIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) Use32BitWorker() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) Use32BitWorkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) VirtualApplication() WindowsWebAppSiteConfigVirtualApplicationList {
	var returns WindowsWebAppSiteConfigVirtualApplicationList
	_jsii_.Get(
		j,
		"virtualApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) VirtualApplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) VnetRouteAllEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) VnetRouteAllEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) WebsocketsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) WebsocketsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) WindowsFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) WorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) WorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCountInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSiteConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigOutputReference_Override(w WindowsWebAppSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetAlwaysOn(val interface{}) {
	_jsii_.Set(
		j,
		"alwaysOn",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetApiDefinitionUrl(val *string) {
	_jsii_.Set(
		j,
		"apiDefinitionUrl",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetApiManagementApiId(val *string) {
	_jsii_.Set(
		j,
		"apiManagementApiId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetAppCommandLine(val *string) {
	_jsii_.Set(
		j,
		"appCommandLine",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetAutoHealEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"autoHealEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetContainerRegistryManagedIdentityClientId(val *string) {
	_jsii_.Set(
		j,
		"containerRegistryManagedIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetContainerRegistryUseManagedIdentity(val interface{}) {
	_jsii_.Set(
		j,
		"containerRegistryUseManagedIdentity",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetDefaultDocuments(val *[]*string) {
	_jsii_.Set(
		j,
		"defaultDocuments",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetFtpsState(val *string) {
	_jsii_.Set(
		j,
		"ftpsState",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetHealthCheckEvictionTimeInMin(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckEvictionTimeInMin",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetHealthCheckPath(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetHttp2Enabled(val interface{}) {
	_jsii_.Set(
		j,
		"http2Enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetInternalValue(val *WindowsWebAppSiteConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetLoadBalancingMode(val *string) {
	_jsii_.Set(
		j,
		"loadBalancingMode",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetLocalMysqlEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"localMysqlEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetManagedPipelineMode(val *string) {
	_jsii_.Set(
		j,
		"managedPipelineMode",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetMinimumTlsVersion(val *string) {
	_jsii_.Set(
		j,
		"minimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetRemoteDebuggingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"remoteDebuggingEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetRemoteDebuggingVersion(val *string) {
	_jsii_.Set(
		j,
		"remoteDebuggingVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetScmMinimumTlsVersion(val *string) {
	_jsii_.Set(
		j,
		"scmMinimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetScmUseMainIpRestriction(val interface{}) {
	_jsii_.Set(
		j,
		"scmUseMainIpRestriction",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetUse32BitWorker(val interface{}) {
	_jsii_.Set(
		j,
		"use32BitWorker",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetVnetRouteAllEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"vnetRouteAllEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetWebsocketsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"websocketsEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigOutputReference) SetWorkerCount(val *float64) {
	_jsii_.Set(
		j,
		"workerCount",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) PutApplicationStack(value *WindowsWebAppSiteConfigApplicationStack) {
	_jsii_.InvokeVoid(
		w,
		"putApplicationStack",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) PutAutoHealSetting(value *WindowsWebAppSiteConfigAutoHealSetting) {
	_jsii_.InvokeVoid(
		w,
		"putAutoHealSetting",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) PutCors(value *WindowsWebAppSiteConfigCors) {
	_jsii_.InvokeVoid(
		w,
		"putCors",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) PutIpRestriction(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putIpRestriction",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) PutScmIpRestriction(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putScmIpRestriction",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) PutVirtualApplication(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putVirtualApplication",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetAlwaysOn() {
	_jsii_.InvokeVoid(
		w,
		"resetAlwaysOn",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetApiDefinitionUrl() {
	_jsii_.InvokeVoid(
		w,
		"resetApiDefinitionUrl",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetApiManagementApiId() {
	_jsii_.InvokeVoid(
		w,
		"resetApiManagementApiId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetAppCommandLine() {
	_jsii_.InvokeVoid(
		w,
		"resetAppCommandLine",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetApplicationStack() {
	_jsii_.InvokeVoid(
		w,
		"resetApplicationStack",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetAutoHealEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetAutoHealEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetAutoHealSetting() {
	_jsii_.InvokeVoid(
		w,
		"resetAutoHealSetting",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetContainerRegistryManagedIdentityClientId() {
	_jsii_.InvokeVoid(
		w,
		"resetContainerRegistryManagedIdentityClientId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetContainerRegistryUseManagedIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetContainerRegistryUseManagedIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		w,
		"resetCors",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetDefaultDocuments() {
	_jsii_.InvokeVoid(
		w,
		"resetDefaultDocuments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetFtpsState() {
	_jsii_.InvokeVoid(
		w,
		"resetFtpsState",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetHealthCheckEvictionTimeInMin() {
	_jsii_.InvokeVoid(
		w,
		"resetHealthCheckEvictionTimeInMin",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		w,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetHttp2Enabled() {
	_jsii_.InvokeVoid(
		w,
		"resetHttp2Enabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetIpRestriction() {
	_jsii_.InvokeVoid(
		w,
		"resetIpRestriction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetLoadBalancingMode() {
	_jsii_.InvokeVoid(
		w,
		"resetLoadBalancingMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetLocalMysqlEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetLocalMysqlEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetManagedPipelineMode() {
	_jsii_.InvokeVoid(
		w,
		"resetManagedPipelineMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetMinimumTlsVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetRemoteDebuggingEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetRemoteDebuggingEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetRemoteDebuggingVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetRemoteDebuggingVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetScmIpRestriction() {
	_jsii_.InvokeVoid(
		w,
		"resetScmIpRestriction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetScmMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetScmMinimumTlsVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetScmUseMainIpRestriction() {
	_jsii_.InvokeVoid(
		w,
		"resetScmUseMainIpRestriction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetUse32BitWorker() {
	_jsii_.InvokeVoid(
		w,
		"resetUse32BitWorker",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetVirtualApplication() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualApplication",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetVnetRouteAllEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetVnetRouteAllEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetWebsocketsEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetWebsocketsEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ResetWorkerCount() {
	_jsii_.InvokeVoid(
		w,
		"resetWorkerCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigScmIpRestriction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#action WindowsWebApp#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#headers WindowsWebApp#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#ip_address WindowsWebApp#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#name WindowsWebApp#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#priority WindowsWebApp#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#service_tag WindowsWebApp#service_tag}.
	ServiceTag *string `field:"optional" json:"serviceTag" yaml:"serviceTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#virtual_network_subnet_id WindowsWebApp#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
}

type WindowsWebAppSiteConfigScmIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#x_azure_fdid WindowsWebApp#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#x_fd_health_probe WindowsWebApp#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#x_forwarded_for WindowsWebApp#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#x_forwarded_host WindowsWebApp#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}

type WindowsWebAppSiteConfigScmIpRestrictionHeadersList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigScmIpRestrictionHeadersList
type jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigScmIpRestrictionHeadersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppSiteConfigScmIpRestrictionHeadersList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigScmIpRestrictionHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigScmIpRestrictionHeadersList_Override(w WindowsWebAppSiteConfigScmIpRestrictionHeadersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigScmIpRestrictionHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) Get(index *float64) WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference {
	var returns WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XAzureFdid() *[]*string
	SetXAzureFdid(val *[]*string)
	XAzureFdidInput() *[]*string
	XFdHealthProbe() *[]*string
	SetXFdHealthProbe(val *[]*string)
	XFdHealthProbeInput() *[]*string
	XForwardedFor() *[]*string
	SetXForwardedFor(val *[]*string)
	XForwardedForInput() *[]*string
	XForwardedHost() *[]*string
	SetXForwardedHost(val *[]*string)
	XForwardedHostInput() *[]*string
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
	ResetXAzureFdid()
	ResetXFdHealthProbe()
	ResetXForwardedFor()
	ResetXForwardedHost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference
type jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) XAzureFdid() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) XAzureFdidInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) XFdHealthProbe() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) XFdHealthProbeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedForInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHostInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference_Override(w WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetXAzureFdid(val *[]*string) {
	_jsii_.Set(
		j,
		"xAzureFdid",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetXFdHealthProbe(val *[]*string) {
	_jsii_.Set(
		j,
		"xFdHealthProbe",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetXForwardedFor(val *[]*string) {
	_jsii_.Set(
		j,
		"xForwardedFor",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) SetXForwardedHost(val *[]*string) {
	_jsii_.Set(
		j,
		"xForwardedHost",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) ResetXAzureFdid() {
	_jsii_.InvokeVoid(
		w,
		"resetXAzureFdid",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) ResetXFdHealthProbe() {
	_jsii_.InvokeVoid(
		w,
		"resetXFdHealthProbe",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) ResetXForwardedFor() {
	_jsii_.InvokeVoid(
		w,
		"resetXForwardedFor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) ResetXForwardedHost() {
	_jsii_.InvokeVoid(
		w,
		"resetXForwardedHost",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigScmIpRestrictionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppSiteConfigScmIpRestrictionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigScmIpRestrictionList
type jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigScmIpRestrictionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppSiteConfigScmIpRestrictionList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigScmIpRestrictionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigScmIpRestrictionList_Override(w WindowsWebAppSiteConfigScmIpRestrictionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigScmIpRestrictionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) Get(index *float64) WindowsWebAppSiteConfigScmIpRestrictionOutputReference {
	var returns WindowsWebAppSiteConfigScmIpRestrictionOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigScmIpRestrictionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() WindowsWebAppSiteConfigScmIpRestrictionHeadersList
	HeadersInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ServiceTag() *string
	SetServiceTag(val *string)
	ServiceTagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
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
	PutHeaders(value interface{})
	ResetAction()
	ResetHeaders()
	ResetIpAddress()
	ResetName()
	ResetPriority()
	ResetServiceTag()
	ResetVirtualNetworkSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigScmIpRestrictionOutputReference
type jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) Headers() WindowsWebAppSiteConfigScmIpRestrictionHeadersList {
	var returns WindowsWebAppSiteConfigScmIpRestrictionHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ServiceTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ServiceTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigScmIpRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSiteConfigScmIpRestrictionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigScmIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigScmIpRestrictionOutputReference_Override(w WindowsWebAppSiteConfigScmIpRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigScmIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetIpAddress(val *string) {
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetServiceTag(val *string) {
	_jsii_.Set(
		j,
		"serviceTag",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) SetVirtualNetworkSubnetId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) PutHeaders(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		w,
		"resetAction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		w,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		w,
		"resetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		w,
		"resetPriority",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ResetServiceTag() {
	_jsii_.InvokeVoid(
		w,
		"resetServiceTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigScmIpRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigVirtualApplication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#physical_path WindowsWebApp#physical_path}.
	PhysicalPath *string `field:"required" json:"physicalPath" yaml:"physicalPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#preload WindowsWebApp#preload}.
	Preload interface{} `field:"required" json:"preload" yaml:"preload"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#virtual_path WindowsWebApp#virtual_path}.
	VirtualPath *string `field:"required" json:"virtualPath" yaml:"virtualPath"`
	// virtual_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#virtual_directory WindowsWebApp#virtual_directory}
	VirtualDirectory interface{} `field:"optional" json:"virtualDirectory" yaml:"virtualDirectory"`
}

type WindowsWebAppSiteConfigVirtualApplicationList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppSiteConfigVirtualApplicationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigVirtualApplicationList
type jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigVirtualApplicationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppSiteConfigVirtualApplicationList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigVirtualApplicationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigVirtualApplicationList_Override(w WindowsWebAppSiteConfigVirtualApplicationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigVirtualApplicationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) Get(index *float64) WindowsWebAppSiteConfigVirtualApplicationOutputReference {
	var returns WindowsWebAppSiteConfigVirtualApplicationOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigVirtualApplicationOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PhysicalPath() *string
	SetPhysicalPath(val *string)
	PhysicalPathInput() *string
	Preload() interface{}
	SetPreload(val interface{})
	PreloadInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualDirectory() WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList
	VirtualDirectoryInput() interface{}
	VirtualPath() *string
	SetVirtualPath(val *string)
	VirtualPathInput() *string
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
	PutVirtualDirectory(value interface{})
	ResetVirtualDirectory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigVirtualApplicationOutputReference
type jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) PhysicalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) PhysicalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) Preload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) PreloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) VirtualDirectory() WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList {
	var returns WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList
	_jsii_.Get(
		j,
		"virtualDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) VirtualDirectoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) VirtualPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) VirtualPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualPathInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigVirtualApplicationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSiteConfigVirtualApplicationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigVirtualApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigVirtualApplicationOutputReference_Override(w WindowsWebAppSiteConfigVirtualApplicationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigVirtualApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) SetPhysicalPath(val *string) {
	_jsii_.Set(
		j,
		"physicalPath",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) SetPreload(val interface{}) {
	_jsii_.Set(
		j,
		"preload",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) SetVirtualPath(val *string) {
	_jsii_.Set(
		j,
		"virtualPath",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) PutVirtualDirectory(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putVirtualDirectory",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) ResetVirtualDirectory() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualDirectory",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigVirtualApplicationVirtualDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#physical_path WindowsWebApp#physical_path}.
	PhysicalPath *string `field:"optional" json:"physicalPath" yaml:"physicalPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#virtual_path WindowsWebApp#virtual_path}.
	VirtualPath *string `field:"optional" json:"virtualPath" yaml:"virtualPath"`
}

type WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList
type jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList_Override(w WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) Get(index *float64) WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference {
	var returns WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PhysicalPath() *string
	SetPhysicalPath(val *string)
	PhysicalPathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualPath() *string
	SetVirtualPath(val *string)
	VirtualPathInput() *string
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
	ResetPhysicalPath()
	ResetVirtualPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference
type jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) PhysicalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) PhysicalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) VirtualPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) VirtualPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualPathInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference_Override(w WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) SetPhysicalPath(val *string) {
	_jsii_.Set(
		j,
		"physicalPath",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) SetVirtualPath(val *string) {
	_jsii_.Set(
		j,
		"virtualPath",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) ResetPhysicalPath() {
	_jsii_.InvokeVoid(
		w,
		"resetPhysicalPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) ResetVirtualPath() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteConfigVirtualApplicationVirtualDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteCredential struct {
}

type WindowsWebAppSiteCredentialList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppSiteCredentialOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteCredentialList
type jsiiProxy_WindowsWebAppSiteCredentialList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteCredentialList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppSiteCredentialList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteCredentialList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteCredentialList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteCredentialList_Override(w WindowsWebAppSiteCredentialList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteCredentialList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialList) Get(index *float64) WindowsWebAppSiteCredentialOutputReference {
	var returns WindowsWebAppSiteCredentialOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppSiteCredentialOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSiteCredential
	SetInternalValue(val *WindowsWebAppSiteCredential)
	Name() *string
	Password() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSiteCredentialOutputReference
type jsiiProxy_WindowsWebAppSiteCredentialOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) InternalValue() *WindowsWebAppSiteCredential {
	var returns *WindowsWebAppSiteCredential
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSiteCredentialOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSiteCredentialOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppSiteCredentialOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSiteCredentialOutputReference_Override(w WindowsWebAppSiteCredentialOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppSiteCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) SetInternalValue(val *WindowsWebAppSiteCredential) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSiteCredentialOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppStickySettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#app_setting_names WindowsWebApp#app_setting_names}.
	AppSettingNames *[]*string `field:"optional" json:"appSettingNames" yaml:"appSettingNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#connection_string_names WindowsWebApp#connection_string_names}.
	ConnectionStringNames *[]*string `field:"optional" json:"connectionStringNames" yaml:"connectionStringNames"`
}

type WindowsWebAppStickySettingsOutputReference interface {
	cdktf.ComplexObject
	AppSettingNames() *[]*string
	SetAppSettingNames(val *[]*string)
	AppSettingNamesInput() *[]*string
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
	ConnectionStringNames() *[]*string
	SetConnectionStringNames(val *[]*string)
	ConnectionStringNamesInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppStickySettings
	SetInternalValue(val *WindowsWebAppStickySettings)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAppSettingNames()
	ResetConnectionStringNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppStickySettingsOutputReference
type jsiiProxy_WindowsWebAppStickySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) AppSettingNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appSettingNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) AppSettingNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appSettingNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) ConnectionStringNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionStringNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) ConnectionStringNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionStringNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) InternalValue() *WindowsWebAppStickySettings {
	var returns *WindowsWebAppStickySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppStickySettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppStickySettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppStickySettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppStickySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppStickySettingsOutputReference_Override(w WindowsWebAppStickySettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppStickySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) SetAppSettingNames(val *[]*string) {
	_jsii_.Set(
		j,
		"appSettingNames",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) SetConnectionStringNames(val *[]*string) {
	_jsii_.Set(
		j,
		"connectionStringNames",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) SetInternalValue(val *WindowsWebAppStickySettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStickySettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) ResetAppSettingNames() {
	_jsii_.InvokeVoid(
		w,
		"resetAppSettingNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) ResetConnectionStringNames() {
	_jsii_.InvokeVoid(
		w,
		"resetConnectionStringNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStickySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppStorageAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#access_key WindowsWebApp#access_key}.
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#account_name WindowsWebApp#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#name WindowsWebApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#share_name WindowsWebApp#share_name}.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#type WindowsWebApp#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#mount_path WindowsWebApp#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

type WindowsWebAppStorageAccountList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) WindowsWebAppStorageAccountOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppStorageAccountList
type jsiiProxy_WindowsWebAppStorageAccountList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWindowsWebAppStorageAccountList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WindowsWebAppStorageAccountList {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppStorageAccountList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppStorageAccountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppStorageAccountList_Override(w WindowsWebAppStorageAccountList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppStorageAccountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppStorageAccountList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountList) Get(index *float64) WindowsWebAppStorageAccountOutputReference {
	var returns WindowsWebAppStorageAccountOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppStorageAccountOutputReference interface {
	cdktf.ComplexObject
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountPath() *string
	SetMountPath(val *string)
	MountPathInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	ShareName() *string
	SetShareName(val *string)
	ShareNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetMountPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppStorageAccountOutputReference
type jsiiProxy_WindowsWebAppStorageAccountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) MountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) MountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) ShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) ShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppStorageAccountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppStorageAccountOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppStorageAccountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppStorageAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppStorageAccountOutputReference_Override(w WindowsWebAppStorageAccountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppStorageAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetMountPath(val *string) {
	_jsii_.Set(
		j,
		"mountPath",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetShareName(val *string) {
	_jsii_.Set(
		j,
		"shareName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppStorageAccountOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) ResetMountPath() {
	_jsii_.InvokeVoid(
		w,
		"resetMountPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppStorageAccountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WindowsWebAppTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#create WindowsWebApp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#delete WindowsWebApp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#read WindowsWebApp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#update WindowsWebApp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type WindowsWebAppTimeoutsOutputReference interface {
	cdktf.ComplexObject
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
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppTimeoutsOutputReference
type jsiiProxy_WindowsWebAppTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WindowsWebAppTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppTimeoutsOutputReference_Override(w WindowsWebAppTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		w,
		"resetCreate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		w,
		"resetDelete",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		w,
		"resetRead",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		w,
		"resetUpdate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

