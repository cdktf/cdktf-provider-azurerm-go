// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference interface {
	cdktf.ComplexObject
	AdaptiveBFrameEnabled() interface{}
	SetAdaptiveBFrameEnabled(val interface{})
	AdaptiveBFrameEnabledInput() interface{}
	BFrames() *float64
	SetBFrames(val *float64)
	BFramesInput() *float64
	Bitrate() *float64
	SetBitrate(val *float64)
	BitrateInput() *float64
	BufferWindow() *string
	SetBufferWindow(val *string)
	BufferWindowInput() *string
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
	Crf() *float64
	SetCrf(val *float64)
	CrfInput() *float64
	// Experimental.
	Fqn() *string
	FrameRate() *string
	SetFrameRate(val *string)
	FrameRateInput() *string
	Height() *string
	SetHeight(val *string)
	HeightInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	MaxBitrate() *float64
	SetMaxBitrate(val *float64)
	MaxBitrateInput() *float64
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	ReferenceFrames() *float64
	SetReferenceFrames(val *float64)
	ReferenceFramesInput() *float64
	Slices() *float64
	SetSlices(val *float64)
	SlicesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Width() *string
	SetWidth(val *string)
	WidthInput() *string
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
	ResetAdaptiveBFrameEnabled()
	ResetBFrames()
	ResetBufferWindow()
	ResetCrf()
	ResetFrameRate()
	ResetHeight()
	ResetLabel()
	ResetLevel()
	ResetMaxBitrate()
	ResetProfile()
	ResetReferenceFrames()
	ResetSlices()
	ResetWidth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference
type jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) AdaptiveBFrameEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adaptiveBFrameEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) AdaptiveBFrameEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adaptiveBFrameEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) BFrames() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bFrames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) BFramesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bFramesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) BufferWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bufferWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) BufferWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bufferWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Crf() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) CrfInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) FrameRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) FrameRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Height() *string {
	var returns *string
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) HeightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) MaxBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ReferenceFrames() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceFrames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ReferenceFramesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceFramesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Slices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) SlicesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Width() *string {
	var returns *string
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) WidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputCustomPresetCodecH265VideoLayerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference_Override(m MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetAdaptiveBFrameEnabled(val interface{}) {
	if err := j.validateSetAdaptiveBFrameEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adaptiveBFrameEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetBFrames(val *float64) {
	if err := j.validateSetBFramesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bFrames",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetBufferWindow(val *string) {
	if err := j.validateSetBufferWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferWindow",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetCrf(val *float64) {
	if err := j.validateSetCrfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crf",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetFrameRate(val *string) {
	if err := j.validateSetFrameRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameRate",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetHeight(val *string) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetMaxBitrate(val *float64) {
	if err := j.validateSetMaxBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetReferenceFrames(val *float64) {
	if err := j.validateSetReferenceFramesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceFrames",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetSlices(val *float64) {
	if err := j.validateSetSlicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slices",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference)SetWidth(val *string) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetAdaptiveBFrameEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetAdaptiveBFrameEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetBFrames() {
	_jsii_.InvokeVoid(
		m,
		"resetBFrames",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetBufferWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetBufferWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetCrf() {
	_jsii_.InvokeVoid(
		m,
		"resetCrf",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetFrameRate() {
	_jsii_.InvokeVoid(
		m,
		"resetFrameRate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		m,
		"resetHeight",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetMaxBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		m,
		"resetProfile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetReferenceFrames() {
	_jsii_.InvokeVoid(
		m,
		"resetReferenceFrames",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetSlices() {
	_jsii_.InvokeVoid(
		m,
		"resetSlices",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		m,
		"resetWidth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH265VideoLayerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

