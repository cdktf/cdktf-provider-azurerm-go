package dataazurermsentinelalertruleanomaly

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/dataazurermsentinelalertruleanomaly/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/data-sources/sentinel_alert_rule_anomaly azurerm_sentinel_alert_rule_anomaly}.
type DataAzurermSentinelAlertRuleAnomaly interface {
	cdktf.TerraformDataSource
	AnomalySettingsVersion() *float64
	AnomalyVersion() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Enabled() cdktf.IResolvable
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	Mode() *string
	MultiSelectObservation() DataAzurermSentinelAlertRuleAnomalyMultiSelectObservationList
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrioritizedExcludeObservation() DataAzurermSentinelAlertRuleAnomalyPrioritizedExcludeObservationList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RequiredDataConnector() DataAzurermSentinelAlertRuleAnomalyRequiredDataConnectorList
	SettingsDefinitionId() *string
	SingleSelectObservation() DataAzurermSentinelAlertRuleAnomalySingleSelectObservationList
	Tactics() *[]*string
	Techniques() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThresholdObservation() DataAzurermSentinelAlertRuleAnomalyThresholdObservationList
	Timeouts() DataAzurermSentinelAlertRuleAnomalyTimeoutsOutputReference
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
	PutTimeouts(value *DataAzurermSentinelAlertRuleAnomalyTimeouts)
	ResetDisplayName()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataAzurermSentinelAlertRuleAnomaly
type jsiiProxy_DataAzurermSentinelAlertRuleAnomaly struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) AnomalySettingsVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anomalySettingsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) AnomalyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) MultiSelectObservation() DataAzurermSentinelAlertRuleAnomalyMultiSelectObservationList {
	var returns DataAzurermSentinelAlertRuleAnomalyMultiSelectObservationList
	_jsii_.Get(
		j,
		"multiSelectObservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) PrioritizedExcludeObservation() DataAzurermSentinelAlertRuleAnomalyPrioritizedExcludeObservationList {
	var returns DataAzurermSentinelAlertRuleAnomalyPrioritizedExcludeObservationList
	_jsii_.Get(
		j,
		"prioritizedExcludeObservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) RequiredDataConnector() DataAzurermSentinelAlertRuleAnomalyRequiredDataConnectorList {
	var returns DataAzurermSentinelAlertRuleAnomalyRequiredDataConnectorList
	_jsii_.Get(
		j,
		"requiredDataConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) SettingsDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) SingleSelectObservation() DataAzurermSentinelAlertRuleAnomalySingleSelectObservationList {
	var returns DataAzurermSentinelAlertRuleAnomalySingleSelectObservationList
	_jsii_.Get(
		j,
		"singleSelectObservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Tactics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tactics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Techniques() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"techniques",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ThresholdObservation() DataAzurermSentinelAlertRuleAnomalyThresholdObservationList {
	var returns DataAzurermSentinelAlertRuleAnomalyThresholdObservationList
	_jsii_.Get(
		j,
		"thresholdObservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) Timeouts() DataAzurermSentinelAlertRuleAnomalyTimeoutsOutputReference {
	var returns DataAzurermSentinelAlertRuleAnomalyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/data-sources/sentinel_alert_rule_anomaly azurerm_sentinel_alert_rule_anomaly} Data Source.
func NewDataAzurermSentinelAlertRuleAnomaly(scope constructs.Construct, id *string, config *DataAzurermSentinelAlertRuleAnomalyConfig) DataAzurermSentinelAlertRuleAnomaly {
	_init_.Initialize()

	if err := validateNewDataAzurermSentinelAlertRuleAnomalyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermSentinelAlertRuleAnomaly{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermSentinelAlertRuleAnomaly.DataAzurermSentinelAlertRuleAnomaly",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/data-sources/sentinel_alert_rule_anomaly azurerm_sentinel_alert_rule_anomaly} Data Source.
func NewDataAzurermSentinelAlertRuleAnomaly_Override(d DataAzurermSentinelAlertRuleAnomaly, scope constructs.Construct, id *string, config *DataAzurermSentinelAlertRuleAnomalyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermSentinelAlertRuleAnomaly.DataAzurermSentinelAlertRuleAnomaly",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetLogAnalyticsWorkspaceId(val *string) {
	if err := j.validateSetLogAnalyticsWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataAzurermSentinelAlertRuleAnomaly_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermSentinelAlertRuleAnomaly_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermSentinelAlertRuleAnomaly.DataAzurermSentinelAlertRuleAnomaly",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermSentinelAlertRuleAnomaly_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermSentinelAlertRuleAnomaly_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermSentinelAlertRuleAnomaly.DataAzurermSentinelAlertRuleAnomaly",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermSentinelAlertRuleAnomaly_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermSentinelAlertRuleAnomaly_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermSentinelAlertRuleAnomaly.DataAzurermSentinelAlertRuleAnomaly",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermSentinelAlertRuleAnomaly_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermSentinelAlertRuleAnomaly.DataAzurermSentinelAlertRuleAnomaly",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) PutTimeouts(value *DataAzurermSentinelAlertRuleAnomalyTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSentinelAlertRuleAnomaly) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

