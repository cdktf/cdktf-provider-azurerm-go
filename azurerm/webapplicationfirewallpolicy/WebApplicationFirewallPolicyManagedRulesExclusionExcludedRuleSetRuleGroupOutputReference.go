// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/webapplicationfirewallpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference interface {
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
	ExcludedRules() *[]*string
	SetExcludedRules(val *[]*string)
	ExcludedRulesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RuleGroupName() *string
	SetRuleGroupName(val *string)
	RuleGroupNameInput() *string
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
	ResetExcludedRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference
type jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) ExcludedRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) ExcludedRulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) RuleGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) RuleGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference {
	_init_.Initialize()

	if err := validateNewWebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.webApplicationFirewallPolicy.WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference_Override(w WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.webApplicationFirewallPolicy.WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference)SetExcludedRules(val *[]*string) {
	if err := j.validateSetExcludedRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedRules",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference)SetRuleGroupName(val *string) {
	if err := j.validateSetRuleGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleGroupName",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) ResetExcludedRules() {
	_jsii_.InvokeVoid(
		w,
		"resetExcludedRules",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

