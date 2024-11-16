// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqldatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/mssqldatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlDatabaseLongTermRetentionPolicyOutputReference interface {
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
	ImmutableBackupsEnabled() interface{}
	SetImmutableBackupsEnabled(val interface{})
	ImmutableBackupsEnabledInput() interface{}
	InternalValue() *MssqlDatabaseLongTermRetentionPolicy
	SetInternalValue(val *MssqlDatabaseLongTermRetentionPolicy)
	MonthlyRetention() *string
	SetMonthlyRetention(val *string)
	MonthlyRetentionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklyRetention() *string
	SetWeeklyRetention(val *string)
	WeeklyRetentionInput() *string
	WeekOfYear() *float64
	SetWeekOfYear(val *float64)
	WeekOfYearInput() *float64
	YearlyRetention() *string
	SetYearlyRetention(val *string)
	YearlyRetentionInput() *string
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
	ResetImmutableBackupsEnabled()
	ResetMonthlyRetention()
	ResetWeeklyRetention()
	ResetWeekOfYear()
	ResetYearlyRetention()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlDatabaseLongTermRetentionPolicyOutputReference
type jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ImmutableBackupsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"immutableBackupsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ImmutableBackupsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"immutableBackupsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) InternalValue() *MssqlDatabaseLongTermRetentionPolicy {
	var returns *MssqlDatabaseLongTermRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) MonthlyRetention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) MonthlyRetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) WeeklyRetention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) WeeklyRetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) WeekOfYear() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) WeekOfYearInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) YearlyRetention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yearlyRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) YearlyRetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yearlyRetentionInput",
		&returns,
	)
	return returns
}


func NewMssqlDatabaseLongTermRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlDatabaseLongTermRetentionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlDatabaseLongTermRetentionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseLongTermRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlDatabaseLongTermRetentionPolicyOutputReference_Override(m MssqlDatabaseLongTermRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseLongTermRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetImmutableBackupsEnabled(val interface{}) {
	if err := j.validateSetImmutableBackupsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"immutableBackupsEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetInternalValue(val *MssqlDatabaseLongTermRetentionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetMonthlyRetention(val *string) {
	if err := j.validateSetMonthlyRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthlyRetention",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetWeeklyRetention(val *string) {
	if err := j.validateSetWeeklyRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyRetention",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetWeekOfYear(val *float64) {
	if err := j.validateSetWeekOfYearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekOfYear",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference)SetYearlyRetention(val *string) {
	if err := j.validateSetYearlyRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yearlyRetention",
		val,
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetImmutableBackupsEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetImmutableBackupsEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetMonthlyRetention() {
	_jsii_.InvokeVoid(
		m,
		"resetMonthlyRetention",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetWeeklyRetention() {
	_jsii_.InvokeVoid(
		m,
		"resetWeeklyRetention",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetWeekOfYear() {
	_jsii_.InvokeVoid(
		m,
		"resetWeekOfYear",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetYearlyRetention() {
	_jsii_.InvokeVoid(
		m,
		"resetYearlyRetention",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

