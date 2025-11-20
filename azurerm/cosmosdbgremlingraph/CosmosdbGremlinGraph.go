// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbgremlingraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/cosmosdbgremlingraph/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cosmosdb_gremlin_graph azurerm_cosmosdb_gremlin_graph}.
type CosmosdbGremlinGraph interface {
	cdktf.TerraformResource
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	AnalyticalStorageTtl() *float64
	SetAnalyticalStorageTtl(val *float64)
	AnalyticalStorageTtlInput() *float64
	AutoscaleSettings() CosmosdbGremlinGraphAutoscaleSettingsOutputReference
	AutoscaleSettingsInput() *CosmosdbGremlinGraphAutoscaleSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConflictResolutionPolicy() CosmosdbGremlinGraphConflictResolutionPolicyOutputReference
	ConflictResolutionPolicyInput() *CosmosdbGremlinGraphConflictResolutionPolicy
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
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
	IndexPolicy() CosmosdbGremlinGraphIndexPolicyOutputReference
	IndexPolicyInput() *CosmosdbGremlinGraphIndexPolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PartitionKeyPath() *string
	SetPartitionKeyPath(val *string)
	PartitionKeyPathInput() *string
	PartitionKeyVersion() *float64
	SetPartitionKeyVersion(val *float64)
	PartitionKeyVersionInput() *float64
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
	Throughput() *float64
	SetThroughput(val *float64)
	ThroughputInput() *float64
	Timeouts() CosmosdbGremlinGraphTimeoutsOutputReference
	TimeoutsInput() interface{}
	UniqueKey() CosmosdbGremlinGraphUniqueKeyList
	UniqueKeyInput() interface{}
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
	PutAutoscaleSettings(value *CosmosdbGremlinGraphAutoscaleSettings)
	PutConflictResolutionPolicy(value *CosmosdbGremlinGraphConflictResolutionPolicy)
	PutIndexPolicy(value *CosmosdbGremlinGraphIndexPolicy)
	PutTimeouts(value *CosmosdbGremlinGraphTimeouts)
	PutUniqueKey(value interface{})
	ResetAnalyticalStorageTtl()
	ResetAutoscaleSettings()
	ResetConflictResolutionPolicy()
	ResetDefaultTtl()
	ResetId()
	ResetIndexPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionKeyVersion()
	ResetThroughput()
	ResetTimeouts()
	ResetUniqueKey()
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

// The jsii proxy struct for CosmosdbGremlinGraph
type jsiiProxy_CosmosdbGremlinGraph struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CosmosdbGremlinGraph) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) AnalyticalStorageTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analyticalStorageTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) AnalyticalStorageTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analyticalStorageTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) AutoscaleSettings() CosmosdbGremlinGraphAutoscaleSettingsOutputReference {
	var returns CosmosdbGremlinGraphAutoscaleSettingsOutputReference
	_jsii_.Get(
		j,
		"autoscaleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) AutoscaleSettingsInput() *CosmosdbGremlinGraphAutoscaleSettings {
	var returns *CosmosdbGremlinGraphAutoscaleSettings
	_jsii_.Get(
		j,
		"autoscaleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) ConflictResolutionPolicy() CosmosdbGremlinGraphConflictResolutionPolicyOutputReference {
	var returns CosmosdbGremlinGraphConflictResolutionPolicyOutputReference
	_jsii_.Get(
		j,
		"conflictResolutionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) ConflictResolutionPolicyInput() *CosmosdbGremlinGraphConflictResolutionPolicy {
	var returns *CosmosdbGremlinGraphConflictResolutionPolicy
	_jsii_.Get(
		j,
		"conflictResolutionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) IndexPolicy() CosmosdbGremlinGraphIndexPolicyOutputReference {
	var returns CosmosdbGremlinGraphIndexPolicyOutputReference
	_jsii_.Get(
		j,
		"indexPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) IndexPolicyInput() *CosmosdbGremlinGraphIndexPolicy {
	var returns *CosmosdbGremlinGraphIndexPolicy
	_jsii_.Get(
		j,
		"indexPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) PartitionKeyPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) PartitionKeyPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) PartitionKeyVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionKeyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) PartitionKeyVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionKeyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Throughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) ThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) Timeouts() CosmosdbGremlinGraphTimeoutsOutputReference {
	var returns CosmosdbGremlinGraphTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) UniqueKey() CosmosdbGremlinGraphUniqueKeyList {
	var returns CosmosdbGremlinGraphUniqueKeyList
	_jsii_.Get(
		j,
		"uniqueKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraph) UniqueKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueKeyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cosmosdb_gremlin_graph azurerm_cosmosdb_gremlin_graph} Resource.
func NewCosmosdbGremlinGraph(scope constructs.Construct, id *string, config *CosmosdbGremlinGraphConfig) CosmosdbGremlinGraph {
	_init_.Initialize()

	if err := validateNewCosmosdbGremlinGraphParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbGremlinGraph{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraph",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cosmosdb_gremlin_graph azurerm_cosmosdb_gremlin_graph} Resource.
func NewCosmosdbGremlinGraph_Override(c CosmosdbGremlinGraph, scope constructs.Construct, id *string, config *CosmosdbGremlinGraphConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraph",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetAccountName(val *string) {
	if err := j.validateSetAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetAnalyticalStorageTtl(val *float64) {
	if err := j.validateSetAnalyticalStorageTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyticalStorageTtl",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetPartitionKeyPath(val *string) {
	if err := j.validateSetPartitionKeyPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionKeyPath",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetPartitionKeyVersion(val *float64) {
	if err := j.validateSetPartitionKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionKeyVersion",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraph)SetThroughput(val *float64) {
	if err := j.validateSetThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughput",
		val,
	)
}

// Generates CDKTF code for importing a CosmosdbGremlinGraph resource upon running "cdktf plan <stack-name>".
func CosmosdbGremlinGraph_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCosmosdbGremlinGraph_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraph",
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
func CosmosdbGremlinGraph_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbGremlinGraph_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraph",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbGremlinGraph_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbGremlinGraph_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraph",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbGremlinGraph_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbGremlinGraph_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraph",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CosmosdbGremlinGraph_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraph",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) PutAutoscaleSettings(value *CosmosdbGremlinGraphAutoscaleSettings) {
	if err := c.validatePutAutoscaleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoscaleSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) PutConflictResolutionPolicy(value *CosmosdbGremlinGraphConflictResolutionPolicy) {
	if err := c.validatePutConflictResolutionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConflictResolutionPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) PutIndexPolicy(value *CosmosdbGremlinGraphIndexPolicy) {
	if err := c.validatePutIndexPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIndexPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) PutTimeouts(value *CosmosdbGremlinGraphTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) PutUniqueKey(value interface{}) {
	if err := c.validatePutUniqueKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUniqueKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetAnalyticalStorageTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalyticalStorageTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetAutoscaleSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoscaleSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetConflictResolutionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetConflictResolutionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetIndexPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetIndexPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetPartitionKeyVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetPartitionKeyVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetThroughput() {
	_jsii_.InvokeVoid(
		c,
		"resetThroughput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ResetUniqueKey() {
	_jsii_.InvokeVoid(
		c,
		"resetUniqueKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraph) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraph) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

