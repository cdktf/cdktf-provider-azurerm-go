package kustoeventgriddataconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/kustoeventgriddataconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection azurerm_kusto_eventgrid_data_connection}.
type KustoEventgridDataConnection interface {
	cdktf.TerraformResource
	BlobStorageEventType() *string
	SetBlobStorageEventType(val *string)
	BlobStorageEventTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	DatabaseRoutingType() *string
	SetDatabaseRoutingType(val *string)
	DatabaseRoutingTypeInput() *string
	DataFormat() *string
	SetDataFormat(val *string)
	DataFormatInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventgridResourceId() *string
	SetEventgridResourceId(val *string)
	EventgridResourceIdInput() *string
	EventhubConsumerGroupName() *string
	SetEventhubConsumerGroupName(val *string)
	EventhubConsumerGroupNameInput() *string
	EventhubId() *string
	SetEventhubId(val *string)
	EventhubIdInput() *string
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
	ManagedIdentityResourceId() *string
	SetManagedIdentityResourceId(val *string)
	ManagedIdentityResourceIdInput() *string
	MappingRuleName() *string
	SetMappingRuleName(val *string)
	MappingRuleNameInput() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SkipFirstRecord() interface{}
	SetSkipFirstRecord(val interface{})
	SkipFirstRecordInput() interface{}
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KustoEventgridDataConnectionTimeoutsOutputReference
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
	PutTimeouts(value *KustoEventgridDataConnectionTimeouts)
	ResetBlobStorageEventType()
	ResetDatabaseRoutingType()
	ResetDataFormat()
	ResetEventgridResourceId()
	ResetId()
	ResetManagedIdentityResourceId()
	ResetMappingRuleName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSkipFirstRecord()
	ResetTableName()
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

// The jsii proxy struct for KustoEventgridDataConnection
type jsiiProxy_KustoEventgridDataConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KustoEventgridDataConnection) BlobStorageEventType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobStorageEventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) BlobStorageEventTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobStorageEventTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) DatabaseRoutingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoutingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) DatabaseRoutingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoutingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) EventgridResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventgridResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) EventgridResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventgridResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) EventhubConsumerGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubConsumerGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) EventhubConsumerGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubConsumerGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) EventhubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) EventhubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) ManagedIdentityResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) ManagedIdentityResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) MappingRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) MappingRuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) SkipFirstRecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFirstRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) SkipFirstRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFirstRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) Timeouts() KustoEventgridDataConnectionTimeoutsOutputReference {
	var returns KustoEventgridDataConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection azurerm_kusto_eventgrid_data_connection} Resource.
