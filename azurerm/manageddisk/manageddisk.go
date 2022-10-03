package manageddisk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/manageddisk/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk azurerm_managed_disk}.
type ManagedDisk interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreateOption() *string
	SetCreateOption(val *string)
	CreateOptionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskAccessId() *string
	SetDiskAccessId(val *string)
	DiskAccessIdInput() *string
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	DiskIopsReadOnly() *float64
	SetDiskIopsReadOnly(val *float64)
	DiskIopsReadOnlyInput() *float64
	DiskIopsReadWrite() *float64
	SetDiskIopsReadWrite(val *float64)
	DiskIopsReadWriteInput() *float64
	DiskMbpsReadOnly() *float64
	SetDiskMbpsReadOnly(val *float64)
	DiskMbpsReadOnlyInput() *float64
	DiskMbpsReadWrite() *float64
	SetDiskMbpsReadWrite(val *float64)
	DiskMbpsReadWriteInput() *float64
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	EncryptionSettings() ManagedDiskEncryptionSettingsOutputReference
	EncryptionSettingsInput() *ManagedDiskEncryptionSettings
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GalleryImageReferenceId() *string
	SetGalleryImageReferenceId(val *string)
	GalleryImageReferenceIdInput() *string
	HyperVGeneration() *string
	SetHyperVGeneration(val *string)
	HyperVGenerationInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageReferenceId() *string
	SetImageReferenceId(val *string)
	ImageReferenceIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogicalSectorSize() *float64
	SetLogicalSectorSize(val *float64)
	LogicalSectorSizeInput() *float64
	MaxShares() *float64
	SetMaxShares(val *float64)
	MaxSharesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAccessPolicy() *string
	SetNetworkAccessPolicy(val *string)
	NetworkAccessPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	OnDemandBurstingEnabled() interface{}
	SetOnDemandBurstingEnabled(val interface{})
	OnDemandBurstingEnabledInput() interface{}
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
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
	SecureVmDiskEncryptionSetId() *string
	SetSecureVmDiskEncryptionSetId(val *string)
	SecureVmDiskEncryptionSetIdInput() *string
	SecurityType() *string
	SetSecurityType(val *string)
	SecurityTypeInput() *string
	SourceResourceId() *string
	SetSourceResourceId(val *string)
	SourceResourceIdInput() *string
	SourceUri() *string
	SetSourceUri(val *string)
	SourceUriInput() *string
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Timeouts() ManagedDiskTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustedLaunchEnabled() interface{}
	SetTrustedLaunchEnabled(val interface{})
	TrustedLaunchEnabledInput() interface{}
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutEncryptionSettings(value *ManagedDiskEncryptionSettings)
	PutTimeouts(value *ManagedDiskTimeouts)
	ResetDiskAccessId()
	ResetDiskEncryptionSetId()
	ResetDiskIopsReadOnly()
	ResetDiskIopsReadWrite()
	ResetDiskMbpsReadOnly()
	ResetDiskMbpsReadWrite()
	ResetDiskSizeGb()
	ResetEdgeZone()
	ResetEncryptionSettings()
	ResetGalleryImageReferenceId()
	ResetHyperVGeneration()
	ResetId()
	ResetImageReferenceId()
	ResetLogicalSectorSize()
	ResetMaxShares()
	ResetNetworkAccessPolicy()
	ResetOnDemandBurstingEnabled()
	ResetOsType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetSecureVmDiskEncryptionSetId()
	ResetSecurityType()
	ResetSourceResourceId()
	ResetSourceUri()
	ResetStorageAccountId()
	ResetTags()
	ResetTier()
	ResetTimeouts()
	ResetTrustedLaunchEnabled()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ManagedDisk
