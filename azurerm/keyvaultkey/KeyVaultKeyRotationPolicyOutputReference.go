// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultkey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/keyvaultkey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultKeyRotationPolicyOutputReference interface {
	cdktf.ComplexObject
	Automatic() KeyVaultKeyRotationPolicyAutomaticOutputReference
	AutomaticInput() *KeyVaultKeyRotationPolicyAutomatic
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
	ExpireAfter() *string
	SetExpireAfter(val *string)
	ExpireAfterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KeyVaultKeyRotationPolicy
	SetInternalValue(val *KeyVaultKeyRotationPolicy)
	NotifyBeforeExpiry() *string
	SetNotifyBeforeExpiry(val *string)
	NotifyBeforeExpiryInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAutomatic(value *KeyVaultKeyRotationPolicyAutomatic)
	ResetAutomatic()
	ResetExpireAfter()
	ResetNotifyBeforeExpiry()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyVaultKeyRotationPolicyOutputReference
type jsiiProxy_KeyVaultKeyRotationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) Automatic() KeyVaultKeyRotationPolicyAutomaticOutputReference {
	var returns KeyVaultKeyRotationPolicyAutomaticOutputReference
	_jsii_.Get(
		j,
		"automatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) AutomaticInput() *KeyVaultKeyRotationPolicyAutomatic {
	var returns *KeyVaultKeyRotationPolicyAutomatic
	_jsii_.Get(
		j,
		"automaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ExpireAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ExpireAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) InternalValue() *KeyVaultKeyRotationPolicy {
	var returns *KeyVaultKeyRotationPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) NotifyBeforeExpiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyBeforeExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) NotifyBeforeExpiryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyBeforeExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKeyVaultKeyRotationPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KeyVaultKeyRotationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewKeyVaultKeyRotationPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultKeyRotationPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultKey.KeyVaultKeyRotationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKeyVaultKeyRotationPolicyOutputReference_Override(k KeyVaultKeyRotationPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultKey.KeyVaultKeyRotationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference)SetExpireAfter(val *string) {
	if err := j.validateSetExpireAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireAfter",
		val,
	)
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference)SetInternalValue(val *KeyVaultKeyRotationPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference)SetNotifyBeforeExpiry(val *string) {
	if err := j.validateSetNotifyBeforeExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyBeforeExpiry",
		val,
	)
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) PutAutomatic(value *KeyVaultKeyRotationPolicyAutomatic) {
	if err := k.validatePutAutomaticParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAutomatic",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ResetAutomatic() {
	_jsii_.InvokeVoid(
		k,
		"resetAutomatic",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ResetExpireAfter() {
	_jsii_.InvokeVoid(
		k,
		"resetExpireAfter",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ResetNotifyBeforeExpiry() {
	_jsii_.InvokeVoid(
		k,
		"resetNotifyBeforeExpiry",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultKeyRotationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

