// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermvpngateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/dataazurermvpngateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList
type jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList {
	_init_.Initialize()

	if err := validateNewDataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermVpnGateway.DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList_Override(d DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermVpnGateway.DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) Get(index *float64) DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermVpnGatewayBgpSettingsInstance0BgpPeeringAddressList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

