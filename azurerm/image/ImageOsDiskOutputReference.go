// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package image

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/image/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImageOsDiskOutputReference interface {
	cdktf.ComplexObject
	BlobUri() *string
	SetBlobUri(val *string)
	BlobUriInput() *string
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
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ImageOsDisk
	SetInternalValue(val *ImageOsDisk)
	ManagedDiskId() *string
	SetManagedDiskId(val *string)
	ManagedDiskIdInput() *string
	OsState() *string
	SetOsState(val *string)
	OsStateInput() *string
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	SizeGb() *float64
	SetSizeGb(val *float64)
	SizeGbInput() *float64
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
	ResetBlobUri()
	ResetCaching()
	ResetDiskEncryptionSetId()
	ResetManagedDiskId()
	ResetOsState()
	ResetOsType()
	ResetSizeGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImageOsDiskOutputReference
type jsiiProxy_ImageOsDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImageOsDiskOutputReference) BlobUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) BlobUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) InternalValue() *ImageOsDisk {
	var returns *ImageOsDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) ManagedDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) ManagedDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) OsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) OsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) SizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) SizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOsDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImageOsDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ImageOsDiskOutputReference {
	_init_.Initialize()

	if err := validateNewImageOsDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImageOsDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.image.ImageOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewImageOsDiskOutputReference_Override(i ImageOsDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.image.ImageOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetBlobUri(val *string) {
	if err := j.validateSetBlobUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blobUri",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetCaching(val *string) {
	if err := j.validateSetCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetDiskEncryptionSetId(val *string) {
	if err := j.validateSetDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetInternalValue(val *ImageOsDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetManagedDiskId(val *string) {
	if err := j.validateSetManagedDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDiskId",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetOsState(val *string) {
	if err := j.validateSetOsStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osState",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetSizeGb(val *float64) {
	if err := j.validateSetSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeGb",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImageOsDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ResetBlobUri() {
	_jsii_.InvokeVoid(
		i,
		"resetBlobUri",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ResetCaching() {
	_jsii_.InvokeVoid(
		i,
		"resetCaching",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		i,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ResetManagedDiskId() {
	_jsii_.InvokeVoid(
		i,
		"resetManagedDiskId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ResetOsState() {
	_jsii_.InvokeVoid(
		i,
		"resetOsState",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ResetOsType() {
	_jsii_.InvokeVoid(
		i,
		"resetOsType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ResetSizeGb() {
	_jsii_.InvokeVoid(
		i,
		"resetSizeGb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageOsDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOsDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

