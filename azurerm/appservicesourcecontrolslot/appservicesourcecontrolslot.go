package appservicesourcecontrolslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/appservicesourcecontrolslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot azurerm_app_service_source_control_slot}.
type AppServiceSourceControlSlot interface {
	cdktf.TerraformResource
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
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
	GithubActionConfiguration() AppServiceSourceControlSlotGithubActionConfigurationOutputReference
	GithubActionConfigurationInput() *AppServiceSourceControlSlotGithubActionConfiguration
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
	RepoUrl() *string
	SetRepoUrl(val *string)
	RepoUrlInput() *string
	RollbackEnabled() interface{}
	SetRollbackEnabled(val interface{})
	RollbackEnabledInput() interface{}
	ScmType() *string
	SlotId() *string
	SetSlotId(val *string)
	SlotIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AppServiceSourceControlSlotTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseLocalGit() interface{}
	SetUseLocalGit(val interface{})
	UseLocalGitInput() interface{}
	UseManualIntegration() interface{}
	SetUseManualIntegration(val interface{})
	UseManualIntegrationInput() interface{}
	UseMercurial() interface{}
	SetUseMercurial(val interface{})
	UseMercurialInput() interface{}
	UsesGithubAction() cdktf.IResolvable
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
	PutGithubActionConfiguration(value *AppServiceSourceControlSlotGithubActionConfiguration)
	PutTimeouts(value *AppServiceSourceControlSlotTimeouts)
	ResetBranch()
	ResetGithubActionConfiguration()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRepoUrl()
	ResetRollbackEnabled()
	ResetTimeouts()
	ResetUseLocalGit()
	ResetUseManualIntegration()
	ResetUseMercurial()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppServiceSourceControlSlot
type jsiiProxy_AppServiceSourceControlSlot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) GithubActionConfiguration() AppServiceSourceControlSlotGithubActionConfigurationOutputReference {
	var returns AppServiceSourceControlSlotGithubActionConfigurationOutputReference
	_jsii_.Get(
		j,
		"githubActionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) GithubActionConfigurationInput() *AppServiceSourceControlSlotGithubActionConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfiguration
	_jsii_.Get(
		j,
		"githubActionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) RepoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) RepoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) RollbackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) RollbackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SlotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SlotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) Timeouts() AppServiceSourceControlSlotTimeoutsOutputReference {
	var returns AppServiceSourceControlSlotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) UseLocalGit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLocalGit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) UseLocalGitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLocalGitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) UseManualIntegration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManualIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) UseManualIntegrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManualIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) UseMercurial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMercurial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) UseMercurialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMercurialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlot) UsesGithubAction() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"usesGithubAction",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot azurerm_app_service_source_control_slot} Resource.
