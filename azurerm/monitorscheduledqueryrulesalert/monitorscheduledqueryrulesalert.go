package monitorscheduledqueryrulesalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/monitorscheduledqueryrulesalert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert azurerm_monitor_scheduled_query_rules_alert}.
type MonitorScheduledQueryRulesAlert interface {
	cdktf.TerraformResource
	Action() MonitorScheduledQueryRulesAlertActionOutputReference
	ActionInput() *MonitorScheduledQueryRulesAlertAction
	AuthorizedResourceIds() *[]*string
	SetAuthorizedResourceIds(val *[]*string)
	AuthorizedResourceIdsInput() *[]*string
	AutoMitigationEnabled() interface{}
	SetAutoMitigationEnabled(val interface{})
	AutoMitigationEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DataSourceId() *string
	SetDataSourceId(val *string)
	DataSourceIdInput() *string
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
	Frequency() *float64
	SetFrequency(val *float64)
	FrequencyInput() *float64
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
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
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	QueryType() *string
	SetQueryType(val *string)
	QueryTypeInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Severity() *float64
	SetSeverity(val *float64)
	SeverityInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Throttling() *float64
	SetThrottling(val *float64)
	ThrottlingInput() *float64
	Timeouts() MonitorScheduledQueryRulesAlertTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeWindow() *float64
	SetTimeWindow(val *float64)
	TimeWindowInput() *float64
	Trigger() MonitorScheduledQueryRulesAlertTriggerOutputReference
	TriggerInput() *MonitorScheduledQueryRulesAlertTrigger
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
	PutAction(value *MonitorScheduledQueryRulesAlertAction)
	PutTimeouts(value *MonitorScheduledQueryRulesAlertTimeouts)
	PutTrigger(value *MonitorScheduledQueryRulesAlertTrigger)
	ResetAuthorizedResourceIds()
	ResetAutoMitigationEnabled()
	ResetDescription()
	ResetEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueryType()
	ResetSeverity()
	ResetTags()
	ResetThrottling()
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

// The jsii proxy struct for MonitorScheduledQueryRulesAlert
type jsiiProxy_MonitorScheduledQueryRulesAlert struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Action() MonitorScheduledQueryRulesAlertActionOutputReference {
	var returns MonitorScheduledQueryRulesAlertActionOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) ActionInput() *MonitorScheduledQueryRulesAlertAction {
	var returns *MonitorScheduledQueryRulesAlertAction
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) AuthorizedResourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedResourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) AuthorizedResourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedResourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) AutoMitigationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMitigationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) AutoMitigationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMitigationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) DataSourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Frequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) FrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) QueryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) QueryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Severity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SeverityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Throttling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) ThrottlingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Timeouts() MonitorScheduledQueryRulesAlertTimeoutsOutputReference {
	var returns MonitorScheduledQueryRulesAlertTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) TimeWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) TimeWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) Trigger() MonitorScheduledQueryRulesAlertTriggerOutputReference {
	var returns MonitorScheduledQueryRulesAlertTriggerOutputReference
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) TriggerInput() *MonitorScheduledQueryRulesAlertTrigger {
	var returns *MonitorScheduledQueryRulesAlertTrigger
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert azurerm_monitor_scheduled_query_rules_alert} Resource.
func NewMonitorScheduledQueryRulesAlert(scope constructs.Construct, id *string, config *MonitorScheduledQueryRulesAlertConfig) MonitorScheduledQueryRulesAlert {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlert{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert azurerm_monitor_scheduled_query_rules_alert} Resource.
func NewMonitorScheduledQueryRulesAlert_Override(m MonitorScheduledQueryRulesAlert, scope constructs.Construct, id *string, config *MonitorScheduledQueryRulesAlertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlert",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetAuthorizedResourceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"authorizedResourceIds",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetAutoMitigationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"autoMitigationEnabled",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetDataSourceId(val *string) {
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetFrequency(val *float64) {
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetQuery(val *string) {
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetQueryType(val *string) {
	_jsii_.Set(
		j,
		"queryType",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetSeverity(val *float64) {
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetThrottling(val *float64) {
	_jsii_.Set(
		j,
		"throttling",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlert) SetTimeWindow(val *float64) {
	_jsii_.Set(
		j,
		"timeWindow",
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
func MonitorScheduledQueryRulesAlert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorScheduledQueryRulesAlert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) PutAction(value *MonitorScheduledQueryRulesAlertAction) {
	_jsii_.InvokeVoid(
		m,
		"putAction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) PutTimeouts(value *MonitorScheduledQueryRulesAlertTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) PutTrigger(value *MonitorScheduledQueryRulesAlertTrigger) {
	_jsii_.InvokeVoid(
		m,
		"putTrigger",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetAuthorizedResourceIds() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthorizedResourceIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetAutoMitigationEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoMitigationEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetQueryType() {
	_jsii_.InvokeVoid(
		m,
		"resetQueryType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetSeverity() {
	_jsii_.InvokeVoid(
		m,
		"resetSeverity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetThrottling() {
	_jsii_.InvokeVoid(
		m,
		"resetThrottling",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#action_group MonitorScheduledQueryRulesAlert#action_group}.
	ActionGroup *[]*string `field:"required" json:"actionGroup" yaml:"actionGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#custom_webhook_payload MonitorScheduledQueryRulesAlert#custom_webhook_payload}.
	CustomWebhookPayload *string `field:"optional" json:"customWebhookPayload" yaml:"customWebhookPayload"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#email_subject MonitorScheduledQueryRulesAlert#email_subject}.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
}

type MonitorScheduledQueryRulesAlertActionOutputReference interface {
	cdktf.ComplexObject
	ActionGroup() *[]*string
	SetActionGroup(val *[]*string)
	ActionGroupInput() *[]*string
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
	CustomWebhookPayload() *string
	SetCustomWebhookPayload(val *string)
	CustomWebhookPayloadInput() *string
	EmailSubject() *string
	SetEmailSubject(val *string)
	EmailSubjectInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorScheduledQueryRulesAlertAction
	SetInternalValue(val *MonitorScheduledQueryRulesAlertAction)
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
	ResetCustomWebhookPayload()
	ResetEmailSubject()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorScheduledQueryRulesAlertActionOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) ActionGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) ActionGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) CustomWebhookPayload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customWebhookPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) CustomWebhookPayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customWebhookPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) EmailSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) EmailSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) InternalValue() *MonitorScheduledQueryRulesAlertAction {
	var returns *MonitorScheduledQueryRulesAlertAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorScheduledQueryRulesAlertActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlertActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertActionOutputReference_Override(m MonitorScheduledQueryRulesAlertActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlertActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) SetActionGroup(val *[]*string) {
	_jsii_.Set(
		j,
		"actionGroup",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) SetCustomWebhookPayload(val *string) {
	_jsii_.Set(
		j,
		"customWebhookPayload",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) SetEmailSubject(val *string) {
	_jsii_.Set(
		j,
		"emailSubject",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) SetInternalValue(val *MonitorScheduledQueryRulesAlertAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) ResetCustomWebhookPayload() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomWebhookPayload",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) ResetEmailSubject() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailSubject",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertConfig struct {
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
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#action MonitorScheduledQueryRulesAlert#action}
	Action *MonitorScheduledQueryRulesAlertAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#data_source_id MonitorScheduledQueryRulesAlert#data_source_id}.
	DataSourceId *string `field:"required" json:"dataSourceId" yaml:"dataSourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#frequency MonitorScheduledQueryRulesAlert#frequency}.
	Frequency *float64 `field:"required" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#location MonitorScheduledQueryRulesAlert#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#name MonitorScheduledQueryRulesAlert#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#query MonitorScheduledQueryRulesAlert#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#resource_group_name MonitorScheduledQueryRulesAlert#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#time_window MonitorScheduledQueryRulesAlert#time_window}.
	TimeWindow *float64 `field:"required" json:"timeWindow" yaml:"timeWindow"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#trigger MonitorScheduledQueryRulesAlert#trigger}
	Trigger *MonitorScheduledQueryRulesAlertTrigger `field:"required" json:"trigger" yaml:"trigger"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#authorized_resource_ids MonitorScheduledQueryRulesAlert#authorized_resource_ids}.
	AuthorizedResourceIds *[]*string `field:"optional" json:"authorizedResourceIds" yaml:"authorizedResourceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#auto_mitigation_enabled MonitorScheduledQueryRulesAlert#auto_mitigation_enabled}.
	AutoMitigationEnabled interface{} `field:"optional" json:"autoMitigationEnabled" yaml:"autoMitigationEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#description MonitorScheduledQueryRulesAlert#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#enabled MonitorScheduledQueryRulesAlert#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#id MonitorScheduledQueryRulesAlert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#query_type MonitorScheduledQueryRulesAlert#query_type}.
	QueryType *string `field:"optional" json:"queryType" yaml:"queryType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#severity MonitorScheduledQueryRulesAlert#severity}.
	Severity *float64 `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#tags MonitorScheduledQueryRulesAlert#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#throttling MonitorScheduledQueryRulesAlert#throttling}.
	Throttling *float64 `field:"optional" json:"throttling" yaml:"throttling"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#timeouts MonitorScheduledQueryRulesAlert#timeouts}
	Timeouts *MonitorScheduledQueryRulesAlertTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type MonitorScheduledQueryRulesAlertTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#create MonitorScheduledQueryRulesAlert#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#delete MonitorScheduledQueryRulesAlert#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#read MonitorScheduledQueryRulesAlert#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#update MonitorScheduledQueryRulesAlert#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type MonitorScheduledQueryRulesAlertTimeoutsOutputReference interface {
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

// The jsii proxy struct for MonitorScheduledQueryRulesAlertTimeoutsOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorScheduledQueryRulesAlertTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlertTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertTimeoutsOutputReference_Override(m MonitorScheduledQueryRulesAlertTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlertTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertTrigger struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#operator MonitorScheduledQueryRulesAlert#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#threshold MonitorScheduledQueryRulesAlert#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// metric_trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#metric_trigger MonitorScheduledQueryRulesAlert#metric_trigger}
	MetricTrigger *MonitorScheduledQueryRulesAlertTriggerMetricTrigger `field:"optional" json:"metricTrigger" yaml:"metricTrigger"`
}

type MonitorScheduledQueryRulesAlertTriggerMetricTrigger struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#metric_column MonitorScheduledQueryRulesAlert#metric_column}.
	MetricColumn *string `field:"required" json:"metricColumn" yaml:"metricColumn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#metric_trigger_type MonitorScheduledQueryRulesAlert#metric_trigger_type}.
	MetricTriggerType *string `field:"required" json:"metricTriggerType" yaml:"metricTriggerType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#operator MonitorScheduledQueryRulesAlert#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#threshold MonitorScheduledQueryRulesAlert#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

type MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference interface {
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
	InternalValue() *MonitorScheduledQueryRulesAlertTriggerMetricTrigger
	SetInternalValue(val *MonitorScheduledQueryRulesAlertTriggerMetricTrigger)
	MetricColumn() *string
	SetMetricColumn(val *string)
	MetricColumnInput() *string
	MetricTriggerType() *string
	SetMetricTriggerType(val *string)
	MetricTriggerTypeInput() *string
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
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
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

// The jsii proxy struct for MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) InternalValue() *MonitorScheduledQueryRulesAlertTriggerMetricTrigger {
	var returns *MonitorScheduledQueryRulesAlertTriggerMetricTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) MetricColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) MetricColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) MetricTriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTriggerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) MetricTriggerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTriggerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference_Override(m MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetInternalValue(val *MonitorScheduledQueryRulesAlertTriggerMetricTrigger) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetMetricColumn(val *string) {
	_jsii_.Set(
		j,
		"metricColumn",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetMetricTriggerType(val *string) {
	_jsii_.Set(
		j,
		"metricTriggerType",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) SetThreshold(val *float64) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertTriggerOutputReference interface {
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
	InternalValue() *MonitorScheduledQueryRulesAlertTrigger
	SetInternalValue(val *MonitorScheduledQueryRulesAlertTrigger)
	MetricTrigger() MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference
	MetricTriggerInput() *MonitorScheduledQueryRulesAlertTriggerMetricTrigger
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
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
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
	PutMetricTrigger(value *MonitorScheduledQueryRulesAlertTriggerMetricTrigger)
	ResetMetricTrigger()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorScheduledQueryRulesAlertTriggerOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) InternalValue() *MonitorScheduledQueryRulesAlertTrigger {
	var returns *MonitorScheduledQueryRulesAlertTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) MetricTrigger() MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference {
	var returns MonitorScheduledQueryRulesAlertTriggerMetricTriggerOutputReference
	_jsii_.Get(
		j,
		"metricTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) MetricTriggerInput() *MonitorScheduledQueryRulesAlertTriggerMetricTrigger {
	var returns *MonitorScheduledQueryRulesAlertTriggerMetricTrigger
	_jsii_.Get(
		j,
		"metricTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorScheduledQueryRulesAlertTriggerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlertTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertTriggerOutputReference_Override(m MonitorScheduledQueryRulesAlertTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlert.MonitorScheduledQueryRulesAlertTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) SetInternalValue(val *MonitorScheduledQueryRulesAlertTrigger) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) SetThreshold(val *float64) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) PutMetricTrigger(value *MonitorScheduledQueryRulesAlertTriggerMetricTrigger) {
	_jsii_.InvokeVoid(
		m,
		"putMetricTrigger",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) ResetMetricTrigger() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricTrigger",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

