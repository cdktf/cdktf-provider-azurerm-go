// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationsoftwareupdateconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/automationsoftwareupdateconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList
type jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTargetNonAzureQueryList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList {
	_init_.Initialize()

	if err := validateNewAutomationSoftwareUpdateConfigurationTargetNonAzureQueryListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTargetNonAzureQueryList_Override(a AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) Get(index *float64) AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

