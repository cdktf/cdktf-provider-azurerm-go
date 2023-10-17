// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractionrulesuppression

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/monitoractionrulesuppression/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorActionRuleSuppressionConditionOutputReference interface {
	cdktf.ComplexObject
	AlertContext() MonitorActionRuleSuppressionConditionAlertContextOutputReference
	AlertContextInput() *MonitorActionRuleSuppressionConditionAlertContext
	AlertRuleId() MonitorActionRuleSuppressionConditionAlertRuleIdOutputReference
	AlertRuleIdInput() *MonitorActionRuleSuppressionConditionAlertRuleId
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
	Description() MonitorActionRuleSuppressionConditionDescriptionOutputReference
	DescriptionInput() *MonitorActionRuleSuppressionConditionDescription
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorActionRuleSuppressionCondition
	SetInternalValue(val *MonitorActionRuleSuppressionCondition)
	Monitor() MonitorActionRuleSuppressionConditionMonitorOutputReference
	MonitorInput() *MonitorActionRuleSuppressionConditionMonitor
	MonitorService() MonitorActionRuleSuppressionConditionMonitorServiceOutputReference
	MonitorServiceInput() *MonitorActionRuleSuppressionConditionMonitorService
	Severity() MonitorActionRuleSuppressionConditionSeverityOutputReference
	SeverityInput() *MonitorActionRuleSuppressionConditionSeverity
	TargetResourceType() MonitorActionRuleSuppressionConditionTargetResourceTypeOutputReference
	TargetResourceTypeInput() *MonitorActionRuleSuppressionConditionTargetResourceType
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
	PutAlertContext(value *MonitorActionRuleSuppressionConditionAlertContext)
	PutAlertRuleId(value *MonitorActionRuleSuppressionConditionAlertRuleId)
	PutDescription(value *MonitorActionRuleSuppressionConditionDescription)
	PutMonitor(value *MonitorActionRuleSuppressionConditionMonitor)
	PutMonitorService(value *MonitorActionRuleSuppressionConditionMonitorService)
	PutSeverity(value *MonitorActionRuleSuppressionConditionSeverity)
	PutTargetResourceType(value *MonitorActionRuleSuppressionConditionTargetResourceType)
	ResetAlertContext()
	ResetAlertRuleId()
	ResetDescription()
	ResetMonitor()
	ResetMonitorService()
	ResetSeverity()
	ResetTargetResourceType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActionRuleSuppressionConditionOutputReference
type jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) AlertContext() MonitorActionRuleSuppressionConditionAlertContextOutputReference {
	var returns MonitorActionRuleSuppressionConditionAlertContextOutputReference
	_jsii_.Get(
		j,
		"alertContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) AlertContextInput() *MonitorActionRuleSuppressionConditionAlertContext {
	var returns *MonitorActionRuleSuppressionConditionAlertContext
	_jsii_.Get(
		j,
		"alertContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) AlertRuleId() MonitorActionRuleSuppressionConditionAlertRuleIdOutputReference {
	var returns MonitorActionRuleSuppressionConditionAlertRuleIdOutputReference
	_jsii_.Get(
		j,
		"alertRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) AlertRuleIdInput() *MonitorActionRuleSuppressionConditionAlertRuleId {
	var returns *MonitorActionRuleSuppressionConditionAlertRuleId
	_jsii_.Get(
		j,
		"alertRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) Description() MonitorActionRuleSuppressionConditionDescriptionOutputReference {
	var returns MonitorActionRuleSuppressionConditionDescriptionOutputReference
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) DescriptionInput() *MonitorActionRuleSuppressionConditionDescription {
	var returns *MonitorActionRuleSuppressionConditionDescription
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) InternalValue() *MonitorActionRuleSuppressionCondition {
	var returns *MonitorActionRuleSuppressionCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) Monitor() MonitorActionRuleSuppressionConditionMonitorOutputReference {
	var returns MonitorActionRuleSuppressionConditionMonitorOutputReference
	_jsii_.Get(
		j,
		"monitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) MonitorInput() *MonitorActionRuleSuppressionConditionMonitor {
	var returns *MonitorActionRuleSuppressionConditionMonitor
	_jsii_.Get(
		j,
		"monitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) MonitorService() MonitorActionRuleSuppressionConditionMonitorServiceOutputReference {
	var returns MonitorActionRuleSuppressionConditionMonitorServiceOutputReference
	_jsii_.Get(
		j,
		"monitorService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) MonitorServiceInput() *MonitorActionRuleSuppressionConditionMonitorService {
	var returns *MonitorActionRuleSuppressionConditionMonitorService
	_jsii_.Get(
		j,
		"monitorServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) Severity() MonitorActionRuleSuppressionConditionSeverityOutputReference {
	var returns MonitorActionRuleSuppressionConditionSeverityOutputReference
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) SeverityInput() *MonitorActionRuleSuppressionConditionSeverity {
	var returns *MonitorActionRuleSuppressionConditionSeverity
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) TargetResourceType() MonitorActionRuleSuppressionConditionTargetResourceTypeOutputReference {
	var returns MonitorActionRuleSuppressionConditionTargetResourceTypeOutputReference
	_jsii_.Get(
		j,
		"targetResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) TargetResourceTypeInput() *MonitorActionRuleSuppressionConditionTargetResourceType {
	var returns *MonitorActionRuleSuppressionConditionTargetResourceType
	_jsii_.Get(
		j,
		"targetResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorActionRuleSuppressionConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorActionRuleSuppressionConditionOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorActionRuleSuppressionConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionRuleSuppression.MonitorActionRuleSuppressionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorActionRuleSuppressionConditionOutputReference_Override(m MonitorActionRuleSuppressionConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionRuleSuppression.MonitorActionRuleSuppressionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference)SetInternalValue(val *MonitorActionRuleSuppressionCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) PutAlertContext(value *MonitorActionRuleSuppressionConditionAlertContext) {
	if err := m.validatePutAlertContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertContext",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) PutAlertRuleId(value *MonitorActionRuleSuppressionConditionAlertRuleId) {
	if err := m.validatePutAlertRuleIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertRuleId",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) PutDescription(value *MonitorActionRuleSuppressionConditionDescription) {
	if err := m.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDescription",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) PutMonitor(value *MonitorActionRuleSuppressionConditionMonitor) {
	if err := m.validatePutMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitor",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) PutMonitorService(value *MonitorActionRuleSuppressionConditionMonitorService) {
	if err := m.validatePutMonitorServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorService",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) PutSeverity(value *MonitorActionRuleSuppressionConditionSeverity) {
	if err := m.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSeverity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) PutTargetResourceType(value *MonitorActionRuleSuppressionConditionTargetResourceType) {
	if err := m.validatePutTargetResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTargetResourceType",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ResetAlertContext() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertContext",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ResetAlertRuleId() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertRuleId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ResetMonitor() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ResetMonitorService() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorService",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		m,
		"resetSeverity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ResetTargetResourceType() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorActionRuleSuppressionConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

