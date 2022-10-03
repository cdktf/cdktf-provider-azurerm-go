package pointtositevpngateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/pointtositevpngateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway azurerm_point_to_site_vpn_gateway}.
type PointToSiteVpnGateway interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionConfiguration() PointToSiteVpnGatewayConnectionConfigurationOutputReference
	ConnectionConfigurationInput() *PointToSiteVpnGatewayConnectionConfiguration
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
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ScaleUnit() *float64
	SetScaleUnit(val *float64)
	ScaleUnitInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PointToSiteVpnGatewayTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualHubId() *string
	SetVirtualHubId(val *string)
	VirtualHubIdInput() *string
	VpnServerConfigurationId() *string
	SetVpnServerConfigurationId(val *string)
	VpnServerConfigurationIdInput() *string
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
	PutConnectionConfiguration(value *PointToSiteVpnGatewayConnectionConfiguration)
	PutTimeouts(value *PointToSiteVpnGatewayTimeouts)
	ResetDnsServers()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
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

// The jsii proxy struct for PointToSiteVpnGateway
type jsiiProxy_PointToSiteVpnGateway struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PointToSiteVpnGateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) ConnectionConfiguration() PointToSiteVpnGatewayConnectionConfigurationOutputReference {
	var returns PointToSiteVpnGatewayConnectionConfigurationOutputReference
	_jsii_.Get(
		j,
		"connectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) ConnectionConfigurationInput() *PointToSiteVpnGatewayConnectionConfiguration {
	var returns *PointToSiteVpnGatewayConnectionConfiguration
	_jsii_.Get(
		j,
		"connectionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) ScaleUnit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) ScaleUnitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) Timeouts() PointToSiteVpnGatewayTimeoutsOutputReference {
	var returns PointToSiteVpnGatewayTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) VirtualHubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) VirtualHubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHubIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) VpnServerConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnServerConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGateway) VpnServerConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnServerConfigurationIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway azurerm_point_to_site_vpn_gateway} Resource.
