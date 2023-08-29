// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputFaceDetectorPresetOutputReference interface {
	cdktf.ComplexObject
	AnalysisResolution() *string
	SetAnalysisResolution(val *string)
	AnalysisResolutionInput() *string
	BlurType() *string
	SetBlurType(val *string)
	BlurTypeInput() *string
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
	ExperimentalOptions() *map[string]*string
	SetExperimentalOptions(val *map[string]*string)
	ExperimentalOptionsInput() *map[string]*string
	FaceRedactorMode() *string
	SetFaceRedactorMode(val *string)
	FaceRedactorModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MediaTransformOutputFaceDetectorPreset
	SetInternalValue(val *MediaTransformOutputFaceDetectorPreset)
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
	ResetAnalysisResolution()
	ResetBlurType()
	ResetExperimentalOptions()
	ResetFaceRedactorMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputFaceDetectorPresetOutputReference
type jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) AnalysisResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) AnalysisResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) BlurType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blurType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) BlurTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blurTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ExperimentalOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"experimentalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ExperimentalOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"experimentalOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) FaceRedactorMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faceRedactorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) FaceRedactorModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faceRedactorModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) InternalValue() *MediaTransformOutputFaceDetectorPreset {
	var returns *MediaTransformOutputFaceDetectorPreset
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputFaceDetectorPresetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaTransformOutputFaceDetectorPresetOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputFaceDetectorPresetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputFaceDetectorPresetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaTransformOutputFaceDetectorPresetOutputReference_Override(m MediaTransformOutputFaceDetectorPresetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputFaceDetectorPresetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetAnalysisResolution(val *string) {
	if err := j.validateSetAnalysisResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analysisResolution",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetBlurType(val *string) {
	if err := j.validateSetBlurTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blurType",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetExperimentalOptions(val *map[string]*string) {
	if err := j.validateSetExperimentalOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experimentalOptions",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetFaceRedactorMode(val *string) {
	if err := j.validateSetFaceRedactorModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faceRedactorMode",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetInternalValue(val *MediaTransformOutputFaceDetectorPreset) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ResetAnalysisResolution() {
	_jsii_.InvokeVoid(
		m,
		"resetAnalysisResolution",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ResetBlurType() {
	_jsii_.InvokeVoid(
		m,
		"resetBlurType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ResetExperimentalOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetExperimentalOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ResetFaceRedactorMode() {
	_jsii_.InvokeVoid(
		m,
		"resetFaceRedactorMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputFaceDetectorPresetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

