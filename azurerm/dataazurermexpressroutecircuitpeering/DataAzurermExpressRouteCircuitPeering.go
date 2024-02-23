// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermexpressroutecircuitpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/dataazurermexpressroutecircuitpeering/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/data-sources/express_route_circuit_peering azurerm_express_route_circuit_peering}.
type DataAzurermExpressRouteCircuitPeering interface {
	cdktf.TerraformDataSource
	AzureAsn() *float64
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
	Ipv4Enabled() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PeerAsn() *float64
	PeeringType() *string
	SetPeeringType(val *string)
	PeeringTypeInput() *string
	PrimaryAzurePort() *string
	PrimaryPeerAddressPrefix() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RouteFilterId() *string
	SecondaryAzurePort() *string
	SecondaryPeerAddressPrefix() *string
	SharedKey() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermExpressRouteCircuitPeeringTimeoutsOutputReference
	TimeoutsInput() interface{}
	VlanId() *float64
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
	PutTimeouts(value *DataAzurermExpressRouteCircuitPeeringTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAzurermExpressRouteCircuitPeering
type jsiiProxy_DataAzurermExpressRouteCircuitPeering struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) AzureAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"azureAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ExpressRouteCircuitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ExpressRouteCircuitNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GatewayManagerEtag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayManagerEtag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) Ipv4Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipv4Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) PeerAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) PeeringType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) PeeringTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) PrimaryAzurePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAzurePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) PrimaryPeerAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPeerAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) RouteFilterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFilterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) SecondaryAzurePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAzurePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) SecondaryPeerAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPeerAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) SharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) Timeouts() DataAzurermExpressRouteCircuitPeeringTimeoutsOutputReference {
	var returns DataAzurermExpressRouteCircuitPeeringTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering) VlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/data-sources/express_route_circuit_peering azurerm_express_route_circuit_peering} Data Source.
func NewDataAzurermExpressRouteCircuitPeering(scope constructs.Construct, id *string, config *DataAzurermExpressRouteCircuitPeeringConfig) DataAzurermExpressRouteCircuitPeering {
	_init_.Initialize()

	if err := validateNewDataAzurermExpressRouteCircuitPeeringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermExpressRouteCircuitPeering{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermExpressRouteCircuitPeering.DataAzurermExpressRouteCircuitPeering",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/data-sources/express_route_circuit_peering azurerm_express_route_circuit_peering} Data Source.
func NewDataAzurermExpressRouteCircuitPeering_Override(d DataAzurermExpressRouteCircuitPeering, scope constructs.Construct, id *string, config *DataAzurermExpressRouteCircuitPeeringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermExpressRouteCircuitPeering.DataAzurermExpressRouteCircuitPeering",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetExpressRouteCircuitName(val *string) {
	if err := j.validateSetExpressRouteCircuitNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressRouteCircuitName",
		val,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetPeeringType(val *string) {
	if err := j.validateSetPeeringTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peeringType",
		val,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermExpressRouteCircuitPeering)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermExpressRouteCircuitPeering resource upon running "cdktf plan <stack-name>".
func DataAzurermExpressRouteCircuitPeering_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermExpressRouteCircuitPeering_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermExpressRouteCircuitPeering.DataAzurermExpressRouteCircuitPeering",
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
func DataAzurermExpressRouteCircuitPeering_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermExpressRouteCircuitPeering_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermExpressRouteCircuitPeering.DataAzurermExpressRouteCircuitPeering",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermExpressRouteCircuitPeering_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermExpressRouteCircuitPeering_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermExpressRouteCircuitPeering.DataAzurermExpressRouteCircuitPeering",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermExpressRouteCircuitPeering_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermExpressRouteCircuitPeering_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermExpressRouteCircuitPeering.DataAzurermExpressRouteCircuitPeering",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermExpressRouteCircuitPeering_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermExpressRouteCircuitPeering.DataAzurermExpressRouteCircuitPeering",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) PutTimeouts(value *DataAzurermExpressRouteCircuitPeeringTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermExpressRouteCircuitPeering) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

