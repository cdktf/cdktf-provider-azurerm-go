package datafactorydataflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/datafactorydataflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDataFlowSinkOutputReference interface {
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
	Dataset() DataFactoryDataFlowSinkDatasetOutputReference
	DatasetInput() *DataFactoryDataFlowSinkDataset
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Flowlet() DataFactoryDataFlowSinkFlowletOutputReference
	FlowletInput() *DataFactoryDataFlowSinkFlowlet
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinkedService() DataFactoryDataFlowSinkLinkedServiceOutputReference
	LinkedServiceInput() *DataFactoryDataFlowSinkLinkedService
	Name() *string
	SetName(val *string)
	NameInput() *string
	RejectedLinkedService() DataFactoryDataFlowSinkRejectedLinkedServiceOutputReference
	RejectedLinkedServiceInput() *DataFactoryDataFlowSinkRejectedLinkedService
	SchemaLinkedService() DataFactoryDataFlowSinkSchemaLinkedServiceOutputReference
	SchemaLinkedServiceInput() *DataFactoryDataFlowSinkSchemaLinkedService
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
	PutDataset(value *DataFactoryDataFlowSinkDataset)
	PutFlowlet(value *DataFactoryDataFlowSinkFlowlet)
	PutLinkedService(value *DataFactoryDataFlowSinkLinkedService)
	PutRejectedLinkedService(value *DataFactoryDataFlowSinkRejectedLinkedService)
	PutSchemaLinkedService(value *DataFactoryDataFlowSinkSchemaLinkedService)
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

// The jsii proxy struct for DataFactoryDataFlowSinkOutputReference
type jsiiProxy_DataFactoryDataFlowSinkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) Dataset() DataFactoryDataFlowSinkDatasetOutputReference {
	var returns DataFactoryDataFlowSinkDatasetOutputReference
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) DatasetInput() *DataFactoryDataFlowSinkDataset {
	var returns *DataFactoryDataFlowSinkDataset
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) Flowlet() DataFactoryDataFlowSinkFlowletOutputReference {
	var returns DataFactoryDataFlowSinkFlowletOutputReference
	_jsii_.Get(
		j,
		"flowlet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) FlowletInput() *DataFactoryDataFlowSinkFlowlet {
	var returns *DataFactoryDataFlowSinkFlowlet
	_jsii_.Get(
		j,
		"flowletInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) LinkedService() DataFactoryDataFlowSinkLinkedServiceOutputReference {
	var returns DataFactoryDataFlowSinkLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"linkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) LinkedServiceInput() *DataFactoryDataFlowSinkLinkedService {
	var returns *DataFactoryDataFlowSinkLinkedService
	_jsii_.Get(
		j,
		"linkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) RejectedLinkedService() DataFactoryDataFlowSinkRejectedLinkedServiceOutputReference {
	var returns DataFactoryDataFlowSinkRejectedLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"rejectedLinkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) RejectedLinkedServiceInput() *DataFactoryDataFlowSinkRejectedLinkedService {
	var returns *DataFactoryDataFlowSinkRejectedLinkedService
	_jsii_.Get(
		j,
		"rejectedLinkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) SchemaLinkedService() DataFactoryDataFlowSinkSchemaLinkedServiceOutputReference {
	var returns DataFactoryDataFlowSinkSchemaLinkedServiceOutputReference
	_jsii_.Get(
		j,
		"schemaLinkedService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) SchemaLinkedServiceInput() *DataFactoryDataFlowSinkSchemaLinkedService {
	var returns *DataFactoryDataFlowSinkSchemaLinkedService
	_jsii_.Get(
		j,
		"schemaLinkedServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryDataFlowSinkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataFactoryDataFlowSinkOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryDataFlowSinkOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryDataFlowSinkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDataFlow.DataFactoryDataFlowSinkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataFactoryDataFlowSinkOutputReference_Override(d DataFactoryDataFlowSinkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDataFlow.DataFactoryDataFlowSinkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDataFlowSinkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) PutDataset(value *DataFactoryDataFlowSinkDataset) {
	if err := d.validatePutDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataset",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) PutFlowlet(value *DataFactoryDataFlowSinkFlowlet) {
	if err := d.validatePutFlowletParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlowlet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) PutLinkedService(value *DataFactoryDataFlowSinkLinkedService) {
	if err := d.validatePutLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) PutRejectedLinkedService(value *DataFactoryDataFlowSinkRejectedLinkedService) {
	if err := d.validatePutRejectedLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRejectedLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) PutSchemaLinkedService(value *DataFactoryDataFlowSinkSchemaLinkedService) {
	if err := d.validatePutSchemaLinkedServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchemaLinkedService",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ResetDataset() {
	_jsii_.InvokeVoid(
		d,
		"resetDataset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ResetFlowlet() {
	_jsii_.InvokeVoid(
		d,
		"resetFlowlet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ResetLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ResetRejectedLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetRejectedLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ResetSchemaLinkedService() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaLinkedService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataFactoryDataFlowSinkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

