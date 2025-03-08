// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualmachineinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/systemcentervirtualmachinemanagervirtualmachineinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference interface {
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
	CpuCount() *float64
	SetCpuCount(val *float64)
	CpuCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DynamicMemoryMaxInMb() *float64
	SetDynamicMemoryMaxInMb(val *float64)
	DynamicMemoryMaxInMbInput() *float64
	DynamicMemoryMinInMb() *float64
	SetDynamicMemoryMinInMb(val *float64)
	DynamicMemoryMinInMbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware
	SetInternalValue(val *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware)
	LimitCpuForMigrationEnabled() interface{}
	SetLimitCpuForMigrationEnabled(val interface{})
	LimitCpuForMigrationEnabledInput() interface{}
	MemoryInMb() *float64
	SetMemoryInMb(val *float64)
	MemoryInMbInput() *float64
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
	ResetCpuCount()
	ResetDynamicMemoryMaxInMb()
	ResetDynamicMemoryMinInMb()
	ResetLimitCpuForMigrationEnabled()
	ResetMemoryInMb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference
type jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) CpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) CpuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) DynamicMemoryMaxInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dynamicMemoryMaxInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) DynamicMemoryMaxInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dynamicMemoryMaxInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) DynamicMemoryMinInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dynamicMemoryMinInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) DynamicMemoryMinInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dynamicMemoryMinInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) InternalValue() *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware {
	var returns *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) LimitCpuForMigrationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limitCpuForMigrationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) LimitCpuForMigrationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limitCpuForMigrationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) MemoryInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) MemoryInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference {
	_init_.Initialize()

	if err := validateNewSystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference_Override(s SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetCpuCount(val *float64) {
	if err := j.validateSetCpuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCount",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetDynamicMemoryMaxInMb(val *float64) {
	if err := j.validateSetDynamicMemoryMaxInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicMemoryMaxInMb",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetDynamicMemoryMinInMb(val *float64) {
	if err := j.validateSetDynamicMemoryMinInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicMemoryMinInMb",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetInternalValue(val *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetLimitCpuForMigrationEnabled(val interface{}) {
	if err := j.validateSetLimitCpuForMigrationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitCpuForMigrationEnabled",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetMemoryInMb(val *float64) {
	if err := j.validateSetMemoryInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryInMb",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ResetCpuCount() {
	_jsii_.InvokeVoid(
		s,
		"resetCpuCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ResetDynamicMemoryMaxInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetDynamicMemoryMaxInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ResetDynamicMemoryMinInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetDynamicMemoryMinInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ResetLimitCpuForMigrationEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetLimitCpuForMigrationEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ResetMemoryInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetMemoryInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

