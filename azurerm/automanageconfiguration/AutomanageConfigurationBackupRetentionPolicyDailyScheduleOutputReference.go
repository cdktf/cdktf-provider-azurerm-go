// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/automanageconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference interface {
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
	InternalValue() *AutomanageConfigurationBackupRetentionPolicyDailySchedule
	SetInternalValue(val *AutomanageConfigurationBackupRetentionPolicyDailySchedule)
	RetentionDuration() AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDurationOutputReference
	RetentionDurationInput() *AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDuration
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutRetentionDuration(value *AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDuration)
	ResetRetentionDuration()
	ResetRetentionTimes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference
type jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) InternalValue() *AutomanageConfigurationBackupRetentionPolicyDailySchedule {
	var returns *AutomanageConfigurationBackupRetentionPolicyDailySchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) RetentionDuration() AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDurationOutputReference {
	var returns AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDurationOutputReference
	_jsii_.Get(
		j,
		"retentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) RetentionDurationInput() *AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDuration {
	var returns *AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDuration
	_jsii_.Get(
		j,
		"retentionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) RetentionTimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retentionTimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) RetentionTimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retentionTimesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewAutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference_Override(a AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference)SetInternalValue(val *AutomanageConfigurationBackupRetentionPolicyDailySchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference)SetRetentionTimes(val *[]*string) {
	if err := j.validateSetRetentionTimesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionTimes",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) PutRetentionDuration(value *AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDuration) {
	if err := a.validatePutRetentionDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetentionDuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) ResetRetentionDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) ResetRetentionTimes() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionTimes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupRetentionPolicyDailyScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

