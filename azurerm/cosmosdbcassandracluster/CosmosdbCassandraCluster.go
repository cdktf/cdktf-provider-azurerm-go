package cosmosdbcassandracluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/cosmosdbcassandracluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/cosmosdb_cassandra_cluster azurerm_cosmosdb_cassandra_cluster}.
type CosmosdbCassandraCluster interface {
	cdktf.TerraformResource
	AuthenticationMethod() *string
	SetAuthenticationMethod(val *string)
	AuthenticationMethodInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificatePems() *[]*string
	SetClientCertificatePems(val *[]*string)
	ClientCertificatePemsInput() *[]*string
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
	DefaultAdminPassword() *string
	SetDefaultAdminPassword(val *string)
	DefaultAdminPasswordInput() *string
	DelegatedManagementSubnetId() *string
	SetDelegatedManagementSubnetId(val *string)
	DelegatedManagementSubnetIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExternalGossipCertificatePems() *[]*string
	SetExternalGossipCertificatePems(val *[]*string)
	ExternalGossipCertificatePemsInput() *[]*string
	ExternalSeedNodeIpAddresses() *[]*string
	SetExternalSeedNodeIpAddresses(val *[]*string)
	ExternalSeedNodeIpAddressesInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HoursBetweenBackups() *float64
	SetHoursBetweenBackups(val *float64)
	HoursBetweenBackupsInput() *float64
	Id() *string
	SetId(val *string)
	Identity() CosmosdbCassandraClusterIdentityOutputReference
	IdentityInput() *CosmosdbCassandraClusterIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
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
	RepairEnabled() interface{}
	SetRepairEnabled(val interface{})
	RepairEnabledInput() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CosmosdbCassandraClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutIdentity(value *CosmosdbCassandraClusterIdentity)
	PutTimeouts(value *CosmosdbCassandraClusterTimeouts)
	ResetAuthenticationMethod()
	ResetClientCertificatePems()
	ResetExternalGossipCertificatePems()
	ResetExternalSeedNodeIpAddresses()
	ResetHoursBetweenBackups()
	ResetId()
	ResetIdentity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRepairEnabled()
	ResetTags()
	ResetTimeouts()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CosmosdbCassandraCluster
type jsiiProxy_CosmosdbCassandraCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CosmosdbCassandraCluster) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ClientCertificatePems() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientCertificatePems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ClientCertificatePemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientCertificatePemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) DefaultAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) DefaultAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) DelegatedManagementSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedManagementSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) DelegatedManagementSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedManagementSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ExternalGossipCertificatePems() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalGossipCertificatePems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ExternalGossipCertificatePemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalGossipCertificatePemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ExternalSeedNodeIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalSeedNodeIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ExternalSeedNodeIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalSeedNodeIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) HoursBetweenBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hoursBetweenBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) HoursBetweenBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hoursBetweenBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Identity() CosmosdbCassandraClusterIdentityOutputReference {
	var returns CosmosdbCassandraClusterIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) IdentityInput() *CosmosdbCassandraClusterIdentity {
	var returns *CosmosdbCassandraClusterIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) RepairEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repairEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) RepairEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repairEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Timeouts() CosmosdbCassandraClusterTimeoutsOutputReference {
	var returns CosmosdbCassandraClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/cosmosdb_cassandra_cluster azurerm_cosmosdb_cassandra_cluster} Resource.
func NewCosmosdbCassandraCluster(scope constructs.Construct, id *string, config *CosmosdbCassandraClusterConfig) CosmosdbCassandraCluster {
	_init_.Initialize()

	if err := validateNewCosmosdbCassandraClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbCassandraCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbCassandraCluster.CosmosdbCassandraCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/cosmosdb_cassandra_cluster azurerm_cosmosdb_cassandra_cluster} Resource.
func NewCosmosdbCassandraCluster_Override(c CosmosdbCassandraCluster, scope constructs.Construct, id *string, config *CosmosdbCassandraClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbCassandraCluster.CosmosdbCassandraCluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetClientCertificatePems(val *[]*string) {
	if err := j.validateSetClientCertificatePemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificatePems",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetDefaultAdminPassword(val *string) {
	if err := j.validateSetDefaultAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAdminPassword",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetDelegatedManagementSubnetId(val *string) {
	if err := j.validateSetDelegatedManagementSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegatedManagementSubnetId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetExternalGossipCertificatePems(val *[]*string) {
	if err := j.validateSetExternalGossipCertificatePemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalGossipCertificatePems",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetExternalSeedNodeIpAddresses(val *[]*string) {
	if err := j.validateSetExternalSeedNodeIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalSeedNodeIpAddresses",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetHoursBetweenBackups(val *float64) {
	if err := j.validateSetHoursBetweenBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hoursBetweenBackups",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetRepairEnabled(val interface{}) {
	if err := j.validateSetRepairEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repairEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraCluster)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
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
func CosmosdbCassandraCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbCassandraCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbCassandraCluster.CosmosdbCassandraCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbCassandraCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbCassandraCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbCassandraCluster.CosmosdbCassandraCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbCassandraCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbCassandraCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbCassandraCluster.CosmosdbCassandraCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CosmosdbCassandraCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cosmosdbCassandraCluster.CosmosdbCassandraCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CosmosdbCassandraCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbCassandraCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) PutIdentity(value *CosmosdbCassandraClusterIdentity) {
	if err := c.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdentity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) PutTimeouts(value *CosmosdbCassandraClusterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetClientCertificatePems() {
	_jsii_.InvokeVoid(
		c,
		"resetClientCertificatePems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetExternalGossipCertificatePems() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalGossipCertificatePems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetExternalSeedNodeIpAddresses() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalSeedNodeIpAddresses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetHoursBetweenBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetHoursBetweenBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetRepairEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetRepairEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ResetVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

