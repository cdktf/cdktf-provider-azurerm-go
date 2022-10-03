package cdnfrontdoorroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/cdnfrontdoorroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route azurerm_cdn_frontdoor_route}.
type CdnFrontdoorRoute interface {
	cdktf.TerraformResource
	Cache() CdnFrontdoorRouteCacheOutputReference
	CacheInput() *CdnFrontdoorRouteCache
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnFrontdoorEndpointId() *string
	SetCdnFrontdoorEndpointId(val *string)
	CdnFrontdoorEndpointIdInput() *string
	CdnFrontdoorOriginGroupId() *string
	SetCdnFrontdoorOriginGroupId(val *string)
	CdnFrontdoorOriginGroupIdInput() *string
	CdnFrontdoorOriginIds() *[]*string
	SetCdnFrontdoorOriginIds(val *[]*string)
	CdnFrontdoorOriginIdsInput() *[]*string
	CdnFrontdoorOriginPath() *string
	SetCdnFrontdoorOriginPath(val *string)
	CdnFrontdoorOriginPathInput() *string
	CdnFrontdoorRuleSetIds() *[]*string
	SetCdnFrontdoorRuleSetIds(val *[]*string)
	CdnFrontdoorRuleSetIdsInput() *[]*string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForwardingProtocol() *string
	SetForwardingProtocol(val *string)
	ForwardingProtocolInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpsRedirectEnabled() interface{}
	SetHttpsRedirectEnabled(val interface{})
	HttpsRedirectEnabledInput() interface{}
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
	PatternsToMatch() *[]*string
	SetPatternsToMatch(val *[]*string)
	PatternsToMatchInput() *[]*string
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
	SupportedProtocols() *[]*string
	SetSupportedProtocols(val *[]*string)
	SupportedProtocolsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CdnFrontdoorRouteTimeoutsOutputReference
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
	PutCache(value *CdnFrontdoorRouteCache)
	PutTimeouts(value *CdnFrontdoorRouteTimeouts)
	ResetCache()
	ResetCdnFrontdoorOriginPath()
	ResetCdnFrontdoorRuleSetIds()
	ResetEnabled()
	ResetForwardingProtocol()
	ResetHttpsRedirectEnabled()
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

// The jsii proxy struct for CdnFrontdoorRoute
type jsiiProxy_CdnFrontdoorRoute struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CdnFrontdoorRoute) Cache() CdnFrontdoorRouteCacheOutputReference {
	var returns CdnFrontdoorRouteCacheOutputReference
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CacheInput() *CdnFrontdoorRouteCache {
	var returns *CdnFrontdoorRouteCache
	_jsii_.Get(
		j,
		"cacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorRuleSetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorRuleSetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorRuleSetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorRuleSetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) ForwardingProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) ForwardingProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) HttpsRedirectEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) HttpsRedirectEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) PatternsToMatch() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) PatternsToMatchInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) SupportedProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) SupportedProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Timeouts() CdnFrontdoorRouteTimeoutsOutputReference {
	var returns CdnFrontdoorRouteTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route azurerm_cdn_frontdoor_route} Resource.
func NewCdnFrontdoorRoute(scope constructs.Construct, id *string, config *CdnFrontdoorRouteConfig) CdnFrontdoorRoute {
	_init_.Initialize()

	j := jsiiProxy_CdnFrontdoorRoute{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route azurerm_cdn_frontdoor_route} Resource.
func NewCdnFrontdoorRoute_Override(c CdnFrontdoorRoute, scope constructs.Construct, id *string, config *CdnFrontdoorRouteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetCdnFrontdoorEndpointId(val *string) {
	_jsii_.Set(
		j,
		"cdnFrontdoorEndpointId",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetCdnFrontdoorOriginGroupId(val *string) {
	_jsii_.Set(
		j,
		"cdnFrontdoorOriginGroupId",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetCdnFrontdoorOriginIds(val *[]*string) {
	_jsii_.Set(
		j,
		"cdnFrontdoorOriginIds",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetCdnFrontdoorOriginPath(val *string) {
	_jsii_.Set(
		j,
		"cdnFrontdoorOriginPath",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetCdnFrontdoorRuleSetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"cdnFrontdoorRuleSetIds",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetForwardingProtocol(val *string) {
	_jsii_.Set(
		j,
		"forwardingProtocol",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetHttpsRedirectEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"httpsRedirectEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetPatternsToMatch(val *[]*string) {
	_jsii_.Set(
		j,
		"patternsToMatch",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute) SetSupportedProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"supportedProtocols",
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
func CdnFrontdoorRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CdnFrontdoorRoute_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) PutCache(value *CdnFrontdoorRouteCache) {
	_jsii_.InvokeVoid(
		c,
		"putCache",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) PutTimeouts(value *CdnFrontdoorRouteTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetCache() {
	_jsii_.InvokeVoid(
		c,
		"resetCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetCdnFrontdoorOriginPath() {
	_jsii_.InvokeVoid(
		c,
		"resetCdnFrontdoorOriginPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetCdnFrontdoorRuleSetIds() {
	_jsii_.InvokeVoid(
		c,
		"resetCdnFrontdoorRuleSetIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetForwardingProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetForwardingProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetHttpsRedirectEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsRedirectEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnFrontdoorRouteCache struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#compression_enabled CdnFrontdoorRoute#compression_enabled}.
	CompressionEnabled interface{} `field:"optional" json:"compressionEnabled" yaml:"compressionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#content_types_to_compress CdnFrontdoorRoute#content_types_to_compress}.
	ContentTypesToCompress *[]*string `field:"optional" json:"contentTypesToCompress" yaml:"contentTypesToCompress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#query_string_caching_behavior CdnFrontdoorRoute#query_string_caching_behavior}.
	QueryStringCachingBehavior *string `field:"optional" json:"queryStringCachingBehavior" yaml:"queryStringCachingBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#query_strings CdnFrontdoorRoute#query_strings}.
	QueryStrings *[]*string `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

type CdnFrontdoorRouteCacheOutputReference interface {
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
	CompressionEnabled() interface{}
	SetCompressionEnabled(val interface{})
	CompressionEnabledInput() interface{}
	ContentTypesToCompress() *[]*string
	SetContentTypesToCompress(val *[]*string)
	ContentTypesToCompressInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnFrontdoorRouteCache
	SetInternalValue(val *CdnFrontdoorRouteCache)
	QueryStringCachingBehavior() *string
	SetQueryStringCachingBehavior(val *string)
	QueryStringCachingBehaviorInput() *string
	QueryStrings() *[]*string
	SetQueryStrings(val *[]*string)
	QueryStringsInput() *[]*string
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
	ResetCompressionEnabled()
	ResetContentTypesToCompress()
	ResetQueryStringCachingBehavior()
	ResetQueryStrings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRouteCacheOutputReference
type jsiiProxy_CdnFrontdoorRouteCacheOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) CompressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) CompressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ContentTypesToCompress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentTypesToCompress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ContentTypesToCompressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentTypesToCompressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) InternalValue() *CdnFrontdoorRouteCache {
	var returns *CdnFrontdoorRouteCache
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) QueryStringCachingBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringCachingBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) QueryStringCachingBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringCachingBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) QueryStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) QueryStringsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRouteCacheOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRouteCacheOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnFrontdoorRouteCacheOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteCacheOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRouteCacheOutputReference_Override(c CdnFrontdoorRouteCacheOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteCacheOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetCompressionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"compressionEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetContentTypesToCompress(val *[]*string) {
	_jsii_.Set(
		j,
		"contentTypesToCompress",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetInternalValue(val *CdnFrontdoorRouteCache) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetQueryStringCachingBehavior(val *string) {
	_jsii_.Set(
		j,
		"queryStringCachingBehavior",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetQueryStrings(val *[]*string) {
	_jsii_.Set(
		j,
		"queryStrings",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ResetCompressionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetCompressionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ResetContentTypesToCompress() {
	_jsii_.InvokeVoid(
		c,
		"resetContentTypesToCompress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ResetQueryStringCachingBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCachingBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ResetQueryStrings() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStrings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnFrontdoorRouteConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#cdn_frontdoor_endpoint_id CdnFrontdoorRoute#cdn_frontdoor_endpoint_id}.
	CdnFrontdoorEndpointId *string `field:"required" json:"cdnFrontdoorEndpointId" yaml:"cdnFrontdoorEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#cdn_frontdoor_origin_group_id CdnFrontdoorRoute#cdn_frontdoor_origin_group_id}.
	CdnFrontdoorOriginGroupId *string `field:"required" json:"cdnFrontdoorOriginGroupId" yaml:"cdnFrontdoorOriginGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#cdn_frontdoor_origin_ids CdnFrontdoorRoute#cdn_frontdoor_origin_ids}.
	CdnFrontdoorOriginIds *[]*string `field:"required" json:"cdnFrontdoorOriginIds" yaml:"cdnFrontdoorOriginIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#name CdnFrontdoorRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#patterns_to_match CdnFrontdoorRoute#patterns_to_match}.
	PatternsToMatch *[]*string `field:"required" json:"patternsToMatch" yaml:"patternsToMatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#supported_protocols CdnFrontdoorRoute#supported_protocols}.
	SupportedProtocols *[]*string `field:"required" json:"supportedProtocols" yaml:"supportedProtocols"`
	// cache block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#cache CdnFrontdoorRoute#cache}
	Cache *CdnFrontdoorRouteCache `field:"optional" json:"cache" yaml:"cache"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#cdn_frontdoor_origin_path CdnFrontdoorRoute#cdn_frontdoor_origin_path}.
	CdnFrontdoorOriginPath *string `field:"optional" json:"cdnFrontdoorOriginPath" yaml:"cdnFrontdoorOriginPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#cdn_frontdoor_rule_set_ids CdnFrontdoorRoute#cdn_frontdoor_rule_set_ids}.
	CdnFrontdoorRuleSetIds *[]*string `field:"optional" json:"cdnFrontdoorRuleSetIds" yaml:"cdnFrontdoorRuleSetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#enabled CdnFrontdoorRoute#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#forwarding_protocol CdnFrontdoorRoute#forwarding_protocol}.
	ForwardingProtocol *string `field:"optional" json:"forwardingProtocol" yaml:"forwardingProtocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#https_redirect_enabled CdnFrontdoorRoute#https_redirect_enabled}.
	HttpsRedirectEnabled interface{} `field:"optional" json:"httpsRedirectEnabled" yaml:"httpsRedirectEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#id CdnFrontdoorRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#timeouts CdnFrontdoorRoute#timeouts}
	Timeouts *CdnFrontdoorRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type CdnFrontdoorRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#create CdnFrontdoorRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#delete CdnFrontdoorRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#read CdnFrontdoorRoute#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_route#update CdnFrontdoorRoute#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type CdnFrontdoorRouteTimeoutsOutputReference interface {
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

// The jsii proxy struct for CdnFrontdoorRouteTimeoutsOutputReference
type jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRouteTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRouteTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRouteTimeoutsOutputReference_Override(c CdnFrontdoorRouteTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		c,
		"resetRead",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

