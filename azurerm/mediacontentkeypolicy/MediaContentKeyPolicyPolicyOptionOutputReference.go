package mediacontentkeypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/mediacontentkeypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaContentKeyPolicyPolicyOptionOutputReference interface {
	cdktf.ComplexObject
	ClearKeyConfigurationEnabled() interface{}
	SetClearKeyConfigurationEnabled(val interface{})
	ClearKeyConfigurationEnabledInput() interface{}
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
	FairplayConfiguration() MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference
	FairplayConfigurationInput() *MediaContentKeyPolicyPolicyOptionFairplayConfiguration
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OpenRestrictionEnabled() interface{}
	SetOpenRestrictionEnabled(val interface{})
	OpenRestrictionEnabledInput() interface{}
	PlayreadyConfigurationLicense() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList
	PlayreadyConfigurationLicenseInput() interface{}
	PlayreadyResponseCustomData() *string
	SetPlayreadyResponseCustomData(val *string)
	PlayreadyResponseCustomDataInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenRestriction() MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference
	TokenRestrictionInput() *MediaContentKeyPolicyPolicyOptionTokenRestriction
	WidevineConfigurationTemplate() *string
	SetWidevineConfigurationTemplate(val *string)
	WidevineConfigurationTemplateInput() *string
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
	PutFairplayConfiguration(value *MediaContentKeyPolicyPolicyOptionFairplayConfiguration)
	PutPlayreadyConfigurationLicense(value interface{})
	PutTokenRestriction(value *MediaContentKeyPolicyPolicyOptionTokenRestriction)
	ResetClearKeyConfigurationEnabled()
	ResetFairplayConfiguration()
	ResetOpenRestrictionEnabled()
	ResetPlayreadyConfigurationLicense()
	ResetPlayreadyResponseCustomData()
	ResetTokenRestriction()
	ResetWidevineConfigurationTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ClearKeyConfigurationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearKeyConfigurationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ClearKeyConfigurationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearKeyConfigurationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) FairplayConfiguration() MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference
	_jsii_.Get(
		j,
		"fairplayConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) FairplayConfigurationInput() *MediaContentKeyPolicyPolicyOptionFairplayConfiguration {
	var returns *MediaContentKeyPolicyPolicyOptionFairplayConfiguration
	_jsii_.Get(
		j,
		"fairplayConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) OpenRestrictionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openRestrictionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) OpenRestrictionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openRestrictionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PlayreadyConfigurationLicense() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList {
	var returns MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList
	_jsii_.Get(
		j,
		"playreadyConfigurationLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PlayreadyConfigurationLicenseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"playreadyConfigurationLicenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PlayreadyResponseCustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playreadyResponseCustomData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PlayreadyResponseCustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playreadyResponseCustomDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) TokenRestriction() MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference
	_jsii_.Get(
		j,
		"tokenRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) TokenRestrictionInput() *MediaContentKeyPolicyPolicyOptionTokenRestriction {
	var returns *MediaContentKeyPolicyPolicyOptionTokenRestriction
	_jsii_.Get(
		j,
		"tokenRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) WidevineConfigurationTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widevineConfigurationTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) WidevineConfigurationTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widevineConfigurationTemplateInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaContentKeyPolicyPolicyOptionOutputReference {
	_init_.Initialize()

	if err := validateNewMediaContentKeyPolicyPolicyOptionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionOutputReference_Override(m MediaContentKeyPolicyPolicyOptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetClearKeyConfigurationEnabled(val interface{}) {
	if err := j.validateSetClearKeyConfigurationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clearKeyConfigurationEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetOpenRestrictionEnabled(val interface{}) {
	if err := j.validateSetOpenRestrictionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openRestrictionEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetPlayreadyResponseCustomData(val *string) {
	if err := j.validateSetPlayreadyResponseCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"playreadyResponseCustomData",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference)SetWidevineConfigurationTemplate(val *string) {
	if err := j.validateSetWidevineConfigurationTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widevineConfigurationTemplate",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PutFairplayConfiguration(value *MediaContentKeyPolicyPolicyOptionFairplayConfiguration) {
	if err := m.validatePutFairplayConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFairplayConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PutPlayreadyConfigurationLicense(value interface{}) {
	if err := m.validatePutPlayreadyConfigurationLicenseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPlayreadyConfigurationLicense",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PutTokenRestriction(value *MediaContentKeyPolicyPolicyOptionTokenRestriction) {
	if err := m.validatePutTokenRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTokenRestriction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetClearKeyConfigurationEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetClearKeyConfigurationEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetFairplayConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetFairplayConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetOpenRestrictionEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenRestrictionEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetPlayreadyConfigurationLicense() {
	_jsii_.InvokeVoid(
		m,
		"resetPlayreadyConfigurationLicense",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetPlayreadyResponseCustomData() {
	_jsii_.InvokeVoid(
		m,
		"resetPlayreadyResponseCustomData",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetTokenRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetWidevineConfigurationTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetWidevineConfigurationTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

