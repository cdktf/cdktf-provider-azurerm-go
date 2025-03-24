// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/springcloudservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/spring_cloud_service azurerm_spring_cloud_service}.
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
	ContainerRegistry() SpringCloudServiceContainerRegistryList
	ContainerRegistryInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultBuildService() SpringCloudServiceDefaultBuildServiceOutputReference
	DefaultBuildServiceInput() *SpringCloudServiceDefaultBuildService
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
	ManagedEnvironmentId() *string
	SetManagedEnvironmentId(val *string)
	ManagedEnvironmentIdInput() *string
	Marketplace() SpringCloudServiceMarketplaceOutputReference
	MarketplaceInput() *SpringCloudServiceMarketplace
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
	SkuTier() *string
	SetSkuTier(val *string)
	SkuTierInput() *string
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
	PutConfigServerGitSetting(value *SpringCloudServiceConfigServerGitSetting)
	PutContainerRegistry(value interface{})
	PutDefaultBuildService(value *SpringCloudServiceDefaultBuildService)
	PutMarketplace(value *SpringCloudServiceMarketplace)
	PutNetwork(value *SpringCloudServiceNetwork)
	PutTimeouts(value *SpringCloudServiceTimeouts)
	PutTrace(value *SpringCloudServiceTrace)
	ResetBuildAgentPoolSize()
	ResetConfigServerGitSetting()
	ResetContainerRegistry()
	ResetDefaultBuildService()
	ResetId()
	ResetLogStreamPublicEndpointEnabled()
	ResetManagedEnvironmentId()
	ResetMarketplace()
	ResetNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceRegistryEnabled()
	ResetSkuName()
	ResetSkuTier()
	ResetTags()
	ResetTimeouts()
	ResetTrace()
	ResetZoneRedundant()
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

func (j *jsiiProxy_SpringCloudService) ContainerRegistry() SpringCloudServiceContainerRegistryList {
	var returns SpringCloudServiceContainerRegistryList
	_jsii_.Get(
		j,
		"containerRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ContainerRegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) DefaultBuildService() SpringCloudServiceDefaultBuildServiceOutputReference {
	var returns SpringCloudServiceDefaultBuildServiceOutputReference
	_jsii_.Get(
		j,
		"defaultBuildService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) DefaultBuildServiceInput() *SpringCloudServiceDefaultBuildService {
	var returns *SpringCloudServiceDefaultBuildService
	_jsii_.Get(
		j,
		"defaultBuildServiceInput",
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

func (j *jsiiProxy_SpringCloudService) ManagedEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) ManagedEnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedEnvironmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) Marketplace() SpringCloudServiceMarketplaceOutputReference {
	var returns SpringCloudServiceMarketplaceOutputReference
	_jsii_.Get(
		j,
		"marketplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) MarketplaceInput() *SpringCloudServiceMarketplace {
	var returns *SpringCloudServiceMarketplace
	_jsii_.Get(
		j,
		"marketplaceInput",
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

func (j *jsiiProxy_SpringCloudService) SkuTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudService) SkuTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuTierInput",
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/spring_cloud_service azurerm_spring_cloud_service} Resource.
func NewSpringCloudService(scope constructs.Construct, id *string, config *SpringCloudServiceConfig) SpringCloudService {
	_init_.Initialize()

	if err := validateNewSpringCloudServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudService{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/spring_cloud_service azurerm_spring_cloud_service} Resource.
func NewSpringCloudService_Override(s SpringCloudService, scope constructs.Construct, id *string, config *SpringCloudServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpringCloudService)SetBuildAgentPoolSize(val *string) {
	if err := j.validateSetBuildAgentPoolSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildAgentPoolSize",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetLogStreamPublicEndpointEnabled(val interface{}) {
	if err := j.validateSetLogStreamPublicEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamPublicEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetManagedEnvironmentId(val *string) {
	if err := j.validateSetManagedEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedEnvironmentId",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetServiceRegistryEnabled(val interface{}) {
	if err := j.validateSetServiceRegistryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRegistryEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetSkuTier(val *string) {
	if err := j.validateSetSkuTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuTier",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SpringCloudService)SetZoneRedundant(val interface{}) {
	if err := j.validateSetZoneRedundantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneRedundant",
		val,
	)
}

// Generates CDKTF code for importing a SpringCloudService resource upon running "cdktf plan <stack-name>".
func SpringCloudService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSpringCloudService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
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
func SpringCloudService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudService",
		"isTerraformResource",
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

func (s *jsiiProxy_SpringCloudService) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SpringCloudService) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpringCloudService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SpringCloudService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudService) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SpringCloudService) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudService) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpringCloudService) PutConfigServerGitSetting(value *SpringCloudServiceConfigServerGitSetting) {
	if err := s.validatePutConfigServerGitSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfigServerGitSetting",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutContainerRegistry(value interface{}) {
	if err := s.validatePutContainerRegistryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putContainerRegistry",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutDefaultBuildService(value *SpringCloudServiceDefaultBuildService) {
	if err := s.validatePutDefaultBuildServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDefaultBuildService",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutMarketplace(value *SpringCloudServiceMarketplace) {
	if err := s.validatePutMarketplaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMarketplace",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutNetwork(value *SpringCloudServiceNetwork) {
	if err := s.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetwork",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutTimeouts(value *SpringCloudServiceTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudService) PutTrace(value *SpringCloudServiceTrace) {
	if err := s.validatePutTraceParameters(value); err != nil {
		panic(err)
	}
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

func (s *jsiiProxy_SpringCloudService) ResetContainerRegistry() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerRegistry",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetDefaultBuildService() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultBuildService",
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

func (s *jsiiProxy_SpringCloudService) ResetManagedEnvironmentId() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedEnvironmentId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudService) ResetMarketplace() {
	_jsii_.InvokeVoid(
		s,
		"resetMarketplace",
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

func (s *jsiiProxy_SpringCloudService) ResetSkuTier() {
	_jsii_.InvokeVoid(
		s,
		"resetSkuTier",
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

func (s *jsiiProxy_SpringCloudService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
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

