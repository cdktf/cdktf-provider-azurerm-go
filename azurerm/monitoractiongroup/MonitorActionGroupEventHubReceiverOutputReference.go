package monitoractiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/monitoractiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorActionGroupEventHubReceiverOutputReference interface {
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
	EventHubId() *string
	SetEventHubId(val *string)
	EventHubIdInput() *string
	EventHubName() *string
	SetEventHubName(val *string)
	EventHubNameInput() *string
	EventHubNamespace() *string
	SetEventHubNamespace(val *string)
	EventHubNamespaceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseCommonAlertSchema() interface{}
	SetUseCommonAlertSchema(val interface{})
	UseCommonAlertSchemaInput() interface{}
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
	ResetEventHubId()
	ResetEventHubName()
	ResetEventHubNamespace()
	ResetSubscriptionId()
	ResetTenantId()
	ResetUseCommonAlertSchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActionGroupEventHubReceiverOutputReference
type jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) EventHubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) EventHubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) EventHubName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) EventHubNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) EventHubNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) EventHubNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) UseCommonAlertSchema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCommonAlertSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) UseCommonAlertSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCommonAlertSchemaInput",
		&returns,
	)
	return returns
}


func NewMonitorActionGroupEventHubReceiverOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorActionGroupEventHubReceiverOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorActionGroupEventHubReceiverOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionGroup.MonitorActionGroupEventHubReceiverOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorActionGroupEventHubReceiverOutputReference_Override(m MonitorActionGroupEventHubReceiverOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionGroup.MonitorActionGroupEventHubReceiverOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetEventHubId(val *string) {
	if err := j.validateSetEventHubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHubId",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetEventHubName(val *string) {
	if err := j.validateSetEventHubNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHubName",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetEventHubNamespace(val *string) {
	if err := j.validateSetEventHubNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHubNamespace",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetSubscriptionId(val *string) {
	if err := j.validateSetSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference)SetUseCommonAlertSchema(val interface{}) {
	if err := j.validateSetUseCommonAlertSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCommonAlertSchema",
		val,
	)
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ResetEventHubId() {
	_jsii_.InvokeVoid(
		m,
		"resetEventHubId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ResetEventHubName() {
	_jsii_.InvokeVoid(
		m,
		"resetEventHubName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ResetEventHubNamespace() {
	_jsii_.InvokeVoid(
		m,
		"resetEventHubNamespace",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		m,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		m,
		"resetTenantId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ResetUseCommonAlertSchema() {
	_jsii_.InvokeVoid(
		m,
		"resetUseCommonAlertSchema",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorActionGroupEventHubReceiverOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

