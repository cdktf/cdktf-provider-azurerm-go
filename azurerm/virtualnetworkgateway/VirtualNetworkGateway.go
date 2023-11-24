// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/virtualnetworkgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/virtual_network_gateway azurerm_virtual_network_gateway}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
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

func (j *jsiiProxy_VirtualNetworkGateway) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/virtual_network_gateway azurerm_virtual_network_gateway} Resource.
func NewVirtualNetworkGateway(scope constructs.Construct, id *string, config *VirtualNetworkGatewayConfig) VirtualNetworkGateway {
	_init_.Initialize()

	if err := validateNewVirtualNetworkGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualNetworkGateway{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/virtual_network_gateway azurerm_virtual_network_gateway} Resource.
func NewVirtualNetworkGateway_Override(v VirtualNetworkGateway, scope constructs.Construct, id *string, config *VirtualNetworkGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetActiveActive(val interface{}) {
	if err := j.validateSetActiveActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeActive",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetDefaultLocalNetworkGatewayId(val *string) {
	if err := j.validateSetDefaultLocalNetworkGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLocalNetworkGatewayId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetEdgeZone(val *string) {
	if err := j.validateSetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetEnableBgp(val interface{}) {
	if err := j.validateSetEnableBgpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBgp",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetGeneration(val *string) {
	if err := j.validateSetGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generation",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetPrivateIpAddressEnabled(val interface{}) {
	if err := j.validateSetPrivateIpAddressEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddressEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGateway)SetVpnType(val *string) {
	if err := j.validateSetVpnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnType",
		val,
	)
}

// Generates CDKTF code for importing a VirtualNetworkGateway resource upon running "cdktf plan <stack-name>".
func VirtualNetworkGateway_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualNetworkGateway_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
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
func VirtualNetworkGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetworkGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetworkGateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGateway",
		"isTerraformResource",
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

func (v *jsiiProxy_VirtualNetworkGateway) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualNetworkGateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualNetworkGateway) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualNetworkGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualNetworkGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualNetworkGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualNetworkGateway) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualNetworkGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualNetworkGateway) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkGateway) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutBgpSettings(value *VirtualNetworkGatewayBgpSettings) {
	if err := v.validatePutBgpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putBgpSettings",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutCustomRoute(value *VirtualNetworkGatewayCustomRoute) {
	if err := v.validatePutCustomRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomRoute",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutIpConfiguration(value interface{}) {
	if err := v.validatePutIpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpConfiguration",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutTimeouts(value *VirtualNetworkGatewayTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGateway) PutVpnClientConfiguration(value *VirtualNetworkGatewayVpnClientConfiguration) {
	if err := v.validatePutVpnClientConfigurationParameters(value); err != nil {
		panic(err)
	}
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

