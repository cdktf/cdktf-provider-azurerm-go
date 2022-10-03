package springcloudappcosmosdbassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/springcloudappcosmosdbassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association azurerm_spring_cloud_app_cosmosdb_association}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association azurerm_spring_cloud_app_cosmosdb_association} Resource.
func NewSpringCloudAppCosmosdbAssociation(scope constructs.Construct, id *string, config *SpringCloudAppCosmosdbAssociationConfig) SpringCloudAppCosmosdbAssociation {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudAppCosmosdbAssociation{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association azurerm_spring_cloud_app_cosmosdb_association} Resource.
func NewSpringCloudAppCosmosdbAssociation_Override(s SpringCloudAppCosmosdbAssociation, scope constructs.Construct, id *string, config *SpringCloudAppCosmosdbAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetApiType(val *string) {
	_jsii_.Set(
		j,
		"apiType",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetCosmosdbAccessKey(val *string) {
	_jsii_.Set(
		j,
		"cosmosdbAccessKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetCosmosdbAccountId(val *string) {
	_jsii_.Set(
		j,
		"cosmosdbAccountId",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetCosmosdbCassandraKeyspaceName(val *string) {
	_jsii_.Set(
		j,
		"cosmosdbCassandraKeyspaceName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetCosmosdbGremlinDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"cosmosdbGremlinDatabaseName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetCosmosdbGremlinGraphName(val *string) {
	_jsii_.Set(
		j,
		"cosmosdbGremlinGraphName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetCosmosdbMongoDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"cosmosdbMongoDatabaseName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetCosmosdbSqlDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"cosmosdbSqlDatabaseName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociation) SetSpringCloudAppId(val *string) {
	_jsii_.Set(
		j,
		"springCloudAppId",
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
func SpringCloudAppCosmosdbAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociation",
		"isConstruct",
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

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociation) PutTimeouts(value *SpringCloudAppCosmosdbAssociationTimeouts) {
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

type SpringCloudAppCosmosdbAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#api_type SpringCloudAppCosmosdbAssociation#api_type}.
	ApiType *string `field:"required" json:"apiType" yaml:"apiType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#cosmosdb_access_key SpringCloudAppCosmosdbAssociation#cosmosdb_access_key}.
	CosmosdbAccessKey *string `field:"required" json:"cosmosdbAccessKey" yaml:"cosmosdbAccessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#cosmosdb_account_id SpringCloudAppCosmosdbAssociation#cosmosdb_account_id}.
	CosmosdbAccountId *string `field:"required" json:"cosmosdbAccountId" yaml:"cosmosdbAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#name SpringCloudAppCosmosdbAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#spring_cloud_app_id SpringCloudAppCosmosdbAssociation#spring_cloud_app_id}.
	SpringCloudAppId *string `field:"required" json:"springCloudAppId" yaml:"springCloudAppId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#cosmosdb_cassandra_keyspace_name SpringCloudAppCosmosdbAssociation#cosmosdb_cassandra_keyspace_name}.
	CosmosdbCassandraKeyspaceName *string `field:"optional" json:"cosmosdbCassandraKeyspaceName" yaml:"cosmosdbCassandraKeyspaceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#cosmosdb_gremlin_database_name SpringCloudAppCosmosdbAssociation#cosmosdb_gremlin_database_name}.
	CosmosdbGremlinDatabaseName *string `field:"optional" json:"cosmosdbGremlinDatabaseName" yaml:"cosmosdbGremlinDatabaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#cosmosdb_gremlin_graph_name SpringCloudAppCosmosdbAssociation#cosmosdb_gremlin_graph_name}.
	CosmosdbGremlinGraphName *string `field:"optional" json:"cosmosdbGremlinGraphName" yaml:"cosmosdbGremlinGraphName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#cosmosdb_mongo_database_name SpringCloudAppCosmosdbAssociation#cosmosdb_mongo_database_name}.
	CosmosdbMongoDatabaseName *string `field:"optional" json:"cosmosdbMongoDatabaseName" yaml:"cosmosdbMongoDatabaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#cosmosdb_sql_database_name SpringCloudAppCosmosdbAssociation#cosmosdb_sql_database_name}.
	CosmosdbSqlDatabaseName *string `field:"optional" json:"cosmosdbSqlDatabaseName" yaml:"cosmosdbSqlDatabaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#id SpringCloudAppCosmosdbAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#timeouts SpringCloudAppCosmosdbAssociation#timeouts}
	Timeouts *SpringCloudAppCosmosdbAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type SpringCloudAppCosmosdbAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#create SpringCloudAppCosmosdbAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#delete SpringCloudAppCosmosdbAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#read SpringCloudAppCosmosdbAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_app_cosmosdb_association#update SpringCloudAppCosmosdbAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type SpringCloudAppCosmosdbAssociationTimeoutsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudAppCosmosdbAssociationTimeoutsOutputReference
type jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewSpringCloudAppCosmosdbAssociationTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudAppCosmosdbAssociationTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudAppCosmosdbAssociationTimeoutsOutputReference_Override(s SpringCloudAppCosmosdbAssociationTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudAppCosmosdbAssociation.SpringCloudAppCosmosdbAssociationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppCosmosdbAssociationTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

