package cosmosdbcassandradatacenter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/cosmosdbcassandradatacenter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter azurerm_cosmosdb_cassandra_datacenter}.
type CosmosdbCassandraDatacenter interface {
	cdktf.TerraformResource
	AvailabilityZonesEnabled() interface{}
	SetAvailabilityZonesEnabled(val interface{})
	AvailabilityZonesEnabledInput() interface{}
	BackupStorageCustomerKeyUri() *string
	SetBackupStorageCustomerKeyUri(val *string)
	BackupStorageCustomerKeyUriInput() *string
	Base64EncodedYamlFragment() *string
	SetBase64EncodedYamlFragment(val *string)
	Base64EncodedYamlFragmentInput() *string
	CassandraClusterId() *string
	SetCassandraClusterId(val *string)
	CassandraClusterIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DelegatedManagementSubnetId() *string
	SetDelegatedManagementSubnetId(val *string)
	DelegatedManagementSubnetIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskCount() *float64
	SetDiskCount(val *float64)
	DiskCountInput() *float64
	DiskSku() *string
	SetDiskSku(val *string)
	DiskSkuInput() *string
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedDiskCustomerKeyUri() *string
	SetManagedDiskCustomerKeyUri(val *string)
	ManagedDiskCustomerKeyUriInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
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
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CosmosdbCassandraDatacenterTimeoutsOutputReference
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
	PutTimeouts(value *CosmosdbCassandraDatacenterTimeouts)
	ResetAvailabilityZonesEnabled()
	ResetBackupStorageCustomerKeyUri()
	ResetBase64EncodedYamlFragment()
	ResetDiskCount()
	ResetDiskSku()
	ResetId()
	ResetManagedDiskCustomerKeyUri()
	ResetNodeCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSkuName()
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

// The jsii proxy struct for CosmosdbCassandraDatacenter
type jsiiProxy_CosmosdbCassandraDatacenter struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) AvailabilityZonesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZonesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) AvailabilityZonesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZonesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) BackupStorageCustomerKeyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupStorageCustomerKeyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) BackupStorageCustomerKeyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupStorageCustomerKeyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Base64EncodedYamlFragment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"base64EncodedYamlFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Base64EncodedYamlFragmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"base64EncodedYamlFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) CassandraClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cassandraClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) CassandraClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cassandraClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) DelegatedManagementSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedManagementSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) DelegatedManagementSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedManagementSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) DiskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) DiskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) DiskSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) DiskSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) ManagedDiskCustomerKeyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskCustomerKeyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) ManagedDiskCustomerKeyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskCustomerKeyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Timeouts() CosmosdbCassandraDatacenterTimeoutsOutputReference {
	var returns CosmosdbCassandraDatacenterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter azurerm_cosmosdb_cassandra_datacenter} Resource.
