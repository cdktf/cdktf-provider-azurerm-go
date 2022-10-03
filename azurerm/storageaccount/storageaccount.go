package storageaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/storageaccount/internal"
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
	PutNetworkRules(value *StorageAccountNetworkRules)
	PutQueueProperties(value *StorageAccountQueueProperties)
	PutRouting(value *StorageAccountRouting)
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

func (j *jsiiProxy_StorageAccount) SetAccessTier(val *string) {
	_jsii_.Set(
		j,
		"accessTier",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetAccountKind(val *string) {
	_jsii_.Set(
		j,
		"accountKind",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetAccountReplicationType(val *string) {
	_jsii_.Set(
		j,
		"accountReplicationType",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetAccountTier(val *string) {
	_jsii_.Set(
		j,
		"accountTier",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetAllowNestedItemsToBePublic(val interface{}) {
	_jsii_.Set(
		j,
		"allowNestedItemsToBePublic",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetCrossTenantReplicationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"crossTenantReplicationEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetDefaultToOauthAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"defaultToOauthAuthentication",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetEdgeZone(val *string) {
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetEnableHttpsTrafficOnly(val interface{}) {
	_jsii_.Set(
		j,
		"enableHttpsTrafficOnly",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetInfrastructureEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"infrastructureEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetIsHnsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"isHnsEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetLargeFileShareEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"largeFileShareEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetMinTlsVersion(val *string) {
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetNfsv3Enabled(val interface{}) {
	_jsii_.Set(
		j,
		"nfsv3Enabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetQueueEncryptionKeyType(val *string) {
	_jsii_.Set(
		j,
		"queueEncryptionKeyType",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetSharedAccessKeyEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"sharedAccessKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetTableEncryptionKeyType(val *string) {
	_jsii_.Set(
		j,
		"tableEncryptionKeyType",
		val,
	)
}

func (j *jsiiProxy_StorageAccount) SetTags(val *map[string]*string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageAccount.StorageAccount",
		"isConstruct",
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
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageAccount) PutAzureFilesAuthentication(value *StorageAccountAzureFilesAuthentication) {
	_jsii_.InvokeVoid(
		s,
		"putAzureFilesAuthentication",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutBlobProperties(value *StorageAccountBlobProperties) {
	_jsii_.InvokeVoid(
		s,
		"putBlobProperties",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutCustomDomain(value *StorageAccountCustomDomain) {
	_jsii_.InvokeVoid(
		s,
		"putCustomDomain",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutCustomerManagedKey(value *StorageAccountCustomerManagedKey) {
	_jsii_.InvokeVoid(
		s,
		"putCustomerManagedKey",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutIdentity(value *StorageAccountIdentity) {
	_jsii_.InvokeVoid(
		s,
		"putIdentity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutNetworkRules(value *StorageAccountNetworkRules) {
	_jsii_.InvokeVoid(
		s,
		"putNetworkRules",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutQueueProperties(value *StorageAccountQueueProperties) {
	_jsii_.InvokeVoid(
		s,
		"putQueueProperties",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutRouting(value *StorageAccountRouting) {
	_jsii_.InvokeVoid(
		s,
		"putRouting",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutShareProperties(value *StorageAccountShareProperties) {
	_jsii_.InvokeVoid(
		s,
		"putShareProperties",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutStaticWebsite(value *StorageAccountStaticWebsite) {
	_jsii_.InvokeVoid(
		s,
		"putStaticWebsite",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccount) PutTimeouts(value *StorageAccountTimeouts) {
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

type StorageAccountAzureFilesAuthentication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#directory_type StorageAccount#directory_type}.
	DirectoryType *string `field:"required" json:"directoryType" yaml:"directoryType"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#active_directory StorageAccount#active_directory}
	ActiveDirectory *StorageAccountAzureFilesAuthenticationActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
}

type StorageAccountAzureFilesAuthenticationActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#domain_guid StorageAccount#domain_guid}.
	DomainGuid *string `field:"required" json:"domainGuid" yaml:"domainGuid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#domain_name StorageAccount#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#domain_sid StorageAccount#domain_sid}.
	DomainSid *string `field:"required" json:"domainSid" yaml:"domainSid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#forest_name StorageAccount#forest_name}.
	ForestName *string `field:"required" json:"forestName" yaml:"forestName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#netbios_domain_name StorageAccount#netbios_domain_name}.
	NetbiosDomainName *string `field:"required" json:"netbiosDomainName" yaml:"netbiosDomainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#storage_sid StorageAccount#storage_sid}.
	StorageSid *string `field:"required" json:"storageSid" yaml:"storageSid"`
}

type StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference interface {
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
	DomainGuid() *string
	SetDomainGuid(val *string)
	DomainGuidInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainSid() *string
	SetDomainSid(val *string)
	DomainSidInput() *string
	ForestName() *string
	SetForestName(val *string)
	ForestNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountAzureFilesAuthenticationActiveDirectory
	SetInternalValue(val *StorageAccountAzureFilesAuthenticationActiveDirectory)
	NetbiosDomainName() *string
	SetNetbiosDomainName(val *string)
	NetbiosDomainNameInput() *string
	StorageSid() *string
	SetStorageSid(val *string)
	StorageSidInput() *string
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

// The jsii proxy struct for StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference
type jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainGuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainGuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainSid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) DomainSidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ForestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ForestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) InternalValue() *StorageAccountAzureFilesAuthenticationActiveDirectory {
	var returns *StorageAccountAzureFilesAuthenticationActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) NetbiosDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) NetbiosDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) StorageSid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) StorageSidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference_Override(s StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetDomainGuid(val *string) {
	_jsii_.Set(
		j,
		"domainGuid",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetDomainSid(val *string) {
	_jsii_.Set(
		j,
		"domainSid",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetForestName(val *string) {
	_jsii_.Set(
		j,
		"forestName",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetInternalValue(val *StorageAccountAzureFilesAuthenticationActiveDirectory) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetNetbiosDomainName(val *string) {
	_jsii_.Set(
		j,
		"netbiosDomainName",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetStorageSid(val *string) {
	_jsii_.Set(
		j,
		"storageSid",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountAzureFilesAuthenticationOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference
	ActiveDirectoryInput() *StorageAccountAzureFilesAuthenticationActiveDirectory
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
	DirectoryType() *string
	SetDirectoryType(val *string)
	DirectoryTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountAzureFilesAuthentication
	SetInternalValue(val *StorageAccountAzureFilesAuthentication)
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
	PutActiveDirectory(value *StorageAccountAzureFilesAuthenticationActiveDirectory)
	ResetActiveDirectory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountAzureFilesAuthenticationOutputReference
type jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ActiveDirectory() StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference {
	var returns StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ActiveDirectoryInput() *StorageAccountAzureFilesAuthenticationActiveDirectory {
	var returns *StorageAccountAzureFilesAuthenticationActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) DirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) DirectoryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) InternalValue() *StorageAccountAzureFilesAuthentication {
	var returns *StorageAccountAzureFilesAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountAzureFilesAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountAzureFilesAuthenticationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountAzureFilesAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountAzureFilesAuthenticationOutputReference_Override(s StorageAccountAzureFilesAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountAzureFilesAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) SetDirectoryType(val *string) {
	_jsii_.Set(
		j,
		"directoryType",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) SetInternalValue(val *StorageAccountAzureFilesAuthentication) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) PutActiveDirectory(value *StorageAccountAzureFilesAuthenticationActiveDirectory) {
	_jsii_.InvokeVoid(
		s,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		s,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountBlobProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#change_feed_enabled StorageAccount#change_feed_enabled}.
	ChangeFeedEnabled interface{} `field:"optional" json:"changeFeedEnabled" yaml:"changeFeedEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#change_feed_retention_in_days StorageAccount#change_feed_retention_in_days}.
	ChangeFeedRetentionInDays *float64 `field:"optional" json:"changeFeedRetentionInDays" yaml:"changeFeedRetentionInDays"`
	// container_delete_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#container_delete_retention_policy StorageAccount#container_delete_retention_policy}
	ContainerDeleteRetentionPolicy *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy `field:"optional" json:"containerDeleteRetentionPolicy" yaml:"containerDeleteRetentionPolicy"`
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#cors_rule StorageAccount#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#default_service_version StorageAccount#default_service_version}.
	DefaultServiceVersion *string `field:"optional" json:"defaultServiceVersion" yaml:"defaultServiceVersion"`
	// delete_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#delete_retention_policy StorageAccount#delete_retention_policy}
	DeleteRetentionPolicy *StorageAccountBlobPropertiesDeleteRetentionPolicy `field:"optional" json:"deleteRetentionPolicy" yaml:"deleteRetentionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#last_access_time_enabled StorageAccount#last_access_time_enabled}.
	LastAccessTimeEnabled interface{} `field:"optional" json:"lastAccessTimeEnabled" yaml:"lastAccessTimeEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#versioning_enabled StorageAccount#versioning_enabled}.
	VersioningEnabled interface{} `field:"optional" json:"versioningEnabled" yaml:"versioningEnabled"`
}

type StorageAccountBlobPropertiesContainerDeleteRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#days StorageAccount#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

type StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference interface {
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
	Days() *float64
	SetDays(val *float64)
	DaysInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy
	SetInternalValue(val *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy)
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
	ResetDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference
type jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) Days() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"days",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) DaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) InternalValue() *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy {
	var returns *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference_Override(s StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) SetDays(val *float64) {
	_jsii_.Set(
		j,
		"days",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) SetInternalValue(val *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) ResetDays() {
	_jsii_.InvokeVoid(
		s,
		"resetDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountBlobPropertiesCorsRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_headers StorageAccount#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_methods StorageAccount#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_origins StorageAccount#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#exposed_headers StorageAccount#exposed_headers}.
	ExposedHeaders *[]*string `field:"required" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#max_age_in_seconds StorageAccount#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"required" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

type StorageAccountBlobPropertiesCorsRuleList interface {
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
	Get(index *float64) StorageAccountBlobPropertiesCorsRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountBlobPropertiesCorsRuleList
type jsiiProxy_StorageAccountBlobPropertiesCorsRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewStorageAccountBlobPropertiesCorsRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StorageAccountBlobPropertiesCorsRuleList {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountBlobPropertiesCorsRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesCorsRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewStorageAccountBlobPropertiesCorsRuleList_Override(s StorageAccountBlobPropertiesCorsRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesCorsRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) Get(index *float64) StorageAccountBlobPropertiesCorsRuleOutputReference {
	var returns StorageAccountBlobPropertiesCorsRuleOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountBlobPropertiesCorsRuleOutputReference interface {
	cdktf.ComplexObject
	AllowedHeaders() *[]*string
	SetAllowedHeaders(val *[]*string)
	AllowedHeadersInput() *[]*string
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	ExposedHeaders() *[]*string
	SetExposedHeaders(val *[]*string)
	ExposedHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxAgeInSeconds() *float64
	SetMaxAgeInSeconds(val *float64)
	MaxAgeInSecondsInput() *float64
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

// The jsii proxy struct for StorageAccountBlobPropertiesCorsRuleOutputReference
type jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) AllowedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) AllowedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) ExposedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) ExposedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) MaxAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) MaxAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountBlobPropertiesCorsRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageAccountBlobPropertiesCorsRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesCorsRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageAccountBlobPropertiesCorsRuleOutputReference_Override(s StorageAccountBlobPropertiesCorsRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesCorsRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetAllowedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedHeaders",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetAllowedMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetAllowedOrigins(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetExposedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"exposedHeaders",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetMaxAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesCorsRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountBlobPropertiesDeleteRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#days StorageAccount#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

type StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference interface {
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
	Days() *float64
	SetDays(val *float64)
	DaysInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountBlobPropertiesDeleteRetentionPolicy
	SetInternalValue(val *StorageAccountBlobPropertiesDeleteRetentionPolicy)
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
	ResetDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference
type jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) Days() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"days",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) DaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) InternalValue() *StorageAccountBlobPropertiesDeleteRetentionPolicy {
	var returns *StorageAccountBlobPropertiesDeleteRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference_Override(s StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) SetDays(val *float64) {
	_jsii_.Set(
		j,
		"days",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) SetInternalValue(val *StorageAccountBlobPropertiesDeleteRetentionPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) ResetDays() {
	_jsii_.InvokeVoid(
		s,
		"resetDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountBlobPropertiesOutputReference interface {
	cdktf.ComplexObject
	ChangeFeedEnabled() interface{}
	SetChangeFeedEnabled(val interface{})
	ChangeFeedEnabledInput() interface{}
	ChangeFeedRetentionInDays() *float64
	SetChangeFeedRetentionInDays(val *float64)
	ChangeFeedRetentionInDaysInput() *float64
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
	ContainerDeleteRetentionPolicy() StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference
	ContainerDeleteRetentionPolicyInput() *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy
	CorsRule() StorageAccountBlobPropertiesCorsRuleList
	CorsRuleInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultServiceVersion() *string
	SetDefaultServiceVersion(val *string)
	DefaultServiceVersionInput() *string
	DeleteRetentionPolicy() StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference
	DeleteRetentionPolicyInput() *StorageAccountBlobPropertiesDeleteRetentionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountBlobProperties
	SetInternalValue(val *StorageAccountBlobProperties)
	LastAccessTimeEnabled() interface{}
	SetLastAccessTimeEnabled(val interface{})
	LastAccessTimeEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VersioningEnabled() interface{}
	SetVersioningEnabled(val interface{})
	VersioningEnabledInput() interface{}
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
	PutContainerDeleteRetentionPolicy(value *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy)
	PutCorsRule(value interface{})
	PutDeleteRetentionPolicy(value *StorageAccountBlobPropertiesDeleteRetentionPolicy)
	ResetChangeFeedEnabled()
	ResetChangeFeedRetentionInDays()
	ResetContainerDeleteRetentionPolicy()
	ResetCorsRule()
	ResetDefaultServiceVersion()
	ResetDeleteRetentionPolicy()
	ResetLastAccessTimeEnabled()
	ResetVersioningEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountBlobPropertiesOutputReference
type jsiiProxy_StorageAccountBlobPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ChangeFeedEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeFeedEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ChangeFeedEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeFeedEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ChangeFeedRetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeFeedRetentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ChangeFeedRetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeFeedRetentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ContainerDeleteRetentionPolicy() StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference {
	var returns StorageAccountBlobPropertiesContainerDeleteRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"containerDeleteRetentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ContainerDeleteRetentionPolicyInput() *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy {
	var returns *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy
	_jsii_.Get(
		j,
		"containerDeleteRetentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) CorsRule() StorageAccountBlobPropertiesCorsRuleList {
	var returns StorageAccountBlobPropertiesCorsRuleList
	_jsii_.Get(
		j,
		"corsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) CorsRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) DefaultServiceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultServiceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) DefaultServiceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultServiceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) DeleteRetentionPolicy() StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference {
	var returns StorageAccountBlobPropertiesDeleteRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"deleteRetentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) DeleteRetentionPolicyInput() *StorageAccountBlobPropertiesDeleteRetentionPolicy {
	var returns *StorageAccountBlobPropertiesDeleteRetentionPolicy
	_jsii_.Get(
		j,
		"deleteRetentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) InternalValue() *StorageAccountBlobProperties {
	var returns *StorageAccountBlobProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) LastAccessTimeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastAccessTimeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) LastAccessTimeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastAccessTimeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) VersioningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) VersioningEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioningEnabledInput",
		&returns,
	)
	return returns
}


func NewStorageAccountBlobPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountBlobPropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountBlobPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountBlobPropertiesOutputReference_Override(s StorageAccountBlobPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountBlobPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetChangeFeedEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"changeFeedEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetChangeFeedRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"changeFeedRetentionInDays",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetDefaultServiceVersion(val *string) {
	_jsii_.Set(
		j,
		"defaultServiceVersion",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetInternalValue(val *StorageAccountBlobProperties) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetLastAccessTimeEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"lastAccessTimeEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountBlobPropertiesOutputReference) SetVersioningEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"versioningEnabled",
		val,
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) PutContainerDeleteRetentionPolicy(value *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy) {
	_jsii_.InvokeVoid(
		s,
		"putContainerDeleteRetentionPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) PutCorsRule(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCorsRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) PutDeleteRetentionPolicy(value *StorageAccountBlobPropertiesDeleteRetentionPolicy) {
	_jsii_.InvokeVoid(
		s,
		"putDeleteRetentionPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetChangeFeedEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeFeedEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetChangeFeedRetentionInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetChangeFeedRetentionInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetContainerDeleteRetentionPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerDeleteRetentionPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetCorsRule() {
	_jsii_.InvokeVoid(
		s,
		"resetCorsRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetDefaultServiceVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultServiceVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetDeleteRetentionPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteRetentionPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetLastAccessTimeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetLastAccessTimeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ResetVersioningEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetVersioningEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountBlobPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#account_replication_type StorageAccount#account_replication_type}.
	AccountReplicationType *string `field:"required" json:"accountReplicationType" yaml:"accountReplicationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#account_tier StorageAccount#account_tier}.
	AccountTier *string `field:"required" json:"accountTier" yaml:"accountTier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#location StorageAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#name StorageAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#resource_group_name StorageAccount#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#access_tier StorageAccount#access_tier}.
	AccessTier *string `field:"optional" json:"accessTier" yaml:"accessTier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#account_kind StorageAccount#account_kind}.
	AccountKind *string `field:"optional" json:"accountKind" yaml:"accountKind"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allow_nested_items_to_be_public StorageAccount#allow_nested_items_to_be_public}.
	AllowNestedItemsToBePublic interface{} `field:"optional" json:"allowNestedItemsToBePublic" yaml:"allowNestedItemsToBePublic"`
	// azure_files_authentication block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#azure_files_authentication StorageAccount#azure_files_authentication}
	AzureFilesAuthentication *StorageAccountAzureFilesAuthentication `field:"optional" json:"azureFilesAuthentication" yaml:"azureFilesAuthentication"`
	// blob_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#blob_properties StorageAccount#blob_properties}
	BlobProperties *StorageAccountBlobProperties `field:"optional" json:"blobProperties" yaml:"blobProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#cross_tenant_replication_enabled StorageAccount#cross_tenant_replication_enabled}.
	CrossTenantReplicationEnabled interface{} `field:"optional" json:"crossTenantReplicationEnabled" yaml:"crossTenantReplicationEnabled"`
	// custom_domain block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#custom_domain StorageAccount#custom_domain}
	CustomDomain *StorageAccountCustomDomain `field:"optional" json:"customDomain" yaml:"customDomain"`
	// customer_managed_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#customer_managed_key StorageAccount#customer_managed_key}
	CustomerManagedKey *StorageAccountCustomerManagedKey `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#default_to_oauth_authentication StorageAccount#default_to_oauth_authentication}.
	DefaultToOauthAuthentication interface{} `field:"optional" json:"defaultToOauthAuthentication" yaml:"defaultToOauthAuthentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#edge_zone StorageAccount#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#enable_https_traffic_only StorageAccount#enable_https_traffic_only}.
	EnableHttpsTrafficOnly interface{} `field:"optional" json:"enableHttpsTrafficOnly" yaml:"enableHttpsTrafficOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#id StorageAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#identity StorageAccount#identity}
	Identity *StorageAccountIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#infrastructure_encryption_enabled StorageAccount#infrastructure_encryption_enabled}.
	InfrastructureEncryptionEnabled interface{} `field:"optional" json:"infrastructureEncryptionEnabled" yaml:"infrastructureEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#is_hns_enabled StorageAccount#is_hns_enabled}.
	IsHnsEnabled interface{} `field:"optional" json:"isHnsEnabled" yaml:"isHnsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#large_file_share_enabled StorageAccount#large_file_share_enabled}.
	LargeFileShareEnabled interface{} `field:"optional" json:"largeFileShareEnabled" yaml:"largeFileShareEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#min_tls_version StorageAccount#min_tls_version}.
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// network_rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#network_rules StorageAccount#network_rules}
	NetworkRules *StorageAccountNetworkRules `field:"optional" json:"networkRules" yaml:"networkRules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#nfsv3_enabled StorageAccount#nfsv3_enabled}.
	Nfsv3Enabled interface{} `field:"optional" json:"nfsv3Enabled" yaml:"nfsv3Enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#public_network_access_enabled StorageAccount#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#queue_encryption_key_type StorageAccount#queue_encryption_key_type}.
	QueueEncryptionKeyType *string `field:"optional" json:"queueEncryptionKeyType" yaml:"queueEncryptionKeyType"`
	// queue_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#queue_properties StorageAccount#queue_properties}
	QueueProperties *StorageAccountQueueProperties `field:"optional" json:"queueProperties" yaml:"queueProperties"`
	// routing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#routing StorageAccount#routing}
	Routing *StorageAccountRouting `field:"optional" json:"routing" yaml:"routing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#shared_access_key_enabled StorageAccount#shared_access_key_enabled}.
	SharedAccessKeyEnabled interface{} `field:"optional" json:"sharedAccessKeyEnabled" yaml:"sharedAccessKeyEnabled"`
	// share_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#share_properties StorageAccount#share_properties}
	ShareProperties *StorageAccountShareProperties `field:"optional" json:"shareProperties" yaml:"shareProperties"`
	// static_website block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#static_website StorageAccount#static_website}
	StaticWebsite *StorageAccountStaticWebsite `field:"optional" json:"staticWebsite" yaml:"staticWebsite"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#table_encryption_key_type StorageAccount#table_encryption_key_type}.
	TableEncryptionKeyType *string `field:"optional" json:"tableEncryptionKeyType" yaml:"tableEncryptionKeyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#tags StorageAccount#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#timeouts StorageAccount#timeouts}
	Timeouts *StorageAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type StorageAccountCustomDomain struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#name StorageAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#use_subdomain StorageAccount#use_subdomain}.
	UseSubdomain interface{} `field:"optional" json:"useSubdomain" yaml:"useSubdomain"`
}

type StorageAccountCustomDomainOutputReference interface {
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
	InternalValue() *StorageAccountCustomDomain
	SetInternalValue(val *StorageAccountCustomDomain)
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
	UseSubdomain() interface{}
	SetUseSubdomain(val interface{})
	UseSubdomainInput() interface{}
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
	ResetUseSubdomain()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountCustomDomainOutputReference
type jsiiProxy_StorageAccountCustomDomainOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) InternalValue() *StorageAccountCustomDomain {
	var returns *StorageAccountCustomDomain
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) UseSubdomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSubdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) UseSubdomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSubdomainInput",
		&returns,
	)
	return returns
}


func NewStorageAccountCustomDomainOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountCustomDomainOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountCustomDomainOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountCustomDomainOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountCustomDomainOutputReference_Override(s StorageAccountCustomDomainOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountCustomDomainOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) SetInternalValue(val *StorageAccountCustomDomain) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomDomainOutputReference) SetUseSubdomain(val interface{}) {
	_jsii_.Set(
		j,
		"useSubdomain",
		val,
	)
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) ResetUseSubdomain() {
	_jsii_.InvokeVoid(
		s,
		"resetUseSubdomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomDomainOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountCustomerManagedKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#key_vault_key_id StorageAccount#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#user_assigned_identity_id StorageAccount#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"required" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

type StorageAccountCustomerManagedKeyOutputReference interface {
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
	InternalValue() *StorageAccountCustomerManagedKey
	SetInternalValue(val *StorageAccountCustomerManagedKey)
	KeyVaultKeyId() *string
	SetKeyVaultKeyId(val *string)
	KeyVaultKeyIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAssignedIdentityId() *string
	SetUserAssignedIdentityId(val *string)
	UserAssignedIdentityIdInput() *string
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

// The jsii proxy struct for StorageAccountCustomerManagedKeyOutputReference
type jsiiProxy_StorageAccountCustomerManagedKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) InternalValue() *StorageAccountCustomerManagedKey {
	var returns *StorageAccountCustomerManagedKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) KeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) KeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) UserAssignedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityIdInput",
		&returns,
	)
	return returns
}


func NewStorageAccountCustomerManagedKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountCustomerManagedKeyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountCustomerManagedKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountCustomerManagedKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountCustomerManagedKeyOutputReference_Override(s StorageAccountCustomerManagedKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountCustomerManagedKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) SetInternalValue(val *StorageAccountCustomerManagedKey) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) SetKeyVaultKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) SetUserAssignedIdentityId(val *string) {
	_jsii_.Set(
		j,
		"userAssignedIdentityId",
		val,
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#type StorageAccount#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#identity_ids StorageAccount#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type StorageAccountIdentityOutputReference interface {
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
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *StorageAccountIdentity
	SetInternalValue(val *StorageAccountIdentity)
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
	ResetIdentityIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountIdentityOutputReference
type jsiiProxy_StorageAccountIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) InternalValue() *StorageAccountIdentity {
	var returns *StorageAccountIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewStorageAccountIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountIdentityOutputReference_Override(s StorageAccountIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) SetInternalValue(val *StorageAccountIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountNetworkRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#default_action StorageAccount#default_action}.
	DefaultAction *string `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#bypass StorageAccount#bypass}.
	Bypass *[]*string `field:"optional" json:"bypass" yaml:"bypass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#ip_rules StorageAccount#ip_rules}.
	IpRules *[]*string `field:"optional" json:"ipRules" yaml:"ipRules"`
	// private_link_access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#private_link_access StorageAccount#private_link_access}
	PrivateLinkAccess interface{} `field:"optional" json:"privateLinkAccess" yaml:"privateLinkAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#virtual_network_subnet_ids StorageAccount#virtual_network_subnet_ids}.
	VirtualNetworkSubnetIds *[]*string `field:"optional" json:"virtualNetworkSubnetIds" yaml:"virtualNetworkSubnetIds"`
}

type StorageAccountNetworkRulesOutputReference interface {
	cdktf.ComplexObject
	Bypass() *[]*string
	SetBypass(val *[]*string)
	BypassInput() *[]*string
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
	DefaultAction() *string
	SetDefaultAction(val *string)
	DefaultActionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountNetworkRules
	SetInternalValue(val *StorageAccountNetworkRules)
	IpRules() *[]*string
	SetIpRules(val *[]*string)
	IpRulesInput() *[]*string
	PrivateLinkAccess() StorageAccountNetworkRulesPrivateLinkAccessList
	PrivateLinkAccessInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkSubnetIds() *[]*string
	SetVirtualNetworkSubnetIds(val *[]*string)
	VirtualNetworkSubnetIdsInput() *[]*string
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
	PutPrivateLinkAccess(value interface{})
	ResetBypass()
	ResetIpRules()
	ResetPrivateLinkAccess()
	ResetVirtualNetworkSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountNetworkRulesOutputReference
type jsiiProxy_StorageAccountNetworkRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) Bypass() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) BypassInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bypassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) DefaultAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) DefaultActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) InternalValue() *StorageAccountNetworkRules {
	var returns *StorageAccountNetworkRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) IpRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) IpRulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) PrivateLinkAccess() StorageAccountNetworkRulesPrivateLinkAccessList {
	var returns StorageAccountNetworkRulesPrivateLinkAccessList
	_jsii_.Get(
		j,
		"privateLinkAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) PrivateLinkAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) VirtualNetworkSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) VirtualNetworkSubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdsInput",
		&returns,
	)
	return returns
}


func NewStorageAccountNetworkRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountNetworkRulesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountNetworkRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountNetworkRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountNetworkRulesOutputReference_Override(s StorageAccountNetworkRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountNetworkRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetBypass(val *[]*string) {
	_jsii_.Set(
		j,
		"bypass",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetDefaultAction(val *string) {
	_jsii_.Set(
		j,
		"defaultAction",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetInternalValue(val *StorageAccountNetworkRules) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetIpRules(val *[]*string) {
	_jsii_.Set(
		j,
		"ipRules",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesOutputReference) SetVirtualNetworkSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"virtualNetworkSubnetIds",
		val,
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) PutPrivateLinkAccess(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putPrivateLinkAccess",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) ResetBypass() {
	_jsii_.InvokeVoid(
		s,
		"resetBypass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) ResetIpRules() {
	_jsii_.InvokeVoid(
		s,
		"resetIpRules",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) ResetPrivateLinkAccess() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateLinkAccess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) ResetVirtualNetworkSubnetIds() {
	_jsii_.InvokeVoid(
		s,
		"resetVirtualNetworkSubnetIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountNetworkRulesPrivateLinkAccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#endpoint_resource_id StorageAccount#endpoint_resource_id}.
	EndpointResourceId *string `field:"required" json:"endpointResourceId" yaml:"endpointResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#endpoint_tenant_id StorageAccount#endpoint_tenant_id}.
	EndpointTenantId *string `field:"optional" json:"endpointTenantId" yaml:"endpointTenantId"`
}

type StorageAccountNetworkRulesPrivateLinkAccessList interface {
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
	Get(index *float64) StorageAccountNetworkRulesPrivateLinkAccessOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountNetworkRulesPrivateLinkAccessList
type jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewStorageAccountNetworkRulesPrivateLinkAccessList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StorageAccountNetworkRulesPrivateLinkAccessList {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountNetworkRulesPrivateLinkAccessList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewStorageAccountNetworkRulesPrivateLinkAccessList_Override(s StorageAccountNetworkRulesPrivateLinkAccessList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountNetworkRulesPrivateLinkAccessList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) Get(index *float64) StorageAccountNetworkRulesPrivateLinkAccessOutputReference {
	var returns StorageAccountNetworkRulesPrivateLinkAccessOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountNetworkRulesPrivateLinkAccessOutputReference interface {
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
	EndpointResourceId() *string
	SetEndpointResourceId(val *string)
	EndpointResourceIdInput() *string
	EndpointTenantId() *string
	SetEndpointTenantId(val *string)
	EndpointTenantIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetEndpointTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountNetworkRulesPrivateLinkAccessOutputReference
type jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) EndpointResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) EndpointResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) EndpointTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) EndpointTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountNetworkRulesPrivateLinkAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageAccountNetworkRulesPrivateLinkAccessOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountNetworkRulesPrivateLinkAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageAccountNetworkRulesPrivateLinkAccessOutputReference_Override(s StorageAccountNetworkRulesPrivateLinkAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountNetworkRulesPrivateLinkAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) SetEndpointResourceId(val *string) {
	_jsii_.Set(
		j,
		"endpointResourceId",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) SetEndpointTenantId(val *string) {
	_jsii_.Set(
		j,
		"endpointTenantId",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) ResetEndpointTenantId() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointTenantId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountQueueProperties struct {
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#cors_rule StorageAccount#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// hour_metrics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#hour_metrics StorageAccount#hour_metrics}
	HourMetrics *StorageAccountQueuePropertiesHourMetrics `field:"optional" json:"hourMetrics" yaml:"hourMetrics"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#logging StorageAccount#logging}
	Logging *StorageAccountQueuePropertiesLogging `field:"optional" json:"logging" yaml:"logging"`
	// minute_metrics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#minute_metrics StorageAccount#minute_metrics}
	MinuteMetrics *StorageAccountQueuePropertiesMinuteMetrics `field:"optional" json:"minuteMetrics" yaml:"minuteMetrics"`
}

type StorageAccountQueuePropertiesCorsRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_headers StorageAccount#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_methods StorageAccount#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_origins StorageAccount#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#exposed_headers StorageAccount#exposed_headers}.
	ExposedHeaders *[]*string `field:"required" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#max_age_in_seconds StorageAccount#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"required" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

type StorageAccountQueuePropertiesCorsRuleList interface {
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
	Get(index *float64) StorageAccountQueuePropertiesCorsRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountQueuePropertiesCorsRuleList
type jsiiProxy_StorageAccountQueuePropertiesCorsRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewStorageAccountQueuePropertiesCorsRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StorageAccountQueuePropertiesCorsRuleList {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountQueuePropertiesCorsRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesCorsRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewStorageAccountQueuePropertiesCorsRuleList_Override(s StorageAccountQueuePropertiesCorsRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesCorsRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) Get(index *float64) StorageAccountQueuePropertiesCorsRuleOutputReference {
	var returns StorageAccountQueuePropertiesCorsRuleOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountQueuePropertiesCorsRuleOutputReference interface {
	cdktf.ComplexObject
	AllowedHeaders() *[]*string
	SetAllowedHeaders(val *[]*string)
	AllowedHeadersInput() *[]*string
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	ExposedHeaders() *[]*string
	SetExposedHeaders(val *[]*string)
	ExposedHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxAgeInSeconds() *float64
	SetMaxAgeInSeconds(val *float64)
	MaxAgeInSecondsInput() *float64
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

// The jsii proxy struct for StorageAccountQueuePropertiesCorsRuleOutputReference
type jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) AllowedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) AllowedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) ExposedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) ExposedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) MaxAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) MaxAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountQueuePropertiesCorsRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageAccountQueuePropertiesCorsRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesCorsRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageAccountQueuePropertiesCorsRuleOutputReference_Override(s StorageAccountQueuePropertiesCorsRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesCorsRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetAllowedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedHeaders",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetAllowedMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetAllowedOrigins(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetExposedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"exposedHeaders",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetMaxAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesCorsRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountQueuePropertiesHourMetrics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#enabled StorageAccount#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#version StorageAccount#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#include_apis StorageAccount#include_apis}.
	IncludeApis interface{} `field:"optional" json:"includeApis" yaml:"includeApis"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#retention_policy_days StorageAccount#retention_policy_days}.
	RetentionPolicyDays *float64 `field:"optional" json:"retentionPolicyDays" yaml:"retentionPolicyDays"`
}

type StorageAccountQueuePropertiesHourMetricsOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeApis() interface{}
	SetIncludeApis(val interface{})
	IncludeApisInput() interface{}
	InternalValue() *StorageAccountQueuePropertiesHourMetrics
	SetInternalValue(val *StorageAccountQueuePropertiesHourMetrics)
	RetentionPolicyDays() *float64
	SetRetentionPolicyDays(val *float64)
	RetentionPolicyDaysInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetIncludeApis()
	ResetRetentionPolicyDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountQueuePropertiesHourMetricsOutputReference
type jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) IncludeApis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeApis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) IncludeApisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeApisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) InternalValue() *StorageAccountQueuePropertiesHourMetrics {
	var returns *StorageAccountQueuePropertiesHourMetrics
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) RetentionPolicyDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPolicyDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) RetentionPolicyDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPolicyDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewStorageAccountQueuePropertiesHourMetricsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountQueuePropertiesHourMetricsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesHourMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountQueuePropertiesHourMetricsOutputReference_Override(s StorageAccountQueuePropertiesHourMetricsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesHourMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetIncludeApis(val interface{}) {
	_jsii_.Set(
		j,
		"includeApis",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetInternalValue(val *StorageAccountQueuePropertiesHourMetrics) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetRetentionPolicyDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionPolicyDays",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) ResetIncludeApis() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeApis",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) ResetRetentionPolicyDays() {
	_jsii_.InvokeVoid(
		s,
		"resetRetentionPolicyDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesHourMetricsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountQueuePropertiesLogging struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#delete StorageAccount#delete}.
	Delete interface{} `field:"required" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#read StorageAccount#read}.
	Read interface{} `field:"required" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#version StorageAccount#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#write StorageAccount#write}.
	Write interface{} `field:"required" json:"write" yaml:"write"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#retention_policy_days StorageAccount#retention_policy_days}.
	RetentionPolicyDays *float64 `field:"optional" json:"retentionPolicyDays" yaml:"retentionPolicyDays"`
}

type StorageAccountQueuePropertiesLoggingOutputReference interface {
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
	Delete() interface{}
	SetDelete(val interface{})
	DeleteInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountQueuePropertiesLogging
	SetInternalValue(val *StorageAccountQueuePropertiesLogging)
	Read() interface{}
	SetRead(val interface{})
	ReadInput() interface{}
	RetentionPolicyDays() *float64
	SetRetentionPolicyDays(val *float64)
	RetentionPolicyDaysInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Write() interface{}
	SetWrite(val interface{})
	WriteInput() interface{}
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
	ResetRetentionPolicyDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountQueuePropertiesLoggingOutputReference
type jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) Delete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) DeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) InternalValue() *StorageAccountQueuePropertiesLogging {
	var returns *StorageAccountQueuePropertiesLogging
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) Read() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) ReadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) RetentionPolicyDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPolicyDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) RetentionPolicyDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPolicyDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) Write() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"write",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) WriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeInput",
		&returns,
	)
	return returns
}


func NewStorageAccountQueuePropertiesLoggingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountQueuePropertiesLoggingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesLoggingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountQueuePropertiesLoggingOutputReference_Override(s StorageAccountQueuePropertiesLoggingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesLoggingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetDelete(val interface{}) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetInternalValue(val *StorageAccountQueuePropertiesLogging) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetRead(val interface{}) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetRetentionPolicyDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionPolicyDays",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) SetWrite(val interface{}) {
	_jsii_.Set(
		j,
		"write",
		val,
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) ResetRetentionPolicyDays() {
	_jsii_.InvokeVoid(
		s,
		"resetRetentionPolicyDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesLoggingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountQueuePropertiesMinuteMetrics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#enabled StorageAccount#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#version StorageAccount#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#include_apis StorageAccount#include_apis}.
	IncludeApis interface{} `field:"optional" json:"includeApis" yaml:"includeApis"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#retention_policy_days StorageAccount#retention_policy_days}.
	RetentionPolicyDays *float64 `field:"optional" json:"retentionPolicyDays" yaml:"retentionPolicyDays"`
}

type StorageAccountQueuePropertiesMinuteMetricsOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeApis() interface{}
	SetIncludeApis(val interface{})
	IncludeApisInput() interface{}
	InternalValue() *StorageAccountQueuePropertiesMinuteMetrics
	SetInternalValue(val *StorageAccountQueuePropertiesMinuteMetrics)
	RetentionPolicyDays() *float64
	SetRetentionPolicyDays(val *float64)
	RetentionPolicyDaysInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetIncludeApis()
	ResetRetentionPolicyDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountQueuePropertiesMinuteMetricsOutputReference
type jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) IncludeApis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeApis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) IncludeApisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeApisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) InternalValue() *StorageAccountQueuePropertiesMinuteMetrics {
	var returns *StorageAccountQueuePropertiesMinuteMetrics
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) RetentionPolicyDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPolicyDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) RetentionPolicyDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPolicyDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewStorageAccountQueuePropertiesMinuteMetricsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountQueuePropertiesMinuteMetricsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesMinuteMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountQueuePropertiesMinuteMetricsOutputReference_Override(s StorageAccountQueuePropertiesMinuteMetricsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesMinuteMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetIncludeApis(val interface{}) {
	_jsii_.Set(
		j,
		"includeApis",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetInternalValue(val *StorageAccountQueuePropertiesMinuteMetrics) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetRetentionPolicyDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionPolicyDays",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) ResetIncludeApis() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeApis",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) ResetRetentionPolicyDays() {
	_jsii_.InvokeVoid(
		s,
		"resetRetentionPolicyDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesMinuteMetricsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountQueuePropertiesOutputReference interface {
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
	CorsRule() StorageAccountQueuePropertiesCorsRuleList
	CorsRuleInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HourMetrics() StorageAccountQueuePropertiesHourMetricsOutputReference
	HourMetricsInput() *StorageAccountQueuePropertiesHourMetrics
	InternalValue() *StorageAccountQueueProperties
	SetInternalValue(val *StorageAccountQueueProperties)
	Logging() StorageAccountQueuePropertiesLoggingOutputReference
	LoggingInput() *StorageAccountQueuePropertiesLogging
	MinuteMetrics() StorageAccountQueuePropertiesMinuteMetricsOutputReference
	MinuteMetricsInput() *StorageAccountQueuePropertiesMinuteMetrics
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
	PutCorsRule(value interface{})
	PutHourMetrics(value *StorageAccountQueuePropertiesHourMetrics)
	PutLogging(value *StorageAccountQueuePropertiesLogging)
	PutMinuteMetrics(value *StorageAccountQueuePropertiesMinuteMetrics)
	ResetCorsRule()
	ResetHourMetrics()
	ResetLogging()
	ResetMinuteMetrics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountQueuePropertiesOutputReference
type jsiiProxy_StorageAccountQueuePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) CorsRule() StorageAccountQueuePropertiesCorsRuleList {
	var returns StorageAccountQueuePropertiesCorsRuleList
	_jsii_.Get(
		j,
		"corsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) CorsRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) HourMetrics() StorageAccountQueuePropertiesHourMetricsOutputReference {
	var returns StorageAccountQueuePropertiesHourMetricsOutputReference
	_jsii_.Get(
		j,
		"hourMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) HourMetricsInput() *StorageAccountQueuePropertiesHourMetrics {
	var returns *StorageAccountQueuePropertiesHourMetrics
	_jsii_.Get(
		j,
		"hourMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) InternalValue() *StorageAccountQueueProperties {
	var returns *StorageAccountQueueProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) Logging() StorageAccountQueuePropertiesLoggingOutputReference {
	var returns StorageAccountQueuePropertiesLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) LoggingInput() *StorageAccountQueuePropertiesLogging {
	var returns *StorageAccountQueuePropertiesLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) MinuteMetrics() StorageAccountQueuePropertiesMinuteMetricsOutputReference {
	var returns StorageAccountQueuePropertiesMinuteMetricsOutputReference
	_jsii_.Get(
		j,
		"minuteMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) MinuteMetricsInput() *StorageAccountQueuePropertiesMinuteMetrics {
	var returns *StorageAccountQueuePropertiesMinuteMetrics
	_jsii_.Get(
		j,
		"minuteMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountQueuePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountQueuePropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountQueuePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountQueuePropertiesOutputReference_Override(s StorageAccountQueuePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountQueuePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) SetInternalValue(val *StorageAccountQueueProperties) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountQueuePropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) PutCorsRule(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCorsRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) PutHourMetrics(value *StorageAccountQueuePropertiesHourMetrics) {
	_jsii_.InvokeVoid(
		s,
		"putHourMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) PutLogging(value *StorageAccountQueuePropertiesLogging) {
	_jsii_.InvokeVoid(
		s,
		"putLogging",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) PutMinuteMetrics(value *StorageAccountQueuePropertiesMinuteMetrics) {
	_jsii_.InvokeVoid(
		s,
		"putMinuteMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) ResetCorsRule() {
	_jsii_.InvokeVoid(
		s,
		"resetCorsRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) ResetHourMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetHourMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) ResetLogging() {
	_jsii_.InvokeVoid(
		s,
		"resetLogging",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) ResetMinuteMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetMinuteMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountQueuePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountRouting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#choice StorageAccount#choice}.
	Choice *string `field:"optional" json:"choice" yaml:"choice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#publish_internet_endpoints StorageAccount#publish_internet_endpoints}.
	PublishInternetEndpoints interface{} `field:"optional" json:"publishInternetEndpoints" yaml:"publishInternetEndpoints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#publish_microsoft_endpoints StorageAccount#publish_microsoft_endpoints}.
	PublishMicrosoftEndpoints interface{} `field:"optional" json:"publishMicrosoftEndpoints" yaml:"publishMicrosoftEndpoints"`
}

type StorageAccountRoutingOutputReference interface {
	cdktf.ComplexObject
	Choice() *string
	SetChoice(val *string)
	ChoiceInput() *string
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
	InternalValue() *StorageAccountRouting
	SetInternalValue(val *StorageAccountRouting)
	PublishInternetEndpoints() interface{}
	SetPublishInternetEndpoints(val interface{})
	PublishInternetEndpointsInput() interface{}
	PublishMicrosoftEndpoints() interface{}
	SetPublishMicrosoftEndpoints(val interface{})
	PublishMicrosoftEndpointsInput() interface{}
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
	ResetChoice()
	ResetPublishInternetEndpoints()
	ResetPublishMicrosoftEndpoints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountRoutingOutputReference
type jsiiProxy_StorageAccountRoutingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) Choice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"choice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) ChoiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"choiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) InternalValue() *StorageAccountRouting {
	var returns *StorageAccountRouting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) PublishInternetEndpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishInternetEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) PublishInternetEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishInternetEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) PublishMicrosoftEndpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishMicrosoftEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) PublishMicrosoftEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishMicrosoftEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountRoutingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountRoutingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountRoutingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountRoutingOutputReference_Override(s StorageAccountRoutingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) SetChoice(val *string) {
	_jsii_.Set(
		j,
		"choice",
		val,
	)
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) SetInternalValue(val *StorageAccountRouting) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) SetPublishInternetEndpoints(val interface{}) {
	_jsii_.Set(
		j,
		"publishInternetEndpoints",
		val,
	)
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) SetPublishMicrosoftEndpoints(val interface{}) {
	_jsii_.Set(
		j,
		"publishMicrosoftEndpoints",
		val,
	)
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountRoutingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) ResetChoice() {
	_jsii_.InvokeVoid(
		s,
		"resetChoice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) ResetPublishInternetEndpoints() {
	_jsii_.InvokeVoid(
		s,
		"resetPublishInternetEndpoints",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) ResetPublishMicrosoftEndpoints() {
	_jsii_.InvokeVoid(
		s,
		"resetPublishMicrosoftEndpoints",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountRoutingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountShareProperties struct {
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#cors_rule StorageAccount#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#retention_policy StorageAccount#retention_policy}
	RetentionPolicy *StorageAccountSharePropertiesRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// smb block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#smb StorageAccount#smb}
	Smb *StorageAccountSharePropertiesSmb `field:"optional" json:"smb" yaml:"smb"`
}

type StorageAccountSharePropertiesCorsRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_headers StorageAccount#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_methods StorageAccount#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#allowed_origins StorageAccount#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#exposed_headers StorageAccount#exposed_headers}.
	ExposedHeaders *[]*string `field:"required" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#max_age_in_seconds StorageAccount#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"required" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

type StorageAccountSharePropertiesCorsRuleList interface {
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
	Get(index *float64) StorageAccountSharePropertiesCorsRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountSharePropertiesCorsRuleList
type jsiiProxy_StorageAccountSharePropertiesCorsRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewStorageAccountSharePropertiesCorsRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StorageAccountSharePropertiesCorsRuleList {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountSharePropertiesCorsRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesCorsRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewStorageAccountSharePropertiesCorsRuleList_Override(s StorageAccountSharePropertiesCorsRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesCorsRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) Get(index *float64) StorageAccountSharePropertiesCorsRuleOutputReference {
	var returns StorageAccountSharePropertiesCorsRuleOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountSharePropertiesCorsRuleOutputReference interface {
	cdktf.ComplexObject
	AllowedHeaders() *[]*string
	SetAllowedHeaders(val *[]*string)
	AllowedHeadersInput() *[]*string
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	ExposedHeaders() *[]*string
	SetExposedHeaders(val *[]*string)
	ExposedHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxAgeInSeconds() *float64
	SetMaxAgeInSeconds(val *float64)
	MaxAgeInSecondsInput() *float64
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

// The jsii proxy struct for StorageAccountSharePropertiesCorsRuleOutputReference
type jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) AllowedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) AllowedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) ExposedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) ExposedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) MaxAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) MaxAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountSharePropertiesCorsRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageAccountSharePropertiesCorsRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesCorsRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageAccountSharePropertiesCorsRuleOutputReference_Override(s StorageAccountSharePropertiesCorsRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesCorsRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetAllowedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedHeaders",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetAllowedMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetAllowedOrigins(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetExposedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"exposedHeaders",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetMaxAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesCorsRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountSharePropertiesOutputReference interface {
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
	CorsRule() StorageAccountSharePropertiesCorsRuleList
	CorsRuleInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountShareProperties
	SetInternalValue(val *StorageAccountShareProperties)
	RetentionPolicy() StorageAccountSharePropertiesRetentionPolicyOutputReference
	RetentionPolicyInput() *StorageAccountSharePropertiesRetentionPolicy
	Smb() StorageAccountSharePropertiesSmbOutputReference
	SmbInput() *StorageAccountSharePropertiesSmb
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
	PutCorsRule(value interface{})
	PutRetentionPolicy(value *StorageAccountSharePropertiesRetentionPolicy)
	PutSmb(value *StorageAccountSharePropertiesSmb)
	ResetCorsRule()
	ResetRetentionPolicy()
	ResetSmb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountSharePropertiesOutputReference
type jsiiProxy_StorageAccountSharePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) CorsRule() StorageAccountSharePropertiesCorsRuleList {
	var returns StorageAccountSharePropertiesCorsRuleList
	_jsii_.Get(
		j,
		"corsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) CorsRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) InternalValue() *StorageAccountShareProperties {
	var returns *StorageAccountShareProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) RetentionPolicy() StorageAccountSharePropertiesRetentionPolicyOutputReference {
	var returns StorageAccountSharePropertiesRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) RetentionPolicyInput() *StorageAccountSharePropertiesRetentionPolicy {
	var returns *StorageAccountSharePropertiesRetentionPolicy
	_jsii_.Get(
		j,
		"retentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) Smb() StorageAccountSharePropertiesSmbOutputReference {
	var returns StorageAccountSharePropertiesSmbOutputReference
	_jsii_.Get(
		j,
		"smb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) SmbInput() *StorageAccountSharePropertiesSmb {
	var returns *StorageAccountSharePropertiesSmb
	_jsii_.Get(
		j,
		"smbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountSharePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountSharePropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountSharePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountSharePropertiesOutputReference_Override(s StorageAccountSharePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) SetInternalValue(val *StorageAccountShareProperties) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) PutCorsRule(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCorsRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) PutRetentionPolicy(value *StorageAccountSharePropertiesRetentionPolicy) {
	_jsii_.InvokeVoid(
		s,
		"putRetentionPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) PutSmb(value *StorageAccountSharePropertiesSmb) {
	_jsii_.InvokeVoid(
		s,
		"putSmb",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) ResetCorsRule() {
	_jsii_.InvokeVoid(
		s,
		"resetCorsRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) ResetRetentionPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetRetentionPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) ResetSmb() {
	_jsii_.InvokeVoid(
		s,
		"resetSmb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountSharePropertiesRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#days StorageAccount#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

type StorageAccountSharePropertiesRetentionPolicyOutputReference interface {
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
	Days() *float64
	SetDays(val *float64)
	DaysInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountSharePropertiesRetentionPolicy
	SetInternalValue(val *StorageAccountSharePropertiesRetentionPolicy)
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
	ResetDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountSharePropertiesRetentionPolicyOutputReference
type jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) Days() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"days",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) DaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) InternalValue() *StorageAccountSharePropertiesRetentionPolicy {
	var returns *StorageAccountSharePropertiesRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountSharePropertiesRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountSharePropertiesRetentionPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountSharePropertiesRetentionPolicyOutputReference_Override(s StorageAccountSharePropertiesRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) SetDays(val *float64) {
	_jsii_.Set(
		j,
		"days",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) SetInternalValue(val *StorageAccountSharePropertiesRetentionPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) ResetDays() {
	_jsii_.InvokeVoid(
		s,
		"resetDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountSharePropertiesSmb struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#authentication_types StorageAccount#authentication_types}.
	AuthenticationTypes *[]*string `field:"optional" json:"authenticationTypes" yaml:"authenticationTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#channel_encryption_type StorageAccount#channel_encryption_type}.
	ChannelEncryptionType *[]*string `field:"optional" json:"channelEncryptionType" yaml:"channelEncryptionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#kerberos_ticket_encryption_type StorageAccount#kerberos_ticket_encryption_type}.
	KerberosTicketEncryptionType *[]*string `field:"optional" json:"kerberosTicketEncryptionType" yaml:"kerberosTicketEncryptionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#multichannel_enabled StorageAccount#multichannel_enabled}.
	MultichannelEnabled interface{} `field:"optional" json:"multichannelEnabled" yaml:"multichannelEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#versions StorageAccount#versions}.
	Versions *[]*string `field:"optional" json:"versions" yaml:"versions"`
}

type StorageAccountSharePropertiesSmbOutputReference interface {
	cdktf.ComplexObject
	AuthenticationTypes() *[]*string
	SetAuthenticationTypes(val *[]*string)
	AuthenticationTypesInput() *[]*string
	ChannelEncryptionType() *[]*string
	SetChannelEncryptionType(val *[]*string)
	ChannelEncryptionTypeInput() *[]*string
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
	InternalValue() *StorageAccountSharePropertiesSmb
	SetInternalValue(val *StorageAccountSharePropertiesSmb)
	KerberosTicketEncryptionType() *[]*string
	SetKerberosTicketEncryptionType(val *[]*string)
	KerberosTicketEncryptionTypeInput() *[]*string
	MultichannelEnabled() interface{}
	SetMultichannelEnabled(val interface{})
	MultichannelEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Versions() *[]*string
	SetVersions(val *[]*string)
	VersionsInput() *[]*string
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
	ResetAuthenticationTypes()
	ResetChannelEncryptionType()
	ResetKerberosTicketEncryptionType()
	ResetMultichannelEnabled()
	ResetVersions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountSharePropertiesSmbOutputReference
type jsiiProxy_StorageAccountSharePropertiesSmbOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) AuthenticationTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) AuthenticationTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ChannelEncryptionType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"channelEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ChannelEncryptionTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"channelEncryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) InternalValue() *StorageAccountSharePropertiesSmb {
	var returns *StorageAccountSharePropertiesSmb
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) KerberosTicketEncryptionType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kerberosTicketEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) KerberosTicketEncryptionTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kerberosTicketEncryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) MultichannelEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multichannelEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) MultichannelEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multichannelEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) Versions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) VersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionsInput",
		&returns,
	)
	return returns
}


func NewStorageAccountSharePropertiesSmbOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountSharePropertiesSmbOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountSharePropertiesSmbOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesSmbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountSharePropertiesSmbOutputReference_Override(s StorageAccountSharePropertiesSmbOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountSharePropertiesSmbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetAuthenticationTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"authenticationTypes",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetChannelEncryptionType(val *[]*string) {
	_jsii_.Set(
		j,
		"channelEncryptionType",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetInternalValue(val *StorageAccountSharePropertiesSmb) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetKerberosTicketEncryptionType(val *[]*string) {
	_jsii_.Set(
		j,
		"kerberosTicketEncryptionType",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetMultichannelEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"multichannelEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) SetVersions(val *[]*string) {
	_jsii_.Set(
		j,
		"versions",
		val,
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetAuthenticationTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthenticationTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetChannelEncryptionType() {
	_jsii_.InvokeVoid(
		s,
		"resetChannelEncryptionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetKerberosTicketEncryptionType() {
	_jsii_.InvokeVoid(
		s,
		"resetKerberosTicketEncryptionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetMultichannelEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetMultichannelEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetVersions() {
	_jsii_.InvokeVoid(
		s,
		"resetVersions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountStaticWebsite struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#error_404_document StorageAccount#error_404_document}.
	Error404Document *string `field:"optional" json:"error404Document" yaml:"error404Document"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#index_document StorageAccount#index_document}.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
}

type StorageAccountStaticWebsiteOutputReference interface {
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
	Error404Document() *string
	SetError404Document(val *string)
	Error404DocumentInput() *string
	// Experimental.
	Fqn() *string
	IndexDocument() *string
	SetIndexDocument(val *string)
	IndexDocumentInput() *string
	InternalValue() *StorageAccountStaticWebsite
	SetInternalValue(val *StorageAccountStaticWebsite)
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
	ResetError404Document()
	ResetIndexDocument()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountStaticWebsiteOutputReference
type jsiiProxy_StorageAccountStaticWebsiteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) Error404Document() *string {
	var returns *string
	_jsii_.Get(
		j,
		"error404Document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) Error404DocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"error404DocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) IndexDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) IndexDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) InternalValue() *StorageAccountStaticWebsite {
	var returns *StorageAccountStaticWebsite
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountStaticWebsiteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountStaticWebsiteOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountStaticWebsiteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountStaticWebsiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountStaticWebsiteOutputReference_Override(s StorageAccountStaticWebsiteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountStaticWebsiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) SetError404Document(val *string) {
	_jsii_.Set(
		j,
		"error404Document",
		val,
	)
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) SetIndexDocument(val *string) {
	_jsii_.Set(
		j,
		"indexDocument",
		val,
	)
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) SetInternalValue(val *StorageAccountStaticWebsite) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountStaticWebsiteOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) ResetError404Document() {
	_jsii_.InvokeVoid(
		s,
		"resetError404Document",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) ResetIndexDocument() {
	_jsii_.InvokeVoid(
		s,
		"resetIndexDocument",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountStaticWebsiteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StorageAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#create StorageAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#delete StorageAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#read StorageAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#update StorageAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type StorageAccountTimeoutsOutputReference interface {
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

// The jsii proxy struct for StorageAccountTimeoutsOutputReference
type jsiiProxy_StorageAccountTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewStorageAccountTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StorageAccountTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountTimeoutsOutputReference_Override(s StorageAccountTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

