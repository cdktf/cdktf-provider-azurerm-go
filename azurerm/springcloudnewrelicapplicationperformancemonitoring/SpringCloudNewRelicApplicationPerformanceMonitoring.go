// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudnewrelicapplicationperformancemonitoring

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/springcloudnewrelicapplicationperformancemonitoring/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring azurerm_spring_cloud_new_relic_application_performance_monitoring}.
type SpringCloudNewRelicApplicationPerformanceMonitoring interface {
	cdktf.TerraformResource
	AgentEnabled() interface{}
	SetAgentEnabled(val interface{})
	AgentEnabledInput() interface{}
	AppName() *string
	SetAppName(val *string)
	AppNameInput() *string
	AppServerPort() *float64
	SetAppServerPort(val *float64)
	AppServerPortInput() *float64
	AuditModeEnabled() interface{}
	SetAuditModeEnabled(val interface{})
	AuditModeEnabledInput() interface{}
	AutoAppNamingEnabled() interface{}
	SetAutoAppNamingEnabled(val interface{})
	AutoAppNamingEnabledInput() interface{}
	AutoTransactionNamingEnabled() interface{}
	SetAutoTransactionNamingEnabled(val interface{})
	AutoTransactionNamingEnabledInput() interface{}
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
	CustomTracingEnabled() interface{}
	SetCustomTracingEnabled(val interface{})
	CustomTracingEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GloballyEnabled() interface{}
	SetGloballyEnabled(val interface{})
	GloballyEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LicenseKey() *string
	SetLicenseKey(val *string)
	LicenseKeyInput() *string
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
	SpringCloudServiceId() *string
	SetSpringCloudServiceId(val *string)
	SpringCloudServiceIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SpringCloudNewRelicApplicationPerformanceMonitoringTimeoutsOutputReference
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
	PutTimeouts(value *SpringCloudNewRelicApplicationPerformanceMonitoringTimeouts)
	ResetAgentEnabled()
	ResetAppServerPort()
	ResetAuditModeEnabled()
	ResetAutoAppNamingEnabled()
	ResetAutoTransactionNamingEnabled()
	ResetCustomTracingEnabled()
	ResetGloballyEnabled()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for SpringCloudNewRelicApplicationPerformanceMonitoring
type jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AgentEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AgentEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AppNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AppServerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appServerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AppServerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appServerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AuditModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AuditModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditModeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AutoAppNamingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAppNamingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AutoAppNamingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAppNamingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AutoTransactionNamingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoTransactionNamingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AutoTransactionNamingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoTransactionNamingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) CustomTracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customTracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) CustomTracingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customTracingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GloballyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globallyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GloballyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globallyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) LicenseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) LicenseKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) SpringCloudServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"springCloudServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) SpringCloudServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"springCloudServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) Timeouts() SpringCloudNewRelicApplicationPerformanceMonitoringTimeoutsOutputReference {
	var returns SpringCloudNewRelicApplicationPerformanceMonitoringTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring azurerm_spring_cloud_new_relic_application_performance_monitoring} Resource.
func NewSpringCloudNewRelicApplicationPerformanceMonitoring(scope constructs.Construct, id *string, config *SpringCloudNewRelicApplicationPerformanceMonitoringConfig) SpringCloudNewRelicApplicationPerformanceMonitoring {
	_init_.Initialize()

	if err := validateNewSpringCloudNewRelicApplicationPerformanceMonitoringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudNewRelicApplicationPerformanceMonitoring.SpringCloudNewRelicApplicationPerformanceMonitoring",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring azurerm_spring_cloud_new_relic_application_performance_monitoring} Resource.
func NewSpringCloudNewRelicApplicationPerformanceMonitoring_Override(s SpringCloudNewRelicApplicationPerformanceMonitoring, scope constructs.Construct, id *string, config *SpringCloudNewRelicApplicationPerformanceMonitoringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudNewRelicApplicationPerformanceMonitoring.SpringCloudNewRelicApplicationPerformanceMonitoring",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetAgentEnabled(val interface{}) {
	if err := j.validateSetAgentEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetAppName(val *string) {
	if err := j.validateSetAppNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetAppServerPort(val *float64) {
	if err := j.validateSetAppServerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appServerPort",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetAuditModeEnabled(val interface{}) {
	if err := j.validateSetAuditModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditModeEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetAutoAppNamingEnabled(val interface{}) {
	if err := j.validateSetAutoAppNamingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAppNamingEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetAutoTransactionNamingEnabled(val interface{}) {
	if err := j.validateSetAutoTransactionNamingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoTransactionNamingEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetCustomTracingEnabled(val interface{}) {
	if err := j.validateSetCustomTracingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTracingEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetGloballyEnabled(val interface{}) {
	if err := j.validateSetGloballyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globallyEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetLicenseKey(val *string) {
	if err := j.validateSetLicenseKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring)SetSpringCloudServiceId(val *string) {
	if err := j.validateSetSpringCloudServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"springCloudServiceId",
		val,
	)
}

// Generates CDKTF code for importing a SpringCloudNewRelicApplicationPerformanceMonitoring resource upon running "cdktf plan <stack-name>".
func SpringCloudNewRelicApplicationPerformanceMonitoring_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSpringCloudNewRelicApplicationPerformanceMonitoring_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudNewRelicApplicationPerformanceMonitoring.SpringCloudNewRelicApplicationPerformanceMonitoring",
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
func SpringCloudNewRelicApplicationPerformanceMonitoring_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudNewRelicApplicationPerformanceMonitoring_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudNewRelicApplicationPerformanceMonitoring.SpringCloudNewRelicApplicationPerformanceMonitoring",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudNewRelicApplicationPerformanceMonitoring_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudNewRelicApplicationPerformanceMonitoring_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudNewRelicApplicationPerformanceMonitoring.SpringCloudNewRelicApplicationPerformanceMonitoring",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudNewRelicApplicationPerformanceMonitoring_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudNewRelicApplicationPerformanceMonitoring_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudNewRelicApplicationPerformanceMonitoring.SpringCloudNewRelicApplicationPerformanceMonitoring",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpringCloudNewRelicApplicationPerformanceMonitoring_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.springCloudNewRelicApplicationPerformanceMonitoring.SpringCloudNewRelicApplicationPerformanceMonitoring",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) PutTimeouts(value *SpringCloudNewRelicApplicationPerformanceMonitoringTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetAgentEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetAgentEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetAppServerPort() {
	_jsii_.InvokeVoid(
		s,
		"resetAppServerPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetAuditModeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetAuditModeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetAutoAppNamingEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoAppNamingEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetAutoTransactionNamingEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoTransactionNamingEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetCustomTracingEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomTracingEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetGloballyEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetGloballyEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetLabels() {
	_jsii_.InvokeVoid(
		s,
		"resetLabels",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudNewRelicApplicationPerformanceMonitoring) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

