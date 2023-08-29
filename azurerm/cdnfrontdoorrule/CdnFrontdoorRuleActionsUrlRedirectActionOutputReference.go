// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/cdnfrontdoorrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorRuleActionsUrlRedirectActionOutputReference interface {
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
	DestinationFragment() *string
	SetDestinationFragment(val *string)
	DestinationFragmentInput() *string
	DestinationHostname() *string
	SetDestinationHostname(val *string)
	DestinationHostnameInput() *string
	DestinationPath() *string
	SetDestinationPath(val *string)
	DestinationPathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnFrontdoorRuleActionsUrlRedirectAction
	SetInternalValue(val *CdnFrontdoorRuleActionsUrlRedirectAction)
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	RedirectProtocol() *string
	SetRedirectProtocol(val *string)
	RedirectProtocolInput() *string
	RedirectType() *string
	SetRedirectType(val *string)
	RedirectTypeInput() *string
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
	ResetDestinationFragment()
	ResetDestinationPath()
	ResetQueryString()
	ResetRedirectProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleActionsUrlRedirectActionOutputReference
type jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) DestinationFragment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) DestinationFragmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) DestinationHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) DestinationHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) DestinationPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) DestinationPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) InternalValue() *CdnFrontdoorRuleActionsUrlRedirectAction {
	var returns *CdnFrontdoorRuleActionsUrlRedirectAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) RedirectProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) RedirectProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) RedirectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) RedirectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleActionsUrlRedirectActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleActionsUrlRedirectActionOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleActionsUrlRedirectActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsUrlRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleActionsUrlRedirectActionOutputReference_Override(c CdnFrontdoorRuleActionsUrlRedirectActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsUrlRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetDestinationFragment(val *string) {
	if err := j.validateSetDestinationFragmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFragment",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetDestinationHostname(val *string) {
	if err := j.validateSetDestinationHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationHostname",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetDestinationPath(val *string) {
	if err := j.validateSetDestinationPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPath",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetInternalValue(val *CdnFrontdoorRuleActionsUrlRedirectAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetRedirectProtocol(val *string) {
	if err := j.validateSetRedirectProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectProtocol",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetRedirectType(val *string) {
	if err := j.validateSetRedirectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectType",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) ResetDestinationFragment() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationFragment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) ResetDestinationPath() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) ResetRedirectProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetRedirectProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

