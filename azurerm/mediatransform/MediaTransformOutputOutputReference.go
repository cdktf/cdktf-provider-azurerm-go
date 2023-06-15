package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputOutputReference interface {
	cdktf.ComplexObject
	AudioAnalyzerPreset() MediaTransformOutputAudioAnalyzerPresetOutputReference
	AudioAnalyzerPresetInput() *MediaTransformOutputAudioAnalyzerPreset
	BuiltinPreset() MediaTransformOutputBuiltinPresetOutputReference
	BuiltinPresetInput() *MediaTransformOutputBuiltinPreset
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
	CustomPreset() MediaTransformOutputCustomPresetOutputReference
	CustomPresetInput() *MediaTransformOutputCustomPreset
	FaceDetectorPreset() MediaTransformOutputFaceDetectorPresetOutputReference
	FaceDetectorPresetInput() *MediaTransformOutputFaceDetectorPreset
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OnErrorAction() *string
	SetOnErrorAction(val *string)
	OnErrorActionInput() *string
	RelativePriority() *string
	SetRelativePriority(val *string)
	RelativePriorityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VideoAnalyzerPreset() MediaTransformOutputVideoAnalyzerPresetOutputReference
	VideoAnalyzerPresetInput() *MediaTransformOutputVideoAnalyzerPreset
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
	PutAudioAnalyzerPreset(value *MediaTransformOutputAudioAnalyzerPreset)
	PutBuiltinPreset(value *MediaTransformOutputBuiltinPreset)
	PutCustomPreset(value *MediaTransformOutputCustomPreset)
	PutFaceDetectorPreset(value *MediaTransformOutputFaceDetectorPreset)
	PutVideoAnalyzerPreset(value *MediaTransformOutputVideoAnalyzerPreset)
	ResetAudioAnalyzerPreset()
	ResetBuiltinPreset()
	ResetCustomPreset()
	ResetFaceDetectorPreset()
	ResetOnErrorAction()
	ResetRelativePriority()
	ResetVideoAnalyzerPreset()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputOutputReference
type jsiiProxy_MediaTransformOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) AudioAnalyzerPreset() MediaTransformOutputAudioAnalyzerPresetOutputReference {
	var returns MediaTransformOutputAudioAnalyzerPresetOutputReference
	_jsii_.Get(
		j,
		"audioAnalyzerPreset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) AudioAnalyzerPresetInput() *MediaTransformOutputAudioAnalyzerPreset {
	var returns *MediaTransformOutputAudioAnalyzerPreset
	_jsii_.Get(
		j,
		"audioAnalyzerPresetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) BuiltinPreset() MediaTransformOutputBuiltinPresetOutputReference {
	var returns MediaTransformOutputBuiltinPresetOutputReference
	_jsii_.Get(
		j,
		"builtinPreset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) BuiltinPresetInput() *MediaTransformOutputBuiltinPreset {
	var returns *MediaTransformOutputBuiltinPreset
	_jsii_.Get(
		j,
		"builtinPresetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) CustomPreset() MediaTransformOutputCustomPresetOutputReference {
	var returns MediaTransformOutputCustomPresetOutputReference
	_jsii_.Get(
		j,
		"customPreset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) CustomPresetInput() *MediaTransformOutputCustomPreset {
	var returns *MediaTransformOutputCustomPreset
	_jsii_.Get(
		j,
		"customPresetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) FaceDetectorPreset() MediaTransformOutputFaceDetectorPresetOutputReference {
	var returns MediaTransformOutputFaceDetectorPresetOutputReference
	_jsii_.Get(
		j,
		"faceDetectorPreset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) FaceDetectorPresetInput() *MediaTransformOutputFaceDetectorPreset {
	var returns *MediaTransformOutputFaceDetectorPreset
	_jsii_.Get(
		j,
		"faceDetectorPresetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) OnErrorAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onErrorAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) OnErrorActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onErrorActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) RelativePriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) RelativePriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) VideoAnalyzerPreset() MediaTransformOutputVideoAnalyzerPresetOutputReference {
	var returns MediaTransformOutputVideoAnalyzerPresetOutputReference
	_jsii_.Get(
		j,
		"videoAnalyzerPreset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputOutputReference) VideoAnalyzerPresetInput() *MediaTransformOutputVideoAnalyzerPreset {
	var returns *MediaTransformOutputVideoAnalyzerPreset
	_jsii_.Get(
		j,
		"videoAnalyzerPresetInput",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaTransformOutputOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaTransformOutputOutputReference_Override(m MediaTransformOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputOutputReference)SetOnErrorAction(val *string) {
	if err := j.validateSetOnErrorActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onErrorAction",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputOutputReference)SetRelativePriority(val *string) {
	if err := j.validateSetRelativePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relativePriority",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) PutAudioAnalyzerPreset(value *MediaTransformOutputAudioAnalyzerPreset) {
	if err := m.validatePutAudioAnalyzerPresetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAudioAnalyzerPreset",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) PutBuiltinPreset(value *MediaTransformOutputBuiltinPreset) {
	if err := m.validatePutBuiltinPresetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBuiltinPreset",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) PutCustomPreset(value *MediaTransformOutputCustomPreset) {
	if err := m.validatePutCustomPresetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCustomPreset",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) PutFaceDetectorPreset(value *MediaTransformOutputFaceDetectorPreset) {
	if err := m.validatePutFaceDetectorPresetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFaceDetectorPreset",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) PutVideoAnalyzerPreset(value *MediaTransformOutputVideoAnalyzerPreset) {
	if err := m.validatePutVideoAnalyzerPresetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVideoAnalyzerPreset",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) ResetAudioAnalyzerPreset() {
	_jsii_.InvokeVoid(
		m,
		"resetAudioAnalyzerPreset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) ResetBuiltinPreset() {
	_jsii_.InvokeVoid(
		m,
		"resetBuiltinPreset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) ResetCustomPreset() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomPreset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) ResetFaceDetectorPreset() {
	_jsii_.InvokeVoid(
		m,
		"resetFaceDetectorPreset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) ResetOnErrorAction() {
	_jsii_.InvokeVoid(
		m,
		"resetOnErrorAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) ResetRelativePriority() {
	_jsii_.InvokeVoid(
		m,
		"resetRelativePriority",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) ResetVideoAnalyzerPreset() {
	_jsii_.InvokeVoid(
		m,
		"resetVideoAnalyzerPreset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

