package synapseworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/synapseworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace azurerm_synapse_workspace}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	ResetLinkingAllowedForAadTenantIds()
	ResetManagedResourceGroupName()
	ResetManagedVirtualNetworkEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetPurviewId()
	ResetSqlAadAdmin()
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

func (j *jsiiProxy_SynapseWorkspace) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace azurerm_synapse_workspace} Resource.
func NewSynapseWorkspace(scope constructs.Construct, id *string, config *SynapseWorkspaceConfig) SynapseWorkspace {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspace{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace azurerm_synapse_workspace} Resource.
func NewSynapseWorkspace_Override(s SynapseWorkspace, scope constructs.Construct, id *string, config *SynapseWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetComputeSubnetId(val *string) {
	_jsii_.Set(
		j,
		"computeSubnetId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetDataExfiltrationProtectionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dataExfiltrationProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetLinkingAllowedForAadTenantIds(val *[]*string) {
	_jsii_.Set(
		j,
		"linkingAllowedForAadTenantIds",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetManagedResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"managedResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetManagedVirtualNetworkEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"managedVirtualNetworkEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetPurviewId(val *string) {
	_jsii_.Set(
		j,
		"purviewId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetSqlAdministratorLogin(val *string) {
	_jsii_.Set(
		j,
		"sqlAdministratorLogin",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetSqlAdministratorLoginPassword(val *string) {
	_jsii_.Set(
		j,
		"sqlAdministratorLoginPassword",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetSqlIdentityControlEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"sqlIdentityControlEnabled",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetStorageDataLakeGen2FilesystemId(val *string) {
	_jsii_.Set(
		j,
		"storageDataLakeGen2FilesystemId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspace) SetTags(val *map[string]*string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspace",
		"isConstruct",
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
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutAadAdmin(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putAadAdmin",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutAzureDevopsRepo(value *SynapseWorkspaceAzureDevopsRepo) {
	_jsii_.InvokeVoid(
		s,
		"putAzureDevopsRepo",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutCustomerManagedKey(value *SynapseWorkspaceCustomerManagedKey) {
	_jsii_.InvokeVoid(
		s,
		"putCustomerManagedKey",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutGithubRepo(value *SynapseWorkspaceGithubRepo) {
	_jsii_.InvokeVoid(
		s,
		"putGithubRepo",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutIdentity(value *SynapseWorkspaceIdentity) {
	_jsii_.InvokeVoid(
		s,
		"putIdentity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutSqlAadAdmin(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putSqlAadAdmin",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseWorkspace) PutTimeouts(value *SynapseWorkspaceTimeouts) {
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

type SynapseWorkspaceAadAdmin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#login SynapseWorkspace#login}.
	Login *string `field:"optional" json:"login" yaml:"login"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#object_id SynapseWorkspace#object_id}.
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#tenant_id SynapseWorkspace#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

type SynapseWorkspaceAadAdminList interface {
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
	Get(index *float64) SynapseWorkspaceAadAdminOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SynapseWorkspaceAadAdminList
type jsiiProxy_SynapseWorkspaceAadAdminList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceAadAdminList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SynapseWorkspaceAadAdminList {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceAadAdminList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceAadAdminList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceAadAdminList_Override(s SynapseWorkspaceAadAdminList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceAadAdminList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminList) Get(index *float64) SynapseWorkspaceAadAdminOutputReference {
	var returns SynapseWorkspaceAadAdminOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SynapseWorkspaceAadAdminOutputReference interface {
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
	Login() *string
	SetLogin(val *string)
	LoginInput() *string
	ObjectId() *string
	SetObjectId(val *string)
	ObjectIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetLogin()
	ResetObjectId()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SynapseWorkspaceAadAdminOutputReference
type jsiiProxy_SynapseWorkspaceAadAdminOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) LoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceAadAdminOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SynapseWorkspaceAadAdminOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceAadAdminOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceAadAdminOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceAadAdminOutputReference_Override(s SynapseWorkspaceAadAdminOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceAadAdminOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) SetLogin(val *string) {
	_jsii_.Set(
		j,
		"login",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) SetObjectId(val *string) {
	_jsii_.Set(
		j,
		"objectId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ResetLogin() {
	_jsii_.InvokeVoid(
		s,
		"resetLogin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ResetObjectId() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		s,
		"resetTenantId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAadAdminOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SynapseWorkspaceAzureDevopsRepo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#account_name SynapseWorkspace#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#branch_name SynapseWorkspace#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#project_name SynapseWorkspace#project_name}.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#repository_name SynapseWorkspace#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#root_folder SynapseWorkspace#root_folder}.
	RootFolder *string `field:"required" json:"rootFolder" yaml:"rootFolder"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#last_commit_id SynapseWorkspace#last_commit_id}.
	LastCommitId *string `field:"optional" json:"lastCommitId" yaml:"lastCommitId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#tenant_id SynapseWorkspace#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

type SynapseWorkspaceAzureDevopsRepoOutputReference interface {
	cdktf.ComplexObject
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	BranchName() *string
	SetBranchName(val *string)
	BranchNameInput() *string
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
	InternalValue() *SynapseWorkspaceAzureDevopsRepo
	SetInternalValue(val *SynapseWorkspaceAzureDevopsRepo)
	LastCommitId() *string
	SetLastCommitId(val *string)
	LastCommitIdInput() *string
	ProjectName() *string
	SetProjectName(val *string)
	ProjectNameInput() *string
	RepositoryName() *string
	SetRepositoryName(val *string)
	RepositoryNameInput() *string
	RootFolder() *string
	SetRootFolder(val *string)
	RootFolderInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetLastCommitId()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SynapseWorkspaceAzureDevopsRepoOutputReference
type jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) BranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) InternalValue() *SynapseWorkspaceAzureDevopsRepo {
	var returns *SynapseWorkspaceAzureDevopsRepo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) LastCommitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) LastCommitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) ProjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) RepositoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) RootFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) RootFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceAzureDevopsRepoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SynapseWorkspaceAzureDevopsRepoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceAzureDevopsRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceAzureDevopsRepoOutputReference_Override(s SynapseWorkspaceAzureDevopsRepoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceAzureDevopsRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetBranchName(val *string) {
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetInternalValue(val *SynapseWorkspaceAzureDevopsRepo) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetLastCommitId(val *string) {
	_jsii_.Set(
		j,
		"lastCommitId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetProjectName(val *string) {
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetRepositoryName(val *string) {
	_jsii_.Set(
		j,
		"repositoryName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetRootFolder(val *string) {
	_jsii_.Set(
		j,
		"rootFolder",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) ResetLastCommitId() {
	_jsii_.InvokeVoid(
		s,
		"resetLastCommitId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		s,
		"resetTenantId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceAzureDevopsRepoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SynapseWorkspaceConfig struct {
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
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#identity SynapseWorkspace#identity}
	Identity *SynapseWorkspaceIdentity `field:"required" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#location SynapseWorkspace#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#name SynapseWorkspace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#resource_group_name SynapseWorkspace#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#sql_administrator_login SynapseWorkspace#sql_administrator_login}.
	SqlAdministratorLogin *string `field:"required" json:"sqlAdministratorLogin" yaml:"sqlAdministratorLogin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#sql_administrator_login_password SynapseWorkspace#sql_administrator_login_password}.
	SqlAdministratorLoginPassword *string `field:"required" json:"sqlAdministratorLoginPassword" yaml:"sqlAdministratorLoginPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#storage_data_lake_gen2_filesystem_id SynapseWorkspace#storage_data_lake_gen2_filesystem_id}.
	StorageDataLakeGen2FilesystemId *string `field:"required" json:"storageDataLakeGen2FilesystemId" yaml:"storageDataLakeGen2FilesystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#aad_admin SynapseWorkspace#aad_admin}.
	AadAdmin interface{} `field:"optional" json:"aadAdmin" yaml:"aadAdmin"`
	// azure_devops_repo block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#azure_devops_repo SynapseWorkspace#azure_devops_repo}
	AzureDevopsRepo *SynapseWorkspaceAzureDevopsRepo `field:"optional" json:"azureDevopsRepo" yaml:"azureDevopsRepo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#compute_subnet_id SynapseWorkspace#compute_subnet_id}.
	ComputeSubnetId *string `field:"optional" json:"computeSubnetId" yaml:"computeSubnetId"`
	// customer_managed_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#customer_managed_key SynapseWorkspace#customer_managed_key}
	CustomerManagedKey *SynapseWorkspaceCustomerManagedKey `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#data_exfiltration_protection_enabled SynapseWorkspace#data_exfiltration_protection_enabled}.
	DataExfiltrationProtectionEnabled interface{} `field:"optional" json:"dataExfiltrationProtectionEnabled" yaml:"dataExfiltrationProtectionEnabled"`
	// github_repo block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#github_repo SynapseWorkspace#github_repo}
	GithubRepo *SynapseWorkspaceGithubRepo `field:"optional" json:"githubRepo" yaml:"githubRepo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#id SynapseWorkspace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#linking_allowed_for_aad_tenant_ids SynapseWorkspace#linking_allowed_for_aad_tenant_ids}.
	LinkingAllowedForAadTenantIds *[]*string `field:"optional" json:"linkingAllowedForAadTenantIds" yaml:"linkingAllowedForAadTenantIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#managed_resource_group_name SynapseWorkspace#managed_resource_group_name}.
	ManagedResourceGroupName *string `field:"optional" json:"managedResourceGroupName" yaml:"managedResourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#managed_virtual_network_enabled SynapseWorkspace#managed_virtual_network_enabled}.
	ManagedVirtualNetworkEnabled interface{} `field:"optional" json:"managedVirtualNetworkEnabled" yaml:"managedVirtualNetworkEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#public_network_access_enabled SynapseWorkspace#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#purview_id SynapseWorkspace#purview_id}.
	PurviewId *string `field:"optional" json:"purviewId" yaml:"purviewId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#sql_aad_admin SynapseWorkspace#sql_aad_admin}.
	SqlAadAdmin interface{} `field:"optional" json:"sqlAadAdmin" yaml:"sqlAadAdmin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#sql_identity_control_enabled SynapseWorkspace#sql_identity_control_enabled}.
	SqlIdentityControlEnabled interface{} `field:"optional" json:"sqlIdentityControlEnabled" yaml:"sqlIdentityControlEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#tags SynapseWorkspace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#timeouts SynapseWorkspace#timeouts}
	Timeouts *SynapseWorkspaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type SynapseWorkspaceCustomerManagedKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#key_versionless_id SynapseWorkspace#key_versionless_id}.
	KeyVersionlessId *string `field:"required" json:"keyVersionlessId" yaml:"keyVersionlessId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#key_name SynapseWorkspace#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
}

type SynapseWorkspaceCustomerManagedKeyOutputReference interface {
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
	InternalValue() *SynapseWorkspaceCustomerManagedKey
	SetInternalValue(val *SynapseWorkspaceCustomerManagedKey)
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	KeyVersionlessId() *string
	SetKeyVersionlessId(val *string)
	KeyVersionlessIdInput() *string
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
	ResetKeyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SynapseWorkspaceCustomerManagedKeyOutputReference
type jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) InternalValue() *SynapseWorkspaceCustomerManagedKey {
	var returns *SynapseWorkspaceCustomerManagedKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) KeyVersionlessId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVersionlessId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) KeyVersionlessIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVersionlessIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceCustomerManagedKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SynapseWorkspaceCustomerManagedKeyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceCustomerManagedKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceCustomerManagedKeyOutputReference_Override(s SynapseWorkspaceCustomerManagedKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceCustomerManagedKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) SetInternalValue(val *SynapseWorkspaceCustomerManagedKey) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) SetKeyName(val *string) {
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) SetKeyVersionlessId(val *string) {
	_jsii_.Set(
		j,
		"keyVersionlessId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceCustomerManagedKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SynapseWorkspaceGithubRepo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#account_name SynapseWorkspace#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#branch_name SynapseWorkspace#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#repository_name SynapseWorkspace#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#root_folder SynapseWorkspace#root_folder}.
	RootFolder *string `field:"required" json:"rootFolder" yaml:"rootFolder"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#git_url SynapseWorkspace#git_url}.
	GitUrl *string `field:"optional" json:"gitUrl" yaml:"gitUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#last_commit_id SynapseWorkspace#last_commit_id}.
	LastCommitId *string `field:"optional" json:"lastCommitId" yaml:"lastCommitId"`
}

type SynapseWorkspaceGithubRepoOutputReference interface {
	cdktf.ComplexObject
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	BranchName() *string
	SetBranchName(val *string)
	BranchNameInput() *string
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
	GitUrl() *string
	SetGitUrl(val *string)
	GitUrlInput() *string
	InternalValue() *SynapseWorkspaceGithubRepo
	SetInternalValue(val *SynapseWorkspaceGithubRepo)
	LastCommitId() *string
	SetLastCommitId(val *string)
	LastCommitIdInput() *string
	RepositoryName() *string
	SetRepositoryName(val *string)
	RepositoryNameInput() *string
	RootFolder() *string
	SetRootFolder(val *string)
	RootFolderInput() *string
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
	ResetGitUrl()
	ResetLastCommitId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SynapseWorkspaceGithubRepoOutputReference
type jsiiProxy_SynapseWorkspaceGithubRepoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) BranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GitUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GitUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) InternalValue() *SynapseWorkspaceGithubRepo {
	var returns *SynapseWorkspaceGithubRepo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) LastCommitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) LastCommitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) RepositoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) RootFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) RootFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceGithubRepoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SynapseWorkspaceGithubRepoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceGithubRepoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceGithubRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceGithubRepoOutputReference_Override(s SynapseWorkspaceGithubRepoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceGithubRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetBranchName(val *string) {
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetGitUrl(val *string) {
	_jsii_.Set(
		j,
		"gitUrl",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetInternalValue(val *SynapseWorkspaceGithubRepo) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetLastCommitId(val *string) {
	_jsii_.Set(
		j,
		"lastCommitId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetRepositoryName(val *string) {
	_jsii_.Set(
		j,
		"repositoryName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetRootFolder(val *string) {
	_jsii_.Set(
		j,
		"rootFolder",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ResetGitUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetGitUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ResetLastCommitId() {
	_jsii_.InvokeVoid(
		s,
		"resetLastCommitId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SynapseWorkspaceIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#type SynapseWorkspace#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

type SynapseWorkspaceIdentityOutputReference interface {
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
	InternalValue() *SynapseWorkspaceIdentity
	SetInternalValue(val *SynapseWorkspaceIdentity)
	PrincipalId() *string
	TenantId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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

// The jsii proxy struct for SynapseWorkspaceIdentityOutputReference
type jsiiProxy_SynapseWorkspaceIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) InternalValue() *SynapseWorkspaceIdentity {
	var returns *SynapseWorkspaceIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SynapseWorkspaceIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceIdentityOutputReference_Override(s SynapseWorkspaceIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) SetInternalValue(val *SynapseWorkspaceIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SynapseWorkspaceSqlAadAdmin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#login SynapseWorkspace#login}.
	Login *string `field:"optional" json:"login" yaml:"login"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#object_id SynapseWorkspace#object_id}.
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#tenant_id SynapseWorkspace#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

type SynapseWorkspaceSqlAadAdminList interface {
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
	Get(index *float64) SynapseWorkspaceSqlAadAdminOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SynapseWorkspaceSqlAadAdminList
type jsiiProxy_SynapseWorkspaceSqlAadAdminList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceSqlAadAdminList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SynapseWorkspaceSqlAadAdminList {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceSqlAadAdminList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceSqlAadAdminList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceSqlAadAdminList_Override(s SynapseWorkspaceSqlAadAdminList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceSqlAadAdminList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminList) Get(index *float64) SynapseWorkspaceSqlAadAdminOutputReference {
	var returns SynapseWorkspaceSqlAadAdminOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SynapseWorkspaceSqlAadAdminOutputReference interface {
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
	Login() *string
	SetLogin(val *string)
	LoginInput() *string
	ObjectId() *string
	SetObjectId(val *string)
	ObjectIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetLogin()
	ResetObjectId()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SynapseWorkspaceSqlAadAdminOutputReference
type jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) LoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceSqlAadAdminOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SynapseWorkspaceSqlAadAdminOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceSqlAadAdminOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceSqlAadAdminOutputReference_Override(s SynapseWorkspaceSqlAadAdminOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceSqlAadAdminOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) SetLogin(val *string) {
	_jsii_.Set(
		j,
		"login",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) SetObjectId(val *string) {
	_jsii_.Set(
		j,
		"objectId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ResetLogin() {
	_jsii_.InvokeVoid(
		s,
		"resetLogin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ResetObjectId() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		s,
		"resetTenantId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceSqlAadAdminOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SynapseWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#create SynapseWorkspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#delete SynapseWorkspace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#read SynapseWorkspace#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace#update SynapseWorkspace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type SynapseWorkspaceTimeoutsOutputReference interface {
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

// The jsii proxy struct for SynapseWorkspaceTimeoutsOutputReference
type jsiiProxy_SynapseWorkspaceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SynapseWorkspaceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SynapseWorkspaceTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceTimeoutsOutputReference_Override(s SynapseWorkspaceTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

