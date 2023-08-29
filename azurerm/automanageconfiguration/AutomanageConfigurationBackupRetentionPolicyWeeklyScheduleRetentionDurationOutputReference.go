// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/automanageconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference interface {
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
	DurationType() *string
	SetDurationType(val *string)
	DurationTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration
	SetInternalValue(val *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration)
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
	ResetCount()
	ResetDurationType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference
type jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) DurationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) DurationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) InternalValue() *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration {
	var returns *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference {
	_init_.Initialize()

	if err := validateNewAutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference_Override(a AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference)SetDurationType(val *string) {
	if err := j.validateSetDurationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationType",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference)SetInternalValue(val *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		a,
		"resetCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) ResetDurationType() {
	_jsii_.InvokeVoid(
		a,
		"resetDurationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

