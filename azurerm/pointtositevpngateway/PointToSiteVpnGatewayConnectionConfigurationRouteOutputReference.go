// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pointtositevpngateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/pointtositevpngateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference interface {
	cdktf.ComplexObject
	AssociatedRouteTableId() *string
	SetAssociatedRouteTableId(val *string)
	AssociatedRouteTableIdInput() *string
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
	InboundRouteMapId() *string
	SetInboundRouteMapId(val *string)
	InboundRouteMapIdInput() *string
	InternalValue() *PointToSiteVpnGatewayConnectionConfigurationRoute
	SetInternalValue(val *PointToSiteVpnGatewayConnectionConfigurationRoute)
	OutboundRouteMapId() *string
	SetOutboundRouteMapId(val *string)
	OutboundRouteMapIdInput() *string
	PropagatedRouteTable() PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference
	PropagatedRouteTableInput() *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable
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
	PutPropagatedRouteTable(value *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable)
	ResetInboundRouteMapId()
	ResetOutboundRouteMapId()
	ResetPropagatedRouteTable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference
type jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) AssociatedRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) AssociatedRouteTableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) InboundRouteMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundRouteMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) InboundRouteMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundRouteMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) InternalValue() *PointToSiteVpnGatewayConnectionConfigurationRoute {
	var returns *PointToSiteVpnGatewayConnectionConfigurationRoute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) OutboundRouteMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundRouteMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) OutboundRouteMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundRouteMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) PropagatedRouteTable() PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference {
	var returns PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference
	_jsii_.Get(
		j,
		"propagatedRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) PropagatedRouteTableInput() *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable {
	var returns *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable
	_jsii_.Get(
		j,
		"propagatedRouteTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPointToSiteVpnGatewayConnectionConfigurationRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference {
	_init_.Initialize()

	if err := validateNewPointToSiteVpnGatewayConnectionConfigurationRouteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPointToSiteVpnGatewayConnectionConfigurationRouteOutputReference_Override(p PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference)SetAssociatedRouteTableId(val *string) {
	if err := j.validateSetAssociatedRouteTableIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedRouteTableId",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference)SetInboundRouteMapId(val *string) {
	if err := j.validateSetInboundRouteMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundRouteMapId",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference)SetInternalValue(val *PointToSiteVpnGatewayConnectionConfigurationRoute) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference)SetOutboundRouteMapId(val *string) {
	if err := j.validateSetOutboundRouteMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundRouteMapId",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) PutPropagatedRouteTable(value *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable) {
	if err := p.validatePutPropagatedRouteTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPropagatedRouteTable",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ResetInboundRouteMapId() {
	_jsii_.InvokeVoid(
		p,
		"resetInboundRouteMapId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ResetOutboundRouteMapId() {
	_jsii_.InvokeVoid(
		p,
		"resetOutboundRouteMapId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ResetPropagatedRouteTable() {
	_jsii_.InvokeVoid(
		p,
		"resetPropagatedRouteTable",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

