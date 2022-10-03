package managementgrouppolicyremediation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/managementgrouppolicyremediation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation azurerm_management_group_policy_remediation}.
type ManagementGroupPolicyRemediation interface {
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
	ManagementGroupId() *string
	SetManagementGroupId(val *string)
	ManagementGroupIdInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ManagementGroupPolicyRemediationTimeoutsOutputReference
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
	PutTimeouts(value *ManagementGroupPolicyRemediationTimeouts)
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

// The jsii proxy struct for ManagementGroupPolicyRemediation
type jsiiProxy_ManagementGroupPolicyRemediation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) FailurePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failurePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) FailurePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failurePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) LocationFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) LocationFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ManagementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ManagementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ParallelDeployments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelDeployments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ParallelDeploymentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelDeploymentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) PolicyAssignmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyAssignmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) PolicyAssignmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyAssignmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) PolicyDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) PolicyDefinitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) PolicyDefinitionReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) PolicyDefinitionReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionReferenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ResourceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resourceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ResourceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resourceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ResourceDiscoveryMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceDiscoveryMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) ResourceDiscoveryModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceDiscoveryModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) Timeouts() ManagementGroupPolicyRemediationTimeoutsOutputReference {
	var returns ManagementGroupPolicyRemediationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation azurerm_management_group_policy_remediation} Resource.
func NewManagementGroupPolicyRemediation(scope constructs.Construct, id *string, config *ManagementGroupPolicyRemediationConfig) ManagementGroupPolicyRemediation {
	_init_.Initialize()

	j := jsiiProxy_ManagementGroupPolicyRemediation{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyRemediation.ManagementGroupPolicyRemediation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation azurerm_management_group_policy_remediation} Resource.
func NewManagementGroupPolicyRemediation_Override(m ManagementGroupPolicyRemediation, scope constructs.Construct, id *string, config *ManagementGroupPolicyRemediationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyRemediation.ManagementGroupPolicyRemediation",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetFailurePercentage(val *float64) {
	_jsii_.Set(
		j,
		"failurePercentage",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetLocationFilters(val *[]*string) {
	_jsii_.Set(
		j,
		"locationFilters",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetManagementGroupId(val *string) {
	_jsii_.Set(
		j,
		"managementGroupId",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetParallelDeployments(val *float64) {
	_jsii_.Set(
		j,
		"parallelDeployments",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetPolicyAssignmentId(val *string) {
	_jsii_.Set(
		j,
		"policyAssignmentId",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetPolicyDefinitionId(val *string) {
	_jsii_.Set(
		j,
		"policyDefinitionId",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetPolicyDefinitionReferenceId(val *string) {
	_jsii_.Set(
		j,
		"policyDefinitionReferenceId",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetResourceCount(val *float64) {
	_jsii_.Set(
		j,
		"resourceCount",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediation) SetResourceDiscoveryMode(val *string) {
	_jsii_.Set(
		j,
		"resourceDiscoveryMode",
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
func ManagementGroupPolicyRemediation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managementGroupPolicyRemediation.ManagementGroupPolicyRemediation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagementGroupPolicyRemediation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.managementGroupPolicyRemediation.ManagementGroupPolicyRemediation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) PutTimeouts(value *ManagementGroupPolicyRemediationTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetFailurePercentage() {
	_jsii_.InvokeVoid(
		m,
		"resetFailurePercentage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetLocationFilters() {
	_jsii_.InvokeVoid(
		m,
		"resetLocationFilters",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetParallelDeployments() {
	_jsii_.InvokeVoid(
		m,
		"resetParallelDeployments",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetPolicyDefinitionId() {
	_jsii_.InvokeVoid(
		m,
		"resetPolicyDefinitionId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetPolicyDefinitionReferenceId() {
	_jsii_.InvokeVoid(
		m,
		"resetPolicyDefinitionReferenceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetResourceCount() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetResourceDiscoveryMode() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceDiscoveryMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ManagementGroupPolicyRemediationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#management_group_id ManagementGroupPolicyRemediation#management_group_id}.
	ManagementGroupId *string `field:"required" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#name ManagementGroupPolicyRemediation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#policy_assignment_id ManagementGroupPolicyRemediation#policy_assignment_id}.
	PolicyAssignmentId *string `field:"required" json:"policyAssignmentId" yaml:"policyAssignmentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#failure_percentage ManagementGroupPolicyRemediation#failure_percentage}.
	FailurePercentage *float64 `field:"optional" json:"failurePercentage" yaml:"failurePercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#id ManagementGroupPolicyRemediation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#location_filters ManagementGroupPolicyRemediation#location_filters}.
	LocationFilters *[]*string `field:"optional" json:"locationFilters" yaml:"locationFilters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#parallel_deployments ManagementGroupPolicyRemediation#parallel_deployments}.
	ParallelDeployments *float64 `field:"optional" json:"parallelDeployments" yaml:"parallelDeployments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#policy_definition_id ManagementGroupPolicyRemediation#policy_definition_id}.
	PolicyDefinitionId *string `field:"optional" json:"policyDefinitionId" yaml:"policyDefinitionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#policy_definition_reference_id ManagementGroupPolicyRemediation#policy_definition_reference_id}.
	PolicyDefinitionReferenceId *string `field:"optional" json:"policyDefinitionReferenceId" yaml:"policyDefinitionReferenceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#resource_count ManagementGroupPolicyRemediation#resource_count}.
	ResourceCount *float64 `field:"optional" json:"resourceCount" yaml:"resourceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#resource_discovery_mode ManagementGroupPolicyRemediation#resource_discovery_mode}.
	ResourceDiscoveryMode *string `field:"optional" json:"resourceDiscoveryMode" yaml:"resourceDiscoveryMode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#timeouts ManagementGroupPolicyRemediation#timeouts}
	Timeouts *ManagementGroupPolicyRemediationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type ManagementGroupPolicyRemediationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#create ManagementGroupPolicyRemediation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#delete ManagementGroupPolicyRemediation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#read ManagementGroupPolicyRemediation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_remediation#update ManagementGroupPolicyRemediation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ManagementGroupPolicyRemediationTimeoutsOutputReference interface {
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

// The jsii proxy struct for ManagementGroupPolicyRemediationTimeoutsOutputReference
type jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewManagementGroupPolicyRemediationTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagementGroupPolicyRemediationTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyRemediation.ManagementGroupPolicyRemediationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagementGroupPolicyRemediationTimeoutsOutputReference_Override(m ManagementGroupPolicyRemediationTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyRemediation.ManagementGroupPolicyRemediationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyRemediationTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

