// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitormetricalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/monitormetricalert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorMetricAlertDynamicCriteriaOutputReference interface {
	cdktf.ComplexObject
	Aggregation() *string
	SetAggregation(val *string)
	AggregationInput() *string
	AlertSensitivity() *string
	SetAlertSensitivity(val *string)
	AlertSensitivityInput() *string
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
	Dimension() MonitorMetricAlertDynamicCriteriaDimensionList
	DimensionInput() interface{}
	EvaluationFailureCount() *float64
	SetEvaluationFailureCount(val *float64)
	EvaluationFailureCountInput() *float64
	EvaluationTotalCount() *float64
	SetEvaluationTotalCount(val *float64)
	EvaluationTotalCountInput() *float64
	// Experimental.
	Fqn() *string
	IgnoreDataBefore() *string
	SetIgnoreDataBefore(val *string)
	IgnoreDataBeforeInput() *string
	InternalValue() *MonitorMetricAlertDynamicCriteria
	SetInternalValue(val *MonitorMetricAlertDynamicCriteria)
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	MetricNamespace() *string
	SetMetricNamespace(val *string)
	MetricNamespaceInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	SkipMetricValidation() interface{}
	SetSkipMetricValidation(val interface{})
	SkipMetricValidationInput() interface{}
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
	PutDimension(value interface{})
	ResetDimension()
	ResetEvaluationFailureCount()
	ResetEvaluationTotalCount()
	ResetIgnoreDataBefore()
	ResetSkipMetricValidation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorMetricAlertDynamicCriteriaOutputReference
type jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) AlertSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) AlertSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) Dimension() MonitorMetricAlertDynamicCriteriaDimensionList {
	var returns MonitorMetricAlertDynamicCriteriaDimensionList
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) DimensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) EvaluationFailureCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationFailureCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) EvaluationFailureCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationFailureCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) EvaluationTotalCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationTotalCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) EvaluationTotalCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationTotalCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) IgnoreDataBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreDataBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) IgnoreDataBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreDataBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) InternalValue() *MonitorMetricAlertDynamicCriteria {
	var returns *MonitorMetricAlertDynamicCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) MetricNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) SkipMetricValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMetricValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) SkipMetricValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMetricValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorMetricAlertDynamicCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorMetricAlertDynamicCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorMetricAlertDynamicCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlertDynamicCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorMetricAlertDynamicCriteriaOutputReference_Override(m MonitorMetricAlertDynamicCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlertDynamicCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetAlertSensitivity(val *string) {
	if err := j.validateSetAlertSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertSensitivity",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetEvaluationFailureCount(val *float64) {
	if err := j.validateSetEvaluationFailureCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationFailureCount",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetEvaluationTotalCount(val *float64) {
	if err := j.validateSetEvaluationTotalCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationTotalCount",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetIgnoreDataBefore(val *string) {
	if err := j.validateSetIgnoreDataBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreDataBefore",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetInternalValue(val *MonitorMetricAlertDynamicCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetMetricNamespace(val *string) {
	if err := j.validateSetMetricNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricNamespace",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetSkipMetricValidation(val interface{}) {
	if err := j.validateSetSkipMetricValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipMetricValidation",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) PutDimension(value interface{}) {
	if err := m.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDimension",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		m,
		"resetDimension",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ResetEvaluationFailureCount() {
	_jsii_.InvokeVoid(
		m,
		"resetEvaluationFailureCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ResetEvaluationTotalCount() {
	_jsii_.InvokeVoid(
		m,
		"resetEvaluationTotalCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ResetIgnoreDataBefore() {
	_jsii_.InvokeVoid(
		m,
		"resetIgnoreDataBefore",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ResetSkipMetricValidation() {
	_jsii_.InvokeVoid(
		m,
		"resetSkipMetricValidation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorMetricAlertDynamicCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