func NewPointToSiteVpnGateway(scope constructs.Construct, id *string, config *PointToSiteVpnGatewayConfig) PointToSiteVpnGateway {
	_init_.Initialize()

	j := jsiiProxy_PointToSiteVpnGateway{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway azurerm_point_to_site_vpn_gateway} Resource.
func NewPointToSiteVpnGateway_Override(p PointToSiteVpnGateway, scope constructs.Construct, id *string, config *PointToSiteVpnGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGateway",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetDnsServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetScaleUnit(val *float64) {
	_jsii_.Set(
		j,
		"scaleUnit",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetVirtualHubId(val *string) {
	_jsii_.Set(
		j,
		"virtualHubId",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGateway) SetVpnServerConfigurationId(val *string) {
	_jsii_.Set(
		j,
		"vpnServerConfigurationId",
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
func PointToSiteVpnGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PointToSiteVpnGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) PutConnectionConfiguration(value *PointToSiteVpnGatewayConnectionConfiguration) {
	_jsii_.InvokeVoid(
		p,
		"putConnectionConfiguration",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) PutTimeouts(value *PointToSiteVpnGatewayTimeouts) {
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) ResetDnsServers() {
	_jsii_.InvokeVoid(
		p,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PointToSiteVpnGatewayConfig struct {
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
	// connection_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#connection_configuration PointToSiteVpnGateway#connection_configuration}
	ConnectionConfiguration *PointToSiteVpnGatewayConnectionConfiguration `field:"required" json:"connectionConfiguration" yaml:"connectionConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#location PointToSiteVpnGateway#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#name PointToSiteVpnGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#resource_group_name PointToSiteVpnGateway#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#scale_unit PointToSiteVpnGateway#scale_unit}.
	ScaleUnit *float64 `field:"required" json:"scaleUnit" yaml:"scaleUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#virtual_hub_id PointToSiteVpnGateway#virtual_hub_id}.
	VirtualHubId *string `field:"required" json:"virtualHubId" yaml:"virtualHubId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#vpn_server_configuration_id PointToSiteVpnGateway#vpn_server_configuration_id}.
	VpnServerConfigurationId *string `field:"required" json:"vpnServerConfigurationId" yaml:"vpnServerConfigurationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#dns_servers PointToSiteVpnGateway#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#id PointToSiteVpnGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#tags PointToSiteVpnGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#timeouts PointToSiteVpnGateway#timeouts}
	Timeouts *PointToSiteVpnGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type PointToSiteVpnGatewayConnectionConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#name PointToSiteVpnGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// vpn_client_address_pool block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#vpn_client_address_pool PointToSiteVpnGateway#vpn_client_address_pool}
	VpnClientAddressPool *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool `field:"required" json:"vpnClientAddressPool" yaml:"vpnClientAddressPool"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#internet_security_enabled PointToSiteVpnGateway#internet_security_enabled}.
	InternetSecurityEnabled interface{} `field:"optional" json:"internetSecurityEnabled" yaml:"internetSecurityEnabled"`
	// route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#route PointToSiteVpnGateway#route}
	Route *PointToSiteVpnGatewayConnectionConfigurationRoute `field:"optional" json:"route" yaml:"route"`
}

type PointToSiteVpnGatewayConnectionConfigurationOutputReference interface {
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
	InternalValue() *PointToSiteVpnGatewayConnectionConfiguration
	SetInternalValue(val *PointToSiteVpnGatewayConnectionConfiguration)
	InternetSecurityEnabled() interface{}
	SetInternetSecurityEnabled(val interface{})
	InternetSecurityEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Route() PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference
	RouteInput() *PointToSiteVpnGatewayConnectionConfigurationRoute
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpnClientAddressPool() PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference
	VpnClientAddressPoolInput() *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool
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
	PutRoute(value *PointToSiteVpnGatewayConnectionConfigurationRoute)
	PutVpnClientAddressPool(value *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool)
	ResetInternetSecurityEnabled()
	ResetRoute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PointToSiteVpnGatewayConnectionConfigurationOutputReference
type jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InternalValue() *PointToSiteVpnGatewayConnectionConfiguration {
	var returns *PointToSiteVpnGatewayConnectionConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InternetSecurityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetSecurityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InternetSecurityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetSecurityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) Route() PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference {
	var returns PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) RouteInput() *PointToSiteVpnGatewayConnectionConfigurationRoute {
	var returns *PointToSiteVpnGatewayConnectionConfigurationRoute
	_jsii_.Get(
		j,
		"routeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) VpnClientAddressPool() PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference {
	var returns PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference
	_jsii_.Get(
		j,
		"vpnClientAddressPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) VpnClientAddressPoolInput() *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool {
	var returns *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool
	_jsii_.Get(
		j,
		"vpnClientAddressPoolInput",
		&returns,
	)
	return returns
}


func NewPointToSiteVpnGatewayConnectionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PointToSiteVpnGatewayConnectionConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPointToSiteVpnGatewayConnectionConfigurationOutputReference_Override(p PointToSiteVpnGatewayConnectionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) SetInternalValue(val *PointToSiteVpnGatewayConnectionConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) SetInternetSecurityEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"internetSecurityEnabled",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) PutRoute(value *PointToSiteVpnGatewayConnectionConfigurationRoute) {
	_jsii_.InvokeVoid(
		p,
		"putRoute",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) PutVpnClientAddressPool(value *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool) {
	_jsii_.InvokeVoid(
		p,
		"putVpnClientAddressPool",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ResetInternetSecurityEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetInternetSecurityEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ResetRoute() {
	_jsii_.InvokeVoid(
		p,
		"resetRoute",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PointToSiteVpnGatewayConnectionConfigurationRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#associated_route_table_id PointToSiteVpnGateway#associated_route_table_id}.
	AssociatedRouteTableId *string `field:"required" json:"associatedRouteTableId" yaml:"associatedRouteTableId"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#propagated_route_table PointToSiteVpnGateway#propagated_route_table}
	PropagatedRouteTable *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
}

type PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference interface {
	cdktf.ComplexObject
	AssociatedRouteTableId() *string
	SetAssociatedRouteTableId(val *string)
	AssociatedRouteTableIdInput() *string
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
	InternalValue() *PointToSiteVpnGatewayConnectionConfigurationRoute
	SetInternalValue(val *PointToSiteVpnGatewayConnectionConfigurationRoute)
	PropagatedRouteTable() PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference
	PropagatedRouteTableInput() *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable
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
	PutPropagatedRouteTable(value *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable)
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

// The jsii proxy struct for PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference
type jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) AssociatedRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) AssociatedRouteTableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedRouteTableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) InternalValue() *PointToSiteVpnGatewayConnectionConfigurationRoute {
	var returns *PointToSiteVpnGatewayConnectionConfigurationRoute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) PropagatedRouteTable() PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference {
	var returns PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference
	_jsii_.Get(
		j,
		"propagatedRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) PropagatedRouteTableInput() *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable {
	var returns *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable
	_jsii_.Get(
		j,
		"propagatedRouteTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPointToSiteVpnGatewayConnectionConfigurationRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPointToSiteVpnGatewayConnectionConfigurationRouteOutputReference_Override(p PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) SetAssociatedRouteTableId(val *string) {
	_jsii_.Set(
		j,
		"associatedRouteTableId",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) SetInternalValue(val *PointToSiteVpnGatewayConnectionConfigurationRoute) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) PutPropagatedRouteTable(value *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable) {
	_jsii_.InvokeVoid(
		p,
		"putPropagatedRouteTable",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ResetPropagatedRouteTable() {
	_jsii_.InvokeVoid(
		p,
		"resetPropagatedRouteTable",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#ids PointToSiteVpnGateway#ids}.
	Ids *[]*string `field:"required" json:"ids" yaml:"ids"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#labels PointToSiteVpnGateway#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

type PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference interface {
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
	Ids() *[]*string
	SetIds(val *[]*string)
	IdsInput() *[]*string
	InternalValue() *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable
	SetInternalValue(val *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable)
	Labels() *[]*string
	SetLabels(val *[]*string)
	LabelsInput() *[]*string
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

// The jsii proxy struct for PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference
type jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) Ids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) IdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) InternalValue() *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable {
	var returns *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) LabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference_Override(p PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) SetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"ids",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) SetInternalValue(val *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) SetLabels(val *[]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		p,
		"resetLabels",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#address_prefixes PointToSiteVpnGateway#address_prefixes}.
	AddressPrefixes *[]*string `field:"required" json:"addressPrefixes" yaml:"addressPrefixes"`
}

type PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference interface {
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
	InternalValue() *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool
	SetInternalValue(val *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool)
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

// The jsii proxy struct for PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference
type jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) AddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) AddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) InternalValue() *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool {
	var returns *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference_Override(p PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) SetAddressPrefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"addressPrefixes",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) SetInternalValue(val *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PointToSiteVpnGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#create PointToSiteVpnGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#delete PointToSiteVpnGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#read PointToSiteVpnGateway#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#update PointToSiteVpnGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type PointToSiteVpnGatewayTimeoutsOutputReference interface {
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

// The jsii proxy struct for PointToSiteVpnGatewayTimeoutsOutputReference
type jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewPointToSiteVpnGatewayTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PointToSiteVpnGatewayTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPointToSiteVpnGatewayTimeoutsOutputReference_Override(p PointToSiteVpnGatewayTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		p,
		"resetCreate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		p,
		"resetDelete",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		p,
		"resetRead",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		p,
		"resetUpdate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

