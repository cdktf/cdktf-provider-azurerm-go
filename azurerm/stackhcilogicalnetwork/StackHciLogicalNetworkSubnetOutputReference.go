// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcilogicalnetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/stackhcilogicalnetwork/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackHciLogicalNetworkSubnetOutputReference interface {
	cdktf.ComplexObject
	AddressPrefix() *string
	SetAddressPrefix(val *string)
	AddressPrefixInput() *string
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
	InternalValue() *StackHciLogicalNetworkSubnet
	SetInternalValue(val *StackHciLogicalNetworkSubnet)
	IpAllocationMethod() *string
	SetIpAllocationMethod(val *string)
	IpAllocationMethodInput() *string
	IpPool() StackHciLogicalNetworkSubnetIpPoolList
	IpPoolInput() interface{}
	Route() StackHciLogicalNetworkSubnetRouteOutputReference
	RouteInput() *StackHciLogicalNetworkSubnetRoute
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VlanId() *float64
	SetVlanId(val *float64)
	VlanIdInput() *float64
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
	PutIpPool(value interface{})
	PutRoute(value *StackHciLogicalNetworkSubnetRoute)
	ResetAddressPrefix()
	ResetIpPool()
	ResetRoute()
	ResetVlanId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StackHciLogicalNetworkSubnetOutputReference
type jsiiProxy_StackHciLogicalNetworkSubnetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) AddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) AddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) InternalValue() *StackHciLogicalNetworkSubnet {
	var returns *StackHciLogicalNetworkSubnet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) IpAllocationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAllocationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) IpAllocationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAllocationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) IpPool() StackHciLogicalNetworkSubnetIpPoolList {
	var returns StackHciLogicalNetworkSubnetIpPoolList
	_jsii_.Get(
		j,
		"ipPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) IpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) Route() StackHciLogicalNetworkSubnetRouteOutputReference {
	var returns StackHciLogicalNetworkSubnetRouteOutputReference
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) RouteInput() *StackHciLogicalNetworkSubnetRoute {
	var returns *StackHciLogicalNetworkSubnetRoute
	_jsii_.Get(
		j,
		"routeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) VlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) VlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanIdInput",
		&returns,
	)
	return returns
}


func NewStackHciLogicalNetworkSubnetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StackHciLogicalNetworkSubnetOutputReference {
	_init_.Initialize()

	if err := validateNewStackHciLogicalNetworkSubnetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackHciLogicalNetworkSubnetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciLogicalNetwork.StackHciLogicalNetworkSubnetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStackHciLogicalNetworkSubnetOutputReference_Override(s StackHciLogicalNetworkSubnetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciLogicalNetwork.StackHciLogicalNetworkSubnetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference)SetAddressPrefix(val *string) {
	if err := j.validateSetAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressPrefix",
		val,
	)
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference)SetInternalValue(val *StackHciLogicalNetworkSubnet) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference)SetIpAllocationMethod(val *string) {
	if err := j.validateSetIpAllocationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAllocationMethod",
		val,
	)
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference)SetVlanId(val *float64) {
	if err := j.validateSetVlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlanId",
		val,
	)
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) PutIpPool(value interface{}) {
	if err := s.validatePutIpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIpPool",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) PutRoute(value *StackHciLogicalNetworkSubnetRoute) {
	if err := s.validatePutRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRoute",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) ResetAddressPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetAddressPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) ResetIpPool() {
	_jsii_.InvokeVoid(
		s,
		"resetIpPool",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) ResetRoute() {
	_jsii_.InvokeVoid(
		s,
		"resetRoute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) ResetVlanId() {
	_jsii_.InvokeVoid(
		s,
		"resetVlanId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StackHciLogicalNetworkSubnetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

