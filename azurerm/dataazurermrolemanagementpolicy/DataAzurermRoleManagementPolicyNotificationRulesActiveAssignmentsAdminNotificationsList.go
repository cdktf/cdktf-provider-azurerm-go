// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermrolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/dataazurermrolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList
type jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList {
	_init_.Initialize()

	if err := validateNewDataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermRoleManagementPolicy.DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList_Override(d DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermRoleManagementPolicy.DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) Get(index *float64) DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotificationsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

