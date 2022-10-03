package networksecurityrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/networksecurityrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule azurerm_network_security_rule}.
type NetworkSecurityRule interface {
	cdktf.TerraformResource
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationAddressPrefix() *string
	SetDestinationAddressPrefix(val *string)
	DestinationAddressPrefixes() *[]*string
	SetDestinationAddressPrefixes(val *[]*string)
	DestinationAddressPrefixesInput() *[]*string
	DestinationAddressPrefixInput() *string
	DestinationApplicationSecurityGroupIds() *[]*string
	SetDestinationApplicationSecurityGroupIds(val *[]*string)
	DestinationApplicationSecurityGroupIdsInput() *[]*string
	DestinationPortRange() *string
	SetDestinationPortRange(val *string)
	DestinationPortRangeInput() *string
	DestinationPortRanges() *[]*string
	SetDestinationPortRanges(val *[]*string)
	DestinationPortRangesInput() *[]*string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
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
	NetworkSecurityGroupName() *string
	SetNetworkSecurityGroupName(val *string)
	NetworkSecurityGroupNameInput() *string
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	SourceAddressPrefix() *string
	SetSourceAddressPrefix(val *string)
	SourceAddressPrefixes() *[]*string
	SetSourceAddressPrefixes(val *[]*string)
	SourceAddressPrefixesInput() *[]*string
	SourceAddressPrefixInput() *string
	SourceApplicationSecurityGroupIds() *[]*string
	SetSourceApplicationSecurityGroupIds(val *[]*string)
	SourceApplicationSecurityGroupIdsInput() *[]*string
	SourcePortRange() *string
	SetSourcePortRange(val *string)
	SourcePortRangeInput() *string
	SourcePortRanges() *[]*string
	SetSourcePortRanges(val *[]*string)
	SourcePortRangesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkSecurityRuleTimeoutsOutputReference
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
	PutTimeouts(value *NetworkSecurityRuleTimeouts)
	ResetDescription()
	ResetDestinationAddressPrefix()
	ResetDestinationAddressPrefixes()
	ResetDestinationApplicationSecurityGroupIds()
	ResetDestinationPortRange()
	ResetDestinationPortRanges()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSourceAddressPrefix()
	ResetSourceAddressPrefixes()
	ResetSourceApplicationSecurityGroupIds()
	ResetSourcePortRange()
	ResetSourcePortRanges()
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

// The jsii proxy struct for NetworkSecurityRule
type jsiiProxy_NetworkSecurityRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkSecurityRule) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationAddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationAddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationApplicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationApplicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationPortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DestinationPortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) NetworkSecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) NetworkSecurityGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceAddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceAddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceApplicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourceApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceApplicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourcePortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourcePortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourcePortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) SourcePortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) Timeouts() NetworkSecurityRuleTimeoutsOutputReference {
	var returns NetworkSecurityRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule azurerm_network_security_rule} Resource.
