// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/automanageconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomanageConfigurationBackupRetentionPolicyOutputReference interface {
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
	DailySchedule() AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference
	DailyScheduleInput() *AutomanageConfigurationBackupRetentionPolicyDailySchedule
	// Experimental.
	Fqn() *string
	InternalValue() *AutomanageConfigurationBackupRetentionPolicy
	SetInternalValue(val *AutomanageConfigurationBackupRetentionPolicy)
	RetentionPolicyType() *string
	SetRetentionPolicyType(val *string)
	RetentionPolicyTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklySchedule() AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference
	WeeklyScheduleInput() *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule
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
	PutDailySchedule(value *AutomanageConfigurationBackupRetentionPolicyDailySchedule)
	PutWeeklySchedule(value *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule)
	ResetDailySchedule()
	ResetRetentionPolicyType()
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomanageConfigurationBackupRetentionPolicyOutputReference
type jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) DailySchedule() AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference {
	var returns AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) DailyScheduleInput() *AutomanageConfigurationBackupRetentionPolicyDailySchedule {
	var returns *AutomanageConfigurationBackupRetentionPolicyDailySchedule
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) InternalValue() *AutomanageConfigurationBackupRetentionPolicy {
	var returns *AutomanageConfigurationBackupRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) RetentionPolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionPolicyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) RetentionPolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionPolicyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) WeeklySchedule() AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference {
	var returns AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) WeeklyScheduleInput() *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule {
	var returns *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewAutomanageConfigurationBackupRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomanageConfigurationBackupRetentionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAutomanageConfigurationBackupRetentionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomanageConfigurationBackupRetentionPolicyOutputReference_Override(a AutomanageConfigurationBackupRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference)SetInternalValue(val *AutomanageConfigurationBackupRetentionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference)SetRetentionPolicyType(val *string) {
	if err := j.validateSetRetentionPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPolicyType",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) PutDailySchedule(value *AutomanageConfigurationBackupRetentionPolicyDailySchedule) {
	if err := a.validatePutDailyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDailySchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) PutWeeklySchedule(value *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule) {
	if err := a.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) ResetRetentionPolicyType() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionPolicyType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

