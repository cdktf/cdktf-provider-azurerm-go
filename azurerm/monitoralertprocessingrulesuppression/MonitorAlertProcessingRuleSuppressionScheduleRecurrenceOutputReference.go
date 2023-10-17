// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingrulesuppression

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/monitoralertprocessingrulesuppression/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference interface {
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
	Daily() MonitorAlertProcessingRuleSuppressionScheduleRecurrenceDailyList
	DailyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleSuppressionScheduleRecurrence
	SetInternalValue(val *MonitorAlertProcessingRuleSuppressionScheduleRecurrence)
	Monthly() MonitorAlertProcessingRuleSuppressionScheduleRecurrenceMonthlyList
	MonthlyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weekly() MonitorAlertProcessingRuleSuppressionScheduleRecurrenceWeeklyList
	WeeklyInput() interface{}
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
	PutDaily(value interface{})
	PutMonthly(value interface{})
	PutWeekly(value interface{})
	ResetDaily()
	ResetMonthly()
	ResetWeekly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference
type jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) Daily() MonitorAlertProcessingRuleSuppressionScheduleRecurrenceDailyList {
	var returns MonitorAlertProcessingRuleSuppressionScheduleRecurrenceDailyList
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) DailyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) InternalValue() *MonitorAlertProcessingRuleSuppressionScheduleRecurrence {
	var returns *MonitorAlertProcessingRuleSuppressionScheduleRecurrence
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) Monthly() MonitorAlertProcessingRuleSuppressionScheduleRecurrenceMonthlyList {
	var returns MonitorAlertProcessingRuleSuppressionScheduleRecurrenceMonthlyList
	_jsii_.Get(
		j,
		"monthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) MonthlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) Weekly() MonitorAlertProcessingRuleSuppressionScheduleRecurrenceWeeklyList {
	var returns MonitorAlertProcessingRuleSuppressionScheduleRecurrenceWeeklyList
	_jsii_.Get(
		j,
		"weekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) WeeklyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleSuppression.MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference_Override(m MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleSuppression.MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference)SetInternalValue(val *MonitorAlertProcessingRuleSuppressionScheduleRecurrence) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) PutDaily(value interface{}) {
	if err := m.validatePutDailyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDaily",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) PutMonthly(value interface{}) {
	if err := m.validatePutMonthlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonthly",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) PutWeekly(value interface{}) {
	if err := m.validatePutWeeklyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWeekly",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		m,
		"resetDaily",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) ResetMonthly() {
	_jsii_.InvokeVoid(
		m,
		"resetMonthly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) ResetWeekly() {
	_jsii_.InvokeVoid(
		m,
		"resetWeekly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionScheduleRecurrenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

