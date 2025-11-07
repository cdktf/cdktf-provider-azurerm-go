// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/sentinelalertrulescheduled/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleScheduledIncidentGroupingOutputReference interface {
	cdktf.ComplexObject
	ByAlertDetails() *[]*string
	SetByAlertDetails(val *[]*string)
	ByAlertDetailsInput() *[]*string
	ByCustomDetails() *[]*string
	SetByCustomDetails(val *[]*string)
	ByCustomDetailsInput() *[]*string
	ByEntities() *[]*string
	SetByEntities(val *[]*string)
	ByEntitiesInput() *[]*string
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
	InternalValue() *SentinelAlertRuleScheduledIncidentGrouping
	SetInternalValue(val *SentinelAlertRuleScheduledIncidentGrouping)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetByAlertDetails()
	ResetByCustomDetails()
	ResetByEntities()
	ResetEnabled()
	ResetEntityMatchingMethod()
	ResetLookbackDuration()
	ResetReopenClosedIncidents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledIncidentGroupingOutputReference
type jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ByAlertDetails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byAlertDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ByAlertDetailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byAlertDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ByCustomDetails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byCustomDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ByCustomDetailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byCustomDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ByEntities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byEntities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ByEntitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byEntitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) EntityMatchingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityMatchingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) EntityMatchingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityMatchingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) InternalValue() *SentinelAlertRuleScheduledIncidentGrouping {
	var returns *SentinelAlertRuleScheduledIncidentGrouping
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) LookbackDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) LookbackDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ReopenClosedIncidents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenClosedIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ReopenClosedIncidentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenClosedIncidentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledIncidentGroupingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SentinelAlertRuleScheduledIncidentGroupingOutputReference {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleScheduledIncidentGroupingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledIncidentGroupingOutputReference_Override(s SentinelAlertRuleScheduledIncidentGroupingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetByAlertDetails(val *[]*string) {
	if err := j.validateSetByAlertDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byAlertDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetByCustomDetails(val *[]*string) {
	if err := j.validateSetByCustomDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byCustomDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetByEntities(val *[]*string) {
	if err := j.validateSetByEntitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byEntities",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetEntityMatchingMethod(val *string) {
	if err := j.validateSetEntityMatchingMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityMatchingMethod",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetInternalValue(val *SentinelAlertRuleScheduledIncidentGrouping) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetLookbackDuration(val *string) {
	if err := j.validateSetLookbackDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackDuration",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetReopenClosedIncidents(val interface{}) {
	if err := j.validateSetReopenClosedIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reopenClosedIncidents",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ResetByAlertDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetByAlertDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ResetByCustomDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetByCustomDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ResetByEntities() {
	_jsii_.InvokeVoid(
		s,
		"resetByEntities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ResetEntityMatchingMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetEntityMatchingMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ResetLookbackDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetLookbackDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ResetReopenClosedIncidents() {
	_jsii_.InvokeVoid(
		s,
		"resetReopenClosedIncidents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentGroupingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

