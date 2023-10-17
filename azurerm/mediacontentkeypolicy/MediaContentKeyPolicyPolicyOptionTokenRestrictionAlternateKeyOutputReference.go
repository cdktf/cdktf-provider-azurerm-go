// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediacontentkeypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mediacontentkeypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RsaTokenKeyExponent() *string
	SetRsaTokenKeyExponent(val *string)
	RsaTokenKeyExponentInput() *string
	RsaTokenKeyModulus() *string
	SetRsaTokenKeyModulus(val *string)
	RsaTokenKeyModulusInput() *string
	SymmetricTokenKey() *string
	SetSymmetricTokenKey(val *string)
	SymmetricTokenKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	X509TokenKeyRaw() *string
	SetX509TokenKeyRaw(val *string)
	X509TokenKeyRawInput() *string
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
	ResetRsaTokenKeyExponent()
	ResetRsaTokenKeyModulus()
	ResetSymmetricTokenKey()
	ResetX509TokenKeyRaw()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) RsaTokenKeyExponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaTokenKeyExponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) RsaTokenKeyExponentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaTokenKeyExponentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) RsaTokenKeyModulus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaTokenKeyModulus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) RsaTokenKeyModulusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaTokenKeyModulusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) SymmetricTokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"symmetricTokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) SymmetricTokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"symmetricTokenKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) X509TokenKeyRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509TokenKeyRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) X509TokenKeyRawInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509TokenKeyRawInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference {
	_init_.Initialize()

	if err := validateNewMediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference_Override(m MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetRsaTokenKeyExponent(val *string) {
	if err := j.validateSetRsaTokenKeyExponentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaTokenKeyExponent",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetRsaTokenKeyModulus(val *string) {
	if err := j.validateSetRsaTokenKeyModulusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaTokenKeyModulus",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetSymmetricTokenKey(val *string) {
	if err := j.validateSetSymmetricTokenKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"symmetricTokenKey",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference)SetX509TokenKeyRaw(val *string) {
	if err := j.validateSetX509TokenKeyRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"x509TokenKeyRaw",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) ResetRsaTokenKeyExponent() {
	_jsii_.InvokeVoid(
		m,
		"resetRsaTokenKeyExponent",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) ResetRsaTokenKeyModulus() {
	_jsii_.InvokeVoid(
		m,
		"resetRsaTokenKeyModulus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) ResetSymmetricTokenKey() {
	_jsii_.InvokeVoid(
		m,
		"resetSymmetricTokenKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) ResetX509TokenKeyRaw() {
	_jsii_.InvokeVoid(
		m,
		"resetX509TokenKeyRaw",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

