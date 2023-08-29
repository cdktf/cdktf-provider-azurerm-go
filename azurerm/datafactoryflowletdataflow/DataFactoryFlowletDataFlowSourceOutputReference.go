// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryflowletdataflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/datafactoryflowletdataflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryFlowletDataFlowSourceOutputReference interface {
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
	Dataset() DataFactoryFlowletDataFlowSourceDatasetOutputReference
	DatasetInput() *DataFactoryFlowletDataFlowSourceDataset
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Flowlet() DataFactoryFlowletDataFlowSourceFlowletOutputReference
	FlowletInput() *DataFactoryFlowletDataFlowSourceFlowlet
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinkedService() DataFactoryFlowletDataFlowSourceLinkedServiceOutputReference
	LinkedServiceInput() *DataFactoryFlowletDataFlowSourceLinkedService
	Name() *string
	SetName(val *string)
	NameInput() *string
	RejectedLinkedService() DataFactoryFlowletDataFlowSourceRejectedLinkedServiceOutputReference
	RejectedLinkedServiceInput() *DataFactoryFlowletDataFlowSourceRejectedLinkedService
	SchemaLinkedService() DataFactoryFlowletDataFlowSourceSchemaLinkedServiceOutputReference
	SchemaLinkedServiceInput() *DataFactoryFlowletDataFlowSourceSchemaLinkedService
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
	PutDataset(value *DataFactoryFlowletDataFlowSourceDataset)
	PutFlowlet(value *DataFactoryFlowletDataFlowSourceFlowlet)
	PutLinkedService(value *DataFactoryFlowletDataFlowSourceLinkedService)
	PutRejectedLinkedService(value *DataFactoryFlowletDataFlowSourceRejectedLinkedService)
	PutSchemaLinkedService(value *DataFactoryFlowletDataFlowSourceSchemaLinkedService)
	ResetDataset()
	ResetDescription()
	ResetFlowlet()
	ResetLinkedService()
	ResetRejectedLinkedService()
	ResetSchemaLinkedService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryFlowletDataFlowSourceOutputReference
type jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) Dataset() DataFactoryFlowletDataFlowSourceDatasetOutputReference {
	var returns DataFactoryFlowletDataFlowSourceDatasetOutputReference
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) DatasetInput() *DataFactoryFlowletDataFlowSourceDataset {
	var returns *DataFactoryFlowletDataFlowSourceDataset
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) Flowlet() DataFactoryFlowletDataFlowSourceFlowletOutputReference {
	var returns DataFactoryFlowletDataFlowSourceFlowletOutputReference
	_jsii_.Get(
		j,
		"flowlet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) FlowletInput() *DataFactoryFlowletDataFlowSourceFlowlet {
	var returns *DataFactoryFlowletDataFlowSourceFlowlet
	_jsii_.Get(
		j,
		"flowletInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) LinkedService() DataFactoryFlowletDataFlowSourceLinkedServiceOutputReference {
	var returns DataFactoryFlowletDataFlowSourceLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"linkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) LinkedServiceInput() *DataFactoryFlowletDataFlowSourceLinkedService {
	var returns *DataFactoryFlowletDataFlowSourceLinkedService
	_jsii_.Get(
		j,
		"linkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) RejectedLinkedService() DataFactoryFlowletDataFlowSourceRejectedLinkedServiceOutputReference {
	var returns DataFactoryFlowletDataFlowSourceRejectedLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"rejectedLinkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) RejectedLinkedServiceInput() *DataFactoryFlowletDataFlowSourceRejectedLinkedService {
	var returns *DataFactoryFlowletDataFlowSourceRejectedLinkedService
	_jsii_.Get(
		j,
		"rejectedLinkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) SchemaLinkedService() DataFactoryFlowletDataFlowSourceSchemaLinkedServiceOutputReference {
	var returns DataFactoryFlowletDataFlowSourceSchemaLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"schemaLinkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) SchemaLinkedServiceInput() *DataFactoryFlowletDataFlowSourceSchemaLinkedService {
	var returns *DataFactoryFlowletDataFlowSourceSchemaLinkedService
	_jsii_.Get(
		j,
		"schemaLinkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryFlowletDataFlowSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataFactoryFlowletDataFlowSourceOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryFlowletDataFlowSourceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryFlowletDataFlow.DataFactoryFlowletDataFlowSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataFactoryFlowletDataFlowSourceOutputReference_Override(d DataFactoryFlowletDataFlowSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryFlowletDataFlow.DataFactoryFlowletDataFlowSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) PutDataset(value *DataFactoryFlowletDataFlowSourceDataset) {
	if err := d.validatePutDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataset",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) PutFlowlet(value *DataFactoryFlowletDataFlowSourceFlowlet) {
	if err := d.validatePutFlowletParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlowlet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) PutLinkedService(value *DataFactoryFlowletDataFlowSourceLinkedService) {
	if err := d.validatePutLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) PutRejectedLinkedService(value *DataFactoryFlowletDataFlowSourceRejectedLinkedService) {
	if err := d.validatePutRejectedLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRejectedLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) PutSchemaLinkedService(value *DataFactoryFlowletDataFlowSourceSchemaLinkedService) {
	if err := d.validatePutSchemaLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchemaLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ResetDataset() {
	_jsii_.InvokeVoid(
		d,
		"resetDataset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ResetFlowlet() {
	_jsii_.InvokeVoid(
		d,
		"resetFlowlet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ResetLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ResetRejectedLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetRejectedLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ResetSchemaLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

