// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulenrt

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/sentinelalertrulenrt/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/sentinel_alert_rule_nrt azurerm_sentinel_alert_rule_nrt}.
type SentinelAlertRuleNrt interface {
	cdktf.TerraformResource
	AlertDetailsOverride() SentinelAlertRuleNrtAlertDetailsOverrideList
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
	EntityMapping() SentinelAlertRuleNrtEntityMappingList
	EntityMappingInput() interface{}
	EventGrouping() SentinelAlertRuleNrtEventGroupingOutputReference
	EventGroupingInput() *SentinelAlertRuleNrtEventGrouping
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
	Incident() SentinelAlertRuleNrtIncidentOutputReference
	IncidentInput() *SentinelAlertRuleNrtIncident
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
	QueryInput() *string
	// Experimental.
	RawOverrides() interface{}
	SentinelEntityMapping() SentinelAlertRuleNrtSentinelEntityMappingList
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
	Timeouts() SentinelAlertRuleNrtTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutEventGrouping(value *SentinelAlertRuleNrtEventGrouping)
	PutIncident(value *SentinelAlertRuleNrtIncident)
	PutSentinelEntityMapping(value interface{})
	PutTimeouts(value *SentinelAlertRuleNrtTimeouts)
	ResetAlertDetailsOverride()
	ResetAlertRuleTemplateGuid()
	ResetAlertRuleTemplateVersion()
	ResetCustomDetails()
	ResetDescription()
	ResetEnabled()
	ResetEntityMapping()
	ResetId()
	ResetIncident()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSentinelEntityMapping()
	ResetSuppressionDuration()
	ResetSuppressionEnabled()
	ResetTactics()
	ResetTechniques()
	ResetTimeouts()
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

// The jsii proxy struct for SentinelAlertRuleNrt
type jsiiProxy_SentinelAlertRuleNrt struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SentinelAlertRuleNrt) AlertDetailsOverride() SentinelAlertRuleNrtAlertDetailsOverrideList {
	var returns SentinelAlertRuleNrtAlertDetailsOverrideList
	_jsii_.Get(
		j,
		"alertDetailsOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) AlertDetailsOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertDetailsOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) AlertRuleTemplateGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertRuleTemplateGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) AlertRuleTemplateGuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertRuleTemplateGuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) AlertRuleTemplateVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertRuleTemplateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) AlertRuleTemplateVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertRuleTemplateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) CustomDetails() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) CustomDetailsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) EntityMapping() SentinelAlertRuleNrtEntityMappingList {
	var returns SentinelAlertRuleNrtEntityMappingList
	_jsii_.Get(
		j,
		"entityMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) EntityMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) EventGrouping() SentinelAlertRuleNrtEventGroupingOutputReference {
	var returns SentinelAlertRuleNrtEventGroupingOutputReference
	_jsii_.Get(
		j,
		"eventGrouping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) EventGroupingInput() *SentinelAlertRuleNrtEventGrouping {
	var returns *SentinelAlertRuleNrtEventGrouping
	_jsii_.Get(
		j,
		"eventGroupingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Incident() SentinelAlertRuleNrtIncidentOutputReference {
	var returns SentinelAlertRuleNrtIncidentOutputReference
	_jsii_.Get(
		j,
		"incident",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) IncidentInput() *SentinelAlertRuleNrtIncident {
	var returns *SentinelAlertRuleNrtIncident
	_jsii_.Get(
		j,
		"incidentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) SentinelEntityMapping() SentinelAlertRuleNrtSentinelEntityMappingList {
	var returns SentinelAlertRuleNrtSentinelEntityMappingList
	_jsii_.Get(
		j,
		"sentinelEntityMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) SentinelEntityMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentinelEntityMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) SuppressionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suppressionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) SuppressionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suppressionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) SuppressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) SuppressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Tactics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tactics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) TacticsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tacticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Techniques() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"techniques",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) TechniquesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"techniquesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) Timeouts() SentinelAlertRuleNrtTimeoutsOutputReference {
	var returns SentinelAlertRuleNrtTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleNrt) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/sentinel_alert_rule_nrt azurerm_sentinel_alert_rule_nrt} Resource.
func NewSentinelAlertRuleNrt(scope constructs.Construct, id *string, config *SentinelAlertRuleNrtConfig) SentinelAlertRuleNrt {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleNrtParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleNrt{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrt",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/sentinel_alert_rule_nrt azurerm_sentinel_alert_rule_nrt} Resource.
func NewSentinelAlertRuleNrt_Override(s SentinelAlertRuleNrt, scope constructs.Construct, id *string, config *SentinelAlertRuleNrtConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrt",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetAlertRuleTemplateGuid(val *string) {
	if err := j.validateSetAlertRuleTemplateGuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertRuleTemplateGuid",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetAlertRuleTemplateVersion(val *string) {
	if err := j.validateSetAlertRuleTemplateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertRuleTemplateVersion",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetCustomDetails(val *map[string]*string) {
	if err := j.validateSetCustomDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDetails",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetLogAnalyticsWorkspaceId(val *string) {
	if err := j.validateSetLogAnalyticsWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetSuppressionDuration(val *string) {
	if err := j.validateSetSuppressionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressionDuration",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetSuppressionEnabled(val interface{}) {
	if err := j.validateSetSuppressionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressionEnabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetTactics(val *[]*string) {
	if err := j.validateSetTacticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tactics",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleNrt)SetTechniques(val *[]*string) {
	if err := j.validateSetTechniquesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"techniques",
		val,
	)
}

// Generates CDKTF code for importing a SentinelAlertRuleNrt resource upon running "cdktf plan <stack-name>".
func SentinelAlertRuleNrt_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSentinelAlertRuleNrt_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrt",
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
func SentinelAlertRuleNrt_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelAlertRuleNrt_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrt",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelAlertRuleNrt_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelAlertRuleNrt_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrt",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelAlertRuleNrt_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelAlertRuleNrt_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrt",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SentinelAlertRuleNrt_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sentinelAlertRuleNrt.SentinelAlertRuleNrt",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrt) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleNrt) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) PutAlertDetailsOverride(value interface{}) {
	if err := s.validatePutAlertDetailsOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAlertDetailsOverride",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) PutEntityMapping(value interface{}) {
	if err := s.validatePutEntityMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEntityMapping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) PutEventGrouping(value *SentinelAlertRuleNrtEventGrouping) {
	if err := s.validatePutEventGroupingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEventGrouping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) PutIncident(value *SentinelAlertRuleNrtIncident) {
	if err := s.validatePutIncidentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIncident",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) PutSentinelEntityMapping(value interface{}) {
	if err := s.validatePutSentinelEntityMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSentinelEntityMapping",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) PutTimeouts(value *SentinelAlertRuleNrtTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetAlertDetailsOverride() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertDetailsOverride",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetAlertRuleTemplateGuid() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertRuleTemplateGuid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetAlertRuleTemplateVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertRuleTemplateVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetCustomDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetEntityMapping() {
	_jsii_.InvokeVoid(
		s,
		"resetEntityMapping",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetIncident() {
	_jsii_.InvokeVoid(
		s,
		"resetIncident",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetSentinelEntityMapping() {
	_jsii_.InvokeVoid(
		s,
		"resetSentinelEntityMapping",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetSuppressionDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetSuppressionDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetSuppressionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSuppressionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetTactics() {
	_jsii_.InvokeVoid(
		s,
		"resetTactics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetTechniques() {
	_jsii_.InvokeVoid(
		s,
		"resetTechniques",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleNrt) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrt) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleNrt) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

