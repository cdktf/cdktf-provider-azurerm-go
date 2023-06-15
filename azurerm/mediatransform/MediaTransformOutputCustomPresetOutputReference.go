package mediatransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/mediatransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaTransformOutputCustomPresetOutputReference interface {
	cdktf.ComplexObject
	Codec() MediaTransformOutputCustomPresetCodecList
	CodecInput() interface{}
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
	Filter() MediaTransformOutputCustomPresetFilterOutputReference
	FilterInput() *MediaTransformOutputCustomPresetFilter
	Format() MediaTransformOutputCustomPresetFormatList
	FormatInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MediaTransformOutputCustomPreset
	SetInternalValue(val *MediaTransformOutputCustomPreset)
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
	PutCodec(value interface{})
	PutFilter(value *MediaTransformOutputCustomPresetFilter)
	PutFormat(value interface{})
	ResetExperimentalOptions()
	ResetFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaTransformOutputCustomPresetOutputReference
type jsiiProxy_MediaTransformOutputCustomPresetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) Codec() MediaTransformOutputCustomPresetCodecList {
	var returns MediaTransformOutputCustomPresetCodecList
	_jsii_.Get(
		j,
		"codec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) CodecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) ExperimentalOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"experimentalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) ExperimentalOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"experimentalOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) Filter() MediaTransformOutputCustomPresetFilterOutputReference {
	var returns MediaTransformOutputCustomPresetFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) FilterInput() *MediaTransformOutputCustomPresetFilter {
	var returns *MediaTransformOutputCustomPresetFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) Format() MediaTransformOutputCustomPresetFormatList {
	var returns MediaTransformOutputCustomPresetFormatList
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) FormatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) InternalValue() *MediaTransformOutputCustomPreset {
	var returns *MediaTransformOutputCustomPreset
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaTransformOutputCustomPresetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaTransformOutputCustomPresetOutputReference {
	_init_.Initialize()

	if err := validateNewMediaTransformOutputCustomPresetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaTransformOutputCustomPresetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaTransformOutputCustomPresetOutputReference_Override(m MediaTransformOutputCustomPresetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaTransform.MediaTransformOutputCustomPresetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference)SetExperimentalOptions(val *map[string]*string) {
	if err := j.validateSetExperimentalOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experimentalOptions",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference)SetInternalValue(val *MediaTransformOutputCustomPreset) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaTransformOutputCustomPresetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) PutCodec(value interface{}) {
	if err := m.validatePutCodecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCodec",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) PutFilter(value *MediaTransformOutputCustomPresetFilter) {
	if err := m.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilter",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) PutFormat(value interface{}) {
	if err := m.validatePutFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFormat",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) ResetExperimentalOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetExperimentalOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaTransformOutputCustomPresetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

