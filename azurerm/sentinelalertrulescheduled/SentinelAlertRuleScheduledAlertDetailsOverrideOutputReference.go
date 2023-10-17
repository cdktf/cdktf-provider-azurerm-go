// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/sentinelalertrulescheduled/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference interface {
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
	DescriptionFormat() *string
	SetDescriptionFormat(val *string)
	DescriptionFormatInput() *string
	DisplayNameFormat() *string
	SetDisplayNameFormat(val *string)
	DisplayNameFormatInput() *string
	DynamicProperty() SentinelAlertRuleScheduledAlertDetailsOverrideDynamicPropertyList
	DynamicPropertyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SeverityColumnName() *string
	SetSeverityColumnName(val *string)
	SeverityColumnNameInput() *string
	TacticsColumnName() *string
	SetTacticsColumnName(val *string)
	TacticsColumnNameInput() *string
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
	PutDynamicProperty(value interface{})
	ResetDescriptionFormat()
	ResetDisplayNameFormat()
	ResetDynamicProperty()
	ResetSeverityColumnName()
	ResetTacticsColumnName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference
type jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DescriptionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DescriptionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DisplayNameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DisplayNameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DynamicProperty() SentinelAlertRuleScheduledAlertDetailsOverrideDynamicPropertyList {
	var returns SentinelAlertRuleScheduledAlertDetailsOverrideDynamicPropertyList
	_jsii_.Get(
		j,
		"dynamicProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DynamicPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SeverityColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SeverityColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) TacticsColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticsColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) TacticsColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticsColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledAlertDetailsOverrideOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleScheduledAlertDetailsOverrideOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledAlertDetailsOverrideOutputReference_Override(s SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetDescriptionFormat(val *string) {
	if err := j.validateSetDescriptionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"descriptionFormat",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetDisplayNameFormat(val *string) {
	if err := j.validateSetDisplayNameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayNameFormat",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetSeverityColumnName(val *string) {
	if err := j.validateSetSeverityColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severityColumnName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetTacticsColumnName(val *string) {
	if err := j.validateSetTacticsColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tacticsColumnName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) PutDynamicProperty(value interface{}) {
	if err := s.validatePutDynamicPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDynamicProperty",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetDescriptionFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDescriptionFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetDisplayNameFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayNameFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetDynamicProperty() {
	_jsii_.InvokeVoid(
		s,
		"resetDynamicProperty",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetSeverityColumnName() {
	_jsii_.InvokeVoid(
		s,
		"resetSeverityColumnName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetTacticsColumnName() {
	_jsii_.InvokeVoid(
		s,
		"resetTacticsColumnName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

