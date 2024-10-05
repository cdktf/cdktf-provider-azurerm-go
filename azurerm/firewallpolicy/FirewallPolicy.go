// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/firewallpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/firewall_policy azurerm_firewall_policy}.
type FirewallPolicy interface {
	cdktf.TerraformResource
	AutoLearnPrivateRangesEnabled() interface{}
	SetAutoLearnPrivateRangesEnabled(val interface{})
	AutoLearnPrivateRangesEnabledInput() interface{}
	BasePolicyId() *string
	SetBasePolicyId(val *string)
	BasePolicyIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChildPolicies() *[]*string
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
	Dns() FirewallPolicyDnsOutputReference
	DnsInput() *FirewallPolicyDns
	ExplicitProxy() FirewallPolicyExplicitProxyOutputReference
	ExplicitProxyInput() *FirewallPolicyExplicitProxy
	Firewalls() *[]*string
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
	Identity() FirewallPolicyIdentityOutputReference
	IdentityInput() *FirewallPolicyIdentity
	IdInput() *string
	Insights() FirewallPolicyInsightsOutputReference
	InsightsInput() *FirewallPolicyInsights
	IntrusionDetection() FirewallPolicyIntrusionDetectionOutputReference
	IntrusionDetectionInput() *FirewallPolicyIntrusionDetection
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
	PrivateIpRanges() *[]*string
	SetPrivateIpRanges(val *[]*string)
	PrivateIpRangesInput() *[]*string
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
	RuleCollectionGroups() *[]*string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	SqlRedirectAllowed() interface{}
	SetSqlRedirectAllowed(val interface{})
	SqlRedirectAllowedInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThreatIntelligenceAllowlist() FirewallPolicyThreatIntelligenceAllowlistStructOutputReference
	ThreatIntelligenceAllowlistInput() *FirewallPolicyThreatIntelligenceAllowlistStruct
	ThreatIntelligenceMode() *string
	SetThreatIntelligenceMode(val *string)
	ThreatIntelligenceModeInput() *string
	Timeouts() FirewallPolicyTimeoutsOutputReference
	TimeoutsInput() interface{}
	TlsCertificate() FirewallPolicyTlsCertificateOutputReference
	TlsCertificateInput() *FirewallPolicyTlsCertificate
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDns(value *FirewallPolicyDns)
	PutExplicitProxy(value *FirewallPolicyExplicitProxy)
	PutIdentity(value *FirewallPolicyIdentity)
	PutInsights(value *FirewallPolicyInsights)
	PutIntrusionDetection(value *FirewallPolicyIntrusionDetection)
	PutThreatIntelligenceAllowlist(value *FirewallPolicyThreatIntelligenceAllowlistStruct)
	PutTimeouts(value *FirewallPolicyTimeouts)
	PutTlsCertificate(value *FirewallPolicyTlsCertificate)
	ResetAutoLearnPrivateRangesEnabled()
	ResetBasePolicyId()
	ResetDns()
	ResetExplicitProxy()
	ResetId()
	ResetIdentity()
	ResetInsights()
	ResetIntrusionDetection()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateIpRanges()
	ResetSku()
	ResetSqlRedirectAllowed()
	ResetTags()
	ResetThreatIntelligenceAllowlist()
	ResetThreatIntelligenceMode()
	ResetTimeouts()
	ResetTlsCertificate()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FirewallPolicy
type jsiiProxy_FirewallPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FirewallPolicy) AutoLearnPrivateRangesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoLearnPrivateRangesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) AutoLearnPrivateRangesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoLearnPrivateRangesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) BasePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) BasePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ChildPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Dns() FirewallPolicyDnsOutputReference {
	var returns FirewallPolicyDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) DnsInput() *FirewallPolicyDns {
	var returns *FirewallPolicyDns
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ExplicitProxy() FirewallPolicyExplicitProxyOutputReference {
	var returns FirewallPolicyExplicitProxyOutputReference
	_jsii_.Get(
		j,
		"explicitProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ExplicitProxyInput() *FirewallPolicyExplicitProxy {
	var returns *FirewallPolicyExplicitProxy
	_jsii_.Get(
		j,
		"explicitProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Firewalls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"firewalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Identity() FirewallPolicyIdentityOutputReference {
	var returns FirewallPolicyIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) IdentityInput() *FirewallPolicyIdentity {
	var returns *FirewallPolicyIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Insights() FirewallPolicyInsightsOutputReference {
	var returns FirewallPolicyInsightsOutputReference
	_jsii_.Get(
		j,
		"insights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) InsightsInput() *FirewallPolicyInsights {
	var returns *FirewallPolicyInsights
	_jsii_.Get(
		j,
		"insightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) IntrusionDetection() FirewallPolicyIntrusionDetectionOutputReference {
	var returns FirewallPolicyIntrusionDetectionOutputReference
	_jsii_.Get(
		j,
		"intrusionDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) IntrusionDetectionInput() *FirewallPolicyIntrusionDetection {
	var returns *FirewallPolicyIntrusionDetection
	_jsii_.Get(
		j,
		"intrusionDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) PrivateIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) PrivateIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) RuleCollectionGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ruleCollectionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) SqlRedirectAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlRedirectAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) SqlRedirectAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlRedirectAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ThreatIntelligenceAllowlist() FirewallPolicyThreatIntelligenceAllowlistStructOutputReference {
	var returns FirewallPolicyThreatIntelligenceAllowlistStructOutputReference
	_jsii_.Get(
		j,
		"threatIntelligenceAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ThreatIntelligenceAllowlistInput() *FirewallPolicyThreatIntelligenceAllowlistStruct {
	var returns *FirewallPolicyThreatIntelligenceAllowlistStruct
	_jsii_.Get(
		j,
		"threatIntelligenceAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ThreatIntelligenceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threatIntelligenceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) ThreatIntelligenceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threatIntelligenceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) Timeouts() FirewallPolicyTimeoutsOutputReference {
	var returns FirewallPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) TlsCertificate() FirewallPolicyTlsCertificateOutputReference {
	var returns FirewallPolicyTlsCertificateOutputReference
	_jsii_.Get(
		j,
		"tlsCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicy) TlsCertificateInput() *FirewallPolicyTlsCertificate {
	var returns *FirewallPolicyTlsCertificate
	_jsii_.Get(
		j,
		"tlsCertificateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/firewall_policy azurerm_firewall_policy} Resource.
func NewFirewallPolicy(scope constructs.Construct, id *string, config *FirewallPolicyConfig) FirewallPolicy {
	_init_.Initialize()

	if err := validateNewFirewallPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/firewall_policy azurerm_firewall_policy} Resource.
func NewFirewallPolicy_Override(f FirewallPolicy, scope constructs.Construct, id *string, config *FirewallPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicy",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetAutoLearnPrivateRangesEnabled(val interface{}) {
	if err := j.validateSetAutoLearnPrivateRangesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoLearnPrivateRangesEnabled",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetBasePolicyId(val *string) {
	if err := j.validateSetBasePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basePolicyId",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetPrivateIpRanges(val *[]*string) {
	if err := j.validateSetPrivateIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpRanges",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetSqlRedirectAllowed(val interface{}) {
	if err := j.validateSetSqlRedirectAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlRedirectAllowed",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicy)SetThreatIntelligenceMode(val *string) {
	if err := j.validateSetThreatIntelligenceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threatIntelligenceMode",
		val,
	)
}

// Generates CDKTF code for importing a FirewallPolicy resource upon running "cdktf plan <stack-name>".
func FirewallPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFirewallPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicy",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func FirewallPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFirewallPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FirewallPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFirewallPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FirewallPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFirewallPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FirewallPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FirewallPolicy) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FirewallPolicy) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FirewallPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FirewallPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FirewallPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FirewallPolicy) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FirewallPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FirewallPolicy) PutDns(value *FirewallPolicyDns) {
	if err := f.validatePutDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDns",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicy) PutExplicitProxy(value *FirewallPolicyExplicitProxy) {
	if err := f.validatePutExplicitProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putExplicitProxy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicy) PutIdentity(value *FirewallPolicyIdentity) {
	if err := f.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putIdentity",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicy) PutInsights(value *FirewallPolicyInsights) {
	if err := f.validatePutInsightsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putInsights",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicy) PutIntrusionDetection(value *FirewallPolicyIntrusionDetection) {
	if err := f.validatePutIntrusionDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putIntrusionDetection",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicy) PutThreatIntelligenceAllowlist(value *FirewallPolicyThreatIntelligenceAllowlistStruct) {
	if err := f.validatePutThreatIntelligenceAllowlistParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putThreatIntelligenceAllowlist",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicy) PutTimeouts(value *FirewallPolicyTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicy) PutTlsCertificate(value *FirewallPolicyTlsCertificate) {
	if err := f.validatePutTlsCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTlsCertificate",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetAutoLearnPrivateRangesEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetAutoLearnPrivateRangesEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetBasePolicyId() {
	_jsii_.InvokeVoid(
		f,
		"resetBasePolicyId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetDns() {
	_jsii_.InvokeVoid(
		f,
		"resetDns",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetExplicitProxy() {
	_jsii_.InvokeVoid(
		f,
		"resetExplicitProxy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetIdentity() {
	_jsii_.InvokeVoid(
		f,
		"resetIdentity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetInsights() {
	_jsii_.InvokeVoid(
		f,
		"resetInsights",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetIntrusionDetection() {
	_jsii_.InvokeVoid(
		f,
		"resetIntrusionDetection",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetPrivateIpRanges() {
	_jsii_.InvokeVoid(
		f,
		"resetPrivateIpRanges",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetSku() {
	_jsii_.InvokeVoid(
		f,
		"resetSku",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetSqlRedirectAllowed() {
	_jsii_.InvokeVoid(
		f,
		"resetSqlRedirectAllowed",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetThreatIntelligenceAllowlist() {
	_jsii_.InvokeVoid(
		f,
		"resetThreatIntelligenceAllowlist",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetThreatIntelligenceMode() {
	_jsii_.InvokeVoid(
		f,
		"resetThreatIntelligenceMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) ResetTlsCertificate() {
	_jsii_.InvokeVoid(
		f,
		"resetTlsCertificate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

