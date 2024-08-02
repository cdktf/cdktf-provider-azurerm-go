// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/containerappenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/container_app_environment azurerm_container_app_environment}.
type ContainerAppEnvironment interface {
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
	CustomDomainVerificationId() *string
	DaprApplicationInsightsConnectionString() *string
	SetDaprApplicationInsightsConnectionString(val *string)
	DaprApplicationInsightsConnectionStringInput() *string
	DefaultDomain() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DockerBridgeCidr() *string
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
	InfrastructureResourceGroupName() *string
	SetInfrastructureResourceGroupName(val *string)
	InfrastructureResourceGroupNameInput() *string
	InfrastructureSubnetId() *string
	SetInfrastructureSubnetId(val *string)
	InfrastructureSubnetIdInput() *string
	InternalLoadBalancerEnabled() interface{}
	SetInternalLoadBalancerEnabled(val interface{})
	InternalLoadBalancerEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	MutualTlsEnabled() interface{}
	SetMutualTlsEnabled(val interface{})
	MutualTlsEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PlatformReservedCidr() *string
	PlatformReservedDnsIpAddress() *string
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
	StaticIpAddress() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContainerAppEnvironmentTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkloadProfile() ContainerAppEnvironmentWorkloadProfileList
	WorkloadProfileInput() interface{}
	ZoneRedundancyEnabled() interface{}
	SetZoneRedundancyEnabled(val interface{})
	ZoneRedundancyEnabledInput() interface{}
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
	PutTimeouts(value *ContainerAppEnvironmentTimeouts)
	PutWorkloadProfile(value interface{})
	ResetDaprApplicationInsightsConnectionString()
	ResetId()
	ResetInfrastructureResourceGroupName()
	ResetInfrastructureSubnetId()
	ResetInternalLoadBalancerEnabled()
	ResetLogAnalyticsWorkspaceId()
	ResetMutualTlsEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
	ResetWorkloadProfile()
	ResetZoneRedundancyEnabled()
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

// The jsii proxy struct for ContainerAppEnvironment
type jsiiProxy_ContainerAppEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerAppEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) DaprApplicationInsightsConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"daprApplicationInsightsConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) DaprApplicationInsightsConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"daprApplicationInsightsConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) DefaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) DockerBridgeCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerBridgeCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) InfrastructureResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) InfrastructureResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) InfrastructureSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) InfrastructureSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) InternalLoadBalancerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalLoadBalancerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) InternalLoadBalancerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalLoadBalancerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) MutualTlsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutualTlsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) MutualTlsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutualTlsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) PlatformReservedCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformReservedCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) PlatformReservedDnsIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformReservedDnsIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) StaticIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) Timeouts() ContainerAppEnvironmentTimeoutsOutputReference {
	var returns ContainerAppEnvironmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) WorkloadProfile() ContainerAppEnvironmentWorkloadProfileList {
	var returns ContainerAppEnvironmentWorkloadProfileList
	_jsii_.Get(
		j,
		"workloadProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) WorkloadProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workloadProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) ZoneRedundancyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundancyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironment) ZoneRedundancyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundancyEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/container_app_environment azurerm_container_app_environment} Resource.
func NewContainerAppEnvironment(scope constructs.Construct, id *string, config *ContainerAppEnvironmentConfig) ContainerAppEnvironment {
	_init_.Initialize()

	if err := validateNewContainerAppEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAppEnvironment{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerAppEnvironment.ContainerAppEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/container_app_environment azurerm_container_app_environment} Resource.
func NewContainerAppEnvironment_Override(c ContainerAppEnvironment, scope constructs.Construct, id *string, config *ContainerAppEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerAppEnvironment.ContainerAppEnvironment",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetDaprApplicationInsightsConnectionString(val *string) {
	if err := j.validateSetDaprApplicationInsightsConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daprApplicationInsightsConnectionString",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetInfrastructureResourceGroupName(val *string) {
	if err := j.validateSetInfrastructureResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetInfrastructureSubnetId(val *string) {
	if err := j.validateSetInfrastructureSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureSubnetId",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetInternalLoadBalancerEnabled(val interface{}) {
	if err := j.validateSetInternalLoadBalancerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalLoadBalancerEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetLogAnalyticsWorkspaceId(val *string) {
	if err := j.validateSetLogAnalyticsWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetMutualTlsEnabled(val interface{}) {
	if err := j.validateSetMutualTlsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutualTlsEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironment)SetZoneRedundancyEnabled(val interface{}) {
	if err := j.validateSetZoneRedundancyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneRedundancyEnabled",
		val,
	)
}

// Generates CDKTF code for importing a ContainerAppEnvironment resource upon running "cdktf plan <stack-name>".
func ContainerAppEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateContainerAppEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerAppEnvironment.ContainerAppEnvironment",
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
func ContainerAppEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerAppEnvironment.ContainerAppEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAppEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerAppEnvironment.ContainerAppEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAppEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerAppEnvironment.ContainerAppEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerAppEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.containerAppEnvironment.ContainerAppEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) PutTimeouts(value *ContainerAppEnvironmentTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) PutWorkloadProfile(value interface{}) {
	if err := c.validatePutWorkloadProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkloadProfile",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetDaprApplicationInsightsConnectionString() {
	_jsii_.InvokeVoid(
		c,
		"resetDaprApplicationInsightsConnectionString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetInfrastructureResourceGroupName() {
	_jsii_.InvokeVoid(
		c,
		"resetInfrastructureResourceGroupName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetInfrastructureSubnetId() {
	_jsii_.InvokeVoid(
		c,
		"resetInfrastructureSubnetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetInternalLoadBalancerEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetInternalLoadBalancerEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetLogAnalyticsWorkspaceId() {
	_jsii_.InvokeVoid(
		c,
		"resetLogAnalyticsWorkspaceId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetMutualTlsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetMutualTlsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetWorkloadProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkloadProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) ResetZoneRedundancyEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetZoneRedundancyEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

