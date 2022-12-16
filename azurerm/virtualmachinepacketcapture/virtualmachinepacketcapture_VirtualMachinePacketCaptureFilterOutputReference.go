package virtualmachinepacketcapture

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/virtualmachinepacketcapture/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachinePacketCaptureFilterOutputReference interface {
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
	LocalIpAddress() *string
	SetLocalIpAddress(val *string)
	LocalIpAddressInput() *string
	LocalPort() *string
	SetLocalPort(val *string)
	LocalPortInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	RemoteIpAddress() *string
	SetRemoteIpAddress(val *string)
	RemoteIpAddressInput() *string
	RemotePort() *string
	SetRemotePort(val *string)
	RemotePortInput() *string
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
	ResetLocalIpAddress()
	ResetLocalPort()
	ResetRemoteIpAddress()
	ResetRemotePort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachinePacketCaptureFilterOutputReference
type jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) LocalIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) LocalIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) LocalPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) LocalPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) RemoteIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) RemoteIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) RemotePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remotePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) RemotePortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remotePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualMachinePacketCaptureFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualMachinePacketCaptureFilterOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachinePacketCaptureFilterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCaptureFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualMachinePacketCaptureFilterOutputReference_Override(v VirtualMachinePacketCaptureFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachinePacketCapture.VirtualMachinePacketCaptureFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetLocalIpAddress(val *string) {
	if err := j.validateSetLocalIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpAddress",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetLocalPort(val *string) {
	if err := j.validateSetLocalPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPort",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetRemoteIpAddress(val *string) {
	if err := j.validateSetRemoteIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpAddress",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetRemotePort(val *string) {
	if err := j.validateSetRemotePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remotePort",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ResetLocalIpAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalIpAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ResetLocalPort() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ResetRemoteIpAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetRemoteIpAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ResetRemotePort() {
	_jsii_.InvokeVoid(
		v,
		"resetRemotePort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachinePacketCaptureFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

