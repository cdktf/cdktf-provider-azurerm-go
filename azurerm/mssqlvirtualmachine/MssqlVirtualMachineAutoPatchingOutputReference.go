// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineAutoPatchingOutputReference interface {
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
	DayOfWeek() *string
	SetDayOfWeek(val *string)
	DayOfWeekInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineAutoPatching
	SetInternalValue(val *MssqlVirtualMachineAutoPatching)
	MaintenanceWindowDurationInMinutes() *float64
	SetMaintenanceWindowDurationInMinutes(val *float64)
	MaintenanceWindowDurationInMinutesInput() *float64
	MaintenanceWindowStartingHour() *float64
	SetMaintenanceWindowStartingHour(val *float64)
	MaintenanceWindowStartingHourInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineAutoPatchingOutputReference
type jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) DayOfWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) InternalValue() *MssqlVirtualMachineAutoPatching {
	var returns *MssqlVirtualMachineAutoPatching
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) MaintenanceWindowDurationInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowDurationInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) MaintenanceWindowDurationInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowDurationInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) MaintenanceWindowStartingHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowStartingHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) MaintenanceWindowStartingHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowStartingHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineAutoPatchingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineAutoPatchingOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineAutoPatchingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoPatchingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineAutoPatchingOutputReference_Override(m MssqlVirtualMachineAutoPatchingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoPatchingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference)SetDayOfWeek(val *string) {
	if err := j.validateSetDayOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference)SetInternalValue(val *MssqlVirtualMachineAutoPatching) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference)SetMaintenanceWindowDurationInMinutes(val *float64) {
	if err := j.validateSetMaintenanceWindowDurationInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowDurationInMinutes",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference)SetMaintenanceWindowStartingHour(val *float64) {
	if err := j.validateSetMaintenanceWindowStartingHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowStartingHour",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

