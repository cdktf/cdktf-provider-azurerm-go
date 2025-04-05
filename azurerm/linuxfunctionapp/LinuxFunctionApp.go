// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/linuxfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/linux_function_app azurerm_linux_function_app}.
type LinuxFunctionApp interface {
	cdktf.TerraformResource
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() LinuxFunctionAppAuthSettingsOutputReference
	AuthSettingsInput() *LinuxFunctionAppAuthSettings
	AuthSettingsV2() LinuxFunctionAppAuthSettingsV2OutputReference
	AuthSettingsV2Input() *LinuxFunctionAppAuthSettingsV2
	Backup() LinuxFunctionAppBackupOutputReference
	BackupInput() *LinuxFunctionAppBackup
	BuiltinLoggingEnabled() interface{}
	SetBuiltinLoggingEnabled(val interface{})
	BuiltinLoggingEnabledInput() interface{}
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
	ConnectionString() LinuxFunctionAppConnectionStringList
	ConnectionStringInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentShareForceDisabled() interface{}
	SetContentShareForceDisabled(val interface{})
	ContentShareForceDisabledInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomDomainVerificationId() *string
	DailyMemoryTimeQuota() *float64
	SetDailyMemoryTimeQuota(val *float64)
	DailyMemoryTimeQuotaInput() *float64
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
	FunctionsExtensionVersion() *string
	SetFunctionsExtensionVersion(val *string)
	FunctionsExtensionVersionInput() *string
	HostingEnvironmentId() *string
	HttpsOnly() interface{}
	SetHttpsOnly(val interface{})
	HttpsOnlyInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() LinuxFunctionAppIdentityOutputReference
	IdentityInput() *LinuxFunctionAppIdentity
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
	SiteConfig() LinuxFunctionAppSiteConfigOutputReference
	SiteConfigInput() *LinuxFunctionAppSiteConfig
	SiteCredential() LinuxFunctionAppSiteCredentialList
	StickySettings() LinuxFunctionAppStickySettingsOutputReference
	StickySettingsInput() *LinuxFunctionAppStickySettings
	StorageAccount() LinuxFunctionAppStorageAccountList
	StorageAccountAccessKey() *string
	SetStorageAccountAccessKey(val *string)
	StorageAccountAccessKeyInput() *string
	StorageAccountInput() interface{}
	StorageAccountName() *string
	SetStorageAccountName(val *string)
	StorageAccountNameInput() *string
	StorageKeyVaultSecretId() *string
	SetStorageKeyVaultSecretId(val *string)
	StorageKeyVaultSecretIdInput() *string
	StorageUsesManagedIdentity() interface{}
	SetStorageUsesManagedIdentity(val interface{})
	StorageUsesManagedIdentityInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LinuxFunctionAppTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkBackupRestoreEnabled() interface{}
	SetVirtualNetworkBackupRestoreEnabled(val interface{})
	VirtualNetworkBackupRestoreEnabledInput() interface{}
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
	VnetImagePullEnabled() interface{}
	SetVnetImagePullEnabled(val interface{})
	VnetImagePullEnabledInput() interface{}
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
	PutAuthSettings(value *LinuxFunctionAppAuthSettings)
	PutAuthSettingsV2(value *LinuxFunctionAppAuthSettingsV2)
	PutBackup(value *LinuxFunctionAppBackup)
	PutConnectionString(value interface{})
	PutIdentity(value *LinuxFunctionAppIdentity)
	PutSiteConfig(value *LinuxFunctionAppSiteConfig)
	PutStickySettings(value *LinuxFunctionAppStickySettings)
	PutStorageAccount(value interface{})
	PutTimeouts(value *LinuxFunctionAppTimeouts)
	ResetAppSettings()
	ResetAuthSettings()
	ResetAuthSettingsV2()
	ResetBackup()
	ResetBuiltinLoggingEnabled()
	ResetClientCertificateEnabled()
	ResetClientCertificateExclusionPaths()
	ResetClientCertificateMode()
	ResetConnectionString()
	ResetContentShareForceDisabled()
	ResetDailyMemoryTimeQuota()
	ResetEnabled()
	ResetFtpPublishBasicAuthenticationEnabled()
	ResetFunctionsExtensionVersion()
	ResetHttpsOnly()
	ResetId()
	ResetIdentity()
	ResetKeyVaultReferenceIdentityId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetStickySettings()
	ResetStorageAccount()
	ResetStorageAccountAccessKey()
	ResetStorageAccountName()
	ResetStorageKeyVaultSecretId()
	ResetStorageUsesManagedIdentity()
	ResetTags()
	ResetTimeouts()
	ResetVirtualNetworkBackupRestoreEnabled()
	ResetVirtualNetworkSubnetId()
	ResetVnetImagePullEnabled()
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

// The jsii proxy struct for LinuxFunctionApp
type jsiiProxy_LinuxFunctionApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LinuxFunctionApp) AppSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) AppSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) AuthSettings() LinuxFunctionAppAuthSettingsOutputReference {
	var returns LinuxFunctionAppAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) AuthSettingsInput() *LinuxFunctionAppAuthSettings {
	var returns *LinuxFunctionAppAuthSettings
	_jsii_.Get(
		j,
		"authSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) AuthSettingsV2() LinuxFunctionAppAuthSettingsV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2OutputReference
	_jsii_.Get(
		j,
		"authSettingsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) AuthSettingsV2Input() *LinuxFunctionAppAuthSettingsV2 {
	var returns *LinuxFunctionAppAuthSettingsV2
	_jsii_.Get(
		j,
		"authSettingsV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Backup() LinuxFunctionAppBackupOutputReference {
	var returns LinuxFunctionAppBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) BackupInput() *LinuxFunctionAppBackup {
	var returns *LinuxFunctionAppBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) BuiltinLoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"builtinLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) BuiltinLoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"builtinLoggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ClientCertificateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ClientCertificateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ClientCertificateExclusionPaths() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ClientCertificateExclusionPathsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ClientCertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ClientCertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ConnectionString() LinuxFunctionAppConnectionStringList {
	var returns LinuxFunctionAppConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ConnectionStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ContentShareForceDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentShareForceDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ContentShareForceDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentShareForceDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) DailyMemoryTimeQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) DailyMemoryTimeQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) FtpPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) FtpPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) FunctionsExtensionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionsExtensionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) FunctionsExtensionVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionsExtensionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) HostingEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Identity() LinuxFunctionAppIdentityOutputReference {
	var returns LinuxFunctionAppIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) IdentityInput() *LinuxFunctionAppIdentity {
	var returns *LinuxFunctionAppIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) KeyVaultReferenceIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) KeyVaultReferenceIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) OutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) PossibleOutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ServicePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) SiteConfig() LinuxFunctionAppSiteConfigOutputReference {
	var returns LinuxFunctionAppSiteConfigOutputReference
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) SiteConfigInput() *LinuxFunctionAppSiteConfig {
	var returns *LinuxFunctionAppSiteConfig
	_jsii_.Get(
		j,
		"siteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) SiteCredential() LinuxFunctionAppSiteCredentialList {
	var returns LinuxFunctionAppSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StickySettings() LinuxFunctionAppStickySettingsOutputReference {
	var returns LinuxFunctionAppStickySettingsOutputReference
	_jsii_.Get(
		j,
		"stickySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StickySettingsInput() *LinuxFunctionAppStickySettings {
	var returns *LinuxFunctionAppStickySettings
	_jsii_.Get(
		j,
		"stickySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageAccount() LinuxFunctionAppStorageAccountList {
	var returns LinuxFunctionAppStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageKeyVaultSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyVaultSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageKeyVaultSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyVaultSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageUsesManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageUsesManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) StorageUsesManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageUsesManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) Timeouts() LinuxFunctionAppTimeoutsOutputReference {
	var returns LinuxFunctionAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) VirtualNetworkBackupRestoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualNetworkBackupRestoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) VirtualNetworkBackupRestoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualNetworkBackupRestoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) VnetImagePullEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetImagePullEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) VnetImagePullEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetImagePullEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) WebdeployPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) WebdeployPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ZipDeployFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipDeployFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionApp) ZipDeployFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipDeployFileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/linux_function_app azurerm_linux_function_app} Resource.
