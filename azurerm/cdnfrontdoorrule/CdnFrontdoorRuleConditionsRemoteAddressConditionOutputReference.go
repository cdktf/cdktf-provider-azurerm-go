// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/cdnfrontdoorrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetOperator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference
type jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleConditionsRemoteAddressConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference_Override(c CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference)SetMatchValues(val *[]*string) {
	if err := j.validateSetMatchValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference)SetNegateCondition(val interface{}) {
	if err := j.validateSetNegateConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		c,
		"resetOperator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsRemoteAddressConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

