// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privateendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/privateendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivateEndpointPrivateServiceConnectionOutputReference interface {
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
	InternalValue() *PrivateEndpointPrivateServiceConnection
	SetInternalValue(val *PrivateEndpointPrivateServiceConnection)
	IsManualConnection() interface{}
	SetIsManualConnection(val interface{})
	IsManualConnectionInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PrivateConnectionResourceAlias() *string
	SetPrivateConnectionResourceAlias(val *string)
	PrivateConnectionResourceAliasInput() *string
	PrivateConnectionResourceId() *string
	SetPrivateConnectionResourceId(val *string)
	PrivateConnectionResourceIdInput() *string
	PrivateIpAddress() *string
	RequestMessage() *string
	SetRequestMessage(val *string)
	RequestMessageInput() *string
	SubresourceNames() *[]*string
	SetSubresourceNames(val *[]*string)
	SubresourceNamesInput() *[]*string
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
	ResetPrivateConnectionResourceAlias()
	ResetPrivateConnectionResourceId()
	ResetRequestMessage()
	ResetSubresourceNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivateEndpointPrivateServiceConnectionOutputReference
type jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) InternalValue() *PrivateEndpointPrivateServiceConnection {
	var returns *PrivateEndpointPrivateServiceConnection
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) IsManualConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isManualConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) IsManualConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isManualConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) PrivateConnectionResourceAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateConnectionResourceAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) PrivateConnectionResourceAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateConnectionResourceAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) PrivateConnectionResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateConnectionResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) PrivateConnectionResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateConnectionResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) RequestMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) RequestMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) SubresourceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subresourceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) SubresourceNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subresourceNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivateEndpointPrivateServiceConnectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivateEndpointPrivateServiceConnectionOutputReference {
	_init_.Initialize()

	if err := validateNewPrivateEndpointPrivateServiceConnectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.privateEndpoint.PrivateEndpointPrivateServiceConnectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivateEndpointPrivateServiceConnectionOutputReference_Override(p PrivateEndpointPrivateServiceConnectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.privateEndpoint.PrivateEndpointPrivateServiceConnectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetInternalValue(val *PrivateEndpointPrivateServiceConnection) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetIsManualConnection(val interface{}) {
	if err := j.validateSetIsManualConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isManualConnection",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetPrivateConnectionResourceAlias(val *string) {
	if err := j.validateSetPrivateConnectionResourceAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateConnectionResourceAlias",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetPrivateConnectionResourceId(val *string) {
	if err := j.validateSetPrivateConnectionResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateConnectionResourceId",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetRequestMessage(val *string) {
	if err := j.validateSetRequestMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMessage",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetSubresourceNames(val *[]*string) {
	if err := j.validateSetSubresourceNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subresourceNames",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) ResetPrivateConnectionResourceAlias() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateConnectionResourceAlias",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) ResetPrivateConnectionResourceId() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateConnectionResourceId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) ResetRequestMessage() {
	_jsii_.InvokeVoid(
		p,
		"resetRequestMessage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) ResetSubresourceNames() {
	_jsii_.InvokeVoid(
		p,
		"resetSubresourceNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEndpointPrivateServiceConnectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

