package monitoralertprocessingruleactiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/monitoralertprocessingruleactiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group azurerm_monitor_alert_processing_rule_action_group}.
type MonitorAlertProcessingRuleActionGroup interface {
	cdktf.TerraformResource
	AddActionGroupIds() *[]*string
	SetAddActionGroupIds(val *[]*string)
	AddActionGroupIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Condition() MonitorAlertProcessingRuleActionGroupConditionOutputReference
	ConditionInput() *MonitorAlertProcessingRuleActionGroupCondition
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Schedule() MonitorAlertProcessingRuleActionGroupScheduleOutputReference
	ScheduleInput() *MonitorAlertProcessingRuleActionGroupSchedule
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCondition(value *MonitorAlertProcessingRuleActionGroupCondition)
	PutSchedule(value *MonitorAlertProcessingRuleActionGroupSchedule)
	PutTimeouts(value *MonitorAlertProcessingRuleActionGroupTimeouts)
	ResetCondition()
	ResetDescription()
	ResetEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchedule()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroup
type jsiiProxy_MonitorAlertProcessingRuleActionGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) AddActionGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addActionGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) AddActionGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addActionGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Condition() MonitorAlertProcessingRuleActionGroupConditionOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ConditionInput() *MonitorAlertProcessingRuleActionGroupCondition {
	var returns *MonitorAlertProcessingRuleActionGroupCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Schedule() MonitorAlertProcessingRuleActionGroupScheduleOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ScheduleInput() *MonitorAlertProcessingRuleActionGroupSchedule {
	var returns *MonitorAlertProcessingRuleActionGroupSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) Timeouts() MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group azurerm_monitor_alert_processing_rule_action_group} Resource.
