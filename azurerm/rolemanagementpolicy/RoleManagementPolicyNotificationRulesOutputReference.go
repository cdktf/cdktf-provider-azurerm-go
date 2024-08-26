// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/rolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleManagementPolicyNotificationRulesOutputReference interface {
	cdktf.ComplexObject
	ActiveAssignments() RoleManagementPolicyNotificationRulesActiveAssignmentsOutputReference
	ActiveAssignmentsInput() *RoleManagementPolicyNotificationRulesActiveAssignments
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
	EligibleActivations() RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference
	EligibleActivationsInput() *RoleManagementPolicyNotificationRulesEligibleActivations
	EligibleAssignments() RoleManagementPolicyNotificationRulesEligibleAssignmentsOutputReference
	EligibleAssignmentsInput() *RoleManagementPolicyNotificationRulesEligibleAssignments
	// Experimental.
	Fqn() *string
	InternalValue() *RoleManagementPolicyNotificationRules
	SetInternalValue(val *RoleManagementPolicyNotificationRules)
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
	PutActiveAssignments(value *RoleManagementPolicyNotificationRulesActiveAssignments)
	PutEligibleActivations(value *RoleManagementPolicyNotificationRulesEligibleActivations)
	PutEligibleAssignments(value *RoleManagementPolicyNotificationRulesEligibleAssignments)
	ResetActiveAssignments()
	ResetEligibleActivations()
	ResetEligibleAssignments()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RoleManagementPolicyNotificationRulesOutputReference
type jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ActiveAssignments() RoleManagementPolicyNotificationRulesActiveAssignmentsOutputReference {
	var returns RoleManagementPolicyNotificationRulesActiveAssignmentsOutputReference
	_jsii_.Get(
		j,
		"activeAssignments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ActiveAssignmentsInput() *RoleManagementPolicyNotificationRulesActiveAssignments {
	var returns *RoleManagementPolicyNotificationRulesActiveAssignments
	_jsii_.Get(
		j,
		"activeAssignmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) EligibleActivations() RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference {
	var returns RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference
	_jsii_.Get(
		j,
		"eligibleActivations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) EligibleActivationsInput() *RoleManagementPolicyNotificationRulesEligibleActivations {
	var returns *RoleManagementPolicyNotificationRulesEligibleActivations
	_jsii_.Get(
		j,
		"eligibleActivationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) EligibleAssignments() RoleManagementPolicyNotificationRulesEligibleAssignmentsOutputReference {
	var returns RoleManagementPolicyNotificationRulesEligibleAssignmentsOutputReference
	_jsii_.Get(
		j,
		"eligibleAssignments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) EligibleAssignmentsInput() *RoleManagementPolicyNotificationRulesEligibleAssignments {
	var returns *RoleManagementPolicyNotificationRulesEligibleAssignments
	_jsii_.Get(
		j,
		"eligibleAssignmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) InternalValue() *RoleManagementPolicyNotificationRules {
	var returns *RoleManagementPolicyNotificationRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoleManagementPolicyNotificationRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RoleManagementPolicyNotificationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewRoleManagementPolicyNotificationRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyNotificationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoleManagementPolicyNotificationRulesOutputReference_Override(r RoleManagementPolicyNotificationRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyNotificationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference)SetInternalValue(val *RoleManagementPolicyNotificationRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) PutActiveAssignments(value *RoleManagementPolicyNotificationRulesActiveAssignments) {
	if err := r.validatePutActiveAssignmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putActiveAssignments",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) PutEligibleActivations(value *RoleManagementPolicyNotificationRulesEligibleActivations) {
	if err := r.validatePutEligibleActivationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEligibleActivations",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) PutEligibleAssignments(value *RoleManagementPolicyNotificationRulesEligibleAssignments) {
	if err := r.validatePutEligibleAssignmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEligibleAssignments",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ResetActiveAssignments() {
	_jsii_.InvokeVoid(
		r,
		"resetActiveAssignments",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ResetEligibleActivations() {
	_jsii_.InvokeVoid(
		r,
		"resetEligibleActivations",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ResetEligibleAssignments() {
	_jsii_.InvokeVoid(
		r,
		"resetEligibleAssignments",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

