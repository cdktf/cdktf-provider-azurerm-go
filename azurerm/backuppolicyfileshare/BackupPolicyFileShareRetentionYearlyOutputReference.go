// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuppolicyfileshare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/backuppolicyfileshare/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupPolicyFileShareRetentionYearlyOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Days() *[]*float64
	SetDays(val *[]*float64)
	DaysInput() *[]*float64
	// Experimental.
	Fqn() *string
	IncludeLastDays() interface{}
	SetIncludeLastDays(val interface{})
	IncludeLastDaysInput() interface{}
	InternalValue() *BackupPolicyFileShareRetentionYearly
	SetInternalValue(val *BackupPolicyFileShareRetentionYearly)
	Months() *[]*string
	SetMonths(val *[]*string)
	MonthsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weekdays() *[]*string
	SetWeekdays(val *[]*string)
	WeekdaysInput() *[]*string
	Weeks() *[]*string
	SetWeeks(val *[]*string)
	WeeksInput() *[]*string
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
	ResetDays()
	ResetIncludeLastDays()
	ResetWeekdays()
	ResetWeeks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyFileShareRetentionYearlyOutputReference
type jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) Days() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"days",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) DaysInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) IncludeLastDays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeLastDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) IncludeLastDaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeLastDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) InternalValue() *BackupPolicyFileShareRetentionYearly {
	var returns *BackupPolicyFileShareRetentionYearly
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) Months() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) MonthsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) Weekdays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) WeekdaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) Weeks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) WeeksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksInput",
		&returns,
	)
	return returns
}


func NewBackupPolicyFileShareRetentionYearlyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyFileShareRetentionYearlyOutputReference {
	_init_.Initialize()

	if err := validateNewBackupPolicyFileShareRetentionYearlyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyFileShare.BackupPolicyFileShareRetentionYearlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyFileShareRetentionYearlyOutputReference_Override(b BackupPolicyFileShareRetentionYearlyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyFileShare.BackupPolicyFileShareRetentionYearlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetDays(val *[]*float64) {
	if err := j.validateSetDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"days",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetIncludeLastDays(val interface{}) {
	if err := j.validateSetIncludeLastDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeLastDays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetInternalValue(val *BackupPolicyFileShareRetentionYearly) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetMonths(val *[]*string) {
	if err := j.validateSetMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"months",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetWeekdays(val *[]*string) {
	if err := j.validateSetWeekdaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekdays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference)SetWeeks(val *[]*string) {
	if err := j.validateSetWeeksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeks",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) ResetDays() {
	_jsii_.InvokeVoid(
		b,
		"resetDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) ResetIncludeLastDays() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeLastDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) ResetWeekdays() {
	_jsii_.InvokeVoid(
		b,
		"resetWeekdays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) ResetWeeks() {
	_jsii_.InvokeVoid(
		b,
		"resetWeeks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyFileShareRetentionYearlyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

