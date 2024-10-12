// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/appservicesourcecontrol/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/app_service_source_control azurerm_app_service_source_control}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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

func (j *jsiiProxy_AppServiceSourceControlA) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/app_service_source_control azurerm_app_service_source_control} Resource.
func NewAppServiceSourceControlA(scope constructs.Construct, id *string, config *AppServiceSourceControlAConfig) AppServiceSourceControlA {
	_init_.Initialize()

	if err := validateNewAppServiceSourceControlAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppServiceSourceControlA{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/app_service_source_control azurerm_app_service_source_control} Resource.
func NewAppServiceSourceControlA_Override(a AppServiceSourceControlA, scope constructs.Construct, id *string, config *AppServiceSourceControlAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetRepoUrl(val *string) {
	if err := j.validateSetRepoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repoUrl",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetRollbackEnabled(val interface{}) {
	if err := j.validateSetRollbackEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rollbackEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetUseLocalGit(val interface{}) {
	if err := j.validateSetUseLocalGitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLocalGit",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetUseManualIntegration(val interface{}) {
	if err := j.validateSetUseManualIntegrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useManualIntegration",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlA)SetUseMercurial(val interface{}) {
	if err := j.validateSetUseMercurialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMercurial",
		val,
	)
}

// Generates CDKTF code for importing a AppServiceSourceControlA resource upon running "cdktf plan <stack-name>".
func AppServiceSourceControlA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppServiceSourceControlA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
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
func AppServiceSourceControlA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppServiceSourceControlA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppServiceSourceControlA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppServiceSourceControlA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppServiceSourceControlA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppServiceSourceControlA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.appServiceSourceControl.AppServiceSourceControlA",
		"isTerraformResource",
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

func (a *jsiiProxy_AppServiceSourceControlA) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) PutGithubActionConfiguration(value *AppServiceSourceControlGithubActionConfiguration) {
	if err := a.validatePutGithubActionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithubActionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlA) PutTimeouts(value *AppServiceSourceControlTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (a *jsiiProxy_AppServiceSourceControlA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
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

