// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/redhatopenshiftcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedhatOpenshiftClusterNetworkProfileOutputReference interface {
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
	InternalValue() *RedhatOpenshiftClusterNetworkProfile
	SetInternalValue(val *RedhatOpenshiftClusterNetworkProfile)
	OutboundType() *string
	SetOutboundType(val *string)
	OutboundTypeInput() *string
	PodCidr() *string
	SetPodCidr(val *string)
	PodCidrInput() *string
	PreconfiguredNetworkSecurityGroupEnabled() interface{}
	SetPreconfiguredNetworkSecurityGroupEnabled(val interface{})
	PreconfiguredNetworkSecurityGroupEnabledInput() interface{}
	ServiceCidr() *string
	SetServiceCidr(val *string)
	ServiceCidrInput() *string
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
	ResetOutboundType()
	ResetPreconfiguredNetworkSecurityGroupEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedhatOpenshiftClusterNetworkProfileOutputReference
type jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) InternalValue() *RedhatOpenshiftClusterNetworkProfile {
	var returns *RedhatOpenshiftClusterNetworkProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) OutboundType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) OutboundTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) PodCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) PodCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) PreconfiguredNetworkSecurityGroupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preconfiguredNetworkSecurityGroupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) PreconfiguredNetworkSecurityGroupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preconfiguredNetworkSecurityGroupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) ServiceCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) ServiceCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRedhatOpenshiftClusterNetworkProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedhatOpenshiftClusterNetworkProfileOutputReference {
	_init_.Initialize()

	if err := validateNewRedhatOpenshiftClusterNetworkProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftClusterNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedhatOpenshiftClusterNetworkProfileOutputReference_Override(r RedhatOpenshiftClusterNetworkProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftClusterNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetInternalValue(val *RedhatOpenshiftClusterNetworkProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetOutboundType(val *string) {
	if err := j.validateSetOutboundTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundType",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetPodCidr(val *string) {
	if err := j.validateSetPodCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podCidr",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetPreconfiguredNetworkSecurityGroupEnabled(val interface{}) {
	if err := j.validateSetPreconfiguredNetworkSecurityGroupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconfiguredNetworkSecurityGroupEnabled",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetServiceCidr(val *string) {
	if err := j.validateSetServiceCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceCidr",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) ResetOutboundType() {
	_jsii_.InvokeVoid(
		r,
		"resetOutboundType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) ResetPreconfiguredNetworkSecurityGroupEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetPreconfiguredNetworkSecurityGroupEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

