package networkpacketcapture

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/networkpacketcapture/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkPacketCaptureFilterOutputReference interface {
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

// The jsii proxy struct for NetworkPacketCaptureFilterOutputReference
type jsiiProxy_NetworkPacketCaptureFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) LocalIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) LocalIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) LocalPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) LocalPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) RemoteIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) RemoteIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) RemotePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remotePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) RemotePortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remotePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkPacketCaptureFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkPacketCaptureFilterOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkPacketCaptureFilterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkPacketCaptureFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCaptureFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkPacketCaptureFilterOutputReference_Override(n NetworkPacketCaptureFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkPacketCapture.NetworkPacketCaptureFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetLocalIpAddress(val *string) {
	if err := j.validateSetLocalIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpAddress",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetLocalPort(val *string) {
	if err := j.validateSetLocalPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPort",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetRemoteIpAddress(val *string) {
	if err := j.validateSetRemoteIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpAddress",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetRemotePort(val *string) {
	if err := j.validateSetRemotePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remotePort",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkPacketCaptureFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ResetLocalIpAddress() {
	_jsii_.InvokeVoid(
		n,
		"resetLocalIpAddress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ResetLocalPort() {
	_jsii_.InvokeVoid(
		n,
		"resetLocalPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ResetRemoteIpAddress() {
	_jsii_.InvokeVoid(
		n,
		"resetRemoteIpAddress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ResetRemotePort() {
	_jsii_.InvokeVoid(
		n,
		"resetRemotePort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPacketCaptureFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

