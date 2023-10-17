// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetbinary

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/datafactorydatasetbinary/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDatasetBinarySftpServerLocationOutputReference interface {
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
	InternalValue() *DataFactoryDatasetBinarySftpServerLocation
	SetInternalValue(val *DataFactoryDatasetBinarySftpServerLocation)
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
	ResetDynamicFilenameEnabled()
	ResetDynamicPathEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryDatasetBinarySftpServerLocationOutputReference
type jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) DynamicFilenameEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicFilenameEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) DynamicFilenameEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicFilenameEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) DynamicPathEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPathEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) DynamicPathEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPathEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) InternalValue() *DataFactoryDatasetBinarySftpServerLocation {
	var returns *DataFactoryDatasetBinarySftpServerLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryDatasetBinarySftpServerLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryDatasetBinarySftpServerLocationOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryDatasetBinarySftpServerLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetBinary.DataFactoryDatasetBinarySftpServerLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryDatasetBinarySftpServerLocationOutputReference_Override(d DataFactoryDatasetBinarySftpServerLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetBinary.DataFactoryDatasetBinarySftpServerLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetDynamicFilenameEnabled(val interface{}) {
	if err := j.validateSetDynamicFilenameEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicFilenameEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetDynamicPathEnabled(val interface{}) {
	if err := j.validateSetDynamicPathEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicPathEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetFilename(val *string) {
	if err := j.validateSetFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetInternalValue(val *DataFactoryDatasetBinarySftpServerLocation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) ResetDynamicFilenameEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicFilenameEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) ResetDynamicPathEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicPathEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataFactoryDatasetBinarySftpServerLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

