// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktophostpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/virtualdesktophostpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/virtual_desktop_host_pool azurerm_virtual_desktop_host_pool}.
type VirtualDesktopHostPool interface {
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
	CustomRdpProperties() *string
	SetCustomRdpProperties(val *string)
	CustomRdpPropertiesInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerType() *string
	SetLoadBalancerType(val *string)
	LoadBalancerTypeInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaximumSessionsAllowed() *float64
	SetMaximumSessionsAllowed(val *float64)
	MaximumSessionsAllowedInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PersonalDesktopAssignmentType() *string
	SetPersonalDesktopAssignmentType(val *string)
	PersonalDesktopAssignmentTypeInput() *string
	PreferredAppGroupType() *string
	SetPreferredAppGroupType(val *string)
	PreferredAppGroupTypeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccess() *string
	SetPublicNetworkAccess(val *string)
	PublicNetworkAccessInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ScheduledAgentUpdates() VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference
	ScheduledAgentUpdatesInput() *VirtualDesktopHostPoolScheduledAgentUpdates
	StartVmOnConnect() interface{}
	SetStartVmOnConnect(val interface{})
	StartVmOnConnectInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualDesktopHostPoolTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ValidateEnvironment() interface{}
	SetValidateEnvironment(val interface{})
	ValidateEnvironmentInput() interface{}
	VmTemplate() *string
	SetVmTemplate(val *string)
	VmTemplateInput() *string
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
	PutScheduledAgentUpdates(value *VirtualDesktopHostPoolScheduledAgentUpdates)
	PutTimeouts(value *VirtualDesktopHostPoolTimeouts)
	ResetCustomRdpProperties()
	ResetDescription()
	ResetFriendlyName()
	ResetId()
	ResetMaximumSessionsAllowed()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersonalDesktopAssignmentType()
	ResetPreferredAppGroupType()
	ResetPublicNetworkAccess()
	ResetScheduledAgentUpdates()
	ResetStartVmOnConnect()
	ResetTags()
	ResetTimeouts()
	ResetValidateEnvironment()
	ResetVmTemplate()
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

// The jsii proxy struct for VirtualDesktopHostPool
type jsiiProxy_VirtualDesktopHostPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualDesktopHostPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) CustomRdpProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customRdpProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) CustomRdpPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customRdpPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) FriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) LoadBalancerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) LoadBalancerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) MaximumSessionsAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumSessionsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) MaximumSessionsAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumSessionsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) PersonalDesktopAssignmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalDesktopAssignmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) PersonalDesktopAssignmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalDesktopAssignmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) PreferredAppGroupType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredAppGroupType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) PreferredAppGroupTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredAppGroupTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) PublicNetworkAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) PublicNetworkAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) ScheduledAgentUpdates() VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference {
	var returns VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference
	_jsii_.Get(
		j,
		"scheduledAgentUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) ScheduledAgentUpdatesInput() *VirtualDesktopHostPoolScheduledAgentUpdates {
	var returns *VirtualDesktopHostPoolScheduledAgentUpdates
	_jsii_.Get(
		j,
		"scheduledAgentUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) StartVmOnConnect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startVmOnConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) StartVmOnConnectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startVmOnConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Timeouts() VirtualDesktopHostPoolTimeoutsOutputReference {
	var returns VirtualDesktopHostPoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) ValidateEnvironment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) ValidateEnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) VmTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPool) VmTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmTemplateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/virtual_desktop_host_pool azurerm_virtual_desktop_host_pool} Resource.
func NewVirtualDesktopHostPool(scope constructs.Construct, id *string, config *VirtualDesktopHostPoolConfig) VirtualDesktopHostPool {
	_init_.Initialize()

	if err := validateNewVirtualDesktopHostPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualDesktopHostPool{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/virtual_desktop_host_pool azurerm_virtual_desktop_host_pool} Resource.
func NewVirtualDesktopHostPool_Override(v VirtualDesktopHostPool, scope constructs.Construct, id *string, config *VirtualDesktopHostPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetCustomRdpProperties(val *string) {
	if err := j.validateSetCustomRdpPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customRdpProperties",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetFriendlyName(val *string) {
	if err := j.validateSetFriendlyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetLoadBalancerType(val *string) {
	if err := j.validateSetLoadBalancerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerType",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetMaximumSessionsAllowed(val *float64) {
	if err := j.validateSetMaximumSessionsAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumSessionsAllowed",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetPersonalDesktopAssignmentType(val *string) {
	if err := j.validateSetPersonalDesktopAssignmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"personalDesktopAssignmentType",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetPreferredAppGroupType(val *string) {
	if err := j.validateSetPreferredAppGroupTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredAppGroupType",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetPublicNetworkAccess(val *string) {
	if err := j.validateSetPublicNetworkAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccess",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetStartVmOnConnect(val interface{}) {
	if err := j.validateSetStartVmOnConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startVmOnConnect",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetValidateEnvironment(val interface{}) {
	if err := j.validateSetValidateEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateEnvironment",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool)SetVmTemplate(val *string) {
	if err := j.validateSetVmTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmTemplate",
		val,
	)
}

// Generates CDKTF code for importing a VirtualDesktopHostPool resource upon running "cdktf plan <stack-name>".
func VirtualDesktopHostPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualDesktopHostPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
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
func VirtualDesktopHostPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualDesktopHostPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualDesktopHostPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualDesktopHostPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualDesktopHostPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualDesktopHostPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualDesktopHostPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualDesktopHostPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualDesktopHostPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualDesktopHostPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualDesktopHostPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualDesktopHostPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualDesktopHostPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualDesktopHostPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualDesktopHostPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualDesktopHostPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualDesktopHostPool) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) PutScheduledAgentUpdates(value *VirtualDesktopHostPoolScheduledAgentUpdates) {
	if err := v.validatePutScheduledAgentUpdatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putScheduledAgentUpdates",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) PutTimeouts(value *VirtualDesktopHostPoolTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetCustomRdpProperties() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomRdpProperties",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetFriendlyName() {
	_jsii_.InvokeVoid(
		v,
		"resetFriendlyName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetMaximumSessionsAllowed() {
	_jsii_.InvokeVoid(
		v,
		"resetMaximumSessionsAllowed",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetPersonalDesktopAssignmentType() {
	_jsii_.InvokeVoid(
		v,
		"resetPersonalDesktopAssignmentType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetPreferredAppGroupType() {
	_jsii_.InvokeVoid(
		v,
		"resetPreferredAppGroupType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetPublicNetworkAccess() {
	_jsii_.InvokeVoid(
		v,
		"resetPublicNetworkAccess",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetScheduledAgentUpdates() {
	_jsii_.InvokeVoid(
		v,
		"resetScheduledAgentUpdates",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetStartVmOnConnect() {
	_jsii_.InvokeVoid(
		v,
		"resetStartVmOnConnect",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetValidateEnvironment() {
	_jsii_.InvokeVoid(
		v,
		"resetValidateEnvironment",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) ResetVmTemplate() {
	_jsii_.InvokeVoid(
		v,
		"resetVmTemplate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

