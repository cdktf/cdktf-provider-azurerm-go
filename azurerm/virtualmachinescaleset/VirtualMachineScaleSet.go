// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/virtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/virtual_machine_scale_set azurerm_virtual_machine_scale_set}.
type VirtualMachineScaleSet interface {
	cdktf.TerraformResource
	AutomaticOsUpgrade() interface{}
	SetAutomaticOsUpgrade(val interface{})
	AutomaticOsUpgradeInput() interface{}
	BootDiagnostics() VirtualMachineScaleSetBootDiagnosticsOutputReference
	BootDiagnosticsInput() *VirtualMachineScaleSetBootDiagnostics
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
	Extension() VirtualMachineScaleSetExtensionList
	ExtensionInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthProbeId() *string
	SetHealthProbeId(val *string)
	HealthProbeIdInput() *string
	Id() *string
	SetId(val *string)
	Identity() VirtualMachineScaleSetIdentityOutputReference
	IdentityInput() *VirtualMachineScaleSetIdentity
	IdInput() *string
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
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
	NetworkProfile() VirtualMachineScaleSetNetworkProfileList
	NetworkProfileInput() interface{}
	// The tree node.
	Node() constructs.Node
	OsProfile() VirtualMachineScaleSetOsProfileOutputReference
	OsProfileInput() *VirtualMachineScaleSetOsProfile
	OsProfileLinuxConfig() VirtualMachineScaleSetOsProfileLinuxConfigOutputReference
	OsProfileLinuxConfigInput() *VirtualMachineScaleSetOsProfileLinuxConfig
	OsProfileSecrets() VirtualMachineScaleSetOsProfileSecretsList
	OsProfileSecretsInput() interface{}
	OsProfileWindowsConfig() VirtualMachineScaleSetOsProfileWindowsConfigOutputReference
	OsProfileWindowsConfigInput() *VirtualMachineScaleSetOsProfileWindowsConfig
	Overprovision() interface{}
	SetOverprovision(val interface{})
	OverprovisionInput() interface{}
	Plan() VirtualMachineScaleSetPlanOutputReference
	PlanInput() *VirtualMachineScaleSetPlan
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProximityPlacementGroupId() *string
	SetProximityPlacementGroupId(val *string)
	ProximityPlacementGroupIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RollingUpgradePolicy() VirtualMachineScaleSetRollingUpgradePolicyOutputReference
	RollingUpgradePolicyInput() *VirtualMachineScaleSetRollingUpgradePolicy
	SinglePlacementGroup() interface{}
	SetSinglePlacementGroup(val interface{})
	SinglePlacementGroupInput() interface{}
	Sku() VirtualMachineScaleSetSkuOutputReference
	SkuInput() *VirtualMachineScaleSetSku
	StorageProfileDataDisk() VirtualMachineScaleSetStorageProfileDataDiskList
	StorageProfileDataDiskInput() interface{}
	StorageProfileImageReference() VirtualMachineScaleSetStorageProfileImageReferenceOutputReference
	StorageProfileImageReferenceInput() *VirtualMachineScaleSetStorageProfileImageReference
	StorageProfileOsDisk() VirtualMachineScaleSetStorageProfileOsDiskOutputReference
	StorageProfileOsDiskInput() *VirtualMachineScaleSetStorageProfileOsDisk
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualMachineScaleSetTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpgradePolicyMode() *string
	SetUpgradePolicyMode(val *string)
	UpgradePolicyModeInput() *string
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	PutBootDiagnostics(value *VirtualMachineScaleSetBootDiagnostics)
	PutExtension(value interface{})
	PutIdentity(value *VirtualMachineScaleSetIdentity)
	PutNetworkProfile(value interface{})
	PutOsProfile(value *VirtualMachineScaleSetOsProfile)
	PutOsProfileLinuxConfig(value *VirtualMachineScaleSetOsProfileLinuxConfig)
	PutOsProfileSecrets(value interface{})
	PutOsProfileWindowsConfig(value *VirtualMachineScaleSetOsProfileWindowsConfig)
	PutPlan(value *VirtualMachineScaleSetPlan)
	PutRollingUpgradePolicy(value *VirtualMachineScaleSetRollingUpgradePolicy)
	PutSku(value *VirtualMachineScaleSetSku)
	PutStorageProfileDataDisk(value interface{})
	PutStorageProfileImageReference(value *VirtualMachineScaleSetStorageProfileImageReference)
	PutStorageProfileOsDisk(value *VirtualMachineScaleSetStorageProfileOsDisk)
	PutTimeouts(value *VirtualMachineScaleSetTimeouts)
	ResetAutomaticOsUpgrade()
	ResetBootDiagnostics()
	ResetEvictionPolicy()
	ResetExtension()
	ResetHealthProbeId()
	ResetId()
	ResetIdentity()
	ResetLicenseType()
	ResetOsProfileLinuxConfig()
	ResetOsProfileSecrets()
	ResetOsProfileWindowsConfig()
	ResetOverprovision()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlan()
	ResetPriority()
	ResetProximityPlacementGroupId()
	ResetRollingUpgradePolicy()
	ResetSinglePlacementGroup()
	ResetStorageProfileDataDisk()
	ResetStorageProfileImageReference()
	ResetTags()
	ResetTimeouts()
	ResetZones()
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

// The jsii proxy struct for VirtualMachineScaleSet
type jsiiProxy_VirtualMachineScaleSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualMachineScaleSet) AutomaticOsUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticOsUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) AutomaticOsUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticOsUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) BootDiagnostics() VirtualMachineScaleSetBootDiagnosticsOutputReference {
	var returns VirtualMachineScaleSetBootDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) BootDiagnosticsInput() *VirtualMachineScaleSetBootDiagnostics {
	var returns *VirtualMachineScaleSetBootDiagnostics
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Extension() VirtualMachineScaleSetExtensionList {
	var returns VirtualMachineScaleSetExtensionList
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) ExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) HealthProbeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) HealthProbeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Identity() VirtualMachineScaleSetIdentityOutputReference {
	var returns VirtualMachineScaleSetIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) IdentityInput() *VirtualMachineScaleSetIdentity {
	var returns *VirtualMachineScaleSetIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) NetworkProfile() VirtualMachineScaleSetNetworkProfileList {
	var returns VirtualMachineScaleSetNetworkProfileList
	_jsii_.Get(
		j,
		"networkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) NetworkProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OsProfile() VirtualMachineScaleSetOsProfileOutputReference {
	var returns VirtualMachineScaleSetOsProfileOutputReference
	_jsii_.Get(
		j,
		"osProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OsProfileInput() *VirtualMachineScaleSetOsProfile {
	var returns *VirtualMachineScaleSetOsProfile
	_jsii_.Get(
		j,
		"osProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OsProfileLinuxConfig() VirtualMachineScaleSetOsProfileLinuxConfigOutputReference {
	var returns VirtualMachineScaleSetOsProfileLinuxConfigOutputReference
	_jsii_.Get(
		j,
		"osProfileLinuxConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OsProfileLinuxConfigInput() *VirtualMachineScaleSetOsProfileLinuxConfig {
	var returns *VirtualMachineScaleSetOsProfileLinuxConfig
	_jsii_.Get(
		j,
		"osProfileLinuxConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OsProfileSecrets() VirtualMachineScaleSetOsProfileSecretsList {
	var returns VirtualMachineScaleSetOsProfileSecretsList
	_jsii_.Get(
		j,
		"osProfileSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OsProfileSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"osProfileSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OsProfileWindowsConfig() VirtualMachineScaleSetOsProfileWindowsConfigOutputReference {
	var returns VirtualMachineScaleSetOsProfileWindowsConfigOutputReference
	_jsii_.Get(
		j,
		"osProfileWindowsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OsProfileWindowsConfigInput() *VirtualMachineScaleSetOsProfileWindowsConfig {
	var returns *VirtualMachineScaleSetOsProfileWindowsConfig
	_jsii_.Get(
		j,
		"osProfileWindowsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Overprovision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overprovision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) OverprovisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overprovisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Plan() VirtualMachineScaleSetPlanOutputReference {
	var returns VirtualMachineScaleSetPlanOutputReference
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) PlanInput() *VirtualMachineScaleSetPlan {
	var returns *VirtualMachineScaleSetPlan
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) RollingUpgradePolicy() VirtualMachineScaleSetRollingUpgradePolicyOutputReference {
	var returns VirtualMachineScaleSetRollingUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"rollingUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) RollingUpgradePolicyInput() *VirtualMachineScaleSetRollingUpgradePolicy {
	var returns *VirtualMachineScaleSetRollingUpgradePolicy
	_jsii_.Get(
		j,
		"rollingUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) SinglePlacementGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singlePlacementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) SinglePlacementGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singlePlacementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Sku() VirtualMachineScaleSetSkuOutputReference {
	var returns VirtualMachineScaleSetSkuOutputReference
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) SkuInput() *VirtualMachineScaleSetSku {
	var returns *VirtualMachineScaleSetSku
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) StorageProfileDataDisk() VirtualMachineScaleSetStorageProfileDataDiskList {
	var returns VirtualMachineScaleSetStorageProfileDataDiskList
	_jsii_.Get(
		j,
		"storageProfileDataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) StorageProfileDataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageProfileDataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) StorageProfileImageReference() VirtualMachineScaleSetStorageProfileImageReferenceOutputReference {
	var returns VirtualMachineScaleSetStorageProfileImageReferenceOutputReference
	_jsii_.Get(
		j,
		"storageProfileImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) StorageProfileImageReferenceInput() *VirtualMachineScaleSetStorageProfileImageReference {
	var returns *VirtualMachineScaleSetStorageProfileImageReference
	_jsii_.Get(
		j,
		"storageProfileImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) StorageProfileOsDisk() VirtualMachineScaleSetStorageProfileOsDiskOutputReference {
	var returns VirtualMachineScaleSetStorageProfileOsDiskOutputReference
	_jsii_.Get(
		j,
		"storageProfileOsDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) StorageProfileOsDiskInput() *VirtualMachineScaleSetStorageProfileOsDisk {
	var returns *VirtualMachineScaleSetStorageProfileOsDisk
	_jsii_.Get(
		j,
		"storageProfileOsDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Timeouts() VirtualMachineScaleSetTimeoutsOutputReference {
	var returns VirtualMachineScaleSetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) UpgradePolicyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradePolicyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) UpgradePolicyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradePolicyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSet) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/virtual_machine_scale_set azurerm_virtual_machine_scale_set} Resource.
func NewVirtualMachineScaleSet(scope constructs.Construct, id *string, config *VirtualMachineScaleSetConfig) VirtualMachineScaleSet {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSet{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/virtual_machine_scale_set azurerm_virtual_machine_scale_set} Resource.
func NewVirtualMachineScaleSet_Override(v VirtualMachineScaleSet, scope constructs.Construct, id *string, config *VirtualMachineScaleSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSet",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetAutomaticOsUpgrade(val interface{}) {
	if err := j.validateSetAutomaticOsUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticOsUpgrade",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetHealthProbeId(val *string) {
	if err := j.validateSetHealthProbeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthProbeId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetOverprovision(val interface{}) {
	if err := j.validateSetOverprovisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overprovision",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetSinglePlacementGroup(val interface{}) {
	if err := j.validateSetSinglePlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singlePlacementGroup",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetUpgradePolicyMode(val *string) {
	if err := j.validateSetUpgradePolicyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradePolicyMode",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSet)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Generates CDKTF code for importing a VirtualMachineScaleSet resource upon running "cdktf plan <stack-name>".
func VirtualMachineScaleSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSet",
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
func VirtualMachineScaleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachineScaleSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachineScaleSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualMachineScaleSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutBootDiagnostics(value *VirtualMachineScaleSetBootDiagnostics) {
	if err := v.validatePutBootDiagnosticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutExtension(value interface{}) {
	if err := v.validatePutExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putExtension",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutIdentity(value *VirtualMachineScaleSetIdentity) {
	if err := v.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIdentity",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutNetworkProfile(value interface{}) {
	if err := v.validatePutNetworkProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNetworkProfile",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutOsProfile(value *VirtualMachineScaleSetOsProfile) {
	if err := v.validatePutOsProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOsProfile",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutOsProfileLinuxConfig(value *VirtualMachineScaleSetOsProfileLinuxConfig) {
	if err := v.validatePutOsProfileLinuxConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOsProfileLinuxConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutOsProfileSecrets(value interface{}) {
	if err := v.validatePutOsProfileSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOsProfileSecrets",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutOsProfileWindowsConfig(value *VirtualMachineScaleSetOsProfileWindowsConfig) {
	if err := v.validatePutOsProfileWindowsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOsProfileWindowsConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutPlan(value *VirtualMachineScaleSetPlan) {
	if err := v.validatePutPlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPlan",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutRollingUpgradePolicy(value *VirtualMachineScaleSetRollingUpgradePolicy) {
	if err := v.validatePutRollingUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRollingUpgradePolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutSku(value *VirtualMachineScaleSetSku) {
	if err := v.validatePutSkuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSku",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutStorageProfileDataDisk(value interface{}) {
	if err := v.validatePutStorageProfileDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageProfileDataDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutStorageProfileImageReference(value *VirtualMachineScaleSetStorageProfileImageReference) {
	if err := v.validatePutStorageProfileImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageProfileImageReference",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutStorageProfileOsDisk(value *VirtualMachineScaleSetStorageProfileOsDisk) {
	if err := v.validatePutStorageProfileOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageProfileOsDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) PutTimeouts(value *VirtualMachineScaleSetTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetAutomaticOsUpgrade() {
	_jsii_.InvokeVoid(
		v,
		"resetAutomaticOsUpgrade",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		v,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetExtension() {
	_jsii_.InvokeVoid(
		v,
		"resetExtension",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetHealthProbeId() {
	_jsii_.InvokeVoid(
		v,
		"resetHealthProbeId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetIdentity() {
	_jsii_.InvokeVoid(
		v,
		"resetIdentity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetLicenseType() {
	_jsii_.InvokeVoid(
		v,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetOsProfileLinuxConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetOsProfileLinuxConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetOsProfileSecrets() {
	_jsii_.InvokeVoid(
		v,
		"resetOsProfileSecrets",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetOsProfileWindowsConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetOsProfileWindowsConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetOverprovision() {
	_jsii_.InvokeVoid(
		v,
		"resetOverprovision",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetPlan() {
	_jsii_.InvokeVoid(
		v,
		"resetPlan",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetPriority() {
	_jsii_.InvokeVoid(
		v,
		"resetPriority",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		v,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetRollingUpgradePolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetRollingUpgradePolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetSinglePlacementGroup() {
	_jsii_.InvokeVoid(
		v,
		"resetSinglePlacementGroup",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetStorageProfileDataDisk() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageProfileDataDisk",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetStorageProfileImageReference() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageProfileImageReference",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) ResetZones() {
	_jsii_.InvokeVoid(
		v,
		"resetZones",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

