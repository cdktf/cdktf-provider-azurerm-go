// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedhardwaresecuritymoduleroledefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/keyvaultmanagedhardwaresecuritymoduleroledefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference interface {
	cdktf.ComplexObject
	Actions() *[]*string
	SetActions(val *[]*string)
	ActionsInput() *[]*string
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
	DataActions() *[]*string
	SetDataActions(val *[]*string)
	DataActionsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NotActions() *[]*string
	SetNotActions(val *[]*string)
	NotActionsInput() *[]*string
	NotDataActions() *[]*string
	SetNotDataActions(val *[]*string)
	NotDataActionsInput() *[]*string
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
	ResetActions()
	ResetDataActions()
	ResetNotActions()
	ResetNotDataActions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference
type jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) DataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) DataActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) NotActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) NotActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) NotDataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notDataActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) NotDataActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notDataActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference {
	_init_.Initialize()

	if err := validateNewKeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleRoleDefinition.KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference_Override(k KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleRoleDefinition.KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetDataActions(val *[]*string) {
	if err := j.validateSetDataActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataActions",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetNotActions(val *[]*string) {
	if err := j.validateSetNotActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notActions",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetNotDataActions(val *[]*string) {
	if err := j.validateSetNotDataActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notDataActions",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		k,
		"resetActions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ResetDataActions() {
	_jsii_.InvokeVoid(
		k,
		"resetDataActions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ResetNotActions() {
	_jsii_.InvokeVoid(
		k,
		"resetNotActions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ResetNotDataActions() {
	_jsii_.InvokeVoid(
		k,
		"resetNotDataActions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleRoleDefinitionPermissionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

