package monitorscheduledqueryrulesalertv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/monitorscheduledqueryrulesalertv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2 azurerm_monitor_scheduled_query_rules_alert_v2}.
type MonitorScheduledQueryRulesAlertV2 interface {
	cdktf.TerraformResource
	Action() MonitorScheduledQueryRulesAlertV2ActionOutputReference
	ActionInput() *MonitorScheduledQueryRulesAlertV2Action
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
	CreatedWithApiVersion() *string
	Criteria() MonitorScheduledQueryRulesAlertV2CriteriaList
	CriteriaInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EvaluationFrequency() *string
	SetEvaluationFrequency(val *string)
	EvaluationFrequencyInput() *string
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
	IsALegacyLogAnalyticsRule() cdktf.IResolvable
	IsWorkspaceAlertsStorageConfigured() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MuteActionsAfterAlertDuration() *string
	SetMuteActionsAfterAlertDuration(val *string)
	MuteActionsAfterAlertDurationInput() *string
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
	QueryTimeRangeOverride() *string
	SetQueryTimeRangeOverride(val *string)
	QueryTimeRangeOverrideInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Severity() *float64
	SetSeverity(val *float64)
	SeverityInput() *float64
	SkipQueryValidation() interface{}
	SetSkipQueryValidation(val interface{})
	SkipQueryValidationInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TargetResourceTypes() *[]*string
	SetTargetResourceTypes(val *[]*string)
	TargetResourceTypesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	WindowDuration() *string
	SetWindowDuration(val *string)
	WindowDurationInput() *string
	WorkspaceAlertsStorageEnabled() interface{}
	SetWorkspaceAlertsStorageEnabled(val interface{})
	WorkspaceAlertsStorageEnabledInput() interface{}
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
	PutAction(value *MonitorScheduledQueryRulesAlertV2Action)
	PutCriteria(value interface{})
	PutTimeouts(value *MonitorScheduledQueryRulesAlertV2Timeouts)
	ResetAction()
	ResetAutoMitigationEnabled()
	ResetDescription()
	ResetDisplayName()
	ResetEnabled()
	ResetEvaluationFrequency()
	ResetId()
	ResetMuteActionsAfterAlertDuration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueryTimeRangeOverride()
	ResetSkipQueryValidation()
	ResetTags()
	ResetTargetResourceTypes()
	ResetTimeouts()
	ResetWorkspaceAlertsStorageEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2
type jsiiProxy_MonitorScheduledQueryRulesAlertV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Action() MonitorScheduledQueryRulesAlertV2ActionOutputReference {
	var returns MonitorScheduledQueryRulesAlertV2ActionOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ActionInput() *MonitorScheduledQueryRulesAlertV2Action {
	var returns *MonitorScheduledQueryRulesAlertV2Action
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) AutoMitigationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMitigationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) AutoMitigationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMitigationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) CreatedWithApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdWithApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Criteria() MonitorScheduledQueryRulesAlertV2CriteriaList {
	var returns MonitorScheduledQueryRulesAlertV2CriteriaList
	_jsii_.Get(
		j,
		"criteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) CriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) EvaluationFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) EvaluationFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) IsALegacyLogAnalyticsRule() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isALegacyLogAnalyticsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) IsWorkspaceAlertsStorageConfigured() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isWorkspaceAlertsStorageConfigured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) MuteActionsAfterAlertDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"muteActionsAfterAlertDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) MuteActionsAfterAlertDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"muteActionsAfterAlertDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) QueryTimeRangeOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTimeRangeOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) QueryTimeRangeOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTimeRangeOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Severity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SeverityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SkipQueryValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipQueryValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SkipQueryValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipQueryValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) TargetResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) TargetResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Timeouts() MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference {
	var returns MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) WindowDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) WindowDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) WorkspaceAlertsStorageEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceAlertsStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) WorkspaceAlertsStorageEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceAlertsStorageEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2 azurerm_monitor_scheduled_query_rules_alert_v2} Resource.
