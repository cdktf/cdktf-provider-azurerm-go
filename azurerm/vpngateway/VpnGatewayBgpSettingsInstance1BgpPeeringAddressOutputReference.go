// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/vpngateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference interface {
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
	CustomIps() *[]*string
	SetCustomIps(val *[]*string)
	CustomIpsInput() *[]*string
	DefaultIps() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *VpnGatewayBgpSettingsInstance1BgpPeeringAddress
	SetInternalValue(val *VpnGatewayBgpSettingsInstance1BgpPeeringAddress)
	IpConfigurationId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TunnelIps() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference
type jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) CustomIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) CustomIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) DefaultIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) InternalValue() *VpnGatewayBgpSettingsInstance1BgpPeeringAddress {
	var returns *VpnGatewayBgpSettingsInstance1BgpPeeringAddress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) IpConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) TunnelIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelIps",
		&returns,
	)
	return returns
}


func NewVpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference {
	_init_.Initialize()

	if err := validateNewVpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGateway.VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference_Override(v VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGateway.VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference)SetCustomIps(val *[]*string) {
	if err := j.validateSetCustomIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customIps",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference)SetInternalValue(val *VpnGatewayBgpSettingsInstance1BgpPeeringAddress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

