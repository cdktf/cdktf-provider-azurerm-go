// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/eventgridnamespace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventgridNamespaceTopicSpacesConfigurationOutputReference interface {
	cdktf.ComplexObject
	AlternativeAuthenticationNameSource() *[]*string
	SetAlternativeAuthenticationNameSource(val *[]*string)
	AlternativeAuthenticationNameSourceInput() *[]*string
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
	DynamicRoutingEnrichment() EventgridNamespaceTopicSpacesConfigurationDynamicRoutingEnrichmentList
	DynamicRoutingEnrichmentInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumClientSessionsPerAuthenticationName() *float64
	SetMaximumClientSessionsPerAuthenticationName(val *float64)
	MaximumClientSessionsPerAuthenticationNameInput() *float64
	MaximumSessionExpiryInHours() *float64
	SetMaximumSessionExpiryInHours(val *float64)
	MaximumSessionExpiryInHoursInput() *float64
	RouteTopicId() *string
	SetRouteTopicId(val *string)
	RouteTopicIdInput() *string
	StaticRoutingEnrichment() EventgridNamespaceTopicSpacesConfigurationStaticRoutingEnrichmentList
	StaticRoutingEnrichmentInput() interface{}
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
	PutDynamicRoutingEnrichment(value interface{})
	PutStaticRoutingEnrichment(value interface{})
	ResetAlternativeAuthenticationNameSource()
	ResetDynamicRoutingEnrichment()
	ResetMaximumClientSessionsPerAuthenticationName()
	ResetMaximumSessionExpiryInHours()
	ResetRouteTopicId()
	ResetStaticRoutingEnrichment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridNamespaceTopicSpacesConfigurationOutputReference
type jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) AlternativeAuthenticationNameSource() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alternativeAuthenticationNameSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) AlternativeAuthenticationNameSourceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alternativeAuthenticationNameSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) DynamicRoutingEnrichment() EventgridNamespaceTopicSpacesConfigurationDynamicRoutingEnrichmentList {
	var returns EventgridNamespaceTopicSpacesConfigurationDynamicRoutingEnrichmentList
	_jsii_.Get(
		j,
		"dynamicRoutingEnrichment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) DynamicRoutingEnrichmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicRoutingEnrichmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) MaximumClientSessionsPerAuthenticationName() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumClientSessionsPerAuthenticationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) MaximumClientSessionsPerAuthenticationNameInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumClientSessionsPerAuthenticationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) MaximumSessionExpiryInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumSessionExpiryInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) MaximumSessionExpiryInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumSessionExpiryInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) RouteTopicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTopicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) RouteTopicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTopicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) StaticRoutingEnrichment() EventgridNamespaceTopicSpacesConfigurationStaticRoutingEnrichmentList {
	var returns EventgridNamespaceTopicSpacesConfigurationStaticRoutingEnrichmentList
	_jsii_.Get(
		j,
		"staticRoutingEnrichment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) StaticRoutingEnrichmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticRoutingEnrichmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridNamespaceTopicSpacesConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridNamespaceTopicSpacesConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEventgridNamespaceTopicSpacesConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridNamespace.EventgridNamespaceTopicSpacesConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridNamespaceTopicSpacesConfigurationOutputReference_Override(e EventgridNamespaceTopicSpacesConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridNamespace.EventgridNamespaceTopicSpacesConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetAlternativeAuthenticationNameSource(val *[]*string) {
	if err := j.validateSetAlternativeAuthenticationNameSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternativeAuthenticationNameSource",
		val,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetMaximumClientSessionsPerAuthenticationName(val *float64) {
	if err := j.validateSetMaximumClientSessionsPerAuthenticationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumClientSessionsPerAuthenticationName",
		val,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetMaximumSessionExpiryInHours(val *float64) {
	if err := j.validateSetMaximumSessionExpiryInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumSessionExpiryInHours",
		val,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetRouteTopicId(val *string) {
	if err := j.validateSetRouteTopicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTopicId",
		val,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) PutDynamicRoutingEnrichment(value interface{}) {
	if err := e.validatePutDynamicRoutingEnrichmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDynamicRoutingEnrichment",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) PutStaticRoutingEnrichment(value interface{}) {
	if err := e.validatePutStaticRoutingEnrichmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStaticRoutingEnrichment",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ResetAlternativeAuthenticationNameSource() {
	_jsii_.InvokeVoid(
		e,
		"resetAlternativeAuthenticationNameSource",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ResetDynamicRoutingEnrichment() {
	_jsii_.InvokeVoid(
		e,
		"resetDynamicRoutingEnrichment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ResetMaximumClientSessionsPerAuthenticationName() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximumClientSessionsPerAuthenticationName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ResetMaximumSessionExpiryInHours() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximumSessionExpiryInHours",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ResetRouteTopicId() {
	_jsii_.InvokeVoid(
		e,
		"resetRouteTopicId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ResetStaticRoutingEnrichment() {
	_jsii_.InvokeVoid(
		e,
		"resetStaticRoutingEnrichment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridNamespaceTopicSpacesConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

