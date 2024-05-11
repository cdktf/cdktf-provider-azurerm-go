// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/managementgrouppolicyassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/management_group_policy_assignment azurerm_management_group_policy_assignment}.
type ManagementGroupPolicyAssignment interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Enforce() interface{}
	SetEnforce(val interface{})
	EnforceInput() interface{}
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
	Identity() ManagementGroupPolicyAssignmentIdentityOutputReference
	IdentityInput() *ManagementGroupPolicyAssignmentIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagementGroupId() *string
	SetManagementGroupId(val *string)
	ManagementGroupIdInput() *string
	Metadata() *string
	SetMetadata(val *string)
	MetadataInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NonComplianceMessage() ManagementGroupPolicyAssignmentNonComplianceMessageList
	NonComplianceMessageInput() interface{}
	NotScopes() *[]*string
	SetNotScopes(val *[]*string)
	NotScopesInput() *[]*string
	Overrides() ManagementGroupPolicyAssignmentOverridesList
	OverridesInput() interface{}
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	PolicyDefinitionId() *string
	SetPolicyDefinitionId(val *string)
	PolicyDefinitionIdInput() *string
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
	ResourceSelectors() ManagementGroupPolicyAssignmentResourceSelectorsList
	ResourceSelectorsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ManagementGroupPolicyAssignmentTimeoutsOutputReference
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
	PutIdentity(value *ManagementGroupPolicyAssignmentIdentity)
	PutNonComplianceMessage(value interface{})
	PutOverrides(value interface{})
	PutResourceSelectors(value interface{})
	PutTimeouts(value *ManagementGroupPolicyAssignmentTimeouts)
	ResetDescription()
	ResetDisplayName()
	ResetEnforce()
	ResetId()
	ResetIdentity()
	ResetLocation()
	ResetMetadata()
	ResetNonComplianceMessage()
	ResetNotScopes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverrides()
	ResetParameters()
	ResetResourceSelectors()
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

// The jsii proxy struct for ManagementGroupPolicyAssignment
type jsiiProxy_ManagementGroupPolicyAssignment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Enforce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) EnforceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Identity() ManagementGroupPolicyAssignmentIdentityOutputReference {
	var returns ManagementGroupPolicyAssignmentIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) IdentityInput() *ManagementGroupPolicyAssignmentIdentity {
	var returns *ManagementGroupPolicyAssignmentIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) ManagementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) ManagementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) NonComplianceMessage() ManagementGroupPolicyAssignmentNonComplianceMessageList {
	var returns ManagementGroupPolicyAssignmentNonComplianceMessageList
	_jsii_.Get(
		j,
		"nonComplianceMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) NonComplianceMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonComplianceMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) NotScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) NotScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Overrides() ManagementGroupPolicyAssignmentOverridesList {
	var returns ManagementGroupPolicyAssignmentOverridesList
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) OverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) PolicyDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) PolicyDefinitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) ResourceSelectors() ManagementGroupPolicyAssignmentResourceSelectorsList {
	var returns ManagementGroupPolicyAssignmentResourceSelectorsList
	_jsii_.Get(
		j,
		"resourceSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) ResourceSelectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) Timeouts() ManagementGroupPolicyAssignmentTimeoutsOutputReference {
	var returns ManagementGroupPolicyAssignmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/management_group_policy_assignment azurerm_management_group_policy_assignment} Resource.
func NewManagementGroupPolicyAssignment(scope constructs.Construct, id *string, config *ManagementGroupPolicyAssignmentConfig) ManagementGroupPolicyAssignment {
	_init_.Initialize()

	if err := validateNewManagementGroupPolicyAssignmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagementGroupPolicyAssignment{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyAssignment.ManagementGroupPolicyAssignment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/management_group_policy_assignment azurerm_management_group_policy_assignment} Resource.
func NewManagementGroupPolicyAssignment_Override(m ManagementGroupPolicyAssignment, scope constructs.Construct, id *string, config *ManagementGroupPolicyAssignmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyAssignment.ManagementGroupPolicyAssignment",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetEnforce(val interface{}) {
	if err := j.validateSetEnforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforce",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetManagementGroupId(val *string) {
	if err := j.validateSetManagementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementGroupId",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetMetadata(val *string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetNotScopes(val *[]*string) {
	if err := j.validateSetNotScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notScopes",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetParameters(val *string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetPolicyDefinitionId(val *string) {
	if err := j.validateSetPolicyDefinitionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDefinitionId",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyAssignment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ManagementGroupPolicyAssignment resource upon running "cdktf plan <stack-name>".
func ManagementGroupPolicyAssignment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateManagementGroupPolicyAssignment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managementGroupPolicyAssignment.ManagementGroupPolicyAssignment",
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
func ManagementGroupPolicyAssignment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagementGroupPolicyAssignment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managementGroupPolicyAssignment.ManagementGroupPolicyAssignment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagementGroupPolicyAssignment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagementGroupPolicyAssignment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managementGroupPolicyAssignment.ManagementGroupPolicyAssignment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagementGroupPolicyAssignment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagementGroupPolicyAssignment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managementGroupPolicyAssignment.ManagementGroupPolicyAssignment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagementGroupPolicyAssignment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.managementGroupPolicyAssignment.ManagementGroupPolicyAssignment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) PutIdentity(value *ManagementGroupPolicyAssignmentIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) PutNonComplianceMessage(value interface{}) {
	if err := m.validatePutNonComplianceMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNonComplianceMessage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) PutOverrides(value interface{}) {
	if err := m.validatePutOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOverrides",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) PutResourceSelectors(value interface{}) {
	if err := m.validatePutResourceSelectorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putResourceSelectors",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) PutTimeouts(value *ManagementGroupPolicyAssignmentTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetDisplayName() {
	_jsii_.InvokeVoid(
		m,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetEnforce() {
	_jsii_.InvokeVoid(
		m,
		"resetEnforce",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetLocation() {
	_jsii_.InvokeVoid(
		m,
		"resetLocation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetNonComplianceMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetNonComplianceMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetNotScopes() {
	_jsii_.InvokeVoid(
		m,
		"resetNotScopes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetOverrides() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrides",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetParameters() {
	_jsii_.InvokeVoid(
		m,
		"resetParameters",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetResourceSelectors() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceSelectors",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyAssignment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

