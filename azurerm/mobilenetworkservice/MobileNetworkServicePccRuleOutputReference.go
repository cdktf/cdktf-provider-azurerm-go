// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mobilenetworkservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkServicePccRuleOutputReference interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	Precedence() *float64
	SetPrecedence(val *float64)
	PrecedenceInput() *float64
	QosPolicy() MobileNetworkServicePccRuleQosPolicyOutputReference
	QosPolicyInput() *MobileNetworkServicePccRuleQosPolicy
	ServiceDataFlowTemplate() MobileNetworkServicePccRuleServiceDataFlowTemplateList
	ServiceDataFlowTemplateInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficControlEnabled() interface{}
	SetTrafficControlEnabled(val interface{})
	TrafficControlEnabledInput() interface{}
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
	PutQosPolicy(value *MobileNetworkServicePccRuleQosPolicy)
	PutServiceDataFlowTemplate(value interface{})
	ResetQosPolicy()
	ResetTrafficControlEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MobileNetworkServicePccRuleOutputReference
type jsiiProxy_MobileNetworkServicePccRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) PrecedenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) QosPolicy() MobileNetworkServicePccRuleQosPolicyOutputReference {
	var returns MobileNetworkServicePccRuleQosPolicyOutputReference
	_jsii_.Get(
		j,
		"qosPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) QosPolicyInput() *MobileNetworkServicePccRuleQosPolicy {
	var returns *MobileNetworkServicePccRuleQosPolicy
	_jsii_.Get(
		j,
		"qosPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) ServiceDataFlowTemplate() MobileNetworkServicePccRuleServiceDataFlowTemplateList {
	var returns MobileNetworkServicePccRuleServiceDataFlowTemplateList
	_jsii_.Get(
		j,
		"serviceDataFlowTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) ServiceDataFlowTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceDataFlowTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) TrafficControlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficControlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference) TrafficControlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficControlEnabledInput",
		&returns,
	)
	return returns
}


func NewMobileNetworkServicePccRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MobileNetworkServicePccRuleOutputReference {
	_init_.Initialize()

	if err := validateNewMobileNetworkServicePccRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkServicePccRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkService.MobileNetworkServicePccRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMobileNetworkServicePccRuleOutputReference_Override(m MobileNetworkServicePccRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkService.MobileNetworkServicePccRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference)SetPrecedence(val *float64) {
	if err := j.validateSetPrecedenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precedence",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServicePccRuleOutputReference)SetTrafficControlEnabled(val interface{}) {
	if err := j.validateSetTrafficControlEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficControlEnabled",
		val,
	)
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) PutQosPolicy(value *MobileNetworkServicePccRuleQosPolicy) {
	if err := m.validatePutQosPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putQosPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) PutServiceDataFlowTemplate(value interface{}) {
	if err := m.validatePutServiceDataFlowTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putServiceDataFlowTemplate",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) ResetQosPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetQosPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) ResetTrafficControlEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetTrafficControlEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServicePccRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

