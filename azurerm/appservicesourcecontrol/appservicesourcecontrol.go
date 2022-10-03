package appservicesourcecontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/appservicesourcecontrol/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control azurerm_app_service_source_control}.
type AppServiceSourceControlA interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
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
	GithubActionConfiguration() AppServiceSourceControlGithubActionConfigurationOutputReference
	GithubActionConfigurationInput() *AppServiceSourceControlGithubActionConfiguration
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AppServiceSourceControlTimeoutsOutputReference
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
	PutGithubActionConfiguration(value *AppServiceSourceControlGithubActionConfiguration)
	PutTimeouts(value *AppServiceSourceControlTimeouts)
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

// The jsii proxy struct for AppServiceSourceControlA
type jsiiProxy_AppServiceSourceControlA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppServiceSourceControlA) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) GithubActionConfiguration() AppServiceSourceControlGithubActionConfigurationOutputReference {
	var returns AppServiceSourceControlGithubActionConfigurationOutputReference
	_jsii_.Get(
		j,
		"githubActionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) GithubActionConfigurationInput() *AppServiceSourceControlGithubActionConfiguration {
	var returns *AppServiceSourceControlGithubActionConfiguration
	_jsii_.Get(
		j,
		"githubActionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) RepoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) RepoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) RollbackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) RollbackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) Timeouts() AppServiceSourceControlTimeoutsOutputReference {
	var returns AppServiceSourceControlTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) UseLocalGit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLocalGit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) UseLocalGitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLocalGitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) UseManualIntegration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManualIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) UseManualIntegrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManualIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) UseMercurial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMercurial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) UseMercurialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMercurialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlA) UsesGithubAction() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"usesGithubAction",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control azurerm_app_service_source_control} Resource.
