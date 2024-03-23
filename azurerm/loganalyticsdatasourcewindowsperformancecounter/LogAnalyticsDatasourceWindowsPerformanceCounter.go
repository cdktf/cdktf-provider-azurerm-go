// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsdatasourcewindowsperformancecounter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/loganalyticsdatasourcewindowsperformancecounter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/log_analytics_datasource_windows_performance_counter azurerm_log_analytics_datasource_windows_performance_counter}.
type LogAnalyticsDatasourceWindowsPerformanceCounter interface {
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
	CounterName() *string
	SetCounterName(val *string)
	CounterNameInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceName() *string
	SetInstanceName(val *string)
	InstanceNameInput() *string
	IntervalSeconds() *float64
	SetIntervalSeconds(val *float64)
	IntervalSecondsInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ObjectName() *string
	SetObjectName(val *string)
	ObjectNameInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LogAnalyticsDatasourceWindowsPerformanceCounterTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkspaceName() *string
	SetWorkspaceName(val *string)
	WorkspaceNameInput() *string
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
	PutTimeouts(value *LogAnalyticsDatasourceWindowsPerformanceCounterTimeouts)
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

// The jsii proxy struct for LogAnalyticsDatasourceWindowsPerformanceCounter
type jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) CounterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"counterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) CounterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"counterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) InstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) InstanceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) IntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) IntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) Timeouts() LogAnalyticsDatasourceWindowsPerformanceCounterTimeoutsOutputReference {
	var returns LogAnalyticsDatasourceWindowsPerformanceCounterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) WorkspaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) WorkspaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/log_analytics_datasource_windows_performance_counter azurerm_log_analytics_datasource_windows_performance_counter} Resource.
func NewLogAnalyticsDatasourceWindowsPerformanceCounter(scope constructs.Construct, id *string, config *LogAnalyticsDatasourceWindowsPerformanceCounterConfig) LogAnalyticsDatasourceWindowsPerformanceCounter {
	_init_.Initialize()

	if err := validateNewLogAnalyticsDatasourceWindowsPerformanceCounterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logAnalyticsDatasourceWindowsPerformanceCounter.LogAnalyticsDatasourceWindowsPerformanceCounter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/log_analytics_datasource_windows_performance_counter azurerm_log_analytics_datasource_windows_performance_counter} Resource.
func NewLogAnalyticsDatasourceWindowsPerformanceCounter_Override(l LogAnalyticsDatasourceWindowsPerformanceCounter, scope constructs.Construct, id *string, config *LogAnalyticsDatasourceWindowsPerformanceCounterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logAnalyticsDatasourceWindowsPerformanceCounter.LogAnalyticsDatasourceWindowsPerformanceCounter",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetCounterName(val *string) {
	if err := j.validateSetCounterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"counterName",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetInstanceName(val *string) {
	if err := j.validateSetInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceName",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetIntervalSeconds(val *float64) {
	if err := j.validateSetIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalSeconds",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetObjectName(val *string) {
	if err := j.validateSetObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectName",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter)SetWorkspaceName(val *string) {
	if err := j.validateSetWorkspaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceName",
		val,
	)
}

// Generates CDKTF code for importing a LogAnalyticsDatasourceWindowsPerformanceCounter resource upon running "cdktf plan <stack-name>".
func LogAnalyticsDatasourceWindowsPerformanceCounter_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLogAnalyticsDatasourceWindowsPerformanceCounter_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logAnalyticsDatasourceWindowsPerformanceCounter.LogAnalyticsDatasourceWindowsPerformanceCounter",
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
func LogAnalyticsDatasourceWindowsPerformanceCounter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogAnalyticsDatasourceWindowsPerformanceCounter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logAnalyticsDatasourceWindowsPerformanceCounter.LogAnalyticsDatasourceWindowsPerformanceCounter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogAnalyticsDatasourceWindowsPerformanceCounter_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogAnalyticsDatasourceWindowsPerformanceCounter_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logAnalyticsDatasourceWindowsPerformanceCounter.LogAnalyticsDatasourceWindowsPerformanceCounter",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogAnalyticsDatasourceWindowsPerformanceCounter_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogAnalyticsDatasourceWindowsPerformanceCounter_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logAnalyticsDatasourceWindowsPerformanceCounter.LogAnalyticsDatasourceWindowsPerformanceCounter",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogAnalyticsDatasourceWindowsPerformanceCounter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.logAnalyticsDatasourceWindowsPerformanceCounter.LogAnalyticsDatasourceWindowsPerformanceCounter",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) PutTimeouts(value *LogAnalyticsDatasourceWindowsPerformanceCounterTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsDatasourceWindowsPerformanceCounter) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

