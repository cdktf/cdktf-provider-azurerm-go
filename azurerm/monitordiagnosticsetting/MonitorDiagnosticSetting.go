// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordiagnosticsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/monitordiagnosticsetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/monitor_diagnostic_setting azurerm_monitor_diagnostic_setting}.
type MonitorDiagnosticSetting interface {
	cdktf.TerraformResource
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
	EnabledLog() MonitorDiagnosticSettingEnabledLogList
	EnabledLogInput() interface{}
	EnabledMetric() MonitorDiagnosticSettingEnabledMetricList
	EnabledMetricInput() interface{}
	EventhubAuthorizationRuleId() *string
	SetEventhubAuthorizationRuleId(val *string)
	EventhubAuthorizationRuleIdInput() *string
	EventhubName() *string
	SetEventhubName(val *string)
	EventhubNameInput() *string
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
	LogAnalyticsDestinationType() *string
	SetLogAnalyticsDestinationType(val *string)
	LogAnalyticsDestinationTypeInput() *string
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	Metric() MonitorDiagnosticSettingMetricList
	MetricInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PartnerSolutionId() *string
	SetPartnerSolutionId(val *string)
	PartnerSolutionIdInput() *string
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
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	TargetResourceId() *string
	SetTargetResourceId(val *string)
	TargetResourceIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MonitorDiagnosticSettingTimeoutsOutputReference
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
	PutEnabledLog(value interface{})
	PutEnabledMetric(value interface{})
	PutMetric(value interface{})
	PutTimeouts(value *MonitorDiagnosticSettingTimeouts)
	ResetEnabledLog()
	ResetEnabledMetric()
	ResetEventhubAuthorizationRuleId()
	ResetEventhubName()
	ResetId()
	ResetLogAnalyticsDestinationType()
	ResetLogAnalyticsWorkspaceId()
	ResetMetric()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartnerSolutionId()
	ResetStorageAccountId()
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

// The jsii proxy struct for MonitorDiagnosticSetting
type jsiiProxy_MonitorDiagnosticSetting struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitorDiagnosticSetting) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) EnabledLog() MonitorDiagnosticSettingEnabledLogList {
	var returns MonitorDiagnosticSettingEnabledLogList
	_jsii_.Get(
		j,
		"enabledLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) EnabledLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) EnabledMetric() MonitorDiagnosticSettingEnabledMetricList {
	var returns MonitorDiagnosticSettingEnabledMetricList
	_jsii_.Get(
		j,
		"enabledMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) EnabledMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) EventhubAuthorizationRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubAuthorizationRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) EventhubAuthorizationRuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubAuthorizationRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) EventhubName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) EventhubNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) LogAnalyticsDestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsDestinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) LogAnalyticsDestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsDestinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Metric() MonitorDiagnosticSettingMetricList {
	var returns MonitorDiagnosticSettingMetricList
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) MetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) PartnerSolutionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerSolutionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) PartnerSolutionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerSolutionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) TargetResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) TargetResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) Timeouts() MonitorDiagnosticSettingTimeoutsOutputReference {
	var returns MonitorDiagnosticSettingTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDiagnosticSetting) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/monitor_diagnostic_setting azurerm_monitor_diagnostic_setting} Resource.
func NewMonitorDiagnosticSetting(scope constructs.Construct, id *string, config *MonitorDiagnosticSettingConfig) MonitorDiagnosticSetting {
	_init_.Initialize()

	if err := validateNewMonitorDiagnosticSettingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorDiagnosticSetting{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDiagnosticSetting.MonitorDiagnosticSetting",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/monitor_diagnostic_setting azurerm_monitor_diagnostic_setting} Resource.
func NewMonitorDiagnosticSetting_Override(m MonitorDiagnosticSetting, scope constructs.Construct, id *string, config *MonitorDiagnosticSettingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDiagnosticSetting.MonitorDiagnosticSetting",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetEventhubAuthorizationRuleId(val *string) {
	if err := j.validateSetEventhubAuthorizationRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubAuthorizationRuleId",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetEventhubName(val *string) {
	if err := j.validateSetEventhubNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubName",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetLogAnalyticsDestinationType(val *string) {
	if err := j.validateSetLogAnalyticsDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsDestinationType",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetLogAnalyticsWorkspaceId(val *string) {
	if err := j.validateSetLogAnalyticsWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetPartnerSolutionId(val *string) {
	if err := j.validateSetPartnerSolutionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerSolutionId",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetStorageAccountId(val *string) {
	if err := j.validateSetStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_MonitorDiagnosticSetting)SetTargetResourceId(val *string) {
	if err := j.validateSetTargetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceId",
		val,
	)
}

// Generates CDKTF code for importing a MonitorDiagnosticSetting resource upon running "cdktf plan <stack-name>".
func MonitorDiagnosticSetting_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMonitorDiagnosticSetting_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorDiagnosticSetting.MonitorDiagnosticSetting",
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
func MonitorDiagnosticSetting_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorDiagnosticSetting_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorDiagnosticSetting.MonitorDiagnosticSetting",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorDiagnosticSetting_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorDiagnosticSetting_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorDiagnosticSetting.MonitorDiagnosticSetting",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorDiagnosticSetting_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorDiagnosticSetting_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorDiagnosticSetting.MonitorDiagnosticSetting",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorDiagnosticSetting_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.monitorDiagnosticSetting.MonitorDiagnosticSetting",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorDiagnosticSetting) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorDiagnosticSetting) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) PutEnabledLog(value interface{}) {
	if err := m.validatePutEnabledLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEnabledLog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) PutEnabledMetric(value interface{}) {
	if err := m.validatePutEnabledMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEnabledMetric",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) PutMetric(value interface{}) {
	if err := m.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMetric",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) PutTimeouts(value *MonitorDiagnosticSettingTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetEnabledLog() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabledLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetEnabledMetric() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabledMetric",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetEventhubAuthorizationRuleId() {
	_jsii_.InvokeVoid(
		m,
		"resetEventhubAuthorizationRuleId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetEventhubName() {
	_jsii_.InvokeVoid(
		m,
		"resetEventhubName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetLogAnalyticsDestinationType() {
	_jsii_.InvokeVoid(
		m,
		"resetLogAnalyticsDestinationType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetLogAnalyticsWorkspaceId() {
	_jsii_.InvokeVoid(
		m,
		"resetLogAnalyticsWorkspaceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetMetric() {
	_jsii_.InvokeVoid(
		m,
		"resetMetric",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetPartnerSolutionId() {
	_jsii_.InvokeVoid(
		m,
		"resetPartnerSolutionId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetStorageAccountId() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDiagnosticSetting) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDiagnosticSetting) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDiagnosticSetting) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

