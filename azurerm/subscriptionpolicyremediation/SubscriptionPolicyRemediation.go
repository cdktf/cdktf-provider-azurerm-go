// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptionpolicyremediation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/subscriptionpolicyremediation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/subscription_policy_remediation azurerm_subscription_policy_remediation}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	PutTimeouts(value *SubscriptionPolicyRemediationTimeouts)
	ResetFailurePercentage()
	ResetId()
	ResetLocationFilters()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParallelDeployments()
	ResetPolicyDefinitionReferenceId()
	ResetResourceCount()
	ResetResourceDiscoveryMode()
	ResetTimeouts()
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

func (j *jsiiProxy_SubscriptionPolicyRemediation) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/subscription_policy_remediation azurerm_subscription_policy_remediation} Resource.
func NewSubscriptionPolicyRemediation(scope constructs.Construct, id *string, config *SubscriptionPolicyRemediationConfig) SubscriptionPolicyRemediation {
	_init_.Initialize()

	if err := validateNewSubscriptionPolicyRemediationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SubscriptionPolicyRemediation{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/subscription_policy_remediation azurerm_subscription_policy_remediation} Resource.
func NewSubscriptionPolicyRemediation_Override(s SubscriptionPolicyRemediation, scope constructs.Construct, id *string, config *SubscriptionPolicyRemediationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetFailurePercentage(val *float64) {
	if err := j.validateSetFailurePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePercentage",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetLocationFilters(val *[]*string) {
	if err := j.validateSetLocationFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationFilters",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetParallelDeployments(val *float64) {
	if err := j.validateSetParallelDeploymentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelDeployments",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetPolicyAssignmentId(val *string) {
	if err := j.validateSetPolicyAssignmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyAssignmentId",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetPolicyDefinitionReferenceId(val *string) {
	if err := j.validateSetPolicyDefinitionReferenceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDefinitionReferenceId",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetResourceCount(val *float64) {
	if err := j.validateSetResourceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceCount",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetResourceDiscoveryMode(val *string) {
	if err := j.validateSetResourceDiscoveryModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceDiscoveryMode",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyRemediation)SetSubscriptionId(val *string) {
	if err := j.validateSetSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

// Generates CDKTF code for importing a SubscriptionPolicyRemediation resource upon running "cdktf plan <stack-name>".
func SubscriptionPolicyRemediation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSubscriptionPolicyRemediation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
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
func SubscriptionPolicyRemediation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSubscriptionPolicyRemediation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SubscriptionPolicyRemediation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSubscriptionPolicyRemediation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SubscriptionPolicyRemediation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSubscriptionPolicyRemediation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.subscriptionPolicyRemediation.SubscriptionPolicyRemediation",
		"isTerraformResource",
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) PutTimeouts(value *SubscriptionPolicyRemediationTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (s *jsiiProxy_SubscriptionPolicyRemediation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyRemediation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
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

