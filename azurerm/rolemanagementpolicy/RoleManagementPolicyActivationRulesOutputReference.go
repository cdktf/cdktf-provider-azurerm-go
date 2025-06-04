// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/rolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleManagementPolicyActivationRulesOutputReference interface {
	cdktf.ComplexObject
	ApprovalStage() RoleManagementPolicyActivationRulesApprovalStageOutputReference
	ApprovalStageInput() *RoleManagementPolicyActivationRulesApprovalStage
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
	InternalValue() *RoleManagementPolicyActivationRules
	SetInternalValue(val *RoleManagementPolicyActivationRules)
	MaximumDuration() *string
	SetMaximumDuration(val *string)
	MaximumDurationInput() *string
	RequireApproval() interface{}
	SetRequireApproval(val interface{})
	RequireApprovalInput() interface{}
	RequiredConditionalAccessAuthenticationContext() *string
	SetRequiredConditionalAccessAuthenticationContext(val *string)
	RequiredConditionalAccessAuthenticationContextInput() *string
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
	PutApprovalStage(value *RoleManagementPolicyActivationRulesApprovalStage)
	ResetApprovalStage()
	ResetMaximumDuration()
	ResetRequireApproval()
	ResetRequiredConditionalAccessAuthenticationContext()
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

// The jsii proxy struct for RoleManagementPolicyActivationRulesOutputReference
type jsiiProxy_RoleManagementPolicyActivationRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ApprovalStage() RoleManagementPolicyActivationRulesApprovalStageOutputReference {
	var returns RoleManagementPolicyActivationRulesApprovalStageOutputReference
	_jsii_.Get(
		j,
		"approvalStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ApprovalStageInput() *RoleManagementPolicyActivationRulesApprovalStage {
	var returns *RoleManagementPolicyActivationRulesApprovalStage
	_jsii_.Get(
		j,
		"approvalStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) InternalValue() *RoleManagementPolicyActivationRules {
	var returns *RoleManagementPolicyActivationRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) MaximumDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) MaximumDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequireApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequireApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequiredConditionalAccessAuthenticationContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredConditionalAccessAuthenticationContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequiredConditionalAccessAuthenticationContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredConditionalAccessAuthenticationContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequireJustification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireJustification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequireJustificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireJustificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequireMultifactorAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMultifactorAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequireMultifactorAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMultifactorAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequireTicketInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTicketInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) RequireTicketInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTicketInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoleManagementPolicyActivationRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RoleManagementPolicyActivationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewRoleManagementPolicyActivationRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RoleManagementPolicyActivationRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyActivationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoleManagementPolicyActivationRulesOutputReference_Override(r RoleManagementPolicyActivationRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyActivationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetInternalValue(val *RoleManagementPolicyActivationRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetMaximumDuration(val *string) {
	if err := j.validateSetMaximumDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumDuration",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetRequireApproval(val interface{}) {
	if err := j.validateSetRequireApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireApproval",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetRequiredConditionalAccessAuthenticationContext(val *string) {
	if err := j.validateSetRequiredConditionalAccessAuthenticationContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredConditionalAccessAuthenticationContext",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetRequireJustification(val interface{}) {
	if err := j.validateSetRequireJustificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireJustification",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetRequireMultifactorAuthentication(val interface{}) {
	if err := j.validateSetRequireMultifactorAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireMultifactorAuthentication",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetRequireTicketInfo(val interface{}) {
	if err := j.validateSetRequireTicketInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTicketInfo",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) PutApprovalStage(value *RoleManagementPolicyActivationRulesApprovalStage) {
	if err := r.validatePutApprovalStageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putApprovalStage",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ResetApprovalStage() {
	_jsii_.InvokeVoid(
		r,
		"resetApprovalStage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ResetMaximumDuration() {
	_jsii_.InvokeVoid(
		r,
		"resetMaximumDuration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ResetRequireApproval() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireApproval",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ResetRequiredConditionalAccessAuthenticationContext() {
	_jsii_.InvokeVoid(
		r,
		"resetRequiredConditionalAccessAuthenticationContext",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ResetRequireJustification() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireJustification",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ResetRequireMultifactorAuthentication() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireMultifactorAuthentication",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ResetRequireTicketInfo() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireTicketInfo",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RoleManagementPolicyActivationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

