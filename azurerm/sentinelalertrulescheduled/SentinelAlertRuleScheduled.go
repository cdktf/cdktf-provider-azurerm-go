// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/sentinelalertrulescheduled/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/sentinel_alert_rule_scheduled azurerm_sentinel_alert_rule_scheduled}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	Incident() SentinelAlertRuleScheduledIncidentOutputReference
	IncidentInput() *SentinelAlertRuleScheduledIncident
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
	SentinelEntityMapping() SentinelAlertRuleScheduledSentinelEntityMappingList
	SentinelEntityMappingInput() interface{}
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
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAlertDetailsOverride(value interface{})
	PutEntityMapping(value interface{})
	PutEventGrouping(value *SentinelAlertRuleScheduledEventGrouping)
	PutIncident(value *SentinelAlertRuleScheduledIncident)
	PutSentinelEntityMapping(value interface{})
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
	ResetIncident()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueryFrequency()
	ResetQueryPeriod()
	ResetSentinelEntityMapping()
	ResetSuppressionDuration()
	ResetSuppressionEnabled()
	ResetTactics()
	ResetTechniques()
	ResetTimeouts()
	ResetTriggerOperator()
	ResetTriggerThreshold()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
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

func (j *jsiiProxy_SentinelAlertRuleScheduled) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_SentinelAlertRuleScheduled) Incident() SentinelAlertRuleScheduledIncidentOutputReference {
	var returns SentinelAlertRuleScheduledIncidentOutputReference
	_jsii_.Get(
		j,
		"incident",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) IncidentInput() *SentinelAlertRuleScheduledIncident {
	var returns *SentinelAlertRuleScheduledIncident
	_jsii_.Get(
		j,
		"incidentInput",
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

func (j *jsiiProxy_SentinelAlertRuleScheduled) SentinelEntityMapping() SentinelAlertRuleScheduledSentinelEntityMappingList {
	var returns SentinelAlertRuleScheduledSentinelEntityMappingList
	_jsii_.Get(
		j,
		"sentinelEntityMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleScheduled) SentinelEntityMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentinelEntityMappingInput",
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/sentinel_alert_rule_scheduled azurerm_sentinel_alert_rule_scheduled} Resource.
func NewSentinelAlertRuleScheduled(scope constructs.Construct, id *string, config *SentinelAlertRuleScheduledConfig) SentinelAlertRuleScheduled {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleScheduledParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleScheduled{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/sentinel_alert_rule_scheduled azurerm_sentinel_alert_rule_scheduled} Resource.
func NewSentinelAlertRuleScheduled_Override(s SentinelAlertRuleScheduled, scope constructs.Construct, id *string, config *SentinelAlertRuleScheduledConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetAlertRuleTemplateGuid(val *string) {
	if err := j.validateSetAlertRuleTemplateGuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertRuleTemplateGuid",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetAlertRuleTemplateVersion(val *string) {
	if err := j.validateSetAlertRuleTemplateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertRuleTemplateVersion",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetCustomDetails(val *map[string]*string) {
	if err := j.validateSetCustomDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetLogAnalyticsWorkspaceId(val *string) {
	if err := j.validateSetLogAnalyticsWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetQueryFrequency(val *string) {
	if err := j.validateSetQueryFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFrequency",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetQueryPeriod(val *string) {
	if err := j.validateSetQueryPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryPeriod",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetSuppressionDuration(val *string) {
	if err := j.validateSetSuppressionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressionDuration",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetSuppressionEnabled(val interface{}) {
	if err := j.validateSetSuppressionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressionEnabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetTactics(val *[]*string) {
	if err := j.validateSetTacticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tactics",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetTechniques(val *[]*string) {
	if err := j.validateSetTechniquesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"techniques",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetTriggerOperator(val *string) {
	if err := j.validateSetTriggerOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerOperator",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleScheduled)SetTriggerThreshold(val *float64) {
	if err := j.validateSetTriggerThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerThreshold",
		val,
	)
}

// Generates CDKTF code for importing a SentinelAlertRuleScheduled resource upon running "cdktf plan <stack-name>".
func SentinelAlertRuleScheduled_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSentinelAlertRuleScheduled_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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

	if err := validateSentinelAlertRuleScheduled_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelAlertRuleScheduled_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelAlertRuleScheduled_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelAlertRuleScheduled_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelAlertRuleScheduled_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleScheduled.SentinelAlertRuleScheduled",
		"isTerraformResource",
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

func (s *jsiiProxy_SentinelAlertRuleScheduled) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutAlertDetailsOverride(value interface{}) {
	if err := s.validatePutAlertDetailsOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAlertDetailsOverride",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutEntityMapping(value interface{}) {
	if err := s.validatePutEntityMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEntityMapping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutEventGrouping(value *SentinelAlertRuleScheduledEventGrouping) {
	if err := s.validatePutEventGroupingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEventGrouping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutIncident(value *SentinelAlertRuleScheduledIncident) {
	if err := s.validatePutIncidentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIncident",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutSentinelEntityMapping(value interface{}) {
	if err := s.validatePutSentinelEntityMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSentinelEntityMapping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) PutTimeouts(value *SentinelAlertRuleScheduledTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetIncident() {
	_jsii_.InvokeVoid(
		s,
		"resetIncident",
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

func (s *jsiiProxy_SentinelAlertRuleScheduled) ResetSentinelEntityMapping() {
	_jsii_.InvokeVoid(
		s,
		"resetSentinelEntityMapping",
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

func (s *jsiiProxy_SentinelAlertRuleScheduled) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleScheduled) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
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

