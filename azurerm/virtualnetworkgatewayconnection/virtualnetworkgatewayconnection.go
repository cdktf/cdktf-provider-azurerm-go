package virtualnetworkgatewayconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/virtualnetworkgatewayconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection} Resource.
func NewVirtualNetworkGatewayConnection(scope constructs.Construct, id *string, config *VirtualNetworkGatewayConnectionConfig) VirtualNetworkGatewayConnection {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayConnection{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection} Resource.
func NewVirtualNetworkGatewayConnection_Override(v VirtualNetworkGatewayConnection, scope constructs.Construct, id *string, config *VirtualNetworkGatewayConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetAuthorizationKey(val *string) {
	_jsii_.Set(
		j,
		"authorizationKey",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetConnectionMode(val *string) {
	_jsii_.Set(
		j,
		"connectionMode",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetConnectionProtocol(val *string) {
	_jsii_.Set(
		j,
		"connectionProtocol",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetDpdTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"dpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetEgressNatRuleIds(val *[]*string) {
	_jsii_.Set(
		j,
		"egressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetEnableBgp(val interface{}) {
	_jsii_.Set(
		j,
		"enableBgp",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetExpressRouteCircuitId(val *string) {
	_jsii_.Set(
		j,
		"expressRouteCircuitId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetExpressRouteGatewayBypass(val interface{}) {
	_jsii_.Set(
		j,
		"expressRouteGatewayBypass",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetIngressNatRuleIds(val *[]*string) {
	_jsii_.Set(
		j,
		"ingressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetLocalAzureIpAddressEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"localAzureIpAddressEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetLocalNetworkGatewayId(val *string) {
	_jsii_.Set(
		j,
		"localNetworkGatewayId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetPeerVirtualNetworkGatewayId(val *string) {
	_jsii_.Set(
		j,
		"peerVirtualNetworkGatewayId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetRoutingWeight(val *float64) {
	_jsii_.Set(
		j,
		"routingWeight",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetSharedKey(val *string) {
	_jsii_.Set(
		j,
		"sharedKey",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetUsePolicyBasedTrafficSelectors(val interface{}) {
	_jsii_.Set(
		j,
		"usePolicyBasedTrafficSelectors",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SetVirtualNetworkGatewayId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkGatewayId",
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
func VirtualNetworkGatewayConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"isConstruct",
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutCustomBgpAddresses(value *VirtualNetworkGatewayConnectionCustomBgpAddresses) {
	_jsii_.InvokeVoid(
		v,
		"putCustomBgpAddresses",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutIpsecPolicy(value *VirtualNetworkGatewayConnectionIpsecPolicy) {
	_jsii_.InvokeVoid(
		v,
		"putIpsecPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutTimeouts(value *VirtualNetworkGatewayConnectionTimeouts) {
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutTrafficSelectorPolicy(value interface{}) {
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

type VirtualNetworkGatewayConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#location VirtualNetworkGatewayConnection#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#name VirtualNetworkGatewayConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#resource_group_name VirtualNetworkGatewayConnection#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#type VirtualNetworkGatewayConnection#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#virtual_network_gateway_id VirtualNetworkGatewayConnection#virtual_network_gateway_id}.
	VirtualNetworkGatewayId *string `field:"required" json:"virtualNetworkGatewayId" yaml:"virtualNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#authorization_key VirtualNetworkGatewayConnection#authorization_key}.
	AuthorizationKey *string `field:"optional" json:"authorizationKey" yaml:"authorizationKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#connection_mode VirtualNetworkGatewayConnection#connection_mode}.
	ConnectionMode *string `field:"optional" json:"connectionMode" yaml:"connectionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#connection_protocol VirtualNetworkGatewayConnection#connection_protocol}.
	ConnectionProtocol *string `field:"optional" json:"connectionProtocol" yaml:"connectionProtocol"`
	// custom_bgp_addresses block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#custom_bgp_addresses VirtualNetworkGatewayConnection#custom_bgp_addresses}
	CustomBgpAddresses *VirtualNetworkGatewayConnectionCustomBgpAddresses `field:"optional" json:"customBgpAddresses" yaml:"customBgpAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#dpd_timeout_seconds VirtualNetworkGatewayConnection#dpd_timeout_seconds}.
	DpdTimeoutSeconds *float64 `field:"optional" json:"dpdTimeoutSeconds" yaml:"dpdTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#egress_nat_rule_ids VirtualNetworkGatewayConnection#egress_nat_rule_ids}.
	EgressNatRuleIds *[]*string `field:"optional" json:"egressNatRuleIds" yaml:"egressNatRuleIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#enable_bgp VirtualNetworkGatewayConnection#enable_bgp}.
	EnableBgp interface{} `field:"optional" json:"enableBgp" yaml:"enableBgp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#express_route_circuit_id VirtualNetworkGatewayConnection#express_route_circuit_id}.
	ExpressRouteCircuitId *string `field:"optional" json:"expressRouteCircuitId" yaml:"expressRouteCircuitId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#express_route_gateway_bypass VirtualNetworkGatewayConnection#express_route_gateway_bypass}.
	ExpressRouteGatewayBypass interface{} `field:"optional" json:"expressRouteGatewayBypass" yaml:"expressRouteGatewayBypass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#id VirtualNetworkGatewayConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#ingress_nat_rule_ids VirtualNetworkGatewayConnection#ingress_nat_rule_ids}.
	IngressNatRuleIds *[]*string `field:"optional" json:"ingressNatRuleIds" yaml:"ingressNatRuleIds"`
	// ipsec_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#ipsec_policy VirtualNetworkGatewayConnection#ipsec_policy}
	IpsecPolicy *VirtualNetworkGatewayConnectionIpsecPolicy `field:"optional" json:"ipsecPolicy" yaml:"ipsecPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#local_azure_ip_address_enabled VirtualNetworkGatewayConnection#local_azure_ip_address_enabled}.
	LocalAzureIpAddressEnabled interface{} `field:"optional" json:"localAzureIpAddressEnabled" yaml:"localAzureIpAddressEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#local_network_gateway_id VirtualNetworkGatewayConnection#local_network_gateway_id}.
	LocalNetworkGatewayId *string `field:"optional" json:"localNetworkGatewayId" yaml:"localNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#peer_virtual_network_gateway_id VirtualNetworkGatewayConnection#peer_virtual_network_gateway_id}.
	PeerVirtualNetworkGatewayId *string `field:"optional" json:"peerVirtualNetworkGatewayId" yaml:"peerVirtualNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#routing_weight VirtualNetworkGatewayConnection#routing_weight}.
	RoutingWeight *float64 `field:"optional" json:"routingWeight" yaml:"routingWeight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#shared_key VirtualNetworkGatewayConnection#shared_key}.
	SharedKey *string `field:"optional" json:"sharedKey" yaml:"sharedKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#tags VirtualNetworkGatewayConnection#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#timeouts VirtualNetworkGatewayConnection#timeouts}
	Timeouts *VirtualNetworkGatewayConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_selector_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#traffic_selector_policy VirtualNetworkGatewayConnection#traffic_selector_policy}
	TrafficSelectorPolicy interface{} `field:"optional" json:"trafficSelectorPolicy" yaml:"trafficSelectorPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#use_policy_based_traffic_selectors VirtualNetworkGatewayConnection#use_policy_based_traffic_selectors}.
	UsePolicyBasedTrafficSelectors interface{} `field:"optional" json:"usePolicyBasedTrafficSelectors" yaml:"usePolicyBasedTrafficSelectors"`
}

type VirtualNetworkGatewayConnectionCustomBgpAddresses struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#primary VirtualNetworkGatewayConnection#primary}.
	Primary *string `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#secondary VirtualNetworkGatewayConnection#secondary}.
	Secondary *string `field:"required" json:"secondary" yaml:"secondary"`
}

type VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference interface {
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
	InternalValue() *VirtualNetworkGatewayConnectionCustomBgpAddresses
	SetInternalValue(val *VirtualNetworkGatewayConnectionCustomBgpAddresses)
	Primary() *string
	SetPrimary(val *string)
	PrimaryInput() *string
	Secondary() *string
	SetSecondary(val *string)
	SecondaryInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference
type jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) InternalValue() *VirtualNetworkGatewayConnectionCustomBgpAddresses {
	var returns *VirtualNetworkGatewayConnectionCustomBgpAddresses
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) Primary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) PrimaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) Secondary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) SecondaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference_Override(v VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) SetInternalValue(val *VirtualNetworkGatewayConnectionCustomBgpAddresses) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) SetPrimary(val *string) {
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) SetSecondary(val *string) {
	_jsii_.Set(
		j,
		"secondary",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayConnectionIpsecPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#dh_group VirtualNetworkGatewayConnection#dh_group}.
	DhGroup *string `field:"required" json:"dhGroup" yaml:"dhGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#ike_encryption VirtualNetworkGatewayConnection#ike_encryption}.
	IkeEncryption *string `field:"required" json:"ikeEncryption" yaml:"ikeEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#ike_integrity VirtualNetworkGatewayConnection#ike_integrity}.
	IkeIntegrity *string `field:"required" json:"ikeIntegrity" yaml:"ikeIntegrity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#ipsec_encryption VirtualNetworkGatewayConnection#ipsec_encryption}.
	IpsecEncryption *string `field:"required" json:"ipsecEncryption" yaml:"ipsecEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#ipsec_integrity VirtualNetworkGatewayConnection#ipsec_integrity}.
	IpsecIntegrity *string `field:"required" json:"ipsecIntegrity" yaml:"ipsecIntegrity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#pfs_group VirtualNetworkGatewayConnection#pfs_group}.
	PfsGroup *string `field:"required" json:"pfsGroup" yaml:"pfsGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#sa_datasize VirtualNetworkGatewayConnection#sa_datasize}.
	SaDatasize *float64 `field:"optional" json:"saDatasize" yaml:"saDatasize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#sa_lifetime VirtualNetworkGatewayConnection#sa_lifetime}.
	SaLifetime *float64 `field:"optional" json:"saLifetime" yaml:"saLifetime"`
}

type VirtualNetworkGatewayConnectionIpsecPolicyOutputReference interface {
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
	DhGroup() *string
	SetDhGroup(val *string)
	DhGroupInput() *string
	// Experimental.
	Fqn() *string
	IkeEncryption() *string
	SetIkeEncryption(val *string)
	IkeEncryptionInput() *string
	IkeIntegrity() *string
	SetIkeIntegrity(val *string)
	IkeIntegrityInput() *string
	InternalValue() *VirtualNetworkGatewayConnectionIpsecPolicy
	SetInternalValue(val *VirtualNetworkGatewayConnectionIpsecPolicy)
	IpsecEncryption() *string
	SetIpsecEncryption(val *string)
	IpsecEncryptionInput() *string
	IpsecIntegrity() *string
	SetIpsecIntegrity(val *string)
	IpsecIntegrityInput() *string
	PfsGroup() *string
	SetPfsGroup(val *string)
	PfsGroupInput() *string
	SaDatasize() *float64
	SetSaDatasize(val *float64)
	SaDatasizeInput() *float64
	SaLifetime() *float64
	SetSaLifetime(val *float64)
	SaLifetimeInput() *float64
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
	ResetSaDatasize()
	ResetSaLifetime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayConnectionIpsecPolicyOutputReference
type jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) DhGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) DhGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) IkeEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) IkeEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) IkeIntegrity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) IkeIntegrityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) InternalValue() *VirtualNetworkGatewayConnectionIpsecPolicy {
	var returns *VirtualNetworkGatewayConnectionIpsecPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) IpsecEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) IpsecEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) IpsecIntegrity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) IpsecIntegrityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) PfsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) PfsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SaDatasize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDatasize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SaDatasizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDatasizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SaLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SaLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayConnectionIpsecPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayConnectionIpsecPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayConnectionIpsecPolicyOutputReference_Override(v VirtualNetworkGatewayConnectionIpsecPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetDhGroup(val *string) {
	_jsii_.Set(
		j,
		"dhGroup",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetIkeEncryption(val *string) {
	_jsii_.Set(
		j,
		"ikeEncryption",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetIkeIntegrity(val *string) {
	_jsii_.Set(
		j,
		"ikeIntegrity",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetInternalValue(val *VirtualNetworkGatewayConnectionIpsecPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetIpsecEncryption(val *string) {
	_jsii_.Set(
		j,
		"ipsecEncryption",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetIpsecIntegrity(val *string) {
	_jsii_.Set(
		j,
		"ipsecIntegrity",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetPfsGroup(val *string) {
	_jsii_.Set(
		j,
		"pfsGroup",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetSaDatasize(val *float64) {
	_jsii_.Set(
		j,
		"saDatasize",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetSaLifetime(val *float64) {
	_jsii_.Set(
		j,
		"saLifetime",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) ResetSaDatasize() {
	_jsii_.InvokeVoid(
		v,
		"resetSaDatasize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) ResetSaLifetime() {
	_jsii_.InvokeVoid(
		v,
		"resetSaLifetime",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#create VirtualNetworkGatewayConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#delete VirtualNetworkGatewayConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#read VirtualNetworkGatewayConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#update VirtualNetworkGatewayConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type VirtualNetworkGatewayConnectionTimeoutsOutputReference interface {
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

// The jsii proxy struct for VirtualNetworkGatewayConnectionTimeoutsOutputReference
type jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayConnectionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayConnectionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayConnectionTimeoutsOutputReference_Override(v VirtualNetworkGatewayConnectionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		v,
		"resetCreate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		v,
		"resetDelete",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		v,
		"resetRead",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		v,
		"resetUpdate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayConnectionTrafficSelectorPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#local_address_cidrs VirtualNetworkGatewayConnection#local_address_cidrs}.
	LocalAddressCidrs *[]*string `field:"required" json:"localAddressCidrs" yaml:"localAddressCidrs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#remote_address_cidrs VirtualNetworkGatewayConnection#remote_address_cidrs}.
	RemoteAddressCidrs *[]*string `field:"required" json:"remoteAddressCidrs" yaml:"remoteAddressCidrs"`
}

type VirtualNetworkGatewayConnectionTrafficSelectorPolicyList interface {
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
	Get(index *float64) VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayConnectionTrafficSelectorPolicyList
type jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayConnectionTrafficSelectorPolicyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualNetworkGatewayConnectionTrafficSelectorPolicyList {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionTrafficSelectorPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayConnectionTrafficSelectorPolicyList_Override(v VirtualNetworkGatewayConnectionTrafficSelectorPolicyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionTrafficSelectorPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) Get(index *float64) VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference {
	var returns VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference interface {
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
	LocalAddressCidrs() *[]*string
	SetLocalAddressCidrs(val *[]*string)
	LocalAddressCidrsInput() *[]*string
	RemoteAddressCidrs() *[]*string
	SetRemoteAddressCidrs(val *[]*string)
	RemoteAddressCidrsInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference
type jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) LocalAddressCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localAddressCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) LocalAddressCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localAddressCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) RemoteAddressCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteAddressCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) RemoteAddressCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteAddressCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference_Override(v VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) SetLocalAddressCidrs(val *[]*string) {
	_jsii_.Set(
		j,
		"localAddressCidrs",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) SetRemoteAddressCidrs(val *[]*string) {
	_jsii_.Set(
		j,
		"remoteAddressCidrs",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnectionTrafficSelectorPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

