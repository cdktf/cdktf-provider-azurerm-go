package storagemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/storagemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy azurerm_storage_management_policy}.
type StorageManagementPolicy interface {
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
	Rule() StorageManagementPolicyRuleList
	RuleInput() interface{}
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StorageManagementPolicyTimeoutsOutputReference
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
	PutRule(value interface{})
	PutTimeouts(value *StorageManagementPolicyTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRule()
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

// The jsii proxy struct for StorageManagementPolicy
type jsiiProxy_StorageManagementPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageManagementPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Rule() StorageManagementPolicyRuleList {
	var returns StorageManagementPolicyRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) Timeouts() StorageManagementPolicyTimeoutsOutputReference {
	var returns StorageManagementPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy azurerm_storage_management_policy} Resource.
func NewStorageManagementPolicy(scope constructs.Construct, id *string, config *StorageManagementPolicyConfig) StorageManagementPolicy {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy azurerm_storage_management_policy} Resource.
func NewStorageManagementPolicy_Override(s StorageManagementPolicy, scope constructs.Construct, id *string, config *StorageManagementPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicy",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicy) SetStorageAccountId(val *string) {
	_jsii_.Set(
		j,
		"storageAccountId",
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
func StorageManagementPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageManagementPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageManagementPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageManagementPolicy) PutRule(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicy) PutTimeouts(value *StorageManagementPolicyTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicy) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicy) ResetRule() {
	_jsii_.InvokeVoid(
		s,
		"resetRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#storage_account_id StorageManagementPolicy#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#id StorageManagementPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#rule StorageManagementPolicy#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#timeouts StorageManagementPolicy#timeouts}
	Timeouts *StorageManagementPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type StorageManagementPolicyRule struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#actions StorageManagementPolicy#actions}
	Actions *StorageManagementPolicyRuleActions `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#enabled StorageManagementPolicy#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#name StorageManagementPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#filters StorageManagementPolicy#filters}
	Filters *StorageManagementPolicyRuleFilters `field:"optional" json:"filters" yaml:"filters"`
}

type StorageManagementPolicyRuleActions struct {
	// base_blob block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#base_blob StorageManagementPolicy#base_blob}
	BaseBlob *StorageManagementPolicyRuleActionsBaseBlob `field:"optional" json:"baseBlob" yaml:"baseBlob"`
	// snapshot block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#snapshot StorageManagementPolicy#snapshot}
	Snapshot *StorageManagementPolicyRuleActionsSnapshot `field:"optional" json:"snapshot" yaml:"snapshot"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#version StorageManagementPolicy#version}
	Version *StorageManagementPolicyRuleActionsVersion `field:"optional" json:"version" yaml:"version"`
}

type StorageManagementPolicyRuleActionsBaseBlob struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#delete_after_days_since_last_access_time_greater_than StorageManagementPolicy#delete_after_days_since_last_access_time_greater_than}.
	DeleteAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceLastAccessTimeGreaterThan" yaml:"deleteAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#delete_after_days_since_modification_greater_than StorageManagementPolicy#delete_after_days_since_modification_greater_than}.
	DeleteAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceModificationGreaterThan" yaml:"deleteAfterDaysSinceModificationGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_archive_after_days_since_last_access_time_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_last_access_time_greater_than}.
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan" yaml:"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_archive_after_days_since_modification_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_modification_greater_than}.
	TierToArchiveAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceModificationGreaterThan" yaml:"tierToArchiveAfterDaysSinceModificationGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_cool_after_days_since_last_access_time_greater_than StorageManagementPolicy#tier_to_cool_after_days_since_last_access_time_greater_than}.
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan" yaml:"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_cool_after_days_since_modification_greater_than StorageManagementPolicy#tier_to_cool_after_days_since_modification_greater_than}.
	TierToCoolAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"tierToCoolAfterDaysSinceModificationGreaterThan" yaml:"tierToCoolAfterDaysSinceModificationGreaterThan"`
}

type StorageManagementPolicyRuleActionsBaseBlobOutputReference interface {
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
	DeleteAfterDaysSinceLastAccessTimeGreaterThan() *float64
	SetDeleteAfterDaysSinceLastAccessTimeGreaterThan(val *float64)
	DeleteAfterDaysSinceLastAccessTimeGreaterThanInput() *float64
	DeleteAfterDaysSinceModificationGreaterThan() *float64
	SetDeleteAfterDaysSinceModificationGreaterThan(val *float64)
	DeleteAfterDaysSinceModificationGreaterThanInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *StorageManagementPolicyRuleActionsBaseBlob
	SetInternalValue(val *StorageManagementPolicyRuleActionsBaseBlob)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan() *float64
	SetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan(val *float64)
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThanInput() *float64
	TierToArchiveAfterDaysSinceModificationGreaterThan() *float64
	SetTierToArchiveAfterDaysSinceModificationGreaterThan(val *float64)
	TierToArchiveAfterDaysSinceModificationGreaterThanInput() *float64
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThan() *float64
	SetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan(val *float64)
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThanInput() *float64
	TierToCoolAfterDaysSinceModificationGreaterThan() *float64
	SetTierToCoolAfterDaysSinceModificationGreaterThan(val *float64)
	TierToCoolAfterDaysSinceModificationGreaterThanInput() *float64
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
	ResetDeleteAfterDaysSinceLastAccessTimeGreaterThan()
	ResetDeleteAfterDaysSinceModificationGreaterThan()
	ResetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan()
	ResetTierToArchiveAfterDaysSinceModificationGreaterThan()
	ResetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan()
	ResetTierToCoolAfterDaysSinceModificationGreaterThan()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleActionsBaseBlobOutputReference
type jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceLastAccessTimeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceLastAccessTimeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceModificationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceModificationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) InternalValue() *StorageManagementPolicyRuleActionsBaseBlob {
	var returns *StorageManagementPolicyRuleActionsBaseBlob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceLastAccessTimeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceModificationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceModificationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceLastAccessTimeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceLastAccessTimeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceModificationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceModificationGreaterThanInput",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleActionsBaseBlobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyRuleActionsBaseBlobOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsBaseBlobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleActionsBaseBlobOutputReference_Override(s StorageManagementPolicyRuleActionsBaseBlobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsBaseBlobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetDeleteAfterDaysSinceLastAccessTimeGreaterThan(val *float64) {
	_jsii_.Set(
		j,
		"deleteAfterDaysSinceLastAccessTimeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetDeleteAfterDaysSinceModificationGreaterThan(val *float64) {
	_jsii_.Set(
		j,
		"deleteAfterDaysSinceModificationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetInternalValue(val *StorageManagementPolicyRuleActionsBaseBlob) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan(val *float64) {
	_jsii_.Set(
		j,
		"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetTierToArchiveAfterDaysSinceModificationGreaterThan(val *float64) {
	_jsii_.Set(
		j,
		"tierToArchiveAfterDaysSinceModificationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan(val *float64) {
	_jsii_.Set(
		j,
		"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) SetTierToCoolAfterDaysSinceModificationGreaterThan(val *float64) {
	_jsii_.Set(
		j,
		"tierToCoolAfterDaysSinceModificationGreaterThan",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetDeleteAfterDaysSinceLastAccessTimeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAfterDaysSinceLastAccessTimeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetDeleteAfterDaysSinceModificationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAfterDaysSinceModificationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToArchiveAfterDaysSinceModificationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToArchiveAfterDaysSinceModificationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ResetTierToCoolAfterDaysSinceModificationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetTierToCoolAfterDaysSinceModificationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyRuleActionsOutputReference interface {
	cdktf.ComplexObject
	BaseBlob() StorageManagementPolicyRuleActionsBaseBlobOutputReference
	BaseBlobInput() *StorageManagementPolicyRuleActionsBaseBlob
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
	// Experimental.
	Fqn() *string
	InternalValue() *StorageManagementPolicyRuleActions
	SetInternalValue(val *StorageManagementPolicyRuleActions)
	Snapshot() StorageManagementPolicyRuleActionsSnapshotOutputReference
	SnapshotInput() *StorageManagementPolicyRuleActionsSnapshot
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() StorageManagementPolicyRuleActionsVersionOutputReference
	VersionInput() *StorageManagementPolicyRuleActionsVersion
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
	PutBaseBlob(value *StorageManagementPolicyRuleActionsBaseBlob)
	PutSnapshot(value *StorageManagementPolicyRuleActionsSnapshot)
	PutVersion(value *StorageManagementPolicyRuleActionsVersion)
	ResetBaseBlob()
	ResetSnapshot()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleActionsOutputReference
type jsiiProxy_StorageManagementPolicyRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) BaseBlob() StorageManagementPolicyRuleActionsBaseBlobOutputReference {
	var returns StorageManagementPolicyRuleActionsBaseBlobOutputReference
	_jsii_.Get(
		j,
		"baseBlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) BaseBlobInput() *StorageManagementPolicyRuleActionsBaseBlob {
	var returns *StorageManagementPolicyRuleActionsBaseBlob
	_jsii_.Get(
		j,
		"baseBlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) InternalValue() *StorageManagementPolicyRuleActions {
	var returns *StorageManagementPolicyRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) Snapshot() StorageManagementPolicyRuleActionsSnapshotOutputReference {
	var returns StorageManagementPolicyRuleActionsSnapshotOutputReference
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) SnapshotInput() *StorageManagementPolicyRuleActionsSnapshot {
	var returns *StorageManagementPolicyRuleActionsSnapshot
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) Version() StorageManagementPolicyRuleActionsVersionOutputReference {
	var returns StorageManagementPolicyRuleActionsVersionOutputReference
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) VersionInput() *StorageManagementPolicyRuleActionsVersion {
	var returns *StorageManagementPolicyRuleActionsVersion
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyRuleActionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleActionsOutputReference_Override(s StorageManagementPolicyRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) SetInternalValue(val *StorageManagementPolicyRuleActions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) PutBaseBlob(value *StorageManagementPolicyRuleActionsBaseBlob) {
	_jsii_.InvokeVoid(
		s,
		"putBaseBlob",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) PutSnapshot(value *StorageManagementPolicyRuleActionsSnapshot) {
	_jsii_.InvokeVoid(
		s,
		"putSnapshot",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) PutVersion(value *StorageManagementPolicyRuleActionsVersion) {
	_jsii_.InvokeVoid(
		s,
		"putVersion",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ResetBaseBlob() {
	_jsii_.InvokeVoid(
		s,
		"resetBaseBlob",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ResetSnapshot() {
	_jsii_.InvokeVoid(
		s,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyRuleActionsSnapshot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#change_tier_to_archive_after_days_since_creation StorageManagementPolicy#change_tier_to_archive_after_days_since_creation}.
	ChangeTierToArchiveAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToArchiveAfterDaysSinceCreation" yaml:"changeTierToArchiveAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#change_tier_to_cool_after_days_since_creation StorageManagementPolicy#change_tier_to_cool_after_days_since_creation}.
	ChangeTierToCoolAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToCoolAfterDaysSinceCreation" yaml:"changeTierToCoolAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#delete_after_days_since_creation_greater_than StorageManagementPolicy#delete_after_days_since_creation_greater_than}.
	DeleteAfterDaysSinceCreationGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceCreationGreaterThan" yaml:"deleteAfterDaysSinceCreationGreaterThan"`
}

