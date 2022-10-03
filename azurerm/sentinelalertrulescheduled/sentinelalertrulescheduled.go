package sentinelalertrulescheduled

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/sentinelalertrulescheduled/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled azurerm_sentinel_alert_rule_scheduled}.
type SentinelAlertRuleScheduled interface {
	cdktf.TerraformResource
	AlertDetailsOverride() SentinelAlertRuleScheduledAlertDetailsOverrideList
	AlertDetailsOverrideInput() interface{}
	AlertRuleTemplateGuid() *string
	SetAlertRuleTemplateGuid(val *string)
	AlertRuleTemplateGuidInput() *string
	AlertRuleTemplateVersion() *string
	SetAlertRuleTemplateVersion(val *string)
	AlertRuleTemplateVersionInput() *string
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
	CustomDetails() *map[string]*string
	SetCustomDetails(val *map[string]*string)
	CustomDetailsInput() *map[string]*string
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
	EntityMapping() SentinelAlertRuleScheduledEntityMappingList
	EntityMappingInput() interface{}
	EventGrouping() SentinelAlertRuleScheduledEventGroupingOutputReference
	EventGroupingInput() *SentinelAlertRuleScheduledEventGrouping
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
	IncidentConfiguration() SentinelAlertRuleScheduledIncidentConfigurationOutputReference
	IncidentConfigurationInput() *SentinelAlertRuleScheduledIncidentConfiguration
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
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
	QueryFrequency() *string
	SetQueryFrequency(val *string)
	QueryFrequencyInput() *string
	QueryInput() *string
	QueryPeriod() *string
	SetQueryPeriod(val *string)
	QueryPeriodInput() *string
	// Experimental.
	RawOverrides() interface{}
	Severity() *string
	SetSeverity(val *string)
	SeverityInput() *string
	SuppressionDuration() *string
	SetSuppressionDuration(val *string)
	SuppressionDurationInput() *string
	SuppressionEnabled() interface{}
	SetSuppressionEnabled(val interface{})
	SuppressionEnabledInput() interface{}
	Tactics() *[]*string
	SetTactics(val *[]*string)
	TacticsInput() *[]*string
	Techniques() *[]*string
	SetTechniques(val *[]*string)
	TechniquesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SentinelAlertRuleScheduledTimeoutsOutputReference
	TimeoutsInput() interface{}
	TriggerOperator() *string
	SetTriggerOperator(val *string)
	TriggerOperatorInput() *string
	TriggerThreshold() *float64
	SetTriggerThreshold(val *float64)
	TriggerThresholdInput() *float64
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
	PutAlertDetailsOverride(value interface{})
	PutEntityMapping(value interface{})
	PutEventGrouping(value *SentinelAlertRuleScheduledEventGrouping)
	PutIncidentConfiguration(value *SentinelAlertRuleScheduledIncidentConfiguration)
	PutTimeouts(value *SentinelAlertRuleScheduledTimeouts)
	ResetAlertDetailsOverride()
	ResetAlertRuleTemplateGuid()
	ResetAlertRuleTemplateVersion()
	ResetCustomDetails()
	ResetDescription()
	ResetEnabled()
	ResetEntityMapping()
	ResetEventGrouping()
	ResetId()
	ResetIncidentConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueryFrequency()
	ResetQueryPeriod()
	ResetSuppressionDuration()
	ResetSuppressionEnabled()
	ResetTactics()
	ResetTechniques()
	ResetTimeouts()
	ResetTriggerOperator()
	ResetTriggerThreshold()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SentinelAlertRuleScheduled
type jsiiProxy_SentinelAlertRuleScheduled struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) AlertDetailsOverride() SentinelAlertRuleScheduledAlertDetailsOverrideList {
	var returns SentinelAlertRuleScheduledAlertDetailsOverrideList
	_jsii_.Get(
		j,
		"alertDetailsOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) AlertDetailsOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertDetailsOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) AlertRuleTemplateGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertRuleTemplateGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) AlertRuleTemplateGuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertRuleTemplateGuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) AlertRuleTemplateVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertRuleTemplateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) AlertRuleTemplateVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertRuleTemplateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) CustomDetails() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) CustomDetailsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) EntityMapping() SentinelAlertRuleScheduledEntityMappingList {
	var returns SentinelAlertRuleScheduledEntityMappingList
	_jsii_.Get(
		j,
		"entityMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) EntityMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) EventGrouping() SentinelAlertRuleScheduledEventGroupingOutputReference {
	var returns SentinelAlertRuleScheduledEventGroupingOutputReference
	_jsii_.Get(
		j,
		"eventGrouping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) EventGroupingInput() *SentinelAlertRuleScheduledEventGrouping {
	var returns *SentinelAlertRuleScheduledEventGrouping
	_jsii_.Get(
		j,
		"eventGroupingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) IncidentConfiguration() SentinelAlertRuleScheduledIncidentConfigurationOutputReference {
	var returns SentinelAlertRuleScheduledIncidentConfigurationOutputReference
	_jsii_.Get(
		j,
		"incidentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) IncidentConfigurationInput() *SentinelAlertRuleScheduledIncidentConfiguration {
	var returns *SentinelAlertRuleScheduledIncidentConfiguration
	_jsii_.Get(
		j,
		"incidentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) QueryFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) QueryFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) QueryPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) QueryPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SuppressionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suppressionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SuppressionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suppressionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SuppressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SuppressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Tactics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tactics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TacticsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tacticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Techniques() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"techniques",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TechniquesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"techniquesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) Timeouts() SentinelAlertRuleScheduledTimeoutsOutputReference {
	var returns SentinelAlertRuleScheduledTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TriggerOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TriggerOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TriggerThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"triggerThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) TriggerThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"triggerThresholdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled azurerm_sentinel_alert_rule_scheduled} Resource.
func NewSentinelAlertRuleScheduled(scope constructs.Construct, id *string, config *SentinelAlertRuleScheduledConfig) SentinelAlertRuleScheduled {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduled{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled azurerm_sentinel_alert_rule_scheduled} Resource.
func NewSentinelAlertRuleScheduled_Override(s SentinelAlertRuleScheduled, scope constructs.Construct, id *string, config *SentinelAlertRuleScheduledConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetAlertRuleTemplateGuid(val *string) {
	_jsii_.Set(
		j,
		"alertRuleTemplateGuid",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetAlertRuleTemplateVersion(val *string) {
	_jsii_.Set(
		j,
		"alertRuleTemplateVersion",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetCustomDetails(val *map[string]*string) {
	_jsii_.Set(
		j,
		"customDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetLogAnalyticsWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetQuery(val *string) {
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetQueryFrequency(val *string) {
	_jsii_.Set(
		j,
		"queryFrequency",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetQueryPeriod(val *string) {
	_jsii_.Set(
		j,
		"queryPeriod",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetSeverity(val *string) {
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetSuppressionDuration(val *string) {
	_jsii_.Set(
		j,
		"suppressionDuration",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetSuppressionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"suppressionEnabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetTactics(val *[]*string) {
	_jsii_.Set(
		j,
		"tactics",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetTechniques(val *[]*string) {
	_jsii_.Set(
		j,
		"techniques",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetTriggerOperator(val *string) {
	_jsii_.Set(
		j,
		"triggerOperator",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SetTriggerThreshold(val *float64) {
	_jsii_.Set(
		j,
		"triggerThreshold",
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
func SentinelAlertRuleScheduled_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SentinelAlertRuleScheduled_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutAlertDetailsOverride(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putAlertDetailsOverride",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutEntityMapping(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putEntityMapping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutEventGrouping(value *SentinelAlertRuleScheduledEventGrouping) {
	_jsii_.InvokeVoid(
		s,
		"putEventGrouping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutIncidentConfiguration(value *SentinelAlertRuleScheduledIncidentConfiguration) {
	_jsii_.InvokeVoid(
		s,
		"putIncidentConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutTimeouts(value *SentinelAlertRuleScheduledTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetAlertDetailsOverride() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertDetailsOverride",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetAlertRuleTemplateGuid() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertRuleTemplateGuid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetAlertRuleTemplateVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertRuleTemplateVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetCustomDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetEntityMapping() {
	_jsii_.InvokeVoid(
		s,
		"resetEntityMapping",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetEventGrouping() {
	_jsii_.InvokeVoid(
		s,
		"resetEventGrouping",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetIncidentConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetIncidentConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetQueryFrequency() {
	_jsii_.InvokeVoid(
		s,
		"resetQueryFrequency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetQueryPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetQueryPeriod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetSuppressionDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetSuppressionDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetSuppressionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSuppressionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetTactics() {
	_jsii_.InvokeVoid(
		s,
		"resetTactics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetTechniques() {
	_jsii_.InvokeVoid(
		s,
		"resetTechniques",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetTriggerOperator() {
	_jsii_.InvokeVoid(
		s,
		"resetTriggerOperator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetTriggerThreshold() {
	_jsii_.InvokeVoid(
		s,
		"resetTriggerThreshold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledAlertDetailsOverride struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#description_format SentinelAlertRuleScheduled#description_format}.
	DescriptionFormat *string `field:"optional" json:"descriptionFormat" yaml:"descriptionFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#display_name_format SentinelAlertRuleScheduled#display_name_format}.
	DisplayNameFormat *string `field:"optional" json:"displayNameFormat" yaml:"displayNameFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#severity_column_name SentinelAlertRuleScheduled#severity_column_name}.
	SeverityColumnName *string `field:"optional" json:"severityColumnName" yaml:"severityColumnName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#tactics_column_name SentinelAlertRuleScheduled#tactics_column_name}.
	TacticsColumnName *string `field:"optional" json:"tacticsColumnName" yaml:"tacticsColumnName"`
}

type SentinelAlertRuleScheduledAlertDetailsOverrideList interface {
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
	Get(index *float64) SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledAlertDetailsOverrideList
type jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledAlertDetailsOverrideList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SentinelAlertRuleScheduledAlertDetailsOverrideList {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledAlertDetailsOverrideList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledAlertDetailsOverrideList_Override(s SentinelAlertRuleScheduledAlertDetailsOverrideList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledAlertDetailsOverrideList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) Get(index *float64) SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference {
	var returns SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference interface {
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
	ResetDescriptionFormat()
	ResetDisplayNameFormat()
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

// The jsii proxy struct for SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference
type jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DescriptionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DescriptionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DisplayNameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) DisplayNameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SeverityColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SeverityColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) TacticsColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticsColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) TacticsColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticsColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledAlertDetailsOverrideOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledAlertDetailsOverrideOutputReference_Override(s SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetDescriptionFormat(val *string) {
	_jsii_.Set(
		j,
		"descriptionFormat",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetDisplayNameFormat(val *string) {
	_jsii_.Set(
		j,
		"displayNameFormat",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetSeverityColumnName(val *string) {
	_jsii_.Set(
		j,
		"severityColumnName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetTacticsColumnName(val *string) {
	_jsii_.Set(
		j,
		"tacticsColumnName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetDescriptionFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDescriptionFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetDisplayNameFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayNameFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetSeverityColumnName() {
	_jsii_.InvokeVoid(
		s,
		"resetSeverityColumnName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ResetTacticsColumnName() {
	_jsii_.InvokeVoid(
		s,
		"resetTacticsColumnName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledAlertDetailsOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#display_name SentinelAlertRuleScheduled#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#log_analytics_workspace_id SentinelAlertRuleScheduled#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#name SentinelAlertRuleScheduled#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#query SentinelAlertRuleScheduled#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#severity SentinelAlertRuleScheduled#severity}.
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// alert_details_override block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#alert_details_override SentinelAlertRuleScheduled#alert_details_override}
	AlertDetailsOverride interface{} `field:"optional" json:"alertDetailsOverride" yaml:"alertDetailsOverride"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#alert_rule_template_guid SentinelAlertRuleScheduled#alert_rule_template_guid}.
	AlertRuleTemplateGuid *string `field:"optional" json:"alertRuleTemplateGuid" yaml:"alertRuleTemplateGuid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#alert_rule_template_version SentinelAlertRuleScheduled#alert_rule_template_version}.
	AlertRuleTemplateVersion *string `field:"optional" json:"alertRuleTemplateVersion" yaml:"alertRuleTemplateVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#custom_details SentinelAlertRuleScheduled#custom_details}.
	CustomDetails *map[string]*string `field:"optional" json:"customDetails" yaml:"customDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#description SentinelAlertRuleScheduled#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#enabled SentinelAlertRuleScheduled#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// entity_mapping block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#entity_mapping SentinelAlertRuleScheduled#entity_mapping}
	EntityMapping interface{} `field:"optional" json:"entityMapping" yaml:"entityMapping"`
	// event_grouping block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#event_grouping SentinelAlertRuleScheduled#event_grouping}
	EventGrouping *SentinelAlertRuleScheduledEventGrouping `field:"optional" json:"eventGrouping" yaml:"eventGrouping"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#id SentinelAlertRuleScheduled#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// incident_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#incident_configuration SentinelAlertRuleScheduled#incident_configuration}
	IncidentConfiguration *SentinelAlertRuleScheduledIncidentConfiguration `field:"optional" json:"incidentConfiguration" yaml:"incidentConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#query_frequency SentinelAlertRuleScheduled#query_frequency}.
	QueryFrequency *string `field:"optional" json:"queryFrequency" yaml:"queryFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#query_period SentinelAlertRuleScheduled#query_period}.
	QueryPeriod *string `field:"optional" json:"queryPeriod" yaml:"queryPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#suppression_duration SentinelAlertRuleScheduled#suppression_duration}.
	SuppressionDuration *string `field:"optional" json:"suppressionDuration" yaml:"suppressionDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#suppression_enabled SentinelAlertRuleScheduled#suppression_enabled}.
	SuppressionEnabled interface{} `field:"optional" json:"suppressionEnabled" yaml:"suppressionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#tactics SentinelAlertRuleScheduled#tactics}.
	Tactics *[]*string `field:"optional" json:"tactics" yaml:"tactics"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#techniques SentinelAlertRuleScheduled#techniques}.
	Techniques *[]*string `field:"optional" json:"techniques" yaml:"techniques"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#timeouts SentinelAlertRuleScheduled#timeouts}
	Timeouts *SentinelAlertRuleScheduledTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#trigger_operator SentinelAlertRuleScheduled#trigger_operator}.
	TriggerOperator *string `field:"optional" json:"triggerOperator" yaml:"triggerOperator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#trigger_threshold SentinelAlertRuleScheduled#trigger_threshold}.
	TriggerThreshold *float64 `field:"optional" json:"triggerThreshold" yaml:"triggerThreshold"`
}

type SentinelAlertRuleScheduledEntityMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#entity_type SentinelAlertRuleScheduled#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#field_mapping SentinelAlertRuleScheduled#field_mapping}
	FieldMapping interface{} `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
}

type SentinelAlertRuleScheduledEntityMappingFieldMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#column_name SentinelAlertRuleScheduled#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#identifier SentinelAlertRuleScheduled#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

type SentinelAlertRuleScheduledEntityMappingFieldMappingList interface {
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
	Get(index *float64) SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledEntityMappingFieldMappingList
type jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledEntityMappingFieldMappingList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SentinelAlertRuleScheduledEntityMappingFieldMappingList {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEntityMappingFieldMappingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledEntityMappingFieldMappingList_Override(s SentinelAlertRuleScheduledEntityMappingFieldMappingList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEntityMappingFieldMappingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) Get(index *float64) SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference {
	var returns SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference interface {
	cdktf.ComplexObject
	ColumnName() *string
	SetColumnName(val *string)
	ColumnNameInput() *string
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
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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

// The jsii proxy struct for SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference
type jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) ColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) ColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference_Override(s SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) SetColumnName(val *string) {
	_jsii_.Set(
		j,
		"columnName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) SetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingFieldMappingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledEntityMappingList interface {
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
	Get(index *float64) SentinelAlertRuleScheduledEntityMappingOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledEntityMappingList
type jsiiProxy_SentinelAlertRuleScheduledEntityMappingList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledEntityMappingList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SentinelAlertRuleScheduledEntityMappingList {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledEntityMappingList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEntityMappingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledEntityMappingList_Override(s SentinelAlertRuleScheduledEntityMappingList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEntityMappingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) Get(index *float64) SentinelAlertRuleScheduledEntityMappingOutputReference {
	var returns SentinelAlertRuleScheduledEntityMappingOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledEntityMappingOutputReference interface {
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
	EntityType() *string
	SetEntityType(val *string)
	EntityTypeInput() *string
	FieldMapping() SentinelAlertRuleScheduledEntityMappingFieldMappingList
	FieldMappingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutFieldMapping(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledEntityMappingOutputReference
type jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) EntityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) EntityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) FieldMapping() SentinelAlertRuleScheduledEntityMappingFieldMappingList {
	var returns SentinelAlertRuleScheduledEntityMappingFieldMappingList
	_jsii_.Get(
		j,
		"fieldMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) FieldMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledEntityMappingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SentinelAlertRuleScheduledEntityMappingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEntityMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledEntityMappingOutputReference_Override(s SentinelAlertRuleScheduledEntityMappingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEntityMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) SetEntityType(val *string) {
	_jsii_.Set(
		j,
		"entityType",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) PutFieldMapping(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putFieldMapping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEntityMappingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledEventGrouping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#aggregation_method SentinelAlertRuleScheduled#aggregation_method}.
	AggregationMethod *string `field:"required" json:"aggregationMethod" yaml:"aggregationMethod"`
}

type SentinelAlertRuleScheduledEventGroupingOutputReference interface {
	cdktf.ComplexObject
	AggregationMethod() *string
	SetAggregationMethod(val *string)
	AggregationMethodInput() *string
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
	InternalValue() *SentinelAlertRuleScheduledEventGrouping
	SetInternalValue(val *SentinelAlertRuleScheduledEventGrouping)
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

// The jsii proxy struct for SentinelAlertRuleScheduledEventGroupingOutputReference
type jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) AggregationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) AggregationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) InternalValue() *SentinelAlertRuleScheduledEventGrouping {
	var returns *SentinelAlertRuleScheduledEventGrouping
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledEventGroupingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SentinelAlertRuleScheduledEventGroupingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEventGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledEventGroupingOutputReference_Override(s SentinelAlertRuleScheduledEventGroupingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledEventGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) SetAggregationMethod(val *string) {
	_jsii_.Set(
		j,
		"aggregationMethod",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) SetInternalValue(val *SentinelAlertRuleScheduledEventGrouping) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledEventGroupingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledIncidentConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#create_incident SentinelAlertRuleScheduled#create_incident}.
	CreateIncident interface{} `field:"required" json:"createIncident" yaml:"createIncident"`
	// grouping block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#grouping SentinelAlertRuleScheduled#grouping}
	Grouping *SentinelAlertRuleScheduledIncidentConfigurationGrouping `field:"required" json:"grouping" yaml:"grouping"`
}

type SentinelAlertRuleScheduledIncidentConfigurationGrouping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#enabled SentinelAlertRuleScheduled#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#entity_matching_method SentinelAlertRuleScheduled#entity_matching_method}.
	EntityMatchingMethod *string `field:"optional" json:"entityMatchingMethod" yaml:"entityMatchingMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#group_by_alert_details SentinelAlertRuleScheduled#group_by_alert_details}.
	GroupByAlertDetails *[]*string `field:"optional" json:"groupByAlertDetails" yaml:"groupByAlertDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#group_by_custom_details SentinelAlertRuleScheduled#group_by_custom_details}.
	GroupByCustomDetails *[]*string `field:"optional" json:"groupByCustomDetails" yaml:"groupByCustomDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#group_by_entities SentinelAlertRuleScheduled#group_by_entities}.
	GroupByEntities *[]*string `field:"optional" json:"groupByEntities" yaml:"groupByEntities"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#lookback_duration SentinelAlertRuleScheduled#lookback_duration}.
	LookbackDuration *string `field:"optional" json:"lookbackDuration" yaml:"lookbackDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#reopen_closed_incidents SentinelAlertRuleScheduled#reopen_closed_incidents}.
	ReopenClosedIncidents interface{} `field:"optional" json:"reopenClosedIncidents" yaml:"reopenClosedIncidents"`
}

type SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EntityMatchingMethod() *string
	SetEntityMatchingMethod(val *string)
	EntityMatchingMethodInput() *string
	// Experimental.
	Fqn() *string
	GroupByAlertDetails() *[]*string
	SetGroupByAlertDetails(val *[]*string)
	GroupByAlertDetailsInput() *[]*string
	GroupByCustomDetails() *[]*string
	SetGroupByCustomDetails(val *[]*string)
	GroupByCustomDetailsInput() *[]*string
	GroupByEntities() *[]*string
	SetGroupByEntities(val *[]*string)
	GroupByEntitiesInput() *[]*string
	InternalValue() *SentinelAlertRuleScheduledIncidentConfigurationGrouping
	SetInternalValue(val *SentinelAlertRuleScheduledIncidentConfigurationGrouping)
	LookbackDuration() *string
	SetLookbackDuration(val *string)
	LookbackDurationInput() *string
	ReopenClosedIncidents() interface{}
	SetReopenClosedIncidents(val interface{})
	ReopenClosedIncidentsInput() interface{}
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
	ResetEnabled()
	ResetEntityMatchingMethod()
	ResetGroupByAlertDetails()
	ResetGroupByCustomDetails()
	ResetGroupByEntities()
	ResetLookbackDuration()
	ResetReopenClosedIncidents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference
type jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) EntityMatchingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityMatchingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) EntityMatchingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityMatchingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByAlertDetails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByAlertDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByAlertDetailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByAlertDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByCustomDetails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByCustomDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByCustomDetailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByCustomDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByEntities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByEntities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GroupByEntitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByEntitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) InternalValue() *SentinelAlertRuleScheduledIncidentConfigurationGrouping {
	var returns *SentinelAlertRuleScheduledIncidentConfigurationGrouping
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) LookbackDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) LookbackDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ReopenClosedIncidents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenClosedIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ReopenClosedIncidentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenClosedIncidentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference_Override(s SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetEntityMatchingMethod(val *string) {
	_jsii_.Set(
		j,
		"entityMatchingMethod",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetGroupByAlertDetails(val *[]*string) {
	_jsii_.Set(
		j,
		"groupByAlertDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetGroupByCustomDetails(val *[]*string) {
	_jsii_.Set(
		j,
		"groupByCustomDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetGroupByEntities(val *[]*string) {
	_jsii_.Set(
		j,
		"groupByEntities",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetInternalValue(val *SentinelAlertRuleScheduledIncidentConfigurationGrouping) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetLookbackDuration(val *string) {
	_jsii_.Set(
		j,
		"lookbackDuration",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetReopenClosedIncidents(val interface{}) {
	_jsii_.Set(
		j,
		"reopenClosedIncidents",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetEntityMatchingMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetEntityMatchingMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetGroupByAlertDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupByAlertDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetGroupByCustomDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupByCustomDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetGroupByEntities() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupByEntities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetLookbackDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetLookbackDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ResetReopenClosedIncidents() {
	_jsii_.InvokeVoid(
		s,
		"resetReopenClosedIncidents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledIncidentConfigurationOutputReference interface {
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
	CreateIncident() interface{}
	SetCreateIncident(val interface{})
	CreateIncidentInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Grouping() SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference
	GroupingInput() *SentinelAlertRuleScheduledIncidentConfigurationGrouping
	InternalValue() *SentinelAlertRuleScheduledIncidentConfiguration
	SetInternalValue(val *SentinelAlertRuleScheduledIncidentConfiguration)
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
	PutGrouping(value *SentinelAlertRuleScheduledIncidentConfigurationGrouping)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleScheduledIncidentConfigurationOutputReference
type jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) CreateIncident() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIncident",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) CreateIncidentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIncidentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) Grouping() SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference {
	var returns SentinelAlertRuleScheduledIncidentConfigurationGroupingOutputReference
	_jsii_.Get(
		j,
		"grouping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GroupingInput() *SentinelAlertRuleScheduledIncidentConfigurationGrouping {
	var returns *SentinelAlertRuleScheduledIncidentConfigurationGrouping
	_jsii_.Get(
		j,
		"groupingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) InternalValue() *SentinelAlertRuleScheduledIncidentConfiguration {
	var returns *SentinelAlertRuleScheduledIncidentConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledIncidentConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SentinelAlertRuleScheduledIncidentConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledIncidentConfigurationOutputReference_Override(s SentinelAlertRuleScheduledIncidentConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledIncidentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) SetCreateIncident(val interface{}) {
	_jsii_.Set(
		j,
		"createIncident",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) SetInternalValue(val *SentinelAlertRuleScheduledIncidentConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) PutGrouping(value *SentinelAlertRuleScheduledIncidentConfigurationGrouping) {
	_jsii_.InvokeVoid(
		s,
		"putGrouping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledIncidentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SentinelAlertRuleScheduledTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#create SentinelAlertRuleScheduled#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#delete SentinelAlertRuleScheduled#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#read SentinelAlertRuleScheduled#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#update SentinelAlertRuleScheduled#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type SentinelAlertRuleScheduledTimeoutsOutputReference interface {
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

// The jsii proxy struct for SentinelAlertRuleScheduledTimeoutsOutputReference
type jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleScheduledTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SentinelAlertRuleScheduledTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleScheduledTimeoutsOutputReference_Override(s SentinelAlertRuleScheduledTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduledTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduledTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

