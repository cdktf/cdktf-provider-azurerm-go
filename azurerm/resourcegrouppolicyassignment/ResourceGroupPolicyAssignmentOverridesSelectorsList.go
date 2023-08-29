// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegrouppolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/resourcegrouppolicyassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourceGroupPolicyAssignmentOverridesSelectorsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ResourceGroupPolicyAssignmentOverridesSelectorsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ResourceGroupPolicyAssignmentOverridesSelectorsList
type jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewResourceGroupPolicyAssignmentOverridesSelectorsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ResourceGroupPolicyAssignmentOverridesSelectorsList {
	_init_.Initialize()

	if err := validateNewResourceGroupPolicyAssignmentOverridesSelectorsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupPolicyAssignment.ResourceGroupPolicyAssignmentOverridesSelectorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewResourceGroupPolicyAssignmentOverridesSelectorsList_Override(r ResourceGroupPolicyAssignmentOverridesSelectorsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupPolicyAssignment.ResourceGroupPolicyAssignmentOverridesSelectorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) Get(index *float64) ResourceGroupPolicyAssignmentOverridesSelectorsOutputReference {
	if err := r.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ResourceGroupPolicyAssignmentOverridesSelectorsOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ResourceGroupPolicyAssignmentOverridesSelectorsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

