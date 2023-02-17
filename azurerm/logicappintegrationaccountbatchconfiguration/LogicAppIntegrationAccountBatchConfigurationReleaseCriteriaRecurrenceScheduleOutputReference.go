package logicappintegrationaccountbatchconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/logicappintegrationaccountbatchconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference interface {
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
	Hours() *[]*float64
	SetHours(val *[]*float64)
	HoursInput() *[]*float64
	InternalValue() *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule
	SetInternalValue(val *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule)
	Minutes() *[]*float64
	SetMinutes(val *[]*float64)
	MinutesInput() *[]*float64
	MonthDays() *[]*float64
	SetMonthDays(val *[]*float64)
	MonthDaysInput() *[]*float64
	Monthly() LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyList
	MonthlyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeekDays() *[]*string
	SetWeekDays(val *[]*string)
	WeekDaysInput() *[]*string
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
	PutMonthly(value interface{})
	ResetHours()
	ResetMinutes()
	ResetMonthDays()
	ResetMonthly()
	ResetWeekDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference
type jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) Hours() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) HoursInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) InternalValue() *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule {
	var returns *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) Minutes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"minutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) MinutesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"minutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) MonthDays() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"monthDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) MonthDaysInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"monthDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) Monthly() LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyList {
	var returns LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyList
	_jsii_.Get(
		j,
		"monthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) MonthlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) WeekDays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) WeekDaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekDaysInput",
		&returns,
	)
	return returns
}


func NewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountBatchConfiguration.LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference_Override(l LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountBatchConfiguration.LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetHours(val *[]*float64) {
	if err := j.validateSetHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hours",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetInternalValue(val *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetMinutes(val *[]*float64) {
	if err := j.validateSetMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minutes",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetMonthDays(val *[]*float64) {
	if err := j.validateSetMonthDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthDays",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference)SetWeekDays(val *[]*string) {
	if err := j.validateSetWeekDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekDays",
		val,
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) PutMonthly(value interface{}) {
	if err := l.validatePutMonthlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMonthly",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ResetHours() {
	_jsii_.InvokeVoid(
		l,
		"resetHours",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ResetMinutes() {
	_jsii_.InvokeVoid(
		l,
		"resetMinutes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ResetMonthDays() {
	_jsii_.InvokeVoid(
		l,
		"resetMonthDays",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ResetMonthly() {
	_jsii_.InvokeVoid(
		l,
		"resetMonthly",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ResetWeekDays() {
	_jsii_.InvokeVoid(
		l,
		"resetWeekDays",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

