package keyvaultcertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/keyvaultcertificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference interface {
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
	DaysBeforeExpiry() *float64
	SetDaysBeforeExpiry(val *float64)
	DaysBeforeExpiryInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *KeyVaultCertificateCertificatePolicyLifetimeActionTrigger
	SetInternalValue(val *KeyVaultCertificateCertificatePolicyLifetimeActionTrigger)
	LifetimePercentage() *float64
	SetLifetimePercentage(val *float64)
	LifetimePercentageInput() *float64
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
	ResetDaysBeforeExpiry()
	ResetLifetimePercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference
type jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) DaysBeforeExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysBeforeExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) DaysBeforeExpiryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysBeforeExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) InternalValue() *KeyVaultCertificateCertificatePolicyLifetimeActionTrigger {
	var returns *KeyVaultCertificateCertificatePolicyLifetimeActionTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) LifetimePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) LifetimePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference {
	_init_.Initialize()

	if err := validateNewKeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultCertificate.KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference_Override(k KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultCertificate.KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference)SetDaysBeforeExpiry(val *float64) {
	if err := j.validateSetDaysBeforeExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysBeforeExpiry",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference)SetInternalValue(val *KeyVaultCertificateCertificatePolicyLifetimeActionTrigger) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference)SetLifetimePercentage(val *float64) {
	if err := j.validateSetLifetimePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetimePercentage",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) ResetDaysBeforeExpiry() {
	_jsii_.InvokeVoid(
		k,
		"resetDaysBeforeExpiry",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) ResetLifetimePercentage() {
	_jsii_.InvokeVoid(
		k,
		"resetLifetimePercentage",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KeyVaultCertificateCertificatePolicyLifetimeActionTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

