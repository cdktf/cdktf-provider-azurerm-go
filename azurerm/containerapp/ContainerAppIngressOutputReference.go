// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/containerapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppIngressOutputReference interface {
	cdktf.ComplexObject
	AllowInsecureConnections() interface{}
	SetAllowInsecureConnections(val interface{})
	AllowInsecureConnectionsInput() interface{}
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
	CustomDomain() ContainerAppIngressCustomDomainOutputReference
	CustomDomainInput() *ContainerAppIngressCustomDomain
	ExternalEnabled() interface{}
	SetExternalEnabled(val interface{})
	ExternalEnabledInput() interface{}
	Fqdn() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerAppIngress
	SetInternalValue(val *ContainerAppIngress)
	TargetPort() *float64
	SetTargetPort(val *float64)
	TargetPortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficWeight() ContainerAppIngressTrafficWeightList
	TrafficWeightInput() interface{}
	Transport() *string
	SetTransport(val *string)
	TransportInput() *string
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
	PutCustomDomain(value *ContainerAppIngressCustomDomain)
	PutTrafficWeight(value interface{})
	ResetAllowInsecureConnections()
	ResetCustomDomain()
	ResetExternalEnabled()
	ResetTransport()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerAppIngressOutputReference
type jsiiProxy_ContainerAppIngressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) AllowInsecureConnections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) AllowInsecureConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) CustomDomain() ContainerAppIngressCustomDomainOutputReference {
	var returns ContainerAppIngressCustomDomainOutputReference
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) CustomDomainInput() *ContainerAppIngressCustomDomain {
	var returns *ContainerAppIngressCustomDomain
	_jsii_.Get(
		j,
		"customDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) ExternalEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) ExternalEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) InternalValue() *ContainerAppIngress {
	var returns *ContainerAppIngress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) TargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) TargetPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) TrafficWeight() ContainerAppIngressTrafficWeightList {
	var returns ContainerAppIngressTrafficWeightList
	_jsii_.Get(
		j,
		"trafficWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) TrafficWeightInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) Transport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppIngressOutputReference) TransportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportInput",
		&returns,
	)
	return returns
}


func NewContainerAppIngressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerAppIngressOutputReference {
	_init_.Initialize()

	if err := validateNewContainerAppIngressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAppIngressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerApp.ContainerAppIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerAppIngressOutputReference_Override(c ContainerAppIngressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerApp.ContainerAppIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetAllowInsecureConnections(val interface{}) {
	if err := j.validateSetAllowInsecureConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInsecureConnections",
		val,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetExternalEnabled(val interface{}) {
	if err := j.validateSetExternalEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetInternalValue(val *ContainerAppIngress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetTargetPort(val *float64) {
	if err := j.validateSetTargetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPort",
		val,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerAppIngressOutputReference)SetTransport(val *string) {
	if err := j.validateSetTransportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transport",
		val,
	)
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) PutCustomDomain(value *ContainerAppIngressCustomDomain) {
	if err := c.validatePutCustomDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomDomain",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) PutTrafficWeight(value interface{}) {
	if err := c.validatePutTrafficWeightParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrafficWeight",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) ResetAllowInsecureConnections() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowInsecureConnections",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) ResetCustomDomain() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomDomain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) ResetExternalEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) ResetTransport() {
	_jsii_.InvokeVoid(
		c,
		"resetTransport",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppIngressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerAppIngressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

