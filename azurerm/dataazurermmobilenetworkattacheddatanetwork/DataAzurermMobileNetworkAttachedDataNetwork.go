// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermmobilenetworkattacheddatanetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/dataazurermmobilenetworkattacheddatanetwork/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/data-sources/mobile_network_attached_data_network azurerm_mobile_network_attached_data_network}.
type DataAzurermMobileNetworkAttachedDataNetwork interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	MobileNetworkDataNetworkName() *string
	SetMobileNetworkDataNetworkName(val *string)
	MobileNetworkDataNetworkNameInput() *string
	MobileNetworkPacketCoreDataPlaneId() *string
	SetMobileNetworkPacketCoreDataPlaneId(val *string)
	MobileNetworkPacketCoreDataPlaneIdInput() *string
	NetworkAddressPortTranslation() DataAzurermMobileNetworkAttachedDataNetworkNetworkAddressPortTranslationList
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermMobileNetworkAttachedDataNetworkTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserEquipmentAddressPoolPrefixes() *[]*string
	UserEquipmentStaticAddressPoolPrefixes() *[]*string
	UserPlaneAccessIpv4Address() *string
	UserPlaneAccessIpv4Gateway() *string
	UserPlaneAccessIpv4Subnet() *string
	UserPlaneAccessName() *string
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
	PutTimeouts(value *DataAzurermMobileNetworkAttachedDataNetworkTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataAzurermMobileNetworkAttachedDataNetwork
type jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) DnsAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) MobileNetworkDataNetworkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkDataNetworkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) MobileNetworkDataNetworkNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkDataNetworkNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) MobileNetworkPacketCoreDataPlaneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkPacketCoreDataPlaneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) MobileNetworkPacketCoreDataPlaneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkPacketCoreDataPlaneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) NetworkAddressPortTranslation() DataAzurermMobileNetworkAttachedDataNetworkNetworkAddressPortTranslationList {
	var returns DataAzurermMobileNetworkAttachedDataNetworkNetworkAddressPortTranslationList
	_jsii_.Get(
		j,
		"networkAddressPortTranslation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) Timeouts() DataAzurermMobileNetworkAttachedDataNetworkTimeoutsOutputReference {
	var returns DataAzurermMobileNetworkAttachedDataNetworkTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) UserEquipmentAddressPoolPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userEquipmentAddressPoolPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) UserEquipmentStaticAddressPoolPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userEquipmentStaticAddressPoolPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) UserPlaneAccessIpv4Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) UserPlaneAccessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessName",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/data-sources/mobile_network_attached_data_network azurerm_mobile_network_attached_data_network} Data Source.
func NewDataAzurermMobileNetworkAttachedDataNetwork(scope constructs.Construct, id *string, config *DataAzurermMobileNetworkAttachedDataNetworkConfig) DataAzurermMobileNetworkAttachedDataNetwork {
	_init_.Initialize()

	if err := validateNewDataAzurermMobileNetworkAttachedDataNetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermMobileNetworkAttachedDataNetwork.DataAzurermMobileNetworkAttachedDataNetwork",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/data-sources/mobile_network_attached_data_network azurerm_mobile_network_attached_data_network} Data Source.
func NewDataAzurermMobileNetworkAttachedDataNetwork_Override(d DataAzurermMobileNetworkAttachedDataNetwork, scope constructs.Construct, id *string, config *DataAzurermMobileNetworkAttachedDataNetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermMobileNetworkAttachedDataNetwork.DataAzurermMobileNetworkAttachedDataNetwork",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork)SetMobileNetworkDataNetworkName(val *string) {
	if err := j.validateSetMobileNetworkDataNetworkNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobileNetworkDataNetworkName",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork)SetMobileNetworkPacketCoreDataPlaneId(val *string) {
	if err := j.validateSetMobileNetworkPacketCoreDataPlaneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobileNetworkPacketCoreDataPlaneId",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermMobileNetworkAttachedDataNetwork resource upon running "cdktf plan <stack-name>".
func DataAzurermMobileNetworkAttachedDataNetwork_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermMobileNetworkAttachedDataNetwork_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermMobileNetworkAttachedDataNetwork.DataAzurermMobileNetworkAttachedDataNetwork",
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
func DataAzurermMobileNetworkAttachedDataNetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermMobileNetworkAttachedDataNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermMobileNetworkAttachedDataNetwork.DataAzurermMobileNetworkAttachedDataNetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermMobileNetworkAttachedDataNetwork_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermMobileNetworkAttachedDataNetwork_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermMobileNetworkAttachedDataNetwork.DataAzurermMobileNetworkAttachedDataNetwork",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermMobileNetworkAttachedDataNetwork_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermMobileNetworkAttachedDataNetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermMobileNetworkAttachedDataNetwork.DataAzurermMobileNetworkAttachedDataNetwork",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermMobileNetworkAttachedDataNetwork_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermMobileNetworkAttachedDataNetwork.DataAzurermMobileNetworkAttachedDataNetwork",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) PutTimeouts(value *DataAzurermMobileNetworkAttachedDataNetworkTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMobileNetworkAttachedDataNetwork) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

