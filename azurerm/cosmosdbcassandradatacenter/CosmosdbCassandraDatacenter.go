package cosmosdbcassandradatacenter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/cosmosdbcassandradatacenter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/cosmosdb_cassandra_datacenter azurerm_cosmosdb_cassandra_datacenter}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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

func (j *jsiiProxy_CosmosdbCassandraDatacenter) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/cosmosdb_cassandra_datacenter azurerm_cosmosdb_cassandra_datacenter} Resource.
func NewCosmosdbCassandraDatacenter(scope constructs.Construct, id *string, config *CosmosdbCassandraDatacenterConfig) CosmosdbCassandraDatacenter {
	_init_.Initialize()

	if err := validateNewCosmosdbCassandraDatacenterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbCassandraDatacenter{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/cosmosdb_cassandra_datacenter azurerm_cosmosdb_cassandra_datacenter} Resource.
func NewCosmosdbCassandraDatacenter_Override(c CosmosdbCassandraDatacenter, scope constructs.Construct, id *string, config *CosmosdbCassandraDatacenterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetAvailabilityZonesEnabled(val interface{}) {
	if err := j.validateSetAvailabilityZonesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZonesEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetBackupStorageCustomerKeyUri(val *string) {
	if err := j.validateSetBackupStorageCustomerKeyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupStorageCustomerKeyUri",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetBase64EncodedYamlFragment(val *string) {
	if err := j.validateSetBase64EncodedYamlFragmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"base64EncodedYamlFragment",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetCassandraClusterId(val *string) {
	if err := j.validateSetCassandraClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cassandraClusterId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetDelegatedManagementSubnetId(val *string) {
	if err := j.validateSetDelegatedManagementSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegatedManagementSubnetId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetDiskCount(val *float64) {
	if err := j.validateSetDiskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskCount",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetDiskSku(val *string) {
	if err := j.validateSetDiskSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSku",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetManagedDiskCustomerKeyUri(val *string) {
	if err := j.validateSetManagedDiskCustomerKeyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDiskCustomerKeyUri",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CosmosdbCassandraDatacenter)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateCosmosdbCassandraDatacenter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbCassandraDatacenter_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbCassandraDatacenter_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbCassandraDatacenter_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbCassandraDatacenter_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbCassandraDatacenter.CosmosdbCassandraDatacenter",
		"isTerraformResource",
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
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbCassandraDatacenter) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosmosdbCassandraDatacenter) PutTimeouts(value *CosmosdbCassandraDatacenterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

