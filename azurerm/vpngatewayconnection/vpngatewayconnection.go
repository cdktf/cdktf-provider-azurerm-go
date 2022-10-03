package vpngatewayconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/vpngatewayconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection azurerm_vpn_gateway_connection}.
type VpnGatewayConnection interface {
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	InternetSecurityEnabled() interface{}
	SetInternetSecurityEnabled(val interface{})
	InternetSecurityEnabledInput() interface{}
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
	RemoteVpnSiteId() *string
	SetRemoteVpnSiteId(val *string)
	RemoteVpnSiteIdInput() *string
	Routing() VpnGatewayConnectionRoutingOutputReference
	RoutingInput() *VpnGatewayConnectionRouting
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VpnGatewayConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrafficSelectorPolicy() VpnGatewayConnectionTrafficSelectorPolicyList
	TrafficSelectorPolicyInput() interface{}
	VpnGatewayId() *string
	SetVpnGatewayId(val *string)
	VpnGatewayIdInput() *string
	VpnLink() VpnGatewayConnectionVpnLinkList
	VpnLinkInput() interface{}
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
	PutRouting(value *VpnGatewayConnectionRouting)
	PutTimeouts(value *VpnGatewayConnectionTimeouts)
	PutTrafficSelectorPolicy(value interface{})
	PutVpnLink(value interface{})
	ResetId()
	ResetInternetSecurityEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRouting()
	ResetTimeouts()
	ResetTrafficSelectorPolicy()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpnGatewayConnection
type jsiiProxy_VpnGatewayConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpnGatewayConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) InternetSecurityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetSecurityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) InternetSecurityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetSecurityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) RemoteVpnSiteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVpnSiteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) RemoteVpnSiteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVpnSiteIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Routing() VpnGatewayConnectionRoutingOutputReference {
	var returns VpnGatewayConnectionRoutingOutputReference
	_jsii_.Get(
		j,
		"routing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) RoutingInput() *VpnGatewayConnectionRouting {
	var returns *VpnGatewayConnectionRouting
	_jsii_.Get(
		j,
		"routingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Timeouts() VpnGatewayConnectionTimeoutsOutputReference {
	var returns VpnGatewayConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TrafficSelectorPolicy() VpnGatewayConnectionTrafficSelectorPolicyList {
	var returns VpnGatewayConnectionTrafficSelectorPolicyList
	_jsii_.Get(
		j,
		"trafficSelectorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TrafficSelectorPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficSelectorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) VpnGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) VpnLink() VpnGatewayConnectionVpnLinkList {
	var returns VpnGatewayConnectionVpnLinkList
	_jsii_.Get(
		j,
		"vpnLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) VpnLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpnLinkInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection azurerm_vpn_gateway_connection} Resource.
func NewVpnGatewayConnection(scope constructs.Construct, id *string, config *VpnGatewayConnectionConfig) VpnGatewayConnection {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnection{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection azurerm_vpn_gateway_connection} Resource.
func NewVpnGatewayConnection_Override(v VpnGatewayConnection, scope constructs.Construct, id *string, config *VpnGatewayConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetInternetSecurityEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"internetSecurityEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetRemoteVpnSiteId(val *string) {
	_jsii_.Set(
		j,
		"remoteVpnSiteId",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection) SetVpnGatewayId(val *string) {
	_jsii_.Set(
		j,
		"vpnGatewayId",
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
func VpnGatewayConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpnGatewayConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) PutRouting(value *VpnGatewayConnectionRouting) {
	_jsii_.InvokeVoid(
		v,
		"putRouting",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) PutTimeouts(value *VpnGatewayConnectionTimeouts) {
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) PutTrafficSelectorPolicy(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putTrafficSelectorPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) PutVpnLink(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putVpnLink",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetInternetSecurityEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetInternetSecurityEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetRouting() {
	_jsii_.InvokeVoid(
		v,
		"resetRouting",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetTrafficSelectorPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetTrafficSelectorPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#name VpnGatewayConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#remote_vpn_site_id VpnGatewayConnection#remote_vpn_site_id}.
	RemoteVpnSiteId *string `field:"required" json:"remoteVpnSiteId" yaml:"remoteVpnSiteId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#vpn_gateway_id VpnGatewayConnection#vpn_gateway_id}.
	VpnGatewayId *string `field:"required" json:"vpnGatewayId" yaml:"vpnGatewayId"`
	// vpn_link block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#vpn_link VpnGatewayConnection#vpn_link}
	VpnLink interface{} `field:"required" json:"vpnLink" yaml:"vpnLink"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#id VpnGatewayConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#internet_security_enabled VpnGatewayConnection#internet_security_enabled}.
	InternetSecurityEnabled interface{} `field:"optional" json:"internetSecurityEnabled" yaml:"internetSecurityEnabled"`
	// routing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#routing VpnGatewayConnection#routing}
	Routing *VpnGatewayConnectionRouting `field:"optional" json:"routing" yaml:"routing"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#timeouts VpnGatewayConnection#timeouts}
	Timeouts *VpnGatewayConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_selector_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#traffic_selector_policy VpnGatewayConnection#traffic_selector_policy}
	TrafficSelectorPolicy interface{} `field:"optional" json:"trafficSelectorPolicy" yaml:"trafficSelectorPolicy"`
}

type VpnGatewayConnectionRouting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#associated_route_table VpnGatewayConnection#associated_route_table}.
	AssociatedRouteTable *string `field:"required" json:"associatedRouteTable" yaml:"associatedRouteTable"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#propagated_route_table VpnGatewayConnection#propagated_route_table}
	PropagatedRouteTable *VpnGatewayConnectionRoutingPropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
}

type VpnGatewayConnectionRoutingOutputReference interface {
	cdktf.ComplexObject
	AssociatedRouteTable() *string
	SetAssociatedRouteTable(val *string)
	AssociatedRouteTableInput() *string
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
	InternalValue() *VpnGatewayConnectionRouting
	SetInternalValue(val *VpnGatewayConnectionRouting)
	PropagatedRouteTable() VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference
	PropagatedRouteTableInput() *VpnGatewayConnectionRoutingPropagatedRouteTable
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
	PutPropagatedRouteTable(value *VpnGatewayConnectionRoutingPropagatedRouteTable)
	ResetPropagatedRouteTable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionRoutingOutputReference
type jsiiProxy_VpnGatewayConnectionRoutingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) AssociatedRouteTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) AssociatedRouteTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) InternalValue() *VpnGatewayConnectionRouting {
	var returns *VpnGatewayConnectionRouting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) PropagatedRouteTable() VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference {
	var returns VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference
	_jsii_.Get(
		j,
		"propagatedRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) PropagatedRouteTableInput() *VpnGatewayConnectionRoutingPropagatedRouteTable {
	var returns *VpnGatewayConnectionRoutingPropagatedRouteTable
	_jsii_.Get(
		j,
		"propagatedRouteTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionRoutingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnGatewayConnectionRoutingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionRoutingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionRoutingOutputReference_Override(v VpnGatewayConnectionRoutingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) SetAssociatedRouteTable(val *string) {
	_jsii_.Set(
		j,
		"associatedRouteTable",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) SetInternalValue(val *VpnGatewayConnectionRouting) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) PutPropagatedRouteTable(value *VpnGatewayConnectionRoutingPropagatedRouteTable) {
	_jsii_.InvokeVoid(
		v,
		"putPropagatedRouteTable",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ResetPropagatedRouteTable() {
	_jsii_.InvokeVoid(
		v,
		"resetPropagatedRouteTable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionRoutingPropagatedRouteTable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#route_table_ids VpnGatewayConnection#route_table_ids}.
	RouteTableIds *[]*string `field:"required" json:"routeTableIds" yaml:"routeTableIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#labels VpnGatewayConnection#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

type VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference interface {
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
	InternalValue() *VpnGatewayConnectionRoutingPropagatedRouteTable
	SetInternalValue(val *VpnGatewayConnectionRoutingPropagatedRouteTable)
	Labels() *[]*string
	SetLabels(val *[]*string)
	LabelsInput() *[]*string
	RouteTableIds() *[]*string
	SetRouteTableIds(val *[]*string)
	RouteTableIdsInput() *[]*string
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
	ResetLabels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference
type jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) InternalValue() *VpnGatewayConnectionRoutingPropagatedRouteTable {
	var returns *VpnGatewayConnectionRoutingPropagatedRouteTable
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) LabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) RouteTableIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTableIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) RouteTableIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTableIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionRoutingPropagatedRouteTableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionRoutingPropagatedRouteTableOutputReference_Override(v VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) SetInternalValue(val *VpnGatewayConnectionRoutingPropagatedRouteTable) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) SetLabels(val *[]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) SetRouteTableIds(val *[]*string) {
	_jsii_.Set(
		j,
		"routeTableIds",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		v,
		"resetLabels",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionRoutingPropagatedRouteTableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#create VpnGatewayConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#delete VpnGatewayConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#read VpnGatewayConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#update VpnGatewayConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type VpnGatewayConnectionTimeoutsOutputReference interface {
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

// The jsii proxy struct for VpnGatewayConnectionTimeoutsOutputReference
type jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnGatewayConnectionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionTimeoutsOutputReference_Override(v VpnGatewayConnectionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		v,
		"resetCreate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		v,
		"resetDelete",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		v,
		"resetRead",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		v,
		"resetUpdate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionTrafficSelectorPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#local_address_ranges VpnGatewayConnection#local_address_ranges}.
	LocalAddressRanges *[]*string `field:"required" json:"localAddressRanges" yaml:"localAddressRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#remote_address_ranges VpnGatewayConnection#remote_address_ranges}.
	RemoteAddressRanges *[]*string `field:"required" json:"remoteAddressRanges" yaml:"remoteAddressRanges"`
}

type VpnGatewayConnectionTrafficSelectorPolicyList interface {
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
	Get(index *float64) VpnGatewayConnectionTrafficSelectorPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionTrafficSelectorPolicyList
type jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionTrafficSelectorPolicyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VpnGatewayConnectionTrafficSelectorPolicyList {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionTrafficSelectorPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionTrafficSelectorPolicyList_Override(v VpnGatewayConnectionTrafficSelectorPolicyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionTrafficSelectorPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) Get(index *float64) VpnGatewayConnectionTrafficSelectorPolicyOutputReference {
	var returns VpnGatewayConnectionTrafficSelectorPolicyOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionTrafficSelectorPolicyOutputReference interface {
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
	LocalAddressRanges() *[]*string
	SetLocalAddressRanges(val *[]*string)
	LocalAddressRangesInput() *[]*string
	RemoteAddressRanges() *[]*string
	SetRemoteAddressRanges(val *[]*string)
	RemoteAddressRangesInput() *[]*string
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

// The jsii proxy struct for VpnGatewayConnectionTrafficSelectorPolicyOutputReference
type jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) LocalAddressRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localAddressRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) LocalAddressRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localAddressRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) RemoteAddressRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteAddressRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) RemoteAddressRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteAddressRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionTrafficSelectorPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpnGatewayConnectionTrafficSelectorPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionTrafficSelectorPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionTrafficSelectorPolicyOutputReference_Override(v VpnGatewayConnectionTrafficSelectorPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionTrafficSelectorPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) SetLocalAddressRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"localAddressRanges",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) SetRemoteAddressRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"remoteAddressRanges",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionTrafficSelectorPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionVpnLink struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#name VpnGatewayConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#vpn_site_link_id VpnGatewayConnection#vpn_site_link_id}.
	VpnSiteLinkId *string `field:"required" json:"vpnSiteLinkId" yaml:"vpnSiteLinkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#bandwidth_mbps VpnGatewayConnection#bandwidth_mbps}.
	BandwidthMbps *float64 `field:"optional" json:"bandwidthMbps" yaml:"bandwidthMbps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#bgp_enabled VpnGatewayConnection#bgp_enabled}.
	BgpEnabled interface{} `field:"optional" json:"bgpEnabled" yaml:"bgpEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#connection_mode VpnGatewayConnection#connection_mode}.
	ConnectionMode *string `field:"optional" json:"connectionMode" yaml:"connectionMode"`
	// custom_bgp_address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#custom_bgp_address VpnGatewayConnection#custom_bgp_address}
	CustomBgpAddress interface{} `field:"optional" json:"customBgpAddress" yaml:"customBgpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#egress_nat_rule_ids VpnGatewayConnection#egress_nat_rule_ids}.
	EgressNatRuleIds *[]*string `field:"optional" json:"egressNatRuleIds" yaml:"egressNatRuleIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ingress_nat_rule_ids VpnGatewayConnection#ingress_nat_rule_ids}.
	IngressNatRuleIds *[]*string `field:"optional" json:"ingressNatRuleIds" yaml:"ingressNatRuleIds"`
	// ipsec_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ipsec_policy VpnGatewayConnection#ipsec_policy}
	IpsecPolicy interface{} `field:"optional" json:"ipsecPolicy" yaml:"ipsecPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#local_azure_ip_address_enabled VpnGatewayConnection#local_azure_ip_address_enabled}.
	LocalAzureIpAddressEnabled interface{} `field:"optional" json:"localAzureIpAddressEnabled" yaml:"localAzureIpAddressEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#policy_based_traffic_selector_enabled VpnGatewayConnection#policy_based_traffic_selector_enabled}.
	PolicyBasedTrafficSelectorEnabled interface{} `field:"optional" json:"policyBasedTrafficSelectorEnabled" yaml:"policyBasedTrafficSelectorEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#protocol VpnGatewayConnection#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ratelimit_enabled VpnGatewayConnection#ratelimit_enabled}.
	RatelimitEnabled interface{} `field:"optional" json:"ratelimitEnabled" yaml:"ratelimitEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#route_weight VpnGatewayConnection#route_weight}.
	RouteWeight *float64 `field:"optional" json:"routeWeight" yaml:"routeWeight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#shared_key VpnGatewayConnection#shared_key}.
	SharedKey *string `field:"optional" json:"sharedKey" yaml:"sharedKey"`
}

type VpnGatewayConnectionVpnLinkCustomBgpAddress struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ip_address VpnGatewayConnection#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ip_configuration_id VpnGatewayConnection#ip_configuration_id}.
	IpConfigurationId *string `field:"required" json:"ipConfigurationId" yaml:"ipConfigurationId"`
}

type VpnGatewayConnectionVpnLinkCustomBgpAddressList interface {
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
	Get(index *float64) VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionVpnLinkCustomBgpAddressList
type jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionVpnLinkCustomBgpAddressList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VpnGatewayConnectionVpnLinkCustomBgpAddressList {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkCustomBgpAddressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionVpnLinkCustomBgpAddressList_Override(v VpnGatewayConnectionVpnLinkCustomBgpAddressList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkCustomBgpAddressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) Get(index *float64) VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference {
	var returns VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference interface {
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
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	IpConfigurationId() *string
	SetIpConfigurationId(val *string)
	IpConfigurationIdInput() *string
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

// The jsii proxy struct for VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference
type jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) IpConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) IpConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfigurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference_Override(v VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) SetIpAddress(val *string) {
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) SetIpConfigurationId(val *string) {
	_jsii_.Set(
		j,
		"ipConfigurationId",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkCustomBgpAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionVpnLinkIpsecPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#dh_group VpnGatewayConnection#dh_group}.
	DhGroup *string `field:"required" json:"dhGroup" yaml:"dhGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#encryption_algorithm VpnGatewayConnection#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"required" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ike_encryption_algorithm VpnGatewayConnection#ike_encryption_algorithm}.
	IkeEncryptionAlgorithm *string `field:"required" json:"ikeEncryptionAlgorithm" yaml:"ikeEncryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ike_integrity_algorithm VpnGatewayConnection#ike_integrity_algorithm}.
	IkeIntegrityAlgorithm *string `field:"required" json:"ikeIntegrityAlgorithm" yaml:"ikeIntegrityAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#integrity_algorithm VpnGatewayConnection#integrity_algorithm}.
	IntegrityAlgorithm *string `field:"required" json:"integrityAlgorithm" yaml:"integrityAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#pfs_group VpnGatewayConnection#pfs_group}.
	PfsGroup *string `field:"required" json:"pfsGroup" yaml:"pfsGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#sa_data_size_kb VpnGatewayConnection#sa_data_size_kb}.
	SaDataSizeKb *float64 `field:"required" json:"saDataSizeKb" yaml:"saDataSizeKb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#sa_lifetime_sec VpnGatewayConnection#sa_lifetime_sec}.
	SaLifetimeSec *float64 `field:"required" json:"saLifetimeSec" yaml:"saLifetimeSec"`
}

type VpnGatewayConnectionVpnLinkIpsecPolicyList interface {
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
	Get(index *float64) VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionVpnLinkIpsecPolicyList
type jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionVpnLinkIpsecPolicyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VpnGatewayConnectionVpnLinkIpsecPolicyList {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkIpsecPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionVpnLinkIpsecPolicyList_Override(v VpnGatewayConnectionVpnLinkIpsecPolicyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkIpsecPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) Get(index *float64) VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference {
	var returns VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference interface {
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
	DhGroup() *string
	SetDhGroup(val *string)
	DhGroupInput() *string
	EncryptionAlgorithm() *string
	SetEncryptionAlgorithm(val *string)
	EncryptionAlgorithmInput() *string
	// Experimental.
	Fqn() *string
	IkeEncryptionAlgorithm() *string
	SetIkeEncryptionAlgorithm(val *string)
	IkeEncryptionAlgorithmInput() *string
	IkeIntegrityAlgorithm() *string
	SetIkeIntegrityAlgorithm(val *string)
	IkeIntegrityAlgorithmInput() *string
	IntegrityAlgorithm() *string
	SetIntegrityAlgorithm(val *string)
	IntegrityAlgorithmInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PfsGroup() *string
	SetPfsGroup(val *string)
	PfsGroupInput() *string
	SaDataSizeKb() *float64
	SetSaDataSizeKb(val *float64)
	SaDataSizeKbInput() *float64
	SaLifetimeSec() *float64
	SetSaLifetimeSec(val *float64)
	SaLifetimeSecInput() *float64
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

// The jsii proxy struct for VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference
type jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) DhGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) DhGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IkeEncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IkeEncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IkeIntegrityAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrityAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IkeIntegrityAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrityAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IntegrityAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrityAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IntegrityAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrityAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) PfsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) PfsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SaDataSizeKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDataSizeKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SaDataSizeKbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDataSizeKbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SaLifetimeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SaLifetimeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionVpnLinkIpsecPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionVpnLinkIpsecPolicyOutputReference_Override(v VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetDhGroup(val *string) {
	_jsii_.Set(
		j,
		"dhGroup",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetEncryptionAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetIkeEncryptionAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"ikeEncryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetIkeIntegrityAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"ikeIntegrityAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetIntegrityAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"integrityAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetPfsGroup(val *string) {
	_jsii_.Set(
		j,
		"pfsGroup",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetSaDataSizeKb(val *float64) {
	_jsii_.Set(
		j,
		"saDataSizeKb",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetSaLifetimeSec(val *float64) {
	_jsii_.Set(
		j,
		"saLifetimeSec",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionVpnLinkList interface {
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
	Get(index *float64) VpnGatewayConnectionVpnLinkOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionVpnLinkList
type jsiiProxy_VpnGatewayConnectionVpnLinkList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionVpnLinkList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VpnGatewayConnectionVpnLinkList {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionVpnLinkList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionVpnLinkList_Override(v VpnGatewayConnectionVpnLinkList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkList) Get(index *float64) VpnGatewayConnectionVpnLinkOutputReference {
	var returns VpnGatewayConnectionVpnLinkOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VpnGatewayConnectionVpnLinkOutputReference interface {
	cdktf.ComplexObject
	BandwidthMbps() *float64
	SetBandwidthMbps(val *float64)
	BandwidthMbpsInput() *float64
	BgpEnabled() interface{}
	SetBgpEnabled(val interface{})
	BgpEnabledInput() interface{}
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
	ConnectionMode() *string
	SetConnectionMode(val *string)
	ConnectionModeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomBgpAddress() VpnGatewayConnectionVpnLinkCustomBgpAddressList
	CustomBgpAddressInput() interface{}
	EgressNatRuleIds() *[]*string
	SetEgressNatRuleIds(val *[]*string)
	EgressNatRuleIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	IngressNatRuleIds() *[]*string
	SetIngressNatRuleIds(val *[]*string)
	IngressNatRuleIdsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpsecPolicy() VpnGatewayConnectionVpnLinkIpsecPolicyList
	IpsecPolicyInput() interface{}
	LocalAzureIpAddressEnabled() interface{}
	SetLocalAzureIpAddressEnabled(val interface{})
	LocalAzureIpAddressEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PolicyBasedTrafficSelectorEnabled() interface{}
	SetPolicyBasedTrafficSelectorEnabled(val interface{})
	PolicyBasedTrafficSelectorEnabledInput() interface{}
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	RatelimitEnabled() interface{}
	SetRatelimitEnabled(val interface{})
	RatelimitEnabledInput() interface{}
	RouteWeight() *float64
	SetRouteWeight(val *float64)
	RouteWeightInput() *float64
	SharedKey() *string
	SetSharedKey(val *string)
	SharedKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpnSiteLinkId() *string
	SetVpnSiteLinkId(val *string)
	VpnSiteLinkIdInput() *string
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
	PutCustomBgpAddress(value interface{})
	PutIpsecPolicy(value interface{})
	ResetBandwidthMbps()
	ResetBgpEnabled()
	ResetConnectionMode()
	ResetCustomBgpAddress()
	ResetEgressNatRuleIds()
	ResetIngressNatRuleIds()
	ResetIpsecPolicy()
	ResetLocalAzureIpAddressEnabled()
	ResetPolicyBasedTrafficSelectorEnabled()
	ResetProtocol()
	ResetRatelimitEnabled()
	ResetRouteWeight()
	ResetSharedKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionVpnLinkOutputReference
type jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) BandwidthMbps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) BandwidthMbpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) BgpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) BgpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ConnectionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) CustomBgpAddress() VpnGatewayConnectionVpnLinkCustomBgpAddressList {
	var returns VpnGatewayConnectionVpnLinkCustomBgpAddressList
	_jsii_.Get(
		j,
		"customBgpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) CustomBgpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customBgpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) EgressNatRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) EgressNatRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) IngressNatRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressNatRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) IngressNatRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressNatRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) IpsecPolicy() VpnGatewayConnectionVpnLinkIpsecPolicyList {
	var returns VpnGatewayConnectionVpnLinkIpsecPolicyList
	_jsii_.Get(
		j,
		"ipsecPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) IpsecPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipsecPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) LocalAzureIpAddressEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAzureIpAddressEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) LocalAzureIpAddressEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAzureIpAddressEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) PolicyBasedTrafficSelectorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyBasedTrafficSelectorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) PolicyBasedTrafficSelectorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyBasedTrafficSelectorEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) RatelimitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ratelimitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) RatelimitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ratelimitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) RouteWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"routeWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) RouteWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"routeWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) VpnSiteLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnSiteLinkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) VpnSiteLinkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnSiteLinkIdInput",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionVpnLinkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpnGatewayConnectionVpnLinkOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionVpnLinkOutputReference_Override(v VpnGatewayConnectionVpnLinkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetBandwidthMbps(val *float64) {
	_jsii_.Set(
		j,
		"bandwidthMbps",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetBgpEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"bgpEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetConnectionMode(val *string) {
	_jsii_.Set(
		j,
		"connectionMode",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetEgressNatRuleIds(val *[]*string) {
	_jsii_.Set(
		j,
		"egressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetIngressNatRuleIds(val *[]*string) {
	_jsii_.Set(
		j,
		"ingressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetLocalAzureIpAddressEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"localAzureIpAddressEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetPolicyBasedTrafficSelectorEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"policyBasedTrafficSelectorEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetRatelimitEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"ratelimitEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetRouteWeight(val *float64) {
	_jsii_.Set(
		j,
		"routeWeight",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetSharedKey(val *string) {
	_jsii_.Set(
		j,
		"sharedKey",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SetVpnSiteLinkId(val *string) {
	_jsii_.Set(
		j,
		"vpnSiteLinkId",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) PutCustomBgpAddress(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putCustomBgpAddress",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) PutIpsecPolicy(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putIpsecPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetBandwidthMbps() {
	_jsii_.InvokeVoid(
		v,
		"resetBandwidthMbps",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetBgpEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetBgpEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetConnectionMode() {
	_jsii_.InvokeVoid(
		v,
		"resetConnectionMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetCustomBgpAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomBgpAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetEgressNatRuleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetEgressNatRuleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetIngressNatRuleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetIngressNatRuleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetIpsecPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetIpsecPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetLocalAzureIpAddressEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalAzureIpAddressEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetPolicyBasedTrafficSelectorEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetPolicyBasedTrafficSelectorEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetRatelimitEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetRatelimitEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetRouteWeight() {
	_jsii_.InvokeVoid(
		v,
		"resetRouteWeight",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetSharedKey() {
	_jsii_.InvokeVoid(
		v,
		"resetSharedKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

