// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcedeploymentscriptazurepowershell

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/resourcedeploymentscriptazurepowershell/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/resource_deployment_script_azure_power_shell azurerm_resource_deployment_script_azure_power_shell}.
type ResourceDeploymentScriptAzurePowerShell interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CleanupPreference() *string
	SetCleanupPreference(val *string)
	CleanupPreferenceInput() *string
	CommandLine() *string
	SetCommandLine(val *string)
	CommandLineInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Container() ResourceDeploymentScriptAzurePowerShellContainerOutputReference
	ContainerInput() *ResourceDeploymentScriptAzurePowerShellContainer
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnvironmentVariable() ResourceDeploymentScriptAzurePowerShellEnvironmentVariableList
	EnvironmentVariableInput() interface{}
	ForceUpdateTag() *string
	SetForceUpdateTag(val *string)
	ForceUpdateTagInput() *string
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
	Identity() ResourceDeploymentScriptAzurePowerShellIdentityOutputReference
	IdentityInput() *ResourceDeploymentScriptAzurePowerShellIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Outputs() *string
	PrimaryScriptUri() *string
	SetPrimaryScriptUri(val *string)
	PrimaryScriptUriInput() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RetentionInterval() *string
	SetRetentionInterval(val *string)
	RetentionIntervalInput() *string
	ScriptContent() *string
	SetScriptContent(val *string)
	ScriptContentInput() *string
	StorageAccount() ResourceDeploymentScriptAzurePowerShellStorageAccountOutputReference
	StorageAccountInput() *ResourceDeploymentScriptAzurePowerShellStorageAccount
	SupportingScriptUris() *[]*string
	SetSupportingScriptUris(val *[]*string)
	SupportingScriptUrisInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	Timeouts() ResourceDeploymentScriptAzurePowerShellTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutContainer(value *ResourceDeploymentScriptAzurePowerShellContainer)
	PutEnvironmentVariable(value interface{})
	PutIdentity(value *ResourceDeploymentScriptAzurePowerShellIdentity)
	PutStorageAccount(value *ResourceDeploymentScriptAzurePowerShellStorageAccount)
	PutTimeouts(value *ResourceDeploymentScriptAzurePowerShellTimeouts)
	ResetCleanupPreference()
	ResetCommandLine()
	ResetContainer()
	ResetEnvironmentVariable()
	ResetForceUpdateTag()
	ResetId()
	ResetIdentity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryScriptUri()
	ResetScriptContent()
	ResetStorageAccount()
	ResetSupportingScriptUris()
	ResetTags()
	ResetTimeout()
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

// The jsii proxy struct for ResourceDeploymentScriptAzurePowerShell
type jsiiProxy_ResourceDeploymentScriptAzurePowerShell struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) CleanupPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) CleanupPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) CommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) CommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Container() ResourceDeploymentScriptAzurePowerShellContainerOutputReference {
	var returns ResourceDeploymentScriptAzurePowerShellContainerOutputReference
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ContainerInput() *ResourceDeploymentScriptAzurePowerShellContainer {
	var returns *ResourceDeploymentScriptAzurePowerShellContainer
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) EnvironmentVariable() ResourceDeploymentScriptAzurePowerShellEnvironmentVariableList {
	var returns ResourceDeploymentScriptAzurePowerShellEnvironmentVariableList
	_jsii_.Get(
		j,
		"environmentVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) EnvironmentVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ForceUpdateTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ForceUpdateTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Identity() ResourceDeploymentScriptAzurePowerShellIdentityOutputReference {
	var returns ResourceDeploymentScriptAzurePowerShellIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) IdentityInput() *ResourceDeploymentScriptAzurePowerShellIdentity {
	var returns *ResourceDeploymentScriptAzurePowerShellIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Outputs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) PrimaryScriptUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryScriptUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) PrimaryScriptUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryScriptUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) RetentionInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) RetentionIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ScriptContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ScriptContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) StorageAccount() ResourceDeploymentScriptAzurePowerShellStorageAccountOutputReference {
	var returns ResourceDeploymentScriptAzurePowerShellStorageAccountOutputReference
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) StorageAccountInput() *ResourceDeploymentScriptAzurePowerShellStorageAccount {
	var returns *ResourceDeploymentScriptAzurePowerShellStorageAccount
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) SupportingScriptUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportingScriptUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) SupportingScriptUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportingScriptUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Timeouts() ResourceDeploymentScriptAzurePowerShellTimeoutsOutputReference {
	var returns ResourceDeploymentScriptAzurePowerShellTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/resource_deployment_script_azure_power_shell azurerm_resource_deployment_script_azure_power_shell} Resource.
