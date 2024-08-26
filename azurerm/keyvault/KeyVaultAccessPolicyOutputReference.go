// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/keyvault/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultAccessPolicyOutputReference interface {
	cdktf.ComplexObject
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	CertificatePermissions() *[]*string
	SetCertificatePermissions(val *[]*string)
	CertificatePermissionsInput() *[]*string
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
	KeyPermissions() *[]*string
	SetKeyPermissions(val *[]*string)
	KeyPermissionsInput() *[]*string
	ObjectId() *string
	SetObjectId(val *string)
	ObjectIdInput() *string
	SecretPermissions() *[]*string
	SetSecretPermissions(val *[]*string)
	SecretPermissionsInput() *[]*string
	StoragePermissions() *[]*string
	SetStoragePermissions(val *[]*string)
	StoragePermissionsInput() *[]*string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetApplicationId()
	ResetCertificatePermissions()
	ResetKeyPermissions()
	ResetObjectId()
	ResetSecretPermissions()
	ResetStoragePermissions()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyVaultAccessPolicyOutputReference
type jsiiProxy_KeyVaultAccessPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) CertificatePermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) CertificatePermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificatePermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) KeyPermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) KeyPermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) ObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) SecretPermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) SecretPermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) StoragePermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storagePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) StoragePermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storagePermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKeyVaultAccessPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KeyVaultAccessPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewKeyVaultAccessPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultAccessPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVault.KeyVaultAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKeyVaultAccessPolicyOutputReference_Override(k KeyVaultAccessPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVault.KeyVaultAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetCertificatePermissions(val *[]*string) {
	if err := j.validateSetCertificatePermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificatePermissions",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetKeyPermissions(val *[]*string) {
	if err := j.validateSetKeyPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPermissions",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetObjectId(val *string) {
	if err := j.validateSetObjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectId",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetSecretPermissions(val *[]*string) {
	if err := j.validateSetSecretPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretPermissions",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetStoragePermissions(val *[]*string) {
	if err := j.validateSetStoragePermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePermissions",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyVaultAccessPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ResetApplicationId() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ResetCertificatePermissions() {
	_jsii_.InvokeVoid(
		k,
		"resetCertificatePermissions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ResetKeyPermissions() {
	_jsii_.InvokeVoid(
		k,
		"resetKeyPermissions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ResetObjectId() {
	_jsii_.InvokeVoid(
		k,
		"resetObjectId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ResetSecretPermissions() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretPermissions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ResetStoragePermissions() {
	_jsii_.InvokeVoid(
		k,
		"resetStoragePermissions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		k,
		"resetTenantId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KeyVaultAccessPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

