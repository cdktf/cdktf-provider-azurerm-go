// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databricksvirtualnetworkpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/databricksvirtualnetworkpeering/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/databricks_virtual_network_peering azurerm_databricks_virtual_network_peering}.
type DatabricksVirtualNetworkPeering interface {
	cdktf.TerraformResource
	AddressSpacePrefixes() *[]*string
	AllowForwardedTraffic() interface{}
	SetAllowForwardedTraffic(val interface{})
	AllowForwardedTrafficInput() interface{}
	AllowGatewayTransit() interface{}
	SetAllowGatewayTransit(val interface{})
	AllowGatewayTransitInput() interface{}
	AllowVirtualNetworkAccess() interface{}
	SetAllowVirtualNetworkAccess(val interface{})
	AllowVirtualNetworkAccessInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	RemoteAddressSpacePrefixes() *[]*string
	SetRemoteAddressSpacePrefixes(val *[]*string)
	RemoteAddressSpacePrefixesInput() *[]*string
	RemoteVirtualNetworkId() *string
	SetRemoteVirtualNetworkId(val *string)
	RemoteVirtualNetworkIdInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DatabricksVirtualNetworkPeeringTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseRemoteGateways() interface{}
	SetUseRemoteGateways(val interface{})
	UseRemoteGatewaysInput() interface{}
	VirtualNetworkId() *string
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	PutTimeouts(value *DatabricksVirtualNetworkPeeringTimeouts)
	ResetAllowForwardedTraffic()
	ResetAllowGatewayTransit()
	ResetAllowVirtualNetworkAccess()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetUseRemoteGateways()
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

// The jsii proxy struct for DatabricksVirtualNetworkPeering
type jsiiProxy_DatabricksVirtualNetworkPeering struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) AddressSpacePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressSpacePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) AllowForwardedTraffic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForwardedTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) AllowForwardedTrafficInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForwardedTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) AllowGatewayTransit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGatewayTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) AllowGatewayTransitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGatewayTransitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) AllowVirtualNetworkAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVirtualNetworkAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) AllowVirtualNetworkAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVirtualNetworkAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) RemoteAddressSpacePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteAddressSpacePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) RemoteAddressSpacePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteAddressSpacePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) RemoteVirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVirtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) RemoteVirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVirtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) Timeouts() DatabricksVirtualNetworkPeeringTimeoutsOutputReference {
	var returns DatabricksVirtualNetworkPeeringTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) UseRemoteGateways() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRemoteGateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) UseRemoteGatewaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRemoteGatewaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/databricks_virtual_network_peering azurerm_databricks_virtual_network_peering} Resource.
func NewDatabricksVirtualNetworkPeering(scope constructs.Construct, id *string, config *DatabricksVirtualNetworkPeeringConfig) DatabricksVirtualNetworkPeering {
	_init_.Initialize()

	if err := validateNewDatabricksVirtualNetworkPeeringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabricksVirtualNetworkPeering{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksVirtualNetworkPeering.DatabricksVirtualNetworkPeering",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/databricks_virtual_network_peering azurerm_databricks_virtual_network_peering} Resource.
func NewDatabricksVirtualNetworkPeering_Override(d DatabricksVirtualNetworkPeering, scope constructs.Construct, id *string, config *DatabricksVirtualNetworkPeeringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksVirtualNetworkPeering.DatabricksVirtualNetworkPeering",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetAllowForwardedTraffic(val interface{}) {
	if err := j.validateSetAllowForwardedTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForwardedTraffic",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetAllowGatewayTransit(val interface{}) {
	if err := j.validateSetAllowGatewayTransitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGatewayTransit",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetAllowVirtualNetworkAccess(val interface{}) {
	if err := j.validateSetAllowVirtualNetworkAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVirtualNetworkAccess",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetRemoteAddressSpacePrefixes(val *[]*string) {
	if err := j.validateSetRemoteAddressSpacePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteAddressSpacePrefixes",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetRemoteVirtualNetworkId(val *string) {
	if err := j.validateSetRemoteVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteVirtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetUseRemoteGateways(val interface{}) {
	if err := j.validateSetUseRemoteGatewaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRemoteGateways",
		val,
	)
}

func (j *jsiiProxy_DatabricksVirtualNetworkPeering)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTF code for importing a DatabricksVirtualNetworkPeering resource upon running "cdktf plan <stack-name>".
func DatabricksVirtualNetworkPeering_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabricksVirtualNetworkPeering_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.databricksVirtualNetworkPeering.DatabricksVirtualNetworkPeering",
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
func DatabricksVirtualNetworkPeering_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksVirtualNetworkPeering_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.databricksVirtualNetworkPeering.DatabricksVirtualNetworkPeering",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabricksVirtualNetworkPeering_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksVirtualNetworkPeering_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.databricksVirtualNetworkPeering.DatabricksVirtualNetworkPeering",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabricksVirtualNetworkPeering_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksVirtualNetworkPeering_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.databricksVirtualNetworkPeering.DatabricksVirtualNetworkPeering",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabricksVirtualNetworkPeering_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.databricksVirtualNetworkPeering.DatabricksVirtualNetworkPeering",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) PutTimeouts(value *DatabricksVirtualNetworkPeeringTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ResetAllowForwardedTraffic() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowForwardedTraffic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ResetAllowGatewayTransit() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowGatewayTransit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ResetAllowVirtualNetworkAccess() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowVirtualNetworkAccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ResetUseRemoteGateways() {
	_jsii_.InvokeVoid(
		d,
		"resetUseRemoteGateways",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksVirtualNetworkPeering) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

