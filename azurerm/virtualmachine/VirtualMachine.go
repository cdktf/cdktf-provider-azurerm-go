// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/virtual_machine azurerm_virtual_machine}.
type VirtualMachine interface {
	cdktf.TerraformResource
	AdditionalCapabilities() VirtualMachineAdditionalCapabilitiesOutputReference
	AdditionalCapabilitiesInput() *VirtualMachineAdditionalCapabilities
	AvailabilitySetId() *string
	SetAvailabilitySetId(val *string)
	AvailabilitySetIdInput() *string
	BootDiagnostics() VirtualMachineBootDiagnosticsOutputReference
	BootDiagnosticsInput() *VirtualMachineBootDiagnostics
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
	DeleteDataDisksOnTermination() interface{}
	SetDeleteDataDisksOnTermination(val interface{})
	DeleteDataDisksOnTerminationInput() interface{}
	DeleteOsDiskOnTermination() interface{}
	SetDeleteOsDiskOnTermination(val interface{})
	DeleteOsDiskOnTerminationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	Identity() VirtualMachineIdentityOutputReference
	IdentityInput() *VirtualMachineIdentity
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
	NetworkInterfaceIds() *[]*string
	SetNetworkInterfaceIds(val *[]*string)
	NetworkInterfaceIdsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	OsProfile() VirtualMachineOsProfileOutputReference
	OsProfileInput() *VirtualMachineOsProfile
	OsProfileLinuxConfig() VirtualMachineOsProfileLinuxConfigOutputReference
	OsProfileLinuxConfigInput() *VirtualMachineOsProfileLinuxConfig
	OsProfileSecrets() VirtualMachineOsProfileSecretsList
	OsProfileSecretsInput() interface{}
	OsProfileWindowsConfig() VirtualMachineOsProfileWindowsConfigOutputReference
	OsProfileWindowsConfigInput() *VirtualMachineOsProfileWindowsConfig
	Plan() VirtualMachinePlanOutputReference
	PlanInput() *VirtualMachinePlan
	PrimaryNetworkInterfaceId() *string
	SetPrimaryNetworkInterfaceId(val *string)
	PrimaryNetworkInterfaceIdInput() *string
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
	StorageDataDisk() VirtualMachineStorageDataDiskList
	StorageDataDiskInput() interface{}
	StorageImageReference() VirtualMachineStorageImageReferenceOutputReference
	StorageImageReferenceInput() *VirtualMachineStorageImageReference
	StorageOsDisk() VirtualMachineStorageOsDiskOutputReference
	StorageOsDiskInput() *VirtualMachineStorageOsDisk
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualMachineTimeoutsOutputReference
	TimeoutsInput() interface{}
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	PutAdditionalCapabilities(value *VirtualMachineAdditionalCapabilities)
	PutBootDiagnostics(value *VirtualMachineBootDiagnostics)
	PutIdentity(value *VirtualMachineIdentity)
	PutOsProfile(value *VirtualMachineOsProfile)
	PutOsProfileLinuxConfig(value *VirtualMachineOsProfileLinuxConfig)
	PutOsProfileSecrets(value interface{})
	PutOsProfileWindowsConfig(value *VirtualMachineOsProfileWindowsConfig)
	PutPlan(value *VirtualMachinePlan)
	PutStorageDataDisk(value interface{})
	PutStorageImageReference(value *VirtualMachineStorageImageReference)
	PutStorageOsDisk(value *VirtualMachineStorageOsDisk)
	PutTimeouts(value *VirtualMachineTimeouts)
	ResetAdditionalCapabilities()
	ResetAvailabilitySetId()
	ResetBootDiagnostics()
	ResetDeleteDataDisksOnTermination()
	ResetDeleteOsDiskOnTermination()
	ResetId()
	ResetIdentity()
	ResetLicenseType()
	ResetOsProfile()
	ResetOsProfileLinuxConfig()
	ResetOsProfileSecrets()
	ResetOsProfileWindowsConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlan()
	ResetPrimaryNetworkInterfaceId()
	ResetProximityPlacementGroupId()
	ResetStorageDataDisk()
	ResetStorageImageReference()
	ResetTags()
	ResetTimeouts()
	ResetZones()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VirtualMachine
type jsiiProxy_VirtualMachine struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualMachine) AdditionalCapabilities() VirtualMachineAdditionalCapabilitiesOutputReference {
	var returns VirtualMachineAdditionalCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"additionalCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) AdditionalCapabilitiesInput() *VirtualMachineAdditionalCapabilities {
	var returns *VirtualMachineAdditionalCapabilities
	_jsii_.Get(
		j,
		"additionalCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) AvailabilitySetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) AvailabilitySetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) BootDiagnostics() VirtualMachineBootDiagnosticsOutputReference {
	var returns VirtualMachineBootDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) BootDiagnosticsInput() *VirtualMachineBootDiagnostics {
	var returns *VirtualMachineBootDiagnostics
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DeleteDataDisksOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteDataDisksOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DeleteDataDisksOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteDataDisksOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DeleteOsDiskOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOsDiskOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DeleteOsDiskOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOsDiskOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Identity() VirtualMachineIdentityOutputReference {
	var returns VirtualMachineIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) IdentityInput() *VirtualMachineIdentity {
	var returns *VirtualMachineIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NetworkInterfaceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OsProfile() VirtualMachineOsProfileOutputReference {
	var returns VirtualMachineOsProfileOutputReference
	_jsii_.Get(
		j,
		"osProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OsProfileInput() *VirtualMachineOsProfile {
	var returns *VirtualMachineOsProfile
	_jsii_.Get(
		j,
		"osProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OsProfileLinuxConfig() VirtualMachineOsProfileLinuxConfigOutputReference {
	var returns VirtualMachineOsProfileLinuxConfigOutputReference
	_jsii_.Get(
		j,
		"osProfileLinuxConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OsProfileLinuxConfigInput() *VirtualMachineOsProfileLinuxConfig {
	var returns *VirtualMachineOsProfileLinuxConfig
	_jsii_.Get(
		j,
		"osProfileLinuxConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OsProfileSecrets() VirtualMachineOsProfileSecretsList {
	var returns VirtualMachineOsProfileSecretsList
	_jsii_.Get(
		j,
		"osProfileSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OsProfileSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"osProfileSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OsProfileWindowsConfig() VirtualMachineOsProfileWindowsConfigOutputReference {
	var returns VirtualMachineOsProfileWindowsConfigOutputReference
	_jsii_.Get(
		j,
		"osProfileWindowsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OsProfileWindowsConfigInput() *VirtualMachineOsProfileWindowsConfig {
	var returns *VirtualMachineOsProfileWindowsConfig
	_jsii_.Get(
		j,
		"osProfileWindowsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Plan() VirtualMachinePlanOutputReference {
	var returns VirtualMachinePlanOutputReference
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) PlanInput() *VirtualMachinePlan {
	var returns *VirtualMachinePlan
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) PrimaryNetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNetworkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) PrimaryNetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNetworkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) StorageDataDisk() VirtualMachineStorageDataDiskList {
	var returns VirtualMachineStorageDataDiskList
	_jsii_.Get(
		j,
		"storageDataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) StorageDataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageDataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) StorageImageReference() VirtualMachineStorageImageReferenceOutputReference {
	var returns VirtualMachineStorageImageReferenceOutputReference
	_jsii_.Get(
		j,
		"storageImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) StorageImageReferenceInput() *VirtualMachineStorageImageReference {
	var returns *VirtualMachineStorageImageReference
	_jsii_.Get(
		j,
		"storageImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) StorageOsDisk() VirtualMachineStorageOsDiskOutputReference {
	var returns VirtualMachineStorageOsDiskOutputReference
	_jsii_.Get(
		j,
		"storageOsDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) StorageOsDiskInput() *VirtualMachineStorageOsDisk {
	var returns *VirtualMachineStorageOsDisk
	_jsii_.Get(
		j,
		"storageOsDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Timeouts() VirtualMachineTimeoutsOutputReference {
	var returns VirtualMachineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/virtual_machine azurerm_virtual_machine} Resource.
func NewVirtualMachine(scope constructs.Construct, id *string, config *VirtualMachineConfig) VirtualMachine {
	_init_.Initialize()

	if err := validateNewVirtualMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachine{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/virtual_machine azurerm_virtual_machine} Resource.
func NewVirtualMachine_Override(v VirtualMachine, scope constructs.Construct, id *string, config *VirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachine",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualMachine)SetAvailabilitySetId(val *string) {
	if err := j.validateSetAvailabilitySetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilitySetId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetDeleteDataDisksOnTermination(val interface{}) {
	if err := j.validateSetDeleteDataDisksOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteDataDisksOnTermination",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetDeleteOsDiskOnTermination(val interface{}) {
	if err := j.validateSetDeleteOsDiskOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOsDiskOnTermination",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetNetworkInterfaceIds(val *[]*string) {
	if err := j.validateSetNetworkInterfaceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceIds",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetPrimaryNetworkInterfaceId(val *string) {
	if err := j.validateSetPrimaryNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryNetworkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Generates CDKTF code for importing a VirtualMachine resource upon running "cdktf plan <stack-name>".
func VirtualMachine_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualMachine_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachine",
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
func VirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualMachine) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualMachine) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachine) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachine) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachine) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachine) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualMachine) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualMachine) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualMachine) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualMachine) PutAdditionalCapabilities(value *VirtualMachineAdditionalCapabilities) {
	if err := v.validatePutAdditionalCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAdditionalCapabilities",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutBootDiagnostics(value *VirtualMachineBootDiagnostics) {
	if err := v.validatePutBootDiagnosticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutIdentity(value *VirtualMachineIdentity) {
	if err := v.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIdentity",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutOsProfile(value *VirtualMachineOsProfile) {
	if err := v.validatePutOsProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOsProfile",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutOsProfileLinuxConfig(value *VirtualMachineOsProfileLinuxConfig) {
	if err := v.validatePutOsProfileLinuxConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOsProfileLinuxConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutOsProfileSecrets(value interface{}) {
	if err := v.validatePutOsProfileSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOsProfileSecrets",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutOsProfileWindowsConfig(value *VirtualMachineOsProfileWindowsConfig) {
	if err := v.validatePutOsProfileWindowsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOsProfileWindowsConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutPlan(value *VirtualMachinePlan) {
	if err := v.validatePutPlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPlan",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutStorageDataDisk(value interface{}) {
	if err := v.validatePutStorageDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageDataDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutStorageImageReference(value *VirtualMachineStorageImageReference) {
	if err := v.validatePutStorageImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageImageReference",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutStorageOsDisk(value *VirtualMachineStorageOsDisk) {
	if err := v.validatePutStorageOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageOsDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutTimeouts(value *VirtualMachineTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) ResetAdditionalCapabilities() {
	_jsii_.InvokeVoid(
		v,
		"resetAdditionalCapabilities",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetAvailabilitySetId() {
	_jsii_.InvokeVoid(
		v,
		"resetAvailabilitySetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		v,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetDeleteDataDisksOnTermination() {
	_jsii_.InvokeVoid(
		v,
		"resetDeleteDataDisksOnTermination",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetDeleteOsDiskOnTermination() {
	_jsii_.InvokeVoid(
		v,
		"resetDeleteOsDiskOnTermination",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetIdentity() {
	_jsii_.InvokeVoid(
		v,
		"resetIdentity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetLicenseType() {
	_jsii_.InvokeVoid(
		v,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetOsProfile() {
	_jsii_.InvokeVoid(
		v,
		"resetOsProfile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetOsProfileLinuxConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetOsProfileLinuxConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetOsProfileSecrets() {
	_jsii_.InvokeVoid(
		v,
		"resetOsProfileSecrets",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetOsProfileWindowsConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetOsProfileWindowsConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetPlan() {
	_jsii_.InvokeVoid(
		v,
		"resetPlan",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetPrimaryNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		v,
		"resetPrimaryNetworkInterfaceId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		v,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetStorageDataDisk() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageDataDisk",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetStorageImageReference() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageImageReference",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetZones() {
	_jsii_.InvokeVoid(
		v,
		"resetZones",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

