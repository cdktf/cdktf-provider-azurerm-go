package subscriptionpolicyremediation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/subscriptionpolicyremediation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation azurerm_subscription_policy_remediation}.
type SubscriptionPolicyRemediation interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FailurePercentage() *float64
	SetFailurePercentage(val *float64)
	FailurePercentageInput() *float64
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
	LocationFilters() *[]*string
	SetLocationFilters(val *[]*string)
	LocationFiltersInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ParallelDeployments() *float64
	SetParallelDeployments(val *float64)
	ParallelDeploymentsInput() *float64
	PolicyAssignmentId() *string
	SetPolicyAssignmentId(val *string)
	PolicyAssignmentIdInput() *string
	PolicyDefinitionId() *string
	SetPolicyDefinitionId(val *string)
	PolicyDefinitionIdInput() *string
	PolicyDefinitionReferenceId() *string
	SetPolicyDefinitionReferenceId(val *string)
	PolicyDefinitionReferenceIdInput() *string
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
	ResourceCount() *float64
	SetResourceCount(val *float64)
	ResourceCountInput() *float64
	ResourceDiscoveryMode() *string
	SetResourceDiscoveryMode(val *string)
	ResourceDiscoveryModeInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SubscriptionPolicyRemediationTimeoutsOutputReference
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
	PutTimeouts(value *SubscriptionPolicyRemediationTimeouts)
	ResetFailurePercentage()
	ResetId()
	ResetLocationFilters()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParallelDeployments()
	ResetPolicyDefinitionId()
	ResetPolicyDefinitionReferenceId()
	ResetResourceCount()
	ResetResourceDiscoveryMode()
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

// The jsii proxy struct for SubscriptionPolicyRemediation
type jsiiProxy_SubscriptionPolicyRemediation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) FailurePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failurePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) FailurePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failurePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) LocationFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) LocationFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) ParallelDeployments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelDeployments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) ParallelDeploymentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelDeploymentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) PolicyAssignmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyAssignmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) PolicyAssignmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyAssignmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) PolicyDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) PolicyDefinitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) PolicyDefinitionReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) PolicyDefinitionReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionReferenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) ResourceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resourceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) ResourceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resourceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) ResourceDiscoveryMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceDiscoveryMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) ResourceDiscoveryModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceDiscoveryModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) Timeouts() SubscriptionPolicyRemediationTimeoutsOutputReference {
	var returns SubscriptionPolicyRemediationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation azurerm_subscription_policy_remediation} Resource.
func NewSubscriptionPolicyRemediation(scope constructs.Construct, id *string, config *SubscriptionPolicyRemediationConfig) SubscriptionPolicyRemediation {
	_init_.Initialize()

	j := jsiiProxy_SubscriptionPolicyRemediation{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation azurerm_subscription_policy_remediation} Resource.
func NewSubscriptionPolicyRemediation_Override(s SubscriptionPolicyRemediation, scope constructs.Construct, id *string, config *SubscriptionPolicyRemediationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetFailurePercentage(val *float64) {
	_jsii_.Set(
		j,
		"failurePercentage",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetLocationFilters(val *[]*string) {
	_jsii_.Set(
		j,
		"locationFilters",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetParallelDeployments(val *float64) {
	_jsii_.Set(
		j,
		"parallelDeployments",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetPolicyAssignmentId(val *string) {
	_jsii_.Set(
		j,
		"policyAssignmentId",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetPolicyDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"policyDefinitionId",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetPolicyDefinitionReferenceId(val *string) {
	_jsii_.Set(
		j,
		"policyDefinitionReferenceId",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetResourceCount(val *float64) {
	_jsii_.Set(
		j,
		"resourceCount",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetResourceDiscoveryMode(val *string) {
	_jsii_.Set(
		j,
		"resourceDiscoveryMode",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation) SetSubscriptionId(val *string) {
	_jsii_.Set(
		j,
		"subscriptionId",
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
func SubscriptionPolicyRemediation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SubscriptionPolicyRemediation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) PutTimeouts(value *SubscriptionPolicyRemediationTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetFailurePercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetFailurePercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetLocationFilters() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationFilters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetParallelDeployments() {
	_jsii_.InvokeVoid(
		s,
		"resetParallelDeployments",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetPolicyDefinitionId() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicyDefinitionId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetPolicyDefinitionReferenceId() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicyDefinitionReferenceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetResourceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetResourceDiscoveryMode() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceDiscoveryMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SubscriptionPolicyRemediationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#name SubscriptionPolicyRemediation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#policy_assignment_id SubscriptionPolicyRemediation#policy_assignment_id}.
	PolicyAssignmentId *string `field:"required" json:"policyAssignmentId" yaml:"policyAssignmentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#subscription_id SubscriptionPolicyRemediation#subscription_id}.
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#failure_percentage SubscriptionPolicyRemediation#failure_percentage}.
	FailurePercentage *float64 `field:"optional" json:"failurePercentage" yaml:"failurePercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#id SubscriptionPolicyRemediation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#location_filters SubscriptionPolicyRemediation#location_filters}.
	LocationFilters *[]*string `field:"optional" json:"locationFilters" yaml:"locationFilters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#parallel_deployments SubscriptionPolicyRemediation#parallel_deployments}.
	ParallelDeployments *float64 `field:"optional" json:"parallelDeployments" yaml:"parallelDeployments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#policy_definition_id SubscriptionPolicyRemediation#policy_definition_id}.
	PolicyDefinitionId *string `field:"optional" json:"policyDefinitionId" yaml:"policyDefinitionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#policy_definition_reference_id SubscriptionPolicyRemediation#policy_definition_reference_id}.
	PolicyDefinitionReferenceId *string `field:"optional" json:"policyDefinitionReferenceId" yaml:"policyDefinitionReferenceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#resource_count SubscriptionPolicyRemediation#resource_count}.
	ResourceCount *float64 `field:"optional" json:"resourceCount" yaml:"resourceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#resource_discovery_mode SubscriptionPolicyRemediation#resource_discovery_mode}.
	ResourceDiscoveryMode *string `field:"optional" json:"resourceDiscoveryMode" yaml:"resourceDiscoveryMode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#timeouts SubscriptionPolicyRemediation#timeouts}
	Timeouts *SubscriptionPolicyRemediationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type SubscriptionPolicyRemediationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#create SubscriptionPolicyRemediation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#delete SubscriptionPolicyRemediation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#read SubscriptionPolicyRemediation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_remediation#update SubscriptionPolicyRemediation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type SubscriptionPolicyRemediationTimeoutsOutputReference interface {
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

// The jsii proxy struct for SubscriptionPolicyRemediationTimeoutsOutputReference
type jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewSubscriptionPolicyRemediationTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SubscriptionPolicyRemediationTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSubscriptionPolicyRemediationTimeoutsOutputReference_Override(s SubscriptionPolicyRemediationTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediationTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

