package digitaltwinstimeseriesdatabaseconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/digitaltwinstimeseriesdatabaseconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_time_series_database_connection azurerm_digital_twins_time_series_database_connection}.
type DigitalTwinsTimeSeriesDatabaseConnection interface {
	cdktf.TerraformResource
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DigitalTwinsId() *string
	SetDigitalTwinsId(val *string)
	DigitalTwinsIdInput() *string
	EventhubConsumerGroupName() *string
	SetEventhubConsumerGroupName(val *string)
	EventhubConsumerGroupNameInput() *string
	EventhubName() *string
	SetEventhubName(val *string)
	EventhubNameInput() *string
	EventhubNamespaceEndpointUri() *string
	SetEventhubNamespaceEndpointUri(val *string)
	EventhubNamespaceEndpointUriInput() *string
	EventhubNamespaceId() *string
	SetEventhubNamespaceId(val *string)
	EventhubNamespaceIdInput() *string
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
	KustoClusterId() *string
	SetKustoClusterId(val *string)
	KustoClusterIdInput() *string
	KustoClusterUri() *string
	SetKustoClusterUri(val *string)
	KustoClusterUriInput() *string
	KustoDatabaseName() *string
	SetKustoDatabaseName(val *string)
	KustoDatabaseNameInput() *string
	KustoTableName() *string
	SetKustoTableName(val *string)
	KustoTableNameInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DigitalTwinsTimeSeriesDatabaseConnectionTimeoutsOutputReference
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
	PutTimeouts(value *DigitalTwinsTimeSeriesDatabaseConnectionTimeouts)
	ResetEventhubConsumerGroupName()
	ResetId()
	ResetKustoTableName()
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

// The jsii proxy struct for DigitalTwinsTimeSeriesDatabaseConnection
type jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) DigitalTwinsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digitalTwinsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) DigitalTwinsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digitalTwinsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) EventhubConsumerGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubConsumerGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) EventhubConsumerGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubConsumerGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) EventhubName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) EventhubNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) EventhubNamespaceEndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNamespaceEndpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) EventhubNamespaceEndpointUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNamespaceEndpointUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) EventhubNamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNamespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) EventhubNamespaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNamespaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) KustoClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) KustoClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) KustoClusterUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoClusterUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) KustoClusterUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoClusterUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) KustoDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) KustoDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) KustoTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) KustoTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) Timeouts() DigitalTwinsTimeSeriesDatabaseConnectionTimeoutsOutputReference {
	var returns DigitalTwinsTimeSeriesDatabaseConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_time_series_database_connection azurerm_digital_twins_time_series_database_connection} Resource.
func NewDigitalTwinsTimeSeriesDatabaseConnection(scope constructs.Construct, id *string, config *DigitalTwinsTimeSeriesDatabaseConnectionConfig) DigitalTwinsTimeSeriesDatabaseConnection {
	_init_.Initialize()

	if err := validateNewDigitalTwinsTimeSeriesDatabaseConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.digitalTwinsTimeSeriesDatabaseConnection.DigitalTwinsTimeSeriesDatabaseConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_time_series_database_connection azurerm_digital_twins_time_series_database_connection} Resource.
func NewDigitalTwinsTimeSeriesDatabaseConnection_Override(d DigitalTwinsTimeSeriesDatabaseConnection, scope constructs.Construct, id *string, config *DigitalTwinsTimeSeriesDatabaseConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.digitalTwinsTimeSeriesDatabaseConnection.DigitalTwinsTimeSeriesDatabaseConnection",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetDigitalTwinsId(val *string) {
	if err := j.validateSetDigitalTwinsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digitalTwinsId",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetEventhubConsumerGroupName(val *string) {
	if err := j.validateSetEventhubConsumerGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubConsumerGroupName",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetEventhubName(val *string) {
	if err := j.validateSetEventhubNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubName",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetEventhubNamespaceEndpointUri(val *string) {
	if err := j.validateSetEventhubNamespaceEndpointUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubNamespaceEndpointUri",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetEventhubNamespaceId(val *string) {
	if err := j.validateSetEventhubNamespaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubNamespaceId",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetKustoClusterId(val *string) {
	if err := j.validateSetKustoClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kustoClusterId",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetKustoClusterUri(val *string) {
	if err := j.validateSetKustoClusterUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kustoClusterUri",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetKustoDatabaseName(val *string) {
	if err := j.validateSetKustoDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kustoDatabaseName",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetKustoTableName(val *string) {
	if err := j.validateSetKustoTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kustoTableName",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func DigitalTwinsTimeSeriesDatabaseConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDigitalTwinsTimeSeriesDatabaseConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.digitalTwinsTimeSeriesDatabaseConnection.DigitalTwinsTimeSeriesDatabaseConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DigitalTwinsTimeSeriesDatabaseConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDigitalTwinsTimeSeriesDatabaseConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.digitalTwinsTimeSeriesDatabaseConnection.DigitalTwinsTimeSeriesDatabaseConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DigitalTwinsTimeSeriesDatabaseConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDigitalTwinsTimeSeriesDatabaseConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.digitalTwinsTimeSeriesDatabaseConnection.DigitalTwinsTimeSeriesDatabaseConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DigitalTwinsTimeSeriesDatabaseConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.digitalTwinsTimeSeriesDatabaseConnection.DigitalTwinsTimeSeriesDatabaseConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) PutTimeouts(value *DigitalTwinsTimeSeriesDatabaseConnectionTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ResetEventhubConsumerGroupName() {
	_jsii_.InvokeVoid(
		d,
		"resetEventhubConsumerGroupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ResetKustoTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetKustoTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitalTwinsTimeSeriesDatabaseConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
