// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/rolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference interface {
	cdktf.ComplexObject
	AdminNotifications() RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsOutputReference
	AdminNotificationsInput() *RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications
	ApproverNotifications() RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference
	ApproverNotificationsInput() *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications
	AssigneeNotifications() RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsOutputReference
	AssigneeNotificationsInput() *RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications
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
	InternalValue() *RoleManagementPolicyNotificationRulesEligibleActivations
	SetInternalValue(val *RoleManagementPolicyNotificationRulesEligibleActivations)
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
	PutAdminNotifications(value *RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications)
	PutApproverNotifications(value *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications)
	PutAssigneeNotifications(value *RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications)
	ResetAdminNotifications()
	ResetApproverNotifications()
	ResetAssigneeNotifications()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference
type jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) AdminNotifications() RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsOutputReference {
	var returns RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsOutputReference
	_jsii_.Get(
		j,
		"adminNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) AdminNotificationsInput() *RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications {
	var returns *RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications
	_jsii_.Get(
		j,
		"adminNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ApproverNotifications() RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference {
	var returns RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference
	_jsii_.Get(
		j,
		"approverNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ApproverNotificationsInput() *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications {
	var returns *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications
	_jsii_.Get(
		j,
		"approverNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) AssigneeNotifications() RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsOutputReference {
	var returns RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsOutputReference
	_jsii_.Get(
		j,
		"assigneeNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) AssigneeNotificationsInput() *RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications {
	var returns *RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications
	_jsii_.Get(
		j,
		"assigneeNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) InternalValue() *RoleManagementPolicyNotificationRulesEligibleActivations {
	var returns *RoleManagementPolicyNotificationRulesEligibleActivations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference {
	_init_.Initialize()

	if err := validateNewRoleManagementPolicyNotificationRulesEligibleActivationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference_Override(r RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetInternalValue(val *RoleManagementPolicyNotificationRulesEligibleActivations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) PutAdminNotifications(value *RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications) {
	if err := r.validatePutAdminNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAdminNotifications",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) PutApproverNotifications(value *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications) {
	if err := r.validatePutApproverNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putApproverNotifications",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) PutAssigneeNotifications(value *RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications) {
	if err := r.validatePutAssigneeNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAssigneeNotifications",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ResetAdminNotifications() {
	_jsii_.InvokeVoid(
		r,
		"resetAdminNotifications",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ResetApproverNotifications() {
	_jsii_.InvokeVoid(
		r,
		"resetApproverNotifications",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ResetAssigneeNotifications() {
	_jsii_.InvokeVoid(
		r,
		"resetAssigneeNotifications",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

