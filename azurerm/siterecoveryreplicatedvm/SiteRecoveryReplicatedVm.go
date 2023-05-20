package siterecoveryreplicatedvm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/siterecoveryreplicatedvm/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/site_recovery_replicated_vm azurerm_site_recovery_replicated_vm}.
type SiteRecoveryReplicatedVm interface {
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManagedDisk() SiteRecoveryReplicatedVmManagedDiskList
	ManagedDiskInput() interface{}
	MultiVmGroupName() *string
	SetMultiVmGroupName(val *string)
	MultiVmGroupNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() SiteRecoveryReplicatedVmNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
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
	RecoveryVaultName() *string
	SetRecoveryVaultName(val *string)
	RecoveryVaultNameInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SourceRecoveryFabricName() *string
	SetSourceRecoveryFabricName(val *string)
	SourceRecoveryFabricNameInput() *string
	SourceRecoveryProtectionContainerName() *string
	SetSourceRecoveryProtectionContainerName(val *string)
	SourceRecoveryProtectionContainerNameInput() *string
	SourceVmId() *string
	SetSourceVmId(val *string)
	SourceVmIdInput() *string
	TargetAvailabilitySetId() *string
	SetTargetAvailabilitySetId(val *string)
	TargetAvailabilitySetIdInput() *string
	TargetBootDiagnosticStorageAccountId() *string
	SetTargetBootDiagnosticStorageAccountId(val *string)
	TargetBootDiagnosticStorageAccountIdInput() *string
	TargetCapacityReservationGroupId() *string
	SetTargetCapacityReservationGroupId(val *string)
	TargetCapacityReservationGroupIdInput() *string
	TargetEdgeZone() *string
	SetTargetEdgeZone(val *string)
	TargetEdgeZoneInput() *string
	TargetNetworkId() *string
	SetTargetNetworkId(val *string)
	TargetNetworkIdInput() *string
	TargetProximityPlacementGroupId() *string
	SetTargetProximityPlacementGroupId(val *string)
	TargetProximityPlacementGroupIdInput() *string
	TargetRecoveryFabricId() *string
	SetTargetRecoveryFabricId(val *string)
	TargetRecoveryFabricIdInput() *string
	TargetRecoveryProtectionContainerId() *string
	SetTargetRecoveryProtectionContainerId(val *string)
	TargetRecoveryProtectionContainerIdInput() *string
	TargetResourceGroupId() *string
	SetTargetResourceGroupId(val *string)
	TargetResourceGroupIdInput() *string
	TargetVirtualMachineScaleSetId() *string
	SetTargetVirtualMachineScaleSetId(val *string)
	TargetVirtualMachineScaleSetIdInput() *string
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
	Timeouts() SiteRecoveryReplicatedVmTimeoutsOutputReference
	TimeoutsInput() interface{}
	UnmanagedDisk() SiteRecoveryReplicatedVmUnmanagedDiskList
	UnmanagedDiskInput() interface{}
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
	PutManagedDisk(value interface{})
	PutNetworkInterface(value interface{})
	PutTimeouts(value *SiteRecoveryReplicatedVmTimeouts)
	PutUnmanagedDisk(value interface{})
	ResetId()
	ResetManagedDisk()
	ResetMultiVmGroupName()
	ResetNetworkInterface()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTargetAvailabilitySetId()
	ResetTargetBootDiagnosticStorageAccountId()
	ResetTargetCapacityReservationGroupId()
	ResetTargetEdgeZone()
	ResetTargetNetworkId()
	ResetTargetProximityPlacementGroupId()
	ResetTargetVirtualMachineScaleSetId()
	ResetTargetZone()
	ResetTestNetworkId()
	ResetTimeouts()
	ResetUnmanagedDisk()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SiteRecoveryReplicatedVm
type jsiiProxy_SiteRecoveryReplicatedVm struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) ManagedDisk() SiteRecoveryReplicatedVmManagedDiskList {
	var returns SiteRecoveryReplicatedVmManagedDiskList
	_jsii_.Get(
		j,
		"managedDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) ManagedDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) MultiVmGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiVmGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) MultiVmGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiVmGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) NetworkInterface() SiteRecoveryReplicatedVmNetworkInterfaceList {
	var returns SiteRecoveryReplicatedVmNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) RecoveryReplicationPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryReplicationPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) RecoveryReplicationPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryReplicationPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) RecoveryVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) RecoveryVaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) SourceRecoveryFabricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRecoveryFabricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) SourceRecoveryFabricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRecoveryFabricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) SourceRecoveryProtectionContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRecoveryProtectionContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) SourceRecoveryProtectionContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRecoveryProtectionContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) SourceVmId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVmId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) SourceVmIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVmIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetAvailabilitySetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAvailabilitySetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetAvailabilitySetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAvailabilitySetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetBootDiagnosticStorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBootDiagnosticStorageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetBootDiagnosticStorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBootDiagnosticStorageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetCapacityReservationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityReservationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetCapacityReservationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityReservationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetEdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEdgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetEdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEdgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetRecoveryFabricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRecoveryFabricId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetRecoveryFabricIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRecoveryFabricIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetRecoveryProtectionContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRecoveryProtectionContainerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetRecoveryProtectionContainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRecoveryProtectionContainerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetResourceGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetResourceGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetVirtualMachineScaleSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVirtualMachineScaleSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetVirtualMachineScaleSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVirtualMachineScaleSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TargetZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TestNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TestNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) Timeouts() SiteRecoveryReplicatedVmTimeoutsOutputReference {
	var returns SiteRecoveryReplicatedVmTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) UnmanagedDisk() SiteRecoveryReplicatedVmUnmanagedDiskList {
	var returns SiteRecoveryReplicatedVmUnmanagedDiskList
	_jsii_.Get(
		j,
		"unmanagedDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm) UnmanagedDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unmanagedDiskInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/site_recovery_replicated_vm azurerm_site_recovery_replicated_vm} Resource.
func NewSiteRecoveryReplicatedVm(scope constructs.Construct, id *string, config *SiteRecoveryReplicatedVmConfig) SiteRecoveryReplicatedVm {
	_init_.Initialize()

	if err := validateNewSiteRecoveryReplicatedVmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SiteRecoveryReplicatedVm{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/site_recovery_replicated_vm azurerm_site_recovery_replicated_vm} Resource.
func NewSiteRecoveryReplicatedVm_Override(s SiteRecoveryReplicatedVm, scope constructs.Construct, id *string, config *SiteRecoveryReplicatedVmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVm",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetMultiVmGroupName(val *string) {
	if err := j.validateSetMultiVmGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiVmGroupName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetRecoveryReplicationPolicyId(val *string) {
	if err := j.validateSetRecoveryReplicationPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryReplicationPolicyId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetRecoveryVaultName(val *string) {
	if err := j.validateSetRecoveryVaultNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryVaultName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetSourceRecoveryFabricName(val *string) {
	if err := j.validateSetSourceRecoveryFabricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRecoveryFabricName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetSourceRecoveryProtectionContainerName(val *string) {
	if err := j.validateSetSourceRecoveryProtectionContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRecoveryProtectionContainerName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetSourceVmId(val *string) {
	if err := j.validateSetSourceVmIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVmId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetAvailabilitySetId(val *string) {
	if err := j.validateSetTargetAvailabilitySetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAvailabilitySetId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetBootDiagnosticStorageAccountId(val *string) {
	if err := j.validateSetTargetBootDiagnosticStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetBootDiagnosticStorageAccountId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetCapacityReservationGroupId(val *string) {
	if err := j.validateSetTargetCapacityReservationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetEdgeZone(val *string) {
	if err := j.validateSetTargetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetEdgeZone",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetNetworkId(val *string) {
	if err := j.validateSetTargetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetNetworkId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetProximityPlacementGroupId(val *string) {
	if err := j.validateSetTargetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetProximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetRecoveryFabricId(val *string) {
	if err := j.validateSetTargetRecoveryFabricIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRecoveryFabricId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetRecoveryProtectionContainerId(val *string) {
	if err := j.validateSetTargetRecoveryProtectionContainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRecoveryProtectionContainerId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetResourceGroupId(val *string) {
	if err := j.validateSetTargetResourceGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceGroupId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetVirtualMachineScaleSetId(val *string) {
	if err := j.validateSetTargetVirtualMachineScaleSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVirtualMachineScaleSetId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTargetZone(val *string) {
	if err := j.validateSetTargetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetZone",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVm)SetTestNetworkId(val *string) {
	if err := j.validateSetTestNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testNetworkId",
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
func SiteRecoveryReplicatedVm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryReplicatedVm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SiteRecoveryReplicatedVm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryReplicatedVm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SiteRecoveryReplicatedVm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryReplicatedVm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SiteRecoveryReplicatedVm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVm) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) PutManagedDisk(value interface{}) {
	if err := s.validatePutManagedDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putManagedDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) PutNetworkInterface(value interface{}) {
	if err := s.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) PutTimeouts(value *SiteRecoveryReplicatedVmTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) PutUnmanagedDisk(value interface{}) {
	if err := s.validatePutUnmanagedDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUnmanagedDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetManagedDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetMultiVmGroupName() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiVmGroupName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTargetAvailabilitySetId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetAvailabilitySetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTargetBootDiagnosticStorageAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetBootDiagnosticStorageAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTargetCapacityReservationGroupId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetCapacityReservationGroupId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTargetEdgeZone() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetEdgeZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTargetNetworkId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetNetworkId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTargetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTargetVirtualMachineScaleSetId() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetVirtualMachineScaleSetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTargetZone() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTestNetworkId() {
	_jsii_.InvokeVoid(
		s,
		"resetTestNetworkId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ResetUnmanagedDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetUnmanagedDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

