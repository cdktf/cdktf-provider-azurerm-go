// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/synapseworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/synapse_workspace azurerm_synapse_workspace}.
type SynapseWorkspace interface {
	cdktf.TerraformResource
	AadAdmin() SynapseWorkspaceAadAdminList
	AadAdminInput() interface{}
	AzureDevopsRepo() SynapseWorkspaceAzureDevopsRepoOutputReference
	AzureDevopsRepoInput() *SynapseWorkspaceAzureDevopsRepo
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputeSubnetId() *string
	SetComputeSubnetId(val *string)
	ComputeSubnetIdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectivityEndpoints() cdktf.StringMap
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomerManagedKey() SynapseWorkspaceCustomerManagedKeyOutputReference
	CustomerManagedKeyInput() *SynapseWorkspaceCustomerManagedKey
	DataExfiltrationProtectionEnabled() interface{}
	SetDataExfiltrationProtectionEnabled(val interface{})
	DataExfiltrationProtectionEnabledInput() interface{}
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
	GithubRepo() SynapseWorkspaceGithubRepoOutputReference
	GithubRepoInput() *SynapseWorkspaceGithubRepo
	Id() *string
	SetId(val *string)
	Identity() SynapseWorkspaceIdentityOutputReference
	IdentityInput() *SynapseWorkspaceIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkingAllowedForAadTenantIds() *[]*string
	SetLinkingAllowedForAadTenantIds(val *[]*string)
	LinkingAllowedForAadTenantIdsInput() *[]*string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedResourceGroupName() *string
	SetManagedResourceGroupName(val *string)
	ManagedResourceGroupNameInput() *string
	ManagedVirtualNetworkEnabled() interface{}
	SetManagedVirtualNetworkEnabled(val interface{})
	ManagedVirtualNetworkEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	PurviewId() *string
	SetPurviewId(val *string)
	PurviewIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SqlAadAdmin() SynapseWorkspaceSqlAadAdminList
	SqlAadAdminInput() interface{}
	SqlAdministratorLogin() *string
	SetSqlAdministratorLogin(val *string)
	SqlAdministratorLoginInput() *string
	SqlAdministratorLoginPassword() *string
	SetSqlAdministratorLoginPassword(val *string)
	SqlAdministratorLoginPasswordInput() *string
	SqlIdentityControlEnabled() interface{}
	SetSqlIdentityControlEnabled(val interface{})
	SqlIdentityControlEnabledInput() interface{}
	StorageDataLakeGen2FilesystemId() *string
	SetStorageDataLakeGen2FilesystemId(val *string)
	StorageDataLakeGen2FilesystemIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SynapseWorkspaceTimeoutsOutputReference
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
	PutAadAdmin(value interface{})
	PutAzureDevopsRepo(value *SynapseWorkspaceAzureDevopsRepo)
	PutCustomerManagedKey(value *SynapseWorkspaceCustomerManagedKey)
	PutGithubRepo(value *SynapseWorkspaceGithubRepo)
	PutIdentity(value *SynapseWorkspaceIdentity)
	PutSqlAadAdmin(value interface{})
	PutTimeouts(value *SynapseWorkspaceTimeouts)
	ResetAadAdmin()
	ResetAzureDevopsRepo()
	ResetComputeSubnetId()
	ResetCustomerManagedKey()
	ResetDataExfiltrationProtectionEnabled()
	ResetGithubRepo()
	ResetId()
	ResetIdentity()
	ResetLinkingAllowedForAadTenantIds()
	ResetManagedResourceGroupName()
	ResetManagedVirtualNetworkEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetPurviewId()
	ResetSqlAadAdmin()
	ResetSqlAdministratorLogin()
	ResetSqlAdministratorLoginPassword()
	ResetSqlIdentityControlEnabled()
	ResetTags()
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

// The jsii proxy struct for SynapseWorkspace
type jsiiProxy_SynapseWorkspace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SynapseWorkspace) AadAdmin() SynapseWorkspaceAadAdminList {
	var returns SynapseWorkspaceAadAdminList
	_jsii_.Get(
		j,
		"aadAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) AadAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aadAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) AzureDevopsRepo() SynapseWorkspaceAzureDevopsRepoOutputReference {
	var returns SynapseWorkspaceAzureDevopsRepoOutputReference
	_jsii_.Get(
		j,
		"azureDevopsRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) AzureDevopsRepoInput() *SynapseWorkspaceAzureDevopsRepo {
	var returns *SynapseWorkspaceAzureDevopsRepo
	_jsii_.Get(
		j,
		"azureDevopsRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ComputeSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ComputeSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ConnectivityEndpoints() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"connectivityEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) CustomerManagedKey() SynapseWorkspaceCustomerManagedKeyOutputReference {
	var returns SynapseWorkspaceCustomerManagedKeyOutputReference
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) CustomerManagedKeyInput() *SynapseWorkspaceCustomerManagedKey {
	var returns *SynapseWorkspaceCustomerManagedKey
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) DataExfiltrationProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataExfiltrationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) DataExfiltrationProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataExfiltrationProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) GithubRepo() SynapseWorkspaceGithubRepoOutputReference {
	var returns SynapseWorkspaceGithubRepoOutputReference
	_jsii_.Get(
		j,
		"githubRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) GithubRepoInput() *SynapseWorkspaceGithubRepo {
	var returns *SynapseWorkspaceGithubRepo
	_jsii_.Get(
		j,
		"githubRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Identity() SynapseWorkspaceIdentityOutputReference {
	var returns SynapseWorkspaceIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) IdentityInput() *SynapseWorkspaceIdentity {
	var returns *SynapseWorkspaceIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) LinkingAllowedForAadTenantIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linkingAllowedForAadTenantIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) LinkingAllowedForAadTenantIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linkingAllowedForAadTenantIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ManagedResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ManagedResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ManagedVirtualNetworkEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedVirtualNetworkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ManagedVirtualNetworkEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedVirtualNetworkEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) PurviewId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purviewId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) PurviewIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purviewIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) SqlAadAdmin() SynapseWorkspaceSqlAadAdminList {
	var returns SynapseWorkspaceSqlAadAdminList
	_jsii_.Get(
		j,
		"sqlAadAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) SqlAadAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlAadAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) SqlAdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlAdministratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) SqlAdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlAdministratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) SqlAdministratorLoginPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlAdministratorLoginPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) SqlAdministratorLoginPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlAdministratorLoginPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) SqlIdentityControlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlIdentityControlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) SqlIdentityControlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlIdentityControlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) StorageDataLakeGen2FilesystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageDataLakeGen2FilesystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) StorageDataLakeGen2FilesystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageDataLakeGen2FilesystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) Timeouts() SynapseWorkspaceTimeoutsOutputReference {
	var returns SynapseWorkspaceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspace) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/synapse_workspace azurerm_synapse_workspace} Resource.
