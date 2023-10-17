// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertruleanomalybuiltin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/sentinelalertruleanomalybuiltin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList interface {
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
	Get(index *float64) SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList
type jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleAnomalyBuiltInRequiredDataConnectorListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyBuiltIn.SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList_Override(s SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyBuiltIn.SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) Get(index *float64) SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyBuiltInRequiredDataConnectorList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

