// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingruleactiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/monitoralertprocessingruleactiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorAlertProcessingRuleActionGroupConditionOutputReference interface {
	cdktf.ComplexObject
	AlertContext() MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference
	AlertContextInput() *MonitorAlertProcessingRuleActionGroupConditionAlertContext
	AlertRuleId() MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference
	AlertRuleIdInput() *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId
	AlertRuleName() MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference
	AlertRuleNameInput() *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName
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
	Description() MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference
	DescriptionInput() *MonitorAlertProcessingRuleActionGroupConditionDescription
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupCondition
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupCondition)
	MonitorCondition() MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference
	MonitorConditionInput() *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition
	MonitorService() MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference
	MonitorServiceInput() *MonitorAlertProcessingRuleActionGroupConditionMonitorService
	Severity() MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference
	SeverityInput() *MonitorAlertProcessingRuleActionGroupConditionSeverity
	SignalType() MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference
	SignalTypeInput() *MonitorAlertProcessingRuleActionGroupConditionSignalType
	TargetResource() MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference
	TargetResourceGroup() MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference
	TargetResourceGroupInput() *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup
	TargetResourceInput() *MonitorAlertProcessingRuleActionGroupConditionTargetResource
	TargetResourceType() MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference
	TargetResourceTypeInput() *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType
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
	PutAlertContext(value *MonitorAlertProcessingRuleActionGroupConditionAlertContext)
	PutAlertRuleId(value *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId)
	PutAlertRuleName(value *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName)
	PutDescription(value *MonitorAlertProcessingRuleActionGroupConditionDescription)
	PutMonitorCondition(value *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition)
	PutMonitorService(value *MonitorAlertProcessingRuleActionGroupConditionMonitorService)
	PutSeverity(value *MonitorAlertProcessingRuleActionGroupConditionSeverity)
	PutSignalType(value *MonitorAlertProcessingRuleActionGroupConditionSignalType)
	PutTargetResource(value *MonitorAlertProcessingRuleActionGroupConditionTargetResource)
	PutTargetResourceGroup(value *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup)
	PutTargetResourceType(value *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType)
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

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) AlertContext() MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference
	_jsii_.Get(
		j,
		"alertContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) AlertContextInput() *MonitorAlertProcessingRuleActionGroupConditionAlertContext {
	var returns *MonitorAlertProcessingRuleActionGroupConditionAlertContext
	_jsii_.Get(
		j,
		"alertContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) AlertRuleId() MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference
	_jsii_.Get(
		j,
		"alertRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) AlertRuleIdInput() *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId {
	var returns *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId
	_jsii_.Get(
		j,
		"alertRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) AlertRuleName() MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference
	_jsii_.Get(
		j,
		"alertRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) AlertRuleNameInput() *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName {
	var returns *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName
	_jsii_.Get(
		j,
		"alertRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) Description() MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) DescriptionInput() *MonitorAlertProcessingRuleActionGroupConditionDescription {
	var returns *MonitorAlertProcessingRuleActionGroupConditionDescription
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupCondition {
	var returns *MonitorAlertProcessingRuleActionGroupCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) MonitorCondition() MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference
	_jsii_.Get(
		j,
		"monitorCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) MonitorConditionInput() *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition {
	var returns *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition
	_jsii_.Get(
		j,
		"monitorConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) MonitorService() MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference
	_jsii_.Get(
		j,
		"monitorService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) MonitorServiceInput() *MonitorAlertProcessingRuleActionGroupConditionMonitorService {
	var returns *MonitorAlertProcessingRuleActionGroupConditionMonitorService
	_jsii_.Get(
		j,
		"monitorServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) Severity() MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) SeverityInput() *MonitorAlertProcessingRuleActionGroupConditionSeverity {
	var returns *MonitorAlertProcessingRuleActionGroupConditionSeverity
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) SignalType() MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference
	_jsii_.Get(
		j,
		"signalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) SignalTypeInput() *MonitorAlertProcessingRuleActionGroupConditionSignalType {
	var returns *MonitorAlertProcessingRuleActionGroupConditionSignalType
	_jsii_.Get(
		j,
		"signalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) TargetResource() MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference
	_jsii_.Get(
		j,
		"targetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) TargetResourceGroup() MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference
	_jsii_.Get(
		j,
		"targetResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) TargetResourceGroupInput() *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup {
	var returns *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup
	_jsii_.Get(
		j,
		"targetResourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) TargetResourceInput() *MonitorAlertProcessingRuleActionGroupConditionTargetResource {
	var returns *MonitorAlertProcessingRuleActionGroupConditionTargetResource
	_jsii_.Get(
		j,
		"targetResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) TargetResourceType() MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference
	_jsii_.Get(
		j,
		"targetResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) TargetResourceTypeInput() *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType {
	var returns *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType
	_jsii_.Get(
		j,
		"targetResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorAlertProcessingRuleActionGroupConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference)SetInternalValue(val *MonitorAlertProcessingRuleActionGroupCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutAlertContext(value *MonitorAlertProcessingRuleActionGroupConditionAlertContext) {
	if err := m.validatePutAlertContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertContext",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutAlertRuleId(value *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId) {
	if err := m.validatePutAlertRuleIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertRuleId",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutAlertRuleName(value *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName) {
	if err := m.validatePutAlertRuleNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertRuleName",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutDescription(value *MonitorAlertProcessingRuleActionGroupConditionDescription) {
	if err := m.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDescription",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutMonitorCondition(value *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition) {
	if err := m.validatePutMonitorConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorCondition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutMonitorService(value *MonitorAlertProcessingRuleActionGroupConditionMonitorService) {
	if err := m.validatePutMonitorServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorService",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutSeverity(value *MonitorAlertProcessingRuleActionGroupConditionSeverity) {
	if err := m.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSeverity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutSignalType(value *MonitorAlertProcessingRuleActionGroupConditionSignalType) {
	if err := m.validatePutSignalTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSignalType",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutTargetResource(value *MonitorAlertProcessingRuleActionGroupConditionTargetResource) {
	if err := m.validatePutTargetResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTargetResource",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutTargetResourceGroup(value *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup) {
	if err := m.validatePutTargetResourceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTargetResourceGroup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutTargetResourceType(value *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType) {
	if err := m.validatePutTargetResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTargetResourceType",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetAlertContext() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertContext",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetAlertRuleId() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertRuleId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetAlertRuleName() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertRuleName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetMonitorCondition() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorCondition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetMonitorService() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorService",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		m,
		"resetSeverity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetSignalType() {
	_jsii_.InvokeVoid(
		m,
		"resetSignalType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetTargetResource() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetTargetResourceGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ResetTargetResourceType() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

