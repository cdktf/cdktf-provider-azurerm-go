// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference interface {
	cdktf.ComplexObject
	Complexity() *string
	SetComplexity(val *string)
	ComplexityInput() *string
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
	InterleaveOutput() *string
	SetInterleaveOutput(val *string)
	InterleaveOutputInput() *string
	InternalValue() *MediaTransformOutputBuiltinPresetPresetConfiguration
	SetInternalValue(val *MediaTransformOutputBuiltinPresetPresetConfiguration)
	KeyFrameIntervalInSeconds() *float64
	SetKeyFrameIntervalInSeconds(val *float64)
	KeyFrameIntervalInSecondsInput() *float64
	MaxBitrateBps() *float64
	SetMaxBitrateBps(val *float64)
	MaxBitrateBpsInput() *float64
	MaxHeight() *float64
	SetMaxHeight(val *float64)
	MaxHeightInput() *float64
	MaxLayers() *float64
	SetMaxLayers(val *float64)
	MaxLayersInput() *float64
	MinBitrateBps() *float64
	SetMinBitrateBps(val *float64)
	MinBitrateBpsInput() *float64
	MinHeight() *float64
	SetMinHeight(val *float64)
	MinHeightInput() *float64
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
	ResetComplexity()
	ResetInterleaveOutput()
	ResetKeyFrameIntervalInSeconds()
	ResetMaxBitrateBps()
	ResetMaxHeight()
	ResetMaxLayers()
	ResetMinBitrateBps()
	ResetMinHeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference
type jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) Complexity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ComplexityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) InterleaveOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interleaveOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) InterleaveOutputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interleaveOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) InternalValue() *MediaTransformOutputBuiltinPresetPresetConfiguration {
	var returns *MediaTransformOutputBuiltinPresetPresetConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) KeyFrameIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyFrameIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) KeyFrameIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyFrameIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MaxBitrateBps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateBps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MaxBitrateBpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateBpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MaxHeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MaxHeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MaxLayers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLayers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MaxLayersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLayersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MinBitrateBps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBitrateBps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MinBitrateBpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minBitrateBpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MinHeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) MinHeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputBuiltinPresetPresetConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputBuiltinPresetPresetConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaTransformOutputBuiltinPresetPresetConfigurationOutputReference_Override(m MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetComplexity(val *string) {
	if err := j.validateSetComplexityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexity",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetInterleaveOutput(val *string) {
	if err := j.validateSetInterleaveOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interleaveOutput",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetInternalValue(val *MediaTransformOutputBuiltinPresetPresetConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetKeyFrameIntervalInSeconds(val *float64) {
	if err := j.validateSetKeyFrameIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFrameIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetMaxBitrateBps(val *float64) {
	if err := j.validateSetMaxBitrateBpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBitrateBps",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetMaxHeight(val *float64) {
	if err := j.validateSetMaxHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeight",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetMaxLayers(val *float64) {
	if err := j.validateSetMaxLayersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLayers",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetMinBitrateBps(val *float64) {
	if err := j.validateSetMinBitrateBpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minBitrateBps",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetMinHeight(val *float64) {
	if err := j.validateSetMinHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minHeight",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ResetComplexity() {
	_jsii_.InvokeVoid(
		m,
		"resetComplexity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ResetInterleaveOutput() {
	_jsii_.InvokeVoid(
		m,
		"resetInterleaveOutput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ResetKeyFrameIntervalInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyFrameIntervalInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ResetMaxBitrateBps() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBitrateBps",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ResetMaxHeight() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxHeight",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ResetMaxLayers() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxLayers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ResetMinBitrateBps() {
	_jsii_.InvokeVoid(
		m,
		"resetMinBitrateBps",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ResetMinHeight() {
	_jsii_.InvokeVoid(
		m,
		"resetMinHeight",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputBuiltinPresetPresetConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

