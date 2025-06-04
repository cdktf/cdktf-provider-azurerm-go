// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulenrt

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/sentinelalertrulenrt/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleNrtAlertDetailsOverrideOutputReference interface {
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
	DynamicProperty() SentinelAlertRuleNrtAlertDetailsOverrideDynamicPropertyList
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

// The jsii proxy struct for SentinelAlertRuleNrtAlertDetailsOverrideOutputReference
type jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) DescriptionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) DescriptionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) DisplayNameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) DisplayNameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) DynamicProperty() SentinelAlertRuleNrtAlertDetailsOverrideDynamicPropertyList {
	var returns SentinelAlertRuleNrtAlertDetailsOverrideDynamicPropertyList
	_jsii_.Get(
		j,
		"dynamicProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) DynamicPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) SeverityColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) SeverityColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) TacticsColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticsColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) TacticsColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticsColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleNrtAlertDetailsOverrideOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SentinelAlertRuleNrtAlertDetailsOverrideOutputReference {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleNrtAlertDetailsOverrideOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrtAlertDetailsOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleNrtAlertDetailsOverrideOutputReference_Override(s SentinelAlertRuleNrtAlertDetailsOverrideOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrtAlertDetailsOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetDescriptionFormat(val *string) {
	if err := j.validateSetDescriptionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"descriptionFormat",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetDisplayNameFormat(val *string) {
	if err := j.validateSetDisplayNameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayNameFormat",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetSeverityColumnName(val *string) {
	if err := j.validateSetSeverityColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severityColumnName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetTacticsColumnName(val *string) {
	if err := j.validateSetTacticsColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tacticsColumnName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) PutDynamicProperty(value interface{}) {
	if err := s.validatePutDynamicPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDynamicProperty",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ResetDescriptionFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDescriptionFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ResetDisplayNameFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayNameFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ResetDynamicProperty() {
	_jsii_.InvokeVoid(
		s,
		"resetDynamicProperty",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ResetSeverityColumnName() {
	_jsii_.InvokeVoid(
		s,
		"resetSeverityColumnName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ResetTacticsColumnName() {
	_jsii_.InvokeVoid(
		s,
		"resetTacticsColumnName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleNrtAlertDetailsOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

