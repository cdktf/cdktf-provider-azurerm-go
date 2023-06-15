package image

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/image/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImageDataDiskOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lun() *float64
	SetLun(val *float64)
	LunInput() *float64
	ManagedDiskId() *string
	SetManagedDiskId(val *string)
	ManagedDiskIdInput() *string
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
	ResetLun()
	ResetManagedDiskId()
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

// The jsii proxy struct for ImageDataDiskOutputReference
type jsiiProxy_ImageDataDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImageDataDiskOutputReference) BlobUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) BlobUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) Lun() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) LunInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) ManagedDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) ManagedDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) SizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) SizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageDataDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImageDataDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ImageDataDiskOutputReference {
	_init_.Initialize()

	if err := validateNewImageDataDiskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImageDataDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.image.ImageDataDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewImageDataDiskOutputReference_Override(i ImageDataDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.image.ImageDataDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetBlobUri(val *string) {
	if err := j.validateSetBlobUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blobUri",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetCaching(val *string) {
	if err := j.validateSetCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetLun(val *float64) {
	if err := j.validateSetLunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lun",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetManagedDiskId(val *string) {
	if err := j.validateSetManagedDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDiskId",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetSizeGb(val *float64) {
	if err := j.validateSetSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeGb",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImageDataDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImageDataDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageDataDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageDataDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) ResetBlobUri() {
	_jsii_.InvokeVoid(
		i,
		"resetBlobUri",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageDataDiskOutputReference) ResetCaching() {
	_jsii_.InvokeVoid(
		i,
		"resetCaching",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageDataDiskOutputReference) ResetLun() {
	_jsii_.InvokeVoid(
		i,
		"resetLun",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageDataDiskOutputReference) ResetManagedDiskId() {
	_jsii_.InvokeVoid(
		i,
		"resetManagedDiskId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageDataDiskOutputReference) ResetSizeGb() {
	_jsii_.InvokeVoid(
		i,
		"resetSizeGb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageDataDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_ImageDataDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