func NewAppServiceSourceControlSlot(scope constructs.Construct, id *string, config *AppServiceSourceControlSlotConfig) AppServiceSourceControlSlot {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlSlot{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot azurerm_app_service_source_control_slot} Resource.
func NewAppServiceSourceControlSlot_Override(a AppServiceSourceControlSlot, scope constructs.Construct, id *string, config *AppServiceSourceControlSlotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlot",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetBranch(val *string) {
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetRepoUrl(val *string) {
	_jsii_.Set(
		j,
		"repoUrl",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetRollbackEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rollbackEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetSlotId(val *string) {
	_jsii_.Set(
		j,
		"slotId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetUseLocalGit(val interface{}) {
	_jsii_.Set(
		j,
		"useLocalGit",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetUseManualIntegration(val interface{}) {
	_jsii_.Set(
		j,
		"useManualIntegration",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlot) SetUseMercurial(val interface{}) {
	_jsii_.Set(
		j,
		"useMercurial",
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
func AppServiceSourceControlSlot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppServiceSourceControlSlot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) PutGithubActionConfiguration(value *AppServiceSourceControlSlotGithubActionConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putGithubActionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) PutTimeouts(value *AppServiceSourceControlSlotTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetBranch() {
	_jsii_.InvokeVoid(
		a,
		"resetBranch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetGithubActionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetGithubActionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetRepoUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRepoUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetRollbackEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRollbackEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetUseLocalGit() {
	_jsii_.InvokeVoid(
		a,
		"resetUseLocalGit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetUseManualIntegration() {
	_jsii_.InvokeVoid(
		a,
		"resetUseManualIntegration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ResetUseMercurial() {
	_jsii_.InvokeVoid(
		a,
		"resetUseMercurial",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSourceControlSlotConfig struct {
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
	// The ID of the Linux or Windows Web App Slot.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#slot_id AppServiceSourceControlSlot#slot_id}
	SlotId *string `field:"required" json:"slotId" yaml:"slotId"`
	// The URL for the repository.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#branch AppServiceSourceControlSlot#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// github_action_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#github_action_configuration AppServiceSourceControlSlot#github_action_configuration}
	GithubActionConfiguration *AppServiceSourceControlSlotGithubActionConfiguration `field:"optional" json:"githubActionConfiguration" yaml:"githubActionConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#id AppServiceSourceControlSlot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The branch name to use for deployments.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#repo_url AppServiceSourceControlSlot#repo_url}
	RepoUrl *string `field:"optional" json:"repoUrl" yaml:"repoUrl"`
	// Should the Deployment Rollback be enabled? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#rollback_enabled AppServiceSourceControlSlot#rollback_enabled}
	RollbackEnabled interface{} `field:"optional" json:"rollbackEnabled" yaml:"rollbackEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#timeouts AppServiceSourceControlSlot#timeouts}
	Timeouts *AppServiceSourceControlSlotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Should the Slot use local Git configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#use_local_git AppServiceSourceControlSlot#use_local_git}
	UseLocalGit interface{} `field:"optional" json:"useLocalGit" yaml:"useLocalGit"`
	// Should code be deployed manually.
	//
	// Set to `true` to disable continuous integration, such as webhooks into online repos such as GitHub. Defaults to `false`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#use_manual_integration AppServiceSourceControlSlot#use_manual_integration}
	UseManualIntegration interface{} `field:"optional" json:"useManualIntegration" yaml:"useManualIntegration"`
	// The repository specified is Mercurial. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#use_mercurial AppServiceSourceControlSlot#use_mercurial}
	UseMercurial interface{} `field:"optional" json:"useMercurial" yaml:"useMercurial"`
}

type AppServiceSourceControlSlotGithubActionConfiguration struct {
	// code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#code_configuration AppServiceSourceControlSlot#code_configuration}
	CodeConfiguration *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#container_configuration AppServiceSourceControlSlot#container_configuration}
	ContainerConfiguration *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
	// Should the service generate the GitHub Action Workflow file. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#generate_workflow_file AppServiceSourceControlSlot#generate_workflow_file}
	GenerateWorkflowFile interface{} `field:"optional" json:"generateWorkflowFile" yaml:"generateWorkflowFile"`
}

type AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration struct {
	// The value to use for the Runtime Stack in the workflow file content for code base apps.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#runtime_stack AppServiceSourceControlSlot#runtime_stack}
	RuntimeStack *string `field:"required" json:"runtimeStack" yaml:"runtimeStack"`
	// The value to use for the Runtime Version in the workflow file content for code base apps.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#runtime_version AppServiceSourceControlSlot#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

type AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference interface {
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
	InternalValue() *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration
	SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration)
	RuntimeStack() *string
	SetRuntimeStack(val *string)
	RuntimeStackInput() *string
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference
type jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) InternalValue() *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) RuntimeStack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) RuntimeStackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeStackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference_Override(a AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) SetRuntimeStack(val *string) {
	_jsii_.Set(
		j,
		"runtimeStack",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) SetRuntimeVersion(val *string) {
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration struct {
	// The image name for the build.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#image_name AppServiceSourceControlSlot#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The server URL for the container registry where the build will be hosted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#registry_url AppServiceSourceControlSlot#registry_url}
	RegistryUrl *string `field:"required" json:"registryUrl" yaml:"registryUrl"`
	// The password used to upload the image to the container registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#registry_password AppServiceSourceControlSlot#registry_password}
	RegistryPassword *string `field:"optional" json:"registryPassword" yaml:"registryPassword"`
	// The username used to upload the image to the container registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#registry_username AppServiceSourceControlSlot#registry_username}
	RegistryUsername *string `field:"optional" json:"registryUsername" yaml:"registryUsername"`
}

type AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference interface {
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
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	InternalValue() *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration
	SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration)
	RegistryPassword() *string
	SetRegistryPassword(val *string)
	RegistryPasswordInput() *string
	RegistryUrl() *string
	SetRegistryUrl(val *string)
	RegistryUrlInput() *string
	RegistryUsername() *string
	SetRegistryUsername(val *string)
	RegistryUsernameInput() *string
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
	ResetRegistryPassword()
	ResetRegistryUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference
type jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) InternalValue() *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference_Override(a AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetRegistryPassword(val *string) {
	_jsii_.Set(
		j,
		"registryPassword",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetRegistryUrl(val *string) {
	_jsii_.Set(
		j,
		"registryUrl",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetRegistryUsername(val *string) {
	_jsii_.Set(
		j,
		"registryUsername",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ResetRegistryPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistryPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ResetRegistryUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistryUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSourceControlSlotGithubActionConfigurationOutputReference interface {
	cdktf.ComplexObject
	CodeConfiguration() AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference
	CodeConfigurationInput() *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration
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
	ContainerConfiguration() AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference
	ContainerConfigurationInput() *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GenerateWorkflowFile() interface{}
	SetGenerateWorkflowFile(val interface{})
	GenerateWorkflowFileInput() interface{}
	InternalValue() *AppServiceSourceControlSlotGithubActionConfiguration
	SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfiguration)
	LinuxAction() cdktf.IResolvable
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
	PutCodeConfiguration(value *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration)
	PutContainerConfiguration(value *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration)
	ResetCodeConfiguration()
	ResetContainerConfiguration()
	ResetGenerateWorkflowFile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSourceControlSlotGithubActionConfigurationOutputReference
type jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) CodeConfiguration() AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference {
	var returns AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference
	_jsii_.Get(
		j,
		"codeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) CodeConfigurationInput() *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration
	_jsii_.Get(
		j,
		"codeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ContainerConfiguration() AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference {
	var returns AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference
	_jsii_.Get(
		j,
		"containerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ContainerConfigurationInput() *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration
	_jsii_.Get(
		j,
		"containerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GenerateWorkflowFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateWorkflowFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GenerateWorkflowFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateWorkflowFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) InternalValue() *AppServiceSourceControlSlotGithubActionConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) LinuxAction() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"linuxAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlSlotGithubActionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlSlotGithubActionConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlSlotGithubActionConfigurationOutputReference_Override(a AppServiceSourceControlSlotGithubActionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) SetGenerateWorkflowFile(val interface{}) {
	_jsii_.Set(
		j,
		"generateWorkflowFile",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) PutCodeConfiguration(value *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putCodeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) PutContainerConfiguration(value *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putContainerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ResetCodeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ResetContainerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ResetGenerateWorkflowFile() {
	_jsii_.InvokeVoid(
		a,
		"resetGenerateWorkflowFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSourceControlSlotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#create AppServiceSourceControlSlot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#delete AppServiceSourceControlSlot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control_slot#read AppServiceSourceControlSlot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

type AppServiceSourceControlSlotTimeoutsOutputReference interface {
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSourceControlSlotTimeoutsOutputReference
type jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlSlotTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlSlotTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlSlotTimeoutsOutputReference_Override(a AppServiceSourceControlSlotTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		a,
		"resetRead",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

