// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/sentinelalertrulescheduled/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleScheduledIncidentOutputReference interface {
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
	CreateIncidentEnabled() interface{}
	SetCreateIncidentEnabled(val interface{})
	CreateIncidentEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Grouping() SentinelAlertRuleScheduledIncidentGroupingOutputReference
	GroupingInput() *SentinelAlertRuleScheduledIncidentGrouping
	InternalValue() *SentinelAlertRuleScheduledIncident
	SetInternalValue(val *SentinelAlertRuleScheduledIncident)
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
	PutGrouping(value *SentinelAlertRuleScheduledIncidentGrouping)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledIncidentOutputReference
type jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) CreateIncidentEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIncidentEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) CreateIncidentEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIncidentEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) Grouping() SentinelAlertRuleScheduledIncidentGroupingOutputReference {
	var returns SentinelAlertRuleScheduledIncidentGroupingOutputReference
	_jsii_.Get(
		j,
		"grouping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GroupingInput() *SentinelAlertRuleScheduledIncidentGrouping {
	var returns *SentinelAlertRuleScheduledIncidentGrouping
	_jsii_.Get(
		j,
		"groupingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) InternalValue() *SentinelAlertRuleScheduledIncident {
	var returns *SentinelAlertRuleScheduledIncident
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledIncidentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SentinelAlertRuleScheduledIncidentOutputReference {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleScheduledIncidentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledIncidentOutputReference_Override(s SentinelAlertRuleScheduledIncidentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference)SetCreateIncidentEnabled(val interface{}) {
	if err := j.validateSetCreateIncidentEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createIncidentEnabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference)SetInternalValue(val *SentinelAlertRuleScheduledIncident) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) PutGrouping(value *SentinelAlertRuleScheduledIncidentGrouping) {
	if err := s.validatePutGroupingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGrouping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

