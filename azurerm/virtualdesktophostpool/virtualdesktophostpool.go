package virtualdesktophostpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/virtualdesktophostpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool azurerm_virtual_desktop_host_pool}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	ResetScheduledAgentUpdates()
	ResetStartVmOnConnect()
	ResetTags()
	ResetTimeouts()
	ResetValidateEnvironment()
	SynthesizeAttributes() *map[string]interface{}
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

func (j *jsiiProxy_VirtualDesktopHostPool) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool azurerm_virtual_desktop_host_pool} Resource.
func NewVirtualDesktopHostPool(scope constructs.Construct, id *string, config *VirtualDesktopHostPoolConfig) VirtualDesktopHostPool {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopHostPool{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool azurerm_virtual_desktop_host_pool} Resource.
func NewVirtualDesktopHostPool_Override(v VirtualDesktopHostPool, scope constructs.Construct, id *string, config *VirtualDesktopHostPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetCustomRdpProperties(val *string) {
	_jsii_.Set(
		j,
		"customRdpProperties",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetFriendlyName(val *string) {
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetLoadBalancerType(val *string) {
	_jsii_.Set(
		j,
		"loadBalancerType",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetMaximumSessionsAllowed(val *float64) {
	_jsii_.Set(
		j,
		"maximumSessionsAllowed",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetPersonalDesktopAssignmentType(val *string) {
	_jsii_.Set(
		j,
		"personalDesktopAssignmentType",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetPreferredAppGroupType(val *string) {
	_jsii_.Set(
		j,
		"preferredAppGroupType",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetStartVmOnConnect(val interface{}) {
	_jsii_.Set(
		j,
		"startVmOnConnect",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPool) SetValidateEnvironment(val interface{}) {
	_jsii_.Set(
		j,
		"validateEnvironment",
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
func VirtualDesktopHostPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPool",
		"isConstruct",
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

func (v *jsiiProxy_VirtualDesktopHostPool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) PutScheduledAgentUpdates(value *VirtualDesktopHostPoolScheduledAgentUpdates) {
	_jsii_.InvokeVoid(
		v,
		"putScheduledAgentUpdates",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPool) PutTimeouts(value *VirtualDesktopHostPoolTimeouts) {
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

type VirtualDesktopHostPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#load_balancer_type VirtualDesktopHostPool#load_balancer_type}.
	LoadBalancerType *string `field:"required" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#location VirtualDesktopHostPool#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#name VirtualDesktopHostPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#resource_group_name VirtualDesktopHostPool#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#type VirtualDesktopHostPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#custom_rdp_properties VirtualDesktopHostPool#custom_rdp_properties}.
	CustomRdpProperties *string `field:"optional" json:"customRdpProperties" yaml:"customRdpProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#description VirtualDesktopHostPool#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#friendly_name VirtualDesktopHostPool#friendly_name}.
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#id VirtualDesktopHostPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#maximum_sessions_allowed VirtualDesktopHostPool#maximum_sessions_allowed}.
	MaximumSessionsAllowed *float64 `field:"optional" json:"maximumSessionsAllowed" yaml:"maximumSessionsAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#personal_desktop_assignment_type VirtualDesktopHostPool#personal_desktop_assignment_type}.
	PersonalDesktopAssignmentType *string `field:"optional" json:"personalDesktopAssignmentType" yaml:"personalDesktopAssignmentType"`
	// Preferred App Group type to display.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#preferred_app_group_type VirtualDesktopHostPool#preferred_app_group_type}
	PreferredAppGroupType *string `field:"optional" json:"preferredAppGroupType" yaml:"preferredAppGroupType"`
	// scheduled_agent_updates block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#scheduled_agent_updates VirtualDesktopHostPool#scheduled_agent_updates}
	ScheduledAgentUpdates *VirtualDesktopHostPoolScheduledAgentUpdates `field:"optional" json:"scheduledAgentUpdates" yaml:"scheduledAgentUpdates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#start_vm_on_connect VirtualDesktopHostPool#start_vm_on_connect}.
	StartVmOnConnect interface{} `field:"optional" json:"startVmOnConnect" yaml:"startVmOnConnect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#tags VirtualDesktopHostPool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#timeouts VirtualDesktopHostPool#timeouts}
	Timeouts *VirtualDesktopHostPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#validate_environment VirtualDesktopHostPool#validate_environment}.
	ValidateEnvironment interface{} `field:"optional" json:"validateEnvironment" yaml:"validateEnvironment"`
}

type VirtualDesktopHostPoolScheduledAgentUpdates struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#enabled VirtualDesktopHostPool#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#schedule VirtualDesktopHostPool#schedule}
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#timezone VirtualDesktopHostPool#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#use_session_host_timezone VirtualDesktopHostPool#use_session_host_timezone}.
	UseSessionHostTimezone interface{} `field:"optional" json:"useSessionHostTimezone" yaml:"useSessionHostTimezone"`
}

type VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference interface {
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
	InternalValue() *VirtualDesktopHostPoolScheduledAgentUpdates
	SetInternalValue(val *VirtualDesktopHostPoolScheduledAgentUpdates)
	Schedule() VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList
	ScheduleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	UseSessionHostTimezone() interface{}
	SetUseSessionHostTimezone(val interface{})
	UseSessionHostTimezoneInput() interface{}
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
	PutSchedule(value interface{})
	ResetEnabled()
	ResetSchedule()
	ResetTimezone()
	ResetUseSessionHostTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference
type jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) InternalValue() *VirtualDesktopHostPoolScheduledAgentUpdates {
	var returns *VirtualDesktopHostPoolScheduledAgentUpdates
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) Schedule() VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList {
	var returns VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) UseSessionHostTimezone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSessionHostTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) UseSessionHostTimezoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSessionHostTimezoneInput",
		&returns,
	)
	return returns
}


func NewVirtualDesktopHostPoolScheduledAgentUpdatesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualDesktopHostPoolScheduledAgentUpdatesOutputReference_Override(v VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) SetInternalValue(val *VirtualDesktopHostPoolScheduledAgentUpdates) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) SetUseSessionHostTimezone(val interface{}) {
	_jsii_.Set(
		j,
		"useSessionHostTimezone",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) PutSchedule(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putSchedule",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		v,
		"resetSchedule",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		v,
		"resetTimezone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ResetUseSessionHostTimezone() {
	_jsii_.InvokeVoid(
		v,
		"resetUseSessionHostTimezone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualDesktopHostPoolScheduledAgentUpdatesSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#day_of_week VirtualDesktopHostPool#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#hour_of_day VirtualDesktopHostPool#hour_of_day}.
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
}

type VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList interface {
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
	Get(index *float64) VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList
type jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList_Override(v VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) Get(index *float64) VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference {
	var returns VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference interface {
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
	DayOfWeek() *string
	SetDayOfWeek(val *string)
	DayOfWeekInput() *string
	// Experimental.
	Fqn() *string
	HourOfDay() *float64
	SetHourOfDay(val *float64)
	HourOfDayInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference
type jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) DayOfWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) HourOfDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) HourOfDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference_Override(v VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) SetDayOfWeek(val *string) {
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) SetHourOfDay(val *float64) {
	_jsii_.Set(
		j,
		"hourOfDay",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualDesktopHostPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#create VirtualDesktopHostPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#delete VirtualDesktopHostPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#read VirtualDesktopHostPool#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#update VirtualDesktopHostPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type VirtualDesktopHostPoolTimeoutsOutputReference interface {
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

// The jsii proxy struct for VirtualDesktopHostPoolTimeoutsOutputReference
type jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewVirtualDesktopHostPoolTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualDesktopHostPoolTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPoolTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualDesktopHostPoolTimeoutsOutputReference_Override(v VirtualDesktopHostPoolTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopHostPool.VirtualDesktopHostPoolTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		v,
		"resetCreate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		v,
		"resetDelete",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		v,
		"resetRead",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		v,
		"resetUpdate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopHostPoolTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

