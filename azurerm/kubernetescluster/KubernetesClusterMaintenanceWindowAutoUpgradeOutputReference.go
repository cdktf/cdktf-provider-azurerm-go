// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference interface {
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
	InternalValue() *KubernetesClusterMaintenanceWindowAutoUpgrade
	SetInternalValue(val *KubernetesClusterMaintenanceWindowAutoUpgrade)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	NotAllowed() KubernetesClusterMaintenanceWindowAutoUpgradeNotAllowedList
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

// The jsii proxy struct for KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference
type jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) DayOfMonth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) DayOfMonthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) DayOfWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) Duration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) DurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) InternalValue() *KubernetesClusterMaintenanceWindowAutoUpgrade {
	var returns *KubernetesClusterMaintenanceWindowAutoUpgrade
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) NotAllowed() KubernetesClusterMaintenanceWindowAutoUpgradeNotAllowedList {
	var returns KubernetesClusterMaintenanceWindowAutoUpgradeNotAllowedList
	_jsii_.Get(
		j,
		"notAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) NotAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) StartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) StartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) UtcOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) UtcOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"utcOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) WeekIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) WeekIndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekIndexInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMaintenanceWindowAutoUpgradeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterMaintenanceWindowAutoUpgradeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterMaintenanceWindowAutoUpgradeOutputReference_Override(k KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetDayOfMonth(val *float64) {
	if err := j.validateSetDayOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfMonth",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetDayOfWeek(val *string) {
	if err := j.validateSetDayOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetDuration(val *float64) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetInternalValue(val *KubernetesClusterMaintenanceWindowAutoUpgrade) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetStartDate(val *string) {
	if err := j.validateSetStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDate",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetUtcOffset(val *string) {
	if err := j.validateSetUtcOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"utcOffset",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference)SetWeekIndex(val *string) {
	if err := j.validateSetWeekIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekIndex",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) PutNotAllowed(value interface{}) {
	if err := k.validatePutNotAllowedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNotAllowed",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ResetDayOfMonth() {
	_jsii_.InvokeVoid(
		k,
		"resetDayOfMonth",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ResetDayOfWeek() {
	_jsii_.InvokeVoid(
		k,
		"resetDayOfWeek",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ResetNotAllowed() {
	_jsii_.InvokeVoid(
		k,
		"resetNotAllowed",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ResetStartDate() {
	_jsii_.InvokeVoid(
		k,
		"resetStartDate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		k,
		"resetStartTime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ResetUtcOffset() {
	_jsii_.InvokeVoid(
		k,
		"resetUtcOffset",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ResetWeekIndex() {
	_jsii_.InvokeVoid(
		k,
		"resetWeekIndex",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAutoUpgradeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

