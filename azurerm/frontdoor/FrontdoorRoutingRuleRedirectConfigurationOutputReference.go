// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/frontdoor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrontdoorRoutingRuleRedirectConfigurationOutputReference interface {
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
	CustomFragment() *string
	SetCustomFragment(val *string)
	CustomFragmentInput() *string
	CustomHost() *string
	SetCustomHost(val *string)
	CustomHostInput() *string
	CustomPath() *string
	SetCustomPath(val *string)
	CustomPathInput() *string
	CustomQueryString() *string
	SetCustomQueryString(val *string)
	CustomQueryStringInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *FrontdoorRoutingRuleRedirectConfiguration
	SetInternalValue(val *FrontdoorRoutingRuleRedirectConfiguration)
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
	ResetCustomFragment()
	ResetCustomHost()
	ResetCustomPath()
	ResetCustomQueryString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorRoutingRuleRedirectConfigurationOutputReference
type jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomFragment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomFragmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomQueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomQueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) InternalValue() *FrontdoorRoutingRuleRedirectConfiguration {
	var returns *FrontdoorRoutingRuleRedirectConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) RedirectProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) RedirectProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) RedirectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) RedirectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorRoutingRuleRedirectConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FrontdoorRoutingRuleRedirectConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFrontdoorRoutingRuleRedirectConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleRedirectConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrontdoorRoutingRuleRedirectConfigurationOutputReference_Override(f FrontdoorRoutingRuleRedirectConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleRedirectConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetCustomFragment(val *string) {
	if err := j.validateSetCustomFragmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customFragment",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetCustomHost(val *string) {
	if err := j.validateSetCustomHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customHost",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetCustomPath(val *string) {
	if err := j.validateSetCustomPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPath",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetCustomQueryString(val *string) {
	if err := j.validateSetCustomQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customQueryString",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetInternalValue(val *FrontdoorRoutingRuleRedirectConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetRedirectProtocol(val *string) {
	if err := j.validateSetRedirectProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectProtocol",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetRedirectType(val *string) {
	if err := j.validateSetRedirectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectType",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ResetCustomFragment() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomFragment",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ResetCustomHost() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomHost",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ResetCustomPath() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ResetCustomQueryString() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomQueryString",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

