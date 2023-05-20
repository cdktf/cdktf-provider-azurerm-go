package kustoiothubdataconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/kustoiothubdataconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/kusto_iothub_data_connection azurerm_kusto_iothub_data_connection}.
type KustoIothubDataConnection interface {
	cdktf.TerraformResource
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
	IdInput() *string
	IothubId() *string
	SetIothubId(val *string)
	IothubIdInput() *string
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
	SharedAccessPolicyName() *string
	SetSharedAccessPolicyName(val *string)
	SharedAccessPolicyNameInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KustoIothubDataConnectionTimeoutsOutputReference
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
	PutTimeouts(value *KustoIothubDataConnectionTimeouts)
	ResetDatabaseRoutingType()
	ResetDataFormat()
	ResetEventSystemProperties()
	ResetId()
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

// The jsii proxy struct for KustoIothubDataConnection
type jsiiProxy_KustoIothubDataConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KustoIothubDataConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) ConsumerGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) ConsumerGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) DatabaseRoutingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoutingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) DatabaseRoutingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoutingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) EventSystemProperties() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSystemProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) EventSystemPropertiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSystemPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) IothubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) IothubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) MappingRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) MappingRuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) SharedAccessPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) SharedAccessPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) Timeouts() KustoIothubDataConnectionTimeoutsOutputReference {
	var returns KustoIothubDataConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoIothubDataConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/kusto_iothub_data_connection azurerm_kusto_iothub_data_connection} Resource.
func NewKustoIothubDataConnection(scope constructs.Construct, id *string, config *KustoIothubDataConnectionConfig) KustoIothubDataConnection {
	_init_.Initialize()

	if err := validateNewKustoIothubDataConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KustoIothubDataConnection{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoIothubDataConnection.KustoIothubDataConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/kusto_iothub_data_connection azurerm_kusto_iothub_data_connection} Resource.
func NewKustoIothubDataConnection_Override(k KustoIothubDataConnection, scope constructs.Construct, id *string, config *KustoIothubDataConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoIothubDataConnection.KustoIothubDataConnection",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetConsumerGroup(val *string) {
	if err := j.validateSetConsumerGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroup",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetDatabaseRoutingType(val *string) {
	if err := j.validateSetDatabaseRoutingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseRoutingType",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetEventSystemProperties(val *[]*string) {
	if err := j.validateSetEventSystemPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSystemProperties",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetIothubId(val *string) {
	if err := j.validateSetIothubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iothubId",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetMappingRuleName(val *string) {
	if err := j.validateSetMappingRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappingRuleName",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetSharedAccessPolicyName(val *string) {
	if err := j.validateSetSharedAccessPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccessPolicyName",
		val,
	)
}

func (j *jsiiProxy_KustoIothubDataConnection)SetTableName(val *string) {
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
func KustoIothubDataConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoIothubDataConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoIothubDataConnection.KustoIothubDataConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KustoIothubDataConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoIothubDataConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoIothubDataConnection.KustoIothubDataConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KustoIothubDataConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoIothubDataConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoIothubDataConnection.KustoIothubDataConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KustoIothubDataConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.kustoIothubDataConnection.KustoIothubDataConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KustoIothubDataConnection) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KustoIothubDataConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KustoIothubDataConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KustoIothubDataConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KustoIothubDataConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KustoIothubDataConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KustoIothubDataConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KustoIothubDataConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KustoIothubDataConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KustoIothubDataConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KustoIothubDataConnection) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) PutTimeouts(value *KustoIothubDataConnectionTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) ResetDatabaseRoutingType() {
	_jsii_.InvokeVoid(
		k,
		"resetDatabaseRoutingType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) ResetDataFormat() {
	_jsii_.InvokeVoid(
		k,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) ResetEventSystemProperties() {
	_jsii_.InvokeVoid(
		k,
		"resetEventSystemProperties",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) ResetMappingRuleName() {
	_jsii_.InvokeVoid(
		k,
		"resetMappingRuleName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) ResetTableName() {
	_jsii_.InvokeVoid(
		k,
		"resetTableName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoIothubDataConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoIothubDataConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoIothubDataConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoIothubDataConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

