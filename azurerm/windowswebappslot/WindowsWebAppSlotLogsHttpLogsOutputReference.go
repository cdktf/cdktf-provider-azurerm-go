// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/windowswebappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppSlotLogsHttpLogsOutputReference interface {
	cdktf.ComplexObject
	AzureBlobStorage() WindowsWebAppSlotLogsHttpLogsAzureBlobStorageOutputReference
	AzureBlobStorageInput() *WindowsWebAppSlotLogsHttpLogsAzureBlobStorage
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
	FileSystem() WindowsWebAppSlotLogsHttpLogsFileSystemOutputReference
	FileSystemInput() *WindowsWebAppSlotLogsHttpLogsFileSystem
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSlotLogsHttpLogs
	SetInternalValue(val *WindowsWebAppSlotLogsHttpLogs)
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
	PutAzureBlobStorage(value *WindowsWebAppSlotLogsHttpLogsAzureBlobStorage)
	PutFileSystem(value *WindowsWebAppSlotLogsHttpLogsFileSystem)
	ResetAzureBlobStorage()
	ResetFileSystem()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSlotLogsHttpLogsOutputReference
type jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) AzureBlobStorage() WindowsWebAppSlotLogsHttpLogsAzureBlobStorageOutputReference {
	var returns WindowsWebAppSlotLogsHttpLogsAzureBlobStorageOutputReference
	_jsii_.Get(
		j,
		"azureBlobStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) AzureBlobStorageInput() *WindowsWebAppSlotLogsHttpLogsAzureBlobStorage {
	var returns *WindowsWebAppSlotLogsHttpLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"azureBlobStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) FileSystem() WindowsWebAppSlotLogsHttpLogsFileSystemOutputReference {
	var returns WindowsWebAppSlotLogsHttpLogsFileSystemOutputReference
	_jsii_.Get(
		j,
		"fileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) FileSystemInput() *WindowsWebAppSlotLogsHttpLogsFileSystem {
	var returns *WindowsWebAppSlotLogsHttpLogsFileSystem
	_jsii_.Get(
		j,
		"fileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) InternalValue() *WindowsWebAppSlotLogsHttpLogs {
	var returns *WindowsWebAppSlotLogsHttpLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSlotLogsHttpLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSlotLogsHttpLogsOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsWebAppSlotLogsHttpLogsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotLogsHttpLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSlotLogsHttpLogsOutputReference_Override(w WindowsWebAppSlotLogsHttpLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotLogsHttpLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference)SetInternalValue(val *WindowsWebAppSlotLogsHttpLogs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) PutAzureBlobStorage(value *WindowsWebAppSlotLogsHttpLogsAzureBlobStorage) {
	if err := w.validatePutAzureBlobStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAzureBlobStorage",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) PutFileSystem(value *WindowsWebAppSlotLogsHttpLogsFileSystem) {
	if err := w.validatePutFileSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putFileSystem",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) ResetAzureBlobStorage() {
	_jsii_.InvokeVoid(
		w,
		"resetAzureBlobStorage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) ResetFileSystem() {
	_jsii_.InvokeVoid(
		w,
		"resetFileSystem",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WindowsWebAppSlotLogsHttpLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

