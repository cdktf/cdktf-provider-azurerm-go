// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/rolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference interface {
	cdktf.ComplexObject
	AdditionalRecipients() *[]*string
	SetAdditionalRecipients(val *[]*string)
	AdditionalRecipientsInput() *[]*string
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
	DefaultRecipients() interface{}
	SetDefaultRecipients(val interface{})
	DefaultRecipientsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications
	SetInternalValue(val *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications)
	NotificationLevel() *string
	SetNotificationLevel(val *string)
	NotificationLevelInput() *string
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
	ResetAdditionalRecipients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference
type jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) AdditionalRecipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) AdditionalRecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) DefaultRecipients() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) DefaultRecipientsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) InternalValue() *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications {
	var returns *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) NotificationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) NotificationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference {
	_init_.Initialize()

	if err := validateNewRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference_Override(r RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.roleManagementPolicy.RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference)SetAdditionalRecipients(val *[]*string) {
	if err := j.validateSetAdditionalRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalRecipients",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference)SetDefaultRecipients(val interface{}) {
	if err := j.validateSetDefaultRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRecipients",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference)SetInternalValue(val *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference)SetNotificationLevel(val *string) {
	if err := j.validateSetNotificationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationLevel",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) ResetAdditionalRecipients() {
	_jsii_.InvokeVoid(
		r,
		"resetAdditionalRecipients",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

