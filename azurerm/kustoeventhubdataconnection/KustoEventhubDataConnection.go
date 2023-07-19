package kustoeventhubdataconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/kustoeventhubdataconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kusto_eventhub_data_connection azurerm_kusto_eventhub_data_connection}.
type KustoEventhubDataConnection interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumerGroup() *string
	SetConsumerGroup(val *string)
	ConsumerGroupInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	EventhubId() *string
	SetEventhubId(val *string)
	EventhubIdInput() *string
	EventSystemProperties() *[]*string
	SetEventSystemProperties(val *[]*string)
	EventSystemPropertiesInput() *[]*string
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
	IdentityId() *string
	SetIdentityId(val *string)
	IdentityIdInput() *string
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
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
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KustoEventhubDataConnectionTimeoutsOutputReference
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
	PutTimeouts(value *KustoEventhubDataConnectionTimeouts)
	ResetCompression()
	ResetDatabaseRoutingType()
	ResetDataFormat()
	ResetEventSystemProperties()
	ResetId()
	ResetIdentityId()
	ResetMappingRuleName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for KustoEventhubDataConnection
type jsiiProxy_KustoEventhubDataConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KustoEventhubDataConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) ConsumerGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) ConsumerGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) DatabaseRoutingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoutingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) DatabaseRoutingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoutingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) EventhubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) EventhubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) EventSystemProperties() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSystemProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) EventSystemPropertiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSystemPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) IdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) IdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) MappingRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) MappingRuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) Timeouts() KustoEventhubDataConnectionTimeoutsOutputReference {
	var returns KustoEventhubDataConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoEventhubDataConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kusto_eventhub_data_connection azurerm_kusto_eventhub_data_connection} Resource.
func NewKustoEventhubDataConnection(scope constructs.Construct, id *string, config *KustoEventhubDataConnectionConfig) KustoEventhubDataConnection {
	_init_.Initialize()

	if err := validateNewKustoEventhubDataConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KustoEventhubDataConnection{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoEventhubDataConnection.KustoEventhubDataConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kusto_eventhub_data_connection azurerm_kusto_eventhub_data_connection} Resource.
func NewKustoEventhubDataConnection_Override(k KustoEventhubDataConnection, scope constructs.Construct, id *string, config *KustoEventhubDataConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoEventhubDataConnection.KustoEventhubDataConnection",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetConsumerGroup(val *string) {
	if err := j.validateSetConsumerGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroup",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetDatabaseRoutingType(val *string) {
	if err := j.validateSetDatabaseRoutingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseRoutingType",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetEventhubId(val *string) {
	if err := j.validateSetEventhubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubId",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetEventSystemProperties(val *[]*string) {
	if err := j.validateSetEventSystemPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSystemProperties",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetIdentityId(val *string) {
	if err := j.validateSetIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityId",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetMappingRuleName(val *string) {
	if err := j.validateSetMappingRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappingRuleName",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KustoEventhubDataConnection)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
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
func KustoEventhubDataConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoEventhubDataConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoEventhubDataConnection.KustoEventhubDataConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KustoEventhubDataConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoEventhubDataConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoEventhubDataConnection.KustoEventhubDataConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KustoEventhubDataConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoEventhubDataConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoEventhubDataConnection.KustoEventhubDataConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KustoEventhubDataConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.kustoEventhubDataConnection.KustoEventhubDataConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) PutTimeouts(value *KustoEventhubDataConnectionTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetCompression() {
	_jsii_.InvokeVoid(
		k,
		"resetCompression",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetDatabaseRoutingType() {
	_jsii_.InvokeVoid(
		k,
		"resetDatabaseRoutingType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetDataFormat() {
	_jsii_.InvokeVoid(
		k,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetEventSystemProperties() {
	_jsii_.InvokeVoid(
		k,
		"resetEventSystemProperties",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetIdentityId() {
	_jsii_.InvokeVoid(
		k,
		"resetIdentityId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetMappingRuleName() {
	_jsii_.InvokeVoid(
		k,
		"resetMappingRuleName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetTableName() {
	_jsii_.InvokeVoid(
		k,
		"resetTableName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoEventhubDataConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoEventhubDataConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

