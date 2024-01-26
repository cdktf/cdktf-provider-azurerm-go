// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgatewayconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/virtualnetworkgatewayconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection}.
type VirtualNetworkGatewayConnection interface {
	cdktf.TerraformResource
	AuthorizationKey() *string
	SetAuthorizationKey(val *string)
	AuthorizationKeyInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionMode() *string
	SetConnectionMode(val *string)
	ConnectionModeInput() *string
	ConnectionProtocol() *string
	SetConnectionProtocol(val *string)
	ConnectionProtocolInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomBgpAddresses() VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference
	CustomBgpAddressesInput() *VirtualNetworkGatewayConnectionCustomBgpAddresses
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DpdTimeoutSeconds() *float64
	SetDpdTimeoutSeconds(val *float64)
	DpdTimeoutSecondsInput() *float64
	EgressNatRuleIds() *[]*string
	SetEgressNatRuleIds(val *[]*string)
	EgressNatRuleIdsInput() *[]*string
	EnableBgp() interface{}
	SetEnableBgp(val interface{})
	EnableBgpInput() interface{}
	ExpressRouteCircuitId() *string
	SetExpressRouteCircuitId(val *string)
	ExpressRouteCircuitIdInput() *string
	ExpressRouteGatewayBypass() interface{}
	SetExpressRouteGatewayBypass(val interface{})
	ExpressRouteGatewayBypassInput() interface{}
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
	IngressNatRuleIds() *[]*string
	SetIngressNatRuleIds(val *[]*string)
	IngressNatRuleIdsInput() *[]*string
	IpsecPolicy() VirtualNetworkGatewayConnectionIpsecPolicyOutputReference
	IpsecPolicyInput() *VirtualNetworkGatewayConnectionIpsecPolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAzureIpAddressEnabled() interface{}
	SetLocalAzureIpAddressEnabled(val interface{})
	LocalAzureIpAddressEnabledInput() interface{}
	LocalNetworkGatewayId() *string
	SetLocalNetworkGatewayId(val *string)
	LocalNetworkGatewayIdInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerVirtualNetworkGatewayId() *string
	SetPeerVirtualNetworkGatewayId(val *string)
	PeerVirtualNetworkGatewayIdInput() *string
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
	RoutingWeight() *float64
	SetRoutingWeight(val *float64)
	RoutingWeightInput() *float64
	SharedKey() *string
	SetSharedKey(val *string)
	SharedKeyInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualNetworkGatewayConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrafficSelectorPolicy() VirtualNetworkGatewayConnectionTrafficSelectorPolicyList
	TrafficSelectorPolicyInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UsePolicyBasedTrafficSelectors() interface{}
	SetUsePolicyBasedTrafficSelectors(val interface{})
	UsePolicyBasedTrafficSelectorsInput() interface{}
	VirtualNetworkGatewayId() *string
	SetVirtualNetworkGatewayId(val *string)
	VirtualNetworkGatewayIdInput() *string
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
	PutCustomBgpAddresses(value *VirtualNetworkGatewayConnectionCustomBgpAddresses)
	PutIpsecPolicy(value *VirtualNetworkGatewayConnectionIpsecPolicy)
	PutTimeouts(value *VirtualNetworkGatewayConnectionTimeouts)
	PutTrafficSelectorPolicy(value interface{})
	ResetAuthorizationKey()
	ResetConnectionMode()
	ResetConnectionProtocol()
	ResetCustomBgpAddresses()
	ResetDpdTimeoutSeconds()
	ResetEgressNatRuleIds()
	ResetEnableBgp()
	ResetExpressRouteCircuitId()
	ResetExpressRouteGatewayBypass()
	ResetId()
	ResetIngressNatRuleIds()
	ResetIpsecPolicy()
	ResetLocalAzureIpAddressEnabled()
	ResetLocalNetworkGatewayId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerVirtualNetworkGatewayId()
	ResetRoutingWeight()
	ResetSharedKey()
	ResetTags()
	ResetTimeouts()
	ResetTrafficSelectorPolicy()
	ResetUsePolicyBasedTrafficSelectors()
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

// The jsii proxy struct for VirtualNetworkGatewayConnection
type jsiiProxy_VirtualNetworkGatewayConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) AuthorizationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) AuthorizationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConnectionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConnectionProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConnectionProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) CustomBgpAddresses() VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference {
	var returns VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference
	_jsii_.Get(
		j,
		"customBgpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) CustomBgpAddressesInput() *VirtualNetworkGatewayConnectionCustomBgpAddresses {
	var returns *VirtualNetworkGatewayConnectionCustomBgpAddresses
	_jsii_.Get(
		j,
		"customBgpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) EgressNatRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) EgressNatRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) EnableBgp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBgp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) EnableBgpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBgpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ExpressRouteCircuitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ExpressRouteCircuitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ExpressRouteGatewayBypass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressRouteGatewayBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ExpressRouteGatewayBypassInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressRouteGatewayBypassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IngressNatRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressNatRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IngressNatRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressNatRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IpsecPolicy() VirtualNetworkGatewayConnectionIpsecPolicyOutputReference {
	var returns VirtualNetworkGatewayConnectionIpsecPolicyOutputReference
	_jsii_.Get(
		j,
		"ipsecPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IpsecPolicyInput() *VirtualNetworkGatewayConnectionIpsecPolicy {
	var returns *VirtualNetworkGatewayConnectionIpsecPolicy
	_jsii_.Get(
		j,
		"ipsecPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocalAzureIpAddressEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAzureIpAddressEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocalAzureIpAddressEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAzureIpAddressEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocalNetworkGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localNetworkGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocalNetworkGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localNetworkGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) PeerVirtualNetworkGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVirtualNetworkGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) PeerVirtualNetworkGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVirtualNetworkGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) RoutingWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"routingWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) RoutingWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"routingWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Timeouts() VirtualNetworkGatewayConnectionTimeoutsOutputReference {
	var returns VirtualNetworkGatewayConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TrafficSelectorPolicy() VirtualNetworkGatewayConnectionTrafficSelectorPolicyList {
	var returns VirtualNetworkGatewayConnectionTrafficSelectorPolicyList
	_jsii_.Get(
		j,
		"trafficSelectorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TrafficSelectorPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficSelectorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) UsePolicyBasedTrafficSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePolicyBasedTrafficSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) UsePolicyBasedTrafficSelectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePolicyBasedTrafficSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) VirtualNetworkGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) VirtualNetworkGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkGatewayIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection} Resource.
func NewVirtualNetworkGatewayConnection(scope constructs.Construct, id *string, config *VirtualNetworkGatewayConnectionConfig) VirtualNetworkGatewayConnection {
	_init_.Initialize()

	if err := validateNewVirtualNetworkGatewayConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualNetworkGatewayConnection{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection} Resource.
func NewVirtualNetworkGatewayConnection_Override(v VirtualNetworkGatewayConnection, scope constructs.Construct, id *string, config *VirtualNetworkGatewayConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetAuthorizationKey(val *string) {
	if err := j.validateSetAuthorizationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationKey",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetConnectionMode(val *string) {
	if err := j.validateSetConnectionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionMode",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetConnectionProtocol(val *string) {
	if err := j.validateSetConnectionProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProtocol",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetDpdTimeoutSeconds(val *float64) {
	if err := j.validateSetDpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetEgressNatRuleIds(val *[]*string) {
	if err := j.validateSetEgressNatRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetEnableBgp(val interface{}) {
	if err := j.validateSetEnableBgpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBgp",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetExpressRouteCircuitId(val *string) {
	if err := j.validateSetExpressRouteCircuitIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressRouteCircuitId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetExpressRouteGatewayBypass(val interface{}) {
	if err := j.validateSetExpressRouteGatewayBypassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressRouteGatewayBypass",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetIngressNatRuleIds(val *[]*string) {
	if err := j.validateSetIngressNatRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetLocalAzureIpAddressEnabled(val interface{}) {
	if err := j.validateSetLocalAzureIpAddressEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAzureIpAddressEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetLocalNetworkGatewayId(val *string) {
	if err := j.validateSetLocalNetworkGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localNetworkGatewayId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetPeerVirtualNetworkGatewayId(val *string) {
	if err := j.validateSetPeerVirtualNetworkGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVirtualNetworkGatewayId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetRoutingWeight(val *float64) {
	if err := j.validateSetRoutingWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingWeight",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetSharedKey(val *string) {
	if err := j.validateSetSharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedKey",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetUsePolicyBasedTrafficSelectors(val interface{}) {
	if err := j.validateSetUsePolicyBasedTrafficSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePolicyBasedTrafficSelectors",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetVirtualNetworkGatewayId(val *string) {
	if err := j.validateSetVirtualNetworkGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkGatewayId",
		val,
	)
}

// Generates CDKTF code for importing a VirtualNetworkGatewayConnection resource upon running "cdktf plan <stack-name>".
func VirtualNetworkGatewayConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualNetworkGatewayConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
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
func VirtualNetworkGatewayConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGatewayConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetworkGatewayConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGatewayConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetworkGatewayConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGatewayConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualNetworkGatewayConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutCustomBgpAddresses(value *VirtualNetworkGatewayConnectionCustomBgpAddresses) {
	if err := v.validatePutCustomBgpAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomBgpAddresses",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutIpsecPolicy(value *VirtualNetworkGatewayConnectionIpsecPolicy) {
	if err := v.validatePutIpsecPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpsecPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutTimeouts(value *VirtualNetworkGatewayConnectionTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutTrafficSelectorPolicy(value interface{}) {
	if err := v.validatePutTrafficSelectorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTrafficSelectorPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetAuthorizationKey() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthorizationKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetConnectionMode() {
	_jsii_.InvokeVoid(
		v,
		"resetConnectionMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetConnectionProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetConnectionProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetCustomBgpAddresses() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomBgpAddresses",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetDpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetDpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetEgressNatRuleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetEgressNatRuleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetEnableBgp() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableBgp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetExpressRouteCircuitId() {
	_jsii_.InvokeVoid(
		v,
		"resetExpressRouteCircuitId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetExpressRouteGatewayBypass() {
	_jsii_.InvokeVoid(
		v,
		"resetExpressRouteGatewayBypass",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetIngressNatRuleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetIngressNatRuleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetIpsecPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetIpsecPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetLocalAzureIpAddressEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalAzureIpAddressEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetLocalNetworkGatewayId() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalNetworkGatewayId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetPeerVirtualNetworkGatewayId() {
	_jsii_.InvokeVoid(
		v,
		"resetPeerVirtualNetworkGatewayId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetRoutingWeight() {
	_jsii_.InvokeVoid(
		v,
		"resetRoutingWeight",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetSharedKey() {
	_jsii_.InvokeVoid(
		v,
		"resetSharedKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetTrafficSelectorPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetTrafficSelectorPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetUsePolicyBasedTrafficSelectors() {
	_jsii_.InvokeVoid(
		v,
		"resetUsePolicyBasedTrafficSelectors",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

