// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineSqlInstanceOutputReference interface {
	cdktf.ComplexObject
	AdhocWorkloadsOptimizationEnabled() interface{}
	SetAdhocWorkloadsOptimizationEnabled(val interface{})
	AdhocWorkloadsOptimizationEnabledInput() interface{}
	Collation() *string
	SetCollation(val *string)
	CollationInput() *string
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
	InstantFileInitializationEnabled() interface{}
	SetInstantFileInitializationEnabled(val interface{})
	InstantFileInitializationEnabledInput() interface{}
	InternalValue() *MssqlVirtualMachineSqlInstance
	SetInternalValue(val *MssqlVirtualMachineSqlInstance)
	LockPagesInMemoryEnabled() interface{}
	SetLockPagesInMemoryEnabled(val interface{})
	LockPagesInMemoryEnabledInput() interface{}
	MaxDop() *float64
	SetMaxDop(val *float64)
	MaxDopInput() *float64
	MaxServerMemoryMb() *float64
	SetMaxServerMemoryMb(val *float64)
	MaxServerMemoryMbInput() *float64
	MinServerMemoryMb() *float64
	SetMinServerMemoryMb(val *float64)
	MinServerMemoryMbInput() *float64
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
	ResetAdhocWorkloadsOptimizationEnabled()
	ResetCollation()
	ResetInstantFileInitializationEnabled()
	ResetLockPagesInMemoryEnabled()
	ResetMaxDop()
	ResetMaxServerMemoryMb()
	ResetMinServerMemoryMb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineSqlInstanceOutputReference
type jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) AdhocWorkloadsOptimizationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adhocWorkloadsOptimizationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) AdhocWorkloadsOptimizationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adhocWorkloadsOptimizationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) CollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) InstantFileInitializationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instantFileInitializationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) InstantFileInitializationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instantFileInitializationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) InternalValue() *MssqlVirtualMachineSqlInstance {
	var returns *MssqlVirtualMachineSqlInstance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) LockPagesInMemoryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockPagesInMemoryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) LockPagesInMemoryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockPagesInMemoryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) MaxDop() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) MaxDopInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) MaxServerMemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxServerMemoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) MaxServerMemoryMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxServerMemoryMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) MinServerMemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minServerMemoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) MinServerMemoryMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minServerMemoryMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineSqlInstanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineSqlInstanceOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineSqlInstanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineSqlInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineSqlInstanceOutputReference_Override(m MssqlVirtualMachineSqlInstanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineSqlInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetAdhocWorkloadsOptimizationEnabled(val interface{}) {
	if err := j.validateSetAdhocWorkloadsOptimizationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adhocWorkloadsOptimizationEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetCollation(val *string) {
	if err := j.validateSetCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetInstantFileInitializationEnabled(val interface{}) {
	if err := j.validateSetInstantFileInitializationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instantFileInitializationEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetInternalValue(val *MssqlVirtualMachineSqlInstance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetLockPagesInMemoryEnabled(val interface{}) {
	if err := j.validateSetLockPagesInMemoryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockPagesInMemoryEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetMaxDop(val *float64) {
	if err := j.validateSetMaxDopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDop",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetMaxServerMemoryMb(val *float64) {
	if err := j.validateSetMaxServerMemoryMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxServerMemoryMb",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetMinServerMemoryMb(val *float64) {
	if err := j.validateSetMinServerMemoryMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minServerMemoryMb",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ResetAdhocWorkloadsOptimizationEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetAdhocWorkloadsOptimizationEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ResetCollation() {
	_jsii_.InvokeVoid(
		m,
		"resetCollation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ResetInstantFileInitializationEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetInstantFileInitializationEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ResetLockPagesInMemoryEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetLockPagesInMemoryEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ResetMaxDop() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxDop",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ResetMaxServerMemoryMb() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxServerMemoryMb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ResetMinServerMemoryMb() {
	_jsii_.InvokeVoid(
		m,
		"resetMinServerMemoryMb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineSqlInstanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

