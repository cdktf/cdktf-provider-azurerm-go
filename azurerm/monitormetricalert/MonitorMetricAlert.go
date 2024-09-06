// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitormetricalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/monitormetricalert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/monitor_metric_alert azurerm_monitor_metric_alert}.
type MonitorMetricAlert interface {
	cdktf.TerraformResource
	Action() MonitorMetricAlertActionList
	ActionInput() interface{}
	ApplicationInsightsWebTestLocationAvailabilityCriteria() MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference
	ApplicationInsightsWebTestLocationAvailabilityCriteriaInput() *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria
	AutoMitigate() interface{}
	SetAutoMitigate(val interface{})
	AutoMitigateInput() interface{}
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
	Criteria() MonitorMetricAlertCriteriaList
	CriteriaInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DynamicCriteria() MonitorMetricAlertDynamicCriteriaOutputReference
	DynamicCriteriaInput() *MonitorMetricAlertDynamicCriteria
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
	SetFrequency(val *string)
	FrequencyInput() *string
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
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Severity() *float64
	SetSeverity(val *float64)
	SeverityInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TargetResourceLocation() *string
	SetTargetResourceLocation(val *string)
	TargetResourceLocationInput() *string
	TargetResourceType() *string
	SetTargetResourceType(val *string)
	TargetResourceTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MonitorMetricAlertTimeoutsOutputReference
	TimeoutsInput() interface{}
	WindowSize() *string
	SetWindowSize(val *string)
	WindowSizeInput() *string
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
	PutAction(value interface{})
	PutApplicationInsightsWebTestLocationAvailabilityCriteria(value *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria)
	PutCriteria(value interface{})
	PutDynamicCriteria(value *MonitorMetricAlertDynamicCriteria)
	PutTimeouts(value *MonitorMetricAlertTimeouts)
	ResetAction()
	ResetApplicationInsightsWebTestLocationAvailabilityCriteria()
	ResetAutoMitigate()
	ResetCriteria()
	ResetDescription()
	ResetDynamicCriteria()
	ResetEnabled()
	ResetFrequency()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSeverity()
	ResetTags()
	ResetTargetResourceLocation()
	ResetTargetResourceType()
	ResetTimeouts()
	ResetWindowSize()
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

// The jsii proxy struct for MonitorMetricAlert
type jsiiProxy_MonitorMetricAlert struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitorMetricAlert) Action() MonitorMetricAlertActionList {
	var returns MonitorMetricAlertActionList
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) ActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) ApplicationInsightsWebTestLocationAvailabilityCriteria() MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference {
	var returns MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference
	_jsii_.Get(
		j,
		"applicationInsightsWebTestLocationAvailabilityCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) ApplicationInsightsWebTestLocationAvailabilityCriteriaInput() *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria {
	var returns *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria
	_jsii_.Get(
		j,
		"applicationInsightsWebTestLocationAvailabilityCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) AutoMitigate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMitigate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) AutoMitigateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMitigateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Criteria() MonitorMetricAlertCriteriaList {
	var returns MonitorMetricAlertCriteriaList
	_jsii_.Get(
		j,
		"criteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) CriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) DynamicCriteria() MonitorMetricAlertDynamicCriteriaOutputReference {
	var returns MonitorMetricAlertDynamicCriteriaOutputReference
	_jsii_.Get(
		j,
		"dynamicCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) DynamicCriteriaInput() *MonitorMetricAlertDynamicCriteria {
	var returns *MonitorMetricAlertDynamicCriteria
	_jsii_.Get(
		j,
		"dynamicCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Severity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) SeverityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TargetResourceLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TargetResourceLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TargetResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TargetResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) Timeouts() MonitorMetricAlertTimeoutsOutputReference {
	var returns MonitorMetricAlertTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) WindowSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlert) WindowSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowSizeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/monitor_metric_alert azurerm_monitor_metric_alert} Resource.
func NewMonitorMetricAlert(scope constructs.Construct, id *string, config *MonitorMetricAlertConfig) MonitorMetricAlert {
	_init_.Initialize()

	if err := validateNewMonitorMetricAlertParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorMetricAlert{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/monitor_metric_alert azurerm_monitor_metric_alert} Resource.
func NewMonitorMetricAlert_Override(m MonitorMetricAlert, scope constructs.Construct, id *string, config *MonitorMetricAlertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlert",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetAutoMitigate(val interface{}) {
	if err := j.validateSetAutoMitigateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMitigate",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetSeverity(val *float64) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetTargetResourceLocation(val *string) {
	if err := j.validateSetTargetResourceLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceLocation",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetTargetResourceType(val *string) {
	if err := j.validateSetTargetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceType",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlert)SetWindowSize(val *string) {
	if err := j.validateSetWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowSize",
		val,
	)
}

// Generates CDKTF code for importing a MonitorMetricAlert resource upon running "cdktf plan <stack-name>".
func MonitorMetricAlert_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMonitorMetricAlert_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlert",
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
func MonitorMetricAlert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorMetricAlert_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorMetricAlert_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorMetricAlert_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlert",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorMetricAlert_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorMetricAlert_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlert",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorMetricAlert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.monitorMetricAlert.MonitorMetricAlert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorMetricAlert) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorMetricAlert) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorMetricAlert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorMetricAlert) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorMetricAlert) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorMetricAlert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorMetricAlert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorMetricAlert) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorMetricAlert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorMetricAlert) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlert) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorMetricAlert) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) PutAction(value interface{}) {
	if err := m.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) PutApplicationInsightsWebTestLocationAvailabilityCriteria(value *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria) {
	if err := m.validatePutApplicationInsightsWebTestLocationAvailabilityCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putApplicationInsightsWebTestLocationAvailabilityCriteria",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) PutCriteria(value interface{}) {
	if err := m.validatePutCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCriteria",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) PutDynamicCriteria(value *MonitorMetricAlertDynamicCriteria) {
	if err := m.validatePutDynamicCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDynamicCriteria",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) PutTimeouts(value *MonitorMetricAlertTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetAction() {
	_jsii_.InvokeVoid(
		m,
		"resetAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetApplicationInsightsWebTestLocationAvailabilityCriteria() {
	_jsii_.InvokeVoid(
		m,
		"resetApplicationInsightsWebTestLocationAvailabilityCriteria",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetAutoMitigate() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoMitigate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetCriteria() {
	_jsii_.InvokeVoid(
		m,
		"resetCriteria",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetDynamicCriteria() {
	_jsii_.InvokeVoid(
		m,
		"resetDynamicCriteria",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetFrequency() {
	_jsii_.InvokeVoid(
		m,
		"resetFrequency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetSeverity() {
	_jsii_.InvokeVoid(
		m,
		"resetSeverity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetTargetResourceLocation() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceLocation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetTargetResourceType() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetResourceType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) ResetWindowSize() {
	_jsii_.InvokeVoid(
		m,
		"resetWindowSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMetricAlert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlert) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlert) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

