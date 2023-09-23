// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/containerapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppTemplateOutputReference interface {
	cdktf.ComplexObject
	AzureQueueScaleRule() ContainerAppTemplateAzureQueueScaleRuleList
	AzureQueueScaleRuleInput() interface{}
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
	Container() ContainerAppTemplateContainerList
	ContainerInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomScaleRule() ContainerAppTemplateCustomScaleRuleList
	CustomScaleRuleInput() interface{}
	// Experimental.
	Fqn() *string
	HttpScaleRule() ContainerAppTemplateHttpScaleRuleList
	HttpScaleRuleInput() interface{}
	InternalValue() *ContainerAppTemplate
	SetInternalValue(val *ContainerAppTemplate)
	MaxReplicas() *float64
	SetMaxReplicas(val *float64)
	MaxReplicasInput() *float64
	MinReplicas() *float64
	SetMinReplicas(val *float64)
	MinReplicasInput() *float64
	RevisionSuffix() *string
	SetRevisionSuffix(val *string)
	RevisionSuffixInput() *string
	TcpScaleRule() ContainerAppTemplateTcpScaleRuleList
	TcpScaleRuleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volume() ContainerAppTemplateVolumeList
	VolumeInput() interface{}
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
	PutAzureQueueScaleRule(value interface{})
	PutContainer(value interface{})
	PutCustomScaleRule(value interface{})
	PutHttpScaleRule(value interface{})
	PutTcpScaleRule(value interface{})
	PutVolume(value interface{})
	ResetAzureQueueScaleRule()
	ResetCustomScaleRule()
	ResetHttpScaleRule()
	ResetMaxReplicas()
	ResetMinReplicas()
	ResetRevisionSuffix()
	ResetTcpScaleRule()
	ResetVolume()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerAppTemplateOutputReference
type jsiiProxy_ContainerAppTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) AzureQueueScaleRule() ContainerAppTemplateAzureQueueScaleRuleList {
	var returns ContainerAppTemplateAzureQueueScaleRuleList
	_jsii_.Get(
		j,
		"azureQueueScaleRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) AzureQueueScaleRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureQueueScaleRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) Container() ContainerAppTemplateContainerList {
	var returns ContainerAppTemplateContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) ContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) CustomScaleRule() ContainerAppTemplateCustomScaleRuleList {
	var returns ContainerAppTemplateCustomScaleRuleList
	_jsii_.Get(
		j,
		"customScaleRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) CustomScaleRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customScaleRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) HttpScaleRule() ContainerAppTemplateHttpScaleRuleList {
	var returns ContainerAppTemplateHttpScaleRuleList
	_jsii_.Get(
		j,
		"httpScaleRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) HttpScaleRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpScaleRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) InternalValue() *ContainerAppTemplate {
	var returns *ContainerAppTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) MaxReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) MaxReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) MinReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) MinReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) RevisionSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) RevisionSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) TcpScaleRule() ContainerAppTemplateTcpScaleRuleList {
	var returns ContainerAppTemplateTcpScaleRuleList
	_jsii_.Get(
		j,
		"tcpScaleRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) TcpScaleRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpScaleRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) Volume() ContainerAppTemplateVolumeList {
	var returns ContainerAppTemplateVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference) VolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


func NewContainerAppTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerAppTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewContainerAppTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAppTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerApp.ContainerAppTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerAppTemplateOutputReference_Override(c ContainerAppTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerApp.ContainerAppTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference)SetInternalValue(val *ContainerAppTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference)SetMaxReplicas(val *float64) {
	if err := j.validateSetMaxReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReplicas",
		val,
	)
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference)SetMinReplicas(val *float64) {
	if err := j.validateSetMinReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicas",
		val,
	)
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference)SetRevisionSuffix(val *string) {
	if err := j.validateSetRevisionSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionSuffix",
		val,
	)
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerAppTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) PutAzureQueueScaleRule(value interface{}) {
	if err := c.validatePutAzureQueueScaleRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureQueueScaleRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) PutContainer(value interface{}) {
	if err := c.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putContainer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) PutCustomScaleRule(value interface{}) {
	if err := c.validatePutCustomScaleRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomScaleRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) PutHttpScaleRule(value interface{}) {
	if err := c.validatePutHttpScaleRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpScaleRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) PutTcpScaleRule(value interface{}) {
	if err := c.validatePutTcpScaleRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTcpScaleRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) PutVolume(value interface{}) {
	if err := c.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ResetAzureQueueScaleRule() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureQueueScaleRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ResetCustomScaleRule() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomScaleRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ResetHttpScaleRule() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpScaleRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ResetMaxReplicas() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxReplicas",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ResetMinReplicas() {
	_jsii_.InvokeVoid(
		c,
		"resetMinReplicas",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ResetRevisionSuffix() {
	_jsii_.InvokeVoid(
		c,
		"resetRevisionSuffix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ResetTcpScaleRule() {
	_jsii_.InvokeVoid(
		c,
		"resetTcpScaleRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerAppTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

