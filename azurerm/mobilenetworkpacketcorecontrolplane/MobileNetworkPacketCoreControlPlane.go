// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkpacketcorecontrolplane

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/mobilenetworkpacketcorecontrolplane/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/mobile_network_packet_core_control_plane azurerm_mobile_network_packet_core_control_plane}.
type MobileNetworkPacketCoreControlPlane interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlaneAccessIpv4Address() *string
	SetControlPlaneAccessIpv4Address(val *string)
	ControlPlaneAccessIpv4AddressInput() *string
	ControlPlaneAccessIpv4Gateway() *string
	SetControlPlaneAccessIpv4Gateway(val *string)
	ControlPlaneAccessIpv4GatewayInput() *string
	ControlPlaneAccessIpv4Subnet() *string
	SetControlPlaneAccessIpv4Subnet(val *string)
	ControlPlaneAccessIpv4SubnetInput() *string
	ControlPlaneAccessName() *string
	SetControlPlaneAccessName(val *string)
	ControlPlaneAccessNameInput() *string
	CoreNetworkTechnology() *string
	SetCoreNetworkTechnology(val *string)
	CoreNetworkTechnologyInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Identity() MobileNetworkPacketCoreControlPlaneIdentityOutputReference
	IdentityInput() *MobileNetworkPacketCoreControlPlaneIdentity
	IdInput() *string
	InteroperabilitySettingsJson() *string
	SetInteroperabilitySettingsJson(val *string)
	InteroperabilitySettingsJsonInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalDiagnosticsAccess() MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccessOutputReference
	LocalDiagnosticsAccessInput() *MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccess
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Platform() MobileNetworkPacketCoreControlPlanePlatformOutputReference
	PlatformInput() *MobileNetworkPacketCoreControlPlanePlatform
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
	SiteIds() *[]*string
	SetSiteIds(val *[]*string)
	SiteIdsInput() *[]*string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	SoftwareVersion() *string
	SetSoftwareVersion(val *string)
	SoftwareVersionInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MobileNetworkPacketCoreControlPlaneTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserEquipmentMtuInBytes() *float64
	SetUserEquipmentMtuInBytes(val *float64)
	UserEquipmentMtuInBytesInput() *float64
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
	PutIdentity(value *MobileNetworkPacketCoreControlPlaneIdentity)
	PutLocalDiagnosticsAccess(value *MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccess)
	PutPlatform(value *MobileNetworkPacketCoreControlPlanePlatform)
	PutTimeouts(value *MobileNetworkPacketCoreControlPlaneTimeouts)
	ResetControlPlaneAccessIpv4Address()
	ResetControlPlaneAccessIpv4Gateway()
	ResetControlPlaneAccessIpv4Subnet()
	ResetControlPlaneAccessName()
	ResetCoreNetworkTechnology()
	ResetId()
	ResetIdentity()
	ResetInteroperabilitySettingsJson()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetSoftwareVersion()
	ResetTags()
	ResetTimeouts()
	ResetUserEquipmentMtuInBytes()
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

// The jsii proxy struct for MobileNetworkPacketCoreControlPlane
type jsiiProxy_MobileNetworkPacketCoreControlPlane struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ControlPlaneAccessIpv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneAccessIpv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ControlPlaneAccessIpv4AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneAccessIpv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ControlPlaneAccessIpv4Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneAccessIpv4Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ControlPlaneAccessIpv4GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneAccessIpv4GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ControlPlaneAccessIpv4Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneAccessIpv4Subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ControlPlaneAccessIpv4SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneAccessIpv4SubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ControlPlaneAccessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneAccessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ControlPlaneAccessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneAccessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) CoreNetworkTechnology() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkTechnology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) CoreNetworkTechnologyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkTechnologyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Identity() MobileNetworkPacketCoreControlPlaneIdentityOutputReference {
	var returns MobileNetworkPacketCoreControlPlaneIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) IdentityInput() *MobileNetworkPacketCoreControlPlaneIdentity {
	var returns *MobileNetworkPacketCoreControlPlaneIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) InteroperabilitySettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interoperabilitySettingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) InteroperabilitySettingsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interoperabilitySettingsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) LocalDiagnosticsAccess() MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccessOutputReference {
	var returns MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccessOutputReference
	_jsii_.Get(
		j,
		"localDiagnosticsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) LocalDiagnosticsAccessInput() *MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccess {
	var returns *MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccess
	_jsii_.Get(
		j,
		"localDiagnosticsAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Platform() MobileNetworkPacketCoreControlPlanePlatformOutputReference {
	var returns MobileNetworkPacketCoreControlPlanePlatformOutputReference
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) PlatformInput() *MobileNetworkPacketCoreControlPlanePlatform {
	var returns *MobileNetworkPacketCoreControlPlanePlatform
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) SiteIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"siteIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) SiteIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"siteIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) SoftwareVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"softwareVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) SoftwareVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"softwareVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) Timeouts() MobileNetworkPacketCoreControlPlaneTimeoutsOutputReference {
	var returns MobileNetworkPacketCoreControlPlaneTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) UserEquipmentMtuInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userEquipmentMtuInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane) UserEquipmentMtuInBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userEquipmentMtuInBytesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/mobile_network_packet_core_control_plane azurerm_mobile_network_packet_core_control_plane} Resource.
