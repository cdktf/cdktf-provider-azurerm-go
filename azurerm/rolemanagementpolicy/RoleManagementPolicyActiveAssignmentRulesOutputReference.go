// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/rolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleManagementPolicyActiveAssignmentRulesOutputReference interface {
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
	ExpirationRequired() interface{}
	SetExpirationRequired(val interface{})
	ExpirationRequiredInput() interface{}
	ExpireAfter() *string
	SetExpireAfter(val *string)
	ExpireAfterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *RoleManagementPolicyActiveAssignmentRules
	SetInternalValue(val *RoleManagementPolicyActiveAssignmentRules)
	RequireJustification() interface{}
	SetRequireJustification(val interface{})
	RequireJustificationInput() interface{}
	RequireMultifactorAuthentication() interface{}
	SetRequireMultifactorAuthentication(val interface{})
	RequireMultifactorAuthenticationInput() interface{}
	RequireTicketInfo() interface{}
	SetRequireTicketInfo(val interface{})
	RequireTicketInfoInput() interface{}
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
	ResetExpirationRequired()
	ResetExpireAfter()
	ResetRequireJustification()
	ResetRequireMultifactorAuthentication()
	ResetRequireTicketInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RoleManagementPolicyActiveAssignmentRulesOutputReference
type jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ExpirationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ExpirationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ExpireAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ExpireAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) InternalValue() *RoleManagementPolicyActiveAssignmentRules {
	var returns *RoleManagementPolicyActiveAssignmentRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) RequireJustification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireJustification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) RequireJustificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireJustificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) RequireMultifactorAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMultifactorAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) RequireMultifactorAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMultifactorAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) RequireTicketInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTicketInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) RequireTicketInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTicketInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoleManagementPolicyActiveAssignmentRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RoleManagementPolicyActiveAssignmentRulesOutputReference {
	_init_.Initialize()

	if err := validateNewRoleManagementPolicyActiveAssignmentRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyActiveAssignmentRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoleManagementPolicyActiveAssignmentRulesOutputReference_Override(r RoleManagementPolicyActiveAssignmentRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyActiveAssignmentRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetExpirationRequired(val interface{}) {
	if err := j.validateSetExpirationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationRequired",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetExpireAfter(val *string) {
	if err := j.validateSetExpireAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireAfter",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetInternalValue(val *RoleManagementPolicyActiveAssignmentRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetRequireJustification(val interface{}) {
	if err := j.validateSetRequireJustificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireJustification",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetRequireMultifactorAuthentication(val interface{}) {
	if err := j.validateSetRequireMultifactorAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireMultifactorAuthentication",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetRequireTicketInfo(val interface{}) {
	if err := j.validateSetRequireTicketInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTicketInfo",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ResetExpirationRequired() {
	_jsii_.InvokeVoid(
		r,
		"resetExpirationRequired",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ResetExpireAfter() {
	_jsii_.InvokeVoid(
		r,
		"resetExpireAfter",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ResetRequireJustification() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireJustification",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ResetRequireMultifactorAuthentication() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireMultifactorAuthentication",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ResetRequireTicketInfo() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireTicketInfo",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActiveAssignmentRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

