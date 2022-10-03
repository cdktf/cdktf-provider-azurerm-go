package cdnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/cdnendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint azurerm_cdn_endpoint}.
type CdnEndpoint interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentTypesToCompress() *[]*string
	SetContentTypesToCompress(val *[]*string)
	ContentTypesToCompressInput() *[]*string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DeliveryRule() CdnEndpointDeliveryRuleList
	DeliveryRuleInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoFilter() CdnEndpointGeoFilterList
	GeoFilterInput() interface{}
	GlobalDeliveryRule() CdnEndpointGlobalDeliveryRuleOutputReference
	GlobalDeliveryRuleInput() *CdnEndpointGlobalDeliveryRule
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsCompressionEnabled() interface{}
	SetIsCompressionEnabled(val interface{})
	IsCompressionEnabledInput() interface{}
	IsHttpAllowed() interface{}
	SetIsHttpAllowed(val interface{})
	IsHttpAllowedInput() interface{}
	IsHttpsAllowed() interface{}
	SetIsHttpsAllowed(val interface{})
	IsHttpsAllowedInput() interface{}
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
	OptimizationType() *string
	SetOptimizationType(val *string)
	OptimizationTypeInput() *string
	Origin() CdnEndpointOriginList
	OriginHostHeader() *string
	SetOriginHostHeader(val *string)
	OriginHostHeaderInput() *string
	OriginInput() interface{}
	OriginPath() *string
	SetOriginPath(val *string)
	OriginPathInput() *string
	ProbePath() *string
	SetProbePath(val *string)
	ProbePathInput() *string
	ProfileName() *string
	SetProfileName(val *string)
	ProfileNameInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QuerystringCachingBehaviour() *string
	SetQuerystringCachingBehaviour(val *string)
	QuerystringCachingBehaviourInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CdnEndpointTimeoutsOutputReference
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
	PutDeliveryRule(value interface{})
	PutGeoFilter(value interface{})
	PutGlobalDeliveryRule(value *CdnEndpointGlobalDeliveryRule)
	PutOrigin(value interface{})
	PutTimeouts(value *CdnEndpointTimeouts)
	ResetContentTypesToCompress()
	ResetDeliveryRule()
	ResetGeoFilter()
	ResetGlobalDeliveryRule()
	ResetId()
	ResetIsCompressionEnabled()
	ResetIsHttpAllowed()
	ResetIsHttpsAllowed()
	ResetOptimizationType()
	ResetOriginHostHeader()
	ResetOriginPath()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProbePath()
	ResetQuerystringCachingBehaviour()
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

