// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryflowletdataflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/datafactoryflowletdataflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryFlowletDataFlowTransformationOutputReference interface {
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
	Dataset() DataFactoryFlowletDataFlowTransformationDatasetOutputReference
	DatasetInput() *DataFactoryFlowletDataFlowTransformationDataset
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Flowlet() DataFactoryFlowletDataFlowTransformationFlowletOutputReference
	FlowletInput() *DataFactoryFlowletDataFlowTransformationFlowlet
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinkedService() DataFactoryFlowletDataFlowTransformationLinkedServiceOutputReference
	LinkedServiceInput() *DataFactoryFlowletDataFlowTransformationLinkedService
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutDataset(value *DataFactoryFlowletDataFlowTransformationDataset)
	PutFlowlet(value *DataFactoryFlowletDataFlowTransformationFlowlet)
	PutLinkedService(value *DataFactoryFlowletDataFlowTransformationLinkedService)
	ResetDataset()
	ResetDescription()
	ResetFlowlet()
	ResetLinkedService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryFlowletDataFlowTransformationOutputReference
type jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) Dataset() DataFactoryFlowletDataFlowTransformationDatasetOutputReference {
	var returns DataFactoryFlowletDataFlowTransformationDatasetOutputReference
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) DatasetInput() *DataFactoryFlowletDataFlowTransformationDataset {
	var returns *DataFactoryFlowletDataFlowTransformationDataset
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) Flowlet() DataFactoryFlowletDataFlowTransformationFlowletOutputReference {
	var returns DataFactoryFlowletDataFlowTransformationFlowletOutputReference
	_jsii_.Get(
		j,
		"flowlet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) FlowletInput() *DataFactoryFlowletDataFlowTransformationFlowlet {
	var returns *DataFactoryFlowletDataFlowTransformationFlowlet
	_jsii_.Get(
		j,
		"flowletInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) LinkedService() DataFactoryFlowletDataFlowTransformationLinkedServiceOutputReference {
	var returns DataFactoryFlowletDataFlowTransformationLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"linkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) LinkedServiceInput() *DataFactoryFlowletDataFlowTransformationLinkedService {
	var returns *DataFactoryFlowletDataFlowTransformationLinkedService
	_jsii_.Get(
		j,
		"linkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryFlowletDataFlowTransformationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataFactoryFlowletDataFlowTransformationOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryFlowletDataFlowTransformationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryFlowletDataFlow.DataFactoryFlowletDataFlowTransformationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataFactoryFlowletDataFlowTransformationOutputReference_Override(d DataFactoryFlowletDataFlowTransformationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryFlowletDataFlow.DataFactoryFlowletDataFlowTransformationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) PutDataset(value *DataFactoryFlowletDataFlowTransformationDataset) {
	if err := d.validatePutDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataset",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) PutFlowlet(value *DataFactoryFlowletDataFlowTransformationFlowlet) {
	if err := d.validatePutFlowletParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlowlet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) PutLinkedService(value *DataFactoryFlowletDataFlowTransformationLinkedService) {
	if err := d.validatePutLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) ResetDataset() {
	_jsii_.InvokeVoid(
		d,
		"resetDataset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) ResetFlowlet() {
	_jsii_.InvokeVoid(
		d,
		"resetFlowlet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) ResetLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowTransformationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

