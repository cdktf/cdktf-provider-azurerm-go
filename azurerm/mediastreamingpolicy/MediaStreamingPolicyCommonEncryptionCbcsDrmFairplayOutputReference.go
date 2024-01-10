// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/mediastreamingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference interface {
	cdktf.ComplexObject
	AllowPersistentLicense() interface{}
	SetAllowPersistentLicense(val interface{})
	AllowPersistentLicenseInput() interface{}
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
	CustomLicenseAcquisitionUrlTemplate() *string
	SetCustomLicenseAcquisitionUrlTemplate(val *string)
	CustomLicenseAcquisitionUrlTemplateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay)
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
	ResetAllowPersistentLicense()
	ResetCustomLicenseAcquisitionUrlTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) AllowPersistentLicense() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPersistentLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) AllowPersistentLicenseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPersistentLicenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) CustomLicenseAcquisitionUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLicenseAcquisitionUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) CustomLicenseAcquisitionUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLicenseAcquisitionUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay {
	var returns *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference {
	_init_.Initialize()

	if err := validateNewMediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference)SetAllowPersistentLicense(val interface{}) {
	if err := j.validateSetAllowPersistentLicenseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPersistentLicense",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference)SetCustomLicenseAcquisitionUrlTemplate(val *string) {
	if err := j.validateSetCustomLicenseAcquisitionUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLicenseAcquisitionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference)SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ResetAllowPersistentLicense() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowPersistentLicense",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ResetCustomLicenseAcquisitionUrlTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomLicenseAcquisitionUrlTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

