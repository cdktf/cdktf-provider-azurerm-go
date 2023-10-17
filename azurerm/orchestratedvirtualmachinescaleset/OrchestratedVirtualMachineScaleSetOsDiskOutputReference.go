// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/orchestratedvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrchestratedVirtualMachineScaleSetOsDiskOutputReference interface {
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
	DiffDiskSettings() OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
	DiffDiskSettingsInput() *OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettings
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *OrchestratedVirtualMachineScaleSetOsDisk
	SetInternalValue(val *OrchestratedVirtualMachineScaleSetOsDisk)
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
	PutDiffDiskSettings(value *OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettings)
	ResetDiffDiskSettings()
	ResetDiskEncryptionSetId()
	ResetDiskSizeGb()
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

// The jsii proxy struct for OrchestratedVirtualMachineScaleSetOsDiskOutputReference
type jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) DiffDiskSettings() OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference {
	var returns OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
	_jsii_.Get(
		j,
		"diffDiskSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) DiffDiskSettingsInput() *OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettings {
	var returns *OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettings
	_jsii_.Get(
		j,
		"diffDiskSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) InternalValue() *OrchestratedVirtualMachineScaleSetOsDisk {
	var returns *OrchestratedVirtualMachineScaleSetOsDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) WriteAcceleratorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) WriteAcceleratorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabledInput",
		&returns,
	)
	return returns
}


func NewOrchestratedVirtualMachineScaleSetOsDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrchestratedVirtualMachineScaleSetOsDiskOutputReference {
	_init_.Initialize()

	if err := validateNewOrchestratedVirtualMachineScaleSetOsDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrchestratedVirtualMachineScaleSetOsDiskOutputReference_Override(o OrchestratedVirtualMachineScaleSetOsDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetCaching(val *string) {
	if err := j.validateSetCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetDiskEncryptionSetId(val *string) {
	if err := j.validateSetDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetInternalValue(val *OrchestratedVirtualMachineScaleSetOsDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetStorageAccountType(val *string) {
	if err := j.validateSetStorageAccountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference)SetWriteAcceleratorEnabled(val interface{}) {
	if err := j.validateSetWriteAcceleratorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeAcceleratorEnabled",
		val,
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) PutDiffDiskSettings(value *OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettings) {
	if err := o.validatePutDiffDiskSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDiffDiskSettings",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) ResetDiffDiskSettings() {
	_jsii_.InvokeVoid(
		o,
		"resetDiffDiskSettings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		o,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) ResetWriteAcceleratorEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetWriteAcceleratorEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

