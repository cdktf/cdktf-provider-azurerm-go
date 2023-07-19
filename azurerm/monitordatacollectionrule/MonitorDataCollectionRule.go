package monitordatacollectionrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/monitordatacollectionrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_data_collection_rule azurerm_monitor_data_collection_rule}.
type MonitorDataCollectionRule interface {
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataCollectionEndpointId() *string
	SetDataCollectionEndpointId(val *string)
	DataCollectionEndpointIdInput() *string
	DataFlow() MonitorDataCollectionRuleDataFlowList
	DataFlowInput() interface{}
	DataSources() MonitorDataCollectionRuleDataSourcesOutputReference
	DataSourcesInput() *MonitorDataCollectionRuleDataSources
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Destinations() MonitorDataCollectionRuleDestinationsOutputReference
	DestinationsInput() *MonitorDataCollectionRuleDestinations
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
	Identity() MonitorDataCollectionRuleIdentityOutputReference
	IdentityInput() *MonitorDataCollectionRuleIdentity
	IdInput() *string
	ImmutableId() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	StreamDeclaration() MonitorDataCollectionRuleStreamDeclarationList
	StreamDeclarationInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MonitorDataCollectionRuleTimeoutsOutputReference
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
	PutDataFlow(value interface{})
	PutDataSources(value *MonitorDataCollectionRuleDataSources)
	PutDestinations(value *MonitorDataCollectionRuleDestinations)
	PutIdentity(value *MonitorDataCollectionRuleIdentity)
	PutStreamDeclaration(value interface{})
	PutTimeouts(value *MonitorDataCollectionRuleTimeouts)
	ResetDataCollectionEndpointId()
	ResetDataSources()
	ResetDescription()
	ResetId()
	ResetIdentity()
	ResetKind()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStreamDeclaration()
	ResetTags()
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

// The jsii proxy struct for MonitorDataCollectionRule
type jsiiProxy_MonitorDataCollectionRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitorDataCollectionRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DataCollectionEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCollectionEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DataCollectionEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCollectionEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DataFlow() MonitorDataCollectionRuleDataFlowList {
	var returns MonitorDataCollectionRuleDataFlowList
	_jsii_.Get(
		j,
		"dataFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DataFlowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DataSources() MonitorDataCollectionRuleDataSourcesOutputReference {
	var returns MonitorDataCollectionRuleDataSourcesOutputReference
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DataSourcesInput() *MonitorDataCollectionRuleDataSources {
	var returns *MonitorDataCollectionRuleDataSources
	_jsii_.Get(
		j,
		"dataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Destinations() MonitorDataCollectionRuleDestinationsOutputReference {
	var returns MonitorDataCollectionRuleDestinationsOutputReference
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) DestinationsInput() *MonitorDataCollectionRuleDestinations {
	var returns *MonitorDataCollectionRuleDestinations
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Identity() MonitorDataCollectionRuleIdentityOutputReference {
	var returns MonitorDataCollectionRuleIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) IdentityInput() *MonitorDataCollectionRuleIdentity {
	var returns *MonitorDataCollectionRuleIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) ImmutableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"immutableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) StreamDeclaration() MonitorDataCollectionRuleStreamDeclarationList {
	var returns MonitorDataCollectionRuleStreamDeclarationList
	_jsii_.Get(
		j,
		"streamDeclaration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) StreamDeclarationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamDeclarationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) Timeouts() MonitorDataCollectionRuleTimeoutsOutputReference {
	var returns MonitorDataCollectionRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_data_collection_rule azurerm_monitor_data_collection_rule} Resource.
func NewMonitorDataCollectionRule(scope constructs.Construct, id *string, config *MonitorDataCollectionRuleConfig) MonitorDataCollectionRule {
	_init_.Initialize()

	if err := validateNewMonitorDataCollectionRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorDataCollectionRule{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_data_collection_rule azurerm_monitor_data_collection_rule} Resource.
func NewMonitorDataCollectionRule_Override(m MonitorDataCollectionRule, scope constructs.Construct, id *string, config *MonitorDataCollectionRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRule",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetDataCollectionEndpointId(val *string) {
	if err := j.validateSetDataCollectionEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCollectionEndpointId",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRule)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func MonitorDataCollectionRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorDataCollectionRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorDataCollectionRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorDataCollectionRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorDataCollectionRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorDataCollectionRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorDataCollectionRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) PutDataFlow(value interface{}) {
	if err := m.validatePutDataFlowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDataFlow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) PutDataSources(value *MonitorDataCollectionRuleDataSources) {
	if err := m.validatePutDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDataSources",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) PutDestinations(value *MonitorDataCollectionRuleDestinations) {
	if err := m.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDestinations",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) PutIdentity(value *MonitorDataCollectionRuleIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) PutStreamDeclaration(value interface{}) {
	if err := m.validatePutStreamDeclarationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStreamDeclaration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) PutTimeouts(value *MonitorDataCollectionRuleTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetDataCollectionEndpointId() {
	_jsii_.InvokeVoid(
		m,
		"resetDataCollectionEndpointId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetDataSources() {
	_jsii_.InvokeVoid(
		m,
		"resetDataSources",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetKind() {
	_jsii_.InvokeVoid(
		m,
		"resetKind",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetStreamDeclaration() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamDeclaration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

