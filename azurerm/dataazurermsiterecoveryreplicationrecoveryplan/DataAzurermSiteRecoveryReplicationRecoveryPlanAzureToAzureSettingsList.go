// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermsiterecoveryreplicationrecoveryplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/dataazurermsiterecoveryreplicationrecoveryplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList interface {
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
	Get(index *float64) DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList
type jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList {
	_init_.Initialize()

	if err := validateNewDataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermSiteRecoveryReplicationRecoveryPlan.DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList_Override(d DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermSiteRecoveryReplicationRecoveryPlan.DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) Get(index *float64) DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

