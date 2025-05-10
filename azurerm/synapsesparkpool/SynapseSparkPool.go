// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsesparkpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/synapsesparkpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_spark_pool azurerm_synapse_spark_pool}.
type SynapseSparkPool interface {
	cdktf.TerraformResource
	AutoPause() SynapseSparkPoolAutoPauseOutputReference
	AutoPauseInput() *SynapseSparkPoolAutoPause
	AutoScale() SynapseSparkPoolAutoScaleOutputReference
	AutoScaleInput() *SynapseSparkPoolAutoScale
	CacheSize() *float64
	SetCacheSize(val *float64)
	CacheSizeInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputeIsolationEnabled() interface{}
	SetComputeIsolationEnabled(val interface{})
	ComputeIsolationEnabledInput() interface{}
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
	DynamicExecutorAllocationEnabled() interface{}
	SetDynamicExecutorAllocationEnabled(val interface{})
	DynamicExecutorAllocationEnabledInput() interface{}
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
	LibraryRequirement() SynapseSparkPoolLibraryRequirementOutputReference
	LibraryRequirementInput() *SynapseSparkPoolLibraryRequirement
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxExecutors() *float64
	SetMaxExecutors(val *float64)
	MaxExecutorsInput() *float64
	MinExecutors() *float64
	SetMinExecutors(val *float64)
	MinExecutorsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodeSize() *string
	SetNodeSize(val *string)
	NodeSizeFamily() *string
	SetNodeSizeFamily(val *string)
	NodeSizeFamilyInput() *string
	NodeSizeInput() *string
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
	SessionLevelPackagesEnabled() interface{}
	SetSessionLevelPackagesEnabled(val interface{})
	SessionLevelPackagesEnabledInput() interface{}
	SparkConfig() SynapseSparkPoolSparkConfigOutputReference
	SparkConfigInput() *SynapseSparkPoolSparkConfig
	SparkEventsFolder() *string
	SetSparkEventsFolder(val *string)
	SparkEventsFolderInput() *string
	SparkLogFolder() *string
	SetSparkLogFolder(val *string)
	SparkLogFolderInput() *string
	SparkVersion() *string
	SetSparkVersion(val *string)
	SparkVersionInput() *string
	SynapseWorkspaceId() *string
	SetSynapseWorkspaceId(val *string)
	SynapseWorkspaceIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SynapseSparkPoolTimeoutsOutputReference
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
	PutAutoPause(value *SynapseSparkPoolAutoPause)
	PutAutoScale(value *SynapseSparkPoolAutoScale)
	PutLibraryRequirement(value *SynapseSparkPoolLibraryRequirement)
	PutSparkConfig(value *SynapseSparkPoolSparkConfig)
	PutTimeouts(value *SynapseSparkPoolTimeouts)
	ResetAutoPause()
	ResetAutoScale()
	ResetCacheSize()
	ResetComputeIsolationEnabled()
	ResetDynamicExecutorAllocationEnabled()
	ResetId()
	ResetLibraryRequirement()
	ResetMaxExecutors()
	ResetMinExecutors()
	ResetNodeCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSessionLevelPackagesEnabled()
	ResetSparkConfig()
	ResetSparkEventsFolder()
	ResetSparkLogFolder()
	ResetTags()
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

// The jsii proxy struct for SynapseSparkPool
type jsiiProxy_SynapseSparkPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SynapseSparkPool) AutoPause() SynapseSparkPoolAutoPauseOutputReference {
	var returns SynapseSparkPoolAutoPauseOutputReference
	_jsii_.Get(
		j,
		"autoPause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) AutoPauseInput() *SynapseSparkPoolAutoPause {
	var returns *SynapseSparkPoolAutoPause
	_jsii_.Get(
		j,
		"autoPauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) AutoScale() SynapseSparkPoolAutoScaleOutputReference {
	var returns SynapseSparkPoolAutoScaleOutputReference
	_jsii_.Get(
		j,
		"autoScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) AutoScaleInput() *SynapseSparkPoolAutoScale {
	var returns *SynapseSparkPoolAutoScale
	_jsii_.Get(
		j,
		"autoScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) CacheSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) CacheSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) ComputeIsolationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeIsolationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) ComputeIsolationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeIsolationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) DynamicExecutorAllocationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicExecutorAllocationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) DynamicExecutorAllocationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicExecutorAllocationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) LibraryRequirement() SynapseSparkPoolLibraryRequirementOutputReference {
	var returns SynapseSparkPoolLibraryRequirementOutputReference
	_jsii_.Get(
		j,
		"libraryRequirement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) LibraryRequirementInput() *SynapseSparkPoolLibraryRequirement {
	var returns *SynapseSparkPoolLibraryRequirement
	_jsii_.Get(
		j,
		"libraryRequirementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) MaxExecutors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExecutors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) MaxExecutorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExecutorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) MinExecutors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) MinExecutorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) NodeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) NodeSizeFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSizeFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) NodeSizeFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSizeFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) NodeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SessionLevelPackagesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionLevelPackagesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SessionLevelPackagesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionLevelPackagesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SparkConfig() SynapseSparkPoolSparkConfigOutputReference {
	var returns SynapseSparkPoolSparkConfigOutputReference
	_jsii_.Get(
		j,
		"sparkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SparkConfigInput() *SynapseSparkPoolSparkConfig {
	var returns *SynapseSparkPoolSparkConfig
	_jsii_.Get(
		j,
		"sparkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SparkEventsFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkEventsFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SparkEventsFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkEventsFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SparkLogFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkLogFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SparkLogFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkLogFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SynapseWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"synapseWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) SynapseWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"synapseWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) Timeouts() SynapseSparkPoolTimeoutsOutputReference {
	var returns SynapseSparkPoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSparkPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_spark_pool azurerm_synapse_spark_pool} Resource.
func NewSynapseSparkPool(scope constructs.Construct, id *string, config *SynapseSparkPoolConfig) SynapseSparkPool {
	_init_.Initialize()

	if err := validateNewSynapseSparkPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SynapseSparkPool{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseSparkPool.SynapseSparkPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_spark_pool azurerm_synapse_spark_pool} Resource.
func NewSynapseSparkPool_Override(s SynapseSparkPool, scope constructs.Construct, id *string, config *SynapseSparkPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseSparkPool.SynapseSparkPool",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetCacheSize(val *float64) {
	if err := j.validateSetCacheSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheSize",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetComputeIsolationEnabled(val interface{}) {
	if err := j.validateSetComputeIsolationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeIsolationEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetDynamicExecutorAllocationEnabled(val interface{}) {
	if err := j.validateSetDynamicExecutorAllocationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicExecutorAllocationEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetMaxExecutors(val *float64) {
	if err := j.validateSetMaxExecutorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxExecutors",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetMinExecutors(val *float64) {
	if err := j.validateSetMinExecutorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minExecutors",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetNodeSize(val *string) {
	if err := j.validateSetNodeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSize",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetNodeSizeFamily(val *string) {
	if err := j.validateSetNodeSizeFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSizeFamily",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetSessionLevelPackagesEnabled(val interface{}) {
	if err := j.validateSetSessionLevelPackagesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionLevelPackagesEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetSparkEventsFolder(val *string) {
	if err := j.validateSetSparkEventsFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEventsFolder",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetSparkLogFolder(val *string) {
	if err := j.validateSetSparkLogFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkLogFolder",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetSynapseWorkspaceId(val *string) {
	if err := j.validateSetSynapseWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"synapseWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_SynapseSparkPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a SynapseSparkPool resource upon running "cdktf plan <stack-name>".
func SynapseSparkPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSynapseSparkPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseSparkPool.SynapseSparkPool",
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
func SynapseSparkPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseSparkPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseSparkPool.SynapseSparkPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SynapseSparkPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseSparkPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseSparkPool.SynapseSparkPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SynapseSparkPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseSparkPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseSparkPool.SynapseSparkPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SynapseSparkPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.synapseSparkPool.SynapseSparkPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SynapseSparkPool) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SynapseSparkPool) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SynapseSparkPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SynapseSparkPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SynapseSparkPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SynapseSparkPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SynapseSparkPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SynapseSparkPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SynapseSparkPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SynapseSparkPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SynapseSparkPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SynapseSparkPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSparkPool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SynapseSparkPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SynapseSparkPool) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SynapseSparkPool) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SynapseSparkPool) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SynapseSparkPool) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SynapseSparkPool) PutAutoPause(value *SynapseSparkPoolAutoPause) {
	if err := s.validatePutAutoPauseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAutoPause",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseSparkPool) PutAutoScale(value *SynapseSparkPoolAutoScale) {
	if err := s.validatePutAutoScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAutoScale",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseSparkPool) PutLibraryRequirement(value *SynapseSparkPoolLibraryRequirement) {
	if err := s.validatePutLibraryRequirementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLibraryRequirement",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseSparkPool) PutSparkConfig(value *SynapseSparkPoolSparkConfig) {
	if err := s.validatePutSparkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSparkConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseSparkPool) PutTimeouts(value *SynapseSparkPoolTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetAutoPause() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoPause",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetAutoScale() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoScale",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetCacheSize() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetComputeIsolationEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetComputeIsolationEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetDynamicExecutorAllocationEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDynamicExecutorAllocationEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetLibraryRequirement() {
	_jsii_.InvokeVoid(
		s,
		"resetLibraryRequirement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetMaxExecutors() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxExecutors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetMinExecutors() {
	_jsii_.InvokeVoid(
		s,
		"resetMinExecutors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetNodeCount() {
	_jsii_.InvokeVoid(
		s,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetSessionLevelPackagesEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionLevelPackagesEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetSparkConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetSparkConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetSparkEventsFolder() {
	_jsii_.InvokeVoid(
		s,
		"resetSparkEventsFolder",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetSparkLogFolder() {
	_jsii_.InvokeVoid(
		s,
		"resetSparkLogFolder",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSparkPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSparkPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSparkPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSparkPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSparkPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSparkPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

