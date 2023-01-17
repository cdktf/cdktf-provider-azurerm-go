package monitoractionruleactiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/monitoractionruleactiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorActionRuleActionGroupConditionOutputReference interface {
	cdktf.ComplexObject
	AlertContext() MonitorActionRuleActionGroupConditionAlertContextOutputReference
	AlertContextInput() *MonitorActionRuleActionGroupConditionAlertContext
	AlertRuleId() MonitorActionRuleActionGroupConditionAlertRuleIdOutputReference
	AlertRuleIdInput() *MonitorActionRuleActionGroupConditionAlertRuleId
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
	Description() MonitorActionRuleActionGroupConditionDescriptionOutputReference
	DescriptionInput() *MonitorActionRuleActionGroupConditionDescription
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorActionRuleActionGroupCondition
	SetInternalValue(val *MonitorActionRuleActionGroupCondition)
	Monitor() MonitorActionRuleActionGroupConditionMonitorOutputReference
	MonitorInput() *MonitorActionRuleActionGroupConditionMonitor
	MonitorService() MonitorActionRuleActionGroupConditionMonitorServiceOutputReference
	MonitorServiceInput() *MonitorActionRuleActionGroupConditionMonitorService
	Severity() MonitorActionRuleActionGroupConditionSeverityOutputReference
	SeverityInput() *MonitorActionRuleActionGroupConditionSeverity
	TargetResourceType() MonitorActionRuleActionGroupConditionTargetResourceTypeOutputReference
	TargetResourceTypeInput() *MonitorActionRuleActionGroupConditionTargetResourceType
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
	PutAlertContext(value *MonitorActionRuleActionGroupConditionAlertContext)
	PutAlertRuleId(value *MonitorActionRuleActionGroupConditionAlertRuleId)
	PutDescription(value *MonitorActionRuleActionGroupConditionDescription)
	PutMonitor(value *MonitorActionRuleActionGroupConditionMonitor)
	PutMonitorService(value *MonitorActionRuleActionGroupConditionMonitorService)
	PutSeverity(value *MonitorActionRuleActionGroupConditionSeverity)
	PutTargetResourceType(value *MonitorActionRuleActionGroupConditionTargetResourceType)
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

// The jsii proxy struct for MonitorActionRuleActionGroupConditionOutputReference
type jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) AlertContext() MonitorActionRuleActionGroupConditionAlertContextOutputReference {
	var returns MonitorActionRuleActionGroupConditionAlertContextOutputReference
	_jsii_.Get(
		j,
		"alertContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) AlertContextInput() *MonitorActionRuleActionGroupConditionAlertContext {
	var returns *MonitorActionRuleActionGroupConditionAlertContext
	_jsii_.Get(
		j,
		"alertContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) AlertRuleId() MonitorActionRuleActionGroupConditionAlertRuleIdOutputReference {
	var returns MonitorActionRuleActionGroupConditionAlertRuleIdOutputReference
	_jsii_.Get(
		j,
		"alertRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) AlertRuleIdInput() *MonitorActionRuleActionGroupConditionAlertRuleId {
	var returns *MonitorActionRuleActionGroupConditionAlertRuleId
	_jsii_.Get(
		j,
		"alertRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) Description() MonitorActionRuleActionGroupConditionDescriptionOutputReference {
	var returns MonitorActionRuleActionGroupConditionDescriptionOutputReference
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) DescriptionInput() *MonitorActionRuleActionGroupConditionDescription {
	var returns *MonitorActionRuleActionGroupConditionDescription
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) InternalValue() *MonitorActionRuleActionGroupCondition {
	var returns *MonitorActionRuleActionGroupCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) Monitor() MonitorActionRuleActionGroupConditionMonitorOutputReference {
	var returns MonitorActionRuleActionGroupConditionMonitorOutputReference
	_jsii_.Get(
		j,
		"monitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) MonitorInput() *MonitorActionRuleActionGroupConditionMonitor {
	var returns *MonitorActionRuleActionGroupConditionMonitor
	_jsii_.Get(
		j,
		"monitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) MonitorService() MonitorActionRuleActionGroupConditionMonitorServiceOutputReference {
	var returns MonitorActionRuleActionGroupConditionMonitorServiceOutputReference
	_jsii_.Get(
		j,
		"monitorService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) MonitorServiceInput() *MonitorActionRuleActionGroupConditionMonitorService {
	var returns *MonitorActionRuleActionGroupConditionMonitorService
	_jsii_.Get(
		j,
		"monitorServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) Severity() MonitorActionRuleActionGroupConditionSeverityOutputReference {
	var returns MonitorActionRuleActionGroupConditionSeverityOutputReference
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) SeverityInput() *MonitorActionRuleActionGroupConditionSeverity {
	var returns *MonitorActionRuleActionGroupConditionSeverity
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) TargetResourceType() MonitorActionRuleActionGroupConditionTargetResourceTypeOutputReference {
	var returns MonitorActionRuleActionGroupConditionTargetResourceTypeOutputReference
	_jsii_.Get(
		j,
		"targetResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) TargetResourceTypeInput() *MonitorActionRuleActionGroupConditionTargetResourceType {
	var returns *MonitorActionRuleActionGroupConditionTargetResourceType
	_jsii_.Get(
		j,
		"targetResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorActionRuleActionGroupConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorActionRuleActionGroupConditionOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorActionRuleActionGroupConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionRuleActionGroup.MonitorActionRuleActionGroupConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorActionRuleActionGroupConditionOutputReference_Override(m MonitorActionRuleActionGroupConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionRuleActionGroup.MonitorActionRuleActionGroupConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference)SetInternalValue(val *MonitorActionRuleActionGroupCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) PutAlertContext(value *MonitorActionRuleActionGroupConditionAlertContext) {
	if err := m.validatePutAlertContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertContext",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) PutAlertRuleId(value *MonitorActionRuleActionGroupConditionAlertRuleId) {
	if err := m.validatePutAlertRuleIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAlertRuleId",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) PutDescription(value *MonitorActionRuleActionGroupConditionDescription) {
	if err := m.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDescription",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) PutMonitor(value *MonitorActionRuleActionGroupConditionMonitor) {
	if err := m.validatePutMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitor",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) PutMonitorService(value *MonitorActionRuleActionGroupConditionMonitorService) {
	if err := m.validatePutMonitorServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorService",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) PutSeverity(value *MonitorActionRuleActionGroupConditionSeverity) {
	if err := m.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSeverity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) PutTargetResourceType(value *MonitorActionRuleActionGroupConditionTargetResourceType) {
	if err := m.validatePutTargetResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTargetResourceType",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ResetAlertContext() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertContext",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ResetAlertRuleId() {
	_jsii_.InvokeVoid(
		m,
		"resetAlertRuleId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ResetMonitor() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ResetMonitorService() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorService",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		m,
		"resetSeverity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ResetTargetResourceType() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorActionRuleActionGroupConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

