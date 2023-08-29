// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/frontdoorfirewallpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList interface {
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
	Get(index *float64) FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList
type jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList {
	_init_.Initialize()

	if err := validateNewFrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoorFirewallPolicy.FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList_Override(f FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoorFirewallPolicy.FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) Get(index *float64) FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference {
	if err := f.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

