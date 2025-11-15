// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcivirtualharddisk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/stackhcivirtualharddisk/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/stack_hci_virtual_hard_disk azurerm_stack_hci_virtual_hard_disk}.
type StackHciVirtualHardDisk interface {
	cdktf.TerraformResource
	BlockSizeInBytes() *float64
	SetBlockSizeInBytes(val *float64)
	BlockSizeInBytesInput() *float64
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
	CustomLocationId() *string
	SetCustomLocationId(val *string)
	CustomLocationIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskFileFormat() *string
	SetDiskFileFormat(val *string)
	DiskFileFormatInput() *string
	DiskSizeInGb() *float64
	SetDiskSizeInGb(val *float64)
	DiskSizeInGbInput() *float64
	DynamicEnabled() interface{}
	SetDynamicEnabled(val interface{})
	DynamicEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HypervGeneration() *string
	SetHypervGeneration(val *string)
	HypervGenerationInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogicalSectorInBytes() *float64
	SetLogicalSectorInBytes(val *float64)
	LogicalSectorInBytesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PhysicalSectorInBytes() *float64
	SetPhysicalSectorInBytes(val *float64)
	PhysicalSectorInBytesInput() *float64
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
	StoragePathId() *string
	SetStoragePathId(val *string)
	StoragePathIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StackHciVirtualHardDiskTimeoutsOutputReference
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
	PutTimeouts(value *StackHciVirtualHardDiskTimeouts)
	ResetBlockSizeInBytes()
	ResetDiskFileFormat()
	ResetDynamicEnabled()
	ResetHypervGeneration()
	ResetId()
	ResetLogicalSectorInBytes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhysicalSectorInBytes()
	ResetStoragePathId()
	ResetTags()
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

// The jsii proxy struct for StackHciVirtualHardDisk
type jsiiProxy_StackHciVirtualHardDisk struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StackHciVirtualHardDisk) BlockSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) BlockSizeInBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeInBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) CustomLocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLocationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) CustomLocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLocationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) DiskFileFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskFileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) DiskFileFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskFileFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) DiskSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) DiskSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) DynamicEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) DynamicEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) HypervGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypervGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) HypervGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypervGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) LogicalSectorInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logicalSectorInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) LogicalSectorInBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logicalSectorInBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) PhysicalSectorInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalSectorInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) PhysicalSectorInBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalSectorInBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) StoragePathId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePathId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) StoragePathIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePathIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) Timeouts() StackHciVirtualHardDiskTimeoutsOutputReference {
	var returns StackHciVirtualHardDiskTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciVirtualHardDisk) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/stack_hci_virtual_hard_disk azurerm_stack_hci_virtual_hard_disk} Resource.
func NewStackHciVirtualHardDisk(scope constructs.Construct, id *string, config *StackHciVirtualHardDiskConfig) StackHciVirtualHardDisk {
	_init_.Initialize()

	if err := validateNewStackHciVirtualHardDiskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackHciVirtualHardDisk{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciVirtualHardDisk.StackHciVirtualHardDisk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/stack_hci_virtual_hard_disk azurerm_stack_hci_virtual_hard_disk} Resource.
func NewStackHciVirtualHardDisk_Override(s StackHciVirtualHardDisk, scope constructs.Construct, id *string, config *StackHciVirtualHardDiskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciVirtualHardDisk.StackHciVirtualHardDisk",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetBlockSizeInBytes(val *float64) {
	if err := j.validateSetBlockSizeInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockSizeInBytes",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetCustomLocationId(val *string) {
	if err := j.validateSetCustomLocationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLocationId",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetDiskFileFormat(val *string) {
	if err := j.validateSetDiskFileFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskFileFormat",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetDiskSizeInGb(val *float64) {
	if err := j.validateSetDiskSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeInGb",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetDynamicEnabled(val interface{}) {
	if err := j.validateSetDynamicEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetHypervGeneration(val *string) {
	if err := j.validateSetHypervGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hypervGeneration",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetLogicalSectorInBytes(val *float64) {
	if err := j.validateSetLogicalSectorInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logicalSectorInBytes",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetPhysicalSectorInBytes(val *float64) {
	if err := j.validateSetPhysicalSectorInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalSectorInBytes",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetStoragePathId(val *string) {
	if err := j.validateSetStoragePathIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePathId",
		val,
	)
}

func (j *jsiiProxy_StackHciVirtualHardDisk)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a StackHciVirtualHardDisk resource upon running "cdktf plan <stack-name>".
func StackHciVirtualHardDisk_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStackHciVirtualHardDisk_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.stackHciVirtualHardDisk.StackHciVirtualHardDisk",
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
func StackHciVirtualHardDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStackHciVirtualHardDisk_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.stackHciVirtualHardDisk.StackHciVirtualHardDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StackHciVirtualHardDisk_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStackHciVirtualHardDisk_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.stackHciVirtualHardDisk.StackHciVirtualHardDisk",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StackHciVirtualHardDisk_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStackHciVirtualHardDisk_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.stackHciVirtualHardDisk.StackHciVirtualHardDisk",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StackHciVirtualHardDisk_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.stackHciVirtualHardDisk.StackHciVirtualHardDisk",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StackHciVirtualHardDisk) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciVirtualHardDisk) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) PutTimeouts(value *StackHciVirtualHardDiskTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetBlockSizeInBytes() {
	_jsii_.InvokeVoid(
		s,
		"resetBlockSizeInBytes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetDiskFileFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDiskFileFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetDynamicEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDynamicEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetHypervGeneration() {
	_jsii_.InvokeVoid(
		s,
		"resetHypervGeneration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetLogicalSectorInBytes() {
	_jsii_.InvokeVoid(
		s,
		"resetLogicalSectorInBytes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetPhysicalSectorInBytes() {
	_jsii_.InvokeVoid(
		s,
		"resetPhysicalSectorInBytes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetStoragePathId() {
	_jsii_.InvokeVoid(
		s,
		"resetStoragePathId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciVirtualHardDisk) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciVirtualHardDisk) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciVirtualHardDisk) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

