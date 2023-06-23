package keyvaultcertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/keyvaultcertificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultCertificateCertificatePolicyOutputReference interface {
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
	InternalValue() *KeyVaultCertificateCertificatePolicy
	SetInternalValue(val *KeyVaultCertificateCertificatePolicy)
	IssuerParameters() KeyVaultCertificateCertificatePolicyIssuerParametersOutputReference
	IssuerParametersInput() *KeyVaultCertificateCertificatePolicyIssuerParameters
	KeyProperties() KeyVaultCertificateCertificatePolicyKeyPropertiesOutputReference
	KeyPropertiesInput() *KeyVaultCertificateCertificatePolicyKeyProperties
	LifetimeAction() KeyVaultCertificateCertificatePolicyLifetimeActionList
	LifetimeActionInput() interface{}
	SecretProperties() KeyVaultCertificateCertificatePolicySecretPropertiesOutputReference
	SecretPropertiesInput() *KeyVaultCertificateCertificatePolicySecretProperties
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	X509CertificateProperties() KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference
	X509CertificatePropertiesInput() *KeyVaultCertificateCertificatePolicyX509CertificateProperties
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
	PutIssuerParameters(value *KeyVaultCertificateCertificatePolicyIssuerParameters)
	PutKeyProperties(value *KeyVaultCertificateCertificatePolicyKeyProperties)
	PutLifetimeAction(value interface{})
	PutSecretProperties(value *KeyVaultCertificateCertificatePolicySecretProperties)
	PutX509CertificateProperties(value *KeyVaultCertificateCertificatePolicyX509CertificateProperties)
	ResetLifetimeAction()
	ResetX509CertificateProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyVaultCertificateCertificatePolicyOutputReference
type jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) InternalValue() *KeyVaultCertificateCertificatePolicy {
	var returns *KeyVaultCertificateCertificatePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) IssuerParameters() KeyVaultCertificateCertificatePolicyIssuerParametersOutputReference {
	var returns KeyVaultCertificateCertificatePolicyIssuerParametersOutputReference
	_jsii_.Get(
		j,
		"issuerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) IssuerParametersInput() *KeyVaultCertificateCertificatePolicyIssuerParameters {
	var returns *KeyVaultCertificateCertificatePolicyIssuerParameters
	_jsii_.Get(
		j,
		"issuerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) KeyProperties() KeyVaultCertificateCertificatePolicyKeyPropertiesOutputReference {
	var returns KeyVaultCertificateCertificatePolicyKeyPropertiesOutputReference
	_jsii_.Get(
		j,
		"keyProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) KeyPropertiesInput() *KeyVaultCertificateCertificatePolicyKeyProperties {
	var returns *KeyVaultCertificateCertificatePolicyKeyProperties
	_jsii_.Get(
		j,
		"keyPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) LifetimeAction() KeyVaultCertificateCertificatePolicyLifetimeActionList {
	var returns KeyVaultCertificateCertificatePolicyLifetimeActionList
	_jsii_.Get(
		j,
		"lifetimeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) LifetimeActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifetimeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) SecretProperties() KeyVaultCertificateCertificatePolicySecretPropertiesOutputReference {
	var returns KeyVaultCertificateCertificatePolicySecretPropertiesOutputReference
	_jsii_.Get(
		j,
		"secretProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) SecretPropertiesInput() *KeyVaultCertificateCertificatePolicySecretProperties {
	var returns *KeyVaultCertificateCertificatePolicySecretProperties
	_jsii_.Get(
		j,
		"secretPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) X509CertificateProperties() KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference {
	var returns KeyVaultCertificateCertificatePolicyX509CertificatePropertiesOutputReference
	_jsii_.Get(
		j,
		"x509CertificateProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) X509CertificatePropertiesInput() *KeyVaultCertificateCertificatePolicyX509CertificateProperties {
	var returns *KeyVaultCertificateCertificatePolicyX509CertificateProperties
	_jsii_.Get(
		j,
		"x509CertificatePropertiesInput",
		&returns,
	)
	return returns
}


func NewKeyVaultCertificateCertificatePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KeyVaultCertificateCertificatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewKeyVaultCertificateCertificatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultCertificate.KeyVaultCertificateCertificatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKeyVaultCertificateCertificatePolicyOutputReference_Override(k KeyVaultCertificateCertificatePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultCertificate.KeyVaultCertificateCertificatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference)SetInternalValue(val *KeyVaultCertificateCertificatePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) PutIssuerParameters(value *KeyVaultCertificateCertificatePolicyIssuerParameters) {
	if err := k.validatePutIssuerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putIssuerParameters",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) PutKeyProperties(value *KeyVaultCertificateCertificatePolicyKeyProperties) {
	if err := k.validatePutKeyPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKeyProperties",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) PutLifetimeAction(value interface{}) {
	if err := k.validatePutLifetimeActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putLifetimeAction",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) PutSecretProperties(value *KeyVaultCertificateCertificatePolicySecretProperties) {
	if err := k.validatePutSecretPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSecretProperties",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) PutX509CertificateProperties(value *KeyVaultCertificateCertificatePolicyX509CertificateProperties) {
	if err := k.validatePutX509CertificatePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putX509CertificateProperties",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) ResetLifetimeAction() {
	_jsii_.InvokeVoid(
		k,
		"resetLifetimeAction",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) ResetX509CertificateProperties() {
	_jsii_.InvokeVoid(
		k,
		"resetX509CertificateProperties",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

