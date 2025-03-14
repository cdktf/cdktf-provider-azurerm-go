// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryvmwarereplicatedvm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/siterecoveryvmwarereplicatedvm/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/site_recovery_vmware_replicated_vm azurerm_site_recovery_vmware_replicated_vm}.
type SiteRecoveryVmwareReplicatedVm interface {
	cdktf.TerraformResource
	ApplianceName() *string
	SetApplianceName(val *string)
	ApplianceNameInput() *string
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
	DefaultLogStorageAccountId() *string
	SetDefaultLogStorageAccountId(val *string)
	DefaultLogStorageAccountIdInput() *string
	DefaultRecoveryDiskType() *string
	SetDefaultRecoveryDiskType(val *string)
	DefaultRecoveryDiskTypeInput() *string
	DefaultTargetDiskEncryptionSetId() *string
	SetDefaultTargetDiskEncryptionSetId(val *string)
	DefaultTargetDiskEncryptionSetIdInput() *string
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
	IdInput() *string
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManagedDisk() SiteRecoveryVmwareReplicatedVmManagedDiskList
	ManagedDiskInput() interface{}
	MultiVmGroupName() *string
	SetMultiVmGroupName(val *string)
	MultiVmGroupNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() SiteRecoveryVmwareReplicatedVmNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	PhysicalServerCredentialName() *string
	SetPhysicalServerCredentialName(val *string)
	PhysicalServerCredentialNameInput() *string
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
	RecoveryReplicationPolicyId() *string
	SetRecoveryReplicationPolicyId(val *string)
	RecoveryReplicationPolicyIdInput() *string
	RecoveryVaultId() *string
	SetRecoveryVaultId(val *string)
	RecoveryVaultIdInput() *string
	SourceVmName() *string
	SetSourceVmName(val *string)
	SourceVmNameInput() *string
	TargetAvailabilitySetId() *string
	SetTargetAvailabilitySetId(val *string)
	TargetAvailabilitySetIdInput() *string
	TargetBootDiagnosticsStorageAccountId() *string
	SetTargetBootDiagnosticsStorageAccountId(val *string)
	TargetBootDiagnosticsStorageAccountIdInput() *string
	TargetNetworkId() *string
	SetTargetNetworkId(val *string)
	TargetNetworkIdInput() *string
	TargetProximityPlacementGroupId() *string
	SetTargetProximityPlacementGroupId(val *string)
	TargetProximityPlacementGroupIdInput() *string
	TargetResourceGroupId() *string
	SetTargetResourceGroupId(val *string)
	TargetResourceGroupIdInput() *string
	TargetVmName() *string
	SetTargetVmName(val *string)
	TargetVmNameInput() *string
	TargetVmSize() *string
	SetTargetVmSize(val *string)
	TargetVmSizeInput() *string
	TargetZone() *string
	SetTargetZone(val *string)
	TargetZoneInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TestNetworkId() *string
	SetTestNetworkId(val *string)
	TestNetworkIdInput() *string
	Timeouts() SiteRecoveryVmwareReplicatedVmTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutManagedDisk(value interface{})
	PutNetworkInterface(value interface{})
	PutTimeouts(value *SiteRecoveryVmwareReplicatedVmTimeouts)
	ResetDefaultLogStorageAccountId()
	ResetDefaultRecoveryDiskType()
	ResetDefaultTargetDiskEncryptionSetId()
	ResetId()
	ResetLicenseType()
	ResetManagedDisk()
	ResetMultiVmGroupName()
	ResetNetworkInterface()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTargetAvailabilitySetId()
	ResetTargetBootDiagnosticsStorageAccountId()
	ResetTargetNetworkId()
	ResetTargetProximityPlacementGroupId()
	ResetTargetVmSize()
	ResetTargetZone()
	ResetTestNetworkId()
	ResetTimeouts()
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

// The jsii proxy struct for SiteRecoveryVmwareReplicatedVm
type jsiiProxy_SiteRecoveryVmwareReplicatedVm struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ApplianceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applianceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ApplianceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applianceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) DefaultLogStorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLogStorageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) DefaultLogStorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLogStorageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) DefaultRecoveryDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRecoveryDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) DefaultRecoveryDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRecoveryDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) DefaultTargetDiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTargetDiskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) DefaultTargetDiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTargetDiskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ManagedDisk() SiteRecoveryVmwareReplicatedVmManagedDiskList {
	var returns SiteRecoveryVmwareReplicatedVmManagedDiskList
	_jsii_.Get(
		j,
		"managedDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ManagedDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) MultiVmGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiVmGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) MultiVmGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiVmGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) NetworkInterface() SiteRecoveryVmwareReplicatedVmNetworkInterfaceList {
	var returns SiteRecoveryVmwareReplicatedVmNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) PhysicalServerCredentialName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalServerCredentialName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) PhysicalServerCredentialNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalServerCredentialNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) RecoveryReplicationPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryReplicationPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) RecoveryReplicationPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryReplicationPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) RecoveryVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) RecoveryVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) SourceVmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) SourceVmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetAvailabilitySetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAvailabilitySetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetAvailabilitySetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAvailabilitySetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetBootDiagnosticsStorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBootDiagnosticsStorageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetBootDiagnosticsStorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBootDiagnosticsStorageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetResourceGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetResourceGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetVmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetVmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetVmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetVmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVmSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TargetZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TestNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TestNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) Timeouts() SiteRecoveryVmwareReplicatedVmTimeoutsOutputReference {
	var returns SiteRecoveryVmwareReplicatedVmTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/site_recovery_vmware_replicated_vm azurerm_site_recovery_vmware_replicated_vm} Resource.