func NewKustoEventgridDataConnection(scope constructs.Construct, id *string, config *KustoEventgridDataConnectionConfig) KustoEventgridDataConnection {
	_init_.Initialize()

	j := jsiiProxy_KustoEventgridDataConnection{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoEventgridDataConnection.KustoEventgridDataConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection azurerm_kusto_eventgrid_data_connection} Resource.
func NewKustoEventgridDataConnection_Override(k KustoEventgridDataConnection, scope constructs.Construct, id *string, config *KustoEventgridDataConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoEventgridDataConnection.KustoEventgridDataConnection",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetBlobStorageEventType(val *string) {
	_jsii_.Set(
		j,
		"blobStorageEventType",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetDatabaseRoutingType(val *string) {
	_jsii_.Set(
		j,
		"databaseRoutingType",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetDataFormat(val *string) {
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetEventgridResourceId(val *string) {
	_jsii_.Set(
		j,
		"eventgridResourceId",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetEventhubConsumerGroupName(val *string) {
	_jsii_.Set(
		j,
		"eventhubConsumerGroupName",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetEventhubId(val *string) {
	_jsii_.Set(
		j,
		"eventhubId",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetManagedIdentityResourceId(val *string) {
	_jsii_.Set(
		j,
		"managedIdentityResourceId",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetMappingRuleName(val *string) {
	_jsii_.Set(
		j,
		"mappingRuleName",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetSkipFirstRecord(val interface{}) {
	_jsii_.Set(
		j,
		"skipFirstRecord",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetStorageAccountId(val *string) {
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnection) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
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
func KustoEventgridDataConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoEventgridDataConnection.KustoEventgridDataConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KustoEventgridDataConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.kustoEventgridDataConnection.KustoEventgridDataConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) PutTimeouts(value *KustoEventgridDataConnectionTimeouts) {
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetBlobStorageEventType() {
	_jsii_.InvokeVoid(
		k,
		"resetBlobStorageEventType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetDatabaseRoutingType() {
	_jsii_.InvokeVoid(
		k,
		"resetDatabaseRoutingType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetDataFormat() {
	_jsii_.InvokeVoid(
		k,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetEventgridResourceId() {
	_jsii_.InvokeVoid(
		k,
		"resetEventgridResourceId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetManagedIdentityResourceId() {
	_jsii_.InvokeVoid(
		k,
		"resetManagedIdentityResourceId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetMappingRuleName() {
	_jsii_.InvokeVoid(
		k,
		"resetMappingRuleName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetSkipFirstRecord() {
	_jsii_.InvokeVoid(
		k,
		"resetSkipFirstRecord",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetTableName() {
	_jsii_.InvokeVoid(
		k,
		"resetTableName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KustoEventgridDataConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#cluster_name KustoEventgridDataConnection#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#database_name KustoEventgridDataConnection#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#eventhub_consumer_group_name KustoEventgridDataConnection#eventhub_consumer_group_name}.
	EventhubConsumerGroupName *string `field:"required" json:"eventhubConsumerGroupName" yaml:"eventhubConsumerGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#eventhub_id KustoEventgridDataConnection#eventhub_id}.
	EventhubId *string `field:"required" json:"eventhubId" yaml:"eventhubId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#location KustoEventgridDataConnection#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#name KustoEventgridDataConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#resource_group_name KustoEventgridDataConnection#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#storage_account_id KustoEventgridDataConnection#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#blob_storage_event_type KustoEventgridDataConnection#blob_storage_event_type}.
	BlobStorageEventType *string `field:"optional" json:"blobStorageEventType" yaml:"blobStorageEventType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#database_routing_type KustoEventgridDataConnection#database_routing_type}.
	DatabaseRoutingType *string `field:"optional" json:"databaseRoutingType" yaml:"databaseRoutingType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#data_format KustoEventgridDataConnection#data_format}.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#eventgrid_resource_id KustoEventgridDataConnection#eventgrid_resource_id}.
	EventgridResourceId *string `field:"optional" json:"eventgridResourceId" yaml:"eventgridResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#id KustoEventgridDataConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#managed_identity_resource_id KustoEventgridDataConnection#managed_identity_resource_id}.
	ManagedIdentityResourceId *string `field:"optional" json:"managedIdentityResourceId" yaml:"managedIdentityResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#mapping_rule_name KustoEventgridDataConnection#mapping_rule_name}.
	MappingRuleName *string `field:"optional" json:"mappingRuleName" yaml:"mappingRuleName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#skip_first_record KustoEventgridDataConnection#skip_first_record}.
	SkipFirstRecord interface{} `field:"optional" json:"skipFirstRecord" yaml:"skipFirstRecord"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#table_name KustoEventgridDataConnection#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#timeouts KustoEventgridDataConnection#timeouts}
	Timeouts *KustoEventgridDataConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type KustoEventgridDataConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#create KustoEventgridDataConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#delete KustoEventgridDataConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#read KustoEventgridDataConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_eventgrid_data_connection#update KustoEventgridDataConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type KustoEventgridDataConnectionTimeoutsOutputReference interface {
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

// The jsii proxy struct for KustoEventgridDataConnectionTimeoutsOutputReference
type jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewKustoEventgridDataConnectionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoEventgridDataConnectionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoEventgridDataConnection.KustoEventgridDataConnectionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoEventgridDataConnectionTimeoutsOutputReference_Override(k KustoEventgridDataConnectionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoEventgridDataConnection.KustoEventgridDataConnectionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		k,
		"resetCreate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		k,
		"resetDelete",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		k,
		"resetRead",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		k,
		"resetUpdate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventgridDataConnectionTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

