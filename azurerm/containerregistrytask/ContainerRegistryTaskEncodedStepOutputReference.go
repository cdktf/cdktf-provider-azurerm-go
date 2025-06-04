// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/containerregistrytask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryTaskEncodedStepOutputReference interface {
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
	ContextAccessToken() *string
	SetContextAccessToken(val *string)
	ContextAccessTokenInput() *string
	ContextPath() *string
	SetContextPath(val *string)
	ContextPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerRegistryTaskEncodedStep
	SetInternalValue(val *ContainerRegistryTaskEncodedStep)
	SecretValues() *map[string]*string
	SetSecretValues(val *map[string]*string)
	SecretValuesInput() *map[string]*string
	TaskContent() *string
	SetTaskContent(val *string)
	TaskContentInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValueContent() *string
	SetValueContent(val *string)
	ValueContentInput() *string
	Values() *map[string]*string
	SetValues(val *map[string]*string)
	ValuesInput() *map[string]*string
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
	ResetContextAccessToken()
	ResetContextPath()
	ResetSecretValues()
	ResetValueContent()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskEncodedStepOutputReference
type jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ContextAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ContextAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ContextPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ContextPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) InternalValue() *ContainerRegistryTaskEncodedStep {
	var returns *ContainerRegistryTaskEncodedStep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SecretValues() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SecretValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) TaskContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) TaskContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ValueContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ValueContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) Values() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskEncodedStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskEncodedStepOutputReference {
	_init_.Initialize()

	if err := validateNewContainerRegistryTaskEncodedStepOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskEncodedStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskEncodedStepOutputReference_Override(c ContainerRegistryTaskEncodedStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskEncodedStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetContextAccessToken(val *string) {
	if err := j.validateSetContextAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextAccessToken",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetContextPath(val *string) {
	if err := j.validateSetContextPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextPath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetInternalValue(val *ContainerRegistryTaskEncodedStep) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetSecretValues(val *map[string]*string) {
	if err := j.validateSetSecretValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretValues",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetTaskContent(val *string) {
	if err := j.validateSetTaskContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskContent",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetValueContent(val *string) {
	if err := j.validateSetValueContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueContent",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference)SetValues(val *map[string]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetContextAccessToken() {
	_jsii_.InvokeVoid(
		c,
		"resetContextAccessToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetContextPath() {
	_jsii_.InvokeVoid(
		c,
		"resetContextPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetSecretValues() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetValueContent() {
	_jsii_.InvokeVoid(
		c,
		"resetValueContent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		c,
		"resetValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

