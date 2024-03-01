// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorscheduledqueryrulesalertv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/monitorscheduledqueryrulesalertv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/monitor_scheduled_query_rules_alert_v2 azurerm_monitor_scheduled_query_rules_alert_v2}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/monitor_scheduled_query_rules_alert_v2 azurerm_monitor_scheduled_query_rules_alert_v2} Resource.
func NewMonitorScheduledQueryRulesAlertV2(scope constructs.Construct, id *string, config *MonitorScheduledQueryRulesAlertV2Config) MonitorScheduledQueryRulesAlertV2 {
	_init_.Initialize()

	if err := validateNewMonitorScheduledQueryRulesAlertV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorScheduledQueryRulesAlertV2{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/monitor_scheduled_query_rules_alert_v2 azurerm_monitor_scheduled_query_rules_alert_v2} Resource.
func NewMonitorScheduledQueryRulesAlertV2_Override(m MonitorScheduledQueryRulesAlertV2, scope constructs.Construct, id *string, config *MonitorScheduledQueryRulesAlertV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetAutoMitigationEnabled(val interface{}) {
	if err := j.validateSetAutoMitigationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMitigationEnabled",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetEvaluationFrequency(val *string) {
	if err := j.validateSetEvaluationFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationFrequency",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetMuteActionsAfterAlertDuration(val *string) {
	if err := j.validateSetMuteActionsAfterAlertDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"muteActionsAfterAlertDuration",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetQueryTimeRangeOverride(val *string) {
	if err := j.validateSetQueryTimeRangeOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryTimeRangeOverride",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetSeverity(val *float64) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetSkipQueryValidation(val interface{}) {
	if err := j.validateSetSkipQueryValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipQueryValidation",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetTargetResourceTypes(val *[]*string) {
	if err := j.validateSetTargetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceTypes",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetWindowDuration(val *string) {
	if err := j.validateSetWindowDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowDuration",
		val,
	)
}

func (j *jsiiProxy_MonitorScheduledQueryRulesAlertV2)SetWorkspaceAlertsStorageEnabled(val interface{}) {
	if err := j.validateSetWorkspaceAlertsStorageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceAlertsStorageEnabled",
		val,
	)
}

// Generates CDKTF code for importing a MonitorScheduledQueryRulesAlertV2 resource upon running "cdktf plan <stack-name>".
func MonitorScheduledQueryRulesAlertV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMonitorScheduledQueryRulesAlertV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
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
func MonitorScheduledQueryRulesAlertV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorScheduledQueryRulesAlertV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorScheduledQueryRulesAlertV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorScheduledQueryRulesAlertV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorScheduledQueryRulesAlertV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorScheduledQueryRulesAlertV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorScheduledQueryRulesAlertV2.MonitorScheduledQueryRulesAlertV2",
		"isTerraformResource",
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) PutAction(value *MonitorScheduledQueryRulesAlertV2Action) {
	if err := m.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) PutCriteria(value interface{}) {
	if err := m.validatePutCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCriteria",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) PutTimeouts(value *MonitorScheduledQueryRulesAlertV2Timeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorScheduledQueryRulesAlertV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
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

