// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/windowsvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsVirtualMachineScaleSetOsDiskOutputReference interface {
	cdktf.ComplexObject
	Caching() *string
	SetCaching(val *string)
	CachingInput() *string
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
	DiffDiskSettings() WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
	DiffDiskSettingsInput() *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsVirtualMachineScaleSetOsDisk
	SetInternalValue(val *WindowsVirtualMachineScaleSetOsDisk)
	SecureVmDiskEncryptionSetId() *string
	SetSecureVmDiskEncryptionSetId(val *string)
	SecureVmDiskEncryptionSetIdInput() *string
	SecurityEncryptionType() *string
	SetSecurityEncryptionType(val *string)
	SecurityEncryptionTypeInput() *string
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WriteAcceleratorEnabled() interface{}
	SetWriteAcceleratorEnabled(val interface{})
	WriteAcceleratorEnabledInput() interface{}
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
	PutDiffDiskSettings(value *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings)
	ResetDiffDiskSettings()
	ResetDiskEncryptionSetId()
	ResetDiskSizeGb()
	ResetSecureVmDiskEncryptionSetId()
	ResetSecurityEncryptionType()
	ResetWriteAcceleratorEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsVirtualMachineScaleSetOsDiskOutputReference
type jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiffDiskSettings() WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference {
	var returns WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
	_jsii_.Get(
		j,
		"diffDiskSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiffDiskSettingsInput() *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings {
	var returns *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings
	_jsii_.Get(
		j,
		"diffDiskSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) InternalValue() *WindowsVirtualMachineScaleSetOsDisk {
	var returns *WindowsVirtualMachineScaleSetOsDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SecureVmDiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureVmDiskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SecureVmDiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureVmDiskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SecurityEncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) SecurityEncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityEncryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) WriteAcceleratorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) WriteAcceleratorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabledInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineScaleSetOsDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsVirtualMachineScaleSetOsDiskOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsVirtualMachineScaleSetOsDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineScaleSetOsDiskOutputReference_Override(w WindowsVirtualMachineScaleSetOsDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSetOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetCaching(val *string) {
	if err := j.validateSetCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetDiskEncryptionSetId(val *string) {
	if err := j.validateSetDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetInternalValue(val *WindowsVirtualMachineScaleSetOsDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetSecureVmDiskEncryptionSetId(val *string) {
	if err := j.validateSetSecureVmDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureVmDiskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetSecurityEncryptionType(val *string) {
	if err := j.validateSetSecurityEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityEncryptionType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetStorageAccountType(val *string) {
	if err := j.validateSetStorageAccountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference)SetWriteAcceleratorEnabled(val interface{}) {
	if err := j.validateSetWriteAcceleratorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeAcceleratorEnabled",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) PutDiffDiskSettings(value *WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings) {
	if err := w.validatePutDiffDiskSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDiffDiskSettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetDiffDiskSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetDiffDiskSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetSecureVmDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		w,
		"resetSecureVmDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetSecurityEncryptionType() {
	_jsii_.InvokeVoid(
		w,
		"resetSecurityEncryptionType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ResetWriteAcceleratorEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetWriteAcceleratorEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSetOsDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

