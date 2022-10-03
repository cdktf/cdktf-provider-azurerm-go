package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/batchpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool azurerm_batch_pool}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	ResetTaskSchedulingPolicy()
	ResetTimeouts()
	ResetUserAccounts()
	ResetWindows()
	SynthesizeAttributes() *map[string]interface{}
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

func (j *jsiiProxy_BatchPool) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool azurerm_batch_pool} Resource.
func NewBatchPool(scope constructs.Construct, id *string, config *BatchPoolConfig) BatchPool {
	_init_.Initialize()

	j := jsiiProxy_BatchPool{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool azurerm_batch_pool} Resource.
func NewBatchPool_Override(b BatchPool, scope constructs.Construct, id *string, config *BatchPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BatchPool) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetInterNodeCommunication(val *string) {
	_jsii_.Set(
		j,
		"interNodeCommunication",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetLicenseType(val *string) {
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetMaxTasksPerNode(val *float64) {
	_jsii_.Set(
		j,
		"maxTasksPerNode",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetMetadata(val *map[string]*string) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetNodeAgentSkuId(val *string) {
	_jsii_.Set(
		j,
		"nodeAgentSkuId",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetOsDiskPlacement(val *string) {
	_jsii_.Set(
		j,
		"osDiskPlacement",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetStopPendingResizeOperation(val interface{}) {
	_jsii_.Set(
		j,
		"stopPendingResizeOperation",
		val,
	)
}

func (j *jsiiProxy_BatchPool) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
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
func BatchPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.batchPool.BatchPool",
		"isConstruct",
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

func (b *jsiiProxy_BatchPool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BatchPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BatchPool) PutAutoScale(value *BatchPoolAutoScale) {
	_jsii_.InvokeVoid(
		b,
		"putAutoScale",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutCertificate(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putCertificate",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutContainerConfiguration(value *BatchPoolContainerConfiguration) {
	_jsii_.InvokeVoid(
		b,
		"putContainerConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutDataDisks(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putDataDisks",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutDiskEncryption(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putDiskEncryption",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutExtensions(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putExtensions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutFixedScale(value *BatchPoolFixedScale) {
	_jsii_.InvokeVoid(
		b,
		"putFixedScale",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutIdentity(value *BatchPoolIdentity) {
	_jsii_.InvokeVoid(
		b,
		"putIdentity",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutMount(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putMount",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutNetworkConfiguration(value *BatchPoolNetworkConfiguration) {
	_jsii_.InvokeVoid(
		b,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutNodePlacement(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putNodePlacement",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutStartTask(value *BatchPoolStartTask) {
	_jsii_.InvokeVoid(
		b,
		"putStartTask",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutStorageImageReference(value *BatchPoolStorageImageReference) {
	_jsii_.InvokeVoid(
		b,
		"putStorageImageReference",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutTaskSchedulingPolicy(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putTaskSchedulingPolicy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutTimeouts(value *BatchPoolTimeouts) {
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutUserAccounts(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putUserAccounts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPool) PutWindows(value interface{}) {
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

type BatchPoolAutoScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#formula BatchPool#formula}.
	Formula *string `field:"required" json:"formula" yaml:"formula"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#evaluation_interval BatchPool#evaluation_interval}.
	EvaluationInterval *string `field:"optional" json:"evaluationInterval" yaml:"evaluationInterval"`
}

type BatchPoolAutoScaleOutputReference interface {
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
	EvaluationInterval() *string
	SetEvaluationInterval(val *string)
	EvaluationIntervalInput() *string
	Formula() *string
	SetFormula(val *string)
	FormulaInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BatchPoolAutoScale
	SetInternalValue(val *BatchPoolAutoScale)
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
	ResetEvaluationInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolAutoScaleOutputReference
type jsiiProxy_BatchPoolAutoScaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) EvaluationInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) EvaluationIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) Formula() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) FormulaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) InternalValue() *BatchPoolAutoScale {
	var returns *BatchPoolAutoScale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolAutoScaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolAutoScaleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolAutoScaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolAutoScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolAutoScaleOutputReference_Override(b BatchPoolAutoScaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolAutoScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) SetEvaluationInterval(val *string) {
	_jsii_.Set(
		j,
		"evaluationInterval",
		val,
	)
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) SetFormula(val *string) {
	_jsii_.Set(
		j,
		"formula",
		val,
	)
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) SetInternalValue(val *BatchPoolAutoScale) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolAutoScaleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) ResetEvaluationInterval() {
	_jsii_.InvokeVoid(
		b,
		"resetEvaluationInterval",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolAutoScaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#id BatchPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#store_location BatchPool#store_location}.
	StoreLocation *string `field:"required" json:"storeLocation" yaml:"storeLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#store_name BatchPool#store_name}.
	StoreName *string `field:"optional" json:"storeName" yaml:"storeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#visibility BatchPool#visibility}.
	Visibility *[]*string `field:"optional" json:"visibility" yaml:"visibility"`
}

type BatchPoolCertificateList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolCertificateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolCertificateList
type jsiiProxy_BatchPoolCertificateList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolCertificateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolCertificateList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolCertificateList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolCertificateList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolCertificateList_Override(b BatchPoolCertificateList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolCertificateList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolCertificateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateList) Get(index *float64) BatchPoolCertificateOutputReference {
	var returns BatchPoolCertificateOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolCertificateOutputReference interface {
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StoreLocation() *string
	SetStoreLocation(val *string)
	StoreLocationInput() *string
	StoreName() *string
	SetStoreName(val *string)
	StoreNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Visibility() *[]*string
	SetVisibility(val *[]*string)
	VisibilityInput() *[]*string
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
	ResetStoreName()
	ResetVisibility()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolCertificateOutputReference
type jsiiProxy_BatchPoolCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) StoreLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) StoreLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) StoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) StoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) Visibility() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) VisibilityInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}


func NewBatchPoolCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolCertificateOutputReference_Override(b BatchPoolCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetStoreLocation(val *string) {
	_jsii_.Set(
		j,
		"storeLocation",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetStoreName(val *string) {
	_jsii_.Set(
		j,
		"storeName",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolCertificateOutputReference) SetVisibility(val *[]*string) {
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) ResetStoreName() {
	_jsii_.InvokeVoid(
		b,
		"resetStoreName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) ResetVisibility() {
	_jsii_.InvokeVoid(
		b,
		"resetVisibility",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#account_name BatchPool#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#node_agent_sku_id BatchPool#node_agent_sku_id}.
	NodeAgentSkuId *string `field:"required" json:"nodeAgentSkuId" yaml:"nodeAgentSkuId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#resource_group_name BatchPool#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// storage_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#storage_image_reference BatchPool#storage_image_reference}
	StorageImageReference *BatchPoolStorageImageReference `field:"required" json:"storageImageReference" yaml:"storageImageReference"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#vm_size BatchPool#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#auto_scale BatchPool#auto_scale}
	AutoScale *BatchPoolAutoScale `field:"optional" json:"autoScale" yaml:"autoScale"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#certificate BatchPool#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#container_configuration BatchPool#container_configuration}
	ContainerConfiguration *BatchPoolContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
	// data_disks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#data_disks BatchPool#data_disks}
	DataDisks interface{} `field:"optional" json:"dataDisks" yaml:"dataDisks"`
	// disk_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#disk_encryption BatchPool#disk_encryption}
	DiskEncryption interface{} `field:"optional" json:"diskEncryption" yaml:"diskEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#display_name BatchPool#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// extensions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#extensions BatchPool#extensions}
	Extensions interface{} `field:"optional" json:"extensions" yaml:"extensions"`
	// fixed_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#fixed_scale BatchPool#fixed_scale}
	FixedScale *BatchPoolFixedScale `field:"optional" json:"fixedScale" yaml:"fixedScale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#id BatchPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#identity BatchPool#identity}
	Identity *BatchPoolIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#inter_node_communication BatchPool#inter_node_communication}.
	InterNodeCommunication *string `field:"optional" json:"interNodeCommunication" yaml:"interNodeCommunication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#license_type BatchPool#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#max_tasks_per_node BatchPool#max_tasks_per_node}.
	MaxTasksPerNode *float64 `field:"optional" json:"maxTasksPerNode" yaml:"maxTasksPerNode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#metadata BatchPool#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// mount block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#mount BatchPool#mount}
	Mount interface{} `field:"optional" json:"mount" yaml:"mount"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#network_configuration BatchPool#network_configuration}
	NetworkConfiguration *BatchPoolNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// node_placement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#node_placement BatchPool#node_placement}
	NodePlacement interface{} `field:"optional" json:"nodePlacement" yaml:"nodePlacement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#os_disk_placement BatchPool#os_disk_placement}.
	OsDiskPlacement *string `field:"optional" json:"osDiskPlacement" yaml:"osDiskPlacement"`
	// start_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#start_task BatchPool#start_task}
	StartTask *BatchPoolStartTask `field:"optional" json:"startTask" yaml:"startTask"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#stop_pending_resize_operation BatchPool#stop_pending_resize_operation}.
	StopPendingResizeOperation interface{} `field:"optional" json:"stopPendingResizeOperation" yaml:"stopPendingResizeOperation"`
	// task_scheduling_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#task_scheduling_policy BatchPool#task_scheduling_policy}
	TaskSchedulingPolicy interface{} `field:"optional" json:"taskSchedulingPolicy" yaml:"taskSchedulingPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#timeouts BatchPool#timeouts}
	Timeouts *BatchPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user_accounts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_accounts BatchPool#user_accounts}
	UserAccounts interface{} `field:"optional" json:"userAccounts" yaml:"userAccounts"`
	// windows block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#windows BatchPool#windows}
	Windows interface{} `field:"optional" json:"windows" yaml:"windows"`
}

type BatchPoolContainerConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#container_image_names BatchPool#container_image_names}.
	ContainerImageNames *[]*string `field:"optional" json:"containerImageNames" yaml:"containerImageNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#container_registries BatchPool#container_registries}.
	ContainerRegistries interface{} `field:"optional" json:"containerRegistries" yaml:"containerRegistries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#type BatchPool#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

type BatchPoolContainerConfigurationContainerRegistries struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#password BatchPool#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#registry_server BatchPool#registry_server}.
	RegistryServer *string `field:"optional" json:"registryServer" yaml:"registryServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_assigned_identity_id BatchPool#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_name BatchPool#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

type BatchPoolContainerConfigurationContainerRegistriesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolContainerConfigurationContainerRegistriesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolContainerConfigurationContainerRegistriesList
type jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolContainerConfigurationContainerRegistriesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolContainerConfigurationContainerRegistriesList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolContainerConfigurationContainerRegistriesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolContainerConfigurationContainerRegistriesList_Override(b BatchPoolContainerConfigurationContainerRegistriesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolContainerConfigurationContainerRegistriesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) Get(index *float64) BatchPoolContainerConfigurationContainerRegistriesOutputReference {
	var returns BatchPoolContainerConfigurationContainerRegistriesOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolContainerConfigurationContainerRegistriesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	RegistryServer() *string
	SetRegistryServer(val *string)
	RegistryServerInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAssignedIdentityId() *string
	SetUserAssignedIdentityId(val *string)
	UserAssignedIdentityIdInput() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
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
	ResetPassword()
	ResetRegistryServer()
	ResetUserAssignedIdentityId()
	ResetUserName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolContainerConfigurationContainerRegistriesOutputReference
type jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) RegistryServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) RegistryServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) UserAssignedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}


func NewBatchPoolContainerConfigurationContainerRegistriesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolContainerConfigurationContainerRegistriesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolContainerConfigurationContainerRegistriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolContainerConfigurationContainerRegistriesOutputReference_Override(b BatchPoolContainerConfigurationContainerRegistriesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolContainerConfigurationContainerRegistriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetRegistryServer(val *string) {
	_jsii_.Set(
		j,
		"registryServer",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetUserAssignedIdentityId(val *string) {
	_jsii_.Set(
		j,
		"userAssignedIdentityId",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		b,
		"resetPassword",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) ResetRegistryServer() {
	_jsii_.InvokeVoid(
		b,
		"resetRegistryServer",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) ResetUserAssignedIdentityId() {
	_jsii_.InvokeVoid(
		b,
		"resetUserAssignedIdentityId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) ResetUserName() {
	_jsii_.InvokeVoid(
		b,
		"resetUserName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationContainerRegistriesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolContainerConfigurationOutputReference interface {
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
	ContainerImageNames() *[]*string
	SetContainerImageNames(val *[]*string)
	ContainerImageNamesInput() *[]*string
	ContainerRegistries() BatchPoolContainerConfigurationContainerRegistriesList
	ContainerRegistriesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BatchPoolContainerConfiguration
	SetInternalValue(val *BatchPoolContainerConfiguration)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutContainerRegistries(value interface{})
	ResetContainerImageNames()
	ResetContainerRegistries()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolContainerConfigurationOutputReference
type jsiiProxy_BatchPoolContainerConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ContainerImageNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerImageNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ContainerImageNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerImageNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ContainerRegistries() BatchPoolContainerConfigurationContainerRegistriesList {
	var returns BatchPoolContainerConfigurationContainerRegistriesList
	_jsii_.Get(
		j,
		"containerRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ContainerRegistriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) InternalValue() *BatchPoolContainerConfiguration {
	var returns *BatchPoolContainerConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBatchPoolContainerConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolContainerConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolContainerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolContainerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolContainerConfigurationOutputReference_Override(b BatchPoolContainerConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolContainerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) SetContainerImageNames(val *[]*string) {
	_jsii_.Set(
		j,
		"containerImageNames",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) SetInternalValue(val *BatchPoolContainerConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolContainerConfigurationOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) PutContainerRegistries(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putContainerRegistries",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ResetContainerImageNames() {
	_jsii_.InvokeVoid(
		b,
		"resetContainerImageNames",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ResetContainerRegistries() {
	_jsii_.InvokeVoid(
		b,
		"resetContainerRegistries",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		b,
		"resetType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolContainerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolDataDisks struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#disk_size_gb BatchPool#disk_size_gb}.
	DiskSizeGb *float64 `field:"required" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#lun BatchPool#lun}.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#caching BatchPool#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#storage_account_type BatchPool#storage_account_type}.
	StorageAccountType *string `field:"optional" json:"storageAccountType" yaml:"storageAccountType"`
}

type BatchPoolDataDisksList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolDataDisksOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolDataDisksList
type jsiiProxy_BatchPoolDataDisksList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolDataDisksList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolDataDisksList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolDataDisksList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolDataDisksList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolDataDisksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolDataDisksList_Override(b BatchPoolDataDisksList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolDataDisksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolDataDisksList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksList) Get(index *float64) BatchPoolDataDisksOutputReference {
	var returns BatchPoolDataDisksOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolDataDisksOutputReference interface {
	cdktf.ComplexObject
	Caching() *string
	SetCaching(val *string)
	CachingInput() *string
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
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lun() *float64
	SetLun(val *float64)
	LunInput() *float64
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
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
	ResetCaching()
	ResetStorageAccountType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolDataDisksOutputReference
type jsiiProxy_BatchPoolDataDisksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) Lun() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) LunInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolDataDisksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolDataDisksOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolDataDisksOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolDataDisksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolDataDisksOutputReference_Override(b BatchPoolDataDisksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolDataDisksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetCaching(val *string) {
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetDiskSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetLun(val *float64) {
	_jsii_.Set(
		j,
		"lun",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetStorageAccountType(val *string) {
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDataDisksOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) ResetCaching() {
	_jsii_.InvokeVoid(
		b,
		"resetCaching",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) ResetStorageAccountType() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageAccountType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDataDisksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#disk_encryption_target BatchPool#disk_encryption_target}.
	DiskEncryptionTarget *string `field:"required" json:"diskEncryptionTarget" yaml:"diskEncryptionTarget"`
}

type BatchPoolDiskEncryptionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolDiskEncryptionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolDiskEncryptionList
type jsiiProxy_BatchPoolDiskEncryptionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolDiskEncryptionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolDiskEncryptionList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolDiskEncryptionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolDiskEncryptionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolDiskEncryptionList_Override(b BatchPoolDiskEncryptionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolDiskEncryptionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolDiskEncryptionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionList) Get(index *float64) BatchPoolDiskEncryptionOutputReference {
	var returns BatchPoolDiskEncryptionOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolDiskEncryptionOutputReference interface {
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
	DiskEncryptionTarget() *string
	SetDiskEncryptionTarget(val *string)
	DiskEncryptionTargetInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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

// The jsii proxy struct for BatchPoolDiskEncryptionOutputReference
type jsiiProxy_BatchPoolDiskEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) DiskEncryptionTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) DiskEncryptionTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolDiskEncryptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolDiskEncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolDiskEncryptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolDiskEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolDiskEncryptionOutputReference_Override(b BatchPoolDiskEncryptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolDiskEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) SetDiskEncryptionTarget(val *string) {
	_jsii_.Set(
		j,
		"diskEncryptionTarget",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolDiskEncryptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolDiskEncryptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolExtensions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#publisher BatchPool#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#type BatchPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#auto_upgrade_minor_version BatchPool#auto_upgrade_minor_version}.
	AutoUpgradeMinorVersion interface{} `field:"optional" json:"autoUpgradeMinorVersion" yaml:"autoUpgradeMinorVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#protected_settings BatchPool#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#provision_after_extensions BatchPool#provision_after_extensions}.
	ProvisionAfterExtensions *[]*string `field:"optional" json:"provisionAfterExtensions" yaml:"provisionAfterExtensions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#settings_json BatchPool#settings_json}.
	SettingsJson *string `field:"optional" json:"settingsJson" yaml:"settingsJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#type_handler_version BatchPool#type_handler_version}.
	TypeHandlerVersion *string `field:"optional" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
}

type BatchPoolExtensionsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolExtensionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolExtensionsList
type jsiiProxy_BatchPoolExtensionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolExtensionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolExtensionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolExtensionsList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolExtensionsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolExtensionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolExtensionsList_Override(b BatchPoolExtensionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolExtensionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolExtensionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsList) Get(index *float64) BatchPoolExtensionsOutputReference {
	var returns BatchPoolExtensionsOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolExtensionsOutputReference interface {
	cdktf.ComplexObject
	AutoUpgradeMinorVersion() interface{}
	SetAutoUpgradeMinorVersion(val interface{})
	AutoUpgradeMinorVersionInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	ProtectedSettings() *string
	SetProtectedSettings(val *string)
	ProtectedSettingsInput() *string
	ProvisionAfterExtensions() *[]*string
	SetProvisionAfterExtensions(val *[]*string)
	ProvisionAfterExtensionsInput() *[]*string
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
	SettingsJson() *string
	SetSettingsJson(val *string)
	SettingsJsonInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeHandlerVersion() *string
	SetTypeHandlerVersion(val *string)
	TypeHandlerVersionInput() *string
	TypeInput() *string
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
	ResetAutoUpgradeMinorVersion()
	ResetProtectedSettings()
	ResetProvisionAfterExtensions()
	ResetSettingsJson()
	ResetTypeHandlerVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolExtensionsOutputReference
type jsiiProxy_BatchPoolExtensionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) AutoUpgradeMinorVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeMinorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) AutoUpgradeMinorVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeMinorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) ProtectedSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) ProtectedSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) ProvisionAfterExtensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"provisionAfterExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) ProvisionAfterExtensionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"provisionAfterExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SettingsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) TypeHandlerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeHandlerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) TypeHandlerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeHandlerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBatchPoolExtensionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolExtensionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolExtensionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolExtensionsOutputReference_Override(b BatchPoolExtensionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetAutoUpgradeMinorVersion(val interface{}) {
	_jsii_.Set(
		j,
		"autoUpgradeMinorVersion",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetProtectedSettings(val *string) {
	_jsii_.Set(
		j,
		"protectedSettings",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetProvisionAfterExtensions(val *[]*string) {
	_jsii_.Set(
		j,
		"provisionAfterExtensions",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetPublisher(val *string) {
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetSettingsJson(val *string) {
	_jsii_.Set(
		j,
		"settingsJson",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_BatchPoolExtensionsOutputReference) SetTypeHandlerVersion(val *string) {
	_jsii_.Set(
		j,
		"typeHandlerVersion",
		val,
	)
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) ResetAutoUpgradeMinorVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoUpgradeMinorVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) ResetProtectedSettings() {
	_jsii_.InvokeVoid(
		b,
		"resetProtectedSettings",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) ResetProvisionAfterExtensions() {
	_jsii_.InvokeVoid(
		b,
		"resetProvisionAfterExtensions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) ResetSettingsJson() {
	_jsii_.InvokeVoid(
		b,
		"resetSettingsJson",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) ResetTypeHandlerVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetTypeHandlerVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolExtensionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolFixedScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#resize_timeout BatchPool#resize_timeout}.
	ResizeTimeout *string `field:"optional" json:"resizeTimeout" yaml:"resizeTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#target_dedicated_nodes BatchPool#target_dedicated_nodes}.
	TargetDedicatedNodes *float64 `field:"optional" json:"targetDedicatedNodes" yaml:"targetDedicatedNodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#target_low_priority_nodes BatchPool#target_low_priority_nodes}.
	TargetLowPriorityNodes *float64 `field:"optional" json:"targetLowPriorityNodes" yaml:"targetLowPriorityNodes"`
}

type BatchPoolFixedScaleOutputReference interface {
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
	InternalValue() *BatchPoolFixedScale
	SetInternalValue(val *BatchPoolFixedScale)
	ResizeTimeout() *string
	SetResizeTimeout(val *string)
	ResizeTimeoutInput() *string
	TargetDedicatedNodes() *float64
	SetTargetDedicatedNodes(val *float64)
	TargetDedicatedNodesInput() *float64
	TargetLowPriorityNodes() *float64
	SetTargetLowPriorityNodes(val *float64)
	TargetLowPriorityNodesInput() *float64
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
	ResetResizeTimeout()
	ResetTargetDedicatedNodes()
	ResetTargetLowPriorityNodes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolFixedScaleOutputReference
type jsiiProxy_BatchPoolFixedScaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) InternalValue() *BatchPoolFixedScale {
	var returns *BatchPoolFixedScale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) ResizeTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resizeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) ResizeTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resizeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TargetDedicatedNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetDedicatedNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TargetDedicatedNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetDedicatedNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TargetLowPriorityNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetLowPriorityNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TargetLowPriorityNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetLowPriorityNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolFixedScaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolFixedScaleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolFixedScaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolFixedScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolFixedScaleOutputReference_Override(b BatchPoolFixedScaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolFixedScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) SetInternalValue(val *BatchPoolFixedScale) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) SetResizeTimeout(val *string) {
	_jsii_.Set(
		j,
		"resizeTimeout",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) SetTargetDedicatedNodes(val *float64) {
	_jsii_.Set(
		j,
		"targetDedicatedNodes",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) SetTargetLowPriorityNodes(val *float64) {
	_jsii_.Set(
		j,
		"targetLowPriorityNodes",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ResetResizeTimeout() {
	_jsii_.InvokeVoid(
		b,
		"resetResizeTimeout",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ResetTargetDedicatedNodes() {
	_jsii_.InvokeVoid(
		b,
		"resetTargetDedicatedNodes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ResetTargetLowPriorityNodes() {
	_jsii_.InvokeVoid(
		b,
		"resetTargetLowPriorityNodes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#identity_ids BatchPool#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#type BatchPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

type BatchPoolIdentityOutputReference interface {
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
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *BatchPoolIdentity
	SetInternalValue(val *BatchPoolIdentity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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

// The jsii proxy struct for BatchPoolIdentityOutputReference
type jsiiProxy_BatchPoolIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) InternalValue() *BatchPoolIdentity {
	var returns *BatchPoolIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBatchPoolIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolIdentityOutputReference_Override(b BatchPoolIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) SetInternalValue(val *BatchPoolIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMount struct {
	// azure_blob_file_system block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#azure_blob_file_system BatchPool#azure_blob_file_system}
	AzureBlobFileSystem *BatchPoolMountAzureBlobFileSystem `field:"optional" json:"azureBlobFileSystem" yaml:"azureBlobFileSystem"`
	// azure_file_share block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#azure_file_share BatchPool#azure_file_share}
	AzureFileShare interface{} `field:"optional" json:"azureFileShare" yaml:"azureFileShare"`
	// cifs_mount block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#cifs_mount BatchPool#cifs_mount}
	CifsMount interface{} `field:"optional" json:"cifsMount" yaml:"cifsMount"`
	// nfs_mount block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#nfs_mount BatchPool#nfs_mount}
	NfsMount interface{} `field:"optional" json:"nfsMount" yaml:"nfsMount"`
}

type BatchPoolMountAzureBlobFileSystem struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#account_name BatchPool#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#container_name BatchPool#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#relative_mount_path BatchPool#relative_mount_path}.
	RelativeMountPath *string `field:"required" json:"relativeMountPath" yaml:"relativeMountPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#account_key BatchPool#account_key}.
	AccountKey *string `field:"optional" json:"accountKey" yaml:"accountKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#blobfuse_options BatchPool#blobfuse_options}.
	BlobfuseOptions *string `field:"optional" json:"blobfuseOptions" yaml:"blobfuseOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#identity_id BatchPool#identity_id}.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#sas_key BatchPool#sas_key}.
	SasKey *string `field:"optional" json:"sasKey" yaml:"sasKey"`
}

type BatchPoolMountAzureBlobFileSystemOutputReference interface {
	cdktf.ComplexObject
	AccountKey() *string
	SetAccountKey(val *string)
	AccountKeyInput() *string
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	BlobfuseOptions() *string
	SetBlobfuseOptions(val *string)
	BlobfuseOptionsInput() *string
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
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IdentityId() *string
	SetIdentityId(val *string)
	IdentityIdInput() *string
	InternalValue() *BatchPoolMountAzureBlobFileSystem
	SetInternalValue(val *BatchPoolMountAzureBlobFileSystem)
	RelativeMountPath() *string
	SetRelativeMountPath(val *string)
	RelativeMountPathInput() *string
	SasKey() *string
	SetSasKey(val *string)
	SasKeyInput() *string
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
	ResetAccountKey()
	ResetBlobfuseOptions()
	ResetIdentityId()
	ResetSasKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountAzureBlobFileSystemOutputReference
type jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) AccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) AccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) BlobfuseOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobfuseOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) BlobfuseOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobfuseOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) IdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) IdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) InternalValue() *BatchPoolMountAzureBlobFileSystem {
	var returns *BatchPoolMountAzureBlobFileSystem
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) RelativeMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) RelativeMountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SasKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SasKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolMountAzureBlobFileSystemOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolMountAzureBlobFileSystemOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountAzureBlobFileSystemOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolMountAzureBlobFileSystemOutputReference_Override(b BatchPoolMountAzureBlobFileSystemOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountAzureBlobFileSystemOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetAccountKey(val *string) {
	_jsii_.Set(
		j,
		"accountKey",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetBlobfuseOptions(val *string) {
	_jsii_.Set(
		j,
		"blobfuseOptions",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetContainerName(val *string) {
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetIdentityId(val *string) {
	_jsii_.Set(
		j,
		"identityId",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetInternalValue(val *BatchPoolMountAzureBlobFileSystem) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetRelativeMountPath(val *string) {
	_jsii_.Set(
		j,
		"relativeMountPath",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetSasKey(val *string) {
	_jsii_.Set(
		j,
		"sasKey",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ResetAccountKey() {
	_jsii_.InvokeVoid(
		b,
		"resetAccountKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ResetBlobfuseOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetBlobfuseOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ResetIdentityId() {
	_jsii_.InvokeVoid(
		b,
		"resetIdentityId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ResetSasKey() {
	_jsii_.InvokeVoid(
		b,
		"resetSasKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMountAzureFileShare struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#account_key BatchPool#account_key}.
	AccountKey *string `field:"required" json:"accountKey" yaml:"accountKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#account_name BatchPool#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#azure_file_url BatchPool#azure_file_url}.
	AzureFileUrl *string `field:"required" json:"azureFileUrl" yaml:"azureFileUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#relative_mount_path BatchPool#relative_mount_path}.
	RelativeMountPath *string `field:"required" json:"relativeMountPath" yaml:"relativeMountPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#mount_options BatchPool#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

type BatchPoolMountAzureFileShareList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolMountAzureFileShareOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountAzureFileShareList
type jsiiProxy_BatchPoolMountAzureFileShareList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolMountAzureFileShareList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolMountAzureFileShareList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountAzureFileShareList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountAzureFileShareList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountAzureFileShareList_Override(b BatchPoolMountAzureFileShareList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountAzureFileShareList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareList) Get(index *float64) BatchPoolMountAzureFileShareOutputReference {
	var returns BatchPoolMountAzureFileShareOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMountAzureFileShareOutputReference interface {
	cdktf.ComplexObject
	AccountKey() *string
	SetAccountKey(val *string)
	AccountKeyInput() *string
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	AzureFileUrl() *string
	SetAzureFileUrl(val *string)
	AzureFileUrlInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountOptions() *string
	SetMountOptions(val *string)
	MountOptionsInput() *string
	RelativeMountPath() *string
	SetRelativeMountPath(val *string)
	RelativeMountPathInput() *string
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
	ResetMountOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountAzureFileShareOutputReference
type jsiiProxy_BatchPoolMountAzureFileShareOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) AccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) AccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) AzureFileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureFileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) AzureFileUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureFileUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) MountOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) MountOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) RelativeMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) RelativeMountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolMountAzureFileShareOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolMountAzureFileShareOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountAzureFileShareOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountAzureFileShareOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountAzureFileShareOutputReference_Override(b BatchPoolMountAzureFileShareOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountAzureFileShareOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetAccountKey(val *string) {
	_jsii_.Set(
		j,
		"accountKey",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetAzureFileUrl(val *string) {
	_jsii_.Set(
		j,
		"azureFileUrl",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetMountOptions(val *string) {
	_jsii_.Set(
		j,
		"mountOptions",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetRelativeMountPath(val *string) {
	_jsii_.Set(
		j,
		"relativeMountPath",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) ResetMountOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMountCifsMount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#password BatchPool#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#relative_mount_path BatchPool#relative_mount_path}.
	RelativeMountPath *string `field:"required" json:"relativeMountPath" yaml:"relativeMountPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#source BatchPool#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_name BatchPool#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#mount_options BatchPool#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

type BatchPoolMountCifsMountList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolMountCifsMountOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountCifsMountList
type jsiiProxy_BatchPoolMountCifsMountList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolMountCifsMountList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolMountCifsMountList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountCifsMountList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountCifsMountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountCifsMountList_Override(b BatchPoolMountCifsMountList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountCifsMountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountCifsMountList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountList) Get(index *float64) BatchPoolMountCifsMountOutputReference {
	var returns BatchPoolMountCifsMountOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMountCifsMountOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountOptions() *string
	SetMountOptions(val *string)
	MountOptionsInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	RelativeMountPath() *string
	SetRelativeMountPath(val *string)
	RelativeMountPathInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
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
	ResetMountOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountCifsMountOutputReference
type jsiiProxy_BatchPoolMountCifsMountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) MountOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) MountOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) RelativeMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) RelativeMountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}


func NewBatchPoolMountCifsMountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolMountCifsMountOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountCifsMountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountCifsMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountCifsMountOutputReference_Override(b BatchPoolMountCifsMountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountCifsMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetMountOptions(val *string) {
	_jsii_.Set(
		j,
		"mountOptions",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetRelativeMountPath(val *string) {
	_jsii_.Set(
		j,
		"relativeMountPath",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetSource(val *string) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountCifsMountOutputReference) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) ResetMountOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountCifsMountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMountList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolMountOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountList
type jsiiProxy_BatchPoolMountList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolMountList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolMountList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolMountList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountList_Override(b BatchPoolMountList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountList) Get(index *float64) BatchPoolMountOutputReference {
	var returns BatchPoolMountOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMountNfsMount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#relative_mount_path BatchPool#relative_mount_path}.
	RelativeMountPath *string `field:"required" json:"relativeMountPath" yaml:"relativeMountPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#source BatchPool#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#mount_options BatchPool#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

type BatchPoolMountNfsMountList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolMountNfsMountOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountNfsMountList
type jsiiProxy_BatchPoolMountNfsMountList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolMountNfsMountList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolMountNfsMountList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountNfsMountList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountNfsMountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountNfsMountList_Override(b BatchPoolMountNfsMountList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountNfsMountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountNfsMountList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountList) Get(index *float64) BatchPoolMountNfsMountOutputReference {
	var returns BatchPoolMountNfsMountOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMountNfsMountOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountOptions() *string
	SetMountOptions(val *string)
	MountOptionsInput() *string
	RelativeMountPath() *string
	SetRelativeMountPath(val *string)
	RelativeMountPathInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
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
	ResetMountOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountNfsMountOutputReference
type jsiiProxy_BatchPoolMountNfsMountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) MountOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) MountOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) RelativeMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) RelativeMountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolMountNfsMountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolMountNfsMountOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountNfsMountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountNfsMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountNfsMountOutputReference_Override(b BatchPoolMountNfsMountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountNfsMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SetMountOptions(val *string) {
	_jsii_.Set(
		j,
		"mountOptions",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SetRelativeMountPath(val *string) {
	_jsii_.Set(
		j,
		"relativeMountPath",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SetSource(val *string) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountNfsMountOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) ResetMountOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountNfsMountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolMountOutputReference interface {
	cdktf.ComplexObject
	AzureBlobFileSystem() BatchPoolMountAzureBlobFileSystemOutputReference
	AzureBlobFileSystemInput() *BatchPoolMountAzureBlobFileSystem
	AzureFileShare() BatchPoolMountAzureFileShareList
	AzureFileShareInput() interface{}
	CifsMount() BatchPoolMountCifsMountList
	CifsMountInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NfsMount() BatchPoolMountNfsMountList
	NfsMountInput() interface{}
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
	PutAzureBlobFileSystem(value *BatchPoolMountAzureBlobFileSystem)
	PutAzureFileShare(value interface{})
	PutCifsMount(value interface{})
	PutNfsMount(value interface{})
	ResetAzureBlobFileSystem()
	ResetAzureFileShare()
	ResetCifsMount()
	ResetNfsMount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountOutputReference
type jsiiProxy_BatchPoolMountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolMountOutputReference) AzureBlobFileSystem() BatchPoolMountAzureBlobFileSystemOutputReference {
	var returns BatchPoolMountAzureBlobFileSystemOutputReference
	_jsii_.Get(
		j,
		"azureBlobFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) AzureBlobFileSystemInput() *BatchPoolMountAzureBlobFileSystem {
	var returns *BatchPoolMountAzureBlobFileSystem
	_jsii_.Get(
		j,
		"azureBlobFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) AzureFileShare() BatchPoolMountAzureFileShareList {
	var returns BatchPoolMountAzureFileShareList
	_jsii_.Get(
		j,
		"azureFileShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) AzureFileShareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureFileShareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) CifsMount() BatchPoolMountCifsMountList {
	var returns BatchPoolMountCifsMountList
	_jsii_.Get(
		j,
		"cifsMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) CifsMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cifsMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) NfsMount() BatchPoolMountNfsMountList {
	var returns BatchPoolMountNfsMountList
	_jsii_.Get(
		j,
		"nfsMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) NfsMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolMountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolMountOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolMountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountOutputReference_Override(b BatchPoolMountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) PutAzureBlobFileSystem(value *BatchPoolMountAzureBlobFileSystem) {
	_jsii_.InvokeVoid(
		b,
		"putAzureBlobFileSystem",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) PutAzureFileShare(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putAzureFileShare",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) PutCifsMount(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putCifsMount",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) PutNfsMount(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putNfsMount",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ResetAzureBlobFileSystem() {
	_jsii_.InvokeVoid(
		b,
		"resetAzureBlobFileSystem",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ResetAzureFileShare() {
	_jsii_.InvokeVoid(
		b,
		"resetAzureFileShare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ResetCifsMount() {
	_jsii_.InvokeVoid(
		b,
		"resetCifsMount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ResetNfsMount() {
	_jsii_.InvokeVoid(
		b,
		"resetNfsMount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#subnet_id BatchPool#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#endpoint_configuration BatchPool#endpoint_configuration}
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#public_address_provisioning_type BatchPool#public_address_provisioning_type}.
	PublicAddressProvisioningType *string `field:"optional" json:"publicAddressProvisioningType" yaml:"publicAddressProvisioningType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#public_ips BatchPool#public_ips}.
	PublicIps *[]*string `field:"optional" json:"publicIps" yaml:"publicIps"`
}

type BatchPoolNetworkConfigurationEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#backend_port BatchPool#backend_port}.
	BackendPort *float64 `field:"required" json:"backendPort" yaml:"backendPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#frontend_port_range BatchPool#frontend_port_range}.
	FrontendPortRange *string `field:"required" json:"frontendPortRange" yaml:"frontendPortRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#protocol BatchPool#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// network_security_group_rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#network_security_group_rules BatchPool#network_security_group_rules}
	NetworkSecurityGroupRules interface{} `field:"optional" json:"networkSecurityGroupRules" yaml:"networkSecurityGroupRules"`
}

type BatchPoolNetworkConfigurationEndpointConfigurationList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolNetworkConfigurationEndpointConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolNetworkConfigurationEndpointConfigurationList
type jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolNetworkConfigurationEndpointConfigurationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolNetworkConfigurationEndpointConfigurationList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolNetworkConfigurationEndpointConfigurationList_Override(b BatchPoolNetworkConfigurationEndpointConfigurationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) Get(index *float64) BatchPoolNetworkConfigurationEndpointConfigurationOutputReference {
	var returns BatchPoolNetworkConfigurationEndpointConfigurationOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#access BatchPool#access}.
	Access *string `field:"required" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#priority BatchPool#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#source_address_prefix BatchPool#source_address_prefix}.
	SourceAddressPrefix *string `field:"required" json:"sourceAddressPrefix" yaml:"sourceAddressPrefix"`
}

type BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList
type jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList_Override(b BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) Get(index *float64) BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference {
	var returns BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference interface {
	cdktf.ComplexObject
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	SourceAddressPrefix() *string
	SetSourceAddressPrefix(val *string)
	SourceAddressPrefixInput() *string
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

// The jsii proxy struct for BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference
type jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SourceAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SourceAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference_Override(b BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SetAccess(val *string) {
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SetSourceAddressPrefix(val *string) {
	_jsii_.Set(
		j,
		"sourceAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolNetworkConfigurationEndpointConfigurationOutputReference interface {
	cdktf.ComplexObject
	BackendPort() *float64
	SetBackendPort(val *float64)
	BackendPortInput() *float64
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
	FrontendPortRange() *string
	SetFrontendPortRange(val *string)
	FrontendPortRangeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSecurityGroupRules() BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList
	NetworkSecurityGroupRulesInput() interface{}
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	PutNetworkSecurityGroupRules(value interface{})
	ResetNetworkSecurityGroupRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolNetworkConfigurationEndpointConfigurationOutputReference
type jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) BackendPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) BackendPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) FrontendPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) FrontendPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) NetworkSecurityGroupRules() BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList {
	var returns BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList
	_jsii_.Get(
		j,
		"networkSecurityGroupRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) NetworkSecurityGroupRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSecurityGroupRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolNetworkConfigurationEndpointConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolNetworkConfigurationEndpointConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolNetworkConfigurationEndpointConfigurationOutputReference_Override(b BatchPoolNetworkConfigurationEndpointConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetBackendPort(val *float64) {
	_jsii_.Set(
		j,
		"backendPort",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetFrontendPortRange(val *string) {
	_jsii_.Set(
		j,
		"frontendPortRange",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) PutNetworkSecurityGroupRules(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putNetworkSecurityGroupRules",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) ResetNetworkSecurityGroupRules() {
	_jsii_.InvokeVoid(
		b,
		"resetNetworkSecurityGroupRules",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolNetworkConfigurationOutputReference interface {
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
	EndpointConfiguration() BatchPoolNetworkConfigurationEndpointConfigurationList
	EndpointConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *BatchPoolNetworkConfiguration
	SetInternalValue(val *BatchPoolNetworkConfiguration)
	PublicAddressProvisioningType() *string
	SetPublicAddressProvisioningType(val *string)
	PublicAddressProvisioningTypeInput() *string
	PublicIps() *[]*string
	SetPublicIps(val *[]*string)
	PublicIpsInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	PutEndpointConfiguration(value interface{})
	ResetEndpointConfiguration()
	ResetPublicAddressProvisioningType()
	ResetPublicIps()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolNetworkConfigurationOutputReference
type jsiiProxy_BatchPoolNetworkConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) EndpointConfiguration() BatchPoolNetworkConfigurationEndpointConfigurationList {
	var returns BatchPoolNetworkConfigurationEndpointConfigurationList
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) EndpointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) InternalValue() *BatchPoolNetworkConfiguration {
	var returns *BatchPoolNetworkConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PublicAddressProvisioningType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicAddressProvisioningType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PublicAddressProvisioningTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicAddressProvisioningTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PublicIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PublicIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolNetworkConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolNetworkConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolNetworkConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolNetworkConfigurationOutputReference_Override(b BatchPoolNetworkConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SetInternalValue(val *BatchPoolNetworkConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SetPublicAddressProvisioningType(val *string) {
	_jsii_.Set(
		j,
		"publicAddressProvisioningType",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SetPublicIps(val *[]*string) {
	_jsii_.Set(
		j,
		"publicIps",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PutEndpointConfiguration(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetPublicAddressProvisioningType() {
	_jsii_.InvokeVoid(
		b,
		"resetPublicAddressProvisioningType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetPublicIps() {
	_jsii_.InvokeVoid(
		b,
		"resetPublicIps",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolNodePlacement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#policy BatchPool#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

type BatchPoolNodePlacementList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolNodePlacementOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolNodePlacementList
type jsiiProxy_BatchPoolNodePlacementList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolNodePlacementList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolNodePlacementList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolNodePlacementList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolNodePlacementList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNodePlacementList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolNodePlacementList_Override(b BatchPoolNodePlacementList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNodePlacementList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNodePlacementList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementList) Get(index *float64) BatchPoolNodePlacementOutputReference {
	var returns BatchPoolNodePlacementOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolNodePlacementOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
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
	ResetPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolNodePlacementOutputReference
type jsiiProxy_BatchPoolNodePlacementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolNodePlacementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolNodePlacementOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolNodePlacementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNodePlacementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolNodePlacementOutputReference_Override(b BatchPoolNodePlacementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNodePlacementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNodePlacementOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) ResetPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNodePlacementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolStartTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#command_line BatchPool#command_line}.
	CommandLine *string `field:"required" json:"commandLine" yaml:"commandLine"`
	// user_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_identity BatchPool#user_identity}
	UserIdentity *BatchPoolStartTaskUserIdentity `field:"required" json:"userIdentity" yaml:"userIdentity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#common_environment_properties BatchPool#common_environment_properties}.
	CommonEnvironmentProperties *map[string]*string `field:"optional" json:"commonEnvironmentProperties" yaml:"commonEnvironmentProperties"`
	// resource_file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#resource_file BatchPool#resource_file}
	ResourceFile interface{} `field:"optional" json:"resourceFile" yaml:"resourceFile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#task_retry_maximum BatchPool#task_retry_maximum}.
	TaskRetryMaximum *float64 `field:"optional" json:"taskRetryMaximum" yaml:"taskRetryMaximum"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#wait_for_success BatchPool#wait_for_success}.
	WaitForSuccess interface{} `field:"optional" json:"waitForSuccess" yaml:"waitForSuccess"`
}

type BatchPoolStartTaskOutputReference interface {
	cdktf.ComplexObject
	CommandLine() *string
	SetCommandLine(val *string)
	CommandLineInput() *string
	CommonEnvironmentProperties() *map[string]*string
	SetCommonEnvironmentProperties(val *map[string]*string)
	CommonEnvironmentPropertiesInput() *map[string]*string
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
	InternalValue() *BatchPoolStartTask
	SetInternalValue(val *BatchPoolStartTask)
	ResourceFile() BatchPoolStartTaskResourceFileList
	ResourceFileInput() interface{}
	TaskRetryMaximum() *float64
	SetTaskRetryMaximum(val *float64)
	TaskRetryMaximumInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserIdentity() BatchPoolStartTaskUserIdentityOutputReference
	UserIdentityInput() *BatchPoolStartTaskUserIdentity
	WaitForSuccess() interface{}
	SetWaitForSuccess(val interface{})
	WaitForSuccessInput() interface{}
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
	PutResourceFile(value interface{})
	PutUserIdentity(value *BatchPoolStartTaskUserIdentity)
	ResetCommonEnvironmentProperties()
	ResetResourceFile()
	ResetTaskRetryMaximum()
	ResetWaitForSuccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolStartTaskOutputReference
type jsiiProxy_BatchPoolStartTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CommonEnvironmentProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commonEnvironmentProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CommonEnvironmentPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commonEnvironmentPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) InternalValue() *BatchPoolStartTask {
	var returns *BatchPoolStartTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ResourceFile() BatchPoolStartTaskResourceFileList {
	var returns BatchPoolStartTaskResourceFileList
	_jsii_.Get(
		j,
		"resourceFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ResourceFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) TaskRetryMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskRetryMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) TaskRetryMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskRetryMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) UserIdentity() BatchPoolStartTaskUserIdentityOutputReference {
	var returns BatchPoolStartTaskUserIdentityOutputReference
	_jsii_.Get(
		j,
		"userIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) UserIdentityInput() *BatchPoolStartTaskUserIdentity {
	var returns *BatchPoolStartTaskUserIdentity
	_jsii_.Get(
		j,
		"userIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) WaitForSuccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) WaitForSuccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSuccessInput",
		&returns,
	)
	return returns
}


func NewBatchPoolStartTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolStartTaskOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolStartTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolStartTaskOutputReference_Override(b BatchPoolStartTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetCommandLine(val *string) {
	_jsii_.Set(
		j,
		"commandLine",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetCommonEnvironmentProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"commonEnvironmentProperties",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetInternalValue(val *BatchPoolStartTask) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetTaskRetryMaximum(val *float64) {
	_jsii_.Set(
		j,
		"taskRetryMaximum",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) SetWaitForSuccess(val interface{}) {
	_jsii_.Set(
		j,
		"waitForSuccess",
		val,
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) PutResourceFile(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putResourceFile",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) PutUserIdentity(value *BatchPoolStartTaskUserIdentity) {
	_jsii_.InvokeVoid(
		b,
		"putUserIdentity",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetCommonEnvironmentProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetCommonEnvironmentProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetResourceFile() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceFile",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetTaskRetryMaximum() {
	_jsii_.InvokeVoid(
		b,
		"resetTaskRetryMaximum",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetWaitForSuccess() {
	_jsii_.InvokeVoid(
		b,
		"resetWaitForSuccess",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolStartTaskResourceFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#auto_storage_container_name BatchPool#auto_storage_container_name}.
	AutoStorageContainerName *string `field:"optional" json:"autoStorageContainerName" yaml:"autoStorageContainerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#blob_prefix BatchPool#blob_prefix}.
	BlobPrefix *string `field:"optional" json:"blobPrefix" yaml:"blobPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#file_mode BatchPool#file_mode}.
	FileMode *string `field:"optional" json:"fileMode" yaml:"fileMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#file_path BatchPool#file_path}.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#http_url BatchPool#http_url}.
	HttpUrl *string `field:"optional" json:"httpUrl" yaml:"httpUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#storage_container_url BatchPool#storage_container_url}.
	StorageContainerUrl *string `field:"optional" json:"storageContainerUrl" yaml:"storageContainerUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_assigned_identity_id BatchPool#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

type BatchPoolStartTaskResourceFileList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolStartTaskResourceFileOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolStartTaskResourceFileList
type jsiiProxy_BatchPoolStartTaskResourceFileList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolStartTaskResourceFileList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolStartTaskResourceFileList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolStartTaskResourceFileList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskResourceFileList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolStartTaskResourceFileList_Override(b BatchPoolStartTaskResourceFileList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskResourceFileList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileList) Get(index *float64) BatchPoolStartTaskResourceFileOutputReference {
	var returns BatchPoolStartTaskResourceFileOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolStartTaskResourceFileOutputReference interface {
	cdktf.ComplexObject
	AutoStorageContainerName() *string
	SetAutoStorageContainerName(val *string)
	AutoStorageContainerNameInput() *string
	BlobPrefix() *string
	SetBlobPrefix(val *string)
	BlobPrefixInput() *string
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
	FileMode() *string
	SetFileMode(val *string)
	FileModeInput() *string
	FilePath() *string
	SetFilePath(val *string)
	FilePathInput() *string
	// Experimental.
	Fqn() *string
	HttpUrl() *string
	SetHttpUrl(val *string)
	HttpUrlInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageContainerUrl() *string
	SetStorageContainerUrl(val *string)
	StorageContainerUrlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAssignedIdentityId() *string
	SetUserAssignedIdentityId(val *string)
	UserAssignedIdentityIdInput() *string
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
	ResetAutoStorageContainerName()
	ResetBlobPrefix()
	ResetFileMode()
	ResetFilePath()
	ResetHttpUrl()
	ResetStorageContainerUrl()
	ResetUserAssignedIdentityId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolStartTaskResourceFileOutputReference
type jsiiProxy_BatchPoolStartTaskResourceFileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) AutoStorageContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoStorageContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) AutoStorageContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoStorageContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) BlobPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) BlobPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) FileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) FileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) FilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) HttpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) HttpUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) StorageContainerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) StorageContainerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) UserAssignedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityIdInput",
		&returns,
	)
	return returns
}


func NewBatchPoolStartTaskResourceFileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolStartTaskResourceFileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolStartTaskResourceFileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskResourceFileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolStartTaskResourceFileOutputReference_Override(b BatchPoolStartTaskResourceFileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskResourceFileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetAutoStorageContainerName(val *string) {
	_jsii_.Set(
		j,
		"autoStorageContainerName",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetBlobPrefix(val *string) {
	_jsii_.Set(
		j,
		"blobPrefix",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetFileMode(val *string) {
	_jsii_.Set(
		j,
		"fileMode",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetFilePath(val *string) {
	_jsii_.Set(
		j,
		"filePath",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetHttpUrl(val *string) {
	_jsii_.Set(
		j,
		"httpUrl",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetStorageContainerUrl(val *string) {
	_jsii_.Set(
		j,
		"storageContainerUrl",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) SetUserAssignedIdentityId(val *string) {
	_jsii_.Set(
		j,
		"userAssignedIdentityId",
		val,
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetAutoStorageContainerName() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoStorageContainerName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetBlobPrefix() {
	_jsii_.InvokeVoid(
		b,
		"resetBlobPrefix",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetFileMode() {
	_jsii_.InvokeVoid(
		b,
		"resetFileMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetFilePath() {
	_jsii_.InvokeVoid(
		b,
		"resetFilePath",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetHttpUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetHttpUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetStorageContainerUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageContainerUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetUserAssignedIdentityId() {
	_jsii_.InvokeVoid(
		b,
		"resetUserAssignedIdentityId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolStartTaskUserIdentity struct {
	// auto_user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#auto_user BatchPool#auto_user}
	AutoUser *BatchPoolStartTaskUserIdentityAutoUser `field:"optional" json:"autoUser" yaml:"autoUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_name BatchPool#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

type BatchPoolStartTaskUserIdentityAutoUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#elevation_level BatchPool#elevation_level}.
	ElevationLevel *string `field:"optional" json:"elevationLevel" yaml:"elevationLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#scope BatchPool#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

type BatchPoolStartTaskUserIdentityAutoUserOutputReference interface {
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
	ElevationLevel() *string
	SetElevationLevel(val *string)
	ElevationLevelInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BatchPoolStartTaskUserIdentityAutoUser
	SetInternalValue(val *BatchPoolStartTaskUserIdentityAutoUser)
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
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
	ResetElevationLevel()
	ResetScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolStartTaskUserIdentityAutoUserOutputReference
type jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ElevationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elevationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ElevationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elevationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) InternalValue() *BatchPoolStartTaskUserIdentityAutoUser {
	var returns *BatchPoolStartTaskUserIdentityAutoUser
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolStartTaskUserIdentityAutoUserOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolStartTaskUserIdentityAutoUserOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskUserIdentityAutoUserOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolStartTaskUserIdentityAutoUserOutputReference_Override(b BatchPoolStartTaskUserIdentityAutoUserOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskUserIdentityAutoUserOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) SetElevationLevel(val *string) {
	_jsii_.Set(
		j,
		"elevationLevel",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) SetInternalValue(val *BatchPoolStartTaskUserIdentityAutoUser) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ResetElevationLevel() {
	_jsii_.InvokeVoid(
		b,
		"resetElevationLevel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		b,
		"resetScope",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityAutoUserOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolStartTaskUserIdentityOutputReference interface {
	cdktf.ComplexObject
	AutoUser() BatchPoolStartTaskUserIdentityAutoUserOutputReference
	AutoUserInput() *BatchPoolStartTaskUserIdentityAutoUser
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
	InternalValue() *BatchPoolStartTaskUserIdentity
	SetInternalValue(val *BatchPoolStartTaskUserIdentity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
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
	PutAutoUser(value *BatchPoolStartTaskUserIdentityAutoUser)
	ResetAutoUser()
	ResetUserName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolStartTaskUserIdentityOutputReference
type jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) AutoUser() BatchPoolStartTaskUserIdentityAutoUserOutputReference {
	var returns BatchPoolStartTaskUserIdentityAutoUserOutputReference
	_jsii_.Get(
		j,
		"autoUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) AutoUserInput() *BatchPoolStartTaskUserIdentityAutoUser {
	var returns *BatchPoolStartTaskUserIdentityAutoUser
	_jsii_.Get(
		j,
		"autoUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) InternalValue() *BatchPoolStartTaskUserIdentity {
	var returns *BatchPoolStartTaskUserIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}


func NewBatchPoolStartTaskUserIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolStartTaskUserIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskUserIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolStartTaskUserIdentityOutputReference_Override(b BatchPoolStartTaskUserIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskUserIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) SetInternalValue(val *BatchPoolStartTaskUserIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) PutAutoUser(value *BatchPoolStartTaskUserIdentityAutoUser) {
	_jsii_.InvokeVoid(
		b,
		"putAutoUser",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) ResetAutoUser() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoUser",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) ResetUserName() {
	_jsii_.InvokeVoid(
		b,
		"resetUserName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskUserIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolStorageImageReference struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#id BatchPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#offer BatchPool#offer}.
	Offer *string `field:"optional" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#publisher BatchPool#publisher}.
	Publisher *string `field:"optional" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#sku BatchPool#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#version BatchPool#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

type BatchPoolStorageImageReferenceOutputReference interface {
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *BatchPoolStorageImageReference
	SetInternalValue(val *BatchPoolStorageImageReference)
	Offer() *string
	SetOffer(val *string)
	OfferInput() *string
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetId()
	ResetOffer()
	ResetPublisher()
	ResetSku()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolStorageImageReferenceOutputReference
type jsiiProxy_BatchPoolStorageImageReferenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) InternalValue() *BatchPoolStorageImageReference {
	var returns *BatchPoolStorageImageReference
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) Offer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) OfferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewBatchPoolStorageImageReferenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolStorageImageReferenceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolStorageImageReferenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStorageImageReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolStorageImageReferenceOutputReference_Override(b BatchPoolStorageImageReferenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStorageImageReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetInternalValue(val *BatchPoolStorageImageReference) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetOffer(val *string) {
	_jsii_.Set(
		j,
		"offer",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetPublisher(val *string) {
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ResetOffer() {
	_jsii_.InvokeVoid(
		b,
		"resetOffer",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ResetPublisher() {
	_jsii_.InvokeVoid(
		b,
		"resetPublisher",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ResetSku() {
	_jsii_.InvokeVoid(
		b,
		"resetSku",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStorageImageReferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolTaskSchedulingPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#node_fill_type BatchPool#node_fill_type}.
	NodeFillType *string `field:"optional" json:"nodeFillType" yaml:"nodeFillType"`
}

type BatchPoolTaskSchedulingPolicyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolTaskSchedulingPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolTaskSchedulingPolicyList
type jsiiProxy_BatchPoolTaskSchedulingPolicyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolTaskSchedulingPolicyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolTaskSchedulingPolicyList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolTaskSchedulingPolicyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolTaskSchedulingPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolTaskSchedulingPolicyList_Override(b BatchPoolTaskSchedulingPolicyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolTaskSchedulingPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyList) Get(index *float64) BatchPoolTaskSchedulingPolicyOutputReference {
	var returns BatchPoolTaskSchedulingPolicyOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolTaskSchedulingPolicyOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NodeFillType() *string
	SetNodeFillType(val *string)
	NodeFillTypeInput() *string
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
	ResetNodeFillType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolTaskSchedulingPolicyOutputReference
type jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) NodeFillType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeFillType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) NodeFillTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeFillTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolTaskSchedulingPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolTaskSchedulingPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolTaskSchedulingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolTaskSchedulingPolicyOutputReference_Override(b BatchPoolTaskSchedulingPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolTaskSchedulingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) SetNodeFillType(val *string) {
	_jsii_.Set(
		j,
		"nodeFillType",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) ResetNodeFillType() {
	_jsii_.InvokeVoid(
		b,
		"resetNodeFillType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#create BatchPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#delete BatchPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#read BatchPool#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#update BatchPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type BatchPoolTimeoutsOutputReference interface {
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

// The jsii proxy struct for BatchPoolTimeoutsOutputReference
type jsiiProxy_BatchPoolTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewBatchPoolTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolTimeoutsOutputReference_Override(b BatchPoolTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		b,
		"resetCreate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		b,
		"resetDelete",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		b,
		"resetRead",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		b,
		"resetUpdate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolUserAccounts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#elevation_level BatchPool#elevation_level}.
	ElevationLevel *string `field:"required" json:"elevationLevel" yaml:"elevationLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#password BatchPool#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// linux_user_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#linux_user_configuration BatchPool#linux_user_configuration}
	LinuxUserConfiguration interface{} `field:"optional" json:"linuxUserConfiguration" yaml:"linuxUserConfiguration"`
	// windows_user_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#windows_user_configuration BatchPool#windows_user_configuration}
	WindowsUserConfiguration interface{} `field:"optional" json:"windowsUserConfiguration" yaml:"windowsUserConfiguration"`
}

type BatchPoolUserAccountsLinuxUserConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#gid BatchPool#gid}.
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#ssh_private_key BatchPool#ssh_private_key}.
	SshPrivateKey *string `field:"optional" json:"sshPrivateKey" yaml:"sshPrivateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#uid BatchPool#uid}.
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

type BatchPoolUserAccountsLinuxUserConfigurationList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolUserAccountsLinuxUserConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolUserAccountsLinuxUserConfigurationList
type jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolUserAccountsLinuxUserConfigurationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolUserAccountsLinuxUserConfigurationList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsLinuxUserConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolUserAccountsLinuxUserConfigurationList_Override(b BatchPoolUserAccountsLinuxUserConfigurationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsLinuxUserConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) Get(index *float64) BatchPoolUserAccountsLinuxUserConfigurationOutputReference {
	var returns BatchPoolUserAccountsLinuxUserConfigurationOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolUserAccountsLinuxUserConfigurationOutputReference interface {
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
	Gid() *float64
	SetGid(val *float64)
	GidInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SshPrivateKey() *string
	SetSshPrivateKey(val *string)
	SshPrivateKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uid() *float64
	SetUid(val *float64)
	UidInput() *float64
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
	ResetGid()
	ResetSshPrivateKey()
	ResetUid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolUserAccountsLinuxUserConfigurationOutputReference
type jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) Gid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SshPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SshPrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) Uid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) UidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}


func NewBatchPoolUserAccountsLinuxUserConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolUserAccountsLinuxUserConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsLinuxUserConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolUserAccountsLinuxUserConfigurationOutputReference_Override(b BatchPoolUserAccountsLinuxUserConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsLinuxUserConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SetGid(val *float64) {
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SetSshPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"sshPrivateKey",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) SetUid(val *float64) {
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) ResetGid() {
	_jsii_.InvokeVoid(
		b,
		"resetGid",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) ResetSshPrivateKey() {
	_jsii_.InvokeVoid(
		b,
		"resetSshPrivateKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) ResetUid() {
	_jsii_.InvokeVoid(
		b,
		"resetUid",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsLinuxUserConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolUserAccountsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolUserAccountsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolUserAccountsList
type jsiiProxy_BatchPoolUserAccountsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolUserAccountsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolUserAccountsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolUserAccountsList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolUserAccountsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolUserAccountsList_Override(b BatchPoolUserAccountsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsList) Get(index *float64) BatchPoolUserAccountsOutputReference {
	var returns BatchPoolUserAccountsOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolUserAccountsOutputReference interface {
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
	ElevationLevel() *string
	SetElevationLevel(val *string)
	ElevationLevelInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinuxUserConfiguration() BatchPoolUserAccountsLinuxUserConfigurationList
	LinuxUserConfigurationInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsUserConfiguration() BatchPoolUserAccountsWindowsUserConfigurationList
	WindowsUserConfigurationInput() interface{}
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
	PutLinuxUserConfiguration(value interface{})
	PutWindowsUserConfiguration(value interface{})
	ResetLinuxUserConfiguration()
	ResetWindowsUserConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolUserAccountsOutputReference
type jsiiProxy_BatchPoolUserAccountsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) ElevationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elevationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) ElevationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elevationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) LinuxUserConfiguration() BatchPoolUserAccountsLinuxUserConfigurationList {
	var returns BatchPoolUserAccountsLinuxUserConfigurationList
	_jsii_.Get(
		j,
		"linuxUserConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) LinuxUserConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxUserConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) WindowsUserConfiguration() BatchPoolUserAccountsWindowsUserConfigurationList {
	var returns BatchPoolUserAccountsWindowsUserConfigurationList
	_jsii_.Get(
		j,
		"windowsUserConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) WindowsUserConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsUserConfigurationInput",
		&returns,
	)
	return returns
}


func NewBatchPoolUserAccountsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolUserAccountsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolUserAccountsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolUserAccountsOutputReference_Override(b BatchPoolUserAccountsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) SetElevationLevel(val *string) {
	_jsii_.Set(
		j,
		"elevationLevel",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) PutLinuxUserConfiguration(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putLinuxUserConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) PutWindowsUserConfiguration(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putWindowsUserConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) ResetLinuxUserConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetLinuxUserConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) ResetWindowsUserConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetWindowsUserConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolUserAccountsWindowsUserConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#login_mode BatchPool#login_mode}.
	LoginMode *string `field:"required" json:"loginMode" yaml:"loginMode"`
}

type BatchPoolUserAccountsWindowsUserConfigurationList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolUserAccountsWindowsUserConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolUserAccountsWindowsUserConfigurationList
type jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolUserAccountsWindowsUserConfigurationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolUserAccountsWindowsUserConfigurationList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsWindowsUserConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolUserAccountsWindowsUserConfigurationList_Override(b BatchPoolUserAccountsWindowsUserConfigurationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsWindowsUserConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) Get(index *float64) BatchPoolUserAccountsWindowsUserConfigurationOutputReference {
	var returns BatchPoolUserAccountsWindowsUserConfigurationOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolUserAccountsWindowsUserConfigurationOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoginMode() *string
	SetLoginMode(val *string)
	LoginModeInput() *string
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

// The jsii proxy struct for BatchPoolUserAccountsWindowsUserConfigurationOutputReference
type jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) LoginMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) LoginModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolUserAccountsWindowsUserConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolUserAccountsWindowsUserConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsWindowsUserConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolUserAccountsWindowsUserConfigurationOutputReference_Override(b BatchPoolUserAccountsWindowsUserConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolUserAccountsWindowsUserConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) SetLoginMode(val *string) {
	_jsii_.Set(
		j,
		"loginMode",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolWindows struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#enable_automatic_updates BatchPool#enable_automatic_updates}.
	EnableAutomaticUpdates interface{} `field:"optional" json:"enableAutomaticUpdates" yaml:"enableAutomaticUpdates"`
}

type BatchPoolWindowsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BatchPoolWindowsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolWindowsList
type jsiiProxy_BatchPoolWindowsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchPoolWindowsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolWindowsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolWindowsList {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolWindowsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolWindowsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolWindowsList_Override(b BatchPoolWindowsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolWindowsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolWindowsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolWindowsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolWindowsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolWindowsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolWindowsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsList) Get(index *float64) BatchPoolWindowsOutputReference {
	var returns BatchPoolWindowsOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchPoolWindowsOutputReference interface {
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
	EnableAutomaticUpdates() interface{}
	SetEnableAutomaticUpdates(val interface{})
	EnableAutomaticUpdatesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetEnableAutomaticUpdates()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolWindowsOutputReference
type jsiiProxy_BatchPoolWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) EnableAutomaticUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) EnableAutomaticUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolWindowsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchPoolWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolWindowsOutputReference_Override(b BatchPoolWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) SetEnableAutomaticUpdates(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutomaticUpdates",
		val,
	)
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolWindowsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) ResetEnableAutomaticUpdates() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableAutomaticUpdates",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

