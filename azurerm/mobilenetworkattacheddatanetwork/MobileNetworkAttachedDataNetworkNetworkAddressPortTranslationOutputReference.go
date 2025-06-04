// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkattacheddatanetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/mobilenetworkattacheddatanetwork/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference interface {
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
	IcmpPinholeTimeoutInSeconds() *float64
	SetIcmpPinholeTimeoutInSeconds(val *float64)
	IcmpPinholeTimeoutInSecondsInput() *float64
	InternalValue() *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation
	SetInternalValue(val *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation)
	PinholeMaximumNumber() *float64
	SetPinholeMaximumNumber(val *float64)
	PinholeMaximumNumberInput() *float64
	PortRange() MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRangeOutputReference
	PortRangeInput() *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRange
	TcpPinholeTimeoutInSeconds() *float64
	SetTcpPinholeTimeoutInSeconds(val *float64)
	TcpPinholeTimeoutInSecondsInput() *float64
	TcpPortReuseMinimumHoldTimeInSeconds() *float64
	SetTcpPortReuseMinimumHoldTimeInSeconds(val *float64)
	TcpPortReuseMinimumHoldTimeInSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UdpPinholeTimeoutInSeconds() *float64
	SetUdpPinholeTimeoutInSeconds(val *float64)
	UdpPinholeTimeoutInSecondsInput() *float64
	UdpPortReuseMinimumHoldTimeInSeconds() *float64
	SetUdpPortReuseMinimumHoldTimeInSeconds(val *float64)
	UdpPortReuseMinimumHoldTimeInSecondsInput() *float64
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
	PutPortRange(value *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRange)
	ResetIcmpPinholeTimeoutInSeconds()
	ResetPinholeMaximumNumber()
	ResetPortRange()
	ResetTcpPinholeTimeoutInSeconds()
	ResetTcpPortReuseMinimumHoldTimeInSeconds()
	ResetUdpPinholeTimeoutInSeconds()
	ResetUdpPortReuseMinimumHoldTimeInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference
type jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) IcmpPinholeTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpPinholeTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) IcmpPinholeTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpPinholeTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) InternalValue() *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation {
	var returns *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) PinholeMaximumNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pinholeMaximumNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) PinholeMaximumNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pinholeMaximumNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) PortRange() MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRangeOutputReference {
	var returns MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRangeOutputReference
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) PortRangeInput() *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRange {
	var returns *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRange
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) TcpPinholeTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpPinholeTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) TcpPinholeTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpPinholeTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) TcpPortReuseMinimumHoldTimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpPortReuseMinimumHoldTimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) TcpPortReuseMinimumHoldTimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpPortReuseMinimumHoldTimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) UdpPinholeTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpPinholeTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) UdpPinholeTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpPinholeTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) UdpPortReuseMinimumHoldTimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpPortReuseMinimumHoldTimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) UdpPortReuseMinimumHoldTimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpPortReuseMinimumHoldTimeInSecondsInput",
		&returns,
	)
	return returns
}


func NewMobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference {
	_init_.Initialize()

	if err := validateNewMobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference_Override(m MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetIcmpPinholeTimeoutInSeconds(val *float64) {
	if err := j.validateSetIcmpPinholeTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpPinholeTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetInternalValue(val *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetPinholeMaximumNumber(val *float64) {
	if err := j.validateSetPinholeMaximumNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinholeMaximumNumber",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetTcpPinholeTimeoutInSeconds(val *float64) {
	if err := j.validateSetTcpPinholeTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpPinholeTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetTcpPortReuseMinimumHoldTimeInSeconds(val *float64) {
	if err := j.validateSetTcpPortReuseMinimumHoldTimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpPortReuseMinimumHoldTimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetUdpPinholeTimeoutInSeconds(val *float64) {
	if err := j.validateSetUdpPinholeTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpPinholeTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference)SetUdpPortReuseMinimumHoldTimeInSeconds(val *float64) {
	if err := j.validateSetUdpPortReuseMinimumHoldTimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpPortReuseMinimumHoldTimeInSeconds",
		val,
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) PutPortRange(value *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRange) {
	if err := m.validatePutPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPortRange",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ResetIcmpPinholeTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetIcmpPinholeTimeoutInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ResetPinholeMaximumNumber() {
	_jsii_.InvokeVoid(
		m,
		"resetPinholeMaximumNumber",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		m,
		"resetPortRange",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ResetTcpPinholeTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTcpPinholeTimeoutInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ResetTcpPortReuseMinimumHoldTimeInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTcpPortReuseMinimumHoldTimeInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ResetUdpPinholeTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetUdpPinholeTimeoutInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ResetUdpPortReuseMinimumHoldTimeInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetUdpPortReuseMinimumHoldTimeInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

