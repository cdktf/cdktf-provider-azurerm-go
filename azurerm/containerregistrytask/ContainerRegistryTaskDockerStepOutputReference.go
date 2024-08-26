// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/containerregistrytask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryTaskDockerStepOutputReference interface {
	cdktf.ComplexObject
	Arguments() *map[string]*string
	SetArguments(val *map[string]*string)
	ArgumentsInput() *map[string]*string
	CacheEnabled() interface{}
	SetCacheEnabled(val interface{})
	CacheEnabledInput() interface{}
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
	DockerfilePath() *string
	SetDockerfilePath(val *string)
	DockerfilePathInput() *string
	// Experimental.
	Fqn() *string
	ImageNames() *[]*string
	SetImageNames(val *[]*string)
	ImageNamesInput() *[]*string
	InternalValue() *ContainerRegistryTaskDockerStep
	SetInternalValue(val *ContainerRegistryTaskDockerStep)
	PushEnabled() interface{}
	SetPushEnabled(val interface{})
	PushEnabledInput() interface{}
	SecretArguments() *map[string]*string
	SetSecretArguments(val *map[string]*string)
	SecretArgumentsInput() *map[string]*string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	ResetArguments()
	ResetCacheEnabled()
	ResetImageNames()
	ResetPushEnabled()
	ResetSecretArguments()
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskDockerStepOutputReference
type jsiiProxy_ContainerRegistryTaskDockerStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) Arguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) CacheEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) CacheEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ContextAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ContextAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ContextPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ContextPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) DockerfilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) DockerfilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ImageNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ImageNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) InternalValue() *ContainerRegistryTaskDockerStep {
	var returns *ContainerRegistryTaskDockerStep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) PushEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) PushEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SecretArguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SecretArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskDockerStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskDockerStepOutputReference {
	_init_.Initialize()

	if err := validateNewContainerRegistryTaskDockerStepOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerRegistryTaskDockerStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskDockerStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskDockerStepOutputReference_Override(c ContainerRegistryTaskDockerStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskDockerStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetArguments(val *map[string]*string) {
	if err := j.validateSetArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetCacheEnabled(val interface{}) {
	if err := j.validateSetCacheEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetContextAccessToken(val *string) {
	if err := j.validateSetContextAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextAccessToken",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetContextPath(val *string) {
	if err := j.validateSetContextPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextPath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetDockerfilePath(val *string) {
	if err := j.validateSetDockerfilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfilePath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetImageNames(val *[]*string) {
	if err := j.validateSetImageNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageNames",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetInternalValue(val *ContainerRegistryTaskDockerStep) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetPushEnabled(val interface{}) {
	if err := j.validateSetPushEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetSecretArguments(val *map[string]*string) {
	if err := j.validateSetSecretArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArguments",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetArguments() {
	_jsii_.InvokeVoid(
		c,
		"resetArguments",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetCacheEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetImageNames() {
	_jsii_.InvokeVoid(
		c,
		"resetImageNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetPushEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPushEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetSecretArguments() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretArguments",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

