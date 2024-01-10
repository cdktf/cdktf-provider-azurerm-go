// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/batchpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchPoolNetworkConfigurationOutputReference interface {
	cdktf.ComplexObject
	AcceleratedNetworkingEnabled() interface{}
	SetAcceleratedNetworkingEnabled(val interface{})
	AcceleratedNetworkingEnabledInput() interface{}
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
	DynamicVnetAssignmentScope() *string
	SetDynamicVnetAssignmentScope(val *string)
	DynamicVnetAssignmentScopeInput() *string
	EndpointConfiguration() BatchPoolNetworkConfigurationEndpointConfigurationList
	EndpointConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *BatchPoolNetworkConfiguration
	SetInternalValue(val *BatchPoolNetworkConfiguration)
	PublicAddressProvisioningType() *string
	SetPublicAddressProvisioningType(val *string)
	PublicAddressProvisioningTypeInput() *string
	PublicIps() *[]*string
	SetPublicIps(val *[]*string)
	PublicIpsInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	PutEndpointConfiguration(value interface{})
	ResetAcceleratedNetworkingEnabled()
	ResetDynamicVnetAssignmentScope()
	ResetEndpointConfiguration()
	ResetPublicAddressProvisioningType()
	ResetPublicIps()
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolNetworkConfigurationOutputReference
type jsiiProxy_BatchPoolNetworkConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) AcceleratedNetworkingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratedNetworkingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) AcceleratedNetworkingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratedNetworkingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) DynamicVnetAssignmentScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamicVnetAssignmentScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) DynamicVnetAssignmentScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamicVnetAssignmentScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) EndpointConfiguration() BatchPoolNetworkConfigurationEndpointConfigurationList {
	var returns BatchPoolNetworkConfigurationEndpointConfigurationList
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) EndpointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) InternalValue() *BatchPoolNetworkConfiguration {
	var returns *BatchPoolNetworkConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PublicAddressProvisioningType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicAddressProvisioningType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PublicAddressProvisioningTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicAddressProvisioningTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PublicIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PublicIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolNetworkConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolNetworkConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBatchPoolNetworkConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchPoolNetworkConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolNetworkConfigurationOutputReference_Override(b BatchPoolNetworkConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetAcceleratedNetworkingEnabled(val interface{}) {
	if err := j.validateSetAcceleratedNetworkingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratedNetworkingEnabled",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetDynamicVnetAssignmentScope(val *string) {
	if err := j.validateSetDynamicVnetAssignmentScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicVnetAssignmentScope",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetInternalValue(val *BatchPoolNetworkConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetPublicAddressProvisioningType(val *string) {
	if err := j.validateSetPublicAddressProvisioningTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAddressProvisioningType",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetPublicIps(val *[]*string) {
	if err := j.validateSetPublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIps",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) PutEndpointConfiguration(value interface{}) {
	if err := b.validatePutEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetAcceleratedNetworkingEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetAcceleratedNetworkingEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetDynamicVnetAssignmentScope() {
	_jsii_.InvokeVoid(
		b,
		"resetDynamicVnetAssignmentScope",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetPublicAddressProvisioningType() {
	_jsii_.InvokeVoid(
		b,
		"resetPublicAddressProvisioningType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetPublicIps() {
	_jsii_.InvokeVoid(
		b,
		"resetPublicIps",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		b,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

