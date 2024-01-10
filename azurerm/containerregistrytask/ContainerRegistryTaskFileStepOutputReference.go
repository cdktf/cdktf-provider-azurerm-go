// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/containerregistrytask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryTaskFileStepOutputReference interface {
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
	InternalValue() *ContainerRegistryTaskFileStep
	SetInternalValue(val *ContainerRegistryTaskFileStep)
	SecretValues() *map[string]*string
	SetSecretValues(val *map[string]*string)
	SecretValuesInput() *map[string]*string
	TaskFilePath() *string
	SetTaskFilePath(val *string)
	TaskFilePathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValueFilePath() *string
	SetValueFilePath(val *string)
	ValueFilePathInput() *string
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
	ResetValueFilePath()
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

// The jsii proxy struct for ContainerRegistryTaskFileStepOutputReference
type jsiiProxy_ContainerRegistryTaskFileStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ContextAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ContextAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ContextPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ContextPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) InternalValue() *ContainerRegistryTaskFileStep {
	var returns *ContainerRegistryTaskFileStep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SecretValues() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SecretValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) TaskFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) TaskFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ValueFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ValueFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) Values() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskFileStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskFileStepOutputReference {
	_init_.Initialize()

	if err := validateNewContainerRegistryTaskFileStepOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerRegistryTaskFileStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskFileStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskFileStepOutputReference_Override(c ContainerRegistryTaskFileStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskFileStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetContextAccessToken(val *string) {
	if err := j.validateSetContextAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextAccessToken",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetContextPath(val *string) {
	if err := j.validateSetContextPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextPath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetInternalValue(val *ContainerRegistryTaskFileStep) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetSecretValues(val *map[string]*string) {
	if err := j.validateSetSecretValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretValues",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetTaskFilePath(val *string) {
	if err := j.validateSetTaskFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskFilePath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetValueFilePath(val *string) {
	if err := j.validateSetValueFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueFilePath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference)SetValues(val *map[string]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetContextAccessToken() {
	_jsii_.InvokeVoid(
		c,
		"resetContextAccessToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetContextPath() {
	_jsii_.InvokeVoid(
		c,
		"resetContextPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetSecretValues() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetValueFilePath() {
	_jsii_.InvokeVoid(
		c,
		"resetValueFilePath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		c,
		"resetValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

