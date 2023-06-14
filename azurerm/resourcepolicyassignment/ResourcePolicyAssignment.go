package resourcepolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/resourcepolicyassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/resource_policy_assignment azurerm_resource_policy_assignment}.
type ResourcePolicyAssignment interface {
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
	Identity() ResourcePolicyAssignmentIdentityOutputReference
	IdentityInput() *ResourcePolicyAssignmentIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Metadata() *string
	SetMetadata(val *string)
	MetadataInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NonComplianceMessage() ResourcePolicyAssignmentNonComplianceMessageList
	NonComplianceMessageInput() interface{}
	NotScopes() *[]*string
	SetNotScopes(val *[]*string)
	NotScopesInput() *[]*string
	Overrides() ResourcePolicyAssignmentOverridesList
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
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResourceSelectors() ResourcePolicyAssignmentResourceSelectorsList
	ResourceSelectorsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ResourcePolicyAssignmentTimeoutsOutputReference
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
	PutIdentity(value *ResourcePolicyAssignmentIdentity)
	PutNonComplianceMessage(value interface{})
	PutOverrides(value interface{})
	PutResourceSelectors(value interface{})
	PutTimeouts(value *ResourcePolicyAssignmentTimeouts)
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
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ResourcePolicyAssignment
type jsiiProxy_ResourcePolicyAssignment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResourcePolicyAssignment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Enforce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) EnforceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Identity() ResourcePolicyAssignmentIdentityOutputReference {
	var returns ResourcePolicyAssignmentIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) IdentityInput() *ResourcePolicyAssignmentIdentity {
	var returns *ResourcePolicyAssignmentIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) NonComplianceMessage() ResourcePolicyAssignmentNonComplianceMessageList {
	var returns ResourcePolicyAssignmentNonComplianceMessageList
	_jsii_.Get(
		j,
		"nonComplianceMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) NonComplianceMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonComplianceMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) NotScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) NotScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Overrides() ResourcePolicyAssignmentOverridesList {
	var returns ResourcePolicyAssignmentOverridesList
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) OverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) PolicyDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) PolicyDefinitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDefinitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) ResourceSelectors() ResourcePolicyAssignmentResourceSelectorsList {
	var returns ResourcePolicyAssignmentResourceSelectorsList
	_jsii_.Get(
		j,
		"resourceSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) ResourceSelectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) Timeouts() ResourcePolicyAssignmentTimeoutsOutputReference {
	var returns ResourcePolicyAssignmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/resource_policy_assignment azurerm_resource_policy_assignment} Resource.
func NewResourcePolicyAssignment(scope constructs.Construct, id *string, config *ResourcePolicyAssignmentConfig) ResourcePolicyAssignment {
	_init_.Initialize()

	if err := validateNewResourcePolicyAssignmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourcePolicyAssignment{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourcePolicyAssignment.ResourcePolicyAssignment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/resource_policy_assignment azurerm_resource_policy_assignment} Resource.
func NewResourcePolicyAssignment_Override(r ResourcePolicyAssignment, scope constructs.Construct, id *string, config *ResourcePolicyAssignmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourcePolicyAssignment.ResourcePolicyAssignment",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetEnforce(val interface{}) {
	if err := j.validateSetEnforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforce",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetMetadata(val *string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetNotScopes(val *[]*string) {
	if err := j.validateSetNotScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notScopes",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetParameters(val *string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetPolicyDefinitionId(val *string) {
	if err := j.validateSetPolicyDefinitionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDefinitionId",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignment)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
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
func ResourcePolicyAssignment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourcePolicyAssignment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourcePolicyAssignment.ResourcePolicyAssignment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourcePolicyAssignment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourcePolicyAssignment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourcePolicyAssignment.ResourcePolicyAssignment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourcePolicyAssignment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourcePolicyAssignment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourcePolicyAssignment.ResourcePolicyAssignment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResourcePolicyAssignment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.resourcePolicyAssignment.ResourcePolicyAssignment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) PutIdentity(value *ResourcePolicyAssignmentIdentity) {
	if err := r.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIdentity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) PutNonComplianceMessage(value interface{}) {
	if err := r.validatePutNonComplianceMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNonComplianceMessage",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) PutOverrides(value interface{}) {
	if err := r.validatePutOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putOverrides",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) PutResourceSelectors(value interface{}) {
	if err := r.validatePutResourceSelectorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResourceSelectors",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) PutTimeouts(value *ResourcePolicyAssignmentTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetDescription() {
	_jsii_.InvokeVoid(
		r,
		"resetDescription",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetDisplayName() {
	_jsii_.InvokeVoid(
		r,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetEnforce() {
	_jsii_.InvokeVoid(
		r,
		"resetEnforce",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetIdentity() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetLocation() {
	_jsii_.InvokeVoid(
		r,
		"resetLocation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetMetadata() {
	_jsii_.InvokeVoid(
		r,
		"resetMetadata",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetNonComplianceMessage() {
	_jsii_.InvokeVoid(
		r,
		"resetNonComplianceMessage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetNotScopes() {
	_jsii_.InvokeVoid(
		r,
		"resetNotScopes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetOverrides() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrides",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetParameters() {
	_jsii_.InvokeVoid(
		r,
		"resetParameters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetResourceSelectors() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceSelectors",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

