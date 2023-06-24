package expressroutecircuitpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/expressroutecircuitpeering/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/express_route_circuit_peering azurerm_express_route_circuit_peering}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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

func (j *jsiiProxy_ExpressRouteCircuitPeering) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/express_route_circuit_peering azurerm_express_route_circuit_peering} Resource.
func NewExpressRouteCircuitPeering(scope constructs.Construct, id *string, config *ExpressRouteCircuitPeeringConfig) ExpressRouteCircuitPeering {
	_init_.Initialize()

	if err := validateNewExpressRouteCircuitPeeringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExpressRouteCircuitPeering{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/express_route_circuit_peering azurerm_express_route_circuit_peering} Resource.
func NewExpressRouteCircuitPeering_Override(e ExpressRouteCircuitPeering, scope constructs.Construct, id *string, config *ExpressRouteCircuitPeeringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetExpressRouteCircuitName(val *string) {
	if err := j.validateSetExpressRouteCircuitNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressRouteCircuitName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetIpv4Enabled(val interface{}) {
	if err := j.validateSetIpv4EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Enabled",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetPeerAsn(val *float64) {
	if err := j.validateSetPeerAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerAsn",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetPeeringType(val *string) {
	if err := j.validateSetPeeringTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peeringType",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetPrimaryPeerAddressPrefix(val *string) {
	if err := j.validateSetPrimaryPeerAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryPeerAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetRouteFilterId(val *string) {
	if err := j.validateSetRouteFilterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeFilterId",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetSecondaryPeerAddressPrefix(val *string) {
	if err := j.validateSetSecondaryPeerAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPeerAddressPrefix",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetSharedKey(val *string) {
	if err := j.validateSetSharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedKey",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeering)SetVlanId(val *float64) {
	if err := j.validateSetVlanIdParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateExpressRouteCircuitPeering_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExpressRouteCircuitPeering_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExpressRouteCircuitPeering_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExpressRouteCircuitPeering_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExpressRouteCircuitPeering_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeering",
		"isTerraformResource",
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
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) PutIpv6(value *ExpressRouteCircuitPeeringIpv6) {
	if err := e.validatePutIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv6",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) PutMicrosoftPeeringConfig(value *ExpressRouteCircuitPeeringMicrosoftPeeringConfig) {
	if err := e.validatePutMicrosoftPeeringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMicrosoftPeeringConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeering) PutTimeouts(value *ExpressRouteCircuitPeeringTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

