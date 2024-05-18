// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/machinelearningworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/machine_learning_workspace azurerm_machine_learning_workspace}.
type MachineLearningWorkspace interface {
	cdktf.TerraformResource
	ApplicationInsightsId() *string
	SetApplicationInsightsId(val *string)
	ApplicationInsightsIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerRegistryId() *string
	SetContainerRegistryId(val *string)
	ContainerRegistryIdInput() *string
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
	DiscoveryUrl() *string
	Encryption() MachineLearningWorkspaceEncryptionOutputReference
	EncryptionInput() *MachineLearningWorkspaceEncryption
	FeatureStore() MachineLearningWorkspaceFeatureStoreOutputReference
	FeatureStoreInput() *MachineLearningWorkspaceFeatureStore
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	FriendlyName() *string
	SetFriendlyName(val *string)
	FriendlyNameInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	HighBusinessImpact() interface{}
	SetHighBusinessImpact(val interface{})
	HighBusinessImpactInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() MachineLearningWorkspaceIdentityOutputReference
	IdentityInput() *MachineLearningWorkspaceIdentity
	IdInput() *string
	ImageBuildComputeName() *string
	SetImageBuildComputeName(val *string)
	ImageBuildComputeNameInput() *string
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedNetwork() MachineLearningWorkspaceManagedNetworkOutputReference
	ManagedNetworkInput() *MachineLearningWorkspaceManagedNetwork
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrimaryUserAssignedIdentity() *string
	SetPrimaryUserAssignedIdentity(val *string)
	PrimaryUserAssignedIdentityInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicAccessBehindVirtualNetworkEnabled() interface{}
	SetPublicAccessBehindVirtualNetworkEnabled(val interface{})
	PublicAccessBehindVirtualNetworkEnabledInput() interface{}
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MachineLearningWorkspaceTimeoutsOutputReference
	TimeoutsInput() interface{}
	V1LegacyModeEnabled() interface{}
	SetV1LegacyModeEnabled(val interface{})
	V1LegacyModeEnabledInput() interface{}
	WorkspaceId() *string
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
	PutEncryption(value *MachineLearningWorkspaceEncryption)
	PutFeatureStore(value *MachineLearningWorkspaceFeatureStore)
	PutIdentity(value *MachineLearningWorkspaceIdentity)
	PutManagedNetwork(value *MachineLearningWorkspaceManagedNetwork)
	PutTimeouts(value *MachineLearningWorkspaceTimeouts)
	ResetContainerRegistryId()
	ResetDescription()
	ResetEncryption()
	ResetFeatureStore()
	ResetFriendlyName()
	ResetHighBusinessImpact()
	ResetId()
	ResetImageBuildComputeName()
	ResetKind()
	ResetManagedNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryUserAssignedIdentity()
	ResetPublicAccessBehindVirtualNetworkEnabled()
	ResetPublicNetworkAccessEnabled()
	ResetSkuName()
	ResetTags()
	ResetTimeouts()
	ResetV1LegacyModeEnabled()
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

// The jsii proxy struct for MachineLearningWorkspace
type jsiiProxy_MachineLearningWorkspace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MachineLearningWorkspace) ApplicationInsightsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ApplicationInsightsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ContainerRegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ContainerRegistryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) DiscoveryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Encryption() MachineLearningWorkspaceEncryptionOutputReference {
	var returns MachineLearningWorkspaceEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) EncryptionInput() *MachineLearningWorkspaceEncryption {
	var returns *MachineLearningWorkspaceEncryption
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) FeatureStore() MachineLearningWorkspaceFeatureStoreOutputReference {
	var returns MachineLearningWorkspaceFeatureStoreOutputReference
	_jsii_.Get(
		j,
		"featureStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) FeatureStoreInput() *MachineLearningWorkspaceFeatureStore {
	var returns *MachineLearningWorkspaceFeatureStore
	_jsii_.Get(
		j,
		"featureStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) FriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) HighBusinessImpact() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highBusinessImpact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) HighBusinessImpactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highBusinessImpactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Identity() MachineLearningWorkspaceIdentityOutputReference {
	var returns MachineLearningWorkspaceIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) IdentityInput() *MachineLearningWorkspaceIdentity {
	var returns *MachineLearningWorkspaceIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ImageBuildComputeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageBuildComputeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ImageBuildComputeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageBuildComputeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ManagedNetwork() MachineLearningWorkspaceManagedNetworkOutputReference {
	var returns MachineLearningWorkspaceManagedNetworkOutputReference
	_jsii_.Get(
		j,
		"managedNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ManagedNetworkInput() *MachineLearningWorkspaceManagedNetwork {
	var returns *MachineLearningWorkspaceManagedNetwork
	_jsii_.Get(
		j,
		"managedNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) PrimaryUserAssignedIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryUserAssignedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) PrimaryUserAssignedIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryUserAssignedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) PublicAccessBehindVirtualNetworkEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessBehindVirtualNetworkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) PublicAccessBehindVirtualNetworkEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessBehindVirtualNetworkEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) Timeouts() MachineLearningWorkspaceTimeoutsOutputReference {
	var returns MachineLearningWorkspaceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) V1LegacyModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"v1LegacyModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) V1LegacyModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"v1LegacyModeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningWorkspace) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/machine_learning_workspace azurerm_machine_learning_workspace} Resource.
func NewMachineLearningWorkspace(scope constructs.Construct, id *string, config *MachineLearningWorkspaceConfig) MachineLearningWorkspace {
	_init_.Initialize()

	if err := validateNewMachineLearningWorkspaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MachineLearningWorkspace{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/machine_learning_workspace azurerm_machine_learning_workspace} Resource.
func NewMachineLearningWorkspace_Override(m MachineLearningWorkspace, scope constructs.Construct, id *string, config *MachineLearningWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspace",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetApplicationInsightsId(val *string) {
	if err := j.validateSetApplicationInsightsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationInsightsId",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetContainerRegistryId(val *string) {
	if err := j.validateSetContainerRegistryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryId",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetFriendlyName(val *string) {
	if err := j.validateSetFriendlyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetHighBusinessImpact(val interface{}) {
	if err := j.validateSetHighBusinessImpactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highBusinessImpact",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetImageBuildComputeName(val *string) {
	if err := j.validateSetImageBuildComputeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageBuildComputeName",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetKeyVaultId(val *string) {
	if err := j.validateSetKeyVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetPrimaryUserAssignedIdentity(val *string) {
	if err := j.validateSetPrimaryUserAssignedIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryUserAssignedIdentity",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetPublicAccessBehindVirtualNetworkEnabled(val interface{}) {
	if err := j.validateSetPublicAccessBehindVirtualNetworkEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccessBehindVirtualNetworkEnabled",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetStorageAccountId(val *string) {
	if err := j.validateSetStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MachineLearningWorkspace)SetV1LegacyModeEnabled(val interface{}) {
	if err := j.validateSetV1LegacyModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"v1LegacyModeEnabled",
		val,
	)
}

// Generates CDKTF code for importing a MachineLearningWorkspace resource upon running "cdktf plan <stack-name>".
func MachineLearningWorkspace_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMachineLearningWorkspace_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspace",
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
func MachineLearningWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningWorkspace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MachineLearningWorkspace_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningWorkspace_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspace",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MachineLearningWorkspace_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningWorkspace_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspace",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MachineLearningWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.machineLearningWorkspace.MachineLearningWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MachineLearningWorkspace) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MachineLearningWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MachineLearningWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MachineLearningWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MachineLearningWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MachineLearningWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MachineLearningWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MachineLearningWorkspace) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MachineLearningWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MachineLearningWorkspace) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspace) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MachineLearningWorkspace) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) PutEncryption(value *MachineLearningWorkspaceEncryption) {
	if err := m.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) PutFeatureStore(value *MachineLearningWorkspaceFeatureStore) {
	if err := m.validatePutFeatureStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFeatureStore",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) PutIdentity(value *MachineLearningWorkspaceIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) PutManagedNetwork(value *MachineLearningWorkspaceManagedNetwork) {
	if err := m.validatePutManagedNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putManagedNetwork",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) PutTimeouts(value *MachineLearningWorkspaceTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetContainerRegistryId() {
	_jsii_.InvokeVoid(
		m,
		"resetContainerRegistryId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetFeatureStore() {
	_jsii_.InvokeVoid(
		m,
		"resetFeatureStore",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetFriendlyName() {
	_jsii_.InvokeVoid(
		m,
		"resetFriendlyName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetHighBusinessImpact() {
	_jsii_.InvokeVoid(
		m,
		"resetHighBusinessImpact",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetImageBuildComputeName() {
	_jsii_.InvokeVoid(
		m,
		"resetImageBuildComputeName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetKind() {
	_jsii_.InvokeVoid(
		m,
		"resetKind",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetManagedNetwork() {
	_jsii_.InvokeVoid(
		m,
		"resetManagedNetwork",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetPrimaryUserAssignedIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryUserAssignedIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetPublicAccessBehindVirtualNetworkEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicAccessBehindVirtualNetworkEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetSkuName() {
	_jsii_.InvokeVoid(
		m,
		"resetSkuName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) ResetV1LegacyModeEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetV1LegacyModeEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningWorkspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspace) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspace) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

