package automationsoftwareupdateconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/automationsoftwareupdateconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationSoftwareUpdateConfigurationScheduleOutputReference interface {
	cdktf.ComplexObject
	AdvancedMonthDays() *[]*float64
	SetAdvancedMonthDays(val *[]*float64)
	AdvancedMonthDaysInput() *[]*float64
	AdvancedWeekDays() *[]*string
	SetAdvancedWeekDays(val *[]*string)
	AdvancedWeekDaysInput() *[]*string
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
	CreationTime() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExpiryTime() *string
	SetExpiryTime(val *string)
	ExpiryTimeInput() *string
	ExpiryTimeOffsetMinutes() *float64
	SetExpiryTimeOffsetMinutes(val *float64)
	ExpiryTimeOffsetMinutesInput() *float64
	// Experimental.
	Fqn() *string
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	LastModifiedTime() *string
	MonthlyOccurrence() AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList
	MonthlyOccurrenceInput() interface{}
	NextRun() *string
	SetNextRun(val *string)
	NextRunInput() *string
	NextRunOffsetMinutes() *float64
	SetNextRunOffsetMinutes(val *float64)
	NextRunOffsetMinutesInput() *float64
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	StartTimeOffsetMinutes() *float64
	SetStartTimeOffsetMinutes(val *float64)
	StartTimeOffsetMinutesInput() *float64
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
	PutMonthlyOccurrence(value interface{})
	ResetAdvancedMonthDays()
	ResetAdvancedWeekDays()
	ResetDescription()
	ResetExpiryTime()
	ResetExpiryTimeOffsetMinutes()
	ResetFrequency()
	ResetInterval()
	ResetIsEnabled()
	ResetMonthlyOccurrence()
	ResetNextRun()
	ResetNextRunOffsetMinutes()
	ResetStartTime()
	ResetStartTimeOffsetMinutes()
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

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationScheduleOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) AdvancedMonthDays() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"advancedMonthDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) AdvancedMonthDaysInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"advancedMonthDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) AdvancedWeekDays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advancedWeekDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) AdvancedWeekDaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advancedWeekDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ExpiryTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ExpiryTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ExpiryTimeOffsetMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiryTimeOffsetMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ExpiryTimeOffsetMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiryTimeOffsetMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) MonthlyOccurrence() AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList {
	var returns AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList
	_jsii_.Get(
		j,
		"monthlyOccurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) MonthlyOccurrenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyOccurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) NextRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) NextRunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) NextRunOffsetMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nextRunOffsetMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) NextRunOffsetMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nextRunOffsetMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) StartTimeOffsetMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeOffsetMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) StartTimeOffsetMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeOffsetMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewAutomationSoftwareUpdateConfigurationScheduleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationScheduleOutputReference_Override(a AutomationSoftwareUpdateConfigurationScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetAdvancedMonthDays(val *[]*float64) {
	if err := j.validateSetAdvancedMonthDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedMonthDays",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetAdvancedWeekDays(val *[]*string) {
	if err := j.validateSetAdvancedWeekDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedWeekDays",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetExpiryTime(val *string) {
	if err := j.validateSetExpiryTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiryTime",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetExpiryTimeOffsetMinutes(val *float64) {
	if err := j.validateSetExpiryTimeOffsetMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiryTimeOffsetMinutes",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetNextRun(val *string) {
	if err := j.validateSetNextRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextRun",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetNextRunOffsetMinutes(val *float64) {
	if err := j.validateSetNextRunOffsetMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextRunOffsetMinutes",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetStartTimeOffsetMinutes(val *float64) {
	if err := j.validateSetStartTimeOffsetMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeOffsetMinutes",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) PutMonthlyOccurrence(value interface{}) {
	if err := a.validatePutMonthlyOccurrenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonthlyOccurrence",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetAdvancedMonthDays() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedMonthDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetAdvancedWeekDays() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedWeekDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetExpiryTime() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiryTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetExpiryTimeOffsetMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiryTimeOffsetMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetFrequency() {
	_jsii_.InvokeVoid(
		a,
		"resetFrequency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetMonthlyOccurrence() {
	_jsii_.InvokeVoid(
		a,
		"resetMonthlyOccurrence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetNextRun() {
	_jsii_.InvokeVoid(
		a,
		"resetNextRun",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetNextRunOffsetMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetNextRunOffsetMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetStartTimeOffsetMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTimeOffsetMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

