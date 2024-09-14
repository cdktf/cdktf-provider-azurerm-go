// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddisk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/manageddisk/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/managed_disk azurerm_managed_disk}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	OptimizedFrequentAttachEnabled() interface{}
	SetOptimizedFrequentAttachEnabled(val interface{})
	OptimizedFrequentAttachEnabledInput() interface{}
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	PerformancePlusEnabled() interface{}
	SetPerformancePlusEnabled(val interface{})
	PerformancePlusEnabledInput() interface{}
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
	UploadSizeBytes() *float64
	SetUploadSizeBytes(val *float64)
	UploadSizeBytesInput() *float64
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	ResetOptimizedFrequentAttachEnabled()
	ResetOsType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformancePlusEnabled()
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
	ResetUploadSizeBytes()
	ResetZone()
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

func (j *jsiiProxy_ManagedDisk) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ManagedDisk) OptimizedFrequentAttachEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optimizedFrequentAttachEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) OptimizedFrequentAttachEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optimizedFrequentAttachEnabledInput",
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

func (j *jsiiProxy_ManagedDisk) PerformancePlusEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performancePlusEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) PerformancePlusEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performancePlusEnabledInput",
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

func (j *jsiiProxy_ManagedDisk) UploadSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uploadSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDisk) UploadSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uploadSizeBytesInput",
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/managed_disk azurerm_managed_disk} Resource.
func NewManagedDisk(scope constructs.Construct, id *string, config *ManagedDiskConfig) ManagedDisk {
	_init_.Initialize()

	if err := validateNewManagedDiskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDisk{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/managed_disk azurerm_managed_disk} Resource.
func NewManagedDisk_Override(m ManagedDisk, scope constructs.Construct, id *string, config *ManagedDiskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagedDisk)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetCreateOption(val *string) {
	if err := j.validateSetCreateOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createOption",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetDiskAccessId(val *string) {
	if err := j.validateSetDiskAccessIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskAccessId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetDiskEncryptionSetId(val *string) {
	if err := j.validateSetDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetDiskIopsReadOnly(val *float64) {
	if err := j.validateSetDiskIopsReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskIopsReadOnly",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetDiskIopsReadWrite(val *float64) {
	if err := j.validateSetDiskIopsReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskIopsReadWrite",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetDiskMbpsReadOnly(val *float64) {
	if err := j.validateSetDiskMbpsReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskMbpsReadOnly",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetDiskMbpsReadWrite(val *float64) {
	if err := j.validateSetDiskMbpsReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskMbpsReadWrite",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetEdgeZone(val *string) {
	if err := j.validateSetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetGalleryImageReferenceId(val *string) {
	if err := j.validateSetGalleryImageReferenceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"galleryImageReferenceId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetHyperVGeneration(val *string) {
	if err := j.validateSetHyperVGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hyperVGeneration",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetImageReferenceId(val *string) {
	if err := j.validateSetImageReferenceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageReferenceId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetLogicalSectorSize(val *float64) {
	if err := j.validateSetLogicalSectorSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logicalSectorSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetMaxShares(val *float64) {
	if err := j.validateSetMaxSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxShares",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetNetworkAccessPolicy(val *string) {
	if err := j.validateSetNetworkAccessPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAccessPolicy",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetOnDemandBurstingEnabled(val interface{}) {
	if err := j.validateSetOnDemandBurstingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandBurstingEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetOptimizedFrequentAttachEnabled(val interface{}) {
	if err := j.validateSetOptimizedFrequentAttachEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizedFrequentAttachEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetPerformancePlusEnabled(val interface{}) {
	if err := j.validateSetPerformancePlusEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performancePlusEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetSecureVmDiskEncryptionSetId(val *string) {
	if err := j.validateSetSecureVmDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureVmDiskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetSecurityType(val *string) {
	if err := j.validateSetSecurityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityType",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetSourceResourceId(val *string) {
	if err := j.validateSetSourceResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceResourceId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetSourceUri(val *string) {
	if err := j.validateSetSourceUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceUri",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetStorageAccountId(val *string) {
	if err := j.validateSetStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetStorageAccountType(val *string) {
	if err := j.validateSetStorageAccountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetTrustedLaunchEnabled(val interface{}) {
	if err := j.validateSetTrustedLaunchEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedLaunchEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetUploadSizeBytes(val *float64) {
	if err := j.validateSetUploadSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadSizeBytes",
		val,
	)
}

func (j *jsiiProxy_ManagedDisk)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a ManagedDisk resource upon running "cdktf plan <stack-name>".
func ManagedDisk_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateManagedDisk_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
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
func ManagedDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDisk_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedDisk_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDisk_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedDisk_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDisk_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managedDisk.ManagedDisk",
		"isTerraformResource",
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

func (m *jsiiProxy_ManagedDisk) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_ManagedDisk) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagedDisk) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_ManagedDisk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedDisk) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_ManagedDisk) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedDisk) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagedDisk) PutEncryptionSettings(value *ManagedDiskEncryptionSettings) {
	if err := m.validatePutEncryptionSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDisk) PutTimeouts(value *ManagedDiskTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (m *jsiiProxy_ManagedDisk) ResetOptimizedFrequentAttachEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetOptimizedFrequentAttachEnabled",
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

func (m *jsiiProxy_ManagedDisk) ResetPerformancePlusEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetPerformancePlusEnabled",
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

func (m *jsiiProxy_ManagedDisk) ResetUploadSizeBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetUploadSizeBytes",
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

func (m *jsiiProxy_ManagedDisk) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDisk) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
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

