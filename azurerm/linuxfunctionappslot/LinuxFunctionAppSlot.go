// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/linuxfunctionappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/linux_function_app_slot azurerm_linux_function_app_slot}.
type LinuxFunctionAppSlot interface {
	cdktf.TerraformResource
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() LinuxFunctionAppSlotAuthSettingsOutputReference
	AuthSettingsInput() *LinuxFunctionAppSlotAuthSettings
	AuthSettingsV2() LinuxFunctionAppSlotAuthSettingsV2OutputReference
	AuthSettingsV2Input() *LinuxFunctionAppSlotAuthSettingsV2
	Backup() LinuxFunctionAppSlotBackupOutputReference
	BackupInput() *LinuxFunctionAppSlotBackup
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
	ConnectionString() LinuxFunctionAppSlotConnectionStringList
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
	FunctionAppId() *string
	SetFunctionAppId(val *string)
	FunctionAppIdInput() *string
	FunctionsExtensionVersion() *string
	SetFunctionsExtensionVersion(val *string)
	FunctionsExtensionVersionInput() *string
	HostingEnvironmentId() *string
	HttpsOnly() interface{}
	SetHttpsOnly(val interface{})
	HttpsOnlyInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() LinuxFunctionAppSlotIdentityOutputReference
	IdentityInput() *LinuxFunctionAppSlotIdentity
	IdInput() *string
	KeyVaultReferenceIdentityId() *string
	SetKeyVaultReferenceIdentityId(val *string)
	KeyVaultReferenceIdentityIdInput() *string
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	ServicePlanId() *string
	SetServicePlanId(val *string)
	ServicePlanIdInput() *string
	SiteConfig() LinuxFunctionAppSlotSiteConfigOutputReference
	SiteConfigInput() *LinuxFunctionAppSlotSiteConfig
	SiteCredential() LinuxFunctionAppSlotSiteCredentialList
	StorageAccount() LinuxFunctionAppSlotStorageAccountList
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
	Timeouts() LinuxFunctionAppSlotTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
	WebdeployPublishBasicAuthenticationEnabled() interface{}
	SetWebdeployPublishBasicAuthenticationEnabled(val interface{})
	WebdeployPublishBasicAuthenticationEnabledInput() interface{}
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAuthSettings(value *LinuxFunctionAppSlotAuthSettings)
	PutAuthSettingsV2(value *LinuxFunctionAppSlotAuthSettingsV2)
	PutBackup(value *LinuxFunctionAppSlotBackup)
	PutConnectionString(value interface{})
	PutIdentity(value *LinuxFunctionAppSlotIdentity)
	PutSiteConfig(value *LinuxFunctionAppSlotSiteConfig)
	PutStorageAccount(value interface{})
	PutTimeouts(value *LinuxFunctionAppSlotTimeouts)
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
	ResetServicePlanId()
	ResetStorageAccount()
	ResetStorageAccountAccessKey()
	ResetStorageAccountName()
	ResetStorageKeyVaultSecretId()
	ResetStorageUsesManagedIdentity()
	ResetTags()
	ResetTimeouts()
	ResetVirtualNetworkSubnetId()
	ResetWebdeployPublishBasicAuthenticationEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LinuxFunctionAppSlot
type jsiiProxy_LinuxFunctionAppSlot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LinuxFunctionAppSlot) AppSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) AppSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) AuthSettings() LinuxFunctionAppSlotAuthSettingsOutputReference {
	var returns LinuxFunctionAppSlotAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) AuthSettingsInput() *LinuxFunctionAppSlotAuthSettings {
	var returns *LinuxFunctionAppSlotAuthSettings
	_jsii_.Get(
		j,
		"authSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) AuthSettingsV2() LinuxFunctionAppSlotAuthSettingsV2OutputReference {
	var returns LinuxFunctionAppSlotAuthSettingsV2OutputReference
	_jsii_.Get(
		j,
		"authSettingsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) AuthSettingsV2Input() *LinuxFunctionAppSlotAuthSettingsV2 {
	var returns *LinuxFunctionAppSlotAuthSettingsV2
	_jsii_.Get(
		j,
		"authSettingsV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Backup() LinuxFunctionAppSlotBackupOutputReference {
	var returns LinuxFunctionAppSlotBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) BackupInput() *LinuxFunctionAppSlotBackup {
	var returns *LinuxFunctionAppSlotBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) BuiltinLoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"builtinLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) BuiltinLoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"builtinLoggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ClientCertificateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ClientCertificateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ClientCertificateExclusionPaths() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ClientCertificateExclusionPathsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ClientCertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ClientCertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ConnectionString() LinuxFunctionAppSlotConnectionStringList {
	var returns LinuxFunctionAppSlotConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ConnectionStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ContentShareForceDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentShareForceDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ContentShareForceDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentShareForceDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) DailyMemoryTimeQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) DailyMemoryTimeQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) FtpPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) FtpPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) FunctionAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) FunctionAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) FunctionsExtensionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionsExtensionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) FunctionsExtensionVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionsExtensionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) HostingEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Identity() LinuxFunctionAppSlotIdentityOutputReference {
	var returns LinuxFunctionAppSlotIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) IdentityInput() *LinuxFunctionAppSlotIdentity {
	var returns *LinuxFunctionAppSlotIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) KeyVaultReferenceIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) KeyVaultReferenceIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) OutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) PossibleOutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) ServicePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) SiteConfig() LinuxFunctionAppSlotSiteConfigOutputReference {
	var returns LinuxFunctionAppSlotSiteConfigOutputReference
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) SiteConfigInput() *LinuxFunctionAppSlotSiteConfig {
	var returns *LinuxFunctionAppSlotSiteConfig
	_jsii_.Get(
		j,
		"siteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) SiteCredential() LinuxFunctionAppSlotSiteCredentialList {
	var returns LinuxFunctionAppSlotSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageAccount() LinuxFunctionAppSlotStorageAccountList {
	var returns LinuxFunctionAppSlotStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageKeyVaultSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyVaultSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageKeyVaultSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyVaultSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageUsesManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageUsesManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) StorageUsesManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageUsesManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) Timeouts() LinuxFunctionAppSlotTimeoutsOutputReference {
	var returns LinuxFunctionAppSlotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) WebdeployPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlot) WebdeployPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/linux_function_app_slot azurerm_linux_function_app_slot} Resource.
