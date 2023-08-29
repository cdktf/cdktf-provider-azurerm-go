// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/batchpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchPoolFixedScaleOutputReference interface {
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
	InternalValue() *BatchPoolFixedScale
	SetInternalValue(val *BatchPoolFixedScale)
	NodeDeallocationMethod() *string
	SetNodeDeallocationMethod(val *string)
	NodeDeallocationMethodInput() *string
	ResizeTimeout() *string
	SetResizeTimeout(val *string)
	ResizeTimeoutInput() *string
	TargetDedicatedNodes() *float64
	SetTargetDedicatedNodes(val *float64)
	TargetDedicatedNodesInput() *float64
	TargetLowPriorityNodes() *float64
	SetTargetLowPriorityNodes(val *float64)
	TargetLowPriorityNodesInput() *float64
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
	ResetNodeDeallocationMethod()
	ResetResizeTimeout()
	ResetTargetDedicatedNodes()
	ResetTargetLowPriorityNodes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolFixedScaleOutputReference
type jsiiProxy_BatchPoolFixedScaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) InternalValue() *BatchPoolFixedScale {
	var returns *BatchPoolFixedScale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) NodeDeallocationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeDeallocationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) NodeDeallocationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeDeallocationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) ResizeTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resizeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) ResizeTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resizeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TargetDedicatedNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetDedicatedNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TargetDedicatedNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetDedicatedNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TargetLowPriorityNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetLowPriorityNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TargetLowPriorityNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetLowPriorityNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolFixedScaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolFixedScaleOutputReference {
	_init_.Initialize()

	if err := validateNewBatchPoolFixedScaleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchPoolFixedScaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolFixedScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolFixedScaleOutputReference_Override(b BatchPoolFixedScaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolFixedScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetInternalValue(val *BatchPoolFixedScale) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetNodeDeallocationMethod(val *string) {
	if err := j.validateSetNodeDeallocationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeDeallocationMethod",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetResizeTimeout(val *string) {
	if err := j.validateSetResizeTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resizeTimeout",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetTargetDedicatedNodes(val *float64) {
	if err := j.validateSetTargetDedicatedNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDedicatedNodes",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetTargetLowPriorityNodes(val *float64) {
	if err := j.validateSetTargetLowPriorityNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetLowPriorityNodes",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolFixedScaleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ResetNodeDeallocationMethod() {
	_jsii_.InvokeVoid(
		b,
		"resetNodeDeallocationMethod",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ResetResizeTimeout() {
	_jsii_.InvokeVoid(
		b,
		"resetResizeTimeout",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ResetTargetDedicatedNodes() {
	_jsii_.InvokeVoid(
		b,
		"resetTargetDedicatedNodes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ResetTargetLowPriorityNodes() {
	_jsii_.InvokeVoid(
		b,
		"resetTargetLowPriorityNodes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BatchPoolFixedScaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

