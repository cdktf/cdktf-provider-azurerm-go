// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/batchpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchPoolMountOutputReference interface {
	cdktf.ComplexObject
	AzureBlobFileSystem() BatchPoolMountAzureBlobFileSystemOutputReference
	AzureBlobFileSystemInput() *BatchPoolMountAzureBlobFileSystem
	AzureFileShare() BatchPoolMountAzureFileShareList
	AzureFileShareInput() interface{}
	CifsMount() BatchPoolMountCifsMountList
	CifsMountInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NfsMount() BatchPoolMountNfsMountList
	NfsMountInput() interface{}
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
	PutAzureBlobFileSystem(value *BatchPoolMountAzureBlobFileSystem)
	PutAzureFileShare(value interface{})
	PutCifsMount(value interface{})
	PutNfsMount(value interface{})
	ResetAzureBlobFileSystem()
	ResetAzureFileShare()
	ResetCifsMount()
	ResetNfsMount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolMountOutputReference
type jsiiProxy_BatchPoolMountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolMountOutputReference) AzureBlobFileSystem() BatchPoolMountAzureBlobFileSystemOutputReference {
	var returns BatchPoolMountAzureBlobFileSystemOutputReference
	_jsii_.Get(
		j,
		"azureBlobFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) AzureBlobFileSystemInput() *BatchPoolMountAzureBlobFileSystem {
	var returns *BatchPoolMountAzureBlobFileSystem
	_jsii_.Get(
		j,
		"azureBlobFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) AzureFileShare() BatchPoolMountAzureFileShareList {
	var returns BatchPoolMountAzureFileShareList
	_jsii_.Get(
		j,
		"azureFileShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) AzureFileShareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureFileShareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) CifsMount() BatchPoolMountCifsMountList {
	var returns BatchPoolMountCifsMountList
	_jsii_.Get(
		j,
		"cifsMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) CifsMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cifsMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) NfsMount() BatchPoolMountNfsMountList {
	var returns BatchPoolMountNfsMountList
	_jsii_.Get(
		j,
		"nfsMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) NfsMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolMountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchPoolMountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchPoolMountOutputReference {
	_init_.Initialize()

	if err := validateNewBatchPoolMountOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchPoolMountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchPoolMountOutputReference_Override(b BatchPoolMountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolMountOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolMountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) PutAzureBlobFileSystem(value *BatchPoolMountAzureBlobFileSystem) {
	if err := b.validatePutAzureBlobFileSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAzureBlobFileSystem",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) PutAzureFileShare(value interface{}) {
	if err := b.validatePutAzureFileShareParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAzureFileShare",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) PutCifsMount(value interface{}) {
	if err := b.validatePutCifsMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCifsMount",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) PutNfsMount(value interface{}) {
	if err := b.validatePutNfsMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNfsMount",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ResetAzureBlobFileSystem() {
	_jsii_.InvokeVoid(
		b,
		"resetAzureBlobFileSystem",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ResetAzureFileShare() {
	_jsii_.InvokeVoid(
		b,
		"resetAzureFileShare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ResetCifsMount() {
	_jsii_.InvokeVoid(
		b,
		"resetCifsMount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) ResetNfsMount() {
	_jsii_.InvokeVoid(
		b,
		"resetNfsMount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolMountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BatchPoolMountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