func NewMonitorScheduledQueryRulesAlertV2(scope constructs.Construct, id *string, config *MonitorScheduledQueryRulesAlertV2Config) MonitorScheduledQueryRulesAlertV2 {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2 azurerm_monitor_scheduled_query_rules_alert_v2} Resource.
func NewMonitorScheduledQueryRulesAlertV2_Override(m MonitorScheduledQueryRulesAlertV2, scope constructs.Construct, id *string, config *MonitorScheduledQueryRulesAlertV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetAutoMitigationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"autoMitigationEnabled",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetEvaluationFrequency(val *string) {
	_jsii_.Set(
		j,
		"evaluationFrequency",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetMuteActionsAfterAlertDuration(val *string) {
	_jsii_.Set(
		j,
		"muteActionsAfterAlertDuration",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetQueryTimeRangeOverride(val *string) {
	_jsii_.Set(
		j,
		"queryTimeRangeOverride",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetSeverity(val *float64) {
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetSkipQueryValidation(val interface{}) {
	_jsii_.Set(
		j,
		"skipQueryValidation",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetTargetResourceTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"targetResourceTypes",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetWindowDuration(val *string) {
	_jsii_.Set(
		j,
		"windowDuration",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SetWorkspaceAlertsStorageEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"workspaceAlertsStorageEnabled",
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
func MonitorScheduledQueryRulesAlertV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorScheduledQueryRulesAlertV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) PutAction(value *MonitorScheduledQueryRulesAlertV2Action) {
	_jsii_.InvokeVoid(
		m,
		"putAction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) PutCriteria(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putCriteria",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) PutTimeouts(value *MonitorScheduledQueryRulesAlertV2Timeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetAction() {
	_jsii_.InvokeVoid(
		m,
		"resetAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetAutoMitigationEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoMitigationEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetDisplayName() {
	_jsii_.InvokeVoid(
		m,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetEvaluationFrequency() {
	_jsii_.InvokeVoid(
		m,
		"resetEvaluationFrequency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetMuteActionsAfterAlertDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetMuteActionsAfterAlertDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetQueryTimeRangeOverride() {
	_jsii_.InvokeVoid(
		m,
		"resetQueryTimeRangeOverride",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetSkipQueryValidation() {
	_jsii_.InvokeVoid(
		m,
		"resetSkipQueryValidation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetTargetResourceTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ResetWorkspaceAlertsStorageEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkspaceAlertsStorageEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertV2Action struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#action_groups MonitorScheduledQueryRulesAlertV2#action_groups}.
	ActionGroups *[]*string `field:"optional" json:"actionGroups" yaml:"actionGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#custom_properties MonitorScheduledQueryRulesAlertV2#custom_properties}.
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
}

type MonitorScheduledQueryRulesAlertV2ActionOutputReference interface {
	cdktf.ComplexObject
	ActionGroups() *[]*string
	SetActionGroups(val *[]*string)
	ActionGroupsInput() *[]*string
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
	CustomProperties() *map[string]*string
	SetCustomProperties(val *map[string]*string)
	CustomPropertiesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorScheduledQueryRulesAlertV2Action
	SetInternalValue(val *MonitorScheduledQueryRulesAlertV2Action)
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
	ResetActionGroups()
	ResetCustomProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2ActionOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) ActionGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) ActionGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) CustomProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) CustomPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) InternalValue() *MonitorScheduledQueryRulesAlertV2Action {
	var returns *MonitorScheduledQueryRulesAlertV2Action
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertV2ActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorScheduledQueryRulesAlertV2ActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2ActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertV2ActionOutputReference_Override(m MonitorScheduledQueryRulesAlertV2ActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2ActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) SetActionGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"actionGroups",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) SetCustomProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"customProperties",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) SetInternalValue(val *MonitorScheduledQueryRulesAlertV2Action) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) ResetActionGroups() {
	_jsii_.InvokeVoid(
		m,
		"resetActionGroups",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) ResetCustomProperties() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomProperties",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2ActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertV2Config struct {
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
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#criteria MonitorScheduledQueryRulesAlertV2#criteria}
	Criteria interface{} `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#location MonitorScheduledQueryRulesAlertV2#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#name MonitorScheduledQueryRulesAlertV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#resource_group_name MonitorScheduledQueryRulesAlertV2#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#scopes MonitorScheduledQueryRulesAlertV2#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#severity MonitorScheduledQueryRulesAlertV2#severity}.
	Severity *float64 `field:"required" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#window_duration MonitorScheduledQueryRulesAlertV2#window_duration}.
	WindowDuration *string `field:"required" json:"windowDuration" yaml:"windowDuration"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#action MonitorScheduledQueryRulesAlertV2#action}
	Action *MonitorScheduledQueryRulesAlertV2Action `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#auto_mitigation_enabled MonitorScheduledQueryRulesAlertV2#auto_mitigation_enabled}.
	AutoMitigationEnabled interface{} `field:"optional" json:"autoMitigationEnabled" yaml:"autoMitigationEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#description MonitorScheduledQueryRulesAlertV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#display_name MonitorScheduledQueryRulesAlertV2#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#enabled MonitorScheduledQueryRulesAlertV2#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#evaluation_frequency MonitorScheduledQueryRulesAlertV2#evaluation_frequency}.
	EvaluationFrequency *string `field:"optional" json:"evaluationFrequency" yaml:"evaluationFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#id MonitorScheduledQueryRulesAlertV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#mute_actions_after_alert_duration MonitorScheduledQueryRulesAlertV2#mute_actions_after_alert_duration}.
	MuteActionsAfterAlertDuration *string `field:"optional" json:"muteActionsAfterAlertDuration" yaml:"muteActionsAfterAlertDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#query_time_range_override MonitorScheduledQueryRulesAlertV2#query_time_range_override}.
	QueryTimeRangeOverride *string `field:"optional" json:"queryTimeRangeOverride" yaml:"queryTimeRangeOverride"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#skip_query_validation MonitorScheduledQueryRulesAlertV2#skip_query_validation}.
	SkipQueryValidation interface{} `field:"optional" json:"skipQueryValidation" yaml:"skipQueryValidation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#tags MonitorScheduledQueryRulesAlertV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#target_resource_types MonitorScheduledQueryRulesAlertV2#target_resource_types}.
	TargetResourceTypes *[]*string `field:"optional" json:"targetResourceTypes" yaml:"targetResourceTypes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#timeouts MonitorScheduledQueryRulesAlertV2#timeouts}
	Timeouts *MonitorScheduledQueryRulesAlertV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#workspace_alerts_storage_enabled MonitorScheduledQueryRulesAlertV2#workspace_alerts_storage_enabled}.
	WorkspaceAlertsStorageEnabled interface{} `field:"optional" json:"workspaceAlertsStorageEnabled" yaml:"workspaceAlertsStorageEnabled"`
}

type MonitorScheduledQueryRulesAlertV2Criteria struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#operator MonitorScheduledQueryRulesAlertV2#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#query MonitorScheduledQueryRulesAlertV2#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#threshold MonitorScheduledQueryRulesAlertV2#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#time_aggregation_method MonitorScheduledQueryRulesAlertV2#time_aggregation_method}.
	TimeAggregationMethod *string `field:"required" json:"timeAggregationMethod" yaml:"timeAggregationMethod"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#dimension MonitorScheduledQueryRulesAlertV2#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// failing_periods block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#failing_periods MonitorScheduledQueryRulesAlertV2#failing_periods}
	FailingPeriods *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods `field:"optional" json:"failingPeriods" yaml:"failingPeriods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#metric_measure_column MonitorScheduledQueryRulesAlertV2#metric_measure_column}.
	MetricMeasureColumn *string `field:"optional" json:"metricMeasureColumn" yaml:"metricMeasureColumn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#resource_id_column MonitorScheduledQueryRulesAlertV2#resource_id_column}.
	ResourceIdColumn *string `field:"optional" json:"resourceIdColumn" yaml:"resourceIdColumn"`
}

type MonitorScheduledQueryRulesAlertV2CriteriaDimension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#name MonitorScheduledQueryRulesAlertV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#operator MonitorScheduledQueryRulesAlertV2#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#values MonitorScheduledQueryRulesAlertV2#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type MonitorScheduledQueryRulesAlertV2CriteriaDimensionList interface {
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
	Get(index *float64) MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2CriteriaDimensionList
type jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertV2CriteriaDimensionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorScheduledQueryRulesAlertV2CriteriaDimensionList {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaDimensionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertV2CriteriaDimensionList_Override(m MonitorScheduledQueryRulesAlertV2CriteriaDimensionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaDimensionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) Get(index *float64) MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference {
	var returns MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
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

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference_Override(m MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaDimensionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#minimum_failing_periods_to_trigger_alert MonitorScheduledQueryRulesAlertV2#minimum_failing_periods_to_trigger_alert}.
	MinimumFailingPeriodsToTriggerAlert *float64 `field:"required" json:"minimumFailingPeriodsToTriggerAlert" yaml:"minimumFailingPeriodsToTriggerAlert"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#number_of_evaluation_periods MonitorScheduledQueryRulesAlertV2#number_of_evaluation_periods}.
	NumberOfEvaluationPeriods *float64 `field:"required" json:"numberOfEvaluationPeriods" yaml:"numberOfEvaluationPeriods"`
}

type MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference interface {
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
	InternalValue() *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods
	SetInternalValue(val *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods)
	MinimumFailingPeriodsToTriggerAlert() *float64
	SetMinimumFailingPeriodsToTriggerAlert(val *float64)
	MinimumFailingPeriodsToTriggerAlertInput() *float64
	NumberOfEvaluationPeriods() *float64
	SetNumberOfEvaluationPeriods(val *float64)
	NumberOfEvaluationPeriodsInput() *float64
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

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) InternalValue() *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods {
	var returns *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) MinimumFailingPeriodsToTriggerAlert() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumFailingPeriodsToTriggerAlert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) MinimumFailingPeriodsToTriggerAlertInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumFailingPeriodsToTriggerAlertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) NumberOfEvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfEvaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) NumberOfEvaluationPeriodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfEvaluationPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference_Override(m MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) SetInternalValue(val *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) SetMinimumFailingPeriodsToTriggerAlert(val *float64) {
	_jsii_.Set(
		j,
		"minimumFailingPeriodsToTriggerAlert",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) SetNumberOfEvaluationPeriods(val *float64) {
	_jsii_.Set(
		j,
		"numberOfEvaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertV2CriteriaList interface {
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
	Get(index *float64) MonitorScheduledQueryRulesAlertV2CriteriaOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2CriteriaList
type jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertV2CriteriaList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorScheduledQueryRulesAlertV2CriteriaList {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertV2CriteriaList_Override(m MonitorScheduledQueryRulesAlertV2CriteriaList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) Get(index *float64) MonitorScheduledQueryRulesAlertV2CriteriaOutputReference {
	var returns MonitorScheduledQueryRulesAlertV2CriteriaOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertV2CriteriaOutputReference interface {
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
	Dimension() MonitorScheduledQueryRulesAlertV2CriteriaDimensionList
	DimensionInput() interface{}
	FailingPeriods() MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference
	FailingPeriodsInput() *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricMeasureColumn() *string
	SetMetricMeasureColumn(val *string)
	MetricMeasureColumnInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	ResourceIdColumn() *string
	SetResourceIdColumn(val *string)
	ResourceIdColumnInput() *string
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
	TimeAggregationMethod() *string
	SetTimeAggregationMethod(val *string)
	TimeAggregationMethodInput() *string
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
	PutDimension(value interface{})
	PutFailingPeriods(value *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods)
	ResetDimension()
	ResetFailingPeriods()
	ResetMetricMeasureColumn()
	ResetResourceIdColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2CriteriaOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Dimension() MonitorScheduledQueryRulesAlertV2CriteriaDimensionList {
	var returns MonitorScheduledQueryRulesAlertV2CriteriaDimensionList
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) DimensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) FailingPeriods() MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference {
	var returns MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriodsOutputReference
	_jsii_.Get(
		j,
		"failingPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) FailingPeriodsInput() *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods {
	var returns *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods
	_jsii_.Get(
		j,
		"failingPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) MetricMeasureColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricMeasureColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) MetricMeasureColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricMeasureColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResourceIdColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResourceIdColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) TimeAggregationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeAggregationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) TimeAggregationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeAggregationMethodInput",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertV2CriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorScheduledQueryRulesAlertV2CriteriaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertV2CriteriaOutputReference_Override(m MonitorScheduledQueryRulesAlertV2CriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2CriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetMetricMeasureColumn(val *string) {
	_jsii_.Set(
		j,
		"metricMeasureColumn",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetQuery(val *string) {
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetResourceIdColumn(val *string) {
	_jsii_.Set(
		j,
		"resourceIdColumn",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetThreshold(val *float64) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) SetTimeAggregationMethod(val *string) {
	_jsii_.Set(
		j,
		"timeAggregationMethod",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) PutDimension(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putDimension",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) PutFailingPeriods(value *MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods) {
	_jsii_.InvokeVoid(
		m,
		"putFailingPeriods",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		m,
		"resetDimension",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResetFailingPeriods() {
	_jsii_.InvokeVoid(
		m,
		"resetFailingPeriods",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResetMetricMeasureColumn() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricMeasureColumn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ResetResourceIdColumn() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceIdColumn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2CriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorScheduledQueryRulesAlertV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#create MonitorScheduledQueryRulesAlertV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#delete MonitorScheduledQueryRulesAlertV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#read MonitorScheduledQueryRulesAlertV2#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert_v2#update MonitorScheduledQueryRulesAlertV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference interface {
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

// The jsii proxy struct for MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference
type jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewMonitorScheduledQueryRulesAlertV2TimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorScheduledQueryRulesAlertV2TimeoutsOutputReference_Override(m MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2TimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