func NewNetworkSecurityRule(scope constructs.Construct, id *string, config *NetworkSecurityRuleConfig) NetworkSecurityRule {
	_init_.Initialize()

	j := jsiiProxy_NetworkSecurityRule{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule azurerm_network_security_rule} Resource.
func NewNetworkSecurityRule_Override(n NetworkSecurityRule, scope constructs.Construct, id *string, config *NetworkSecurityRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetAccess(val *string) {
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetDestinationAddressPrefix(val *string) {
	_jsii_.Set(
		j,
		"destinationAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetDestinationAddressPrefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationAddressPrefixes",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetDestinationApplicationSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationApplicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetDestinationPortRange(val *string) {
	_jsii_.Set(
		j,
		"destinationPortRange",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetDestinationPortRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationPortRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetDirection(val *string) {
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetNetworkSecurityGroupName(val *string) {
	_jsii_.Set(
		j,
		"networkSecurityGroupName",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetSourceAddressPrefix(val *string) {
	_jsii_.Set(
		j,
		"sourceAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetSourceAddressPrefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceAddressPrefixes",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetSourceApplicationSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceApplicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetSourcePortRange(val *string) {
	_jsii_.Set(
		j,
		"sourcePortRange",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRule) SetSourcePortRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"sourcePortRanges",
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
func NetworkSecurityRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkSecurityRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) PutTimeouts(value *NetworkSecurityRuleTimeouts) {
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationAddressPrefix() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationAddressPrefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationAddressPrefixes() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationAddressPrefixes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetDestinationPortRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationPortRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourceAddressPrefix() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceAddressPrefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourceAddressPrefixes() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceAddressPrefixes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourceApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		n,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetSourcePortRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetSourcePortRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type NetworkSecurityRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#access NetworkSecurityRule#access}.
	Access *string `field:"required" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#direction NetworkSecurityRule#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#name NetworkSecurityRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#network_security_group_name NetworkSecurityRule#network_security_group_name}.
	NetworkSecurityGroupName *string `field:"required" json:"networkSecurityGroupName" yaml:"networkSecurityGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#priority NetworkSecurityRule#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#protocol NetworkSecurityRule#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#resource_group_name NetworkSecurityRule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#description NetworkSecurityRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#destination_address_prefix NetworkSecurityRule#destination_address_prefix}.
	DestinationAddressPrefix *string `field:"optional" json:"destinationAddressPrefix" yaml:"destinationAddressPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#destination_address_prefixes NetworkSecurityRule#destination_address_prefixes}.
	DestinationAddressPrefixes *[]*string `field:"optional" json:"destinationAddressPrefixes" yaml:"destinationAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#destination_application_security_group_ids NetworkSecurityRule#destination_application_security_group_ids}.
	DestinationApplicationSecurityGroupIds *[]*string `field:"optional" json:"destinationApplicationSecurityGroupIds" yaml:"destinationApplicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#destination_port_range NetworkSecurityRule#destination_port_range}.
	DestinationPortRange *string `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#destination_port_ranges NetworkSecurityRule#destination_port_ranges}.
	DestinationPortRanges *[]*string `field:"optional" json:"destinationPortRanges" yaml:"destinationPortRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#id NetworkSecurityRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#source_address_prefix NetworkSecurityRule#source_address_prefix}.
	SourceAddressPrefix *string `field:"optional" json:"sourceAddressPrefix" yaml:"sourceAddressPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#source_address_prefixes NetworkSecurityRule#source_address_prefixes}.
	SourceAddressPrefixes *[]*string `field:"optional" json:"sourceAddressPrefixes" yaml:"sourceAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#source_application_security_group_ids NetworkSecurityRule#source_application_security_group_ids}.
	SourceApplicationSecurityGroupIds *[]*string `field:"optional" json:"sourceApplicationSecurityGroupIds" yaml:"sourceApplicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#source_port_range NetworkSecurityRule#source_port_range}.
	SourcePortRange *string `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#source_port_ranges NetworkSecurityRule#source_port_ranges}.
	SourcePortRanges *[]*string `field:"optional" json:"sourcePortRanges" yaml:"sourcePortRanges"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#timeouts NetworkSecurityRule#timeouts}
	Timeouts *NetworkSecurityRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type NetworkSecurityRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#create NetworkSecurityRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#delete NetworkSecurityRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#read NetworkSecurityRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_security_rule#update NetworkSecurityRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type NetworkSecurityRuleTimeoutsOutputReference interface {
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

// The jsii proxy struct for NetworkSecurityRuleTimeoutsOutputReference
type jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewNetworkSecurityRuleTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkSecurityRuleTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRuleTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkSecurityRuleTimeoutsOutputReference_Override(n NetworkSecurityRuleTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkSecurityRule.NetworkSecurityRuleTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		n,
		"resetCreate",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		n,
		"resetDelete",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		n,
		"resetRead",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		n,
		"resetUpdate",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityRuleTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

