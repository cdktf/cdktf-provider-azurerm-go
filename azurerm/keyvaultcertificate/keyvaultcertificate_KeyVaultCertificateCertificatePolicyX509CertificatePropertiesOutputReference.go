package keyvaultcertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/keyvaultcertificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference interface {
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
	ExtendedKeyUsage() *[]*string
	SetExtendedKeyUsage(val *[]*string)
	ExtendedKeyUsageInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KeyVaultCertificateCertificatePolicyX509CertificateProperties
	SetInternalValue(val *KeyVaultCertificateCertificatePolicyX509CertificateProperties)
	KeyUsage() *[]*string
	SetKeyUsage(val *[]*string)
	KeyUsageInput() *[]*string
	Subject() *string
	SetSubject(val *string)
	SubjectAlternativeNames() KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesOutputReference
	SubjectAlternativeNamesInput() *KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames
	SubjectInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidityInMonths() *float64
	SetValidityInMonths(val *float64)
	ValidityInMonthsInput() *float64
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
	PutSubjectAlternativeNames(value *KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)
	ResetExtendedKeyUsage()
	ResetSubjectAlternativeNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference
type jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ExtendedKeyUsage() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendedKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ExtendedKeyUsageInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendedKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) InternalValue() *KeyVaultCertificateCertificatePolicyX509CertificateProperties {
	var returns *KeyVaultCertificateCertificatePolicyX509CertificateProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) KeyUsage() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) KeyUsageInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) SubjectAlternativeNames() KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesOutputReference {
	var returns KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesOutputReference
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) SubjectAlternativeNamesInput() *KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames {
	var returns *KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames
	_jsii_.Get(
		j,
		"subjectAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ValidityInMonths() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validityInMonths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ValidityInMonthsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validityInMonthsInput",
		&returns,
	)
	return returns
}


func NewKeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewKeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultCertificate.KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference_Override(k KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultCertificate.KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetExtendedKeyUsage(val *[]*string) {
	if err := j.validateSetExtendedKeyUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedKeyUsage",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetInternalValue(val *KeyVaultCertificateCertificatePolicyX509CertificateProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetKeyUsage(val *[]*string) {
	if err := j.validateSetKeyUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyUsage",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference)SetValidityInMonths(val *float64) {
	if err := j.validateSetValidityInMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validityInMonths",
		val,
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) PutSubjectAlternativeNames(value *KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames) {
	if err := k.validatePutSubjectAlternativeNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSubjectAlternativeNames",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ResetExtendedKeyUsage() {
	_jsii_.InvokeVoid(
		k,
		"resetExtendedKeyUsage",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ResetSubjectAlternativeNames() {
	_jsii_.InvokeVoid(
		k,
		"resetSubjectAlternativeNames",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

