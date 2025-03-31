// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcedeploymentscriptazurecli

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/resourcedeploymentscriptazurecli/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.25.0/docs/resources/resource_deployment_script_azure_cli azurerm_resource_deployment_script_azure_cli}.
type ResourceDeploymentScriptAzureCli interface {
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
	Container() ResourceDeploymentScriptAzureCliContainerOutputReference
	ContainerInput() *ResourceDeploymentScriptAzureCliContainer
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnvironmentVariable() ResourceDeploymentScriptAzureCliEnvironmentVariableList
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
	Identity() ResourceDeploymentScriptAzureCliIdentityOutputReference
	IdentityInput() *ResourceDeploymentScriptAzureCliIdentity
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
	StorageAccount() ResourceDeploymentScriptAzureCliStorageAccountOutputReference
	StorageAccountInput() *ResourceDeploymentScriptAzureCliStorageAccount
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
	Timeouts() ResourceDeploymentScriptAzureCliTimeoutsOutputReference
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
	PutContainer(value *ResourceDeploymentScriptAzureCliContainer)
	PutEnvironmentVariable(value interface{})
	PutIdentity(value *ResourceDeploymentScriptAzureCliIdentity)
	PutStorageAccount(value *ResourceDeploymentScriptAzureCliStorageAccount)
	PutTimeouts(value *ResourceDeploymentScriptAzureCliTimeouts)
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

// The jsii proxy struct for ResourceDeploymentScriptAzureCli
type jsiiProxy_ResourceDeploymentScriptAzureCli struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) CleanupPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) CleanupPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) CommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) CommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Container() ResourceDeploymentScriptAzureCliContainerOutputReference {
	var returns ResourceDeploymentScriptAzureCliContainerOutputReference
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ContainerInput() *ResourceDeploymentScriptAzureCliContainer {
	var returns *ResourceDeploymentScriptAzureCliContainer
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) EnvironmentVariable() ResourceDeploymentScriptAzureCliEnvironmentVariableList {
	var returns ResourceDeploymentScriptAzureCliEnvironmentVariableList
	_jsii_.Get(
		j,
		"environmentVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) EnvironmentVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ForceUpdateTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ForceUpdateTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Identity() ResourceDeploymentScriptAzureCliIdentityOutputReference {
	var returns ResourceDeploymentScriptAzureCliIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) IdentityInput() *ResourceDeploymentScriptAzureCliIdentity {
	var returns *ResourceDeploymentScriptAzureCliIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Outputs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) PrimaryScriptUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryScriptUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) PrimaryScriptUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryScriptUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) RetentionInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) RetentionIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ScriptContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) ScriptContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) StorageAccount() ResourceDeploymentScriptAzureCliStorageAccountOutputReference {
	var returns ResourceDeploymentScriptAzureCliStorageAccountOutputReference
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) StorageAccountInput() *ResourceDeploymentScriptAzureCliStorageAccount {
	var returns *ResourceDeploymentScriptAzureCliStorageAccount
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) SupportingScriptUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportingScriptUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) SupportingScriptUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportingScriptUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Timeouts() ResourceDeploymentScriptAzureCliTimeoutsOutputReference {
	var returns ResourceDeploymentScriptAzureCliTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.25.0/docs/resources/resource_deployment_script_azure_cli azurerm_resource_deployment_script_azure_cli} Resource.
func NewResourceDeploymentScriptAzureCli(scope constructs.Construct, id *string, config *ResourceDeploymentScriptAzureCliConfig) ResourceDeploymentScriptAzureCli {
	_init_.Initialize()

	if err := validateNewResourceDeploymentScriptAzureCliParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceDeploymentScriptAzureCli{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzureCli.ResourceDeploymentScriptAzureCli",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.25.0/docs/resources/resource_deployment_script_azure_cli azurerm_resource_deployment_script_azure_cli} Resource.
func NewResourceDeploymentScriptAzureCli_Override(r ResourceDeploymentScriptAzureCli, scope constructs.Construct, id *string, config *ResourceDeploymentScriptAzureCliConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzureCli.ResourceDeploymentScriptAzureCli",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetCleanupPreference(val *string) {
	if err := j.validateSetCleanupPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanupPreference",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetCommandLine(val *string) {
	if err := j.validateSetCommandLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commandLine",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetForceUpdateTag(val *string) {
	if err := j.validateSetForceUpdateTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateTag",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetPrimaryScriptUri(val *string) {
	if err := j.validateSetPrimaryScriptUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryScriptUri",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetRetentionInterval(val *string) {
	if err := j.validateSetRetentionIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInterval",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetScriptContent(val *string) {
	if err := j.validateSetScriptContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptContent",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetSupportingScriptUris(val *[]*string) {
	if err := j.validateSetSupportingScriptUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportingScriptUris",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCli)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a ResourceDeploymentScriptAzureCli resource upon running "cdktf plan <stack-name>".
func ResourceDeploymentScriptAzureCli_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateResourceDeploymentScriptAzureCli_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzureCli.ResourceDeploymentScriptAzureCli",
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
func ResourceDeploymentScriptAzureCli_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceDeploymentScriptAzureCli_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzureCli.ResourceDeploymentScriptAzureCli",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceDeploymentScriptAzureCli_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceDeploymentScriptAzureCli_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzureCli.ResourceDeploymentScriptAzureCli",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceDeploymentScriptAzureCli_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceDeploymentScriptAzureCli_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzureCli.ResourceDeploymentScriptAzureCli",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResourceDeploymentScriptAzureCli_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.resourceDeploymentScriptAzureCli.ResourceDeploymentScriptAzureCli",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) PutContainer(value *ResourceDeploymentScriptAzureCliContainer) {
	if err := r.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putContainer",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) PutEnvironmentVariable(value interface{}) {
	if err := r.validatePutEnvironmentVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEnvironmentVariable",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) PutIdentity(value *ResourceDeploymentScriptAzureCliIdentity) {
	if err := r.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIdentity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) PutStorageAccount(value *ResourceDeploymentScriptAzureCliStorageAccount) {
	if err := r.validatePutStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) PutTimeouts(value *ResourceDeploymentScriptAzureCliTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetCleanupPreference() {
	_jsii_.InvokeVoid(
		r,
		"resetCleanupPreference",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetCommandLine() {
	_jsii_.InvokeVoid(
		r,
		"resetCommandLine",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetContainer() {
	_jsii_.InvokeVoid(
		r,
		"resetContainer",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetEnvironmentVariable() {
	_jsii_.InvokeVoid(
		r,
		"resetEnvironmentVariable",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetForceUpdateTag() {
	_jsii_.InvokeVoid(
		r,
		"resetForceUpdateTag",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetIdentity() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetPrimaryScriptUri() {
	_jsii_.InvokeVoid(
		r,
		"resetPrimaryScriptUri",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetScriptContent() {
	_jsii_.InvokeVoid(
		r,
		"resetScriptContent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetSupportingScriptUris() {
	_jsii_.InvokeVoid(
		r,
		"resetSupportingScriptUris",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCli) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

