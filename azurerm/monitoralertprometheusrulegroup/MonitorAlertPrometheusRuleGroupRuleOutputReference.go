// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprometheusrulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/monitoralertprometheusrulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorAlertPrometheusRuleGroupRuleOutputReference interface {
	cdktf.ComplexObject
	Action() MonitorAlertPrometheusRuleGroupRuleActionList
	ActionInput() interface{}
	Alert() *string
	SetAlert(val *string)
	AlertInput() *string
	AlertResolution() MonitorAlertPrometheusRuleGroupRuleAlertResolutionOutputReference
	AlertResolutionInput() *MonitorAlertPrometheusRuleGroupRuleAlertResolution
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
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
	Expression() *string
	SetExpression(val *string)
	ExpressionInput() *string
	For() *string
	SetFor(val *string)
	ForInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Record() *string
	SetRecord(val *string)
	RecordInput() *string
	Severity() *float64
	SetSeverity(val *float64)
	SeverityInput() *float64
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
	PutAction(value interface{})
	PutAlertResolution(value *MonitorAlertPrometheusRuleGroupRuleAlertResolution)
	ResetAction()
	ResetAlert()
	ResetAlertResolution()
	ResetAnnotations()
	ResetEnabled()
	ResetFor()
	ResetLabels()
	ResetRecord()
	ResetSeverity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertPrometheusRuleGroupRuleOutputReference
type jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Action() MonitorAlertPrometheusRuleGroupRuleActionList {
	var returns MonitorAlertPrometheusRuleGroupRuleActionList
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Alert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) AlertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) AlertResolution() MonitorAlertPrometheusRuleGroupRuleAlertResolutionOutputReference {
	var returns MonitorAlertPrometheusRuleGroupRuleAlertResolutionOutputReference
	_jsii_.Get(
		j,
		"alertResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) AlertResolutionInput() *MonitorAlertPrometheusRuleGroupRuleAlertResolution {
	var returns *MonitorAlertPrometheusRuleGroupRuleAlertResolution
	_jsii_.Get(
		j,
		"alertResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) For() *string {
	var returns *string
	_jsii_.Get(
		j,
		"for",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ForInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Record() *string {
	var returns *string
	_jsii_.Get(
		j,
		"record",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) RecordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Severity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) SeverityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorAlertPrometheusRuleGroupRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorAlertPrometheusRuleGroupRuleOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorAlertPrometheusRuleGroupRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertPrometheusRuleGroup.MonitorAlertPrometheusRuleGroupRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorAlertPrometheusRuleGroupRuleOutputReference_Override(m MonitorAlertPrometheusRuleGroupRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertPrometheusRuleGroup.MonitorAlertPrometheusRuleGroupRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetAlert(val *string) {
	if err := j.validateSetAlertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alert",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetExpression(val *string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetFor(val *string) {
	if err := j.validateSetForParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"for",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetRecord(val *string) {
	if err := j.validateSetRecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"record",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetSeverity(val *float64) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) PutAction(value interface{}) {
	if err := m.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) PutAlertResolution(value *MonitorAlertPrometheusRuleGroupRuleAlertResolution) {
	if err := m.validatePutAlertResolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertResolution",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		m,
		"resetAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetAlert() {
	_jsii_.InvokeVoid(
		m,
		"resetAlert",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetAlertResolution() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertResolution",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetAnnotations() {
	_jsii_.InvokeVoid(
		m,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetFor() {
	_jsii_.InvokeVoid(
		m,
		"resetFor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		m,
		"resetLabels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetRecord() {
	_jsii_.InvokeVoid(
		m,
		"resetRecord",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		m,
		"resetSeverity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertPrometheusRuleGroupRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

