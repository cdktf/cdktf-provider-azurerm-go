package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputCustomPresetFilterOutputReference interface {
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
	CropRectangle() MediaTransformOutputCustomPresetFilterCropRectangleOutputReference
	CropRectangleInput() *MediaTransformOutputCustomPresetFilterCropRectangle
	Deinterlace() MediaTransformOutputCustomPresetFilterDeinterlaceOutputReference
	DeinterlaceInput() *MediaTransformOutputCustomPresetFilterDeinterlace
	FadeIn() MediaTransformOutputCustomPresetFilterFadeInOutputReference
	FadeInInput() *MediaTransformOutputCustomPresetFilterFadeIn
	FadeOut() MediaTransformOutputCustomPresetFilterFadeOutOutputReference
	FadeOutInput() *MediaTransformOutputCustomPresetFilterFadeOut
	// Experimental.
	Fqn() *string
	InternalValue() *MediaTransformOutputCustomPresetFilter
	SetInternalValue(val *MediaTransformOutputCustomPresetFilter)
	Overlay() MediaTransformOutputCustomPresetFilterOverlayList
	OverlayInput() interface{}
	Rotation() *string
	SetRotation(val *string)
	RotationInput() *string
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
	PutCropRectangle(value *MediaTransformOutputCustomPresetFilterCropRectangle)
	PutDeinterlace(value *MediaTransformOutputCustomPresetFilterDeinterlace)
	PutFadeIn(value *MediaTransformOutputCustomPresetFilterFadeIn)
	PutFadeOut(value *MediaTransformOutputCustomPresetFilterFadeOut)
	PutOverlay(value interface{})
	ResetCropRectangle()
	ResetDeinterlace()
	ResetFadeIn()
	ResetFadeOut()
	ResetOverlay()
	ResetRotation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputCustomPresetFilterOutputReference
type jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) CropRectangle() MediaTransformOutputCustomPresetFilterCropRectangleOutputReference {
	var returns MediaTransformOutputCustomPresetFilterCropRectangleOutputReference
	_jsii_.Get(
		j,
		"cropRectangle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) CropRectangleInput() *MediaTransformOutputCustomPresetFilterCropRectangle {
	var returns *MediaTransformOutputCustomPresetFilterCropRectangle
	_jsii_.Get(
		j,
		"cropRectangleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) Deinterlace() MediaTransformOutputCustomPresetFilterDeinterlaceOutputReference {
	var returns MediaTransformOutputCustomPresetFilterDeinterlaceOutputReference
	_jsii_.Get(
		j,
		"deinterlace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) DeinterlaceInput() *MediaTransformOutputCustomPresetFilterDeinterlace {
	var returns *MediaTransformOutputCustomPresetFilterDeinterlace
	_jsii_.Get(
		j,
		"deinterlaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) FadeIn() MediaTransformOutputCustomPresetFilterFadeInOutputReference {
	var returns MediaTransformOutputCustomPresetFilterFadeInOutputReference
	_jsii_.Get(
		j,
		"fadeIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) FadeInInput() *MediaTransformOutputCustomPresetFilterFadeIn {
	var returns *MediaTransformOutputCustomPresetFilterFadeIn
	_jsii_.Get(
		j,
		"fadeInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) FadeOut() MediaTransformOutputCustomPresetFilterFadeOutOutputReference {
	var returns MediaTransformOutputCustomPresetFilterFadeOutOutputReference
	_jsii_.Get(
		j,
		"fadeOut",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) FadeOutInput() *MediaTransformOutputCustomPresetFilterFadeOut {
	var returns *MediaTransformOutputCustomPresetFilterFadeOut
	_jsii_.Get(
		j,
		"fadeOutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) InternalValue() *MediaTransformOutputCustomPresetFilter {
	var returns *MediaTransformOutputCustomPresetFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) Overlay() MediaTransformOutputCustomPresetFilterOverlayList {
	var returns MediaTransformOutputCustomPresetFilterOverlayList
	_jsii_.Get(
		j,
		"overlay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) OverlayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overlayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) Rotation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) RotationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputCustomPresetFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaTransformOutputCustomPresetFilterOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputCustomPresetFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaTransformOutputCustomPresetFilterOutputReference_Override(m MediaTransformOutputCustomPresetFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference)SetInternalValue(val *MediaTransformOutputCustomPresetFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference)SetRotation(val *string) {
	if err := j.validateSetRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotation",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) PutCropRectangle(value *MediaTransformOutputCustomPresetFilterCropRectangle) {
	if err := m.validatePutCropRectangleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCropRectangle",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) PutDeinterlace(value *MediaTransformOutputCustomPresetFilterDeinterlace) {
	if err := m.validatePutDeinterlaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDeinterlace",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) PutFadeIn(value *MediaTransformOutputCustomPresetFilterFadeIn) {
	if err := m.validatePutFadeInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFadeIn",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) PutFadeOut(value *MediaTransformOutputCustomPresetFilterFadeOut) {
	if err := m.validatePutFadeOutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFadeOut",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) PutOverlay(value interface{}) {
	if err := m.validatePutOverlayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOverlay",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ResetCropRectangle() {
	_jsii_.InvokeVoid(
		m,
		"resetCropRectangle",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ResetDeinterlace() {
	_jsii_.InvokeVoid(
		m,
		"resetDeinterlace",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ResetFadeIn() {
	_jsii_.InvokeVoid(
		m,
		"resetFadeIn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ResetFadeOut() {
	_jsii_.InvokeVoid(
		m,
		"resetFadeOut",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ResetOverlay() {
	_jsii_.InvokeVoid(
		m,
		"resetOverlay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ResetRotation() {
	_jsii_.InvokeVoid(
		m,
		"resetRotation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

