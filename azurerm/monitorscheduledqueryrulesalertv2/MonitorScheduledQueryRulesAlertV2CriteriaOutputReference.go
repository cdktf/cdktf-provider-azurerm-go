// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorscheduledqueryrulesalertv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/monitorscheduledqueryrulesalertv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorScheduledQueryRulesAlertV2CriteriaOutputReference interface {
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
	Dimension() MonitorScheduledQueryRulesAlertV2CriteriaDimensionList
	DimensionInput() interface{}
	FailingPeriods() MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference
	FailingPeriodsInput() *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricMeasureColumn() *string
	SetMetricMeasureColumn(val *string)
	MetricMeasureColumnInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	ResourceIdColumn() *string
	SetResourceIdColumn(val *string)
	ResourceIdColumnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
	TimeAggregationMethod() *string
	SetTimeAggregationMethod(val *string)
	TimeAggregationMethodInput() *string
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
	PutDimension(value interface{})
	PutFailingPeriods(value *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods)
	ResetDimension()
	ResetFailingPeriods()
	ResetMetricMeasureColumn()
	ResetResourceIdColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2CriteriaOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Dimension() MonitorScheduledQueryRulesAlertV2CriteriaDimensionList {
	var returns MonitorScheduledQueryRulesAlertV2CriteriaDimensionList
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) DimensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) FailingPeriods() MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference {
	var returns MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference
	_jsii_.Get(
		j,
		"failingPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) FailingPeriodsInput() *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods {
	var returns *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods
	_jsii_.Get(
		j,
		"failingPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) MetricMeasureColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricMeasureColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) MetricMeasureColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricMeasureColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResourceIdColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResourceIdColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) TimeAggregationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeAggregationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) TimeAggregationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeAggregationMethodInput",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertV2CriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorScheduledQueryRulesAlertV2CriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorScheduledQueryRulesAlertV2CriteriaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertV2CriteriaOutputReference_Override(m MonitorScheduledQueryRulesAlertV2CriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetMetricMeasureColumn(val *string) {
	if err := j.validateSetMetricMeasureColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricMeasureColumn",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetResourceIdColumn(val *string) {
	if err := j.validateSetResourceIdColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdColumn",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference)SetTimeAggregationMethod(val *string) {
	if err := j.validateSetTimeAggregationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeAggregationMethod",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) PutDimension(value interface{}) {
	if err := m.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDimension",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) PutFailingPeriods(value *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods) {
	if err := m.validatePutFailingPeriodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFailingPeriods",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		m,
		"resetDimension",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResetFailingPeriods() {
	_jsii_.InvokeVoid(
		m,
		"resetFailingPeriods",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResetMetricMeasureColumn() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricMeasureColumn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResetResourceIdColumn() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceIdColumn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

