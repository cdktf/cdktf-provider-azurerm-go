package virtualnetworkgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/virtualnetworkgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway azurerm_virtual_network_gateway}.
type VirtualNetworkGateway interface {
	cdktf.TerraformResource
	ActiveActive() interface{}
	SetActiveActive(val interface{})
	ActiveActiveInput() interface{}
	BgpSettings() VirtualNetworkGatewayBgpSettingsOutputReference
	BgpSettingsInput() *VirtualNetworkGatewayBgpSettings
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
	CustomRoute() VirtualNetworkGatewayCustomRouteOutputReference
	CustomRouteInput() *VirtualNetworkGatewayCustomRoute
	DefaultLocalNetworkGatewayId() *string
	SetDefaultLocalNetworkGatewayId(val *string)
	DefaultLocalNetworkGatewayIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	EnableBgp() interface{}
	SetEnableBgp(val interface{})
	EnableBgpInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Generation() *string
	SetGeneration(val *string)
	GenerationInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpConfiguration() VirtualNetworkGatewayIpConfigurationList
	IpConfigurationInput() interface{}
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
	PrivateIpAddressEnabled() interface{}
	SetPrivateIpAddressEnabled(val interface{})
	PrivateIpAddressEnabledInput() interface{}
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
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualNetworkGatewayTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VpnClientConfiguration() VirtualNetworkGatewayVpnClientConfigurationOutputReference
	VpnClientConfigurationInput() *VirtualNetworkGatewayVpnClientConfiguration
	VpnType() *string
	SetVpnType(val *string)
	VpnTypeInput() *string
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
	PutBgpSettings(value *VirtualNetworkGatewayBgpSettings)
	PutCustomRoute(value *VirtualNetworkGatewayCustomRoute)
	PutIpConfiguration(value interface{})
	PutTimeouts(value *VirtualNetworkGatewayTimeouts)
	PutVpnClientConfiguration(value *VirtualNetworkGatewayVpnClientConfiguration)
	ResetActiveActive()
	ResetBgpSettings()
	ResetCustomRoute()
	ResetDefaultLocalNetworkGatewayId()
	ResetEdgeZone()
	ResetEnableBgp()
	ResetGeneration()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateIpAddressEnabled()
	ResetTags()
	ResetTimeouts()
	ResetVpnClientConfiguration()
	ResetVpnType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VirtualNetworkGateway
type jsiiProxy_VirtualNetworkGateway struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualNetworkGateway) ActiveActive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeActive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) ActiveActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeActiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) BgpSettings() VirtualNetworkGatewayBgpSettingsOutputReference {
	var returns VirtualNetworkGatewayBgpSettingsOutputReference
	_jsii_.Get(
		j,
		"bgpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) BgpSettingsInput() *VirtualNetworkGatewayBgpSettings {
	var returns *VirtualNetworkGatewayBgpSettings
	_jsii_.Get(
		j,
		"bgpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) CustomRoute() VirtualNetworkGatewayCustomRouteOutputReference {
	var returns VirtualNetworkGatewayCustomRouteOutputReference
	_jsii_.Get(
		j,
		"customRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) CustomRouteInput() *VirtualNetworkGatewayCustomRoute {
	var returns *VirtualNetworkGatewayCustomRoute
	_jsii_.Get(
		j,
		"customRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) DefaultLocalNetworkGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLocalNetworkGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) DefaultLocalNetworkGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLocalNetworkGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) EnableBgp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBgp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) EnableBgpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBgpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Generation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) GenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) IpConfiguration() VirtualNetworkGatewayIpConfigurationList {
	var returns VirtualNetworkGatewayIpConfigurationList
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) IpConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) PrivateIpAddressEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpAddressEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) PrivateIpAddressEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpAddressEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Timeouts() VirtualNetworkGatewayTimeoutsOutputReference {
	var returns VirtualNetworkGatewayTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) VpnClientConfiguration() VirtualNetworkGatewayVpnClientConfigurationOutputReference {
	var returns VirtualNetworkGatewayVpnClientConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpnClientConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) VpnClientConfigurationInput() *VirtualNetworkGatewayVpnClientConfiguration {
	var returns *VirtualNetworkGatewayVpnClientConfiguration
	_jsii_.Get(
		j,
		"vpnClientConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) VpnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGateway) VpnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway azurerm_virtual_network_gateway} Resource.
func NewVirtualNetworkGateway(scope constructs.Construct, id *string, config *VirtualNetworkGatewayConfig) VirtualNetworkGateway {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGateway{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway azurerm_virtual_network_gateway} Resource.
func NewVirtualNetworkGateway_Override(v VirtualNetworkGateway, scope constructs.Construct, id *string, config *VirtualNetworkGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetActiveActive(val interface{}) {
	_jsii_.Set(
		j,
		"activeActive",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetDefaultLocalNetworkGatewayId(val *string) {
	_jsii_.Set(
		j,
		"defaultLocalNetworkGatewayId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetEdgeZone(val *string) {
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetEnableBgp(val interface{}) {
	_jsii_.Set(
		j,
		"enableBgp",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetGeneration(val *string) {
	_jsii_.Set(
		j,
		"generation",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetPrivateIpAddressEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"privateIpAddressEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway) SetVpnType(val *string) {
	_jsii_.Set(
		j,
		"vpnType",
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
func VirtualNetworkGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualNetworkGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutBgpSettings(value *VirtualNetworkGatewayBgpSettings) {
	_jsii_.InvokeVoid(
		v,
		"putBgpSettings",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutCustomRoute(value *VirtualNetworkGatewayCustomRoute) {
	_jsii_.InvokeVoid(
		v,
		"putCustomRoute",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutIpConfiguration(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putIpConfiguration",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutTimeouts(value *VirtualNetworkGatewayTimeouts) {
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutVpnClientConfiguration(value *VirtualNetworkGatewayVpnClientConfiguration) {
	_jsii_.InvokeVoid(
		v,
		"putVpnClientConfiguration",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetActiveActive() {
	_jsii_.InvokeVoid(
		v,
		"resetActiveActive",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetBgpSettings() {
	_jsii_.InvokeVoid(
		v,
		"resetBgpSettings",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetCustomRoute() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomRoute",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetDefaultLocalNetworkGatewayId() {
	_jsii_.InvokeVoid(
		v,
		"resetDefaultLocalNetworkGatewayId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		v,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetEnableBgp() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableBgp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetGeneration() {
	_jsii_.InvokeVoid(
		v,
		"resetGeneration",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetPrivateIpAddressEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetPrivateIpAddressEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetVpnClientConfiguration() {
	_jsii_.InvokeVoid(
		v,
		"resetVpnClientConfiguration",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) ResetVpnType() {
	_jsii_.InvokeVoid(
		v,
		"resetVpnType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayBgpSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#asn VirtualNetworkGateway#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// peering_addresses block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#peering_addresses VirtualNetworkGateway#peering_addresses}
	PeeringAddresses interface{} `field:"optional" json:"peeringAddresses" yaml:"peeringAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#peer_weight VirtualNetworkGateway#peer_weight}.
	PeerWeight *float64 `field:"optional" json:"peerWeight" yaml:"peerWeight"`
}

type VirtualNetworkGatewayBgpSettingsOutputReference interface {
	cdktf.ComplexObject
	Asn() *float64
	SetAsn(val *float64)
	AsnInput() *float64
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
	InternalValue() *VirtualNetworkGatewayBgpSettings
	SetInternalValue(val *VirtualNetworkGatewayBgpSettings)
	PeeringAddresses() VirtualNetworkGatewayBgpSettingsPeeringAddressesList
	PeeringAddressesInput() interface{}
	PeerWeight() *float64
	SetPeerWeight(val *float64)
	PeerWeightInput() *float64
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
	PutPeeringAddresses(value interface{})
	ResetAsn()
	ResetPeeringAddresses()
	ResetPeerWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayBgpSettingsOutputReference
type jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) Asn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) AsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) InternalValue() *VirtualNetworkGatewayBgpSettings {
	var returns *VirtualNetworkGatewayBgpSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) PeeringAddresses() VirtualNetworkGatewayBgpSettingsPeeringAddressesList {
	var returns VirtualNetworkGatewayBgpSettingsPeeringAddressesList
	_jsii_.Get(
		j,
		"peeringAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) PeeringAddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"peeringAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) PeerWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) PeerWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayBgpSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayBgpSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayBgpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayBgpSettingsOutputReference_Override(v VirtualNetworkGatewayBgpSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayBgpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) SetAsn(val *float64) {
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) SetInternalValue(val *VirtualNetworkGatewayBgpSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) SetPeerWeight(val *float64) {
	_jsii_.Set(
		j,
		"peerWeight",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) PutPeeringAddresses(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putPeeringAddresses",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) ResetAsn() {
	_jsii_.InvokeVoid(
		v,
		"resetAsn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) ResetPeeringAddresses() {
	_jsii_.InvokeVoid(
		v,
		"resetPeeringAddresses",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) ResetPeerWeight() {
	_jsii_.InvokeVoid(
		v,
		"resetPeerWeight",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayBgpSettingsPeeringAddresses struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#apipa_addresses VirtualNetworkGateway#apipa_addresses}.
	ApipaAddresses *[]*string `field:"optional" json:"apipaAddresses" yaml:"apipaAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#ip_configuration_name VirtualNetworkGateway#ip_configuration_name}.
	IpConfigurationName *string `field:"optional" json:"ipConfigurationName" yaml:"ipConfigurationName"`
}

type VirtualNetworkGatewayBgpSettingsPeeringAddressesList interface {
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
	Get(index *float64) VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayBgpSettingsPeeringAddressesList
type jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayBgpSettingsPeeringAddressesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualNetworkGatewayBgpSettingsPeeringAddressesList {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayBgpSettingsPeeringAddressesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayBgpSettingsPeeringAddressesList_Override(v VirtualNetworkGatewayBgpSettingsPeeringAddressesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayBgpSettingsPeeringAddressesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) Get(index *float64) VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference {
	var returns VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference interface {
	cdktf.ComplexObject
	ApipaAddresses() *[]*string
	SetApipaAddresses(val *[]*string)
	ApipaAddressesInput() *[]*string
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
	DefaultAddresses() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpConfigurationName() *string
	SetIpConfigurationName(val *string)
	IpConfigurationNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TunnelIpAddresses() *[]*string
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
	ResetApipaAddresses()
	ResetIpConfigurationName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference
type jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) ApipaAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apipaAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) ApipaAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apipaAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) DefaultAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) IpConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) IpConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) TunnelIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelIpAddresses",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference_Override(v VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) SetApipaAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"apipaAddresses",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) SetIpConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"ipConfigurationName",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) ResetApipaAddresses() {
	_jsii_.InvokeVoid(
		v,
		"resetApipaAddresses",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) ResetIpConfigurationName() {
	_jsii_.InvokeVoid(
		v,
		"resetIpConfigurationName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayBgpSettingsPeeringAddressesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayConfig struct {
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
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#ip_configuration VirtualNetworkGateway#ip_configuration}
	IpConfiguration interface{} `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#location VirtualNetworkGateway#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#resource_group_name VirtualNetworkGateway#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#sku VirtualNetworkGateway#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#type VirtualNetworkGateway#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#active_active VirtualNetworkGateway#active_active}.
	ActiveActive interface{} `field:"optional" json:"activeActive" yaml:"activeActive"`
	// bgp_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#bgp_settings VirtualNetworkGateway#bgp_settings}
	BgpSettings *VirtualNetworkGatewayBgpSettings `field:"optional" json:"bgpSettings" yaml:"bgpSettings"`
	// custom_route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#custom_route VirtualNetworkGateway#custom_route}
	CustomRoute *VirtualNetworkGatewayCustomRoute `field:"optional" json:"customRoute" yaml:"customRoute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#default_local_network_gateway_id VirtualNetworkGateway#default_local_network_gateway_id}.
	DefaultLocalNetworkGatewayId *string `field:"optional" json:"defaultLocalNetworkGatewayId" yaml:"defaultLocalNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#edge_zone VirtualNetworkGateway#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#enable_bgp VirtualNetworkGateway#enable_bgp}.
	EnableBgp interface{} `field:"optional" json:"enableBgp" yaml:"enableBgp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#generation VirtualNetworkGateway#generation}.
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#id VirtualNetworkGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#private_ip_address_enabled VirtualNetworkGateway#private_ip_address_enabled}.
	PrivateIpAddressEnabled interface{} `field:"optional" json:"privateIpAddressEnabled" yaml:"privateIpAddressEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#tags VirtualNetworkGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#timeouts VirtualNetworkGateway#timeouts}
	Timeouts *VirtualNetworkGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpn_client_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#vpn_client_configuration VirtualNetworkGateway#vpn_client_configuration}
	VpnClientConfiguration *VirtualNetworkGatewayVpnClientConfiguration `field:"optional" json:"vpnClientConfiguration" yaml:"vpnClientConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#vpn_type VirtualNetworkGateway#vpn_type}.
	VpnType *string `field:"optional" json:"vpnType" yaml:"vpnType"`
}

type VirtualNetworkGatewayCustomRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#address_prefixes VirtualNetworkGateway#address_prefixes}.
	AddressPrefixes *[]*string `field:"optional" json:"addressPrefixes" yaml:"addressPrefixes"`
}

type VirtualNetworkGatewayCustomRouteOutputReference interface {
	cdktf.ComplexObject
	AddressPrefixes() *[]*string
	SetAddressPrefixes(val *[]*string)
	AddressPrefixesInput() *[]*string
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
	InternalValue() *VirtualNetworkGatewayCustomRoute
	SetInternalValue(val *VirtualNetworkGatewayCustomRoute)
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
	ResetAddressPrefixes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayCustomRouteOutputReference
type jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) AddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) AddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) InternalValue() *VirtualNetworkGatewayCustomRoute {
	var returns *VirtualNetworkGatewayCustomRoute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayCustomRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayCustomRouteOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayCustomRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayCustomRouteOutputReference_Override(v VirtualNetworkGatewayCustomRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayCustomRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) SetAddressPrefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"addressPrefixes",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) SetInternalValue(val *VirtualNetworkGatewayCustomRoute) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) ResetAddressPrefixes() {
	_jsii_.InvokeVoid(
		v,
		"resetAddressPrefixes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayCustomRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#public_ip_address_id VirtualNetworkGateway#public_ip_address_id}.
	PublicIpAddressId *string `field:"required" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#subnet_id VirtualNetworkGateway#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#private_ip_address_allocation VirtualNetworkGateway#private_ip_address_allocation}.
	PrivateIpAddressAllocation *string `field:"optional" json:"privateIpAddressAllocation" yaml:"privateIpAddressAllocation"`
}

type VirtualNetworkGatewayIpConfigurationList interface {
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
	Get(index *float64) VirtualNetworkGatewayIpConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayIpConfigurationList
type jsiiProxy_VirtualNetworkGatewayIpConfigurationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayIpConfigurationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualNetworkGatewayIpConfigurationList {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayIpConfigurationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayIpConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayIpConfigurationList_Override(v VirtualNetworkGatewayIpConfigurationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayIpConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) Get(index *float64) VirtualNetworkGatewayIpConfigurationOutputReference {
	var returns VirtualNetworkGatewayIpConfigurationOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayIpConfigurationOutputReference interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	PrivateIpAddressAllocation() *string
	SetPrivateIpAddressAllocation(val *string)
	PrivateIpAddressAllocationInput() *string
	PublicIpAddressId() *string
	SetPublicIpAddressId(val *string)
	PublicIpAddressIdInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	ResetName()
	ResetPrivateIpAddressAllocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayIpConfigurationOutputReference
type jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) PrivateIpAddressAllocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) PrivateIpAddressAllocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressAllocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) PublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) PublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayIpConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualNetworkGatewayIpConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayIpConfigurationOutputReference_Override(v VirtualNetworkGatewayIpConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetPrivateIpAddressAllocation(val *string) {
	_jsii_.Set(
		j,
		"privateIpAddressAllocation",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetPublicIpAddressId(val *string) {
	_jsii_.Set(
		j,
		"publicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) ResetPrivateIpAddressAllocation() {
	_jsii_.InvokeVoid(
		v,
		"resetPrivateIpAddressAllocation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayIpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#create VirtualNetworkGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#delete VirtualNetworkGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#read VirtualNetworkGateway#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#update VirtualNetworkGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type VirtualNetworkGatewayTimeoutsOutputReference interface {
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

// The jsii proxy struct for VirtualNetworkGatewayTimeoutsOutputReference
type jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayTimeoutsOutputReference_Override(v VirtualNetworkGatewayTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		v,
		"resetCreate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		v,
		"resetDelete",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		v,
		"resetRead",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		v,
		"resetUpdate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayVpnClientConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#address_space VirtualNetworkGateway#address_space}.
	AddressSpace *[]*string `field:"required" json:"addressSpace" yaml:"addressSpace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#aad_audience VirtualNetworkGateway#aad_audience}.
	AadAudience *string `field:"optional" json:"aadAudience" yaml:"aadAudience"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#aad_issuer VirtualNetworkGateway#aad_issuer}.
	AadIssuer *string `field:"optional" json:"aadIssuer" yaml:"aadIssuer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#aad_tenant VirtualNetworkGateway#aad_tenant}.
	AadTenant *string `field:"optional" json:"aadTenant" yaml:"aadTenant"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#radius_server_address VirtualNetworkGateway#radius_server_address}.
	RadiusServerAddress *string `field:"optional" json:"radiusServerAddress" yaml:"radiusServerAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#radius_server_secret VirtualNetworkGateway#radius_server_secret}.
	RadiusServerSecret *string `field:"optional" json:"radiusServerSecret" yaml:"radiusServerSecret"`
	// revoked_certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#revoked_certificate VirtualNetworkGateway#revoked_certificate}
	RevokedCertificate interface{} `field:"optional" json:"revokedCertificate" yaml:"revokedCertificate"`
	// root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#root_certificate VirtualNetworkGateway#root_certificate}
	RootCertificate interface{} `field:"optional" json:"rootCertificate" yaml:"rootCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#vpn_auth_types VirtualNetworkGateway#vpn_auth_types}.
	VpnAuthTypes *[]*string `field:"optional" json:"vpnAuthTypes" yaml:"vpnAuthTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#vpn_client_protocols VirtualNetworkGateway#vpn_client_protocols}.
	VpnClientProtocols *[]*string `field:"optional" json:"vpnClientProtocols" yaml:"vpnClientProtocols"`
}

type VirtualNetworkGatewayVpnClientConfigurationOutputReference interface {
	cdktf.ComplexObject
	AadAudience() *string
	SetAadAudience(val *string)
	AadAudienceInput() *string
	AadIssuer() *string
	SetAadIssuer(val *string)
	AadIssuerInput() *string
	AadTenant() *string
	SetAadTenant(val *string)
	AadTenantInput() *string
	AddressSpace() *[]*string
	SetAddressSpace(val *[]*string)
	AddressSpaceInput() *[]*string
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
	InternalValue() *VirtualNetworkGatewayVpnClientConfiguration
	SetInternalValue(val *VirtualNetworkGatewayVpnClientConfiguration)
	RadiusServerAddress() *string
	SetRadiusServerAddress(val *string)
	RadiusServerAddressInput() *string
	RadiusServerSecret() *string
	SetRadiusServerSecret(val *string)
	RadiusServerSecretInput() *string
	RevokedCertificate() VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList
	RevokedCertificateInput() interface{}
	RootCertificate() VirtualNetworkGatewayVpnClientConfigurationRootCertificateList
	RootCertificateInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpnAuthTypes() *[]*string
	SetVpnAuthTypes(val *[]*string)
	VpnAuthTypesInput() *[]*string
	VpnClientProtocols() *[]*string
	SetVpnClientProtocols(val *[]*string)
	VpnClientProtocolsInput() *[]*string
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
	PutRevokedCertificate(value interface{})
	PutRootCertificate(value interface{})
	ResetAadAudience()
	ResetAadIssuer()
	ResetAadTenant()
	ResetRadiusServerAddress()
	ResetRadiusServerSecret()
	ResetRevokedCertificate()
	ResetRootCertificate()
	ResetVpnAuthTypes()
	ResetVpnClientProtocols()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayVpnClientConfigurationOutputReference
type jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadAudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadTenant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadTenantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadTenantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AddressSpace() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AddressSpaceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) InternalValue() *VirtualNetworkGatewayVpnClientConfiguration {
	var returns *VirtualNetworkGatewayVpnClientConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiusServerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiusServerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiusServerSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiusServerSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RevokedCertificate() VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList {
	var returns VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList
	_jsii_.Get(
		j,
		"revokedCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RevokedCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokedCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RootCertificate() VirtualNetworkGatewayVpnClientConfigurationRootCertificateList {
	var returns VirtualNetworkGatewayVpnClientConfigurationRootCertificateList
	_jsii_.Get(
		j,
		"rootCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RootCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VpnAuthTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnAuthTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VpnAuthTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnAuthTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VpnClientProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnClientProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VpnClientProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnClientProtocolsInput",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayVpnClientConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayVpnClientConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayVpnClientConfigurationOutputReference_Override(v VirtualNetworkGatewayVpnClientConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetAadAudience(val *string) {
	_jsii_.Set(
		j,
		"aadAudience",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetAadIssuer(val *string) {
	_jsii_.Set(
		j,
		"aadIssuer",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetAadTenant(val *string) {
	_jsii_.Set(
		j,
		"aadTenant",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetAddressSpace(val *[]*string) {
	_jsii_.Set(
		j,
		"addressSpace",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetInternalValue(val *VirtualNetworkGatewayVpnClientConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetRadiusServerAddress(val *string) {
	_jsii_.Set(
		j,
		"radiusServerAddress",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetRadiusServerSecret(val *string) {
	_jsii_.Set(
		j,
		"radiusServerSecret",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetVpnAuthTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"vpnAuthTypes",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) SetVpnClientProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"vpnClientProtocols",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) PutRevokedCertificate(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putRevokedCertificate",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) PutRootCertificate(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putRootCertificate",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetAadAudience() {
	_jsii_.InvokeVoid(
		v,
		"resetAadAudience",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetAadIssuer() {
	_jsii_.InvokeVoid(
		v,
		"resetAadIssuer",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetAadTenant() {
	_jsii_.InvokeVoid(
		v,
		"resetAadTenant",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRadiusServerAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetRadiusServerAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRadiusServerSecret() {
	_jsii_.InvokeVoid(
		v,
		"resetRadiusServerSecret",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRevokedCertificate() {
	_jsii_.InvokeVoid(
		v,
		"resetRevokedCertificate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRootCertificate() {
	_jsii_.InvokeVoid(
		v,
		"resetRootCertificate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetVpnAuthTypes() {
	_jsii_.InvokeVoid(
		v,
		"resetVpnAuthTypes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetVpnClientProtocols() {
	_jsii_.InvokeVoid(
		v,
		"resetVpnClientProtocols",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayVpnClientConfigurationRevokedCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#thumbprint VirtualNetworkGateway#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
}

type VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList interface {
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
	Get(index *float64) VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList
type jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList_Override(v VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) Get(index *float64) VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference {
	var returns VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
	SetThumbprint(val *string)
	ThumbprintInput() *string
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

// The jsii proxy struct for VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference
type jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) ThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintInput",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference_Override(v VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) SetThumbprint(val *string) {
	_jsii_.Set(
		j,
		"thumbprint",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayVpnClientConfigurationRootCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#public_cert_data VirtualNetworkGateway#public_cert_data}.
	PublicCertData *string `field:"required" json:"publicCertData" yaml:"publicCertData"`
}

type VirtualNetworkGatewayVpnClientConfigurationRootCertificateList interface {
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
	Get(index *float64) VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayVpnClientConfigurationRootCertificateList
type jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayVpnClientConfigurationRootCertificateList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualNetworkGatewayVpnClientConfigurationRootCertificateList {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationRootCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayVpnClientConfigurationRootCertificateList_Override(v VirtualNetworkGatewayVpnClientConfigurationRootCertificateList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationRootCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) Get(index *float64) VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference {
	var returns VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	PublicCertData() *string
	SetPublicCertData(val *string)
	PublicCertDataInput() *string
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

// The jsii proxy struct for VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference
type jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) PublicCertData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCertData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) PublicCertDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCertDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference_Override(v VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) SetPublicCertData(val *string) {
	_jsii_.Set(
		j,
		"publicCertData",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationRootCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

