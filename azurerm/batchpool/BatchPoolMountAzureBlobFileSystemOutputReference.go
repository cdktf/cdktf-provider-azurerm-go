// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/batchpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchPoolMountAzureBlobFileSystemOutputReference interface {
	cdktf.ComplexObject
	AccountKey() *string
	SetAccountKey(val *string)
	AccountKeyInput() *string
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	BlobfuseOptions() *string
	SetBlobfuseOptions(val *string)
	BlobfuseOptionsInput() *string
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
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IdentityId() *string
	SetIdentityId(val *string)
	IdentityIdInput() *string
	InternalValue() *BatchPoolMountAzureBlobFileSystem
	SetInternalValue(val *BatchPoolMountAzureBlobFileSystem)
	RelativeMountPath() *string
	SetRelativeMountPath(val *string)
	RelativeMountPathInput() *string
	SasKey() *string
	SetSasKey(val *string)
	SasKeyInput() *string
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
	ResetAccountKey()
	ResetBlobfuseOptions()
	ResetIdentityId()
	ResetSasKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountAzureBlobFileSystemOutputReference
type jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) AccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) AccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) BlobfuseOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobfuseOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) BlobfuseOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobfuseOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) IdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) IdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) InternalValue() *BatchPoolMountAzureBlobFileSystem {
	var returns *BatchPoolMountAzureBlobFileSystem
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) RelativeMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) RelativeMountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeMountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SasKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) SasKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolMountAzureBlobFileSystemOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolMountAzureBlobFileSystemOutputReference {
	_init_.Initialize()

	if err := validateNewBatchPoolMountAzureBlobFileSystemOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountAzureBlobFileSystemOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolMountAzureBlobFileSystemOutputReference_Override(b BatchPoolMountAzureBlobFileSystemOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountAzureBlobFileSystemOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetAccountKey(val *string) {
	if err := j.validateSetAccountKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountKey",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetAccountName(val *string) {
	if err := j.validateSetAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetBlobfuseOptions(val *string) {
	if err := j.validateSetBlobfuseOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blobfuseOptions",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetIdentityId(val *string) {
	if err := j.validateSetIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityId",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetInternalValue(val *BatchPoolMountAzureBlobFileSystem) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetRelativeMountPath(val *string) {
	if err := j.validateSetRelativeMountPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relativeMountPath",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetSasKey(val *string) {
	if err := j.validateSetSasKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sasKey",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ResetAccountKey() {
	_jsii_.InvokeVoid(
		b,
		"resetAccountKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ResetBlobfuseOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetBlobfuseOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ResetIdentityId() {
	_jsii_.InvokeVoid(
		b,
		"resetIdentityId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ResetSasKey() {
	_jsii_.InvokeVoid(
		b,
		"resetSasKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountAzureBlobFileSystemOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

