// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpacketcapture

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/networkpacketcapture/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/network_packet_capture azurerm_network_packet_capture}.
type NetworkPacketCapture interface {
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
	Filter() NetworkPacketCaptureFilterList
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
	MaximumCaptureDuration() *float64
	SetMaximumCaptureDuration(val *float64)
	MaximumCaptureDurationInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkWatcherName() *string
	SetNetworkWatcherName(val *string)
	NetworkWatcherNameInput() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	StorageLocation() NetworkPacketCaptureStorageLocationOutputReference
	StorageLocationInput() *NetworkPacketCaptureStorageLocation
	TargetResourceId() *string
	SetTargetResourceId(val *string)
	TargetResourceIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkPacketCaptureTimeoutsOutputReference
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
	PutFilter(value interface{})
	PutStorageLocation(value *NetworkPacketCaptureStorageLocation)
	PutTimeouts(value *NetworkPacketCaptureTimeouts)
	ResetFilter()
	ResetId()
	ResetMaximumBytesPerPacket()
	ResetMaximumBytesPerSession()
	ResetMaximumCaptureDuration()
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

// The jsii proxy struct for NetworkPacketCapture
type jsiiProxy_NetworkPacketCapture struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkPacketCapture) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Filter() NetworkPacketCaptureFilterList {
	var returns NetworkPacketCaptureFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) MaximumBytesPerPacket() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerPacket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) MaximumBytesPerPacketInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerPacketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) MaximumBytesPerSession() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) MaximumBytesPerSessionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) MaximumCaptureDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCaptureDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) MaximumCaptureDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCaptureDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) NetworkWatcherName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkWatcherName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) NetworkWatcherNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkWatcherNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) StorageLocation() NetworkPacketCaptureStorageLocationOutputReference {
	var returns NetworkPacketCaptureStorageLocationOutputReference
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) StorageLocationInput() *NetworkPacketCaptureStorageLocation {
	var returns *NetworkPacketCaptureStorageLocation
	_jsii_.Get(
		j,
		"storageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) TargetResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) TargetResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) Timeouts() NetworkPacketCaptureTimeoutsOutputReference {
	var returns NetworkPacketCaptureTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCapture) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/network_packet_capture azurerm_network_packet_capture} Resource.
func NewNetworkPacketCapture(scope constructs.Construct, id *string, config *NetworkPacketCaptureConfig) NetworkPacketCapture {
	_init_.Initialize()

	if err := validateNewNetworkPacketCaptureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkPacketCapture{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCapture",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/network_packet_capture azurerm_network_packet_capture} Resource.
func NewNetworkPacketCapture_Override(n NetworkPacketCapture, scope constructs.Construct, id *string, config *NetworkPacketCaptureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCapture",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetMaximumBytesPerPacket(val *float64) {
	if err := j.validateSetMaximumBytesPerPacketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBytesPerPacket",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetMaximumBytesPerSession(val *float64) {
	if err := j.validateSetMaximumBytesPerSessionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBytesPerSession",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetMaximumCaptureDuration(val *float64) {
	if err := j.validateSetMaximumCaptureDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCaptureDuration",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetNetworkWatcherName(val *string) {
	if err := j.validateSetNetworkWatcherNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkWatcherName",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCapture)SetTargetResourceId(val *string) {
	if err := j.validateSetTargetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceId",
		val,
	)
}

// Generates CDKTF code for importing a NetworkPacketCapture resource upon running "cdktf plan <stack-name>".
func NetworkPacketCapture_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkPacketCapture_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCapture",
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
func NetworkPacketCapture_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkPacketCapture_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCapture",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkPacketCapture_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkPacketCapture_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCapture",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkPacketCapture_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkPacketCapture_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCapture",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkPacketCapture_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCapture",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) PutFilter(value interface{}) {
	if err := n.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFilter",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) PutStorageLocation(value *NetworkPacketCaptureStorageLocation) {
	if err := n.validatePutStorageLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStorageLocation",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) PutTimeouts(value *NetworkPacketCaptureTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPacketCapture) ResetFilter() {
	_jsii_.InvokeVoid(
		n,
		"resetFilter",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCapture) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCapture) ResetMaximumBytesPerPacket() {
	_jsii_.InvokeVoid(
		n,
		"resetMaximumBytesPerPacket",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCapture) ResetMaximumBytesPerSession() {
	_jsii_.InvokeVoid(
		n,
		"resetMaximumBytesPerSession",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCapture) ResetMaximumCaptureDuration() {
	_jsii_.InvokeVoid(
		n,
		"resetMaximumCaptureDuration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCapture) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCapture) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCapture) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCapture) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

