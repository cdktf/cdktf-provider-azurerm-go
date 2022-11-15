package windowsfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/windowsfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsFunctionAppBackupScheduleOutputReference interface {
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
	FrequencyInterval() *float64
	SetFrequencyInterval(val *float64)
	FrequencyIntervalInput() *float64
	FrequencyUnit() *string
	SetFrequencyUnit(val *string)
	FrequencyUnitInput() *string
	InternalValue() *WindowsFunctionAppBackupSchedule
	SetInternalValue(val *WindowsFunctionAppBackupSchedule)
	KeepAtLeastOneBackup() interface{}
	SetKeepAtLeastOneBackup(val interface{})
	KeepAtLeastOneBackupInput() interface{}
	LastExecutionTime() *string
	RetentionPeriodDays() *float64
	SetRetentionPeriodDays(val *float64)
	RetentionPeriodDaysInput() *float64
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
	ResetKeepAtLeastOneBackup()
	ResetRetentionPeriodDays()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsFunctionAppBackupScheduleOutputReference
type jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) FrequencyInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) FrequencyIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) FrequencyUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) FrequencyUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) InternalValue() *WindowsFunctionAppBackupSchedule {
	var returns *WindowsFunctionAppBackupSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) KeepAtLeastOneBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepAtLeastOneBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) KeepAtLeastOneBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepAtLeastOneBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) LastExecutionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastExecutionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) RetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) RetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsFunctionAppBackupScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsFunctionAppBackupScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsFunctionAppBackupScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsFunctionApp.WindowsFunctionAppBackupScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsFunctionAppBackupScheduleOutputReference_Override(w WindowsFunctionAppBackupScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsFunctionApp.WindowsFunctionAppBackupScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetFrequencyInterval(val *float64) {
	if err := j.validateSetFrequencyIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequencyInterval",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetFrequencyUnit(val *string) {
	if err := j.validateSetFrequencyUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequencyUnit",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetInternalValue(val *WindowsFunctionAppBackupSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetKeepAtLeastOneBackup(val interface{}) {
	if err := j.validateSetKeepAtLeastOneBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAtLeastOneBackup",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetRetentionPeriodDays(val *float64) {
	if err := j.validateSetRetentionPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriodDays",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) ResetKeepAtLeastOneBackup() {
	_jsii_.InvokeVoid(
		w,
		"resetKeepAtLeastOneBackup",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) ResetRetentionPeriodDays() {
	_jsii_.InvokeVoid(
		w,
		"resetRetentionPeriodDays",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		w,
		"resetStartTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppBackupScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

