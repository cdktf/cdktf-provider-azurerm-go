// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkattacheddatanetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/mobilenetworkattacheddatanetwork/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/mobile_network_attached_data_network azurerm_mobile_network_attached_data_network}.
type MobileNetworkAttachedDataNetwork interface {
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
	DnsAddresses() *[]*string
	SetDnsAddresses(val *[]*string)
	DnsAddressesInput() *[]*string
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MobileNetworkDataNetworkName() *string
	SetMobileNetworkDataNetworkName(val *string)
	MobileNetworkDataNetworkNameInput() *string
	MobileNetworkPacketCoreDataPlaneId() *string
	SetMobileNetworkPacketCoreDataPlaneId(val *string)
	MobileNetworkPacketCoreDataPlaneIdInput() *string
	NetworkAddressPortTranslation() MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference
	NetworkAddressPortTranslationInput() *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MobileNetworkAttachedDataNetworkTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserEquipmentAddressPoolPrefixes() *[]*string
	SetUserEquipmentAddressPoolPrefixes(val *[]*string)
	UserEquipmentAddressPoolPrefixesInput() *[]*string
	UserEquipmentStaticAddressPoolPrefixes() *[]*string
	SetUserEquipmentStaticAddressPoolPrefixes(val *[]*string)
	UserEquipmentStaticAddressPoolPrefixesInput() *[]*string
	UserPlaneAccessIpv4Address() *string
	SetUserPlaneAccessIpv4Address(val *string)
	UserPlaneAccessIpv4AddressInput() *string
	UserPlaneAccessIpv4Gateway() *string
	SetUserPlaneAccessIpv4Gateway(val *string)
	UserPlaneAccessIpv4GatewayInput() *string
	UserPlaneAccessIpv4Subnet() *string
	SetUserPlaneAccessIpv4Subnet(val *string)
	UserPlaneAccessIpv4SubnetInput() *string
	UserPlaneAccessName() *string
	SetUserPlaneAccessName(val *string)
	UserPlaneAccessNameInput() *string
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
	PutNetworkAddressPortTranslation(value *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation)
	PutTimeouts(value *MobileNetworkAttachedDataNetworkTimeouts)
	ResetId()
	ResetNetworkAddressPortTranslation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
	ResetUserEquipmentAddressPoolPrefixes()
	ResetUserEquipmentStaticAddressPoolPrefixes()
	ResetUserPlaneAccessIpv4Address()
	ResetUserPlaneAccessIpv4Gateway()
	ResetUserPlaneAccessIpv4Subnet()
	ResetUserPlaneAccessName()
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

// The jsii proxy struct for MobileNetworkAttachedDataNetwork
type jsiiProxy_MobileNetworkAttachedDataNetwork struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) DnsAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) DnsAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) MobileNetworkDataNetworkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkDataNetworkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) MobileNetworkDataNetworkNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkDataNetworkNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) MobileNetworkPacketCoreDataPlaneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkPacketCoreDataPlaneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) MobileNetworkPacketCoreDataPlaneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkPacketCoreDataPlaneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) NetworkAddressPortTranslation() MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference {
	var returns MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationOutputReference
	_jsii_.Get(
		j,
		"networkAddressPortTranslation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) NetworkAddressPortTranslationInput() *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation {
	var returns *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation
	_jsii_.Get(
		j,
		"networkAddressPortTranslationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) Timeouts() MobileNetworkAttachedDataNetworkTimeoutsOutputReference {
	var returns MobileNetworkAttachedDataNetworkTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserEquipmentAddressPoolPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userEquipmentAddressPoolPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserEquipmentAddressPoolPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userEquipmentAddressPoolPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserEquipmentStaticAddressPoolPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userEquipmentStaticAddressPoolPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserEquipmentStaticAddressPoolPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userEquipmentStaticAddressPoolPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4SubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserPlaneAccessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork) UserPlaneAccessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/mobile_network_attached_data_network azurerm_mobile_network_attached_data_network} Resource.
func NewMobileNetworkAttachedDataNetwork(scope constructs.Construct, id *string, config *MobileNetworkAttachedDataNetworkConfig) MobileNetworkAttachedDataNetwork {
	_init_.Initialize()

	if err := validateNewMobileNetworkAttachedDataNetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkAttachedDataNetwork{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetwork",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/mobile_network_attached_data_network azurerm_mobile_network_attached_data_network} Resource.
func NewMobileNetworkAttachedDataNetwork_Override(m MobileNetworkAttachedDataNetwork, scope constructs.Construct, id *string, config *MobileNetworkAttachedDataNetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetwork",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetDnsAddresses(val *[]*string) {
	if err := j.validateSetDnsAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsAddresses",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetMobileNetworkDataNetworkName(val *string) {
	if err := j.validateSetMobileNetworkDataNetworkNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobileNetworkDataNetworkName",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetMobileNetworkPacketCoreDataPlaneId(val *string) {
	if err := j.validateSetMobileNetworkPacketCoreDataPlaneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobileNetworkPacketCoreDataPlaneId",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetUserEquipmentAddressPoolPrefixes(val *[]*string) {
	if err := j.validateSetUserEquipmentAddressPoolPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEquipmentAddressPoolPrefixes",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetUserEquipmentStaticAddressPoolPrefixes(val *[]*string) {
	if err := j.validateSetUserEquipmentStaticAddressPoolPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEquipmentStaticAddressPoolPrefixes",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetUserPlaneAccessIpv4Address(val *string) {
	if err := j.validateSetUserPlaneAccessIpv4AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPlaneAccessIpv4Address",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetUserPlaneAccessIpv4Gateway(val *string) {
	if err := j.validateSetUserPlaneAccessIpv4GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPlaneAccessIpv4Gateway",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetUserPlaneAccessIpv4Subnet(val *string) {
	if err := j.validateSetUserPlaneAccessIpv4SubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPlaneAccessIpv4Subnet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkAttachedDataNetwork)SetUserPlaneAccessName(val *string) {
	if err := j.validateSetUserPlaneAccessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPlaneAccessName",
		val,
	)
}

// Generates CDKTF code for importing a MobileNetworkAttachedDataNetwork resource upon running "cdktf plan <stack-name>".
func MobileNetworkAttachedDataNetwork_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMobileNetworkAttachedDataNetwork_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetwork",
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
func MobileNetworkAttachedDataNetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkAttachedDataNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MobileNetworkAttachedDataNetwork_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkAttachedDataNetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetwork",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MobileNetworkAttachedDataNetwork_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkAttachedDataNetwork_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetwork",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MobileNetworkAttachedDataNetwork_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mobileNetworkAttachedDataNetwork.MobileNetworkAttachedDataNetwork",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) PutNetworkAddressPortTranslation(value *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation) {
	if err := m.validatePutNetworkAddressPortTranslationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNetworkAddressPortTranslation",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) PutTimeouts(value *MobileNetworkAttachedDataNetworkTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetNetworkAddressPortTranslation() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworkAddressPortTranslation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetUserEquipmentAddressPoolPrefixes() {
	_jsii_.InvokeVoid(
		m,
		"resetUserEquipmentAddressPoolPrefixes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetUserEquipmentStaticAddressPoolPrefixes() {
	_jsii_.InvokeVoid(
		m,
		"resetUserEquipmentStaticAddressPoolPrefixes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetUserPlaneAccessIpv4Address() {
	_jsii_.InvokeVoid(
		m,
		"resetUserPlaneAccessIpv4Address",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetUserPlaneAccessIpv4Gateway() {
	_jsii_.InvokeVoid(
		m,
		"resetUserPlaneAccessIpv4Gateway",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetUserPlaneAccessIpv4Subnet() {
	_jsii_.InvokeVoid(
		m,
		"resetUserPlaneAccessIpv4Subnet",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ResetUserPlaneAccessName() {
	_jsii_.InvokeVoid(
		m,
		"resetUserPlaneAccessName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkAttachedDataNetwork) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

