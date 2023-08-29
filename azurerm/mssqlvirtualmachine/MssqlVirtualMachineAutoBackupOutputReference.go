// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineAutoBackupOutputReference interface {
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
	EncryptionEnabled() interface{}
	SetEncryptionEnabled(val interface{})
	EncryptionEnabledInput() interface{}
	EncryptionPassword() *string
	SetEncryptionPassword(val *string)
	EncryptionPasswordInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineAutoBackup
	SetInternalValue(val *MssqlVirtualMachineAutoBackup)
	ManualSchedule() MssqlVirtualMachineAutoBackupManualScheduleOutputReference
	ManualScheduleInput() *MssqlVirtualMachineAutoBackupManualSchedule
	RetentionPeriodInDays() *float64
	SetRetentionPeriodInDays(val *float64)
	RetentionPeriodInDaysInput() *float64
	StorageAccountAccessKey() *string
	SetStorageAccountAccessKey(val *string)
	StorageAccountAccessKeyInput() *string
	StorageBlobEndpoint() *string
	SetStorageBlobEndpoint(val *string)
	StorageBlobEndpointInput() *string
	SystemDatabasesBackupEnabled() interface{}
	SetSystemDatabasesBackupEnabled(val interface{})
	SystemDatabasesBackupEnabledInput() interface{}
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
	PutManualSchedule(value *MssqlVirtualMachineAutoBackupManualSchedule)
	ResetEncryptionEnabled()
	ResetEncryptionPassword()
	ResetManualSchedule()
	ResetSystemDatabasesBackupEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineAutoBackupOutputReference
type jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) EncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) EncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) EncryptionPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) EncryptionPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) InternalValue() *MssqlVirtualMachineAutoBackup {
	var returns *MssqlVirtualMachineAutoBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ManualSchedule() MssqlVirtualMachineAutoBackupManualScheduleOutputReference {
	var returns MssqlVirtualMachineAutoBackupManualScheduleOutputReference
	_jsii_.Get(
		j,
		"manualSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ManualScheduleInput() *MssqlVirtualMachineAutoBackupManualSchedule {
	var returns *MssqlVirtualMachineAutoBackupManualSchedule
	_jsii_.Get(
		j,
		"manualScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) RetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) RetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) StorageBlobEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) StorageBlobEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SystemDatabasesBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemDatabasesBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SystemDatabasesBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemDatabasesBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineAutoBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineAutoBackupOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineAutoBackupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineAutoBackupOutputReference_Override(m MssqlVirtualMachineAutoBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetEncryptionEnabled(val interface{}) {
	if err := j.validateSetEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetEncryptionPassword(val *string) {
	if err := j.validateSetEncryptionPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionPassword",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetInternalValue(val *MssqlVirtualMachineAutoBackup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetRetentionPeriodInDays(val *float64) {
	if err := j.validateSetRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetStorageAccountAccessKey(val *string) {
	if err := j.validateSetStorageAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetStorageBlobEndpoint(val *string) {
	if err := j.validateSetStorageBlobEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageBlobEndpoint",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetSystemDatabasesBackupEnabled(val interface{}) {
	if err := j.validateSetSystemDatabasesBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemDatabasesBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) PutManualSchedule(value *MssqlVirtualMachineAutoBackupManualSchedule) {
	if err := m.validatePutManualScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putManualSchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ResetEncryptionEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ResetEncryptionPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ResetManualSchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetManualSchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ResetSystemDatabasesBackupEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetSystemDatabasesBackupEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