type jsiiProxy_ManagedDisk struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ManagedDisk) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) CreateOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) CreateOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskAccessId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskAccessId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskAccessIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskAccessIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskIopsReadOnly() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskIopsReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskIopsReadOnlyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskIopsReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskIopsReadWrite() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskIopsReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskIopsReadWriteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskIopsReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskMbpsReadOnly() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskMbpsReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskMbpsReadOnlyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskMbpsReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskMbpsReadWrite() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskMbpsReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskMbpsReadWriteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskMbpsReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) EncryptionSettings() ManagedDiskEncryptionSettingsOutputReference {
	var returns ManagedDiskEncryptionSettingsOutputReference
	_jsii_.Get(
		j,
		"encryptionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) EncryptionSettingsInput() *ManagedDiskEncryptionSettings {
	var returns *ManagedDiskEncryptionSettings
	_jsii_.Get(
		j,
		"encryptionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) GalleryImageReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"galleryImageReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) GalleryImageReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"galleryImageReferenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) HyperVGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hyperVGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) HyperVGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hyperVGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) ImageReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) ImageReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageReferenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) LogicalSectorSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logicalSectorSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) LogicalSectorSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logicalSectorSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) MaxShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) MaxSharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) NetworkAccessPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) NetworkAccessPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) OnDemandBurstingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandBurstingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) OnDemandBurstingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandBurstingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) SecureVmDiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureVmDiskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) SecureVmDiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureVmDiskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) SecurityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) SecurityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) SourceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) SourceResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) SourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) SourceUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Timeouts() ManagedDiskTimeoutsOutputReference {
	var returns ManagedDiskTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) TrustedLaunchEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedLaunchEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) TrustedLaunchEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedLaunchEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk azurerm_managed_disk} Resource.
