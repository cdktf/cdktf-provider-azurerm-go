// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetbinary

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/datafactorydatasetbinary/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference interface {
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
	Container() *string
	SetContainer(val *string)
	ContainerInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DynamicContainerEnabled() interface{}
	SetDynamicContainerEnabled(val interface{})
	DynamicContainerEnabledInput() interface{}
	DynamicFilenameEnabled() interface{}
	SetDynamicFilenameEnabled(val interface{})
	DynamicFilenameEnabledInput() interface{}
	DynamicPathEnabled() interface{}
	SetDynamicPathEnabled(val interface{})
	DynamicPathEnabledInput() interface{}
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataFactoryDatasetBinaryAzureBlobStorageLocation
	SetInternalValue(val *DataFactoryDatasetBinaryAzureBlobStorageLocation)
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	ResetDynamicContainerEnabled()
	ResetDynamicFilenameEnabled()
	ResetDynamicPathEnabled()
	ResetFilename()
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference
type jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) Container() *string {
	var returns *string
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) DynamicContainerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicContainerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) DynamicContainerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicContainerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) DynamicFilenameEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicFilenameEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) DynamicFilenameEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicFilenameEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) DynamicPathEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPathEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) DynamicPathEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPathEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) InternalValue() *DataFactoryDatasetBinaryAzureBlobStorageLocation {
	var returns *DataFactoryDatasetBinaryAzureBlobStorageLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryDatasetBinaryAzureBlobStorageLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetBinary.DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference_Override(d DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetBinary.DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetContainer(val *string) {
	if err := j.validateSetContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetDynamicContainerEnabled(val interface{}) {
	if err := j.validateSetDynamicContainerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicContainerEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetDynamicFilenameEnabled(val interface{}) {
	if err := j.validateSetDynamicFilenameEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicFilenameEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetDynamicPathEnabled(val interface{}) {
	if err := j.validateSetDynamicPathEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicPathEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetFilename(val *string) {
	if err := j.validateSetFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetInternalValue(val *DataFactoryDatasetBinaryAzureBlobStorageLocation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ResetDynamicContainerEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicContainerEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ResetDynamicFilenameEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicFilenameEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ResetDynamicPathEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicPathEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ResetFilename() {
	_jsii_.InvokeVoid(
		d,
		"resetFilename",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		d,
		"resetPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinaryAzureBlobStorageLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

