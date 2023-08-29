// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingrulesuppression

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/monitoralertprocessingrulesuppression/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorAlertProcessingRuleSuppressionConditionOutputReference interface {
	cdktf.ComplexObject
	AlertContext() MonitorAlertProcessingRuleSuppressionConditionAlertContextOutputReference
	AlertContextInput() *MonitorAlertProcessingRuleSuppressionConditionAlertContext
	AlertRuleId() MonitorAlertProcessingRuleSuppressionConditionAlertRuleIdOutputReference
	AlertRuleIdInput() *MonitorAlertProcessingRuleSuppressionConditionAlertRuleId
	AlertRuleName() MonitorAlertProcessingRuleSuppressionConditionAlertRuleNameOutputReference
	AlertRuleNameInput() *MonitorAlertProcessingRuleSuppressionConditionAlertRuleName
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
	Description() MonitorAlertProcessingRuleSuppressionConditionDescriptionOutputReference
	DescriptionInput() *MonitorAlertProcessingRuleSuppressionConditionDescription
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleSuppressionCondition
	SetInternalValue(val *MonitorAlertProcessingRuleSuppressionCondition)
	MonitorCondition() MonitorAlertProcessingRuleSuppressionConditionMonitorConditionOutputReference
	MonitorConditionInput() *MonitorAlertProcessingRuleSuppressionConditionMonitorCondition
	MonitorService() MonitorAlertProcessingRuleSuppressionConditionMonitorServiceOutputReference
	MonitorServiceInput() *MonitorAlertProcessingRuleSuppressionConditionMonitorService
	Severity() MonitorAlertProcessingRuleSuppressionConditionSeverityOutputReference
	SeverityInput() *MonitorAlertProcessingRuleSuppressionConditionSeverity
	SignalType() MonitorAlertProcessingRuleSuppressionConditionSignalTypeOutputReference
	SignalTypeInput() *MonitorAlertProcessingRuleSuppressionConditionSignalType
	TargetResource() MonitorAlertProcessingRuleSuppressionConditionTargetResourceOutputReference
	TargetResourceGroup() MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroupOutputReference
	TargetResourceGroupInput() *MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroup
	TargetResourceInput() *MonitorAlertProcessingRuleSuppressionConditionTargetResource
	TargetResourceType() MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeOutputReference
	TargetResourceTypeInput() *MonitorAlertProcessingRuleSuppressionConditionTargetResourceType
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
	PutAlertContext(value *MonitorAlertProcessingRuleSuppressionConditionAlertContext)
	PutAlertRuleId(value *MonitorAlertProcessingRuleSuppressionConditionAlertRuleId)
	PutAlertRuleName(value *MonitorAlertProcessingRuleSuppressionConditionAlertRuleName)
	PutDescription(value *MonitorAlertProcessingRuleSuppressionConditionDescription)
	PutMonitorCondition(value *MonitorAlertProcessingRuleSuppressionConditionMonitorCondition)
	PutMonitorService(value *MonitorAlertProcessingRuleSuppressionConditionMonitorService)
	PutSeverity(value *MonitorAlertProcessingRuleSuppressionConditionSeverity)
	PutSignalType(value *MonitorAlertProcessingRuleSuppressionConditionSignalType)
	PutTargetResource(value *MonitorAlertProcessingRuleSuppressionConditionTargetResource)
	PutTargetResourceGroup(value *MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroup)
	PutTargetResourceType(value *MonitorAlertProcessingRuleSuppressionConditionTargetResourceType)
	ResetAlertContext()
	ResetAlertRuleId()
	ResetAlertRuleName()
	ResetDescription()
	ResetMonitorCondition()
	ResetMonitorService()
	ResetSeverity()
	ResetSignalType()
	ResetTargetResource()
	ResetTargetResourceGroup()
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

// The jsii proxy struct for MonitorAlertProcessingRuleSuppressionConditionOutputReference
type jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) AlertContext() MonitorAlertProcessingRuleSuppressionConditionAlertContextOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionAlertContextOutputReference
	_jsii_.Get(
		j,
		"alertContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) AlertContextInput() *MonitorAlertProcessingRuleSuppressionConditionAlertContext {
	var returns *MonitorAlertProcessingRuleSuppressionConditionAlertContext
	_jsii_.Get(
		j,
		"alertContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) AlertRuleId() MonitorAlertProcessingRuleSuppressionConditionAlertRuleIdOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionAlertRuleIdOutputReference
	_jsii_.Get(
		j,
		"alertRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) AlertRuleIdInput() *MonitorAlertProcessingRuleSuppressionConditionAlertRuleId {
	var returns *MonitorAlertProcessingRuleSuppressionConditionAlertRuleId
	_jsii_.Get(
		j,
		"alertRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) AlertRuleName() MonitorAlertProcessingRuleSuppressionConditionAlertRuleNameOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionAlertRuleNameOutputReference
	_jsii_.Get(
		j,
		"alertRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) AlertRuleNameInput() *MonitorAlertProcessingRuleSuppressionConditionAlertRuleName {
	var returns *MonitorAlertProcessingRuleSuppressionConditionAlertRuleName
	_jsii_.Get(
		j,
		"alertRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) Description() MonitorAlertProcessingRuleSuppressionConditionDescriptionOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionDescriptionOutputReference
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) DescriptionInput() *MonitorAlertProcessingRuleSuppressionConditionDescription {
	var returns *MonitorAlertProcessingRuleSuppressionConditionDescription
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) InternalValue() *MonitorAlertProcessingRuleSuppressionCondition {
	var returns *MonitorAlertProcessingRuleSuppressionCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) MonitorCondition() MonitorAlertProcessingRuleSuppressionConditionMonitorConditionOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionMonitorConditionOutputReference
	_jsii_.Get(
		j,
		"monitorCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) MonitorConditionInput() *MonitorAlertProcessingRuleSuppressionConditionMonitorCondition {
	var returns *MonitorAlertProcessingRuleSuppressionConditionMonitorCondition
	_jsii_.Get(
		j,
		"monitorConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) MonitorService() MonitorAlertProcessingRuleSuppressionConditionMonitorServiceOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionMonitorServiceOutputReference
	_jsii_.Get(
		j,
		"monitorService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) MonitorServiceInput() *MonitorAlertProcessingRuleSuppressionConditionMonitorService {
	var returns *MonitorAlertProcessingRuleSuppressionConditionMonitorService
	_jsii_.Get(
		j,
		"monitorServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) Severity() MonitorAlertProcessingRuleSuppressionConditionSeverityOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionSeverityOutputReference
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) SeverityInput() *MonitorAlertProcessingRuleSuppressionConditionSeverity {
	var returns *MonitorAlertProcessingRuleSuppressionConditionSeverity
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) SignalType() MonitorAlertProcessingRuleSuppressionConditionSignalTypeOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionSignalTypeOutputReference
	_jsii_.Get(
		j,
		"signalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) SignalTypeInput() *MonitorAlertProcessingRuleSuppressionConditionSignalType {
	var returns *MonitorAlertProcessingRuleSuppressionConditionSignalType
	_jsii_.Get(
		j,
		"signalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) TargetResource() MonitorAlertProcessingRuleSuppressionConditionTargetResourceOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionTargetResourceOutputReference
	_jsii_.Get(
		j,
		"targetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) TargetResourceGroup() MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroupOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroupOutputReference
	_jsii_.Get(
		j,
		"targetResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) TargetResourceGroupInput() *MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroup {
	var returns *MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroup
	_jsii_.Get(
		j,
		"targetResourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) TargetResourceInput() *MonitorAlertProcessingRuleSuppressionConditionTargetResource {
	var returns *MonitorAlertProcessingRuleSuppressionConditionTargetResource
	_jsii_.Get(
		j,
		"targetResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) TargetResourceType() MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeOutputReference {
	var returns MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeOutputReference
	_jsii_.Get(
		j,
		"targetResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) TargetResourceTypeInput() *MonitorAlertProcessingRuleSuppressionConditionTargetResourceType {
	var returns *MonitorAlertProcessingRuleSuppressionConditionTargetResourceType
	_jsii_.Get(
		j,
		"targetResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleSuppressionConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleSuppressionConditionOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorAlertProcessingRuleSuppressionConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleSuppression.MonitorAlertProcessingRuleSuppressionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleSuppressionConditionOutputReference_Override(m MonitorAlertProcessingRuleSuppressionConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleSuppression.MonitorAlertProcessingRuleSuppressionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference)SetInternalValue(val *MonitorAlertProcessingRuleSuppressionCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutAlertContext(value *MonitorAlertProcessingRuleSuppressionConditionAlertContext) {
	if err := m.validatePutAlertContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertContext",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutAlertRuleId(value *MonitorAlertProcessingRuleSuppressionConditionAlertRuleId) {
	if err := m.validatePutAlertRuleIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertRuleId",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutAlertRuleName(value *MonitorAlertProcessingRuleSuppressionConditionAlertRuleName) {
	if err := m.validatePutAlertRuleNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertRuleName",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutDescription(value *MonitorAlertProcessingRuleSuppressionConditionDescription) {
	if err := m.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDescription",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutMonitorCondition(value *MonitorAlertProcessingRuleSuppressionConditionMonitorCondition) {
	if err := m.validatePutMonitorConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorCondition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutMonitorService(value *MonitorAlertProcessingRuleSuppressionConditionMonitorService) {
	if err := m.validatePutMonitorServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorService",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutSeverity(value *MonitorAlertProcessingRuleSuppressionConditionSeverity) {
	if err := m.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSeverity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutSignalType(value *MonitorAlertProcessingRuleSuppressionConditionSignalType) {
	if err := m.validatePutSignalTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSignalType",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutTargetResource(value *MonitorAlertProcessingRuleSuppressionConditionTargetResource) {
	if err := m.validatePutTargetResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTargetResource",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutTargetResourceGroup(value *MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroup) {
	if err := m.validatePutTargetResourceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTargetResourceGroup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) PutTargetResourceType(value *MonitorAlertProcessingRuleSuppressionConditionTargetResourceType) {
	if err := m.validatePutTargetResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTargetResourceType",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetAlertContext() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertContext",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetAlertRuleId() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertRuleId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetAlertRuleName() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertRuleName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetMonitorCondition() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorCondition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetMonitorService() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorService",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		m,
		"resetSeverity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetSignalType() {
	_jsii_.InvokeVoid(
		m,
		"resetSignalType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetTargetResource() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetTargetResourceGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ResetTargetResourceType() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleSuppressionConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

