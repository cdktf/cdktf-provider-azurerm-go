// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimeazuressis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/datafactoryintegrationruntimeazuressis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference interface {
	cdktf.ComplexObject
	CommandKey() DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupCommandKeyList
	CommandKeyInput() interface{}
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
	Component() DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupComponentList
	ComponentInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Environment() *map[string]*string
	SetEnvironment(val *map[string]*string)
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup
	SetInternalValue(val *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup)
	PowershellVersion() *string
	SetPowershellVersion(val *string)
	PowershellVersionInput() *string
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
	PutCommandKey(value interface{})
	PutComponent(value interface{})
	ResetCommandKey()
	ResetComponent()
	ResetEnvironment()
	ResetPowershellVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference
type jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) CommandKey() DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupCommandKeyList {
	var returns DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupCommandKeyList
	_jsii_.Get(
		j,
		"commandKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) CommandKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commandKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) Component() DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupComponentList {
	var returns DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupComponentList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ComponentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) InternalValue() *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup {
	var returns *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) PowershellVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powershellVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) PowershellVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powershellVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference_Override(d DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference)SetInternalValue(val *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference)SetPowershellVersion(val *string) {
	if err := j.validateSetPowershellVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powershellVersion",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) PutCommandKey(value interface{}) {
	if err := d.validatePutCommandKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCommandKey",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) PutComponent(value interface{}) {
	if err := d.validatePutComponentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putComponent",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ResetCommandKey() {
	_jsii_.InvokeVoid(
		d,
		"resetCommandKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ResetComponent() {
	_jsii_.InvokeVoid(
		d,
		"resetComponent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ResetPowershellVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPowershellVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

