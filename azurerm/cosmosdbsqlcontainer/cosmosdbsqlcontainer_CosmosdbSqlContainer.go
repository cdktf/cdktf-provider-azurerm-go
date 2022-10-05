package cosmosdbsqlcontainer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/cosmosdbsqlcontainer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_container azurerm_cosmosdb_sql_container}.
type CosmosdbSqlContainer interface {
	cdktf.TerraformResource
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	AnalyticalStorageTtl() *float64
	SetAnalyticalStorageTtl(val *float64)
	AnalyticalStorageTtlInput() *float64
	AutoscaleSettings() CosmosdbSqlContainerAutoscaleSettingsOutputReference
	AutoscaleSettingsInput() *CosmosdbSqlContainerAutoscaleSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConflictResolutionPolicy() CosmosdbSqlContainerConflictResolutionPolicyOutputReference
	ConflictResolutionPolicyInput() *CosmosdbSqlContainerConflictResolutionPolicy
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	IndexingPolicy() CosmosdbSqlContainerIndexingPolicyOutputReference
	IndexingPolicyInput() *CosmosdbSqlContainerIndexingPolicy
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
	Timeouts() CosmosdbSqlContainerTimeoutsOutputReference
	TimeoutsInput() interface{}
	UniqueKey() CosmosdbSqlContainerUniqueKeyList
	UniqueKeyInput() interface{}
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
	PutAutoscaleSettings(value *CosmosdbSqlContainerAutoscaleSettings)
	PutConflictResolutionPolicy(value *CosmosdbSqlContainerConflictResolutionPolicy)
	PutIndexingPolicy(value *CosmosdbSqlContainerIndexingPolicy)
	PutTimeouts(value *CosmosdbSqlContainerTimeouts)
	PutUniqueKey(value interface{})
	ResetAnalyticalStorageTtl()
	ResetAutoscaleSettings()
	ResetConflictResolutionPolicy()
	ResetDefaultTtl()
	ResetId()
	ResetIndexingPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionKeyVersion()
	ResetThroughput()
	ResetTimeouts()
	ResetUniqueKey()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CosmosdbSqlContainer
type jsiiProxy_CosmosdbSqlContainer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CosmosdbSqlContainer) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) AnalyticalStorageTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analyticalStorageTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) AnalyticalStorageTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analyticalStorageTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) AutoscaleSettings() CosmosdbSqlContainerAutoscaleSettingsOutputReference {
	var returns CosmosdbSqlContainerAutoscaleSettingsOutputReference
	_jsii_.Get(
		j,
		"autoscaleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) AutoscaleSettingsInput() *CosmosdbSqlContainerAutoscaleSettings {
	var returns *CosmosdbSqlContainerAutoscaleSettings
	_jsii_.Get(
		j,
		"autoscaleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) ConflictResolutionPolicy() CosmosdbSqlContainerConflictResolutionPolicyOutputReference {
	var returns CosmosdbSqlContainerConflictResolutionPolicyOutputReference
	_jsii_.Get(
		j,
		"conflictResolutionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) ConflictResolutionPolicyInput() *CosmosdbSqlContainerConflictResolutionPolicy {
	var returns *CosmosdbSqlContainerConflictResolutionPolicy
	_jsii_.Get(
		j,
		"conflictResolutionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) IndexingPolicy() CosmosdbSqlContainerIndexingPolicyOutputReference {
	var returns CosmosdbSqlContainerIndexingPolicyOutputReference
	_jsii_.Get(
		j,
		"indexingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) IndexingPolicyInput() *CosmosdbSqlContainerIndexingPolicy {
	var returns *CosmosdbSqlContainerIndexingPolicy
	_jsii_.Get(
		j,
		"indexingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) PartitionKeyPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) PartitionKeyPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) PartitionKeyVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionKeyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) PartitionKeyVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionKeyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Throughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) ThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) Timeouts() CosmosdbSqlContainerTimeoutsOutputReference {
	var returns CosmosdbSqlContainerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) UniqueKey() CosmosdbSqlContainerUniqueKeyList {
	var returns CosmosdbSqlContainerUniqueKeyList
	_jsii_.Get(
		j,
		"uniqueKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainer) UniqueKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueKeyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_container azurerm_cosmosdb_sql_container} Resource.
func NewCosmosdbSqlContainer(scope constructs.Construct, id *string, config *CosmosdbSqlContainerConfig) CosmosdbSqlContainer {
	_init_.Initialize()

	if err := validateNewCosmosdbSqlContainerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbSqlContainer{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbSqlContainer.CosmosdbSqlContainer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_container azurerm_cosmosdb_sql_container} Resource.
func NewCosmosdbSqlContainer_Override(c CosmosdbSqlContainer, scope constructs.Construct, id *string, config *CosmosdbSqlContainerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbSqlContainer.CosmosdbSqlContainer",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetAccountName(val *string) {
	if err := j.validateSetAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetAnalyticalStorageTtl(val *float64) {
	if err := j.validateSetAnalyticalStorageTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyticalStorageTtl",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetPartitionKeyPath(val *string) {
	if err := j.validateSetPartitionKeyPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionKeyPath",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetPartitionKeyVersion(val *float64) {
	if err := j.validateSetPartitionKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionKeyVersion",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainer)SetThroughput(val *float64) {
	if err := j.validateSetThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughput",
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
func CosmosdbSqlContainer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbSqlContainer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbSqlContainer.CosmosdbSqlContainer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CosmosdbSqlContainer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cosmosdbSqlContainer.CosmosdbSqlContainer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainer) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CosmosdbSqlContainer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbSqlContainer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CosmosdbSqlContainer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CosmosdbSqlContainer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CosmosdbSqlContainer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CosmosdbSqlContainer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CosmosdbSqlContainer) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CosmosdbSqlContainer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CosmosdbSqlContainer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbSqlContainer) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) PutAutoscaleSettings(value *CosmosdbSqlContainerAutoscaleSettings) {
	if err := c.validatePutAutoscaleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoscaleSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) PutConflictResolutionPolicy(value *CosmosdbSqlContainerConflictResolutionPolicy) {
	if err := c.validatePutConflictResolutionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConflictResolutionPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) PutIndexingPolicy(value *CosmosdbSqlContainerIndexingPolicy) {
	if err := c.validatePutIndexingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIndexingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) PutTimeouts(value *CosmosdbSqlContainerTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) PutUniqueKey(value interface{}) {
	if err := c.validatePutUniqueKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUniqueKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetAnalyticalStorageTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalyticalStorageTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetAutoscaleSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoscaleSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetConflictResolutionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetConflictResolutionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetIndexingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetIndexingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetPartitionKeyVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetPartitionKeyVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetThroughput() {
	_jsii_.InvokeVoid(
		c,
		"resetThroughput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) ResetUniqueKey() {
	_jsii_.InvokeVoid(
		c,
		"resetUniqueKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

