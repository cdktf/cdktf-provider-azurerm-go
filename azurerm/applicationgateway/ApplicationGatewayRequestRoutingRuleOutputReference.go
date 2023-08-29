// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayRequestRoutingRuleOutputReference interface {
	cdktf.ComplexObject
	BackendAddressPoolId() *string
	BackendAddressPoolName() *string
	SetBackendAddressPoolName(val *string)
	BackendAddressPoolNameInput() *string
	BackendHttpSettingsId() *string
	BackendHttpSettingsName() *string
	SetBackendHttpSettingsName(val *string)
	BackendHttpSettingsNameInput() *string
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
	HttpListenerId() *string
	HttpListenerName() *string
	SetHttpListenerName(val *string)
	HttpListenerNameInput() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	RedirectConfigurationId() *string
	RedirectConfigurationName() *string
	SetRedirectConfigurationName(val *string)
	RedirectConfigurationNameInput() *string
	RewriteRuleSetId() *string
	RewriteRuleSetName() *string
	SetRewriteRuleSetName(val *string)
	RewriteRuleSetNameInput() *string
	RuleType() *string
	SetRuleType(val *string)
	RuleTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlPathMapId() *string
	UrlPathMapName() *string
	SetUrlPathMapName(val *string)
	UrlPathMapNameInput() *string
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
	ResetBackendAddressPoolName()
	ResetBackendHttpSettingsName()
	ResetPriority()
	ResetRedirectConfigurationName()
	ResetRewriteRuleSetName()
	ResetUrlPathMapName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayRequestRoutingRuleOutputReference
type jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) BackendAddressPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendAddressPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) BackendAddressPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendAddressPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) BackendAddressPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendAddressPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) BackendHttpSettingsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendHttpSettingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) BackendHttpSettingsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendHttpSettingsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) BackendHttpSettingsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendHttpSettingsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) HttpListenerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpListenerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) HttpListenerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpListenerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) HttpListenerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpListenerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) RedirectConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) RedirectConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) RedirectConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) RewriteRuleSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteRuleSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) RewriteRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteRuleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) RewriteRuleSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteRuleSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) RuleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) RuleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) UrlPathMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlPathMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) UrlPathMapName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlPathMapName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) UrlPathMapNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlPathMapNameInput",
		&returns,
	)
	return returns
}


func NewApplicationGatewayRequestRoutingRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewayRequestRoutingRuleOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayRequestRoutingRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayRequestRoutingRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewayRequestRoutingRuleOutputReference_Override(a ApplicationGatewayRequestRoutingRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayRequestRoutingRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetBackendAddressPoolName(val *string) {
	if err := j.validateSetBackendAddressPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendAddressPoolName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetBackendHttpSettingsName(val *string) {
	if err := j.validateSetBackendHttpSettingsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendHttpSettingsName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetHttpListenerName(val *string) {
	if err := j.validateSetHttpListenerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpListenerName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetRedirectConfigurationName(val *string) {
	if err := j.validateSetRedirectConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectConfigurationName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetRewriteRuleSetName(val *string) {
	if err := j.validateSetRewriteRuleSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rewriteRuleSetName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetRuleType(val *string) {
	if err := j.validateSetRuleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleType",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference)SetUrlPathMapName(val *string) {
	if err := j.validateSetUrlPathMapNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlPathMapName",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ResetBackendAddressPoolName() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendAddressPoolName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ResetBackendHttpSettingsName() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendHttpSettingsName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ResetRedirectConfigurationName() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectConfigurationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ResetRewriteRuleSetName() {
	_jsii_.InvokeVoid(
		a,
		"resetRewriteRuleSetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ResetUrlPathMapName() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlPathMapName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationGatewayRequestRoutingRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