func NewLinuxFunctionAppSlot(scope constructs.Construct, id *string, config *LinuxFunctionAppSlotConfig) LinuxFunctionAppSlot {
	_init_.Initialize()

	if err := validateNewLinuxFunctionAppSlotParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxFunctionAppSlot{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/linux_function_app_slot azurerm_linux_function_app_slot} Resource.
func NewLinuxFunctionAppSlot_Override(l LinuxFunctionAppSlot, scope constructs.Construct, id *string, config *LinuxFunctionAppSlotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlot",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetAppSettings(val *map[string]*string) {
	if err := j.validateSetAppSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetBuiltinLoggingEnabled(val interface{}) {
	if err := j.validateSetBuiltinLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtinLoggingEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetClientCertificateEnabled(val interface{}) {
	if err := j.validateSetClientCertificateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetClientCertificateExclusionPaths(val *string) {
	if err := j.validateSetClientCertificateExclusionPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateExclusionPaths",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetClientCertificateMode(val *string) {
	if err := j.validateSetClientCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateMode",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetContentShareForceDisabled(val interface{}) {
	if err := j.validateSetContentShareForceDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentShareForceDisabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetDailyMemoryTimeQuota(val *float64) {
	if err := j.validateSetDailyMemoryTimeQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyMemoryTimeQuota",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetFtpPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetFtpPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetFunctionAppId(val *string) {
	if err := j.validateSetFunctionAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionAppId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetFunctionsExtensionVersion(val *string) {
	if err := j.validateSetFunctionsExtensionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionsExtensionVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetHttpsOnly(val interface{}) {
	if err := j.validateSetHttpsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetKeyVaultReferenceIdentityId(val *string) {
	if err := j.validateSetKeyVaultReferenceIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultReferenceIdentityId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetServicePlanId(val *string) {
	if err := j.validateSetServicePlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePlanId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetStorageAccountAccessKey(val *string) {
	if err := j.validateSetStorageAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetStorageAccountName(val *string) {
	if err := j.validateSetStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetStorageKeyVaultSecretId(val *string) {
	if err := j.validateSetStorageKeyVaultSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageKeyVaultSecretId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetStorageUsesManagedIdentity(val interface{}) {
	if err := j.validateSetStorageUsesManagedIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageUsesManagedIdentity",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetVirtualNetworkSubnetId(val *string) {
	if err := j.validateSetVirtualNetworkSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlot)SetWebdeployPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetWebdeployPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		val,
	)
}

// Generates CDKTF code for importing a LinuxFunctionAppSlot resource upon running "cdktf plan <stack-name>".
func LinuxFunctionAppSlot_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLinuxFunctionAppSlot_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlot",
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
func LinuxFunctionAppSlot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxFunctionAppSlot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxFunctionAppSlot_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxFunctionAppSlot_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlot",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxFunctionAppSlot_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxFunctionAppSlot_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlot",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LinuxFunctionAppSlot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlot) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxFunctionAppSlot) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) PutAuthSettings(value *LinuxFunctionAppSlotAuthSettings) {
	if err := l.validatePutAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) PutAuthSettingsV2(value *LinuxFunctionAppSlotAuthSettingsV2) {
	if err := l.validatePutAuthSettingsV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAuthSettingsV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) PutBackup(value *LinuxFunctionAppSlotBackup) {
	if err := l.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBackup",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) PutConnectionString(value interface{}) {
	if err := l.validatePutConnectionStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) PutIdentity(value *LinuxFunctionAppSlotIdentity) {
	if err := l.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) PutSiteConfig(value *LinuxFunctionAppSlotSiteConfig) {
	if err := l.validatePutSiteConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) PutStorageAccount(value interface{}) {
	if err := l.validatePutStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) PutTimeouts(value *LinuxFunctionAppSlotTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetAppSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetAppSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetAuthSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetAuthSettingsV2() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthSettingsV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetBackup() {
	_jsii_.InvokeVoid(
		l,
		"resetBackup",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetBuiltinLoggingEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetBuiltinLoggingEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetClientCertificateEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetClientCertificateExclusionPaths() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateExclusionPaths",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetClientCertificateMode() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetConnectionString() {
	_jsii_.InvokeVoid(
		l,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetContentShareForceDisabled() {
	_jsii_.InvokeVoid(
		l,
		"resetContentShareForceDisabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetDailyMemoryTimeQuota() {
	_jsii_.InvokeVoid(
		l,
		"resetDailyMemoryTimeQuota",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetFtpPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetFtpPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetFunctionsExtensionVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetFunctionsExtensionVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetKeyVaultReferenceIdentityId() {
	_jsii_.InvokeVoid(
		l,
		"resetKeyVaultReferenceIdentityId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetServicePlanId() {
	_jsii_.InvokeVoid(
		l,
		"resetServicePlanId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetStorageAccountAccessKey() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageAccountAccessKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetStorageAccountName() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageAccountName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetStorageKeyVaultSecretId() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageKeyVaultSecretId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetStorageUsesManagedIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageUsesManagedIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		l,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ResetWebdeployPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetWebdeployPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

