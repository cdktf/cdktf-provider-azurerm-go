// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudappdynamicsapplicationperformancemonitoring

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/springcloudappdynamicsapplicationperformancemonitoring/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/spring_cloud_app_dynamics_application_performance_monitoring azurerm_spring_cloud_app_dynamics_application_performance_monitoring}.
type SpringCloudAppDynamicsApplicationPerformanceMonitoring interface {
	cdktf.TerraformResource
	AgentAccountAccessKey() *string
	SetAgentAccountAccessKey(val *string)
	AgentAccountAccessKeyInput() *string
	AgentAccountName() *string
	SetAgentAccountName(val *string)
	AgentAccountNameInput() *string
	AgentApplicationName() *string
	SetAgentApplicationName(val *string)
	AgentApplicationNameInput() *string
	AgentNodeName() *string
	SetAgentNodeName(val *string)
	AgentNodeNameInput() *string
	AgentTierName() *string
	SetAgentTierName(val *string)
	AgentTierNameInput() *string
	AgentUniqueHostId() *string
	SetAgentUniqueHostId(val *string)
	AgentUniqueHostIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControllerHostName() *string
	SetControllerHostName(val *string)
	ControllerHostNameInput() *string
	ControllerPort() *float64
	SetControllerPort(val *float64)
	ControllerPortInput() *float64
	ControllerSslEnabled() interface{}
	SetControllerSslEnabled(val interface{})
	ControllerSslEnabledInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	Timeouts() SpringCloudAppDynamicsApplicationPerformanceMonitoringTimeoutsOutputReference
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
	PutTimeouts(value *SpringCloudAppDynamicsApplicationPerformanceMonitoringTimeouts)
	ResetAgentApplicationName()
	ResetAgentNodeName()
	ResetAgentTierName()
	ResetAgentUniqueHostId()
	ResetControllerPort()
	ResetControllerSslEnabled()
	ResetGloballyEnabled()
	ResetId()
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

// The jsii proxy struct for SpringCloudAppDynamicsApplicationPerformanceMonitoring
type jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentApplicationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentApplicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentNodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentNodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentNodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentNodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentTierName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentTierName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentTierNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentTierNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentUniqueHostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentUniqueHostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AgentUniqueHostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentUniqueHostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ControllerHostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerHostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ControllerHostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerHostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ControllerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ControllerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ControllerSslEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controllerSslEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ControllerSslEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controllerSslEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GloballyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globallyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GloballyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globallyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) SpringCloudServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"springCloudServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) SpringCloudServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"springCloudServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) Timeouts() SpringCloudAppDynamicsApplicationPerformanceMonitoringTimeoutsOutputReference {
	var returns SpringCloudAppDynamicsApplicationPerformanceMonitoringTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/spring_cloud_app_dynamics_application_performance_monitoring azurerm_spring_cloud_app_dynamics_application_performance_monitoring} Resource.
func NewSpringCloudAppDynamicsApplicationPerformanceMonitoring(scope constructs.Construct, id *string, config *SpringCloudAppDynamicsApplicationPerformanceMonitoringConfig) SpringCloudAppDynamicsApplicationPerformanceMonitoring {
	_init_.Initialize()

	if err := validateNewSpringCloudAppDynamicsApplicationPerformanceMonitoringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudAppDynamicsApplicationPerformanceMonitoring.SpringCloudAppDynamicsApplicationPerformanceMonitoring",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/spring_cloud_app_dynamics_application_performance_monitoring azurerm_spring_cloud_app_dynamics_application_performance_monitoring} Resource.
func NewSpringCloudAppDynamicsApplicationPerformanceMonitoring_Override(s SpringCloudAppDynamicsApplicationPerformanceMonitoring, scope constructs.Construct, id *string, config *SpringCloudAppDynamicsApplicationPerformanceMonitoringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudAppDynamicsApplicationPerformanceMonitoring.SpringCloudAppDynamicsApplicationPerformanceMonitoring",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetAgentAccountAccessKey(val *string) {
	if err := j.validateSetAgentAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetAgentAccountName(val *string) {
	if err := j.validateSetAgentAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentAccountName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetAgentApplicationName(val *string) {
	if err := j.validateSetAgentApplicationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentApplicationName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetAgentNodeName(val *string) {
	if err := j.validateSetAgentNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentNodeName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetAgentTierName(val *string) {
	if err := j.validateSetAgentTierNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentTierName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetAgentUniqueHostId(val *string) {
	if err := j.validateSetAgentUniqueHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentUniqueHostId",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetControllerHostName(val *string) {
	if err := j.validateSetControllerHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controllerHostName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetControllerPort(val *float64) {
	if err := j.validateSetControllerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controllerPort",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetControllerSslEnabled(val interface{}) {
	if err := j.validateSetControllerSslEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controllerSslEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetGloballyEnabled(val interface{}) {
	if err := j.validateSetGloballyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globallyEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring)SetSpringCloudServiceId(val *string) {
	if err := j.validateSetSpringCloudServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"springCloudServiceId",
		val,
	)
}

// Generates CDKTF code for importing a SpringCloudAppDynamicsApplicationPerformanceMonitoring resource upon running "cdktf plan <stack-name>".
func SpringCloudAppDynamicsApplicationPerformanceMonitoring_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSpringCloudAppDynamicsApplicationPerformanceMonitoring_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppDynamicsApplicationPerformanceMonitoring.SpringCloudAppDynamicsApplicationPerformanceMonitoring",
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
func SpringCloudAppDynamicsApplicationPerformanceMonitoring_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudAppDynamicsApplicationPerformanceMonitoring_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppDynamicsApplicationPerformanceMonitoring.SpringCloudAppDynamicsApplicationPerformanceMonitoring",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudAppDynamicsApplicationPerformanceMonitoring_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudAppDynamicsApplicationPerformanceMonitoring_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppDynamicsApplicationPerformanceMonitoring.SpringCloudAppDynamicsApplicationPerformanceMonitoring",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudAppDynamicsApplicationPerformanceMonitoring_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudAppDynamicsApplicationPerformanceMonitoring_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppDynamicsApplicationPerformanceMonitoring.SpringCloudAppDynamicsApplicationPerformanceMonitoring",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpringCloudAppDynamicsApplicationPerformanceMonitoring_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.springCloudAppDynamicsApplicationPerformanceMonitoring.SpringCloudAppDynamicsApplicationPerformanceMonitoring",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) PutTimeouts(value *SpringCloudAppDynamicsApplicationPerformanceMonitoringTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetAgentApplicationName() {
	_jsii_.InvokeVoid(
		s,
		"resetAgentApplicationName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetAgentNodeName() {
	_jsii_.InvokeVoid(
		s,
		"resetAgentNodeName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetAgentTierName() {
	_jsii_.InvokeVoid(
		s,
		"resetAgentTierName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetAgentUniqueHostId() {
	_jsii_.InvokeVoid(
		s,
		"resetAgentUniqueHostId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetControllerPort() {
	_jsii_.InvokeVoid(
		s,
		"resetControllerPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetControllerSslEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetControllerSslEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetGloballyEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetGloballyEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppDynamicsApplicationPerformanceMonitoring) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

