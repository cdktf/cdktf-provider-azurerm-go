// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicymysqlflexibleserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dataprotectionbackuppolicymysqlflexibleserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference interface {
	cdktf.ComplexObject
	AbsoluteCriteria() *string
	SetAbsoluteCriteria(val *string)
	AbsoluteCriteriaInput() *string
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
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteria
	SetInternalValue(val *DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteria)
	MonthsOfYear() *[]*string
	SetMonthsOfYear(val *[]*string)
	MonthsOfYearInput() *[]*string
	ScheduledBackupTimes() *[]*string
	SetScheduledBackupTimes(val *[]*string)
	ScheduledBackupTimesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeksOfMonth() *[]*string
	SetWeeksOfMonth(val *[]*string)
	WeeksOfMonthInput() *[]*string
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
	ResetAbsoluteCriteria()
	ResetDaysOfWeek()
	ResetMonthsOfYear()
	ResetScheduledBackupTimes()
	ResetWeeksOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference
type jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) AbsoluteCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) AbsoluteCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) InternalValue() *DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteria {
	var returns *DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) MonthsOfYear() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsOfYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) MonthsOfYearInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsOfYearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ScheduledBackupTimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scheduledBackupTimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ScheduledBackupTimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scheduledBackupTimesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) WeeksOfMonth() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) WeeksOfMonthInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


func NewDataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewDataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyMysqlFlexibleServer.DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference_Override(d DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyMysqlFlexibleServer.DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetAbsoluteCriteria(val *string) {
	if err := j.validateSetAbsoluteCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"absoluteCriteria",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetInternalValue(val *DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetMonthsOfYear(val *[]*string) {
	if err := j.validateSetMonthsOfYearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthsOfYear",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetScheduledBackupTimes(val *[]*string) {
	if err := j.validateSetScheduledBackupTimesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledBackupTimes",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference)SetWeeksOfMonth(val *[]*string) {
	if err := j.validateSetWeeksOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ResetAbsoluteCriteria() {
	_jsii_.InvokeVoid(
		d,
		"resetAbsoluteCriteria",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		d,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ResetMonthsOfYear() {
	_jsii_.InvokeVoid(
		d,
		"resetMonthsOfYear",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ResetScheduledBackupTimes() {
	_jsii_.InvokeVoid(
		d,
		"resetScheduledBackupTimes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		d,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

