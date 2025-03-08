// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinepacketcapture

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/virtualmachinepacketcapture/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/virtual_machine_packet_capture azurerm_virtual_machine_packet_capture}.
type VirtualMachinePacketCapture interface {
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
	Filter() VirtualMachinePacketCaptureFilterList
	FilterInput() interface{}
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
	MaximumBytesPerPacket() *float64
	SetMaximumBytesPerPacket(val *float64)
	MaximumBytesPerPacketInput() *float64
	MaximumBytesPerSession() *float64
	SetMaximumBytesPerSession(val *float64)
	MaximumBytesPerSessionInput() *float64
	MaximumCaptureDurationInSeconds() *float64
	SetMaximumCaptureDurationInSeconds(val *float64)
	MaximumCaptureDurationInSecondsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkWatcherId() *string
	SetNetworkWatcherId(val *string)
	NetworkWatcherIdInput() *string
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
	StorageLocation() VirtualMachinePacketCaptureStorageLocationOutputReference
	StorageLocationInput() *VirtualMachinePacketCaptureStorageLocation
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualMachinePacketCaptureTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualMachineId() *string
	SetVirtualMachineId(val *string)
	VirtualMachineIdInput() *string
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
	PutFilter(value interface{})
	PutStorageLocation(value *VirtualMachinePacketCaptureStorageLocation)
	PutTimeouts(value *VirtualMachinePacketCaptureTimeouts)
	ResetFilter()
	ResetId()
	ResetMaximumBytesPerPacket()
	ResetMaximumBytesPerSession()
	ResetMaximumCaptureDurationInSeconds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for VirtualMachinePacketCapture
type jsiiProxy_VirtualMachinePacketCapture struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualMachinePacketCapture) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Filter() VirtualMachinePacketCaptureFilterList {
	var returns VirtualMachinePacketCaptureFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) MaximumBytesPerPacket() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerPacket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) MaximumBytesPerPacketInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerPacketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) MaximumBytesPerSession() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) MaximumBytesPerSessionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) MaximumCaptureDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCaptureDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) MaximumCaptureDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCaptureDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) NetworkWatcherId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkWatcherId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) NetworkWatcherIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkWatcherIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) StorageLocation() VirtualMachinePacketCaptureStorageLocationOutputReference {
	var returns VirtualMachinePacketCaptureStorageLocationOutputReference
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) StorageLocationInput() *VirtualMachinePacketCaptureStorageLocation {
	var returns *VirtualMachinePacketCaptureStorageLocation
	_jsii_.Get(
		j,
		"storageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) Timeouts() VirtualMachinePacketCaptureTimeoutsOutputReference {
	var returns VirtualMachinePacketCaptureTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) VirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCapture) VirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/virtual_machine_packet_capture azurerm_virtual_machine_packet_capture} Resource.
func NewVirtualMachinePacketCapture(scope constructs.Construct, id *string, config *VirtualMachinePacketCaptureConfig) VirtualMachinePacketCapture {
	_init_.Initialize()

	if err := validateNewVirtualMachinePacketCaptureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachinePacketCapture{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCapture",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/virtual_machine_packet_capture azurerm_virtual_machine_packet_capture} Resource.
func NewVirtualMachinePacketCapture_Override(v VirtualMachinePacketCapture, scope constructs.Construct, id *string, config *VirtualMachinePacketCaptureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCapture",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetMaximumBytesPerPacket(val *float64) {
	if err := j.validateSetMaximumBytesPerPacketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBytesPerPacket",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetMaximumBytesPerSession(val *float64) {
	if err := j.validateSetMaximumBytesPerSessionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBytesPerSession",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetMaximumCaptureDurationInSeconds(val *float64) {
	if err := j.validateSetMaximumCaptureDurationInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCaptureDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetNetworkWatcherId(val *string) {
	if err := j.validateSetNetworkWatcherIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkWatcherId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCapture)SetVirtualMachineId(val *string) {
	if err := j.validateSetVirtualMachineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineId",
		val,
	)
}

// Generates CDKTF code for importing a VirtualMachinePacketCapture resource upon running "cdktf plan <stack-name>".
func VirtualMachinePacketCapture_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualMachinePacketCapture_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCapture",
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
func VirtualMachinePacketCapture_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachinePacketCapture_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCapture",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachinePacketCapture_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachinePacketCapture_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCapture",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachinePacketCapture_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachinePacketCapture_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCapture",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualMachinePacketCapture_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCapture",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCapture) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachinePacketCapture) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) PutFilter(value interface{}) {
	if err := v.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putFilter",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) PutStorageLocation(value *VirtualMachinePacketCaptureStorageLocation) {
	if err := v.validatePutStorageLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageLocation",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) PutTimeouts(value *VirtualMachinePacketCaptureTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ResetFilter() {
	_jsii_.InvokeVoid(
		v,
		"resetFilter",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ResetMaximumBytesPerPacket() {
	_jsii_.InvokeVoid(
		v,
		"resetMaximumBytesPerPacket",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ResetMaximumBytesPerSession() {
	_jsii_.InvokeVoid(
		v,
		"resetMaximumBytesPerSession",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ResetMaximumCaptureDurationInSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetMaximumCaptureDurationInSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCapture) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCapture) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCapture) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