func NewAppServiceSourceControlA(scope constructs.Construct, id *string, config *AppServiceSourceControlAConfig) AppServiceSourceControlA {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlA{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control azurerm_app_service_source_control} Resource.
func NewAppServiceSourceControlA_Override(a AppServiceSourceControlA, scope constructs.Construct, id *string, config *AppServiceSourceControlAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetBranch(val *string) {
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetRepoUrl(val *string) {
	_jsii_.Set(
		j,
		"repoUrl",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetRollbackEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rollbackEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetUseLocalGit(val interface{}) {
	_jsii_.Set(
		j,
		"useLocalGit",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetUseManualIntegration(val interface{}) {
	_jsii_.Set(
		j,
		"useManualIntegration",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA) SetUseMercurial(val interface{}) {
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
func AppServiceSourceControlA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppServiceSourceControlA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) PutGithubActionConfiguration(value *AppServiceSourceControlGithubActionConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putGithubActionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) PutTimeouts(value *AppServiceSourceControlTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetBranch() {
	_jsii_.InvokeVoid(
		a,
		"resetBranch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetGithubActionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetGithubActionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetRepoUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRepoUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetRollbackEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRollbackEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetUseLocalGit() {
	_jsii_.InvokeVoid(
		a,
		"resetUseLocalGit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetUseManualIntegration() {
	_jsii_.InvokeVoid(
		a,
		"resetUseManualIntegration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) ResetUseMercurial() {
	_jsii_.InvokeVoid(
		a,
		"resetUseMercurial",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSourceControlAConfig struct {
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
	// The ID of the Windows or Linux Web App.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#app_id AppServiceSourceControlA#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The branch name to use for deployments.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#branch AppServiceSourceControlA#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// github_action_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#github_action_configuration AppServiceSourceControlA#github_action_configuration}
	GithubActionConfiguration *AppServiceSourceControlGithubActionConfiguration `field:"optional" json:"githubActionConfiguration" yaml:"githubActionConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#id AppServiceSourceControlA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The URL for the repository.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#repo_url AppServiceSourceControlA#repo_url}
	RepoUrl *string `field:"optional" json:"repoUrl" yaml:"repoUrl"`
	// Should the Deployment Rollback be enabled? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#rollback_enabled AppServiceSourceControlA#rollback_enabled}
	RollbackEnabled interface{} `field:"optional" json:"rollbackEnabled" yaml:"rollbackEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#timeouts AppServiceSourceControlA#timeouts}
	Timeouts *AppServiceSourceControlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Should the App use local Git configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#use_local_git AppServiceSourceControlA#use_local_git}
	UseLocalGit interface{} `field:"optional" json:"useLocalGit" yaml:"useLocalGit"`
	// Should code be deployed manually.
	//
	// Set to `false` to enable continuous integration, such as webhooks into online repos such as GitHub. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#use_manual_integration AppServiceSourceControlA#use_manual_integration}
	UseManualIntegration interface{} `field:"optional" json:"useManualIntegration" yaml:"useManualIntegration"`
	// The repository specified is Mercurial. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#use_mercurial AppServiceSourceControlA#use_mercurial}
	UseMercurial interface{} `field:"optional" json:"useMercurial" yaml:"useMercurial"`
}

type AppServiceSourceControlGithubActionConfiguration struct {
	// code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#code_configuration AppServiceSourceControlA#code_configuration}
	CodeConfiguration *AppServiceSourceControlGithubActionConfigurationCodeConfiguration `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#container_configuration AppServiceSourceControlA#container_configuration}
	ContainerConfiguration *AppServiceSourceControlGithubActionConfigurationContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
	// Should the service generate the GitHub Action Workflow file. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#generate_workflow_file AppServiceSourceControlA#generate_workflow_file}
	GenerateWorkflowFile interface{} `field:"optional" json:"generateWorkflowFile" yaml:"generateWorkflowFile"`
}

type AppServiceSourceControlGithubActionConfigurationCodeConfiguration struct {
	// The value to use for the Runtime Stack in the workflow file content for code base apps.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#runtime_stack AppServiceSourceControlA#runtime_stack}
	RuntimeStack *string `field:"required" json:"runtimeStack" yaml:"runtimeStack"`
	// The value to use for the Runtime Version in the workflow file content for code base apps.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#runtime_version AppServiceSourceControlA#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

type AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference interface {
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
	InternalValue() *AppServiceSourceControlGithubActionConfigurationCodeConfiguration
	SetInternalValue(val *AppServiceSourceControlGithubActionConfigurationCodeConfiguration)
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

// The jsii proxy struct for AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference
type jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) InternalValue() *AppServiceSourceControlGithubActionConfigurationCodeConfiguration {
	var returns *AppServiceSourceControlGithubActionConfigurationCodeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) RuntimeStack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) RuntimeStackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeStackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference_Override(a AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) SetInternalValue(val *AppServiceSourceControlGithubActionConfigurationCodeConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) SetRuntimeStack(val *string) {
	_jsii_.Set(
		j,
		"runtimeStack",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) SetRuntimeVersion(val *string) {
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSourceControlGithubActionConfigurationContainerConfiguration struct {
	// The image name for the build.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#image_name AppServiceSourceControlA#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The server URL for the container registry where the build will be hosted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#registry_url AppServiceSourceControlA#registry_url}
	RegistryUrl *string `field:"required" json:"registryUrl" yaml:"registryUrl"`
	// The password used to upload the image to the container registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#registry_password AppServiceSourceControlA#registry_password}
	RegistryPassword *string `field:"optional" json:"registryPassword" yaml:"registryPassword"`
	// The username used to upload the image to the container registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#registry_username AppServiceSourceControlA#registry_username}
	RegistryUsername *string `field:"optional" json:"registryUsername" yaml:"registryUsername"`
}

type AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference interface {
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
	InternalValue() *AppServiceSourceControlGithubActionConfigurationContainerConfiguration
	SetInternalValue(val *AppServiceSourceControlGithubActionConfigurationContainerConfiguration)
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

// The jsii proxy struct for AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference
type jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) InternalValue() *AppServiceSourceControlGithubActionConfigurationContainerConfiguration {
	var returns *AppServiceSourceControlGithubActionConfigurationContainerConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) RegistryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) RegistryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) RegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) RegistryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) RegistryUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) RegistryUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference_Override(a AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetInternalValue(val *AppServiceSourceControlGithubActionConfigurationContainerConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetRegistryPassword(val *string) {
	_jsii_.Set(
		j,
		"registryPassword",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetRegistryUrl(val *string) {
	_jsii_.Set(
		j,
		"registryUrl",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetRegistryUsername(val *string) {
	_jsii_.Set(
		j,
		"registryUsername",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) ResetRegistryPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistryPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) ResetRegistryUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistryUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSourceControlGithubActionConfigurationOutputReference interface {
	cdktf.ComplexObject
	CodeConfiguration() AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference
	CodeConfigurationInput() *AppServiceSourceControlGithubActionConfigurationCodeConfiguration
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
	ContainerConfiguration() AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference
	ContainerConfigurationInput() *AppServiceSourceControlGithubActionConfigurationContainerConfiguration
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
	InternalValue() *AppServiceSourceControlGithubActionConfiguration
	SetInternalValue(val *AppServiceSourceControlGithubActionConfiguration)
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
	PutCodeConfiguration(value *AppServiceSourceControlGithubActionConfigurationCodeConfiguration)
	PutContainerConfiguration(value *AppServiceSourceControlGithubActionConfigurationContainerConfiguration)
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

// The jsii proxy struct for AppServiceSourceControlGithubActionConfigurationOutputReference
type jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) CodeConfiguration() AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference {
	var returns AppServiceSourceControlGithubActionConfigurationCodeConfigurationOutputReference
	_jsii_.Get(
		j,
		"codeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) CodeConfigurationInput() *AppServiceSourceControlGithubActionConfigurationCodeConfiguration {
	var returns *AppServiceSourceControlGithubActionConfigurationCodeConfiguration
	_jsii_.Get(
		j,
		"codeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ContainerConfiguration() AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference {
	var returns AppServiceSourceControlGithubActionConfigurationContainerConfigurationOutputReference
	_jsii_.Get(
		j,
		"containerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ContainerConfigurationInput() *AppServiceSourceControlGithubActionConfigurationContainerConfiguration {
	var returns *AppServiceSourceControlGithubActionConfigurationContainerConfiguration
	_jsii_.Get(
		j,
		"containerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GenerateWorkflowFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateWorkflowFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GenerateWorkflowFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateWorkflowFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) InternalValue() *AppServiceSourceControlGithubActionConfiguration {
	var returns *AppServiceSourceControlGithubActionConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) LinuxAction() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"linuxAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlGithubActionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlGithubActionConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlGithubActionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlGithubActionConfigurationOutputReference_Override(a AppServiceSourceControlGithubActionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlGithubActionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) SetGenerateWorkflowFile(val interface{}) {
	_jsii_.Set(
		j,
		"generateWorkflowFile",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) SetInternalValue(val *AppServiceSourceControlGithubActionConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) PutCodeConfiguration(value *AppServiceSourceControlGithubActionConfigurationCodeConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putCodeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) PutContainerConfiguration(value *AppServiceSourceControlGithubActionConfigurationContainerConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putContainerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ResetCodeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ResetContainerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ResetGenerateWorkflowFile() {
	_jsii_.InvokeVoid(
		a,
		"resetGenerateWorkflowFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlGithubActionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSourceControlTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#create AppServiceSourceControlA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#delete AppServiceSourceControlA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_source_control#read AppServiceSourceControlA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

type AppServiceSourceControlTimeoutsOutputReference interface {
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

// The jsii proxy struct for AppServiceSourceControlTimeoutsOutputReference
type jsiiProxy_AppServiceSourceControlTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSourceControlTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlTimeoutsOutputReference_Override(a AppServiceSourceControlTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		a,
		"resetRead",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