func NewResourceDeploymentScriptAzurePowerShell(scope constructs.Construct, id *string, config *ResourceDeploymentScriptAzurePowerShellConfig) ResourceDeploymentScriptAzurePowerShell {
	_init_.Initialize()

	if err := validateNewResourceDeploymentScriptAzurePowerShellParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceDeploymentScriptAzurePowerShell{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzurePowerShell.ResourceDeploymentScriptAzurePowerShell",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/resource_deployment_script_azure_power_shell azurerm_resource_deployment_script_azure_power_shell} Resource.
func NewResourceDeploymentScriptAzurePowerShell_Override(r ResourceDeploymentScriptAzurePowerShell, scope constructs.Construct, id *string, config *ResourceDeploymentScriptAzurePowerShellConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzurePowerShell.ResourceDeploymentScriptAzurePowerShell",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetCleanupPreference(val *string) {
	if err := j.validateSetCleanupPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanupPreference",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetCommandLine(val *string) {
	if err := j.validateSetCommandLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commandLine",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetForceUpdateTag(val *string) {
	if err := j.validateSetForceUpdateTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateTag",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetPrimaryScriptUri(val *string) {
	if err := j.validateSetPrimaryScriptUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryScriptUri",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetRetentionInterval(val *string) {
	if err := j.validateSetRetentionIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInterval",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetScriptContent(val *string) {
	if err := j.validateSetScriptContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptContent",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetSupportingScriptUris(val *[]*string) {
	if err := j.validateSetSupportingScriptUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportingScriptUris",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzurePowerShell)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a ResourceDeploymentScriptAzurePowerShell resource upon running "cdktf plan <stack-name>".
func ResourceDeploymentScriptAzurePowerShell_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateResourceDeploymentScriptAzurePowerShell_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzurePowerShell.ResourceDeploymentScriptAzurePowerShell",
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
func ResourceDeploymentScriptAzurePowerShell_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceDeploymentScriptAzurePowerShell_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzurePowerShell.ResourceDeploymentScriptAzurePowerShell",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceDeploymentScriptAzurePowerShell_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceDeploymentScriptAzurePowerShell_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzurePowerShell.ResourceDeploymentScriptAzurePowerShell",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceDeploymentScriptAzurePowerShell_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceDeploymentScriptAzurePowerShell_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzurePowerShell.ResourceDeploymentScriptAzurePowerShell",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResourceDeploymentScriptAzurePowerShell_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzurePowerShell.ResourceDeploymentScriptAzurePowerShell",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) PutContainer(value *ResourceDeploymentScriptAzurePowerShellContainer) {
	if err := r.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putContainer",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) PutEnvironmentVariable(value interface{}) {
	if err := r.validatePutEnvironmentVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEnvironmentVariable",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) PutIdentity(value *ResourceDeploymentScriptAzurePowerShellIdentity) {
	if err := r.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIdentity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) PutStorageAccount(value *ResourceDeploymentScriptAzurePowerShellStorageAccount) {
	if err := r.validatePutStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) PutTimeouts(value *ResourceDeploymentScriptAzurePowerShellTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetCleanupPreference() {
	_jsii_.InvokeVoid(
		r,
		"resetCleanupPreference",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetCommandLine() {
	_jsii_.InvokeVoid(
		r,
		"resetCommandLine",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetContainer() {
	_jsii_.InvokeVoid(
		r,
		"resetContainer",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetEnvironmentVariable() {
	_jsii_.InvokeVoid(
		r,
		"resetEnvironmentVariable",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetForceUpdateTag() {
	_jsii_.InvokeVoid(
		r,
		"resetForceUpdateTag",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetIdentity() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetPrimaryScriptUri() {
	_jsii_.InvokeVoid(
		r,
		"resetPrimaryScriptUri",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetScriptContent() {
	_jsii_.InvokeVoid(
		r,
		"resetScriptContent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetSupportingScriptUris() {
	_jsii_.InvokeVoid(
		r,
		"resetSupportingScriptUris",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzurePowerShell) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

