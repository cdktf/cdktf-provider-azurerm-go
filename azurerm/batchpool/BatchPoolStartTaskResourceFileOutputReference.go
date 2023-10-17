// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/batchpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchPoolStartTaskResourceFileOutputReference interface {
	cdktf.ComplexObject
	AutoStorageContainerName() *string
	SetAutoStorageContainerName(val *string)
	AutoStorageContainerNameInput() *string
	BlobPrefix() *string
	SetBlobPrefix(val *string)
	BlobPrefixInput() *string
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
	FileMode() *string
	SetFileMode(val *string)
	FileModeInput() *string
	FilePath() *string
	SetFilePath(val *string)
	FilePathInput() *string
	// Experimental.
	Fqn() *string
	HttpUrl() *string
	SetHttpUrl(val *string)
	HttpUrlInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageContainerUrl() *string
	SetStorageContainerUrl(val *string)
	StorageContainerUrlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAssignedIdentityId() *string
	SetUserAssignedIdentityId(val *string)
	UserAssignedIdentityIdInput() *string
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
	ResetAutoStorageContainerName()
	ResetBlobPrefix()
	ResetFileMode()
	ResetFilePath()
	ResetHttpUrl()
	ResetStorageContainerUrl()
	ResetUserAssignedIdentityId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolStartTaskResourceFileOutputReference
type jsiiProxy_BatchPoolStartTaskResourceFileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) AutoStorageContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoStorageContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) AutoStorageContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoStorageContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) BlobPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) BlobPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) FileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) FileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) FilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) HttpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) HttpUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) StorageContainerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) StorageContainerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) UserAssignedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityIdInput",
		&returns,
	)
	return returns
}


func NewBatchPoolStartTaskResourceFileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolStartTaskResourceFileOutputReference {
	_init_.Initialize()

	if err := validateNewBatchPoolStartTaskResourceFileOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchPoolStartTaskResourceFileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskResourceFileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolStartTaskResourceFileOutputReference_Override(b BatchPoolStartTaskResourceFileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskResourceFileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetAutoStorageContainerName(val *string) {
	if err := j.validateSetAutoStorageContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStorageContainerName",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetBlobPrefix(val *string) {
	if err := j.validateSetBlobPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blobPrefix",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetFileMode(val *string) {
	if err := j.validateSetFileModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileMode",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetFilePath(val *string) {
	if err := j.validateSetFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePath",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetHttpUrl(val *string) {
	if err := j.validateSetHttpUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpUrl",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetStorageContainerUrl(val *string) {
	if err := j.validateSetStorageContainerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageContainerUrl",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference)SetUserAssignedIdentityId(val *string) {
	if err := j.validateSetUserAssignedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAssignedIdentityId",
		val,
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetAutoStorageContainerName() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoStorageContainerName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetBlobPrefix() {
	_jsii_.InvokeVoid(
		b,
		"resetBlobPrefix",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetFileMode() {
	_jsii_.InvokeVoid(
		b,
		"resetFileMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetFilePath() {
	_jsii_.InvokeVoid(
		b,
		"resetFilePath",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetHttpUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetHttpUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetStorageContainerUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageContainerUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ResetUserAssignedIdentityId() {
	_jsii_.InvokeVoid(
		b,
		"resetUserAssignedIdentityId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskResourceFileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

