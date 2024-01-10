// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/hdinsightinteractivequerycluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference interface {
	cdktf.ComplexObject
	Capacity() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference
	CapacityInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity
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
	InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale
	SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale)
	Recurrence() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference
	RecurrenceInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence
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
	PutCapacity(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity)
	PutRecurrence(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence)
	ResetCapacity()
	ResetRecurrence()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) Capacity() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) CapacityInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) Recurrence() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference
	_jsii_.Get(
		j,
		"recurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) RecurrenceInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence
	_jsii_.Get(
		j,
		"recurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference {
	_init_.Initialize()

	if err := validateNewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference_Override(h HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference)SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) PutCapacity(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity) {
	if err := h.validatePutCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putCapacity",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) PutRecurrence(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence) {
	if err := h.validatePutRecurrenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putRecurrence",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ResetCapacity() {
	_jsii_.InvokeVoid(
		h,
		"resetCapacity",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ResetRecurrence() {
	_jsii_.InvokeVoid(
		h,
		"resetRecurrence",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

