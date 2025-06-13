// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudappcosmosdbassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/springcloudappcosmosdbassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_app_cosmosdb_association azurerm_spring_cloud_app_cosmosdb_association}.
type SpringCloudAppCosmosdbAssociation interface {
	cdktf.TerraformResource
	ApiType() *string
	SetApiType(val *string)
	ApiTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CosmosdbAccessKey() *string
	SetCosmosdbAccessKey(val *string)
	CosmosdbAccessKeyInput() *string
	CosmosdbAccountId() *string
	SetCosmosdbAccountId(val *string)
	CosmosdbAccountIdInput() *string
	CosmosdbCassandraKeyspaceName() *string
	SetCosmosdbCassandraKeyspaceName(val *string)
	CosmosdbCassandraKeyspaceNameInput() *string
	CosmosdbGremlinDatabaseName() *string
	SetCosmosdbGremlinDatabaseName(val *string)
	CosmosdbGremlinDatabaseNameInput() *string
	CosmosdbGremlinGraphName() *string
	SetCosmosdbGremlinGraphName(val *string)
	CosmosdbGremlinGraphNameInput() *string
	CosmosdbMongoDatabaseName() *string
	SetCosmosdbMongoDatabaseName(val *string)
	CosmosdbMongoDatabaseNameInput() *string
	CosmosdbSqlDatabaseName() *string
	SetCosmosdbSqlDatabaseName(val *string)
	CosmosdbSqlDatabaseNameInput() *string
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
	SpringCloudAppId() *string
	SetSpringCloudAppId(val *string)
	SpringCloudAppIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SpringCloudAppCosmosdbAssociationTimeoutsOutputReference
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
	PutTimeouts(value *SpringCloudAppCosmosdbAssociationTimeouts)
	ResetCosmosdbCassandraKeyspaceName()
	ResetCosmosdbGremlinDatabaseName()
	ResetCosmosdbGremlinGraphName()
	ResetCosmosdbMongoDatabaseName()
	ResetCosmosdbSqlDatabaseName()
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

// The jsii proxy struct for SpringCloudAppCosmosdbAssociation
type jsiiProxy_SpringCloudAppCosmosdbAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) ApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) ApiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbCassandraKeyspaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbCassandraKeyspaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbCassandraKeyspaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbCassandraKeyspaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbGremlinDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbGremlinDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbGremlinDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbGremlinDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbGremlinGraphName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbGremlinGraphName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbGremlinGraphNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbGremlinGraphNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbMongoDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbMongoDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbMongoDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbMongoDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbSqlDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbSqlDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) CosmosdbSqlDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbSqlDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SpringCloudAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"springCloudAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SpringCloudAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"springCloudAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Timeouts() SpringCloudAppCosmosdbAssociationTimeoutsOutputReference {
	var returns SpringCloudAppCosmosdbAssociationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_app_cosmosdb_association azurerm_spring_cloud_app_cosmosdb_association} Resource.
func NewSpringCloudAppCosmosdbAssociation(scope constructs.Construct, id *string, config *SpringCloudAppCosmosdbAssociationConfig) SpringCloudAppCosmosdbAssociation {
	_init_.Initialize()

	if err := validateNewSpringCloudAppCosmosdbAssociationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudAppCosmosdbAssociation{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/spring_cloud_app_cosmosdb_association azurerm_spring_cloud_app_cosmosdb_association} Resource.
func NewSpringCloudAppCosmosdbAssociation_Override(s SpringCloudAppCosmosdbAssociation, scope constructs.Construct, id *string, config *SpringCloudAppCosmosdbAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetApiType(val *string) {
	if err := j.validateSetApiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiType",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetCosmosdbAccessKey(val *string) {
	if err := j.validateSetCosmosdbAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbAccessKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetCosmosdbAccountId(val *string) {
	if err := j.validateSetCosmosdbAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbAccountId",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetCosmosdbCassandraKeyspaceName(val *string) {
	if err := j.validateSetCosmosdbCassandraKeyspaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbCassandraKeyspaceName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetCosmosdbGremlinDatabaseName(val *string) {
	if err := j.validateSetCosmosdbGremlinDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbGremlinDatabaseName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetCosmosdbGremlinGraphName(val *string) {
	if err := j.validateSetCosmosdbGremlinGraphNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbGremlinGraphName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetCosmosdbMongoDatabaseName(val *string) {
	if err := j.validateSetCosmosdbMongoDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbMongoDatabaseName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetCosmosdbSqlDatabaseName(val *string) {
	if err := j.validateSetCosmosdbSqlDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbSqlDatabaseName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation)SetSpringCloudAppId(val *string) {
	if err := j.validateSetSpringCloudAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"springCloudAppId",
		val,
	)
}

// Generates CDKTF code for importing a SpringCloudAppCosmosdbAssociation resource upon running "cdktf plan <stack-name>".
func SpringCloudAppCosmosdbAssociation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSpringCloudAppCosmosdbAssociation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
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
func SpringCloudAppCosmosdbAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudAppCosmosdbAssociation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudAppCosmosdbAssociation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudAppCosmosdbAssociation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudAppCosmosdbAssociation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudAppCosmosdbAssociation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpringCloudAppCosmosdbAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) PutTimeouts(value *SpringCloudAppCosmosdbAssociationTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ResetCosmosdbCassandraKeyspaceName() {
	_jsii_.InvokeVoid(
		s,
		"resetCosmosdbCassandraKeyspaceName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ResetCosmosdbGremlinDatabaseName() {
	_jsii_.InvokeVoid(
		s,
		"resetCosmosdbGremlinDatabaseName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ResetCosmosdbGremlinGraphName() {
	_jsii_.InvokeVoid(
		s,
		"resetCosmosdbGremlinGraphName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ResetCosmosdbMongoDatabaseName() {
	_jsii_.InvokeVoid(
		s,
		"resetCosmosdbMongoDatabaseName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ResetCosmosdbSqlDatabaseName() {
	_jsii_.InvokeVoid(
		s,
		"resetCosmosdbSqlDatabaseName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

