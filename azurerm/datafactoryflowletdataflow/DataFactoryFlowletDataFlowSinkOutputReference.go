// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryflowletdataflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/datafactoryflowletdataflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryFlowletDataFlowSinkOutputReference interface {
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
	Dataset() DataFactoryFlowletDataFlowSinkDatasetOutputReference
	DatasetInput() *DataFactoryFlowletDataFlowSinkDataset
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Flowlet() DataFactoryFlowletDataFlowSinkFlowletOutputReference
	FlowletInput() *DataFactoryFlowletDataFlowSinkFlowlet
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinkedService() DataFactoryFlowletDataFlowSinkLinkedServiceOutputReference
	LinkedServiceInput() *DataFactoryFlowletDataFlowSinkLinkedService
	Name() *string
	SetName(val *string)
	NameInput() *string
	RejectedLinkedService() DataFactoryFlowletDataFlowSinkRejectedLinkedServiceOutputReference
	RejectedLinkedServiceInput() *DataFactoryFlowletDataFlowSinkRejectedLinkedService
	SchemaLinkedService() DataFactoryFlowletDataFlowSinkSchemaLinkedServiceOutputReference
	SchemaLinkedServiceInput() *DataFactoryFlowletDataFlowSinkSchemaLinkedService
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
	PutDataset(value *DataFactoryFlowletDataFlowSinkDataset)
	PutFlowlet(value *DataFactoryFlowletDataFlowSinkFlowlet)
	PutLinkedService(value *DataFactoryFlowletDataFlowSinkLinkedService)
	PutRejectedLinkedService(value *DataFactoryFlowletDataFlowSinkRejectedLinkedService)
	PutSchemaLinkedService(value *DataFactoryFlowletDataFlowSinkSchemaLinkedService)
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

// The jsii proxy struct for DataFactoryFlowletDataFlowSinkOutputReference
type jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) Dataset() DataFactoryFlowletDataFlowSinkDatasetOutputReference {
	var returns DataFactoryFlowletDataFlowSinkDatasetOutputReference
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) DatasetInput() *DataFactoryFlowletDataFlowSinkDataset {
	var returns *DataFactoryFlowletDataFlowSinkDataset
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) Flowlet() DataFactoryFlowletDataFlowSinkFlowletOutputReference {
	var returns DataFactoryFlowletDataFlowSinkFlowletOutputReference
	_jsii_.Get(
		j,
		"flowlet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) FlowletInput() *DataFactoryFlowletDataFlowSinkFlowlet {
	var returns *DataFactoryFlowletDataFlowSinkFlowlet
	_jsii_.Get(
		j,
		"flowletInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) LinkedService() DataFactoryFlowletDataFlowSinkLinkedServiceOutputReference {
	var returns DataFactoryFlowletDataFlowSinkLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"linkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) LinkedServiceInput() *DataFactoryFlowletDataFlowSinkLinkedService {
	var returns *DataFactoryFlowletDataFlowSinkLinkedService
	_jsii_.Get(
		j,
		"linkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) RejectedLinkedService() DataFactoryFlowletDataFlowSinkRejectedLinkedServiceOutputReference {
	var returns DataFactoryFlowletDataFlowSinkRejectedLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"rejectedLinkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) RejectedLinkedServiceInput() *DataFactoryFlowletDataFlowSinkRejectedLinkedService {
	var returns *DataFactoryFlowletDataFlowSinkRejectedLinkedService
	_jsii_.Get(
		j,
		"rejectedLinkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) SchemaLinkedService() DataFactoryFlowletDataFlowSinkSchemaLinkedServiceOutputReference {
	var returns DataFactoryFlowletDataFlowSinkSchemaLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"schemaLinkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) SchemaLinkedServiceInput() *DataFactoryFlowletDataFlowSinkSchemaLinkedService {
	var returns *DataFactoryFlowletDataFlowSinkSchemaLinkedService
	_jsii_.Get(
		j,
		"schemaLinkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryFlowletDataFlowSinkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataFactoryFlowletDataFlowSinkOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryFlowletDataFlowSinkOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryFlowletDataFlow.DataFactoryFlowletDataFlowSinkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataFactoryFlowletDataFlowSinkOutputReference_Override(d DataFactoryFlowletDataFlowSinkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryFlowletDataFlow.DataFactoryFlowletDataFlowSinkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) PutDataset(value *DataFactoryFlowletDataFlowSinkDataset) {
	if err := d.validatePutDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataset",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) PutFlowlet(value *DataFactoryFlowletDataFlowSinkFlowlet) {
	if err := d.validatePutFlowletParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlowlet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) PutLinkedService(value *DataFactoryFlowletDataFlowSinkLinkedService) {
	if err := d.validatePutLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) PutRejectedLinkedService(value *DataFactoryFlowletDataFlowSinkRejectedLinkedService) {
	if err := d.validatePutRejectedLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRejectedLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) PutSchemaLinkedService(value *DataFactoryFlowletDataFlowSinkSchemaLinkedService) {
	if err := d.validatePutSchemaLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchemaLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ResetDataset() {
	_jsii_.InvokeVoid(
		d,
		"resetDataset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ResetFlowlet() {
	_jsii_.InvokeVoid(
		d,
		"resetFlowlet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ResetLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ResetRejectedLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetRejectedLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ResetSchemaLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataFactoryFlowletDataFlowSinkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

