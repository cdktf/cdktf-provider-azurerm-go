// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsstandardwebtest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/applicationinsightsstandardwebtest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationInsightsStandardWebTestValidationRulesOutputReference interface {
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
	Content() ApplicationInsightsStandardWebTestValidationRulesContentOutputReference
	ContentInput() *ApplicationInsightsStandardWebTestValidationRulesContent
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExpectedStatusCode() *float64
	SetExpectedStatusCode(val *float64)
	ExpectedStatusCodeInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ApplicationInsightsStandardWebTestValidationRules
	SetInternalValue(val *ApplicationInsightsStandardWebTestValidationRules)
	SslCertRemainingLifetime() *float64
	SetSslCertRemainingLifetime(val *float64)
	SslCertRemainingLifetimeInput() *float64
	SslCheckEnabled() interface{}
	SetSslCheckEnabled(val interface{})
	SslCheckEnabledInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutContent(value *ApplicationInsightsStandardWebTestValidationRulesContent)
	ResetContent()
	ResetExpectedStatusCode()
	ResetSslCertRemainingLifetime()
	ResetSslCheckEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationInsightsStandardWebTestValidationRulesOutputReference
type jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) Content() ApplicationInsightsStandardWebTestValidationRulesContentOutputReference {
	var returns ApplicationInsightsStandardWebTestValidationRulesContentOutputReference
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ContentInput() *ApplicationInsightsStandardWebTestValidationRulesContent {
	var returns *ApplicationInsightsStandardWebTestValidationRulesContent
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ExpectedStatusCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expectedStatusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ExpectedStatusCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expectedStatusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) InternalValue() *ApplicationInsightsStandardWebTestValidationRules {
	var returns *ApplicationInsightsStandardWebTestValidationRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) SslCertRemainingLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sslCertRemainingLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) SslCertRemainingLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sslCertRemainingLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) SslCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) SslCheckEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslCheckEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationInsightsStandardWebTestValidationRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationInsightsStandardWebTestValidationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationInsightsStandardWebTestValidationRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationInsightsStandardWebTest.ApplicationInsightsStandardWebTestValidationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationInsightsStandardWebTestValidationRulesOutputReference_Override(a ApplicationInsightsStandardWebTestValidationRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationInsightsStandardWebTest.ApplicationInsightsStandardWebTestValidationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference)SetExpectedStatusCode(val *float64) {
	if err := j.validateSetExpectedStatusCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedStatusCode",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference)SetInternalValue(val *ApplicationInsightsStandardWebTestValidationRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference)SetSslCertRemainingLifetime(val *float64) {
	if err := j.validateSetSslCertRemainingLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertRemainingLifetime",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference)SetSslCheckEnabled(val interface{}) {
	if err := j.validateSetSslCheckEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) PutContent(value *ApplicationInsightsStandardWebTestValidationRulesContent) {
	if err := a.validatePutContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		a,
		"resetContent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ResetExpectedStatusCode() {
	_jsii_.InvokeVoid(
		a,
		"resetExpectedStatusCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ResetSslCertRemainingLifetime() {
	_jsii_.InvokeVoid(
		a,
		"resetSslCertRemainingLifetime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ResetSslCheckEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSslCheckEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsightsStandardWebTestValidationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

