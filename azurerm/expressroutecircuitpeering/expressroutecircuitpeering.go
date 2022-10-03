package expressroutecircuitpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/expressroutecircuitpeering/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering azurerm_express_route_circuit_peering}.
type ExpressRouteCircuitPeering interface {
	cdktf.TerraformResource
	AzureAsn() *float64
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
	ExpressRouteCircuitName() *string
	SetExpressRouteCircuitName(val *string)
	ExpressRouteCircuitNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayManagerEtag() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Ipv4Enabled() interface{}
	SetIpv4Enabled(val interface{})
	Ipv4EnabledInput() interface{}
	Ipv6() ExpressRouteCircuitPeeringIpv6OutputReference
	Ipv6Input() *ExpressRouteCircuitPeeringIpv6
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MicrosoftPeeringConfig() ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference
	MicrosoftPeeringConfigInput() *ExpressRouteCircuitPeeringMicrosoftPeeringConfig
	// The tree node.
	Node() constructs.Node
	PeerAsn() *float64
	SetPeerAsn(val *float64)
	PeerAsnInput() *float64
	PeeringType() *string
	SetPeeringType(val *string)
	PeeringTypeInput() *string
	PrimaryAzurePort() *string
	PrimaryPeerAddressPrefix() *string
	SetPrimaryPeerAddressPrefix(val *string)
	PrimaryPeerAddressPrefixInput() *string
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
	RouteFilterId() *string
	SetRouteFilterId(val *string)
	RouteFilterIdInput() *string
	SecondaryAzurePort() *string
	SecondaryPeerAddressPrefix() *string
	SetSecondaryPeerAddressPrefix(val *string)
	SecondaryPeerAddressPrefixInput() *string
	SharedKey() *string
	SetSharedKey(val *string)
	SharedKeyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ExpressRouteCircuitPeeringTimeoutsOutputReference
	TimeoutsInput() interface{}
	VlanId() *float64
	SetVlanId(val *float64)
	VlanIdInput() *float64
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
	PutIpv6(value *ExpressRouteCircuitPeeringIpv6)
	PutMicrosoftPeeringConfig(value *ExpressRouteCircuitPeeringMicrosoftPeeringConfig)
	PutTimeouts(value *ExpressRouteCircuitPeeringTimeouts)
	ResetId()
	ResetIpv4Enabled()
	ResetIpv6()
	ResetMicrosoftPeeringConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerAsn()
	ResetPrimaryPeerAddressPrefix()
	ResetRouteFilterId()
	ResetSecondaryPeerAddressPrefix()
	ResetSharedKey()
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

// The jsii proxy struct for ExpressRouteCircuitPeering
type jsiiProxy_ExpressRouteCircuitPeering struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) AzureAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"azureAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) ExpressRouteCircuitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) ExpressRouteCircuitNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) GatewayManagerEtag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayManagerEtag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Ipv4Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Ipv4EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Ipv6() ExpressRouteCircuitPeeringIpv6OutputReference {
	var returns ExpressRouteCircuitPeeringIpv6OutputReference
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Ipv6Input() *ExpressRouteCircuitPeeringIpv6 {
	var returns *ExpressRouteCircuitPeeringIpv6
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference {
	var returns ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference
	_jsii_.Get(
		j,
		"microsoftPeeringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) MicrosoftPeeringConfigInput() *ExpressRouteCircuitPeeringMicrosoftPeeringConfig {
	var returns *ExpressRouteCircuitPeeringMicrosoftPeeringConfig
	_jsii_.Get(
		j,
		"microsoftPeeringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) PeerAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) PeerAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) PeeringType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) PeeringTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) PrimaryAzurePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAzurePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) PrimaryPeerAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPeerAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) PrimaryPeerAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPeerAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) RouteFilterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFilterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) RouteFilterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFilterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SecondaryAzurePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAzurePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SecondaryPeerAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPeerAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SecondaryPeerAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPeerAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) Timeouts() ExpressRouteCircuitPeeringTimeoutsOutputReference {
	var returns ExpressRouteCircuitPeeringTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) VlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) VlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering azurerm_express_route_circuit_peering} Resource.
func NewExpressRouteCircuitPeering(scope constructs.Construct, id *string, config *ExpressRouteCircuitPeeringConfig) ExpressRouteCircuitPeering {
	_init_.Initialize()

	j := jsiiProxy_ExpressRouteCircuitPeering{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering azurerm_express_route_circuit_peering} Resource.
func NewExpressRouteCircuitPeering_Override(e ExpressRouteCircuitPeering, scope constructs.Construct, id *string, config *ExpressRouteCircuitPeeringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetExpressRouteCircuitName(val *string) {
	_jsii_.Set(
		j,
		"expressRouteCircuitName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetIpv4Enabled(val interface{}) {
	_jsii_.Set(
		j,
		"ipv4Enabled",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetPeerAsn(val *float64) {
	_jsii_.Set(
		j,
		"peerAsn",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetPeeringType(val *string) {
	_jsii_.Set(
		j,
		"peeringType",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetPrimaryPeerAddressPrefix(val *string) {
	_jsii_.Set(
		j,
		"primaryPeerAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetRouteFilterId(val *string) {
	_jsii_.Set(
		j,
		"routeFilterId",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetSecondaryPeerAddressPrefix(val *string) {
	_jsii_.Set(
		j,
		"secondaryPeerAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetSharedKey(val *string) {
	_jsii_.Set(
		j,
		"sharedKey",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering) SetVlanId(val *float64) {
	_jsii_.Set(
		j,
		"vlanId",
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
func ExpressRouteCircuitPeering_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExpressRouteCircuitPeering_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) PutIpv6(value *ExpressRouteCircuitPeeringIpv6) {
	_jsii_.InvokeVoid(
		e,
		"putIpv6",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) PutMicrosoftPeeringConfig(value *ExpressRouteCircuitPeeringMicrosoftPeeringConfig) {
	_jsii_.InvokeVoid(
		e,
		"putMicrosoftPeeringConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) PutTimeouts(value *ExpressRouteCircuitPeeringTimeouts) {
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetIpv4Enabled() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv4Enabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetIpv6() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetMicrosoftPeeringConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetMicrosoftPeeringConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetPeerAsn() {
	_jsii_.InvokeVoid(
		e,
		"resetPeerAsn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetPrimaryPeerAddressPrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetPrimaryPeerAddressPrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetRouteFilterId() {
	_jsii_.InvokeVoid(
		e,
		"resetRouteFilterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetSecondaryPeerAddressPrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetSecondaryPeerAddressPrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetSharedKey() {
	_jsii_.InvokeVoid(
		e,
		"resetSharedKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ExpressRouteCircuitPeeringConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#express_route_circuit_name ExpressRouteCircuitPeering#express_route_circuit_name}.
	ExpressRouteCircuitName *string `field:"required" json:"expressRouteCircuitName" yaml:"expressRouteCircuitName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#peering_type ExpressRouteCircuitPeering#peering_type}.
	PeeringType *string `field:"required" json:"peeringType" yaml:"peeringType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#resource_group_name ExpressRouteCircuitPeering#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#vlan_id ExpressRouteCircuitPeering#vlan_id}.
	VlanId *float64 `field:"required" json:"vlanId" yaml:"vlanId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#id ExpressRouteCircuitPeering#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#ipv4_enabled ExpressRouteCircuitPeering#ipv4_enabled}.
	Ipv4Enabled interface{} `field:"optional" json:"ipv4Enabled" yaml:"ipv4Enabled"`
	// ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#ipv6 ExpressRouteCircuitPeering#ipv6}
	Ipv6 *ExpressRouteCircuitPeeringIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
	// microsoft_peering_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#microsoft_peering_config ExpressRouteCircuitPeering#microsoft_peering_config}
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringMicrosoftPeeringConfig `field:"optional" json:"microsoftPeeringConfig" yaml:"microsoftPeeringConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#peer_asn ExpressRouteCircuitPeering#peer_asn}.
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#primary_peer_address_prefix ExpressRouteCircuitPeering#primary_peer_address_prefix}.
	PrimaryPeerAddressPrefix *string `field:"optional" json:"primaryPeerAddressPrefix" yaml:"primaryPeerAddressPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#route_filter_id ExpressRouteCircuitPeering#route_filter_id}.
	RouteFilterId *string `field:"optional" json:"routeFilterId" yaml:"routeFilterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#secondary_peer_address_prefix ExpressRouteCircuitPeering#secondary_peer_address_prefix}.
	SecondaryPeerAddressPrefix *string `field:"optional" json:"secondaryPeerAddressPrefix" yaml:"secondaryPeerAddressPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#shared_key ExpressRouteCircuitPeering#shared_key}.
	SharedKey *string `field:"optional" json:"sharedKey" yaml:"sharedKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#timeouts ExpressRouteCircuitPeering#timeouts}
	Timeouts *ExpressRouteCircuitPeeringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type ExpressRouteCircuitPeeringIpv6 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#primary_peer_address_prefix ExpressRouteCircuitPeering#primary_peer_address_prefix}.
	PrimaryPeerAddressPrefix *string `field:"required" json:"primaryPeerAddressPrefix" yaml:"primaryPeerAddressPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#secondary_peer_address_prefix ExpressRouteCircuitPeering#secondary_peer_address_prefix}.
	SecondaryPeerAddressPrefix *string `field:"required" json:"secondaryPeerAddressPrefix" yaml:"secondaryPeerAddressPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#enabled ExpressRouteCircuitPeering#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// microsoft_peering block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#microsoft_peering ExpressRouteCircuitPeering#microsoft_peering}
	MicrosoftPeering *ExpressRouteCircuitPeeringIpv6MicrosoftPeering `field:"optional" json:"microsoftPeering" yaml:"microsoftPeering"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#route_filter_id ExpressRouteCircuitPeering#route_filter_id}.
	RouteFilterId *string `field:"optional" json:"routeFilterId" yaml:"routeFilterId"`
}

type ExpressRouteCircuitPeeringIpv6MicrosoftPeering struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#advertised_public_prefixes ExpressRouteCircuitPeering#advertised_public_prefixes}.
	AdvertisedPublicPrefixes *[]*string `field:"optional" json:"advertisedPublicPrefixes" yaml:"advertisedPublicPrefixes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#customer_asn ExpressRouteCircuitPeering#customer_asn}.
	CustomerAsn *float64 `field:"optional" json:"customerAsn" yaml:"customerAsn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#routing_registry_name ExpressRouteCircuitPeering#routing_registry_name}.
	RoutingRegistryName *string `field:"optional" json:"routingRegistryName" yaml:"routingRegistryName"`
}

type ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference interface {
	cdktf.ComplexObject
	AdvertisedPublicPrefixes() *[]*string
	SetAdvertisedPublicPrefixes(val *[]*string)
	AdvertisedPublicPrefixesInput() *[]*string
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
	CustomerAsn() *float64
	SetCustomerAsn(val *float64)
	CustomerAsnInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ExpressRouteCircuitPeeringIpv6MicrosoftPeering
	SetInternalValue(val *ExpressRouteCircuitPeeringIpv6MicrosoftPeering)
	RoutingRegistryName() *string
	SetRoutingRegistryName(val *string)
	RoutingRegistryNameInput() *string
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
	ResetAdvertisedPublicPrefixes()
	ResetCustomerAsn()
	ResetRoutingRegistryName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference
type jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) AdvertisedPublicPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedPublicPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) AdvertisedPublicPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedPublicPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) CustomerAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) CustomerAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) InternalValue() *ExpressRouteCircuitPeeringIpv6MicrosoftPeering {
	var returns *ExpressRouteCircuitPeeringIpv6MicrosoftPeering
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) RoutingRegistryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRegistryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) RoutingRegistryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRegistryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference_Override(e ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) SetAdvertisedPublicPrefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"advertisedPublicPrefixes",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) SetCustomerAsn(val *float64) {
	_jsii_.Set(
		j,
		"customerAsn",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) SetInternalValue(val *ExpressRouteCircuitPeeringIpv6MicrosoftPeering) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) SetRoutingRegistryName(val *string) {
	_jsii_.Set(
		j,
		"routingRegistryName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) ResetAdvertisedPublicPrefixes() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvertisedPublicPrefixes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) ResetCustomerAsn() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomerAsn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) ResetRoutingRegistryName() {
	_jsii_.InvokeVoid(
		e,
		"resetRoutingRegistryName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ExpressRouteCircuitPeeringIpv6OutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ExpressRouteCircuitPeeringIpv6
	SetInternalValue(val *ExpressRouteCircuitPeeringIpv6)
	MicrosoftPeering() ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference
	MicrosoftPeeringInput() *ExpressRouteCircuitPeeringIpv6MicrosoftPeering
	PrimaryPeerAddressPrefix() *string
	SetPrimaryPeerAddressPrefix(val *string)
	PrimaryPeerAddressPrefixInput() *string
	RouteFilterId() *string
	SetRouteFilterId(val *string)
	RouteFilterIdInput() *string
	SecondaryPeerAddressPrefix() *string
	SetSecondaryPeerAddressPrefix(val *string)
	SecondaryPeerAddressPrefixInput() *string
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
	PutMicrosoftPeering(value *ExpressRouteCircuitPeeringIpv6MicrosoftPeering)
	ResetEnabled()
	ResetMicrosoftPeering()
	ResetRouteFilterId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ExpressRouteCircuitPeeringIpv6OutputReference
type jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) InternalValue() *ExpressRouteCircuitPeeringIpv6 {
	var returns *ExpressRouteCircuitPeeringIpv6
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) MicrosoftPeering() ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference {
	var returns ExpressRouteCircuitPeeringIpv6MicrosoftPeeringOutputReference
	_jsii_.Get(
		j,
		"microsoftPeering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) MicrosoftPeeringInput() *ExpressRouteCircuitPeeringIpv6MicrosoftPeering {
	var returns *ExpressRouteCircuitPeeringIpv6MicrosoftPeering
	_jsii_.Get(
		j,
		"microsoftPeeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) PrimaryPeerAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPeerAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) PrimaryPeerAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPeerAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) RouteFilterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFilterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) RouteFilterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFilterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SecondaryPeerAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPeerAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SecondaryPeerAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPeerAddressPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExpressRouteCircuitPeeringIpv6OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExpressRouteCircuitPeeringIpv6OutputReference {
	_init_.Initialize()

	j := jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringIpv6OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExpressRouteCircuitPeeringIpv6OutputReference_Override(e ExpressRouteCircuitPeeringIpv6OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringIpv6OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetInternalValue(val *ExpressRouteCircuitPeeringIpv6) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetPrimaryPeerAddressPrefix(val *string) {
	_jsii_.Set(
		j,
		"primaryPeerAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetRouteFilterId(val *string) {
	_jsii_.Set(
		j,
		"routeFilterId",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetSecondaryPeerAddressPrefix(val *string) {
	_jsii_.Set(
		j,
		"secondaryPeerAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) PutMicrosoftPeering(value *ExpressRouteCircuitPeeringIpv6MicrosoftPeering) {
	_jsii_.InvokeVoid(
		e,
		"putMicrosoftPeering",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ResetMicrosoftPeering() {
	_jsii_.InvokeVoid(
		e,
		"resetMicrosoftPeering",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ResetRouteFilterId() {
	_jsii_.InvokeVoid(
		e,
		"resetRouteFilterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringIpv6OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ExpressRouteCircuitPeeringMicrosoftPeeringConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#advertised_public_prefixes ExpressRouteCircuitPeering#advertised_public_prefixes}.
	AdvertisedPublicPrefixes *[]*string `field:"required" json:"advertisedPublicPrefixes" yaml:"advertisedPublicPrefixes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#customer_asn ExpressRouteCircuitPeering#customer_asn}.
	CustomerAsn *float64 `field:"optional" json:"customerAsn" yaml:"customerAsn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#routing_registry_name ExpressRouteCircuitPeering#routing_registry_name}.
	RoutingRegistryName *string `field:"optional" json:"routingRegistryName" yaml:"routingRegistryName"`
}

type ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference interface {
	cdktf.ComplexObject
	AdvertisedPublicPrefixes() *[]*string
	SetAdvertisedPublicPrefixes(val *[]*string)
	AdvertisedPublicPrefixesInput() *[]*string
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
	CustomerAsn() *float64
	SetCustomerAsn(val *float64)
	CustomerAsnInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ExpressRouteCircuitPeeringMicrosoftPeeringConfig
	SetInternalValue(val *ExpressRouteCircuitPeeringMicrosoftPeeringConfig)
	RoutingRegistryName() *string
	SetRoutingRegistryName(val *string)
	RoutingRegistryNameInput() *string
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
	ResetCustomerAsn()
	ResetRoutingRegistryName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference
type jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) AdvertisedPublicPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedPublicPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) AdvertisedPublicPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedPublicPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) CustomerAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) CustomerAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) InternalValue() *ExpressRouteCircuitPeeringMicrosoftPeeringConfig {
	var returns *ExpressRouteCircuitPeeringMicrosoftPeeringConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) RoutingRegistryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRegistryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) RoutingRegistryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRegistryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference_Override(e ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) SetAdvertisedPublicPrefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"advertisedPublicPrefixes",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) SetCustomerAsn(val *float64) {
	_jsii_.Set(
		j,
		"customerAsn",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) SetInternalValue(val *ExpressRouteCircuitPeeringMicrosoftPeeringConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) SetRoutingRegistryName(val *string) {
	_jsii_.Set(
		j,
		"routingRegistryName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ResetCustomerAsn() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomerAsn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ResetRoutingRegistryName() {
	_jsii_.InvokeVoid(
		e,
		"resetRoutingRegistryName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ExpressRouteCircuitPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#create ExpressRouteCircuitPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#delete ExpressRouteCircuitPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#read ExpressRouteCircuitPeering#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#update ExpressRouteCircuitPeering#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ExpressRouteCircuitPeeringTimeoutsOutputReference interface {
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

// The jsii proxy struct for ExpressRouteCircuitPeeringTimeoutsOutputReference
type jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewExpressRouteCircuitPeeringTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExpressRouteCircuitPeeringTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExpressRouteCircuitPeeringTimeoutsOutputReference_Override(e ExpressRouteCircuitPeeringTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		e,
		"resetCreate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		e,
		"resetDelete",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		e,
		"resetRead",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

