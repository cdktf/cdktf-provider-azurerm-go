package firewallpolicyrulecollectiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/firewallpolicyrulecollectiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group azurerm_firewall_policy_rule_collection_group}.
type FirewallPolicyRuleCollectionGroup interface {
	cdktf.TerraformResource
	ApplicationRuleCollection() FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList
	ApplicationRuleCollectionInput() interface{}
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
	FirewallPolicyId() *string
	SetFirewallPolicyId(val *string)
	FirewallPolicyIdInput() *string
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
	NatRuleCollection() FirewallPolicyRuleCollectionGroupNatRuleCollectionList
	NatRuleCollectionInput() interface{}
	NetworkRuleCollection() FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList
	NetworkRuleCollectionInput() interface{}
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	Timeouts() FirewallPolicyRuleCollectionGroupTimeoutsOutputReference
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
	PutApplicationRuleCollection(value interface{})
	PutNatRuleCollection(value interface{})
	PutNetworkRuleCollection(value interface{})
	PutTimeouts(value *FirewallPolicyRuleCollectionGroupTimeouts)
	ResetApplicationRuleCollection()
	ResetId()
	ResetNatRuleCollection()
	ResetNetworkRuleCollection()
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

// The jsii proxy struct for FirewallPolicyRuleCollectionGroup
type jsiiProxy_FirewallPolicyRuleCollectionGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) ApplicationRuleCollection() FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList {
	var returns FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList
	_jsii_.Get(
		j,
		"applicationRuleCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) ApplicationRuleCollectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationRuleCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) FirewallPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) FirewallPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) NatRuleCollection() FirewallPolicyRuleCollectionGroupNatRuleCollectionList {
	var returns FirewallPolicyRuleCollectionGroupNatRuleCollectionList
	_jsii_.Get(
		j,
		"natRuleCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) NatRuleCollectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"natRuleCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) NetworkRuleCollection() FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList {
	var returns FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList
	_jsii_.Get(
		j,
		"networkRuleCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) NetworkRuleCollectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkRuleCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) Timeouts() FirewallPolicyRuleCollectionGroupTimeoutsOutputReference {
	var returns FirewallPolicyRuleCollectionGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group azurerm_firewall_policy_rule_collection_group} Resource.
func NewFirewallPolicyRuleCollectionGroup(scope constructs.Construct, id *string, config *FirewallPolicyRuleCollectionGroupConfig) FirewallPolicyRuleCollectionGroup {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroup{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group azurerm_firewall_policy_rule_collection_group} Resource.
func NewFirewallPolicyRuleCollectionGroup_Override(f FirewallPolicyRuleCollectionGroup, scope constructs.Construct, id *string, config *FirewallPolicyRuleCollectionGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroup",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetFirewallPolicyId(val *string) {
	_jsii_.Set(
		j,
		"firewallPolicyId",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroup) SetProvisioners(val *[]interface{}) {
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
func FirewallPolicyRuleCollectionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FirewallPolicyRuleCollectionGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) PutApplicationRuleCollection(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putApplicationRuleCollection",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) PutNatRuleCollection(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putNatRuleCollection",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) PutNetworkRuleCollection(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putNetworkRuleCollection",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) PutTimeouts(value *FirewallPolicyRuleCollectionGroupTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ResetApplicationRuleCollection() {
	_jsii_.InvokeVoid(
		f,
		"resetApplicationRuleCollection",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ResetNatRuleCollection() {
	_jsii_.InvokeVoid(
		f,
		"resetNatRuleCollection",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ResetNetworkRuleCollection() {
	_jsii_.InvokeVoid(
		f,
		"resetNetworkRuleCollection",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#action FirewallPolicyRuleCollectionGroup#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#priority FirewallPolicyRuleCollectionGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#rule FirewallPolicyRuleCollectionGroup#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList
type jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionList_Override(f FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) Get(index *float64) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference {
	var returns FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Rule() FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList
	RuleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutRule(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) Rule() FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList {
	var returns FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference_Override(f FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) PutRule(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putRule",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#description FirewallPolicyRuleCollectionGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_addresses FirewallPolicyRuleCollectionGroup#destination_addresses}.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_fqdns FirewallPolicyRuleCollectionGroup#destination_fqdns}.
	DestinationFqdns *[]*string `field:"optional" json:"destinationFqdns" yaml:"destinationFqdns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_fqdn_tags FirewallPolicyRuleCollectionGroup#destination_fqdn_tags}.
	DestinationFqdnTags *[]*string `field:"optional" json:"destinationFqdnTags" yaml:"destinationFqdnTags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_urls FirewallPolicyRuleCollectionGroup#destination_urls}.
	DestinationUrls *[]*string `field:"optional" json:"destinationUrls" yaml:"destinationUrls"`
	// protocols block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#protocols FirewallPolicyRuleCollectionGroup#protocols}
	Protocols interface{} `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#source_addresses FirewallPolicyRuleCollectionGroup#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#source_ip_groups FirewallPolicyRuleCollectionGroup#source_ip_groups}.
	SourceIpGroups *[]*string `field:"optional" json:"sourceIpGroups" yaml:"sourceIpGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#terminate_tls FirewallPolicyRuleCollectionGroup#terminate_tls}.
	TerminateTls interface{} `field:"optional" json:"terminateTls" yaml:"terminateTls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#web_categories FirewallPolicyRuleCollectionGroup#web_categories}.
	WebCategories *[]*string `field:"optional" json:"webCategories" yaml:"webCategories"`
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList
type jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList_Override(f FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) Get(index *float64) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference {
	var returns FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationAddresses() *[]*string
	SetDestinationAddresses(val *[]*string)
	DestinationAddressesInput() *[]*string
	DestinationFqdns() *[]*string
	SetDestinationFqdns(val *[]*string)
	DestinationFqdnsInput() *[]*string
	DestinationFqdnTags() *[]*string
	SetDestinationFqdnTags(val *[]*string)
	DestinationFqdnTagsInput() *[]*string
	DestinationUrls() *[]*string
	SetDestinationUrls(val *[]*string)
	DestinationUrlsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocols() FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList
	ProtocolsInput() interface{}
	SourceAddresses() *[]*string
	SetSourceAddresses(val *[]*string)
	SourceAddressesInput() *[]*string
	SourceIpGroups() *[]*string
	SetSourceIpGroups(val *[]*string)
	SourceIpGroupsInput() *[]*string
	TerminateTls() interface{}
	SetTerminateTls(val interface{})
	TerminateTlsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebCategories() *[]*string
	SetWebCategories(val *[]*string)
	WebCategoriesInput() *[]*string
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
	PutProtocols(value interface{})
	ResetDescription()
	ResetDestinationAddresses()
	ResetDestinationFqdns()
	ResetDestinationFqdnTags()
	ResetDestinationUrls()
	ResetProtocols()
	ResetSourceAddresses()
	ResetSourceIpGroups()
	ResetTerminateTls()
	ResetWebCategories()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationFqdnTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdnTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationFqdnTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdnTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) DestinationUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Protocols() FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList {
	var returns FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ProtocolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SourceIpGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SourceIpGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) TerminateTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) TerminateTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) WebCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) WebCategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webCategoriesInput",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference_Override(f FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetDestinationAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetDestinationFqdns(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationFqdns",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetDestinationFqdnTags(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationFqdnTags",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetDestinationUrls(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationUrls",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetSourceAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetSourceIpGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceIpGroups",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetTerminateTls(val interface{}) {
	_jsii_.Set(
		j,
		"terminateTls",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) SetWebCategories(val *[]*string) {
	_jsii_.Set(
		j,
		"webCategories",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) PutProtocols(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putProtocols",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDestinationFqdns() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationFqdns",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDestinationFqdnTags() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationFqdnTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetDestinationUrls() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationUrls",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		f,
		"resetProtocols",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetSourceIpGroups() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceIpGroups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetTerminateTls() {
	_jsii_.InvokeVoid(
		f,
		"resetTerminateTls",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ResetWebCategories() {
	_jsii_.InvokeVoid(
		f,
		"resetWebCategories",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocols struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#port FirewallPolicyRuleCollectionGroup#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#type FirewallPolicyRuleCollectionGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList
type jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList_Override(f FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) Get(index *float64) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference {
	var returns FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference_Override(f FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#firewall_policy_id FirewallPolicyRuleCollectionGroup#firewall_policy_id}.
	FirewallPolicyId *string `field:"required" json:"firewallPolicyId" yaml:"firewallPolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#priority FirewallPolicyRuleCollectionGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// application_rule_collection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#application_rule_collection FirewallPolicyRuleCollectionGroup#application_rule_collection}
	ApplicationRuleCollection interface{} `field:"optional" json:"applicationRuleCollection" yaml:"applicationRuleCollection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#id FirewallPolicyRuleCollectionGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// nat_rule_collection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#nat_rule_collection FirewallPolicyRuleCollectionGroup#nat_rule_collection}
	NatRuleCollection interface{} `field:"optional" json:"natRuleCollection" yaml:"natRuleCollection"`
	// network_rule_collection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#network_rule_collection FirewallPolicyRuleCollectionGroup#network_rule_collection}
	NetworkRuleCollection interface{} `field:"optional" json:"networkRuleCollection" yaml:"networkRuleCollection"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#timeouts FirewallPolicyRuleCollectionGroup#timeouts}
	Timeouts *FirewallPolicyRuleCollectionGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type FirewallPolicyRuleCollectionGroupNatRuleCollection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#action FirewallPolicyRuleCollectionGroup#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#priority FirewallPolicyRuleCollectionGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#rule FirewallPolicyRuleCollectionGroup#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

type FirewallPolicyRuleCollectionGroupNatRuleCollectionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNatRuleCollectionList
type jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirewallPolicyRuleCollectionGroupNatRuleCollectionList {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionList_Override(f FirewallPolicyRuleCollectionGroupNatRuleCollectionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) Get(index *float64) FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference {
	var returns FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Rule() FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList
	RuleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutRule(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) Rule() FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList {
	var returns FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference_Override(f FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) PutRule(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putRule",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupNatRuleCollectionRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#protocols FirewallPolicyRuleCollectionGroup#protocols}.
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#translated_port FirewallPolicyRuleCollectionGroup#translated_port}.
	TranslatedPort *float64 `field:"required" json:"translatedPort" yaml:"translatedPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_address FirewallPolicyRuleCollectionGroup#destination_address}.
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_ports FirewallPolicyRuleCollectionGroup#destination_ports}.
	DestinationPorts *[]*string `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#source_addresses FirewallPolicyRuleCollectionGroup#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#source_ip_groups FirewallPolicyRuleCollectionGroup#source_ip_groups}.
	SourceIpGroups *[]*string `field:"optional" json:"sourceIpGroups" yaml:"sourceIpGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#translated_address FirewallPolicyRuleCollectionGroup#translated_address}.
	TranslatedAddress *string `field:"optional" json:"translatedAddress" yaml:"translatedAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#translated_fqdn FirewallPolicyRuleCollectionGroup#translated_fqdn}.
	TranslatedFqdn *string `field:"optional" json:"translatedFqdn" yaml:"translatedFqdn"`
}

type FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList
type jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList_Override(f FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) Get(index *float64) FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference {
	var returns FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationAddress() *string
	SetDestinationAddress(val *string)
	DestinationAddressInput() *string
	DestinationPorts() *[]*string
	SetDestinationPorts(val *[]*string)
	DestinationPortsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	SourceAddresses() *[]*string
	SetSourceAddresses(val *[]*string)
	SourceAddressesInput() *[]*string
	SourceIpGroups() *[]*string
	SetSourceIpGroups(val *[]*string)
	SourceIpGroupsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TranslatedAddress() *string
	SetTranslatedAddress(val *string)
	TranslatedAddressInput() *string
	TranslatedFqdn() *string
	SetTranslatedFqdn(val *string)
	TranslatedFqdnInput() *string
	TranslatedPort() *float64
	SetTranslatedPort(val *float64)
	TranslatedPortInput() *float64
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
	ResetDestinationAddress()
	ResetDestinationPorts()
	ResetSourceAddresses()
	ResetSourceIpGroups()
	ResetTranslatedAddress()
	ResetTranslatedFqdn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DestinationPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DestinationPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SourceIpGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SourceIpGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"translatedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"translatedPortInput",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference_Override(f FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetDestinationAddress(val *string) {
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetDestinationPorts(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationPorts",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetSourceAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetSourceIpGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceIpGroups",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetTranslatedAddress(val *string) {
	_jsii_.Set(
		j,
		"translatedAddress",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetTranslatedFqdn(val *string) {
	_jsii_.Set(
		j,
		"translatedFqdn",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SetTranslatedPort(val *float64) {
	_jsii_.Set(
		j,
		"translatedPort",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetDestinationPorts() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationPorts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetSourceIpGroups() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceIpGroups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetTranslatedAddress() {
	_jsii_.InvokeVoid(
		f,
		"resetTranslatedAddress",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetTranslatedFqdn() {
	_jsii_.InvokeVoid(
		f,
		"resetTranslatedFqdn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupNetworkRuleCollection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#action FirewallPolicyRuleCollectionGroup#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#priority FirewallPolicyRuleCollectionGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#rule FirewallPolicyRuleCollectionGroup#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

type FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList
type jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionList_Override(f FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) Get(index *float64) FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference {
	var returns FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Rule() FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList
	RuleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutRule(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) Rule() FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList {
	var returns FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference_Override(f FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) PutRule(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putRule",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_ports FirewallPolicyRuleCollectionGroup#destination_ports}.
	DestinationPorts *[]*string `field:"required" json:"destinationPorts" yaml:"destinationPorts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#protocols FirewallPolicyRuleCollectionGroup#protocols}.
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_addresses FirewallPolicyRuleCollectionGroup#destination_addresses}.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_fqdns FirewallPolicyRuleCollectionGroup#destination_fqdns}.
	DestinationFqdns *[]*string `field:"optional" json:"destinationFqdns" yaml:"destinationFqdns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#destination_ip_groups FirewallPolicyRuleCollectionGroup#destination_ip_groups}.
	DestinationIpGroups *[]*string `field:"optional" json:"destinationIpGroups" yaml:"destinationIpGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#source_addresses FirewallPolicyRuleCollectionGroup#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#source_ip_groups FirewallPolicyRuleCollectionGroup#source_ip_groups}.
	SourceIpGroups *[]*string `field:"optional" json:"sourceIpGroups" yaml:"sourceIpGroups"`
}

type FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList
type jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList_Override(f FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) Get(index *float64) FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference {
	var returns FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationAddresses() *[]*string
	SetDestinationAddresses(val *[]*string)
	DestinationAddressesInput() *[]*string
	DestinationFqdns() *[]*string
	SetDestinationFqdns(val *[]*string)
	DestinationFqdnsInput() *[]*string
	DestinationIpGroups() *[]*string
	SetDestinationIpGroups(val *[]*string)
	DestinationIpGroupsInput() *[]*string
	DestinationPorts() *[]*string
	SetDestinationPorts(val *[]*string)
	DestinationPortsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	SourceAddresses() *[]*string
	SetSourceAddresses(val *[]*string)
	SourceAddressesInput() *[]*string
	SourceIpGroups() *[]*string
	SetSourceIpGroups(val *[]*string)
	SourceIpGroupsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetDestinationAddresses()
	ResetDestinationFqdns()
	ResetDestinationIpGroups()
	ResetSourceAddresses()
	ResetSourceIpGroups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) DestinationFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) DestinationFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) DestinationIpGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIpGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) DestinationIpGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIpGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) DestinationPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) DestinationPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SourceIpGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SourceIpGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference_Override(f FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetDestinationAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetDestinationFqdns(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationFqdns",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetDestinationIpGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationIpGroups",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetDestinationPorts(val *[]*string) {
	_jsii_.Set(
		j,
		"destinationPorts",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetSourceAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetSourceIpGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceIpGroups",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ResetDestinationFqdns() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationFqdns",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ResetDestinationIpGroups() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationIpGroups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ResetSourceIpGroups() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceIpGroups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FirewallPolicyRuleCollectionGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#create FirewallPolicyRuleCollectionGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#delete FirewallPolicyRuleCollectionGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#read FirewallPolicyRuleCollectionGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#update FirewallPolicyRuleCollectionGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type FirewallPolicyRuleCollectionGroupTimeoutsOutputReference interface {
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

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupTimeoutsOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FirewallPolicyRuleCollectionGroupTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupTimeoutsOutputReference_Override(f FirewallPolicyRuleCollectionGroupTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		f,
		"resetRead",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

