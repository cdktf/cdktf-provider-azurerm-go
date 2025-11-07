// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webpubsubhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/webpubsubhub/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebPubsubHubEventHandlerOutputReference interface {
	cdktf.ComplexObject
	Auth() WebPubsubHubEventHandlerAuthOutputReference
	AuthInput() *WebPubsubHubEventHandlerAuth
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
	SystemEvents() *[]*string
	SetSystemEvents(val *[]*string)
	SystemEventsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlTemplate() *string
	SetUrlTemplate(val *string)
	UrlTemplateInput() *string
	UserEventPattern() *string
	SetUserEventPattern(val *string)
	UserEventPatternInput() *string
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
	PutAuth(value *WebPubsubHubEventHandlerAuth)
	ResetAuth()
	ResetSystemEvents()
	ResetUserEventPattern()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WebPubsubHubEventHandlerOutputReference
type jsiiProxy_WebPubsubHubEventHandlerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) Auth() WebPubsubHubEventHandlerAuthOutputReference {
	var returns WebPubsubHubEventHandlerAuthOutputReference
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) AuthInput() *WebPubsubHubEventHandlerAuth {
	var returns *WebPubsubHubEventHandlerAuth
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) SystemEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) SystemEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) UrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) UrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) UserEventPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userEventPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference) UserEventPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userEventPatternInput",
		&returns,
	)
	return returns
}


func NewWebPubsubHubEventHandlerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WebPubsubHubEventHandlerOutputReference {
	_init_.Initialize()

	if err := validateNewWebPubsubHubEventHandlerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebPubsubHubEventHandlerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.webPubsubHub.WebPubsubHubEventHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWebPubsubHubEventHandlerOutputReference_Override(w WebPubsubHubEventHandlerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.webPubsubHub.WebPubsubHubEventHandlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference)SetSystemEvents(val *[]*string) {
	if err := j.validateSetSystemEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemEvents",
		val,
	)
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference)SetUrlTemplate(val *string) {
	if err := j.validateSetUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlTemplate",
		val,
	)
}

func (j *jsiiProxy_WebPubsubHubEventHandlerOutputReference)SetUserEventPattern(val *string) {
	if err := j.validateSetUserEventPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEventPattern",
		val,
	)
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) PutAuth(value *WebPubsubHubEventHandlerAuth) {
	if err := w.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAuth",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		w,
		"resetAuth",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) ResetSystemEvents() {
	_jsii_.InvokeVoid(
		w,
		"resetSystemEvents",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) ResetUserEventPattern() {
	_jsii_.InvokeVoid(
		w,
		"resetUserEventPattern",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubHubEventHandlerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