func NewSynapseWorkspace(scope constructs.Construct, id *string, config *SynapseWorkspaceConfig) SynapseWorkspace {
	_init_.Initialize()

	if err := validateNewSynapseWorkspaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SynapseWorkspace{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/synapse_workspace azurerm_synapse_workspace} Resource.
func NewSynapseWorkspace_Override(s SynapseWorkspace, scope constructs.Construct, id *string, config *SynapseWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetComputeSubnetId(val *string) {
	if err := j.validateSetComputeSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeSubnetId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetDataExfiltrationProtectionEnabled(val interface{}) {
	if err := j.validateSetDataExfiltrationProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataExfiltrationProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetLinkingAllowedForAadTenantIds(val *[]*string) {
	if err := j.validateSetLinkingAllowedForAadTenantIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkingAllowedForAadTenantIds",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetManagedResourceGroupName(val *string) {
	if err := j.validateSetManagedResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetManagedVirtualNetworkEnabled(val interface{}) {
	if err := j.validateSetManagedVirtualNetworkEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedVirtualNetworkEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetPurviewId(val *string) {
	if err := j.validateSetPurviewIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purviewId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetSqlAdministratorLogin(val *string) {
	if err := j.validateSetSqlAdministratorLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlAdministratorLogin",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetSqlAdministratorLoginPassword(val *string) {
	if err := j.validateSetSqlAdministratorLoginPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlAdministratorLoginPassword",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetSqlIdentityControlEnabled(val interface{}) {
	if err := j.validateSetSqlIdentityControlEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlIdentityControlEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetStorageDataLakeGen2FilesystemId(val *string) {
	if err := j.validateSetStorageDataLakeGen2FilesystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageDataLakeGen2FilesystemId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func SynapseWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseWorkspace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SynapseWorkspace_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseWorkspace_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SynapseWorkspace_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseWorkspace_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SynapseWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SynapseWorkspace) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutAadAdmin(value interface{}) {
	if err := s.validatePutAadAdminParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAadAdmin",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutAzureDevopsRepo(value *SynapseWorkspaceAzureDevopsRepo) {
	if err := s.validatePutAzureDevopsRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureDevopsRepo",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutCustomerManagedKey(value *SynapseWorkspaceCustomerManagedKey) {
	if err := s.validatePutCustomerManagedKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomerManagedKey",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutGithubRepo(value *SynapseWorkspaceGithubRepo) {
	if err := s.validatePutGithubRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGithubRepo",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutIdentity(value *SynapseWorkspaceIdentity) {
	if err := s.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdentity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutSqlAadAdmin(value interface{}) {
	if err := s.validatePutSqlAadAdminParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSqlAadAdmin",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutTimeouts(value *SynapseWorkspaceTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetAadAdmin() {
	_jsii_.InvokeVoid(
		s,
		"resetAadAdmin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetAzureDevopsRepo() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureDevopsRepo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetComputeSubnetId() {
	_jsii_.InvokeVoid(
		s,
		"resetComputeSubnetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetDataExfiltrationProtectionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDataExfiltrationProtectionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetGithubRepo() {
	_jsii_.InvokeVoid(
		s,
		"resetGithubRepo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetIdentity() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetLinkingAllowedForAadTenantIds() {
	_jsii_.InvokeVoid(
		s,
		"resetLinkingAllowedForAadTenantIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetManagedResourceGroupName() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedResourceGroupName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetManagedVirtualNetworkEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedVirtualNetworkEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetPurviewId() {
	_jsii_.InvokeVoid(
		s,
		"resetPurviewId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetSqlAadAdmin() {
	_jsii_.InvokeVoid(
		s,
		"resetSqlAadAdmin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetSqlAdministratorLogin() {
	_jsii_.InvokeVoid(
		s,
		"resetSqlAdministratorLogin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetSqlAdministratorLoginPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetSqlAdministratorLoginPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetSqlIdentityControlEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSqlIdentityControlEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