func NewCosmosdbCassandraDatacenter(scope constructs.Construct, id *string, config *CosmosdbCassandraDatacenterConfig) CosmosdbCassandraDatacenter {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbCassandraDatacenter{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter azurerm_cosmosdb_cassandra_datacenter} Resource.
func NewCosmosdbCassandraDatacenter_Override(c CosmosdbCassandraDatacenter, scope constructs.Construct, id *string, config *CosmosdbCassandraDatacenterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetAvailabilityZonesEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"availabilityZonesEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetBackupStorageCustomerKeyUri(val *string) {
	_jsii_.Set(
		j,
		"backupStorageCustomerKeyUri",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetBase64EncodedYamlFragment(val *string) {
	_jsii_.Set(
		j,
		"base64EncodedYamlFragment",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetCassandraClusterId(val *string) {
	_jsii_.Set(
		j,
		"cassandraClusterId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetDelegatedManagementSubnetId(val *string) {
	_jsii_.Set(
		j,
		"delegatedManagementSubnetId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetDiskCount(val *float64) {
	_jsii_.Set(
		j,
		"diskCount",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetDiskSku(val *string) {
	_jsii_.Set(
		j,
		"diskSku",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetManagedDiskCustomerKeyUri(val *string) {
	_jsii_.Set(
		j,
		"managedDiskCustomerKeyUri",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetNodeCount(val *float64) {
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter) SetSkuName(val *string) {
	_jsii_.Set(
		j,
		"skuName",
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
func CosmosdbCassandraDatacenter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CosmosdbCassandraDatacenter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) PutTimeouts(value *CosmosdbCassandraDatacenterTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetAvailabilityZonesEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailabilityZonesEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetBackupStorageCustomerKeyUri() {
	_jsii_.InvokeVoid(
		c,
		"resetBackupStorageCustomerKeyUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetBase64EncodedYamlFragment() {
	_jsii_.InvokeVoid(
		c,
		"resetBase64EncodedYamlFragment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetDiskCount() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetDiskSku() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskSku",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetManagedDiskCustomerKeyUri() {
	_jsii_.InvokeVoid(
		c,
		"resetManagedDiskCustomerKeyUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetNodeCount() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetSkuName() {
	_jsii_.InvokeVoid(
		c,
		"resetSkuName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbCassandraDatacenterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#cassandra_cluster_id CosmosdbCassandraDatacenter#cassandra_cluster_id}.
	CassandraClusterId *string `field:"required" json:"cassandraClusterId" yaml:"cassandraClusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#delegated_management_subnet_id CosmosdbCassandraDatacenter#delegated_management_subnet_id}.
	DelegatedManagementSubnetId *string `field:"required" json:"delegatedManagementSubnetId" yaml:"delegatedManagementSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#location CosmosdbCassandraDatacenter#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#name CosmosdbCassandraDatacenter#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#availability_zones_enabled CosmosdbCassandraDatacenter#availability_zones_enabled}.
	AvailabilityZonesEnabled interface{} `field:"optional" json:"availabilityZonesEnabled" yaml:"availabilityZonesEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#backup_storage_customer_key_uri CosmosdbCassandraDatacenter#backup_storage_customer_key_uri}.
	BackupStorageCustomerKeyUri *string `field:"optional" json:"backupStorageCustomerKeyUri" yaml:"backupStorageCustomerKeyUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#base64_encoded_yaml_fragment CosmosdbCassandraDatacenter#base64_encoded_yaml_fragment}.
	Base64EncodedYamlFragment *string `field:"optional" json:"base64EncodedYamlFragment" yaml:"base64EncodedYamlFragment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#disk_count CosmosdbCassandraDatacenter#disk_count}.
	DiskCount *float64 `field:"optional" json:"diskCount" yaml:"diskCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#disk_sku CosmosdbCassandraDatacenter#disk_sku}.
	DiskSku *string `field:"optional" json:"diskSku" yaml:"diskSku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#id CosmosdbCassandraDatacenter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#managed_disk_customer_key_uri CosmosdbCassandraDatacenter#managed_disk_customer_key_uri}.
	ManagedDiskCustomerKeyUri *string `field:"optional" json:"managedDiskCustomerKeyUri" yaml:"managedDiskCustomerKeyUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#node_count CosmosdbCassandraDatacenter#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#sku_name CosmosdbCassandraDatacenter#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#timeouts CosmosdbCassandraDatacenter#timeouts}
	Timeouts *CosmosdbCassandraDatacenterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type CosmosdbCassandraDatacenterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#create CosmosdbCassandraDatacenter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#delete CosmosdbCassandraDatacenter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#read CosmosdbCassandraDatacenter#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_datacenter#update CosmosdbCassandraDatacenter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type CosmosdbCassandraDatacenterTimeoutsOutputReference interface {
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

// The jsii proxy struct for CosmosdbCassandraDatacenterTimeoutsOutputReference
type jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewCosmosdbCassandraDatacenterTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbCassandraDatacenterTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbCassandraDatacenterTimeoutsOutputReference_Override(c CosmosdbCassandraDatacenterTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		c,
		"resetRead",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbCassandraDatacenterTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

