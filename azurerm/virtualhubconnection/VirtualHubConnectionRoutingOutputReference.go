// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualhubconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/virtualhubconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualHubConnectionRoutingOutputReference interface {
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
	InternalValue() *VirtualHubConnectionRouting
	SetInternalValue(val *VirtualHubConnectionRouting)
	OutboundRouteMapId() *string
	SetOutboundRouteMapId(val *string)
	OutboundRouteMapIdInput() *string
	PropagatedRouteTable() VirtualHubConnectionRoutingPropagatedRouteTableOutputReference
	PropagatedRouteTableInput() *VirtualHubConnectionRoutingPropagatedRouteTable
	StaticVnetLocalRouteOverrideCriteria() *string
	SetStaticVnetLocalRouteOverrideCriteria(val *string)
	StaticVnetLocalRouteOverrideCriteriaInput() *string
	StaticVnetPropagateStaticRoutesEnabled() interface{}
	SetStaticVnetPropagateStaticRoutesEnabled(val interface{})
	StaticVnetPropagateStaticRoutesEnabledInput() interface{}
	StaticVnetRoute() VirtualHubConnectionRoutingStaticVnetRouteList
	StaticVnetRouteInput() interface{}
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
	PutPropagatedRouteTable(value *VirtualHubConnectionRoutingPropagatedRouteTable)
	PutStaticVnetRoute(value interface{})
	ResetAssociatedRouteTableId()
	ResetInboundRouteMapId()
	ResetOutboundRouteMapId()
	ResetPropagatedRouteTable()
	ResetStaticVnetLocalRouteOverrideCriteria()
	ResetStaticVnetPropagateStaticRoutesEnabled()
	ResetStaticVnetRoute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualHubConnectionRoutingOutputReference
type jsiiProxy_VirtualHubConnectionRoutingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) AssociatedRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) AssociatedRouteTableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) InboundRouteMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundRouteMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) InboundRouteMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundRouteMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) InternalValue() *VirtualHubConnectionRouting {
	var returns *VirtualHubConnectionRouting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) OutboundRouteMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundRouteMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) OutboundRouteMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundRouteMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) PropagatedRouteTable() VirtualHubConnectionRoutingPropagatedRouteTableOutputReference {
	var returns VirtualHubConnectionRoutingPropagatedRouteTableOutputReference
	_jsii_.Get(
		j,
		"propagatedRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) PropagatedRouteTableInput() *VirtualHubConnectionRoutingPropagatedRouteTable {
	var returns *VirtualHubConnectionRoutingPropagatedRouteTable
	_jsii_.Get(
		j,
		"propagatedRouteTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) StaticVnetLocalRouteOverrideCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticVnetLocalRouteOverrideCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) StaticVnetLocalRouteOverrideCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticVnetLocalRouteOverrideCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) StaticVnetPropagateStaticRoutesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticVnetPropagateStaticRoutesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) StaticVnetPropagateStaticRoutesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticVnetPropagateStaticRoutesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) StaticVnetRoute() VirtualHubConnectionRoutingStaticVnetRouteList {
	var returns VirtualHubConnectionRoutingStaticVnetRouteList
	_jsii_.Get(
		j,
		"staticVnetRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) StaticVnetRouteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticVnetRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualHubConnectionRoutingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualHubConnectionRoutingOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualHubConnectionRoutingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualHubConnectionRoutingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualHubConnection.VirtualHubConnectionRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualHubConnectionRoutingOutputReference_Override(v VirtualHubConnectionRoutingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualHubConnection.VirtualHubConnectionRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetAssociatedRouteTableId(val *string) {
	if err := j.validateSetAssociatedRouteTableIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedRouteTableId",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetInboundRouteMapId(val *string) {
	if err := j.validateSetInboundRouteMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundRouteMapId",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetInternalValue(val *VirtualHubConnectionRouting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetOutboundRouteMapId(val *string) {
	if err := j.validateSetOutboundRouteMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundRouteMapId",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetStaticVnetLocalRouteOverrideCriteria(val *string) {
	if err := j.validateSetStaticVnetLocalRouteOverrideCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticVnetLocalRouteOverrideCriteria",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetStaticVnetPropagateStaticRoutesEnabled(val interface{}) {
	if err := j.validateSetStaticVnetPropagateStaticRoutesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticVnetPropagateStaticRoutesEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualHubConnectionRoutingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) PutPropagatedRouteTable(value *VirtualHubConnectionRoutingPropagatedRouteTable) {
	if err := v.validatePutPropagatedRouteTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPropagatedRouteTable",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) PutStaticVnetRoute(value interface{}) {
	if err := v.validatePutStaticVnetRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStaticVnetRoute",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ResetAssociatedRouteTableId() {
	_jsii_.InvokeVoid(
		v,
		"resetAssociatedRouteTableId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ResetInboundRouteMapId() {
	_jsii_.InvokeVoid(
		v,
		"resetInboundRouteMapId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ResetOutboundRouteMapId() {
	_jsii_.InvokeVoid(
		v,
		"resetOutboundRouteMapId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ResetPropagatedRouteTable() {
	_jsii_.InvokeVoid(
		v,
		"resetPropagatedRouteTable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ResetStaticVnetLocalRouteOverrideCriteria() {
	_jsii_.InvokeVoid(
		v,
		"resetStaticVnetLocalRouteOverrideCriteria",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ResetStaticVnetPropagateStaticRoutesEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetStaticVnetPropagateStaticRoutesEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ResetStaticVnetRoute() {
	_jsii_.InvokeVoid(
		v,
		"resetStaticVnetRoute",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualHubConnectionRoutingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

