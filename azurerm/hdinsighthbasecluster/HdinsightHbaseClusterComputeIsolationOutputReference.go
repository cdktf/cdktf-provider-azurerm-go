// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthbasecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/hdinsighthbasecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightHbaseClusterComputeIsolationOutputReference interface {
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
	ComputeIsolationEnabled() interface{}
	SetComputeIsolationEnabled(val interface{})
	ComputeIsolationEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HostSku() *string
	SetHostSku(val *string)
	HostSkuInput() *string
	InternalValue() *HdinsightHbaseClusterComputeIsolation
	SetInternalValue(val *HdinsightHbaseClusterComputeIsolation)
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
	ResetComputeIsolationEnabled()
	ResetHostSku()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHbaseClusterComputeIsolationOutputReference
type jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) ComputeIsolationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeIsolationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) ComputeIsolationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeIsolationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) HostSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) HostSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) InternalValue() *HdinsightHbaseClusterComputeIsolation {
	var returns *HdinsightHbaseClusterComputeIsolation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHbaseClusterComputeIsolationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHbaseClusterComputeIsolationOutputReference {
	_init_.Initialize()

	if err := validateNewHdinsightHbaseClusterComputeIsolationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHbaseCluster.HdinsightHbaseClusterComputeIsolationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHbaseClusterComputeIsolationOutputReference_Override(h HdinsightHbaseClusterComputeIsolationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHbaseCluster.HdinsightHbaseClusterComputeIsolationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference)SetComputeIsolationEnabled(val interface{}) {
	if err := j.validateSetComputeIsolationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeIsolationEnabled",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference)SetHostSku(val *string) {
	if err := j.validateSetHostSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostSku",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference)SetInternalValue(val *HdinsightHbaseClusterComputeIsolation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) ResetComputeIsolationEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetComputeIsolationEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) ResetHostSku() {
	_jsii_.InvokeVoid(
		h,
		"resetHostSku",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterComputeIsolationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

