// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertruleanomalyduplicate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/sentinelalertruleanomalyduplicate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/sentinel_alert_rule_anomaly_duplicate azurerm_sentinel_alert_rule_anomaly_duplicate}.
type SentinelAlertRuleAnomalyDuplicate interface {
	cdktf.TerraformResource
	AnomalySettingsVersion() *float64
	AnomalyVersion() *string
	BuiltInRuleId() *string
	SetBuiltInRuleId(val *string)
	BuiltInRuleIdInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	Frequency() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsDefaultSettings() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	MultiSelectObservation() SentinelAlertRuleAnomalyDuplicateMultiSelectObservationList
	MultiSelectObservationInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	PrioritizedExcludeObservation() SentinelAlertRuleAnomalyDuplicatePrioritizedExcludeObservationList
	PrioritizedExcludeObservationInput() interface{}
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
	RequiredDataConnector() SentinelAlertRuleAnomalyDuplicateRequiredDataConnectorList
	SettingsDefinitionId() *string
	SingleSelectObservation() SentinelAlertRuleAnomalyDuplicateSingleSelectObservationList
	SingleSelectObservationInput() interface{}
	Tactics() *[]*string
	Techniques() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThresholdObservation() SentinelAlertRuleAnomalyDuplicateThresholdObservationList
	ThresholdObservationInput() interface{}
	Timeouts() SentinelAlertRuleAnomalyDuplicateTimeoutsOutputReference
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
	PutMultiSelectObservation(value interface{})
	PutPrioritizedExcludeObservation(value interface{})
	PutSingleSelectObservation(value interface{})
	PutThresholdObservation(value interface{})
	PutTimeouts(value *SentinelAlertRuleAnomalyDuplicateTimeouts)
	ResetId()
	ResetMultiSelectObservation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrioritizedExcludeObservation()
	ResetSingleSelectObservation()
	ResetThresholdObservation()
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

// The jsii proxy struct for SentinelAlertRuleAnomalyDuplicate
type jsiiProxy_SentinelAlertRuleAnomalyDuplicate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) AnomalySettingsVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anomalySettingsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) AnomalyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) BuiltInRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) BuiltInRuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) IsDefaultSettings() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isDefaultSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) MultiSelectObservation() SentinelAlertRuleAnomalyDuplicateMultiSelectObservationList {
	var returns SentinelAlertRuleAnomalyDuplicateMultiSelectObservationList
	_jsii_.Get(
		j,
		"multiSelectObservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) MultiSelectObservationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiSelectObservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) PrioritizedExcludeObservation() SentinelAlertRuleAnomalyDuplicatePrioritizedExcludeObservationList {
	var returns SentinelAlertRuleAnomalyDuplicatePrioritizedExcludeObservationList
	_jsii_.Get(
		j,
		"prioritizedExcludeObservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) PrioritizedExcludeObservationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizedExcludeObservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) RequiredDataConnector() SentinelAlertRuleAnomalyDuplicateRequiredDataConnectorList {
	var returns SentinelAlertRuleAnomalyDuplicateRequiredDataConnectorList
	_jsii_.Get(
		j,
		"requiredDataConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) SettingsDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) SingleSelectObservation() SentinelAlertRuleAnomalyDuplicateSingleSelectObservationList {
	var returns SentinelAlertRuleAnomalyDuplicateSingleSelectObservationList
	_jsii_.Get(
		j,
		"singleSelectObservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) SingleSelectObservationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleSelectObservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Tactics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tactics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Techniques() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"techniques",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ThresholdObservation() SentinelAlertRuleAnomalyDuplicateThresholdObservationList {
	var returns SentinelAlertRuleAnomalyDuplicateThresholdObservationList
	_jsii_.Get(
		j,
		"thresholdObservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ThresholdObservationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdObservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) Timeouts() SentinelAlertRuleAnomalyDuplicateTimeoutsOutputReference {
	var returns SentinelAlertRuleAnomalyDuplicateTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/sentinel_alert_rule_anomaly_duplicate azurerm_sentinel_alert_rule_anomaly_duplicate} Resource.
func NewSentinelAlertRuleAnomalyDuplicate(scope constructs.Construct, id *string, config *SentinelAlertRuleAnomalyDuplicateConfig) SentinelAlertRuleAnomalyDuplicate {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleAnomalyDuplicateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleAnomalyDuplicate{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyDuplicate.SentinelAlertRuleAnomalyDuplicate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/sentinel_alert_rule_anomaly_duplicate azurerm_sentinel_alert_rule_anomaly_duplicate} Resource.
func NewSentinelAlertRuleAnomalyDuplicate_Override(s SentinelAlertRuleAnomalyDuplicate, scope constructs.Construct, id *string, config *SentinelAlertRuleAnomalyDuplicateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyDuplicate.SentinelAlertRuleAnomalyDuplicate",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetBuiltInRuleId(val *string) {
	if err := j.validateSetBuiltInRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInRuleId",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetLogAnalyticsWorkspaceId(val *string) {
	if err := j.validateSetLogAnalyticsWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyDuplicate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func SentinelAlertRuleAnomalyDuplicate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelAlertRuleAnomalyDuplicate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyDuplicate.SentinelAlertRuleAnomalyDuplicate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelAlertRuleAnomalyDuplicate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelAlertRuleAnomalyDuplicate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyDuplicate.SentinelAlertRuleAnomalyDuplicate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelAlertRuleAnomalyDuplicate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelAlertRuleAnomalyDuplicate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyDuplicate.SentinelAlertRuleAnomalyDuplicate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SentinelAlertRuleAnomalyDuplicate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyDuplicate.SentinelAlertRuleAnomalyDuplicate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) PutMultiSelectObservation(value interface{}) {
	if err := s.validatePutMultiSelectObservationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMultiSelectObservation",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) PutPrioritizedExcludeObservation(value interface{}) {
	if err := s.validatePutPrioritizedExcludeObservationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrioritizedExcludeObservation",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) PutSingleSelectObservation(value interface{}) {
	if err := s.validatePutSingleSelectObservationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSingleSelectObservation",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) PutThresholdObservation(value interface{}) {
	if err := s.validatePutThresholdObservationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThresholdObservation",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) PutTimeouts(value *SentinelAlertRuleAnomalyDuplicateTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ResetMultiSelectObservation() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiSelectObservation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ResetPrioritizedExcludeObservation() {
	_jsii_.InvokeVoid(
		s,
		"resetPrioritizedExcludeObservation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ResetSingleSelectObservation() {
	_jsii_.InvokeVoid(
		s,
		"resetSingleSelectObservation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ResetThresholdObservation() {
	_jsii_.InvokeVoid(
		s,
		"resetThresholdObservation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyDuplicate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