// The jsii proxy struct for CdnEndpoint
type jsiiProxy_CdnEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CdnEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ContentTypesToCompress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentTypesToCompress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ContentTypesToCompressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentTypesToCompressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) DeliveryRule() CdnEndpointDeliveryRuleList {
	var returns CdnEndpointDeliveryRuleList
	_jsii_.Get(
		j,
		"deliveryRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) DeliveryRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) GeoFilter() CdnEndpointGeoFilterList {
	var returns CdnEndpointGeoFilterList
	_jsii_.Get(
		j,
		"geoFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) GeoFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) GlobalDeliveryRule() CdnEndpointGlobalDeliveryRuleOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleOutputReference
	_jsii_.Get(
		j,
		"globalDeliveryRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) GlobalDeliveryRuleInput() *CdnEndpointGlobalDeliveryRule {
	var returns *CdnEndpointGlobalDeliveryRule
	_jsii_.Get(
		j,
		"globalDeliveryRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsCompressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCompressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsCompressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCompressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsHttpAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHttpAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsHttpAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHttpAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsHttpsAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHttpsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsHttpsAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHttpsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OptimizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optimizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OptimizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optimizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Origin() CdnEndpointOriginList {
	var returns CdnEndpointOriginList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ProbePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ProbePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) QuerystringCachingBehaviour() *string {
	var returns *string
	_jsii_.Get(
		j,
		"querystringCachingBehaviour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) QuerystringCachingBehaviourInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"querystringCachingBehaviourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Timeouts() CdnEndpointTimeoutsOutputReference {
	var returns CdnEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint azurerm_cdn_endpoint} Resource.
func NewCdnEndpoint(scope constructs.Construct, id *string, config *CdnEndpointConfig) CdnEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint azurerm_cdn_endpoint} Resource.
func NewCdnEndpoint_Override(c CdnEndpoint, scope constructs.Construct, id *string, config *CdnEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetContentTypesToCompress(val *[]*string) {
	_jsii_.Set(
		j,
		"contentTypesToCompress",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetIsCompressionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"isCompressionEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetIsHttpAllowed(val interface{}) {
	_jsii_.Set(
		j,
		"isHttpAllowed",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetIsHttpsAllowed(val interface{}) {
	_jsii_.Set(
		j,
		"isHttpsAllowed",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetOptimizationType(val *string) {
	_jsii_.Set(
		j,
		"optimizationType",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetOriginHostHeader(val *string) {
	_jsii_.Set(
		j,
		"originHostHeader",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetOriginPath(val *string) {
	_jsii_.Set(
		j,
		"originPath",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetProbePath(val *string) {
	_jsii_.Set(
		j,
		"probePath",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetProfileName(val *string) {
	_jsii_.Set(
		j,
		"profileName",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetQuerystringCachingBehaviour(val *string) {
	_jsii_.Set(
		j,
		"querystringCachingBehaviour",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint) SetTags(val *map[string]*string) {
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
func CdnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CdnEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CdnEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CdnEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutDeliveryRule(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putDeliveryRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutGeoFilter(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putGeoFilter",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutGlobalDeliveryRule(value *CdnEndpointGlobalDeliveryRule) {
	_jsii_.InvokeVoid(
		c,
		"putGlobalDeliveryRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutOrigin(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putOrigin",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutTimeouts(value *CdnEndpointTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetContentTypesToCompress() {
	_jsii_.InvokeVoid(
		c,
		"resetContentTypesToCompress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetDeliveryRule() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetGeoFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetGeoFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetGlobalDeliveryRule() {
	_jsii_.InvokeVoid(
		c,
		"resetGlobalDeliveryRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetIsCompressionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetIsCompressionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetIsHttpAllowed() {
	_jsii_.InvokeVoid(
		c,
		"resetIsHttpAllowed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetIsHttpsAllowed() {
	_jsii_.InvokeVoid(
		c,
		"resetIsHttpsAllowed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetOptimizationType() {
	_jsii_.InvokeVoid(
		c,
		"resetOptimizationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetOriginHostHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginHostHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetOriginPath() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetProbePath() {
	_jsii_.InvokeVoid(
		c,
		"resetProbePath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetQuerystringCachingBehaviour() {
	_jsii_.InvokeVoid(
		c,
		"resetQuerystringCachingBehaviour",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#location CdnEndpoint#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#origin CdnEndpoint#origin}
	Origin interface{} `field:"required" json:"origin" yaml:"origin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#profile_name CdnEndpoint#profile_name}.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#resource_group_name CdnEndpoint#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#content_types_to_compress CdnEndpoint#content_types_to_compress}.
	ContentTypesToCompress *[]*string `field:"optional" json:"contentTypesToCompress" yaml:"contentTypesToCompress"`
	// delivery_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#delivery_rule CdnEndpoint#delivery_rule}
	DeliveryRule interface{} `field:"optional" json:"deliveryRule" yaml:"deliveryRule"`
	// geo_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#geo_filter CdnEndpoint#geo_filter}
	GeoFilter interface{} `field:"optional" json:"geoFilter" yaml:"geoFilter"`
	// global_delivery_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#global_delivery_rule CdnEndpoint#global_delivery_rule}
	GlobalDeliveryRule *CdnEndpointGlobalDeliveryRule `field:"optional" json:"globalDeliveryRule" yaml:"globalDeliveryRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#id CdnEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#is_compression_enabled CdnEndpoint#is_compression_enabled}.
	IsCompressionEnabled interface{} `field:"optional" json:"isCompressionEnabled" yaml:"isCompressionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#is_http_allowed CdnEndpoint#is_http_allowed}.
	IsHttpAllowed interface{} `field:"optional" json:"isHttpAllowed" yaml:"isHttpAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#is_https_allowed CdnEndpoint#is_https_allowed}.
	IsHttpsAllowed interface{} `field:"optional" json:"isHttpsAllowed" yaml:"isHttpsAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#optimization_type CdnEndpoint#optimization_type}.
	OptimizationType *string `field:"optional" json:"optimizationType" yaml:"optimizationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#origin_host_header CdnEndpoint#origin_host_header}.
	OriginHostHeader *string `field:"optional" json:"originHostHeader" yaml:"originHostHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#origin_path CdnEndpoint#origin_path}.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#probe_path CdnEndpoint#probe_path}.
	ProbePath *string `field:"optional" json:"probePath" yaml:"probePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#querystring_caching_behaviour CdnEndpoint#querystring_caching_behaviour}.
	QuerystringCachingBehaviour *string `field:"optional" json:"querystringCachingBehaviour" yaml:"querystringCachingBehaviour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#tags CdnEndpoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#timeouts CdnEndpoint#timeouts}
	Timeouts *CdnEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type CdnEndpointDeliveryRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#order CdnEndpoint#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
	// cache_expiration_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#cache_expiration_action CdnEndpoint#cache_expiration_action}
	CacheExpirationAction *CdnEndpointDeliveryRuleCacheExpirationAction `field:"optional" json:"cacheExpirationAction" yaml:"cacheExpirationAction"`
	// cache_key_query_string_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#cache_key_query_string_action CdnEndpoint#cache_key_query_string_action}
	CacheKeyQueryStringAction *CdnEndpointDeliveryRuleCacheKeyQueryStringAction `field:"optional" json:"cacheKeyQueryStringAction" yaml:"cacheKeyQueryStringAction"`
	// cookies_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#cookies_condition CdnEndpoint#cookies_condition}
	CookiesCondition interface{} `field:"optional" json:"cookiesCondition" yaml:"cookiesCondition"`
	// device_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#device_condition CdnEndpoint#device_condition}
	DeviceCondition *CdnEndpointDeliveryRuleDeviceCondition `field:"optional" json:"deviceCondition" yaml:"deviceCondition"`
	// http_version_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#http_version_condition CdnEndpoint#http_version_condition}
	HttpVersionCondition interface{} `field:"optional" json:"httpVersionCondition" yaml:"httpVersionCondition"`
	// modify_request_header_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#modify_request_header_action CdnEndpoint#modify_request_header_action}
	ModifyRequestHeaderAction interface{} `field:"optional" json:"modifyRequestHeaderAction" yaml:"modifyRequestHeaderAction"`
	// modify_response_header_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#modify_response_header_action CdnEndpoint#modify_response_header_action}
	ModifyResponseHeaderAction interface{} `field:"optional" json:"modifyResponseHeaderAction" yaml:"modifyResponseHeaderAction"`
	// post_arg_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#post_arg_condition CdnEndpoint#post_arg_condition}
	PostArgCondition interface{} `field:"optional" json:"postArgCondition" yaml:"postArgCondition"`
	// query_string_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#query_string_condition CdnEndpoint#query_string_condition}
	QueryStringCondition interface{} `field:"optional" json:"queryStringCondition" yaml:"queryStringCondition"`
	// remote_address_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#remote_address_condition CdnEndpoint#remote_address_condition}
	RemoteAddressCondition interface{} `field:"optional" json:"remoteAddressCondition" yaml:"remoteAddressCondition"`
	// request_body_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#request_body_condition CdnEndpoint#request_body_condition}
	RequestBodyCondition interface{} `field:"optional" json:"requestBodyCondition" yaml:"requestBodyCondition"`
	// request_header_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#request_header_condition CdnEndpoint#request_header_condition}
	RequestHeaderCondition interface{} `field:"optional" json:"requestHeaderCondition" yaml:"requestHeaderCondition"`
	// request_method_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#request_method_condition CdnEndpoint#request_method_condition}
	RequestMethodCondition *CdnEndpointDeliveryRuleRequestMethodCondition `field:"optional" json:"requestMethodCondition" yaml:"requestMethodCondition"`
	// request_scheme_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#request_scheme_condition CdnEndpoint#request_scheme_condition}
	RequestSchemeCondition *CdnEndpointDeliveryRuleRequestSchemeCondition `field:"optional" json:"requestSchemeCondition" yaml:"requestSchemeCondition"`
	// request_uri_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#request_uri_condition CdnEndpoint#request_uri_condition}
	RequestUriCondition interface{} `field:"optional" json:"requestUriCondition" yaml:"requestUriCondition"`
	// url_file_extension_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#url_file_extension_condition CdnEndpoint#url_file_extension_condition}
	UrlFileExtensionCondition interface{} `field:"optional" json:"urlFileExtensionCondition" yaml:"urlFileExtensionCondition"`
	// url_file_name_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#url_file_name_condition CdnEndpoint#url_file_name_condition}
	UrlFileNameCondition interface{} `field:"optional" json:"urlFileNameCondition" yaml:"urlFileNameCondition"`
	// url_path_condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#url_path_condition CdnEndpoint#url_path_condition}
	UrlPathCondition interface{} `field:"optional" json:"urlPathCondition" yaml:"urlPathCondition"`
	// url_redirect_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#url_redirect_action CdnEndpoint#url_redirect_action}
	UrlRedirectAction *CdnEndpointDeliveryRuleUrlRedirectAction `field:"optional" json:"urlRedirectAction" yaml:"urlRedirectAction"`
	// url_rewrite_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#url_rewrite_action CdnEndpoint#url_rewrite_action}
	UrlRewriteAction *CdnEndpointDeliveryRuleUrlRewriteAction `field:"optional" json:"urlRewriteAction" yaml:"urlRewriteAction"`
}

type CdnEndpointDeliveryRuleCacheExpirationAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#duration CdnEndpoint#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

type CdnEndpointDeliveryRuleCacheExpirationActionOutputReference interface {
	cdktf.ComplexObject
	Behavior() *string
	SetBehavior(val *string)
	BehaviorInput() *string
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnEndpointDeliveryRuleCacheExpirationAction
	SetInternalValue(val *CdnEndpointDeliveryRuleCacheExpirationAction)
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
	ResetDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleCacheExpirationActionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) InternalValue() *CdnEndpointDeliveryRuleCacheExpirationAction {
	var returns *CdnEndpointDeliveryRuleCacheExpirationAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleCacheExpirationActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointDeliveryRuleCacheExpirationActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCacheExpirationActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleCacheExpirationActionOutputReference_Override(c CdnEndpointDeliveryRuleCacheExpirationActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCacheExpirationActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) SetBehavior(val *string) {
	_jsii_.Set(
		j,
		"behavior",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) SetDuration(val *string) {
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) SetInternalValue(val *CdnEndpointDeliveryRuleCacheExpirationAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleCacheKeyQueryStringAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#parameters CdnEndpoint#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

type CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference interface {
	cdktf.ComplexObject
	Behavior() *string
	SetBehavior(val *string)
	BehaviorInput() *string
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
	InternalValue() *CdnEndpointDeliveryRuleCacheKeyQueryStringAction
	SetInternalValue(val *CdnEndpointDeliveryRuleCacheKeyQueryStringAction)
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
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
	ResetParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) InternalValue() *CdnEndpointDeliveryRuleCacheKeyQueryStringAction {
	var returns *CdnEndpointDeliveryRuleCacheKeyQueryStringAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference_Override(c CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) SetBehavior(val *string) {
	_jsii_.Set(
		j,
		"behavior",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) SetInternalValue(val *CdnEndpointDeliveryRuleCacheKeyQueryStringAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) SetParameters(val *string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleCookiesCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#selector CdnEndpoint#selector}.
	Selector *string `field:"required" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRuleCookiesConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleCookiesConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleCookiesConditionList
type jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleCookiesConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleCookiesConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCookiesConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleCookiesConditionList_Override(c CdnEndpointDeliveryRuleCookiesConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCookiesConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) Get(index *float64) CdnEndpointDeliveryRuleCookiesConditionOutputReference {
	var returns CdnEndpointDeliveryRuleCookiesConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleCookiesConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Selector() *string
	SetSelector(val *string)
	SelectorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleCookiesConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) Selector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleCookiesConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleCookiesConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCookiesConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleCookiesConditionOutputReference_Override(c CdnEndpointDeliveryRuleCookiesConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCookiesConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetSelector(val *string) {
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCookiesConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleDeviceCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"required" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

type CdnEndpointDeliveryRuleDeviceConditionOutputReference interface {
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
	InternalValue() *CdnEndpointDeliveryRuleDeviceCondition
	SetInternalValue(val *CdnEndpointDeliveryRuleDeviceCondition)
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
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
	ResetNegateCondition()
	ResetOperator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleDeviceConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) InternalValue() *CdnEndpointDeliveryRuleDeviceCondition {
	var returns *CdnEndpointDeliveryRuleDeviceCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleDeviceConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointDeliveryRuleDeviceConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleDeviceConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleDeviceConditionOutputReference_Override(c CdnEndpointDeliveryRuleDeviceConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleDeviceConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) SetInternalValue(val *CdnEndpointDeliveryRuleDeviceCondition) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		c,
		"resetOperator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleDeviceConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleHttpVersionCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"required" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

type CdnEndpointDeliveryRuleHttpVersionConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleHttpVersionConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleHttpVersionConditionList
type jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleHttpVersionConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleHttpVersionConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleHttpVersionConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleHttpVersionConditionList_Override(c CdnEndpointDeliveryRuleHttpVersionConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleHttpVersionConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) Get(index *float64) CdnEndpointDeliveryRuleHttpVersionConditionOutputReference {
	var returns CdnEndpointDeliveryRuleHttpVersionConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleHttpVersionConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
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
	ResetNegateCondition()
	ResetOperator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleHttpVersionConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleHttpVersionConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleHttpVersionConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleHttpVersionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleHttpVersionConditionOutputReference_Override(c CdnEndpointDeliveryRuleHttpVersionConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleHttpVersionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		c,
		"resetOperator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleHttpVersionConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleList
type jsiiProxy_CdnEndpointDeliveryRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleList_Override(c CdnEndpointDeliveryRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleList) Get(index *float64) CdnEndpointDeliveryRuleOutputReference {
	var returns CdnEndpointDeliveryRuleOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleModifyRequestHeaderAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#action CdnEndpoint#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#value CdnEndpoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

type CdnEndpointDeliveryRuleModifyRequestHeaderActionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleModifyRequestHeaderActionList
type jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleModifyRequestHeaderActionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleModifyRequestHeaderActionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleModifyRequestHeaderActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleModifyRequestHeaderActionList_Override(c CdnEndpointDeliveryRuleModifyRequestHeaderActionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleModifyRequestHeaderActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) Get(index *float64) CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference {
	var returns CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference_Override(c CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		c,
		"resetValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyRequestHeaderActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleModifyResponseHeaderAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#action CdnEndpoint#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#value CdnEndpoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

type CdnEndpointDeliveryRuleModifyResponseHeaderActionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleModifyResponseHeaderActionList
type jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleModifyResponseHeaderActionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleModifyResponseHeaderActionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleModifyResponseHeaderActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleModifyResponseHeaderActionList_Override(c CdnEndpointDeliveryRuleModifyResponseHeaderActionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleModifyResponseHeaderActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) Get(index *float64) CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference {
	var returns CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference_Override(c CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		c,
		"resetValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleModifyResponseHeaderActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleOutputReference interface {
	cdktf.ComplexObject
	CacheExpirationAction() CdnEndpointDeliveryRuleCacheExpirationActionOutputReference
	CacheExpirationActionInput() *CdnEndpointDeliveryRuleCacheExpirationAction
	CacheKeyQueryStringAction() CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference
	CacheKeyQueryStringActionInput() *CdnEndpointDeliveryRuleCacheKeyQueryStringAction
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
	CookiesCondition() CdnEndpointDeliveryRuleCookiesConditionList
	CookiesConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeviceCondition() CdnEndpointDeliveryRuleDeviceConditionOutputReference
	DeviceConditionInput() *CdnEndpointDeliveryRuleDeviceCondition
	// Experimental.
	Fqn() *string
	HttpVersionCondition() CdnEndpointDeliveryRuleHttpVersionConditionList
	HttpVersionConditionInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModifyRequestHeaderAction() CdnEndpointDeliveryRuleModifyRequestHeaderActionList
	ModifyRequestHeaderActionInput() interface{}
	ModifyResponseHeaderAction() CdnEndpointDeliveryRuleModifyResponseHeaderActionList
	ModifyResponseHeaderActionInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	PostArgCondition() CdnEndpointDeliveryRulePostArgConditionList
	PostArgConditionInput() interface{}
	QueryStringCondition() CdnEndpointDeliveryRuleQueryStringConditionList
	QueryStringConditionInput() interface{}
	RemoteAddressCondition() CdnEndpointDeliveryRuleRemoteAddressConditionList
	RemoteAddressConditionInput() interface{}
	RequestBodyCondition() CdnEndpointDeliveryRuleRequestBodyConditionList
	RequestBodyConditionInput() interface{}
	RequestHeaderCondition() CdnEndpointDeliveryRuleRequestHeaderConditionList
	RequestHeaderConditionInput() interface{}
	RequestMethodCondition() CdnEndpointDeliveryRuleRequestMethodConditionOutputReference
	RequestMethodConditionInput() *CdnEndpointDeliveryRuleRequestMethodCondition
	RequestSchemeCondition() CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference
	RequestSchemeConditionInput() *CdnEndpointDeliveryRuleRequestSchemeCondition
	RequestUriCondition() CdnEndpointDeliveryRuleRequestUriConditionList
	RequestUriConditionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlFileExtensionCondition() CdnEndpointDeliveryRuleUrlFileExtensionConditionList
	UrlFileExtensionConditionInput() interface{}
	UrlFileNameCondition() CdnEndpointDeliveryRuleUrlFileNameConditionList
	UrlFileNameConditionInput() interface{}
	UrlPathCondition() CdnEndpointDeliveryRuleUrlPathConditionList
	UrlPathConditionInput() interface{}
	UrlRedirectAction() CdnEndpointDeliveryRuleUrlRedirectActionOutputReference
	UrlRedirectActionInput() *CdnEndpointDeliveryRuleUrlRedirectAction
	UrlRewriteAction() CdnEndpointDeliveryRuleUrlRewriteActionOutputReference
	UrlRewriteActionInput() *CdnEndpointDeliveryRuleUrlRewriteAction
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
	PutCacheExpirationAction(value *CdnEndpointDeliveryRuleCacheExpirationAction)
	PutCacheKeyQueryStringAction(value *CdnEndpointDeliveryRuleCacheKeyQueryStringAction)
	PutCookiesCondition(value interface{})
	PutDeviceCondition(value *CdnEndpointDeliveryRuleDeviceCondition)
	PutHttpVersionCondition(value interface{})
	PutModifyRequestHeaderAction(value interface{})
	PutModifyResponseHeaderAction(value interface{})
	PutPostArgCondition(value interface{})
	PutQueryStringCondition(value interface{})
	PutRemoteAddressCondition(value interface{})
	PutRequestBodyCondition(value interface{})
	PutRequestHeaderCondition(value interface{})
	PutRequestMethodCondition(value *CdnEndpointDeliveryRuleRequestMethodCondition)
	PutRequestSchemeCondition(value *CdnEndpointDeliveryRuleRequestSchemeCondition)
	PutRequestUriCondition(value interface{})
	PutUrlFileExtensionCondition(value interface{})
	PutUrlFileNameCondition(value interface{})
	PutUrlPathCondition(value interface{})
	PutUrlRedirectAction(value *CdnEndpointDeliveryRuleUrlRedirectAction)
	PutUrlRewriteAction(value *CdnEndpointDeliveryRuleUrlRewriteAction)
	ResetCacheExpirationAction()
	ResetCacheKeyQueryStringAction()
	ResetCookiesCondition()
	ResetDeviceCondition()
	ResetHttpVersionCondition()
	ResetModifyRequestHeaderAction()
	ResetModifyResponseHeaderAction()
	ResetPostArgCondition()
	ResetQueryStringCondition()
	ResetRemoteAddressCondition()
	ResetRequestBodyCondition()
	ResetRequestHeaderCondition()
	ResetRequestMethodCondition()
	ResetRequestSchemeCondition()
	ResetRequestUriCondition()
	ResetUrlFileExtensionCondition()
	ResetUrlFileNameCondition()
	ResetUrlPathCondition()
	ResetUrlRedirectAction()
	ResetUrlRewriteAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CacheExpirationAction() CdnEndpointDeliveryRuleCacheExpirationActionOutputReference {
	var returns CdnEndpointDeliveryRuleCacheExpirationActionOutputReference
	_jsii_.Get(
		j,
		"cacheExpirationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CacheExpirationActionInput() *CdnEndpointDeliveryRuleCacheExpirationAction {
	var returns *CdnEndpointDeliveryRuleCacheExpirationAction
	_jsii_.Get(
		j,
		"cacheExpirationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CacheKeyQueryStringAction() CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference {
	var returns CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference
	_jsii_.Get(
		j,
		"cacheKeyQueryStringAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CacheKeyQueryStringActionInput() *CdnEndpointDeliveryRuleCacheKeyQueryStringAction {
	var returns *CdnEndpointDeliveryRuleCacheKeyQueryStringAction
	_jsii_.Get(
		j,
		"cacheKeyQueryStringActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CookiesCondition() CdnEndpointDeliveryRuleCookiesConditionList {
	var returns CdnEndpointDeliveryRuleCookiesConditionList
	_jsii_.Get(
		j,
		"cookiesCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CookiesConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) DeviceCondition() CdnEndpointDeliveryRuleDeviceConditionOutputReference {
	var returns CdnEndpointDeliveryRuleDeviceConditionOutputReference
	_jsii_.Get(
		j,
		"deviceCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) DeviceConditionInput() *CdnEndpointDeliveryRuleDeviceCondition {
	var returns *CdnEndpointDeliveryRuleDeviceCondition
	_jsii_.Get(
		j,
		"deviceConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) HttpVersionCondition() CdnEndpointDeliveryRuleHttpVersionConditionList {
	var returns CdnEndpointDeliveryRuleHttpVersionConditionList
	_jsii_.Get(
		j,
		"httpVersionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) HttpVersionConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpVersionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ModifyRequestHeaderAction() CdnEndpointDeliveryRuleModifyRequestHeaderActionList {
	var returns CdnEndpointDeliveryRuleModifyRequestHeaderActionList
	_jsii_.Get(
		j,
		"modifyRequestHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ModifyRequestHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyRequestHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ModifyResponseHeaderAction() CdnEndpointDeliveryRuleModifyResponseHeaderActionList {
	var returns CdnEndpointDeliveryRuleModifyResponseHeaderActionList
	_jsii_.Get(
		j,
		"modifyResponseHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ModifyResponseHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyResponseHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PostArgCondition() CdnEndpointDeliveryRulePostArgConditionList {
	var returns CdnEndpointDeliveryRulePostArgConditionList
	_jsii_.Get(
		j,
		"postArgCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PostArgConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postArgConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) QueryStringCondition() CdnEndpointDeliveryRuleQueryStringConditionList {
	var returns CdnEndpointDeliveryRuleQueryStringConditionList
	_jsii_.Get(
		j,
		"queryStringCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) QueryStringConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RemoteAddressCondition() CdnEndpointDeliveryRuleRemoteAddressConditionList {
	var returns CdnEndpointDeliveryRuleRemoteAddressConditionList
	_jsii_.Get(
		j,
		"remoteAddressCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RemoteAddressConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteAddressConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestBodyCondition() CdnEndpointDeliveryRuleRequestBodyConditionList {
	var returns CdnEndpointDeliveryRuleRequestBodyConditionList
	_jsii_.Get(
		j,
		"requestBodyCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestBodyConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestHeaderCondition() CdnEndpointDeliveryRuleRequestHeaderConditionList {
	var returns CdnEndpointDeliveryRuleRequestHeaderConditionList
	_jsii_.Get(
		j,
		"requestHeaderCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestHeaderConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestMethodCondition() CdnEndpointDeliveryRuleRequestMethodConditionOutputReference {
	var returns CdnEndpointDeliveryRuleRequestMethodConditionOutputReference
	_jsii_.Get(
		j,
		"requestMethodCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestMethodConditionInput() *CdnEndpointDeliveryRuleRequestMethodCondition {
	var returns *CdnEndpointDeliveryRuleRequestMethodCondition
	_jsii_.Get(
		j,
		"requestMethodConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestSchemeCondition() CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference {
	var returns CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference
	_jsii_.Get(
		j,
		"requestSchemeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestSchemeConditionInput() *CdnEndpointDeliveryRuleRequestSchemeCondition {
	var returns *CdnEndpointDeliveryRuleRequestSchemeCondition
	_jsii_.Get(
		j,
		"requestSchemeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestUriCondition() CdnEndpointDeliveryRuleRequestUriConditionList {
	var returns CdnEndpointDeliveryRuleRequestUriConditionList
	_jsii_.Get(
		j,
		"requestUriCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestUriConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestUriConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlFileExtensionCondition() CdnEndpointDeliveryRuleUrlFileExtensionConditionList {
	var returns CdnEndpointDeliveryRuleUrlFileExtensionConditionList
	_jsii_.Get(
		j,
		"urlFileExtensionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlFileExtensionConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlFileExtensionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlFileNameCondition() CdnEndpointDeliveryRuleUrlFileNameConditionList {
	var returns CdnEndpointDeliveryRuleUrlFileNameConditionList
	_jsii_.Get(
		j,
		"urlFileNameCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlFileNameConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlFileNameConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlPathCondition() CdnEndpointDeliveryRuleUrlPathConditionList {
	var returns CdnEndpointDeliveryRuleUrlPathConditionList
	_jsii_.Get(
		j,
		"urlPathCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlPathConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlPathConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlRedirectAction() CdnEndpointDeliveryRuleUrlRedirectActionOutputReference {
	var returns CdnEndpointDeliveryRuleUrlRedirectActionOutputReference
	_jsii_.Get(
		j,
		"urlRedirectAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlRedirectActionInput() *CdnEndpointDeliveryRuleUrlRedirectAction {
	var returns *CdnEndpointDeliveryRuleUrlRedirectAction
	_jsii_.Get(
		j,
		"urlRedirectActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlRewriteAction() CdnEndpointDeliveryRuleUrlRewriteActionOutputReference {
	var returns CdnEndpointDeliveryRuleUrlRewriteActionOutputReference
	_jsii_.Get(
		j,
		"urlRewriteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlRewriteActionInput() *CdnEndpointDeliveryRuleUrlRewriteAction {
	var returns *CdnEndpointDeliveryRuleUrlRewriteAction
	_jsii_.Get(
		j,
		"urlRewriteActionInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleOutputReference_Override(c CdnEndpointDeliveryRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) SetOrder(val *float64) {
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutCacheExpirationAction(value *CdnEndpointDeliveryRuleCacheExpirationAction) {
	_jsii_.InvokeVoid(
		c,
		"putCacheExpirationAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutCacheKeyQueryStringAction(value *CdnEndpointDeliveryRuleCacheKeyQueryStringAction) {
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyQueryStringAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutCookiesCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putCookiesCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutDeviceCondition(value *CdnEndpointDeliveryRuleDeviceCondition) {
	_jsii_.InvokeVoid(
		c,
		"putDeviceCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutHttpVersionCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putHttpVersionCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutModifyRequestHeaderAction(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putModifyRequestHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutModifyResponseHeaderAction(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putModifyResponseHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutPostArgCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putPostArgCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutQueryStringCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putQueryStringCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRemoteAddressCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putRemoteAddressCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestBodyCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putRequestBodyCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestHeaderCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putRequestHeaderCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestMethodCondition(value *CdnEndpointDeliveryRuleRequestMethodCondition) {
	_jsii_.InvokeVoid(
		c,
		"putRequestMethodCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestSchemeCondition(value *CdnEndpointDeliveryRuleRequestSchemeCondition) {
	_jsii_.InvokeVoid(
		c,
		"putRequestSchemeCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestUriCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putRequestUriCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlFileExtensionCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putUrlFileExtensionCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlFileNameCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putUrlFileNameCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlPathCondition(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putUrlPathCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlRedirectAction(value *CdnEndpointDeliveryRuleUrlRedirectAction) {
	_jsii_.InvokeVoid(
		c,
		"putUrlRedirectAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlRewriteAction(value *CdnEndpointDeliveryRuleUrlRewriteAction) {
	_jsii_.InvokeVoid(
		c,
		"putUrlRewriteAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetCacheExpirationAction() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheExpirationAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetCacheKeyQueryStringAction() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyQueryStringAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetCookiesCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetCookiesCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetDeviceCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetHttpVersionCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersionCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetModifyRequestHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyRequestHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetModifyResponseHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyResponseHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetPostArgCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetPostArgCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetQueryStringCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRemoteAddressCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteAddressCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestBodyCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestBodyCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestHeaderCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeaderCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestMethodCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestMethodCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestSchemeCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestSchemeCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestUriCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestUriCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlFileExtensionCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlFileExtensionCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlFileNameCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlFileNameCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlPathCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlPathCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlRedirectAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRedirectAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlRewriteAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewriteAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRulePostArgCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#selector CdnEndpoint#selector}.
	Selector *string `field:"required" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRulePostArgConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRulePostArgConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRulePostArgConditionList
type jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRulePostArgConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRulePostArgConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRulePostArgConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRulePostArgConditionList_Override(c CdnEndpointDeliveryRulePostArgConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRulePostArgConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) Get(index *float64) CdnEndpointDeliveryRulePostArgConditionOutputReference {
	var returns CdnEndpointDeliveryRulePostArgConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRulePostArgConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Selector() *string
	SetSelector(val *string)
	SelectorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRulePostArgConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) Selector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRulePostArgConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRulePostArgConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRulePostArgConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRulePostArgConditionOutputReference_Override(c CdnEndpointDeliveryRulePostArgConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRulePostArgConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetSelector(val *string) {
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleQueryStringCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRuleQueryStringConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleQueryStringConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleQueryStringConditionList
type jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleQueryStringConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleQueryStringConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleQueryStringConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleQueryStringConditionList_Override(c CdnEndpointDeliveryRuleQueryStringConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleQueryStringConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) Get(index *float64) CdnEndpointDeliveryRuleQueryStringConditionOutputReference {
	var returns CdnEndpointDeliveryRuleQueryStringConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleQueryStringConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleQueryStringConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleQueryStringConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleQueryStringConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleQueryStringConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleQueryStringConditionOutputReference_Override(c CdnEndpointDeliveryRuleQueryStringConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleQueryStringConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleQueryStringConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRemoteAddressCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
}

type CdnEndpointDeliveryRuleRemoteAddressConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRemoteAddressConditionList
type jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRemoteAddressConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleRemoteAddressConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRemoteAddressConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRemoteAddressConditionList_Override(c CdnEndpointDeliveryRuleRemoteAddressConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRemoteAddressConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) Get(index *float64) CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference {
	var returns CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
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
	ResetMatchValues()
	ResetNegateCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRemoteAddressConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRemoteAddressConditionOutputReference_Override(c CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRemoteAddressConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRequestBodyCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRuleRequestBodyConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleRequestBodyConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRequestBodyConditionList
type jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRequestBodyConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleRequestBodyConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestBodyConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRequestBodyConditionList_Override(c CdnEndpointDeliveryRuleRequestBodyConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestBodyConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) Get(index *float64) CdnEndpointDeliveryRuleRequestBodyConditionOutputReference {
	var returns CdnEndpointDeliveryRuleRequestBodyConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRequestBodyConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRequestBodyConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRequestBodyConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleRequestBodyConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestBodyConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRequestBodyConditionOutputReference_Override(c CdnEndpointDeliveryRuleRequestBodyConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestBodyConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRequestHeaderCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#selector CdnEndpoint#selector}.
	Selector *string `field:"required" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRuleRequestHeaderConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRequestHeaderConditionList
type jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRequestHeaderConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleRequestHeaderConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestHeaderConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRequestHeaderConditionList_Override(c CdnEndpointDeliveryRuleRequestHeaderConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestHeaderConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) Get(index *float64) CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference {
	var returns CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Selector() *string
	SetSelector(val *string)
	SelectorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) Selector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRequestHeaderConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRequestHeaderConditionOutputReference_Override(c CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetSelector(val *string) {
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestHeaderConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRequestMethodCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"required" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

type CdnEndpointDeliveryRuleRequestMethodConditionOutputReference interface {
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
	InternalValue() *CdnEndpointDeliveryRuleRequestMethodCondition
	SetInternalValue(val *CdnEndpointDeliveryRuleRequestMethodCondition)
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
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
	ResetNegateCondition()
	ResetOperator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRequestMethodConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) InternalValue() *CdnEndpointDeliveryRuleRequestMethodCondition {
	var returns *CdnEndpointDeliveryRuleRequestMethodCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRequestMethodConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointDeliveryRuleRequestMethodConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestMethodConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRequestMethodConditionOutputReference_Override(c CdnEndpointDeliveryRuleRequestMethodConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestMethodConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) SetInternalValue(val *CdnEndpointDeliveryRuleRequestMethodCondition) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		c,
		"resetOperator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestMethodConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRequestSchemeCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"required" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

type CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference interface {
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
	InternalValue() *CdnEndpointDeliveryRuleRequestSchemeCondition
	SetInternalValue(val *CdnEndpointDeliveryRuleRequestSchemeCondition)
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
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
	ResetNegateCondition()
	ResetOperator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) InternalValue() *CdnEndpointDeliveryRuleRequestSchemeCondition {
	var returns *CdnEndpointDeliveryRuleRequestSchemeCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRequestSchemeConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRequestSchemeConditionOutputReference_Override(c CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) SetInternalValue(val *CdnEndpointDeliveryRuleRequestSchemeCondition) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		c,
		"resetOperator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRequestUriCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRuleRequestUriConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleRequestUriConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRequestUriConditionList
type jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRequestUriConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleRequestUriConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestUriConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRequestUriConditionList_Override(c CdnEndpointDeliveryRuleRequestUriConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestUriConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) Get(index *float64) CdnEndpointDeliveryRuleRequestUriConditionOutputReference {
	var returns CdnEndpointDeliveryRuleRequestUriConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleRequestUriConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleRequestUriConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleRequestUriConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleRequestUriConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestUriConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleRequestUriConditionOutputReference_Override(c CdnEndpointDeliveryRuleRequestUriConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleRequestUriConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestUriConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleUrlFileExtensionCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRuleUrlFileExtensionConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleUrlFileExtensionConditionList
type jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleUrlFileExtensionConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleUrlFileExtensionConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlFileExtensionConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleUrlFileExtensionConditionList_Override(c CdnEndpointDeliveryRuleUrlFileExtensionConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlFileExtensionConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) Get(index *float64) CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference {
	var returns CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference_Override(c CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileExtensionConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleUrlFileNameCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRuleUrlFileNameConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleUrlFileNameConditionList
type jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleUrlFileNameConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleUrlFileNameConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlFileNameConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleUrlFileNameConditionList_Override(c CdnEndpointDeliveryRuleUrlFileNameConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlFileNameConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) Get(index *float64) CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference {
	var returns CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleUrlFileNameConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleUrlFileNameConditionOutputReference_Override(c CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlFileNameConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleUrlPathCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

type CdnEndpointDeliveryRuleUrlPathConditionList interface {
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
	Get(index *float64) CdnEndpointDeliveryRuleUrlPathConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleUrlPathConditionList
type jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleUrlPathConditionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointDeliveryRuleUrlPathConditionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlPathConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleUrlPathConditionList_Override(c CdnEndpointDeliveryRuleUrlPathConditionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlPathConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) Get(index *float64) CdnEndpointDeliveryRuleUrlPathConditionOutputReference {
	var returns CdnEndpointDeliveryRuleUrlPathConditionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleUrlPathConditionOutputReference interface {
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
	MatchValues() *[]*string
	SetMatchValues(val *[]*string)
	MatchValuesInput() *[]*string
	NegateCondition() interface{}
	SetNegateCondition(val interface{})
	NegateConditionInput() interface{}
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transforms() *[]*string
	SetTransforms(val *[]*string)
	TransformsInput() *[]*string
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
	ResetMatchValues()
	ResetNegateCondition()
	ResetTransforms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleUrlPathConditionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) MatchValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) MatchValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) NegateCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) NegateConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) TransformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transformsInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleUrlPathConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleUrlPathConditionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlPathConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleUrlPathConditionOutputReference_Override(c CdnEndpointDeliveryRuleUrlPathConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlPathConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetMatchValues(val *[]*string) {
	_jsii_.Set(
		j,
		"matchValues",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetNegateCondition(val interface{}) {
	_jsii_.Set(
		j,
		"negateCondition",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetOperator(val *string) {
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) ResetMatchValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) ResetNegateCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetNegateCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) ResetTransforms() {
	_jsii_.InvokeVoid(
		c,
		"resetTransforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlPathConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleUrlRedirectAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#redirect_type CdnEndpoint#redirect_type}.
	RedirectType *string `field:"required" json:"redirectType" yaml:"redirectType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#fragment CdnEndpoint#fragment}.
	Fragment *string `field:"optional" json:"fragment" yaml:"fragment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#hostname CdnEndpoint#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#path CdnEndpoint#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#protocol CdnEndpoint#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#query_string CdnEndpoint#query_string}.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
}

type CdnEndpointDeliveryRuleUrlRedirectActionOutputReference interface {
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
	Fragment() *string
	SetFragment(val *string)
	FragmentInput() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InternalValue() *CdnEndpointDeliveryRuleUrlRedirectAction
	SetInternalValue(val *CdnEndpointDeliveryRuleUrlRedirectAction)
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
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
	ResetFragment()
	ResetHostname()
	ResetPath()
	ResetProtocol()
	ResetQueryString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleUrlRedirectActionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) Fragment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) FragmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) InternalValue() *CdnEndpointDeliveryRuleUrlRedirectAction {
	var returns *CdnEndpointDeliveryRuleUrlRedirectAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) RedirectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) RedirectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleUrlRedirectActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointDeliveryRuleUrlRedirectActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleUrlRedirectActionOutputReference_Override(c CdnEndpointDeliveryRuleUrlRedirectActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetFragment(val *string) {
	_jsii_.Set(
		j,
		"fragment",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetInternalValue(val *CdnEndpointDeliveryRuleUrlRedirectAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetQueryString(val *string) {
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetRedirectType(val *string) {
	_jsii_.Set(
		j,
		"redirectType",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ResetFragment() {
	_jsii_.InvokeVoid(
		c,
		"resetFragment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRedirectActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointDeliveryRuleUrlRewriteAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#destination CdnEndpoint#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#source_pattern CdnEndpoint#source_pattern}.
	SourcePattern *string `field:"required" json:"sourcePattern" yaml:"sourcePattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#preserve_unmatched_path CdnEndpoint#preserve_unmatched_path}.
	PreserveUnmatchedPath interface{} `field:"optional" json:"preserveUnmatchedPath" yaml:"preserveUnmatchedPath"`
}

type CdnEndpointDeliveryRuleUrlRewriteActionOutputReference interface {
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
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnEndpointDeliveryRuleUrlRewriteAction
	SetInternalValue(val *CdnEndpointDeliveryRuleUrlRewriteAction)
	PreserveUnmatchedPath() interface{}
	SetPreserveUnmatchedPath(val interface{})
	PreserveUnmatchedPathInput() interface{}
	SourcePattern() *string
	SetSourcePattern(val *string)
	SourcePatternInput() *string
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
	ResetPreserveUnmatchedPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleUrlRewriteActionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) InternalValue() *CdnEndpointDeliveryRuleUrlRewriteAction {
	var returns *CdnEndpointDeliveryRuleUrlRewriteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) PreserveUnmatchedPath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveUnmatchedPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) PreserveUnmatchedPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveUnmatchedPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SourcePattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SourcePatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleUrlRewriteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointDeliveryRuleUrlRewriteActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlRewriteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleUrlRewriteActionOutputReference_Override(c CdnEndpointDeliveryRuleUrlRewriteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleUrlRewriteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SetDestination(val *string) {
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SetInternalValue(val *CdnEndpointDeliveryRuleUrlRewriteAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SetPreserveUnmatchedPath(val interface{}) {
	_jsii_.Set(
		j,
		"preserveUnmatchedPath",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SetSourcePattern(val *string) {
	_jsii_.Set(
		j,
		"sourcePattern",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) ResetPreserveUnmatchedPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPreserveUnmatchedPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleUrlRewriteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGeoFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#action CdnEndpoint#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#country_codes CdnEndpoint#country_codes}.
	CountryCodes *[]*string `field:"required" json:"countryCodes" yaml:"countryCodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#relative_path CdnEndpoint#relative_path}.
	RelativePath *string `field:"required" json:"relativePath" yaml:"relativePath"`
}

type CdnEndpointGeoFilterList interface {
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
	Get(index *float64) CdnEndpointGeoFilterOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGeoFilterList
type jsiiProxy_CdnEndpointGeoFilterList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointGeoFilterList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointGeoFilterList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGeoFilterList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGeoFilterList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointGeoFilterList_Override(c CdnEndpointGeoFilterList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGeoFilterList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGeoFilterList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterList) Get(index *float64) CdnEndpointGeoFilterOutputReference {
	var returns CdnEndpointGeoFilterOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGeoFilterOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	CountryCodes() *[]*string
	SetCountryCodes(val *[]*string)
	CountryCodesInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RelativePath() *string
	SetRelativePath(val *string)
	RelativePathInput() *string
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

// The jsii proxy struct for CdnEndpointGeoFilterOutputReference
type jsiiProxy_CdnEndpointGeoFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) CountryCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"countryCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) CountryCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"countryCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) RelativePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) RelativePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointGeoFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointGeoFilterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGeoFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGeoFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointGeoFilterOutputReference_Override(c CdnEndpointGeoFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGeoFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) SetCountryCodes(val *[]*string) {
	_jsii_.Set(
		j,
		"countryCodes",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) SetRelativePath(val *string) {
	_jsii_.Set(
		j,
		"relativePath",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGeoFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGeoFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRule struct {
	// cache_expiration_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#cache_expiration_action CdnEndpoint#cache_expiration_action}
	CacheExpirationAction *CdnEndpointGlobalDeliveryRuleCacheExpirationAction `field:"optional" json:"cacheExpirationAction" yaml:"cacheExpirationAction"`
	// cache_key_query_string_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#cache_key_query_string_action CdnEndpoint#cache_key_query_string_action}
	CacheKeyQueryStringAction *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction `field:"optional" json:"cacheKeyQueryStringAction" yaml:"cacheKeyQueryStringAction"`
	// modify_request_header_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#modify_request_header_action CdnEndpoint#modify_request_header_action}
	ModifyRequestHeaderAction interface{} `field:"optional" json:"modifyRequestHeaderAction" yaml:"modifyRequestHeaderAction"`
	// modify_response_header_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#modify_response_header_action CdnEndpoint#modify_response_header_action}
	ModifyResponseHeaderAction interface{} `field:"optional" json:"modifyResponseHeaderAction" yaml:"modifyResponseHeaderAction"`
	// url_redirect_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#url_redirect_action CdnEndpoint#url_redirect_action}
	UrlRedirectAction *CdnEndpointGlobalDeliveryRuleUrlRedirectAction `field:"optional" json:"urlRedirectAction" yaml:"urlRedirectAction"`
	// url_rewrite_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#url_rewrite_action CdnEndpoint#url_rewrite_action}
	UrlRewriteAction *CdnEndpointGlobalDeliveryRuleUrlRewriteAction `field:"optional" json:"urlRewriteAction" yaml:"urlRewriteAction"`
}

type CdnEndpointGlobalDeliveryRuleCacheExpirationAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#duration CdnEndpoint#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

type CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference interface {
	cdktf.ComplexObject
	Behavior() *string
	SetBehavior(val *string)
	BehaviorInput() *string
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnEndpointGlobalDeliveryRuleCacheExpirationAction
	SetInternalValue(val *CdnEndpointGlobalDeliveryRuleCacheExpirationAction)
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
	ResetDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference
type jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) InternalValue() *CdnEndpointGlobalDeliveryRuleCacheExpirationAction {
	var returns *CdnEndpointGlobalDeliveryRuleCacheExpirationAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference_Override(c CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) SetBehavior(val *string) {
	_jsii_.Set(
		j,
		"behavior",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) SetDuration(val *string) {
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) SetInternalValue(val *CdnEndpointGlobalDeliveryRuleCacheExpirationAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#parameters CdnEndpoint#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

type CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference interface {
	cdktf.ComplexObject
	Behavior() *string
	SetBehavior(val *string)
	BehaviorInput() *string
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
	InternalValue() *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction
	SetInternalValue(val *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction)
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
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
	ResetParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference
type jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) InternalValue() *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction {
	var returns *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference_Override(c CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) SetBehavior(val *string) {
	_jsii_.Set(
		j,
		"behavior",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) SetInternalValue(val *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) SetParameters(val *string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRuleModifyRequestHeaderAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#action CdnEndpoint#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#value CdnEndpoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

type CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList interface {
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
	Get(index *float64) CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList
type jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList_Override(c CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) Get(index *float64) CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference
type jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference_Override(c CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		c,
		"resetValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRuleModifyResponseHeaderAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#action CdnEndpoint#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#value CdnEndpoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

type CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList interface {
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
	Get(index *float64) CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList
type jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList_Override(c CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) Get(index *float64) CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference
type jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference_Override(c CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		c,
		"resetValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRuleOutputReference interface {
	cdktf.ComplexObject
	CacheExpirationAction() CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference
	CacheExpirationActionInput() *CdnEndpointGlobalDeliveryRuleCacheExpirationAction
	CacheKeyQueryStringAction() CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference
	CacheKeyQueryStringActionInput() *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction
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
	InternalValue() *CdnEndpointGlobalDeliveryRule
	SetInternalValue(val *CdnEndpointGlobalDeliveryRule)
	ModifyRequestHeaderAction() CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList
	ModifyRequestHeaderActionInput() interface{}
	ModifyResponseHeaderAction() CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList
	ModifyResponseHeaderActionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRedirectAction() CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference
	UrlRedirectActionInput() *CdnEndpointGlobalDeliveryRuleUrlRedirectAction
	UrlRewriteAction() CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference
	UrlRewriteActionInput() *CdnEndpointGlobalDeliveryRuleUrlRewriteAction
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
	PutCacheExpirationAction(value *CdnEndpointGlobalDeliveryRuleCacheExpirationAction)
	PutCacheKeyQueryStringAction(value *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction)
	PutModifyRequestHeaderAction(value interface{})
	PutModifyResponseHeaderAction(value interface{})
	PutUrlRedirectAction(value *CdnEndpointGlobalDeliveryRuleUrlRedirectAction)
	PutUrlRewriteAction(value *CdnEndpointGlobalDeliveryRuleUrlRewriteAction)
	ResetCacheExpirationAction()
	ResetCacheKeyQueryStringAction()
	ResetModifyRequestHeaderAction()
	ResetModifyResponseHeaderAction()
	ResetUrlRedirectAction()
	ResetUrlRewriteAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleOutputReference
type jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CacheExpirationAction() CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference
	_jsii_.Get(
		j,
		"cacheExpirationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CacheExpirationActionInput() *CdnEndpointGlobalDeliveryRuleCacheExpirationAction {
	var returns *CdnEndpointGlobalDeliveryRuleCacheExpirationAction
	_jsii_.Get(
		j,
		"cacheExpirationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CacheKeyQueryStringAction() CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference
	_jsii_.Get(
		j,
		"cacheKeyQueryStringAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CacheKeyQueryStringActionInput() *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction {
	var returns *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction
	_jsii_.Get(
		j,
		"cacheKeyQueryStringActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) InternalValue() *CdnEndpointGlobalDeliveryRule {
	var returns *CdnEndpointGlobalDeliveryRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ModifyRequestHeaderAction() CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList {
	var returns CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList
	_jsii_.Get(
		j,
		"modifyRequestHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ModifyRequestHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyRequestHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ModifyResponseHeaderAction() CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList {
	var returns CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList
	_jsii_.Get(
		j,
		"modifyResponseHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ModifyResponseHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyResponseHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) UrlRedirectAction() CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference
	_jsii_.Get(
		j,
		"urlRedirectAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) UrlRedirectActionInput() *CdnEndpointGlobalDeliveryRuleUrlRedirectAction {
	var returns *CdnEndpointGlobalDeliveryRuleUrlRedirectAction
	_jsii_.Get(
		j,
		"urlRedirectActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) UrlRewriteAction() CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference
	_jsii_.Get(
		j,
		"urlRewriteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) UrlRewriteActionInput() *CdnEndpointGlobalDeliveryRuleUrlRewriteAction {
	var returns *CdnEndpointGlobalDeliveryRuleUrlRewriteAction
	_jsii_.Get(
		j,
		"urlRewriteActionInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointGlobalDeliveryRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleOutputReference_Override(c CdnEndpointGlobalDeliveryRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) SetInternalValue(val *CdnEndpointGlobalDeliveryRule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutCacheExpirationAction(value *CdnEndpointGlobalDeliveryRuleCacheExpirationAction) {
	_jsii_.InvokeVoid(
		c,
		"putCacheExpirationAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutCacheKeyQueryStringAction(value *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction) {
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyQueryStringAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutModifyRequestHeaderAction(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putModifyRequestHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutModifyResponseHeaderAction(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putModifyResponseHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutUrlRedirectAction(value *CdnEndpointGlobalDeliveryRuleUrlRedirectAction) {
	_jsii_.InvokeVoid(
		c,
		"putUrlRedirectAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutUrlRewriteAction(value *CdnEndpointGlobalDeliveryRuleUrlRewriteAction) {
	_jsii_.InvokeVoid(
		c,
		"putUrlRewriteAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetCacheExpirationAction() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheExpirationAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetCacheKeyQueryStringAction() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyQueryStringAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetModifyRequestHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyRequestHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetModifyResponseHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyResponseHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetUrlRedirectAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRedirectAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetUrlRewriteAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewriteAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRuleUrlRedirectAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#redirect_type CdnEndpoint#redirect_type}.
	RedirectType *string `field:"required" json:"redirectType" yaml:"redirectType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#fragment CdnEndpoint#fragment}.
	Fragment *string `field:"optional" json:"fragment" yaml:"fragment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#hostname CdnEndpoint#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#path CdnEndpoint#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#protocol CdnEndpoint#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#query_string CdnEndpoint#query_string}.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
}

type CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference interface {
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
	Fragment() *string
	SetFragment(val *string)
	FragmentInput() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InternalValue() *CdnEndpointGlobalDeliveryRuleUrlRedirectAction
	SetInternalValue(val *CdnEndpointGlobalDeliveryRuleUrlRedirectAction)
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
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
	ResetFragment()
	ResetHostname()
	ResetPath()
	ResetProtocol()
	ResetQueryString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference
type jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) Fragment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) FragmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) InternalValue() *CdnEndpointGlobalDeliveryRuleUrlRedirectAction {
	var returns *CdnEndpointGlobalDeliveryRuleUrlRedirectAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) RedirectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) RedirectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference_Override(c CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetFragment(val *string) {
	_jsii_.Set(
		j,
		"fragment",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetInternalValue(val *CdnEndpointGlobalDeliveryRuleUrlRedirectAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetQueryString(val *string) {
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetRedirectType(val *string) {
	_jsii_.Set(
		j,
		"redirectType",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ResetFragment() {
	_jsii_.InvokeVoid(
		c,
		"resetFragment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointGlobalDeliveryRuleUrlRewriteAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#destination CdnEndpoint#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#source_pattern CdnEndpoint#source_pattern}.
	SourcePattern *string `field:"required" json:"sourcePattern" yaml:"sourcePattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#preserve_unmatched_path CdnEndpoint#preserve_unmatched_path}.
	PreserveUnmatchedPath interface{} `field:"optional" json:"preserveUnmatchedPath" yaml:"preserveUnmatchedPath"`
}

type CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference interface {
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
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnEndpointGlobalDeliveryRuleUrlRewriteAction
	SetInternalValue(val *CdnEndpointGlobalDeliveryRuleUrlRewriteAction)
	PreserveUnmatchedPath() interface{}
	SetPreserveUnmatchedPath(val interface{})
	PreserveUnmatchedPathInput() interface{}
	SourcePattern() *string
	SetSourcePattern(val *string)
	SourcePatternInput() *string
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
	ResetPreserveUnmatchedPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference
type jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) InternalValue() *CdnEndpointGlobalDeliveryRuleUrlRewriteAction {
	var returns *CdnEndpointGlobalDeliveryRuleUrlRewriteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) PreserveUnmatchedPath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveUnmatchedPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) PreserveUnmatchedPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveUnmatchedPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SourcePattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SourcePatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference_Override(c CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SetDestination(val *string) {
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SetInternalValue(val *CdnEndpointGlobalDeliveryRuleUrlRewriteAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SetPreserveUnmatchedPath(val interface{}) {
	_jsii_.Set(
		j,
		"preserveUnmatchedPath",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SetSourcePattern(val *string) {
	_jsii_.Set(
		j,
		"sourcePattern",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) ResetPreserveUnmatchedPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPreserveUnmatchedPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointOrigin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#host_name CdnEndpoint#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#http_port CdnEndpoint#http_port}.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#https_port CdnEndpoint#https_port}.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
}

type CdnEndpointOriginList interface {
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
	Get(index *float64) CdnEndpointOriginOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointOriginList
type jsiiProxy_CdnEndpointOriginList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CdnEndpointOriginList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnEndpointOriginList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnEndpointOriginList {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointOriginList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointOriginList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnEndpointOriginList_Override(c CdnEndpointOriginList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointOriginList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointOriginList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointOriginList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginList) Get(index *float64) CdnEndpointOriginOutputReference {
	var returns CdnEndpointOriginOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointOriginOutputReference interface {
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
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	HttpsPort() *float64
	SetHttpsPort(val *float64)
	HttpsPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetHttpPort()
	ResetHttpsPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointOriginOutputReference
type jsiiProxy_CdnEndpointOriginOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) HttpsPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) HttpsPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointOriginOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointOriginOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointOriginOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointOriginOutputReference_Override(c CdnEndpointOriginOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetHostName(val *string) {
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetHttpPort(val *float64) {
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetHttpsPort(val *float64) {
	_jsii_.Set(
		j,
		"httpsPort",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointOriginOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) ResetHttpPort() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) ResetHttpsPort() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointOriginOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CdnEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#create CdnEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#delete CdnEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#read CdnEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#update CdnEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type CdnEndpointTimeoutsOutputReference interface {
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

// The jsii proxy struct for CdnEndpointTimeoutsOutputReference
type jsiiProxy_CdnEndpointTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CdnEndpointTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointTimeoutsOutputReference_Override(c CdnEndpointTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		c,
		"resetRead",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

