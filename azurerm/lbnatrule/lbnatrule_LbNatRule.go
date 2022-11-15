package lbnatrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/lbnatrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/lb_nat_rule azurerm_lb_nat_rule}.
type LbNatRule interface {
	cdktf.TerraformResource
	BackendAddressPoolId() *string
	SetBackendAddressPoolId(val *string)
	BackendAddressPoolIdInput() *string
	BackendIpConfigurationId() *string
	BackendPort() *float64
	SetBackendPort(val *float64)
	BackendPortInput() *float64
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableFloatingIp() interface{}
	SetEnableFloatingIp(val interface{})
	EnableFloatingIpInput() interface{}
	EnableTcpReset() interface{}
	SetEnableTcpReset(val interface{})
	EnableTcpResetInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FrontendIpConfigurationId() *string
	FrontendIpConfigurationName() *string
	SetFrontendIpConfigurationName(val *string)
	FrontendIpConfigurationNameInput() *string
	FrontendPort() *float64
	SetFrontendPort(val *float64)
	FrontendPortEnd() *float64
	SetFrontendPortEnd(val *float64)
	FrontendPortEndInput() *float64
	FrontendPortInput() *float64
	FrontendPortStart() *float64
	SetFrontendPortStart(val *float64)
	FrontendPortStartInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdleTimeoutInMinutes() *float64
	SetIdleTimeoutInMinutes(val *float64)
	IdleTimeoutInMinutesInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadbalancerId() *string
	SetLoadbalancerId(val *string)
	LoadbalancerIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LbNatRuleTimeoutsOutputReference
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
	PutTimeouts(value *LbNatRuleTimeouts)
	ResetBackendAddressPoolId()
	ResetEnableFloatingIp()
	ResetEnableTcpReset()
	ResetFrontendPort()
	ResetFrontendPortEnd()
	ResetFrontendPortStart()
	ResetId()
	ResetIdleTimeoutInMinutes()
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

// The jsii proxy struct for LbNatRule
type jsiiProxy_LbNatRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbNatRule) BackendAddressPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendAddressPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) BackendAddressPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendAddressPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) BackendIpConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendIpConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) BackendPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) BackendPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) EnableFloatingIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFloatingIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) EnableFloatingIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFloatingIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) EnableTcpReset() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTcpReset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) EnableTcpResetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTcpResetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendIpConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendIpConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendIpConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendPortEnd() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPortEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendPortEndInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPortEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendPortStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPortStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) FrontendPortStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPortStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) LoadbalancerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) LoadbalancerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) Timeouts() LbNatRuleTimeoutsOutputReference {
	var returns LbNatRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbNatRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/lb_nat_rule azurerm_lb_nat_rule} Resource.
func NewLbNatRule(scope constructs.Construct, id *string, config *LbNatRuleConfig) LbNatRule {
	_init_.Initialize()

	if err := validateNewLbNatRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbNatRule{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.lbNatRule.LbNatRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/lb_nat_rule azurerm_lb_nat_rule} Resource.
func NewLbNatRule_Override(l LbNatRule, scope constructs.Construct, id *string, config *LbNatRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.lbNatRule.LbNatRule",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbNatRule)SetBackendAddressPoolId(val *string) {
	if err := j.validateSetBackendAddressPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendAddressPoolId",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetBackendPort(val *float64) {
	if err := j.validateSetBackendPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendPort",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetEnableFloatingIp(val interface{}) {
	if err := j.validateSetEnableFloatingIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFloatingIp",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetEnableTcpReset(val interface{}) {
	if err := j.validateSetEnableTcpResetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTcpReset",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetFrontendIpConfigurationName(val *string) {
	if err := j.validateSetFrontendIpConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendIpConfigurationName",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetFrontendPort(val *float64) {
	if err := j.validateSetFrontendPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendPort",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetFrontendPortEnd(val *float64) {
	if err := j.validateSetFrontendPortEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendPortEnd",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetFrontendPortStart(val *float64) {
	if err := j.validateSetFrontendPortStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendPortStart",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetLoadbalancerId(val *string) {
	if err := j.validateSetLoadbalancerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadbalancerId",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LbNatRule)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
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
func LbNatRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbNatRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.lbNatRule.LbNatRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbNatRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbNatRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.lbNatRule.LbNatRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbNatRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbNatRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.lbNatRule.LbNatRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbNatRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.lbNatRule.LbNatRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbNatRule) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbNatRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbNatRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbNatRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbNatRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbNatRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbNatRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbNatRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbNatRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbNatRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbNatRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbNatRule) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbNatRule) PutTimeouts(value *LbNatRuleTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbNatRule) ResetBackendAddressPoolId() {
	_jsii_.InvokeVoid(
		l,
		"resetBackendAddressPoolId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetEnableFloatingIp() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableFloatingIp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetEnableTcpReset() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableTcpReset",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetFrontendPort() {
	_jsii_.InvokeVoid(
		l,
		"resetFrontendPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetFrontendPortEnd() {
	_jsii_.InvokeVoid(
		l,
		"resetFrontendPortEnd",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetFrontendPortStart() {
	_jsii_.InvokeVoid(
		l,
		"resetFrontendPortStart",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		l,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbNatRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNatRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNatRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbNatRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

