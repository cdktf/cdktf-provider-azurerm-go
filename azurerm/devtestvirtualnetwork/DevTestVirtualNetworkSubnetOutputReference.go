// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestvirtualnetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/devtestvirtualnetwork/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevTestVirtualNetworkSubnetOutputReference interface {
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
	InternalValue() *DevTestVirtualNetworkSubnet
	SetInternalValue(val *DevTestVirtualNetworkSubnet)
	Name() *string
	SharedPublicIpAddress() DevTestVirtualNetworkSubnetSharedPublicIpAddressOutputReference
	SharedPublicIpAddressInput() *DevTestVirtualNetworkSubnetSharedPublicIpAddress
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseInVirtualMachineCreation() *string
	SetUseInVirtualMachineCreation(val *string)
	UseInVirtualMachineCreationInput() *string
	UsePublicIpAddress() *string
	SetUsePublicIpAddress(val *string)
	UsePublicIpAddressInput() *string
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
	PutSharedPublicIpAddress(value *DevTestVirtualNetworkSubnetSharedPublicIpAddress)
	ResetSharedPublicIpAddress()
	ResetUseInVirtualMachineCreation()
	ResetUsePublicIpAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DevTestVirtualNetworkSubnetOutputReference
type jsiiProxy_DevTestVirtualNetworkSubnetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) InternalValue() *DevTestVirtualNetworkSubnet {
	var returns *DevTestVirtualNetworkSubnet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) SharedPublicIpAddress() DevTestVirtualNetworkSubnetSharedPublicIpAddressOutputReference {
	var returns DevTestVirtualNetworkSubnetSharedPublicIpAddressOutputReference
	_jsii_.Get(
		j,
		"sharedPublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) SharedPublicIpAddressInput() *DevTestVirtualNetworkSubnetSharedPublicIpAddress {
	var returns *DevTestVirtualNetworkSubnetSharedPublicIpAddress
	_jsii_.Get(
		j,
		"sharedPublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) UseInVirtualMachineCreation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useInVirtualMachineCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) UseInVirtualMachineCreationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useInVirtualMachineCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) UsePublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) UsePublicIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePublicIpAddressInput",
		&returns,
	)
	return returns
}


func NewDevTestVirtualNetworkSubnetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DevTestVirtualNetworkSubnetOutputReference {
	_init_.Initialize()

	if err := validateNewDevTestVirtualNetworkSubnetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevTestVirtualNetworkSubnetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestVirtualNetwork.DevTestVirtualNetworkSubnetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDevTestVirtualNetworkSubnetOutputReference_Override(d DevTestVirtualNetworkSubnetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.devTestVirtualNetwork.DevTestVirtualNetworkSubnetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference)SetInternalValue(val *DevTestVirtualNetworkSubnet) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference)SetUseInVirtualMachineCreation(val *string) {
	if err := j.validateSetUseInVirtualMachineCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useInVirtualMachineCreation",
		val,
	)
}

func (j *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference)SetUsePublicIpAddress(val *string) {
	if err := j.validateSetUsePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePublicIpAddress",
		val,
	)
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) PutSharedPublicIpAddress(value *DevTestVirtualNetworkSubnetSharedPublicIpAddress) {
	if err := d.validatePutSharedPublicIpAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSharedPublicIpAddress",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) ResetSharedPublicIpAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetSharedPublicIpAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) ResetUseInVirtualMachineCreation() {
	_jsii_.InvokeVoid(
		d,
		"resetUseInVirtualMachineCreation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) ResetUsePublicIpAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetUsePublicIpAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevTestVirtualNetworkSubnetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

