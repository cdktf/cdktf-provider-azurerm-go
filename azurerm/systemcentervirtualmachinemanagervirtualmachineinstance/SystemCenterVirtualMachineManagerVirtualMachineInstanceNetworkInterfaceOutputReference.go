// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualmachineinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/systemcentervirtualmachinemanagervirtualmachineinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference interface {
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
	Ipv4AddressType() *string
	SetIpv4AddressType(val *string)
	Ipv4AddressTypeInput() *string
	Ipv6AddressType() *string
	SetIpv6AddressType(val *string)
	Ipv6AddressTypeInput() *string
	MacAddressType() *string
	SetMacAddressType(val *string)
	MacAddressTypeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetIpv4AddressType()
	ResetIpv6AddressType()
	ResetMacAddressType()
	ResetVirtualNetworkId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference
type jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) Ipv4AddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4AddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) Ipv4AddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4AddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) Ipv6AddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) Ipv6AddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) MacAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) MacAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}


func NewSystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewSystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference_Override(s SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetIpv4AddressType(val *string) {
	if err := j.validateSetIpv4AddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4AddressType",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetIpv6AddressType(val *string) {
	if err := j.validateSetIpv6AddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressType",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetMacAddressType(val *string) {
	if err := j.validateSetMacAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macAddressType",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) ResetIpv4AddressType() {
	_jsii_.InvokeVoid(
		s,
		"resetIpv4AddressType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) ResetIpv6AddressType() {
	_jsii_.InvokeVoid(
		s,
		"resetIpv6AddressType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) ResetMacAddressType() {
	_jsii_.InvokeVoid(
		s,
		"resetMacAddressType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		s,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

