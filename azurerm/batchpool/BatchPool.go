// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/batchpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/batch_pool azurerm_batch_pool}.
type BatchPool interface {
	cdktf.TerraformResource
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	AutoScale() BatchPoolAutoScaleOutputReference
	AutoScaleInput() *BatchPoolAutoScale
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() BatchPoolCertificateList
	CertificateInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerConfiguration() BatchPoolContainerConfigurationOutputReference
	ContainerConfigurationInput() *BatchPoolContainerConfiguration
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataDisks() BatchPoolDataDisksList
	DataDisksInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskEncryption() BatchPoolDiskEncryptionList
	DiskEncryptionInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Extensions() BatchPoolExtensionsList
	ExtensionsInput() interface{}
	FixedScale() BatchPoolFixedScaleOutputReference
	FixedScaleInput() *BatchPoolFixedScale
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
	Identity() BatchPoolIdentityOutputReference
	IdentityInput() *BatchPoolIdentity
	IdInput() *string
	InterNodeCommunication() *string
	SetInterNodeCommunication(val *string)
	InterNodeCommunicationInput() *string
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxTasksPerNode() *float64
	SetMaxTasksPerNode(val *float64)
	MaxTasksPerNodeInput() *float64
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Mount() BatchPoolMountList
	MountInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfiguration() BatchPoolNetworkConfigurationOutputReference
	NetworkConfigurationInput() *BatchPoolNetworkConfiguration
	// The tree node.
	Node() constructs.Node
	NodeAgentSkuId() *string
	SetNodeAgentSkuId(val *string)
	NodeAgentSkuIdInput() *string
	NodePlacement() BatchPoolNodePlacementList
	NodePlacementInput() interface{}
	OsDiskPlacement() *string
	SetOsDiskPlacement(val *string)
	OsDiskPlacementInput() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	StartTask() BatchPoolStartTaskOutputReference
	StartTaskInput() *BatchPoolStartTask
	StopPendingResizeOperation() interface{}
	SetStopPendingResizeOperation(val interface{})
	StopPendingResizeOperationInput() interface{}
	StorageImageReference() BatchPoolStorageImageReferenceOutputReference
	StorageImageReferenceInput() *BatchPoolStorageImageReference
	TargetNodeCommunicationMode() *string
	SetTargetNodeCommunicationMode(val *string)
	TargetNodeCommunicationModeInput() *string
	TaskSchedulingPolicy() BatchPoolTaskSchedulingPolicyList
	TaskSchedulingPolicyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BatchPoolTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserAccounts() BatchPoolUserAccountsList
	UserAccountsInput() interface{}
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
	Windows() BatchPoolWindowsList
	WindowsInput() interface{}
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
	PutAutoScale(value *BatchPoolAutoScale)
	PutCertificate(value interface{})
	PutContainerConfiguration(value *BatchPoolContainerConfiguration)
	PutDataDisks(value interface{})
	PutDiskEncryption(value interface{})
	PutExtensions(value interface{})
	PutFixedScale(value *BatchPoolFixedScale)
	PutIdentity(value *BatchPoolIdentity)
	PutMount(value interface{})
	PutNetworkConfiguration(value *BatchPoolNetworkConfiguration)
	PutNodePlacement(value interface{})
	PutStartTask(value *BatchPoolStartTask)
	PutStorageImageReference(value *BatchPoolStorageImageReference)
	PutTaskSchedulingPolicy(value interface{})
	PutTimeouts(value *BatchPoolTimeouts)
	PutUserAccounts(value interface{})
	PutWindows(value interface{})
	ResetAutoScale()
	ResetCertificate()
	ResetContainerConfiguration()
	ResetDataDisks()
	ResetDiskEncryption()
	ResetDisplayName()
	ResetExtensions()
	ResetFixedScale()
	ResetId()
	ResetIdentity()
	ResetInterNodeCommunication()
	ResetLicenseType()
	ResetMaxTasksPerNode()
	ResetMetadata()
	ResetMount()
	ResetNetworkConfiguration()
	ResetNodePlacement()
	ResetOsDiskPlacement()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartTask()
	ResetStopPendingResizeOperation()
	ResetTargetNodeCommunicationMode()
	ResetTaskSchedulingPolicy()
	ResetTimeouts()
	ResetUserAccounts()
	ResetWindows()
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

// The jsii proxy struct for BatchPool
type jsiiProxy_BatchPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BatchPool) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) AutoScale() BatchPoolAutoScaleOutputReference {
	var returns BatchPoolAutoScaleOutputReference
	_jsii_.Get(
		j,
		"autoScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) AutoScaleInput() *BatchPoolAutoScale {
	var returns *BatchPoolAutoScale
	_jsii_.Get(
		j,
		"autoScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Certificate() BatchPoolCertificateList {
	var returns BatchPoolCertificateList
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) ContainerConfiguration() BatchPoolContainerConfigurationOutputReference {
	var returns BatchPoolContainerConfigurationOutputReference
	_jsii_.Get(
		j,
		"containerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) ContainerConfigurationInput() *BatchPoolContainerConfiguration {
	var returns *BatchPoolContainerConfiguration
	_jsii_.Get(
		j,
		"containerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) DataDisks() BatchPoolDataDisksList {
	var returns BatchPoolDataDisksList
	_jsii_.Get(
		j,
		"dataDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) DataDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) DiskEncryption() BatchPoolDiskEncryptionList {
	var returns BatchPoolDiskEncryptionList
	_jsii_.Get(
		j,
		"diskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) DiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Extensions() BatchPoolExtensionsList {
	var returns BatchPoolExtensionsList
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) ExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) FixedScale() BatchPoolFixedScaleOutputReference {
	var returns BatchPoolFixedScaleOutputReference
	_jsii_.Get(
		j,
		"fixedScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) FixedScaleInput() *BatchPoolFixedScale {
	var returns *BatchPoolFixedScale
	_jsii_.Get(
		j,
		"fixedScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Identity() BatchPoolIdentityOutputReference {
	var returns BatchPoolIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) IdentityInput() *BatchPoolIdentity {
	var returns *BatchPoolIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) InterNodeCommunication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interNodeCommunication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) InterNodeCommunicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interNodeCommunicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) MaxTasksPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTasksPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) MaxTasksPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTasksPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Mount() BatchPoolMountList {
	var returns BatchPoolMountList
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) MountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) NetworkConfiguration() BatchPoolNetworkConfigurationOutputReference {
	var returns BatchPoolNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) NetworkConfigurationInput() *BatchPoolNetworkConfiguration {
	var returns *BatchPoolNetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) NodeAgentSkuId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAgentSkuId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) NodeAgentSkuIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAgentSkuIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) NodePlacement() BatchPoolNodePlacementList {
	var returns BatchPoolNodePlacementList
	_jsii_.Get(
		j,
		"nodePlacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) NodePlacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodePlacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) OsDiskPlacement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskPlacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) OsDiskPlacementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskPlacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) StartTask() BatchPoolStartTaskOutputReference {
	var returns BatchPoolStartTaskOutputReference
	_jsii_.Get(
		j,
		"startTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) StartTaskInput() *BatchPoolStartTask {
	var returns *BatchPoolStartTask
	_jsii_.Get(
		j,
		"startTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) StopPendingResizeOperation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopPendingResizeOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) StopPendingResizeOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopPendingResizeOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) StorageImageReference() BatchPoolStorageImageReferenceOutputReference {
	var returns BatchPoolStorageImageReferenceOutputReference
	_jsii_.Get(
		j,
		"storageImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) StorageImageReferenceInput() *BatchPoolStorageImageReference {
	var returns *BatchPoolStorageImageReference
	_jsii_.Get(
		j,
		"storageImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) TargetNodeCommunicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNodeCommunicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) TargetNodeCommunicationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNodeCommunicationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) TaskSchedulingPolicy() BatchPoolTaskSchedulingPolicyList {
	var returns BatchPoolTaskSchedulingPolicyList
	_jsii_.Get(
		j,
		"taskSchedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) TaskSchedulingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskSchedulingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Timeouts() BatchPoolTimeoutsOutputReference {
	var returns BatchPoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) UserAccounts() BatchPoolUserAccountsList {
	var returns BatchPoolUserAccountsList
	_jsii_.Get(
		j,
		"userAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) UserAccountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) Windows() BatchPoolWindowsList {
	var returns BatchPoolWindowsList
	_jsii_.Get(
		j,
		"windows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPool) WindowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/batch_pool azurerm_batch_pool} Resource.
func NewBatchPool(scope constructs.Construct, id *string, config *BatchPoolConfig) BatchPool {
	_init_.Initialize()

	if err := validateNewBatchPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchPool{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/batch_pool azurerm_batch_pool} Resource.
func NewBatchPool_Override(b BatchPool, scope constructs.Construct, id *string, config *BatchPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BatchPool)SetAccountName(val *string) {
	if err := j.validateSetAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetInterNodeCommunication(val *string) {
	if err := j.validateSetInterNodeCommunicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interNodeCommunication",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetMaxTasksPerNode(val *float64) {
	if err := j.validateSetMaxTasksPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTasksPerNode",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetNodeAgentSkuId(val *string) {
	if err := j.validateSetNodeAgentSkuIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeAgentSkuId",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetOsDiskPlacement(val *string) {
	if err := j.validateSetOsDiskPlacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskPlacement",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetStopPendingResizeOperation(val interface{}) {
	if err := j.validateSetStopPendingResizeOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopPendingResizeOperation",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetTargetNodeCommunicationMode(val *string) {
	if err := j.validateSetTargetNodeCommunicationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetNodeCommunicationMode",
		val,
	)
}

func (j *jsiiProxy_BatchPool)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

// Generates CDKTF code for importing a BatchPool resource upon running "cdktf plan <stack-name>".
func BatchPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBatchPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
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
func BatchPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BatchPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BatchPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BatchPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BatchPool) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BatchPool) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BatchPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BatchPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BatchPool) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BatchPool) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BatchPool) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BatchPool) PutAutoScale(value *BatchPoolAutoScale) {
	if err := b.validatePutAutoScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAutoScale",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutCertificate(value interface{}) {
	if err := b.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCertificate",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutContainerConfiguration(value *BatchPoolContainerConfiguration) {
	if err := b.validatePutContainerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putContainerConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutDataDisks(value interface{}) {
	if err := b.validatePutDataDisksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDataDisks",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutDiskEncryption(value interface{}) {
	if err := b.validatePutDiskEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDiskEncryption",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutExtensions(value interface{}) {
	if err := b.validatePutExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExtensions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutFixedScale(value *BatchPoolFixedScale) {
	if err := b.validatePutFixedScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFixedScale",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutIdentity(value *BatchPoolIdentity) {
	if err := b.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putIdentity",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutMount(value interface{}) {
	if err := b.validatePutMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMount",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutNetworkConfiguration(value *BatchPoolNetworkConfiguration) {
	if err := b.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutNodePlacement(value interface{}) {
	if err := b.validatePutNodePlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNodePlacement",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutStartTask(value *BatchPoolStartTask) {
	if err := b.validatePutStartTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStartTask",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutStorageImageReference(value *BatchPoolStorageImageReference) {
	if err := b.validatePutStorageImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStorageImageReference",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutTaskSchedulingPolicy(value interface{}) {
	if err := b.validatePutTaskSchedulingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTaskSchedulingPolicy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutTimeouts(value *BatchPoolTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutUserAccounts(value interface{}) {
	if err := b.validatePutUserAccountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUserAccounts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutWindows(value interface{}) {
	if err := b.validatePutWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putWindows",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) ResetAutoScale() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoScale",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetCertificate() {
	_jsii_.InvokeVoid(
		b,
		"resetCertificate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetContainerConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetContainerConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetDataDisks() {
	_jsii_.InvokeVoid(
		b,
		"resetDataDisks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetDiskEncryption() {
	_jsii_.InvokeVoid(
		b,
		"resetDiskEncryption",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetDisplayName() {
	_jsii_.InvokeVoid(
		b,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetExtensions() {
	_jsii_.InvokeVoid(
		b,
		"resetExtensions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetFixedScale() {
	_jsii_.InvokeVoid(
		b,
		"resetFixedScale",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetIdentity() {
	_jsii_.InvokeVoid(
		b,
		"resetIdentity",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetInterNodeCommunication() {
	_jsii_.InvokeVoid(
		b,
		"resetInterNodeCommunication",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetLicenseType() {
	_jsii_.InvokeVoid(
		b,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetMaxTasksPerNode() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxTasksPerNode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetMetadata() {
	_jsii_.InvokeVoid(
		b,
		"resetMetadata",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetMount() {
	_jsii_.InvokeVoid(
		b,
		"resetMount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetNodePlacement() {
	_jsii_.InvokeVoid(
		b,
		"resetNodePlacement",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetOsDiskPlacement() {
	_jsii_.InvokeVoid(
		b,
		"resetOsDiskPlacement",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetStartTask() {
	_jsii_.InvokeVoid(
		b,
		"resetStartTask",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetStopPendingResizeOperation() {
	_jsii_.InvokeVoid(
		b,
		"resetStopPendingResizeOperation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetTargetNodeCommunicationMode() {
	_jsii_.InvokeVoid(
		b,
		"resetTargetNodeCommunicationMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetTaskSchedulingPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetTaskSchedulingPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetUserAccounts() {
	_jsii_.InvokeVoid(
		b,
		"resetUserAccounts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) ResetWindows() {
	_jsii_.InvokeVoid(
		b,
		"resetWindows",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

