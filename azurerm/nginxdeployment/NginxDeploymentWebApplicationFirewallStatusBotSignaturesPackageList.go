// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/nginxdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList interface {
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
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList
type jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewNginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList {
	_init_.Initialize()

	if err := validateNewNginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList_Override(n NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := n.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		n,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) Get(index *float64) NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageOutputReference {
	if err := n.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusBotSignaturesPackageList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

