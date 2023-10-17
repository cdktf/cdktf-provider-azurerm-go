// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package expressroutecircuitpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/expressroutecircuitpeering/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExpressRouteCircuitPeeringIpv6OutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ExpressRouteCircuitPeeringIpv6
	SetInternalValue(val *ExpressRouteCircuitPeeringIpv6)
	MicrosoftPeering() ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference
	MicrosoftPeeringInput() *ExpressRouteCircuitPeeringIpv6MicrosoftPeering
	PrimaryPeerAddressPrefix() *string
	SetPrimaryPeerAddressPrefix(val *string)
	PrimaryPeerAddressPrefixInput() *string
	RouteFilterId() *string
	SetRouteFilterId(val *string)
	RouteFilterIdInput() *string
	SecondaryPeerAddressPrefix() *string
	SetSecondaryPeerAddressPrefix(val *string)
	SecondaryPeerAddressPrefixInput() *string
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
	PutMicrosoftPeering(value *ExpressRouteCircuitPeeringIpv6MicrosoftPeering)
	ResetEnabled()
	ResetMicrosoftPeering()
	ResetRouteFilterId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ExpressRouteCircuitPeeringIpv6OutputReference
type jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) InternalValue() *ExpressRouteCircuitPeeringIpv6 {
	var returns *ExpressRouteCircuitPeeringIpv6
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) MicrosoftPeering() ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference {
	var returns ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference
	_jsii_.Get(
		j,
		"microsoftPeering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) MicrosoftPeeringInput() *ExpressRouteCircuitPeeringIpv6MicrosoftPeering {
	var returns *ExpressRouteCircuitPeeringIpv6MicrosoftPeering
	_jsii_.Get(
		j,
		"microsoftPeeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) PrimaryPeerAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPeerAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) PrimaryPeerAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPeerAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) RouteFilterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFilterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) RouteFilterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFilterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SecondaryPeerAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPeerAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SecondaryPeerAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPeerAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExpressRouteCircuitPeeringIpv6OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExpressRouteCircuitPeeringIpv6OutputReference {
	_init_.Initialize()

	if err := validateNewExpressRouteCircuitPeeringIpv6OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringIpv6OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExpressRouteCircuitPeeringIpv6OutputReference_Override(e ExpressRouteCircuitPeeringIpv6OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringIpv6OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetInternalValue(val *ExpressRouteCircuitPeeringIpv6) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetPrimaryPeerAddressPrefix(val *string) {
	if err := j.validateSetPrimaryPeerAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryPeerAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetRouteFilterId(val *string) {
	if err := j.validateSetRouteFilterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeFilterId",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetSecondaryPeerAddressPrefix(val *string) {
	if err := j.validateSetSecondaryPeerAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPeerAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) PutMicrosoftPeering(value *ExpressRouteCircuitPeeringIpv6MicrosoftPeering) {
	if err := e.validatePutMicrosoftPeeringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMicrosoftPeering",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ResetMicrosoftPeering() {
	_jsii_.InvokeVoid(
		e,
		"resetMicrosoftPeering",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ResetRouteFilterId() {
	_jsii_.InvokeVoid(
		e,
		"resetRouteFilterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