type StorageManagementPolicyRuleActionsSnapshotOutputReference interface {
	cdktf.ComplexObject
	ChangeTierToArchiveAfterDaysSinceCreation() *float64
	SetChangeTierToArchiveAfterDaysSinceCreation(val *float64)
	ChangeTierToArchiveAfterDaysSinceCreationInput() *float64
	ChangeTierToCoolAfterDaysSinceCreation() *float64
	SetChangeTierToCoolAfterDaysSinceCreation(val *float64)
	ChangeTierToCoolAfterDaysSinceCreationInput() *float64
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
	DeleteAfterDaysSinceCreationGreaterThan() *float64
	SetDeleteAfterDaysSinceCreationGreaterThan(val *float64)
	DeleteAfterDaysSinceCreationGreaterThanInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *StorageManagementPolicyRuleActionsSnapshot
	SetInternalValue(val *StorageManagementPolicyRuleActionsSnapshot)
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
	ResetChangeTierToArchiveAfterDaysSinceCreation()
	ResetChangeTierToCoolAfterDaysSinceCreation()
	ResetDeleteAfterDaysSinceCreationGreaterThan()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleActionsSnapshotOutputReference
type jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ChangeTierToArchiveAfterDaysSinceCreation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToArchiveAfterDaysSinceCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ChangeTierToArchiveAfterDaysSinceCreationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToArchiveAfterDaysSinceCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ChangeTierToCoolAfterDaysSinceCreation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToCoolAfterDaysSinceCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ChangeTierToCoolAfterDaysSinceCreationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToCoolAfterDaysSinceCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) DeleteAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) DeleteAfterDaysSinceCreationGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreationGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) InternalValue() *StorageManagementPolicyRuleActionsSnapshot {
	var returns *StorageManagementPolicyRuleActionsSnapshot
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleActionsSnapshotOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyRuleActionsSnapshotOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsSnapshotOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleActionsSnapshotOutputReference_Override(s StorageManagementPolicyRuleActionsSnapshotOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsSnapshotOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) SetChangeTierToArchiveAfterDaysSinceCreation(val *float64) {
	_jsii_.Set(
		j,
		"changeTierToArchiveAfterDaysSinceCreation",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) SetChangeTierToCoolAfterDaysSinceCreation(val *float64) {
	_jsii_.Set(
		j,
		"changeTierToCoolAfterDaysSinceCreation",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) SetDeleteAfterDaysSinceCreationGreaterThan(val *float64) {
	_jsii_.Set(
		j,
		"deleteAfterDaysSinceCreationGreaterThan",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) SetInternalValue(val *StorageManagementPolicyRuleActionsSnapshot) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ResetChangeTierToArchiveAfterDaysSinceCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeTierToArchiveAfterDaysSinceCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ResetChangeTierToCoolAfterDaysSinceCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeTierToCoolAfterDaysSinceCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ResetDeleteAfterDaysSinceCreationGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAfterDaysSinceCreationGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyRuleActionsVersion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#change_tier_to_archive_after_days_since_creation StorageManagementPolicy#change_tier_to_archive_after_days_since_creation}.
	ChangeTierToArchiveAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToArchiveAfterDaysSinceCreation" yaml:"changeTierToArchiveAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#change_tier_to_cool_after_days_since_creation StorageManagementPolicy#change_tier_to_cool_after_days_since_creation}.
	ChangeTierToCoolAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToCoolAfterDaysSinceCreation" yaml:"changeTierToCoolAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#delete_after_days_since_creation StorageManagementPolicy#delete_after_days_since_creation}.
	DeleteAfterDaysSinceCreation *float64 `field:"optional" json:"deleteAfterDaysSinceCreation" yaml:"deleteAfterDaysSinceCreation"`
}

type StorageManagementPolicyRuleActionsVersionOutputReference interface {
	cdktf.ComplexObject
	ChangeTierToArchiveAfterDaysSinceCreation() *float64
	SetChangeTierToArchiveAfterDaysSinceCreation(val *float64)
	ChangeTierToArchiveAfterDaysSinceCreationInput() *float64
	ChangeTierToCoolAfterDaysSinceCreation() *float64
	SetChangeTierToCoolAfterDaysSinceCreation(val *float64)
	ChangeTierToCoolAfterDaysSinceCreationInput() *float64
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
	DeleteAfterDaysSinceCreation() *float64
	SetDeleteAfterDaysSinceCreation(val *float64)
	DeleteAfterDaysSinceCreationInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *StorageManagementPolicyRuleActionsVersion
	SetInternalValue(val *StorageManagementPolicyRuleActionsVersion)
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
	ResetChangeTierToArchiveAfterDaysSinceCreation()
	ResetChangeTierToCoolAfterDaysSinceCreation()
	ResetDeleteAfterDaysSinceCreation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleActionsVersionOutputReference
type jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ChangeTierToArchiveAfterDaysSinceCreation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToArchiveAfterDaysSinceCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ChangeTierToArchiveAfterDaysSinceCreationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToArchiveAfterDaysSinceCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ChangeTierToCoolAfterDaysSinceCreation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToCoolAfterDaysSinceCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ChangeTierToCoolAfterDaysSinceCreationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeTierToCoolAfterDaysSinceCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) DeleteAfterDaysSinceCreation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) DeleteAfterDaysSinceCreationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) InternalValue() *StorageManagementPolicyRuleActionsVersion {
	var returns *StorageManagementPolicyRuleActionsVersion
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleActionsVersionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyRuleActionsVersionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleActionsVersionOutputReference_Override(s StorageManagementPolicyRuleActionsVersionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) SetChangeTierToArchiveAfterDaysSinceCreation(val *float64) {
	_jsii_.Set(
		j,
		"changeTierToArchiveAfterDaysSinceCreation",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) SetChangeTierToCoolAfterDaysSinceCreation(val *float64) {
	_jsii_.Set(
		j,
		"changeTierToCoolAfterDaysSinceCreation",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) SetDeleteAfterDaysSinceCreation(val *float64) {
	_jsii_.Set(
		j,
		"deleteAfterDaysSinceCreation",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) SetInternalValue(val *StorageManagementPolicyRuleActionsVersion) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ResetChangeTierToArchiveAfterDaysSinceCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeTierToArchiveAfterDaysSinceCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ResetChangeTierToCoolAfterDaysSinceCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeTierToCoolAfterDaysSinceCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ResetDeleteAfterDaysSinceCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAfterDaysSinceCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyRuleFilters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#blob_types StorageManagementPolicy#blob_types}.
	BlobTypes *[]*string `field:"required" json:"blobTypes" yaml:"blobTypes"`
	// match_blob_index_tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#match_blob_index_tag StorageManagementPolicy#match_blob_index_tag}
	MatchBlobIndexTag interface{} `field:"optional" json:"matchBlobIndexTag" yaml:"matchBlobIndexTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#prefix_match StorageManagementPolicy#prefix_match}.
	PrefixMatch *[]*string `field:"optional" json:"prefixMatch" yaml:"prefixMatch"`
}

type StorageManagementPolicyRuleFiltersMatchBlobIndexTag struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#name StorageManagementPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#value StorageManagementPolicy#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#operation StorageManagementPolicy#operation}.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
}

type StorageManagementPolicyRuleFiltersMatchBlobIndexTagList interface {
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
	Get(index *float64) StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleFiltersMatchBlobIndexTagList
type jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleFiltersMatchBlobIndexTagList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StorageManagementPolicyRuleFiltersMatchBlobIndexTagList {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersMatchBlobIndexTagList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleFiltersMatchBlobIndexTagList_Override(s StorageManagementPolicyRuleFiltersMatchBlobIndexTagList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersMatchBlobIndexTagList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) Get(index *float64) StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference {
	var returns StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Operation() *string
	SetOperation(val *string)
	OperationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetOperation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference
type jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) OperationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference_Override(s StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) SetOperation(val *string) {
	_jsii_.Set(
		j,
		"operation",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) ResetOperation() {
	_jsii_.InvokeVoid(
		s,
		"resetOperation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyRuleFiltersOutputReference interface {
	cdktf.ComplexObject
	BlobTypes() *[]*string
	SetBlobTypes(val *[]*string)
	BlobTypesInput() *[]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() *StorageManagementPolicyRuleFilters
	SetInternalValue(val *StorageManagementPolicyRuleFilters)
	MatchBlobIndexTag() StorageManagementPolicyRuleFiltersMatchBlobIndexTagList
	MatchBlobIndexTagInput() interface{}
	PrefixMatch() *[]*string
	SetPrefixMatch(val *[]*string)
	PrefixMatchInput() *[]*string
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
	PutMatchBlobIndexTag(value interface{})
	ResetMatchBlobIndexTag()
	ResetPrefixMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleFiltersOutputReference
type jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) BlobTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blobTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) BlobTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blobTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) InternalValue() *StorageManagementPolicyRuleFilters {
	var returns *StorageManagementPolicyRuleFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) MatchBlobIndexTag() StorageManagementPolicyRuleFiltersMatchBlobIndexTagList {
	var returns StorageManagementPolicyRuleFiltersMatchBlobIndexTagList
	_jsii_.Get(
		j,
		"matchBlobIndexTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) MatchBlobIndexTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchBlobIndexTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) PrefixMatch() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prefixMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) PrefixMatchInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prefixMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyRuleFiltersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleFiltersOutputReference_Override(s StorageManagementPolicyRuleFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) SetBlobTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"blobTypes",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) SetInternalValue(val *StorageManagementPolicyRuleFilters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) SetPrefixMatch(val *[]*string) {
	_jsii_.Set(
		j,
		"prefixMatch",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) PutMatchBlobIndexTag(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putMatchBlobIndexTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) ResetMatchBlobIndexTag() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchBlobIndexTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) ResetPrefixMatch() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefixMatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyRuleList interface {
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
	Get(index *float64) StorageManagementPolicyRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleList
type jsiiProxy_StorageManagementPolicyRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StorageManagementPolicyRuleList {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleList_Override(s StorageManagementPolicyRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleList) Get(index *float64) StorageManagementPolicyRuleOutputReference {
	var returns StorageManagementPolicyRuleOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyRuleOutputReference interface {
	cdktf.ComplexObject
	Actions() StorageManagementPolicyRuleActionsOutputReference
	ActionsInput() *StorageManagementPolicyRuleActions
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
	Filters() StorageManagementPolicyRuleFiltersOutputReference
	FiltersInput() *StorageManagementPolicyRuleFilters
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutActions(value *StorageManagementPolicyRuleActions)
	PutFilters(value *StorageManagementPolicyRuleFilters)
	ResetFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleOutputReference
type jsiiProxy_StorageManagementPolicyRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) Actions() StorageManagementPolicyRuleActionsOutputReference {
	var returns StorageManagementPolicyRuleActionsOutputReference
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) ActionsInput() *StorageManagementPolicyRuleActions {
	var returns *StorageManagementPolicyRuleActions
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) Filters() StorageManagementPolicyRuleFiltersOutputReference {
	var returns StorageManagementPolicyRuleFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) FiltersInput() *StorageManagementPolicyRuleFilters {
	var returns *StorageManagementPolicyRuleFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageManagementPolicyRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleOutputReference_Override(s StorageManagementPolicyRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) PutActions(value *StorageManagementPolicyRuleActions) {
	_jsii_.InvokeVoid(
		s,
		"putActions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) PutFilters(value *StorageManagementPolicyRuleFilters) {
	_jsii_.InvokeVoid(
		s,
		"putFilters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		s,
		"resetFilters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageManagementPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#create StorageManagementPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#delete StorageManagementPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#read StorageManagementPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#update StorageManagementPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type StorageManagementPolicyTimeoutsOutputReference interface {
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

// The jsii proxy struct for StorageManagementPolicyTimeoutsOutputReference
type jsiiProxy_StorageManagementPolicyTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageManagementPolicyTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyTimeoutsOutputReference_Override(s StorageManagementPolicyTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

