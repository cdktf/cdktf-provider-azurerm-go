// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/linuxwebappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/linux_web_app_slot azurerm_linux_web_app_slot}.
type LinuxWebAppSlot interface {
	cdktf.TerraformResource
	AppMetadata() cdktf.StringMap
	AppServiceId() *string
	SetAppServiceId(val *string)
	AppServiceIdInput() *string
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() LinuxWebAppSlotAuthSettingsOutputReference
	AuthSettingsInput() *LinuxWebAppSlotAuthSettings
	AuthSettingsV2() LinuxWebAppSlotAuthSettingsV2OutputReference
	AuthSettingsV2Input() *LinuxWebAppSlotAuthSettingsV2
	Backup() LinuxWebAppSlotBackupOutputReference
	BackupInput() *LinuxWebAppSlotBackup
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
	ConnectionString() LinuxWebAppSlotConnectionStringList
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
	Identity() LinuxWebAppSlotIdentityOutputReference
	IdentityInput() *LinuxWebAppSlotIdentity
	IdInput() *string
	KeyVaultReferenceIdentityId() *string
	SetKeyVaultReferenceIdentityId(val *string)
	KeyVaultReferenceIdentityIdInput() *string
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logs() LinuxWebAppSlotLogsOutputReference
	LogsInput() *LinuxWebAppSlotLogs
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
	SiteConfig() LinuxWebAppSlotSiteConfigOutputReference
	SiteConfigInput() *LinuxWebAppSlotSiteConfig
	SiteCredential() LinuxWebAppSlotSiteCredentialList
	StorageAccount() LinuxWebAppSlotStorageAccountList
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
	Timeouts() LinuxWebAppSlotTimeoutsOutputReference
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
	PutAuthSettings(value *LinuxWebAppSlotAuthSettings)
	PutAuthSettingsV2(value *LinuxWebAppSlotAuthSettingsV2)
	PutBackup(value *LinuxWebAppSlotBackup)
	PutConnectionString(value interface{})
	PutIdentity(value *LinuxWebAppSlotIdentity)
	PutLogs(value *LinuxWebAppSlotLogs)
	PutSiteConfig(value *LinuxWebAppSlotSiteConfig)
	PutStorageAccount(value interface{})
	PutTimeouts(value *LinuxWebAppSlotTimeouts)
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
	ResetServicePlanId()
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

// The jsii proxy struct for LinuxWebAppSlot
type jsiiProxy_LinuxWebAppSlot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LinuxWebAppSlot) AppMetadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"appMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) AppServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) AppServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) AppSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) AppSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) AuthSettings() LinuxWebAppSlotAuthSettingsOutputReference {
	var returns LinuxWebAppSlotAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) AuthSettingsInput() *LinuxWebAppSlotAuthSettings {
	var returns *LinuxWebAppSlotAuthSettings
	_jsii_.Get(
		j,
		"authSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) AuthSettingsV2() LinuxWebAppSlotAuthSettingsV2OutputReference {
	var returns LinuxWebAppSlotAuthSettingsV2OutputReference
	_jsii_.Get(
		j,
		"authSettingsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) AuthSettingsV2Input() *LinuxWebAppSlotAuthSettingsV2 {
	var returns *LinuxWebAppSlotAuthSettingsV2
	_jsii_.Get(
		j,
		"authSettingsV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Backup() LinuxWebAppSlotBackupOutputReference {
	var returns LinuxWebAppSlotBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) BackupInput() *LinuxWebAppSlotBackup {
	var returns *LinuxWebAppSlotBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ClientAffinityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAffinityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ClientAffinityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAffinityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ClientCertificateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ClientCertificateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ClientCertificateExclusionPaths() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ClientCertificateExclusionPathsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateExclusionPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ClientCertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ClientCertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ConnectionString() LinuxWebAppSlotConnectionStringList {
	var returns LinuxWebAppSlotConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ConnectionStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) FtpPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) FtpPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ftpPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) HostingEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Identity() LinuxWebAppSlotIdentityOutputReference {
	var returns LinuxWebAppSlotIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) IdentityInput() *LinuxWebAppSlotIdentity {
	var returns *LinuxWebAppSlotIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) KeyVaultReferenceIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) KeyVaultReferenceIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Logs() LinuxWebAppSlotLogsOutputReference {
	var returns LinuxWebAppSlotLogsOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) LogsInput() *LinuxWebAppSlotLogs {
	var returns *LinuxWebAppSlotLogs
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) OutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) PossibleOutboundIpAddressList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ServicePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) SiteConfig() LinuxWebAppSlotSiteConfigOutputReference {
	var returns LinuxWebAppSlotSiteConfigOutputReference
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) SiteConfigInput() *LinuxWebAppSlotSiteConfig {
	var returns *LinuxWebAppSlotSiteConfig
	_jsii_.Get(
		j,
		"siteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) SiteCredential() LinuxWebAppSlotSiteCredentialList {
	var returns LinuxWebAppSlotSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) StorageAccount() LinuxWebAppSlotStorageAccountList {
	var returns LinuxWebAppSlotStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) Timeouts() LinuxWebAppSlotTimeoutsOutputReference {
	var returns LinuxWebAppSlotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) WebdeployPublishBasicAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) WebdeployPublishBasicAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webdeployPublishBasicAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ZipDeployFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipDeployFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlot) ZipDeployFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipDeployFileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/linux_web_app_slot azurerm_linux_web_app_slot} Resource.
