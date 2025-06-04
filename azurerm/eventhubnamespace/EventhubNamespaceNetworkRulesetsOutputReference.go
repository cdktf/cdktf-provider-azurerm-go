// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventhubnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/eventhubnamespace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventhubNamespaceNetworkRulesetsOutputReference interface {
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
	DefaultAction() *string
	SetDefaultAction(val *string)
	DefaultActionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpRule() EventhubNamespaceNetworkRulesetsIpRuleList
	IpRuleInput() interface{}
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedServiceAccessEnabled() interface{}
	SetTrustedServiceAccessEnabled(val interface{})
	TrustedServiceAccessEnabledInput() interface{}
	VirtualNetworkRule() EventhubNamespaceNetworkRulesetsVirtualNetworkRuleList
	VirtualNetworkRuleInput() interface{}
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
	PutIpRule(value interface{})
	PutVirtualNetworkRule(value interface{})
	ResetDefaultAction()
	ResetIpRule()
	ResetPublicNetworkAccessEnabled()
	ResetTrustedServiceAccessEnabled()
	ResetVirtualNetworkRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventhubNamespaceNetworkRulesetsOutputReference
type jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) DefaultAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) DefaultActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) IpRule() EventhubNamespaceNetworkRulesetsIpRuleList {
	var returns EventhubNamespaceNetworkRulesetsIpRuleList
	_jsii_.Get(
		j,
		"ipRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) IpRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) TrustedServiceAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedServiceAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) TrustedServiceAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedServiceAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) VirtualNetworkRule() EventhubNamespaceNetworkRulesetsVirtualNetworkRuleList {
	var returns EventhubNamespaceNetworkRulesetsVirtualNetworkRuleList
	_jsii_.Get(
		j,
		"virtualNetworkRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) VirtualNetworkRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualNetworkRuleInput",
		&returns,
	)
	return returns
}


func NewEventhubNamespaceNetworkRulesetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventhubNamespaceNetworkRulesetsOutputReference {
	_init_.Initialize()

	if err := validateNewEventhubNamespaceNetworkRulesetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventhubNamespace.EventhubNamespaceNetworkRulesetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventhubNamespaceNetworkRulesetsOutputReference_Override(e EventhubNamespaceNetworkRulesetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventhubNamespace.EventhubNamespaceNetworkRulesetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference)SetDefaultAction(val *string) {
	if err := j.validateSetDefaultActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAction",
		val,
	)
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference)SetTrustedServiceAccessEnabled(val interface{}) {
	if err := j.validateSetTrustedServiceAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedServiceAccessEnabled",
		val,
	)
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) PutIpRule(value interface{}) {
	if err := e.validatePutIpRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpRule",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) PutVirtualNetworkRule(value interface{}) {
	if err := e.validatePutVirtualNetworkRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVirtualNetworkRule",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ResetDefaultAction() {
	_jsii_.InvokeVoid(
		e,
		"resetDefaultAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ResetIpRule() {
	_jsii_.InvokeVoid(
		e,
		"resetIpRule",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ResetTrustedServiceAccessEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetTrustedServiceAccessEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ResetVirtualNetworkRule() {
	_jsii_.InvokeVoid(
		e,
		"resetVirtualNetworkRule",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventhubNamespaceNetworkRulesetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

