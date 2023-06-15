package logicappintegrationaccountbatchconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/logicappintegrationaccountbatchconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference interface {
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
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInput() *string
	InternalValue() *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrence
	SetInternalValue(val *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrence)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	Schedule() LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference
	ScheduleInput() *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	PutSchedule(value *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule)
	ResetEndTime()
	ResetSchedule()
	ResetStartTime()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference
type jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) InternalValue() *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrence {
	var returns *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrence
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) Schedule() LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference {
	var returns LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ScheduleInput() *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule {
	var returns *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference {
	_init_.Initialize()

	if err := validateNewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountBatchConfiguration.LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference_Override(l LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountBatchConfiguration.LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetInternalValue(val *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrence) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) PutSchedule(value *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule) {
	if err := l.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSchedule",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		l,
		"resetEndTime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		l,
		"resetSchedule",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		l,
		"resetStartTime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