func NewMonitorAlertProcessingRuleActionGroup(scope constructs.Construct, id *string, config *MonitorAlertProcessingRuleActionGroupConfig) MonitorAlertProcessingRuleActionGroup {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroup{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group azurerm_monitor_alert_processing_rule_action_group} Resource.
func NewMonitorAlertProcessingRuleActionGroup_Override(m MonitorAlertProcessingRuleActionGroup, scope constructs.Construct, id *string, config *MonitorAlertProcessingRuleActionGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroup",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetAddActionGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"addActionGroupIds",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func MonitorAlertProcessingRuleActionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorAlertProcessingRuleActionGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) PutCondition(value *MonitorAlertProcessingRuleActionGroupCondition) {
	_jsii_.InvokeVoid(
		m,
		"putCondition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) PutSchedule(value *MonitorAlertProcessingRuleActionGroupSchedule) {
	_jsii_.InvokeVoid(
		m,
		"putSchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) PutTimeouts(value *MonitorAlertProcessingRuleActionGroupTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResetCondition() {
	_jsii_.InvokeVoid(
		m,
		"resetCondition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResetSchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetSchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupCondition struct {
	// alert_context block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#alert_context MonitorAlertProcessingRuleActionGroup#alert_context}
	AlertContext *MonitorAlertProcessingRuleActionGroupConditionAlertContext `field:"optional" json:"alertContext" yaml:"alertContext"`
	// alert_rule_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#alert_rule_id MonitorAlertProcessingRuleActionGroup#alert_rule_id}
	AlertRuleId *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId `field:"optional" json:"alertRuleId" yaml:"alertRuleId"`
	// alert_rule_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#alert_rule_name MonitorAlertProcessingRuleActionGroup#alert_rule_name}
	AlertRuleName *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName `field:"optional" json:"alertRuleName" yaml:"alertRuleName"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#description MonitorAlertProcessingRuleActionGroup#description}
	Description *MonitorAlertProcessingRuleActionGroupConditionDescription `field:"optional" json:"description" yaml:"description"`
	// monitor_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#monitor_condition MonitorAlertProcessingRuleActionGroup#monitor_condition}
	MonitorCondition *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition `field:"optional" json:"monitorCondition" yaml:"monitorCondition"`
	// monitor_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#monitor_service MonitorAlertProcessingRuleActionGroup#monitor_service}
	MonitorService *MonitorAlertProcessingRuleActionGroupConditionMonitorService `field:"optional" json:"monitorService" yaml:"monitorService"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#severity MonitorAlertProcessingRuleActionGroup#severity}
	Severity *MonitorAlertProcessingRuleActionGroupConditionSeverity `field:"optional" json:"severity" yaml:"severity"`
	// signal_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#signal_type MonitorAlertProcessingRuleActionGroup#signal_type}
	SignalType *MonitorAlertProcessingRuleActionGroupConditionSignalType `field:"optional" json:"signalType" yaml:"signalType"`
	// target_resource block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#target_resource MonitorAlertProcessingRuleActionGroup#target_resource}
	TargetResource *MonitorAlertProcessingRuleActionGroupConditionTargetResource `field:"optional" json:"targetResource" yaml:"targetResource"`
	// target_resource_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#target_resource_group MonitorAlertProcessingRuleActionGroup#target_resource_group}
	TargetResourceGroup *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup `field:"optional" json:"targetResourceGroup" yaml:"targetResourceGroup"`
	// target_resource_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#target_resource_type MonitorAlertProcessingRuleActionGroup#target_resource_type}
	TargetResourceType *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType `field:"optional" json:"targetResourceType" yaml:"targetResourceType"`
}

type MonitorAlertProcessingRuleActionGroupConditionAlertContext struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionAlertContext
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionAlertContext)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionAlertContext {
	var returns *MonitorAlertProcessingRuleActionGroupConditionAlertContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionAlertContext) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionAlertRuleId struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId {
	var returns *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleIdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionAlertRuleName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName {
	var returns *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionAlertRuleNameOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionDescription struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionDescription
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionDescription)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionDescription {
	var returns *MonitorAlertProcessingRuleActionGroupConditionDescription
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionDescription) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionDescriptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionMonitorCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition {
	var returns *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionMonitorService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionMonitorService
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionMonitorService)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionMonitorService {
	var returns *MonitorAlertProcessingRuleActionGroupConditionMonitorService
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionMonitorService) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionMonitorServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupCondition) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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
	_jsii_.InvokeVoid(
		m,
		"putAlertContext",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutAlertRuleId(value *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId) {
	_jsii_.InvokeVoid(
		m,
		"putAlertRuleId",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutAlertRuleName(value *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName) {
	_jsii_.InvokeVoid(
		m,
		"putAlertRuleName",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutDescription(value *MonitorAlertProcessingRuleActionGroupConditionDescription) {
	_jsii_.InvokeVoid(
		m,
		"putDescription",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutMonitorCondition(value *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition) {
	_jsii_.InvokeVoid(
		m,
		"putMonitorCondition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutMonitorService(value *MonitorAlertProcessingRuleActionGroupConditionMonitorService) {
	_jsii_.InvokeVoid(
		m,
		"putMonitorService",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutSeverity(value *MonitorAlertProcessingRuleActionGroupConditionSeverity) {
	_jsii_.InvokeVoid(
		m,
		"putSeverity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutSignalType(value *MonitorAlertProcessingRuleActionGroupConditionSignalType) {
	_jsii_.InvokeVoid(
		m,
		"putSignalType",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutTargetResource(value *MonitorAlertProcessingRuleActionGroupConditionTargetResource) {
	_jsii_.InvokeVoid(
		m,
		"putTargetResource",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutTargetResourceGroup(value *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup) {
	_jsii_.InvokeVoid(
		m,
		"putTargetResourceGroup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionOutputReference) PutTargetResourceType(value *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType) {
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

type MonitorAlertProcessingRuleActionGroupConditionSeverity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionSeverity
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionSeverity)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionSeverity {
	var returns *MonitorAlertProcessingRuleActionGroupConditionSeverity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionSeverity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSeverityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionSignalType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionSignalType
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionSignalType)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionSignalType {
	var returns *MonitorAlertProcessingRuleActionGroupConditionSignalType
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionSignalType) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionSignalTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionTargetResource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup {
	var returns *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionTargetResource
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionTargetResource)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionTargetResource {
	var returns *MonitorAlertProcessingRuleActionGroupConditionTargetResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionTargetResource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConditionTargetResourceType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType {
	var returns *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference_Override(m MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#add_action_group_ids MonitorAlertProcessingRuleActionGroup#add_action_group_ids}.
	AddActionGroupIds *[]*string `field:"required" json:"addActionGroupIds" yaml:"addActionGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#name MonitorAlertProcessingRuleActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#resource_group_name MonitorAlertProcessingRuleActionGroup#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#scopes MonitorAlertProcessingRuleActionGroup#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#condition MonitorAlertProcessingRuleActionGroup#condition}
	Condition *MonitorAlertProcessingRuleActionGroupCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#description MonitorAlertProcessingRuleActionGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#enabled MonitorAlertProcessingRuleActionGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#id MonitorAlertProcessingRuleActionGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#schedule MonitorAlertProcessingRuleActionGroup#schedule}
	Schedule *MonitorAlertProcessingRuleActionGroupSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#tags MonitorAlertProcessingRuleActionGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#timeouts MonitorAlertProcessingRuleActionGroup#timeouts}
	Timeouts *MonitorAlertProcessingRuleActionGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type MonitorAlertProcessingRuleActionGroupSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#effective_from MonitorAlertProcessingRuleActionGroup#effective_from}.
	EffectiveFrom *string `field:"optional" json:"effectiveFrom" yaml:"effectiveFrom"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#effective_until MonitorAlertProcessingRuleActionGroup#effective_until}.
	EffectiveUntil *string `field:"optional" json:"effectiveUntil" yaml:"effectiveUntil"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#recurrence MonitorAlertProcessingRuleActionGroup#recurrence}
	Recurrence *MonitorAlertProcessingRuleActionGroupScheduleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#time_zone MonitorAlertProcessingRuleActionGroup#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

type MonitorAlertProcessingRuleActionGroupScheduleOutputReference interface {
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
	EffectiveFrom() *string
	SetEffectiveFrom(val *string)
	EffectiveFromInput() *string
	EffectiveUntil() *string
	SetEffectiveUntil(val *string)
	EffectiveUntilInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupSchedule
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupSchedule)
	Recurrence() MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference
	RecurrenceInput() *MonitorAlertProcessingRuleActionGroupScheduleRecurrence
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	PutRecurrence(value *MonitorAlertProcessingRuleActionGroupScheduleRecurrence)
	ResetEffectiveFrom()
	ResetEffectiveUntil()
	ResetRecurrence()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupScheduleOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) EffectiveFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) EffectiveFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) EffectiveUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) EffectiveUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveUntilInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupSchedule {
	var returns *MonitorAlertProcessingRuleActionGroupSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) Recurrence() MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference
	_jsii_.Get(
		j,
		"recurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) RecurrenceInput() *MonitorAlertProcessingRuleActionGroupScheduleRecurrence {
	var returns *MonitorAlertProcessingRuleActionGroupScheduleRecurrence
	_jsii_.Get(
		j,
		"recurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupScheduleOutputReference_Override(m MonitorAlertProcessingRuleActionGroupScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) SetEffectiveFrom(val *string) {
	_jsii_.Set(
		j,
		"effectiveFrom",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) SetEffectiveUntil(val *string) {
	_jsii_.Set(
		j,
		"effectiveUntil",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupSchedule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) PutRecurrence(value *MonitorAlertProcessingRuleActionGroupScheduleRecurrence) {
	_jsii_.InvokeVoid(
		m,
		"putRecurrence",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) ResetEffectiveFrom() {
	_jsii_.InvokeVoid(
		m,
		"resetEffectiveFrom",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) ResetEffectiveUntil() {
	_jsii_.InvokeVoid(
		m,
		"resetEffectiveUntil",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) ResetRecurrence() {
	_jsii_.InvokeVoid(
		m,
		"resetRecurrence",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrence struct {
	// daily block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#daily MonitorAlertProcessingRuleActionGroup#daily}
	Daily interface{} `field:"optional" json:"daily" yaml:"daily"`
	// monthly block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#monthly MonitorAlertProcessingRuleActionGroup#monthly}
	Monthly interface{} `field:"optional" json:"monthly" yaml:"monthly"`
	// weekly block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#weekly MonitorAlertProcessingRuleActionGroup#weekly}
	Weekly interface{} `field:"optional" json:"weekly" yaml:"weekly"`
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDaily struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#end_time MonitorAlertProcessingRuleActionGroup#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#start_time MonitorAlertProcessingRuleActionGroup#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList
type jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList_Override(m MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) Get(index *float64) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference interface {
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
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference_Override(m MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) SetEndTime(val *string) {
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#days_of_month MonitorAlertProcessingRuleActionGroup#days_of_month}.
	DaysOfMonth *[]*float64 `field:"required" json:"daysOfMonth" yaml:"daysOfMonth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#end_time MonitorAlertProcessingRuleActionGroup#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#start_time MonitorAlertProcessingRuleActionGroup#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList
type jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList_Override(m MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) Get(index *float64) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference interface {
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
	DaysOfMonth() *[]*float64
	SetDaysOfMonth(val *[]*float64)
	DaysOfMonthInput() *[]*float64
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
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
	ResetEndTime()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) DaysOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) DaysOfMonthInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysOfMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference_Override(m MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) SetDaysOfMonth(val *[]*float64) {
	_jsii_.Set(
		j,
		"daysOfMonth",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) SetEndTime(val *string) {
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		m,
		"resetEndTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference interface {
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
	Daily() MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList
	DailyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAlertProcessingRuleActionGroupScheduleRecurrence
	SetInternalValue(val *MonitorAlertProcessingRuleActionGroupScheduleRecurrence)
	Monthly() MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList
	MonthlyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weekly() MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList
	WeeklyInput() interface{}
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
	PutDaily(value interface{})
	PutMonthly(value interface{})
	PutWeekly(value interface{})
	ResetDaily()
	ResetMonthly()
	ResetWeekly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) Daily() MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList {
	var returns MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDailyList
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) DailyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) InternalValue() *MonitorAlertProcessingRuleActionGroupScheduleRecurrence {
	var returns *MonitorAlertProcessingRuleActionGroupScheduleRecurrence
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) Monthly() MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList {
	var returns MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthlyList
	_jsii_.Get(
		j,
		"monthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) MonthlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) Weekly() MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList {
	var returns MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList
	_jsii_.Get(
		j,
		"weekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) WeeklyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference_Override(m MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) SetInternalValue(val *MonitorAlertProcessingRuleActionGroupScheduleRecurrence) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) PutDaily(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putDaily",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) PutMonthly(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putMonthly",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) PutWeekly(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putWeekly",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		m,
		"resetDaily",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) ResetMonthly() {
	_jsii_.InvokeVoid(
		m,
		"resetMonthly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) ResetWeekly() {
	_jsii_.InvokeVoid(
		m,
		"resetWeekly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeekly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#days_of_week MonitorAlertProcessingRuleActionGroup#days_of_week}.
	DaysOfWeek *[]*string `field:"required" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#end_time MonitorAlertProcessingRuleActionGroup#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#start_time MonitorAlertProcessingRuleActionGroup#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList
type jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList_Override(m MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) Get(index *float64) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference {
	var returns MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference interface {
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
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
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
	ResetEndTime()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference_Override(m MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) SetDaysOfWeek(val *[]*string) {
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) SetEndTime(val *string) {
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		m,
		"resetEndTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeeklyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorAlertProcessingRuleActionGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#create MonitorAlertProcessingRuleActionGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#delete MonitorAlertProcessingRuleActionGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#read MonitorAlertProcessingRuleActionGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#update MonitorAlertProcessingRuleActionGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference interface {
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
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference
type jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewMonitorAlertProcessingRuleActionGroupTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAlertProcessingRuleActionGroupTimeoutsOutputReference_Override(m MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorAlertProcessingRuleActionGroup.MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAlertProcessingRuleActionGroupTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

