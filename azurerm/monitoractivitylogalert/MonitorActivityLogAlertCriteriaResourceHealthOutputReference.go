// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractivitylogalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/monitoractivitylogalert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorActivityLogAlertCriteriaResourceHealthOutputReference interface {
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
	Current() *[]*string
	SetCurrent(val *[]*string)
	CurrentInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorActivityLogAlertCriteriaResourceHealth
	SetInternalValue(val *MonitorActivityLogAlertCriteriaResourceHealth)
	Previous() *[]*string
	SetPrevious(val *[]*string)
	PreviousInput() *[]*string
	Reason() *[]*string
	SetReason(val *[]*string)
	ReasonInput() *[]*string
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
	ResetCurrent()
	ResetPrevious()
	ResetReason()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertCriteriaResourceHealthOutputReference
type jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Current() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"current",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) CurrentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"currentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) InternalValue() *MonitorActivityLogAlertCriteriaResourceHealth {
	var returns *MonitorActivityLogAlertCriteriaResourceHealth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Previous() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) PreviousInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Reason() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ReasonInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertCriteriaResourceHealthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorActivityLogAlertCriteriaResourceHealthOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorActivityLogAlertCriteriaResourceHealthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaResourceHealthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertCriteriaResourceHealthOutputReference_Override(m MonitorActivityLogAlertCriteriaResourceHealthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaResourceHealthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference)SetCurrent(val *[]*string) {
	if err := j.validateSetCurrentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"current",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference)SetInternalValue(val *MonitorActivityLogAlertCriteriaResourceHealth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference)SetPrevious(val *[]*string) {
	if err := j.validateSetPreviousParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"previous",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference)SetReason(val *[]*string) {
	if err := j.validateSetReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reason",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ResetCurrent() {
	_jsii_.InvokeVoid(
		m,
		"resetCurrent",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ResetPrevious() {
	_jsii_.InvokeVoid(
		m,
		"resetPrevious",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ResetReason() {
	_jsii_.InvokeVoid(
		m,
		"resetReason",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

