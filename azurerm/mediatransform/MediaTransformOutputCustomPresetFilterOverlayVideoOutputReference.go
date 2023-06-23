package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference interface {
	cdktf.ComplexObject
	AudioGainLevel() *float64
	SetAudioGainLevel(val *float64)
	AudioGainLevelInput() *float64
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
	CropRectangle() MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangleOutputReference
	CropRectangleInput() *MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangle
	End() *string
	SetEnd(val *string)
	EndInput() *string
	FadeInDuration() *string
	SetFadeInDuration(val *string)
	FadeInDurationInput() *string
	FadeOutDuration() *string
	SetFadeOutDuration(val *string)
	FadeOutDurationInput() *string
	// Experimental.
	Fqn() *string
	InputLabel() *string
	SetInputLabel(val *string)
	InputLabelInput() *string
	InternalValue() *MediaTransformOutputCustomPresetFilterOverlayVideo
	SetInternalValue(val *MediaTransformOutputCustomPresetFilterOverlayVideo)
	Opacity() *float64
	SetOpacity(val *float64)
	OpacityInput() *float64
	Position() MediaTransformOutputCustomPresetFilterOverlayVideoPositionOutputReference
	PositionInput() *MediaTransformOutputCustomPresetFilterOverlayVideoPosition
	Start() *string
	SetStart(val *string)
	StartInput() *string
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
	PutCropRectangle(value *MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangle)
	PutPosition(value *MediaTransformOutputCustomPresetFilterOverlayVideoPosition)
	ResetAudioGainLevel()
	ResetCropRectangle()
	ResetEnd()
	ResetFadeInDuration()
	ResetFadeOutDuration()
	ResetOpacity()
	ResetPosition()
	ResetStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference
type jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) AudioGainLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"audioGainLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) AudioGainLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"audioGainLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) CropRectangle() MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangleOutputReference {
	var returns MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangleOutputReference
	_jsii_.Get(
		j,
		"cropRectangle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) CropRectangleInput() *MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangle {
	var returns *MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangle
	_jsii_.Get(
		j,
		"cropRectangleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) FadeInDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fadeInDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) FadeInDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fadeInDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) FadeOutDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fadeOutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) FadeOutDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fadeOutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) InputLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) InputLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) InternalValue() *MediaTransformOutputCustomPresetFilterOverlayVideo {
	var returns *MediaTransformOutputCustomPresetFilterOverlayVideo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) Opacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"opacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) OpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"opacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) Position() MediaTransformOutputCustomPresetFilterOverlayVideoPositionOutputReference {
	var returns MediaTransformOutputCustomPresetFilterOverlayVideoPositionOutputReference
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) PositionInput() *MediaTransformOutputCustomPresetFilterOverlayVideoPosition {
	var returns *MediaTransformOutputCustomPresetFilterOverlayVideoPosition
	_jsii_.Get(
		j,
		"positionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputCustomPresetFilterOverlayVideoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputCustomPresetFilterOverlayVideoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaTransformOutputCustomPresetFilterOverlayVideoOutputReference_Override(m MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetAudioGainLevel(val *float64) {
	if err := j.validateSetAudioGainLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioGainLevel",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetEnd(val *string) {
	if err := j.validateSetEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"end",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetFadeInDuration(val *string) {
	if err := j.validateSetFadeInDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fadeInDuration",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetFadeOutDuration(val *string) {
	if err := j.validateSetFadeOutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fadeOutDuration",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetInputLabel(val *string) {
	if err := j.validateSetInputLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLabel",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetInternalValue(val *MediaTransformOutputCustomPresetFilterOverlayVideo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetOpacity(val *float64) {
	if err := j.validateSetOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opacity",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetStart(val *string) {
	if err := j.validateSetStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"start",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) PutCropRectangle(value *MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangle) {
	if err := m.validatePutCropRectangleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCropRectangle",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) PutPosition(value *MediaTransformOutputCustomPresetFilterOverlayVideoPosition) {
	if err := m.validatePutPositionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPosition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ResetAudioGainLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetAudioGainLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ResetCropRectangle() {
	_jsii_.InvokeVoid(
		m,
		"resetCropRectangle",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ResetEnd() {
	_jsii_.InvokeVoid(
		m,
		"resetEnd",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ResetFadeInDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetFadeInDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ResetFadeOutDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetFadeOutDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ResetOpacity() {
	_jsii_.InvokeVoid(
		m,
		"resetOpacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ResetPosition() {
	_jsii_.InvokeVoid(
		m,
		"resetPosition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ResetStart() {
	_jsii_.InvokeVoid(
		m,
		"resetStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOverlayVideoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

