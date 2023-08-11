package paloaltonextgenerationfirewallvirtualhublocalrulestack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/paloaltonextgenerationfirewallvirtualhublocalrulestack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack azurerm_palo_alto_next_generation_firewall_virtual_hub_local_rulestack}.
type PaloAltoNextGenerationFirewallVirtualHubLocalRulestack interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationNat() PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDestinationNatList
	DestinationNatInput() interface{}
	DnsSettings() PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettingsOutputReference
	DnsSettingsInput() *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettings
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
	NetworkProfile() PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference
	NetworkProfileInput() *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile
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
	RulestackId() *string
	SetRulestackId(val *string)
	RulestackIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PaloAltoNextGenerationFirewallVirtualHubLocalRulestackTimeoutsOutputReference
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
	PutDestinationNat(value interface{})
	PutDnsSettings(value *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettings)
	PutNetworkProfile(value *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile)
	PutTimeouts(value *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackTimeouts)
	ResetDestinationNat()
	ResetDnsSettings()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for PaloAltoNextGenerationFirewallVirtualHubLocalRulestack
type jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) DestinationNat() PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDestinationNatList {
	var returns PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDestinationNatList
	_jsii_.Get(
		j,
		"destinationNat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) DestinationNatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationNatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) DnsSettings() PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettingsOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettingsOutputReference
	_jsii_.Get(
		j,
		"dnsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) DnsSettingsInput() *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettings {
	var returns *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettings
	_jsii_.Get(
		j,
		"dnsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) NetworkProfile() PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference
	_jsii_.Get(
		j,
		"networkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) NetworkProfileInput() *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile {
	var returns *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile
	_jsii_.Get(
		j,
		"networkProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) RulestackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulestackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) RulestackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulestackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) Timeouts() PaloAltoNextGenerationFirewallVirtualHubLocalRulestackTimeoutsOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualHubLocalRulestackTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack azurerm_palo_alto_next_generation_firewall_virtual_hub_local_rulestack} Resource.
func NewPaloAltoNextGenerationFirewallVirtualHubLocalRulestack(scope constructs.Construct, id *string, config *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackConfig) PaloAltoNextGenerationFirewallVirtualHubLocalRulestack {
	_init_.Initialize()

	if err := validateNewPaloAltoNextGenerationFirewallVirtualHubLocalRulestackParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubLocalRulestack.PaloAltoNextGenerationFirewallVirtualHubLocalRulestack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack azurerm_palo_alto_next_generation_firewall_virtual_hub_local_rulestack} Resource.
func NewPaloAltoNextGenerationFirewallVirtualHubLocalRulestack_Override(p PaloAltoNextGenerationFirewallVirtualHubLocalRulestack, scope constructs.Construct, id *string, config *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubLocalRulestack.PaloAltoNextGenerationFirewallVirtualHubLocalRulestack",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetRulestackId(val *string) {
	if err := j.validateSetRulestackIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rulestackId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack)SetTags(val *map[string]*string) {
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
func PaloAltoNextGenerationFirewallVirtualHubLocalRulestack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoNextGenerationFirewallVirtualHubLocalRulestack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubLocalRulestack.PaloAltoNextGenerationFirewallVirtualHubLocalRulestack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PaloAltoNextGenerationFirewallVirtualHubLocalRulestack_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoNextGenerationFirewallVirtualHubLocalRulestack_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubLocalRulestack.PaloAltoNextGenerationFirewallVirtualHubLocalRulestack",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PaloAltoNextGenerationFirewallVirtualHubLocalRulestack_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoNextGenerationFirewallVirtualHubLocalRulestack_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubLocalRulestack.PaloAltoNextGenerationFirewallVirtualHubLocalRulestack",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PaloAltoNextGenerationFirewallVirtualHubLocalRulestack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubLocalRulestack.PaloAltoNextGenerationFirewallVirtualHubLocalRulestack",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) PutDestinationNat(value interface{}) {
	if err := p.validatePutDestinationNatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDestinationNat",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) PutDnsSettings(value *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettings) {
	if err := p.validatePutDnsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDnsSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) PutNetworkProfile(value *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile) {
	if err := p.validatePutNetworkProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNetworkProfile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) PutTimeouts(value *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ResetDestinationNat() {
	_jsii_.InvokeVoid(
		p,
		"resetDestinationNat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ResetDnsSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetDnsSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

