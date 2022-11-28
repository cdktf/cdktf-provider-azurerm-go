package storageaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/storageaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account azurerm_storage_account}.
type StorageAccount interface {
	cdktf.TerraformResource
	AccessTier() *string
	SetAccessTier(val *string)
	AccessTierInput() *string
	AccountKind() *string
	SetAccountKind(val *string)
	AccountKindInput() *string
	AccountReplicationType() *string
	SetAccountReplicationType(val *string)
	AccountReplicationTypeInput() *string
	AccountTier() *string
	SetAccountTier(val *string)
	AccountTierInput() *string
	AllowNestedItemsToBePublic() interface{}
	SetAllowNestedItemsToBePublic(val interface{})
	AllowNestedItemsToBePublicInput() interface{}
	AzureFilesAuthentication() StorageAccountAzureFilesAuthenticationOutputReference
	AzureFilesAuthenticationInput() *StorageAccountAzureFilesAuthentication
	BlobProperties() StorageAccountBlobPropertiesOutputReference
	BlobPropertiesInput() *StorageAccountBlobProperties
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
	CrossTenantReplicationEnabled() interface{}
	SetCrossTenantReplicationEnabled(val interface{})
	CrossTenantReplicationEnabledInput() interface{}
	CustomDomain() StorageAccountCustomDomainOutputReference
	CustomDomainInput() *StorageAccountCustomDomain
	CustomerManagedKey() StorageAccountCustomerManagedKeyOutputReference
	CustomerManagedKeyInput() *StorageAccountCustomerManagedKey
	DefaultToOauthAuthentication() interface{}
	SetDefaultToOauthAuthentication(val interface{})
	DefaultToOauthAuthenticationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	EnableHttpsTrafficOnly() interface{}
	SetEnableHttpsTrafficOnly(val interface{})
	EnableHttpsTrafficOnlyInput() interface{}
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
	Identity() StorageAccountIdentityOutputReference
	IdentityInput() *StorageAccountIdentity
	IdInput() *string
	ImmutabilityPolicy() StorageAccountImmutabilityPolicyOutputReference
	ImmutabilityPolicyInput() *StorageAccountImmutabilityPolicy
	InfrastructureEncryptionEnabled() interface{}
	SetInfrastructureEncryptionEnabled(val interface{})
	InfrastructureEncryptionEnabledInput() interface{}
	IsHnsEnabled() interface{}
	SetIsHnsEnabled(val interface{})
	IsHnsEnabledInput() interface{}
	LargeFileShareEnabled() interface{}
	SetLargeFileShareEnabled(val interface{})
	LargeFileShareEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MinTlsVersion() *string
	SetMinTlsVersion(val *string)
	MinTlsVersionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkRules() StorageAccountNetworkRulesOutputReference
	NetworkRulesInput() *StorageAccountNetworkRules
	Nfsv3Enabled() interface{}
	SetNfsv3Enabled(val interface{})
	Nfsv3EnabledInput() interface{}
	// The tree node.
	Node() constructs.Node
	PrimaryAccessKey() *string
	PrimaryBlobConnectionString() *string
	PrimaryBlobEndpoint() *string
	PrimaryBlobHost() *string
	PrimaryConnectionString() *string
	PrimaryDfsEndpoint() *string
	PrimaryDfsHost() *string
	PrimaryFileEndpoint() *string
	PrimaryFileHost() *string
	PrimaryLocation() *string
	PrimaryQueueEndpoint() *string
	PrimaryQueueHost() *string
	PrimaryTableEndpoint() *string
	PrimaryTableHost() *string
	PrimaryWebEndpoint() *string
	PrimaryWebHost() *string
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
	QueueEncryptionKeyType() *string
	SetQueueEncryptionKeyType(val *string)
	QueueEncryptionKeyTypeInput() *string
	QueueProperties() StorageAccountQueuePropertiesOutputReference
	QueuePropertiesInput() *StorageAccountQueueProperties
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Routing() StorageAccountRoutingOutputReference
	RoutingInput() *StorageAccountRouting
	SasPolicy() StorageAccountSasPolicyOutputReference
	SasPolicyInput() *StorageAccountSasPolicy
	SecondaryAccessKey() *string
	SecondaryBlobConnectionString() *string
	SecondaryBlobEndpoint() *string
	SecondaryBlobHost() *string
	SecondaryConnectionString() *string
	SecondaryDfsEndpoint() *string
	SecondaryDfsHost() *string
	SecondaryFileEndpoint() *string
	SecondaryFileHost() *string
	SecondaryLocation() *string
	SecondaryQueueEndpoint() *string
	SecondaryQueueHost() *string
	SecondaryTableEndpoint() *string
	SecondaryTableHost() *string
	SecondaryWebEndpoint() *string
	SecondaryWebHost() *string
	SharedAccessKeyEnabled() interface{}
	SetSharedAccessKeyEnabled(val interface{})
	SharedAccessKeyEnabledInput() interface{}
	ShareProperties() StorageAccountSharePropertiesOutputReference
	SharePropertiesInput() *StorageAccountShareProperties
	StaticWebsite() StorageAccountStaticWebsiteOutputReference
	StaticWebsiteInput() *StorageAccountStaticWebsite
	TableEncryptionKeyType() *string
	SetTableEncryptionKeyType(val *string)
	TableEncryptionKeyTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StorageAccountTimeoutsOutputReference
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
	PutAzureFilesAuthentication(value *StorageAccountAzureFilesAuthentication)
	PutBlobProperties(value *StorageAccountBlobProperties)
	PutCustomDomain(value *StorageAccountCustomDomain)
	PutCustomerManagedKey(value *StorageAccountCustomerManagedKey)
	PutIdentity(value *StorageAccountIdentity)
	PutImmutabilityPolicy(value *StorageAccountImmutabilityPolicy)
	PutNetworkRules(value *StorageAccountNetworkRules)
	PutQueueProperties(value *StorageAccountQueueProperties)
	PutRouting(value *StorageAccountRouting)
	PutSasPolicy(value *StorageAccountSasPolicy)
	PutShareProperties(value *StorageAccountShareProperties)
	PutStaticWebsite(value *StorageAccountStaticWebsite)
	PutTimeouts(value *StorageAccountTimeouts)
	ResetAccessTier()
	ResetAccountKind()
	ResetAllowNestedItemsToBePublic()
	ResetAzureFilesAuthentication()
	ResetBlobProperties()
	ResetCrossTenantReplicationEnabled()
	ResetCustomDomain()
	ResetCustomerManagedKey()
	ResetDefaultToOauthAuthentication()
	ResetEdgeZone()
	ResetEnableHttpsTrafficOnly()
	ResetId()
	ResetIdentity()
	ResetImmutabilityPolicy()
	ResetInfrastructureEncryptionEnabled()
	ResetIsHnsEnabled()
	ResetLargeFileShareEnabled()
	ResetMinTlsVersion()
	ResetNetworkRules()
	ResetNfsv3Enabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetQueueEncryptionKeyType()
	ResetQueueProperties()
	ResetRouting()
	ResetSasPolicy()
	ResetSharedAccessKeyEnabled()
	ResetShareProperties()
	ResetStaticWebsite()
	ResetTableEncryptionKeyType()
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

// The jsii proxy struct for StorageAccount
type jsiiProxy_StorageAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageAccount) AccessTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AccessTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AccountKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AccountKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AccountReplicationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountReplicationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AccountReplicationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountReplicationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AccountTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AccountTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AllowNestedItemsToBePublic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNestedItemsToBePublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AllowNestedItemsToBePublicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNestedItemsToBePublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AzureFilesAuthentication() StorageAccountAzureFilesAuthenticationOutputReference {
	var returns StorageAccountAzureFilesAuthenticationOutputReference
	_jsii_.Get(
		j,
		"azureFilesAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) AzureFilesAuthenticationInput() *StorageAccountAzureFilesAuthentication {
	var returns *StorageAccountAzureFilesAuthentication
	_jsii_.Get(
		j,
		"azureFilesAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) BlobProperties() StorageAccountBlobPropertiesOutputReference {
	var returns StorageAccountBlobPropertiesOutputReference
	_jsii_.Get(
		j,
		"blobProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) BlobPropertiesInput() *StorageAccountBlobProperties {
	var returns *StorageAccountBlobProperties
	_jsii_.Get(
		j,
		"blobPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) CrossTenantReplicationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossTenantReplicationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) CrossTenantReplicationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossTenantReplicationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) CustomDomain() StorageAccountCustomDomainOutputReference {
	var returns StorageAccountCustomDomainOutputReference
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) CustomDomainInput() *StorageAccountCustomDomain {
	var returns *StorageAccountCustomDomain
	_jsii_.Get(
		j,
		"customDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) CustomerManagedKey() StorageAccountCustomerManagedKeyOutputReference {
	var returns StorageAccountCustomerManagedKeyOutputReference
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) CustomerManagedKeyInput() *StorageAccountCustomerManagedKey {
	var returns *StorageAccountCustomerManagedKey
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) DefaultToOauthAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultToOauthAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) DefaultToOauthAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultToOauthAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) EnableHttpsTrafficOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpsTrafficOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) EnableHttpsTrafficOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpsTrafficOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Identity() StorageAccountIdentityOutputReference {
	var returns StorageAccountIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) IdentityInput() *StorageAccountIdentity {
	var returns *StorageAccountIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) ImmutabilityPolicy() StorageAccountImmutabilityPolicyOutputReference {
	var returns StorageAccountImmutabilityPolicyOutputReference
	_jsii_.Get(
		j,
		"immutabilityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) ImmutabilityPolicyInput() *StorageAccountImmutabilityPolicy {
	var returns *StorageAccountImmutabilityPolicy
	_jsii_.Get(
		j,
		"immutabilityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) InfrastructureEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infrastructureEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) InfrastructureEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infrastructureEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) IsHnsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHnsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) IsHnsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHnsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) LargeFileShareEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"largeFileShareEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) LargeFileShareEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"largeFileShareEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) MinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) NetworkRules() StorageAccountNetworkRulesOutputReference {
	var returns StorageAccountNetworkRulesOutputReference
	_jsii_.Get(
		j,
		"networkRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) NetworkRulesInput() *StorageAccountNetworkRules {
	var returns *StorageAccountNetworkRules
	_jsii_.Get(
		j,
		"networkRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Nfsv3Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv3Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Nfsv3EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv3EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryBlobConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryBlobEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryBlobHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryDfsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDfsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryDfsHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDfsHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryFileEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFileEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryFileHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFileHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryQueueEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryQueueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryQueueHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryQueueHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryTableEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTableEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryTableHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTableHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryWebEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryWebEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PrimaryWebHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryWebHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) QueueEncryptionKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueEncryptionKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) QueueEncryptionKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueEncryptionKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) QueueProperties() StorageAccountQueuePropertiesOutputReference {
	var returns StorageAccountQueuePropertiesOutputReference
	_jsii_.Get(
		j,
		"queueProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) QueuePropertiesInput() *StorageAccountQueueProperties {
	var returns *StorageAccountQueueProperties
	_jsii_.Get(
		j,
		"queuePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Routing() StorageAccountRoutingOutputReference {
	var returns StorageAccountRoutingOutputReference
	_jsii_.Get(
		j,
		"routing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) RoutingInput() *StorageAccountRouting {
	var returns *StorageAccountRouting
	_jsii_.Get(
		j,
		"routingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SasPolicy() StorageAccountSasPolicyOutputReference {
	var returns StorageAccountSasPolicyOutputReference
	_jsii_.Get(
		j,
		"sasPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SasPolicyInput() *StorageAccountSasPolicy {
	var returns *StorageAccountSasPolicy
	_jsii_.Get(
		j,
		"sasPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryBlobConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryBlobEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryBlobHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryDfsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryDfsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryDfsHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryDfsHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryFileEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFileEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryFileHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFileHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryQueueEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryQueueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryQueueHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryQueueHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryTableEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryTableEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryTableHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryTableHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryWebEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryWebEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SecondaryWebHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryWebHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SharedAccessKeyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedAccessKeyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SharedAccessKeyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedAccessKeyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) ShareProperties() StorageAccountSharePropertiesOutputReference {
	var returns StorageAccountSharePropertiesOutputReference
	_jsii_.Get(
		j,
		"shareProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) SharePropertiesInput() *StorageAccountShareProperties {
	var returns *StorageAccountShareProperties
	_jsii_.Get(
		j,
		"sharePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) StaticWebsite() StorageAccountStaticWebsiteOutputReference {
	var returns StorageAccountStaticWebsiteOutputReference
	_jsii_.Get(
		j,
		"staticWebsite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) StaticWebsiteInput() *StorageAccountStaticWebsite {
	var returns *StorageAccountStaticWebsite
	_jsii_.Get(
		j,
		"staticWebsiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) TableEncryptionKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableEncryptionKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) TableEncryptionKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableEncryptionKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) Timeouts() StorageAccountTimeoutsOutputReference {
	var returns StorageAccountTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccount) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account azurerm_storage_account} Resource.
func NewStorageAccount(scope constructs.Construct, id *string, config *StorageAccountConfig) StorageAccount {
	_init_.Initialize()

	if err := validateNewStorageAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageAccount{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account azurerm_storage_account} Resource.
func NewStorageAccount_Override(s StorageAccount, scope constructs.Construct, id *string, config *StorageAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccount",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageAccount)SetAccessTier(val *string) {
	if err := j.validateSetAccessTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTier",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetAccountKind(val *string) {
	if err := j.validateSetAccountKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountKind",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetAccountReplicationType(val *string) {
	if err := j.validateSetAccountReplicationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountReplicationType",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetAccountTier(val *string) {
	if err := j.validateSetAccountTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountTier",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetAllowNestedItemsToBePublic(val interface{}) {
	if err := j.validateSetAllowNestedItemsToBePublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowNestedItemsToBePublic",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetCrossTenantReplicationEnabled(val interface{}) {
	if err := j.validateSetCrossTenantReplicationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossTenantReplicationEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetDefaultToOauthAuthentication(val interface{}) {
	if err := j.validateSetDefaultToOauthAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultToOauthAuthentication",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetEdgeZone(val *string) {
	if err := j.validateSetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetEnableHttpsTrafficOnly(val interface{}) {
	if err := j.validateSetEnableHttpsTrafficOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttpsTrafficOnly",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetInfrastructureEncryptionEnabled(val interface{}) {
	if err := j.validateSetInfrastructureEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetIsHnsEnabled(val interface{}) {
	if err := j.validateSetIsHnsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isHnsEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetLargeFileShareEnabled(val interface{}) {
	if err := j.validateSetLargeFileShareEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"largeFileShareEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetMinTlsVersion(val *string) {
	if err := j.validateSetMinTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetNfsv3Enabled(val interface{}) {
	if err := j.validateSetNfsv3EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv3Enabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetQueueEncryptionKeyType(val *string) {
	if err := j.validateSetQueueEncryptionKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueEncryptionKeyType",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetSharedAccessKeyEnabled(val interface{}) {
	if err := j.validateSetSharedAccessKeyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccessKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetTableEncryptionKeyType(val *string) {
	if err := j.validateSetTableEncryptionKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableEncryptionKeyType",
		val,
	)
}

func (j *jsiiProxy_StorageAccount)SetTags(val *map[string]*string) {
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
func StorageAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageAccount.StorageAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageAccount.StorageAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageAccount.StorageAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.storageAccount.StorageAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageAccount) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageAccount) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageAccount) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageAccount) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageAccount) PutAzureFilesAuthentication(value *StorageAccountAzureFilesAuthentication) {
	if err := s.validatePutAzureFilesAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureFilesAuthentication",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutBlobProperties(value *StorageAccountBlobProperties) {
	if err := s.validatePutBlobPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBlobProperties",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutCustomDomain(value *StorageAccountCustomDomain) {
	if err := s.validatePutCustomDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomDomain",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutCustomerManagedKey(value *StorageAccountCustomerManagedKey) {
	if err := s.validatePutCustomerManagedKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomerManagedKey",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutIdentity(value *StorageAccountIdentity) {
	if err := s.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdentity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutImmutabilityPolicy(value *StorageAccountImmutabilityPolicy) {
	if err := s.validatePutImmutabilityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putImmutabilityPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutNetworkRules(value *StorageAccountNetworkRules) {
	if err := s.validatePutNetworkRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkRules",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutQueueProperties(value *StorageAccountQueueProperties) {
	if err := s.validatePutQueuePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putQueueProperties",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutRouting(value *StorageAccountRouting) {
	if err := s.validatePutRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRouting",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutSasPolicy(value *StorageAccountSasPolicy) {
	if err := s.validatePutSasPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSasPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutShareProperties(value *StorageAccountShareProperties) {
	if err := s.validatePutSharePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putShareProperties",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutStaticWebsite(value *StorageAccountStaticWebsite) {
	if err := s.validatePutStaticWebsiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStaticWebsite",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutTimeouts(value *StorageAccountTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) ResetAccessTier() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessTier",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetAccountKind() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountKind",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetAllowNestedItemsToBePublic() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowNestedItemsToBePublic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetAzureFilesAuthentication() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureFilesAuthentication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetBlobProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetBlobProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetCrossTenantReplicationEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCrossTenantReplicationEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetCustomDomain() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomDomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetDefaultToOauthAuthentication() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultToOauthAuthentication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		s,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetEnableHttpsTrafficOnly() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableHttpsTrafficOnly",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetIdentity() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetImmutabilityPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetImmutabilityPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetInfrastructureEncryptionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetInfrastructureEncryptionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetIsHnsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetIsHnsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetLargeFileShareEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetLargeFileShareEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetMinTlsVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetMinTlsVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetNetworkRules() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkRules",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetNfsv3Enabled() {
	_jsii_.InvokeVoid(
		s,
		"resetNfsv3Enabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetQueueEncryptionKeyType() {
	_jsii_.InvokeVoid(
		s,
		"resetQueueEncryptionKeyType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetQueueProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetQueueProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetRouting() {
	_jsii_.InvokeVoid(
		s,
		"resetRouting",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetSasPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetSasPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetSharedAccessKeyEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedAccessKeyEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetShareProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetShareProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetStaticWebsite() {
	_jsii_.InvokeVoid(
		s,
		"resetStaticWebsite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetTableEncryptionKeyType() {
	_jsii_.InvokeVoid(
		s,
		"resetTableEncryptionKeyType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

