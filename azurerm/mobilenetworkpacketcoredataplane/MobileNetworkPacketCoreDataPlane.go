// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkpacketcoredataplane

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/mobilenetworkpacketcoredataplane/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/mobile_network_packet_core_data_plane azurerm_mobile_network_packet_core_data_plane}.
type MobileNetworkPacketCoreDataPlane interface {
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
	MobileNetworkPacketCoreControlPlaneId() *string
	SetMobileNetworkPacketCoreControlPlaneId(val *string)
	MobileNetworkPacketCoreControlPlaneIdInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MobileNetworkPacketCoreDataPlaneTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *MobileNetworkPacketCoreDataPlaneTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for MobileNetworkPacketCoreDataPlane
type jsiiProxy_MobileNetworkPacketCoreDataPlane struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) MobileNetworkPacketCoreControlPlaneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkPacketCoreControlPlaneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) MobileNetworkPacketCoreControlPlaneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkPacketCoreControlPlaneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) Timeouts() MobileNetworkPacketCoreDataPlaneTimeoutsOutputReference {
	var returns MobileNetworkPacketCoreDataPlaneTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) UserPlaneAccessIpv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) UserPlaneAccessIpv4AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) UserPlaneAccessIpv4Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) UserPlaneAccessIpv4GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) UserPlaneAccessIpv4Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4Subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) UserPlaneAccessIpv4SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessIpv4SubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) UserPlaneAccessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane) UserPlaneAccessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPlaneAccessNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/mobile_network_packet_core_data_plane azurerm_mobile_network_packet_core_data_plane} Resource.
func NewMobileNetworkPacketCoreDataPlane(scope constructs.Construct, id *string, config *MobileNetworkPacketCoreDataPlaneConfig) MobileNetworkPacketCoreDataPlane {
	_init_.Initialize()

	if err := validateNewMobileNetworkPacketCoreDataPlaneParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkPacketCoreDataPlane{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreDataPlane.MobileNetworkPacketCoreDataPlane",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/mobile_network_packet_core_data_plane azurerm_mobile_network_packet_core_data_plane} Resource.
func NewMobileNetworkPacketCoreDataPlane_Override(m MobileNetworkPacketCoreDataPlane, scope constructs.Construct, id *string, config *MobileNetworkPacketCoreDataPlaneConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreDataPlane.MobileNetworkPacketCoreDataPlane",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetMobileNetworkPacketCoreControlPlaneId(val *string) {
	if err := j.validateSetMobileNetworkPacketCoreControlPlaneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobileNetworkPacketCoreControlPlaneId",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetUserPlaneAccessIpv4Address(val *string) {
	if err := j.validateSetUserPlaneAccessIpv4AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPlaneAccessIpv4Address",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetUserPlaneAccessIpv4Gateway(val *string) {
	if err := j.validateSetUserPlaneAccessIpv4GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPlaneAccessIpv4Gateway",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetUserPlaneAccessIpv4Subnet(val *string) {
	if err := j.validateSetUserPlaneAccessIpv4SubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPlaneAccessIpv4Subnet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkPacketCoreDataPlane)SetUserPlaneAccessName(val *string) {
	if err := j.validateSetUserPlaneAccessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPlaneAccessName",
		val,
	)
}

// Generates CDKTF code for importing a MobileNetworkPacketCoreDataPlane resource upon running "cdktf plan <stack-name>".
func MobileNetworkPacketCoreDataPlane_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMobileNetworkPacketCoreDataPlane_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreDataPlane.MobileNetworkPacketCoreDataPlane",
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
func MobileNetworkPacketCoreDataPlane_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkPacketCoreDataPlane_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreDataPlane.MobileNetworkPacketCoreDataPlane",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MobileNetworkPacketCoreDataPlane_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkPacketCoreDataPlane_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreDataPlane.MobileNetworkPacketCoreDataPlane",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MobileNetworkPacketCoreDataPlane_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkPacketCoreDataPlane_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreDataPlane.MobileNetworkPacketCoreDataPlane",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MobileNetworkPacketCoreDataPlane_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mobileNetworkPacketCoreDataPlane.MobileNetworkPacketCoreDataPlane",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) PutTimeouts(value *MobileNetworkPacketCoreDataPlaneTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ResetUserPlaneAccessIpv4Address() {
	_jsii_.InvokeVoid(
		m,
		"resetUserPlaneAccessIpv4Address",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ResetUserPlaneAccessIpv4Gateway() {
	_jsii_.InvokeVoid(
		m,
		"resetUserPlaneAccessIpv4Gateway",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ResetUserPlaneAccessIpv4Subnet() {
	_jsii_.InvokeVoid(
		m,
		"resetUserPlaneAccessIpv4Subnet",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ResetUserPlaneAccessName() {
	_jsii_.InvokeVoid(
		m,
		"resetUserPlaneAccessName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkPacketCoreDataPlane) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

