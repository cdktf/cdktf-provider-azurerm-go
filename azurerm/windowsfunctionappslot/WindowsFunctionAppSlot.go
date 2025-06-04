// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/windowsfunctionappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/windows_function_app_slot azurerm_windows_function_app_slot}.
type WindowsFunctionAppSlot interface {
	cdktf.TerraformResource
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() WindowsFunctionAppSlotAuthSettingsOutputReference
	AuthSettingsInput() *WindowsFunctionAppSlotAuthSettings
	AuthSettingsV2() WindowsFunctionAppSlotAuthSettingsV2OutputReference
	AuthSettingsV2Input() *WindowsFunctionAppSlotAuthSettingsV2
	Backup() WindowsFunctionAppSlotBackupOutputReference
	BackupInput() *WindowsFunctionAppSlotBackup
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
	ConnectionString() WindowsFunctionAppSlotConnectionStringList
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
	Identity() WindowsFunctionAppSlotIdentityOutputReference
	IdentityInput() *WindowsFunctionAppSlotIdentity
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
	SiteConfig() WindowsFunctionAppSlotSiteConfigOutputReference
	SiteConfigInput() *WindowsFunctionAppSlotSiteConfig
	SiteCredential() WindowsFunctionAppSlotSiteCredentialList
	StorageAccount() WindowsFunctionAppSlotStorageAccountList
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
	Timeouts() WindowsFunctionAppSlotTimeoutsOutputReference
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
	PutAuthSettings(value *WindowsFunctionAppSlotAuthSettings)
	PutAuthSettingsV2(value *WindowsFunctionAppSlotAuthSettingsV2)
	PutBackup(value *WindowsFunctionAppSlotBackup)
	PutConnectionString(value interface{})
	PutIdentity(value *WindowsFunctionAppSlotIdentity)
	PutSiteConfig(value *WindowsFunctionAppSlotSiteConfig)
	PutStorageAccount(value interface{})
	PutTimeouts(value *WindowsFunctionAppSlotTimeouts)
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
	ResetVirtualNetworkBackupRestoreEnabled()
	ResetVirtualNetworkSubnetId()
	ResetVnetImagePullEnabled()
	ResetWebdeployPublishBasicAuthenticationEnabled()
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

// The jsii proxy struct for WindowsFunctionAppSlot
type jsiiProxy_WindowsFunctionAppSlot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WindowsFunctionAppSlot) AppSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) AppSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) AuthSettings() WindowsFunctionAppSlotAuthSettingsOutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) AuthSettingsInput() *WindowsFunctionAppSlotAuthSettings {
	var returns *WindowsFunctionAppSlotAuthSettings
	_jsii_.Get(
		j,
		"authSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) AuthSettingsV2() WindowsFunctionAppSlotAuthSettingsV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2OutputReference
	_jsii_.Get(
		j,
		"authSettingsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) AuthSettingsV2Input() *WindowsFunctionAppSlotAuthSettingsV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2
	_jsii_.Get(
		j,
		"authSettingsV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Backup() WindowsFunctionAppSlotBackupOutputReference {
	var returns WindowsFunctionAppSlotBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) BackupInput() *WindowsFunctionAppSlotBackup {
	var returns *WindowsFunctionAppSlotBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) BuiltinLoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"builtinLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) BuiltinLoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"builtinLoggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ClientCertificateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ClientCertificateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ClientCertificateExclusionPaths() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ClientCertificateExclusionPathsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ClientCertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ClientCertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ConnectionString() WindowsFunctionAppSlotConnectionStringList {
	var returns WindowsFunctionAppSlotConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ConnectionStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ContentShareForceDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentShareForceDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ContentShareForceDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentShareForceDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) DailyMemoryTimeQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) DailyMemoryTimeQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) FtpPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) FtpPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) FunctionAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) FunctionAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) FunctionsExtensionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionsExtensionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) FunctionsExtensionVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionsExtensionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) HostingEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Identity() WindowsFunctionAppSlotIdentityOutputReference {
	var returns WindowsFunctionAppSlotIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) IdentityInput() *WindowsFunctionAppSlotIdentity {
	var returns *WindowsFunctionAppSlotIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) KeyVaultReferenceIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) KeyVaultReferenceIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) OutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) PossibleOutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) ServicePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) SiteConfig() WindowsFunctionAppSlotSiteConfigOutputReference {
	var returns WindowsFunctionAppSlotSiteConfigOutputReference
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) SiteConfigInput() *WindowsFunctionAppSlotSiteConfig {
	var returns *WindowsFunctionAppSlotSiteConfig
	_jsii_.Get(
		j,
		"siteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) SiteCredential() WindowsFunctionAppSlotSiteCredentialList {
	var returns WindowsFunctionAppSlotSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageAccount() WindowsFunctionAppSlotStorageAccountList {
	var returns WindowsFunctionAppSlotStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageKeyVaultSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyVaultSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageKeyVaultSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyVaultSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageUsesManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageUsesManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) StorageUsesManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageUsesManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) Timeouts() WindowsFunctionAppSlotTimeoutsOutputReference {
	var returns WindowsFunctionAppSlotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) VirtualNetworkBackupRestoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualNetworkBackupRestoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) VirtualNetworkBackupRestoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualNetworkBackupRestoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) VnetImagePullEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetImagePullEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) VnetImagePullEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetImagePullEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) WebdeployPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlot) WebdeployPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/windows_function_app_slot azurerm_windows_function_app_slot} Resource.
