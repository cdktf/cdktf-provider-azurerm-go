// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescalesetpacketcapture

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/virtualmachinescalesetpacketcapture/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference interface {
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
	FilePath() *string
	SetFilePath(val *string)
	FilePathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *VirtualMachineScaleSetPacketCaptureStorageLocation
	SetInternalValue(val *VirtualMachineScaleSetPacketCaptureStorageLocation)
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	StoragePath() *string
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
	ResetFilePath()
	ResetStorageAccountId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference
type jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) FilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) InternalValue() *VirtualMachineScaleSetPacketCaptureStorageLocation {
	var returns *VirtualMachineScaleSetPacketCaptureStorageLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) StoragePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualMachineScaleSetPacketCaptureStorageLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetPacketCaptureStorageLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSetPacketCapture.VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualMachineScaleSetPacketCaptureStorageLocationOutputReference_Override(v VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSetPacketCapture.VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference)SetFilePath(val *string) {
	if err := j.validateSetFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePath",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference)SetInternalValue(val *VirtualMachineScaleSetPacketCaptureStorageLocation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference)SetStorageAccountId(val *string) {
	if err := j.validateSetStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) ResetFilePath() {
	_jsii_.InvokeVoid(
		v,
		"resetFilePath",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) ResetStorageAccountId() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageAccountId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetPacketCaptureStorageLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