func NewMobileNetworkPacketCoreControlPlane(scope constructs.Construct, id *string, config *MobileNetworkPacketCoreControlPlaneConfig) MobileNetworkPacketCoreControlPlane {
	_init_.Initialize()

	if err := validateNewMobileNetworkPacketCoreControlPlaneParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkPacketCoreControlPlane{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlane",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/mobile_network_packet_core_control_plane azurerm_mobile_network_packet_core_control_plane} Resource.
func NewMobileNetworkPacketCoreControlPlane_Override(m MobileNetworkPacketCoreControlPlane, scope constructs.Construct, id *string, config *MobileNetworkPacketCoreControlPlaneConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlane",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetControlPlaneAccessIpv4Address(val *string) {
	if err := j.validateSetControlPlaneAccessIpv4AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneAccessIpv4Address",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetControlPlaneAccessIpv4Gateway(val *string) {
	if err := j.validateSetControlPlaneAccessIpv4GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneAccessIpv4Gateway",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetControlPlaneAccessIpv4Subnet(val *string) {
	if err := j.validateSetControlPlaneAccessIpv4SubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneAccessIpv4Subnet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetControlPlaneAccessName(val *string) {
	if err := j.validateSetControlPlaneAccessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneAccessName",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetCoreNetworkTechnology(val *string) {
	if err := j.validateSetCoreNetworkTechnologyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNetworkTechnology",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetInteroperabilitySettingsJson(val *string) {
	if err := j.validateSetInteroperabilitySettingsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interoperabilitySettingsJson",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetSiteIds(val *[]*string) {
	if err := j.validateSetSiteIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteIds",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetSoftwareVersion(val *string) {
	if err := j.validateSetSoftwareVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softwareVersion",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreControlPlane)SetUserEquipmentMtuInBytes(val *float64) {
	if err := j.validateSetUserEquipmentMtuInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEquipmentMtuInBytes",
		val,
	)
}

// Generates CDKTF code for importing a MobileNetworkPacketCoreControlPlane resource upon running "cdktf plan <stack-name>".
func MobileNetworkPacketCoreControlPlane_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMobileNetworkPacketCoreControlPlane_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlane",
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
func MobileNetworkPacketCoreControlPlane_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkPacketCoreControlPlane_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlane",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MobileNetworkPacketCoreControlPlane_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkPacketCoreControlPlane_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlane",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MobileNetworkPacketCoreControlPlane_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkPacketCoreControlPlane_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlane",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MobileNetworkPacketCoreControlPlane_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreControlPlane.MobileNetworkPacketCoreControlPlane",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) PutIdentity(value *MobileNetworkPacketCoreControlPlaneIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) PutLocalDiagnosticsAccess(value *MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccess) {
	if err := m.validatePutLocalDiagnosticsAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLocalDiagnosticsAccess",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) PutPlatform(value *MobileNetworkPacketCoreControlPlanePlatform) {
	if err := m.validatePutPlatformParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPlatform",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) PutTimeouts(value *MobileNetworkPacketCoreControlPlaneTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetControlPlaneAccessIpv4Address() {
	_jsii_.InvokeVoid(
		m,
		"resetControlPlaneAccessIpv4Address",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetControlPlaneAccessIpv4Gateway() {
	_jsii_.InvokeVoid(
		m,
		"resetControlPlaneAccessIpv4Gateway",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetControlPlaneAccessIpv4Subnet() {
	_jsii_.InvokeVoid(
		m,
		"resetControlPlaneAccessIpv4Subnet",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetControlPlaneAccessName() {
	_jsii_.InvokeVoid(
		m,
		"resetControlPlaneAccessName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetCoreNetworkTechnology() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreNetworkTechnology",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetInteroperabilitySettingsJson() {
	_jsii_.InvokeVoid(
		m,
		"resetInteroperabilitySettingsJson",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetPlatform() {
	_jsii_.InvokeVoid(
		m,
		"resetPlatform",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetSoftwareVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetSoftwareVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ResetUserEquipmentMtuInBytes() {
	_jsii_.InvokeVoid(
		m,
		"resetUserEquipmentMtuInBytes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreControlPlane) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

