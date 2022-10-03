package frontdoor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/frontdoor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor azurerm_frontdoor}.
type Frontdoor interface {
	cdktf.TerraformResource
	BackendPool() FrontdoorBackendPoolList
	BackendPoolHealthProbe() FrontdoorBackendPoolHealthProbeList
	BackendPoolHealthProbeInput() interface{}
	BackendPoolHealthProbes() cdktf.StringMap
	BackendPoolInput() interface{}
	BackendPoolLoadBalancing() FrontdoorBackendPoolLoadBalancingList
	BackendPoolLoadBalancingInput() interface{}
	BackendPoolLoadBalancingSettings() cdktf.StringMap
	BackendPools() cdktf.StringMap
	BackendPoolSettings() FrontdoorBackendPoolSettingsList
	BackendPoolSettingsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cname() *string
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
	ExplicitResourceOrder() FrontdoorExplicitResourceOrderList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	FriendlyName() *string
	SetFriendlyName(val *string)
	FriendlyNameInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	FrontendEndpoint() FrontdoorFrontendEndpointList
	FrontendEndpointInput() interface{}
	FrontendEndpoints() cdktf.StringMap
	HeaderFrontdoorId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerEnabled() interface{}
	SetLoadBalancerEnabled(val interface{})
	LoadBalancerEnabledInput() interface{}
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
	RoutingRule() FrontdoorRoutingRuleList
	RoutingRuleInput() interface{}
	RoutingRules() cdktf.StringMap
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FrontdoorTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutBackendPool(value interface{})
	PutBackendPoolHealthProbe(value interface{})
	PutBackendPoolLoadBalancing(value interface{})
	PutBackendPoolSettings(value interface{})
	PutFrontendEndpoint(value interface{})
	PutRoutingRule(value interface{})
	PutTimeouts(value *FrontdoorTimeouts)
	ResetBackendPoolSettings()
	ResetFriendlyName()
	ResetId()
	ResetLoadBalancerEnabled()
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

// The jsii proxy struct for Frontdoor
type jsiiProxy_Frontdoor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Frontdoor) BackendPool() FrontdoorBackendPoolList {
	var returns FrontdoorBackendPoolList
	_jsii_.Get(
		j,
		"backendPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolHealthProbe() FrontdoorBackendPoolHealthProbeList {
	var returns FrontdoorBackendPoolHealthProbeList
	_jsii_.Get(
		j,
		"backendPoolHealthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolHealthProbeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendPoolHealthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolHealthProbes() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"backendPoolHealthProbes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolLoadBalancing() FrontdoorBackendPoolLoadBalancingList {
	var returns FrontdoorBackendPoolLoadBalancingList
	_jsii_.Get(
		j,
		"backendPoolLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolLoadBalancingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendPoolLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolLoadBalancingSettings() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"backendPoolLoadBalancingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPools() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"backendPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolSettings() FrontdoorBackendPoolSettingsList {
	var returns FrontdoorBackendPoolSettingsList
	_jsii_.Get(
		j,
		"backendPoolSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) BackendPoolSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendPoolSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Cname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) ExplicitResourceOrder() FrontdoorExplicitResourceOrderList {
	var returns FrontdoorExplicitResourceOrderList
	_jsii_.Get(
		j,
		"explicitResourceOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) FriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) FrontendEndpoint() FrontdoorFrontendEndpointList {
	var returns FrontdoorFrontendEndpointList
	_jsii_.Get(
		j,
		"frontendEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) FrontendEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) FrontendEndpoints() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"frontendEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) HeaderFrontdoorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerFrontdoorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) LoadBalancerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) LoadBalancerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) RoutingRule() FrontdoorRoutingRuleList {
	var returns FrontdoorRoutingRuleList
	_jsii_.Get(
		j,
		"routingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) RoutingRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) RoutingRules() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"routingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) Timeouts() FrontdoorTimeoutsOutputReference {
	var returns FrontdoorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Frontdoor) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor azurerm_frontdoor} Resource.
func NewFrontdoor(scope constructs.Construct, id *string, config *FrontdoorConfig) Frontdoor {
	_init_.Initialize()

	j := jsiiProxy_Frontdoor{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor azurerm_frontdoor} Resource.
func NewFrontdoor_Override(f Frontdoor, scope constructs.Construct, id *string, config *FrontdoorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_Frontdoor) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetFriendlyName(val *string) {
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetLoadBalancerEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"loadBalancerEnabled",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_Frontdoor) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
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
func Frontdoor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Frontdoor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_Frontdoor) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_Frontdoor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_Frontdoor) PutBackendPool(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putBackendPool",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutBackendPoolHealthProbe(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putBackendPoolHealthProbe",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutBackendPoolLoadBalancing(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putBackendPoolLoadBalancing",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutBackendPoolSettings(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putBackendPoolSettings",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutFrontendEndpoint(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putFrontendEndpoint",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutRoutingRule(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putRoutingRule",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutTimeouts(value *FrontdoorTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) ResetBackendPoolSettings() {
	_jsii_.InvokeVoid(
		f,
		"resetBackendPoolSettings",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Frontdoor) ResetFriendlyName() {
	_jsii_.InvokeVoid(
		f,
		"resetFriendlyName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Frontdoor) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Frontdoor) ResetLoadBalancerEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetLoadBalancerEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Frontdoor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Frontdoor) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Frontdoor) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Frontdoor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPool struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend Frontdoor#backend}
	Backend interface{} `field:"required" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#health_probe_name Frontdoor#health_probe_name}.
	HealthProbeName *string `field:"required" json:"healthProbeName" yaml:"healthProbeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#load_balancing_name Frontdoor#load_balancing_name}.
	LoadBalancingName *string `field:"required" json:"loadBalancingName" yaml:"loadBalancingName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

type FrontdoorBackendPoolBackend struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#address Frontdoor#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#host_header Frontdoor#host_header}.
	HostHeader *string `field:"required" json:"hostHeader" yaml:"hostHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#http_port Frontdoor#http_port}.
	HttpPort *float64 `field:"required" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#https_port Frontdoor#https_port}.
	HttpsPort *float64 `field:"required" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#enabled Frontdoor#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#priority Frontdoor#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#weight Frontdoor#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

type FrontdoorBackendPoolBackendList interface {
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
	Get(index *float64) FrontdoorBackendPoolBackendOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolBackendList
type jsiiProxy_FrontdoorBackendPoolBackendList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolBackendList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorBackendPoolBackendList {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolBackendList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolBackendList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolBackendList_Override(f FrontdoorBackendPoolBackendList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolBackendList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendList) Get(index *float64) FrontdoorBackendPoolBackendOutputReference {
	var returns FrontdoorBackendPoolBackendOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolBackendOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HostHeader() *string
	SetHostHeader(val *string)
	HostHeaderInput() *string
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	HttpsPort() *float64
	SetHttpsPort(val *float64)
	HttpsPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	ResetEnabled()
	ResetPriority()
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolBackendOutputReference
type jsiiProxy_FrontdoorBackendPoolBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) HostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) HostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) HttpsPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) HttpsPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorBackendPoolBackendOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolBackendOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolBackendOutputReference_Override(f FrontdoorBackendPoolBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetAddress(val *string) {
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetHostHeader(val *string) {
	_jsii_.Set(
		j,
		"hostHeader",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetHttpPort(val *float64) {
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetHttpsPort(val *float64) {
	_jsii_.Set(
		j,
		"httpsPort",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		f,
		"resetPriority",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		f,
		"resetWeight",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolHealthProbe struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#enabled Frontdoor#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#interval_in_seconds Frontdoor#interval_in_seconds}.
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#path Frontdoor#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#probe_method Frontdoor#probe_method}.
	ProbeMethod *string `field:"optional" json:"probeMethod" yaml:"probeMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#protocol Frontdoor#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

type FrontdoorBackendPoolHealthProbeList interface {
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
	Get(index *float64) FrontdoorBackendPoolHealthProbeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolHealthProbeList
type jsiiProxy_FrontdoorBackendPoolHealthProbeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolHealthProbeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorBackendPoolHealthProbeList {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolHealthProbeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolHealthProbeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolHealthProbeList_Override(f FrontdoorBackendPoolHealthProbeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolHealthProbeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeList) Get(index *float64) FrontdoorBackendPoolHealthProbeOutputReference {
	var returns FrontdoorBackendPoolHealthProbeOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolHealthProbeOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IntervalInSeconds() *float64
	SetIntervalInSeconds(val *float64)
	IntervalInSecondsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	ProbeMethod() *string
	SetProbeMethod(val *string)
	ProbeMethodInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	ResetEnabled()
	ResetIntervalInSeconds()
	ResetPath()
	ResetProbeMethod()
	ResetProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolHealthProbeOutputReference
type jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) IntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) IntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ProbeMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ProbeMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolHealthProbeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorBackendPoolHealthProbeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolHealthProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolHealthProbeOutputReference_Override(f FrontdoorBackendPoolHealthProbeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolHealthProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetIntervalInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"intervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetProbeMethod(val *string) {
	_jsii_.Set(
		j,
		"probeMethod",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ResetIntervalInSeconds() {
	_jsii_.InvokeVoid(
		f,
		"resetIntervalInSeconds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		f,
		"resetPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ResetProbeMethod() {
	_jsii_.InvokeVoid(
		f,
		"resetProbeMethod",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		f,
		"resetProtocol",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolHealthProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolList interface {
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
	Get(index *float64) FrontdoorBackendPoolOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolList
type jsiiProxy_FrontdoorBackendPoolList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorBackendPoolList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorBackendPoolList {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolList_Override(f FrontdoorBackendPoolList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolList) Get(index *float64) FrontdoorBackendPoolOutputReference {
	var returns FrontdoorBackendPoolOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolLoadBalancing struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#additional_latency_milliseconds Frontdoor#additional_latency_milliseconds}.
	AdditionalLatencyMilliseconds *float64 `field:"optional" json:"additionalLatencyMilliseconds" yaml:"additionalLatencyMilliseconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#sample_size Frontdoor#sample_size}.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#successful_samples_required Frontdoor#successful_samples_required}.
	SuccessfulSamplesRequired *float64 `field:"optional" json:"successfulSamplesRequired" yaml:"successfulSamplesRequired"`
}

type FrontdoorBackendPoolLoadBalancingList interface {
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
	Get(index *float64) FrontdoorBackendPoolLoadBalancingOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolLoadBalancingList
type jsiiProxy_FrontdoorBackendPoolLoadBalancingList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolLoadBalancingList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorBackendPoolLoadBalancingList {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolLoadBalancingList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolLoadBalancingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolLoadBalancingList_Override(f FrontdoorBackendPoolLoadBalancingList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolLoadBalancingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) Get(index *float64) FrontdoorBackendPoolLoadBalancingOutputReference {
	var returns FrontdoorBackendPoolLoadBalancingOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolLoadBalancingOutputReference interface {
	cdktf.ComplexObject
	AdditionalLatencyMilliseconds() *float64
	SetAdditionalLatencyMilliseconds(val *float64)
	AdditionalLatencyMillisecondsInput() *float64
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
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	SampleSize() *float64
	SetSampleSize(val *float64)
	SampleSizeInput() *float64
	SuccessfulSamplesRequired() *float64
	SetSuccessfulSamplesRequired(val *float64)
	SuccessfulSamplesRequiredInput() *float64
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
	ResetAdditionalLatencyMilliseconds()
	ResetSampleSize()
	ResetSuccessfulSamplesRequired()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolLoadBalancingOutputReference
type jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) AdditionalLatencyMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalLatencyMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) AdditionalLatencyMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalLatencyMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SampleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SampleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SuccessfulSamplesRequired() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulSamplesRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SuccessfulSamplesRequiredInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulSamplesRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolLoadBalancingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorBackendPoolLoadBalancingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolLoadBalancingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolLoadBalancingOutputReference_Override(f FrontdoorBackendPoolLoadBalancingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolLoadBalancingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetAdditionalLatencyMilliseconds(val *float64) {
	_jsii_.Set(
		j,
		"additionalLatencyMilliseconds",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetSampleSize(val *float64) {
	_jsii_.Set(
		j,
		"sampleSize",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetSuccessfulSamplesRequired(val *float64) {
	_jsii_.Set(
		j,
		"successfulSamplesRequired",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) ResetAdditionalLatencyMilliseconds() {
	_jsii_.InvokeVoid(
		f,
		"resetAdditionalLatencyMilliseconds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) ResetSampleSize() {
	_jsii_.InvokeVoid(
		f,
		"resetSampleSize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) ResetSuccessfulSamplesRequired() {
	_jsii_.InvokeVoid(
		f,
		"resetSuccessfulSamplesRequired",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolOutputReference interface {
	cdktf.ComplexObject
	Backend() FrontdoorBackendPoolBackendList
	BackendInput() interface{}
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
	HealthProbeName() *string
	SetHealthProbeName(val *string)
	HealthProbeNameInput() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancingName() *string
	SetLoadBalancingName(val *string)
	LoadBalancingNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutBackend(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolOutputReference
type jsiiProxy_FrontdoorBackendPoolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) Backend() FrontdoorBackendPoolBackendList {
	var returns FrontdoorBackendPoolBackendList
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) BackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) HealthProbeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) HealthProbeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) LoadBalancingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) LoadBalancingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorBackendPoolOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolOutputReference_Override(f FrontdoorBackendPoolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) SetHealthProbeName(val *string) {
	_jsii_.Set(
		j,
		"healthProbeName",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) SetLoadBalancingName(val *string) {
	_jsii_.Set(
		j,
		"loadBalancingName",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) PutBackend(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putBackend",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#enforce_backend_pools_certificate_name_check Frontdoor#enforce_backend_pools_certificate_name_check}.
	EnforceBackendPoolsCertificateNameCheck interface{} `field:"required" json:"enforceBackendPoolsCertificateNameCheck" yaml:"enforceBackendPoolsCertificateNameCheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend_pools_send_receive_timeout_seconds Frontdoor#backend_pools_send_receive_timeout_seconds}.
	BackendPoolsSendReceiveTimeoutSeconds *float64 `field:"optional" json:"backendPoolsSendReceiveTimeoutSeconds" yaml:"backendPoolsSendReceiveTimeoutSeconds"`
}

type FrontdoorBackendPoolSettingsList interface {
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
	Get(index *float64) FrontdoorBackendPoolSettingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolSettingsList
type jsiiProxy_FrontdoorBackendPoolSettingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolSettingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorBackendPoolSettingsList {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolSettingsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolSettingsList_Override(f FrontdoorBackendPoolSettingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsList) Get(index *float64) FrontdoorBackendPoolSettingsOutputReference {
	var returns FrontdoorBackendPoolSettingsOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorBackendPoolSettingsOutputReference interface {
	cdktf.ComplexObject
	BackendPoolsSendReceiveTimeoutSeconds() *float64
	SetBackendPoolsSendReceiveTimeoutSeconds(val *float64)
	BackendPoolsSendReceiveTimeoutSecondsInput() *float64
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
	EnforceBackendPoolsCertificateNameCheck() interface{}
	SetEnforceBackendPoolsCertificateNameCheck(val interface{})
	EnforceBackendPoolsCertificateNameCheckInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetBackendPoolsSendReceiveTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorBackendPoolSettingsOutputReference
type jsiiProxy_FrontdoorBackendPoolSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) BackendPoolsSendReceiveTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPoolsSendReceiveTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) BackendPoolsSendReceiveTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPoolsSendReceiveTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) EnforceBackendPoolsCertificateNameCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceBackendPoolsCertificateNameCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) EnforceBackendPoolsCertificateNameCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceBackendPoolsCertificateNameCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorBackendPoolSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorBackendPoolSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorBackendPoolSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorBackendPoolSettingsOutputReference_Override(f FrontdoorBackendPoolSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorBackendPoolSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) SetBackendPoolsSendReceiveTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"backendPoolsSendReceiveTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) SetEnforceBackendPoolsCertificateNameCheck(val interface{}) {
	_jsii_.Set(
		j,
		"enforceBackendPoolsCertificateNameCheck",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) ResetBackendPoolsSendReceiveTimeoutSeconds() {
	_jsii_.InvokeVoid(
		f,
		"resetBackendPoolsSendReceiveTimeoutSeconds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorBackendPoolSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorConfig struct {
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
	// backend_pool block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend_pool Frontdoor#backend_pool}
	BackendPool interface{} `field:"required" json:"backendPool" yaml:"backendPool"`
	// backend_pool_health_probe block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend_pool_health_probe Frontdoor#backend_pool_health_probe}
	BackendPoolHealthProbe interface{} `field:"required" json:"backendPoolHealthProbe" yaml:"backendPoolHealthProbe"`
	// backend_pool_load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend_pool_load_balancing Frontdoor#backend_pool_load_balancing}
	BackendPoolLoadBalancing interface{} `field:"required" json:"backendPoolLoadBalancing" yaml:"backendPoolLoadBalancing"`
	// frontend_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#frontend_endpoint Frontdoor#frontend_endpoint}
	FrontendEndpoint interface{} `field:"required" json:"frontendEndpoint" yaml:"frontendEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#resource_group_name Frontdoor#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// routing_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#routing_rule Frontdoor#routing_rule}
	RoutingRule interface{} `field:"required" json:"routingRule" yaml:"routingRule"`
	// backend_pool_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend_pool_settings Frontdoor#backend_pool_settings}
	BackendPoolSettings interface{} `field:"optional" json:"backendPoolSettings" yaml:"backendPoolSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#friendly_name Frontdoor#friendly_name}.
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#id Frontdoor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#load_balancer_enabled Frontdoor#load_balancer_enabled}.
	LoadBalancerEnabled interface{} `field:"optional" json:"loadBalancerEnabled" yaml:"loadBalancerEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#tags Frontdoor#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#timeouts Frontdoor#timeouts}
	Timeouts *FrontdoorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type FrontdoorExplicitResourceOrder struct {
}

type FrontdoorExplicitResourceOrderList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) FrontdoorExplicitResourceOrderOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorExplicitResourceOrderList
type jsiiProxy_FrontdoorExplicitResourceOrderList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorExplicitResourceOrderList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorExplicitResourceOrderList {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorExplicitResourceOrderList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorExplicitResourceOrderList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorExplicitResourceOrderList_Override(f FrontdoorExplicitResourceOrderList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorExplicitResourceOrderList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderList) Get(index *float64) FrontdoorExplicitResourceOrderOutputReference {
	var returns FrontdoorExplicitResourceOrderOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorExplicitResourceOrderOutputReference interface {
	cdktf.ComplexObject
	BackendPoolHealthProbeIds() *[]*string
	BackendPoolIds() *[]*string
	BackendPoolLoadBalancingIds() *[]*string
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
	FrontendEndpointIds() *[]*string
	InternalValue() *FrontdoorExplicitResourceOrder
	SetInternalValue(val *FrontdoorExplicitResourceOrder)
	RoutingRuleIds() *[]*string
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

// The jsii proxy struct for FrontdoorExplicitResourceOrderOutputReference
type jsiiProxy_FrontdoorExplicitResourceOrderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) BackendPoolHealthProbeIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backendPoolHealthProbeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) BackendPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backendPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) BackendPoolLoadBalancingIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backendPoolLoadBalancingIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) FrontendEndpointIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frontendEndpointIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) InternalValue() *FrontdoorExplicitResourceOrder {
	var returns *FrontdoorExplicitResourceOrder
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) RoutingRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routingRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorExplicitResourceOrderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorExplicitResourceOrderOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorExplicitResourceOrderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorExplicitResourceOrderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorExplicitResourceOrderOutputReference_Override(f FrontdoorExplicitResourceOrderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorExplicitResourceOrderOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) SetInternalValue(val *FrontdoorExplicitResourceOrder) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorFrontendEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#host_name Frontdoor#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#session_affinity_enabled Frontdoor#session_affinity_enabled}.
	SessionAffinityEnabled interface{} `field:"optional" json:"sessionAffinityEnabled" yaml:"sessionAffinityEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#session_affinity_ttl_seconds Frontdoor#session_affinity_ttl_seconds}.
	SessionAffinityTtlSeconds *float64 `field:"optional" json:"sessionAffinityTtlSeconds" yaml:"sessionAffinityTtlSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#web_application_firewall_policy_link_id Frontdoor#web_application_firewall_policy_link_id}.
	WebApplicationFirewallPolicyLinkId *string `field:"optional" json:"webApplicationFirewallPolicyLinkId" yaml:"webApplicationFirewallPolicyLinkId"`
}

type FrontdoorFrontendEndpointList interface {
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
	Get(index *float64) FrontdoorFrontendEndpointOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorFrontendEndpointList
type jsiiProxy_FrontdoorFrontendEndpointList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorFrontendEndpointList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorFrontendEndpointList {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorFrontendEndpointList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorFrontendEndpointList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorFrontendEndpointList_Override(f FrontdoorFrontendEndpointList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorFrontendEndpointList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorFrontendEndpointList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointList) Get(index *float64) FrontdoorFrontendEndpointOutputReference {
	var returns FrontdoorFrontendEndpointOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorFrontendEndpointOutputReference interface {
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
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	SessionAffinityEnabled() interface{}
	SetSessionAffinityEnabled(val interface{})
	SessionAffinityEnabledInput() interface{}
	SessionAffinityTtlSeconds() *float64
	SetSessionAffinityTtlSeconds(val *float64)
	SessionAffinityTtlSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebApplicationFirewallPolicyLinkId() *string
	SetWebApplicationFirewallPolicyLinkId(val *string)
	WebApplicationFirewallPolicyLinkIdInput() *string
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
	ResetSessionAffinityEnabled()
	ResetSessionAffinityTtlSeconds()
	ResetWebApplicationFirewallPolicyLinkId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorFrontendEndpointOutputReference
type jsiiProxy_FrontdoorFrontendEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SessionAffinityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionAffinityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SessionAffinityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionAffinityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SessionAffinityTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SessionAffinityTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) WebApplicationFirewallPolicyLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webApplicationFirewallPolicyLinkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) WebApplicationFirewallPolicyLinkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webApplicationFirewallPolicyLinkIdInput",
		&returns,
	)
	return returns
}


func NewFrontdoorFrontendEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorFrontendEndpointOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorFrontendEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorFrontendEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorFrontendEndpointOutputReference_Override(f FrontdoorFrontendEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorFrontendEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetHostName(val *string) {
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetSessionAffinityEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"sessionAffinityEnabled",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetSessionAffinityTtlSeconds(val *float64) {
	_jsii_.Set(
		j,
		"sessionAffinityTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFrontendEndpointOutputReference) SetWebApplicationFirewallPolicyLinkId(val *string) {
	_jsii_.Set(
		j,
		"webApplicationFirewallPolicyLinkId",
		val,
	)
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) ResetSessionAffinityEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetSessionAffinityEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) ResetSessionAffinityTtlSeconds() {
	_jsii_.InvokeVoid(
		f,
		"resetSessionAffinityTtlSeconds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) ResetWebApplicationFirewallPolicyLinkId() {
	_jsii_.InvokeVoid(
		f,
		"resetWebApplicationFirewallPolicyLinkId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFrontendEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorRoutingRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#accepted_protocols Frontdoor#accepted_protocols}.
	AcceptedProtocols *[]*string `field:"required" json:"acceptedProtocols" yaml:"acceptedProtocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#frontend_endpoints Frontdoor#frontend_endpoints}.
	FrontendEndpoints *[]*string `field:"required" json:"frontendEndpoints" yaml:"frontendEndpoints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#patterns_to_match Frontdoor#patterns_to_match}.
	PatternsToMatch *[]*string `field:"required" json:"patternsToMatch" yaml:"patternsToMatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#enabled Frontdoor#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// forwarding_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#forwarding_configuration Frontdoor#forwarding_configuration}
	ForwardingConfiguration *FrontdoorRoutingRuleForwardingConfiguration `field:"optional" json:"forwardingConfiguration" yaml:"forwardingConfiguration"`
	// redirect_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#redirect_configuration Frontdoor#redirect_configuration}
	RedirectConfiguration *FrontdoorRoutingRuleRedirectConfiguration `field:"optional" json:"redirectConfiguration" yaml:"redirectConfiguration"`
}

type FrontdoorRoutingRuleForwardingConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend_pool_name Frontdoor#backend_pool_name}.
	BackendPoolName *string `field:"required" json:"backendPoolName" yaml:"backendPoolName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#cache_duration Frontdoor#cache_duration}.
	CacheDuration *string `field:"optional" json:"cacheDuration" yaml:"cacheDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#cache_enabled Frontdoor#cache_enabled}.
	CacheEnabled interface{} `field:"optional" json:"cacheEnabled" yaml:"cacheEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#cache_query_parameters Frontdoor#cache_query_parameters}.
	CacheQueryParameters *[]*string `field:"optional" json:"cacheQueryParameters" yaml:"cacheQueryParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#cache_query_parameter_strip_directive Frontdoor#cache_query_parameter_strip_directive}.
	CacheQueryParameterStripDirective *string `field:"optional" json:"cacheQueryParameterStripDirective" yaml:"cacheQueryParameterStripDirective"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#cache_use_dynamic_compression Frontdoor#cache_use_dynamic_compression}.
	CacheUseDynamicCompression interface{} `field:"optional" json:"cacheUseDynamicCompression" yaml:"cacheUseDynamicCompression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#custom_forwarding_path Frontdoor#custom_forwarding_path}.
	CustomForwardingPath *string `field:"optional" json:"customForwardingPath" yaml:"customForwardingPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#forwarding_protocol Frontdoor#forwarding_protocol}.
	ForwardingProtocol *string `field:"optional" json:"forwardingProtocol" yaml:"forwardingProtocol"`
}

type FrontdoorRoutingRuleForwardingConfigurationOutputReference interface {
	cdktf.ComplexObject
	BackendPoolName() *string
	SetBackendPoolName(val *string)
	BackendPoolNameInput() *string
	CacheDuration() *string
	SetCacheDuration(val *string)
	CacheDurationInput() *string
	CacheEnabled() interface{}
	SetCacheEnabled(val interface{})
	CacheEnabledInput() interface{}
	CacheQueryParameters() *[]*string
	SetCacheQueryParameters(val *[]*string)
	CacheQueryParametersInput() *[]*string
	CacheQueryParameterStripDirective() *string
	SetCacheQueryParameterStripDirective(val *string)
	CacheQueryParameterStripDirectiveInput() *string
	CacheUseDynamicCompression() interface{}
	SetCacheUseDynamicCompression(val interface{})
	CacheUseDynamicCompressionInput() interface{}
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
	CustomForwardingPath() *string
	SetCustomForwardingPath(val *string)
	CustomForwardingPathInput() *string
	ForwardingProtocol() *string
	SetForwardingProtocol(val *string)
	ForwardingProtocolInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *FrontdoorRoutingRuleForwardingConfiguration
	SetInternalValue(val *FrontdoorRoutingRuleForwardingConfiguration)
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
	ResetCacheDuration()
	ResetCacheEnabled()
	ResetCacheQueryParameters()
	ResetCacheQueryParameterStripDirective()
	ResetCacheUseDynamicCompression()
	ResetCustomForwardingPath()
	ResetForwardingProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorRoutingRuleForwardingConfigurationOutputReference
type jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) BackendPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) BackendPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheQueryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheQueryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheQueryParameterStripDirective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheQueryParameterStripDirective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheQueryParameterStripDirectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheQueryParameterStripDirectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheUseDynamicCompression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheUseDynamicCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheUseDynamicCompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheUseDynamicCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CustomForwardingPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customForwardingPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CustomForwardingPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customForwardingPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ForwardingProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ForwardingProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) InternalValue() *FrontdoorRoutingRuleForwardingConfiguration {
	var returns *FrontdoorRoutingRuleForwardingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorRoutingRuleForwardingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FrontdoorRoutingRuleForwardingConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleForwardingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrontdoorRoutingRuleForwardingConfigurationOutputReference_Override(f FrontdoorRoutingRuleForwardingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleForwardingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetBackendPoolName(val *string) {
	_jsii_.Set(
		j,
		"backendPoolName",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetCacheDuration(val *string) {
	_jsii_.Set(
		j,
		"cacheDuration",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetCacheEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cacheEnabled",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetCacheQueryParameters(val *[]*string) {
	_jsii_.Set(
		j,
		"cacheQueryParameters",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetCacheQueryParameterStripDirective(val *string) {
	_jsii_.Set(
		j,
		"cacheQueryParameterStripDirective",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetCacheUseDynamicCompression(val interface{}) {
	_jsii_.Set(
		j,
		"cacheUseDynamicCompression",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetCustomForwardingPath(val *string) {
	_jsii_.Set(
		j,
		"customForwardingPath",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetForwardingProtocol(val *string) {
	_jsii_.Set(
		j,
		"forwardingProtocol",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetInternalValue(val *FrontdoorRoutingRuleForwardingConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheDuration() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheDuration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheQueryParameters() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheQueryParameters",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheQueryParameterStripDirective() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheQueryParameterStripDirective",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheUseDynamicCompression() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheUseDynamicCompression",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCustomForwardingPath() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomForwardingPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetForwardingProtocol() {
	_jsii_.InvokeVoid(
		f,
		"resetForwardingProtocol",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorRoutingRuleList interface {
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
	Get(index *float64) FrontdoorRoutingRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorRoutingRuleList
type jsiiProxy_FrontdoorRoutingRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFrontdoorRoutingRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FrontdoorRoutingRuleList {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorRoutingRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFrontdoorRoutingRuleList_Override(f FrontdoorRoutingRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleList) Get(index *float64) FrontdoorRoutingRuleOutputReference {
	var returns FrontdoorRoutingRuleOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorRoutingRuleOutputReference interface {
	cdktf.ComplexObject
	AcceptedProtocols() *[]*string
	SetAcceptedProtocols(val *[]*string)
	AcceptedProtocolsInput() *[]*string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ForwardingConfiguration() FrontdoorRoutingRuleForwardingConfigurationOutputReference
	ForwardingConfigurationInput() *FrontdoorRoutingRuleForwardingConfiguration
	// Experimental.
	Fqn() *string
	FrontendEndpoints() *[]*string
	SetFrontendEndpoints(val *[]*string)
	FrontendEndpointsInput() *[]*string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PatternsToMatch() *[]*string
	SetPatternsToMatch(val *[]*string)
	PatternsToMatchInput() *[]*string
	RedirectConfiguration() FrontdoorRoutingRuleRedirectConfigurationOutputReference
	RedirectConfigurationInput() *FrontdoorRoutingRuleRedirectConfiguration
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
	PutForwardingConfiguration(value *FrontdoorRoutingRuleForwardingConfiguration)
	PutRedirectConfiguration(value *FrontdoorRoutingRuleRedirectConfiguration)
	ResetEnabled()
	ResetForwardingConfiguration()
	ResetRedirectConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorRoutingRuleOutputReference
type jsiiProxy_FrontdoorRoutingRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) AcceptedProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceptedProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) AcceptedProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceptedProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) ForwardingConfiguration() FrontdoorRoutingRuleForwardingConfigurationOutputReference {
	var returns FrontdoorRoutingRuleForwardingConfigurationOutputReference
	_jsii_.Get(
		j,
		"forwardingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) ForwardingConfigurationInput() *FrontdoorRoutingRuleForwardingConfiguration {
	var returns *FrontdoorRoutingRuleForwardingConfiguration
	_jsii_.Get(
		j,
		"forwardingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) FrontendEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frontendEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) FrontendEndpointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frontendEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) PatternsToMatch() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) PatternsToMatchInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) RedirectConfiguration() FrontdoorRoutingRuleRedirectConfigurationOutputReference {
	var returns FrontdoorRoutingRuleRedirectConfigurationOutputReference
	_jsii_.Get(
		j,
		"redirectConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) RedirectConfigurationInput() *FrontdoorRoutingRuleRedirectConfiguration {
	var returns *FrontdoorRoutingRuleRedirectConfiguration
	_jsii_.Get(
		j,
		"redirectConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorRoutingRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorRoutingRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorRoutingRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorRoutingRuleOutputReference_Override(f FrontdoorRoutingRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetAcceptedProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"acceptedProtocols",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetFrontendEndpoints(val *[]*string) {
	_jsii_.Set(
		j,
		"frontendEndpoints",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetPatternsToMatch(val *[]*string) {
	_jsii_.Set(
		j,
		"patternsToMatch",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) PutForwardingConfiguration(value *FrontdoorRoutingRuleForwardingConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putForwardingConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) PutRedirectConfiguration(value *FrontdoorRoutingRuleRedirectConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putRedirectConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) ResetForwardingConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetForwardingConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) ResetRedirectConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetRedirectConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorRoutingRuleRedirectConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#redirect_protocol Frontdoor#redirect_protocol}.
	RedirectProtocol *string `field:"required" json:"redirectProtocol" yaml:"redirectProtocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#redirect_type Frontdoor#redirect_type}.
	RedirectType *string `field:"required" json:"redirectType" yaml:"redirectType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#custom_fragment Frontdoor#custom_fragment}.
	CustomFragment *string `field:"optional" json:"customFragment" yaml:"customFragment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#custom_host Frontdoor#custom_host}.
	CustomHost *string `field:"optional" json:"customHost" yaml:"customHost"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#custom_path Frontdoor#custom_path}.
	CustomPath *string `field:"optional" json:"customPath" yaml:"customPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#custom_query_string Frontdoor#custom_query_string}.
	CustomQueryString *string `field:"optional" json:"customQueryString" yaml:"customQueryString"`
}

type FrontdoorRoutingRuleRedirectConfigurationOutputReference interface {
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
	CustomFragment() *string
	SetCustomFragment(val *string)
	CustomFragmentInput() *string
	CustomHost() *string
	SetCustomHost(val *string)
	CustomHostInput() *string
	CustomPath() *string
	SetCustomPath(val *string)
	CustomPathInput() *string
	CustomQueryString() *string
	SetCustomQueryString(val *string)
	CustomQueryStringInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *FrontdoorRoutingRuleRedirectConfiguration
	SetInternalValue(val *FrontdoorRoutingRuleRedirectConfiguration)
	RedirectProtocol() *string
	SetRedirectProtocol(val *string)
	RedirectProtocolInput() *string
	RedirectType() *string
	SetRedirectType(val *string)
	RedirectTypeInput() *string
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
	ResetCustomFragment()
	ResetCustomHost()
	ResetCustomPath()
	ResetCustomQueryString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorRoutingRuleRedirectConfigurationOutputReference
type jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomFragment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomFragmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomQueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) CustomQueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) InternalValue() *FrontdoorRoutingRuleRedirectConfiguration {
	var returns *FrontdoorRoutingRuleRedirectConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) RedirectProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) RedirectProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) RedirectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) RedirectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorRoutingRuleRedirectConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FrontdoorRoutingRuleRedirectConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleRedirectConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrontdoorRoutingRuleRedirectConfigurationOutputReference_Override(f FrontdoorRoutingRuleRedirectConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleRedirectConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetCustomFragment(val *string) {
	_jsii_.Set(
		j,
		"customFragment",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetCustomHost(val *string) {
	_jsii_.Set(
		j,
		"customHost",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetCustomPath(val *string) {
	_jsii_.Set(
		j,
		"customPath",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetCustomQueryString(val *string) {
	_jsii_.Set(
		j,
		"customQueryString",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetInternalValue(val *FrontdoorRoutingRuleRedirectConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetRedirectProtocol(val *string) {
	_jsii_.Set(
		j,
		"redirectProtocol",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetRedirectType(val *string) {
	_jsii_.Set(
		j,
		"redirectType",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ResetCustomFragment() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomFragment",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ResetCustomHost() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomHost",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ResetCustomPath() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ResetCustomQueryString() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomQueryString",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleRedirectConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#create Frontdoor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#delete Frontdoor#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#read Frontdoor#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#update Frontdoor#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type FrontdoorTimeoutsOutputReference interface {
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

// The jsii proxy struct for FrontdoorTimeoutsOutputReference
type jsiiProxy_FrontdoorTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewFrontdoorTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FrontdoorTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrontdoorTimeoutsOutputReference_Override(f FrontdoorTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		f,
		"resetRead",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

