// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulenrt

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/sentinelalertrulenrt/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleNrtIncidentGroupingOutputReference interface {
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
	InternalValue() *SentinelAlertRuleNrtIncidentGrouping
	SetInternalValue(val *SentinelAlertRuleNrtIncidentGrouping)
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
	ResetByAlertDetails()
	ResetByCustomDetails()
	ResetByEntities()
	ResetEnabled()
	ResetEntityMatchingMethod()
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

// The jsii proxy struct for SentinelAlertRuleNrtIncidentGroupingOutputReference
type jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ByAlertDetails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byAlertDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ByAlertDetailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byAlertDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ByCustomDetails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byCustomDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ByCustomDetailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byCustomDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ByEntities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byEntities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ByEntitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"byEntitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) EntityMatchingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityMatchingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) EntityMatchingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityMatchingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) InternalValue() *SentinelAlertRuleNrtIncidentGrouping {
	var returns *SentinelAlertRuleNrtIncidentGrouping
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) LookbackDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) LookbackDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ReopenClosedIncidents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenClosedIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ReopenClosedIncidentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenClosedIncidentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleNrtIncidentGroupingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SentinelAlertRuleNrtIncidentGroupingOutputReference {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleNrtIncidentGroupingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrtIncidentGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleNrtIncidentGroupingOutputReference_Override(s SentinelAlertRuleNrtIncidentGroupingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrtIncidentGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetByAlertDetails(val *[]*string) {
	if err := j.validateSetByAlertDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byAlertDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetByCustomDetails(val *[]*string) {
	if err := j.validateSetByCustomDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byCustomDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetByEntities(val *[]*string) {
	if err := j.validateSetByEntitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byEntities",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetEntityMatchingMethod(val *string) {
	if err := j.validateSetEntityMatchingMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityMatchingMethod",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetInternalValue(val *SentinelAlertRuleNrtIncidentGrouping) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetLookbackDuration(val *string) {
	if err := j.validateSetLookbackDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackDuration",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetReopenClosedIncidents(val interface{}) {
	if err := j.validateSetReopenClosedIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reopenClosedIncidents",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ResetByAlertDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetByAlertDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ResetByCustomDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetByCustomDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ResetByEntities() {
	_jsii_.InvokeVoid(
		s,
		"resetByEntities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ResetEntityMatchingMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetEntityMatchingMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ResetLookbackDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetLookbackDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ResetReopenClosedIncidents() {
	_jsii_.InvokeVoid(
		s,
		"resetReopenClosedIncidents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleNrtIncidentGroupingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

