// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbbackendaddresspool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/lbbackendaddresspool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/lb_backend_address_pool azurerm_lb_backend_address_pool}.
type LbBackendAddressPool interface {
	cdktf.TerraformResource
	BackendIpConfigurations() *[]*string
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
	InboundNatRules() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadbalancerId() *string
	SetLoadbalancerId(val *string)
	LoadbalancerIdInput() *string
	LoadBalancingRules() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutboundRules() *[]*string
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
	SynchronousMode() *string
	SetSynchronousMode(val *string)
	SynchronousModeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LbBackendAddressPoolTimeoutsOutputReference
	TimeoutsInput() interface{}
	TunnelInterface() LbBackendAddressPoolTunnelInterfaceList
	TunnelInterfaceInput() interface{}
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
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
	PutTimeouts(value *LbBackendAddressPoolTimeouts)
	PutTunnelInterface(value interface{})
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSynchronousMode()
	ResetTimeouts()
	ResetTunnelInterface()
	ResetVirtualNetworkId()
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

// The jsii proxy struct for LbBackendAddressPool
type jsiiProxy_LbBackendAddressPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbBackendAddressPool) BackendIpConfigurations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backendIpConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) InboundNatRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundNatRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) LoadbalancerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) LoadbalancerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) LoadBalancingRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) OutboundRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) SynchronousMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"synchronousMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) SynchronousModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"synchronousModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) Timeouts() LbBackendAddressPoolTimeoutsOutputReference {
	var returns LbBackendAddressPoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) TunnelInterface() LbBackendAddressPoolTunnelInterfaceList {
	var returns LbBackendAddressPoolTunnelInterfaceList
	_jsii_.Get(
		j,
		"tunnelInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) TunnelInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnelInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbBackendAddressPool) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/lb_backend_address_pool azurerm_lb_backend_address_pool} Resource.
func NewLbBackendAddressPool(scope constructs.Construct, id *string, config *LbBackendAddressPoolConfig) LbBackendAddressPool {
	_init_.Initialize()

	if err := validateNewLbBackendAddressPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbBackendAddressPool{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.lbBackendAddressPool.LbBackendAddressPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/lb_backend_address_pool azurerm_lb_backend_address_pool} Resource.
func NewLbBackendAddressPool_Override(l LbBackendAddressPool, scope constructs.Construct, id *string, config *LbBackendAddressPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.lbBackendAddressPool.LbBackendAddressPool",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetLoadbalancerId(val *string) {
	if err := j.validateSetLoadbalancerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadbalancerId",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetSynchronousMode(val *string) {
	if err := j.validateSetSynchronousModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"synchronousMode",
		val,
	)
}

func (j *jsiiProxy_LbBackendAddressPool)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

// Generates CDKTF code for importing a LbBackendAddressPool resource upon running "cdktf plan <stack-name>".
func LbBackendAddressPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLbBackendAddressPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.lbBackendAddressPool.LbBackendAddressPool",
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
func LbBackendAddressPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbBackendAddressPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.lbBackendAddressPool.LbBackendAddressPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbBackendAddressPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbBackendAddressPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.lbBackendAddressPool.LbBackendAddressPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbBackendAddressPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbBackendAddressPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.lbBackendAddressPool.LbBackendAddressPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbBackendAddressPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.lbBackendAddressPool.LbBackendAddressPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) PutTimeouts(value *LbBackendAddressPoolTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) PutTunnelInterface(value interface{}) {
	if err := l.validatePutTunnelInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTunnelInterface",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbBackendAddressPool) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbBackendAddressPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbBackendAddressPool) ResetSynchronousMode() {
	_jsii_.InvokeVoid(
		l,
		"resetSynchronousMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbBackendAddressPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbBackendAddressPool) ResetTunnelInterface() {
	_jsii_.InvokeVoid(
		l,
		"resetTunnelInterface",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbBackendAddressPool) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		l,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbBackendAddressPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbBackendAddressPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

