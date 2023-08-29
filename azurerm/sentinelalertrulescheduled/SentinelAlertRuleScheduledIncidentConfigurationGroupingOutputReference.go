// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/sentinelalertrulescheduled/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference interface {
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
	EntityMatchingMethod() *string
	SetEntityMatchingMethod(val *string)
	EntityMatchingMethodInput() *string
	// Experimental.
	Fqn() *string
	GroupByAlertDetails() *[]*string
	SetGroupByAlertDetails(val *[]*string)
	GroupByAlertDetailsInput() *[]*string
	GroupByCustomDetails() *[]*string
	SetGroupByCustomDetails(val *[]*string)
	GroupByCustomDetailsInput() *[]*string
	GroupByEntities() *[]*string
	SetGroupByEntities(val *[]*string)
	GroupByEntitiesInput() *[]*string
	InternalValue() *SentinelAlertRuleScheduledIncidentConfigurationGrouping
	SetInternalValue(val *SentinelAlertRuleScheduledIncidentConfigurationGrouping)
	LookbackDuration() *string
	SetLookbackDuration(val *string)
	LookbackDurationInput() *string
	ReopenClosedIncidents() interface{}
	SetReopenClosedIncidents(val interface{})
	ReopenClosedIncidentsInput() interface{}
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
	ResetEnabled()
	ResetEntityMatchingMethod()
	ResetGroupByAlertDetails()
	ResetGroupByCustomDetails()
	ResetGroupByEntities()
	ResetLookbackDuration()
	ResetReopenClosedIncidents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference
type jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) EntityMatchingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityMatchingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) EntityMatchingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityMatchingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByAlertDetails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByAlertDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByAlertDetailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByAlertDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByCustomDetails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByCustomDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByCustomDetailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByCustomDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByEntities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByEntities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByEntitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByEntitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) InternalValue() *SentinelAlertRuleScheduledIncidentConfigurationGrouping {
	var returns *SentinelAlertRuleScheduledIncidentConfigurationGrouping
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) LookbackDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) LookbackDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ReopenClosedIncidents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenClosedIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ReopenClosedIncidentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenClosedIncidentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference_Override(s SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetEntityMatchingMethod(val *string) {
	if err := j.validateSetEntityMatchingMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityMatchingMethod",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetGroupByAlertDetails(val *[]*string) {
	if err := j.validateSetGroupByAlertDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupByAlertDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetGroupByCustomDetails(val *[]*string) {
	if err := j.validateSetGroupByCustomDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupByCustomDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetGroupByEntities(val *[]*string) {
	if err := j.validateSetGroupByEntitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupByEntities",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetInternalValue(val *SentinelAlertRuleScheduledIncidentConfigurationGrouping) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetLookbackDuration(val *string) {
	if err := j.validateSetLookbackDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackDuration",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetReopenClosedIncidents(val interface{}) {
	if err := j.validateSetReopenClosedIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reopenClosedIncidents",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetEntityMatchingMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetEntityMatchingMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetGroupByAlertDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupByAlertDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetGroupByCustomDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupByCustomDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetGroupByEntities() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupByEntities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetLookbackDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetLookbackDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetReopenClosedIncidents() {
	_jsii_.InvokeVoid(
		s,
		"resetReopenClosedIncidents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

