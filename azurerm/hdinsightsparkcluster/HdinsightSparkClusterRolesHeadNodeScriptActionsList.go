// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightsparkcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/hdinsightsparkcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightSparkClusterRolesHeadNodeScriptActionsList interface {
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
	Get(index *float64) HdinsightSparkClusterRolesHeadNodeScriptActionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightSparkClusterRolesHeadNodeScriptActionsList
type jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightSparkClusterRolesHeadNodeScriptActionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightSparkClusterRolesHeadNodeScriptActionsList {
	_init_.Initialize()

	if err := validateNewHdinsightSparkClusterRolesHeadNodeScriptActionsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightSparkCluster.HdinsightSparkClusterRolesHeadNodeScriptActionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightSparkClusterRolesHeadNodeScriptActionsList_Override(h HdinsightSparkClusterRolesHeadNodeScriptActionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightSparkCluster.HdinsightSparkClusterRolesHeadNodeScriptActionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) Get(index *float64) HdinsightSparkClusterRolesHeadNodeScriptActionsOutputReference {
	if err := h.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns HdinsightSparkClusterRolesHeadNodeScriptActionsOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightSparkClusterRolesHeadNodeScriptActionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

