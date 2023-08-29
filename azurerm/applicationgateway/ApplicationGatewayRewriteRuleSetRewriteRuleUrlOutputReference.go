// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference interface {
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
	Components() *string
	SetComponents(val *string)
	ComponentsInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ApplicationGatewayRewriteRuleSetRewriteRuleUrl
	SetInternalValue(val *ApplicationGatewayRewriteRuleSetRewriteRuleUrl)
	Path() *string
	SetPath(val *string)
	PathInput() *string
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	Reroute() interface{}
	SetReroute(val interface{})
	RerouteInput() interface{}
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
	ResetComponents()
	ResetPath()
	ResetQueryString()
	ResetReroute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference
type jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) Components() *string {
	var returns *string
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ComponentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) InternalValue() *ApplicationGatewayRewriteRuleSetRewriteRuleUrl {
	var returns *ApplicationGatewayRewriteRuleSetRewriteRuleUrl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) Reroute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reroute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) RerouteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rerouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference_Override(a ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetComponents(val *string) {
	if err := j.validateSetComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetInternalValue(val *ApplicationGatewayRewriteRuleSetRewriteRuleUrl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetReroute(val interface{}) {
	if err := j.validateSetRerouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reroute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ResetComponents() {
	_jsii_.InvokeVoid(
		a,
		"resetComponents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ResetReroute() {
	_jsii_.InvokeVoid(
		a,
		"resetReroute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

