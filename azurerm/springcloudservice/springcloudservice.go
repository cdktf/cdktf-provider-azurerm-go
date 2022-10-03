package springcloudservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/springcloudservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service azurerm_spring_cloud_service}.
type SpringCloudService interface {
	cdktf.TerraformResource
	BuildAgentPoolSize() *string
	SetBuildAgentPoolSize(val *string)
	BuildAgentPoolSizeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigServerGitSetting() SpringCloudServiceConfigServerGitSettingOutputReference
	ConfigServerGitSettingInput() *SpringCloudServiceConfigServerGitSetting
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogStreamPublicEndpointEnabled() interface{}
	SetLogStreamPublicEndpointEnabled(val interface{})
	LogStreamPublicEndpointEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() SpringCloudServiceNetworkOutputReference
	NetworkInput() *SpringCloudServiceNetwork
	// The tree node.
	Node() constructs.Node
	OutboundPublicIpAddresses() *[]*string
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
	RequiredNetworkTrafficRules() SpringCloudServiceRequiredNetworkTrafficRulesList
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ServiceRegistryEnabled() interface{}
	SetServiceRegistryEnabled(val interface{})
	ServiceRegistryEnabledInput() interface{}
	ServiceRegistryId() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SpringCloudServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Trace() SpringCloudServiceTraceOutputReference
	TraceInput() *SpringCloudServiceTrace
	ZoneRedundant() interface{}
	SetZoneRedundant(val interface{})
	ZoneRedundantInput() interface{}
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
	PutConfigServerGitSetting(value *SpringCloudServiceConfigServerGitSetting)
	PutNetwork(value *SpringCloudServiceNetwork)
	PutTimeouts(value *SpringCloudServiceTimeouts)
	PutTrace(value *SpringCloudServiceTrace)
	ResetBuildAgentPoolSize()
	ResetConfigServerGitSetting()
	ResetId()
	ResetLogStreamPublicEndpointEnabled()
	ResetNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceRegistryEnabled()
	ResetSkuName()
	ResetTags()
	ResetTimeouts()
	ResetTrace()
	ResetZoneRedundant()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SpringCloudService
type jsiiProxy_SpringCloudService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpringCloudService) BuildAgentPoolSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildAgentPoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) BuildAgentPoolSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildAgentPoolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ConfigServerGitSetting() SpringCloudServiceConfigServerGitSettingOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingOutputReference
	_jsii_.Get(
		j,
		"configServerGitSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ConfigServerGitSettingInput() *SpringCloudServiceConfigServerGitSetting {
	var returns *SpringCloudServiceConfigServerGitSetting
	_jsii_.Get(
		j,
		"configServerGitSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) LogStreamPublicEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamPublicEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) LogStreamPublicEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStreamPublicEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Network() SpringCloudServiceNetworkOutputReference {
	var returns SpringCloudServiceNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) NetworkInput() *SpringCloudServiceNetwork {
	var returns *SpringCloudServiceNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) OutboundPublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundPublicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) RequiredNetworkTrafficRules() SpringCloudServiceRequiredNetworkTrafficRulesList {
	var returns SpringCloudServiceRequiredNetworkTrafficRulesList
	_jsii_.Get(
		j,
		"requiredNetworkTrafficRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ServiceRegistryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceRegistryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ServiceRegistryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceRegistryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ServiceRegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRegistryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Timeouts() SpringCloudServiceTimeoutsOutputReference {
	var returns SpringCloudServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Trace() SpringCloudServiceTraceOutputReference {
	var returns SpringCloudServiceTraceOutputReference
	_jsii_.Get(
		j,
		"trace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) TraceInput() *SpringCloudServiceTrace {
	var returns *SpringCloudServiceTrace
	_jsii_.Get(
		j,
		"traceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ZoneRedundant() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ZoneRedundantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundantInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service azurerm_spring_cloud_service} Resource.
func NewSpringCloudService(scope constructs.Construct, id *string, config *SpringCloudServiceConfig) SpringCloudService {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudService{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service azurerm_spring_cloud_service} Resource.
func NewSpringCloudService_Override(s SpringCloudService, scope constructs.Construct, id *string, config *SpringCloudServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpringCloudService) SetBuildAgentPoolSize(val *string) {
	_jsii_.Set(
		j,
		"buildAgentPoolSize",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetLogStreamPublicEndpointEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"logStreamPublicEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetServiceRegistryEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"serviceRegistryEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetSkuName(val *string) {
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService) SetZoneRedundant(val interface{}) {
	_jsii_.Set(
		j,
		"zoneRedundant",
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
func SpringCloudService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpringCloudService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpringCloudService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpringCloudService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpringCloudService) PutConfigServerGitSetting(value *SpringCloudServiceConfigServerGitSetting) {
	_jsii_.InvokeVoid(
		s,
		"putConfigServerGitSetting",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutNetwork(value *SpringCloudServiceNetwork) {
	_jsii_.InvokeVoid(
		s,
		"putNetwork",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutTimeouts(value *SpringCloudServiceTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutTrace(value *SpringCloudServiceTrace) {
	_jsii_.InvokeVoid(
		s,
		"putTrace",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) ResetBuildAgentPoolSize() {
	_jsii_.InvokeVoid(
		s,
		"resetBuildAgentPoolSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetConfigServerGitSetting() {
	_jsii_.InvokeVoid(
		s,
		"resetConfigServerGitSetting",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetLogStreamPublicEndpointEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetLogStreamPublicEndpointEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetNetwork() {
	_jsii_.InvokeVoid(
		s,
		"resetNetwork",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetServiceRegistryEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceRegistryEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetSkuName() {
	_jsii_.InvokeVoid(
		s,
		"resetSkuName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetTrace() {
	_jsii_.InvokeVoid(
		s,
		"resetTrace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetZoneRedundant() {
	_jsii_.InvokeVoid(
		s,
		"resetZoneRedundant",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#location SpringCloudService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#name SpringCloudService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#resource_group_name SpringCloudService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#build_agent_pool_size SpringCloudService#build_agent_pool_size}.
	BuildAgentPoolSize *string `field:"optional" json:"buildAgentPoolSize" yaml:"buildAgentPoolSize"`
	// config_server_git_setting block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#config_server_git_setting SpringCloudService#config_server_git_setting}
	ConfigServerGitSetting *SpringCloudServiceConfigServerGitSetting `field:"optional" json:"configServerGitSetting" yaml:"configServerGitSetting"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#id SpringCloudService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#log_stream_public_endpoint_enabled SpringCloudService#log_stream_public_endpoint_enabled}.
	LogStreamPublicEndpointEnabled interface{} `field:"optional" json:"logStreamPublicEndpointEnabled" yaml:"logStreamPublicEndpointEnabled"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#network SpringCloudService#network}
	Network *SpringCloudServiceNetwork `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#service_registry_enabled SpringCloudService#service_registry_enabled}.
	ServiceRegistryEnabled interface{} `field:"optional" json:"serviceRegistryEnabled" yaml:"serviceRegistryEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#sku_name SpringCloudService#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#tags SpringCloudService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#timeouts SpringCloudService#timeouts}
	Timeouts *SpringCloudServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// trace block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#trace SpringCloudService#trace}
	Trace *SpringCloudServiceTrace `field:"optional" json:"trace" yaml:"trace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#zone_redundant SpringCloudService#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

type SpringCloudServiceConfigServerGitSetting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#uri SpringCloudService#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// http_basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#http_basic_auth SpringCloudService#http_basic_auth}
	HttpBasicAuth *SpringCloudServiceConfigServerGitSettingHttpBasicAuth `field:"optional" json:"httpBasicAuth" yaml:"httpBasicAuth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#label SpringCloudService#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// repository block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#repository SpringCloudService#repository}
	Repository interface{} `field:"optional" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#search_paths SpringCloudService#search_paths}.
	SearchPaths *[]*string `field:"optional" json:"searchPaths" yaml:"searchPaths"`
	// ssh_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#ssh_auth SpringCloudService#ssh_auth}
	SshAuth *SpringCloudServiceConfigServerGitSettingSshAuth `field:"optional" json:"sshAuth" yaml:"sshAuth"`
}

type SpringCloudServiceConfigServerGitSettingHttpBasicAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#password SpringCloudService#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#username SpringCloudService#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference interface {
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
	InternalValue() *SpringCloudServiceConfigServerGitSettingHttpBasicAuth
	SetInternalValue(val *SpringCloudServiceConfigServerGitSettingHttpBasicAuth)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) InternalValue() *SpringCloudServiceConfigServerGitSettingHttpBasicAuth {
	var returns *SpringCloudServiceConfigServerGitSettingHttpBasicAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference_Override(s SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) SetInternalValue(val *SpringCloudServiceConfigServerGitSettingHttpBasicAuth) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceConfigServerGitSettingOutputReference interface {
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
	HttpBasicAuth() SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference
	HttpBasicAuthInput() *SpringCloudServiceConfigServerGitSettingHttpBasicAuth
	InternalValue() *SpringCloudServiceConfigServerGitSetting
	SetInternalValue(val *SpringCloudServiceConfigServerGitSetting)
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Repository() SpringCloudServiceConfigServerGitSettingRepositoryList
	RepositoryInput() interface{}
	SearchPaths() *[]*string
	SetSearchPaths(val *[]*string)
	SearchPathsInput() *[]*string
	SshAuth() SpringCloudServiceConfigServerGitSettingSshAuthOutputReference
	SshAuthInput() *SpringCloudServiceConfigServerGitSettingSshAuth
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	PutHttpBasicAuth(value *SpringCloudServiceConfigServerGitSettingHttpBasicAuth)
	PutRepository(value interface{})
	PutSshAuth(value *SpringCloudServiceConfigServerGitSettingSshAuth)
	ResetHttpBasicAuth()
	ResetLabel()
	ResetRepository()
	ResetSearchPaths()
	ResetSshAuth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) HttpBasicAuth() SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference
	_jsii_.Get(
		j,
		"httpBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) HttpBasicAuthInput() *SpringCloudServiceConfigServerGitSettingHttpBasicAuth {
	var returns *SpringCloudServiceConfigServerGitSettingHttpBasicAuth
	_jsii_.Get(
		j,
		"httpBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) InternalValue() *SpringCloudServiceConfigServerGitSetting {
	var returns *SpringCloudServiceConfigServerGitSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Repository() SpringCloudServiceConfigServerGitSettingRepositoryList {
	var returns SpringCloudServiceConfigServerGitSettingRepositoryList
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) RepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SearchPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SearchPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SshAuth() SpringCloudServiceConfigServerGitSettingSshAuthOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingSshAuthOutputReference
	_jsii_.Get(
		j,
		"sshAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SshAuthInput() *SpringCloudServiceConfigServerGitSettingSshAuth {
	var returns *SpringCloudServiceConfigServerGitSettingSshAuth
	_jsii_.Get(
		j,
		"sshAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceConfigServerGitSettingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingOutputReference_Override(s SpringCloudServiceConfigServerGitSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SetInternalValue(val *SpringCloudServiceConfigServerGitSetting) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SetLabel(val *string) {
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SetSearchPaths(val *[]*string) {
	_jsii_.Set(
		j,
		"searchPaths",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) PutHttpBasicAuth(value *SpringCloudServiceConfigServerGitSettingHttpBasicAuth) {
	_jsii_.InvokeVoid(
		s,
		"putHttpBasicAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) PutRepository(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putRepository",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) PutSshAuth(value *SpringCloudServiceConfigServerGitSettingSshAuth) {
	_jsii_.InvokeVoid(
		s,
		"putSshAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetHttpBasicAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpBasicAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetRepository() {
	_jsii_.InvokeVoid(
		s,
		"resetRepository",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetSearchPaths() {
	_jsii_.InvokeVoid(
		s,
		"resetSearchPaths",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetSshAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetSshAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceConfigServerGitSettingRepository struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#name SpringCloudService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#uri SpringCloudService#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// http_basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#http_basic_auth SpringCloudService#http_basic_auth}
	HttpBasicAuth *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth `field:"optional" json:"httpBasicAuth" yaml:"httpBasicAuth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#label SpringCloudService#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#pattern SpringCloudService#pattern}.
	Pattern *[]*string `field:"optional" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#search_paths SpringCloudService#search_paths}.
	SearchPaths *[]*string `field:"optional" json:"searchPaths" yaml:"searchPaths"`
	// ssh_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#ssh_auth SpringCloudService#ssh_auth}
	SshAuth *SpringCloudServiceConfigServerGitSettingRepositorySshAuth `field:"optional" json:"sshAuth" yaml:"sshAuth"`
}

type SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#password SpringCloudService#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#username SpringCloudService#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference interface {
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
	InternalValue() *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth
	SetInternalValue(val *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) InternalValue() *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth {
	var returns *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference_Override(s SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) SetInternalValue(val *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceConfigServerGitSettingRepositoryList interface {
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
	Get(index *float64) SpringCloudServiceConfigServerGitSettingRepositoryOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingRepositoryList
type jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingRepositoryList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SpringCloudServiceConfigServerGitSettingRepositoryList {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositoryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingRepositoryList_Override(s SpringCloudServiceConfigServerGitSettingRepositoryList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositoryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) Get(index *float64) SpringCloudServiceConfigServerGitSettingRepositoryOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingRepositoryOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceConfigServerGitSettingRepositoryOutputReference interface {
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
	HttpBasicAuth() SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference
	HttpBasicAuthInput() *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Pattern() *[]*string
	SetPattern(val *[]*string)
	PatternInput() *[]*string
	SearchPaths() *[]*string
	SetSearchPaths(val *[]*string)
	SearchPathsInput() *[]*string
	SshAuth() SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference
	SshAuthInput() *SpringCloudServiceConfigServerGitSettingRepositorySshAuth
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	PutHttpBasicAuth(value *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth)
	PutSshAuth(value *SpringCloudServiceConfigServerGitSettingRepositorySshAuth)
	ResetHttpBasicAuth()
	ResetLabel()
	ResetPattern()
	ResetSearchPaths()
	ResetSshAuth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingRepositoryOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) HttpBasicAuth() SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference
	_jsii_.Get(
		j,
		"httpBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) HttpBasicAuthInput() *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth {
	var returns *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth
	_jsii_.Get(
		j,
		"httpBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Pattern() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) PatternInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SearchPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SearchPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SshAuth() SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference
	_jsii_.Get(
		j,
		"sshAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SshAuthInput() *SpringCloudServiceConfigServerGitSettingRepositorySshAuth {
	var returns *SpringCloudServiceConfigServerGitSettingRepositorySshAuth
	_jsii_.Get(
		j,
		"sshAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SpringCloudServiceConfigServerGitSettingRepositoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingRepositoryOutputReference_Override(s SpringCloudServiceConfigServerGitSettingRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetLabel(val *string) {
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetPattern(val *[]*string) {
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetSearchPaths(val *[]*string) {
	_jsii_.Set(
		j,
		"searchPaths",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) PutHttpBasicAuth(value *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth) {
	_jsii_.InvokeVoid(
		s,
		"putHttpBasicAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) PutSshAuth(value *SpringCloudServiceConfigServerGitSettingRepositorySshAuth) {
	_jsii_.InvokeVoid(
		s,
		"putSshAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetHttpBasicAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpBasicAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetPattern() {
	_jsii_.InvokeVoid(
		s,
		"resetPattern",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetSearchPaths() {
	_jsii_.InvokeVoid(
		s,
		"resetSearchPaths",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetSshAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetSshAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceConfigServerGitSettingRepositorySshAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#private_key SpringCloudService#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#host_key SpringCloudService#host_key}.
	HostKey *string `field:"optional" json:"hostKey" yaml:"hostKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#host_key_algorithm SpringCloudService#host_key_algorithm}.
	HostKeyAlgorithm *string `field:"optional" json:"hostKeyAlgorithm" yaml:"hostKeyAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#strict_host_key_checking_enabled SpringCloudService#strict_host_key_checking_enabled}.
	StrictHostKeyCheckingEnabled interface{} `field:"optional" json:"strictHostKeyCheckingEnabled" yaml:"strictHostKeyCheckingEnabled"`
}

type SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference interface {
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
	HostKey() *string
	SetHostKey(val *string)
	HostKeyAlgorithm() *string
	SetHostKeyAlgorithm(val *string)
	HostKeyAlgorithmInput() *string
	HostKeyInput() *string
	InternalValue() *SpringCloudServiceConfigServerGitSettingRepositorySshAuth
	SetInternalValue(val *SpringCloudServiceConfigServerGitSettingRepositorySshAuth)
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	StrictHostKeyCheckingEnabled() interface{}
	SetStrictHostKeyCheckingEnabled(val interface{})
	StrictHostKeyCheckingEnabledInput() interface{}
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
	ResetHostKey()
	ResetHostKeyAlgorithm()
	ResetStrictHostKeyCheckingEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) HostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) HostKeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) HostKeyAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) HostKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) InternalValue() *SpringCloudServiceConfigServerGitSettingRepositorySshAuth {
	var returns *SpringCloudServiceConfigServerGitSettingRepositorySshAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) StrictHostKeyCheckingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictHostKeyCheckingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) StrictHostKeyCheckingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictHostKeyCheckingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference_Override(s SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetHostKey(val *string) {
	_jsii_.Set(
		j,
		"hostKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetHostKeyAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"hostKeyAlgorithm",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetInternalValue(val *SpringCloudServiceConfigServerGitSettingRepositorySshAuth) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetStrictHostKeyCheckingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"strictHostKeyCheckingEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ResetHostKey() {
	_jsii_.InvokeVoid(
		s,
		"resetHostKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ResetHostKeyAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetHostKeyAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ResetStrictHostKeyCheckingEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetStrictHostKeyCheckingEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceConfigServerGitSettingSshAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#private_key SpringCloudService#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#host_key SpringCloudService#host_key}.
	HostKey *string `field:"optional" json:"hostKey" yaml:"hostKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#host_key_algorithm SpringCloudService#host_key_algorithm}.
	HostKeyAlgorithm *string `field:"optional" json:"hostKeyAlgorithm" yaml:"hostKeyAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#strict_host_key_checking_enabled SpringCloudService#strict_host_key_checking_enabled}.
	StrictHostKeyCheckingEnabled interface{} `field:"optional" json:"strictHostKeyCheckingEnabled" yaml:"strictHostKeyCheckingEnabled"`
}

type SpringCloudServiceConfigServerGitSettingSshAuthOutputReference interface {
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
	HostKey() *string
	SetHostKey(val *string)
	HostKeyAlgorithm() *string
	SetHostKeyAlgorithm(val *string)
	HostKeyAlgorithmInput() *string
	HostKeyInput() *string
	InternalValue() *SpringCloudServiceConfigServerGitSettingSshAuth
	SetInternalValue(val *SpringCloudServiceConfigServerGitSettingSshAuth)
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	StrictHostKeyCheckingEnabled() interface{}
	SetStrictHostKeyCheckingEnabled(val interface{})
	StrictHostKeyCheckingEnabledInput() interface{}
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
	ResetHostKey()
	ResetHostKeyAlgorithm()
	ResetStrictHostKeyCheckingEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingSshAuthOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) HostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) HostKeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) HostKeyAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) HostKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) InternalValue() *SpringCloudServiceConfigServerGitSettingSshAuth {
	var returns *SpringCloudServiceConfigServerGitSettingSshAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) StrictHostKeyCheckingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictHostKeyCheckingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) StrictHostKeyCheckingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictHostKeyCheckingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingSshAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceConfigServerGitSettingSshAuthOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingSshAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingSshAuthOutputReference_Override(s SpringCloudServiceConfigServerGitSettingSshAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingSshAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetHostKey(val *string) {
	_jsii_.Set(
		j,
		"hostKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetHostKeyAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"hostKeyAlgorithm",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetInternalValue(val *SpringCloudServiceConfigServerGitSettingSshAuth) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetStrictHostKeyCheckingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"strictHostKeyCheckingEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) ResetHostKey() {
	_jsii_.InvokeVoid(
		s,
		"resetHostKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) ResetHostKeyAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetHostKeyAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) ResetStrictHostKeyCheckingEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetStrictHostKeyCheckingEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingSshAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceNetwork struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#app_subnet_id SpringCloudService#app_subnet_id}.
	AppSubnetId *string `field:"required" json:"appSubnetId" yaml:"appSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#cidr_ranges SpringCloudService#cidr_ranges}.
	CidrRanges *[]*string `field:"required" json:"cidrRanges" yaml:"cidrRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#service_runtime_subnet_id SpringCloudService#service_runtime_subnet_id}.
	ServiceRuntimeSubnetId *string `field:"required" json:"serviceRuntimeSubnetId" yaml:"serviceRuntimeSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#app_network_resource_group SpringCloudService#app_network_resource_group}.
	AppNetworkResourceGroup *string `field:"optional" json:"appNetworkResourceGroup" yaml:"appNetworkResourceGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#read_timeout_seconds SpringCloudService#read_timeout_seconds}.
	ReadTimeoutSeconds *float64 `field:"optional" json:"readTimeoutSeconds" yaml:"readTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#service_runtime_network_resource_group SpringCloudService#service_runtime_network_resource_group}.
	ServiceRuntimeNetworkResourceGroup *string `field:"optional" json:"serviceRuntimeNetworkResourceGroup" yaml:"serviceRuntimeNetworkResourceGroup"`
}

type SpringCloudServiceNetworkOutputReference interface {
	cdktf.ComplexObject
	AppNetworkResourceGroup() *string
	SetAppNetworkResourceGroup(val *string)
	AppNetworkResourceGroupInput() *string
	AppSubnetId() *string
	SetAppSubnetId(val *string)
	AppSubnetIdInput() *string
	CidrRanges() *[]*string
	SetCidrRanges(val *[]*string)
	CidrRangesInput() *[]*string
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
	InternalValue() *SpringCloudServiceNetwork
	SetInternalValue(val *SpringCloudServiceNetwork)
	ReadTimeoutSeconds() *float64
	SetReadTimeoutSeconds(val *float64)
	ReadTimeoutSecondsInput() *float64
	ServiceRuntimeNetworkResourceGroup() *string
	SetServiceRuntimeNetworkResourceGroup(val *string)
	ServiceRuntimeNetworkResourceGroupInput() *string
	ServiceRuntimeSubnetId() *string
	SetServiceRuntimeSubnetId(val *string)
	ServiceRuntimeSubnetIdInput() *string
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
	ResetAppNetworkResourceGroup()
	ResetReadTimeoutSeconds()
	ResetServiceRuntimeNetworkResourceGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceNetworkOutputReference
type jsiiProxy_SpringCloudServiceNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) AppNetworkResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNetworkResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) AppNetworkResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNetworkResourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) AppSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) AppSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) CidrRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) CidrRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) InternalValue() *SpringCloudServiceNetwork {
	var returns *SpringCloudServiceNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) ReadTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) ReadTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) ServiceRuntimeNetworkResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRuntimeNetworkResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) ServiceRuntimeNetworkResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRuntimeNetworkResourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) ServiceRuntimeSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRuntimeSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) ServiceRuntimeSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRuntimeSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceNetworkOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceNetworkOutputReference_Override(s SpringCloudServiceNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetAppNetworkResourceGroup(val *string) {
	_jsii_.Set(
		j,
		"appNetworkResourceGroup",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetAppSubnetId(val *string) {
	_jsii_.Set(
		j,
		"appSubnetId",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetCidrRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"cidrRanges",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetInternalValue(val *SpringCloudServiceNetwork) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetReadTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"readTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetServiceRuntimeNetworkResourceGroup(val *string) {
	_jsii_.Set(
		j,
		"serviceRuntimeNetworkResourceGroup",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetServiceRuntimeSubnetId(val *string) {
	_jsii_.Set(
		j,
		"serviceRuntimeSubnetId",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceNetworkOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) ResetAppNetworkResourceGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetAppNetworkResourceGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) ResetReadTimeoutSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetReadTimeoutSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) ResetServiceRuntimeNetworkResourceGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceRuntimeNetworkResourceGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceRequiredNetworkTrafficRules struct {
}

type SpringCloudServiceRequiredNetworkTrafficRulesList interface {
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
	Get(index *float64) SpringCloudServiceRequiredNetworkTrafficRulesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceRequiredNetworkTrafficRulesList
type jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceRequiredNetworkTrafficRulesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SpringCloudServiceRequiredNetworkTrafficRulesList {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceRequiredNetworkTrafficRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSpringCloudServiceRequiredNetworkTrafficRulesList_Override(s SpringCloudServiceRequiredNetworkTrafficRulesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceRequiredNetworkTrafficRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) Get(index *float64) SpringCloudServiceRequiredNetworkTrafficRulesOutputReference {
	var returns SpringCloudServiceRequiredNetworkTrafficRulesOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceRequiredNetworkTrafficRulesOutputReference interface {
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
	Direction() *string
	Fqdns() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SpringCloudServiceRequiredNetworkTrafficRules
	SetInternalValue(val *SpringCloudServiceRequiredNetworkTrafficRules)
	IpAddresses() *[]*string
	Port() *float64
	Protocol() *string
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

// The jsii proxy struct for SpringCloudServiceRequiredNetworkTrafficRulesOutputReference
type jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) Fqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) InternalValue() *SpringCloudServiceRequiredNetworkTrafficRules {
	var returns *SpringCloudServiceRequiredNetworkTrafficRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceRequiredNetworkTrafficRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SpringCloudServiceRequiredNetworkTrafficRulesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceRequiredNetworkTrafficRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSpringCloudServiceRequiredNetworkTrafficRulesOutputReference_Override(s SpringCloudServiceRequiredNetworkTrafficRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceRequiredNetworkTrafficRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) SetInternalValue(val *SpringCloudServiceRequiredNetworkTrafficRules) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#create SpringCloudService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#delete SpringCloudService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#read SpringCloudService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#update SpringCloudService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type SpringCloudServiceTimeoutsOutputReference interface {
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

// The jsii proxy struct for SpringCloudServiceTimeoutsOutputReference
type jsiiProxy_SpringCloudServiceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceTimeoutsOutputReference_Override(s SpringCloudServiceTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SpringCloudServiceTrace struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#connection_string SpringCloudService#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#sample_rate SpringCloudService#sample_rate}.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

type SpringCloudServiceTraceOutputReference interface {
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
	ConnectionString() *string
	SetConnectionString(val *string)
	ConnectionStringInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SpringCloudServiceTrace
	SetInternalValue(val *SpringCloudServiceTrace)
	SampleRate() *float64
	SetSampleRate(val *float64)
	SampleRateInput() *float64
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
	ResetConnectionString()
	ResetSampleRate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceTraceOutputReference
type jsiiProxy_SpringCloudServiceTraceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) ConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) InternalValue() *SpringCloudServiceTrace {
	var returns *SpringCloudServiceTrace
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceTraceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceTraceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SpringCloudServiceTraceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceTraceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceTraceOutputReference_Override(s SpringCloudServiceTraceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceTraceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SetConnectionString(val *string) {
	_jsii_.Set(
		j,
		"connectionString",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SetInternalValue(val *SpringCloudServiceTrace) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SetSampleRate(val *float64) {
	_jsii_.Set(
		j,
		"sampleRate",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceTraceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) ResetConnectionString() {
	_jsii_.InvokeVoid(
		s,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) ResetSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceTraceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