func NewLinuxWebAppSlot(scope constructs.Construct, id *string, config *LinuxWebAppSlotConfig) LinuxWebAppSlot {
	_init_.Initialize()

	if err := validateNewLinuxWebAppSlotParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxWebAppSlot{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/linux_web_app_slot azurerm_linux_web_app_slot} Resource.
func NewLinuxWebAppSlot_Override(l LinuxWebAppSlot, scope constructs.Construct, id *string, config *LinuxWebAppSlotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlot",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetAppServiceId(val *string) {
	if err := j.validateSetAppServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appServiceId",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetAppSettings(val *map[string]*string) {
	if err := j.validateSetAppSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetClientAffinityEnabled(val interface{}) {
	if err := j.validateSetClientAffinityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAffinityEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetClientCertificateEnabled(val interface{}) {
	if err := j.validateSetClientCertificateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetClientCertificateExclusionPaths(val *string) {
	if err := j.validateSetClientCertificateExclusionPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateExclusionPaths",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetClientCertificateMode(val *string) {
	if err := j.validateSetClientCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateMode",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetFtpPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetFtpPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetHttpsOnly(val interface{}) {
	if err := j.validateSetHttpsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetKeyVaultReferenceIdentityId(val *string) {
	if err := j.validateSetKeyVaultReferenceIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultReferenceIdentityId",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetServicePlanId(val *string) {
	if err := j.validateSetServicePlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePlanId",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetVirtualNetworkSubnetId(val *string) {
	if err := j.validateSetVirtualNetworkSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetWebdeployPublishBasicAuthenticationEnabled(val interface{}) {
	if err := j.validateSetWebdeployPublishBasicAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webdeployPublishBasicAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlot)SetZipDeployFile(val *string) {
	if err := j.validateSetZipDeployFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipDeployFile",
		val,
	)
}

// Generates CDKTF code for importing a LinuxWebAppSlot resource upon running "cdktf plan <stack-name>".
func LinuxWebAppSlot_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLinuxWebAppSlot_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlot",
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
func LinuxWebAppSlot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxWebAppSlot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxWebAppSlot_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxWebAppSlot_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlot",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxWebAppSlot_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxWebAppSlot_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlot",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LinuxWebAppSlot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxWebAppSlot) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxWebAppSlot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxWebAppSlot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxWebAppSlot) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxWebAppSlot) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxWebAppSlot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxWebAppSlot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxWebAppSlot) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxWebAppSlot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxWebAppSlot) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlot) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxWebAppSlot) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutAuthSettings(value *LinuxWebAppSlotAuthSettings) {
	if err := l.validatePutAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutAuthSettingsV2(value *LinuxWebAppSlotAuthSettingsV2) {
	if err := l.validatePutAuthSettingsV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAuthSettingsV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutBackup(value *LinuxWebAppSlotBackup) {
	if err := l.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBackup",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutConnectionString(value interface{}) {
	if err := l.validatePutConnectionStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutIdentity(value *LinuxWebAppSlotIdentity) {
	if err := l.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutLogs(value *LinuxWebAppSlotLogs) {
	if err := l.validatePutLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLogs",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutSiteConfig(value *LinuxWebAppSlotSiteConfig) {
	if err := l.validatePutSiteConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutStorageAccount(value interface{}) {
	if err := l.validatePutStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) PutTimeouts(value *LinuxWebAppSlotTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetAppSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetAppSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetAuthSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetAuthSettingsV2() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthSettingsV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetBackup() {
	_jsii_.InvokeVoid(
		l,
		"resetBackup",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetClientAffinityEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetClientAffinityEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetClientCertificateEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetClientCertificateExclusionPaths() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateExclusionPaths",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetClientCertificateMode() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCertificateMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetConnectionString() {
	_jsii_.InvokeVoid(
		l,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetFtpPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetFtpPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetKeyVaultReferenceIdentityId() {
	_jsii_.InvokeVoid(
		l,
		"resetKeyVaultReferenceIdentityId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetLogs() {
	_jsii_.InvokeVoid(
		l,
		"resetLogs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetServicePlanId() {
	_jsii_.InvokeVoid(
		l,
		"resetServicePlanId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		l,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		l,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetWebdeployPublishBasicAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetWebdeployPublishBasicAuthenticationEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) ResetZipDeployFile() {
	_jsii_.InvokeVoid(
		l,
		"resetZipDeployFile",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlot) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlot) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

