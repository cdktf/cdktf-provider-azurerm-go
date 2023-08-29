// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputCustomPresetCodecJpgImageOutputReference interface {
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
	InternalValue() *MediaTransformOutputCustomPresetCodecJpgImage
	SetInternalValue(val *MediaTransformOutputCustomPresetCodecJpgImage)
	KeyFrameInterval() *string
	SetKeyFrameInterval(val *string)
	KeyFrameIntervalInput() *string
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Layer() MediaTransformOutputCustomPresetCodecJpgImageLayerList
	LayerInput() interface{}
	Range() *string
	SetRange(val *string)
	RangeInput() *string
	SpriteColumn() *float64
	SetSpriteColumn(val *float64)
	SpriteColumnInput() *float64
	Start() *string
	SetStart(val *string)
	StartInput() *string
	Step() *string
	SetStep(val *string)
	StepInput() *string
	StretchMode() *string
	SetStretchMode(val *string)
	StretchModeInput() *string
	SyncMode() *string
	SetSyncMode(val *string)
	SyncModeInput() *string
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
	PutLayer(value interface{})
	ResetKeyFrameInterval()
	ResetLabel()
	ResetLayer()
	ResetRange()
	ResetSpriteColumn()
	ResetStep()
	ResetStretchMode()
	ResetSyncMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputCustomPresetCodecJpgImageOutputReference
type jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) InternalValue() *MediaTransformOutputCustomPresetCodecJpgImage {
	var returns *MediaTransformOutputCustomPresetCodecJpgImage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) KeyFrameInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFrameInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) KeyFrameIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFrameIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) Layer() MediaTransformOutputCustomPresetCodecJpgImageLayerList {
	var returns MediaTransformOutputCustomPresetCodecJpgImageLayerList
	_jsii_.Get(
		j,
		"layer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) LayerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"layerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) SpriteColumn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spriteColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) SpriteColumnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spriteColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) Step() *string {
	var returns *string
	_jsii_.Get(
		j,
		"step",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) StepInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) StretchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stretchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) StretchModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stretchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) SyncMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) SyncModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputCustomPresetCodecJpgImageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaTransformOutputCustomPresetCodecJpgImageOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputCustomPresetCodecJpgImageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetCodecJpgImageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaTransformOutputCustomPresetCodecJpgImageOutputReference_Override(m MediaTransformOutputCustomPresetCodecJpgImageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetCodecJpgImageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetInternalValue(val *MediaTransformOutputCustomPresetCodecJpgImage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetKeyFrameInterval(val *string) {
	if err := j.validateSetKeyFrameIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFrameInterval",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetRange(val *string) {
	if err := j.validateSetRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"range",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetSpriteColumn(val *float64) {
	if err := j.validateSetSpriteColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spriteColumn",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetStart(val *string) {
	if err := j.validateSetStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"start",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetStep(val *string) {
	if err := j.validateSetStepParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"step",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetStretchMode(val *string) {
	if err := j.validateSetStretchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stretchMode",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetSyncMode(val *string) {
	if err := j.validateSetSyncModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncMode",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) PutLayer(value interface{}) {
	if err := m.validatePutLayerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLayer",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ResetKeyFrameInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyFrameInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ResetLayer() {
	_jsii_.InvokeVoid(
		m,
		"resetLayer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ResetRange() {
	_jsii_.InvokeVoid(
		m,
		"resetRange",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ResetSpriteColumn() {
	_jsii_.InvokeVoid(
		m,
		"resetSpriteColumn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ResetStep() {
	_jsii_.InvokeVoid(
		m,
		"resetStep",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ResetStretchMode() {
	_jsii_.InvokeVoid(
		m,
		"resetStretchMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ResetSyncMode() {
	_jsii_.InvokeVoid(
		m,
		"resetSyncMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecJpgImageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

