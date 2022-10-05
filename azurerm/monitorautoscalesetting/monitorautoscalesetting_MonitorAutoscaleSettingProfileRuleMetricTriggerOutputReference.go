package monitorautoscalesetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/monitorautoscalesetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference interface {
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
	Dimensions() MonitorAutoscaleSettingProfileRuleMetricTriggerDimensionsList
	DimensionsInput() interface{}
	DivideByInstanceCount() interface{}
	SetDivideByInstanceCount(val interface{})
	DivideByInstanceCountInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAutoscaleSettingProfileRuleMetricTrigger
	SetInternalValue(val *MonitorAutoscaleSettingProfileRuleMetricTrigger)
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	MetricNamespace() *string
	SetMetricNamespace(val *string)
	MetricNamespaceInput() *string
	MetricResourceId() *string
	SetMetricResourceId(val *string)
	MetricResourceIdInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Statistic() *string
	SetStatistic(val *string)
	StatisticInput() *string
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
	TimeAggregation() *string
	SetTimeAggregation(val *string)
	TimeAggregationInput() *string
	TimeGrain() *string
	SetTimeGrain(val *string)
	TimeGrainInput() *string
	TimeWindow() *string
	SetTimeWindow(val *string)
	TimeWindowInput() *string
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
	PutDimensions(value interface{})
	ResetDimensions()
	ResetDivideByInstanceCount()
	ResetMetricNamespace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference
type jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) Dimensions() MonitorAutoscaleSettingProfileRuleMetricTriggerDimensionsList {
	var returns MonitorAutoscaleSettingProfileRuleMetricTriggerDimensionsList
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) DimensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) DivideByInstanceCount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"divideByInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) DivideByInstanceCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"divideByInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) InternalValue() *MonitorAutoscaleSettingProfileRuleMetricTrigger {
	var returns *MonitorAutoscaleSettingProfileRuleMetricTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) MetricNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) MetricResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) MetricResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) TimeAggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeAggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) TimeAggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeAggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) TimeGrain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGrain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) TimeGrainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGrainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) TimeWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) TimeWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}


func NewMonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorAutoscaleSettingProfileRuleMetricTriggerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAutoscaleSetting.MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference_Override(m MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAutoscaleSetting.MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetDivideByInstanceCount(val interface{}) {
	if err := j.validateSetDivideByInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"divideByInstanceCount",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetInternalValue(val *MonitorAutoscaleSettingProfileRuleMetricTrigger) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetMetricNamespace(val *string) {
	if err := j.validateSetMetricNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricNamespace",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetMetricResourceId(val *string) {
	if err := j.validateSetMetricResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricResourceId",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetTimeAggregation(val *string) {
	if err := j.validateSetTimeAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeAggregation",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetTimeGrain(val *string) {
	if err := j.validateSetTimeGrainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeGrain",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference)SetTimeWindow(val *string) {
	if err := j.validateSetTimeWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindow",
		val,
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) PutDimensions(value interface{}) {
	if err := m.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDimensions",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		m,
		"resetDimensions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) ResetDivideByInstanceCount() {
	_jsii_.InvokeVoid(
		m,
		"resetDivideByInstanceCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) ResetMetricNamespace() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricNamespace",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorAutoscaleSettingProfileRuleMetricTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

