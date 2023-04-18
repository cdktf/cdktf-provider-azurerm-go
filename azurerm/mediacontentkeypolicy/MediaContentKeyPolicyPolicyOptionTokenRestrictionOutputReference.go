package mediacontentkeypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/mediacontentkeypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference interface {
	cdktf.ComplexObject
	AlternateKey() MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyList
	AlternateKeyInput() interface{}
	Audience() *string
	SetAudience(val *string)
	AudienceInput() *string
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
	InternalValue() *MediaContentKeyPolicyPolicyOptionTokenRestriction
	SetInternalValue(val *MediaContentKeyPolicyPolicyOptionTokenRestriction)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	OpenIdConnectDiscoveryDocument() *string
	SetOpenIdConnectDiscoveryDocument(val *string)
	OpenIdConnectDiscoveryDocumentInput() *string
	PrimaryRsaTokenKeyExponent() *string
	SetPrimaryRsaTokenKeyExponent(val *string)
	PrimaryRsaTokenKeyExponentInput() *string
	PrimaryRsaTokenKeyModulus() *string
	SetPrimaryRsaTokenKeyModulus(val *string)
	PrimaryRsaTokenKeyModulusInput() *string
	PrimarySymmetricTokenKey() *string
	SetPrimarySymmetricTokenKey(val *string)
	PrimarySymmetricTokenKeyInput() *string
	PrimaryX509TokenKeyRaw() *string
	SetPrimaryX509TokenKeyRaw(val *string)
	PrimaryX509TokenKeyRawInput() *string
	RequiredClaim() MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList
	RequiredClaimInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
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
	PutAlternateKey(value interface{})
	PutRequiredClaim(value interface{})
	ResetAlternateKey()
	ResetAudience()
	ResetIssuer()
	ResetOpenIdConnectDiscoveryDocument()
	ResetPrimaryRsaTokenKeyExponent()
	ResetPrimaryRsaTokenKeyModulus()
	ResetPrimarySymmetricTokenKey()
	ResetPrimaryX509TokenKeyRaw()
	ResetRequiredClaim()
	ResetTokenType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) AlternateKey() MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyList {
	var returns MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKeyList
	_jsii_.Get(
		j,
		"alternateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) AlternateKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alternateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) AudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) InternalValue() *MediaContentKeyPolicyPolicyOptionTokenRestriction {
	var returns *MediaContentKeyPolicyPolicyOptionTokenRestriction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) OpenIdConnectDiscoveryDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectDiscoveryDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) OpenIdConnectDiscoveryDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectDiscoveryDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryRsaTokenKeyExponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryRsaTokenKeyExponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryRsaTokenKeyExponentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryRsaTokenKeyExponentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryRsaTokenKeyModulus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryRsaTokenKeyModulus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryRsaTokenKeyModulusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryRsaTokenKeyModulusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimarySymmetricTokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primarySymmetricTokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimarySymmetricTokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primarySymmetricTokenKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryX509TokenKeyRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryX509TokenKeyRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryX509TokenKeyRawInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryX509TokenKeyRawInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) RequiredClaim() MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList {
	var returns MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList
	_jsii_.Get(
		j,
		"requiredClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) RequiredClaimInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference {
	_init_.Initialize()

	if err := validateNewMediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference_Override(m MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetAudience(val *string) {
	if err := j.validateSetAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetInternalValue(val *MediaContentKeyPolicyPolicyOptionTokenRestriction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetOpenIdConnectDiscoveryDocument(val *string) {
	if err := j.validateSetOpenIdConnectDiscoveryDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openIdConnectDiscoveryDocument",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetPrimaryRsaTokenKeyExponent(val *string) {
	if err := j.validateSetPrimaryRsaTokenKeyExponentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryRsaTokenKeyExponent",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetPrimaryRsaTokenKeyModulus(val *string) {
	if err := j.validateSetPrimaryRsaTokenKeyModulusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryRsaTokenKeyModulus",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetPrimarySymmetricTokenKey(val *string) {
	if err := j.validateSetPrimarySymmetricTokenKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primarySymmetricTokenKey",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetPrimaryX509TokenKeyRaw(val *string) {
	if err := j.validateSetPrimaryX509TokenKeyRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryX509TokenKeyRaw",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PutAlternateKey(value interface{}) {
	if err := m.validatePutAlternateKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlternateKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PutRequiredClaim(value interface{}) {
	if err := m.validatePutRequiredClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRequiredClaim",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetAlternateKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAlternateKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetAudience() {
	_jsii_.InvokeVoid(
		m,
		"resetAudience",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		m,
		"resetIssuer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetOpenIdConnectDiscoveryDocument() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenIdConnectDiscoveryDocument",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetPrimaryRsaTokenKeyExponent() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryRsaTokenKeyExponent",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetPrimaryRsaTokenKeyModulus() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryRsaTokenKeyModulus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetPrimarySymmetricTokenKey() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimarySymmetricTokenKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetPrimaryX509TokenKeyRaw() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryX509TokenKeyRaw",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetRequiredClaim() {
	_jsii_.InvokeVoid(
		m,
		"resetRequiredClaim",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetTokenType() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

