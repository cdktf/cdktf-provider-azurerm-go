package logicapptriggerrecurrence

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/logicapptriggerrecurrence/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppTriggerRecurrenceScheduleOutputReference interface {
	cdktf.ComplexObject
	AtTheseHours() *[]*float64
	SetAtTheseHours(val *[]*float64)
	AtTheseHoursInput() *[]*float64
	AtTheseMinutes() *[]*float64
	SetAtTheseMinutes(val *[]*float64)
	AtTheseMinutesInput() *[]*float64
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
	InternalValue() *LogicAppTriggerRecurrenceSchedule
	SetInternalValue(val *LogicAppTriggerRecurrenceSchedule)
	OnTheseDays() *[]*string
	SetOnTheseDays(val *[]*string)
	OnTheseDaysInput() *[]*string
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
	ResetAtTheseHours()
	ResetAtTheseMinutes()
	ResetOnTheseDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogicAppTriggerRecurrenceScheduleOutputReference
type jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) AtTheseHours() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"atTheseHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) AtTheseHoursInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"atTheseHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) AtTheseMinutes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"atTheseMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) AtTheseMinutesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"atTheseMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) InternalValue() *LogicAppTriggerRecurrenceSchedule {
	var returns *LogicAppTriggerRecurrenceSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) OnTheseDays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onTheseDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) OnTheseDaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onTheseDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLogicAppTriggerRecurrenceScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogicAppTriggerRecurrenceScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewLogicAppTriggerRecurrenceScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppTriggerRecurrence.LogicAppTriggerRecurrenceScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogicAppTriggerRecurrenceScheduleOutputReference_Override(l LogicAppTriggerRecurrenceScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppTriggerRecurrence.LogicAppTriggerRecurrenceScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference)SetAtTheseHours(val *[]*float64) {
	if err := j.validateSetAtTheseHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atTheseHours",
		val,
	)
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference)SetAtTheseMinutes(val *[]*float64) {
	if err := j.validateSetAtTheseMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atTheseMinutes",
		val,
	)
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference)SetInternalValue(val *LogicAppTriggerRecurrenceSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference)SetOnTheseDays(val *[]*string) {
	if err := j.validateSetOnTheseDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onTheseDays",
		val,
	)
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) ResetAtTheseHours() {
	_jsii_.InvokeVoid(
		l,
		"resetAtTheseHours",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) ResetAtTheseMinutes() {
	_jsii_.InvokeVoid(
		l,
		"resetAtTheseMinutes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) ResetOnTheseDays() {
	_jsii_.InvokeVoid(
		l,
		"resetOnTheseDays",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LogicAppTriggerRecurrenceScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

