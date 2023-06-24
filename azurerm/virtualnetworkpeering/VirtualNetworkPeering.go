package virtualnetworkpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/virtualnetworkpeering/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/virtual_network_peering azurerm_virtual_network_peering}.
type VirtualNetworkPeering interface {
	cdktf.TerraformResource
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
	Timeouts() VirtualNetworkPeeringTimeoutsOutputReference
	TimeoutsInput() interface{}
	Triggers() *map[string]*string
	SetTriggers(val *map[string]*string)
	TriggersInput() *map[string]*string
	UseRemoteGateways() interface{}
	SetUseRemoteGateways(val interface{})
	UseRemoteGatewaysInput() interface{}
	VirtualNetworkName() *string
	SetVirtualNetworkName(val *string)
	VirtualNetworkNameInput() *string
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
	PutTimeouts(value *VirtualNetworkPeeringTimeouts)
	ResetAllowForwardedTraffic()
	ResetAllowGatewayTransit()
	ResetAllowVirtualNetworkAccess()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetTriggers()
	ResetUseRemoteGateways()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VirtualNetworkPeering
type jsiiProxy_VirtualNetworkPeering struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualNetworkPeering) AllowForwardedTraffic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForwardedTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) AllowForwardedTrafficInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForwardedTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) AllowGatewayTransit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGatewayTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) AllowGatewayTransitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGatewayTransitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) AllowVirtualNetworkAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVirtualNetworkAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) AllowVirtualNetworkAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVirtualNetworkAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) RemoteVirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVirtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) RemoteVirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVirtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Timeouts() VirtualNetworkPeeringTimeoutsOutputReference {
	var returns VirtualNetworkPeeringTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) UseRemoteGateways() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRemoteGateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) UseRemoteGatewaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRemoteGatewaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) VirtualNetworkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkPeering) VirtualNetworkNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/virtual_network_peering azurerm_virtual_network_peering} Resource.
func NewVirtualNetworkPeering(scope constructs.Construct, id *string, config *VirtualNetworkPeeringConfig) VirtualNetworkPeering {
	_init_.Initialize()

	if err := validateNewVirtualNetworkPeeringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualNetworkPeering{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkPeering.VirtualNetworkPeering",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/virtual_network_peering azurerm_virtual_network_peering} Resource.
func NewVirtualNetworkPeering_Override(v VirtualNetworkPeering, scope constructs.Construct, id *string, config *VirtualNetworkPeeringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkPeering.VirtualNetworkPeering",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetAllowForwardedTraffic(val interface{}) {
	if err := j.validateSetAllowForwardedTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForwardedTraffic",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetAllowGatewayTransit(val interface{}) {
	if err := j.validateSetAllowGatewayTransitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGatewayTransit",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetAllowVirtualNetworkAccess(val interface{}) {
	if err := j.validateSetAllowVirtualNetworkAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVirtualNetworkAccess",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetRemoteVirtualNetworkId(val *string) {
	if err := j.validateSetRemoteVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteVirtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetUseRemoteGateways(val interface{}) {
	if err := j.validateSetUseRemoteGatewaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRemoteGateways",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkPeering)SetVirtualNetworkName(val *string) {
	if err := j.validateSetVirtualNetworkNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkName",
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
func VirtualNetworkPeering_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkPeering_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkPeering.VirtualNetworkPeering",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetworkPeering_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkPeering_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkPeering.VirtualNetworkPeering",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetworkPeering_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkPeering_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualNetworkPeering.VirtualNetworkPeering",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualNetworkPeering_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualNetworkPeering.VirtualNetworkPeering",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualNetworkPeering) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualNetworkPeering) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkPeering) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualNetworkPeering) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualNetworkPeering) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualNetworkPeering) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualNetworkPeering) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualNetworkPeering) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualNetworkPeering) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualNetworkPeering) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkPeering) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) PutTimeouts(value *VirtualNetworkPeeringTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) ResetAllowForwardedTraffic() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowForwardedTraffic",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) ResetAllowGatewayTransit() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowGatewayTransit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) ResetAllowVirtualNetworkAccess() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowVirtualNetworkAccess",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) ResetTriggers() {
	_jsii_.InvokeVoid(
		v,
		"resetTriggers",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) ResetUseRemoteGateways() {
	_jsii_.InvokeVoid(
		v,
		"resetUseRemoteGateways",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkPeering) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkPeering) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkPeering) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkPeering) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

