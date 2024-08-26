// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimeazuressis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/datafactoryintegrationruntimeazuressis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale
	SetInternalValue(val *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale)
	NumberOfExternalNodes() *float64
	SetNumberOfExternalNodes(val *float64)
	NumberOfExternalNodesInput() *float64
	NumberOfPipelineNodes() *float64
	SetNumberOfPipelineNodes(val *float64)
	NumberOfPipelineNodesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeToLive() *float64
	SetTimeToLive(val *float64)
	TimeToLiveInput() *float64
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
	ResetNumberOfExternalNodes()
	ResetNumberOfPipelineNodes()
	ResetTimeToLive()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference
type jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) InternalValue() *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale {
	var returns *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) NumberOfExternalNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfExternalNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) NumberOfExternalNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfExternalNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) NumberOfPipelineNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPipelineNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) NumberOfPipelineNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPipelineNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) TimeToLive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeToLive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) TimeToLiveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeToLiveInput",
		&returns,
	)
	return returns
}


func NewDataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference_Override(d DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference)SetInternalValue(val *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference)SetNumberOfExternalNodes(val *float64) {
	if err := j.validateSetNumberOfExternalNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfExternalNodes",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference)SetNumberOfPipelineNodes(val *float64) {
	if err := j.validateSetNumberOfPipelineNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfPipelineNodes",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference)SetTimeToLive(val *float64) {
	if err := j.validateSetTimeToLiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToLive",
		val,
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) ResetNumberOfExternalNodes() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfExternalNodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) ResetNumberOfPipelineNodes() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfPipelineNodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) ResetTimeToLive() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeToLive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

