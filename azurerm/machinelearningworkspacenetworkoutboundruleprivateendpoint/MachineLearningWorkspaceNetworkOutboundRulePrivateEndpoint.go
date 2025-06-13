// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningworkspacenetworkoutboundruleprivateendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/machinelearningworkspacenetworkoutboundruleprivateendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint azurerm_machine_learning_workspace_network_outbound_rule_private_endpoint}.
type MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint interface {
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
	ServiceResourceId() *string
	SetServiceResourceId(val *string)
	ServiceResourceIdInput() *string
	SparkEnabled() interface{}
	SetSparkEnabled(val interface{})
	SparkEnabledInput() interface{}
	SubResourceTarget() *string
	SetSubResourceTarget(val *string)
	SubResourceTargetInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	PutTimeouts(value *MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSparkEnabled()
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

// The jsii proxy struct for MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint
type jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ServiceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ServiceResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) SparkEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) SparkEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparkEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) SubResourceTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subResourceTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) SubResourceTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subResourceTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) Timeouts() MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointTimeoutsOutputReference {
	var returns MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint azurerm_machine_learning_workspace_network_outbound_rule_private_endpoint} Resource.
func NewMachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint(scope constructs.Construct, id *string, config *MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointConfig) MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint {
	_init_.Initialize()

	if err := validateNewMachineLearningWorkspaceNetworkOutboundRulePrivateEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.machineLearningWorkspaceNetworkOutboundRulePrivateEndpoint.MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint azurerm_machine_learning_workspace_network_outbound_rule_private_endpoint} Resource.
func NewMachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_Override(m MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint, scope constructs.Construct, id *string, config *MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.machineLearningWorkspaceNetworkOutboundRulePrivateEndpoint.MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetServiceResourceId(val *string) {
	if err := j.validateSetServiceResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceResourceId",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetSparkEnabled(val interface{}) {
	if err := j.validateSetSparkEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnabled",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetSubResourceTarget(val *string) {
	if err := j.validateSetSubResourceTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subResourceTarget",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTF code for importing a MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint resource upon running "cdktf plan <stack-name>".
func MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.machineLearningWorkspaceNetworkOutboundRulePrivateEndpoint.MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint",
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
func MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.machineLearningWorkspaceNetworkOutboundRulePrivateEndpoint.MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.machineLearningWorkspaceNetworkOutboundRulePrivateEndpoint.MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.machineLearningWorkspaceNetworkOutboundRulePrivateEndpoint.MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.machineLearningWorkspaceNetworkOutboundRulePrivateEndpoint.MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) PutTimeouts(value *MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ResetSparkEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetSparkEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

