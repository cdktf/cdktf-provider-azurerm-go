// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngatewayconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/vpngatewayconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnGatewayConnectionRoutingOutputReference interface {
	cdktf.ComplexObject
	AssociatedRouteTable() *string
	SetAssociatedRouteTable(val *string)
	AssociatedRouteTableInput() *string
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
	InternalValue() *VpnGatewayConnectionRouting
	SetInternalValue(val *VpnGatewayConnectionRouting)
	OutboundRouteMapId() *string
	SetOutboundRouteMapId(val *string)
	OutboundRouteMapIdInput() *string
	PropagatedRouteTable() VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference
	PropagatedRouteTableInput() *VpnGatewayConnectionRoutingPropagatedRouteTable
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
	PutPropagatedRouteTable(value *VpnGatewayConnectionRoutingPropagatedRouteTable)
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

// The jsii proxy struct for VpnGatewayConnectionRoutingOutputReference
type jsiiProxy_VpnGatewayConnectionRoutingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) AssociatedRouteTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) AssociatedRouteTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) InboundRouteMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundRouteMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) InboundRouteMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundRouteMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) InternalValue() *VpnGatewayConnectionRouting {
	var returns *VpnGatewayConnectionRouting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) OutboundRouteMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundRouteMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) OutboundRouteMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundRouteMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) PropagatedRouteTable() VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference {
	var returns VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference
	_jsii_.Get(
		j,
		"propagatedRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) PropagatedRouteTableInput() *VpnGatewayConnectionRoutingPropagatedRouteTable {
	var returns *VpnGatewayConnectionRoutingPropagatedRouteTable
	_jsii_.Get(
		j,
		"propagatedRouteTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionRoutingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnGatewayConnectionRoutingOutputReference {
	_init_.Initialize()

	if err := validateNewVpnGatewayConnectionRoutingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnGatewayConnectionRoutingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionRoutingOutputReference_Override(v VpnGatewayConnectionRoutingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference)SetAssociatedRouteTable(val *string) {
	if err := j.validateSetAssociatedRouteTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedRouteTable",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference)SetInboundRouteMapId(val *string) {
	if err := j.validateSetInboundRouteMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundRouteMapId",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference)SetInternalValue(val *VpnGatewayConnectionRouting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference)SetOutboundRouteMapId(val *string) {
	if err := j.validateSetOutboundRouteMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundRouteMapId",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) PutPropagatedRouteTable(value *VpnGatewayConnectionRoutingPropagatedRouteTable) {
	if err := v.validatePutPropagatedRouteTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPropagatedRouteTable",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ResetInboundRouteMapId() {
	_jsii_.InvokeVoid(
		v,
		"resetInboundRouteMapId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ResetOutboundRouteMapId() {
	_jsii_.InvokeVoid(
		v,
		"resetOutboundRouteMapId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ResetPropagatedRouteTable() {
	_jsii_.InvokeVoid(
		v,
		"resetPropagatedRouteTable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

