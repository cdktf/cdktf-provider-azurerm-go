package managementgrouppolicyexemption

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/managementgrouppolicyexemption/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption azurerm_management_group_policy_exemption}.
type ManagementGroupPolicyExemption interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExemptionCategory() *string
	SetExemptionCategory(val *string)
	ExemptionCategoryInput() *string
	ExpiresOn() *string
	SetExpiresOn(val *string)
	ExpiresOnInput() *string
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
	PolicyAssignmentId() *string
	SetPolicyAssignmentId(val *string)
	PolicyAssignmentIdInput() *string
	PolicyDefinitionReferenceIds() *[]*string
	SetPolicyDefinitionReferenceIds(val *[]*string)
	PolicyDefinitionReferenceIdsInput() *[]*string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ManagementGroupPolicyExemptionTimeoutsOutputReference
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
	PutTimeouts(value *ManagementGroupPolicyExemptionTimeouts)
	ResetDescription()
	ResetDisplayName()
	ResetExpiresOn()
	ResetId()
	ResetMetadata()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyDefinitionReferenceIds()
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

// The jsii proxy struct for ManagementGroupPolicyExemption
type jsiiProxy_ManagementGroupPolicyExemption struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) ExemptionCategory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exemptionCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) ExemptionCategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exemptionCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) ExpiresOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) ExpiresOnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) ManagementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) ManagementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) PolicyAssignmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyAssignmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) PolicyAssignmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyAssignmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) PolicyDefinitionReferenceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyDefinitionReferenceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) PolicyDefinitionReferenceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyDefinitionReferenceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) Timeouts() ManagementGroupPolicyExemptionTimeoutsOutputReference {
	var returns ManagementGroupPolicyExemptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption azurerm_management_group_policy_exemption} Resource.
func NewManagementGroupPolicyExemption(scope constructs.Construct, id *string, config *ManagementGroupPolicyExemptionConfig) ManagementGroupPolicyExemption {
	_init_.Initialize()

	j := jsiiProxy_ManagementGroupPolicyExemption{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyExemption.ManagementGroupPolicyExemption",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption azurerm_management_group_policy_exemption} Resource.
func NewManagementGroupPolicyExemption_Override(m ManagementGroupPolicyExemption, scope constructs.Construct, id *string, config *ManagementGroupPolicyExemptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyExemption.ManagementGroupPolicyExemption",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetExemptionCategory(val *string) {
	_jsii_.Set(
		j,
		"exemptionCategory",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetExpiresOn(val *string) {
	_jsii_.Set(
		j,
		"expiresOn",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetManagementGroupId(val *string) {
	_jsii_.Set(
		j,
		"managementGroupId",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetMetadata(val *string) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetPolicyAssignmentId(val *string) {
	_jsii_.Set(
		j,
		"policyAssignmentId",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetPolicyDefinitionReferenceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"policyDefinitionReferenceIds",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemption) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
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
func ManagementGroupPolicyExemption_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.managementGroupPolicyExemption.ManagementGroupPolicyExemption",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagementGroupPolicyExemption_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.managementGroupPolicyExemption.ManagementGroupPolicyExemption",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) PutTimeouts(value *ManagementGroupPolicyExemptionTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ResetDisplayName() {
	_jsii_.InvokeVoid(
		m,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ResetExpiresOn() {
	_jsii_.InvokeVoid(
		m,
		"resetExpiresOn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ResetMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ResetPolicyDefinitionReferenceIds() {
	_jsii_.InvokeVoid(
		m,
		"resetPolicyDefinitionReferenceIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemption) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ManagementGroupPolicyExemptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#exemption_category ManagementGroupPolicyExemption#exemption_category}.
	ExemptionCategory *string `field:"required" json:"exemptionCategory" yaml:"exemptionCategory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#management_group_id ManagementGroupPolicyExemption#management_group_id}.
	ManagementGroupId *string `field:"required" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#name ManagementGroupPolicyExemption#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#policy_assignment_id ManagementGroupPolicyExemption#policy_assignment_id}.
	PolicyAssignmentId *string `field:"required" json:"policyAssignmentId" yaml:"policyAssignmentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#description ManagementGroupPolicyExemption#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#display_name ManagementGroupPolicyExemption#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#expires_on ManagementGroupPolicyExemption#expires_on}.
	ExpiresOn *string `field:"optional" json:"expiresOn" yaml:"expiresOn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#id ManagementGroupPolicyExemption#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#metadata ManagementGroupPolicyExemption#metadata}.
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#policy_definition_reference_ids ManagementGroupPolicyExemption#policy_definition_reference_ids}.
	PolicyDefinitionReferenceIds *[]*string `field:"optional" json:"policyDefinitionReferenceIds" yaml:"policyDefinitionReferenceIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#timeouts ManagementGroupPolicyExemption#timeouts}
	Timeouts *ManagementGroupPolicyExemptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type ManagementGroupPolicyExemptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#create ManagementGroupPolicyExemption#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#delete ManagementGroupPolicyExemption#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#read ManagementGroupPolicyExemption#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_exemption#update ManagementGroupPolicyExemption#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ManagementGroupPolicyExemptionTimeoutsOutputReference interface {
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

// The jsii proxy struct for ManagementGroupPolicyExemptionTimeoutsOutputReference
type jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewManagementGroupPolicyExemptionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagementGroupPolicyExemptionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyExemption.ManagementGroupPolicyExemptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagementGroupPolicyExemptionTimeoutsOutputReference_Override(m ManagementGroupPolicyExemptionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managementGroupPolicyExemption.ManagementGroupPolicyExemptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagementGroupPolicyExemptionTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