func NewManagedDisk(scope constructs.Construct, id *string, config *ManagedDiskConfig) ManagedDisk {
	_init_.Initialize()

	j := jsiiProxy_ManagedDisk{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk azurerm_managed_disk} Resource.
func NewManagedDisk_Override(m ManagedDisk, scope constructs.Construct, id *string, config *ManagedDiskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagedDisk) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetCreateOption(val *string) {
	_jsii_.Set(
		j,
		"createOption",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetDiskAccessId(val *string) {
	_jsii_.Set(
		j,
		"diskAccessId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetDiskEncryptionSetId(val *string) {
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetDiskIopsReadOnly(val *float64) {
	_jsii_.Set(
		j,
		"diskIopsReadOnly",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetDiskIopsReadWrite(val *float64) {
	_jsii_.Set(
		j,
		"diskIopsReadWrite",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetDiskMbpsReadOnly(val *float64) {
	_jsii_.Set(
		j,
		"diskMbpsReadOnly",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetDiskMbpsReadWrite(val *float64) {
	_jsii_.Set(
		j,
		"diskMbpsReadWrite",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetDiskSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetEdgeZone(val *string) {
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetGalleryImageReferenceId(val *string) {
	_jsii_.Set(
		j,
		"galleryImageReferenceId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetHyperVGeneration(val *string) {
	_jsii_.Set(
		j,
		"hyperVGeneration",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetImageReferenceId(val *string) {
	_jsii_.Set(
		j,
		"imageReferenceId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetLogicalSectorSize(val *float64) {
	_jsii_.Set(
		j,
		"logicalSectorSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetMaxShares(val *float64) {
	_jsii_.Set(
		j,
		"maxShares",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetNetworkAccessPolicy(val *string) {
	_jsii_.Set(
		j,
		"networkAccessPolicy",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetOnDemandBurstingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"onDemandBurstingEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetOsType(val *string) {
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetSecureVmDiskEncryptionSetId(val *string) {
	_jsii_.Set(
		j,
		"secureVmDiskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetSecurityType(val *string) {
	_jsii_.Set(
		j,
		"securityType",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetSourceResourceId(val *string) {
	_jsii_.Set(
		j,
		"sourceResourceId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetSourceUri(val *string) {
	_jsii_.Set(
		j,
		"sourceUri",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetStorageAccountId(val *string) {
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetStorageAccountType(val *string) {
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetTier(val *string) {
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetTrustedLaunchEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"trustedLaunchEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk) SetZone(val *string) {
	_jsii_.Set(
		j,
		"zone",
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
func ManagedDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagedDisk_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagedDisk) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagedDisk) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagedDisk) PutEncryptionSettings(value *ManagedDiskEncryptionSettings) {
	_jsii_.InvokeVoid(
		m,
		"putEncryptionSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDisk) PutTimeouts(value *ManagedDiskTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDisk) ResetDiskAccessId() {
	_jsii_.InvokeVoid(
		m,
		"resetDiskAccessId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		m,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetDiskIopsReadOnly() {
	_jsii_.InvokeVoid(
		m,
		"resetDiskIopsReadOnly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetDiskIopsReadWrite() {
	_jsii_.InvokeVoid(
		m,
		"resetDiskIopsReadWrite",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetDiskMbpsReadOnly() {
	_jsii_.InvokeVoid(
		m,
		"resetDiskMbpsReadOnly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetDiskMbpsReadWrite() {
	_jsii_.InvokeVoid(
		m,
		"resetDiskMbpsReadWrite",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		m,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		m,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetEncryptionSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetGalleryImageReferenceId() {
	_jsii_.InvokeVoid(
		m,
		"resetGalleryImageReferenceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetHyperVGeneration() {
	_jsii_.InvokeVoid(
		m,
		"resetHyperVGeneration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetImageReferenceId() {
	_jsii_.InvokeVoid(
		m,
		"resetImageReferenceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetLogicalSectorSize() {
	_jsii_.InvokeVoid(
		m,
		"resetLogicalSectorSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetMaxShares() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxShares",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetNetworkAccessPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworkAccessPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetOnDemandBurstingEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetOnDemandBurstingEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetOsType() {
	_jsii_.InvokeVoid(
		m,
		"resetOsType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetSecureVmDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		m,
		"resetSecureVmDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetSecurityType() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetSourceResourceId() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceResourceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetSourceUri() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceUri",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetStorageAccountId() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetTier() {
	_jsii_.InvokeVoid(
		m,
		"resetTier",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetTrustedLaunchEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetTrustedLaunchEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) ResetZone() {
	_jsii_.InvokeVoid(
		m,
		"resetZone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDisk) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ManagedDiskConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#create_option ManagedDisk#create_option}.
	CreateOption *string `field:"required" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#location ManagedDisk#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#name ManagedDisk#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#resource_group_name ManagedDisk#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#storage_account_type ManagedDisk#storage_account_type}.
	StorageAccountType *string `field:"required" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_access_id ManagedDisk#disk_access_id}.
	DiskAccessId *string `field:"optional" json:"diskAccessId" yaml:"diskAccessId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_encryption_set_id ManagedDisk#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_iops_read_only ManagedDisk#disk_iops_read_only}.
	DiskIopsReadOnly *float64 `field:"optional" json:"diskIopsReadOnly" yaml:"diskIopsReadOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_iops_read_write ManagedDisk#disk_iops_read_write}.
	DiskIopsReadWrite *float64 `field:"optional" json:"diskIopsReadWrite" yaml:"diskIopsReadWrite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_mbps_read_only ManagedDisk#disk_mbps_read_only}.
	DiskMbpsReadOnly *float64 `field:"optional" json:"diskMbpsReadOnly" yaml:"diskMbpsReadOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_mbps_read_write ManagedDisk#disk_mbps_read_write}.
	DiskMbpsReadWrite *float64 `field:"optional" json:"diskMbpsReadWrite" yaml:"diskMbpsReadWrite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_size_gb ManagedDisk#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#edge_zone ManagedDisk#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// encryption_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#encryption_settings ManagedDisk#encryption_settings}
	EncryptionSettings *ManagedDiskEncryptionSettings `field:"optional" json:"encryptionSettings" yaml:"encryptionSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#gallery_image_reference_id ManagedDisk#gallery_image_reference_id}.
	GalleryImageReferenceId *string `field:"optional" json:"galleryImageReferenceId" yaml:"galleryImageReferenceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#hyper_v_generation ManagedDisk#hyper_v_generation}.
	HyperVGeneration *string `field:"optional" json:"hyperVGeneration" yaml:"hyperVGeneration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#id ManagedDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#image_reference_id ManagedDisk#image_reference_id}.
	ImageReferenceId *string `field:"optional" json:"imageReferenceId" yaml:"imageReferenceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#logical_sector_size ManagedDisk#logical_sector_size}.
	LogicalSectorSize *float64 `field:"optional" json:"logicalSectorSize" yaml:"logicalSectorSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#max_shares ManagedDisk#max_shares}.
	MaxShares *float64 `field:"optional" json:"maxShares" yaml:"maxShares"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#network_access_policy ManagedDisk#network_access_policy}.
	NetworkAccessPolicy *string `field:"optional" json:"networkAccessPolicy" yaml:"networkAccessPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#on_demand_bursting_enabled ManagedDisk#on_demand_bursting_enabled}.
	OnDemandBurstingEnabled interface{} `field:"optional" json:"onDemandBurstingEnabled" yaml:"onDemandBurstingEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#os_type ManagedDisk#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#public_network_access_enabled ManagedDisk#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#secure_vm_disk_encryption_set_id ManagedDisk#secure_vm_disk_encryption_set_id}.
	SecureVmDiskEncryptionSetId *string `field:"optional" json:"secureVmDiskEncryptionSetId" yaml:"secureVmDiskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#security_type ManagedDisk#security_type}.
	SecurityType *string `field:"optional" json:"securityType" yaml:"securityType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#source_resource_id ManagedDisk#source_resource_id}.
	SourceResourceId *string `field:"optional" json:"sourceResourceId" yaml:"sourceResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#source_uri ManagedDisk#source_uri}.
	SourceUri *string `field:"optional" json:"sourceUri" yaml:"sourceUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#storage_account_id ManagedDisk#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#tags ManagedDisk#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#tier ManagedDisk#tier}.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#timeouts ManagedDisk#timeouts}
	Timeouts *ManagedDiskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#trusted_launch_enabled ManagedDisk#trusted_launch_enabled}.
	TrustedLaunchEnabled interface{} `field:"optional" json:"trustedLaunchEnabled" yaml:"trustedLaunchEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#zone ManagedDisk#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

type ManagedDiskEncryptionSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#enabled ManagedDisk#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_encryption_key ManagedDisk#disk_encryption_key}
	DiskEncryptionKey *ManagedDiskEncryptionSettingsDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// key_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#key_encryption_key ManagedDisk#key_encryption_key}
	KeyEncryptionKey *ManagedDiskEncryptionSettingsKeyEncryptionKey `field:"optional" json:"keyEncryptionKey" yaml:"keyEncryptionKey"`
}

type ManagedDiskEncryptionSettingsDiskEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#secret_url ManagedDisk#secret_url}.
	SecretUrl *string `field:"required" json:"secretUrl" yaml:"secretUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#source_vault_id ManagedDisk#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

type ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference interface {
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
	InternalValue() *ManagedDiskEncryptionSettingsDiskEncryptionKey
	SetInternalValue(val *ManagedDiskEncryptionSettingsDiskEncryptionKey)
	SecretUrl() *string
	SetSecretUrl(val *string)
	SecretUrlInput() *string
	SourceVaultId() *string
	SetSourceVaultId(val *string)
	SourceVaultIdInput() *string
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

// The jsii proxy struct for ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference
type jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) InternalValue() *ManagedDiskEncryptionSettingsDiskEncryptionKey {
	var returns *ManagedDiskEncryptionSettingsDiskEncryptionKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SecretUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SecretUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SourceVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SourceVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference_Override(m ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SetInternalValue(val *ManagedDiskEncryptionSettingsDiskEncryptionKey) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SetSecretUrl(val *string) {
	_jsii_.Set(
		j,
		"secretUrl",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SetSourceVaultId(val *string) {
	_jsii_.Set(
		j,
		"sourceVaultId",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ManagedDiskEncryptionSettingsKeyEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#key_url ManagedDisk#key_url}.
	KeyUrl *string `field:"required" json:"keyUrl" yaml:"keyUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#source_vault_id ManagedDisk#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

type ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference interface {
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
	InternalValue() *ManagedDiskEncryptionSettingsKeyEncryptionKey
	SetInternalValue(val *ManagedDiskEncryptionSettingsKeyEncryptionKey)
	KeyUrl() *string
	SetKeyUrl(val *string)
	KeyUrlInput() *string
	SourceVaultId() *string
	SetSourceVaultId(val *string)
	SourceVaultIdInput() *string
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

// The jsii proxy struct for ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference
type jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) InternalValue() *ManagedDiskEncryptionSettingsKeyEncryptionKey {
	var returns *ManagedDiskEncryptionSettingsKeyEncryptionKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) KeyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) KeyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SourceVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SourceVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference_Override(m ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SetInternalValue(val *ManagedDiskEncryptionSettingsKeyEncryptionKey) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SetKeyUrl(val *string) {
	_jsii_.Set(
		j,
		"keyUrl",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SetSourceVaultId(val *string) {
	_jsii_.Set(
		j,
		"sourceVaultId",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ManagedDiskEncryptionSettingsOutputReference interface {
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
	DiskEncryptionKey() ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference
	DiskEncryptionKeyInput() *ManagedDiskEncryptionSettingsDiskEncryptionKey
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ManagedDiskEncryptionSettings
	SetInternalValue(val *ManagedDiskEncryptionSettings)
	KeyEncryptionKey() ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference
	KeyEncryptionKeyInput() *ManagedDiskEncryptionSettingsKeyEncryptionKey
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
	PutDiskEncryptionKey(value *ManagedDiskEncryptionSettingsDiskEncryptionKey)
	PutKeyEncryptionKey(value *ManagedDiskEncryptionSettingsKeyEncryptionKey)
	ResetDiskEncryptionKey()
	ResetKeyEncryptionKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDiskEncryptionSettingsOutputReference
type jsiiProxy_ManagedDiskEncryptionSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) DiskEncryptionKey() ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference {
	var returns ManagedDiskEncryptionSettingsDiskEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"diskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) DiskEncryptionKeyInput() *ManagedDiskEncryptionSettingsDiskEncryptionKey {
	var returns *ManagedDiskEncryptionSettingsDiskEncryptionKey
	_jsii_.Get(
		j,
		"diskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) InternalValue() *ManagedDiskEncryptionSettings {
	var returns *ManagedDiskEncryptionSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) KeyEncryptionKey() ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference {
	var returns ManagedDiskEncryptionSettingsKeyEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"keyEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) KeyEncryptionKeyInput() *ManagedDiskEncryptionSettingsKeyEncryptionKey {
	var returns *ManagedDiskEncryptionSettingsKeyEncryptionKey
	_jsii_.Get(
		j,
		"keyEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDiskEncryptionSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDiskEncryptionSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ManagedDiskEncryptionSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDiskEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDiskEncryptionSettingsOutputReference_Override(m ManagedDiskEncryptionSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDiskEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) SetInternalValue(val *ManagedDiskEncryptionSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) PutDiskEncryptionKey(value *ManagedDiskEncryptionSettingsDiskEncryptionKey) {
	_jsii_.InvokeVoid(
		m,
		"putDiskEncryptionKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) PutKeyEncryptionKey(value *ManagedDiskEncryptionSettingsKeyEncryptionKey) {
	_jsii_.InvokeVoid(
		m,
		"putKeyEncryptionKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) ResetDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		m,
		"resetDiskEncryptionKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) ResetKeyEncryptionKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyEncryptionKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskEncryptionSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ManagedDiskTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#create ManagedDisk#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#delete ManagedDisk#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#read ManagedDisk#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#update ManagedDisk#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ManagedDiskTimeoutsOutputReference interface {
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

// The jsii proxy struct for ManagedDiskTimeoutsOutputReference
type jsiiProxy_ManagedDiskTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewManagedDiskTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDiskTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ManagedDiskTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDiskTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDiskTimeoutsOutputReference_Override(m ManagedDiskTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDiskTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDiskTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDiskTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

