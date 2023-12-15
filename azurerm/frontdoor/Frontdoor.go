// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/frontdoor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/frontdoor azurerm_frontdoor}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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

func (j *jsiiProxy_Frontdoor) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/frontdoor azurerm_frontdoor} Resource.
func NewFrontdoor(scope constructs.Construct, id *string, config *FrontdoorConfig) Frontdoor {
	_init_.Initialize()

	if err := validateNewFrontdoorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Frontdoor{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/frontdoor azurerm_frontdoor} Resource.
func NewFrontdoor_Override(f Frontdoor, scope constructs.Construct, id *string, config *FrontdoorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_Frontdoor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetFriendlyName(val *string) {
	if err := j.validateSetFriendlyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetLoadBalancerEnabled(val interface{}) {
	if err := j.validateSetLoadBalancerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerEnabled",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_Frontdoor)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a Frontdoor resource upon running "cdktf plan <stack-name>".
func Frontdoor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFrontdoor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
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
func Frontdoor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFrontdoor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Frontdoor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFrontdoor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Frontdoor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFrontdoor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.frontdoor.Frontdoor",
		"isTerraformResource",
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

func (f *jsiiProxy_Frontdoor) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_Frontdoor) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_Frontdoor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_Frontdoor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Frontdoor) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_Frontdoor) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_Frontdoor) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_Frontdoor) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_Frontdoor) PutBackendPool(value interface{}) {
	if err := f.validatePutBackendPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putBackendPool",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutBackendPoolHealthProbe(value interface{}) {
	if err := f.validatePutBackendPoolHealthProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putBackendPoolHealthProbe",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutBackendPoolLoadBalancing(value interface{}) {
	if err := f.validatePutBackendPoolLoadBalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putBackendPoolLoadBalancing",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutBackendPoolSettings(value interface{}) {
	if err := f.validatePutBackendPoolSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putBackendPoolSettings",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutFrontendEndpoint(value interface{}) {
	if err := f.validatePutFrontendEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFrontendEndpoint",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutRoutingRule(value interface{}) {
	if err := f.validatePutRoutingRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putRoutingRule",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Frontdoor) PutTimeouts(value *FrontdoorTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

