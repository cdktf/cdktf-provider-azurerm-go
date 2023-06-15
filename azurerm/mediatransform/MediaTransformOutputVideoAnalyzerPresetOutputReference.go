package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputVideoAnalyzerPresetOutputReference interface {
	cdktf.ComplexObject
	AudioAnalysisMode() *string
	SetAudioAnalysisMode(val *string)
	AudioAnalysisModeInput() *string
	AudioLanguage() *string
	SetAudioLanguage(val *string)
	AudioLanguageInput() *string
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
	// Experimental.
	Fqn() *string
	InsightsType() *string
	SetInsightsType(val *string)
	InsightsTypeInput() *string
	InternalValue() *MediaTransformOutputVideoAnalyzerPreset
	SetInternalValue(val *MediaTransformOutputVideoAnalyzerPreset)
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
	ResetAudioAnalysisMode()
	ResetAudioLanguage()
	ResetExperimentalOptions()
	ResetInsightsType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputVideoAnalyzerPresetOutputReference
type jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) AudioAnalysisMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioAnalysisMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) AudioAnalysisModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioAnalysisModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) AudioLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) AudioLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ExperimentalOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"experimentalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ExperimentalOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"experimentalOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) InsightsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) InsightsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) InternalValue() *MediaTransformOutputVideoAnalyzerPreset {
	var returns *MediaTransformOutputVideoAnalyzerPreset
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputVideoAnalyzerPresetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaTransformOutputVideoAnalyzerPresetOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputVideoAnalyzerPresetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputVideoAnalyzerPresetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaTransformOutputVideoAnalyzerPresetOutputReference_Override(m MediaTransformOutputVideoAnalyzerPresetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputVideoAnalyzerPresetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetAudioAnalysisMode(val *string) {
	if err := j.validateSetAudioAnalysisModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioAnalysisMode",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetAudioLanguage(val *string) {
	if err := j.validateSetAudioLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioLanguage",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetExperimentalOptions(val *map[string]*string) {
	if err := j.validateSetExperimentalOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experimentalOptions",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetInsightsType(val *string) {
	if err := j.validateSetInsightsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insightsType",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetInternalValue(val *MediaTransformOutputVideoAnalyzerPreset) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ResetAudioAnalysisMode() {
	_jsii_.InvokeVoid(
		m,
		"resetAudioAnalysisMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ResetAudioLanguage() {
	_jsii_.InvokeVoid(
		m,
		"resetAudioLanguage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ResetExperimentalOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetExperimentalOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ResetInsightsType() {
	_jsii_.InvokeVoid(
		m,
		"resetInsightsType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputVideoAnalyzerPresetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

