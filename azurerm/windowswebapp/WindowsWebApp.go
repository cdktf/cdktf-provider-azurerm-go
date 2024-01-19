// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/windowswebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/windows_web_app azurerm_windows_web_app}.
type WindowsWebApp interface {
	cdktf.TerraformResource
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() WindowsWebAppAuthSettingsOutputReference
	AuthSettingsInput() *WindowsWebAppAuthSettings
	AuthSettingsV2() WindowsWebAppAuthSettingsV2OutputReference
	AuthSettingsV2Input() *WindowsWebAppAuthSettingsV2
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
	ConnectionString() WindowsWebAppConnectionStringList
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
	FtpPublishBasicAuthenticationEnabled() interface{}
	SetFtpPublishBasicAuthenticationEnabled(val interface{})
	FtpPublishBasicAuthenticationEnabledInput() interface{}
	HostingEnvironmentId() *string
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
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
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
	PutAuthSettings(value *WindowsWebAppAuthSettings)
	PutAuthSettingsV2(value *WindowsWebAppAuthSettingsV2)
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
	ResetAuthSettingsV2()
	ResetBackup()
	ResetClientAffinityEnabled()
	ResetClientCertificateEnabled()
	ResetClientCertificateExclusionPaths()
	ResetClientCertificateMode()
	ResetConnectionString()
	ResetEnabled()
	ResetFtpPublishBasicAuthenticationEnabled()
	ResetHttpsOnly()
	ResetId()
	ResetIdentity()
	ResetKeyVaultReferenceIdentityId()
	ResetLogs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetStickySettings()
	ResetStorageAccount()
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

func (j *jsiiProxy_WindowsWebApp) AuthSettingsV2() WindowsWebAppAuthSettingsV2OutputReference {
	var returns WindowsWebAppAuthSettingsV2OutputReference
	_jsii_.Get(
		j,
		"authSettingsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) AuthSettingsV2Input() *WindowsWebAppAuthSettingsV2 {
	var returns *WindowsWebAppAuthSettingsV2
	_jsii_.Get(
		j,
		"authSettingsV2Input",
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

func (j *jsiiProxy_WindowsWebApp) ClientCertificateExclusionPaths() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) ClientCertificateExclusionPathsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPathsInput",
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

func (j *jsiiProxy_WindowsWebApp) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_WindowsWebApp) FtpPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) FtpPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) HostingEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingEnvironmentId",
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

func (j *jsiiProxy_WindowsWebApp) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
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

func (j *jsiiProxy_WindowsWebApp) WebdeployPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebApp) WebdeployPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabledInput",
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/windows_web_app azurerm_windows_web_app} Resource.
func NewWindowsWebApp(scope constructs.Construct, id *string, config *WindowsWebAppConfig) WindowsWebApp {
	_init_.Initialize()

	if err := validateNewWindowsWebAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebApp{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/windows_web_app azurerm_windows_web_app} Resource.
func NewWindowsWebApp_Override(w WindowsWebApp, scope constructs.Construct, id *string, config *WindowsWebAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetAppSettings(val *map[string]*string) {
	if err := j.validateSetAppSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetClientAffinityEnabled(val interface{}) {
	if err := j.validateSetClientAffinityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAffinityEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetClientCertificateEnabled(val interface{}) {
	if err := j.validateSetClientCertificateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetClientCertificateExclusionPaths(val *string) {
	if err := j.validateSetClientCertificateExclusionPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateExclusionPaths",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetClientCertificateMode(val *string) {
	if err := j.validateSetClientCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateMode",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetFtpPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetFtpPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetHttpsOnly(val interface{}) {
	if err := j.validateSetHttpsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetKeyVaultReferenceIdentityId(val *string) {
	if err := j.validateSetKeyVaultReferenceIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultReferenceIdentityId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetServicePlanId(val *string) {
	if err := j.validateSetServicePlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePlanId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetVirtualNetworkSubnetId(val *string) {
	if err := j.validateSetVirtualNetworkSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetWebdeployPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetWebdeployPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebApp)SetZipDeployFile(val *string) {
	if err := j.validateSetZipDeployFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipDeployFile",
		val,
	)
}

// Generates CDKTF code for importing a WindowsWebApp resource upon running "cdktf plan <stack-name>".
func WindowsWebApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWindowsWebApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
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
func WindowsWebApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsWebApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsWebApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsWebApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsWebApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsWebApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebApp",
		"isTerraformResource",
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

func (w *jsiiProxy_WindowsWebApp) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WindowsWebApp) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WindowsWebApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WindowsWebApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WindowsWebApp) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WindowsWebApp) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WindowsWebApp) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutAuthSettings(value *WindowsWebAppAuthSettings) {
	if err := w.validatePutAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutAuthSettingsV2(value *WindowsWebAppAuthSettingsV2) {
	if err := w.validatePutAuthSettingsV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAuthSettingsV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutBackup(value *WindowsWebAppBackup) {
	if err := w.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBackup",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutConnectionString(value interface{}) {
	if err := w.validatePutConnectionStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutIdentity(value *WindowsWebAppIdentity) {
	if err := w.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIdentity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutLogs(value *WindowsWebAppLogs) {
	if err := w.validatePutLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLogs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutSiteConfig(value *WindowsWebAppSiteConfig) {
	if err := w.validatePutSiteConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutStickySettings(value *WindowsWebAppStickySettings) {
	if err := w.validatePutStickySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStickySettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutStorageAccount(value interface{}) {
	if err := w.validatePutStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebApp) PutTimeouts(value *WindowsWebAppTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (w *jsiiProxy_WindowsWebApp) ResetAuthSettingsV2() {
	_jsii_.InvokeVoid(
		w,
		"resetAuthSettingsV2",
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

func (w *jsiiProxy_WindowsWebApp) ResetClientCertificateExclusionPaths() {
	_jsii_.InvokeVoid(
		w,
		"resetClientCertificateExclusionPaths",
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

func (w *jsiiProxy_WindowsWebApp) ResetFtpPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetFtpPublishBasicAuthenticationEnabled",
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

func (w *jsiiProxy_WindowsWebApp) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetPublicNetworkAccessEnabled",
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

func (w *jsiiProxy_WindowsWebApp) ResetWebdeployPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetWebdeployPublishBasicAuthenticationEnabled",
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

func (w *jsiiProxy_WindowsWebApp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebApp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
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

