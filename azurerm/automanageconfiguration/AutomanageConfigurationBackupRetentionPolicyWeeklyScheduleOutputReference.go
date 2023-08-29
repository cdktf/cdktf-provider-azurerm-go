// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/automanageconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference interface {
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
	InternalValue() *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule
	SetInternalValue(val *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule)
	RetentionDuration() AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference
	RetentionDurationInput() *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration
	RetentionTimes() *[]*string
	SetRetentionTimes(val *[]*string)
	RetentionTimesInput() *[]*string
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
	PutRetentionDuration(value *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration)
	ResetRetentionDuration()
	ResetRetentionTimes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference
type jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) InternalValue() *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule {
	var returns *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) RetentionDuration() AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference {
	var returns AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDurationOutputReference
	_jsii_.Get(
		j,
		"retentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) RetentionDurationInput() *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration {
	var returns *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration
	_jsii_.Get(
		j,
		"retentionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) RetentionTimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retentionTimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) RetentionTimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retentionTimesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewAutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference_Override(a AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference)SetInternalValue(val *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference)SetRetentionTimes(val *[]*string) {
	if err := j.validateSetRetentionTimesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionTimes",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) PutRetentionDuration(value *AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration) {
	if err := a.validatePutRetentionDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetentionDuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) ResetRetentionDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) ResetRetentionTimes() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionTimes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

