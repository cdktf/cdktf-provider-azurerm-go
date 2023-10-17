// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractionrulesuppression

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/monitoractionrulesuppression/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorActionRuleSuppressionSuppressionScheduleOutputReference interface {
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
	EndDateUtc() *string
	SetEndDateUtc(val *string)
	EndDateUtcInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorActionRuleSuppressionSuppressionSchedule
	SetInternalValue(val *MonitorActionRuleSuppressionSuppressionSchedule)
	RecurrenceMonthly() *[]*float64
	SetRecurrenceMonthly(val *[]*float64)
	RecurrenceMonthlyInput() *[]*float64
	RecurrenceWeekly() *[]*string
	SetRecurrenceWeekly(val *[]*string)
	RecurrenceWeeklyInput() *[]*string
	StartDateUtc() *string
	SetStartDateUtc(val *string)
	StartDateUtcInput() *string
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
	ResetRecurrenceMonthly()
	ResetRecurrenceWeekly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActionRuleSuppressionSuppressionScheduleOutputReference
type jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) EndDateUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) EndDateUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) InternalValue() *MonitorActionRuleSuppressionSuppressionSchedule {
	var returns *MonitorActionRuleSuppressionSuppressionSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) RecurrenceMonthly() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"recurrenceMonthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) RecurrenceMonthlyInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"recurrenceMonthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) RecurrenceWeekly() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recurrenceWeekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) RecurrenceWeeklyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recurrenceWeeklyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) StartDateUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) StartDateUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorActionRuleSuppressionSuppressionScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorActionRuleSuppressionSuppressionScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorActionRuleSuppressionSuppressionScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionRuleSuppression.MonitorActionRuleSuppressionSuppressionScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorActionRuleSuppressionSuppressionScheduleOutputReference_Override(m MonitorActionRuleSuppressionSuppressionScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionRuleSuppression.MonitorActionRuleSuppressionSuppressionScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetEndDateUtc(val *string) {
	if err := j.validateSetEndDateUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDateUtc",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetInternalValue(val *MonitorActionRuleSuppressionSuppressionSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetRecurrenceMonthly(val *[]*float64) {
	if err := j.validateSetRecurrenceMonthlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceMonthly",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetRecurrenceWeekly(val *[]*string) {
	if err := j.validateSetRecurrenceWeeklyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceWeekly",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetStartDateUtc(val *string) {
	if err := j.validateSetStartDateUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDateUtc",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) ResetRecurrenceMonthly() {
	_jsii_.InvokeVoid(
		m,
		"resetRecurrenceMonthly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) ResetRecurrenceWeekly() {
	_jsii_.InvokeVoid(
		m,
		"resetRecurrenceWeekly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionSuppressionScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