func NewWindowsFunctionAppSlot(scope constructs.Construct, id *string, config *WindowsFunctionAppSlotConfig) WindowsFunctionAppSlot {
	_init_.Initialize()

	if err := validateNewWindowsFunctionAppSlotParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsFunctionAppSlot{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/windows_function_app_slot azurerm_windows_function_app_slot} Resource.
func NewWindowsFunctionAppSlot_Override(w WindowsFunctionAppSlot, scope constructs.Construct, id *string, config *WindowsFunctionAppSlotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlot",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetAppSettings(val *map[string]*string) {
	if err := j.validateSetAppSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetBuiltinLoggingEnabled(val interface{}) {
	if err := j.validateSetBuiltinLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtinLoggingEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetClientCertificateEnabled(val interface{}) {
	if err := j.validateSetClientCertificateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetClientCertificateExclusionPaths(val *string) {
	if err := j.validateSetClientCertificateExclusionPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateExclusionPaths",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetClientCertificateMode(val *string) {
	if err := j.validateSetClientCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateMode",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetContentShareForceDisabled(val interface{}) {
	if err := j.validateSetContentShareForceDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentShareForceDisabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetDailyMemoryTimeQuota(val *float64) {
	if err := j.validateSetDailyMemoryTimeQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyMemoryTimeQuota",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetFtpPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetFtpPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetFunctionAppId(val *string) {
	if err := j.validateSetFunctionAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionAppId",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetFunctionsExtensionVersion(val *string) {
	if err := j.validateSetFunctionsExtensionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionsExtensionVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetHttpsOnly(val interface{}) {
	if err := j.validateSetHttpsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetKeyVaultReferenceIdentityId(val *string) {
	if err := j.validateSetKeyVaultReferenceIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultReferenceIdentityId",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetServicePlanId(val *string) {
	if err := j.validateSetServicePlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePlanId",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetStorageAccountAccessKey(val *string) {
	if err := j.validateSetStorageAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetStorageAccountName(val *string) {
	if err := j.validateSetStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetStorageKeyVaultSecretId(val *string) {
	if err := j.validateSetStorageKeyVaultSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageKeyVaultSecretId",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetStorageUsesManagedIdentity(val interface{}) {
	if err := j.validateSetStorageUsesManagedIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageUsesManagedIdentity",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetVirtualNetworkBackupRestoreEnabled(val interface{}) {
	if err := j.validateSetVirtualNetworkBackupRestoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkBackupRestoreEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetVirtualNetworkSubnetId(val *string) {
	if err := j.validateSetVirtualNetworkSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetVnetImagePullEnabled(val interface{}) {
	if err := j.validateSetVnetImagePullEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetImagePullEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlot)SetWebdeployPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetWebdeployPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		val,
	)
}

// Generates CDKTF code for importing a WindowsFunctionAppSlot resource upon running "cdktf plan <stack-name>".
func WindowsFunctionAppSlot_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWindowsFunctionAppSlot_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlot",
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
func WindowsFunctionAppSlot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsFunctionAppSlot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsFunctionAppSlot_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsFunctionAppSlot_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlot",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsFunctionAppSlot_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsFunctionAppSlot_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlot",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WindowsFunctionAppSlot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlot) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsFunctionAppSlot) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) PutAuthSettings(value *WindowsFunctionAppSlotAuthSettings) {
	if err := w.validatePutAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) PutAuthSettingsV2(value *WindowsFunctionAppSlotAuthSettingsV2) {
	if err := w.validatePutAuthSettingsV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAuthSettingsV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) PutBackup(value *WindowsFunctionAppSlotBackup) {
	if err := w.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBackup",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) PutConnectionString(value interface{}) {
	if err := w.validatePutConnectionStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) PutIdentity(value *WindowsFunctionAppSlotIdentity) {
	if err := w.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIdentity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) PutSiteConfig(value *WindowsFunctionAppSlotSiteConfig) {
	if err := w.validatePutSiteConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) PutStorageAccount(value interface{}) {
	if err := w.validatePutStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) PutTimeouts(value *WindowsFunctionAppSlotTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetAppSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetAppSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetAuthSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetAuthSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetAuthSettingsV2() {
	_jsii_.InvokeVoid(
		w,
		"resetAuthSettingsV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetBackup() {
	_jsii_.InvokeVoid(
		w,
		"resetBackup",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetBuiltinLoggingEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetBuiltinLoggingEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetClientCertificateEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetClientCertificateEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetClientCertificateExclusionPaths() {
	_jsii_.InvokeVoid(
		w,
		"resetClientCertificateExclusionPaths",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetClientCertificateMode() {
	_jsii_.InvokeVoid(
		w,
		"resetClientCertificateMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetConnectionString() {
	_jsii_.InvokeVoid(
		w,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetContentShareForceDisabled() {
	_jsii_.InvokeVoid(
		w,
		"resetContentShareForceDisabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetDailyMemoryTimeQuota() {
	_jsii_.InvokeVoid(
		w,
		"resetDailyMemoryTimeQuota",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetFtpPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetFtpPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetFunctionsExtensionVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetFunctionsExtensionVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetKeyVaultReferenceIdentityId() {
	_jsii_.InvokeVoid(
		w,
		"resetKeyVaultReferenceIdentityId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetServicePlanId() {
	_jsii_.InvokeVoid(
		w,
		"resetServicePlanId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		w,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetStorageAccountAccessKey() {
	_jsii_.InvokeVoid(
		w,
		"resetStorageAccountAccessKey",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetStorageAccountName() {
	_jsii_.InvokeVoid(
		w,
		"resetStorageAccountName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetStorageKeyVaultSecretId() {
	_jsii_.InvokeVoid(
		w,
		"resetStorageKeyVaultSecretId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetStorageUsesManagedIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetStorageUsesManagedIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetVirtualNetworkBackupRestoreEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualNetworkBackupRestoreEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetVnetImagePullEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetVnetImagePullEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ResetWebdeployPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetWebdeployPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlot) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

