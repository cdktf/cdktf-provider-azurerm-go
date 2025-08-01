// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermoracledbsystemshapes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dataazurermoracledbsystemshapes/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference interface {
	cdktf.ComplexObject
	AreServerTypesSupported() cdktf.IResolvable
	AvailableCoreCount() *float64
	AvailableCoreCountPerNode() *float64
	AvailableDataStorageInTbs() *float64
	AvailableDataStoragePerServerInTbs() *float64
	AvailableDbNodePerNodeInGbs() *float64
	AvailableDbNodeStorageInGbs() *float64
	AvailableMemoryInGbs() *float64
	AvailableMemoryPerNodeInGbs() *float64
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
	ComputeModel() *string
	CoreCountIncrement() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisplayName() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAzurermOracleDbSystemShapesDbSystemShapes
	SetInternalValue(val *DataAzurermOracleDbSystemShapesDbSystemShapes)
	MaximumNodeCount() *float64
	MaximumStorageCount() *float64
	MinimumCoreCount() *float64
	MinimumCoreCountPerNode() *float64
	MinimumDataStorageInTbs() *float64
	MinimumDbNodeStoragePerNodeInGbs() *float64
	MinimumMemoryPerNodeInGbs() *float64
	MinimumNodeCount() *float64
	MinimumStorageCount() *float64
	RuntimeMinimumCoreCount() *float64
	ShapeFamily() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference
type jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AreServerTypesSupported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"areServerTypesSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AvailableCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AvailableCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AvailableDataStorageInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDataStorageInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AvailableDataStoragePerServerInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDataStoragePerServerInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AvailableDbNodePerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDbNodePerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AvailableDbNodeStorageInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDbNodeStorageInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AvailableMemoryInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) AvailableMemoryPerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryPerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) CoreCountIncrement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreCountIncrement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) InternalValue() *DataAzurermOracleDbSystemShapesDbSystemShapes {
	var returns *DataAzurermOracleDbSystemShapesDbSystemShapes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MaximumNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MaximumStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MinimumCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MinimumCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MinimumDataStorageInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumDataStorageInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MinimumDbNodeStoragePerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumDbNodeStoragePerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MinimumMemoryPerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumMemoryPerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MinimumNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) MinimumStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) RuntimeMinimumCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runtimeMinimumCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) ShapeFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAzurermOracleDbSystemShapesDbSystemShapesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermOracleDbSystemShapesDbSystemShapesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermOracleDbSystemShapes.DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermOracleDbSystemShapesDbSystemShapesOutputReference_Override(d DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermOracleDbSystemShapes.DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference)SetInternalValue(val *DataAzurermOracleDbSystemShapesDbSystemShapes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermOracleDbSystemShapesDbSystemShapesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

