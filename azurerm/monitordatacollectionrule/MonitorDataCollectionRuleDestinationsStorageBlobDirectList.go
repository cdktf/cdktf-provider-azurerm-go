// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/monitordatacollectionrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorDataCollectionRuleDestinationsStorageBlobDirectList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) MonitorDataCollectionRuleDestinationsStorageBlobDirectOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorDataCollectionRuleDestinationsStorageBlobDirectList
type jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorDataCollectionRuleDestinationsStorageBlobDirectList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorDataCollectionRuleDestinationsStorageBlobDirectList {
	_init_.Initialize()

	if err := validateNewMonitorDataCollectionRuleDestinationsStorageBlobDirectListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRuleDestinationsStorageBlobDirectList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorDataCollectionRuleDestinationsStorageBlobDirectList_Override(m MonitorDataCollectionRuleDestinationsStorageBlobDirectList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRuleDestinationsStorageBlobDirectList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) Get(index *float64) MonitorDataCollectionRuleDestinationsStorageBlobDirectOutputReference {
	if err := m.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns MonitorDataCollectionRuleDestinationsStorageBlobDirectOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsStorageBlobDirectList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

