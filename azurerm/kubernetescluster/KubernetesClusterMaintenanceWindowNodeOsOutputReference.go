// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterMaintenanceWindowNodeOsOutputReference interface {
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
	DayOfMonth() *float64
	SetDayOfMonth(val *float64)
	DayOfMonthInput() *float64
	DayOfWeek() *string
	SetDayOfWeek(val *string)
	DayOfWeekInput() *string
	Duration() *float64
	SetDuration(val *float64)
	DurationInput() *float64
	// Experimental.
	Fqn() *string
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInput() *string
	InternalValue() *KubernetesClusterMaintenanceWindowNodeOs
	SetInternalValue(val *KubernetesClusterMaintenanceWindowNodeOs)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	NotAllowed() KubernetesClusterMaintenanceWindowNodeOsNotAllowedList
	NotAllowedInput() interface{}
	StartDate() *string
	SetStartDate(val *string)
	StartDateInput() *string
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
	UtcOffset() *string
	SetUtcOffset(val *string)
	UtcOffsetInput() *string
	WeekIndex() *string
	SetWeekIndex(val *string)
	WeekIndexInput() *string
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
	PutNotAllowed(value interface{})
	ResetDayOfMonth()
	ResetDayOfWeek()
	ResetNotAllowed()
	ResetStartDate()
	ResetStartTime()
	ResetUtcOffset()
	ResetWeekIndex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterMaintenanceWindowNodeOsOutputReference
type jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) DayOfMonth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) DayOfMonthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) DayOfWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) Duration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) DurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) InternalValue() *KubernetesClusterMaintenanceWindowNodeOs {
	var returns *KubernetesClusterMaintenanceWindowNodeOs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) NotAllowed() KubernetesClusterMaintenanceWindowNodeOsNotAllowedList {
	var returns KubernetesClusterMaintenanceWindowNodeOsNotAllowedList
	_jsii_.Get(
		j,
		"notAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) NotAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) StartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) StartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) UtcOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) UtcOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) WeekIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) WeekIndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekIndexInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMaintenanceWindowNodeOsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterMaintenanceWindowNodeOsOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterMaintenanceWindowNodeOsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowNodeOsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterMaintenanceWindowNodeOsOutputReference_Override(k KubernetesClusterMaintenanceWindowNodeOsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowNodeOsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetDayOfMonth(val *float64) {
	if err := j.validateSetDayOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfMonth",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetDayOfWeek(val *string) {
	if err := j.validateSetDayOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetDuration(val *float64) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetInternalValue(val *KubernetesClusterMaintenanceWindowNodeOs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetStartDate(val *string) {
	if err := j.validateSetStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDate",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetUtcOffset(val *string) {
	if err := j.validateSetUtcOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utcOffset",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference)SetWeekIndex(val *string) {
	if err := j.validateSetWeekIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekIndex",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) PutNotAllowed(value interface{}) {
	if err := k.validatePutNotAllowedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNotAllowed",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ResetDayOfMonth() {
	_jsii_.InvokeVoid(
		k,
		"resetDayOfMonth",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ResetDayOfWeek() {
	_jsii_.InvokeVoid(
		k,
		"resetDayOfWeek",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ResetNotAllowed() {
	_jsii_.InvokeVoid(
		k,
		"resetNotAllowed",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ResetStartDate() {
	_jsii_.InvokeVoid(
		k,
		"resetStartDate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		k,
		"resetStartTime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ResetUtcOffset() {
	_jsii_.InvokeVoid(
		k,
		"resetUtcOffset",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ResetWeekIndex() {
	_jsii_.InvokeVoid(
		k,
		"resetWeekIndex",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNodeOsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