func NewLinuxFunctionApp(scope constructs.Construct, id *string, config *LinuxFunctionAppConfig) LinuxFunctionApp {
	_init_.Initialize()

	if err := validateNewLinuxFunctionAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxFunctionApp{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/linux_function_app azurerm_linux_function_app} Resource.
func NewLinuxFunctionApp_Override(l LinuxFunctionApp, scope constructs.Construct, id *string, config *LinuxFunctionAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionApp",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetAppSettings(val *map[string]*string) {
	if err := j.validateSetAppSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetBuiltinLoggingEnabled(val interface{}) {
	if err := j.validateSetBuiltinLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtinLoggingEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetClientCertificateEnabled(val interface{}) {
	if err := j.validateSetClientCertificateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetClientCertificateExclusionPaths(val *string) {
	if err := j.validateSetClientCertificateExclusionPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateExclusionPaths",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetClientCertificateMode(val *string) {
	if err := j.validateSetClientCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateMode",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetContentShareForceDisabled(val interface{}) {
	if err := j.validateSetContentShareForceDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentShareForceDisabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetDailyMemoryTimeQuota(val *float64) {
	if err := j.validateSetDailyMemoryTimeQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyMemoryTimeQuota",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetFtpPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetFtpPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetFunctionsExtensionVersion(val *string) {
	if err := j.validateSetFunctionsExtensionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionsExtensionVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetHttpsOnly(val interface{}) {
	if err := j.validateSetHttpsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetKeyVaultReferenceIdentityId(val *string) {
	if err := j.validateSetKeyVaultReferenceIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultReferenceIdentityId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetServicePlanId(val *string) {
	if err := j.validateSetServicePlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePlanId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetStorageAccountAccessKey(val *string) {
	if err := j.validateSetStorageAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetStorageAccountName(val *string) {
	if err := j.validateSetStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetStorageKeyVaultSecretId(val *string) {
	if err := j.validateSetStorageKeyVaultSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageKeyVaultSecretId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetStorageUsesManagedIdentity(val interface{}) {
	if err := j.validateSetStorageUsesManagedIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageUsesManagedIdentity",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetVirtualNetworkBackupRestoreEnabled(val interface{}) {
	if err := j.validateSetVirtualNetworkBackupRestoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkBackupRestoreEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetVirtualNetworkSubnetId(val *string) {
	if err := j.validateSetVirtualNetworkSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetVnetImagePullEnabled(val interface{}) {
	if err := j.validateSetVnetImagePullEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetImagePullEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetWebdeployPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetWebdeployPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionApp)SetZipDeployFile(val *string) {
	if err := j.validateSetZipDeployFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipDeployFile",
		val,
	)
}

// Generates CDKTF code for importing a LinuxFunctionApp resource upon running "cdktf plan <stack-name>".
func LinuxFunctionApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLinuxFunctionApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionApp",
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
func LinuxFunctionApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxFunctionApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxFunctionApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxFunctionApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxFunctionApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxFunctionApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionApp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LinuxFunctionApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxFunctionApp) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxFunctionApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxFunctionApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxFunctionApp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxFunctionApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxFunctionApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxFunctionApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxFunctionApp) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxFunctionApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxFunctionApp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionApp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionApp) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutAuthSettings(value *LinuxFunctionAppAuthSettings) {
	if err := l.validatePutAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutAuthSettingsV2(value *LinuxFunctionAppAuthSettingsV2) {
	if err := l.validatePutAuthSettingsV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAuthSettingsV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutBackup(value *LinuxFunctionAppBackup) {
	if err := l.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBackup",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutConnectionString(value interface{}) {
	if err := l.validatePutConnectionStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutIdentity(value *LinuxFunctionAppIdentity) {
	if err := l.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutSiteConfig(value *LinuxFunctionAppSiteConfig) {
	if err := l.validatePutSiteConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutStickySettings(value *LinuxFunctionAppStickySettings) {
	if err := l.validatePutStickySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStickySettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutStorageAccount(value interface{}) {
	if err := l.validatePutStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) PutTimeouts(value *LinuxFunctionAppTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetAppSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetAppSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetAuthSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetAuthSettingsV2() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthSettingsV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetBackup() {
	_jsii_.InvokeVoid(
		l,
		"resetBackup",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetBuiltinLoggingEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetBuiltinLoggingEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetClientCertificateEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetClientCertificateExclusionPaths() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateExclusionPaths",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetClientCertificateMode() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetConnectionString() {
	_jsii_.InvokeVoid(
		l,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetContentShareForceDisabled() {
	_jsii_.InvokeVoid(
		l,
		"resetContentShareForceDisabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetDailyMemoryTimeQuota() {
	_jsii_.InvokeVoid(
		l,
		"resetDailyMemoryTimeQuota",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetFtpPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetFtpPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetFunctionsExtensionVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetFunctionsExtensionVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetKeyVaultReferenceIdentityId() {
	_jsii_.InvokeVoid(
		l,
		"resetKeyVaultReferenceIdentityId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetStickySettings() {
	_jsii_.InvokeVoid(
		l,
		"resetStickySettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetStorageAccountAccessKey() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageAccountAccessKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetStorageAccountName() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageAccountName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetStorageKeyVaultSecretId() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageKeyVaultSecretId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetStorageUsesManagedIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageUsesManagedIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetVirtualNetworkBackupRestoreEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetVirtualNetworkBackupRestoreEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		l,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetVnetImagePullEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetVnetImagePullEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetWebdeployPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetWebdeployPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) ResetZipDeployFile() {
	_jsii_.InvokeVoid(
		l,
		"resetZipDeployFile",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionApp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionApp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

