// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputCustomPresetCodecH264VideoOutputReference interface {
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
	InternalValue() *MediaTransformOutputCustomPresetCodecH264Video
	SetInternalValue(val *MediaTransformOutputCustomPresetCodecH264Video)
	KeyFrameInterval() *string
	SetKeyFrameInterval(val *string)
	KeyFrameIntervalInput() *string
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Layer() MediaTransformOutputCustomPresetCodecH264VideoLayerList
	LayerInput() interface{}
	RateControlMode() *string
	SetRateControlMode(val *string)
	RateControlModeInput() *string
	SceneChangeDetectionEnabled() interface{}
	SetSceneChangeDetectionEnabled(val interface{})
	SceneChangeDetectionEnabledInput() interface{}
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
	ResetComplexity()
	ResetKeyFrameInterval()
	ResetLabel()
	ResetLayer()
	ResetRateControlMode()
	ResetSceneChangeDetectionEnabled()
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

// The jsii proxy struct for MediaTransformOutputCustomPresetCodecH264VideoOutputReference
type jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) Complexity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ComplexityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) InternalValue() *MediaTransformOutputCustomPresetCodecH264Video {
	var returns *MediaTransformOutputCustomPresetCodecH264Video
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) KeyFrameInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFrameInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) KeyFrameIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFrameIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) Layer() MediaTransformOutputCustomPresetCodecH264VideoLayerList {
	var returns MediaTransformOutputCustomPresetCodecH264VideoLayerList
	_jsii_.Get(
		j,
		"layer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) LayerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"layerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) RateControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) RateControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) SceneChangeDetectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sceneChangeDetectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) SceneChangeDetectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sceneChangeDetectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) StretchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stretchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) StretchModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stretchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) SyncMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) SyncModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputCustomPresetCodecH264VideoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaTransformOutputCustomPresetCodecH264VideoOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputCustomPresetCodecH264VideoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetCodecH264VideoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaTransformOutputCustomPresetCodecH264VideoOutputReference_Override(m MediaTransformOutputCustomPresetCodecH264VideoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetCodecH264VideoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetComplexity(val *string) {
	if err := j.validateSetComplexityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexity",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetInternalValue(val *MediaTransformOutputCustomPresetCodecH264Video) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetKeyFrameInterval(val *string) {
	if err := j.validateSetKeyFrameIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFrameInterval",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetRateControlMode(val *string) {
	if err := j.validateSetRateControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateControlMode",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetSceneChangeDetectionEnabled(val interface{}) {
	if err := j.validateSetSceneChangeDetectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sceneChangeDetectionEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetStretchMode(val *string) {
	if err := j.validateSetStretchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stretchMode",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetSyncMode(val *string) {
	if err := j.validateSetSyncModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncMode",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) PutLayer(value interface{}) {
	if err := m.validatePutLayerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLayer",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ResetComplexity() {
	_jsii_.InvokeVoid(
		m,
		"resetComplexity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ResetKeyFrameInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyFrameInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ResetLayer() {
	_jsii_.InvokeVoid(
		m,
		"resetLayer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ResetRateControlMode() {
	_jsii_.InvokeVoid(
		m,
		"resetRateControlMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ResetSceneChangeDetectionEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetSceneChangeDetectionEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ResetStretchMode() {
	_jsii_.InvokeVoid(
		m,
		"resetStretchMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ResetSyncMode() {
	_jsii_.InvokeVoid(
		m,
		"resetSyncMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetCodecH264VideoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