func NewSiteRecoveryVmwareReplicatedVm(scope constructs.Construct, id *string, config *SiteRecoveryVmwareReplicatedVmConfig) SiteRecoveryVmwareReplicatedVm {
	_init_.Initialize()

	if err := validateNewSiteRecoveryVmwareReplicatedVmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SiteRecoveryVmwareReplicatedVm{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryVmwareReplicatedVm.SiteRecoveryVmwareReplicatedVm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/site_recovery_vmware_replicated_vm azurerm_site_recovery_vmware_replicated_vm} Resource.
func NewSiteRecoveryVmwareReplicatedVm_Override(s SiteRecoveryVmwareReplicatedVm, scope constructs.Construct, id *string, config *SiteRecoveryVmwareReplicatedVmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryVmwareReplicatedVm.SiteRecoveryVmwareReplicatedVm",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetApplianceName(val *string) {
	if err := j.validateSetApplianceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applianceName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetDefaultLogStorageAccountId(val *string) {
	if err := j.validateSetDefaultLogStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLogStorageAccountId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetDefaultRecoveryDiskType(val *string) {
	if err := j.validateSetDefaultRecoveryDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRecoveryDiskType",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetDefaultTargetDiskEncryptionSetId(val *string) {
	if err := j.validateSetDefaultTargetDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTargetDiskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetMultiVmGroupName(val *string) {
	if err := j.validateSetMultiVmGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiVmGroupName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetPhysicalServerCredentialName(val *string) {
	if err := j.validateSetPhysicalServerCredentialNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalServerCredentialName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetRecoveryReplicationPolicyId(val *string) {
	if err := j.validateSetRecoveryReplicationPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryReplicationPolicyId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetRecoveryVaultId(val *string) {
	if err := j.validateSetRecoveryVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryVaultId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetSourceVmName(val *string) {
	if err := j.validateSetSourceVmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVmName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTargetAvailabilitySetId(val *string) {
	if err := j.validateSetTargetAvailabilitySetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAvailabilitySetId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTargetBootDiagnosticsStorageAccountId(val *string) {
	if err := j.validateSetTargetBootDiagnosticsStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetBootDiagnosticsStorageAccountId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTargetNetworkId(val *string) {
	if err := j.validateSetTargetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetNetworkId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTargetProximityPlacementGroupId(val *string) {
	if err := j.validateSetTargetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetProximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTargetResourceGroupId(val *string) {
	if err := j.validateSetTargetResourceGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceGroupId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTargetVmName(val *string) {
	if err := j.validateSetTargetVmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVmName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTargetVmSize(val *string) {
	if err := j.validateSetTargetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVmSize",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTargetZone(val *string) {
	if err := j.validateSetTargetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetZone",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryVmwareReplicatedVm)SetTestNetworkId(val *string) {
	if err := j.validateSetTestNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testNetworkId",
		val,
	)
}

// Generates CDKTF code for importing a SiteRecoveryVmwareReplicatedVm resource upon running "cdktf plan <stack-name>".
func SiteRecoveryVmwareReplicatedVm_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSiteRecoveryVmwareReplicatedVm_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryVmwareReplicatedVm.SiteRecoveryVmwareReplicatedVm",
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
func SiteRecoveryVmwareReplicatedVm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryVmwareReplicatedVm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryVmwareReplicatedVm.SiteRecoveryVmwareReplicatedVm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SiteRecoveryVmwareReplicatedVm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryVmwareReplicatedVm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryVmwareReplicatedVm.SiteRecoveryVmwareReplicatedVm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SiteRecoveryVmwareReplicatedVm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryVmwareReplicatedVm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryVmwareReplicatedVm.SiteRecoveryVmwareReplicatedVm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SiteRecoveryVmwareReplicatedVm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.siteRecoveryVmwareReplicatedVm.SiteRecoveryVmwareReplicatedVm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) PutManagedDisk(value interface{}) {
	if err := s.validatePutManagedDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putManagedDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) PutNetworkInterface(value interface{}) {
	if err := s.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) PutTimeouts(value *SiteRecoveryVmwareReplicatedVmTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetDefaultLogStorageAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultLogStorageAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetDefaultRecoveryDiskType() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultRecoveryDiskType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetDefaultTargetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultTargetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetLicenseType() {
	_jsii_.InvokeVoid(
		s,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetManagedDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetMultiVmGroupName() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiVmGroupName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetTargetAvailabilitySetId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetAvailabilitySetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetTargetBootDiagnosticsStorageAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetBootDiagnosticsStorageAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetTargetNetworkId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetNetworkId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetTargetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetTargetVmSize() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetVmSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetTargetZone() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetTestNetworkId() {
	_jsii_.InvokeVoid(
		s,
		"resetTestNetworkId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryVmwareReplicatedVm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

