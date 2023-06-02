package virtualmachinescalesetpacketcapture

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/virtualmachinescalesetpacketcapture/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_machine_scale_set_packet_capture azurerm_virtual_machine_scale_set_packet_capture}.
type VirtualMachineScaleSetPacketCapture interface {
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
	Filter() VirtualMachineScaleSetPacketCaptureFilterList
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
	MachineScope() VirtualMachineScaleSetPacketCaptureMachineScopeOutputReference
	MachineScopeInput() *VirtualMachineScaleSetPacketCaptureMachineScope
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
	StorageLocation() VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference
	StorageLocationInput() *VirtualMachineScaleSetPacketCaptureStorageLocation
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualMachineScaleSetPacketCaptureTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualMachineScaleSetId() *string
	SetVirtualMachineScaleSetId(val *string)
	VirtualMachineScaleSetIdInput() *string
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
	PutFilter(value interface{})
	PutMachineScope(value *VirtualMachineScaleSetPacketCaptureMachineScope)
	PutStorageLocation(value *VirtualMachineScaleSetPacketCaptureStorageLocation)
	PutTimeouts(value *VirtualMachineScaleSetPacketCaptureTimeouts)
	ResetFilter()
	ResetId()
	ResetMachineScope()
	ResetMaximumBytesPerPacket()
	ResetMaximumBytesPerSession()
	ResetMaximumCaptureDurationInSeconds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VirtualMachineScaleSetPacketCapture
type jsiiProxy_VirtualMachineScaleSetPacketCapture struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Filter() VirtualMachineScaleSetPacketCaptureFilterList {
	var returns VirtualMachineScaleSetPacketCaptureFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) MachineScope() VirtualMachineScaleSetPacketCaptureMachineScopeOutputReference {
	var returns VirtualMachineScaleSetPacketCaptureMachineScopeOutputReference
	_jsii_.Get(
		j,
		"machineScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) MachineScopeInput() *VirtualMachineScaleSetPacketCaptureMachineScope {
	var returns *VirtualMachineScaleSetPacketCaptureMachineScope
	_jsii_.Get(
		j,
		"machineScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) MaximumBytesPerPacket() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerPacket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) MaximumBytesPerPacketInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerPacketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) MaximumBytesPerSession() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) MaximumBytesPerSessionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBytesPerSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) MaximumCaptureDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCaptureDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) MaximumCaptureDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCaptureDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) NetworkWatcherId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkWatcherId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) NetworkWatcherIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkWatcherIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) StorageLocation() VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference {
	var returns VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) StorageLocationInput() *VirtualMachineScaleSetPacketCaptureStorageLocation {
	var returns *VirtualMachineScaleSetPacketCaptureStorageLocation
	_jsii_.Get(
		j,
		"storageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) Timeouts() VirtualMachineScaleSetPacketCaptureTimeoutsOutputReference {
	var returns VirtualMachineScaleSetPacketCaptureTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) VirtualMachineScaleSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture) VirtualMachineScaleSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_machine_scale_set_packet_capture azurerm_virtual_machine_scale_set_packet_capture} Resource.
func NewVirtualMachineScaleSetPacketCapture(scope constructs.Construct, id *string, config *VirtualMachineScaleSetPacketCaptureConfig) VirtualMachineScaleSetPacketCapture {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetPacketCaptureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSetPacketCapture{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSetPacketCapture.VirtualMachineScaleSetPacketCapture",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_machine_scale_set_packet_capture azurerm_virtual_machine_scale_set_packet_capture} Resource.
func NewVirtualMachineScaleSetPacketCapture_Override(v VirtualMachineScaleSetPacketCapture, scope constructs.Construct, id *string, config *VirtualMachineScaleSetPacketCaptureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSetPacketCapture.VirtualMachineScaleSetPacketCapture",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetMaximumBytesPerPacket(val *float64) {
	if err := j.validateSetMaximumBytesPerPacketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBytesPerPacket",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetMaximumBytesPerSession(val *float64) {
	if err := j.validateSetMaximumBytesPerSessionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBytesPerSession",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetMaximumCaptureDurationInSeconds(val *float64) {
	if err := j.validateSetMaximumCaptureDurationInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCaptureDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetNetworkWatcherId(val *string) {
	if err := j.validateSetNetworkWatcherIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkWatcherId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCapture)SetVirtualMachineScaleSetId(val *string) {
	if err := j.validateSetVirtualMachineScaleSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineScaleSetId",
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
func VirtualMachineScaleSetPacketCapture_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSetPacketCapture_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineScaleSetPacketCapture.VirtualMachineScaleSetPacketCapture",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachineScaleSetPacketCapture_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSetPacketCapture_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineScaleSetPacketCapture.VirtualMachineScaleSetPacketCapture",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachineScaleSetPacketCapture_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSetPacketCapture_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineScaleSetPacketCapture.VirtualMachineScaleSetPacketCapture",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualMachineScaleSetPacketCapture_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualMachineScaleSetPacketCapture.VirtualMachineScaleSetPacketCapture",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) PutFilter(value interface{}) {
	if err := v.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putFilter",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) PutMachineScope(value *VirtualMachineScaleSetPacketCaptureMachineScope) {
	if err := v.validatePutMachineScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMachineScope",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) PutStorageLocation(value *VirtualMachineScaleSetPacketCaptureStorageLocation) {
	if err := v.validatePutStorageLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageLocation",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) PutTimeouts(value *VirtualMachineScaleSetPacketCaptureTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ResetFilter() {
	_jsii_.InvokeVoid(
		v,
		"resetFilter",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ResetMachineScope() {
	_jsii_.InvokeVoid(
		v,
		"resetMachineScope",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ResetMaximumBytesPerPacket() {
	_jsii_.InvokeVoid(
		v,
		"resetMaximumBytesPerPacket",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ResetMaximumBytesPerSession() {
	_jsii_.InvokeVoid(
		v,
		"resetMaximumBytesPerSession",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ResetMaximumCaptureDurationInSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetMaximumCaptureDurationInSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCapture) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

